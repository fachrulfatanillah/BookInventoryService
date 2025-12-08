package route

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"BookInventoryService/controller"
	"BookInventoryService/middleware"

)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Book Inventory Service API",
		})
	})

	r.POST("api/users/register", controller.CreateUser)
	r.POST("api/users/login", controller.Login)

	auth := r.Group("/api", middleware.AuthMiddleware())
	{
		auth.POST("/categories", controller.CreateCategory)
		auth.GET("/categories", controller.GetAllCategories)
		auth.GET("/categories/:id", controller.GetCategoryByID)
		auth.PUT("/categories/:id", controller.UpdateCategory)
	}
}
