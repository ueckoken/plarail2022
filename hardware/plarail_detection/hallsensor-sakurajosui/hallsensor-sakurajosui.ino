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

struct Station sakurajosui[6];

const char* ssid = "";
const char* password = "";

HTTPClient http;

void setup() {
    Serial.begin(115200);

    WiFi.begin(ssid, password);
    while (WiFi.status() != WL_CONNECTED) {
      delay(500);
      Serial.println("connecting");
    }
    if (WiFi.status() == WL_CONNECTED) { Serial.println("conected");}

    http.begin("http://:8081/sensor");
    http.addHeader("Content-Type", "application/json");

    setupStation(&sakurajosui[0], 34, sakurajosui_d1);
    setupStation(&sakurajosui[1], 35, sakurajosui_d2);
    setupStation(&sakurajosui[2], 32, sakurajosui_d3);
    setupStation(&sakurajosui[3], 33, sakurajosui_d4);
    setupStation(&sakurajosui[4], 25, sakurajosui_d5);
    setupStation(&sakurajosui[5], 26, sakurajosui_d6);
}

void loop() {
    for (int i = 0; i < 6; i++) {
        detect(&sakurajosui[i]);
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
