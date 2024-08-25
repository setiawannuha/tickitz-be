package repository

import (
	models "khalifgfrz/coffee-shop-be-go/internal/models/moviesAdd"

	"github.com/jmoiron/sqlx"
)

type LocationMovieTimeRepoInterface interface {
	CreateLocationMovie(data *models.LocationMovieTime) (string, error)
	GetMovieLocTimeById(id string) (*models.LocationMovieTime, error)
	DeleteLocationMovie(id string) (string, error)
}

type RepoLocationMovie struct {
	*sqlx.DB
}

func NewLocationMovieRepository(db *sqlx.DB) *RepoLocationMovie {
	return &RepoLocationMovie{db}
}

func (r *RepoLocationMovie) CreateLocationMovie(data *models.LocationMovieTime) (string, error) {
	query := `
    INSERT INTO public.location_movie_time (
      "location_id",
			"movie_time_id"
		) VALUES (
			:location_id,
			:movie_time_id)`

	var results models.LocationMovieTime
	rows, err := r.DB.NamedQuery(query, data)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.StructScan(&results)
		if err != nil {
			return "", err
		}
	}

	return "Data created", nil
}

func (r *RepoLocationMovie) GetMovieLocTimeById(id string) (*models.LocationMovieTime, error) {
	query := `
    SELECT * FROM public.location_movie_time
    WHERE "movie_time_id" = $1`

	var data models.LocationMovieTime
	if err := r.DB.QueryRowx(query, id).StructScan(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *RepoLocationMovie) DeleteLocationMovie(id string) (string, error) {
	query := `DELETE FROM public.location_movie_time WHERE "movie_time_id" = :movie_time_id`

	var results models.GenreMovie
	rows, err := r.DB.Queryx(query, id)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.StructScan(&results)
		if err != nil {
			return "", err
		}
	}

	return "Data deleted", nil
}
