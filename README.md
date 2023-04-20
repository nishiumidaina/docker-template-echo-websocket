# docker-echo-websocket-template
Echoでソケット通信を使用できるテンプレート。  
## 構成
[Echo](https://echo.labstack.com)：フレームワーク  
[golang-migrate](https://github.com/golang-migrate/migrate)：マイグレーションツール  
```
docker-echo-websocket-template
 ├──controllers
 |   ├──router.go(ルーティングの設定)
 |   └──websocket
 |      └──websocket_controller.go
 ├──database(マイグレーション・SQLファイル)
 ├──docker
 |   └──go
 |      └──Dockerfile
 ├──models
 |   └──databese.go(DBの接続設定)
 ├──.air.toml
 ├──.env
 ├──.env.example
 ├──.gitignore
 ├──docker-compose.yaml
 ├──go.mod
 ├──go.sum
 └──main.go
```
## 環境構築
1.コンテナを起動
```
docker compose up -d --build
```
2.コンテナに入る
```
docker container exec -it docker-echo-websocket-template-app-1 bash
```
3.ビルド, サーバー起動
```
air
```
4.終了
```
exit
```