package repositories

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/models"
	"EleccionesUcu/utils"
	"database/sql"
)

type tableMySQLRepo struct {
	db *sql.DB
}

func NewTableRepository(db *sql.DB) interfaces.TableRepository {
	return &tableMySQLRepo{db: db}
}

func (r *tableMySQLRepo) GetAll() ([]models.Table, error) {
	rows, err := r.db.Query("SELECT id, circuit_id FROM TABLES")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []models.Table
	for rows.Next() {
		var t models.Table
		if err := rows.Scan(&t.ID, &t.CircuitID); err != nil {
			return nil, err
		}
		tables = append(tables, t)
	}
	return tables, nil
}

func (r *tableMySQLRepo) GetById(id int) (*models.Table, error) {
	row := r.db.QueryRow("SELECT id, circuit_id FROM TABLES WHERE id = ?", id)

	var t models.Table
	if err := row.Scan(&t.ID, &t.CircuitID); err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *tableMySQLRepo) Add(table models.Table) (*models.Table, error) {
	result, err := r.db.Exec("INSERT INTO TABLES(circuit_id) VALUES (?)", table.CircuitID)
	err = utils.ForeignKeyNotFoundError(err)
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()
	table.ID = int(id)
	return &table, nil
}

func (r *tableMySQLRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM TABLES WHERE id = ?", id)
	return err
}
