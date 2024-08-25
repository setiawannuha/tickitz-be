package handlers

import (
	"fmt"
	"khalifgfrz/coffee-shop-be-go/internal/models"
	"khalifgfrz/coffee-shop-be-go/internal/models/moviesAdd"
	"khalifgfrz/coffee-shop-be-go/internal/repository"
	"khalifgfrz/coffee-shop-be-go/pkg"
	"math/rand"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HandlerMovie struct {
	repository.MovieRepoInterface
	repository.GenreMovieRepoInterface
	repository.AiringDateRepoInterface
	repository.MovieTimeRepoInteface
	repository.LocationMovieTimeRepoInterface
	pkg.Cloudinary
	DB *gorm.DB
}

func NewMovieRepository(
	mr repository.MovieRepoInterface,
	gmr repository.GenreMovieRepoInterface,
	ad repository.AiringDateRepoInterface,
	mt repository.MovieTimeRepoInteface,
	lmt repository.LocationMovieTimeRepoInterface,
	cld pkg.Cloudinary,
	db *gorm.DB,
) *HandlerMovie {
	return &HandlerMovie{mr, gmr, ad, mt, lmt, cld, db}
}

func (h *HandlerMovie) InsertMovies(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)

	var movies models.MoviesBody

	tx := h.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := ctx.ShouldBind(&movies); err != nil {
		tx.Rollback()
		response.BadRequest("Create movie failed, invalid input", err.Error())
		return
	}

	_, err := govalidator.ValidateStruct(&movies)
	if err != nil {
		tx.Rollback()
		response.BadRequest("Create movie failed", err.Error())
		return
	}

	file, header, err := ctx.Request.FormFile("image")
	if err != nil {
		tx.Rollback()
		response.BadRequest("Create movie failed, upload file is required", nil)
		return
	}

	mimeType := header.Header.Get("Content-Type")
	if mimeType != "image/jpg" && mimeType != "image/jpeg" && mimeType != "image/png" {
		tx.Rollback()
		response.BadRequest("Create movie failed, upload file failed, file type is not supported", nil)
		return
	}

	if header.Size > 2*1024*1024 {
		tx.Rollback()
		response.BadRequest("Create movie failed, upload file failed, file size exceeds 2 MB", nil)
		return
	}

	randomNumber := rand.Int()
	fileName := fmt.Sprintf("movie-image-%d", randomNumber)
	uploadResult, err := h.UploadFile(ctx, file, fileName)
	if err != nil {
		tx.Rollback()
		response.BadRequest("Create movie failed, upload file failed", err.Error())
		return
	}

	imageURL := uploadResult.SecureURL
	movies.Image = imageURL

	results, err := h.CreateMovie(&movies)
	if err != nil {
		tx.Rollback()
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			if strings.Contains(err.Error(), "unique_name") {
				response.BadRequest("Create movie failed, movie name already exists", err.Error())
				return
			}
		}
		response.InternalServerError("Internal Server Error", err.Error())
		return
	}

	for _, genre := range movies.Genres {
		genreMovie := moviesAdd.GenreMovie{
			Movie_id: results.Id,
			Genre_id: genre.ID,
		}
		if _, err := h.InsertGenreMovie(&genreMovie); err != nil {
			tx.Rollback()
			response.InternalServerError("Failed to insert genre data", err.Error())
			return
		}
	}

	// // 5. Insert related data into the AiringDate table
	for _, airingDate := range movies.AiringDate {
		existingAiringDate, err := h.GetAiringDateByInput(&airingDate)
		if err != nil {
			tx.Rollback()
			response.InternalServerError("Failed to check airing date", err.Error())
			return
		}

		if existingAiringDate != nil {
			// Use existing ID
			airingDate.Id = existingAiringDate.Id
		} else {
			// Insert new airing date
			airingDate := moviesAdd.AiringDate{
				Start_date: airingDate.Start_date,
				End_date:   airingDate.End_date,
			}
			_, err := h.InsertAiringDate(&airingDate)
			if err != nil {
				tx.Rollback()
				response.InternalServerError("Failed to insert airing date", err.Error())
				return
			}
		}
	}

	for _, movieTime := range movies.MovieTime {
		movieTimeRecord := moviesAdd.MovieTime{
			Movie_id:            results.Id,
			Airing_time_date_id: movieTime.Airing_time_date_id,
		}
		if _, err := h.CreateMovieTime(&movieTimeRecord); err != nil {
			tx.Rollback()
			response.InternalServerError("Failed to insert movie time", err.Error())
			return
		}
	}

	for _, location := range movies.LocationMovieTime {
		locationMovieTime := moviesAdd.LocationMovieTime{
			Location_id: location.Location_id,
		}
		if _, err := h.CreateLocationMovie(&locationMovieTime); err != nil {
			tx.Rollback()
			response.InternalServerError("Failed to insert location movie time", err.Error())
			return
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		response.InternalServerError("Internal Server Error", "Failed to commit transaction")
		return
	}

	response.Created("Data created", results)
}

func (h *HandlerMovie) GetMovies(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)

	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		response.BadRequest("Invalid or missing 'page' parameter", err.Error())
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		response.BadRequest("Invalid or missing 'limit' parameter", err.Error())
		return
	}

	query := models.MoviesQuery{
		Page:  page,
		Limit: limit,
	}

	if err := ctx.ShouldBindQuery(&query); err != nil {
		response.BadRequest("Invalid query parameter", err.Error())
		return
	}

	movies, total, err := h.GetAllMovies(&query)
	if err != nil {
		response.InternalServerError("Internal Server Error", err.Error())
		return
	}

	if len(*movies) == 0 {
		response.NotFound("Movie Not Found", "No movies available for the given criteria")
		return
	}

	totalPages := (total + query.Limit - 1) / query.Limit
	meta := &pkg.Meta{
		Total:     total,
		TotalPage: totalPages,
		Page:      query.Page,
		NextPage:  0,
		PrevPage:  0,
	}

	if query.Page+1 <= totalPages {
		meta.NextPage = query.Page + 1
	}

	if query.Page > 1 {
		meta.PrevPage = query.Page - 1
	}

	response.GetAllSuccess("Data fetched", movies, meta)
}
