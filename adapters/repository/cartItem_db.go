package repository

import (
	"go-shop-api/core/domain"
	"go-shop-api/core/ports"

	"gorm.io/gorm"
)

type cartItemRepositoryDB struct {
	db *gorm.DB
}

func NewCartItemRepositoryDB(db *gorm.DB) ports.CartItemRepository {
	return &cartItemRepositoryDB{db: db}
}

// CreateCartItem implements ports.CartItemRepository.
func (c *cartItemRepositoryDB) CreateCartItem(cartItem *domain.CartItem) error {
	err := c.db.Create(&cartItem).Error
	if err != nil {
		return err
	}
	return nil
}

// FindByProductId implements ports.CartItemRepository.
func (c *cartItemRepositoryDB) FindByProductId(productId uint) (*domain.Product, error) {
	var product domain.Product
	err := c.db.Model(&domain.Product{}).Where("id = ?", productId).First(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

// FindByProductIds implements ports.CartItemRepository.
func (c *cartItemRepositoryDB) FindByProductIds(productIds []uint) ([]domain.Product, error) {
	var products []domain.Product

	err := c.db.Preload("ProductImage").Where("id IN (?)", productIds).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

// FindByUserId implements ports.CartItemRepository.
func (c *cartItemRepositoryDB) FindByUserId(userId uint) (*domain.User, error) {
	var user domain.User
	err := c.db.Model(&domain.User{}).Where("id = ?", userId).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// DeleteCartItem implements ports.CartItemRepository.
func (c *cartItemRepositoryDB) DeleteCartItem(cartItem *domain.CartItem) error {
	return nil
}

// FindByUser implements ports.CartItemRepository.
func (c *cartItemRepositoryDB) FindCartItemByUserId(userId uint) ([]domain.CartItem, error) {

	var cartItems []domain.CartItem
	err := c.db.Where(&domain.CartItem{UserID: userId}).Find(&cartItems).Error
	if err != nil {
		return nil, err
	}
	return cartItems, nil
}
