package usecase

import (
	"fmt"

	"github.com/ismailash/be-enigma-laundry/model/entity"
	"github.com/ismailash/be-enigma-laundry/repository"
)

type CustomerUseCase interface {
	FindById(id string) (entity.Customer, error)
}

type customerUseCase struct {
	customerRepo repository.CustomerRepository
}

func NewCustomerUseCase(customerRepo repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{customerRepo: customerRepo}
}

func (u *customerUseCase) FindById(id string) (entity.Customer, error) {
	customer, err := u.customerRepo.Get(id)
	if err != nil {
		return entity.Customer{}, fmt.Errorf("customer with id %s not found", id)
	}
	return customer, nil
}
