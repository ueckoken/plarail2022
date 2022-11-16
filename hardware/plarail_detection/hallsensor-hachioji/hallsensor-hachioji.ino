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

Hallsensor hachioji_1 = Hallsensor(hachioji_d1, 26);
Hallsensor hachioji_2 = Hallsensor(hachioji_d2, 27);

void loop() {
    hachioji_1.detect();
    hachioji_2.detect();
}
