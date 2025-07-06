package models

import "time"

type ListVoteModel struct {
	ID         int
	VoteDate   time.Time
	ListNumber int
}
