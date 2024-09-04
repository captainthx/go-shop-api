package routes

import (
	"go-shop-api/adapters/handler"
	"go-shop-api/adapters/repository"
	"go-shop-api/core/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Product Customer
// @Description Product Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Success 200 {object} common.Pagination
// @Router /v1/product [get]

func RegisterProductCusRoutes(router *gin.Engine, db *gorm.DB) {
	prodRepo := repository.NewProductRepositoryDB(db)
	prodService := service.NewProductService(prodRepo)
	prodHandler := handler.NewHttpProductHandler(prodService)

	// product router
	prodCustumer := router.Group("/v1/product")

	prodCustumer.GET("/", prodHandler.GetProductList)
}
