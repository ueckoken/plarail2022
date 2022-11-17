from Components import StopPoint, Train
from SignalSystem import SignalSystem
from State import State


class ATS:
    def __init__(
        self,
        state: State,
        signalSystem: SignalSystem,
        MAXSPEED: float,
        MERGIN: float = 20,
        BREAKING_DISTANCE: float = 10,
    ):
        """
        Parameters
        ----------
        MAXSPEED: 車両の最高速度[cm/s]
        MERGIN: ATSで非常停止するとき、信号機から何cm手前に停止するか
        BREAKING_DISTANCE: 最高速度で走行中にinputを0にしたときの制動距離[cm]
        """
        self.__state = state
        self.__signalSystem = signalSystem
        self.__MAXSPEED = MAXSPEED
        self.__MERGIN = MERGIN
        self.__BREAKING_DISTANCE = BREAKING_DISTANCE
        self.__enabled: dict[int, bool] = {}  # 各列車の、ATS有効,無効が入る辞書. key=trainId, value=enabled
        self.__activated: dict[int, bool] = {}  # ATSが作動したときにTrueが入る. key=trainId, value=activated
        for train in state.trainList:
            self.__enabled[train.id] = True
            self.__activated[train.id] = False

    # 指定した列車に対して速度を指令する. このとき、衝突しないような速度に変えて送信する
    def setSpeedCommand(self, trainId: int, speedCommand: float):
        train = self.__state.getTrainById(trainId)
        if train is None:
            print(f"[ATS] getTrainById({trainId}) return None")
            return

        # ATS有効時
        if self.__enabled[trainId]:
            atsStopPoint = self.getATSStopPoint(train)
            distance = self.__state.getDistance(
                train.currentSection, train.mileage, atsStopPoint.section, atsStopPoint.mileage
            )  # ATS停止点までの距離を計算
            if train.stopPoint is None:  # 停止点が代入されていない場合、ATSで計算した停止点を代入する
                train.stopPoint = atsStopPoint

            # ATS停止点までの距離から、出してよいスピード(speedLimit)を求める
            if distance > self.__BREAKING_DISTANCE:
                speedLimit = self.__MAXSPEED + 1
            elif distance > 0:
                speedLimit = (distance) / self.__BREAKING_DISTANCE * (self.__MAXSPEED + 1)
            else:
                speedLimit = 0

            if speedLimit < speedCommand:  # 非常停止できる速度を超えた速度が指示された場合、速度を強制的にspeedLimitに落とす
                train.targetSpeed = speedLimit
                if not self.__activated[train.id]:  # ATS作動フラグ(activated)を立てる
                    print(f"[ATS.setSpeedCommand] ATS activated on train {train.id} !")
                    self.__activated[train.id] = True
            else:
                train.targetSpeed = speedCommand
                if self.__activated[train.id]:
                    if (
                        atsStopPoint.section.id != train.currentSection.id
                    ):  # 目の前が赤信号ではなくなったら、ATS作動フラグを解除
                        print(f"[ATS.setSpeedCommand] ATS deactivated on train {train.id}")
                        self.__activated[train.id] = False

        # ATS無効時、指示された速度をそのまま設定
        else:
            train.targetSpeed = speedCommand

    def getATSStopPoint(self, train: Train) -> StopPoint:
        """
        赤信号によって生じる停止点を取得する。閉塞の手前や、開通していないポイントの手前が該当する

        Parameters
        ----------
        train: 列車
        """
        testSection = train.currentSection
        while True:
            # 青信号なら次のセクションへ
            if (
                self.__signalSystem.getSignal(
                    testSection.id, testSection.targetJunction.getOutSection().id
                ).value
                == "G"
            ):
                testSection = testSection.targetJunction.getOutSection()
            # 赤信号ならこのセクションの終わりを停止点として返す(停止余裕距離を引く)
            else:
                return StopPoint(testSection, testSection.length - self.__MERGIN)

    # ATSの有効/無効を切替
    def setEnabled(self, trainId: int, enabled: bool):
        self.__enabled[trainId] = enabled

    # ATS作動状況を取得
    def getActivated(self, trainId: int):
        return self.__activated[trainId]
