package repositories

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/models"
	"database/sql"
	"errors"
)

type citizenMySQLRepo struct {
	db *sql.DB
}

func NewCitizenRepository(db *sql.DB) interfaces.CitizenRepository {
	return &citizenMySQLRepo{db: db}
}

func (r *citizenMySQLRepo) GetAll() ([]models.Citizen, error) {
	rows, err := r.db.Query("SELECT id, first_name, last_name, birth_date, credential FROM CITIZENS")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var citizens []models.Citizen

	for rows.Next() {
		var citizen models.Citizen
		if err := rows.Scan(&citizen.ID, &citizen.FirstName, &citizen.LastName, &citizen.BirthDate, &citizen.Credential); err != nil {
			return nil, err
		}
		citizens = append(citizens, citizen)
	}

	return citizens, nil
}

func (r *citizenMySQLRepo) GetByID(id int) (models.Citizen, error) {
	query := "SELECT id, first_name, last_name, birth_date, credential FROM CITIZENS WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var citizen models.Citizen
	err := row.Scan(&citizen.ID, &citizen.FirstName, &citizen.LastName, &citizen.BirthDate, &citizen.Credential)

	if errors.Is(err, sql.ErrNoRows) {
		return models.Citizen{}, err
	} else if err != nil {
		return models.Citizen{}, err
	}

	return citizen, nil
}

func (r *citizenMySQLRepo) Create(c models.Citizen) (int, error) {
	var id int
	err := r.db.QueryRow(
		"INSERT INTO citizens (first_name, last_name, birth_date, credential) VALUES (?, ?, ?, ?) RETURNING id",
		c.FirstName, c.LastName, c.BirthDate, c.Credential).Scan(&id)
	return id, err
}

func (r *citizenMySQLRepo) Update(id int, citizen models.Citizen) error {
	query := "UPDATE CITIZENS SET first_name = ?, last_name = ?, birth_date = ?, credential = ? WHERE id = ?"
	_, err := r.db.Exec(query, citizen.FirstName, citizen.LastName, citizen.BirthDate, citizen.Credential, id)
	return err
}

func (r *citizenMySQLRepo) Delete(id int) error {
	query := "DELETE FROM CITIZENS WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}
