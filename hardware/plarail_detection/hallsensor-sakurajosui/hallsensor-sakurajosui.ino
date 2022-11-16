#include <ArduinoJson.h>
#include <WiFi.h>
#include <HTTPClient.h>
#include <Hallsensor.hpp>

const char* ssid = "";
const char* password = "";

void setup() {
    Serial.begin(115200);
    WiFi.begin(ssid, password);
}

Hallsensor sakurajosui_1 = Hallsensor(sakurajosui_d1, 34);
Hallsensor sakurajosui_2 = Hallsensor(sakurajosui_d2, 35);
Hallsensor sakurajosui_3 = Hallsensor(sakurajosui_d3, 32);
Hallsensor sakurajosui_4 = Hallsensor(sakurajosui_d4, 33);
Hallsensor sakurajosui_5 = Hallsensor(sakurajosui_d5, 25);
Hallsensor sakurajosui_6 = Hallsensor(sakurajosui_d6, 26);

void loop() {
    sakurajosui_1.detect();
    sakurajosui_2.detect();
    sakurajosui_3.detect();
    sakurajosui_4.detect();
    sakurajosui_5.detect();
    sakurajosui_6.detect();
}
