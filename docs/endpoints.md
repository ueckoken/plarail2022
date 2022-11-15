# 鯖エンドポイント

変更の可能性があります

## GKE

### 認証なし

- "control.chofufes2022.ueckoken.club"
`./backend/external`が動いています。クライアントはここにむけてwebsocketを張ってください。`/ws`にwebsocketのエンドポイントがあります。
- "chofufes2022.ueckoken.club"
`./frontend/site`が動いています。ここにメインのページがデプロイされます。
- "auth.chofufes2022.ueckoken.club"
認証画面です。認証が必要なページに入るには先にここを通ってください。
- "receiver-test.chofufes2022.ueckoken.club"
`./frontend/skyway_receiver/index.html`が動いています。webrtcの受信側ページです。

### 認証あり

- "grafana.chofufes2022.ueckoken.club"
grafanaというメトリクス可視化ツールが動いています。ID、パスワードはslackみてください。
- "prometheus.chofufes2022.ueckoken.club"
prometheusというメトリクス収集ツールが動いています。基本的に見なくていいです。
- "alert.chofufes2022.ueckoken.club"
使おうと思いましたがやめました。
- "webrtc-sender.chofufes2022.ueckoken.club"
`./frontend/momo_sender/index.html`が動いています。webrtcの配信者側ページです。


## 学内

- "internal.chofufes2022.ueckoken.club"
予定です。まだ取っていません。
- "speed.chofufes2022.ueckoken.club"
予定です。まだ取っていません。
