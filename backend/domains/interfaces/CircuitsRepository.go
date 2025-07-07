package interfaces

import (
	"EleccionesUcu/models"
)

type CircuitsRepository interface {
	GetAll() ([]models.Circuit, error)
	GetById(id int) (*models.Circuit, error)
	GetVotesByParty(circuitID int) ([]models.PartyVote, error)
	GetVotes(circuitID int) ([]models.CircuitResult, error)
	GetVotesByAllCandidates(circuitID int) ([]models.CircuitResultByAllCandidates, error)
	GetCircuitByCitizenId(citizenId int) (*models.Circuit, error)
	AddCircuit(circuit models.Circuit) (*models.Circuit, error)
	Update(c models.Circuit) (*models.Circuit, error)
	Delete(id int) error
}
