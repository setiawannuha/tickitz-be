package routers

import (
	"khalifgfrz/coffee-shop-be-go/internal/handlers"
	"khalifgfrz/coffee-shop-be-go/internal/middlewares"
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
	router.PATCH("/settings", middleware.Auth("admin", "user"), handler.Update)
	// router.GET("/", handler.FetchAll)
	router.GET("/profile", middleware.Auth("admin", "user"), handler.FetchDetail)
	router.DELETE("/delete", middleware.Auth("admin", "user"), handler.Delete)

}
