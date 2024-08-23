package service

import (
	"go-shop-api/adapters/errs"
	"go-shop-api/config"
	"go-shop-api/core/ports"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type fileService struct{}

func NewFileService() ports.FileService {
	return &fileService{}
}

// UpLoadFile implements ports.FileService.
func (f *fileService) UpLoadFile(file multipart.FileHeader, c *gin.Context) (*ports.UpLodaFileResponse, error) {

	if file.Size == 0 {
		return nil, errs.NewBadRequestError("file is empty")
	}

	// check if the file size is greater than 2 MB
	if file.Size > 2<<20 {
		return nil, errs.NewBadRequestError("file size must be less 2MB")
	}

	fileType := file.Header.Get("Content-Type")
	var fileExtension string
	switch fileType {
	case "image/jpeg":
		fileExtension = ".jpg"
	case "image/png":
		fileExtension = ".png"
	case "image/webp":
		fileExtension = ".webp"
	default:
		return nil, errs.NewBadRequestError("file type not supported")
	}

	filename := uuid.New().String() + fileExtension
	filePath := filepath.Join(config.UploadPath, filename)

	err := c.SaveUploadedFile(&file, filePath)
	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}
	// Construct the file URL
	fileUrl := config.ImageBaseUrl + config.ImageBasePath + filename

	uplodaResponse := &ports.UpLodaFileResponse{
		FileName: filename,
		FileUrl:  fileUrl,
		Size:     float32(file.Size),
	}

	return uplodaResponse, nil
}

// ServeFile implements ports.FileService.
func (f *fileService) ServeFile(fileName string) (string, error) {
	// Construct the file path
	filePath := filepath.Join(config.UploadPath, fileName)

	// Check if the file exists
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return "", errs.NewNotFoundError("file not found")
		}
		return "", errs.NewBadRequestError(err.Error())
	}

	return filePath, nil
}
