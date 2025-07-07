package repositories

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/models"
	"EleccionesUcu/utils"
	"database/sql"
	"errors"
	"strconv"
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
func (r *circuitMySQLRepo) GetVotesPersonById(citizenID int) (*models.PersonVoteModel, error) {
	query := "SELECT * FROM PERSON_VOTES WHERE citizen_id = ?"
	row := r.db.QueryRow(query, citizenID)

	var c models.PersonVoteModel

	err := row.Scan(&c.ID, &c.VoteDate, &c.IsObserved, &c.VoteType, &c.CitizenID, &c.CircuitID)

	if errors.Is(sql.ErrNoRows, err) {
		return nil, err
	}
	return &c, nil
}

func (r *circuitMySQLRepo) GetVotesByParty(circuitID int) ([]models.PartyVote, error) {
	query := `
		SELECT p.name AS partido, COUNT(*) AS votos
		FROM LIST_VOTES lv
		JOIN PARTY_LISTS pl ON lv.list_number = pl.list_number
		JOIN PARTIES p ON pl.party_id = p.id
		WHERE lv.circuit_id = ?
		GROUP BY p.name;
	`

	rows, err := r.db.Query(query, circuitID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.PartyVote
	for rows.Next() {
		var vote models.PartyVote
		if err := rows.Scan(&vote.PartyName, &vote.VoteCount); err != nil {
			return nil, err
		}
		results = append(results, vote)
	}
	return results, nil
}

func (r *circuitMySQLRepo) GetVotes(circuitID int) ([]models.CircuitResult, error) {
	query := `
	WITH NormalVotes AS (
    SELECT COUNT(*) AS count
    FROM PERSON_VOTES
    WHERE circuit_id = 300 AND vote_type = 'Normal'
),
     ListVotes AS (
         SELECT
             CONCAT('Lista ', lv.list_number) AS list,
             p.name AS party_name,
             COUNT(*) AS vote_count
         FROM LIST_VOTES lv
                  JOIN PARTY_LISTS pl ON lv.list_number = pl.list_number
                  JOIN PARTIES p ON pl.party_id = p.id
         WHERE lv.circuit_id = 300
         GROUP BY lv.list_number, p.name
     ),
     OtherVotes AS (
         SELECT
             vote_type AS list,
             vote_type AS party_name,
             COUNT(*) AS vote_count
         FROM PERSON_VOTES
         WHERE circuit_id = ? AND vote_type IN ('En Blanco', 'Anulado')
         GROUP BY vote_type
     )
SELECT * FROM ListVotes
UNION ALL
SELECT * FROM OtherVotes;

	`

	rows, err := r.db.Query(query, circuitID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.CircuitResult
	for rows.Next() {
		var vote models.CircuitResult
		if err := rows.Scan(&vote.List, &vote.PartyName, &vote.VoteCount); err != nil {
			return nil, err
		}
		results = append(results, vote)
	}
	return results, nil
}

func (r *circuitMySQLRepo) GetVotesByAllCandidates(circuitID int) ([]models.CircuitResultByAllCandidates, error) {
	query := `
	WITH ListVotes AS (
    SELECT
        p.name AS party,
        CONCAT(c.first_name, ' ', c.last_name) AS candidate,
        COUNT(*) AS vote_count
    FROM LIST_VOTES lv
             JOIN PARTY_LISTS pl ON lv.list_number = pl.list_number
             JOIN PARTIES p ON pl.party_id = p.id
             LEFT JOIN LEADERS l ON p.id = l.party_id AND l.election_year = 2025
             LEFT JOIN CITIZENS c ON l.citizen_id = c.id
    WHERE lv.circuit_id = 300
    GROUP BY p.name, c.first_name, c.last_name
),
     OtherVotes AS (
         SELECT
             vote_type AS party,
             vote_type AS candidate,
             COUNT(*) AS vote_count
         FROM PERSON_VOTES
         WHERE circuit_id = ? AND vote_type IN ('En Blanco', 'Anulado')
         GROUP BY vote_type
     )

SELECT * FROM ListVotes
UNION ALL
SELECT * FROM OtherVotes
	;

	`

	rows, err := r.db.Query(query, circuitID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.CircuitResultByAllCandidates
	for rows.Next() {
		var vote models.CircuitResultByAllCandidates
		if err := rows.Scan(&vote.Party, &vote.Candidate, &vote.VoteCount); err != nil {
			return nil, err
		}
		results = append(results, vote)
	}
	return results, nil
}

func (r *circuitMySQLRepo) GetCircuitByCitizenId(citizenId int) (*models.Circuit, error) {
	// Primero obtenemos el credential del ciudadano
	var credential string
	err := r.db.QueryRow("SELECT credential FROM CITIZENS WHERE id = ?", citizenId).Scan(&credential)
	if err != nil {
		return nil, err
	}

	// Sacamos los 3 primeros caracteres y convertimos a int
	if len(credential) < 4 {
		return nil, errors.New("credential too short")
	}
	numericPart := credential[3:] // desde el 4to caracter (index 3)
	credInt, err := strconv.Atoi(numericPart)
	if err != nil {
		return nil, err
	}

	// Ahora buscamos el circuito cuyo rango incluya ese número
	query := `
	SELECT id, location, is_accessible, credential_start, credential_end, polling_place_id
	FROM CIRCUITS
	WHERE ? BETWEEN credential_start AND credential_end
	LIMIT 1;
	`
	row := r.db.QueryRow(query, credInt)
	var c models.Circuit
	err = row.Scan(&c.ID, &c.Location, &c.Accessible, &c.CredentialStart, &c.CredentialEnd, &c.PollingPlaceId)
	if err != nil {
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

func (r *circuitMySQLRepo) AddVotePerson(vote models.PersonVoteModel) (*models.PersonVoteModel, error) {
	query := "INSERT INTO PERSON_VOTES (vote_date, is_observed, vote_type, citizen_id, circuit_id) VALUES (?, ?, ?, ? , ?)"

	_, err := r.db.Exec(query, vote.VoteDate, vote.IsObserved, vote.VoteType, vote.CitizenID, vote.CircuitID)
	if err != nil {
		return nil, err
	}

	return &vote, nil
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
