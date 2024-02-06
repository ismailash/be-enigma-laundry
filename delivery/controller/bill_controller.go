package controller

import (
	"github.com/ismailash/be-enigma-laundry/utils/model_util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	req "github.com/ismailash/be-enigma-laundry/model/dto/req"
	"github.com/ismailash/be-enigma-laundry/usecase"
	"github.com/ismailash/be-enigma-laundry/utils/common"
)

type BillController struct {
	uc usecase.BillUseCase
	rg *gin.RouterGroup
}

func NewBillController(uc usecase.BillUseCase, rg *gin.RouterGroup) *BillController {
	return &BillController{uc: uc, rg: rg}
}

func (c *BillController) Route() {
	br := c.rg.Group("/bills")
	br.POST("/", c.createHandler)
	br.GET("/:id", c.getHandler)
	br.POST("/:id", c.getBillsWithPaginationHandler)
}

func (c *BillController) createHandler(ctx *gin.Context) {
	var billReq req.BillReqDTO
	if err := ctx.ShouldBindJSON(&billReq); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response, err := c.uc.RegisterNewBill(billReq)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "New bill has been created successfully", response)
}

func (c *BillController) getHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "id cannot be empty")
		return
	}

	log.Println("CTRL DISINI BANGGGGG >> ", id)

	res, err := c.uc.FindById(id)
	log.Println("CTRL DISINI BANGGGGG res >> ", res)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "OK", res)
}

func (c *BillController) getBillsWithPaginationHandler(ctx *gin.Context) {
	var paging model_util.Paging

	if err := ctx.ShouldBindJSON(&paging); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	bills, err := c.uc.GetBillsWithPagination(paging)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	var billsInterface []any
	for _, bill := range bills {
		billsInterface = append(billsInterface, bill)
	}

	common.SendPagedResponse(ctx, "OK", billsInterface, paging)
}
