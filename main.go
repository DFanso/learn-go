package main

import (
	"fmt"
	"log"

	"github.com/dfanso/learn-go/config"
	"github.com/dfanso/learn-go/database"
	"github.com/dfanso/learn-go/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

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

	godotenv.Load()
	port := config.GoDotEnvVariable("GIN_PORT")

	// Initialize the database
	database.Connect()

	// Create a Gin router with default middleware (logger and recovery)
	router := gin.Default()

	// Set trusted proxies for the reverse proxy
	err := router.SetTrustedProxies([]string{config.GoDotEnvVariable("GIN_ALLOW_ORIGIN")})
	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	// Setup API routes
	routes.SetupRoutes(router)

	// Swagger endpoint (http://localhost:8080/swagger/index.html)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		panic("Failed to run server: " + err.Error())
	}
}
