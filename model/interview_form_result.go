package model

type InterviewFormResult struct {
	BaseModel
	InterviewFormId      string `json:"interviewFormId"`
	InterviewFormPointId string `json:"interviewFormPointId"`
	Note                 string `json:"note"`
	CandidateResumeId    string `json:"CandidateResumeId"`
	InterviewerId        string `json:"interviewerId"`
	Result               string `json:"result"`
	InterviewerNote      string `json:"interviewerNote"`
}
