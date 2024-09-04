package routes

import (
	"go-shop-api/adapters/handler"
	"go-shop-api/adapters/repository"
	"go-shop-api/core/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Auth
// @Description Auth
// @Tags Auth
// @Param request body domain.User true "Create User Request"
// @Accept json
// @Produce json
// @Success 200 {object} response.LoginResponse
// @Router /v1/auth/sign-up [post]
// @Router /v1/auth/sign-in [post]
func RegisterAuthRoutes(router *gin.Engine, db *gorm.DB) {
	authRepo := repository.NewAuthRepositoryDB(db)
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewHttpAuthHandler(authService)

	auth := router.Group("/v1/auth")
	{
		auth.POST("/sign-up", authHandler.SignUp)
		auth.POST("/sign-in", authHandler.SignIn)
	}
}
