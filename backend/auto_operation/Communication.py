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
    arduino: Optional[serial.Serial]
    esp32Map: dict[int, serial.Serial]
    deltaMap: dict[int, float]
    sensorSignalBuffer: queue.Queue[int]

    def __init__(self, pidParamMap: dict[int, Train.PIDParam]) -> None:
        self.simulationMode = False
        self.simulationSpeedMap = {}
        self.pidParamMap = pidParamMap
        self.prevUpdate = 0.0
        self.arduino = None
        self.esp32Map = {}
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
            self.esp32Map[0] = serial.Serial("/dev/cu.ESP32-Dr", 115200)
            self.esp32Map[1] = serial.Serial("/dev/cu.ESP32-E6", 115200)
            self.deltaMap[0] = 0.0
            self.deltaMap[1] = 0.0
            self.deltaMap[2] = 0.0
            self.deltaMap[3] = 0.0
            self.arduino = serial.Serial("/dev/ttyS0", 9600)
        self.update()

    def update(self) -> None:
        now = time.time()
        dt = now - self.prevUpdate
        self.prevUpdate = now

        if self.simulationMode:
            for trainId in self.deltaMap.keys():
                self.deltaMap[trainId] += self.simulationSpeedMap[trainId] * dt

            if self.arduino != None:
                while self.arduino.in_waiting > 0:
                    self.sensorSignalBuffer.put(self.arduino.read())

        else:
            for trainId in self.esp32Map.keys():
                esp32 = self.esp32Map[trainId]
                if esp32 != None:
                    while esp32.in_waiting > 0:
                        # ホールセンサ信号が来たら、車輪0.5回転分deltaを進める
                        self.deltaMap[trainId] += 2 * pi * self.pidParamMap[trainId].r / 2
                        # 同時刻に複数の信号が来る不具合のため、1回のループですべて消費する
                        while esp32.in_waiting > 0:
                            esp32.read()
                else:  # 実機がない場合はsimulationを更新
                    self.deltaMap[trainId] += self.simulationSpeedMap[trainId] * dt

            if self.arduino != None:
                while self.arduino.in_waiting > 0:
                    self.sensorSignalBuffer.put(self.arduino.read())

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
            esp32 = self.esp32Map[trainId]
            if esp32 != None:
                if speed > 0.1:
                    INPUT_MIN = self.pidParamMap[trainId].INPUT_MIN
                    KP = self.pidParamMap[trainId].kp
                    input = int(INPUT_MIN + speed * KP)  # kp制御のみ
                else:
                    input = 0
                if esp32 != None:
                    esp32.write(input.to_bytes(1, "little"))
            else:  # ESP32の実機がないときはsimulationを更新
                self.simulationSpeedMap[trainId] = speed

    # 指定したポイントに切替命令を送る
    def sendToggle(self, servoId: int, servoState: Junction.ServoState) -> None:
        if self.arduino != None:
            servoStateCode = 0
            if servoState == Junction.ServoState.NoServo:
                return
            elif servoState == Junction.ServoState.Straight:
                servoStateCode = 0
            elif servoState == Junction.ServoState.Curve:
                servoStateCode = 1
            else:
                return
            self.arduino.write(servoId.to_bytes(1, "little"))
            self.arduino.write(servoStateCode.to_bytes(1, "little"))
            print(f"[Communication.sendToggle] servoId {servoId} toggle to {servoStateCode}")
