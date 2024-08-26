package handlers

import (
	models "khalifgfrz/coffee-shop-be-go/internal/models/moviesAdd"
	"khalifgfrz/coffee-shop-be-go/internal/repository"
	"khalifgfrz/coffee-shop-be-go/pkg"

	"github.com/gin-gonic/gin"
)

type HandlerLocationMovieTime struct {
	repository.LocationMovieTimeRepoInterface
}

func NewLocationMovieTime(r repository.LocationMovieTimeRepoInterface) *HandlerLocationMovieTime {
	return &HandlerLocationMovieTime{r}
}

func (h *HandlerLocationMovieTime) PostLocationMovieTime(ctx *gin.Context) {

	response := pkg.NewResponse(ctx)

	data := models.LocationMovieTime{}

	if err := ctx.ShouldBind(&data); err != nil {
		response.BadRequest("Insert Location  failed", err.Error())
		return
	}

	results, err := h.CreateLocationMovie(&data)
	if err != nil {
		response.InternalServerError("Internar server error", err.Error())
		return
	}
	response.Created("Location has been created", results)
}

func (h *HandlerLocationMovieTime) GetLocationMovieById(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)

	id := ctx.Param("id")

	data, err := h.GetMovieLocTimeById(id)
	if err != nil {
		response.InternalServerError("Internal Server Error", err.Error())
	}
	if data == nil {
		response.NotFound("Data Not Found", "No datas available for the given criteria")
		return
	}

	response.Success("Data retrieved successfully", data)
}

func (h *HandlerLocationMovieTime) DeleteLocationMovies(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	id := ctx.Param("id")

	results, err := h.DeleteLocationMovie(id)
	if err != nil {
		response.InternalServerError("Internar server error", err.Error())
		return
	}
	response.Success("Genre movie has been deleted", results)
}
