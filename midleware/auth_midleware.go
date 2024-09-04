package midleware

import (
	"go-shop-api/core/domain"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RequireAuth(c *gin.Context) {
	user := &domain.User{}
	authHeader := c.GetHeader("Authorization")

	// Check if the Authorization header is missing or empty
	if authHeader == "" || len(authHeader) <= len("Bearer ") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	// Extract the token from the Authorization header
	token := authHeader[len("Bearer "):]
	// Parse the token
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}
	// get token claims
	claims := parsedToken.Claims.(jwt.MapClaims)
	// Extract user claims
	user.ID = uint(claims["auth"].(float64))
	roleStr := claims["role"].(string)
	user.Role = domain.Role(roleStr)

	// Store the user data in the Gin context
	c.Set("user", user)
	// Proceed to the next middleware or handler
	c.Next()
}

func AdminOnly(c *gin.Context) {
	user := c.MustGet("user").(*domain.User)
	if user.Role != domain.Admin {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "Forbidden",
		})
		return
	}
	c.Next()
}
