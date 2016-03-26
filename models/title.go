package models

import (
	"time"
	"github.com/wcl48/valval"
)

type Title struct {
	ID   int
	Messages  []Message
	Name      string `sql:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func TitleValidate(title Title) error {
	Validator := valval.Object(valval.M{
		"Name": valval.String(
			valval.MinLength(1),
			valval.MaxLength(20),
		),
	})

	return Validator.Validate(title)
}
