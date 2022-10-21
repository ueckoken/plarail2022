import React, { FC, useState } from "react"
import Rail from "./svgParts/Rail"
import Platform from "./svgParts/Platform"
import SwitchPoint from "./svgParts/SwitchPoint"
import StopPoint from "./svgParts/StopPoint"
import { Point, TrainData } from "../types/svgPartsTypes"
import { BunkiRailId, StationId, StopRailId } from "../types/control-messages"

interface Prop {
  datas: {
    stop: Record<StopRailId, boolean>
    switchState: Record<BunkiRailId, boolean>
    train1: TrainData
  }
  onStopPointOrSwitchPointClick?: (stationId: StationId) => any
}

type StopPointPosition = {
  id: StopRailId
  position: Point
}

type MapRange = {
  range: Rectangle
  id: string
  name: string
}

const STOP_PONINTS: StopPointPosition[] = [
  { position: { x: 940, y: 560 }, id: "motoyawata_s1" },
  { position: { x: 1000, y: 560 }, id: "motoyawata_s2" },
  { position: { x: 120, y: 20 }, id: "kitano_s6" },
  { position: { x: 240, y: 20 }, id: "kitano_s4" },
  { position: { x: 240, y: 60 }, id: "kitano_s3" },
  { position: { x: 240, y: 80 }, id: "kitano_s2" },
  { position: { x: 240, y: 120 }, id: "kitano_s1" },
  { position: { x: 520, y: 20 }, id: "chofu_s4" },
  { position: { x: 520, y: 60 }, id: "chofu_s3" },
  { position: { x: 520, y: 120 }, id: "chofu_s2" },
  { position: { x: 520, y: 160 }, id: "chofu_s1" },
  { position: { x: 700, y: 60 }, id: "meidaimae_s2" },
  { position: { x: 700, y: 80 }, id: "meidaimae_s1" },
  { position: { x: 800, y: 70 }, id: "sasazuka_s5" },
  { position: { x: 880, y: 20 }, id: "sasazuka_s4" },
  { position: { x: 880, y: 60 }, id: "sasazuka_s3" },
  { position: { x: 880, y: 80 }, id: "sasazuka_s2" },
  { position: { x: 880, y: 120 }, id: "sasazuka_s1" },
  { position: { x: 120, y: 140 }, id: "kitano_s5" },
  { position: { x: 180, y: 200 }, id: "takao_s1" },
  { position: { x: 120, y: 380 }, id: "takao_s2" },
  { position: { x: 320, y: 180 }, id: "chofu_s5" },
  { position: { x: 940, y: 180 }, id: "kudanshita_s5" },
  { position: { x: 1000, y: 180 }, id: "kudanshita_s6" },
  { position: { x: 880, y: 380 }, id: "iwamotocho_s1" },
  { position: { x: 940, y: 380 }, id: "iwamotocho_s2" },
  { position: { x: 1000, y: 380 }, id: "iwamotocho_s4" },
]

type SwitchPointPotiionAndAngle = {
  id: BunkiRailId
  position: Point
  fromAngle: number
  leftOutAngle: number
  rightOutAngle: number
}
const SWITCH_POINTS: SwitchPointPotiionAndAngle[] = [
  {
    position: { x: 160, y: 20 },
    fromAngle: 180,
    leftOutAngle: 0,
    rightOutAngle: 45,
    id: "kitano_b2",
  },
  {
    position: { x: 320, y: 120 },
    fromAngle: 0,
    leftOutAngle: 180,
    rightOutAngle: 225,
    id: "kitano_b1",
  },
  {
    position: { x: 380, y: 20 },
    fromAngle: 180,
    leftOutAngle: 45,
    rightOutAngle: 0,
    id: "chofu_b5",
  },
  {
    position: { x: 440, y: 60 },
    fromAngle: 180,
    leftOutAngle: -45,
    rightOutAngle: 0,
    id: "chofu_b4",
  },
  {
    position: { x: 420, y: 160 },
    fromAngle: 0,
    leftOutAngle: 225,
    rightOutAngle: 180,
    id: "chofu_b3",
  },
  {
    position: { x: 480, y: 120 },
    fromAngle: 0,
    leftOutAngle: 180,
    rightOutAngle: 135,
    id: "chofu_b2",
  },
  {
    position: { x: 600, y: 120 },
    fromAngle: 0,
    leftOutAngle: 135,
    rightOutAngle: 180,
    id: "chofu_b1",
  },
  {
    position: { x: 800, y: 20 },
    fromAngle: 180,
    leftOutAngle: 0,
    rightOutAngle: 45,
    id: "sasazuka_b2",
  },
  {
    position: { x: 840, y: 80 },
    fromAngle: 0,
    leftOutAngle: 180,
    rightOutAngle: 135,
    id: "sasazuka_b1",
  },
  {
    position: { x: 1000, y: 280 },
    fromAngle: 270,
    leftOutAngle: 90,
    rightOutAngle: 135,
    id: "iwamotocho_b4",
  },
  {
    position: { x: 940, y: 480 },
    fromAngle: 90,
    leftOutAngle: 270,
    rightOutAngle: 225,
    id: "iwamotocho_b1",
  },
]

type Rectangle = {
  leftUp: Point
  rightDown: Point
}

const MAP_RANGES: MapRange[] = [
  {
    range: {
      leftUp: { x: 0, y: 0 },
      rightDown: { x: 1120, y: 620 },
    },
    id: "all",
    name: "全体地図",
  },
  {
    range: {
      leftUp: { x: 0, y: 0 },
      rightDown: { x: 180, y: 140 },
    },
    id: "keiohachioji",
    name: "京王八王子付近",
  },
  {
    range: {
      leftUp: { x: 120, y: 0 },
      rightDown: { x: 360, y: 140 },
    },
    id: "kitano",
    name: "北野付近",
  },
  {
    range: {
      leftUp: { x: 360, y: 0 },
      rightDown: { x: 620, y: 180 },
    },
    id: "chofu",
    name: "調布付近",
  },
  {
    range: {
      leftUp: { x: 620, y: 0 },
      rightDown: { x: 800, y: 140 },
    },
    id: "meidaimae",
    name: "明大前付近",
  },
  {
    range: {
      leftUp: { x: 780, y: 0 },
      rightDown: { x: 1020, y: 140 },
    },
    id: "sasazuka",
    name: "笹塚付近",
  },
  {
    range: {
      leftUp: { x: 780, y: 0 },
      rightDown: { x: 1020, y: 140 },
    },
    id: "shinjuku",
    name: "新宿付近",
  },
  {
    range: {
      leftUp: { x: 100, y: 100 },
      rightDown: { x: 200, y: 280 },
    },
    id: "takao",
    name: "高尾付近",
  },
  {
    range: {
      leftUp: { x: 100, y: 300 },
      rightDown: { x: 200, y: 440 },
    },
    id: "takaosanguchi",
    name: "高尾山口付近",
  },
  {
    range: {
      leftUp: { x: 300, y: 140 },
      rightDown: { x: 400, y: 340 },
    },
    id: "wakabadai",
    name: "若葉台付近",
  },
  {
    range: {
      leftUp: { x: 300, y: 320 },
      rightDown: { x: 400, y: 440 },
    },
    id: "hashimoto",
    name: "橋本付近",
  },
  {
    range: {
      leftUp: { x: 920, y: 140 },
      rightDown: { x: 1020, y: 260 },
    },
    id: "kudanshita",
    name: "九段下付近",
  },
  {
    range: {
      leftUp: { x: 860, y: 260 },
      rightDown: { x: 1020, y: 500 },
    },
    id: "iwamotocho",
    name: "岩本町付近",
  },
  {
    range: {
      leftUp: { x: 920, y: 500 },
      rightDown: { x: 1020, y: 620 },
    },
    id: "motoyawata",
    name: "本八幡付近",
  },
]

const RailroadMap: FC<Prop> = ({
  datas: { stop, switchState, train1 },
  onStopPointOrSwitchPointClick,
}) => {
  const [mapRange, setMapRange] = useState<MapRange>(MAP_RANGES[0])
  return (
    <>
      <svg
        width="100%"
        viewBox={`${mapRange.range.leftUp.x} ${mapRange.range.leftUp.y} ${
          mapRange.range.rightDown.x - mapRange.range.leftUp.x
        } ${mapRange.range.rightDown.y - mapRange.range.leftUp.y}`}
      >
        <rect x={0} y={0} width={1120} height={620} fill="lightgray" />

        <Platform name="京王八王子" position={{ x: 60, y: 70 }} />
        <Platform name="北野1" position={{ x: 240, y: 40 }} />
        <Platform name="北野2" position={{ x: 240, y: 100 }} />
        <Platform name="調布1" position={{ x: 520, y: 40 }} />
        <Platform name="調布2" position={{ x: 520, y: 140 }} />
        <Platform name="明大前1" position={{ x: 700, y: 40 }} />
        <Platform name="明大前2" position={{ x: 700, y: 100 }} />
        <Platform name="笹塚1" position={{ x: 880, y: 40 }} />
        <Platform name="笹塚2" position={{ x: 880, y: 100 }} />
        <Platform name="新宿" position={{ x: 1060, y: 70 }} />

        <Platform
          name="高尾"
          position={{ x: 150, y: 200 }}
          isHorizontal={false}
        />
        <Platform
          name="高尾山口"
          position={{ x: 150, y: 380 }}
          isHorizontal={false}
        />
        <Platform
          name="若葉台1"
          position={{ x: 320, y: 260 }}
          isHorizontal={false}
        />
        <Platform
          name="若葉台2"
          position={{ x: 380, y: 260 }}
          isHorizontal={false}
        />
        <Platform
          name="橋本"
          position={{ x: 350, y: 380 }}
          isHorizontal={false}
        />

        <Platform
          name="九段下"
          position={{ x: 970, y: 180 }}
          isHorizontal={false}
        />
        <Platform
          name="岩本町1"
          position={{ x: 910, y: 380 }}
          isHorizontal={false}
        />
        <Platform
          name="岩本町2"
          position={{ x: 970, y: 380 }}
          isHorizontal={false}
        />
        <Platform
          name="本八幡"
          position={{ x: 970, y: 560 }}
          isHorizontal={false}
        />

        <Rail
          positions={[
            { x: 120, y: 120 },
            { x: 20, y: 120 },
            { x: 20, y: 20 },
            { x: 140, y: 20 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 120, y: 120 },
            { x: 20, y: 120 },
            { x: 20, y: 20 },
            { x: 140, y: 20 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 140, y: 20 },
            { x: 160, y: 20 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 140, y: 20 },
            { x: 140, y: 100 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 140, y: 100 },
            { x: 120, y: 120 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 120, y: 120 },
            { x: 120, y: 240 },
            { x: 150, y: 270 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 150, y: 270 },
            { x: 150, y: 310 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 150, y: 310 },
            { x: 120, y: 340 },
            { x: 120, y: 420 },
            { x: 180, y: 420 },
            { x: 180, y: 340 },
            { x: 150, y: 310 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 150, y: 270 },
            { x: 180, y: 240 },
            { x: 180, y: 120 },
            { x: 320, y: 120 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 160, y: 20 },
            { x: 320, y: 20 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 160, y: 20 },
            { x: 200, y: 60 },
            { x: 280, y: 60 },
            { x: 320, y: 20 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 320, y: 20 },
            { x: 380, y: 20 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 140, y: 100 },
            { x: 160, y: 80 },
            { x: 280, y: 80 },
            { x: 320, y: 120 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 320, y: 120 },
            { x: 380, y: 120 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 380, y: 120 },
            { x: 420, y: 160 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 420, y: 160 },
            { x: 440, y: 160 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 440, y: 160 },
            { x: 480, y: 120 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 380, y: 120 },
            { x: 480, y: 120 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 480, y: 120 },
            { x: 600, y: 120 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 440, y: 160 },
            { x: 560, y: 160 },
            { x: 600, y: 120 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 380, y: 20 },
            { x: 480, y: 20 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 380, y: 20 },
            { x: 420, y: 60 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 420, y: 60 },
            { x: 440, y: 60 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 440, y: 60 },
            { x: 480, y: 20 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 480, y: 20 },
            { x: 600, y: 20 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 440, y: 60 },
            { x: 560, y: 60 },
            { x: 600, y: 20 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 420, y: 60 },
            { x: 320, y: 160 },
            { x: 320, y: 200 },
            { x: 340, y: 220 },
            { x: 340, y: 300 },
            { x: 320, y: 320 },
            { x: 320, y: 420 },
            { x: 380, y: 420 },
            { x: 380, y: 320 },
            { x: 360, y: 300 },
            { x: 360, y: 220 },
            { x: 380, y: 200 },
            { x: 380, y: 160 },
            { x: 420, y: 160 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 600, y: 20 },
            { x: 620, y: 20 },
            { x: 660, y: 60 },
            { x: 740, y: 60 },
            { x: 780, y: 20 },
            { x: 800, y: 20 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 600, y: 120 },
            { x: 620, y: 120 },
            { x: 660, y: 80 },
            { x: 740, y: 80 },
            { x: 780, y: 120 },
            { x: 800, y: 120 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 800, y: 20 },
            { x: 840, y: 60 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 800, y: 120 },
            { x: 840, y: 80 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 800, y: 20 },
            { x: 1100, y: 20 },
            { x: 1100, y: 120 },
            { x: 800, y: 120 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 840, y: 60 },
            { x: 800, y: 60 },
            { x: 800, y: 80 },
            { x: 840, y: 80 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 840, y: 60 },
            { x: 1000, y: 60 },
            { x: 1000, y: 280 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 840, y: 80 },
            { x: 940, y: 80 },
            { x: 940, y: 280 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 940, y: 280 },
            { x: 880, y: 340 },
            { x: 880, y: 420 },
            { x: 940, y: 480 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 940, y: 280 },
            { x: 940, y: 340 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 1000, y: 280 },
            { x: 940, y: 340 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 940, y: 340 },
            { x: 940, y: 420 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 940, y: 420 },
            { x: 1000, y: 480 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 1000, y: 280 },
            { x: 1000, y: 480 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 940, y: 420 },
            { x: 940, y: 480 },
          ]}
          trains={[]}
        />
        <Rail
          positions={[
            { x: 940, y: 480 },
            { x: 940, y: 600 },
            { x: 1000, y: 600 },
            { x: 1000, y: 480 },
          ]}
          trains={[]}
        />

        {STOP_PONINTS.map(({ position, id }) => (
          <StopPoint
            position={position}
            isStop={stop[id]}
            key={id}
            onClick={() => onStopPointOrSwitchPointClick?.(id)}
          />
        ))}

        {SWITCH_POINTS.map(
          ({ id, position, fromAngle, leftOutAngle, rightOutAngle }) => (
            <SwitchPoint
              position={position}
              fromAngle={fromAngle}
              leftOutAngle={leftOutAngle}
              rightOutAngle={rightOutAngle}
              isLeft={!switchState[id]}
              onClick={() => onStopPointOrSwitchPointClick?.(id)}
              key={id}
            />
          )
        )}
      </svg>
      {MAP_RANGES.map((map, i) => (
        <button
          key={`mapchange-${map.id}`}
          onClick={() => setMapRange(MAP_RANGES[i])}
        >
          {map.name}
        </button>
      ))}
    </>
  )
}

export default RailroadMap
