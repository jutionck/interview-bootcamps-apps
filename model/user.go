package model

import validation "github.com/go-ozzo/ozzo-validation"

type User struct {
	BaseModel
	Username   string `json:"username"`
	Password   string `json:"password,omitempty"`
	Role       string `json:"role"`
	ResetToken string `json:"resetToken,omitempty"`
	IsActive   bool   `json:"isActive"`
}

func (u User) IsValidRole() bool {
	return u.Role == "admin" || u.Role == "recruiter" || u.Role == "interviewer"
}

func (u User) ValidateRequire() error {
	return validation.ValidateStruct(
		&u,
		validation.Field(&u.Username, validation.Required),
		validation.Field(&u.Password, validation.Required),
		validation.Field(&u.Role, validation.Required),
	)
}
