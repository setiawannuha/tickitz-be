package routers

import (
	"khalifgfrz/coffee-shop-be-go/internal/handlers"
	"khalifgfrz/coffee-shop-be-go/internal/repository"
	"khalifgfrz/coffee-shop-be-go/pkg"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

func movieRouter(g *gin.Engine, d *sqlx.DB) {
	router := g.Group("/movie")

	var movieRepo repository.MovieRepoInterface = repository.NewMovieRepository(d)
	var genreRepo repository.GenreMovieRepoInterface = repository.NewGenreMovieRepository(d)
	var airingDateRepo repository.AiringDateRepoInterface = repository.NewAiringDateRepository(d)
	var movieTimeRepo repository.MovieTimeRepoInteface = repository.NewMovieTimeRepository(d)
	var locationMovieTimeRepo repository.LocationMovieTimeRepoInterface = repository.NewLocationMovieRepository(d)
	var gorm *gorm.DB
	var cld pkg.Cloudinary = *pkg.NewCloudinaryUtil()
	handler := handlers.NewMovieRepository(movieRepo, genreRepo, airingDateRepo, movieTimeRepo, locationMovieTimeRepo, cld, gorm)

	router.POST("/insert", handler.InsertMovies)

}
