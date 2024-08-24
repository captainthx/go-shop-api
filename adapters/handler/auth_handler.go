package handler

import (
	"go-shop-api/core/domain"
	"go-shop-api/core/ports"
	"net/http"

	"github.com/gin-gonic/gin"
)

// all the handler interfaces are defined here

type HttpUserHandler struct {
	service ports.AuthService
}

func NewHttpAuthHandler(service ports.AuthService) *HttpUserHandler {
	return &HttpUserHandler{service: service}
}

func (h *HttpUserHandler) SignUp(c *gin.Context) {
	user := new(domain.User)
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	err := h.service.CreateUser(user)
	if err != nil {
		handlerError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}

func (h *HttpUserHandler) SignIn(c *gin.Context) {

	user := new(domain.User)
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	result, err := h.service.LogIn(user.Username, user.Password)

	if err != nil {
		handlerError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": result,
	})
}
