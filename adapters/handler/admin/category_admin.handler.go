package adminHandler

import (
	"go-shop-api/adapters/handler"
	"go-shop-api/core/domain"
	adminPorts "go-shop-api/core/ports/admin"
	"net/http"

	"github.com/gin-gonic/gin"
)

type httpCategoryAdminHandler struct {
	service adminPorts.CategoryAdminService
}

func NewHttpCategoryAdminHandler(service adminPorts.CategoryAdminService) *httpCategoryAdminHandler {
	return &httpCategoryAdminHandler{service: service}
}

func (h *httpCategoryAdminHandler) CreateCateory(c *gin.Context) {
	category := new(domain.Category)
	if err := c.ShouldBindJSON(category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	err := h.service.CreateCategory(category)
	if err != nil {
		handler.HandlerError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Category created successfully",
	})
}

func (h *httpCategoryAdminHandler) GetCategoryList(c *gin.Context) {
	result, err := h.service.GetCategory()
	if err != nil {
		handler.HandlerError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
