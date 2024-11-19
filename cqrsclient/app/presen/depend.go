package presen

import (
	"github.com/yosuke7040/cqrsqlient/infra"
	"github.com/yosuke7040/cqrsqlient/presen/gin"
	"github.com/yosuke7040/cqrsqlient/presen/gin/handler"
	"github.com/yosuke7040/cqrsqlient/presen/gin/helper"

	"go.uber.org/fx"
)

// プレゼンテーション層の依存定義
var PresonDepend = fx.Options(
	infra.InfraDepend, // インフラストラクチャ層の依存定義
	fx.Provide(
		helper.NewCommandHelper,   // ヘルパの依存定義
		helper.NewQueryHelper,     // ヘルパの依存定義
		handler.NewCommandHandler, // ハンドラの依存定義
		handler.NewQueryHandler,   // ハンドラの依存定義
		gin.NewRouter,             // ルータの依存定義
	),
	fx.Invoke(gin.RegisterHooks), // GINの起動と停止機能の実行
)
