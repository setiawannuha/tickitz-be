package routers

import (
	"setiawannuha/tickitz-be/internal/handlers"
	middleware "setiawannuha/tickitz-be/internal/middlewares"
	"setiawannuha/tickitz-be/internal/repository"
	"setiawannuha/tickitz-be/pkg"

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
	router.PATCH("/settings", middleware.Auth("user"), handler.Update)
	// router.GET("/", handler.FetchAll)
	router.GET("/profile", middleware.Auth("admin", "user"), handler.FetchDetail)
	// router.DELETE("/delete", middleware.Auth("admin", "user"), handler.Delete)
}
