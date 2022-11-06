interface PlatformProps {
  position: { x: number; y: number };
}

const Platform: React.FC<PlatformProps> = ({ position }) => {
  const width = 30;
  const height = 10;
  return (
    <rect
      x={position.x - width / 2}
      y={position.y - height / 2}
      width={width}
      height={height}
      fill="white"
    />
  );
};

export default Platform;
