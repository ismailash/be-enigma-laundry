package usecase

import (
	"fmt"

	"github.com/ismailash/be-enigma-laundry/model/entity"
	"github.com/ismailash/be-enigma-laundry/repository"
)

type ProductUseCase interface {
	FindById(id string) (entity.Product, error)
}

type productUseCase struct {
	productRepo repository.ProductRepository
}

func NewProductUseCase(productRepo repository.ProductRepository) ProductUseCase {
	return &productUseCase{productRepo: productRepo}
}

func (u *productUseCase) FindById(id string) (entity.Product, error) {
	product, err := u.productRepo.Get(id)
	if err != nil {
		return entity.Product{}, fmt.Errorf("product with id %s not found", id)
	}
	return product, nil
}
