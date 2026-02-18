package controllers

import (
	"go-api/internal/models"
	"go-api/internal/repositories"
	"go-api/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type updateProductRequest struct {
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}

func UpdateProductController(ctx *gin.Context, id uint, repository repositories.IProductRepository) {
	var req updateProductRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	product := &models.Product{
		Name:  req.Name,
		Price: req.Price,
	}

	useCase := usecases.NewUpdateProductUseCase(repository)
	err := useCase.Execute(ctx, id, product)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
