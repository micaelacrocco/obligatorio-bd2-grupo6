package models

import "time"

type Candidate struct {
	CitizenID     int
	ListNumber    int
	StartDate     time.Time
	EndDate       time.Time
	CandidacyType string
}
