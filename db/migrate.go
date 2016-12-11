package main

import (
	"github.com/fukumone/alpaca/config"
	"github.com/fukumone/alpaca/models"
)

func main(){
	db := config.Database()
	db.CreateTable(&models.Title{})
	db.CreateTable(&models.Message{})
}
