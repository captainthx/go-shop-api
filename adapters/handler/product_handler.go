package handler

import (
	"go-shop-api/core/ports"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type httpProductHandler struct {
	service ports.ProductService
}

func NewHttpProductHandler(service ports.ProductService) *httpProductHandler {
	return &httpProductHandler{service: service}
}

func (h *httpProductHandler) GetProductList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	sort := c.Query("sort")

	result, err := h.service.GetProductList(page, limit, sort)
	if err != nil {
		HandlerError(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}
