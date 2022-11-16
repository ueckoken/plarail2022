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

Hallsensor shinjuku_1 = Hallsensor(shinjuku_d1, 26);
Hallsensor shinjuku_2 = Hallsensor(shinjuku_d1, 27);

void loop() {
    shinjuku_1.detect();
    shinjuku_2.detect();
}
