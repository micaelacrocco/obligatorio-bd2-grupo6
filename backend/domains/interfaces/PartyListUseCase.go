package interfaces

import "EleccionesUcu/dtos"

type PartyListUseCase interface {
	GetAll() ([]dtos.PartyListDto, error)
	Add(dto dtos.PartyListDto) (*dtos.PartyListDto, error)
	Update(dto dtos.PartyListDto) (*dtos.PartyListDto, error)
	Delete(listNumber int) error
}
