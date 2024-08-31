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

type httpOrderhandler struct {
	service ports.OrderService
}

func NewHttpOrderHandler(service ports.OrderService) *httpOrderhandler {
	return &httpOrderhandler{service: service}
}

func (h *httpOrderhandler) CreateOrder(c *gin.Context) {
	request := new(request.NewOrderReuqest)
	if err := c.ShouldBindJSON(request); err != nil {
		HandlerError(c, err)
		return
	}
	user := c.MustGet("user").(*domain.User)
	request.UserID = user.ID

	response, err := h.service.CreateOrder(request)
	if err != nil {
		logs.Error(err)
		HandlerError(c, err)
		return
	}
	c.JSON(http.StatusCreated, response)
}

func (h *httpOrderhandler) GetOrderHistoryList(c *gin.Context) {
	user := c.MustGet("user").(*domain.User)
	orderHistories, err := h.service.GetOrderHistory(user)
	if err != nil {
		logs.Error(err)
		HandlerError(c, err)
		return
	}
	c.JSON(http.StatusOK, orderHistories)
}

func (h *httpOrderhandler) CancelOrder(c *gin.Context) {
	orderId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		HandlerError(c, err)
		return
	}

	if err := h.service.CancelOrder(uint(orderId)); err != nil {
		logs.Error(err)
		HandlerError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Order canceled successfully",
	})
}
