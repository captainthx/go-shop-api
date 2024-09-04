package routes

import (
	adminHandler "go-shop-api/adapters/handler/admin"
	adminRepository "go-shop-api/adapters/repository/admin"
	adminService "go-shop-api/core/service/admin"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterProductAdminRoutes(router *gin.Engine, db *gorm.DB) {
	prodAdminRepo := adminRepository.NewProductAdminRepositoryDB(db)
	prodAdminService := adminService.NewProductAdminService(prodAdminRepo)
	prodAdminHandler := adminHandler.NewHttpProductAdminHandler(prodAdminService)

	prodAdmin := router.Group("/v1/admin/product")
	prodAdmin.POST("/", prodAdminHandler.CreateProduct)
}
