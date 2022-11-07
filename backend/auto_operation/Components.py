from enum import Enum


class Section:
    def __init__(self, id: int, sourceJunction: 'Junction', targetJunction: 'Junction', sourceServoState: 'Junction.ServoState', targetServoState: 'Junction.ServoState', length: float):
        self.id = id
        self.length = length
        self.station = None
        self.stationPosition = 0
        self.sourceJunction = sourceJunction
        self.sourceJunction.addOutSection(self, sourceServoState)
        self.targetJunction = targetJunction
        self.targetJunction.addInSection(self, targetServoState)

    def putStation(self, station: 'Station', stationPosition: float):
        self.station = station
        self.stationPosition = stationPosition


class Junction:
    class ServoState(Enum):
        NoServo = 0
        Straight = 1
        Curve = 2

        @staticmethod
        def invert(input: 'Junction.ServoState'):
            if input == Junction.ServoState.Straight:
                return Junction.ServoState.Curve
            elif input == Junction.ServoState.Curve:
                return Junction.ServoState.Straight
            else:
                return Junction.ServoState.NoServo

    def __init__(self, id: int, servoId: int):
        self.id = id
        self.servoId = servoId
        self.inSectionStraight = None
        self.inSectionCurve = None
        self.outSectionStraight = None
        self.outSectionCurve = None
        self.inServoState = Junction.ServoState.Straight
        self.outServoState = Junction.ServoState.Straight
        self.belongStation = None
        self.toggleRequested = False

    def addInSection(self, section, servoState):
        if servoState == Junction.ServoState.Straight:
            self.inSectionStraight = section
        elif servoState == Junction.ServoState.Curve:
            self.inSectionCurve = section
        else:  # NoServoの時はStraight側に接続
            self.inSectionStraight = section
            self.inSectionCurve = None
            self.inServoState = Junction.ServoState.NoServo

    def addOutSection(self, section, servoState):
        if servoState == Junction.ServoState.Straight:
            self.outSectionStraight = section
        elif servoState == Junction.ServoState.Curve:
            self.outSectionCurve = section
        else:  # NoServoの時はStraight側に接続
            self.outSectionStraight = section
            self.outSectionCurve = None
            self.outServoState = Junction.ServoState.NoServo

    def toggle(self):
        if self.inSectionStraight and self.inSectionCurve:  # IN側に2本入ってくる分岐点の場合
            self.inServoState = Junction.ServoState.invert(self.inServoState)  # 反転
            self.toggleRequested = True
        elif self.outSectionStraight and self.outSectionCurve:  # OUT側に2本入ってくる分岐点の場合
            self.outServoState = Junction.ServoState.invert(self.outServoState)  # 反転
            self.toggleRequested = True

    def setServoState(self, servoState: ServoState):
        if self.inSectionStraight and self.inSectionCurve:  # IN側に2本入ってくる分岐点の場合inServoStateをセット
            self.inServoState = servoState
        if self.outSectionStraight and self.outSectionCurve:  # OUT側に2本入ってくる分岐点の場合outServoStateをセット
            self.outServoState = servoState

    def getOutSection(self) -> Section:
        if self.outServoState == Junction.ServoState.Curve:
            return self.outSectionCurve
        else:
            return self.outSectionStraight

    def getInSection(self) -> Section:
        if self.inServoState == Junction.ServoState.Curve:
            return self.inSectionCurve
        else:
            return self.inSectionStraight


class Sensor:
    def __init__(self, id: int, belongSection: Section, position: float):
        self.id = id
        self.belongSection = belongSection
        self.position = position


class Station:
    def __init__(self, id: int, name: str):
        self.id = id
        self.name = name


class Train:
    class PIDParam:
        def __init__(self, r: float, INPUT_MIN: int, INPUT_MAX: int, INPUT_START: int, kp: float, ki: float, kd: float):
            self.r = r
            self.INPUT_MIN = INPUT_MIN
            self.INPUT_MAX = INPUT_MAX
            self.INPUT_START = INPUT_START
            self.kp = kp
            self.ki = ki
            self.kd = kd

    def __init__(self, id: int, initialSection: Section, initialPosition: float, pidParam: PIDParam):
        self.id = id
        self.targetSpeed = 0.0
        self.currentSection = initialSection
        self.mileage = initialPosition
        self.prevMileage = initialPosition
        self.pidParam = pidParam

    # 引数：進んだ距離
    def move(self, delta: float):
        self.prevMileage = self.mileage
        self.mileage += delta
        if (self.mileage >= self.currentSection.length):  # junctionを通過したとき
            self.mileage = self.mileage - self.currentSection.length
            self.currentSection = self.currentSection.targetJunction.getOutSection()
