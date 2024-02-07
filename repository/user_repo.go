package repository

import (
	"database/sql"
	"github.com/ismailash/be-enigma-laundry/utils/common/query"

	"github.com/ismailash/be-enigma-laundry/model/entity"
)

type UserRepository interface {
	GetById(id string) (entity.User, error)
	GetByUsername(username string) (entity.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetById(id string) (entity.User, error) {
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

func (r *userRepository) GetByUsername(username string) (entity.User, error) {
	var user entity.User

	err := r.db.QueryRow(query.GetUserByUsername, username).Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.Role)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
