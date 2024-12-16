package service

import (
	"go-wire-example/cache"
	"go-wire-example/repo"

	"github.com/google/wire"
)

// ServiceSet 配置 service 的依赖项
var ServiceSet = wire.NewSet(
	NewUserService,
	NewProductService,
	repo.RepoSet,
	cache.CacheSet,
)
