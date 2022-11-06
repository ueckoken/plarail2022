const byte min_Serial_servo_id[3] = {0, 0, 1}; //各子艦のサーボのidの中で最小のものを書いた配列。最後の要素はservo_idの最大値+1(番兵)
const byte min_Serial_sensor_id[3] = {0, 1, 2}; //各子艦のセンサーのidの中で最小のものを書いた配列。最後の要素はsensor_idの最大値+1(番兵)

byte data = 0;//受信データ格納用

void to_child(byte servo_id, byte servo_state){ //母艦から子艦へのデータの送信
  //if (servo_id >= min_Serial_servo_id[3]) {
  //  Serial.print("error");
  //}
  //else if (servo_id >= min_Serial_servo_id[2]) {
  //  Serial3.write(servo_id-min_Serial_servo_id[2]);
  //  Serial3.write(servo_state);
  //}
  if (servo_id >= min_Serial_servo_id[1]) { //この行は元々else ifだったので直す時に注意
    Serial2.write(servo_id-min_Serial_servo_id[1]);
    Serial2.write(servo_state);
  }
  else {
    Serial1.write(servo_id-min_Serial_servo_id[0]);
    Serial1.write(servo_state);
  }
}

void to_pc(byte sensor_id, byte num){ //母艦からPCへのデータの送信
  Serial.write((byte)(sensor_id+min_Serial_sensor_id[num-1]));
}


void setup(){
  Serial.begin(9600);
  Serial1.begin(9600);
  Serial2.begin(9600);
  //Serial3.begin(9600);
}

void loop(){
  //シリアルで受け取った信号をもとに子艦に信号を受け流す。
  while(Serial.available() >= 2){
    byte servo_id = Serial.read();
    byte servo_state = Serial.read();
    to_child(servo_id, servo_state);
  }
  //シリアル1~3で受け取った信号をもとにPCに信号を受け流す。
  //どこかのセンサーがバグった時に最低限他のセンサーからの信号が読み取れるような実装になっています。
  //while(Serial1.available() > 0 || Serial2.available() > 0 || Serial3.available() > 0){
  while(Serial1.available() > 0 || Serial2.available() > 0){
    if (Serial1.available() > 0) {
      data = Serial1.read();
      to_pc(data, 1);
    }
    if (Serial2.available() > 0) {
      data = Serial2.read();
      to_pc(data, 2);
    }
    /*if (Serial3.available() > 0) {
      data = Serial3.read();
      to_pc(data, 3);
    }*/
  }
}
