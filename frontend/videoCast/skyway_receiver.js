const SkywayPeer = require("skyway-js")

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
  if ("srcObject" in element) {
    element.srcObject = null
  } else {
    element.src = null
  }
}
const webSocket = new WebSocket(SW_WSURL)
let sendFuncs = []
const rooms = {}
const skywayPeer = new SkywayPeer({
  key: SKYWAY_APIKEY,
  debug: SKYWAY_DEBUG_LEVEL,
})
let isConnectWebSocket = false
let isConnectSkywayPeer = false
webSocket.onopen = (event) => {
  isConnectWebSocket = true
  if (!isConnectSkywayPeer) {
    return
  }
  for (const func of sendFuncs) {
    func()
  }
  sendFuncs = []
}

skywayPeer.on("open", () => {
  isConnectSkywayPeer = true
  if (!isConnectWebSocket) {
    return
  }
  for (const func of sendFuncs) {
    func()
  }
  sendFuncs = []
})
webSocket.onmessage = (event) => {
  const message = JSON.parse(event.data)
  console.log(message)
  const peerId = message["peer_id"]
  const roomId = message["room_id"]
  const skywayRoomId = message["skyway_room_id"]
  const room = rooms[roomId]
  console.log("joinroom")
  room["skyway_room"] = skywayPeer.joinRoom(skywayRoomId, {
    mode: "sfu",
  })
  room["skyway_room"].on("stream", (stream) => {
    const streamPeerId = stream.peerId
    console.log("on stream")
    if (streamPeerId == peerId) {
      playVideo(room["video_element"], stream)
    }
  })
}

window.openVideoConnection = (id, roomId) => {
  rooms[roomId] = {
    video_element: document.getElementById(id),
    peer: null,
  }
  const sendFunc = () => {
    webSocket.send(
      JSON.stringify({
        msg_type: "connect_receiver",
        room_id: roomId,
      })
    )
  }
  if (webSocket.readyState == webSocket.OPEN) {
    sendFunc()
  } else {
    sendFuncs.push(sendFunc)
  }
}
window.closeVideoConnection = (roomId) => {
  const room = rooms[roomId]
  stopVideo(room["video_element"])
  room["skyway_room"].close()
  delete rooms[roomId]

  webSocket.send(
    JSON.stringify({
      msg_type: "exit_room",
      room_id: roomId,
    })
  )
}
