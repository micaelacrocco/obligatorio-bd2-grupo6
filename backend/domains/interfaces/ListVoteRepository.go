package interfaces

import "EleccionesUcu/models"

type ListVoteRepository interface {
	GetAll() ([]models.ListVote, error)
	Add(vote models.ListVote) (*models.ListVote, error)
	Update(vote models.ListVote) (*models.ListVote, error)
	Delete(id int) error
}
