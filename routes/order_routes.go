package routes

import (
	"go-shop-api/adapters/handler"
	"go-shop-api/adapters/repository"
	"go-shop-api/core/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterOrderRoutes(router *gin.Engine, db *gorm.DB) {
	orderRepo := repository.NewOrderRepositoryDB(db)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := handler.NewHttpOrderHandler(orderService)

	// order router
	order := router.Group("/v1/order")
	{
		order.GET("/", orderHandler.GetOrderHistoryList)
		order.GET("/search", orderHandler.GetOrderListByStatus)
		order.POST("/", orderHandler.CreateOrder)
		order.POST("/cancel/:id", orderHandler.CancelOrder)
	}

}
