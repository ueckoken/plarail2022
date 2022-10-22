import * as Ayame from "@open-ayame/ayame-web-sdk"
import Connection from "@open-ayame/ayame-web-sdk/dist/connection"
import { VideoCodecOption } from "@open-ayame/ayame-web-sdk/dist/connection/options"
import SkywayPeer, { SfuRoom } from "skyway-js"

// https://github.com/OpenAyame/ayame-web-sdk/blob/2022.1.0/src/connection/base.ts#L240-L251
type AddStreamEvent = RTCTrackEvent & {
  stream: MediaStream
}
declare var AYAME_SIGNALING_KEY: string
declare var SW_WSURL: string
declare var SKYWAY_APIKEY: string
declare var SKYWAY_DEBUG_LEVEL: 0 | 1 | 2 | 3
declare var SENDER_TOKEN: string

const remoteVideo = document.querySelector<HTMLVideoElement>("#remote_video")
const localVideo = document.querySelector<HTMLVideoElement>("#local_video")
const roomIdInput = document.querySelector<HTMLInputElement>("#room_id")
const videoCodecInput =
  document.querySelector<HTMLSelectElement>("#video_codec")
const connectMomoBtn =
  document.querySelector<HTMLInputElement>("#connect_momo_btn")
const disconnectMomoBtn = document.querySelector<HTMLInputElement>(
  "#disconnect_momo_btn"
)
const connectReceiverBtn = document.querySelector<HTMLInputElement>(
  "#connect_receiver_btn"
)
const disconnectReceiverBtn = document.querySelector<HTMLInputElement>(
  "#disconnect_receiver_btn"
)
const connectCameraBtn = document.querySelector<HTMLInputElement>(
  "#connect_camera_btn"
)
const disconnectCameraBtn = document.querySelector<HTMLInputElement>(
  "#disconnect_camera_btn"
)

if (
  remoteVideo === null ||
  localVideo === null ||
  roomIdInput === null ||
  videoCodecInput === null ||
  disconnectMomoBtn === null ||
  connectMomoBtn === null ||
  connectReceiverBtn === null ||
  connectCameraBtn === null ||
  disconnectCameraBtn === null ||
  disconnectReceiverBtn === null
) {
  throw new Error("failed to initialize")
}
function createRandomString(num: number) {
  const S = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
  return Array.from(crypto.getRandomValues(new Uint8Array(num)))
    .map((n) => S[n % S.length])
    .join("")
}
// https://github.com/OpenAyame/ayame-web-sdk/blob/2022.1.0/src/connection/options.ts#L28
const isSupportedVideoCodec = (codec: string): codec is VideoCodecOption =>
  codec === "VP8" ||
  codec === "VP9" ||
  codec === "AV1" ||
  codec === "H264" ||
  codec === "H265"

remoteVideo.controls = true
localVideo.controls = true

const signalingUrl = "wss://ayame-labo.shiguredo.jp/signaling"

const options = Ayame.defaultOptions
options.signalingKey = AYAME_SIGNALING_KEY

let conn: Connection | null = null
disconnectMomoBtn.addEventListener("click", () => {
  conn?.disconnect()
})
let videoCodec: VideoCodecOption | null = null
videoCodecInput.addEventListener("change", (e) => {
  if (!(e.currentTarget instanceof HTMLSelectElement)) {
    return
  }
  const value = e.currentTarget.value
  if (value === "none") {
    videoCodec = null
    return
  }
  if (!isSupportedVideoCodec(value)) {
    return
  }
  videoCodec = value
})
let momoStream: MediaStream | null
connectMomoBtn.addEventListener("click", async () => {
  if (videoCodec) {
    options.video.codec = videoCodec
  }
  console.log(options.video.codec)
  conn = Ayame.connection(signalingUrl, "ruu413@ayame-momo", options, true)
  await conn.connect(null)
  conn.on("open", ({ authzMetadata }: { authzMetadata: any }) =>
    console.log(authzMetadata)
  )
  conn.on("disconnect", () => {
    momoStream = null
    stopVideo(remoteVideo)
  })
  conn.on("addstream", (e: AddStreamEvent) => {
    debugger
    console.log("momo", e.stream.getVideoTracks()[0].getSettings())

    console.log("addstream")
    momoStream = e.stream
    playVideo(remoteVideo, e.stream)
  })
})

function playVideo(element: HTMLVideoElement, stream: MediaStream) {
  element.srcObject = stream
  element.play()
  element.volume = 0
}
function stopVideo(element: HTMLVideoElement) {
  element.pause()
  if ("srcObject" in element) {
    element.srcObject = null
  } else {
    element.src = ""
  }
}
const webSocket = new WebSocket(SW_WSURL)

type Room = {
  room_id: string
  stream: MediaStream
  skyway_room_id: string
  skyway_room: SfuRoom
}
let room: Room | null = null
const skywayPeer = new SkywayPeer({
  key: SKYWAY_APIKEY,
  debug: SKYWAY_DEBUG_LEVEL,
})
function connectRoom(roomId: string, stream: MediaStream) {
  const skywayRoomId = roomId + createRandomString(16)
  const skywayRoom = skywayPeer.joinRoom<SfuRoom>(skywayRoomId, {
    mode: "sfu",
    stream,
  })
  const room: Room = {
    room_id: roomId,
    stream,
    skyway_room_id: skywayRoomId,
    skyway_room: skywayRoom,
  }
  skywayRoom.on("open", () => {
    console.log("connect_sender", room)
    webSocket.send(
      JSON.stringify({
        msg_type: "connect_sender",
        peer_id: skywayPeer.id,
        room_id: room["room_id"],
        skyway_room_id: room["skyway_room_id"],
        sender_token: SENDER_TOKEN,
      })
    )
  })
  return room
}

connectReceiverBtn.addEventListener("click", () => {
  let stream = null
  const selector = document.querySelector<HTMLSelectElement>("#select_source")
  if (selector === null) {
    return
  }
  const selectedIdx = selector.selectedIndex
  const selected = selector.options[selectedIdx].value
  if (selected == "momo") {
    stream = momoStream
  } else {
    stream = cameraStream
  }
  if (stream === null) {
    return
  }
  room = connectRoom(roomIdInput.value, stream)
})

disconnectReceiverBtn.addEventListener("click", () => {
  if (room === null) {
    return
  }
  room.skyway_room.close()
  console.log("disconnect ", room)
  webSocket.send(
    JSON.stringify({
      msg_type: "exit_room",
      room_id: room.room_id,
    })
  )
  room = null
})

let cameraStream: MediaStream | null = null
connectCameraBtn.addEventListener("click", async () => {
  const cameraFramerateInput =
    document.querySelector<HTMLInputElement>("#camera_framerate")
  const cameraWidthInput =
    document.querySelector<HTMLInputElement>("#camera_width")
  const cameraHeightInput =
    document.querySelector<HTMLInputElement>("#camera_height")
  if (
    cameraFramerateInput === null ||
    cameraWidthInput === null ||
    cameraHeightInput === null
  ) {
    return
  }
  cameraStream = await navigator.mediaDevices.getUserMedia({
    video: {
      frameRate: {
        ideal: Number(cameraFramerateInput.value),
      },
      width: {
        ideal: Number(cameraWidthInput.value),
      },
      height: {
        ideal: Number(cameraHeightInput.value),
      },
      facingMode: { ideal: "environment" }, // リアカメラを利用する
    },
    audio: false,
  })
  console.log("camera", cameraStream.getVideoTracks()[0].getSettings())

  playVideo(localVideo, cameraStream)
})
disconnectCameraBtn.addEventListener("click", () => {
  if (cameraStream) {
    cameraStream.getTracks().forEach((track) => track.stop())
    cameraStream = null
  }
  stopVideo(localVideo)
})

webSocket.onmessage = (event) => {
  const message = JSON.parse(event.data)
  console.log(message)
  const msg_type = message["msg_type"]
  if (msg_type == "request_reconnect_sender") {
    const oldRoom = room
    if (oldRoom === null) {
      return
    }
    setTimeout(() => {
      oldRoom.skyway_room.close()
    }, 30000)
    // 前の通信残すようにしてるけどあまり効果ない?
    room = connectRoom(oldRoom.room_id, oldRoom.stream)
  }
}
