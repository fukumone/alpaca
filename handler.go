package main

import (
	"net/http"
	"html/template"

	"github.com/t-fukui/alpaca/models"
)

var tpl *template.Template

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	Messages := [] models.Message{}
	db.Find(&Messages)
	tpl = template.Must(template.ParseFiles("templates/index.html"))
	tpl.Execute(w, Messages)
}
