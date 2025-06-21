package service

import (
	"blog-service/model"
	"blog-service/store"

	"gofr.dev/pkg/gofr"
)

type BlogService interface {
	Create(*gofr.Context, model.Blog) (model.Blog, error)
	GetAll(*gofr.Context) ([]model.Blog, error)
	GetByID(*gofr.Context, string) (model.Blog, error)
}

type blogService struct {
	store store.BlogStore
}

func New(s store.BlogStore) BlogService {
	return &blogService{store: s}
}

func (b *blogService) Create(ctx *gofr.Context, blog model.Blog) (model.Blog, error) {
	return b.store.Create(ctx, blog)
}

func (b *blogService) GetAll(ctx *gofr.Context) ([]model.Blog, error) {
	return b.store.GetAll(ctx)
}

func (b *blogService) GetByID(ctx *gofr.Context, id string) (model.Blog, error) {
	return b.store.GetByID(ctx, id)
}
