```shell
go get google.golang.org/grpc
go get github.com/onsi/ginkgo/v2
go get github.com/onsi/gomega
go get go.uber.org/fx
go get github.com/fullness-MFurukawa/samplepb@v1.0.0
go get github.com/gin-gonic/gin
go get github.com/swaggo/swag/cmd/swag
go get github.com/go-openapi/swag
go get github.com/swaggo/gin-swagger
go get github.com/swaggo/files

go install github.com/swaggo/swag/cmd/swag@latest

go mod tidy

swag init --parseDependency -g main.go
```
