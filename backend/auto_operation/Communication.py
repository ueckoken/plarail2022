import queue
import time
from dataclasses import dataclass
from typing import Optional

import serial

from Components import Junction, Train

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
    sensorSignalBuffer: queue.Queue[int]

    def __init__(self, pidParamMap: dict[int, Train.PIDParam]) -> None:
        self.simulationMode = False
        self.simulationSpeedMap = {}
        self.pidParamMap = pidParamMap
        self.prevUpdate = 0.0
        self.deltaMap = {}
        self.sensorSignalBuffer = queue.Queue()

    def setup(self, simulationMode: bool) -> None:
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
        self.update()

    def update(self) -> None:
        now = time.time()
        dt = now - self.prevUpdate
        self.prevUpdate = now

        if self.simulationMode:
            for trainId in self.deltaMap.keys():
                self.deltaMap[trainId] += self.simulationSpeedMap[trainId] * dt

        else:
            # 何もしない
            pass

    def receiveTrainDelta(self, trainId: int) -> float:
        retval = self.deltaMap[trainId]
        self.deltaMap[trainId] = 0.0
        return retval

    def availableSensorSignal(self) -> int:
        return self.sensorSignalBuffer.qsize()

    def receiveSensorSignal(self) -> int:
        return self.sensorSignalBuffer.get()

    # 指定した車両にspeedを送る. PID制御もここで行う
    def sendSpeed(self, trainId: int, speed: float) -> None:
        if self.simulationMode:
            self.simulationSpeedMap[trainId] = speed
        else:
            # 何もしない
            pass

    # 指定したポイントに切替命令を送る
    def sendToggle(self, servoId: int, servoState: Junction.ServoState) -> None:
        # 何もしない
        pass
