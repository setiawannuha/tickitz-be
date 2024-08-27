package routers

import (
	"khalifgfrz/coffee-shop-be-go/internal/handlers"
	"khalifgfrz/coffee-shop-be-go/internal/repository"
	"khalifgfrz/coffee-shop-be-go/pkg"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func movieRouter(g *gin.Engine, d *sqlx.DB) {
	router := g.Group("/movie")

	var movieRepo repository.MovieRepoInterface = repository.NewMovieRepository(d)
	var genreRepo repository.GenreMovieRepoInterface = repository.NewGenreMovieRepository(d)
	var airingDateRepo repository.AiringDateRepoInterface = repository.NewAiringDateRepository(d)
	var airingTimeDateRepo repository.AiringTimeDateRepoInterface = repository.NewAiringTimeDateRepository(d, nil)
	var movieTimeRepo repository.MovieTimeRepoInteface = repository.NewMovieTimeRepository(d)
	var locationMovieTimeRepo repository.LocationMovieTimeRepoInterface = repository.NewLocationMovieRepository(d)
	var cld pkg.Cloudinary = *pkg.NewCloudinaryUtil()
	handler := handlers.NewMovieRepository(movieRepo, genreRepo, airingDateRepo, airingTimeDateRepo, movieTimeRepo, locationMovieTimeRepo, cld, d)

	router.POST("/insert", handler.InsertMovies)
	router.GET("/", handler.GetMovies)
	router.GET("/:id", handler.GetMovieDetails)
	router.PATCH("/:id", handler.UpdateMovies)
	router.PATCH("/banner/:id", handler.BannerUpdate)
	router.DELETE("/:id", handler.MoviesDelete)

}
