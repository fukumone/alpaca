package main

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/t-fukui/alpaca/config"
)

var db gorm.DB

func main() {
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/new", NewHandler)
	http.HandleFunc("/create", CreateHandler)
	// http.Handle("/edit", &EditHandler{})

	// Webサーバーを起動
	if err := http.ListenAndServe(":3000", Log(http.DefaultServeMux)); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func init() {
	db = config.Database()
}
