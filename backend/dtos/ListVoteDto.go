package dtos

type ListVoteDto struct {
	ID         int    `json:"id"`
	VoteDate   string `json:"vote_date"` // formato: "YYYY-MM-DD"
	ListNumber int    `json:"list_number"`
}
