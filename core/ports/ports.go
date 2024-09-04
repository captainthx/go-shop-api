package ports

import (
	"go-shop-api/common"
	"go-shop-api/core/domain"
	"go-shop-api/core/model/response"
	request "go-shop-api/core/model/resquest"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	FindByUserId(id uint) (*domain.User, error)
	FindOrderByUser(user *domain.User) ([]domain.Order, error)
	UpdateAvartar(user *domain.User) error
}

type UserService interface {
	UpdateUserAvatar(request *request.UpdateUserAvatarRequest) error
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
	FindCartItemByUserId(userId uint) ([]domain.CartItem, error)
	FindByUserId(userId uint) (*domain.User, error)
	FindByProductId(productId uint) (*domain.Product, error)
	FindByProductIds(productIds []uint) ([]domain.Product, error)
	CreateCartItem(cartItem *domain.CartItem) error
	DeleteCartItem(cartItemId uint) error
	UpdateCartItem(cartItem *domain.CartItem) error
	UpdateProductQuantity(product *domain.Product) error
	FindByCartId(cartId uint) (*domain.CartItem, error)
}

type CartItemService interface {
	GetCartItemList(user *domain.User) ([]response.CartItemResponse, error)
	AddCartItem(request *request.NewCartItemRequest) error
	UpdateCartItem(request *request.UpdQauntityCartItem) (*response.CartItemResponse, error)
	DeleteCartItem(cartId uint) error
}

type OrderRepository interface {
	CreateOrder(*domain.Order) error
	CreateOrderItems(orderItem []domain.OrderItem) error
	FindOrderByUserId(userId uint) ([]domain.Order, error)
	FindOrderByUserIdAndStatus(userId uint, orderStatus domain.OrderStatus) ([]domain.Order, error)
	FindOrderByID(orderId uint) (*domain.Order, error)
	FindOrderByStatus(orderStatus domain.OrderStatus) ([]domain.Order, error)
	FindCartItemByUserId(userId uint) ([]domain.CartItem, error)
	FindProductByIds(productIds []uint) ([]domain.Product, error)
	FindProudctById(productId uint) (*domain.Product, error)
	FindOrderItemByOrderNumber(orderNumber uuid.UUID) ([]domain.OrderItem, error)
	DeleteCartItemByUserId(userId uint) error
	UpdateOrder(order *domain.Order) error
	UpdateProductQuantityById(product *domain.Product) error
}

type OrderService interface {
	CreateOrder(request *request.NewOrderReuqest) (*response.OrderResponse, error)
	GetOrderByStatus(request *request.FindOrderByStatusRequest) ([]response.OrderHistoryResponse, error)
	GetOrderHistory(user *domain.User) ([]response.OrderHistoryResponse, error)
	CancelOrder(orderId uint) error
}

type TransactionRepository interface {
	CreateTransaction(payment *domain.Transaction) error
	UpdateTransaction(payment *domain.Transaction) error
	UpdateOrder(order *domain.Order) error
	FindTransactionByOrderNumber(orderNumber uuid.UUID) (*domain.Transaction, error)
	FindOrderByOrderNumber(orderNumber uuid.UUID) (*domain.Order, error)
}

type TransactionService interface {
	CreateTransaction(request *request.NewTransactionRequest) error
	UpdateTransaction(request *request.UpdateTransactionRequest) (*domain.Transaction, error)
}
