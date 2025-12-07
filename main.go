package main

import (
	"github.com/gin-gonic/gin"

	"BookInventoryService/database"
	"BookInventoryService/model"

	"log"
)

func main() {
	database.ConnectDB()

	database.DB.AutoMigrate(&model.Book{}, &model.Category{}, &model.User{})
	log.Println("✅ Database migrated!")

	r := gin.Default()
	r.Run(":8080")
}