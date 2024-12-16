package cache

import "github.com/google/wire"

var CacheSet = wire.NewSet(
	NewUserCache,
	NewProductCache,
)
