package dtos

type PersonVoteDTO struct {
	ID         int    `json:"id"`
	VoteDate   string `json:"vote_date"`
	IsObserved bool   `json:"is_observed"`
	VoteType   string `json:"vote_type"`
	CitizenID  int    `json:"citizen_id"`
	CircuitID  int    `json:"circuit_id"`
}
