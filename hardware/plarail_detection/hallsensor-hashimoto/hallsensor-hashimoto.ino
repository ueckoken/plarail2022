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

Hallsensor hashimoto_1 = Hallsensor(hashimoto_d1, 26);
Hallsensor hashimoto_2 = Hallsensor(hashimoto_d2, 27);

void loop() {
    hashimoto_1.detect();
    hashimoto_2.detect();
}
