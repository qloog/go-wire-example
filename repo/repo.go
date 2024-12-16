package repo

import "github.com/google/wire"

var RepoSet = wire.NewSet(
	NewUserRepo,
	NewProductRepo,
)
