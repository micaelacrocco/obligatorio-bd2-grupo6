package repositories

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/models"
	"EleccionesUcu/utils"
	"database/sql"
	"time"
)

type candidateMySQLRepo struct {
	db *sql.DB
}

func NewCandidateRepository(db *sql.DB) interfaces.CandidateRepository {
	return &candidateMySQLRepo{db: db}
}

func (r *candidateMySQLRepo) GetAll() ([]models.Candidate, error) {
	rows, err := r.db.Query("SELECT citizen_id, list_number, start_date, end_date, candidacy_type FROM CANDIDATES")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.Candidate
	for rows.Next() {
		var c models.Candidate
		var startDate, endDate string
		if err := rows.Scan(&c.CitizenID, &c.ListNumber, &startDate, &endDate, &c.CandidacyType); err != nil {
			return nil, err
		}
		c.StartDate, _ = time.Parse("2006-01-02", startDate)
		c.EndDate, _ = time.Parse("2006-01-02", endDate)
		result = append(result, c)
	}
	return result, nil
}

func (r *candidateMySQLRepo) GetByCitizenID(id int) ([]models.Candidate, error) {
	rows, err := r.db.Query("SELECT citizen_id, list_number, start_date, end_date, candidacy_type FROM CANDIDATES WHERE citizen_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.Candidate
	for rows.Next() {
		var c models.Candidate
		var startDate, endDate string
		if err := rows.Scan(&c.CitizenID, &c.ListNumber, &startDate, &endDate, &c.CandidacyType); err != nil {
			return nil, err
		}
		c.StartDate, _ = time.Parse("2006-01-02", startDate)
		c.EndDate, _ = time.Parse("2006-01-02", endDate)
		result = append(result, c)
	}
	return result, nil
}

func (r *candidateMySQLRepo) Add(candidate models.Candidate) (*models.Candidate, error) {
	_, err := r.db.Exec("INSERT INTO CANDIDATES (citizen_id, list_number, start_date, end_date, candidacy_type) VALUES (?, ?, ?, ?, ?)",
		candidate.CitizenID, candidate.ListNumber, candidate.StartDate.Format("2006-01-02"), candidate.EndDate.Format("2006-01-02"), candidate.CandidacyType)
	err = utils.ForeignKeyNotFoundError(err)
	if err != nil {
		return nil, err
	}
	return &candidate, nil
}

func (r *candidateMySQLRepo) Delete(citizenID int, listNumber int) error {
	_, err := r.db.Exec("DELETE FROM CANDIDATES WHERE citizen_id = ? AND list_number = ?", citizenID, listNumber)
	return err
}
