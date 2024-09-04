package routes

import (
	"go-shop-api/adapters/handler"
	"go-shop-api/core/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterFileRoutes(router *gin.Engine, db *gorm.DB) {

	fileService := service.NewFileService()
	fileHandler := handler.NewHttpFileHandler(fileService)

	file := router.Group("/v1/file")
	{

		file.POST("/upload", fileHandler.UploadFile)
		file.GET("/serve/:fileName", fileHandler.ServeFile)
	}
}
