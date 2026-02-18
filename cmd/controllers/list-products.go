package controllers

import (
	"go-api/internal/repositories"
	"go-api/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListProductsController(ctx *gin.Context, repository repositories.IProductRepository) {
	useCase := usecases.NewListProductsUseCase(repository)
	products, err := useCase.Execute(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    products,
	})
}
