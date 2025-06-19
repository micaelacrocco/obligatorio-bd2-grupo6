package interfaces

import "EleccionesUcu/dtos"

type CircuitsUseCase interface {
	GetAll() ([]dtos.CircuitDto, error)
}
