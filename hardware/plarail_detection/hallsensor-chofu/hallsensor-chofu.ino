#include <ArduinoJson.h>
#include <WiFi.h>
#include <HTTPClient.h>
#include "Hallsensor.hpp"

struct Station {
    int pin;
    int state;
    unsigned long prevtime;
    char jsonString[255];
};

struct Station chofu[5];

const char* ssid = "";
const char* password = "";
const char* json2grpcAddr = "http://:54322/sensor";
const IPAddress ipaddr(192, 168, 0, 0);  // ESP32 static IP address
const IPAddress gateway(192, 168, 0, 0);  // default gateway
const IPAddress subnet(255, 255, 254, 0);  // subnet mask

HTTPClient http;

void setup() {
    Serial.begin(115200);
    if (!WiFi.config(ipaddr, gateway, subnet)) {
       Serial.println("Failed to configure!");
       while (1);
    }
    WiFi.begin(ssid, password);
    while (WiFi.status() != WL_CONNECTED) {
      delay(500);
      Serial.println("connecting");
    }
    if (WiFi.status() == WL_CONNECTED) { Serial.println("conected");}

    http.begin(json2grpcAddr);
    http.addHeader("Content-Type", "application/json");

    setupStation(&chofu[0], 34, chofu_d1);
    setupStation(&chofu[1], 35, chofu_d2);
    setupStation(&chofu[2], 32, chofu_d3);
    setupStation(&chofu[3], 33, chofu_d4);
    setupStation(&chofu[4], 25, chofu_d5);
}

void loop() {
    for (int i = 0; i < 5; i++) {
        detect(&chofu[i]);
    }
}



StaticJsonDocument<JSON_OBJECT_SIZE(1)> doc;

void setupStation(struct Station *s, const int p, enum SensorName d){
    s->pin = p;
    pinMode(s->pin, INPUT);
    s->state = 1;
    s->prevtime = 0;
    doc["Sensor"] = d;
    serializeJson(doc, s->jsonString, sizeof(s->jsonString));
}

void sendPost(struct Station *s) {
    int ttl = 5;
    int status_code = http.POST((uint8_t *)s->jsonString, strlen(s->jsonString));
    while (status_code == -2 && ttl > 0){
        status_code = http.POST((uint8_t *)s->jsonString, strlen(s->jsonString));
        ttl--;
    }
    if (status_code == 200){
        s->prevtime = millis();
        Serial.println("[POST]Send to server");
        Serial.println(http.getString());
    }
    else{
        Serial.println(status_code);
        Serial.println("[POST]failed to send to server");
    }
}

void detect(struct Station *s) {
    if(digitalRead(s->pin) == HIGH){
        if(s->state == 0 && millis() - s->prevtime > 3000) {
            sendPost(s);
        }
        s->state = 1;
    } else {
        if(s->state == 1 && millis() - s->prevtime > 3000) {
            sendPost(s);
        }
        s->state = 0;
    }
}
