package routes

import (
	adminHandler "go-shop-api/adapters/handler/admin"
	adminRepository "go-shop-api/adapters/repository/admin"
	adminService "go-shop-api/core/service/admin"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Product Admin
// @Description Product Admin
// @Shemes http
// @Tags Product Admin
// @Param request body domain.Product true "Create Product Request"
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 201 {object} string
// @Router /v1/admin/product [post]
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func RegisterProductAdminRoutes(router *gin.Engine, db *gorm.DB) {
	prodAdminRepo := adminRepository.NewProductAdminRepositoryDB(db)
	prodAdminService := adminService.NewProductAdminService(prodAdminRepo)
	prodAdminHandler := adminHandler.NewHttpProductAdminHandler(prodAdminService)

	prodAdmin := router.Group("/v1/admin/product")
	prodAdmin.POST("/", prodAdminHandler.CreateProduct)
}
