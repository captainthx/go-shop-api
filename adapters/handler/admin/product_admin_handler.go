package adminHandler

import (
	"go-shop-api/adapters/handler"
	"go-shop-api/core/domain"
	adminPorts "go-shop-api/core/ports/admin"
	"net/http"

	"github.com/gin-gonic/gin"
)

type httpProducAdminHandler struct {
	service adminPorts.ProductAdminService
}

func NewHttpProductAdminHandler(service adminPorts.ProductAdminService) *httpProducAdminHandler {
	return &httpProducAdminHandler{service: service}
}

func (h *httpProducAdminHandler) CreateProduct(c *gin.Context) {
	product := new(domain.Product)
	if err := c.ShouldBindJSON(product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	err := h.service.CreateProduct(product)
	if err != nil {
		handler.HandlerError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Product created successfully",
	})

}
