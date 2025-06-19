package repositories

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/models"
	"database/sql"
)

type circuitMySQLRepo struct {
	db *sql.DB
}

func NewCircuitRepository(db *sql.DB) interfaces.CircuitsRepository {
	return &circuitMySQLRepo{db: db}
}

func (r *circuitMySQLRepo) GetAll() ([]models.Circuit, error) {
	rows, err := r.db.Query("SELECT c.id, c.location, c.is_accessible, c.credential_start, c.credential_end, c.polling_place_id from CIRCUITS c")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var circuits []models.Circuit

	for rows.Next() {
		var c models.Circuit
		if err := rows.Scan(&c.ID, &c.Location, &c.Accessible, &c.CredentialStart, &c.CredentialEnd, &c.PollingPlaceId); err != nil {
			return nil, err
		}

		circuits = append(circuits, c)
	}
	return circuits, nil
}
