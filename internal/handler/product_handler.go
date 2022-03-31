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

type ProductHandler struct {
	productService service.IProductService
}

func NewProductHandler(productService service.IProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (ph *ProductHandler) CreateProduct(c *gin.Context) {
	var (
		request entity.ProductRequest
		err     error
	)

	err = c.ShouldBind(&request)
	if err != nil {
		response.ResponseErrorCustom(c, err, "Bad Request", http.StatusBadRequest)
		return
	}

	res, err := ph.productService.CreateProduct(&request)
	if err != nil {
		response.ResponseErrorCustom(c, err, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if res == nil {
		res = &entity.ProductResponse{}
	}

	response.ResponseSuccessWithData(c, res)
}

func (ph *ProductHandler) GetAllProducts(c *gin.Context) {
	urlValues := c.Request.URL.Query()
	limit, _ := strconv.Atoi(urlValues.Get("limit"))
	skip, _ := strconv.Atoi(urlValues.Get("skip"))
	categoryId, _ := strconv.Atoi(urlValues.Get("categoryId"))
	q := urlValues.Get("q")

	res, err := ph.productService.GetAllProducts(limit, skip, categoryId, q)
	if err != nil {
		response.ResponseErrorCustom(c, err, err.Error(), http.StatusInternalServerError)
		return
	}

	if res.Products == nil {
		res.Products = make([]entity.ProductView, 0)
	}

	response.ResponseSuccessWithData(c, res)
}

func (ph *ProductHandler) GetDetailProduct(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		response.ResponseErrorCustom(c, err, err.Error(), http.StatusBadGateway)
		return
	}

	res, err := ph.productService.GetDetailProduct(productId)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			response.ResponseErrorCustom(c, err, "Product Not Found", http.StatusNotFound)
			return
		}
		response.ResponseErrorCustom(c, err, err.Error(), http.StatusInternalServerError)
		return
	}

	if res == nil {
		res = &entity.ProductView{}
	}

	response.ResponseSuccessWithData(c, res)
}

func (ph *ProductHandler) UpdateProduct(c *gin.Context) {
	var (
		request entity.ProductRequest
		err     error
	)

	productId, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		response.ResponseErrorCustom(c, err, err.Error(), http.StatusBadGateway)
		return
	}

	err = c.ShouldBind(&request)
	if err != nil {
		response.ResponseErrorCustom(c, err, "Bad Request", http.StatusBadRequest)
		return
	}

	err = ph.productService.UpdateProduct(productId, &request)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			response.ResponseErrorCustom(c, err, "Product Not Found", http.StatusNotFound)
			return
		}
		response.ResponseErrorCustom(c, err, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ResponseSuccess(c)
}

func (ph *ProductHandler) DeleteProduct(c *gin.Context) {
	var (
		err error
	)

	productId, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		response.ResponseErrorCustom(c, err, err.Error(), http.StatusBadGateway)
		return
	}

	err = ph.productService.DeleteProduct(productId)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			response.ResponseErrorCustom(c, err, "Product Not Found", http.StatusNotFound)
			return
		}
		response.ResponseErrorCustom(c, err, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ResponseSuccess(c)
}
