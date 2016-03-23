package models

import (
	"github.com/jinzhu/gorm"
	"github.com/wcl48/valval"
)

type Message struct {
	gorm.Model
	Title     Title
	TitleId   int64
	Name      string `sql:"size:255"`
	Body      string `sql:"size:255"`
}

func MessageValidate(message Message) error {
	Validator := valval.Object(valval.M{
		"Name": valval.String(
			valval.MinLength(1),
			valval.MaxLength(20),
		),
	})

	return Validator.Validate(message)
}
