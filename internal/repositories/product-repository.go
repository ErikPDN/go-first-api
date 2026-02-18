package repositories

import (
	"context"
	"database/sql"
	"errors"
	"go-api/internal/models"
)

var ErrProductNotFound = errors.New("product not found")

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Save(ctx context.Context, product *models.Product) error {
	query := `
		INSERT INTO product (name, price, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	return r.db.QueryRowContext(
		ctx,
		query,
		product.Name,
		product.Price,
		product.CreatedAt,
		product.UpdatedAt,
	).Scan(&product.ID)
}

func (r *ProductRepository) ListAll(ctx context.Context) ([]*models.Product, error) {
	query := `
		SELECT id, name, price, created_at, updated_at
		FROM product
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.Product

	for rows.Next() {
		var product models.Product

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) ListByID(ctx context.Context, id uint) (*models.Product, error) {
	query := `
		SELECT id, name, price, created_at, updated_at
		FROM product
		WHERE id = $1
	`

	var product models.Product

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, ErrProductNotFound
	}

	if err != nil {
		return nil, err
	}

	return &product, nil
}
