import { useContext } from "react";
import { BlocksContext } from "../contexts";

interface SectionProps {
  id: string;
  points: { x: number; y: number }[];
}

const Section: React.FC<SectionProps> = ({ id, points }) => {
  const blocks = useContext(BlocksContext);

  const blocked = Boolean(blocks[id]);

  const shrinkedPointFirst = {
    x: points[0].x + Math.sign(points[1].x - points[0].x) * 2,
    y: points[0].y + Math.sign(points[1].y - points[0].y) * 2,
  };

  const n = points.length;

  const shrinkedPointLast = {
    x: points[n - 1].x + Math.sign(points[n - 2].x - points[n - 1].x) * 2,
    y: points[n - 1].y + Math.sign(points[n - 2].y - points[n - 1].y) * 2,
  };

  const shrinkedPoints = [
    shrinkedPointFirst,
    ...points.slice(1, -1),
    shrinkedPointLast,
  ];

  return (
    <polyline
      points={shrinkedPoints.map((p) => `${p.x},${p.y}`).join(" ")}
      fill="none"
      stroke={blocked ? "red" : "white"}
      strokeWidth={blocked ? 2 : 1}
      strokeLinecap="square"
    />
  );
};

export default Section;
