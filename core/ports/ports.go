package ports

import (
	"go-shop-api/core/domain"
	"go-shop-api/core/model/response"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	CreateUser(user *domain.User) error
	LogIn(username string, password string) (*response.LoginResponse, error)
}

type AuthRepository interface {
	Create(user *domain.User) error
	FindByUserName(username string) (*domain.User, error)
}

type UserRepository interface {
	FindByID(id uint) (*domain.User, error)
	FindOrderByUser(user *domain.User) ([]domain.Order, error)
}

type FileService interface {
	UpLoadFile(file multipart.FileHeader, c *gin.Context) (*response.UpLodaFileResponse, error)
	ServeFile(fileName string) (string, error)
}
