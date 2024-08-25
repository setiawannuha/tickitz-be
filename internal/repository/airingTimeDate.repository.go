package repository

import (
	models "khalifgfrz/coffee-shop-be-go/internal/models/moviesAdd"

	"github.com/jmoiron/sqlx"
)

type AiringTimeDateRepoInterface interface {
	GetAiringTimeDate() ([]models.AiringTimeDate, error)
	InsertAiringTimeDate(airingTimeDate *models.AiringTimeDate) (*models.AiringTimeDate, error)
}

type RepoAiringTimeDate struct {
	*sqlx.DB
}

func NewAiringTimeDateRepository(db *sqlx.DB) *RepoAiringTimeDate {
	return &RepoAiringTimeDate{db}
}

//mthods get airing_time from database

func (r *RepoAiringTimeDate) GetAiringTimeDate() ([]models.AiringTimeDate, error) {
	query := `SELECT * FROM airing_time_date`

	var airingTimeDates []models.AiringTimeDate
	err := r.DB.Select(&airingTimeDates, query)

	return airingTimeDates, err
}

//methods insert airing_time_date

func (r *RepoAiringTimeDate) InsertAiringTimeDate(airingTimeDate *models.AiringTimeDate) (*models.AiringTimeDate, error) {
	query := `INSERT INTO airing_time_date (airing_time_id, airing_date_id) VALUES (:airing_time_id,:airing_date_id) RETURNING *`

	var results models.AiringTimeDate
	rows, err := r.DB.NamedQuery(query, airingTimeDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.StructScan(&results)
		if err != nil {
			return nil, err
		}
	}

	return &results, nil

}
