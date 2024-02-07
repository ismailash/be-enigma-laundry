package usecase

import (
	"errors"
	"fmt"

	"github.com/ismailash/be-enigma-laundry/model/entity"
	"github.com/ismailash/be-enigma-laundry/repository"
)

type UserUseCase interface {
	FindById(id string) (entity.User, error)
	FindByUsernamePassword(username string, password string) (entity.User, error)
}

type userUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{userRepo: userRepo}
}

func (u *userUseCase) FindById(id string) (entity.User, error) {
	user, err := u.userRepo.GetById(id)
	if err != nil {
		return entity.User{}, fmt.Errorf("user with id %s not found", id)
	}
	return user, nil
}

func (u *userUseCase) FindByUsernamePassword(username string, password string) (entity.User, error) {
	user, err := u.userRepo.GetByUsername(username)
	if err != nil {
		return entity.User{}, errors.New("invalid username or password")
	}

	// compare password
	if user.Password != password {
		return entity.User{}, err
	}

	// set user password jadi kosong agar tidak ditampilkan di response
	user.Password = ""
	return user, nil
}
