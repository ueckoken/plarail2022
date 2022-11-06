# 閉塞区間について

閉塞区間の情報はexternalにて保持をして、その情報をATSサーバーから更新を行います。
また、管理画面から閉塞区間の上書きが指示された場合は、その情報をATSサーバーに伝えます。

ATSサーバーでは、現状のポイントの状態と閉塞区間の情報をもとにして、適切なポイント制御とValidateを実施します。

## 閉塞区間の名前について

閉塞区間は駅間ごとに設ける。
複線区間に関しては、上り線と下り線ごとに閉塞を設ける(命名規則は上りは`up`、下は`down`とする)
閉塞区間の名称は、、2つの駅の名前を`_`で繋げたもので順番は列車が降って行く順番になる。

```
shinjuku_sakurajosui_up
shinjuku_sakurajosui_down
sakurajosui_chofu_up
sakurajosui_chofu_down
chofu_hachioji_up
chofu_hachioji_down
chofu_hashimoto_up
chofu_hashimoto_down
```

また、駅構内に関しても同様に閉塞を用いて管理する。命名は以下の通り、
```
shinjuku_b1
shinjuku_b2
sakurajosui_b1
sakurajosui_b2
sakurajosui_b3
sakurajosui_b4
chofu_b1
chofu_b2
chofu_b3
chofu_b4
hashimoto_b1
hashimoto_b2
hachioji_b1
hachioji_b2
```

なお、若葉台車両基地に入線する部分に関しては特別な閉塞は設けない。
