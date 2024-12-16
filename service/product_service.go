package service

import (
	"go-wire-example/cache"
	"go-wire-example/repo"
)

type ProductService struct {
	Repo  *repo.ProductRepo
	Cache *cache.ProductCache
}

func NewProductService(repo *repo.ProductRepo, cache *cache.ProductCache) *ProductService {
	return &ProductService{
		Repo:  repo,
		Cache: cache,
	}
}
