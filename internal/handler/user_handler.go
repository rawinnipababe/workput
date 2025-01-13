package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

// Mock database
var users = map[int]*User{
	1: {ID: 1, Name: "รวินท์นิภา ดำรบูรณะกุลชัย"},
}

// @Summary Update user by ID
// @Description Update details of a user by ID
// @Tags Users
// @Accept  json
// @Produce json
// @Param   id    path      int     true  "User ID"
// @Param   user  body      User    true  "User data"
// @Success 200   {object}  User
// @Failure 400   {object}  ErrorResponse
// @Failure 404   {object}  ErrorResponse
// @Router  /users/{id} [put]
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"id": id, "name": "รวินท์นิภา ดำรบูรณะกุลชัย"})
}
func UpdateUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid user ID"})
		return
	}

	var updatedUser User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request body"})
		return
	}

	user, exists := users[id]
	if !exists {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "User not found"})
		return
	}

	// Update user details
	user.Name = updatedUser.Name

	c.JSON(http.StatusOK, user)
}
