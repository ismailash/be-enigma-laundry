package repository

import (
	"database/sql"

	"github.com/ismailash/be-enigma-laundry/model/entity"
)

type CustomerRepository interface {
	Get(id string) (entity.Customer, error)
}

type customerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) Get(id string) (entity.Customer, error) {
	var customer entity.Customer
	query := `SELECT id, name, phone_number, address, created_at, updated_at FROM customers WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(
		&customer.Id,
		&customer.Name,
		&customer.PhoneNumber,
		&customer.Address,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	)
	if err != nil {
		return entity.Customer{}, err
	}

	return customer, nil
}
