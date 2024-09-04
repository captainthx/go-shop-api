package service

import (
	"go-shop-api/adapters/errs"
	"go-shop-api/core/domain"
	request "go-shop-api/core/model/resquest"
	"go-shop-api/core/ports"
	"go-shop-api/logs"
	"time"
)

type transactionService struct {
	repo ports.TransactionRepository
}

func NewTransactionService(repo ports.TransactionRepository) ports.TransactionService {
	return &transactionService{repo: repo}
}

// CreateTransaction implements ports.TransactionService.
func (t *transactionService) CreateTransaction(request *request.NewTransactionRequest) error {

	if request.Amount < 0 {
		return errs.NewBadRequestError("Amount must be greater than 0")
	}

	// wait for implement  pay ment gateway
	_, err := t.repo.FindOrderByOrderNumber(request.OrderNumber)
	if err != nil {
		return errs.NewNotFoundError("Order not found")
	}

	transaction := &domain.Transaction{
		Amount:      request.Amount,
		OrderNumber: request.OrderNumber,
		PayTime:     time.Now(),
	}

	if err := t.repo.CreateTransaction(transaction); err != nil {
		return errs.NewUnexpectedError("Could not create transaction")
	}
	return nil
}

// UpdateTransaction implements ports.TransactionService.
func (t *transactionService) UpdateTransaction(request *request.UpdateTransactionRequest) (*domain.Transaction, error) {

	order, err := t.repo.FindOrderByOrderNumber(request.OrderNumber)
	if err != nil {
		return nil, errs.NewNotFoundError("Order not found")
	}

	order.Status = domain.Success

	if err := t.repo.UpdateOrder(order); err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError("Could not update order")
	}

	transaction, err := t.repo.FindTransactionByOrderNumber(request.OrderNumber)

	if err != nil {
		logs.Error(err)
		return nil, errs.NewNotFoundError("Transaction not found")
	}

	transaction.PayTime = request.PayTime

	if err := t.repo.UpdateTransaction(transaction); err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError("Could not update transaction")
	}

	return nil, nil
}
