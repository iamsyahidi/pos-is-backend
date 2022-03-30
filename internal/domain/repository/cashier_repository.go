package repository

import (
	"pos-is-backend/internal/domain/entity"

	"github.com/jinzhu/gorm"
)

type CashierRepository struct {
	db *gorm.DB
}

type ICashierRepository interface {
	CreateCashier(request *entity.Cashier) (*entity.Cashier, error)
	GetAllCashier(limit, skip int) (*[]entity.Cashier, error)
	GetDetailCashier(cashierId int) (*entity.Cashier, error)
	UpdateCashier(cashierId int, name string) error
	DeleteCashier(cashierId int, name string) error
	TotalCashier() (int, error)
}

// NewCashierRepository(db *gorm.DB) *CashierRepository ...
func NewCashierRepository(db *gorm.DB) *CashierRepository {
	return &CashierRepository{
		db: db,
	}
}

// CreateCashier(request *entity.Cashier) (*entity.Cashier, error) ...
func (cr *CashierRepository) CreateCashier(request *entity.Cashier) (*entity.Cashier, error) {
	tx := cr.db.Begin()
	if err := tx.Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	err := tx.Debug().Model(&entity.Cashier{}).Create(request).Error
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

// GetAllCashier(limit, skip int) (*[]entity.Cashier, error) ...
func (cr *CashierRepository) GetAllCashier(limit, skip int) (*[]entity.Cashier, error) {
	cashiers := []entity.Cashier{}

	err := cr.db.Debug().Model(&entity.Cashier{}).Offset(skip).Limit(limit).Find(&cashiers).Error
	if err != nil {
		return nil, err
	}

	return &cashiers, err
}

// GetDetailCashier(cashierId int) (*entity.Cashier, error) ...
func (cr *CashierRepository) GetDetailCashier(cashierId int) (*entity.Cashier, error) {
	var (
		err     error
		cashier entity.Cashier
	)

	err = cr.db.Debug().Model(&entity.Cashier{}).Where("id = ?", cashierId).Take(&cashier).Error
	if err != nil {
		return nil, err
	}

	return &cashier, err
}

// UpdateCashier(cashierId int, name string) error ...
func (cr *CashierRepository) UpdateCashier(cashierId int, name string) error {
	var err error

	tx := cr.db.Begin()
	if err = tx.Error; err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Debug().Model(&entity.Cashier{}).Where("id = ?", cashierId).Updates(entity.Cashier{
		Name: name,
	}).Error
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

// DeleteCashier(cashierId int, name string) error ...
func (cr *CashierRepository) DeleteCashier(cashierId int, name string) error {
	var err error = nil

	err = cr.db.Debug().Model(&entity.Cashier{}).Where("id = ? and name = ?", cashierId, name).Delete(&entity.Cashier{}).Error
	if err != nil {
		return err
	}
	return nil
}

// TotalCashier() (int, error) ...
func (cr *CashierRepository) TotalCashier() (int, error) {
	var (
		err   error
		total int
	)

	err = cr.db.Debug().Model(&entity.Cashier{}).Count(&total).Error
	if err != nil {
		return 0, err
	}

	return total, nil
}
