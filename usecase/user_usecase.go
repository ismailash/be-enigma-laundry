package usecase

import (
	"fmt"

	"github.com/ismailash/be-enigma-laundry/model/entity"
	"github.com/ismailash/be-enigma-laundry/repository"
)

type UserUseCase interface {
	FindById(id string) (entity.User, error)
}

type userUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{userRepo: userRepo}
}

func (u *userUseCase) FindById(id string) (entity.User, error) {
	user, err := u.userRepo.Get(id)
	if err != nil {
		return entity.User{}, fmt.Errorf("user with id %s not found", id)
	}
	return user, nil
}
