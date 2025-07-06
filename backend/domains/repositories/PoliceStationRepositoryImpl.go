package repositories

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/models"
	"EleccionesUcu/utils"
	"database/sql"
)

type policeStationMySQLRepo struct {
	db *sql.DB
}

func NewPoliceStationRepository(db *sql.DB) interfaces.PoliceStationRepository {
	return &policeStationMySQLRepo{db: db}
}

func (r *policeStationMySQLRepo) GetAll() ([]models.PoliceStation, error) {
	rows, err := r.db.Query("SELECT id, station_number, address, department_id FROM POLICE_STATIONS")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stations []models.PoliceStation
	for rows.Next() {
		var s models.PoliceStation
		if err := rows.Scan(&s.ID, &s.StationNumber, &s.Address, &s.DepartmentID); err != nil {
			return nil, err
		}
		stations = append(stations, s)
	}
	return stations, nil
}

func (r *policeStationMySQLRepo) Add(station models.PoliceStation) (*models.PoliceStation, error) {
	result, err := r.db.Exec("INSERT INTO POLICE_STATIONS(station_number, address, department_id) VALUES (?, ?, ?)",
		station.StationNumber, station.Address, station.DepartmentID)
	err = utils.ForeignKeyNotFoundError(err)
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()
	station.ID = int(id)
	return &station, nil
}

func (r *policeStationMySQLRepo) Update(station models.PoliceStation) (*models.PoliceStation, error) {
	result, err := r.db.Exec("UPDATE POLICE_STATIONS SET station_number = ?, address = ?, department_id = ? WHERE id = ?",
		station.StationNumber, station.Address, station.DepartmentID, station.ID)
	if err != nil {
		return nil, err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return nil, sql.ErrNoRows
	}
	return &station, nil
}

func (r *policeStationMySQLRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM POLICE_STATIONS WHERE id = ?", id)
	return err
}
