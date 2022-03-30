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
	if err := db.DropTableIfExists(&entity.Cashier{}).Error; err != nil {
		return nil, err
	}

	//Create table
	if err := db.AutoMigrate(&entity.Cashier{}).Error; err != nil {
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

	db.Create(&cashier1)
	db.Create(&cashier2)
	db.Create(&cashier3)

	return db, nil
}
