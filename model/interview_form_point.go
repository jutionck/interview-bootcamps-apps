package model

type InterviewFormPoint struct {
	BaseModel
	InterviewFormId string `json:"interviewFormId"`
	Point           int    `json:"point"`
	Title           string `json:"title"`
}
