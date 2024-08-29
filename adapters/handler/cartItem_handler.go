package handler

import (
	"go-shop-api/core/domain"
	request "go-shop-api/core/model/resquest"
	"go-shop-api/core/ports"
	"go-shop-api/logs"
	"net/http"
	"strconv"

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

func (h *httpCartItemHandler) UpdateCartItem(c *gin.Context) {
	request := new(request.UpdQauntityCartItem)
	if err := c.ShouldBindJSON(request); err != nil {
		HandlerError(c, err)
		return
	}

	cartItem, err := h.service.UpdateCartItem(request)
	if err != nil {
		logs.Error(err)
		HandlerError(c, err)
		return
	}
	c.JSON(http.StatusCreated, cartItem)
}

func (h *httpCartItemHandler) DeleteCartItem(c *gin.Context) {
	carItemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		HandlerError(c, err)
		return
	}
	if err := h.service.DeleteCartItem(uint(carItemId)); err != nil {
		logs.Error(err)
		HandlerError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Cart item deleted successfully",
	})
}
