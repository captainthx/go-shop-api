package response

import (
	"go-shop-api/core/domain"

	"github.com/google/uuid"
)

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type UpLodaFileResponse struct {
	FileName string  `json:"fileName"`
	FileUrl  string  `json:"fileUrl"`
	Size     float32 `json:"size"`
}

type ProductResponse struct {
	ID            uint                   `json:"id"`
	Name          string                 `json:"name"`
	Price         float64                `json:"price"`
	Quantity      int                    `json:"quantity"`
	CreatedAt     string                 `json:"createdAt"`
	ProductImages []ProductImageResponse `json:"productImage"`
}

type ProductImageResponse struct {
	ProductID uint   `json:"productId"`
	Url       string `json:"url"`
}

type CartItemResponse struct {
	ID        uint            `json:"id"`
	Product   ProductResponse `json:"product"`
	CreatedAt string          `json:"createdAt"`
	UpdatedAt string          `json:"updatedAt"`
}

type OrderResponse struct {
	ID          uint               `json:"id"`
	OrderNumber uuid.UUID          `json:"orderNumber"`
	Status      domain.OrderStatus `json:"status"`
	TotalPay    float64            `json:"totalPay"`
}

type OrderHistoryResponse struct {
	ID          uint                `json:"id"`
	OrderNumber uuid.UUID           `json:"orderNumber"`
	Status      domain.OrderStatus  `json:"status"`
	TotalPay    float64             `json:"totalPay"`
	OrderItems  []OrderItemResponse `json:"orderItems"`
}

type OrderItemResponse struct {
	ProductName string  `json:"productName"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}
