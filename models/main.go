package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Password string
}

type Message struct {
	gorm.Model
	Content string
	UserID  int
	User    User
}

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:root@tcp(127.0.0.1:3306)/gocrud?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect db")
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Message{})

	DB = db
}
