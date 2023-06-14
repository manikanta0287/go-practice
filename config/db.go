package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var dsn *gorm.DB

func Database() {

	dsn := "root:Mani*123@tcp(127.0.0.1:3306)/testingdb?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("database connected successfully...")
	}
}
