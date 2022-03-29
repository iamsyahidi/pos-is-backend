package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Cashier struct {
	Id        int        `gorm:"primary_key;auto_increment;"  json:"id"`
	Name      string     `gorm:"size:100;not null;" json:"name"`
	Passcode  string     `gorm:"size:100;not null;" json:"passcode"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}

type CashierView struct {
	CashierId int    `json:"cashierId"`
	Name      string `json:"name"`
}

type CashierRequest struct {
	Name string `json:"name" validate:"required"`
}

type CashierResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Passcode  string    `json:"passcode"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ListCashierResponse struct {
	Cashiers []CashierView `json:"cashiers"`
	Meta     MetaView      `json:"meta"`
}

var validate = validator.New()

func (c *CashierRequest) Validate() error {
	validationErr := validate.Struct(c)
	return validationErr
}
