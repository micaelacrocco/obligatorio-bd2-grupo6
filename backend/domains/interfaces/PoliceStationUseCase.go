package interfaces

import "EleccionesUcu/dtos"

type PoliceStationUseCase interface {
	GetAll() ([]dtos.PoliceStationDto, error)
	Add(dto dtos.PoliceStationDto) (*dtos.PoliceStationDto, error)
	Update(dto dtos.PoliceStationDto) (*dtos.PoliceStationDto, error)
	Delete(id int) error
}
