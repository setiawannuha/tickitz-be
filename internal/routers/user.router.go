package routers

import (
	"khalifgfrz/coffee-shop-be-go/internal/handlers"
	"khalifgfrz/coffee-shop-be-go/internal/repository"
	"khalifgfrz/coffee-shop-be-go/pkg"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func authRouter(g *gin.Engine, d *sqlx.DB) {
	router := g.Group("/user")

	var userRepo repository.UserRepositoryInterface = repository.NewUserRepository(d)
	var authRepo repository.AuthRepositoryInterface = repository.NewAuthRepository(d)
	var cld pkg.Cloudinary = *pkg.NewCloudinaryUtil()
	handler := handlers.NewAuthHandler(userRepo, authRepo, cld)

	router.POST("/register", handler.Register)
	router.POST("/login", handler.Login)
	router.PATCH("/settings" , handler.Update)
	// router.GET("/", handler.FetchAll)
	router.GET("/profile", handler.FetchDetail)
	router.DELETE("/delete" , handler.Delete)
	
}