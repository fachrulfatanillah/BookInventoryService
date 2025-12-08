package controller

import (
	"net/http"
	"time"
	"strconv"

	"BookInventoryService/database"
	"BookInventoryService/model"
	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	createdBy, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized request",
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

func GetAllCategories(c *gin.Context) {
	_, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized request",
		})
		return
	}

	var categories []model.Category

	if err := database.DB.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve categories",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Categories fetched successfully",
		"data":         categories,
	})
}

func GetCategoryByID(c *gin.Context) {

	_, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized request",
		})
		return
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid category ID",
		})
		return
	}

	var category model.Category

	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Category not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Success",
		"category": category,
	})
}