# plarail2022-soft

## これはなに？

このリポジトリでは電気通信大学工学研究部 調布祭展示企画のためのソースコードを管理します。

## コントリビュータのみなさんへ

GitHub Actions による CI/CD を走らせるために main ブランチへの直接コミットを禁止しています。 プルリクエストを発行してください。

PR の Reviewer には似たような PR で指定されている人を指定してください。内容ごとにタグを付けているので調べるときはぜひ活用ください。

PR が Approve されたら PR の Author が直接マージして下さい。

## 動かす方法

- [環境変数ドキュメント](./docs/environmentValList.md)に従って環境変数を設定する。

## CI/CD パイプラインについて

このリポジトリでは CI/CD によるビルド、デプロイの自動化を行なっています。
github actions によりイメージのビルドを行い、fluxcd により新しいイメージを自動でサーバに適用しています。

### 使い方

どんどん main にマージしていってください。./manifests 以下への変更は即時で fluxcd により反映されます。

docker イメージの更新やるぞ!と思ったら main から deployment にマージしてください。github actions による docker image のビルドが走ります。

新しいイメージができたら fluxcd が自動で検出して main に更新 commit を加えるところまでやります。鯖への反映は 2 分ぐらいかかります。

## 各種ドキュメント

- [環境変数に関するドキュメント](./docs/environmentValList.md)
- [ESP32 に送信する JSON に関するドキュメント](./docs/api.md)
- [External<->Internal で用いるプロトコルに関するドキュメント(未整備)](./docs/protocolBuf.md)
- [エンドポイントに関するドキュメント](./docs/endpoints.md)
- [配線図](./docs/map.md)
- [システムアーキテクチャ](./docs/archtecture.png)
