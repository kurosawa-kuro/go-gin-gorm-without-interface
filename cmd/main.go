package main

import (
	"fmt"
	"log"
	"os"

	"go-gin-gorm-without-interface/internal/controllers"
	"go-gin-gorm-without-interface/internal/database"
	"go-gin-gorm-without-interface/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	// configs ディレクトリからの .env ファイル読み込み
	err := godotenv.Load(fmt.Sprintf("configs/.env.%s", env))
	if err != nil {
		log.Fatalf("Error loading .env.%s file", env)
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
