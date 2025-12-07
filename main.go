package main

import (
	"github.com/gin-gonic/gin"

	"BookInventoryService/database"
	"BookInventoryService/model"
	"BookInventoryService/route"

	"log"
)

func main() {
	database.ConnectDB()

	database.DB.AutoMigrate(&model.Book{}, &model.Category{}, &model.User{})
	log.Println("Database migrated!")

	r := gin.Default()

	route.RegisterRoutes(r)

	r.Run(":8080")
}