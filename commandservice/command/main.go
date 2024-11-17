package main

import (
	"commandservice/presen"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		presen.CommandDepend,
	).Run()
}
