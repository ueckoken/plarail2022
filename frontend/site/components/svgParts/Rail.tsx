import React, { FC } from "react"
import { Point, TrainData } from "../../types/svgPartsTypes"

interface Props {
  positions: [Point, Point, ...Point[]]
  trains: TrainData[],
  isClosed: boolean
}

const Rail: FC<Props> = ({ positions, trains, isClosed }) => {
  const pointsText = positions
    .map((point: Point) => `${point.x}, ${point.y}`)
    .join("\n")
  return (
    <g>
      <polyline
        points={pointsText}
        stroke={isClosed ? "red" : "black"}
        strokeWidth={isClosed ? "2" : "1"}
        fill="none"
      />
    </g>
  )
}

export default Rail
