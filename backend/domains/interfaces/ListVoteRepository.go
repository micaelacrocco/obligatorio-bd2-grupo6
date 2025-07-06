package interfaces

import "EleccionesUcu/models"

type ListVoteRepository interface {
	GetAll() ([]models.ListVoteModel, error)
	Add(vote models.ListVoteModel) (*models.ListVoteModel, error)
	Update(vote models.ListVoteModel) (*models.ListVoteModel, error)
	Delete(id int) error
}
