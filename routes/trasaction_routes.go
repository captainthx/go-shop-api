package routes

import (
	"go-shop-api/adapters/handler"
	"go-shop-api/adapters/repository"
	"go-shop-api/core/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Transaction
// @Description Transaction
// @Shemes http
// @Tags Transaction
// @Param request body request.NewTransactionRequest true "Create Transaction Request"
// @Param request body request.UpdateTransactionRequest true "Update Transaction Request"
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 201 {object} string
// @Router /v1/transaction [post]
// @Router /v1/transaction [put]
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func RegisterTransactionRoutes(router *gin.Engine, db *gorm.DB) {
	transactionRepo := repository.NewTransactionRepositoryDB(db)
	transactionService := service.NewTransactionService(transactionRepo)
	transactionHandler := handler.NewHttpTransactionHandler(transactionService)

	{
		transaction := router.Group("/v1/transaction")
		transaction.POST("/", transactionHandler.CreateTransaction)
		transaction.PUT("/", transactionHandler.UpdateTransaction)
	}
}
