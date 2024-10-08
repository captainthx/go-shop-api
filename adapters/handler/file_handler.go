package handler

import (
	"go-shop-api/core/ports"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpFileHandler struct {
	service ports.FileService
}

func NewHttpFileHandler(service ports.FileService) *HttpFileHandler {
	return &HttpFileHandler{service: service}
}

func (h *HttpFileHandler) UploadFile(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		HandlerError(c, err)
		return
	}

	result, err := h.service.UpLoadFile(*file, c)
	if err != nil {
		HandlerError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "File uploaded successfully",
		"result":  result,
	})
}

func (h *HttpFileHandler) ServeFile(c *gin.Context) {
	fileName := c.Param("fileName")
	filePath, err := h.service.ServeFile(fileName)
	if err != nil {
		HandlerError(c, err)
		return
	}

	// Serve the image file
	c.File(filePath)
}
