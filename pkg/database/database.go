package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

func InitMySQL() (*gorm.DB, error) {
	username := viper.GetString("MYSQL_USER")
	password := viper.GetString("MYSQL_PASSWORD")
	host := viper.GetString("MYSQL_HOST")
	port := viper.GetInt("MYSQL_PORT")
	dbname := viper.GetString("MYSQL_DBNAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.DB().Ping(); err != nil {
		return nil, err
	}
	db.LogMode(true)

	// //Drop table if exists
	// if err := db.DropTableIfExists(&entity.User{}, &entity.Role{}, &entity.Post{}).Error; err != nil {
	// 	return nil, err
	// }

	// //Create table
	// if err := db.AutoMigrate(&entity.Role{}, &entity.User{}, &entity.Post{}).Error; err != nil {
	// 	return nil, err
	// }

	// if err := db.Model(&entity.User{}).AddForeignKey("role_id", "roles(id)", "CASCADE", "CASCADE").Error; err != nil {
	// 	return nil, err
	// }

	// if err := db.Model(&entity.Post{}).AddForeignKey("author_id", "users(id)", "CASCADE", "CASCADE").Error; err != nil {
	// 	return nil, err
	// }

	return db, nil
}
