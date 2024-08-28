package handler

import (
	"go-shop-api/core/domain"
	request "go-shop-api/core/model/resquest"
	"go-shop-api/core/ports"
	"go-shop-api/logs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type httpCartItemHandler struct {
	service ports.CartItemService
}

func NewHttpCartItemHandler(service ports.CartItemService) *httpCartItemHandler {
	return &httpCartItemHandler{service}
}

func (h *httpCartItemHandler) AddCartItem(c *gin.Context) {
	request := new(request.NewCartItemRequest)
	if err := c.ShouldBindJSON(request); err != nil {
		HandlerError(c, err)
		return
	}
	user := c.MustGet("user").(*domain.User)
	request.UserID = user.ID

	err := h.service.AddCartItem(request)
	if err != nil {
		logs.Error(err)
		HandlerError(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Cart item added successfully",
	})
}

func (h *httpCartItemHandler) GetCartItems(c *gin.Context) {
	user := c.MustGet("user").(*domain.User)

	cartItems, err := h.service.GetCartItemList(user)
	if err != nil {
		HandlerError(c, err)
		return
	}

	c.JSON(http.StatusOK, cartItems)
}
