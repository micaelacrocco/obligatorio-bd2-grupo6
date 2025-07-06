package interfaces

import (
	"EleccionesUcu/dtos"
)

type TableMembersUseCase interface {
	GetAll() ([]dtos.TableMembersDto, error)
	GetCitizenIsTableMember(citizenID int, tableID int) (bool, error)
	Add(tableMember dtos.TableMembersDto) (*dtos.TableMembersDto, error)
	Delete(citizenID int, tableMemberID int) error
}
