package interfaces

import "EleccionesUcu/dtos"

type PoliceAgentUseCase interface {
	GetAll() ([]dtos.PoliceAgentDto, error)
	GetByCitizenID(ci int) (*dtos.PoliceAgentDto, error)
	Add(dto dtos.PoliceAgentDto) (*dtos.PoliceAgentDto, error)
	Update(dto dtos.PoliceAgentDto) (*dtos.PoliceAgentDto, error)
	Delete(ci int) error
}
