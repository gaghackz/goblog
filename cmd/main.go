package main

import (
	"blog-service/factory"
	"blog-service/handler"
	"blog-service/service"

	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	// Dependencies
	store := factory.NewBlogStore(app)
	blogService := service.New(store)
	blogHandler := handler.New(blogService)

	// Routes
	app.GET("/blogs", blogHandler.GetAll)
	app.GET("/blogs/{id}", blogHandler.GetByID)
	app.POST("/blogs", blogHandler.Create)

	app.Run()
}
