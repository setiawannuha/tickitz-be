package repository

import (
	models "khalifgfrz/coffee-shop-be-go/internal/models/moviesAdd"

	"github.com/jmoiron/sqlx"
)

type AiringDateRepoInterface interface {
	InsertAiringDate(dates *models.AiringDate) ([]models.AiringDate, error)
	GetAiringDate() (*models.AiringDate, error)
	GetAiringDateByInput(input *models.AiringDate) (*models.AiringDate, error)
}

type RepoAiringDate struct {
	*sqlx.DB
}

func NewAiringDateRepository(db *sqlx.DB) *RepoAiringDate {
	return &RepoAiringDate{db}
}

// insert dynamically
func (r *RepoAiringDate) CreateAiringDate(dates *models.AiringDate) (*models.AiringDate, error) {
	// insert start_date and end_date into database
	query := `INSERT INTO airing_date (start_date, end_date) VALUES (:start_date,:end_date) RETURNING *`

	var results models.AiringDate
	rows, err := r.DB.NamedQuery(query, dates)
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

// methods get all airing_date
func (r *RepoAiringDate) GetAiringDate() (*models.AiringDate, error) {
	query := `SELECT * FROM airing_date`

	var airingDate models.AiringDate
	err := r.DB.Get(&airingDate, query)

	return &airingDate, err
}

func (r *RepoAiringDate) GetAiringDateByInput(input *models.AiringDate) (*models.AiringDate, error) {
	query := `SELECT * FROM airing_date
	WHERE start_date = :start_date AND end_date = :end_date
	RETURNING *`

	var airingDate models.AiringDate
	err := r.DB.Get(&airingDate, query)

	return &airingDate, err
}
