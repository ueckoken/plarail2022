#include <ArduinoJson.h>
#include <WiFi.h>
#include "Hallsensor.hpp"

const char* ssid = "";
const char* password = "";

WiFiServer server(80);

void setup() {
    Serial.begin(115200);
    WiFi.begin(ssid, password);
    while (WiFi.status() != WL_CONNECTED) {
      delay(500);
      Serial.println("connecting");
    }
    if (WiFi.status() == WL_CONNECTED) { Serial.println("conected");}
}

Hallsensor chofu_1 = Hallsensor(chofu_d1, 34);
Hallsensor chofu_2 = Hallsensor(chofu_d2, 35);
Hallsensor chofu_3 = Hallsensor(chofu_d3, 32);
Hallsensor chofu_4 = Hallsensor(chofu_d4, 33);
Hallsensor chofu_5 = Hallsensor(chofu_d5, 25);

void loop() {
    chofu_1.detect();
    chofu_2.detect();
    chofu_3.detect();
    chofu_4.detect();
    chofu_5.detect();
}
