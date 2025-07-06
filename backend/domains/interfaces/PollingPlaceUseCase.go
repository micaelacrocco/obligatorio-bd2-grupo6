package interfaces

import "EleccionesUcu/dtos"

type PollingPlaceUseCase interface {
	GetAll() ([]dtos.PollingPlaceDto, error)
	GetByID(id int) (*dtos.PollingPlaceDto, error)
	Add(dtos.PollingPlaceDto) (*dtos.PollingPlaceDto, error)
	Delete(id int) error
}
