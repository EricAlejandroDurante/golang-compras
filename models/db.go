package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("mysql", "admin:12345678@tcp(127.0.0.1:3306)/tarea_1_sd")
	if err != nil {
		panic(err)
	}
	DB = database
}
