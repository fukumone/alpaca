package main

import (
	"net/http"
	"html/template"
	"github.com/t-fukui/alpaca/models"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	Messages := []models.Message{}
	db.Debug().Find(&Messages)
	tpl := template.Must(template.ParseFiles("templates/index.html"))
	tpl.Execute(w, &Messages)
}
