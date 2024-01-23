import asyncio
import json
from typing import cast
from unittest import IsolatedAsyncioTestCase
from unittest.mock import patch
from websockets.server import serve, WebSocketServer
from websockets.client import connect, WebSocketClientProtocol
from one_to_multiple_cast_skyway import (
    handler,
    ConnectSenderMessage,
    ConnectReceiverMessage,
    RemoteAddress,
)

DUMMY_SENDER_TOKEN = "abcdefg"
WAIT_STATE_CHANGE_SEC = 0.5


class OneToMultipleCastSkywayTest(IsolatedAsyncioTestCase):
    server: WebSocketServer
    ws_uri: str

    async def asyncSetUp(self) -> None:
        self.server = await serve(handler, "127.0.0.1", 0)
        addr = cast(RemoteAddress, next(iter(self.server.sockets)).getsockname())
        if addr is None:
            raise Exception("cannot get server host and port")
        host, port = addr
        self.ws_uri = f"ws://{host}:{port}/"

    async def asyncTearDown(self) -> None:
        self.server.close()
        await self.server.wait_closed()

    @patch("one_to_multiple_cast_skyway.SENDER_TOKEN", DUMMY_SENDER_TOKEN)
    async def test_connect_sender_with_invalid_sender_token(self) -> None:
        from one_to_multiple_cast_skyway import rooms

        DUMMY_ROOM_ID = "A101"
        DUMMY_SKYWAY_ROOM_ID = "B101"
        DUMMY_PEER_ID = "you-write-code-slowly-because-you-don't-write-test"
        async with connect(self.ws_uri) as sender_socket:
            await sender_socket.send(
                json.dumps(
                    ConnectSenderMessage(
                        msg_type="connect_sender",
                        room_id=DUMMY_ROOM_ID,
                        skyway_room_id=DUMMY_SKYWAY_ROOM_ID,
                        peer_id=DUMMY_PEER_ID,
                        sender_token="invalid_token",
                    )
                )
            )
            await asyncio.sleep(WAIT_STATE_CHANGE_SEC)
            self.assertNotIn(DUMMY_ROOM_ID, rooms)
        await asyncio.sleep(WAIT_STATE_CHANGE_SEC)
        self.assertNotIn(DUMMY_ROOM_ID, rooms)

    @patch("one_to_multiple_cast_skyway.SENDER_TOKEN", DUMMY_SENDER_TOKEN)
    async def test_connect_sender_with_valid_sender_token(self) -> None:
        from one_to_multiple_cast_skyway import rooms

        DUMMY_ROOM_ID = "A102"
        DUMMY_SKYWAY_ROOM_ID = "B102"
        DUMMY_PEER_ID = "you-would-go-slow-if-you-don't-write-test"
        async with connect(self.ws_uri) as sender_socket:
            await sender_socket.send(
                json.dumps(
                    ConnectSenderMessage(
                        msg_type="connect_sender",
                        room_id=DUMMY_ROOM_ID,
                        skyway_room_id=DUMMY_SKYWAY_ROOM_ID,
                        peer_id=DUMMY_PEER_ID,
                        sender_token=DUMMY_SENDER_TOKEN,
                    )
                )
            )
            await asyncio.sleep(WAIT_STATE_CHANGE_SEC)
            self.assertIn(DUMMY_ROOM_ID, rooms)
            self.assertIsNotNone(rooms[DUMMY_ROOM_ID]["sender_socket"])
            self.assertEqual(len(rooms[DUMMY_ROOM_ID]["connections"]), 1)
            self.assertEqual(
                rooms[DUMMY_ROOM_ID]["cumulative_activated_connect_num"], 0
            )
            self.assertEqual(
                rooms[DUMMY_ROOM_ID]["skyway_room_id"], DUMMY_SKYWAY_ROOM_ID
            )
            self.assertEqual(rooms[DUMMY_ROOM_ID]["peer_id"], DUMMY_PEER_ID)
        await asyncio.sleep(WAIT_STATE_CHANGE_SEC)
        self.assertNotIn(DUMMY_ROOM_ID, rooms)

    @patch("one_to_multiple_cast_skyway.SENDER_TOKEN", DUMMY_SENDER_TOKEN)
    async def test_connect_receiver(self) -> None:
        from one_to_multiple_cast_skyway import rooms

        DUMMY_ROOM_ID = "A201"
        async with connect(self.ws_uri) as receiver_socket:
            await receiver_socket.send(
                json.dumps(
                    ConnectReceiverMessage(
                        msg_type="connect_receiver",
                        room_id=DUMMY_ROOM_ID,
                    )
                )
            )
            await asyncio.sleep(WAIT_STATE_CHANGE_SEC)
            self.assertIn(DUMMY_ROOM_ID, rooms)
            self.assertIsNone(rooms[DUMMY_ROOM_ID]["sender_socket"])
            self.assertEqual(len(rooms[DUMMY_ROOM_ID]["connections"]), 1)
            self.assertEqual(
                rooms[DUMMY_ROOM_ID]["cumulative_activated_connect_num"], 0
            )

        await asyncio.sleep(WAIT_STATE_CHANGE_SEC)
        self.assertNotIn(DUMMY_ROOM_ID, rooms)

    @patch("one_to_multiple_cast_skyway.SENDER_TOKEN", DUMMY_SENDER_TOKEN)
    async def test_connect_multiple_receiver(self) -> None:
        from one_to_multiple_cast_skyway import rooms

        DUMMY_ROOM_ID = "B202"
        CONNECTIONS_LEN = 10
        receiver_sockets: set[WebSocketClientProtocol] = set()
        for i in range(CONNECTIONS_LEN):
            receiver_socket = await connect(self.ws_uri)
            receiver_sockets.add(receiver_socket)
            await receiver_socket.send(
                json.dumps(
                    ConnectReceiverMessage(
                        msg_type="connect_receiver",
                        room_id=DUMMY_ROOM_ID,
                    )
                )
            )
            await asyncio.sleep(WAIT_STATE_CHANGE_SEC)
            self.assertIn(DUMMY_ROOM_ID, rooms)
            self.assertIsNone(rooms[DUMMY_ROOM_ID]["sender_socket"])
            connections_len_expected = i + 1
            self.assertEqual(
                len(rooms[DUMMY_ROOM_ID]["connections"]), connections_len_expected
            )
            self.assertEqual(
                rooms[DUMMY_ROOM_ID]["cumulative_activated_connect_num"], 0
            )

        self.assertEqual(len(rooms[DUMMY_ROOM_ID]["connections"]), CONNECTIONS_LEN)

        for i, receiver_socket in enumerate(receiver_sockets, start=1):
            await receiver_socket.close()
            await asyncio.sleep(WAIT_STATE_CHANGE_SEC)
            connections_len_expected = CONNECTIONS_LEN - i
            if connections_len_expected != 0:
                self.assertEqual(
                    len(rooms[DUMMY_ROOM_ID]["connections"]), connections_len_expected
                )
                self.assertEqual(
                    rooms[DUMMY_ROOM_ID]["cumulative_activated_connect_num"], 0
                )
            else:
                self.assertNotIn(DUMMY_ROOM_ID, rooms)
