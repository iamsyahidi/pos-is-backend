package entity

import (
	"time"
)

type Category struct {
	Id        int        `gorm:"primary_key;auto_increment;"  json:"id"`
	Name      string     `gorm:"size:100;not null;" json:"name"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}

type CategoryRequest struct {
	Name string `json:"name"`
}

type CategoryResponses struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CategoryView struct {
	CategoryId int    `json:"categoryId"`
	Name       string `json:"name"`
}

type ListCategoryResponse struct {
	Categories []CategoryView `json:"categories"`
	Meta       MetaView       `json:"meta"`
}

func (c *CategoryRequest) Validate() error {
	validationErr := validate.Struct(c)
	return validationErr
}
