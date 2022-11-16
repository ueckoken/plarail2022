#include <ArduinoJson.h>
#include <WiFi.h>
#include <HTTPClient.h>
#define LED_PIN 13
#define HL_PIN 27

const char* ssid = "";
const char* password = "";

StaticJsonDocument<JSON_OBJECT_SIZE(1)> doc;
char json_string[255];
    
void setupjson(){
    doc["Sensor"] = 1;

    serializeJson(doc, json_string, sizeof(json_string));
}

HTTPClient http;

void setup() {
    Serial.begin(115200);
    WiFi.begin(ssid, password);
    setupjson();
    pinMode(LED_PIN, OUTPUT);
    pinMode(HL_PIN, INPUT);
    http.begin("http://localhost:8080/sensor");
    http.addHeader("Content-Type", "application/json");
}

int state = 1;
unsigned long prevtime = 0;

void loop() {
    if(digitalRead(HL_PIN) == HIGH){
        digitalWrite(LED_PIN, HIGH);
        if(state == 0 && millis() - prevtime > 3000) {
            int ttl = 5;
            int status_code = http.POST((uint8_t *)json_string, strlen(json_string));
            while (status_code == -2 && ttl > 0){
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
            }
        }
        state = 1;
    } else {
        digitalWrite(LED_PIN, LOW);
        if(state == 1 && millis() - prevtime > 3000) {
            int ttl = 5;
            int status_code = http.POST((uint8_t *)json_string, strlen(json_string));
            while (status_code == -2 && ttl > 0){
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
            }
        }
        state = 0;
    }
}
