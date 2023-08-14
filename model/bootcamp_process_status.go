package model

import validation "github.com/go-ozzo/ozzo-validation"

type BootcampProcessStatus struct {
	BaseModel
	Name   string `json:"name"`
	Status string `json:"status"`
}

func (b BootcampProcessStatus) IsValidStatus() bool {
	return b.Status == "schedule" || b.Status == "interview"
}

func (b BootcampProcessStatus) ValidateRequire() error {
	return validation.ValidateStruct(
		&b,
		validation.Field(&b.Name, validation.Required),
		validation.Field(&b.Status, validation.Required),
	)
}
