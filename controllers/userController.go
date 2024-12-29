package controllers

import (
	"fmt"
	"net/http"

	"github.com/dfanso/learn-go/api"
	"github.com/dfanso/learn-go/database"
	"github.com/dfanso/learn-go/models"
	"github.com/gin-gonic/gin"
)

// GetUsers godoc
// @Summary      Retrieve all Users
// @Description  Get a list of all Users
// @Tags         users
// @Produce      json
// @Success      200  {array}   models.User
// @Failure      500  {object}  api.ErrorResponse
// @Router       /users [get]
func GetUsers(c *gin.Context) {
	var users []models.User
	result := database.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Error: result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// CreateUser godoc
// @Summary      Create a new User
// @Description  Add a new user to the database
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      models.User  true  "User to create"
// @Success      201   {object}  models.User
// @Failure      400   {object}  api.ErrorResponse
// @Failure      500   {object}  api.ErrorResponse
// @Router       /users [post]
func CreateUser(c *gin.Context) {
	fmt.Print("Create User")
	var user models.User
	fmt.Println("User: ", user)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{Error: err.Error()})
		return
	}
	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Error: result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

// GetUser godoc
// @Summary      Retrieve a single user
// @Description  Get details of a user by ID
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  models.User
// @Failure      404  {object}  api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /users/{id} [get]
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			c.JSON(http.StatusNotFound, api.ErrorResponse{Error: "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, api.ErrorResponse{Error: result.Error.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser godoc
// @Summary      Update a user
// @Description  Update details of a user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id    path      int         true  "User ID"
// @Param        user  body      models.User true  "Updated user information"
// @Success      200   {object}  models.User
// @Failure      400   {object}  api.ErrorResponse
// @Failure      404   {object}  api.ErrorResponse
// @Failure      500   {object}  api.ErrorResponse
// @Router       /users/{id} [put]
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			c.JSON(http.StatusNotFound, api.ErrorResponse{Error: "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, api.ErrorResponse{Error: result.Error.Error()})
		}
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{Error: err.Error()})
		return
	}

	saveResult := database.DB.Save(&user)
	if saveResult.Error != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Error: saveResult.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary      Delete a user
// @Description  Remove a user from the database by ID
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  api.MessageResponse
// @Failure      404  {object}  api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			c.JSON(http.StatusNotFound, api.ErrorResponse{Error: "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, api.ErrorResponse{Error: result.Error.Error()})
		}
		return
	}

	deleteResult := database.DB.Delete(&user)
	if deleteResult.Error != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Error: deleteResult.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, api.MessageResponse{Message: "User deleted successfully"})
}
