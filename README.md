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

## 実環境で稼働させる場合

Python3.3以上がインストールされているUbuntuを想定。

### 必要なもの

- ビルド後のバイナリ
    - `/var/www`ディレクトリに配置（ディレクトリがなければ作成）
- [`configFiles`](./configFiles)ディレクトリ
    - バイナリと同じディレクトリに配置
- プロセスマネージャー[Circus](https://circus.readthedocs.io/en/latest/)
    - 詳細は後述
    - ホットリロードに必要
- [`circus.ini`](./circus.ini)
    - バイナリと同じディレクトリに配置
    - Circusの設定ファイル
- [`go-webapp-template.service`](./go-webapp-template.service)
    - `/etc/systemd/system/`ディレクトリに配置
    - CircusのSystemd対応（自動起動）に使用

### 準備

#### ファイルを配置

ファイルを以下のディレクトリに配置
（ディレクトリがなければ作成）

- ビルド後のバイナリ、`configFiles`、`circus.ini` - `/var/www`
- `go-webapp-template.service` - `/etc/systemd/system`

#### 環境変数の設定

`/etc/environment`に以下の内容を追加

開発環境

```
ENV=development
```

本番環境

```
ENV=production
GIN_MODE=release
```

#### Circusインストール

```bash
pip3 install circus
```

### 自動起動設定

システム起動時にアプリケーションを起動する
（[ドキュメント](https://circus.readthedocs.io/en/latest/for-ops/deployment/) 参照）

```bash
sudo systemctl enable circus
```

いろんなコマンド

```bash
# 開始
sudo systemctl start go-webapp-template

# 終了
sudo systemctl stop go-webapp-template

# graceful reload
sudo systemctl reload go-webapp-template
```

### ローカル環境でCircusを使う場合

このリポジトリを`git clone`したディレクトリで以下を実行

```bash
# フォアグラウンドで起動する場合
circusd circus-local.ini

# デーモン化
circusd --daemon circus-local.ini
```

いろんなコマンド

```bash
# プロセス起動状況確認
circus-top

# graceful reload
# （circus-topが起動した状態で別ターミナルから実行すると、プロセスが入れ替わる様子がわかる）
circusctl reload

# 一時停止
circusctl stop

# 再開
circusctl start

# 完全停止（Circus自体は生きている）
circusctl rm go-webapp-template

# circus停止
circusctl quit
```
