package interfaces

import (
	"EleccionesUcu/dtos"
)

type CircuitsUseCase interface {
	GetAll() ([]dtos.CircuitDto, error)
	GetById(id int) (*dtos.CircuitDto, error)
	GetVotesByParty(circuitID int) ([]dtos.PartyVoteDto, error)
	GetVotes(circuitID int) ([]dtos.CircuitResultDto, error)
	GetVotesByAllCandidates(circuitID int) ([]dtos.CircuitResultByAllCandidates, error)
	AddCircuit(circuit dtos.CircuitDto) (*dtos.CircuitDto, error)
	Update(dto dtos.CircuitDto) (*dtos.CircuitDto, error)
	Delete(id int) error
}
