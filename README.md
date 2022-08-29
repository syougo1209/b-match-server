## .env.sampleをコピーして.envを作成
```shell script
$ cp .env.sample .env
```

## migration
sqldefをインストール
```
go install github.com/k0kubun/sqldef/cmd/mysqldef@latest
```

schemaに変更があるときはschema.sqlを編集して、下記コマンドを実行
```shell script
$ make migrate
```
