export type ConnectMessage = {
  msg_type: "connect_receiver"
  room_id: string
}
export type ExitMessage = {
  msg_type: "exit_room"
  room_id: string
}
export type SendMessage = {
  skyway_room_id: string
  room_id: string
  peer_id: string
}
export type WebRTCMessage = ConnectMessage | ExitMessage | SendMessage
