package dtos

type CircuitResultByAllCandidates struct {
	Party      string  `json:"party"`
	Candidate  string  `json:"candidate"`
	VoteCount  int     `json:"vote_count"`
	Percentage float64 `json:"percentage"`
}
