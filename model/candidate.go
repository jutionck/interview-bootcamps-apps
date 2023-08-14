package model

import validation "github.com/go-ozzo/ozzo-validation"

type Candidate struct {
	BaseModel
	Name string `json:"name"`
}

func (u Candidate) ValidateRequire() error {
	return validation.ValidateStruct(
		&u,
		validation.Field(&u.Name, validation.Required),
	)
}
