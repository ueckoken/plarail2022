import time
from dataclasses import dataclass
from typing import Optional

from Components import Junction, Train
from Connection import AtsServicer, Connection

# ESP32 や Arduino との通信をまとめる。
# シミュレーションモードを使うと接続が無くてもある程度動作確認できる。


class Communication:
    @dataclass
    class TrainSignal:
        trainId: int

    simulationMode: bool
    simulationSpeedMap: dict[int, float]
    pidParamMap: dict[int, Train.PIDParam]
    prevUpdate: float
    deltaMap: dict[int, float]
    connection: Optional[Connection]

    def __init__(self, pidParamMap: dict[int, Train.PIDParam]) -> None:
        self.simulationMode = False
        self.simulationSpeedMap = {}
        self.pidParamMap = pidParamMap
        self.prevUpdate = 0.0
        self.deltaMap = {}
        self.connection = None

    def setup(self, simulationMode: bool, connection: Optional[Connection]) -> None:
        self.simulationMode = simulationMode
        if self.simulationMode:
            self.simulationSpeedMap[0] = 0.0
            self.simulationSpeedMap[1] = 0.0
            self.simulationSpeedMap[2] = 0.0
            self.simulationSpeedMap[3] = 0.0
            self.deltaMap[0] = 0.0
            self.deltaMap[1] = 0.0
            self.deltaMap[2] = 0.0
            self.deltaMap[3] = 0.0
        else:
            self.deltaMap[0] = 0.0
            self.deltaMap[1] = 0.0
            self.deltaMap[2] = 0.0
            self.deltaMap[3] = 0.0
            self.connection = connection
            if connection:
                connection.startServerThread()
        self.update()

    def update(self) -> None:
        now = time.time()
        dt = now - self.prevUpdate
        self.prevUpdate = now

        if self.simulationMode:
            # シミュレーションモードでは車両の位置を進める
            for trainId in self.deltaMap.keys():
                self.deltaMap[trainId] += self.simulationSpeedMap[trainId] * dt

        else:
            # 何もしない。
            # もともとは車輪の回転情報を受け取っていたが削除された。
            pass

    def receiveTrainDelta(self, trainId: int) -> float:
        retval = self.deltaMap[trainId]
        self.deltaMap[trainId] = 0.0
        return retval

    def availableSensorSignal(self) -> int:
        if self.connection:
            return self.connection.atsServicer.sensorQueue.qsize()
        else:
            return 0

    def receiveSensorSignal(self) -> Optional[AtsServicer.SensorData]:
        if self.connection:
            return self.connection.atsServicer.sensorQueue.get()
        else:
            return None

    # 指定した車両にspeedを送る. PID制御もここで行う
    def sendSpeed(self, trainId: int, speed: float) -> None:
        if self.simulationMode:
            self.simulationSpeedMap[trainId] = speed
        else:
            # 何もしない。
            # もともとは車両に速度指令を送ることができたが削除された。
            pass

    # 指定したポイントに切替命令を送る
    def sendToggle(self, servoId: int, servoState: Junction.ServoState) -> None:
        # 何もしない。
        # もともとはポイントに切替命令を送ることができたが削除された。
        pass
