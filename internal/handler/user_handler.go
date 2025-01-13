package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// User โครงสร้างข้อมูลผู้ใช้
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ErrorResponse โครงสร้างข้อมูลสำหรับข้อผิดพลาด
type ErrorResponse struct {
	Message string `json:"message"`
}

// UpdateUserRequest โครงสร้างสำหรับคำขอ PUT
type UpdateUserRequest struct {
	Name  string `json:"name" binding:"required"`        // ฟิลด์ที่จำเป็น
	Email string `json:"email" binding:"required,email"` // อีเมลต้องอยู่ในรูปแบบที่ถูกต้อง
}

// Mock database
var users = []User{
	{ID: 1, Name: "John Doe", Email: "john@example.com"},
	{ID: 2, Name: "Jane Smith", Email: "jane@example.com"},
}
var nextID = 3

// AuthMiddleware ตรวจสอบ Bearer Token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") || len(authHeader) < 8 {
			c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// @Summary      Update a user by ID
// @Description  Update a user's details by ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user_id  path      int              true  "User ID"
// @Param        user     body      UpdateUserRequest true  "User Information"
// @Success      200      {object}  User
// @Failure      400      {object}  ErrorResponse
// @Failure      401      {object}  ErrorResponse
// @Failure      404      {object}  ErrorResponse
// @Failure      409      {object}  ErrorResponse
// @Router       /users/{user_id} [put]
func UpdateUser(c *gin.Context) {
	// ดึง user_id จาก path
	idParam := c.Param("user_id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid user ID"})
		return
	}

	// อ่านและตรวจสอบ JSON Body
	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	// ค้นหาผู้ใช้
	var user *User
	for i := range users {
		if users[i].ID == id {
			user = &users[i]
			break
		}
	}

	if user == nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "User not found"})
		return
	}

	// ตรวจสอบว่า email ใหม่ซ้ำหรือไม่
	for _, u := range users {
		if u.Email == req.Email && u.ID != id {
			c.JSON(http.StatusConflict, ErrorResponse{Message: "Email already exists"})
			return
		}
	}

	// อัพเดทข้อมูลผู้ใช้
	user.Name = req.Name
	user.Email = req.Email

	// ส่งคืนข้อมูลผู้ใช้อัพเดทแล้ว
	c.JSON(http.StatusOK, user)
}
