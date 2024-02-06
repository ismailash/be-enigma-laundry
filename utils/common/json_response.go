package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ismailash/be-enigma-laundry/model/dto/res"
)

func SendCreateResponse(ctx *gin.Context, description string, data any) {
	ctx.JSON(http.StatusCreated, res.SingleResponse{
		Status: res.Status{
			Code:        http.StatusCreated,
			Description: description,
		},
		Data: data,
	})
}

// single dan paged response

func SendSingleResponse(ctx *gin.Context, description string, data any) {
	ctx.JSON(http.StatusOK, res.SingleResponse{
		Status: res.Status{
			Code:        http.StatusOK,
			Description: description,
		},
		Data: data,
	})
}

func SendErrorResponse(ctx *gin.Context, code int, description string) {
	ctx.JSON(code, res.SingleResponse{
		Status: res.Status{
			Code:        code,
			Description: description,
		},
	})
}
