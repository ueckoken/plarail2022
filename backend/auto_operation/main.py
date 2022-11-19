import threading
import time

import pydantic
from flask import Flask, render_template
from flask_cors import CORS
from flask_socketio import SocketIO

from Components import Junction, Point, Section, Stop
from Connection import Connection
from Operation import Operation


class Conf(pydantic.BaseSettings):
    esp_eye_endpoint: str
    secret_key: str
    auto_operation_server_address: str
    external_server_address: str
    simulation_mode: bool = False


try:
    conf = Conf()
except pydantic.ValidationError as e:
    print(e.json)
    raise e

# 自動運転システムの初期化
operation = Operation()
operation.state.communication.setup(
    simulationMode=conf.simulation_mode,
    connection=Connection(
        autoOperationServerAddress=conf.auto_operation_server_address,
        externalServerAddress=conf.external_server_address,
    ),
)
operation.ato.setEnabled(1, True)

# Flaskウェブサーバの初期化
app = Flask(__name__)
app.config["SECRET_KEY"] = conf.secret_key
socketio = SocketIO(
    app,
    cors_allowed_origins="*",  # ビジュアライザからのアクセスを許可する
    async_mode="gevent",  # WindowsでCtrl-Cが効かない問題への対処
)
CORS(app)  # ビジュアライザからのアクセスを許可する


# 自動運転システムが0.1secおきに実行する作業
def operation_loop():
    while True:
        operation.update()
        print(
            f"[main.operation_loop] t0.section: {operation.state.getTrainById(0).currentSection.id}, t0.mil: {operation.state.getTrainById(0).mileage:.2f}, t0.spd: {operation.state.getTrainById(0).targetSpeed:.2f}, t1.section: {operation.state.getTrainById(1).currentSection.id}, t1.mil: {operation.state.getTrainById(1).mileage:.2f}, t1.spd: {operation.state.getTrainById(1).targetSpeed:.2f}"
        )
        print(
            f"[main.operation_loop] t2.section: {operation.state.getTrainById(2).currentSection.id}, t2.mil: {operation.state.getTrainById(2).mileage:.2f}, t2.spd: {operation.state.getTrainById(2).targetSpeed:.2f}, t3.section: {operation.state.getTrainById(3).currentSection.id}, t3.mil: {operation.state.getTrainById(3).mileage:.2f}, t3.spd: {operation.state.getTrainById(3).targetSpeed:.2f}"
        )
        time.sleep(0.1)


# ブラウザにwebsocketで0.1secおきに信号を送る関数
def send_signal_to_browser():
    connection = operation.state.communication.connection

    if connection is None:
        return

    while True:
        socketio.sleep(0.1)
        # 閉塞を送る。
        blocks: dict[Section.SectionId, bool] = {}
        for section in operation.state.sectionList:
            blocks[section.id] = False
        for train in operation.state.trainList:
            blocks[train.currentSection.id] = True
        # どこのストップレールを上げるべきか送る。
        stops: dict[Stop.StopId, bool] = {}
        for section in operation.state.sectionList:
            stopId = operation.state.sectionIdToStopId[section.id]
            stops[stopId] = False
        for train in operation.state.trainList:
            if train.stopPoint:  # 列車には停止すべき点が存在しない場合もある(ATSを無効化した場合など)。停止点を持っている場合はsectionIDを送る
                sectionId = train.stopPoint.section.id
                stopId = operation.state.sectionIdToStopId[sectionId]
                stops[stopId] = True
        # ポイント切り替えを送る。分岐側に向ける=true, 直進側に向ける=false
        points: dict[Point.PointId, bool] = {}
        for junction in operation.state.junctionList:
            if junction.pointId is not None:  # サーボがついてるjunctionについて
                pointId = junction.pointId
                if junction.inSectionCurve:  # in2->out1という分岐器の場合
                    points[pointId] = (
                        True if (junction.inServoState == Junction.ServoState.Curve) else False
                    )
                if junction.outSectionCurve:  # in1->out2という分岐器の場合
                    points[pointId] = (
                        True if (junction.outServoState == Junction.ServoState.Curve) else False
                    )

        # 閉塞を送る
        for blockId, state in blocks.items():
            connection.sendBlock(
                blockId=blockId, state="BLOCKSTATE_CLOSE" if state else "BLOCKSTATE_OPEN"
            )
        # ストップレールを送る
        for stopId, state in stops.items():
            connection.sendStop(
                stationId=stopId, state="POINTSTATE_ON" if state else "POINTSTATE_OFF"
            )
        # ポイント切り替えを送る
        for pointId, state in points.items():
            connection.sendPoint(
                pointId=pointId, state="POINTSTATE_ON" if state else "POINTSTATE_OFF"
            )

        # websocketで送信
        socketio.emit(
            "signal_taiken",
            {"signal": "R", "distance": 0, "blocks": blocks, "stops": stops},
        )


# ブラウザからwebsocketで速度指令を受け取って、体験車の速度を変更する関数
@socketio.on("speed")
def receive_speed_from_browser(json):
    # 体験車のID:1, 受け取ったspeed: json['speed']で速度制御
    operation.ato.setSpeed(1, json["speed"])


@app.route("/")
def index():
    # ブラウザにwebページのデータを返す
    return render_template(
        "index.html",
        esp_eye_ip_addr=conf.esp_eye_endpoint,
        max_speed=Operation.MAXSPEED,
    )


if __name__ == "__main__":
    operationThread = threading.Thread(target=operation_loop, daemon=True)
    operationThread.start()  # 自動運転のオペレーションを開始
    signalThread = threading.Thread(target=send_signal_to_browser, daemon=True)
    signalThread.start()  # ブラウザへデータを送信するタスクの開始
    socketio.run(app, host="0.0.0.0", port=50050)  # Flaskソケットを起動
