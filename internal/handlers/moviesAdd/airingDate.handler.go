package handlers

import (
	models "khalifgfrz/coffee-shop-be-go/internal/models/moviesAdd"
	"khalifgfrz/coffee-shop-be-go/internal/repository"
	"khalifgfrz/coffee-shop-be-go/pkg"

	"github.com/gin-gonic/gin"
)

type HandlerAiringDate struct {
	repository.AiringDateRepoInterface
}

func NewAiringDate(r repository.AiringDateRepoInterface) *HandlerAiringDate {
	return &HandlerAiringDate{r}
}

func (h *HandlerAiringDate) PostAiringDate(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)

	airingDates := models.AiringDate{}

	if err := ctx.ShouldBind(&airingDates); err != nil {
		response.BadRequest("Insert airing date failed", err.Error())
		return
	}

	results, err := h.InsertAiringDate(&airingDates)
	if err != nil {
		response.InternalServerError("Internal Server Error", err.Error())
		return
	}

	response.Created("Airing date inserted successfully", results)
}

func (h *HandlerAiringDate) GetAllAiringDate(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)

	data, err := h.GetAiringDate()
	if err != nil {
		response.InternalServerError("Internal Server Error", err.Error())
	}
	if data == nil {
		response.NotFound("Data Not Found", "No datas available for the given criteria")
		return
	}

	response.Success("Data retrieved successfully", data)
}

func (h *HandlerAiringDate) GetAiringDateInput(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)

	input := models.AiringDate{}
	if err := ctx.ShouldBind(&input); err != nil {
		response.BadRequest("Failed to retrieve data, invalid input", err.Error())
		return
	}

	data, err := h.GetAiringDateByInput(&input)
	if err != nil {
		response.InternalServerError("Internal Server Error", err.Error())
	}
	if data == nil {
		response.NotFound("Data Not Found", "No datas available for the given criteria")
		return
	}

	response.Success("Data retrieved successfully", data)
}
