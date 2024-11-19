package presen

import (
	"github.com/yosuke7040/queryservice/infra/gorm"
	"github.com/yosuke7040/queryservice/presen/builder"
	"github.com/yosuke7040/queryservice/presen/prepare"
	"github.com/yosuke7040/queryservice/presen/server"

	"go.uber.org/fx"
)

var QueryDepend = fx.Options(
	gorm.RepDepend,
	// プレゼンテーション層の依存定義
	fx.Provide(
		builder.NewresultBuilderImpl,
		server.NewcategoryServer,
		server.NewproductServerImpl,
		prepare.NewQueryServer,
	),
	// メソッドの起動
	fx.Invoke(prepare.QueryServiceLifecycle),
)
