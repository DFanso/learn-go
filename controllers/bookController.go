package controllers

import (
	"log"
	"net/http"

	"github.com/dfanso/learn-go/api"
	"github.com/dfanso/learn-go/database"
	"github.com/dfanso/learn-go/models"
	"github.com/gin-gonic/gin"
)

// GetBooks godoc
// @Summary      Retrieve all books
// @Description  Get a list of all books
// @Tags         books
// @Produce      json
// @Success      200  {array}   models.Book
// @Failure      500  {object}  api.ErrorResponse
// @Router       /books [get]
func GetBooks(c *gin.Context) {
	Context := c
	if Context != nil {
		log.Printf("Context: %v", Context.Request.URL)
	}

	var books []models.Book
	result := database.DB.Find(&books)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Error: result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

// CreateBook godoc
// @Summary      Create a new book
// @Description  Add a new book to the database
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        book  body      models.Book  true  "Book to create"
// @Success      201   {object}  models.Book
// @Failure      400   {object}  api.ErrorResponse
// @Failure      500   {object}  api.ErrorResponse
// @Router       /books [post]
func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{Error: err.Error()})
		return
	}
	result := database.DB.Create(&book)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Error: result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, book)
}

// GetBook godoc
// @Summary      Retrieve a single book
// @Description  Get details of a book by ID
// @Tags         books
// @Produce      json
// @Param        id   path      int  true  "Book ID"
// @Success      200  {object}  models.Book
// @Failure      404  {object}  api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /books/{id} [get]
func GetBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	result := database.DB.First(&book, id)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			c.JSON(http.StatusNotFound, api.ErrorResponse{Error: "Book not found"})
		} else {
			c.JSON(http.StatusInternalServerError, api.ErrorResponse{Error: result.Error.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, book)
}

// UpdateBook godoc
// @Summary      Update a book
// @Description  Update details of a book by ID
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id    path      int         true  "Book ID"
// @Param        book  body      models.Book true  "Updated book information"
// @Success      200   {object}  models.Book
// @Failure      400   {object}  api.ErrorResponse
// @Failure      404   {object}  api.ErrorResponse
// @Failure      500   {object}  api.ErrorResponse
// @Router       /books/{id} [put]
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	result := database.DB.First(&book, id)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			c.JSON(http.StatusNotFound, api.ErrorResponse{Error: "Book not found"})
		} else {
			c.JSON(http.StatusInternalServerError, api.ErrorResponse{Error: result.Error.Error()})
		}
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{Error: err.Error()})
		return
	}

	saveResult := database.DB.Save(&book)
	if saveResult.Error != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Error: saveResult.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

// DeleteBook godoc
// @Summary      Delete a book
// @Description  Remove a book from the database by ID
// @Tags         books
// @Produce      json
// @Param        id   path      int  true  "Book ID"
// @Success      200  {object}  api.MessageResponse
// @Failure      404  {object}  api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /books/{id} [delete]
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	result := database.DB.First(&book, id)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			c.JSON(http.StatusNotFound, api.ErrorResponse{Error: "Book not found"})
		} else {
			c.JSON(http.StatusInternalServerError, api.ErrorResponse{Error: result.Error.Error()})
		}
		return
	}

	deleteResult := database.DB.Delete(&book)
	if deleteResult.Error != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Error: deleteResult.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, api.MessageResponse{Message: "Book deleted successfully"})
}
