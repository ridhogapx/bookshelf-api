package config

import (
	"bookshelf-api/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn string = "root:@tcp(127.0.0.1:3306)/bookshelf?charset=utf8mb4&parseTime=True&loc=Local"
var DB *gorm.DB

func init() {
	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database!")
	}

	DB.AutoMigrate(&model.User{})
}

func Msg() string {
	return "Migrating database..."
}
