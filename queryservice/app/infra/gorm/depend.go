package gorm

import (
	"github.com/yosuke7040/queryservice/infra/gorm/adapter"
	"github.com/yosuke7040/queryservice/infra/gorm/handler"
	"github.com/yosuke7040/queryservice/infra/gorm/repository"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

// データベース接続
var DBModule = fx.Provide(func() (*gorm.DB, error) {
	return handler.ConnectDB()
})

// gorm.DB、Adapter、Repositoryの依存定義
var RepDepend = fx.Options(
	DBModule,
	fx.Provide(
		// Adapterインターフェス実装のコンストラクタを指定
		adapter.NewcategoryAdapterImpl,
		adapter.NewproductAdapterImpl,
		// Repositoryインターフェス実装のコンストラクタを指定
		repository.NewcategoryRepositoryGORM, // カテゴリ用Repository
		repository.NewproductRepositoryGORM,  // 商品用Repository
	),
)
