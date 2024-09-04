package repository

import (
	"go-shop-api/core/domain"
	"go-shop-api/core/ports"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type transactionRepositoryDB struct {
	db *gorm.DB
}

func NewTransactionRepositoryDB(db *gorm.DB) ports.TransactionRepository {
	return &transactionRepositoryDB{db: db}
}

// CreateTransaction implements ports.TransactionRepository.
func (t *transactionRepositoryDB) CreateTransaction(payment *domain.Transaction) error {
	err := t.db.Create(payment).Error
	if err != nil {
		return nil
	}
	return nil
}

// FindTransactionByOrderNumber implements ports.TransactionRepository.
func (t *transactionRepositoryDB) FindTransactionByOrderNumber(orderNumber uuid.UUID) (*domain.Transaction, error) {
	var transaction domain.Transaction
	err := t.db.Where(&domain.Transaction{OrderNumber: orderNumber}).First(&transaction).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

// FindOrderByOrderNumber implements ports.TransactionRepository.
func (t *transactionRepositoryDB) FindOrderByOrderNumber(orderNumber uuid.UUID) (*domain.Order, error) {
	var order domain.Order
	err := t.db.Where(&domain.Order{OrderNumber: orderNumber}).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// UpdateOrder implements ports.TransactionRepository.
func (t *transactionRepositoryDB) UpdateOrder(order *domain.Order) error {
	err := t.db.Where(&domain.Order{OrderNumber: order.OrderNumber}).Updates(&order).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateTransaction implements ports.TransactionRepository.
func (t *transactionRepositoryDB) UpdateTransaction(payment *domain.Transaction) error {
	err := t.db.Where(&domain.Transaction{OrderNumber: payment.OrderNumber}).Updates(&payment).Error
	if err != nil {
		return nil
	}
	return nil
}
