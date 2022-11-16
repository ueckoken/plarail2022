from dataclasses import dataclass
from enum import Enum
from typing import Literal, Optional


class Stop:
    StopId = Literal[
        "shinjuku_s1",
        "shinjuku_s2",
        "sakurajosui_s0",
        "sakurajosui_s1",
        "sakurajosui_s2",
        "sakurajosui_s3",
        "sakurajosui_s4",
        "sakurajosui_s5",
        "chofu_s0",
        "chofu_s1",
        "chofu_s2",
        "chofu_s3",
        "chofu_s4",
        "hashimoto_s1",
        "hashimoto_s2",
        "hachioji_s1",
        "hachioji_s2",
    ]


class Point:
    PointId = Literal[
        "sakurajosui_p1",
        "sakurajosui_p2",
        "sakurajosui_p3",
        "sakurajosui_p4",
        "chofu_p1",
        "chofu_p2",
    ]


class Section:
    SectionId = Literal[
        "shinjuku_b1",
        "shinjuku_b2",
        "sakurajosui_b1",
        "sakurajosui_b2",
        "sakurajosui_b3",
        "sakurajosui_b4",
        "sakurajosui_b5",
        "sakurajosui_b6",
        "chofu_b1",
        "chofu_b2",
        "chofu_b3",
        "chofu_b4",
        "chofu_b5",
        "hashimoto_b1",
        "hashimoto_b2",
        "hachioji_b1",
        "hachioji_b2",
    ]

    id: SectionId
    length: float
    station: Optional["Station"]
    stationPosition: float
    sourceJunction: "Junction"
    targetJunction: "Junction"

    def __init__(
        self,
        id: SectionId,
        sourceJunction: "Junction",
        targetJunction: "Junction",
        sourceServoState: "Junction.ServoState",
        targetServoState: "Junction.ServoState",
        length: float,
    ):
        self.id = id
        self.length = length
        self.station = None
        self.stationPosition = 0
        self.sourceJunction = sourceJunction
        self.sourceJunction.addOutSection(self, sourceServoState)
        self.targetJunction = targetJunction
        self.targetJunction.addInSection(self, targetServoState)

    def putStation(self, station: "Station", stationPosition: float):
        self.station = station
        self.stationPosition = stationPosition


class Junction:
    class ServoState(Enum):
        NoServo = 0
        Straight = 1
        Curve = 2

        @staticmethod
        def invert(input: "Junction.ServoState") -> "Junction.ServoState":
            if input == Junction.ServoState.Straight:
                return Junction.ServoState.Curve
            elif input == Junction.ServoState.Curve:
                return Junction.ServoState.Straight
            else:
                return Junction.ServoState.NoServo

    JunctionId = Literal[
        "shinjuku_j1",
        "shinjuku_j2",
        "sakurajosui_j1",
        "sakurajosui_j2",
        "sakurajosui_j3",
        "sakurajosui_j4",
        "chofu_j1",
        "chofu_j2",
        "chofu_j3",
        "chofu_j4",
        "hashimoto_j1",
        "hashimoto_j2",
        "hachioji_j1",
        "hachioji_j2",
    ]

    id: JunctionId
    servoId: int
    inSectionStraight: Optional["Section"]
    inSectionCurve: Optional["Section"]
    outSectionStraight: Optional["Section"]
    outSectionCurve: Optional["Section"]
    inServoState: ServoState
    outServoState: ServoState
    belongStation: Optional["Station"]
    toggleRequested: bool

    def __init__(self, id: JunctionId, servoId: int) -> None:
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

    def addInSection(self, section, servoState) -> None:
        if servoState == Junction.ServoState.Straight:
            self.inSectionStraight = section
        elif servoState == Junction.ServoState.Curve:
            self.inSectionCurve = section
        else:  # NoServoの時はStraight側に接続
            self.inSectionStraight = section
            self.inSectionCurve = None
            self.inServoState = Junction.ServoState.NoServo

    def addOutSection(self, section, servoState) -> None:
        if servoState == Junction.ServoState.Straight:
            self.outSectionStraight = section
        elif servoState == Junction.ServoState.Curve:
            self.outSectionCurve = section
        else:  # NoServoの時はStraight側に接続
            self.outSectionStraight = section
            self.outSectionCurve = None
            self.outServoState = Junction.ServoState.NoServo

    def toggle(self) -> None:
        if self.inSectionStraight and self.inSectionCurve:  # IN側に2本入ってくる分岐点の場合
            self.inServoState = Junction.ServoState.invert(self.inServoState)  # 反転
            self.toggleRequested = True
        elif self.outSectionStraight and self.outSectionCurve:  # OUT側に2本入ってくる分岐点の場合
            self.outServoState = Junction.ServoState.invert(self.outServoState)  # 反転
            self.toggleRequested = True

    def setServoState(self, servoState: ServoState) -> None:
        if self.inSectionStraight and self.inSectionCurve:  # IN側に2本入ってくる分岐点の場合inServoStateをセット
            self.inServoState = servoState
        if self.outSectionStraight and self.outSectionCurve:  # OUT側に2本入ってくる分岐点の場合outServoStateをセット
            self.outServoState = servoState

    def getOutSection(self) -> Optional[Section]:
        if self.outServoState == Junction.ServoState.Curve:
            return self.outSectionCurve
        else:
            return self.outSectionStraight

    def getInSection(self) -> Optional[Section]:
        if self.inServoState == Junction.ServoState.Curve:
            return self.inSectionCurve
        else:
            return self.inSectionStraight


@dataclass
class Sensor:
    SensorId = Literal[
        "shinjuku_d1",
        "shinjuku_d2",
        "sakurajosui_d1",
        "sakurajosui_d2",
        "sakurajosui_d3",
        "sakurajosui_d4",
        "sakurajosui_d5",
        "sakurajosui_d6",
        "chofu_d1",
        "chofu_d2",
        "chofu_d3",
        "chofu_d4",
        "chofu_d5",
        "hashimoto_d1",
        "hashimoto_d2",
        "hachioji_d1",
        "hachioji_d2",
    ]

    id: SensorId
    belongSection: "Section"
    position: float


@dataclass
class Station:
    StationId = Literal[
        "shinjuku_up",
        "shinjuku_down",
        "sakurajosui_up",
        "sakurajosui_down",
        "chofu_up",
        "chofu_down",
        "hashimoto_up",
        "hashimoto_down",
        "hachioji_up",
        "hachioji_down",
    ]

    id: StationId
    name: str


class Train:
    @dataclass
    class PIDParam:
        r: float
        INPUT_MIN: int
        INPUT_MAX: int
        INPUT_START: int
        kp: float
        ki: float
        kd: float

    id: int
    targetSpeed: float
    currentSection: Optional["Section"]
    mileage: float
    prevMileage: float
    pidParam: "PIDParam"

    def __init__(
        self,
        id: int,
        initialSection: Section,
        initialPosition: float,
        pidParam: PIDParam,
    ) -> None:
        self.id = id
        self.targetSpeed = 0.0
        self.currentSection = initialSection
        self.mileage = initialPosition
        self.prevMileage = initialPosition
        self.pidParam = pidParam
        self.stopPoint: StopPoint | None = None  # 停止点

    # 引数：進んだ距離
    def move(self, delta: float) -> None:
        self.prevMileage = self.mileage
        self.mileage += delta
        assert self.currentSection is not None
        if self.mileage >= self.currentSection.length:  # junctionを通過したとき
            self.mileage = self.mileage - self.currentSection.length
            self.currentSection = self.currentSection.targetJunction.getOutSection()


@dataclass
class StopPoint:
    section: Section
    mileage: float
