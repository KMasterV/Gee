package common

import (
	"Gee/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	// driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "gorm"
	username := "root"
	password := "root"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)

	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("faild to connect database, err:" + err.Error())
	}
	db.AutoMigrate(&model.User{})
	DB = db
	return DB
}

// func GetDB() *gorm.DB {
// 	return DB
// }
