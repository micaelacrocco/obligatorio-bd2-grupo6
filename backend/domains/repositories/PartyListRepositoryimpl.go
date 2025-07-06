package repositories

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/models"
	"EleccionesUcu/utils"
	"database/sql"
)

type partyListMySQLRepo struct {
	db *sql.DB
}

func NewPartyListRepository(db *sql.DB) interfaces.PartyListRepository {
	return &partyListMySQLRepo{db: db}
}

func (r *partyListMySQLRepo) GetAll() ([]models.PartyList, error) {
	rows, err := r.db.Query("SELECT list_number, party_id FROM PARTY_LISTS")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lists []models.PartyList
	for rows.Next() {
		var l models.PartyList
		if err := rows.Scan(&l.ListNumber, &l.PartyID); err != nil {
			return nil, err
		}
		lists = append(lists, l)
	}
	return lists, nil
}

func (r *partyListMySQLRepo) Add(list models.PartyList) (*models.PartyList, error) {
	_, err := r.db.Exec("INSERT INTO PARTY_LISTS(list_number, party_id) VALUES (?, ?)", list.ListNumber, list.PartyID)
	err = utils.ForeignKeyNotFoundError(err)
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *partyListMySQLRepo) Update(list models.PartyList) (*models.PartyList, error) {
	result, err := r.db.Exec("UPDATE PARTY_LISTS SET party_id = ? WHERE list_number = ?", list.PartyID, list.ListNumber)
	if err != nil {
		return nil, err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return nil, sql.ErrNoRows
	}
	return &list, nil
}

func (r *partyListMySQLRepo) Delete(listNumber int) error {
	_, err := r.db.Exec("DELETE FROM PARTY_LISTS WHERE list_number = ?", listNumber)
	return err
}
