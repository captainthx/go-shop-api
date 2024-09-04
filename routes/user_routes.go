package routes

import (
	"go-shop-api/adapters/handler"
	"go-shop-api/adapters/repository"
	"go-shop-api/core/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(router *gin.Engine, db *gorm.DB) {
	userRepo := repository.NewUserRepositoryDB(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewHttpUserHandler(userService)

	// user router

	user := router.Group("/v1/user")

	user.PUT("/avatar", userHandler.UpdateUserAvatar)
}
