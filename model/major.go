package model

import validation "github.com/go-ozzo/ozzo-validation"

type Major struct {
	BaseModel
	Name string `json:"name"`
}

func (u Major) ValidateRequire() error {
	return validation.ValidateStruct(
		&u,
		validation.Field(&u.Name, validation.Required),
	)
}
