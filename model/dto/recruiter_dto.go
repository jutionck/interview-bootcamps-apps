package dto

type RecruiterRequestDto struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	UserId string `json:"userId"`
}
