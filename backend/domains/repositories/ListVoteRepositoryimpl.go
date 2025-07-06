package repositories

import (
	"EleccionesUcu/domains/interfaces"
	"EleccionesUcu/models"
	"EleccionesUcu/utils"
	"database/sql"
	"time"
)

type listVoteMySQLRepo struct {
	db *sql.DB
}

func NewListVoteRepository(db *sql.DB) interfaces.ListVoteRepository {
	return &listVoteMySQLRepo{db: db}
}

func (r *listVoteMySQLRepo) GetAll() ([]models.ListVoteModel, error) {
	rows, err := r.db.Query("SELECT id, vote_date, list_number FROM LIST_VOTES")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var votes []models.ListVoteModel
	for rows.Next() {
		var v models.ListVoteModel
		var voteDateStr string

		if err := rows.Scan(&v.ID, &voteDateStr, &v.ListNumber); err != nil {
			return nil, err
		}

		v.VoteDate, _ = time.Parse("2006-01-02", voteDateStr)
		votes = append(votes, v)
	}
	return votes, nil
}

func (r *listVoteMySQLRepo) Add(vote models.ListVoteModel) (*models.ListVoteModel, error) {
	// Convert time.Time to string before insert
	voteDateStr := vote.VoteDate.Format("2006-01-02")

	_, err := r.db.Exec("INSERT INTO LIST_VOTES(vote_date, list_number) VALUES (?, ?)", voteDateStr, vote.ListNumber)
	err = utils.ForeignKeyNotFoundError(err)
	if err != nil {
		return nil, err
	}
	return &vote, nil
}

func (r *listVoteMySQLRepo) Update(vote models.ListVoteModel) (*models.ListVoteModel, error) {
	voteDateStr := vote.VoteDate.Format("2006-01-02")

	result, err := r.db.Exec("UPDATE LIST_VOTES SET vote_date = ?, list_number = ? WHERE id = ?", voteDateStr, vote.ListNumber, vote.ID)
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
