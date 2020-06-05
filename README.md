# go-webapp-template

![Verify](https://github.com/shimataro/go-webapp-template/workflows/Verify/badge.svg)

Goを使ったウェブアプリケーションのテンプレート

## 必要なソフトウェア

* Go 1.12以上

## ビルド

```bash
go build
```

## テスト

```bash
go vet        # 静的解析
go test ./... # テスト
```

## 実行

### ローカル環境で実行

必要なもの

- ビルド後のバイナリ or ソースファイル
- [`configFiles`](./configFiles)ディレクトリ

環境変数`ENV`で環境を指定
（指定しないと起動できない）

```bash
# 開発環境
ENV=development go run main.go

# 本番環境
ENV=production go run main.go
```

<http://localhost:8080/api/v1/users>にアクセス

```bash
curl -H "X-Requested-With: x" http://localhost:8080/api/v1/users
```

### すでに開いているファイルディスクリプタをListenする場合

```bash
# 開発環境
ENV=development go run main.go --fd FILE_DESCRIPTOR

# 本番環境
ENV=production go run main.go --fd FILE_DESCRIPTOR
```
