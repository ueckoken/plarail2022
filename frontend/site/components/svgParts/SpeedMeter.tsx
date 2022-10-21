import { FC, memo } from "react"
import { degToRad } from "../../utils/util"

type Props = {
  cx: number
  cy: number
  r: number
  value: number
  max: number
}

const MIN_RAD = degToRad(-45)
const MAX_RAD = degToRad(225)

const SpeedMeter: FC<Props> = ({ cx, cy, r, value, max }) => {
  const minLinePositionX = cx - Math.cos(MIN_RAD) * r
  const minLinePositionY = cy - Math.sin(MIN_RAD) * r
  const maxLinePositionX = cx - Math.cos(MAX_RAD) * r
  const maxLinePositionY = cy - Math.sin(MAX_RAD) * r
  const ratio = value / max
  const rad = (MAX_RAD - MIN_RAD) * ratio
  const strokeWidth = r / 30
  return (
    <g strokeWidth={strokeWidth}>
      <circle cx={cx} cy={cy} r={r} fill="white" stroke="black" />
      <line
        x1={cx}
        y1={cy}
        x2={minLinePositionX}
        y2={minLinePositionY}
        stroke="black"
      />
      <line
        x1={cx}
        y1={cy}
        x2={maxLinePositionX}
        y2={maxLinePositionY}
        stroke="black"
      />
      <circle cx={cx} cy={cy} r={r * 0.7} fill="white" />
      <line
        x1={cx}
        y1={cy}
        x2={minLinePositionX}
        y2={minLinePositionY}
        stroke="black"
        strokeWidth={3 * strokeWidth}
        style={{
          transform: `rotate(${rad}rad)`,
          transformOrigin: `${cx}px ${cy}px`,
          transition: "transform 2s linear 0s",
        }}
      />
    </g>
  )
}

export default memo(SpeedMeter)
