import typing
from typing import Optional

from Communication import Communication
from Components import Junction, Section, Sensor, Station, Stop, Train


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
        # TODO: もっといい感じにストップレールのIDを導出する
        self.sectionIdToStopId: dict[Section.SectionId, Stop.StopId] = {}

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
        self.sensorList.append(
            Sensor("shinjuku_d1", self.getSectionById("shinjuku_b1"), State.STRAIGHT_UNIT * 2)
        )
        self.sensorList.append(
            Sensor("shinjuku_d2", self.getSectionById("shinjuku_b2"), State.STRAIGHT_UNIT * 2)
        )
        self.sensorList.append(
            Sensor("sakurajosui_d1", self.getSectionById("sakurajosui_b1"), State.STRAIGHT_UNIT * 2)
        )
        self.sensorList.append(
            Sensor("sakurajosui_d2", self.getSectionById("sakurajosui_b2"), State.STRAIGHT_UNIT * 2)
        )
        self.sensorList.append(
            Sensor("sakurajosui_d3", self.getSectionById("sakurajosui_b3"), State.STRAIGHT_UNIT * 2)
        )
        self.sensorList.append(
            Sensor("sakurajosui_d4", self.getSectionById("sakurajosui_b4"), State.STRAIGHT_UNIT * 2)
        )
        self.sensorList.append(
            Sensor("sakurajosui_d5", self.getSectionById("sakurajosui_b5"), State.STRAIGHT_UNIT * 2)
        )
        self.sensorList.append(
            Sensor("sakurajosui_d6", self.getSectionById("sakurajosui_b6"), State.STRAIGHT_UNIT * 2)
        )
        self.sensorList.append(
            Sensor("chofu_d1", self.getSectionById("chofu_b1"), State.STRAIGHT_UNIT * 2)
        )
        self.sensorList.append(
            Sensor("chofu_d2", self.getSectionById("chofu_b2"), State.STRAIGHT_UNIT * 2)
        )
        self.sensorList.append(
            Sensor("chofu_d3", self.getSectionById("chofu_b3"), State.STRAIGHT_UNIT * 2)
        )
        self.sensorList.append(
            Sensor("chofu_d4", self.getSectionById("chofu_b4"), State.STRAIGHT_UNIT * 2)
        )
        self.sensorList.append(
            Sensor("chofu_d5", self.getSectionById("chofu_b5"), State.STRAIGHT_UNIT * 2)
        )
        self.sensorList.append(
            Sensor("hashimoto_d1", self.getSectionById("hashimoto_b1"), State.STRAIGHT_UNIT * 2)
        )
        self.sensorList.append(
            Sensor("hashimoto_d2", self.getSectionById("hashimoto_b2"), State.STRAIGHT_UNIT * 2)
        )
        self.sensorList.append(
            Sensor("hachioji_d1", self.getSectionById("hachioji_b1"), State.STRAIGHT_UNIT * 2)
        )
        self.sensorList.append(
            Sensor("hachioji_d2", self.getSectionById("hachioji_b2"), State.STRAIGHT_UNIT * 2)
        )

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

        # 区間とストップレールの対応
        # TODO: もっといい感じにストップレールのIDを導出する
        self.sectionIdToStopId.update(
            {
                "shinjuku_b1": "sakurajosui_s0",
                "shinjuku_b2": "shinjuku_s1",
                "sakurajosui_b1": "sakurajosui_s1",
                "sakurajosui_b2": "sakurajosui_s2",
                "sakurajosui_b3": "sakurajosui_s3",
                "sakurajosui_b4": "sakurajosui_s4",
                "sakurajosui_b5": "chofu_s0",
                "sakurajosui_b6": "shinjuku_s2",
                "chofu_b1": "chofu_s1",
                "chofu_b2": "chofu_s2",
                "chofu_b3": "hashimoto_s1",
                "chofu_b4": "hachioji_s2",
                "chofu_b5": "sakurajosui_s5",
                "hashimoto_b1": "hashimoto_s2",
                "hashimoto_b2": "chofu_s3",
                "hachioji_b1": "chofu_s4",
                "hachioji_b2": "hachioji_s1",
            }
        )
        # 停止点に重複がないことを確認
        assert len(self.sectionIdToStopId.values()) == len(set(self.sectionIdToStopId.values()))

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
        while self.communication.availableSensorSignal() > 0:
            sensorData = self.communication.receiveSensorSignal()
            if sensorData is None:
                continue
            id = typing.cast(Sensor.SensorId, sensorData.sensorId)  # TODO: IDを検証する
            sensor = self.getSensorById(id)
            # センサに近づいてくる列車で一番近いものを取得
            approachingTrain = self.getApproachingTrain(sensor.belongSection, sensor.position)
            if approachingTrain is not None:
                approachingTrain.currentSection = sensor.belongSection
                approachingTrain.mileage = sensor.position + 1.0  # センサより僅かに先に進める（センサを通過したことを示すため）
                print(f"[State.update] sensor {sensor.id}: train{train.id} position updated")
            else:
                print(f"[State.update] sensor {sensor.id}: train is not detected")

    # Stateに格納されている状態を現実世界に送信する. 各種計算後に実行すること
    def sendCommand(self):
        # 車両への指令送信
        for train in self.trainList:
            self.communication.sendSpeed(train.id, train.targetSpeed)

        # ポイントへの指令送信
        for junction in self.junctionList:
            if junction.servoId > -1 and junction.toggleRequested:
                self.communication.sendToggle(junction.servoId, junction.outServoState)
                junction.toggleRequested = False
                # inServoStateは、実際にはサーボモーターがついていないので送信しない

    def getJunctionById(self, id: Junction.JunctionId) -> Junction:
        return list(filter(lambda item: item.id == id, self.junctionList))[0]

    def getSectionById(self, id: Section.SectionId) -> Section:
        return list(filter(lambda item: item.id == id, self.sectionList))[0]

    def getSensorById(self, id: int) -> Sensor:
        # https://github.com/python/mypy/issues/12682
        return list(filter(lambda item: item.id == int.from_bytes(id, "little"), self.sensorList))[0]  # type: ignore

    def getStationById(self, id: Station.StationId) -> Station:
        return list(filter(lambda item: item.id == id, self.stationList))[0]

    def getStationBySectionId(self, sectionId: Section.SectionId) -> Optional[Station]:
        return self.getSectionById(sectionId).station

    def getTrainById(self, id: int) -> Train:
        return list(filter(lambda item: item.id == id, self.trainList))[0]

    # 指定されたsectionにいる列車を返す。列車がいなければNoneを返す
    def getTrainInSection(self, section: Section) -> Optional[Train]:
        trains = list(filter(lambda train: train.currentSection.id == section.id, self.trainList))  # type: ignore
        if trains:
            return trains[0]
        else:
            return None

    # 指定された地点に近づいてくる列車で一番近いものを取得
    def getApproachingTrain(self, section: Section, mileage: float) -> Optional[Train]:
        testSection = section
        testedSectionId = []  # 一度確認したセクションIDを記録しておき、2回通ったら抜けられないとみてNoneを返す
        while True:
            approachingTrain = self.getTrainInSection(testSection)
            # testSectionに列車が存在しない、または存在しても引数で指定された地点より進んだ位置にいる（＝遠ざかっている）場合、探索セクションをひとつ前へ
            if approachingTrain is None or (
                approachingTrain.currentSection.id == section.id
                and approachingTrain.mileage > mileage
            ):
                testedSectionId.append(testSection.id)  # 一度確認したセクションIDを記録しておく
                testSection = testSection.sourceJunction.getInSection()  # ひとつ前のセクションに検索範囲を移す
                if testSection.id in testedSectionId:  # すでに確認済みの場合、一周して戻ってしまったので、列車はいない。Noneを返す
                    return None
            # 存在すればその列車を返す
            else:
                return approachingTrain

    # 線路上のある点からある点までの距離を返す
    # 2つの地点が同じsectionに存在する場合、s1>s2だと負の値を返す。ポイントの状態的に、ある点からある点にたどり着くのが不可能な場合、noneを返す
    def getDistance(self, s1: Section, mileage1: float, s2: Section, mileage2: float) -> float:
        distance = 0.0
        testSection: Section = s1
        # 一度通過したjucntionのidを記録しておく。同じjunctionを2回通った場合は到達不能と判定しnanを返す
        passedJunctionId: list[Junction.JunctionId] = []

        while testSection.id != s2.id:
            # junctionを2回目に通過した場合、一周してしまうのでnanを返す
            if testSection.targetJunction.id in passedJunctionId:
                return float("nan")
            else:
                passedJunctionId.append(testSection.targetJunction.id)  # 一度通過したjunctionのidを記録
                distance += testSection.length
                testSection = testSection.targetJunction.getOutSection()

        return distance - mileage1 + mileage2
