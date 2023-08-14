package dto

type InterviewerRequestDto struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	UserId string `json:"userId"`
}
