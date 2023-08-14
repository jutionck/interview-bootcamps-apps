package model

import validation "github.com/go-ozzo/ozzo-validation"

type University struct {
	BaseModel
	Name string `json:"name"`
}

func (u University) ValidateRequire() error {
	return validation.ValidateStruct(
		&u,
		validation.Field(&u.Name, validation.Required),
	)
}
