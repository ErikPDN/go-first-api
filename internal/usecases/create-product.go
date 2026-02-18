package usecases

import (
	"context"
	"go-api/internal/models"
	"go-api/internal/repositories"
)

type createProductUseCase struct {
	repository repositories.IProductRepository
}

func NewCreateProductUseCase(repository repositories.IProductRepository) *createProductUseCase {
	return &createProductUseCase{repository: repository}
}

func (uc *createProductUseCase) Execute(ctx context.Context, name string, price float64) (*models.Product, error) {
	product, err := models.NewProductModel(name, price)

	if err != nil {
		return nil, err
	}

	err = uc.repository.Save(ctx, product)
	if err != nil {
		return nil, err
	}

	return product, nil
}
