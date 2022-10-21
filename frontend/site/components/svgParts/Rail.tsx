import React, { FC } from "react"
import { Point, TrainData } from "../../types/svgPartsTypes"

interface Props {
  positions: [Point, Point, ...Point[]]
  trains: TrainData[]
}

const Rail: FC<Props> = ({ positions, trains }) => {
  const pointsText = positions
    .map((point: Point) => `${point.x}, ${point.y}`)
    .join("\n")
  return (
    <g>
      <polyline
        points={pointsText}
        stroke={trains.length !== 0 ? "blue" : "black"}
        fill="none"
      />
    </g>
  )
}

export default Rail
