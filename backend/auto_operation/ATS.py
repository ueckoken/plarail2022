from Components import *
from State import *
from SignalSystem import *


class ATS:
    def __init__(self, state: State, signalSystem: SignalSystem, MAXSPEED: float, MERGIN: float=20, BREAKING_DISTANCE: float=10):
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
        signal = self.__signalSystem.getSignal(train.currentSection.id, train.currentSection.targetJunction.getOutSection().id)
        if signal.value == 'R':  # 赤信号の場合、信号までの距離から、非常停止できる最高速度を計算
            distance = self.__state.getDistance(train.currentSection, train.mileage, train.currentSection, train.currentSection.length)
            if distance > self.__BREAKING_DISTANCE + self.__MERGIN:
                speedLimit = self.__MAXSPEED + 1
            elif distance > self.__MERGIN:
                speedLimit = (distance - self.__MERGIN) / self.__BREAKING_DISTANCE * (self.__MAXSPEED + 1)
            else:
                speedLimit = 0
            if speedLimit < speedCommand:  # 非常停止できる速度を超えた速度が指示された場合にATS作動. 速度を強制的にspeedLimitに落とす
                if self.__activated[train.id] == False:
                    print(f"[ATS.setSpeedCommand] ATS activated on train {train.id} !")
                train.targetSpeed = speedLimit
                self.__activated[train.id] = True
                return
        else:
            if self.__activated[train.id] == True:
                print(f"[ATS.setSpeedCommand] ATS deactivated on train {train.id}")
            self.__activated[train.id] = False  # 信号が青になったらATS解除
            
        train.targetSpeed = speedCommand

    # ATSの有効/無効を切替
    def setEnabled(self, trainId: int, enabled: bool):
        self.__enabled[trainId] = enabled

    # ATS作動状況を取得
    def getActivated(self, trainId: int):
        return self.__activated[trainId]
