package ports

import (
	"go-shop-api/common"
	"go-shop-api/core/domain"
	"go-shop-api/core/model/response"
	request "go-shop-api/core/model/resquest"
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

type ProductRepository interface {
	FindAll(pagination *common.Pagination) (*common.Pagination, error)
}

type ProductService interface {
	GetProductList(page, limit int, sort string) (*common.Pagination, error)
}

type CartItemRepository interface {
	FindByUser(user *domain.User) ([]domain.CartItem, error)
	FindByUserId(userId uint) (*domain.User, error)
	FindByProductId(productId uint) (*domain.Product, error)
	CreateCartItem(cartItem *domain.CartItem) error
	DeleteCartItem(cartItem *domain.CartItem) error
}

type CartItemService interface {
	GetCartItemList(user *domain.User) ([]domain.CartItem, error)
	AddCartItem(request *request.NewCartItemRequest) error
}
