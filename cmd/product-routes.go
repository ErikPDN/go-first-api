package main

import (
	"go-api/cmd/controllers"
	"go-api/internal/repositories"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {
	repo := repositories.NewProductRepository(db)

	productRoute := router.Group("/products")

	productRoute.POST("/", func(ctx *gin.Context) {
		controllers.CreateProductController(ctx, repo)
	})

}
