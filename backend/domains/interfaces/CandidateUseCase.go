package interfaces

import "EleccionesUcu/dtos"

type CandidateUseCase interface {
	GetAll() ([]dtos.CandidateDto, error)
	GetByCitizenID(id int) ([]dtos.CandidateDto, error)
	Add(dto dtos.CandidateDto) (*dtos.CandidateDto, error)
	Delete(citizenID int, listNumber int) error
}
