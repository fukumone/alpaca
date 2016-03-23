package models

import (
	"github.com/jinzhu/gorm"
	"github.com/wcl48/valval"
)

type Title struct {
	gorm.Model
	Messages  []Message
	Name      string `sql:"size:255"`
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
