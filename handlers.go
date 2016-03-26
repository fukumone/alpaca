package main

import (
	"fmt"
	"net/http"
	"strings"
	"strconv"
	"text/template"

	"github.com/t-fukui/alpaca/models"
	"github.com/wcl48/valval"
)

var tpl *template.Template

type TitleFormData struct {
	Title models.Title
	Mess    string
}

//////////////////////////
///// Titleアクション /////
//////////////////////////
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	Titles := []models.Title{}
	db.Debug().Find(&Titles)
	tpl = template.Must(template.ParseFiles("templates/layout.html", "templates/index.html"))
	tpl.Execute(w, &Titles)
}

func NewHandler(w http.ResponseWriter, r *http.Request) {
	tpl = template.Must(template.ParseFiles("templates/layout.html", "templates/new.html"))
	tpl.Execute(w, TitleFormData{models.Title{}, ""})
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	Title := models.Title{Name: r.FormValue("Name")}

	if err := models.TitleValidate(Title); err != nil {
		var Mess string
		errs := valval.Errors(err)
		for _, errInfo := range errs {
			Mess += fmt.Sprint(errInfo.Error)
		}
		tpl = template.Must(template.ParseFiles("templates/layout.html", "templates/new.html"))
		tpl.Execute(w, TitleFormData{Title, Mess})
	} else {
		db.Debug().Create(&Title)
		http.Redirect(w, r, "/", 301)
	}
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	Title := models.Title{}
	id := strings.Split(r.URL.Path, "/")[2]
	db.Debug().First(&Title, id)
	tpl = template.Must(template.ParseFiles("templates/layout.html", "templates/edit.html"))
	tpl.Execute(w, TitleFormData{Title, ""})
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	Title := models.Title{}
	id := strings.Split(r.URL.Path, "/")[2]
	db.Debug().First(&Title, id)
	Title.Name = r.FormValue("Name")
	if err := models.TitleValidate(Title); err != nil {
		var Mess string
		errs := valval.Errors(err)
		for _, errInfo := range errs {
			Mess += fmt.Sprint(errInfo.Error)
		}
		tpl = template.Must(template.ParseFiles("templates/layout.html", "templates/edit.html"))
		tpl.Execute(w, TitleFormData{Title, Mess})
	} else {
		db.Debug().Save(&Title)
		http.Redirect(w, r, "/", 301)
	}
}

////////////////////////////
///// Messageアクション /////
////////////////////////////
type MessageFormData struct {
	Message models.Message
	Title   models.Title
	Mess    string
}

type MessageIndexData struct {
	Messages []models.Message
	Title   models.Title
}

func MessagesIndexHandler(w http.ResponseWriter, r *http.Request) {
	Title := models.Title{}
	TitleID := strings.Split(r.URL.Path, "/")[2]
	db.Debug().First(&Title, TitleID)

	Messages := []models.Message{}
	db.Debug().Where("title_id = ?", TitleID).Find(&Messages)
	tpl = template.Must(template.ParseFiles("templates/layout.html", "templates/messages/index.html"))
	tpl.Execute(w, MessageIndexData{Messages, Title})
}

func MessageNewHandler(w http.ResponseWriter, r *http.Request) {
	Title := models.Title{}
	Titleid := strings.Split(r.URL.Path, "/")[2]
	db.Debug().First(&Title, Titleid)

	tpl = template.Must(template.ParseFiles("templates/layout.html", "templates/messages/new.html"))
	tpl.Execute(w, MessageFormData{models.Message{}, Title, ""})
}

func MessageCreateHandler(w http.ResponseWriter, r *http.Request) {
	Title := models.Title{}
	id := strings.Split(r.URL.Path, "/")[2]
	db.Debug().First(&Title, id)

	TitleId, _ := strconv.Atoi(id)
	Message := models.Message{TitleId: int64(TitleId),
		Name: r.FormValue("Name"),
		Body: r.FormValue("Body")}

	if err := models.MessageValidate(Message); err != nil {
		var Mess string
		errs := valval.Errors(err)
		for _, errInfo := range errs {
			Mess += fmt.Sprint(errInfo.Error)
		}
		tpl = template.Must(template.ParseFiles("templates/layout.html", "templates/messages/new.html"))
		tpl.Execute(w, MessageFormData{Message, Title, Mess})
	} else {
		db.Debug().Create(&Message)
		path := fmt.Sprintf("/title/%s/messages", id)
		http.Redirect(w, r, path, 301)
	}
}

func MessageEditHandler(w http.ResponseWriter, r *http.Request) {
	Title := models.Title{}
	TitleID := strings.Split(r.URL.Path, "/")[2]
	db.Debug().First(&Title, TitleID)

	Message := models.Message{}
	MessageID := strings.Split(r.URL.Path, "/")[5]
	db.Debug().First(&Message, MessageID)

	tpl = template.Must(template.ParseFiles("templates/layout.html", "templates/messages/edit.html"))
	tpl.Execute(w, MessageFormData{Message, Title, ""})
}

func MessageUpdateHandler(w http.ResponseWriter, r *http.Request) {
	Title := models.Title{}
	TitleID := strings.Split(r.URL.Path, "/")[2]
	db.Debug().First(&Title, TitleID)

	Message := models.Message{}
	MessageID := strings.Split(r.URL.Path, "/")[5]
	db.Debug().First(&Message, MessageID)

	Message.Name = r.FormValue("Name")
	Message.Body = r.FormValue("Body")

	if err := models.MessageValidate(Message); err != nil {
		var Mess string
		errs := valval.Errors(err)
		for _, errInfo := range errs {
			Mess += fmt.Sprint(errInfo.Error)
		}
		tpl = template.Must(template.ParseFiles("templates/layout.html", "templates/messages/edit.html"))
		tpl.Execute(w, MessageFormData{Message, Title, Mess})
	} else {
		db.Debug().Save(&Message)
		path := fmt.Sprintf("/title/%s/messages", TitleID)
		http.Redirect(w, r, path, 301)
	}
}
