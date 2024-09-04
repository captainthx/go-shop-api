package routes

import (
	adminHandler "go-shop-api/adapters/handler/admin"
	adminRepository "go-shop-api/adapters/repository/admin"
	adminService "go-shop-api/core/service/admin"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAuthAdminRoutes(router *gin.Engine, db *gorm.DB) {
	authAdminRepo := adminRepository.NewAuthAdminRepositoryDB(db)
	authAdminService := adminService.NewAuthAdminService(authAdminRepo)
	authAdminHandler := adminHandler.NewHttpAdminHandler(authAdminService)

	authAdmin := router.Group("/v1/admin/auth")

	authAdmin.POST("/sign-up", authAdminHandler.SignUp)
	authAdmin.POST("/sign-in", authAdminHandler.SignIn)
}
