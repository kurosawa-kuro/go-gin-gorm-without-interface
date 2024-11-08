package main

import (
	"fmt"
	"log"
	"os"

	"go-gin-gorm-without-interface/internal/controller"
	"go-gin-gorm-without-interface/internal/database"
	"go-gin-gorm-without-interface/internal/models"
	"go-gin-gorm-without-interface/internal/services"

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

	// サービスを初期化
	micropostService := services.NewMicropostService(db)

	// サービスをコントローラに注入
	micropostController := controller.NewMicropostController(micropostService)

	// ルーティング
	r.POST("/microposts", micropostController.Create)
	r.GET("/microposts", micropostController.GetAll)
	r.GET("/microposts/:id", micropostController.GetByID)
	r.PUT("/microposts/:id", micropostController.Update)
	r.DELETE("/microposts/:id", micropostController.Delete)

	r.Run()
}
