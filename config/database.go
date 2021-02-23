package config

import (
	"fmt"
	"goorder/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	dsn := "root:root@tcp(127.0.0.1:3306)/gotoko?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&models.Order{}, &models.Item{})
	fmt.Println("Connected to Database")

	DB = db
}
