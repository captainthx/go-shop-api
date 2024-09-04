package routes

import (
	"go-shop-api/adapters/handler"
	"go-shop-api/adapters/repository"
	"go-shop-api/core/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Order
// @Description Order
// @Shemes http
// @Tags Order
// @Param request body request.NewOrderReuqest true "Create Order Request"
// @Param id path string true "Order ID"
// @Param status query string false "Order Status"
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Router /v1/order [post]
// @Router /v1/order [get]
// @Router /v1/order/search [get]
// @Router /v1/order/cancel/{id} [post]
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
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
