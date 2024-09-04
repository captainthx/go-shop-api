package request

import (
	"go-shop-api/core/domain"
	"time"

	"github.com/google/uuid"
)

// Request is a struct that contains the request body.

type NewCartItemRequest struct {
	UserID    uint `json:"user_id"`
	ProductID uint `json:"productId" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required"`
}

type UpdQauntityCartItem struct {
	CartItemId uint `json:"cartItemId"`
	Quantity   int  `json:"quantity"`
}

type NewOrderReuqest struct {
	UserID   uint    `json:"userId"`
	TotalPay float64 `json:"totalPay"`
}
type FindOrderByStatusRequest struct {
	Status domain.OrderStatus `json:"status"`
	UserID uint               `json:"userId"`
}

type NewTransactionRequest struct {
	OrderId     uint      `json:"orderId"`
	OrderNumber uuid.UUID `json:"orderNumber"`
	Amount      float64
}

type UpdateTransactionRequest struct {
	OrderNumber uuid.UUID `json:"orderNumber"`
	PayTime     time.Time `json:"payTime"`
}

type UpdateUserAvatarRequest struct {
	UserId   uint   `json:"userId"`
	ImageUrl string `json:"image_url"`
}
