package repositories

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/models"
	"EleccionesUcu/utils"
	"database/sql"
)

type listVoteMySQLRepo struct {
	db *sql.DB
}

func NewListVoteRepository(db *sql.DB) interfaces.ListVoteRepository {
	return &listVoteMySQLRepo{db: db}
}

func (r *listVoteMySQLRepo) GetAll() ([]models.ListVote, error) {
	rows, err := r.db.Query("SELECT id, vote_date, list_number, circuit_id FROM LIST_VOTES")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var votes []models.ListVote
	for rows.Next() {
		var v models.ListVote

		if err := rows.Scan(&v.ID, &v.VoteDate, &v.ListNumber, &v.CircuitID); err != nil {
			return nil, err
		}

		votes = append(votes, v)
	}
	return votes, nil
}

func (r *listVoteMySQLRepo) Add(vote models.ListVote) (*models.ListVote, error) {
	// Convert time.Time to string before insert

	_, err := r.db.Exec("INSERT INTO LIST_VOTES(vote_date, list_number, circuit_id) VALUES (?, ?, ?)", vote.VoteDate, vote.ListNumber, vote.CircuitID)
	err = utils.ForeignKeyNotFoundError(err)
	if err != nil {
		return nil, err
	}
	return &vote, nil
}

func (r *listVoteMySQLRepo) Update(vote models.ListVote) (*models.ListVote, error) {

	result, err := r.db.Exec("UPDATE LIST_VOTES SET vote_date = ?, list_number = ?, circuit_id = ? WHERE id = ?", vote.VoteDate, vote.ListNumber, vote.ID, vote.CircuitID)
	if err != nil {
		return nil, err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return nil, sql.ErrNoRows
	}
	return &vote, nil
}

func (r *listVoteMySQLRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM LIST_VOTES WHERE id = ?", id)
	return err
}
