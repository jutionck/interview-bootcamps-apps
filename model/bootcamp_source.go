package model

import validation "github.com/go-ozzo/ozzo-validation"

type BootcampSource struct {
	BaseModel
	Name string `json:"name"`
}

func (u BootcampSource) ValidateRequire() error {
	return validation.ValidateStruct(
		&u,
		validation.Field(&u.Name, validation.Required),
	)
}
