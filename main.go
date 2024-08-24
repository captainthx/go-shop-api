package main

import (
	"fmt"
	"go-shop-api/adapters/handler"
	adminHandler "go-shop-api/adapters/handler/admin"
	"go-shop-api/adapters/repository"
	adminRepository "go-shop-api/adapters/repository/admin"
	"go-shop-api/config"
	"go-shop-api/core/domain"
	"go-shop-api/core/service"
	adminService "go-shop-api/core/service/admin"
	"go-shop-api/logs"
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
		logs.Error(err)
	}
	// Initialize the configuration
	config.Init()
	// create dsn for mysql gorm
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DbUsername, config.DbPassword, config.DbHost, config.DbPort, config.DbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		logs.Error(err)
	}

	// Create or modify the database tables based on the model structs found in the imported package
	err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&domain.User{},
		&domain.Category{},
		&domain.Product{},
		&domain.ProductImage{},
		&domain.Cart{},
		&domain.CartItem{},
		&domain.Order{},
		&domain.OrderItem{},
		&domain.Transaction{})
	if err != nil {
		logs.Error(err)
		return
	}
	initRoute(db)

}

func initRoute(db *gorm.DB) {
	gin.SetMode(config.Mode)

	router := gin.Default()
	router.MaxMultipartMemory = 2 << 20 // 2 MB

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// file router

	fileService := service.NewFileService()
	fileHandler := handler.NewHttpFileHandler(fileService)

	file := router.Group("/v1/file")

	file.POST("/upload", fileHandler.UploadFile)
	file.GET("/serve/:fileName", fileHandler.ServeFile)

	// auth user router
	authRepo := repository.NewauthRepositoryDB(db)
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewHttpAuthHandler(authService)

	auth := router.Group("/v1/auth")

	auth.POST("/sign-up", authHandler.SignUp)
	auth.POST("/sign-in", authHandler.SignIn)

	// auth admin router
	authAdminRepo := adminRepository.NewAuthAdminRepositoryDB(db)
	authAdminService := adminService.NewAuthAdminService(authAdminRepo)
	authAdminHandler := adminHandler.NewHttpAdminHandler(authAdminService)

	authAdmin := router.Group("/v1/admin/auth")

	authAdmin.POST("/sign-up", authAdminHandler.SignUp)
	authAdmin.POST("/sign-in", authAdminHandler.SignIn)

	// product router
	prodcutCus := router.Group("/v1/product")

	prodcutCus.GET("/", func(ctx *gin.Context) {})

	// Protected routes
	router.Use(RequireAuth)

	router.GET("/", func(c *gin.Context) {
		user := c.MustGet("user").(*domain.User)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello," + user.Role,
			"user":    user,
		})
	})

	// Admin routes
	router.Use(adminOnly)
	router.GET("/admin", func(c *gin.Context) {
		user := c.MustGet("user").(*domain.User)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Admin",
			"user":    user,
		})
	})

	err := router.Run(":" + config.ServerPort)
	if err != nil {
		logs.Error(err)
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

func adminOnly(c *gin.Context) {
	user := c.MustGet("user").(*domain.User)

	if user.Role != domain.Role("admin") {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "Forbidden",
		})
		return
	}
	c.Next()
}
