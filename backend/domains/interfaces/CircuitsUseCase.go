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
	GetVotesPersonById(citizenID int) (*dtos.PersonVoteDTO, error)
	AddCircuit(circuit dtos.CircuitDto) (*dtos.CircuitDto, error)
	AddVotePerson(vote dtos.PersonVoteDTO) (*dtos.PersonVoteDTO, error)
	Update(dto dtos.CircuitDto) (*dtos.CircuitDto, error)
	Delete(id int) error
}
