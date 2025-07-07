package dtos

type CircuitResultDto struct {
	List       string  `json:"list"`
	PartyName  string  `json:"party_name"`
	VoteCount  int     `json:"vote_count"`
	Percentage float64 `json:"percentage"`
}
