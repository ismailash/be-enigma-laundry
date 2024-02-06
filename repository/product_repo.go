package repository

import (
	"database/sql"

	"github.com/ismailash/be-enigma-laundry/model/entity"
)

type ProductRepository interface {
	Get(id string) (entity.Product, error)
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Get(id string) (entity.Product, error) {
	var product entity.Product
	query := `SELECT id, name, price, type, created_at, updated_at FROM products WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(
		&product.Id,
		&product.Name,
		&product.Price,
		&product.Type,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		return entity.Product{}, err
	}

	return product, nil
}
