package model

import validation "github.com/go-ozzo/ozzo-validation"

type Interviewer struct {
	BaseModel
	Name  string `json:"name"`
	Email string `json:"email"`
	User  User   `json:"user"`
}

func (u Interviewer) ValidateRequire() error {
	return validation.ValidateStruct(
		&u,
		validation.Field(&u.Name, validation.Required),
		validation.Field(&u.Email, validation.Required),
	)
}
