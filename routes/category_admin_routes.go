package routes

import (
	adminHandler "go-shop-api/adapters/handler/admin"
	adminRepository "go-shop-api/adapters/repository/admin"
	adminService "go-shop-api/core/service/admin"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Category Admin
// @Description Category Admin
// @Tags  Category Admin
// @Param request body domain.Category true "Create Category Request"
// @Shemes http
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Router /v1/admin/category [post]
// @Router /v1/admin/category [get]
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func RegisterCategoryAdminRoutes(router *gin.Engine, db *gorm.DB) {
	categoryRepo := adminRepository.NewCategoryAdminRepositoryDB(db)
	categoryService := adminService.NewCategoryAdminService(categoryRepo)
	categoryHandler := adminHandler.NewHttpCategoryAdminHandler(categoryService)

	{
		categoryRoute := router.Group("/v1/admin/category")
		categoryRoute.POST("/", categoryHandler.CreateCateory)
		categoryRoute.GET("/", categoryHandler.GetCategoryList)
	}
}
