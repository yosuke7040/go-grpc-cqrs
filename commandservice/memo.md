```shell

go get -u go.uber.org/fx
go get github.com/fullness-MFurukawa/samplepb@v1.0.0
go get -u github.com/onsi/ginkgo/v2
go get -u github.com/onsi/gomega
go get -u github.com/google/uuid
go get -u github.com/go-sql-driver/mysql
go get -u github.com/volatiletech/sqlboiler/v4
go get -u github.com/volatiletech/null/v8
go get -u github.com/BurntSushi/toml

go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest

grpcurl -plaintext localhost:8082 list
grpcurl -plaintext localhost:8082 list proto.CategoryCommand
grpcurl -plaintext -d '{"crud": "1" , "name" : "食料品" }' localhost:8082 proto.CategoryCommand.Create

grpcurl -cacert ./commandservice.pem commandservice:8082 list

go mod tidy

commandのdocker composeとsampledbのdocker composeを起動しておく
sampledbでcreate_object.sqlを実行しておく（phpMyAdminで投げれる）
dommand-serviceのコンテナ内で実行↓
sqlboiler mysql -c config/database.toml -o models -p models --no-tests --wipe
```
