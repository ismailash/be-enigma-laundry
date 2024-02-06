package usecase

import (
	"fmt"
	"github.com/ismailash/be-enigma-laundry/utils/model_util"
	"log"

	req "github.com/ismailash/be-enigma-laundry/model/dto/req"
	"github.com/ismailash/be-enigma-laundry/model/entity"
	"github.com/ismailash/be-enigma-laundry/repository"
)

type BillUseCase interface {
	RegisterNewBill(billReq req.BillReqDTO) (entity.Bill, error)
	FindById(id string) (entity.Bill, error)
	GetBillsWithPagination(paging model_util.Paging) ([]entity.Bill, error)
}

type billUseCase struct {
	billRepo        repository.BillRepository
	userUseCase     UserUseCase
	customerUseCase CustomerUseCase
	productUseCase  ProductUseCase
}

func NewBillUseCase(
	billRepo repository.BillRepository,
	userUseCase UserUseCase,
	customerUseCase CustomerUseCase,
	productUseCase ProductUseCase,
) BillUseCase {
	return &billUseCase{
		billRepo:        billRepo,
		userUseCase:     userUseCase,
		customerUseCase: customerUseCase,
		productUseCase:  productUseCase,
	}
}

func (u *billUseCase) RegisterNewBill(billReq req.BillReqDTO) (entity.Bill, error) {
	customer, err := u.customerUseCase.FindById(billReq.CustomerId)
	if err != nil {
		return entity.Bill{}, err
	}

	user, err := u.userUseCase.FindById(billReq.UserId)
	if err != nil {
		return entity.Bill{}, err
	}

	var billDetails []entity.BillDetail
	for _, v := range billReq.BillDetails {
		product, err := u.productUseCase.FindById(v.Product.Id)
		if err != nil {
			return entity.Bill{}, err
		}

		billDetails = append(billDetails, entity.BillDetail{Product: product, Qty: v.Qty, Price: product.Price})
	}

	newBill := entity.Bill{
		Customer:    customer,
		User:        user,
		BillDetails: billDetails,
	}

	bill, err := u.billRepo.Create(newBill)
	if err != nil {
		return entity.Bill{}, err
	}

	return bill, nil
}

func (u *billUseCase) FindById(id string) (entity.Bill, error) {
	log.Println("USECASE DISINIII BANGGGG >> ", id)
	bill, err := u.billRepo.Get(id)
	if err != nil {
		return entity.Bill{}, fmt.Errorf("bill with id %s not found", id)
	}
	return bill, nil
}

func (u *billUseCase) GetBillsWithPagination(paging model_util.Paging) ([]entity.Bill, error) {
	bills, err := u.billRepo.GetWithPagination(paging)
	if err != nil {
		return nil, err
	}

	return bills, nil
}
