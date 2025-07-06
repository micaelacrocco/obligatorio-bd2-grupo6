package repositories

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/models"
	"EleccionesUcu/utils"
	"database/sql"
)

type policeAgentMySQLRepo struct {
	db *sql.DB
}

func NewPoliceAgentRepository(db *sql.DB) interfaces.PoliceAgentRepository {
	return &policeAgentMySQLRepo{db: db}
}

func (r *policeAgentMySQLRepo) GetAll() ([]models.PoliceAgent, error) {
	rows, err := r.db.Query("SELECT citizen_id, police_station_id, polling_place_id FROM POLICE_AGENTS")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var agents []models.PoliceAgent
	for rows.Next() {
		var a models.PoliceAgent
		if err := rows.Scan(&a.CitizenID, &a.PoliceStationID, &a.PollingPlaceID); err != nil {
			return nil, err
		}
		agents = append(agents, a)
	}
	return agents, nil
}

func (r *policeAgentMySQLRepo) GetByCitizenID(ci int) (*models.PoliceAgent, error) {
	row := r.db.QueryRow("SELECT citizen_id, police_station_id, polling_place_id FROM POLICE_AGENTS WHERE citizen_id = ?", ci)

	var a models.PoliceAgent
	err := row.Scan(&a.CitizenID, &a.PoliceStationID, &a.PollingPlaceID)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *policeAgentMySQLRepo) Add(agent models.PoliceAgent) (*models.PoliceAgent, error) {
	_, err := r.db.Exec("INSERT INTO POLICE_AGENTS(citizen_id, police_station_id, polling_place_id) VALUES (?, ?, ?)",
		agent.CitizenID, agent.PoliceStationID, agent.PollingPlaceID)
	err = utils.ForeignKeyNotFoundError(err)
	if err != nil {
		return nil, err
	}
	return &agent, nil
}

func (r *policeAgentMySQLRepo) Update(agent models.PoliceAgent) (*models.PoliceAgent, error) {
	_, err := r.db.Exec(
		"UPDATE POLICE_AGENTS SET police_station_id = ?, polling_place_id = ? WHERE citizen_id = ?",
		agent.PoliceStationID, agent.PollingPlaceID, agent.CitizenID)
	if err != nil {
		return nil, err
	}
	return &agent, nil
}

func (r *policeAgentMySQLRepo) Delete(ci int) error {
	_, err := r.db.Exec("DELETE FROM POLICE_AGENTS WHERE citizen_id = ?", ci)
	return err
}
