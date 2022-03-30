package service

import (
	"pos-is-backend/internal/domain/entity"
	"pos-is-backend/internal/domain/repository"
	"time"
)

type CategoryService struct {
	categoryRepo repository.ICategoryRepository
}

type ICategoryService interface {
	CreateCategory(*entity.CategoryRequest) (*entity.CategoryResponses, error)
	GetAllCategory(limit, skip int) (*entity.ListCategoryResponse, error)
	GetDetailCategory(categoryId int) (*entity.CategoryView, error)
	UpdateCategory(categoryId int, name string) error
	DeleteCategory(categoryId int) error
}

func NewCategoryService(categoryRepo repository.ICategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepo: categoryRepo,
	}
}

func (cs *CategoryService) CreateCategory(request *entity.CategoryRequest) (*entity.CategoryResponses, error) {
	req := entity.Category{
		Name: request.Name,
	}

	res, err := cs.categoryRepo.CreateCategory(&req)
	if err != nil {
		return nil, err
	}

	return &entity.CategoryResponses{
		Id:        res.Id,
		Name:      res.Name,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}

func (cs *CategoryService) GetAllCategory(limit, skip int) (*entity.ListCategoryResponse, error) {
	var (
		res          entity.ListCategoryResponse
		err          error
		total        int
		categoryView []entity.CategoryView
	)

	res.Meta.Limit = limit
	res.Meta.Skip = skip
	total, err = cs.categoryRepo.TotalCategory()
	if err != nil {
		return &res, err
	}

	res.Meta.Total = total
	cashiers, err := cs.categoryRepo.GetAllCategory(limit, skip)
	if err != nil {
		res.Categories = make([]entity.CategoryView, 0)
		return &res, nil
	}

	for _, v := range *cashiers {
		categoryView = append(categoryView, entity.CategoryView{
			CategoryId: v.Id,
			Name:       v.Name,
		})
	}

	res.Categories = categoryView

	return &res, nil
}

func (cs *CategoryService) GetDetailCategory(categoryId int) (*entity.CategoryView, error) {
	var (
		res entity.CategoryView
		err error
	)

	category, err := cs.categoryRepo.GetDetailCategory(categoryId)
	if err != nil {
		return &res, err
	}

	res.CategoryId = category.Id
	res.Name = category.Name

	return &res, nil
}

func (cs *CategoryService) UpdateCategory(categoryId int, name string) error {
	_, err := cs.categoryRepo.GetDetailCategory(categoryId)
	if err != nil {
		return err
	}

	err = cs.categoryRepo.UpdateCategory(&entity.Category{
		Id:        categoryId,
		Name:      name,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return err
	}

	return nil
}

func (cs *CategoryService) DeleteCategory(categoryId int) error {
	_, err := cs.categoryRepo.GetDetailCategory(categoryId)
	if err != nil {
		return err
	}

	err = cs.categoryRepo.DeleteCategory(categoryId)
	if err != nil {
		return err
	}

	return nil
}
