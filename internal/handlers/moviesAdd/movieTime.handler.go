package handlers

import (
	models "khalifgfrz/coffee-shop-be-go/internal/models/moviesAdd"
	"khalifgfrz/coffee-shop-be-go/internal/repository"
	"khalifgfrz/coffee-shop-be-go/pkg"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HandlerMovieTime struct {
	repository.MovieTimeRepoInteface
}

func NewMovieTime(r repository.MovieTimeRepoInteface) *HandlerMovieTime {
	return &HandlerMovieTime{r}
}

func (h *HandlerMovieTime) PostMovieTimes(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)

	movieTimes := models.MovieTime{}

	if err := ctx.ShouldBind(&movieTimes); err != nil {
		response.BadRequest("Create movie time failed, invalid input", err.Error())
		return
	}

	results, err := h.CreateMovieTime(&movieTimes)
	if err != nil {
		response.InternalServerError("Internal Server Error", err.Error())
		return
	}

	response.Created("Movie time has been created", results)
}

func (h *HandlerMovieTime) GetMovieTimeById(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)

	id := ctx.Param("id")

	data, err := h.GetTimeByMovieId(id)
	if err != nil {
		response.InternalServerError("Internal Server Error", err.Error())
	}
	if data == nil {
		response.NotFound("Data Not Found", "No datas available for the given criteria")
		return
	}

	response.Success("Data retrieved successfully", data)
}

func (h *HandlerMovieTime) PatchMovieTime(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)

	id := ctx.Param("id")

	movieTime := models.MovieTime{}

	if err := ctx.ShouldBindJSON(&movieTime); err != nil {
		response.BadRequest("Update movie time failed, invalid input", err.Error())
		return
	}

	result, err := h.UpdateMovieTime(id, &movieTime)
	if err != nil {
		response.InternalServerError("Internal Server Error", err.Error())
		return
	}

	response.Success("Data updated successfully", result)
}

func (h *HandlerMovieTime) DeleteMovieTimes(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.BadRequest("Delete genre movie failed", err.Error())
		return
	}

	results, err := h.DeleteMovieTime(id)
	if err != nil {
		response.InternalServerError("Internar server error", err.Error())
		return
	}
	response.Success("Genre movie has been deleted", results)
}
