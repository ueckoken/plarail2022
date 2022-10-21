const Ayame = require("@open-ayame/ayame-web-sdk")
const remoteVideo = document.getElementById("remote_video")
const localVideo = document.getElementById("local_video")
const roomIdInput = document.getElementById("room_id")
async function getDeviceStream(option) {
  console.log("wrap navigator.getUserMadia with Promise")
  return new Promise(function (resolve, reject) {
    navigator.getUserMedia(option, resolve, reject)
  })
}
function createRandomString(num) {
  const S = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
  return Array.from(crypto.getRandomValues(new Uint8Array(num)))
    .map((n) => S[n % S.length])
    .join("")
}
remoteVideo.controls = true
localVideo.controls = true

const signalingUrl = "wss://ayame-labo.shiguredo.jp/signaling"

const options = Ayame.defaultOptions
options.signalingKey = AYAME_SIGNALING_KEY

let conn = null
window.disconnectMomo = () => {
  if (conn) {
    conn.disconnect()
  }
}
let videoCodec = null
window.onChangeVideoCodec = () => {
  videoCodec = document.getElementById("video_codec").value
  if (videoCodec == "none") {
    videoCodec = null
  }
}
let momoStream = null
window.connectMomo = async () => {
  options.video.codec = videoCodec
  console.log(options.video.codec)
  conn = Ayame.connection(signalingUrl, "ruu413@ayame-momo", options, true)
  await conn.connect(null)
  conn.on("open", ({ authzMetadata }) => console.log(authzMetadata))
  conn.on("disconnect", (e) => {
    momoStream = null
    stopVideo(remoteVideo)
  })
  conn.on("addstream", (e) => {
    console.log("momo", e.stream.getVideoTracks()[0].getSettings())

    console.log("addstream")
    momoStream = e.stream
    playVideo(remoteVideo, e.stream)
  })
}

function playVideo(element, stream) {
  if ("srcObject" in element) {
    element.srcObject = stream
  } else {
    element.src = window.URL.createObjectURL(stream)
  }
  element.play()
  element.volume = 0
}
function stopVideo(element) {
  element.pause()
  if ("srcObject" in element) {
    element.srcObject = null
  } else {
    element.src = null
  }
}
const SkywayPeer = require("skyway-js")
const webSocket = new WebSocket(SW_WSURL)

let room = {}
const skywayPeer = new SkywayPeer({
  key: SKYWAY_APIKEY,
  debug: SKYWAY_DEBUG_LEVEL,
})
function connectRoom(roomId, stream) {
  let room = {}
  room["room_id"] = roomId
  room["stream"] = stream
  room["skyway_room_id"] = room["room_id"] + createRandomString(16)
  room["skyway_room"] = skywayPeer.joinRoom(room["skyway_room_id"], {
    mode: "sfu",
    stream: room["stream"],
  })
  room["skyway_room"].on("open", () => {
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

window.connectReceiver = () => {
  let stream = null
  const selector = document.getElementById("select_source")
  const selectedIdx = selector.selectedIndex
  const selected = selector.options[selectedIdx].value
  if (selected == "momo") {
    stream = momoStream
  } else {
    stream = cameraStream
  }
  room = connectRoom(roomIdInput.value, stream)
}

window.disconnectReceiver = () => {
  room["skyway_room"].close()
  console.log("disconnect ", room)
  webSocket.send(
    JSON.stringify({
      msg_type: "exit_room",
      room_id: room["room_id"],
    })
  )
  room = {}
}

let cameraStream = null
window.connectCamera = async () => {
  cameraStream = await getDeviceStream({
    video: {
      frameRate: {
        ideal: Number(document.getElementById("camera_framerate").value),
      },
      width: {
        ideal: Number(document.getElementById("camera_width").value),
      },
      height: {
        ideal: Number(document.getElementById("camera_height").value),
      },
      facingMode: { ideal: "environment" }, // リアカメラを利用する
    },
    audio: false,
  })
  console.log("camera", cameraStream.getVideoTracks()[0].getSettings())

  playVideo(localVideo, cameraStream)
}
window.disconnectCamera = () => {
  if (cameraStream) {
    cameraStream.getTracks().forEach((track) => track.stop())
    cameraStream = null
  }
  stopVideo(localVideo)
}

webSocket.onmessage = (event) => {
  const message = JSON.parse(event.data)
  console.log(message)
  const msg_type = message["msg_type"]
  if (msg_type == "request_reconnect_sender") {
    const oldRoom = room
    setTimeout(() => {
      oldRoom["skyway_room"].close()
    }, 30000)
    // 前の通信残すようにしてるけどあまり効果ない?
    room = connectRoom(room["room_id"], room["stream"])
  }
}
