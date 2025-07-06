package interfaces

import "EleccionesUcu/dtos"

type DepartmentUseCase interface {
	GetAll() ([]dtos.DepartmentDto, error)
	Add(dto dtos.DepartmentDto) (*dtos.DepartmentDto, error)
	Delete(id int) error
}
