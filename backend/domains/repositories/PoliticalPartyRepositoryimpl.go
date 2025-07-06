package repositories

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/models"
	"EleccionesUcu/utils"
	"database/sql"
)

type politicalPartyMySQLRepo struct {
	db *sql.DB
}

func NewPoliticalPartyRepository(db *sql.DB) interfaces.PoliticalPartyRepository {
	return &politicalPartyMySQLRepo{db: db}
}

func (r *politicalPartyMySQLRepo) GetAll() ([]models.PoliticalParty, error) {
	rows, err := r.db.Query("SELECT p.id, p.name FROM PARTIES p")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var parties []models.PoliticalParty

	for rows.Next() {
		var p models.PoliticalParty
		if err := rows.Scan(&p.ID, &p.Name); err != nil {
			return nil, err
		}
		parties = append(parties, p)
	}
	return parties, nil
}

func (r *politicalPartyMySQLRepo) Add(party models.PoliticalParty) (*models.PoliticalParty, error) {
	query := "INSERT INTO PARTIES(name) VALUES (?)"
	result, err := r.db.Exec(query, party.Name)

	err = utils.ForeignKeyNotFoundError(err)
	if err != nil {
		return nil, err
	}

	insertedID, err := result.LastInsertId()
	if err == nil {
		party.ID = int(insertedID)
	}

	return &party, nil
}

func (r *politicalPartyMySQLRepo) Update(party models.PoliticalParty) (*models.PoliticalParty, error) {
	query := "UPDATE PARTIES SET name = ? WHERE id = ?"
	result, err := r.db.Exec(query, party.Name, party.ID)

	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}

	return &party, nil
}

func (r *politicalPartyMySQLRepo) Delete(id int) error {
	query := "DELETE FROM PARTIES p WHERE p.id = ?"
	_, err := r.db.Exec(query, id)
	return err
}
