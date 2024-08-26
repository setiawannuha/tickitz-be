package repository

import (
	"fmt"
	"khalifgfrz/coffee-shop-be-go/internal/models"
	"strings"

	"github.com/jmoiron/sqlx"
)

type MovieRepoInterface interface {
	CreateMovie(data *models.MoviesBody) (*models.Movies, error)
	GetAllMovies(query *models.MoviesQuery) (*models.MovieResponse, int, error)
	GetDetailMovie(id string) (*models.Movies, error)
	UpdateMovie(id string, data *models.Movies) (string, error)
	DeleteMovie(id string) (string, error)
}

type RepoMovies struct {
	*sqlx.DB
}

func NewMovieRepository(db *sqlx.DB) *RepoMovies {
	return &RepoMovies{db}
}

func (r *RepoMovies) CreateMovie(data *models.MoviesBody) (*models.Movies, error) {
	query := `
        INSERT INTO public.movies (
            "title",
            "image", 
            "director", 
            "casts",
            "duration",
            "release_date",
            "synopsis",
						"is_deleted"
        ) VALUES (
            :title,
            :image,
            :director,
            :casts,
            :duration,
            :release_date,
            :synopsis,
						FALSE
        ) RETURNING id, title, image, director, casts, duration, release_date, synopsis, created_at;
    `
	var result models.Movies
	rows, err := r.DB.NamedQuery(query, data)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.StructScan(&result)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func (r *RepoMovies) GetAllMovies(query *models.MoviesQuery) (*models.MovieResponse, int, error) {
	baseQuery := `
  	SELECT 
      "m".id, 
      "m".title, 
      "m".image, 
      COALESCE(STRING_AGG("g"."name", ', '), '') AS genres,
      "m".director, 
      "m".casts, 
      "m".duration, 
      "m".release_date, 
      "m".synopsis, 
      "m".is_deleted, 
      "m".created_at, 
      "m".updated_at
  	FROM public.movies "m"
  	LEFT JOIN public.genre_movies gm ON "m"."id" = "gm"."movie_id"
  	LEFT JOIN public.genres "g" ON "gm"."genre_id" = "g"."id"
    `
	countQuery := `
      SELECT COUNT(DISTINCT "m"."id") 
      FROM public.movies m
      LEFT JOIN public.genre_movies gm ON "m"."id" = "gm"."movie_id"
      LEFT JOIN public.genres "g" ON "gm"."genre_id" = "g"."id"
      WHERE "m"."is_deleted" = FALSE
    `
	whereClauses := []string{}
	var values []interface{}

	if query.Search != nil {
		searchTerm := "%" + *query.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf(`m.title ILIKE $%d`, len(values)+1))
		values = append(values, searchTerm)
	}

	if query.Filter != nil {
		filterTerm := *query.Filter
		whereClauses = append(whereClauses, fmt.Sprintf(`"g"."name" = $%d`, len(values)+1))
		values = append(values, filterTerm)
	}

	if len(whereClauses) > 0 {
		whereQuery := " WHERE " + strings.Join(whereClauses, " AND ")
		baseQuery += whereQuery
		countQuery += whereQuery
	}

	baseQuery += ` GROUP BY "m"."id"`

	if query.Page > 0 && query.Limit > 0 {
		limit := query.Limit
		offset := (query.Page - 1) * limit
		baseQuery += fmt.Sprintf(` LIMIT $%d OFFSET $%d`, len(values)+1, len(values)+2)
		values = append(values, limit, offset)
	}

	var data models.MovieResponse
	if err := r.Select(&data, baseQuery, values...); err != nil {
		return nil, 0, err
	}

	var total int
	if err := r.Get(&total, countQuery, values[:len(values)-2]...); err != nil {
		return nil, 0, err
	}

	return &data, total, nil
}

func (r *RepoMovies) GetDetailMovie(id string) (*models.Movies, error) {
	query := `
    SELECT 
      "m".id, 
      "m".title, 
      "m".image, 
      COALESCE(STRING_AGG("g"."name", ', '), '') AS genres,
      "m".director, 
      "m".casts, 
      "m".duration, 
      "m".release_date, 
      "m".synopsis, 
      "m".is_deleted, 
      "m".created_at, 
      "m".updated_at
  	FROM public.movies "m"
  	LEFT JOIN public.genre_movies gm ON "m"."id" = "gm"."movie_id"
  	LEFT JOIN public.genres "g" ON "gm"."genre_id" = "g"."id"
  `

	var result models.Movies
	if err := r.Get(&result, query, id); err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *RepoMovies) UpdateMovie(id string, data *models.Movies) (string, error) {
	query := `
    UPDATE public.movies
    SET
      "title" = COALESCE(NULLIF(:title, ''), "title"),
      "image" = COALESCE(NULLIF(:image, ''), "image"),
      "director" = COALESCE(NULLIF(:director, ''), "director"),
      "casts" = COALESCE(NULLIF(:casts, ''), "casts"),
      "duration" = COALESCE(NULLIF(:duration, ''), "duration"),
      "release_date" = COALESCE(NULLIF(:release_date, ''), "release_date"),
      "synopsis" = COALESCE(NULLIF(:synopsis, ''), "synopsis"),
			"updated_at" = now()
    WHERE "id" = :id
    RETURNING *
  `
	data.Id = id

	var updatedMovie models.Movies
	rows, err := r.DB.NamedQuery(query, data)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.StructScan(&updatedMovie)
		if err != nil {
			return "", err
		}
	}

	return "Data updated", nil
}

func (r *RepoMovies) DeleteMovie(id string) (string, error) {
	query := `
    DELETE FROM public.movies
    WHERE "id" = $1
    RETURNING *
  `

	var deletedMovie models.Movies
	if err := r.DB.QueryRowx(query, id).StructScan(&deletedMovie); err != nil {
		return "", err
	}

	return "Data deleted", nil
}
