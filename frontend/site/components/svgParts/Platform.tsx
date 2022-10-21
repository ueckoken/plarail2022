import React, { FC } from "react"
import { Point } from "../../types/svgPartsTypes"

interface Props {
  name: string
  position: Point
  isHorizontal?: boolean
  size?: number
}

const Platform: FC<Props> = ({ name, position, isHorizontal = true, size }) => {
  const stationSize = size ? size : 40
  const width = isHorizontal ? stationSize : 20
  const height = isHorizontal ? 20 : stationSize
  return (
    <g>
      <rect
        x={position.x - width / 2}
        y={position.y - height / 2}
        width={width}
        height={height}
        fill="white"
        stroke="black"
      />
      <text
        x={position.x}
        y={position.y}
        width={width}
        height={height}
        fontFamily="monospace"
        fontSize={20}
        writingMode={isHorizontal ? "horizontal-tb" : "vertical-lr"}
        textAnchor="middle"
        dominantBaseline="central"
      >
        {name}
      </text>
    </g>
  )
}

export default Platform
