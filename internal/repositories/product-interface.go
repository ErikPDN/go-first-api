package repositories

import (
	"context"
	model "go-api/internal/models"
)

type IProductRepository interface {
	Save(ctx context.Context, product *model.Product) error
	ListAll(ctx context.Context) ([]*model.Product, error)
	ListByID(ctx context.Context, id uint) (*model.Product, error)
	UpdateByID(ctx context.Context, id uint, product *model.Product) error
	DeleteByID(ctx context.Context, id uint) error
}
