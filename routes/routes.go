package routes

import (
	"github.com/dfanso/learn-go/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		books := api.Group("/books")
		{
			books.GET("/", controllers.GetBooks)         // GET /api/books/
			books.POST("/", controllers.CreateBook)      // POST /api/books/
			books.GET("/:id", controllers.GetBook)       // GET /api/books/:id
			books.PUT("/:id", controllers.UpdateBook)    // PUT /api/books/:id
			books.DELETE("/:id", controllers.DeleteBook) // DELETE /api/books/:id
		}

		users := api.Group("/users")
		{
			users.GET("/", controllers.GetUsers)         // GET /api/users/
			users.POST("/", controllers.CreateUser)      // POST /api/users/
			users.GET("/:id", controllers.GetUser)       // GET /api/users/:id
			users.PUT("/:id", controllers.UpdateUser)    // PUT /api/users/:id
			users.DELETE("/:id", controllers.DeleteUser) // DELETE /api/users/:id

		}
	}
}
