package main

import (
	"fmt"
	"go-shop-api/adapters/handler"
	"go-shop-api/adapters/repository"
	"go-shop-api/core/domain"
	"go-shop-api/core/service"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	initTimezone()

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	databaseName := os.Getenv("DB_NAME")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	// create dsn for mysql gorm
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, databaseName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		panic(err)
	}

	// Create or modify the database tables based on the model structs found in the imported package
	err = db.AutoMigrate(&domain.User{}, &domain.Product{}, &domain.ProductImage{}, &domain.Order{}, &domain.OrderItem{}, &domain.Transaction{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
		return // exit if migration fails
	}

	initGin(db)

}

func initGin(db *gorm.DB) {

	router := gin.Default()

	authRepo := repository.NewauthRepositoryDB(db)
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewHttpAuthHandler(authService)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	auth := router.Group("/v1/auth")

	auth.POST("/sign-up", authHandler.SignUp)
	auth.POST("/sign-in", authHandler.SignIn)

	// Protected routes

	router.Use(RequireAuth)

	router.GET("/", func(c *gin.Context) {
		user := c.MustGet("user").(*domain.User)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello," + user.Role,
			"user":    user,
		})
	})

	err := router.Run(":3000")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
		panic(err)
	}

}

func initTimezone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}

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
