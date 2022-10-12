# WebRTCで1->複数に配信する
# 1人がsender

import asyncio
import websockets
import json
import ssl
import os

ADDRESS = "0.0.0.0"
PORT = 8081
MAX_CONNECT_NUM = 980
# CERT = "C://Users/asika/OneDrive/ドキュメント/webRTC/vscode_live_server.cert.pem"
# KEY = "C://Users/asika/OneDrive/ドキュメント/webRTC/vscode_live_server.key.pem"


connection_num = 0
connections = []
if "SENDER_TOKEN" in os.environ:
    sender_token = os.environ["SENDER_TOKEN"]  # ["127.0.0.1"]
else:
    sender_token = None
print(sender_token)

rooms = {}
# keyが部屋id, valueが{"sender_socket", "peer_id", "connections"}
print("a")
# 後で通信している個人が本物か見分けるのもじっそうしないといけないきがする


# offerer == receiver
# answerer == sender

# idで判別して複数動画同時配信したい
lock = asyncio.Lock()


async def handler(websocket, path):
    global connection_num, connections, rooms, lock
    remote_address = websocket.remote_address
    print(remote_address)
    async with lock:
        connections.append(websocket)
    while True:
        # 受信
        try:
            received_packet = await websocket.recv()
        except:
            break
        dictionary = json.loads(received_packet)
        # msg_type, room_idで構成
        # connect_senderの時のみ+peer_id
        promises = []
        msg_type = dictionary["msg_type"]
        room_id = dictionary["room_id"]

        if room_id not in rooms:
            room = {
                "sender_socket": None,
                "peer_id": None,
                "skyway_room_id": None,
                "connections": [websocket],
                "connect_num": 0,
                # ルームの累積接続数が1000行くと通信が弾かれるのでその前にルームを切り替え
            }
            async with lock:
                rooms[room_id] = room
        else:
            room = rooms[room_id]
            if websocket not in room["connections"]:
                async with lock:
                    room["connections"].append(websocket)
        # 現在の通信のwebsocketが入ったroom_idのroomが存在することを保証

        if msg_type == "connect_sender":
            if sender_token is None or sender_token == dictionary["sender_token"]:
                async with lock:  # if room["sender_socket"] is None:
                    print("sender_connect")
                    room["sender_socket"] = websocket
                    room["skyway_room_id"] = dictionary["skyway_room_id"]
                    room["connect_num"] = 0
                    # senderは上書き
                    room["peer_id"] = dictionary["peer_id"]
                    for connection in room["connections"]:
                        if connection is websocket:
                            continue
                        print("send")
                        promise = connection.send(
                            json.dumps(
                                {
                                    "msg_type": "connect_receiver",
                                    "room_id": room_id,
                                    "skyway_room_id": room["skyway_room_id"],
                                    "peer_id": room["peer_id"],
                                }
                            )
                        )
                        room["connect_num"] += 1
                        promises.append(promise)
        elif msg_type == "connect_receiver":
            print("connect_receiver")
            if room["sender_socket"] is not None:
                print("send")
                if room["connect_num"] < MAX_CONNECT_NUM:
                    promise = websocket.send(
                        json.dumps(
                            {
                                "msg_type": "connect_receiver",
                                "room_id": room_id,
                                "skyway_room_id": room["skyway_room_id"],
                                "peer_id": room["peer_id"],
                            }
                        )
                    )

                    async with lock:
                        room["connect_num"] += 1
                    promises.append(promise)
                else:
                    # ルームの累積接続数が溢れそうだったら新しい部屋をsenderに作ってもらう
                    promise = room["sender_socket"].send(
                        json.dumps(
                            {"msg_type": "request_reconnect_sender", "room_id": room_id}
                        )
                    )
                    promises.append(promise)
                    async with lock:
                        room["connect_num"] = 0
        elif msg_type == "exit_room":
            if websocket in room["connections"]:
                async with lock:
                    room["connections"].remove(websocket)
            if room["sender_socket"] is websocket:
                async with lock:
                    room["sender_socket"] = None
                    room["peer_id"] = None
                    room["skyway_room_id"] = None
                    room["connect_num"] = 0
        print("{}: {}".format(path, dictionary))
        for p in promises:
            try:
                a = await p
            except:
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
async def main():
    async with websockets.serve(handler, ADDRESS, PORT):  # , ssl=ssl_context)
        await asyncio.Future()  # run forever


if __name__ == "__main__":
    asyncio.run(main())
