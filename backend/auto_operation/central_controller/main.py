import atexit
from Operation import *
from flask import Flask, render_template, Response
from flask_socketio import SocketIO
from flask_cors import CORS
import threading

ESP_EYE_IP_ADDR = "esp32-1D31E4.mshome.net"

# 自動運転システムの初期化
operation = Operation()
operation.state.communication.setup(simulationMode=True)
operation.ato.setEnabled(1, True)

# Flaskウェブサーバの初期化
app = Flask(__name__)
app.config['SECRET_KEY'] = 'secret!'
socketio = SocketIO(
    app,
    cors_allowed_origins="*",  # ビジュアライザからのアクセスを許可する
    async_mode='gevent'  # WindowsでCtrl-Cが効かない問題への対処
)
CORS(app)  # ビジュアライザからのアクセスを許可する

# 自動運転システムが0.1secおきに実行する作業
def operation_loop():
    while True:
        operation.update()
        train_taiken = operation.state.getTrainById(1)  # ラズパイ体験車(id=1)を取得
        print(f"[main.operation_loop] t0.section: {operation.state.getTrainById(0).currentSection.id}, t0.mil: {operation.state.getTrainById(0).mileage:.2f}, t0.spd: {operation.state.getTrainById(0).targetSpeed:.2f}, t1.section: {operation.state.getTrainById(1).currentSection.id}, t1.mil: {operation.state.getTrainById(1).mileage:.2f}, t1.spd: {operation.state.getTrainById(1).targetSpeed:.2f}")
        time.sleep(0.1)

# ブラウザにwebsocketで0.1secおきに信号を送る関数
def send_signal_to_browser():
    while True:
        socketio.sleep(0.1)
        train_taiken = operation.state.getTrainById(1)  # ラズパイ体験車(id=1)を取得
        signal = operation.signalSystem.getSignal(train_taiken.currentSection.id, train_taiken.currentSection.targetJunction.getOutSection().id)  # 体験車から見た信号機を取得
        distance = operation.ato.getDistanceUntilStop(train_taiken)  # 停止位置までの距離を取得
        blocks = {}  # 閉塞を送る。区間0と3に列車がいるなら、{'s0': True, 's3': True} のような文字列
        for train in operation.state.trainList:
            blocks["s" + str(train.currentSection.id)] = True

        # websocketで送信
        socketio.emit('signal_taiken', {
            'signal': signal.value,
            'distance': int(distance),
            'blocks': blocks
        })

# ブラウザからwebsocketで速度指令を受け取って、体験車の速度を変更する関数
@socketio.on('speed')
def receive_speed_from_browser(json):
    # 体験車のID:1, 受け取ったspeed: json['speed']で速度制御
    operation.ato.setSpeed(1, json['speed'])


@app.route('/')
def index():
    # ブラウザにwebページのデータを返す
    return render_template(
        'index.html',
        esp_eye_ip_addr=ESP_EYE_IP_ADDR,
        max_speed=Operation.MAXSPEED,
    )


@app.before_first_request
def before_first_request():
    # ブラウザへデータを送信するタスクの開始
    socketio.start_background_task(target=send_signal_to_browser)


@atexit.register
def reset():
    state = operation.state
    state.communication.sendToggle(0, Junction.ServoState.Straight)
    state.communication.sendToggle(1, Junction.ServoState.Straight)
    state.communication.sendSpeed(0, 0)
    state.communication.sendSpeed(1, 0)
    print("successfully reset")


if __name__ == "__main__":
    thread1 = threading.Thread(target=operation_loop, daemon=True)
    thread1.start()  # 自動運転のオペレーションを開始
    socketio.run(app, host='0.0.0.0', port=50050)  # Flaskソケットを起動
