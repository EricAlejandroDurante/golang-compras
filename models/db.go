package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("mysql", "root:1Q2w3e4r5t6y@tcp(127.0.0.1:3306)/tarea_1_sd")
	if err != nil {
		panic(err)
	}
	DB = database
}
