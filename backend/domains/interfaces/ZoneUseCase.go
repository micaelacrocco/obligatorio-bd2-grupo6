package interfaces

import "EleccionesUcu/dtos"

type ZoneUseCase interface {
	GetAll() ([]dtos.ZoneDto, error)
	GetById(id int) (*dtos.ZoneDto, error)
	Add(dto dtos.ZoneDto) (*dtos.ZoneDto, error)
	Delete(id int) error
}
