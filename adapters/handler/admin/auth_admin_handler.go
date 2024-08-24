package adminHandler

import (
	"go-shop-api/adapters/handler"
	"go-shop-api/core/domain"
	adminPorts "go-shop-api/core/ports/admin"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpAdminHandler struct {
	service adminPorts.AuthAdminService
}

func NewHttpAdminHandler(service adminPorts.AuthAdminService) *HttpAdminHandler {
	return &HttpAdminHandler{service: service}
}

func (h *HttpAdminHandler) SignUp(c *gin.Context) {
	user := new(domain.User)
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	err := h.service.CreateAdmin(user)
	if err != nil {
		handler.HandlerError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Admin created successfully",
	})
}

func (h *HttpAdminHandler) SignIn(c *gin.Context) {

	user := new(domain.User)
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	result, err := h.service.LogIn(user.Username, user.Password)

	if err != nil {
		handler.HandlerError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": result,
	})
}
