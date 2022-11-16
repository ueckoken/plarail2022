import React, {
  FC,
  useEffect,
  useState,
  useRef,
  useCallback,
  CSSProperties,
} from "react"
import Peer, { SfuRoom, MeshRoom } from "skyway-js"
interface PeerIdProp {
  peerId: string
}
type MediaStreamWithPeerId = MediaStream & PeerIdProp
interface Prop {
  roomIds: string[]
  styles: CSSProperties[]
}
const SW_WSURL = "wss://webrtc.chofufes2022.ueckoken.club/"
const SKYWAY_APIKEY =
  process.env.SKYWAY_APIKEY === undefined
    ? "2eb379e0-0374-4e3c-9674-d7415b0b7f27"
    : process.env.SKYWAY_APIKEY
const SKYWAY_DEBUG_LEVEL = 2

const skyWayPeer: Peer = new Peer({
  key: SKYWAY_APIKEY,
  debug: SKYWAY_DEBUG_LEVEL,
});

const VideoCast: FC<Prop> = ({ roomIds, styles }) => {
  const [isPeerAvailable, setIsPeerAvailable] = useState<boolean>(false)
  useEffect(() => {
    skyWayPeer?.on("open", (id) => {
      console.log(`opened skyway peer id:${id}`)
      setIsPeerAvailable(true)
    })
    skyWayPeer?.on("close", () => {
      console.log("closed skyway peer")
      setIsPeerAvailable(false)
    })
    return () => {
      skyWayPeer?.destroy()
      setIsPeerAvailable(false)
    }
  }, [])

  return (
    <React.Fragment>
      {roomIds.map((roomId, index) => {
        if (
          skyWayPeer === undefined ||
          !isPeerAvailable
        ) {
          return (
            <p key={roomId}>Peer not available</p>
          )
        }
        return (
          <RoomViewer
            roomId={roomId}
            peer={skyWayPeer}
            key={roomId}
            style={styles[index]}
          />
        )
      })}
    </React.Fragment>
  )
}

type RoomViewerProps = {
  roomId: string
  peer: Peer
  style: CSSProperties
}

const RoomViewer: FC<RoomViewerProps> = ({ roomId, peer, style }) => {
  const [skywayRoom, setSkywayRoom] = useState<SfuRoom | MeshRoom | null>(null)
  const [peerId, setPeerId] = useState<string | null>(null)

  useEffect(() => {
    console.log("joinroom ",roomId)
    const skywayRoom = peer.joinRoom(roomId, {
      mode:"sfu"
    })
    skywayRoom.on("open", () => {
      console.log("sfuroom onopen")
    })
    setSkywayRoom(skywayRoom)
    setPeerId(peer.id)

    return () => {
      skywayRoom?.close()
      console.log("closed room", roomId)
    }
  }, [roomId, peer,  setSkywayRoom, setPeerId])
  if (skywayRoom === null || peerId === null) {
    return <p>NO STREAM</p>
  }
  return <SkywayRoomViewer room={skywayRoom} peerId={peerId} style={style} />
}

type SkywayRoomViewerProps = {
  room: SfuRoom | MeshRoom
  peerId: string
  style: CSSProperties
}

const SkywayRoomViewer: FC<SkywayRoomViewerProps> = ({
  room,
  style,
}) => {
  const ref = useRef<HTMLVideoElement>(null)
  const [castingStream, setCastingStream] = useState<MediaStreamWithPeerId>()
  const onStream = useCallback(
    (stream: MediaStreamWithPeerId) => {
      console.log("on stream", stream)
      setCastingStream(stream)
    },
    [setCastingStream]
  )
  useEffect(() => {
    room.on("stream", onStream)
    return () => {
      room.off("stream", onStream)
    }
  }, [room, onStream])
  useEffect(() => {
    const video = ref.current
    if (video === null || castingStream === undefined) {
      return
    }
    if (video.srcObject !== castingStream) {
      video.srcObject = castingStream
      video.playsInline = true
      try {
        video.play()
      } catch (err) {
        console.error(err)
      }
      video.volume = 0
    }
  }, [castingStream])
  return <video style={style} ref={ref} autoPlay muted></video>
}

export default VideoCast
