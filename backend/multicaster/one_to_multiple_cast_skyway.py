# WebRTCで1->複数に配信する
# 1人がsender

import asyncio
from websockets.server import serve
from websockets.exceptions import ConnectionClosedError
from websockets.server import WebSocketServerProtocol
import json
import ssl
import os
from typing import Literal, Optional, TypedDict, Union, cast

ADDRESS = "0.0.0.0"
PORT = 8081
MAX_CONNECT_NUM = 980
# CERT = "C://Users/asika/OneDrive/ドキュメント/webRTC/vscode_live_server.cert.pem"
# KEY = "C://Users/asika/OneDrive/ドキュメント/webRTC/vscode_live_server.key.pem"
SENDER_TOKEN = os.environ.get("SENDER_TOKEN")  # ["127.0.0.1"]
print(SENDER_TOKEN)


connections = []


class ConnectSenderMessage(TypedDict):
    msg_type: Literal["connect_sender"]
    room_id: str
    peer_id: str
    skyway_room_id: str
    sender_token: str


class ConnectReceiverMessage(TypedDict):
    msg_type: Literal["connect_receiver"]
    room_id: str


class ExitRoomMessage(TypedDict):
    msg_type: Literal["exit_room"]
    room_id: str


class RequestReconnectSenderMessage(TypedDict):
    msg_type: Literal["request_reconnect_sender"]
    room_id: str


class ConnectReceiverReplyMessage(TypedDict):
    msg_type: Literal["connect_receiver"]
    room_id: str
    peer_id: str
    skyway_room_id: str


Message = Union[
    ConnectSenderMessage,
    ConnectReceiverMessage,
    ExitRoomMessage,
    RequestReconnectSenderMessage,
    ConnectReceiverReplyMessage,
]


class Room(TypedDict):
    sender_socket: Optional[WebSocketServerProtocol]
    peer_id: Optional[str]
    skyway_room_id: Optional[str]
    connections: list[WebSocketServerProtocol]
    connect_num: int


RemoteAddress = Optional[tuple[str, int]]
"""ref: https://github.com/aaugustin/websockets/blob/10.3/src/websockets/legacy/protocol.py#L392-L399
"""

rooms: dict[str, Room] = {}
# keyが部屋id, valueが{"sender_socket", "peer_id", "connections"}
print("a")
# 後で通信している個人が本物か見分けるのもじっそうしないといけないきがする


# offerer == receiver
# answerer == sender

# idで判別して複数動画同時配信したい
lock = asyncio.Lock()


async def handler(websocket: WebSocketServerProtocol, path: str) -> None:
    global connections, rooms, lock
    remote_address = cast(RemoteAddress, websocket.remote_address)
    print(remote_address)
    async with lock:
        connections.append(websocket)

    try:
        async for raw_message in websocket:  # 受信
            message: Message = json.loads(raw_message)
            # msg_type, room_idで構成
            # connect_senderの時のみ+peer_id
            promises = []
            room_id = message["room_id"]

            if room_id not in rooms:
                room = Room(
                    sender_socket=None,
                    peer_id=None,
                    skyway_room_id=None,
                    connections=[websocket],
                    connect_num=0,
                    # ルームの累積接続数が1000行くと通信が弾かれるのでその前にルームを切り替え
                )
                async with lock:
                    rooms[room_id] = room
            else:
                room = rooms[room_id]
                if websocket not in room["connections"]:
                    async with lock:
                        room["connections"].append(websocket)
            # 現在の通信のwebsocketが入ったroom_idのroomが存在することを保証

            if message["msg_type"] == "connect_sender":
                if room["skyway_room_id"] is None or room["peer_id"] is None:
                    continue
                if SENDER_TOKEN is None or SENDER_TOKEN == message["sender_token"]:
                    async with lock:  # if room["sender_socket"] is None:
                        print("sender_connect")
                        room["sender_socket"] = websocket
                        room["skyway_room_id"] = message["skyway_room_id"]
                        room["connect_num"] = 0
                        # senderは上書き
                        room["peer_id"] = message["peer_id"]
                        for connection in room["connections"]:
                            if connection is websocket:
                                continue
                            print("send")
                            promise = connection.send(
                                json.dumps(
                                    ConnectReceiverReplyMessage(
                                        msg_type="connect_receiver",
                                        room_id=room_id,
                                        skyway_room_id=room["skyway_room_id"],
                                        peer_id=room["peer_id"],
                                    )
                                )
                            )
                            room["connect_num"] += 1
                            promises.append(promise)
            elif message["msg_type"] == "connect_receiver":
                print("connect_receiver")
                if room["sender_socket"] is not None:
                    if room["skyway_room_id"] is None or room["peer_id"] is None:
                        continue
                    print("send")
                    if room["connect_num"] < MAX_CONNECT_NUM:
                        promise = websocket.send(
                            json.dumps(
                                ConnectReceiverReplyMessage(
                                    msg_type="connect_receiver",
                                    room_id=room_id,
                                    skyway_room_id=room["skyway_room_id"],
                                    peer_id=room["peer_id"],
                                )
                            )
                        )

                        async with lock:
                            room["connect_num"] += 1
                        promises.append(promise)
                    else:
                        # ルームの累積接続数が溢れそうだったら新しい部屋をsenderに作ってもらう
                        promise = room["sender_socket"].send(
                            json.dumps(
                                RequestReconnectSenderMessage(
                                    msg_type="request_reconnect_sender", room_id=room_id
                                )
                            )
                        )
                        promises.append(promise)
                        async with lock:
                            room["connect_num"] = 0
            elif message["msg_type"] == "exit_room":
                if websocket in room["connections"]:
                    async with lock:
                        room["connections"].remove(websocket)
                if room["sender_socket"] is websocket:
                    async with lock:
                        room["sender_socket"] = None
                        room["peer_id"] = None
                        room["skyway_room_id"] = None
                        room["connect_num"] = 0
            print("{}: {}".format(path, message))
            for p in promises:
                try:
                    await p
                except:
                    pass
    except ConnectionClosedError:
        pass

    # 接続が切れたらその接続を削除
    for room_id, room in rooms.items():
        print(room)
        if websocket in room["connections"]:
            async with lock:
                room["connections"].remove(websocket)
        if room["sender_socket"] is websocket:
            async with lock:
                room["sender_socket"] = None
                room["peer_id"] = None

    async with lock:
        rooms = {
            room_id: room
            for room_id, room in rooms.items()
            if len(room["connections"]) != 0
        }
    # ルームに人がいなくなったら削除

    async with lock:
        connections.remove(websocket)


# ssl_context = ssl.SSLContext(ssl.PROTOCOL_TLS_SERVER)
# ssl_context.load_cert_chain(CERT, keyfile=KEY)

# サーバー立ち上げ
async def main() -> None:
    async with serve(handler, ADDRESS, PORT):  # , ssl=ssl_context)
        await asyncio.Future()  # run forever


if __name__ == "__main__":
    asyncio.run(main())
