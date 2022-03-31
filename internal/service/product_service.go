package service

import (
	"pos-is-backend/internal/domain/entity"
	"pos-is-backend/internal/domain/repository"
)

type ProductService struct {
	productRepo repository.IProductRepository
}

type IProductService interface {
	CreateProduct(request *entity.ProductRequest) (*entity.ProductResponse, error)
	GetAllProducts(limit, skip, categoryId int, q string) (*entity.ListProductResponse, error)
	GetDetailProduct(productId int) (*entity.ProductView, error)
	UpdateProduct(productId int, request *entity.ProductRequest) error
	DeleteProduct(productId int) error
}

func NewProductService(productRepo repository.IProductRepository) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

func (ps *ProductService) CreateProduct(request *entity.ProductRequest) (*entity.ProductResponse, error) {
	req := entity.Product{
		CategoriesId: request.CategoryId,
		Name:         request.Name,
		Sku:          request.Sku,
		Image:        request.Image,
		Price:        request.Price,
		Stock:        request.Price,
	}

	res, err := ps.productRepo.CreateProduct(&req)
	if err != nil {
		return nil, err
	}

	return &entity.ProductResponse{
		Id:           req.Id,
		CategoriesId: res.CategoriesId,
		Name:         res.Name,
		Sku:          res.Sku,
		Image:        res.Image,
		Price:        res.Price,
		Stock:        res.Stock,
		CreatedAt:    res.CreatedAt,
		UpdatedAt:    res.UpdatedAt,
	}, nil
}

func (ps *ProductService) GetAllProducts(limit, skip, categoryId int, q string) (*entity.ListProductResponse, error) {
	var (
		res   entity.ListProductResponse
		err   error
		total int
	)

	res.Meta.Limit = limit
	res.Meta.Skip = skip
	total, err = ps.productRepo.TotalProduct(categoryId, q)
	if err != nil {
		return &res, err
	}

	res.Meta.Total = total
	products, err := ps.productRepo.GetAllProducts(limit, skip, categoryId, q)
	if err != nil {
		res.Products = make([]entity.ProductView, 0)
		return &res, nil
	}

	res.Products = products.Products

	return &res, nil
}

func (ps *ProductService) GetDetailProduct(productId int) (*entity.ProductView, error) {
	var (
		res entity.ProductView
		err error
	)

	product, category, err := ps.productRepo.GetProductByID(productId)
	if err != nil {
		return &res, err
	}

	res.Category = *category
	res.Discount = nil
	res.Image = product.Image
	res.Name = product.Name
	res.Price = product.Price
	res.ProductId = product.Id
	res.Sku = product.Sku
	res.Stock = product.Stock

	return &res, nil
}

func (ps *ProductService) UpdateProduct(productId int, request *entity.ProductRequest) error {
	_, _, err := ps.productRepo.GetProductByID(productId)
	if err != nil {
		return err
	}

	err = ps.productRepo.UpdateProduct(&entity.Product{
		Id:           productId,
		CategoriesId: request.CategoryId,
		Name:         request.Name,
		Sku:          request.Sku,
		Image:        request.Image,
		Price:        request.Price,
		Stock:        request.Stock,
	})
	if err != nil {
		return err
	}

	return nil
}

func (ps *ProductService) DeleteProduct(productId int) error {
	_, _, err := ps.productRepo.GetProductByID(productId)
	if err != nil {
		return err
	}

	err = ps.productRepo.DeleteProduct(productId)
	if err != nil {
		return err
	}

	return nil
}
