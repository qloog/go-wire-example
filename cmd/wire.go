//go:build wireinject
// +build wireinject

package main

import (
	"go-wire-example/handler"

	"github.com/google/wire"
)

type App struct {
	UserHandler    *handler.UserHandler
	ProductHandler *handler.ProductHandler
}

func InitializeApp() (*App, error) {
	wire.Build(
		handler.HandlerSet,         // 汇总所有 Handler 的依赖
		wire.Struct(new(App), "*"), // 自动注入到 App 结构
	)

	return &App{}, nil
}
