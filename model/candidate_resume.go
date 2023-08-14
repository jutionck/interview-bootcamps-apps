package model

type CandidateResume struct {
	BaseModel
	Candidate       Candidate  `json:"candidate"`
	University      University `json:"university"`
	Major           Major      `json:"major"`
	HackerrankPoint int        `json:"hackerrankPoint"`
}
