package repositories

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/models"
	"database/sql"
)

type pollingPlaceMySQLRepo struct {
	db *sql.DB
}

func NewPollingPlaceRepository(db *sql.DB) interfaces.PollingPlaceRepository {
	return &pollingPlaceMySQLRepo{db: db}
}

func (r *pollingPlaceMySQLRepo) GetAll() ([]models.PollingPlace, error) {
	rows, err := r.db.Query("SELECT id, name, type, address, zone_id FROM POLLING_PLACES")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.PollingPlace
	for rows.Next() {
		var p models.PollingPlace
		if err := rows.Scan(&p.ID, &p.Name, &p.Type, &p.Address, &p.ZoneID); err != nil {
			return nil, err
		}
		result = append(result, p)
	}
	return result, nil
}

func (r *pollingPlaceMySQLRepo) GetByID(id int) (*models.PollingPlace, error) {
	row := r.db.QueryRow("SELECT id, name, type, address, zone_id FROM POLLING_PLACES WHERE id = ?", id)
	var p models.PollingPlace
	if err := row.Scan(&p.ID, &p.Name, &p.Type, &p.Address, &p.ZoneID); err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *pollingPlaceMySQLRepo) Add(p models.PollingPlace) (*models.PollingPlace, error) {
	res, err := r.db.Exec("INSERT INTO POLLING_PLACES (name, type, address, zone_id) VALUES (?, ?, ?, ?)",
		p.Name, p.Type, p.Address, p.ZoneID)
	if err != nil {
		return nil, err
	}
	id, _ := res.LastInsertId()
	p.ID = int(id)
	return &p, nil
}

func (r *pollingPlaceMySQLRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM POLLING_PLACES WHERE id = ?", id)
	return err
}
