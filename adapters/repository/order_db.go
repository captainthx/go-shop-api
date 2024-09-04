package repository

import (
	"go-shop-api/core/domain"
	"go-shop-api/core/ports"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type orderRepositoryDB struct {
	db *gorm.DB
}

func NewOrderRepositoryDB(db *gorm.DB) ports.OrderRepository {
	return &orderRepositoryDB{db: db}
}

// CreateOrder implements ports.OrderRepository.
func (o *orderRepositoryDB) CreateOrder(order *domain.Order) error {
	err := o.db.Create(&order).Error
	if err != nil {
		return err
	}
	return nil
}

// CreateOrderItems implements ports.OrderRepository.
func (o *orderRepositoryDB) CreateOrderItems(orderItem []domain.OrderItem) error {
	err := o.db.Create(&orderItem).Error
	if err != nil {
		return err
	}
	return nil
}

// FindOrderByUserId implements ports.OrderRepository.
func (o *orderRepositoryDB) FindOrderByUserId(userId uint) ([]domain.Order, error) {
	var orders []domain.Order
	err := o.db.Model(&domain.Order{}).Where("user_id = ?", userId).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// FindOrderByUserIdAndStatus implements ports.OrderRepository.
func (o *orderRepositoryDB) FindOrderByUserIdAndStatus(userId uint, orderStatus domain.OrderStatus) ([]domain.Order, error) {
	var orders []domain.Order

	err := o.db.Model(&domain.Order{}).Where("user_id = ? AND status = ?", userId, orderStatus).Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}

// FindCartItemByCartId implements ports.OrderRepository.
func (o *orderRepositoryDB) FindCartItemByUserId(userId uint) ([]domain.CartItem, error) {
	var cartItems []domain.CartItem
	err := o.db.Model(&domain.CartItem{}).Where("user_id = ?", userId).Find(&cartItems).Error
	if err != nil {
		return nil, err
	}
	return cartItems, nil
}

// FindProductByIds implements ports.OrderRepository.
func (o *orderRepositoryDB) FindProductByIds(productIds []uint) ([]domain.Product, error) {
	var products []domain.Product

	err := o.db.Where("id IN (?)", productIds).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

// FindProudctById implements ports.OrderRepository.
func (o *orderRepositoryDB) FindProudctById(productId uint) (*domain.Product, error) {
	var product domain.Product
	err := o.db.Model(&domain.Product{}).Where("id = ?", productId).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// FindOrderByID implements ports.OrderRepository.
func (o *orderRepositoryDB) FindOrderByID(orderId uint) (*domain.Order, error) {
	var order domain.Order
	err := o.db.Model(&domain.Order{}).Where("id =? ", orderId).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// FindOrderByStatus implements ports.OrderRepository.
func (o *orderRepositoryDB) FindOrderByStatus(orderStatus domain.OrderStatus) ([]domain.Order, error) {
	var orders []domain.Order
	err := o.db.Model(&domain.Order{}).Where("status = ?", orderStatus).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// FindOrderItemByUserId implements ports.OrderRepository.
func (o *orderRepositoryDB) FindOrderItemByOrderNumber(orderNumber uuid.UUID) ([]domain.OrderItem, error) {
	var orderItems []domain.OrderItem
	err := o.db.Model(&domain.OrderItem{}).Where("order_number =? ", orderNumber).Find(&orderItems).Error
	if err != nil {
		return nil, err
	}
	return orderItems, nil
}

// DeleteCartItemByUserId implements ports.OrderRepository.
func (o *orderRepositoryDB) DeleteCartItemByUserId(userId uint) error {
	err := o.db.Where("user_id = ?", userId).Delete(&domain.CartItem{}).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateOrder implements ports.OrderRepository.
func (o *orderRepositoryDB) UpdateOrder(order *domain.Order) error {
	err := o.db.Model(&domain.Order{}).Where("id = ?", order.ID).Updates(&order).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateProductQuantityById implements ports.OrderRepository.
func (o *orderRepositoryDB) UpdateProductQuantityById(product *domain.Product) error {
	err := o.db.Model(&domain.Product{}).Where("id = ?", product.ID).Updates(&product).Error
	if err != nil {
		return err
	}
	return nil
}
