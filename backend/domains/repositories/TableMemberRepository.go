package repositories

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/models"
	"EleccionesUcu/utils"
	"database/sql"
)

type MySqlRepository struct {
	db *sql.DB
}

func NewTableMemberRepository(db *sql.DB) interfaces.TableMembersRepository {
	return &MySqlRepository{db: db}
}

func (r MySqlRepository) GetAll() ([]models.TableMembers, error) {
	rows, err := r.db.Query("SELECT t.table_id, t.citizen_id, t.integration_date, t.duty FROM TABLE_MEMBERS t")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tableMembers []models.TableMembers

	for rows.Next() {
		var t models.TableMembers
		if err := rows.Scan(&t.TableID, &t.CitizenID, &t.IntegrationDate, &t.Duty); err != nil {
			return nil, err
		}
		tableMembers = append(tableMembers, t)
	}
	return tableMembers, nil
}

func (r MySqlRepository) GetCitizenIsTableMember(citizenID int, tableID int) (*models.TableMembers, error) {
	query := "SELECT t.table_id, t.citizen_id, t.integration_date, t.duty FROM TABLE_MEMBERS t WHERE t.table_id = ? AND t.citizen_id = ?"
	row := r.db.QueryRow(query, citizenID, tableID)

	var t models.TableMembers

	if err := row.Scan(&t.TableID, &t.CitizenID, &t.IntegrationDate, &t.Duty); err != nil {
		return nil, err
	}
	return &t, nil
}

func (r MySqlRepository) Add(tableMember models.TableMembers) (*models.TableMembers, error) {
	query := "INSERT INTO TABLE_MEMBERS(table_id, citizen_id, integration_date, duty) VALUES(?, ?, ?, ?)"
	_, err := r.db.Exec(query, tableMember.TableID, tableMember.CitizenID, tableMember.IntegrationDate, tableMember.Duty)
	err = utils.ForeignKeyNotFoundError(err)

	if err != nil {
		return nil, err
	}

	return &tableMember, nil
}

func (r MySqlRepository) Delete(citizenID int, tableID int) error {
	_, err := r.db.Exec("DELETE FROM TABLE_MEMBERS t WHERE t.citizen_id = ? AND t.table_id = ?", citizenID, tableID)
	return err
}
