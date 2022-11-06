from Components import *
from State import *
from DiaPlanner import *
from PointInterlock import *

# ダイヤ情報をもとにポイントの自動切換えを行う
class PointSwitcher:
    def __init__(self, state: State, diaPlanner: DiaPlanner, pointInterLock: PointInterlock):
        self.__state = state
        self.__diaPlanner = diaPlanner
        self.__pointInterlock = pointInterLock
        self.__autoToggle = True

    # ポイント自動切り替えを有効/無効化
    def setAutoToggleEnabled(self, enabled: bool) -> None:
        self.__autoToggle = enabled

    # 状態に応じてポイントを自動で切り替える. 毎update時に呼ぶこと
    def update(self) -> None:
        for junction in self.__state.junctionList:
            # if junction.id == 6:
            #     print(f"[PointSwitcher.update(j{junction.id})] t0: ({self.__state.getTrainById(0).currentSection.id}, {self.__state.getTrainById(0).mileage:.2f}), t1: ({self.__state.getTrainById(1).currentSection.id}, {self.__state.getTrainById(1).mileage:.2f}), j2: {self.__state.getJunctionById(2).getOutSection().id}, j6: {self.__state.getJunctionById(6).getInSection().id}")

            # in1->out2 という分岐器の場合、out側を到着列車が入線する番線に合わせる
            if junction.inSectionCurve == None and junction.outSectionCurve != None:
                train = self.__getNearestTrain(junction)  # junctionに一番先に到着する列車を取得
                if train:
                    # print(f"{junction.id}, {train.id}")
                    dia = self.__diaPlanner.getDia(train.id, junction.belongStation.id)  # このjunctionが存在する駅のダイヤ情報を取得
                    if dia.arriveSectionId != junction.getOutSection().id:
                        self.__pointInterlock.requestToggle(junction.id)
                        # print(f"[PointSwitcher.update] junction {junction.id} toggle requested to section {dia.arriveSectionId}")
            # in2->out1 という分岐器の場合、in側を出発列車の存在する番線に合わせる
            elif junction.inSectionCurve != None and junction.outSectionCurve == None:
                train = self.__getNearestTrain(junction)  # junctionを一番先に通る列車を取得
                if train:
                    # print(f"{junction.id}, {train.id}")
                    dia = self.__diaPlanner.getDia(train.id, junction.belongStation.id)  # このjunctionが存在する駅のダイヤ情報を取得
                    if dia.arriveSectionId != junction.getInSection().id:
                        self.__pointInterlock.requestToggle(junction.id)
                        # print(f"[PointSwitcher.update] junction {junction.id} toggle requested to section {dia.arriveSectionId}")

    def __getNearestTrain(self, junction: Junction, maxSearchNum: int = -1, originalJunction: Junction = None) -> Train:
        """
        指定したjunctionに一番先に到着する列車を取得する。指定したjunctionの手前に向かってセクションを辿っていき、列車を見つけたらそれを返す
        
        Parameters
        ----------
        junction : Junction
            指定したjunction。このjuntionに一番先に到着する列車を返す。存在しない場合は[]を返す
        maxSearchNum: int
            列車を何個手前のセクションまで確認するかどうか。-1を指定すると無限
        originalJunction : Junction
            再帰の停止条件。外から呼び出すときには何も代入しなくてよい
        """

        # まず、指定したjunctionから線路を辿り、junctionに先に到着する可能性のある列車をすべて取得する
        # 駅の出口など、inSectionが複数あるjunctionを指定した場合、trainsは2つ以上の候補がある
        trains: list[Train] = []

        # 何も見つけられずに最初の地点に戻ってきてしまった場合、終了
        if originalJunction != None and junction.id == originalJunction.id:
            return trains

        if originalJunction == None:
            originalJunction = junction
        
        # maxSearchNumが0になった場合、終了
        if maxSearchNum == 0:
            return trains
        
        while True:
            train = self.__state.getTrainInSection(junction.inSectionStraight)
            if train:
                # if junction.id == 6:
                #     print(f"train {train.id} exists on section {junction.inSectionStraight.id}")
                trains.append(train)
            else:
                # if junction.id == 6:
                #     print(f"train doesn't exist on section {junction.inSectionStraight.id}")
                nextJunction = junction.inSectionStraight.sourceJunction
                trains.append(self.__getNearestTrain(nextJunction, maxSearchNum - 1, originalJunction))

            # inSectionがひとつだけの分岐であれば、ここで終了
            if junction.inSectionCurve == None:
                # if junction.id == 6:
                #     print("no curve")
                break
            # inSectionが2つある場合、Curve側も調べる
            else:
                train = self.__state.getTrainInSection(junction.inSectionCurve)
                if train:
                    # if junction.id == 6:
                    #     print(f"train {train.id} exists on section {junction.inSectionCurve.id}")
                    trains.append(train)
                else:
                    # if junction.id == 6:
                    #     print(f"train doesn't exist on section {junction.inSectionCurve.id}")
                    nextJunction = junction.inSectionCurve.sourceJunction
                    trains.append(self.__getNearestTrain(nextJunction, maxSearchNum - 1, originalJunction))
                break

        trains = list(set(trains))  # 重複を削除

        # 候補となる列車が1つの場合はそれを返す
        if len(trains) == 0:
            # if junction.id == 6:
            #     print("no train")
            return None
        elif len(trains) == 1:
            # if junction.id == 6:
            #     print("one train")
            return trains[0]
        # 複数の候補がある場合、ダイヤと照らしあわせることで最も先にポイントを通過する列車を絞り込む
        else:
            # if junction.id == 6:
            #     print("multiple trains")
            station = self.__getNearestStation(junction)  # junction直前の駅を取得
            trainsWantToGo = list(filter(lambda t: self.__diaPlanner.getDia(t.id, station.id).wait == False, trains))  # 駅で退避するつもりのないtrainをfilter
            if len(trainsWantToGo) == 0:    # 全列車が退避したい場合、どれを先に出すか決めようがないので、とりあえず0番を返す
                # if junction.id == 6:
                #     print("multiple trains A")
                return trains[0]
            elif len(trainsWantToGo) == 1:  # 退避するつもりのない(追い抜きたい)列車が1つのとき、それを先に行かせる
                # if junction.id == 6:
                #     print("multiple trains B")
                return trainsWantToGo[0]
            else:                           # 退避するつもりのない列車が2つ以上のとき、最もjunctionに近いものを返す
                # if junction.id == 6:
                #     print("multiple trains C")
                trainsWantToGo.sort(key=lambda t: self.__state.getDistance(t.currentSection, t.mileage, junction.getOutSection(), 0))
                return trainsWantToGo[0]

    # 指定したjunctionの直前にある駅を取得
    def __getNearestStation(self, junction: Junction) -> Station:
        searchSection = junction.inSectionStraight
        while searchSection.station == None:
            searchSection = searchSection.sourceJunction.inSectionStraight
        return searchSection.station
