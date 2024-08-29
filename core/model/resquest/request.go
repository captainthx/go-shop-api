package request

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
