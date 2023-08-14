package model

import "time"

type InterviewBook struct {
	BaseModel
	InterviewDate        time.Time             `json:"interviewDate"`
	Interviewer          Interviewer           `json:"interviewer"`
	Recruiter            Recruiter             `json:"recruiter"`
	BootcampSource       BootcampSource        `json:"bootcampSource"`
	InterviewBookDetails []InterviewBookDetail `json:"interviewBookDetails"`
}

type InterviewBookDetail struct {
	BaseModel
	InterviewBookId   string `json:"interviewBookId"`
	InterviewTime     time.Time
	CandidateResume   CandidateResume `json:"CandidateResume"`
	InterviewFile     string          `json:"interviewFile"`
	MeetingLink       string          `json:"meetingLink"`
	InterviewStatusId string          `json:"interviewStatusId"`
}
