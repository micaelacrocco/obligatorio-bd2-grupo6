package usecases

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/models"
	"crypto/sha256"
	"encoding/hex"
)

type userUseCase struct {
	repo interfaces.UserRepository
}

func NewUserUseCase(r interfaces.UserRepository) interfaces.UserUseCase {
	return &userUseCase{repo: r}
}

/* ---- IMPLEMENTA LA INTERFAZ EXACTA ---- */

func (u *userUseCase) FindByCitizenID(citizenID int) (models.User, error) {
	return u.repo.FindByCitizenID(citizenID)
}

func (u *userUseCase) CheckPassword(plain string, user models.User) bool {
	hash := sha256.Sum256([]byte(plain))
	return hex.EncodeToString(hash[:]) == user.PasswordHashed
}
