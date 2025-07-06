package dtos

type CandidateDto struct {
	CitizenID     int    `json:"citizen_id"`
	ListNumber    int    `json:"list_number"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
	CandidacyType string `json:"candidacy_type"`
}
