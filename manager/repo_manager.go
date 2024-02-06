package manager

import "github.com/ismailash/be-enigma-laundry/repository"

type RepoManager interface {
	NewUserRepo() repository.UserRepository
	NewCustomerRepo() repository.CustomerRepository
	NewProductRepo() repository.ProductRepository
	NewBillRepo() repository.BillRepository
}

type repoManager struct {
	infra InfraManager
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}

func (r *repoManager) NewBillRepo() repository.BillRepository {
	return repository.NewBillRepository(r.infra.Conn())
}

func (r *repoManager) NewUserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.Conn())
}

func (r *repoManager) NewCustomerRepo() repository.CustomerRepository {
	return repository.NewCustomerRepository(r.infra.Conn())
}

func (r *repoManager) NewProductRepo() repository.ProductRepository {
	return repository.NewProductRepository(r.infra.Conn())
}
