package handler

import (
	"go-shop-api/core/domain"
	request "go-shop-api/core/model/resquest"
	"go-shop-api/core/ports"
	"go-shop-api/logs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type httpUserHandler struct {
	service ports.UserService
}

func NewHttpUserHandler(service ports.UserService) *httpUserHandler {
	return &httpUserHandler{service: service}
}

func (h *httpUserHandler) UpdateUserAvatar(c *gin.Context) {
	request := new(request.UpdateUserAvatarRequest)
	if err := c.ShouldBindJSON(request); err != nil {
		HandlerError(c, err)
		return
	}
	user := c.MustGet("user").(*domain.User)

	request.UserId = user.ID
	err := h.service.UpdateUserAvatar(request)
	if err != nil {
		logs.Error(err)
		HandlerError(c, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "update avatar success",
	})
}
