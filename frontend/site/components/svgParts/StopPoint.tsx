import React, { FC } from "react"
import { Point } from "../../types/svgPartsTypes"

interface Props {
  position: Point
  isStop: boolean
  onClick?: () => void
}

const StopPoint: FC<Props> = ({ position, isStop, onClick }) => (
  <g
    onClick={onClick}
    style={{
      cursor: onClick ? "pointer" : "initial",
    }}
  >
    <line
      x1={position.x}
      y1={position.y - 5}
      x2={position.x}
      y2={position.y + 5}
      stroke="black"
      strokeWidth={20}
      strokeLinecap="round"
    />
    <circle
      cx={position.x}
      cy={position.y - 6}
      r={5}
      fill={isStop ? "grey" : "limegreen"}
    />
    <circle
      cx={position.x}
      cy={position.y + 6}
      r={5}
      fill={isStop ? "red" : "grey"}
    />
  </g>
)

export default StopPoint
