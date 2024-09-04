package handler

import (
	request "go-shop-api/core/model/resquest"
	"go-shop-api/core/ports"

	"github.com/gin-gonic/gin"
)

type httpTransactionHandler struct {
	service ports.TransactionService
}

func NewHttpTransactionHandler(service ports.TransactionService) *httpTransactionHandler {
	return &httpTransactionHandler{service: service}
}

func (h *httpTransactionHandler) CreateTransaction(c *gin.Context) {
	request := new(request.NewTransactionRequest)
	if err := c.ShouldBindJSON(request); err != nil {
		HandlerError(c, err)
		return
	}

	err := h.service.CreateTransaction(request)
	if err != nil {
		HandlerError(c, err)
		return
	}
}

func (h *httpTransactionHandler) UpdateTransaction(c *gin.Context) {
	request := new(request.UpdateTransactionRequest)
	if err := c.ShouldBindJSON(request); err != nil {
		HandlerError(c, err)
		return
	}

	_, err := h.service.UpdateTransaction(request)
	if err != nil {
		HandlerError(c, err)
		return
	}
	c.JSON(200, gin.H{"message": "transaction success"})
}
