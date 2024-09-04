package routes

import (
	"go-shop-api/adapters/handler"
	"go-shop-api/adapters/repository"
	"go-shop-api/core/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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
