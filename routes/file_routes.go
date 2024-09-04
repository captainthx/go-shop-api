package routes

import (
	"go-shop-api/adapters/handler"
	"go-shop-api/core/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary File
// @Description File Upload
// @Shemes http
// @Tags File
// @Param file formData file true "File to upload"
// @Success 200 {object} string
// @Router /v1/file/upload [post]
// @Router /v1/file/serve/{fileName} [get]
func RegisterFileRoutes(router *gin.Engine, db *gorm.DB) {

	fileService := service.NewFileService()
	fileHandler := handler.NewHttpFileHandler(fileService)

	file := router.Group("/v1/file")
	{

		file.POST("/upload", fileHandler.UploadFile)
		file.GET("/serve/:fileName", fileHandler.ServeFile)
	}
}
