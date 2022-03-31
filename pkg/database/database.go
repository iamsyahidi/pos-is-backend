package database

import (
	"fmt"
	"os"
	"pos-is-backend/internal/domain/entity"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitMySQL() (*gorm.DB, error) {
	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbname := os.Getenv("MYSQL_DBNAME")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.DB().Ping(); err != nil {
		return nil, err
	}
	db.LogMode(true)

	//Drop table if exists
	if err := db.DropTableIfExists(&entity.Cashier{}, &entity.Category{}, &entity.Product{}).Error; err != nil {
		return nil, err
	}

	//Create table
	if err := db.AutoMigrate(&entity.Cashier{}, &entity.Category{}, &entity.Product{}).Error; err != nil {
		return nil, err
	}

	cashier1 := entity.Cashier{
		Name:     "Kasir 10",
		Passcode: "123456",
	}
	cashier2 := entity.Cashier{
		Name:     "Kasir 12",
		Passcode: "123456",
	}
	cashier3 := entity.Cashier{
		Name:     "Kasir 14",
		Passcode: "123456",
	}

	category1 := entity.Category{
		Name: "Kategori 1",
	}
	category2 := entity.Category{
		Name: "Kategori 2",
	}
	category3 := entity.Category{
		Name: "Kategori 3",
	}

	product1 := entity.Product{
		CategoriesId: 1,
		Name:         "Lifebuoy Body Foam",
		Sku:          "ID001",
		Image:        "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2020/11/26/001f1c6e-d068-484f-9333-c3fa4129ef26.jpg",
		Price:        25000,
		Stock:        4,
	}
	product2 := entity.Product{
		CategoriesId: 2,
		Name:         "Lifebuoy Body Foam2",
		Sku:          "ID002",
		Image:        "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2020/11/26/001f1c6e-d068-484f-9333-c3fa4129ef26.jpg",
		Price:        25001,
		Stock:        4,
	}
	product3 := entity.Product{
		CategoriesId: 3,
		Name:         "Lifebuoy Body Foam3",
		Sku:          "ID001",
		Image:        "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2020/11/26/001f1c6e-d068-484f-9333-c3fa4129ef26.jpg",
		Price:        25002,
		Stock:        4,
	}
	product4 := entity.Product{
		CategoriesId: 1,
		Name:         "Lifebuoy Body Foam4",
		Sku:          "ID002",
		Image:        "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2020/11/26/001f1c6e-d068-484f-9333-c3fa4129ef26.jpg",
		Price:        25003,
		Stock:        4,
	}
	product5 := entity.Product{
		CategoriesId: 2,
		Name:         "Lifebuoy Body Foam5",
		Sku:          "ID001",
		Image:        "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2020/11/26/001f1c6e-d068-484f-9333-c3fa4129ef26.jpg",
		Price:        25004,
		Stock:        4,
	}
	product6 := entity.Product{
		CategoriesId: 3,
		Name:         "Lifebuoy Body Foam6",
		Sku:          "ID002",
		Image:        "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2020/11/26/001f1c6e-d068-484f-9333-c3fa4129ef26.jpg",
		Price:        25005,
		Stock:        4,
	}

	db.Create(&cashier1)
	db.Create(&cashier2)
	db.Create(&cashier3)
	db.Create(&category1)
	db.Create(&category2)
	db.Create(&category3)
	db.Create(&product1)
	db.Create(&product2)
	db.Create(&product3)
	db.Create(&product4)
	db.Create(&product5)
	db.Create(&product6)

	return db, nil
}
