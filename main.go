package main

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/bmizerany/pat"
	"github.com/t-fukui/alpaca/config"
)

var db gorm.DB

func main() {
	css := http.FileServer(http.Dir("assets/css"))
	http.Handle("/assets/css/", http.StripPrefix("/assets/css/", css))
	js := http.FileServer(http.Dir("assets/js"))
	http.Handle("/assets/js/", http.StripPrefix("/assets/js/", js))

	mux := pat.New()
	mux.Get("/", http.HandlerFunc(IndexHandler))
	mux.Get("/new", http.HandlerFunc(NewHandler))
	mux.Post("/create", http.HandlerFunc(CreateHandler))
	mux.Get("/edit/:id", http.HandlerFunc(EditHandler))
	mux.Patch("/update/:id", http.HandlerFunc(UpdateHandler))

	http.Handle("/", mux)
	// Webサーバーを起動
	if err := http.ListenAndServe(":3000", Log(http.DefaultServeMux)); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func init() {
	db = config.Database()
}
