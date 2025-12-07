package route

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"BookInventoryService/controller"

)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Book Inventory Service API",
		})
	})

	r.POST("/api/users", controller.CreateUser)
}
