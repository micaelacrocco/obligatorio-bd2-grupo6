package repositories

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/models"
	"database/sql"
)

type userMySQLRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) interfaces.UserRepository {
	return &userMySQLRepo{db: db}
}

func (r *userMySQLRepo) FindByCitizenID(citizenID int) (models.User, error) {
	var user models.User
	query := "SELECT id, password_hashed, user_type, citizen_id FROM users WHERE citizen_id = ?"
	err := r.db.QueryRow(query, citizenID).Scan(
		&user.ID,
		&user.PasswordHashed,
		&user.UserType,
		&user.CitizenID,
	)
	return user, err
}
