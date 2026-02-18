package controllers

import (
	"go-api/internal/repositories"
	"go-api/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createProductRequest struct {
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}

func CreateProductController(ctx *gin.Context, repository repositories.IProductRepository) {
	var req createProductRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	useCase := usecases.NewCreateProductUseCase(repository)
	product, err := useCase.Execute(ctx, req.Name, req.Price)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"id":      product.ID,
	})
}
