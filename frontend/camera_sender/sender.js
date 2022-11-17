const canvas = document.getElementById("canvas")
// create canvas element
const tempCanvas = document.createElement("canvas")
const context = canvas.getContext("2d")
const esp_eye_ip = ESP_EYE_IP
const Peer = require("skyway-js")

function render() {
  const image = new Image()
  image.onload = function () {
    const context = canvas.getContext("2d")
    context.save()
    context.translate(80, 60)
    context.rotate((180 * Math.PI) / 180)
    context.drawImage(image, -(image.width / 2), -(image.height / 2))
    context.restore()
  }
  // ↓ MJPG-streamerの静止画像をロード
  image.src = "http://" + esp_eye_ip + "/mjpeg/1"
  image.crossOrigin = "anonymous"
  requestAnimationFrame(render)
}
render()

// キャプチャしたい canvas 要素を取得
let canvasElt = document.querySelector("canvas")

// ストリームの取得
let myStream = canvasElt.captureStream(25) // 25 FPS

//Peer作成
const peer = new Peer({
  key: "2eb379e0-0374-4e3c-9674-d7415b0b7f27",
  debug: 3,
})

//PeerID取得
peer.on("open", () => {
  document.getElementById("my-id").textContent = "my-id: " + peer.id
})

// ルームに入室する処理
document.getElementById("make-call").onclick = () => {
  const roomID = document.getElementById("room-id").value
  if (!peer.open) {
    return
  }

  const room = peer.joinRoom(roomID, {
    mode: "sfu",
    stream: myStream,
  })
}

// error 処理
peer.on("error", console.error)