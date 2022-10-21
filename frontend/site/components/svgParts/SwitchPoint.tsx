import React, { FC } from "react"
import { Point } from "../../types/svgPartsTypes"
import { degToRad, radToDeg } from "../../utils/util"

interface Props {
  position: Point
  fromAngle: number
  leftOutAngle: number
  rightOutAngle: number
  isLeft: boolean
  onClick?: () => void
}

const SwitchPoint: FC<Props> = ({
  position,
  fromAngle,
  leftOutAngle,
  rightOutAngle,
  isLeft,
  onClick,
}) => {
  const fromAnglePointX = position.x + Math.cos(degToRad(fromAngle)) * 10
  const fromAnglePointY = position.y + Math.sin(degToRad(fromAngle)) * 10
  const outAngle = isLeft ? leftOutAngle : rightOutAngle
  return (
    <g
      onClick={onClick}
      style={{
        cursor: onClick ? "pointer" : "initial",
      }}
    >
      <circle
        cx={position.x}
        cy={position.y}
        r={10}
        fill="white"
        stroke="black"
      />
      <line
        x1={position.x}
        y1={position.y}
        x2={fromAnglePointX}
        y2={fromAnglePointY}
        stroke="black"
      />
      <line
        x1={position.x}
        y1={position.y}
        x2={position.x + 10}
        y2={position.y}
        stroke="black"
        style={{
          transform: `rotate(${outAngle}deg)`,
          transformOrigin: `${position.x}px ${position.y}px`,
          transition: "transform 0.3s linear 0s",
        }}
      />
    </g>
  )
}

export default SwitchPoint
