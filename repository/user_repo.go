package repository

import (
	"database/sql"

	"github.com/ismailash/be-enigma-laundry/model/entity"
)

type UserRepository interface {
	Get(id string) (entity.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Get(id string) (entity.User, error) {
	var user entity.User
	query := `SELECT id, name, email, username, password, role, created_at, updated_at FROM users WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
