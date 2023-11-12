package main

import (
	"go_crud/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root@tcp(127.0.0.1:3306)/go_crud?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.User{})

	// Create
	db.Create(&model.User{Fname: "Tae", Lname: "ToodDum", Username: "tanuson@toodDum.com", Avatar: "https://www.melivecode.com/users/1.png"})
	db.Create(&model.User{Fname: "Cat", Lname: "Meow", Username: "Catmeow@toodDum.com", Avatar: "https://www.melivecode.com/users/2.png"})
}
