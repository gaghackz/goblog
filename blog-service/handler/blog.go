package handler

import (
	"blog-service/model"
	"blog-service/service"

	"gofr.dev/pkg/gofr"
)

type Handler struct {
	service service.BlogService
}

func New(s service.BlogService) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Create(ctx *gofr.Context) (interface{}, error) {
	var blog model.Blog
	if err := ctx.Bind(&blog); err != nil {
		return nil, err
	}
	return h.service.Create(ctx, blog)
}

func (h *Handler) GetAll(ctx *gofr.Context) (interface{}, error) {
	return h.service.GetAll(ctx)
}

func (h *Handler) GetByID(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	return h.service.GetByID(ctx, id)
}