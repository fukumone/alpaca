package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"models"
)

func main(){
	db := config.Database()
	db.CreateTable(&models.Message{})
}
