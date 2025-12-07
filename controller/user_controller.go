package controller

import (
	"net/http"
	"time"

	"BookInventoryService/database"
	"BookInventoryService/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	createdBy := c.PostForm("created_by")

	if username == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "username and password are required",
		})
		return
	}

	if len(password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "password must be at least 6 characters",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to hash password",
		})
		return
	}

	user := model.User{
		Username:   username,
		Password:   string(hashedPassword),
		CreatedAt:  time.Now(),
		CreatedBy:  createdBy,
		ModifiedAt: time.Now(),
		ModifiedBy: createdBy,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user": gin.H{
			"username":   user.Username,
			"created_at": user.CreatedAt,
		},
	})
}