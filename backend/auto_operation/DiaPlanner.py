from dataclasses import dataclass

from Components import Section, Station
from State import State

# ダイヤを管理、更新する
# Diaクラスの意味:
#   { trainId: 0,
#     stationId: 1,
#     wait: False,
#     stopTime: 0,
#     arriveSectionId: 2,
#     destSectionId: 4 }
#   -> 列車0は駅1で最低0秒停車=通過. section2に到着しsection4へ出発
#   { trainId: 1,
#     stationId: 1,
#     wait: True,
#     stopTime: 10,
#     arriveSectionId: 3,
#     destSectionId: 4 }
#   -> 列車1は駅1で退避. 最低10秒停車. section3に到着しsection4へ出発
#
#   まず、退避はない前提で組む。八王子に行くか調布に行くかは、arriveSectionIdとdestSectionIdを指定して決める


@dataclass
class Dia:
    trainId: int
    stationId: Station.StationId
    wait: bool
    stopTime: int
    arriveSectionId: Section.SectionId
    destSectionId: Section.SectionId


class DiaPlanner:
    def __init__(self, state: State) -> None:
        self.__state = state
        self.__autoUpdate = True  # 自動更新の有効/無効
        self.__diaList: list[Dia] = []  # ダイヤリスト
        for train in self.__state.trainList:  # trainIdとstationIdからダイヤリストを生成
            for station in self.__state.stationList:
                self.__diaList.append(Dia(train.id, station.id, False, 10, 0, 0))

    # ダイヤ自動更新の初期値を記述
    def setup(self) -> None:
        # 初期値として、列車0は八王子経由の特急とする
        self.setDia(0, "shinjuku_down", False, 5, "shinjuku_b2", "shinjuku_b1")
        self.setDia(0, "sakurajosui_down", False, 0, "sakurajosui_b2", "sakurajosui_b5")
        self.setDia(0, "chofu_down", False, 5, "chofu_b2", "chofu_b4")
        self.setDia(0, "hachioji_down", False, 5, "chofu_b4", "hachioji_b1")
        self.setDia(0, "hachioji_up", False, 5, "hachioji_b1", "hachioji_b2")
        self.setDia(0, "chofu_up", False, 5, "hachioji_b2", "chofu_b5")
        self.setDia(0, "sakurajosui_up", False, 0, "sakurajosui_b3", "sakurajosui_b6")
        self.setDia(0, "shinjuku_up", False, 5, "sakurajosui_b6", "shinjuku_b2")

        # 初期値として、列車1は八王子経由の普通とする
        self.setDia(1, "shinjuku_down", False, 5, "shinjuku_b2", "shinjuku_b1")
        self.setDia(1, "sakurajosui_down", True, 5, "sakurajosui_b1", "sakurajosui_b5")
        self.setDia(1, "chofu_down", False, 5, "chofu_b2", "chofu_b4")
        self.setDia(1, "hachioji_down", False, 5, "chofu_b4", "hachioji_b1")
        self.setDia(1, "hachioji_up", False, 5, "hachioji_b1", "hachioji_b2")
        self.setDia(1, "chofu_up", False, 5, "hachioji_b2", "chofu_b5")
        self.setDia(1, "sakurajosui_up", True, 5, "sakurajosui_b4", "sakurajosui_b6")
        self.setDia(1, "shinjuku_up", False, 5, "sakurajosui_b6", "shinjuku_b2")

        # 初期値として、列車2は橋本経由の特急とする
        self.setDia(2, "shinjuku_down", False, 5, "shinjuku_b2", "shinjuku_b1")
        self.setDia(2, "sakurajosui_down", False, 0, "sakurajosui_b2", "sakurajosui_b5")
        self.setDia(2, "chofu_down", False, 5, "chofu_b1", "chofu_b3")
        self.setDia(2, "hashimoto_down", False, 5, "chofu_b3", "hashimoto_b1")
        self.setDia(2, "hashimoto_up", False, 5, "hashimoto_b1", "hashimoto_b2")
        self.setDia(2, "chofu_up", False, 5, "hashimoto_b2", "chofu_b5")
        self.setDia(2, "sakurajosui_up", False, 0, "sakurajosui_b3", "sakurajosui_b6")
        self.setDia(2, "shinjuku_up", False, 5, "sakurajosui_b6", "shinjuku_b2")

        # 初期値として、列車3は橋本経由の普通とする
        self.setDia(3, "shinjuku_down", False, 5, "shinjuku_b2", "shinjuku_b1")
        self.setDia(3, "sakurajosui_down", True, 5, "sakurajosui_b1", "sakurajosui_b5")
        self.setDia(3, "chofu_down", False, 5, "chofu_b1", "chofu_b3")
        self.setDia(3, "hashimoto_down", False, 5, "chofu_b3", "hashimoto_b1")
        self.setDia(3, "hashimoto_up", False, 5, "hashimoto_b1", "hashimoto_b2")
        self.setDia(3, "chofu_up", False, 5, "hashimoto_b2", "chofu_b5")
        self.setDia(3, "sakurajosui_up", True, 5, "sakurajosui_b4", "sakurajosui_b6")
        self.setDia(3, "shinjuku_up", False, 5, "sakurajosui_b6", "shinjuku_b2")

    # ダイヤ自動更新のルールを記述. 毎update時によぶ
    def update(self) -> None:
        if self.__autoUpdate:
            pass  # まずはダイヤ更新なしで走らせてみる
            # # 今回は、駅1の待避線(section3)に退避列車(wait=True)がいる状況でsection4に列車が出て行ったとき、
            # # 追い抜き成功と判断し、train0と1の退避フラグを反転させる
            # waitingTrain = self.__state.getTrainInSection(self.__state.getSectionById(3))  # section3の列車を取得
            # if waitingTrain != None and self.getDia(waitingTrain.id, 1).wait == True:
            #     for train in self.__state.trainList:
            #         # sectionが変化した瞬間だけ、mileageがprevmileageより小さくなることを利用し、section4に入った瞬間を検知
            #         if train.currentSection.id == 4 and train.mileage < train.prevMileage:
            #             # 列車0,1のフラグを反転させる
            #             if self.getDia(0, 1).wait == True:  # 列車0は直前まで退避していた
            #                 self.setDia(0, 1, False, 0, 2, 4)  # 列車0を追い抜きに
            #                 self.setDia(1, 1, True, 10, 3, 4)  # 列車1を退避に
            #             else:
            #                 self.setDia(0, 1, True, 10, 3, 4)
            #                 self.setDia(1, 1, False, 0, 2, 4)
            #             print("[DiaPlanner.update] wait flag switched!")

            # 桜上水下り線で退避していた列車の退避終了処理
            waitingTrain = self.__state.getTrainInSection(
                self.__state.getSectionById("sakurajosui_b1")
            )  # 桜上水下り退避線にいる列車
            if waitingTrain is not None and self.getDia(waitingTrain.id, "sakurajosui_down").wait:
                for train in self.__state.trainList:
                    if train is None:
                        continue
                    # sectionが変化した瞬間だけ、mileageがprevmileageより小さくなることを利用し、桜上水を出発した下り列車がいるか判断。いれば、追い抜きに成功したので、退避していたほうの退避フラグをリセット
                    if (
                        train.currentSection is not None
                        and train.currentSection.id == "sakurajosui_b5"
                        and train.mileage < train.prevMileage
                    ):
                        self.setDia(
                            waitingTrain.id,
                            "sakurajosui_down",
                            False,
                            5,
                            "sakurajosui_b1",
                            "sakurajosui_b5",
                        )
                        print(f"[DiaPlanner.update] train {waitingTrain.id} wait finish!")
            # 退避を終えて出発した各停列車の退避フラグを再びTrueに戻す
            sakurajosuiDepartedTrain = self.__state.getTrainInSection(
                self.__state.getSectionById("sakurajosui_b5")
            )  # 桜上水下りを出発後の列車を取得
            if sakurajosuiDepartedTrain is not None and (
                sakurajosuiDepartedTrain.id == 1 or sakurajosuiDepartedTrain.id == 3
            ):  # 各停列車なら
                self.setDia(
                    sakurajosuiDepartedTrain.id,
                    "sakurajosui_down",
                    True,
                    5,
                    "sakurajosui_b1",
                    "sakurajosui_b5",
                )

            # 桜上水上り線で退避していた列車の退避終了処理
            waitingTrain = self.__state.getTrainInSection(
                self.__state.getSectionById("sakurajosui_b4")
            )  # 桜上水上り退避線にいる列車
            if waitingTrain is not None and self.getDia(waitingTrain.id, "sakurajosui_up").wait:
                for train in self.__state.trainList:
                    # sectionが変化した瞬間だけ、mileageがprevmileageより小さくなることを利用し、桜上水を出発した上り列車がいるか判断。いれば、追い抜きに成功したので、退避していたほうの退避フラグをリセット
                    if (
                        train.currentSection is not None
                        and train.currentSection.id == "sakurajosui_b6"
                        and train.mileage < train.prevMileage
                    ):
                        self.setDia(
                            waitingTrain.id,
                            "sakurajosui_up",
                            False,
                            5,
                            "sakurajosui_b4",
                            "sakurajosui_b6",
                        )
                        print(f"[DiaPlanner.update] train {waitingTrain.id} wait finish!")
            # 退避を終えて出発した各停列車の退避フラグを再びTrueに戻す
            sakurajosuiDepartedTrain = self.__state.getTrainInSection(
                self.__state.getSectionById("sakurajosui_b6")
            )  # 桜上水上りを出発後の列車を取得
            if sakurajosuiDepartedTrain is not None and (
                sakurajosuiDepartedTrain.id == 1 or sakurajosuiDepartedTrain.id == 3
            ):  # 各停列車なら
                self.setDia(
                    sakurajosuiDepartedTrain.id,
                    "sakurajosui_up",
                    True,
                    5,
                    "sakurajosui_b4",
                    "sakurajosui_b6",
                )

    # 指定した列車の、指定した駅に対するダイヤを取得
    def getDia(self, trainId: int, stationId: Station.StationId) -> Dia:
        result = list(
            filter(
                lambda x: (x.trainId == trainId and x.stationId == stationId),
                self.__diaList,
            )
        )
        return result[0]

    # 指定した列車の、指定した駅に対するダイヤを更新
    def setDia(
        self,
        trainId: int,
        stationId: Station.StationId,
        wait: bool,
        stopTime: int,
        arriveSectionId: Section.SectionId,
        destSectionId: Section.SectionId,
    ) -> None:
        for dia in self.__diaList:
            if dia.trainId == trainId and dia.stationId == stationId:
                dia.wait = wait
                dia.stopTime = stopTime
                dia.arriveSectionId = arriveSectionId
                dia.destSectionId = destSectionId
                break

    # ダイヤ自動更新を有効/無効化
    def setAutoUpdateEnabled(self, enabled: bool) -> None:
        self.__autoUpdate = enabled
