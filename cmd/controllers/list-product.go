package controllers

import (
	"go-api/internal/repositories"
	"go-api/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListProductController(ctx *gin.Context, repository repositories.IProductRepository, id uint) {
	useCase := usecases.NewListProductUseCase(repository)
	product, err := useCase.Execute(ctx, id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    product,
	})
}
