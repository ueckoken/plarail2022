#ifndef Hallsensor_H
#define Hallsensor_H


enum SensorName {
    unknown = 0,
    
    shinjuku_d1 = 1,
    shinjuku_d2 = 2,
    
    sakurajosui_d1 = 11,
    sakurajosui_d2 = 12,
    sakurajosui_d3 = 13,
    sakurajosui_d4 = 14,
    sakurajosui_d5 = 15,
    sakurajosui_d6 = 16,
    
    chofu_d1 = 21,
    chofu_d2 = 22,
    chofu_d3 = 23,
    chofu_d4 = 24,
    chofu_d5 = 25,
    
    hashimoto_d1 = 31,
    hashimoto_d2 = 32,
    
    hachioji_d1 = 41,
    hachioji_d2 = 42,
};

// クラスの定義
class Hallsensor
{
public:
    Hallsensor(enum SensorName s, int h);
    enum SensorName sen;
    int hpin;
    StaticJsonDocument<JSON_OBJECT_SIZE(1)> doc;
    char json_string[255];
    void detect();
    void postJson();
    int ttl;
    int state = 1;
    unsigned long prevtime = 0;
    int status_code;
};

#endif
