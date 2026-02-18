package repositories

import (
	"context"
	model "go-api/internal/models"
)

type IProductRepository interface {
	Save(ctx context.Context, product *model.Product) error
	ListAll(ctx context.Context) ([]*model.Product, error)
	ListByID(ctx context.Context, id uint) (*model.Product, error)
}
