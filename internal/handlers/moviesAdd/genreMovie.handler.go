package handlers

import (
	models "khalifgfrz/coffee-shop-be-go/internal/models/moviesAdd"
	"khalifgfrz/coffee-shop-be-go/internal/repository"
	"khalifgfrz/coffee-shop-be-go/pkg"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HandlerGenreMovie struct {
	repository.GenreMovieRepoInterface
}

func NewGenreMovie(r repository.GenreMovieRepoInterface) *HandlerGenreMovie {
	return &HandlerGenreMovie{r}
}

func (h *HandlerGenreMovie) PostGenreMovie(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)

	genreMovies := models.GenreMovie{}

	if err := ctx.ShouldBind(&genreMovies); err != nil {
		response.BadRequest("Insert genre movie failed", err.Error())
		return
	}

	results, err := h.InsertGenreMovie(&genreMovies)
	if err != nil {
		response.InternalServerError("Internar server error", err.Error())
		return
	}
	response.Created("Genre movie has been created", results)
}

func (h *HandlerGenreMovie) PatchGenreMovies(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	var genreMovie models.GenreMovie
	if err := ctx.ShouldBind(&genreMovie); err != nil {
		response.BadRequest("Update genre movie failed", err.Error())
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.BadRequest("Update genre movie failed", err.Error())
		return
	}
	results, err := h.UpdateGenreMovie(id, &genreMovie)
	if err != nil {
		response.InternalServerError("Internar server error", err.Error())
		return
	}
	response.Success("Genre movie has been updated", results)
}

func (h *HandlerGenreMovie) DeleteGenreMovies(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	id := ctx.Param("id")

	results, err := h.DeleteGenreMovie(id)
	if err != nil {
		response.InternalServerError("Internar server error", err.Error())
		return
	}
	response.Success("Genre movie has been deleted", results)
}
