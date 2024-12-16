package repo

import (
	"context"
	"go-wire-example/cache"
)

type UserRepo struct {
	Cache *cache.UserCache
}

func NewUserRepo(cache *cache.UserCache) *UserRepo {
	return &UserRepo{
		Cache: cache,
	}
}

func (r *UserRepo) GetUserByUsername(ctx context.Context, username string) (*UserResponse, error) {
	// 假设从数据库查询用户
	return &UserResponse{
		Username: username,
		Password: "123456",
	}, nil
}

type UserResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
