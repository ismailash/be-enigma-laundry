package manager

import "github.com/ismailash/be-enigma-laundry/usecase"

type UseCaseManager interface {
	NewUserUseCase() usecase.UserUseCase
	NewCustomerUseCase() usecase.CustomerUseCase
	NewProductUseCase() usecase.ProductUseCase
	NewBillUseCase() usecase.BillUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{repo: repo}
}

func (u *useCaseManager) NewBillUseCase() usecase.BillUseCase {
	return usecase.NewBillUseCase(u.repo.NewBillRepo(), u.NewUserUseCase(), u.NewCustomerUseCase(), u.NewProductUseCase())
}

func (u *useCaseManager) NewUserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repo.NewUserRepo())
}

func (u *useCaseManager) NewCustomerUseCase() usecase.CustomerUseCase {
	return usecase.NewCustomerUseCase(u.repo.NewCustomerRepo())
}

func (u *useCaseManager) NewProductUseCase() usecase.ProductUseCase {
	return usecase.NewProductUseCase(u.repo.NewProductRepo())
}
