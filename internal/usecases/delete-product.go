package usecases

import (
	"context"
	"go-api/internal/repositories"
)

type deleteProductUseCase struct {
	repository repositories.IProductRepository
}

func NewDeleteProductUseCase(repository repositories.IProductRepository) *deleteProductUseCase {
	return &deleteProductUseCase{repository: repository}
}

func (uc *deleteProductUseCase) Execute(ctx context.Context, id uint) error {
	err := uc.repository.DeleteByID(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
