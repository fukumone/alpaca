package main

import (
	"fmt"
	"net/http"
	"html/template"

	"github.com/wcl48/valval"
	"github.com/t-fukui/alpaca/models"
)

type FormData struct {
	Message models.Message
	Mess string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	Messages := []models.Message{}
	db.Debug().Find(&Messages)
	tpl := template.Must(template.ParseFiles("templates/index.html"))
	tpl.Execute(w, &Messages)
}

func NewHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/new.html"))
	tpl.Execute(w, FormData{models.Message{}, ""})
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	Message := models.Message{Name: r.FormValue("Name"), Title: r.FormValue("Title"), Body: r.FormValue("Body")}
	if err := models.MessageValidate(Message); err != nil {
		var Mess string
		errs := valval.Errors(err)
		for _, errInfo := range errs {
			Mess += fmt.Sprint(errInfo.Error)
		}
		tpl := template.Must(template.ParseFiles("templates/new.html"))
		tpl.Execute(w, FormData{Message, Mess})
	} else {
		db.Create(&Message)
		http.Redirect(w, r, "/index", 301)
	}
}
