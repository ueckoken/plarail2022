from numpy import pi

from Communication import *
from Components import *


class State:
    STRAIGHT_UNIT = 21.5  # 直線レールの長さ[cm]
    CURVE_UNIT = 16.9  # 曲線レールの長さ[cm]

    # __init__で線路形状と車両の配置を定義する
    def __init__(self) -> None:
        self.junctionList: list[Junction] = []
        self.sectionList: list[Section] = []
        self.sensorList: list[Sensor] = []
        self.stationList: list[Station] = []
        self.trainList: list[Train] = []

        # Junction(id, servoId)
        self.junctionList.append(Junction("shinjuku_j1", -1))
        self.junctionList.append(Junction("shinjuku_j2", -1))
        self.junctionList.append(Junction("sakurajosui_j1", -1))
        self.junctionList.append(Junction("sakurajosui_j2", -1))
        self.junctionList.append(Junction("sakurajosui_j3", -1))
        self.junctionList.append(Junction("sakurajosui_j4", -1))
        self.junctionList.append(Junction("chofu_j1", -1))
        self.junctionList.append(Junction("chofu_j2", -1))
        self.junctionList.append(Junction("chofu_j3", -1))
        self.junctionList.append(Junction("chofu_j4", -1))
        self.junctionList.append(Junction("hashimoto_j1", -1))
        self.junctionList.append(Junction("hashimoto_j2", -1))
        self.junctionList.append(Junction("hachioji_j1", -1))
        self.junctionList.append(Junction("hachioji_j2", -1))

        # Section(id, sourceJuncction, targetJuncction, sourceServoState, targetServoState, length)
        self.sectionList.append(
            Section(
                "shinjuku_b1",
                self.getJunctionById("shinjuku_j1"),
                self.getJunctionById("sakurajosui_j1"),
                Junction.ServoState.NoServo,
                Junction.ServoState.NoServo,
                State.STRAIGHT_UNIT * 10,
            )
        )
        self.sectionList.append(
            Section(
                "shinjuku_b2",
                self.getJunctionById("shinjuku_j2"),
                self.getJunctionById("shinjuku_j1"),
                Junction.ServoState.NoServo,
                Junction.ServoState.NoServo,
                State.STRAIGHT_UNIT * 10,
            )
        )
        self.sectionList.append(
            Section(
                "sakurajosui_b1",
                self.getJunctionById("sakurajosui_j1"),
                self.getJunctionById("sakurajosui_j3"),
                Junction.ServoState.Curve,
                Junction.ServoState.Curve,
                State.STRAIGHT_UNIT * 10,
            )
        )
        self.sectionList.append(
            Section(
                "sakurajosui_b2",
                self.getJunctionById("sakurajosui_j1"),
                self.getJunctionById("sakurajosui_j3"),
                Junction.ServoState.Straight,
                Junction.ServoState.Straight,
                State.STRAIGHT_UNIT * 10,
            )
        )
        self.sectionList.append(
            Section(
                "sakurajosui_b3",
                self.getJunctionById("sakurajosui_j2"),
                self.getJunctionById("sakurajosui_j4"),
                Junction.ServoState.Straight,
                Junction.ServoState.Straight,
                State.STRAIGHT_UNIT * 10,
            )
        )
        self.sectionList.append(
            Section(
                "sakurajosui_b4",
                self.getJunctionById("sakurajosui_j2"),
                self.getJunctionById("sakurajosui_j4"),
                Junction.ServoState.Curve,
                Junction.ServoState.Curve,
                State.STRAIGHT_UNIT * 10,
            )
        )
        self.sectionList.append(
            Section(
                "sakurajosui_b5",
                self.getJunctionById("sakurajosui_j3"),
                self.getJunctionById("chofu_j1"),
                Junction.ServoState.NoServo,
                Junction.ServoState.NoServo,
                State.STRAIGHT_UNIT * 10,
            )
        )
        self.sectionList.append(
            Section(
                "sakurajosui_b6",
                self.getJunctionById("sakurajosui_j4"),
                self.getJunctionById("shinjuku_j2"),
                Junction.ServoState.NoServo,
                Junction.ServoState.NoServo,
                State.STRAIGHT_UNIT * 10,
            )
        )
        self.sectionList.append(
            Section(
                "chofu_b1",
                self.getJunctionById("chofu_j1"),
                self.getJunctionById("chofu_j3"),
                Junction.ServoState.Straight,
                Junction.ServoState.NoServo,
                State.STRAIGHT_UNIT * 10,
            )
        )
        self.sectionList.append(
            Section(
                "chofu_b2",
                self.getJunctionById("chofu_j1"),
                self.getJunctionById("chofu_j4"),
                Junction.ServoState.Curve,
                Junction.ServoState.NoServo,
                State.STRAIGHT_UNIT * 10,
            )
        )
        self.sectionList.append(
            Section(
                "chofu_b3",
                self.getJunctionById("chofu_j3"),
                self.getJunctionById("hashimoto_j1"),
                Junction.ServoState.NoServo,
                Junction.ServoState.NoServo,
                State.STRAIGHT_UNIT * 10,
            )
        )
        self.sectionList.append(
            Section(
                "chofu_b4",
                self.getJunctionById("chofu_j4"),
                self.getJunctionById("hachioji_j1"),
                Junction.ServoState.NoServo,
                Junction.ServoState.NoServo,
                State.STRAIGHT_UNIT * 10,
            )
        )
        self.sectionList.append(
            Section(
                "chofu_b5",
                self.getJunctionById("chofu_j2"),
                self.getJunctionById("sakurajosui_j2"),
                Junction.ServoState.NoServo,
                Junction.ServoState.NoServo,
                State.STRAIGHT_UNIT * 10,
            )
        )
        self.sectionList.append(
            Section(
                "hashimoto_b1",
                self.getJunctionById("hashimoto_j1"),
                self.getJunctionById("hashimoto_j2"),
                Junction.ServoState.NoServo,
                Junction.ServoState.NoServo,
                State.STRAIGHT_UNIT * 10,
            )
        )
        self.sectionList.append(
            Section(
                "hashimoto_b2",
                self.getJunctionById("hashimoto_j2"),
                self.getJunctionById("chofu_j2"),
                Junction.ServoState.NoServo,
                Junction.ServoState.Curve,
                State.STRAIGHT_UNIT * 10,
            )
        )
        self.sectionList.append(
            Section(
                "hachioji_b1",
                self.getJunctionById("hachioji_j1"),
                self.getJunctionById("hachioji_j2"),
                Junction.ServoState.NoServo,
                Junction.ServoState.NoServo,
                State.STRAIGHT_UNIT * 10,
            )
        )
        self.sectionList.append(
            Section(
                "hachioji_b2",
                self.getJunctionById("hachioji_j2"),
                self.getJunctionById("chofu_j2"),
                Junction.ServoState.NoServo,
                Junction.ServoState.Straight,
                State.STRAIGHT_UNIT * 10,
            )
        )

        # Sensor(id, section, position)
        # self.sensorList.append(Sensor(0, self.getSectionById(1), State.STRAIGHT_UNIT * 2.5 + State.CURVE_UNIT * 2))
        # self.sensorList.append(Sensor(1, self.getSectionById(4), State.STRAIGHT_UNIT * 2 + State.CURVE_UNIT * 2))

        # Station(id, name)
        self.stationList.append(Station("shinjuku_up", "shinjuku_up"))
        self.stationList.append(Station("shinjuku_down", "shinjuku_down"))
        self.stationList.append(Station("sakurajosui_up", "sakurajosui_up"))
        self.stationList.append(Station("sakurajosui_down", "sakurajosui_down"))
        self.stationList.append(Station("chofu_up", "chofu_up"))
        self.stationList.append(Station("chofu_down", "chofu_down"))
        self.stationList.append(Station("hashimoto_up", "hashimoto_up"))
        self.stationList.append(Station("hashimoto_down", "hashimoto_down"))
        self.stationList.append(Station("hachioji_up", "hachioji_up"))
        self.stationList.append(Station("hachioji_down", "hachioji_down"))

        # section.putStation(station, stationPosition)
        self.getSectionById("sakurajosui_b6").putStation(
            self.getStationById("shinjuku_up"), State.STRAIGHT_UNIT * 8
        )  # 新宿上り
        self.getSectionById("shinjuku_b2").putStation(
            self.getStationById("shinjuku_down"), State.STRAIGHT_UNIT * 8
        )  # 新宿下り
        self.getSectionById("sakurajosui_b1").putStation(
            self.getStationById("sakurajosui_down"), State.STRAIGHT_UNIT * 8
        )  # 桜上水下り1番線
        self.getSectionById("sakurajosui_b2").putStation(
            self.getStationById("sakurajosui_down"), State.STRAIGHT_UNIT * 8
        )  # 桜上水下り2番線
        self.getSectionById("chofu_b1").putStation(
            self.getStationById("chofu_down"), State.STRAIGHT_UNIT * 8
        )  # 調布下り1番線
        self.getSectionById("chofu_b2").putStation(
            self.getStationById("chofu_down"), State.STRAIGHT_UNIT * 8
        )  # 調布下り2番線
        self.getSectionById("chofu_b4").putStation(
            self.getStationById("hachioji_down"), State.STRAIGHT_UNIT * 8
        )  # 八王子下り
        self.getSectionById("hachioji_b1").putStation(
            self.getStationById("hachioji_up"), State.STRAIGHT_UNIT * 8
        )  # 八王子上り
        self.getSectionById("chofu_b3").putStation(
            self.getStationById("hashimoto_down"), State.STRAIGHT_UNIT * 8
        )  # 橋本下り
        self.getSectionById("hashimoto_b1").putStation(
            self.getStationById("hashimoto_up"), State.STRAIGHT_UNIT * 8
        )  # 橋本上り
        self.getSectionById("hachioji_b2").putStation(
            self.getStationById("chofu_up"), State.STRAIGHT_UNIT * 8
        )  # 調布上り4番線
        self.getSectionById("hashimoto_b2").putStation(
            self.getStationById("chofu_up"), State.STRAIGHT_UNIT * 8
        )  # 調布上り3番線
        self.getSectionById("sakurajosui_b3").putStation(
            self.getStationById("sakurajosui_up"), State.STRAIGHT_UNIT * 8
        )  # 桜上水上り3番線
        self.getSectionById("sakurajosui_b4").putStation(
            self.getStationById("sakurajosui_up"), State.STRAIGHT_UNIT * 8
        )  # 桜上水上り4番線

        # junction.belogStation
        self.getJunctionById("chofu_j1").belongStation = self.getStationById(
            "chofu_down"
        )  # junction2は調布下り
        self.getJunctionById("chofu_j2").belongStation = self.getStationById(
            "chofu_up"
        )  # junction6は調布上り
        self.getJunctionById("sakurajosui_j1").belongStation = self.getStationById(
            "sakurajosui_down"
        )
        self.getJunctionById("sakurajosui_j3").belongStation = self.getStationById(
            "sakurajosui_down"
        )
        self.getJunctionById("sakurajosui_j2").belongStation = self.getStationById("sakurajosui_up")
        self.getJunctionById("sakurajosui_j4").belongStation = self.getStationById("sakurajosui_up")

        # PIDParams(r: float, INPUT_MIN: int, INPUT_MAX: int, INPUT_START: int, kp: float, ki: float, kd: float)
        pidParam0 = Train.PIDParam(
            1.25, 40, 55, 68, 0.70, 0, 0
        )  # Dr. (maxinput: 40 + 0.70*40cm/s = 68)
        pidParam1 = Train.PIDParam(
            1.18, 55, 90, 70, 0.40, 0, 0
        )  # Raspi (maxinput: 50 + 0.50*40cm/s= 70)

        # Train(initialSection, initialPosition)
        self.trainList.append(
            Train(
                0,
                self.getSectionById("shinjuku_b2"),
                State.STRAIGHT_UNIT * 4,
                pidParam0,
            )
        )  # 列車0を新宿b2に配置
        self.trainList.append(
            Train(
                1,
                self.getSectionById("hachioji_b1"),
                State.STRAIGHT_UNIT * 4,
                pidParam1,
            )
        )  # 列車1を八王子b1(八王子の手前)に配置:DiaPlannerで八王子行に指定しているため
        self.trainList.append(
            Train(
                2,
                self.getSectionById("hashimoto_b1"),
                State.STRAIGHT_UNIT * 4,
                pidParam1,
            )
        )  # 列車2を橋本b2に配置
        self.trainList.append(
            Train(
                3,
                self.getSectionById("sakurajosui_b4"),
                State.STRAIGHT_UNIT * 4,
                pidParam1,
            )
        )  # 列車3を桜上水b4に配置

        # start communication
        self.communication = Communication({0: pidParam0, 1: pidParam1})

        # 初回の着発番線に合わせてここにtoggleを書く
        self.communication.sendToggle(
            self.getJunctionById("sakurajosui_j1").servoId, Junction.ServoState.Straight
        )
        self.communication.sendToggle(
            self.getJunctionById("sakurajosui_j2").servoId, Junction.ServoState.Straight
        )
        self.communication.sendToggle(
            self.getJunctionById("chofu_j1").servoId, Junction.ServoState.Curve
        )
        self.communication.sendToggle(
            self.getJunctionById("chofu_j2").servoId, Junction.ServoState.Straight
        )

    # 現実世界の状態を取得しStateに反映する. 定期的に実行すること
    def update(self):
        # 情報取得
        self.communication.update()

        # 列車位置更新
        for train in self.trainList:
            delta = self.communication.receiveTrainDelta(train.id)
            train.move(delta)

        # センサによる補正
        # while self.communication.availableSensorSignal() > 0:
        #     id = self.communication.receiveSensorSignal()
        #     sensor = self.getSensorById(id)
        #     # sensorと同じセクションにいるtrainを取得して位置を補正
        #     train = self.getTrainInSection(sensor.belongSection)
        #     if train != None:
        #         train.move(sensor.position - train.mileage)
        #         print(f"[State.update] sensor {sensor.id}: train{train.id} position calibrated")
        #     else:
        #         print(f"[State.update] sensor {sensor.id}: train is not detected")

    # Stateに格納されている状態を現実世界に送信する. 各種計算後に実行すること
    def sendCommand(self):
        # 車両への指令送信
        for train in self.trainList:
            self.communication.sendSpeed(train.id, train.targetSpeed)

        # ポイントへの指令送信
        for junction in self.junctionList:
            if junction.servoId > -1 and junction.toggleRequested == True:
                self.communication.sendToggle(junction.servoId, junction.outServoState)
                junction.toggleRequested = False
                # inServoStateは、実際にはサーボモーターがついていないので送信しない

    def getJunctionById(self, id: Junction.JunctionId) -> Junction:
        return list(filter(lambda item: item.id == id, self.junctionList))[0]

    def getSectionById(self, id: Section.SectionId) -> Section:
        return list(filter(lambda item: item.id == id, self.sectionList))[0]

    def getSensorById(self, id: int) -> Sensor:
        return list(filter(lambda item: item.id == int.from_bytes(id, "little"), self.sensorList))[
            0
        ]

    def getStationById(self, id: Station.StationId) -> Station:
        return list(filter(lambda item: item.id == id, self.stationList))[0]

    def getStationBySectionId(self, sectionId: Section.SectionId) -> Station:
        return self.getSectionById(sectionId).station

    def getTrainById(self, id: int) -> Train:
        return list(filter(lambda item: item.id == id, self.trainList))[0]

    # 指定されたsectionにいる列車を返す。列車がいなければNoneを返す
    def getTrainInSection(self, section: Section) -> Train:
        trains = list(filter(lambda train: train.currentSection.id == section.id, self.trainList))
        if trains != []:
            return trains[0]
        else:
            return None

    # 線路上のある点からある点までの距離を返す
    # 2つの地点が同じsectionに存在する場合、s1>s2だと負の値を返す。ポイントの状態的に、ある点からある点にたどり着くのが不可能な場合、noneを返す
    def getDistance(self, s1: Section, mileage1: float, s2: Section, mileage2: float) -> float:
        distance = 0
        testSection = s1
        passedJunctionId: list[int] = []  # 一度通過したjucntionのidを記録しておく。同じjunctionを2回通った場合は到達不能と判定しnanを返す

        while testSection.id != s2.id:
            if testSection.targetJunction.id in passedJunctionId:  # junctionを2回目に通過した場合、一周してしまうのでnanを返す
                return float('nan')
            else:
                passedJunctionId.append(testSection.targetJunction.id)  # 一度通過したjunctionのidを記録
                distance += testSection.length
                testSection = testSection.targetJunction.getOutSection()
            
        return distance - mileage1 + mileage2
