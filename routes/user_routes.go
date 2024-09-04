package routes

import (
	"go-shop-api/adapters/handler"
	"go-shop-api/adapters/repository"
	"go-shop-api/core/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary User
// @Description User
// @Shemes http
// @Tags User
// @Param request body request.UpdateUserAvatarRequest true "Create User Request"
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} string
// @Router /v1/user/avatar [put]
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func RegisterUserRoutes(router *gin.Engine, db *gorm.DB) {
	userRepo := repository.NewUserRepositoryDB(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewHttpUserHandler(userService)

	// user router

	user := router.Group("/v1/user")

	user.PUT("/avatar", userHandler.UpdateUserAvatar)
}
