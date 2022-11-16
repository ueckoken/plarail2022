#include <ArduinoJson.h>
#include <HTTPClient.h>
#include "Hallsensor.hpp"

HTTPClient http;

Hallsensor::Hallsensor(enum SensorName s, int h) {
    sen = s;
    hpin = h;
    pinMode(hpin, INPUT);
    doc["Sensor"] = sen;
    serializeJson(doc, json_string, sizeof(json_string));
    http.begin("http://lacalhost8080:8081/sensor");
    http.addHeader("Content-Type", "application/json");
}

void Hallsensor::postJson() {
    ttl = 5;
    status_code = http.POST((uint8_t *)json_string, strlen(json_string));
    while (status_code == -2 && ttl > 0) {
        status_code = http.POST((uint8_t *)json_string, strlen(json_string));
        ttl--;
    }
    if (status_code == 200){
        prevtime = millis();
        Serial.println("[POST]Send to server");
    }
    else{
        Serial.println(status_code);
        Serial.println("[POST]failed to send to server");
        Serial.println(json_string);
    }
}

void Hallsensor::detect() {
    if(digitalRead(hpin) == HIGH) {
        if( state == 0 && millis() - prevtime > 3000 ) {
            postJson();
        }
        state = 1;
    } else {
        if(state == 1 && millis() - prevtime > 3000) {
            postJson();
        }
        state = 0;
    }
}
