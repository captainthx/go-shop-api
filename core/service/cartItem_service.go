package service

import (
	"go-shop-api/adapters/errs"
	"go-shop-api/core/domain"
	request "go-shop-api/core/model/resquest"
	"go-shop-api/core/ports"

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
func (c *cartItemService) GetCartItemList(user *domain.User) ([]domain.CartItem, error) {
	return nil, nil
}
