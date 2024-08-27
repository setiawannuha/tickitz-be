package routers

import (
	"khalifgfrz/coffee-shop-be-go/internal/handlers"
	middleware "khalifgfrz/coffee-shop-be-go/internal/middlewares"
	"khalifgfrz/coffee-shop-be-go/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func orderRouter(g *gin.Engine, d *sqlx.DB) {
	router := g.Group("/order")

	var orderRepo repository.OrderRepositoryInterface = repository.NewOrderRepository(d)
	var orderDetailsRepo repository.OrderDetailsRepositoryInterface = repository.NewOrderDetailsRepository(d)
	handler := handlers.NewOrderHandler(orderRepo, orderDetailsRepo)

	router.POST("/", middleware.Auth("user"), handler.CreateOrder)
	router.GET("/history", middleware.Auth("user"), handler.FetchHistory)
	router.GET("/", middleware.Auth("admin"), handler.FetchAll)
	router.GET("/:id", middleware.Auth("user"), handler.FetchDetail)
}
