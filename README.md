# fishing-api-server

## 環境

- Go 1.21.5
- MySQL 8.0.33

## 環境構築

下記の流れに従って、環境構築を行なってください。

#### clone

```
git@github.com:y-ttkt/fishing-api-server.git
```

#### install

```
make install
```

#### コンテナ作成
```
make up
```

#### Goコンテナへの接続
```
make shell
```


## 各種コマンド

| コマンド           | 説明                                       |
| ------------------ | ------------------------------------------ |
| `make install`     | Goモジュール＆ツールのインストール         |
| `make up`          | Dockerコンテナの起動（バックグラウンド）     |
| `make down`        | Dockerコンテナの停止＆ネットワーク削除     |
| `make shell`       | Goコンテナ内にシェルで入る                 |
| `make migrate`     | DBマイグレーションを実行                   |
| `make test`        | ユニットテスト＆統合テストを実行           |

