package main

import (
	"log"

	"go-gin-gorm-without-interface/controllers"
	"go-gin-gorm-without-interface/database"
	"go-gin-gorm-without-interface/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env.development"); err != nil {
		log.Fatal("Error loading .env.development file")
	}

	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// マイグレーション
	db.AutoMigrate(&models.Micropost{})

	r := gin.Default()

	mc := controllers.NewMicropostController(db)

	// ルーティング
	r.POST("/microposts", mc.Create)
	r.GET("/microposts", mc.GetAll)
	r.GET("/microposts/:id", mc.GetByID)
	r.PUT("/microposts/:id", mc.Update)
	r.DELETE("/microposts/:id", mc.Delete)

	r.Run()
}
