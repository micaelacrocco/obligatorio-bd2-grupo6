package repositories

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/models"
	"EleccionesUcu/utils"
	"database/sql"
)

type zoneMySQLRepo struct {
	db *sql.DB
}

func NewZoneRepository(db *sql.DB) interfaces.ZoneRepository {
	return &zoneMySQLRepo{db: db}
}

func (r *zoneMySQLRepo) GetAll() ([]models.Zone, error) {
	rows, err := r.db.Query("SELECT id, name, address, department_id FROM ZONES")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var zones []models.Zone
	for rows.Next() {
		var z models.Zone
		if err := rows.Scan(&z.ID, &z.Name, &z.Address, &z.DepartmentID); err != nil {
			return nil, err
		}
		zones = append(zones, z)
	}
	return zones, nil
}

func (r *zoneMySQLRepo) GetById(id int) (*models.Zone, error) {
	row := r.db.QueryRow("SELECT id, name, address, department_id FROM ZONES WHERE id = ?", id)

	var z models.Zone
	err := row.Scan(&z.ID, &z.Name, &z.Address, &z.DepartmentID)
	if err != nil {
		return nil, err
	}
	return &z, nil
}

func (r *zoneMySQLRepo) Add(zone models.Zone) (*models.Zone, error) {
	result, err := r.db.Exec("INSERT INTO ZONES(name, address, department_id) VALUES (?, ?, ?)", zone.Name, zone.Address, zone.DepartmentID)
	err = utils.ForeignKeyNotFoundError(err)
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()
	zone.ID = int(id)
	return &zone, nil
}

func (r *zoneMySQLRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM ZONES WHERE id = ?", id)
	return err
}
