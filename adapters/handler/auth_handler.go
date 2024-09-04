package handler

import (
	"go-shop-api/core/domain"
	"go-shop-api/core/ports"
	"net/http"

	"github.com/gin-gonic/gin"
)

// all the handler interfaces are defined here

type httpAuthHandler struct {
	service ports.AuthService
}

func NewHttpAuthHandler(service ports.AuthService) *httpAuthHandler {
	return &httpAuthHandler{service: service}
}

func (h *httpAuthHandler) SignUp(c *gin.Context) {
	user := new(domain.User)
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	err := h.service.CreateUser(user)
	if err != nil {
		HandlerError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}

func (h *httpAuthHandler) SignIn(c *gin.Context) {

	user := new(domain.User)
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	result, err := h.service.LogIn(user.Username, user.Password)

	if err != nil {
		HandlerError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": result,
	})
}
