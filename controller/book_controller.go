package controller

import (
	"net/http"
	"strconv"
	"time"

	"BookInventoryService/database"
	"BookInventoryService/model"
	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	createdBy, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized request",
		})
		return
	}

	title := c.PostForm("title")
	description := c.PostForm("description")
	imageURL := c.PostForm("image_url")
	releaseYearStr := c.PostForm("release_year")
	priceStr := c.PostForm("price")
	totalPageStr := c.PostForm("total_page")
	categoryIDStr := c.PostForm("category_id")

	if title == "" || description == "" || totalPageStr == "" || categoryIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Title, description, total_page, and category_id are required",
		})
		return
	}

	releaseYear, err := strconv.Atoi(releaseYearStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid release_year format",
		})
		return
	}

	if releaseYear < 1980 || releaseYear > 2024 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "release_year must be between 1980 and 2024",
		})
		return
	}

	price, _ := strconv.Atoi(priceStr)
	totalPage, err := strconv.Atoi(totalPageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid total_page format",
		})
		return
	}

	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid category_id format",
		})
		return
	}

	var category model.Category
	if err := database.DB.First(&category, categoryID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid category_id — category not found",
		})
		return
	}

	thickness := "tipis"
	if totalPage > 100 {
		thickness = "tebal"
	}

	book := model.Book{
		Title:       title,
		Description: description,
		ImageURL:    imageURL,
		ReleaseYear: releaseYear,
		Price:       price,
		TotalPage:   totalPage,
		Thickness:   thickness,
		CategoryID:  uint(categoryID),
		CreatedAt:   time.Now(),
		CreatedBy:   createdBy.(string),
		ModifiedAt:  time.Now(),
		ModifiedBy:  createdBy.(string),
	}

	if err := database.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create book",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Book created successfully",
		"data":    book,
	})
}

func GetAllBooks(c *gin.Context) {
	var books []model.Book

	if err := database.DB.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch books",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Books retrieved successfully",
		"data":    books,
	})
}