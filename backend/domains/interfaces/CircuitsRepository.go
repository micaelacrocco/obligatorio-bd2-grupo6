package interfaces

import (
	"EleccionesUcu/models"
)

type CircuitsRepository interface {
	GetAll() ([]models.Circuit, error)
	GetById(id int) (*models.Circuit, error)
	GetVotesByParty(circuitID int) ([]models.PartyVote, error)
	AddCircuit(circuit models.Circuit) (*models.Circuit, error)
	Update(c models.Circuit) (*models.Circuit, error)
	Delete(id int) error
}
