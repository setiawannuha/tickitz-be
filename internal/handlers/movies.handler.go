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
	"github.com/jmoiron/sqlx"
)

type HandlerMovie struct {
	repository.MovieRepoInterface
	repository.GenreMovieRepoInterface
	repository.AiringDateRepoInterface
	repository.AiringTimeDateRepoInterface
	repository.MovieTimeRepoInteface
	repository.LocationMovieTimeRepoInterface
	pkg.Cloudinary
	DB *sqlx.DB
}

func NewMovieRepository(
	mr repository.MovieRepoInterface,
	gmr repository.GenreMovieRepoInterface,
	ad repository.AiringDateRepoInterface,
	atd repository.AiringTimeDateRepoInterface,
	mt repository.MovieTimeRepoInteface,
	lmt repository.LocationMovieTimeRepoInterface,
	cld pkg.Cloudinary,
	db *sqlx.DB,
) *HandlerMovie {
	return &HandlerMovie{mr, gmr, ad, atd, mt, lmt, cld, db}
}

func SplitCommaSeparatedInts(input string) ([]int, error) {
	var result []int
	if input == "" {
		return result, nil
	}

	parts := strings.Split(input, ",")
	for _, part := range parts {
		trimmedPart := strings.TrimSpace(part)
		value, err := strconv.Atoi(trimmedPart)
		if err != nil {
			return nil, err
		}
		result = append(result, value)
	}
	return result, nil
}

func (h *HandlerMovie) InsertMovies(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)

	var movies models.MoviesBody

	if h.DB == nil {
		response.InternalServerError("Database connection is not initialized", nil)
		return
	}

	tx, err := h.DB.Beginx()
	if err != nil {
		response.InternalServerError("Failed to start transaction", err.Error())
		return
	}
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

	genres, err := SplitCommaSeparatedInts(movies.Genres)
	if err != nil {
		tx.Rollback()
		response.BadRequest("Invalid genre ID format", err.Error())
		return
	}

	airingTimes, err := SplitCommaSeparatedInts(movies.AiringTime)
	if err != nil {
		tx.Rollback()
		response.BadRequest("Invalid airing time format", err.Error())
		return
	}

	locations, err := SplitCommaSeparatedInts(movies.Locations)
	if err != nil {
		tx.Rollback()
		response.BadRequest("Invalid location format", err.Error())
		return
	}

	_, err = govalidator.ValidateStruct(&movies)
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

	// Step 1: Create the movie and get the ID
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

	// Step 2: Link Movie ID with Genre IDs based on genres []int
	for _, genreId := range genres {
		genreMovie := moviesAdd.GenreMovie{
			Movie_id: results.Id,
			Genre_id: genreId,
		}
		if _, err := h.InsertGenreMovie(&genreMovie); err != nil {
			tx.Rollback()
			response.InternalServerError("Failed to insert genre data", err.Error())
			return
		}
	}

	// Step 3: Create AiringDate based on airingDate []string (start and end)
	for _, dateRange := range movies.AiringDate {
		dates := strings.Split(dateRange, ",")
		var airingDate moviesAdd.AiringDate

		if len(dates) == 2 {
			airingDate = moviesAdd.AiringDate{
				Start_date: strings.TrimSpace(dates[0]),
				End_date:   strings.TrimSpace(dates[1]),
			}
		} else if len(dates) == 1 {
			airingDate = moviesAdd.AiringDate{
				Start_date: strings.TrimSpace(dates[0]),
				End_date:   strings.TrimSpace(dates[0]),
			}
		} else {
			tx.Rollback()
			response.BadRequest("Invalid date range provided", "Expected one or two dates in the format [yyyy-mm-dd] or [yyyy-mm-dd, yyyy-mm-dd]")
			return
		}

		// Check if airingDate already exists and insert or retrieve its ID
		existingAiringDate, err := h.GetAiringDateByInput(&airingDate)
		if err != nil {
			tx.Rollback()
			response.InternalServerError("Failed to check airing date", err.Error())
			return
		}

		if existingAiringDate != nil {
			airingDate.Id = existingAiringDate.Id
		} else {
			insertedAiringDates, err := h.CreateAiringDate(&airingDate)
			if err != nil {
				tx.Rollback()
				response.InternalServerError("Failed to insert airing date", err.Error())
				return
			}
			if len(insertedAiringDates) > 0 {
				airingDate.Id = insertedAiringDates[0].Id
			} else {
				tx.Rollback()
				response.InternalServerError("No airing date was inserted", "Expected a valid ID")
				return
			}
		}

		// Step 4: Link AiringDate ID with AiringTime IDs based on airingTime []int
		for _, airingTimeId := range airingTimes {
			newAiringTimeDate := moviesAdd.AiringTimeDate{
				Airing_time_id: airingTimeId,
				Date_id:        airingDate.Id,
			}
			insertedAiringTimeDateId, err := h.InsertAiringTimeDate(&newAiringTimeDate)
			if err != nil {
				tx.Rollback()
				response.InternalServerError("Failed to insert airing time date", err.Error())
				return
			}

			// Step 5: Link Movie ID with AiringTimeDate ID and get MovieTime ID
			movieTime := moviesAdd.MovieTime{
				Movie_id:            results.Id,
				Airing_time_date_id: insertedAiringTimeDateId.Id,
			}
			insertedMovieTime, err := h.CreateMovieTime(&movieTime)
			if err != nil {
				tx.Rollback()
				response.InternalServerError("Failed to insert movie time", err.Error())
				return
			}

			// Step 6: Link Location IDs with Movie ID based on location []int (id)
			for _, locationId := range locations {
				movieLocation := moviesAdd.LocationMovieTime{
					Movie_time_id: insertedMovieTime.ID,
					Location_id:   locationId,
				}
				if _, err := h.CreateLocationMovie(&movieLocation); err != nil {
					tx.Rollback()
					response.InternalServerError("Failed to insert movie location", err.Error())
					return
				}
			}
		}

		// Commit transaction only if all operations are successful
		if err := tx.Commit(); err != nil {
			tx.Rollback()
			response.InternalServerError("Internal Server Error", "Failed to commit transaction")
			return
		}

		// Return successful response
		response.Created("Data created", results)
	}
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
