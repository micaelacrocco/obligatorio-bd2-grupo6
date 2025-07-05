package interfaces

import "EleccionesUcu/dtos"

type CitizenUseCase interface {
	GetByID(id int) (dtos.CitizenDto, error)
	GetAll() ([]dtos.CitizenDto, error)
	AddCitizen(dto dtos.CitizenDto) (dtos.CitizenDto, error)
	Update(id int, dto dtos.CitizenDto) error
	Delete(id int) error
}