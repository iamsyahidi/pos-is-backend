package handler

import (
	"net/http"
	"pos-is-backend/internal/domain/entity"
	"pos-is-backend/internal/service"
	"pos-is-backend/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CashierHandler struct {
	cashierService service.ICashierService
}

func NewCashierHandler(cashierService service.ICashierService) *CashierHandler {
	return &CashierHandler{
		cashierService: cashierService,
	}
}

func (ch *CashierHandler) CreateCashier(c *gin.Context) {
	var (
		request entity.CashierRequest
		err     error
	)

	err = c.ShouldBind(&request)
	if err != nil {
		response.ResponseErrorCustom(c, err, "Bad Request", http.StatusBadRequest)
		return
	}

	err = request.Validate()
	if err != nil {
		response.ResponseErrorCustom(c, err, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := ch.cashierService.CreateCashier(&request)
	if err != nil {
		response.ResponseErrorCustom(c, err, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if res == nil {
		res = &entity.CashierResponse{}
	}

	response.ResponseSuccessWithData(c, res)
}

func (ch *CashierHandler) GetAllCashier(c *gin.Context) {
	urlValues := c.Request.URL.Query()
	limit, err := strconv.Atoi(urlValues.Get("limit"))
	if err != nil {
		response.ResponseErrorCustom(c, err, err.Error(), http.StatusBadGateway)
		return
	}

	skip, err := strconv.Atoi(urlValues.Get("skip"))
	if err != nil {
		response.ResponseErrorCustom(c, err, err.Error(), http.StatusBadGateway)
		return
	}

	res, err := ch.cashierService.GetAllCashier(limit, skip)
	if err != nil {
		response.ResponseErrorCustom(c, err, err.Error(), http.StatusInternalServerError)
		return
	}

	if res == nil {
		res = &entity.ListCashierResponse{}
	}

	response.ResponseSuccessWithData(c, res)
}

func (ch *CashierHandler) GetDetailCashier(c *gin.Context) {
	cashierId, err := strconv.Atoi(c.Param("cashierId"))
	if err != nil {
		response.ResponseErrorCustom(c, err, err.Error(), http.StatusBadGateway)
		return
	}

	res, err := ch.cashierService.GetDetailCashier(cashierId)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			response.ResponseErrorCustom(c, err, "Cashier Not Found", http.StatusNotFound)
			return
		}
		response.ResponseErrorCustom(c, err, err.Error(), http.StatusInternalServerError)
		return
	}

	if res == nil {
		res = &entity.CashierView{}
	}

	response.ResponseSuccessWithData(c, res)
}

func (ch *CashierHandler) UpdateCashier(c *gin.Context) {
	var (
		request entity.CashierRequest
		err     error
	)

	cashierId, err := strconv.Atoi(c.Param("cashierId"))
	if err != nil {
		response.ResponseErrorCustom(c, err, err.Error(), http.StatusBadGateway)
		return
	}

	err = c.ShouldBind(&request)
	if err != nil {
		response.ResponseErrorCustom(c, err, "Bad Request", http.StatusBadRequest)
		return
	}

	err = ch.cashierService.UpdateCashier(cashierId, request.Name)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			response.ResponseErrorCustom(c, err, "Cashier Not Found", http.StatusNotFound)
			return
		}
		response.ResponseErrorCustom(c, err, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ResponseSuccess(c)
}

func (ch *CashierHandler) DeleteCashier(c *gin.Context) {
	var (
		request entity.CashierRequest
		err     error
	)

	cashierId, err := strconv.Atoi(c.Param("cashierId"))
	if err != nil {
		response.ResponseErrorCustom(c, err, err.Error(), http.StatusBadGateway)
		return
	}

	err = c.ShouldBind(&request)
	if err != nil {
		response.ResponseErrorCustom(c, err, "Bad Request", http.StatusBadRequest)
		return
	}

	err = ch.cashierService.DeleteCashier(cashierId, request.Name)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			response.ResponseErrorCustom(c, err, "Cashier Not Found", http.StatusNotFound)
			return
		}
		response.ResponseErrorCustom(c, err, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ResponseSuccess(c)
}
