# IPアドレス割り当てレンジ

192.168.1.1~100 自動割り当て分に予約

192.168.1.101~200 ESP32/Raspberry Pi

# IP固定マニュアル

192.168.1.1にブラウザで入ってNVR500の設定ページを開きます。

コマンド打つ画面に入ります。

# IPアドレスの固定手順

IPアドレスを固定したいESP32/Raspberry Pi以外のデバイスをなるべく繋がないでください。

IPアドレス固定対象のデバイスをWIFIに接続します。

`show status dhcp`でDHCPのステータスを表示します。

新しく追加されたデバイスがおそらくIPアドレス固定対象です。クライアントIDをコピーします。

もし(ff) 01 23 ... のようなクライアントIDであれば、括弧を削除して ff 01 23 ...という形にしてください。

`dhcp scope bind 1 192.168.1.$(なにか) $(クライアントID)`というコマンドを打ってください。これで固定されます。

IPアドレスは上のレンジに沿っていれば雑に決めていいです。連番とかで。

ESP32にマスキングテープを貼って、IPアドレスをマスキングテープに貼って分かるようにしてください。

ESP32の電源を消して付け直して、もう一度`show status dhcp`を叩いてください。目的のIPアドレスが設定されていることを確認してください。