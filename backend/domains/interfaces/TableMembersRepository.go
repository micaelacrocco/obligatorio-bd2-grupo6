package interfaces

import "EleccionesUcu/models"

type TableMembersRepository interface {
	GetAll() ([]models.TableMembers, error)
	GetCitizenIsTableMember(citizenID int, tableID int) (*models.TableMembers, error)
	Add(tableMember models.TableMembers) (*models.TableMembers, error)
	Delete(citizenID int, tableID int) error
}
