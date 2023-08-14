package model

type InterviewForm struct {
	BaseModel
	Title       string `json:"title"`
	Description string `json:"description"`
}
