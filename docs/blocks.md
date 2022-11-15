# 閉塞区間について

閉塞区間の情報はexternalにて保持をして、その情報をATSサーバーから更新を行います。
また、管理画面から閉塞区間の上書きが指示された場合は、その情報をATSサーバーに伝えます。

ATSサーバーでは、現状のポイントの状態と閉塞区間の情報をもとにして、適切なポイント制御とValidateを実施します。

## 閉塞区間の名前について

閉塞区間の区切りは配線図に記載してある。
閉塞区間の区切れ目には必ずストップレールが設置されており、閉塞区間は[センサー]-[ストップレール]といった構成になる。
命名規則は以下の通り。

```
shinjuku_b1
shinjuku_b2
sakurajosui_b1
sakurajosui_b2
sakurajosui_b3
sakurajosui_b4
sakurajosui_b5
sakurajosui_b6
chofu_b1
chofu_b2
chofu_b3
chofu_b4
chofu_b5
hashimoto_b1
hashimoto_b2
hachioji_b1
hachioji_b2
```

なお、若葉台車両基地に入線する部分に関しては特別な閉塞は設けない。
