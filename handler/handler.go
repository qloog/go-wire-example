package handler

import (
	"go-wire-example/service"

	"github.com/google/wire"
)

var HandlerSet = wire.NewSet(
	// Note: NewXxxHandler equeal wire.Struct(new(XxxHandler), "*")
	NewUserHandler,
	// wire.Struct(new(UserHandler), "*"),
	NewProductHandler,
	service.ServiceSet,
)

type Handler struct {
	UserHandler    *UserHandler
	ProductHandler *ProductHandler
}

var (
	Handle *Handler
)
