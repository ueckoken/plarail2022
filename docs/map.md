# 2022 年調布祭プラレール企画配線図

## 配線図

[配線図 PDF](./chofufes2022-map.pdf)

![配線図](./chofufes2022-map.jpg)

駅付近にある`s`, `p` はそれぞれ制御するストップレール、分岐レールを表します。
設定ファイルなどで指定するときは以下のルールで指定します。

```text
<station_name>_<s or p><number>
```

例えば調布駅にあるストップレール `s1`を指定するときは `chofu_s1`と指定します。


名称一覧
```
新宿駅
shinjuku_p1
shinjuku_p1
shinjuku_s1
shinjuku_s2
shinjuku_d1
shinjuku_d2
桜上水駅
sakurajosui_p1
sakurajosui_p2
sakurajosui_s1
sakurajosui_s2
sakurajosui_s3
sakurajosui_s4
sakurajosui_d1
sakurajosui_d2
sakurajosui_d3
sakurajosui_d4
調布駅
chofu_p1
chofu_p2
chofu_p3
chofu_p4
chofu_p5
chofu_s1
chofu_s2
chofu_s3
chofu_s4
chofu_d1
chofu_d2
chofu_d3
chofu_d4
chofu_d5
chofu_d6
橋本
hashimoto_s1
hashimoto_s2
hashimoto_d1
hashimoto_d2
京王八王子
hachioji_s1
hachioji_s2
hachioji_d1
hachioji_d2
若葉台(車両基地)
wakabadai_p1
wakabadai_p2
wakabadai_s1
wakabadai_s2
wakabadai_d1
wakabadai_d2
```

サーボモータが ON, OFF になったときの挙動は [./api.md](https://github.com/ueckoken/plarail2021-soft/blob/main/docs/api.md#client---control-external)を参考にしてください。
