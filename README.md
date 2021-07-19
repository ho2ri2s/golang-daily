## frontend
ルートディレクトリ配下で
```
cd frontend
```

```
yarn serve
```
ローカル環境:
http://localhost:8081/ で起動

## server
1. docker-composeでimageの起動
```
docker-compose up -d
```
-d をつけるとbackgroundでの起動になる。

2. docker image上でサーバーの起動
``` 
docker-compose exec app go run main.go
```