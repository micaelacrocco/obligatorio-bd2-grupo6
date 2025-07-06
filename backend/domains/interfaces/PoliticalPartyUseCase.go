package interfaces

import (
	"EleccionesUcu/dtos"
)

type PoliticalPartyUseCase interface {
	GetAll() ([]dtos.PoliticalPartyDTO, error)
	Add(party dtos.PoliticalPartyDTO) (*dtos.PoliticalPartyDTO, error)
	Update(party dtos.PoliticalPartyDTO) (*dtos.PoliticalPartyDTO, error)
	Delete(id int) error
}
