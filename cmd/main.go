package main

import (
	"users/internal/handler"

	_ "users/docs" // ให้ Swag สร้างเอกสารใน Folder docs โดยอัตโนมัติ

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Simple API Example
// @version         1.0
// @description     This is a simple example of using Gin with Swagger.
// @host            localhost:8080
// @BasePath        /api/v1
func main() {
	r := gin.Default()

	// Swagger endpoint
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// User API routes
	api := r.Group("/api/v1")
	{
		// Middleware สำหรับตรวจสอบ Bearer Token
		api.Use(handler.AuthMiddleware())

		// Routes สำหรับ API
		api.PUT("/users/:user_id", handler.UpdateUser) // อัพเดทข้อมูลผู้ใช้
	}

	// Start server
	r.Run(":8080")
}
