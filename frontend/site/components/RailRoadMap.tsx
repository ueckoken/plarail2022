import React, { FC, useState } from "react"
import Rail from "./svgParts/Rail"
import Platform from "./svgParts/Platform"
import SwitchPoint from "./svgParts/SwitchPoint"
import StopPoint from "./svgParts/StopPoint"
import { Point, TrainData } from "../types/svgPartsTypes"
import {
  BlocklId,
  BunkiRailId,
  StationId,
  StopRailId,
} from "../types/control-messages"

interface Prop {
  datas: {
    stop: Record<StopRailId, boolean>
    switchState: Record<BunkiRailId, boolean>
    train1: TrainData
    blockState: Record<BlocklId, boolean>
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
  { position: { x: 700, y: 40 }, id: "shinjuku_s1" },
  { position: { x: 700, y: 100 }, id: "shinjuku_s2" },
  { position: { x: 950, y: 250 }, id: "sakurajosui_s0" },
  { position: { x: 1000, y: 400 }, id: "sakurajosui_s1" },
  { position: { x: 950, y: 400 }, id: "sakurajosui_s2" },
  { position: { x: 900, y: 400 }, id: "sakurajosui_s3" },
  { position: { x: 850, y: 400 }, id: "sakurajosui_s4" },
  { position: { x: 900, y: 500 }, id: "sakurajosui_s5" },
  { position: { x: 350, y: 550 }, id: "chofu_s0" },
  { position: { x: 350, y: 400 }, id: "chofu_s1" },
  { position: { x: 400, y: 400 }, id: "chofu_s2" },
  { position: { x: 300, y: 400 }, id: "chofu_s4" },
  { position: { x: 240, y: 400 }, id: "chofu_s3" },
  { position: { x: 90, y: 40 }, id: "hachioji_s2" },
  { position: { x: 90, y: 100 }, id: "hachioji_s1" },
  { position: { x: 90, y: 300 }, id: "hashimoto_s1" },
  { position: { x: 90, y: 240 }, id: "hashimoto_s2" },
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
    position: { x: 950, y: 300 },
    fromAngle: 270,
    leftOutAngle: 90,
    rightOutAngle: 0,
    id: "sakurajosui_p1",
  },
  {
    position: { x: 350, y: 500 },
    fromAngle: 90,
    leftOutAngle: 270,
    rightOutAngle: 0,
    id: "chofu_p1",
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
      rightDown: { x: 360, y: 140 },
    },
    id: "keiohachioji",
    name: "京王八王子付近",
  },
  {
    range: {
      leftUp: { x: 0, y: 200 },
      rightDown: { x: 360, y: 350 },
    },
    id: "kitano",
    name: "橋本付近",
  },
  {
    range: {
      leftUp: { x: 200, y: 300 },
      rightDown: { x: 620, y: 600 },
    },
    id: "chofu",
    name: "調布付近",
  },
  {
    range: {
      leftUp: { x: 800, y: 200 },
      rightDown: { x: 1200, y: 540 },
    },
    id: "meidaimae",
    name: "桜上水付近",
  },
  {
    range: {
      leftUp: { x: 500, y: 0 },
      rightDown: { x: 1020, y: 140 },
    },
    id: "sasazuka",
    name: "新宿付近",
  },
]

const RailroadMap: FC<Prop> = ({
  datas: { stop, switchState, train1, blockState },
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

        <Platform name="京王八王子" position={{ x: 90, y: 70 }} />
        <Platform name="橋本" position={{ x: 90, y: 270 }} />
        <Platform
          name="調布1"
          position={{ x: 375, y: 400 }}
          isHorizontal={false}
        />
        <Platform
          name="調布2"
          position={{ x: 270, y: 400 }}
          isHorizontal={false}
        />
        <Platform
          name="桜上水2"
          position={{ x: 875, y: 400 }}
          isHorizontal={false}
        />
        <Platform
          name="桜上水1"
          position={{ x: 975, y: 400 }}
          isHorizontal={false}
        />
        <Platform name="新宿" position={{ x: 700, y: 70 }} />

        <Rail
          // hachioji_b1
          positions={[
            { x: 60, y: 100 },
            { x: 20, y: 100 },
            { x: 20, y: 40 },
            { x: 120, y: 40 },
          ]}
          trains={[]}
          isClosed={blockState["hachioji_b1"]}
        />

        <Rail
          // hachioji_b2
          positions={[
            { x: 120, y: 40 },
            { x: 300, y: 40 },
            { x: 300, y: 500 },
          ]}
          trains={[]}
          isClosed={blockState["hachioji_b2"]}
        />

        <Rail
          // chofu_b3
          positions={[
            { x: 300, y: 470 },
            { x: 300, y: 550 },
            { x: 900, y: 550 },
            { x: 900, y: 470 },
          ]}
          trains={[]}
          isClosed={blockState["chofu_b5"]}
        />

        <Rail
          // sakurajosui_b3
          positions={[
            { x: 900, y: 470 },
            { x: 900, y: 300 },
          ]}
          trains={[]}
          isClosed={blockState["sakurajosui_b3"]}
        />

        <Rail
          // sakurajosui_b4
          positions={[
            { x: 900, y: 470 },
            { x: 850, y: 470 },
            { x: 850, y: 300 },
            { x: 900, y: 300 },
          ]}
          trains={[]}
          isClosed={blockState["sakurajosui_b4"]}
        />

        <Rail
          // sakurajosui_b6
          positions={[
            { x: 900, y: 300 },
            { x: 900, y: 100 },
            { x: 650, y: 100 },
          ]}
          trains={[]}
          isClosed={blockState["sakurajosui_b6"]}
        />

        <Rail
          // shinjuku_b2
          positions={[
            { x: 650, y: 100 },
            { x: 600, y: 100 },
            { x: 600, y: 40 },
            { x: 750, y: 40 },
          ]}
          trains={[]}
          isClosed={blockState["shinjuku_b2"]}
        />

        <Rail
          // shinjuku_b1
          positions={[
            { x: 750, y: 40 },
            { x: 950, y: 40 },
            { x: 950, y: 300 },
          ]}
          trains={[]}
          isClosed={blockState["shinjuku_b1"]}
        />

        <Rail
          // sakurajosui_b1
          positions={[
            { x: 950, y: 300 },
            { x: 1000, y: 300 },
            { x: 1000, y: 470 },
            { x: 950, y: 470 },
          ]}
          trains={[]}
          isClosed={blockState["sakurajosui_b1"]}
        />

        <Rail
          // sakurajosui_b2
          positions={[
            { x: 950, y: 300 },
            { x: 950, y: 500 },
          ]}
          trains={[]}
          isClosed={blockState["sakurajosui_b2"]}
        />

        <Rail
          // sakurajosui_b5
          positions={[
            { x: 950, y: 500 },
            { x: 950, y: 600 },
            { x: 350, y: 600 },
            { x: 350, y: 480 },
          ]}
          trains={[]}
          isClosed={blockState["sakurajosui_b5"]}
        />

        <Rail
          // chofu_b1
          positions={[
            { x: 350, y: 500 },
            { x: 350, y: 350 },
          ]}
          trains={[]}
          isClosed={blockState["chofu_b1"]}
        />

        <Rail
          // chofu_b4
          positions={[
            { x: 400, y: 350 },
            { x: 400, y: 100 },
            { x: 50, y: 100 },
          ]}
          trains={[]}
          isClosed={blockState["chofu_b4"]}
        />

        <Rail
          // chofu_b2
          positions={[
            { x: 350, y: 500 },
            { x: 400, y: 500 },
            { x: 400, y: 350 },
          ]}
          trains={[]}
          isClosed={blockState["chofu_b2"]}
        />

        <Rail
          // chofu_b3
          positions={[
            { x: 350, y: 350 },
            { x: 350, y: 300 },
            { x: 50, y: 300 },
          ]}
          trains={[]}
          isClosed={blockState["chofu_b3"]}
        />

        <Rail
          // hashimoto_b1
          positions={[
            { x: 50, y: 300 },
            { x: 20, y: 300 },
            { x: 20, y: 240 },
            { x: 130, y: 240 },
          ]}
          trains={[]}
          isClosed={blockState["hashimoto_b1"]}
        />

        <Rail
          // hashimoto_b2
          positions={[
            { x: 130, y: 240 },
            { x: 240, y: 240 },
            { x: 240, y: 500 },
            { x: 300, y: 500 },
          ]}
          trains={[]}
          isClosed={blockState["hashimoto_b2"]}
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
