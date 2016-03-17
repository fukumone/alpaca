package main

import (
	"net/http"
	"html/template"
	"github.com/t-fukui/alpaca/models"
)

type FormData struct {
	Message models.Message
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	Messages := []models.Message{}
	db.Debug().Find(&Messages)
	tpl := template.Must(template.ParseFiles("templates/index.html"))
	tpl.Execute(w, &Messages)
}

func NewHandler(w http.ResponseWriter, r *http.Request){
	tpl := template.Must(template.ParseFiles("templates/new.html"))
	tpl.Execute(w, FormData{models.Message{}})
}
