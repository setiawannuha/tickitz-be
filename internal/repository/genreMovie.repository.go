package repository

import (
	models "khalifgfrz/coffee-shop-be-go/internal/models/moviesAdd"

	"github.com/jmoiron/sqlx"
)

type GenreMovieRepoInterface interface {
	InsertGenreMovie(data *models.GenreMovie) (*models.GenreMovie, error)
	UpdateGenreMovie(id int, data *models.GenreMovie) (*models.GenreMovie, error)
	DeleteGenreMovie(id string) (string, error)
}

type RepoGenreMovie struct {
	*sqlx.DB
}

func NewGenreMovieRepository(db *sqlx.DB) *RepoGenreMovie {
	return &RepoGenreMovie{db}
}

func (r *RepoGenreMovie) InsertGenreMovie(data *models.GenreMovie) (*models.GenreMovie, error) {
	query := `
    INSERT INTO public.genre_movies (
      "genre_id",
      "movie_id"
    ) VALUES (
      :genre_id,
      :movie_id
    ) RETURNING id, genre_id, movie_id, created_at, updated_at;
  `
	var results models.GenreMovie
	rows, err := r.DB.NamedQuery(query, data)
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

func (r *RepoGenreMovie) UpdateGenreMovie(id int, data *models.GenreMovie) (*models.GenreMovie, error) {
	query := `
    UPDATE public.genre_movies
    SET "genre_id" = :genre_id,
        "movie_id" = :movie_id
    WHERE id = :id
    RETURNING *;
  `

	data.ID = id

	var results models.GenreMovie
	rows, err := r.DB.NamedQuery(query, data)
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

func (r *RepoGenreMovie) DeleteGenreMovie(id string) (string, error) {
	query := `
    DELETE FROM public.genre_movies
    WHERE id = $1
    RETURNING *;
  `

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
