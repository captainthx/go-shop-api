package routes

import (
	"go-shop-api/adapters/handler"
	"go-shop-api/adapters/repository"
	"go-shop-api/core/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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
