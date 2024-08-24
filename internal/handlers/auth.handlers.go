package handlers

import (
	"fmt"
	"khalifgfrz/coffee-shop-be-go/internal/models"
	"khalifgfrz/coffee-shop-be-go/internal/repository"
	"khalifgfrz/coffee-shop-be-go/pkg"
	"math/rand"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	repository.UserRepositoryInterface
	repository.AuthRepositoryInterface
	pkg.Cloudinary
}

func NewAuthHandler(userRepo repository.UserRepositoryInterface , authRepo repository.AuthRepositoryInterface , cld pkg.Cloudinary) *AuthHandler {
	return &AuthHandler{userRepo , authRepo ,cld}
}

func (h *AuthHandler) Register(ctx *gin.Context){
	response := pkg.NewResponse(ctx)
	body := models.User{}

	if err := ctx.ShouldBind(&body); err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}

	_ ,err := govalidator.ValidateStruct(&body)
	if err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}

	body.Password , err = pkg.HashPassword(body.Password)
	if err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}
	
	body.Role = "USER"

	result , err := h.CreateData(&body)
	if err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}
	response.Created("create data success", result)
}

func (h *AuthHandler) Login(ctx *gin.Context){
	response := pkg.NewResponse(ctx)
	body := models.UserLogin{}
	
	if err := ctx.ShouldBind(&body); err != nil {
		response.BadRequest("Login failed", err.Error())
		return
	}

	_ ,err := govalidator.ValidateStruct(&body)
	if err != nil {
		response.BadRequest("Login failed", err.Error())
		return
	}

	result, err := h.GetByEmail(body.Email)
	if err != nil {
		response.BadRequest("Login failed", err.Error())
		return
	}
	err = pkg.VerifyPassword(result.Password, body.Password)
	if err != nil {
		response.Unauthorized("wrong password", err.Error())
		return
	}

	jwt := pkg.NewJWT(result.Id , result.Email )
	token , err := jwt.GenerateToken()
	if err != nil {
		response.Unauthorized("failed generate token", err.Error())
		return
	}
	
	response.Success("Login successful", map[string]interface{}{
		"token": token,
		"id":    result.Id,
	})
}

func (h *AuthHandler) Update(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	body := models.User{}
	id := ctx.Param("id")

	if err := ctx.ShouldBind(&body); err != nil {
		response.InternalServerError("update data failed", err.Error())
		return
	}

	/* _, err := govalidator.ValidateStruct(&body)
	if err != nil {
		response.BadRequest("update data failed", err.Error())
		return
	} */

	file, header, err := ctx.Request.FormFile("profile")
	if err != nil {
		if err != http.ErrMissingFile {
			response.BadRequest("update data failed, upload file failed", err.Error())
			return
		}
	} else {
		fmt.Println(header.Size)
		mimeType := header.Header.Get("Content-Type")
		if mimeType != "image/jpg" && mimeType != "image/png" {
			response.BadRequest("update data failed, upload file failed, wrong file type", err)
			return
		}

		randomNumber := rand.Int()
		fileName := fmt.Sprintf("go-profile-%d", randomNumber)
		uploadResult, err := h.UploadFile(ctx, file, fileName)
		if err != nil {
			response.BadRequest("update data failed, upload file failed", err.Error())
			return
		}
		body.Image = uploadResult.SecureURL
	}

	if body.Password != "" {
		body.Password, err = pkg.HashPassword(body.Password)
		if err != nil {
			response.BadRequest("create data failed", err.Error())
			return
		}
	}

	result, err := h.UpdateData(&body, id)
	if err != nil {
		response.InternalServerError("update data failed", err.Error())
		return
	}

	response.Success("update data success", result)
}

func (h *AuthHandler) FetchAll(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)

	result, err := h.GetAllData()
	if err != nil {
		response.InternalServerError("get data failed", err.Error())
		return
	}

	response.Success("get data success", result)
}

func (h *AuthHandler) FetchDetail(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	id := ctx.Param("id")

	result, err := h.GetDetailData(id)
	if err != nil {
		response.InternalServerError("get data failed", err.Error())
		return
	}

	response.Success("get data success", result)
}

func (h *AuthHandler) Delete(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	id := ctx.Param("id")

	result, err := h.DeleteData(id)
	if err != nil {
		response.InternalServerError("Delete data failed", err.Error())
		return
	}

	response.Success("Delete data success", result)
}