from Components import *
from State import *

from typing import Literal


class Signal:
    def __init__(self, sourceSectionId: int, targetSectionId: int, value: Literal['R', 'G']):
        self.sourceSectionId = sourceSectionId
        self.targetSectionId = targetSectionId
        self.value = value  # 信号機の色(赤と緑)


class SignalSystem:
    def __init__(self, state: State):
        self.__state = state

    # sourceSection から targetSection へ進む信号を取得
    def getSignal(self, sourceSectionId: int, targetSectionId: int) -> Signal:
        sourceSection = self.__state.getSectionById(sourceSectionId)
        junction = sourceSection.targetJunction
        # sourceからtargetへの経路が存在しない場合はNoneを返す
        if junction.outSectionStraight == None or junction.outSectionStraight.id != targetSectionId:
            if junction.outSectionCurve == None or junction.outSectionCurve.id != targetSectionId:
                return None
        # 信号判定
        if junction.getInSection().id == sourceSectionId:  # 分岐器が開通している
            train = self.__state.getTrainInSection(junction.getOutSection())  # 前方セクションにいる列車を取得
            if train == None:  # 前方セクションに在線なし
                return Signal(sourceSectionId, targetSectionId, 'G')
        return Signal(sourceSectionId, targetSectionId, 'R')

    # すべての信号を取得
    def getAllSignal(self) -> list[Signal]:
        # 全sectionの組み合わせについて判定し、Noneを削除
        result = list(map(lambda s, t: self.getSignal(s.id, t.id), self.__state.sectionList, self.__state.sectionList))
        return [i for i in result if i != None]
