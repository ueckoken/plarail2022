// komabasai2019ブランチのarduino/stationのコードを改変して利用しています。
//サーボ1、センサー1
#include <VarSpeedServo.h>

//サーボ関係の定数、変数
// const int num_servo = 2;//サーボの数(ポイントの数)
const int num_servo = 1;

VarSpeedServo servo[num_servo];  //各サーボを入れる配列

const int servoSpeed =
    100;  // 1から255。サーボを回転させる速さ。これを大きくすると、サーボ稼働時にセンサ入力が送られてしまう

const int servo_angle_straight[num_servo] = {25};
const int servo_angle_curve[num_servo]    = {75};
const byte straight                       = 0;
const byte curve                          = 1;
byte servo_status[num_servo]              = {straight};

byte data = 0;  //受信データ格納用

void servo_change(
    byte servo_id,
    byte
        servo_state) {  // servoの向きを切り替える関数。引数は変えたいサーボのidと向き
  if (servo_state == straight) {
    servo[servo_id].write(servo_angle_straight[servo_id], servoSpeed, true);
    servo_status[servo_id] = straight;
  } else {
    servo[servo_id].write(servo_angle_curve[servo_id], servoSpeed, true);
    servo_status[servo_id] = curve;
  }
}

// CdS関係の変数、定数
const int num_sensor            = 2;         // CdSの個数
const int sensorPin[num_sensor] = {A0, A2};  // CdSセンサーの計測
int sensor_baseline[num_sensor] = {
    150,
    150};  // CdSセンサーの読み取り値がこれを下回ったら通過と判定する。起動時に決定する仕様に変更する必要あり。
unsigned long before_passing_time[num_sensor] = {0, 0};  //前回通過した時の時間
unsigned long time;  //現在の時間。millis()を受ける。
unsigned long time_for_passing =
    5000;  //通過に要する時間。前回の通過判定からこの時間だけは通過判定がなされない。
int value;

// CdSセンサーのデータを元に車両が来ているかいないか判定してPCにその情報を送る関数。引数は読み取りたいセンサーのid
void CdS_process(int sensor_id) {
  value = analogRead(sensorPin[sensor_id]);  // CdSセンサーで明るさを計測
  // Serial.print(sensor_id);
  // Serial.print(", ");
  // Serial.println(value);
  //絶対値バージョン
  if (value <
      sensor_baseline[sensor_id]) {  //基準を下回る明るさだったら通過と判定
    time = millis();
    if (before_passing_time[sensor_id] == 0 ||
        time - before_passing_time[sensor_id] >
            time_for_passing) {  //一度も車両が通過していないか前回通過時から一定時間経っていれば通過と判定。
      // Serial.print("pass: ");
      // Serial.println(sensor_id); //シリアルモニタを見てデバッグ等したい時用
      Serial.write((byte)sensor_id);  //実際に使う用
      before_passing_time[sensor_id] = time;
    }
  }
}

void setup() {
  Serial.begin(9600);
  servo[0].attach(3);
  for (int servo_id = 0; servo_id < num_servo; servo_id++) {
    servo[servo_id].write(servo_angle_straight[servo_id], servoSpeed, true);
    servo_status[servo_id] = straight;
  }
  for (int i = 0; i < num_sensor; i++) {
    value = analogRead(sensorPin[i]);
    sensor_baseline[i] =
        value /
        3;  //理論上はvalue *=
            // 8/35でいけると思いますが若干基準を緩くしてあります。状況に応じて調節。
  }
}

void loop() {
  //サーボなしでテスト
  while (Serial.available() >=
         2) {  //シリアルで受け取った信号をもとにサーボを動かす
    byte servo_id    = Serial.read();
    byte servo_state = Serial.read();
    servo_change(servo_id, servo_state);
  }
  for (int i = 0; i < num_sensor; i++) {
    CdS_process(i);  // CdSセンサーから車両の通過を検知した場合にはPCに送る
  }
}
