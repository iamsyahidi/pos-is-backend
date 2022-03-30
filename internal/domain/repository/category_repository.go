package repository

import (
	"pos-is-backend/internal/domain/entity"

	"github.com/jinzhu/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

type ICategoryRepository interface {
	CreateCategory(request *entity.Category) (*entity.Category, error)
	GetAllCategory(limit, skip int) (*[]entity.Category, error)
	GetDetailCategory(categoryId int) (*entity.Category, error)
	UpdateCategory(request *entity.Category) error
	DeleteCategory(categoryId int) error
	TotalCategory() (int, error)
}

// NewCategoryRepository(db *gorm.DB) *CategoryRepository ...
func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

// CreateCategory(request *entity.Category) (*entity.Category, error) ...
func (cr *CategoryRepository) CreateCategory(request *entity.Category) (*entity.Category, error) {
	tx := cr.db.Begin()
	if err := tx.Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	err := tx.Debug().Model(&entity.Category{}).Create(request).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return request, err
}

// GetAllCategory(limit, skip int) (*[]entity.Category, error) ...
func (cr *CategoryRepository) GetAllCategory(limit, skip int) (*[]entity.Category, error) {
	var err error
	categories := []entity.Category{}

	if limit > 0 {
		err = cr.db.Debug().Model(&entity.Category{}).Offset(skip).Limit(limit).Find(&categories).Error
		if err != nil {
			return nil, err
		}
	} else {
		err = cr.db.Debug().Model(&entity.Category{}).Find(&categories).Error
		if err != nil {
			return nil, err
		}
	}

	return &categories, err
}

// GetDetailCategory(categoryId int) (*entity.Category, error) ...
func (cr *CategoryRepository) GetDetailCategory(categoryId int) (*entity.Category, error) {
	var (
		err     error
		cashier entity.Category
	)

	err = cr.db.Debug().Model(&entity.Category{}).Where("id = ?", categoryId).Take(&cashier).Error
	if err != nil {
		return nil, err
	}

	return &cashier, err
}

// UpdateCategory(request *entity.Category) error ...
func (cr *CategoryRepository) UpdateCategory(request *entity.Category) error {
	var err error

	tx := cr.db.Begin()
	if err = tx.Error; err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Debug().Model(&entity.Category{}).Save(&request).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// DeleteCategory(categoryId int) error ...
func (cr *CategoryRepository) DeleteCategory(categoryId int) error {
	var err error = nil

	err = cr.db.Debug().Model(&entity.Category{}).Where("id = ?", categoryId).Delete(&entity.Category{}).Error
	if err != nil {
		return err
	}
	return nil
}

// TotalCategory() (int, error) ...
func (cr *CategoryRepository) TotalCategory() (int, error) {
	var (
		err   error
		total int
	)

	err = cr.db.Debug().Model(&entity.Category{}).Count(&total).Error
	if err != nil {
		return 0, err
	}

	return total, nil
}
