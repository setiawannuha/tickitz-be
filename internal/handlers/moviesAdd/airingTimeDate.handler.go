package handlers

import (
	models "khalifgfrz/coffee-shop-be-go/internal/models/moviesAdd"
	"khalifgfrz/coffee-shop-be-go/internal/repository"
	"khalifgfrz/coffee-shop-be-go/pkg"

	"github.com/gin-gonic/gin"
)

type HandlerAiringTimeDate struct {
	repository.AiringTimeDateRepoInterface
}

func NewAiringTimeDate(r repository.AiringTimeDateRepoInterface) *HandlerAiringTimeDate {
	return &HandlerAiringTimeDate{r}
}

func (h *HandlerAiringTimeDate) PostAiringTimeDate(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)

	airingTimeDates := models.AiringTimeDate{}

	if err := ctx.ShouldBind(&airingTimeDates); err != nil {
		response.BadRequest("Insert airing date failed", err.Error())
		return
	}

	results, err := h.InsertAiringTimeDate(&airingTimeDates)
	if err != nil {
		response.InternalServerError("Internal Server Error", err.Error())
		return
	}

	response.Created("Airing time date inserted successfully", results)
}

func (h *HandlerAiringTimeDate) GetAllAiringTimeDate(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)

	data, err := h.GetAiringTimeDate()
	if err != nil {
		response.InternalServerError("Internal Server Error", err.Error())
	}
	if len(data) == 0 {
		response.NotFound("Data Not Found", "No datas available for the given criteria")
		return
	}

	response.Success("Data retrieved successfully", data)
}
