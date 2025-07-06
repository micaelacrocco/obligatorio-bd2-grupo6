package interfaces

import "EleccionesUcu/dtos"

type ListVoteUseCase interface {
	GetAll() ([]dtos.ListVoteDto, error)
	Add(dto dtos.ListVoteDto) (*dtos.ListVoteDto, error)
	Update(dto dtos.ListVoteDto) (*dtos.ListVoteDto, error)
	Delete(id int) error
}
