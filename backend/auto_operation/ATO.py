import time

from ATS import ATS
from Components import StopPoint, Train
from DiaPlanner import DiaPlanner
from SignalSystem import SignalSystem
from State import State

# ダイヤ情報に基づく自動運転を行うクラス


class ATO:
    def __init__(
        self,
        state: State,
        signalSystem: SignalSystem,
        ats: ATS,
        diaPlanner: DiaPlanner,
        MAXSPEED: float,
        MERGIN: float = 35,
    ) -> None:
        self.__ats = ats
        self.__signalSystem = signalSystem
        self.__state = state
        self.__diaPlanner = diaPlanner
        self.__prevUpdate = 0.0
        self.__enabled = {}  # 各列車の、ATO有効,無効が入る辞書. key=trainId, valued=enabled
        self.__arriveTime = {}  # 各列車が、直近に駅に到着した時刻を記録する辞書. key=trainId, value=float
        self.__MAXSPEED = MAXSPEED
        self.__MERGIN = MERGIN
        for train in state.trainList:
            self.__enabled[train.id] = True
            self.__arriveTime[train.id] = 0.0

    # ダイヤ情報をもとに、速度指令を自動的に更新する
    def update(self) -> None:
        now = time.time()
        dt = now - self.__prevUpdate
        self.__prevUpdate = now

        for train in self.__state.trainList:
            # 【到着判定】停止位置をまたいだとき、駅に到着or通過したと判定し、arriveTimeを更新する
            if train.currentSection.station is not None:
                stationPosition = train.currentSection.stationPosition
                if train.prevMileage < stationPosition and stationPosition <= train.mileage:
                    self.__arriveTime[train.id] = now

            # ATO有効時、停止点を計算し、出してよいスピードを求める
            if self.__enabled[train.id]:
                train.stopPoint = self.getATOStopPoint(train)
                distance = self.__state.getDistance(
                    train.currentSection,
                    train.mileage,
                    train.stopPoint.section,
                    train.stopPoint.mileage,
                )
                if distance > 100:
                    speedLimit = self.__MAXSPEED
                elif distance > 0:
                    speedLimit = (0.9 * distance / 100 + 0.1) * self.__MAXSPEED
                else:
                    speedLimit = 0.0
                speedCommand = min(
                    train.targetSpeed + self.__MAXSPEED * dt / 5, speedLimit
                )  # 加速時は緩やかに加速する(5秒で最高速度に到達)
            else:
                speedCommand = self.__MAXSPEED

            # 速度指令値をATSにセット
            self.__ats.setSpeedCommand(train.id, speedCommand)

    def getATOStopPoint(self, train: Train) -> StopPoint:
        testSection = train.currentSection
        while True:
            # 現在のセクションに駅がある
            if testSection.station is not None:
                diaOfThisStation = self.__diaPlanner.getDia(train.id, testSection.station.id)  # ダイヤ
                # 当該駅に列車がすでに到着/通過済みの場合
                if (
                    train.currentSection.id == testSection.id
                    and train.mileage >= testSection.stationPosition
                ):
                    stopDuration = time.time() - self.__arriveTime[train.id]  # 停車からの経過時間
                    departSignal = self.__signalSystem.getSignal(
                        train.currentSection.id,
                        train.currentSection.targetJunction.getOutSection().id,
                    )  # 出発信号機
                    # print(f"trainId={train.id}, stopDuration={stopDuration}, departSignal={departSignal.value}")
                    # 到着した駅が退避駅でない & 最低停車時間を過ぎた & 信号が青 なら次のセクションへ進む
                    if (
                        not diaOfThisStation.wait
                        and stopDuration >= diaOfThisStation.stopTime
                        and departSignal.value == "G"
                    ):
                        testSection = testSection.targetJunction.getOutSection()
                    # それ以外のときは現在のセクションにある駅までの距離が求める距離となる
                    else:
                        return StopPoint(testSection, testSection.stationPosition)

                # まだ当該駅に到着/通過していない場合
                else:
                    # 当該駅で退避または1秒以上停車するなら、当該駅までの距離を返す
                    if diaOfThisStation.wait or diaOfThisStation.stopTime > 1:
                        return StopPoint(testSection, testSection.stationPosition)
                    # 当該駅が通過駅なら次のセクションへ進む
                    else:
                        testSection = testSection.targetJunction.getOutSection()

            # 現在のセクションに駅がない
            else:
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

    # ATO無効の列車に対して、外部から速度を指令する. MAXSPEEDの制限のみが効く
    def setSpeed(self, trainId: int, speedCommand: int) -> None:
        self.__ats.setSpeedCommand(trainId, min(self.__MAXSPEED, speedCommand))

    # ATOの有効/無効を切り替える
    def setEnabled(self, trainId: int, enabled: bool):
        self.__enabled[trainId] = enabled
