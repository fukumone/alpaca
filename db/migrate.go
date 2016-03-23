package main

import (
	"github.com/t-fukui/alpaca/config"
	"github.com/t-fukui/alpaca/models"
)

func main(){
	db := config.Database()
	db.CreateTable(&models.Title{})
	db.CreateTable(&models.Message{})
}
