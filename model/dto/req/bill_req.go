package req

import "github.com/ismailash/be-enigma-laundry/model/entity"

type BillReqDTO struct {
	Id          string              `json:"id"`
	CustomerId  string              `json:"customerId"`
	UserId      string              `json:"userId"`
	BillDetails []entity.BillDetail `json:"billDetails"`
}
