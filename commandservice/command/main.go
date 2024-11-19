package main

import (
	"github.com/yosuke7040/commandservice/presen"

	"go.uber.org/fx"
)

func main() {
	// fxを起動する
	fx.New(
		presen.CommandDepend, // 依存性を定義する
	).Run()
}
