package models

import "time"

type ListVote struct {
	ID         int
	VoteDate   time.Time
	ListNumber int
	CircuitID  int
}
