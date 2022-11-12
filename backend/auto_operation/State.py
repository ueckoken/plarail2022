from numpy import pi
from Components import *
from Communication import *


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
        self.junctionList.append(Junction("0", -1))
        self.junctionList.append(Junction("1", -1))
        self.junctionList.append(Junction("2", 0))  # 対応するサーボID=0
        self.junctionList.append(Junction("3", -1))
        self.junctionList.append(Junction("4", -1))
        self.junctionList.append(Junction("5", -1))
        self.junctionList.append(Junction("6", -1))
        self.junctionList.append(Junction("7", -1))
        self.junctionList.append(Junction("8", -1))
        self.junctionList.append(Junction("9", -1))
        self.junctionList.append(Junction("10", -1))
        self.junctionList.append(Junction("11", -1))

        # Section(id, sourceJuncction, targetJuncction, sourceServoState, targetServoState, length)
        self.sectionList.append(Section("0", self.getJunctionById("11"), self.getJunctionById("0"), Junction.ServoState.NoServo, Junction.ServoState.NoServo, State.STRAIGHT_UNIT * 10))
        self.sectionList.append(Section("1", self.getJunctionById("0"), self.getJunctionById("1"), Junction.ServoState.NoServo, Junction.ServoState.NoServo, State.STRAIGHT_UNIT * 10))
        self.sectionList.append(Section("2", self.getJunctionById("1"), self.getJunctionById("2"), Junction.ServoState.NoServo, Junction.ServoState.NoServo, State.STRAIGHT_UNIT * 10))
        self.sectionList.append(Section("3", self.getJunctionById("2"), self.getJunctionById("3"), Junction.ServoState.Straight, Junction.ServoState.NoServo, State.STRAIGHT_UNIT * 10))
        self.sectionList.append(Section("4", self.getJunctionById("3"), self.getJunctionById("4"), Junction.ServoState.NoServo, Junction.ServoState.NoServo, State.STRAIGHT_UNIT * 10))
        self.sectionList.append(Section("5", self.getJunctionById("4"), self.getJunctionById("5"), Junction.ServoState.NoServo, Junction.ServoState.NoServo, State.STRAIGHT_UNIT * 10))
        self.sectionList.append(Section("6", self.getJunctionById("5"), self.getJunctionById("6"), Junction.ServoState.NoServo, Junction.ServoState.Curve, State.STRAIGHT_UNIT * 10))
        self.sectionList.append(Section("7", self.getJunctionById("2"), self.getJunctionById("7"), Junction.ServoState.Curve, Junction.ServoState.NoServo, State.STRAIGHT_UNIT * 10))
        self.sectionList.append(Section("8", self.getJunctionById("7"), self.getJunctionById("8"), Junction.ServoState.NoServo, Junction.ServoState.NoServo, State.STRAIGHT_UNIT * 10))
        self.sectionList.append(Section("9", self.getJunctionById("8"), self.getJunctionById("9"), Junction.ServoState.NoServo, Junction.ServoState.NoServo, State.STRAIGHT_UNIT * 10))
        self.sectionList.append(Section("10", self.getJunctionById("9"), self.getJunctionById("6"), Junction.ServoState.NoServo, Junction.ServoState.Straight, State.STRAIGHT_UNIT * 10))
        self.sectionList.append(Section("11", self.getJunctionById("6"), self.getJunctionById("10"), Junction.ServoState.NoServo, Junction.ServoState.NoServo, State.STRAIGHT_UNIT * 10))
        self.sectionList.append(Section("12", self.getJunctionById("10"), self.getJunctionById("11"), Junction.ServoState.NoServo, Junction.ServoState.NoServo, State.STRAIGHT_UNIT * 10))

        # Sensor(id, section, position)
        # self.sensorList.append(Sensor(0, self.getSectionById(1), State.STRAIGHT_UNIT * 2.5 + State.CURVE_UNIT * 2))
        # self.sensorList.append(Sensor(1, self.getSectionById(4), State.STRAIGHT_UNIT * 2 + State.CURVE_UNIT * 2))

        # Station(id, name)
        self.stationList.append(Station("0", "Shinjuku-down"))
        self.stationList.append(Station("1", "Sakurajosui-down"))
        self.stationList.append(Station("2", "Chofu-down"))
        self.stationList.append(Station("3", "Hashimoto-down"))
        self.stationList.append(Station("4", "Hachioji-down"))
        self.stationList.append(Station("5", "Hachioji-up"))
        self.stationList.append(Station("6", "Hashimoto-up"))
        self.stationList.append(Station("7", "Chofu-up"))
        self.stationList.append(Station("8", "Sakurajosui-up"))
        self.stationList.append(Station("9", "Shinjuku-up"))

        # section.putStation(station, stationPosition)
        self.getSectionById("0").putStation(self.getStationById("0"), State.STRAIGHT_UNIT * 8)  # section0: 新宿下り
        self.getSectionById("1").putStation(self.getStationById("1"), State.STRAIGHT_UNIT * 8)  # section1: 桜上水下り
        self.getSectionById("3").putStation(self.getStationById("2"), State.STRAIGHT_UNIT * 8)  # section3: 調布下り1番:橋本方面
        self.getSectionById("4").putStation(self.getStationById("3"), State.STRAIGHT_UNIT * 8)  # section4: 橋本下り
        self.getSectionById("5").putStation(self.getStationById("6"), State.STRAIGHT_UNIT * 8)  # section5: 橋本上り
        self.getSectionById("6").putStation(self.getStationById("7"), State.STRAIGHT_UNIT * 8)  # section6: 調布上り3番:橋本方面
        self.getSectionById("7").putStation(self.getStationById("2"), State.STRAIGHT_UNIT * 8)  # section7: 調布下り2番:八王子方面
        self.getSectionById("8").putStation(self.getStationById("4"), State.STRAIGHT_UNIT * 8)  # section8: 八王子下り
        self.getSectionById("9").putStation(self.getStationById("5"), State.STRAIGHT_UNIT * 8)  # section9: 八王子上り
        self.getSectionById("10").putStation(self.getStationById("7"), State.STRAIGHT_UNIT * 8)  # section10: 調布上り4番:八王子方面
        self.getSectionById("11").putStation(self.getStationById("8"), State.STRAIGHT_UNIT * 8)  # section11: 桜上水上り
        self.getSectionById("12").putStation(self.getStationById("9"), State.STRAIGHT_UNIT * 8)  # section12: 新宿上り

        # junction.belogStation
        self.getJunctionById("2").belongStation = self.getStationById("2")  # junction2は調布下り
        self.getJunctionById("6").belongStation = self.getStationById("7")  # junction6は調布上り

        # PIDParams(r: float, INPUT_MIN: int, INPUT_MAX: int, INPUT_START: int, kp: float, ki: float, kd: float)
        pidParam0 = Train.PIDParam(1.25, 40, 55, 68, 0.70, 0, 0)  # Dr. (maxinput: 40 + 0.70*40cm/s = 68)
        pidParam1 = Train.PIDParam(1.18, 55, 90, 70, 0.40, 0, 0)  # Raspi (maxinput: 50 + 0.50*40cm/s= 70)

        # Train(initialSection, initialPosition)
        self.trainList.append(Train(0, self.getSectionById("0"), State.STRAIGHT_UNIT * 4, pidParam0))  # 列車0をsection0に配置
        self.trainList.append(Train(1, self.getSectionById("8"), State.STRAIGHT_UNIT * 4, pidParam1))  # 列車1をsection8(八王子の手前)に配置:DiaPlannerで八王子行に指定しているため

        # start communication
        self.communication = Communication({0: pidParam0, 1: pidParam1})

        # 初回の着発番線に合わせてここにtoggleを書く
        self.communication.sendToggle(self.getJunctionById("2").servoId, Junction.ServoState.Straight)
        self.communication.sendToggle(self.getJunctionById("6").servoId, Junction.ServoState.Straight)

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

    def getJunctionById(self, id: str) -> Junction:
        return list(filter(lambda item: item.id == id, self.junctionList))[0]

    def getSectionById(self, id: str) -> Section:
        return list(filter(lambda item: item.id == id, self.sectionList))[0]

    def getSensorById(self, id: int) -> Sensor:
        return list(filter(lambda item: item.id == int.from_bytes(id,'little'), self.sensorList))[0]

    def getStationById(self, id: str) -> Station:
        return list(filter(lambda item: item.id == id, self.stationList))[0]

    def getStationBySectionId(self, sectionId: str) -> Station:
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
    # 2つの地点が同じsectionに存在する場合、s1>s2だと負の値を返す
    def getDistance(self, s1: Section, mileage1: float, s2: Section, mileage2: float, originalStartSection: Section = None) -> float:
        distance = 0
        testSection = s1

        # 何も見つけられずに最初の地点に戻ってきてしまった場合、終了
        if originalStartSection != None and s1.id == originalStartSection.id:
            return s1.length

        if originalStartSection == None:
            originalStartSection = s1
        while testSection.id != s2.id:
            distance += testSection.length
            if testSection.targetJunction.outSectionCurve == None:
                testSection = testSection.targetJunction.getOutSection()
            else:
                distanceFrom2OutJucntionToS2ViaStraight = self.getDistance(testSection.targetJunction.outSectionStraight, 0, s2, mileage2, originalStartSection)
                distanceFrom2OutJucntionToS2ViaCurve = self.getDistance(testSection.targetJunction.outSectionCurve, 0, s2, mileage2, originalStartSection)
                return distance - mileage1 + min(distanceFrom2OutJucntionToS2ViaStraight, distanceFrom2OutJucntionToS2ViaCurve)
            if testSection.id == originalStartSection.id:
                break  # 1周して戻ってきた場合は終了
        return distance - mileage1 + mileage2
