package dtos

type PartyVoteDto struct {
	PartyName  string  `json:"party"`
	Votes      int     `json:"votes"`
	Percentage float64 `json:"percentage,omitempty"`
}
