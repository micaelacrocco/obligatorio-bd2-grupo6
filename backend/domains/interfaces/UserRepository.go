package interfaces

import "EleccionesUcu/models"

type UserRepository interface {
	FindByCitizenID(citizenID int) (models.User, error)
}
