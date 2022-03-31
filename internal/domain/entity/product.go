package entity

import "time"

type Product struct {
	Id           int        `gorm:"primary_key;auto_increment;"  json:"id"`
	CategoriesId int        `gorm:"not null;" json:"categoriesId"`
	Name         string     `gorm:"size:100;not null;" json:"name"`
	Sku          string     `gorm:"size:100;null" json:"sku"`
	Image        string     `gorm:"size:255;null" json:"image"`
	Price        int        `gorm:"null;"  json:"price"`
	Stock        int        `gorm:"null;"  json:"stock"`
	CreatedAt    time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt    time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeletedAt    *time.Time `json:"deletedAt,omitempty"`
}

type ProductRequest struct {
	CategoryId int    `json:"categoryId"`
	Name       string `json:"name"`
	Sku        string `json:"sku"`
	Image      string `json:"image"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
}

type ProductResponse struct {
	Id           int       `json:"id"`
	CategoriesId int       `json:"categoriesId"`
	Name         string    `json:"name"`
	Sku          string    `json:"sku"`
	Image        string    `json:"image"`
	Price        int       `json:"price"`
	Stock        int       `json:"stock"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type ProductView struct {
	ProductId int          `json:"productId"`
	Category  CategoryView `json:"category"`
	Name      string       `json:"name"`
	Sku       string       `json:"sku"`
	Image     string       `json:"image"`
	Price     int          `json:"price"`
	Stock     int          `json:"stock"`
	Discount  *Discount    `json:"discount"`
}

type ListProductResponse struct {
	Products []ProductView `json:"products"`
	Meta     MetaView      `json:"meta"`
}
