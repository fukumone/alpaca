package models

import (
	"time"
	"github.com/wcl48/valval"
)

type Message struct {
	ID   int
	Title     Title
	TitleId   int64
	Name      string `sql:"size:255"`
	Body      string `sql:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
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
