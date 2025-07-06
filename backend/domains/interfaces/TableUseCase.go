package interfaces

import "EleccionesUcu/dtos"

type TableUseCase interface {
	GetAll() ([]dtos.TableDto, error)
	GetById(id int) (*dtos.TableDto, error)
	Add(dto dtos.TableDto) (*dtos.TableDto, error)
	Delete(id int) error
}
