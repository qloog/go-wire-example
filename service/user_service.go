package service

import (
	"context"
	"errors"
	"go-wire-example/repo"
)

type UserService struct {
	Repo *repo.UserRepo
}

func NewUserService(repo *repo.UserRepo) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func (s *UserService) Login(ctx context.Context, username, password string) (interface{}, error) {
	// 假设从数据库中验证用户
	user, err := s.Repo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if user.Password != password {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
