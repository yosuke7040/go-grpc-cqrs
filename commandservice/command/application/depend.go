package application

import (
	"github.com/yosuke7040/commandservice/application/impl"
	"github.com/yosuke7040/commandservice/infra/sqlboiler"

	"go.uber.org/fx"
)

// アプリケーション層の依存定義
var SrvDepend = fx.Options(
	sqlboiler.RepDepend, // SQLBoilderを利用したリポジトリインターフェイス実装
	fx.Provide(
		// サービスインターフェイス実装のコンストラクタ
		impl.NewcategoryServiceImpl,
		impl.NewproductServiceImpl,
	),
)
