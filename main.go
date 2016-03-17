package main

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/t-fukui/alpaca/config"
)

var db gorm.DB

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/new", NewHandler)

	// Webサーバーを起動
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func init() {
	db = config.Database()
}
