package usecases

import (
	"go-api/internal/models"
	"go-api/internal/repositories"

	"github.com/gin-gonic/gin"
)

type listProductsUseCase struct {
	repository repositories.IProductRepository
}

func NewListProductsUseCase(repository repositories.IProductRepository) *listProductsUseCase {
	return &listProductsUseCase{repository: repository}
}

func (un *listProductsUseCase) Execute(ctx *gin.Context) ([]*models.Product, error) {
	products, err := un.repository.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}
