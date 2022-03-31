package repository

import (
	"pos-is-backend/internal/domain/entity"
	"strings"

	"github.com/jinzhu/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

type IProductRepository interface {
	CreateProduct(*entity.Product) (*entity.Product, error)
	GetAllProducts(limit, skip int, categoryId int, q string) (*entity.ListProductResponse, error)
	GetProductByID(productID int) (*entity.Product, *entity.CategoryView, error)
	UpdateProduct(product *entity.Product) error
	DeleteProduct(productID int) error
	TotalProduct(categoryId int, q string) (int, error)
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (p *ProductRepository) CreateProduct(product *entity.Product) (*entity.Product, error) {
	tx := p.db.Begin()
	if err := tx.Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	err := tx.Debug().Model(&entity.Product{}).Create(&product).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return product, nil
}

func (p *ProductRepository) GetAllProducts(limit, skip int, categoryId int, q string) (*entity.ListProductResponse, error) {
	var err error
	products := []entity.Product{}
	category := entity.CategoryView{}
	res := entity.ListProductResponse{}
	pView := []entity.ProductView{}

	if limit > 0 {
		if categoryId != 0 {
			err = p.db.Debug().Model(&entity.Product{}).Limit(limit).Offset(skip).Where("lower(name) like ? and categories_id = ?", "%"+strings.ToLower(q)+"%", categoryId).Find(&products).Error
			if err != nil {
				return nil, err
			}
		} else {
			err = p.db.Debug().Model(&entity.Product{}).Limit(limit).Offset(skip).Where("lower(name) like ?", "%"+strings.ToLower(q)+"%").Find(&products).Error
			if err != nil {
				return nil, err
			}
		}
	} else {
		if categoryId != 0 {
			err = p.db.Debug().Model(&entity.Product{}).Where("lower(name) like ? and categories_id = ?", "%"+strings.ToLower(q)+"%", categoryId).Find(&products).Error
			if err != nil {
				return nil, err
			}
		} else {
			err = p.db.Debug().Model(&entity.Product{}).Where("lower(name) like ?", "%"+strings.ToLower(q)+"%").Find(&products).Error
			if err != nil {
				return nil, err
			}
		}
	}

	if len(products) > 0 {
		for _, v := range products {
			err = p.db.Debug().Model(&entity.Category{}).Where("id = ?", v.CategoriesId).Select("id, name").Scan(&category).Error
			if err != nil {
				return nil, err
			}
			pView = append(pView, entity.ProductView{
				ProductId: v.Id,
				Category: entity.CategoryView{
					CategoryId: category.CategoryId,
					Name:       category.Name,
				},
				Name:     v.Name,
				Sku:      v.Sku,
				Image:    v.Image,
				Price:    v.Price,
				Stock:    v.Stock,
				Discount: nil,
			})
		}
	}

	res.Products = pView

	return &res, nil
}

func (p *ProductRepository) GetProductByID(productID int) (*entity.Product, *entity.CategoryView, error) {
	var (
		err      error
		product  entity.Product
		category entity.CategoryView
	)

	err = p.db.Debug().Model(&entity.Product{}).Where("id = ?", productID).Take(&product).Error
	if err != nil {
		return nil, nil, err
	}
	if product.Id != 0 {
		err = p.db.Debug().Model(&entity.Category{}).Where("id = ?", product.CategoriesId).Select("id, name").Scan(&category).Error
		if err != nil {
			return nil, nil, err
		}
	}
	return &product, &category, nil
}

func (p *ProductRepository) UpdateProduct(product *entity.Product) error {
	var err error

	tx := p.db.Begin()
	if err = tx.Error; err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Debug().Model(&entity.Product{}).Save(&product).Error
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

func (p *ProductRepository) DeleteProduct(productID int) error {
	var err error = nil

	err = p.db.Debug().Model(&entity.Product{}).Where("id = ?", productID).Delete(&entity.Product{}).Error
	if err != nil {
		return err
	}
	return nil
}

// TotalProduct() (int, error) ...
func (p *ProductRepository) TotalProduct(categoryId int, q string) (int, error) {
	var (
		err   error
		total int
	)

	if categoryId != 0 {
		err = p.db.Debug().Model(&entity.Product{}).Where("lower(name) like ? and categories_id = ? ", "%"+strings.ToLower(q)+"%", categoryId).Count(&total).Error
		if err != nil {
			return 0, err
		}
	} else {
		err = p.db.Debug().Model(&entity.Product{}).Where("lower(name) like ?", "%"+strings.ToLower(q)+"%").Count(&total).Error
		if err != nil {
			return 0, err
		}
	}

	return total, nil
}
