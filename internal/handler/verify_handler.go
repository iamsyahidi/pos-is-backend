package handler

import (
	"net/http"
	"pos-is-backend/internal/domain/entity"
	"pos-is-backend/internal/service"
	"pos-is-backend/pkg/auth"
	"pos-is-backend/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type VerifyHandler struct {
	verifyService service.IVerifyService
}

func NewVerifyHandler(verifyService service.IVerifyService) *VerifyHandler {
	return &VerifyHandler{
		verifyService: verifyService,
	}
}

func (vh *VerifyHandler) LoginPasscode(c *gin.Context) {
	var (
		request entity.Verify
		err     error
	)

	err = c.ShouldBindJSON(&request)
	if err != nil {
		response.ResponseErrorCustom(c, err, "Bad Request", http.StatusBadRequest)
		return
	}

	cashierId, err := strconv.Atoi(c.Param("cashierId"))
	if err != nil {
		response.ResponseErrorCustom(c, err, err.Error(), http.StatusBadGateway)
		return
	}

	res, err := vh.verifyService.LoginPasscode(cashierId, request.Passcode)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			response.ResponseErrorCustom(c, err, "Passcode Not Match", http.StatusNotFound)
			return
		}
		response.ResponseErrorCustom(c, err, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if res == nil {
		res = &auth.TokenDetails{}
	}

	response.ResponseSuccessWithData(c, res)
}

func (vh *VerifyHandler) GetPasscode(c *gin.Context) {
	cashierId, err := strconv.Atoi(c.Param("cashierId"))
	if err != nil {
		response.ResponseErrorCustom(c, err, err.Error(), http.StatusBadGateway)
		return
	}

	res, err := vh.verifyService.GetPasscode(cashierId)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			response.ResponseErrorCustom(c, err, "Cashier Not Found", http.StatusNotFound)
			return
		}
		response.ResponseErrorCustom(c, err, err.Error(), http.StatusInternalServerError)
		return
	}

	if res == nil {
		res = &entity.Verify{}
	}

	response.ResponseSuccessWithData(c, res)
}

func (vh *VerifyHandler) LogoutPasscode(c *gin.Context) {
	var (
		request entity.Verify
		err     error
	)

	err = c.ShouldBindJSON(&request)
	if err != nil {
		response.ResponseErrorCustom(c, err, "Bad Request", http.StatusBadRequest)
		return
	}

	cashierId, err := strconv.Atoi(c.Param("cashierId"))
	if err != nil {
		response.ResponseErrorCustom(c, err, err.Error(), http.StatusBadGateway)
		return
	}

	err = vh.verifyService.LogoutPasscode(cashierId, request.Passcode)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			response.ResponseErrorCustom(c, err, "Passcode Not Match", http.StatusNotFound)
			return
		}
		response.ResponseErrorCustom(c, err, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response.ResponseSuccess(c)
}
