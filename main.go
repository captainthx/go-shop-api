package main

import (
	"fmt"
	"go-shop-api/config"
	"go-shop-api/core/domain"
	"go-shop-api/logs"
	"go-shop-api/midleware"
	"go-shop-api/routes"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	initTimezone()

	err := godotenv.Load(".env")
	if err != nil {
		logs.Error(err)
	}
	// Initialize the configuration
	config.Init()

	// create dsn for mysql gorm
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DbUsername, config.DbPassword, config.DbHost, config.DbPort, config.DbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
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
		&domain.CartItem{},
		&domain.Order{},
		&domain.OrderItem{},
		&domain.Transaction{})
	if err != nil {
		logs.Error(err)
		return
	}
	router := initRoute(db)

	err = router.Run(":" + config.ServerPort)
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

func initRoute(db *gorm.DB) *gin.Engine {
	gin.SetMode(config.Mode)

	router := gin.Default()
	router.MaxMultipartMemory = 2 << 20 // 2 MB

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong!",
		})
	})

	// Public routes
	routes.RegisterFileRoutes(router, db)
	routes.RegisterAuthRoutes(router, db)
	routes.RegisterAuthAdminRoutes(router, db)
	routes.RegisterProductCusRoutes(router, db)

	// Protected routes
	router.Use(midleware.RequireAuth)
	routes.RegisterCartRoutes(router, db)
	routes.RegisterOrderRoutes(router, db)
	routes.RegisterTransactionRoutes(router, db)

	// Admin routes
	router.Use(midleware.AdminOnly)
	routes.RegisterProductAdminRoutes(router, db)
	routes.RegisterCategoryAdminRoutes(router, db)

	return router
}
