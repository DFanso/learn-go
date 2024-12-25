package main

import (
	"github.com/dfanso/learn-go/database"
	"github.com/dfanso/learn-go/routes"
	"github.com/gin-gonic/gin"

	// Swagger files
	_ "github.com/dfanso/learn-go/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Go CRUD API
// @version         1.0
// @description     This is a simple CRUD API for managing books.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /api

func main() {

	// Initialize the database
	database.Connect()

	// Create a Gin router with default middleware (logger and recovery)
	router := gin.Default()

	// Setup API routes
	routes.SetupRoutes(router)

	// Swagger endpoint (http://localhost:8080/swagger/index.html)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := router.Run(":8080"); err != nil {
		panic("Failed to run server: " + err.Error())
	}
}
