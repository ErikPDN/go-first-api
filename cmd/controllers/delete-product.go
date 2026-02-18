package controllers

import (
	"go-api/internal/repositories"
	"go-api/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteProductController(ctx *gin.Context, id uint, repository repositories.IProductRepository) {
	useCase := usecases.NewDeleteProductUseCase(repository)
	err := useCase.Execute(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Product deleted successfully",
	})
}
