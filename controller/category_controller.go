package controller

import (
	"net/http"
	"time"

	"BookInventoryService/database"
	"BookInventoryService/model"
	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	createdBy, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not authorized",
		})
		return
	}

	name := c.PostForm("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Category name is required",
		})
		return
	}

	category := model.Category{
		Name:       name,
		CreatedAt:  time.Now(),
		CreatedBy:  createdBy.(string),
		ModifiedAt: time.Now(),
		ModifiedBy: createdBy.(string),
	}

	if err := database.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create category",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Category created successfully",
		"data":    category,
	})
}