import { FC } from "react"

type Props = {
  isBack: boolean
  cx: number
  cy: number
  r: number
  onChange: (isBack: boolean) => any
}

const ReverseHandle: FC<Props> = ({ cx, cy, r, isBack, onChange }) => (
  <>
    <rect
      x={cx - r * 1.75}
      y={cy - r * 4}
      width={r * 5}
      height={r * 8}
      fill="black"
    />
    <line
      x1={cx}
      y1={cy - r * 1.75}
      x2={cx}
      y2={cy + r * 1.75}
      stroke="#555"
      strokeWidth={r * 0.75}
      strokeLinecap="round"
    />
    <circle
      cx={cx}
      cy={cy - (isBack ? -1 : 1) * r}
      r={r}
      fill="black"
      stroke="rgba(255, 255, 255, 0.8)"
      strokeWidth={r / 10}
      onClick={() => {
        onChange(!isBack)
      }}
      style={{
        cursor: "pointer",
      }}
    />
    <g fontSize={r} dominantBaseline="middle" fill="white">
      <text x={cx + r * 1.5} y={cy - r * 1.5}>
        前
      </text>
      <text x={cx + r * 1.5} y={cy + r * 1.5}>
        後
      </text>
    </g>
  </>
)

export default ReverseHandle
