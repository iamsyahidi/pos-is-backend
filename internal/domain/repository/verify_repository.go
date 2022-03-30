package repository

import (
	"pos-is-backend/internal/domain/entity"

	"github.com/jinzhu/gorm"
)

type VerifyRepository struct {
	db *gorm.DB
}

type IVerifyRepository interface {
	LoginPasscode(cashierId int, passcode string) (int, error)
	GetPasscode(cashierId int) (*entity.Verify, error)
	LogoutPasscode(cashierId int, passcode string) (int, error)
}

// NewVerifyRepository(db *gorm.DB) *VerifyRepository ...
func NewVerifyRepository(db *gorm.DB) *VerifyRepository {
	return &VerifyRepository{
		db: db,
	}
}

// LoginPasscode(cashierId int, passcode string) error ...
func (cr *VerifyRepository) LoginPasscode(cashierId int, passcode string) (int, error) {
	var (
		err   error
		total int
	)

	err = cr.db.Debug().Model(&entity.Cashier{}).Where("id = ? and passcode = ?", cashierId, passcode).Count(&total).Error
	if err != nil {
		return 0, err
	}

	return total, nil
}

// GetPasscode(cashierId int) (*entity.Verify, error) ...
func (cr *VerifyRepository) GetPasscode(cashierId int) (*entity.Verify, error) {
	var verify = &entity.Verify{}

	err := cr.db.Debug().Model(&entity.Cashier{}).Where("id = ?", cashierId).Select("passcode").Scan(&verify).Error
	if err != nil {
		return nil, err
	}

	return verify, err
}

// LogoutPasscode(cashierId int, passcode string) (int, error) ...
func (cr *VerifyRepository) LogoutPasscode(cashierId int, passcode string) (int, error) {
	var (
		err   error
		total int
	)

	err = cr.db.Debug().Model(&entity.Cashier{}).Where("id = ? and passcode = ?", cashierId, passcode).Count(&total).Error
	if err != nil {
		return 0, err
	}

	return total, nil
}
