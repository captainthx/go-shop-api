package service

import (
	"go-shop-api/adapters/errs"
	"go-shop-api/core/domain"
	"go-shop-api/core/model/response"
	request "go-shop-api/core/model/resquest"
	"go-shop-api/core/ports"
	"go-shop-api/logs"
	"math"

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

	product, err := c.repo.FindByProductId(request.ProductID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errs.NewBadRequestError("Product not found")
		}
		return errs.NewUnexpectedError(err.Error())
	}

	if request.Quantity <= 0 {
		return errs.NewBadRequestError("Quantity must be greater than 0")
	}
	if product.Quantity-request.Quantity < 0 {
		return errs.NewBadRequestError("Product quantity is not enough")
	}

	cartItem := domain.CartItem{
		UserID:    request.UserID,
		ProductID: request.ProductID,
		Quantity:  request.Quantity,
	}
	if err := c.repo.CreateCartItem(&cartItem); err != nil {
		return errs.NewUnexpectedError(err.Error())
	}

	product.Quantity -= request.Quantity

	err = c.repo.UpdateProductQuantity(product)
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
	productMap, err := getProductMap(products)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError(err.Error())
	}
	// create a slice of cart item response
	var cartItemResponse = mapCartItemsToResponse(cartItems, productMap)

	return cartItemResponse, nil
}

// UpdateCartItem implements ports.CartItemService.
func (c *cartItemService) UpdateCartItem(request *request.UpdQauntityCartItem) (*response.CartItemResponse, error) {

	// validate
	cartItem, err := c.repo.FindByCartId(request.CartItemId)
	if err != nil {
		logs.Error(err)
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewBadRequestError("Cart item not found")
		}
		return nil, errs.NewUnexpectedError(err.Error())
	}

	product, err := c.repo.FindByProductId(cartItem.ProductID)
	if err != nil {
		logs.Error(err)
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewBadRequestError("Product not found")
		}
		return nil, errs.NewUnexpectedError(err.Error())
	}

	if math.Signbit(float64(request.Quantity)) {
		if cartItem.Quantity+request.Quantity < 0 {
			return nil, errs.NewBadRequestError("Cannot reduce quantity below 0")
		}
		cartItem.Quantity += request.Quantity
		err = c.repo.UpdateCartItem(cartItem)
		if err != nil {
			logs.Error(err)
			return nil, errs.NewUnexpectedError(err.Error())
		}

		product.Quantity += int(math.Abs(float64(request.Quantity)))
		err = c.repo.UpdateProductQuantity(product)
		if err != nil {
			logs.Error(err)
			return nil, errs.NewUnexpectedError(err.Error())
		}
	} else {
		if product.Quantity-request.Quantity < 0 {
			return nil, errs.NewBadRequestError("product quantity is not enough")
		}

		cartItem.Quantity += request.Quantity
		err = c.repo.UpdateCartItem(cartItem)
		if err != nil {
			logs.Error(err)
			return nil, errs.NewUnexpectedError(err.Error())
		}

		product.Quantity -= request.Quantity
		err := c.repo.UpdateProductQuantity(product)
		if err != nil {
			logs.Error(err)
			return nil, errs.NewUnexpectedError(err.Error())
		}
	}

	updatedCartItem, err := c.repo.FindByCartId(request.CartItemId)

	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError(err.Error())
	}

	// response
	cartItemResponse := response.CartItemResponse{
		ID: cartItem.ID,
		Product: response.ProductResponse{
			ID:            product.ID,
			Name:          product.Name,
			Price:         product.Price,
			Quantity:      updatedCartItem.Quantity,
			ProductImages: mapProductImagesToResponse(product.ProductImage),
		},
		UpdatedAt: updatedCartItem.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	return &cartItemResponse, nil
}

// DeleteCartItem implements ports.CartItemService.
func (c *cartItemService) DeleteCartItem(cartId uint) error {
	cartItem, err := c.repo.FindByCartId(cartId)

	if err != nil {
		logs.Error(err)
		if err == gorm.ErrRecordNotFound {
			return errs.NewBadRequestError("Cart item not found")
		}
		return errs.NewUnexpectedError(err.Error())
	}

	product, err := c.repo.FindByProductId(cartItem.ProductID)
	if err != nil {
		logs.Error(err)
		if err == gorm.ErrRecordNotFound {
			return errs.NewBadRequestError("Product not found")
		}
		return errs.NewUnexpectedError(err.Error())
	}

	product.Quantity += cartItem.Quantity

	if err := c.repo.UpdateProductQuantity(product); err != nil {
		logs.Error(err)
		return errs.NewUnexpectedError(err.Error())
	}

	if err := c.repo.DeleteCartItem(cartId); err != nil {
		logs.Error(err)
		return errs.NewUnexpectedError(err.Error())
	}
	return nil
}

// Helper function to get product map
func getProductMap(products []domain.Product) (map[uint]domain.Product, error) {
	productMap := make(map[uint]domain.Product)
	for _, product := range products {
		productMap[product.ID] = product
	}
	return productMap, nil
}

// Helper function to map cart items to response
func mapCartItemsToResponse(cartItems []domain.CartItem, productMap map[uint]domain.Product) []response.CartItemResponse {
	var cartItemResponse = make([]response.CartItemResponse, 0, len(cartItems))

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
	return cartItemResponse
}
