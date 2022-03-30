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

type CategoryHandler struct {
	categoryService service.ICategoryService
}

func NewCategoryHandler(categoryService service.ICategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
	}
}

func (ch *CategoryHandler) CreateCategory(c *gin.Context) {
	var (
		request entity.CategoryRequest
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

	res, err := ch.categoryService.CreateCategory(&request)
	if err != nil {
		response.ResponseErrorCustom(c, err, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if res == nil {
		res = &entity.CategoryResponses{}
	}

	response.ResponseSuccessWithData(c, res)
}

func (ch *CategoryHandler) GetAllCategory(c *gin.Context) {
	urlValues := c.Request.URL.Query()
	limit, _ := strconv.Atoi(urlValues.Get("limit"))
	skip, _ := strconv.Atoi(urlValues.Get("skip"))

	res, err := ch.categoryService.GetAllCategory(limit, skip)
	if err != nil {
		response.ResponseErrorCustom(c, err, err.Error(), http.StatusInternalServerError)
		return
	}

	if res.Categories == nil {
		res.Categories = make([]entity.CategoryView, 0)
	}

	response.ResponseSuccessWithData(c, res)
}

func (ch *CategoryHandler) GetDetailCategory(c *gin.Context) {
	categoryId, err := strconv.Atoi(c.Param("categoryId"))
	if err != nil {
		response.ResponseErrorCustom(c, err, err.Error(), http.StatusBadGateway)
		return
	}

	res, err := ch.categoryService.GetDetailCategory(categoryId)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			response.ResponseErrorCustom(c, err, "Category Not Found", http.StatusNotFound)
			return
		}
		response.ResponseErrorCustom(c, err, err.Error(), http.StatusInternalServerError)
		return
	}

	if res == nil {
		res = &entity.CategoryView{}
	}

	response.ResponseSuccessWithData(c, res)
}

func (ch *CategoryHandler) UpdateCategory(c *gin.Context) {
	var (
		request entity.CategoryRequest
		err     error
	)

	categoryId, err := strconv.Atoi(c.Param("categoryId"))
	if err != nil {
		response.ResponseErrorCustom(c, err, err.Error(), http.StatusBadGateway)
		return
	}

	err = c.ShouldBind(&request)
	if err != nil {
		response.ResponseErrorCustom(c, err, "Bad Request", http.StatusBadRequest)
		return
	}

	err = ch.categoryService.UpdateCategory(categoryId, request.Name)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			response.ResponseErrorCustom(c, err, "Category Not Found", http.StatusNotFound)
			return
		}
		response.ResponseErrorCustom(c, err, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ResponseSuccess(c)
}

func (ch *CategoryHandler) DeleteCategory(c *gin.Context) {
	var (
		request entity.CategoryRequest
		err     error
	)

	categoryId, err := strconv.Atoi(c.Param("categoryId"))
	if err != nil {
		response.ResponseErrorCustom(c, err, err.Error(), http.StatusBadGateway)
		return
	}

	err = c.ShouldBind(&request)
	if err != nil {
		response.ResponseErrorCustom(c, err, "Bad Request", http.StatusBadRequest)
		return
	}

	err = ch.categoryService.DeleteCategory(categoryId)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			response.ResponseErrorCustom(c, err, "Category Not Found", http.StatusNotFound)
			return
		}
		response.ResponseErrorCustom(c, err, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ResponseSuccess(c)
}
