# generate-id-use-autoincrement

DDDでの実装にあたり、EntityのIDの採番をRDBのAUTO_INCREMENTで行う場合のサンプル

# Requirement
- golang 1.20.2
- golang-migrate/migrate（sqlite3）

# Installation
## migrateを実行する
golang-migrate/migrateのsqlite3版をインストールし、次のコマンドを実行します。
```
./scripts/migrate/up.sh
```

# Usage
次の機能を用意しています
- ユーザーの作成
- ユーザーの検索

## ユーザーの作成
次のコマンドでuserを作成できます。  
nameを必須パラメータとして指定します。
```
go run cmd/main.go create-user {name}
```

## ユーザーの検索
次のコマンドでuserを検索できます。  
idを必須パラメータとして指定します。
```
go run cmd/main.go find-user {id}
```
