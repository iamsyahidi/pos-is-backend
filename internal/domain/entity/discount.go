package entity

import "time"

type Discount struct {
	Id              int        `gorm:"primary_key;auto_increment;" json:"id"`
	Qty             int        `gorm:"null;" json:"qty"`
	Type            string     `gorm:"null;" json:"type"`
	Result          string     `gorm:"null;" json:"result"`
	ExpiredAt       time.Time  `gorm:"null;" json:"expiredAt"`
	ExpiredAtFormat time.Time  `gorm:"null;" json:"expiredAtFormat"`
	StringFormat    string     `gorm:"null;" json:"stringFormat"`
	CreatedAt       time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt       time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeletedAt       *time.Time `json:"deletedAt,omitempty"`
}
