package interfaces

import "EleccionesUcu/models"

type UserUseCase interface {
	FindByCitizenID(citizenID int) (models.User, error) // debe llamarse FindByCitizenID para que coincida con handler
	CheckPassword(plain string, user models.User) bool
}
