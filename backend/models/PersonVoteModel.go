package models

type PersonVoteModel struct {
	ID         int
	VoteDate   string
	IsObserved bool
	VoteType   string
	CitizenID  int
	CircuitID  int
}
