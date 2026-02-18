package usecases

import (
	"context"
	"go-api/internal/models"
	"go-api/internal/repositories"
)

type listProductUseCase struct {
	repository repositories.IProductRepository
}

func NewListProductUseCase(repository repositories.IProductRepository) *listProductUseCase {
	return &listProductUseCase{repository: repository}
}

func (un *listProductUseCase) Execute(ctx context.Context, id uint) (*models.Product, error) {
	product, err := un.repository.ListByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
