# 自動運転プログラム
## 使い方
以下、auto_operation/system_on_raspi の中を見る

### シミュレーションモードの設定
- main.py の中に、`operation.state.communication.setup(simulationMode=True)`と書いてある部分がある。simulationModeを`True` にすると、ESP32が接続されていなくても動く

### Arduino のポート設定
- Communication.py の、`self.arduino = serial.Serial("COM8", 9600)` と書いてある部分でポートを指定する。simulationModeがTrueの場合とFalseの場合で2箇所あるので注意

### ESP32 のポート設定
- simulationMode = False にした場合には、ESP32も指定する
-  Communication.py の、`self.esp32Map[0] = serial.Serial("/dev/cu.ESP32-ESP32SPP", 115200)` と書いてある部分に各車両がつながっているポートを指定する

### 線路形状の確認
- auto_operation/system_on_raspi/State.py の中に線路形状と列車の初期位置が書き込まれている。必要に応じて修正する

### 実行
```
python3 main.py
```

## 中身
### child_ship/child_ship.ino (谷口)
  - 子艦の Arduino Nano に書き込む、サーボとセンサの制御プログラム
  - シリアル通信で junctionId を受け取り、ポイントを切り替える。
  - CdSが車両の通過を検知したとき、通過したCdSの sensorId を送る。

### mother_ship/mother_ship.ino (谷口)
  - 母艦の Arduino Mega に書き込むプログラム
  - シリアル通信で junctionId を受け取り、ポイントを切り替える。
  - CdSが車両の通過を検知したとき、通過したCdSの sensorId を送る。

### central_controller
自動運転システム

- templates
  - 運転体験ウェブページのCSSが入っている

- static
  - 運転体験ウェブページのHTMLが入っている

- Communication.py (三浦)
  - ESP32およびArduinoとの通信プログラム。
  - ESP32へinputを送り、ESP32からホールセンサの信号を受け取る。
  - Arduinoとの通信。切り替えるポイントの junctionId を送り、検知された sensorId を受け取る。
  - simulationMode = true にすると、ESP32の実機が無くてもシミュレーションできる。
  - **環境によってCOMポート番号の書き換えなどが必要**
- FlaskWebServer.py (松田)
  - 列車の現在地等を表示するウェブページを配信するウェブサーバー
- main.py (松田)
  - コマンドラインから実行する。自動運転システム(Operation)の稼働と、WEBサーバとしてGUIへの情報配信(Server)を行う
- Server.py (松田)
  - GUIに情報を送る
- Operation.py (松田)
  - 自動運転システムの最上位プログラム
  - Communication から受け取ったホールセンサ信号とセンサ信号(sensorId)を基にシミュレーションを更新。
  - 各車両の速度指令と、切り替えるポイントを計算し、内部モジュールのCommunicationを通してハードウェアに送信。
- その他の.pyファイル
  - 自動運転に用いられるモジュール。詳細は[五月祭2022列車制御システム構成図](https://docs.google.com/presentation/d/1fT75t7o4gi9V8cbwn1BwtUComzHEs0q3-pRfCk-om7c/edit?usp=sharing)を参照。
  - **State.pyは、線路形状に応じて書き換える**
  - **DiaPlanner.pyは、最初に列車を置いた場所に応じて書き換える**
