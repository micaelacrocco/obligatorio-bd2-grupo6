package repositories

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/models"
	"EleccionesUcu/utils"
	"database/sql"
	"errors"
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

func (r *circuitMySQLRepo) GetById(id int) (*models.Circuit, error) {
	query := "SELECT c.id, c.location, c.is_accessible, c.credential_start, c.credential_end, c.polling_place_id from CIRCUITS c WHERE c.id = ?"
	row := r.db.QueryRow(query, id)

	var c models.Circuit

	err := row.Scan(&c.ID, &c.Location, &c.Accessible, &c.CredentialStart, &c.CredentialEnd, &c.PollingPlaceId)

	if errors.Is(sql.ErrNoRows, err) {
		return nil, err
	}
	return &c, nil
}

func (r *circuitMySQLRepo) AddCircuit(circuit models.Circuit) (*models.Circuit, error) {
	query := "INSERT INTO CIRCUITS(id, location, is_accessible, credential_start, credential_end, polling_place_id) VALUES(?, ?, ?, ?, ?, ?)"
	_, err := r.db.Exec(query, circuit.ID, circuit.Location, circuit.Accessible, circuit.CredentialStart, circuit.CredentialEnd, circuit.PollingPlaceId)

	err = utils.ForeignKeyNotFoundError(err)
	if err != nil {
		return nil, err
	}

	return &circuit, nil
}

func (r *circuitMySQLRepo) Update(circuit models.Circuit) (*models.Circuit, error) {
	query := "UPDATE CIRCUITS SET location = ?, is_accessible = ?, credential_start = ?, credential_end = ?, polling_place_id = ? WHERE id = ?"
	_, err := r.db.Exec(query, circuit.Location, circuit.Accessible, circuit.CredentialStart, circuit.CredentialEnd, circuit.PollingPlaceId, circuit.ID)
	if err != nil {
		return nil, err
	}
	return &circuit, nil
}

func (r *circuitMySQLRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM CIRCUITS WHERE id = ?", id)
	return err
}
