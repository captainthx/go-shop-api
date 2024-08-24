package handler

import (
	"go-shop-api/adapters/errs"

	"github.com/gin-gonic/gin"
)

func HandlerError(c *gin.Context, err error) {

	switch e := err.(type) {
	case errs.AppError:
		c.JSON(e.Code, gin.H{
			"error": e.Message,
		})
	case error:
		c.JSON(500, gin.H{
			"error": e.Error(),
		})
	}
}
