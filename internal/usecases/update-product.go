package usecases

import (
	"go-api/internal/models"
	"go-api/internal/repositories"

	"github.com/gin-gonic/gin"
)

type updateProductUseCase struct {
	repository repositories.IProductRepository
}

func NewUpdateProductUseCase(repository repositories.IProductRepository) *updateProductUseCase {
	return &updateProductUseCase{repository: repository}
}

func (un *updateProductUseCase) Execute(ctx *gin.Context, id uint, product *models.Product) error {
	err := un.repository.UpdateByID(ctx, id, product)
	if err != nil {
		return err
	}
	return nil
}
