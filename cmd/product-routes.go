package main

import (
	"go-api/cmd/controllers"
	"go-api/internal/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {
	repo := repositories.NewProductRepository(db)

	productRoute := router.Group("/products")

	productRoute.POST("/", func(ctx *gin.Context) {
		controllers.CreateProductController(ctx, repo)
	})

	productRoute.GET("/:id", func(ctx *gin.Context) {
		id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.AbortWithStatusJSON(400, gin.H{
				"success": false,
				"error":   "invalid id",
			})
			return
		}

		controllers.ListProductController(ctx, repo, uint(id))
	})

	productRoute.GET("/", func(ctx *gin.Context) {
		controllers.ListProductsController(ctx, repo)
	})

	productRoute.PUT("/:id", func(ctx *gin.Context) {
		id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.AbortWithStatusJSON(400, gin.H{
				"success": false,
				"error":   "invalid id",
			})
			return
		}

		controllers.UpdateProductController(ctx, uint(id), repo)
	})

	productRoute.DELETE("/:id", func(ctx *gin.Context) {
		id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.AbortWithStatusJSON(400, gin.H{
				"success": false,
				"error":   "invalid id",
			})
			return
		}

		controllers.DeleteProductController(ctx, uint(id), repo)
	})
}
