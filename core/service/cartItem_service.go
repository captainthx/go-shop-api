package service

import (
	"go-shop-api/adapters/errs"
	"go-shop-api/core/domain"
	"go-shop-api/core/model/response"
	request "go-shop-api/core/model/resquest"
	"go-shop-api/core/ports"
	"go-shop-api/logs"

	"gorm.io/gorm"
)

type cartItemService struct {
	repo ports.CartItemRepository
}

func NewCartItemService(repo ports.CartItemRepository) ports.CartItemService {
	return &cartItemService{repo: repo}
}

// AddCartItem implements ports.CartItemService.
func (c *cartItemService) AddCartItem(request *request.NewCartItemRequest) error {

	if _, err := c.repo.FindByUserId(request.UserID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return errs.NewBadRequestError("User not found")
		}
		return errs.NewUnexpectedError(err.Error())
	}

	if _, err := c.repo.FindByProductId(request.ProductID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return errs.NewBadRequestError("Product not found")
		}
		return errs.NewUnexpectedError(err.Error())
	}

	if request.Quantity <= 0 {
		return errs.NewBadRequestError("Quantity must be greater than 0")
	}

	cartItem := domain.CartItem{
		UserID:    request.UserID,
		ProductID: request.ProductID,
		Quantity:  request.Quantity,
	}
	err := c.repo.CreateCartItem(&cartItem)
	if err != nil {
		return errs.NewUnexpectedError(err.Error())
	}

	return nil
}

// GetCartItemList implements ports.CartItemService.
func (c *cartItemService) GetCartItemList(user *domain.User) ([]response.CartItemResponse, error) {
	cartItems, err := c.repo.FindCartItemByUserId(user.ID)
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	if len(cartItems) == 0 {
		return []response.CartItemResponse{}, nil
	}

	productIds := make([]uint, 0, len(cartItems))
	for _, cartItem := range cartItems {
		productIds = append(productIds, cartItem.ProductID)
	}

	products, err := c.repo.FindByProductIds(productIds)

	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError(err.Error())
	}

	// create a map of product id to product
	productMap := make(map[uint]domain.Product)
	for _, product := range products {
		productMap[product.ID] = product
	}
	// create a slice of cart item response
	var cartItemResponse = make([]response.CartItemResponse, 0, len(cartItems))

	// map cart item to response
	for _, cartItem := range cartItems {
		if product, ok := productMap[cartItem.ProductID]; ok {
			cartItemResponse = append(cartItemResponse, response.CartItemResponse{
				ID: cartItem.ID,
				Product: response.ProductResponse{
					ID:            product.ID,
					Name:          product.Name,
					Price:         product.Price,
					Quantity:      cartItem.Quantity,
					ProductImages: mapProductImagesToResponse(product.ProductImage),
				},
				CreatedAt: cartItem.CreatedAt.Format("2006-01-02 15:04:05"),
			})
		}
	}
	return cartItemResponse, nil
}
