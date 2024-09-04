package routes

import (
	"go-shop-api/adapters/handler"
	"go-shop-api/adapters/repository"
	"go-shop-api/core/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Cart
// @Description Cart
// @Shemes http
// @Tags Cart
// @Param request body request.NewCartItemRequest true "Create Cart Item Request"
// @Param request body request.UpdQauntityCartItem true "Update Cart Item Request"
// @Param id path string true "Cart Item ID"
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.CartItemResponse
// @Router /v1/cart [post]
// @Router /v1/cart [get]
// @Router /v1/cart/update [put]
// @Router /v1/cart/{id} [delete]
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func RegisterCartRoutes(router *gin.Engine, db *gorm.DB) {
	cartItemRepo := repository.NewCartItemRepositoryDB(db)
	cartItemService := service.NewCartItemService(cartItemRepo)
	cartItemHandler := handler.NewHttpCartItemHandler(cartItemService)

	cart := router.Group("/v1/cart")
	{
		cart.POST("/", cartItemHandler.AddCartItem)
		cart.GET("/", cartItemHandler.GetCartItems)
		cart.PUT("/update", cartItemHandler.UpdateCartItem)
		cart.DELETE("/:id", cartItemHandler.DeleteCartItem)
	}

}
