package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Micropostモデルの定義
type Micropost struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
}

func main() {
	// .env.development ファイルを読み込む
	if err := godotenv.Load(".env.development"); err != nil {
		log.Fatal("Error loading .env.development file")
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// マイグレーション
	db.AutoMigrate(&Micropost{})

	r := gin.Default()

	// CRUD エンドポイント
	r.POST("/microposts", func(c *gin.Context) {
		var micropost Micropost
		if err := c.ShouldBindJSON(&micropost); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		db.Create(&micropost)
		c.JSON(200, micropost)
	})

	r.GET("/microposts", func(c *gin.Context) {
		var microposts []Micropost
		db.Find(&microposts)
		c.JSON(200, microposts)
	})

	r.GET("/microposts/:id", func(c *gin.Context) {
		var micropost Micropost
		if err := db.First(&micropost, c.Param("id")).Error; err != nil {
			c.JSON(404, gin.H{"error": "Record not found!"})
			return
		}
		c.JSON(200, micropost)
	})

	r.PUT("/microposts/:id", func(c *gin.Context) {
		var micropost Micropost
		if err := db.First(&micropost, c.Param("id")).Error; err != nil {
			c.JSON(404, gin.H{"error": "Record not found!"})
			return
		}

		var newMicropost Micropost
		if err := c.ShouldBindJSON(&newMicropost); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		db.Model(&micropost).Updates(newMicropost)
		c.JSON(200, micropost)
	})

	r.DELETE("/microposts/:id", func(c *gin.Context) {
		var micropost Micropost
		if err := db.First(&micropost, c.Param("id")).Error; err != nil {
			c.JSON(404, gin.H{"error": "Record not found!"})
			return
		}
		db.Delete(&micropost)
		c.JSON(200, gin.H{"message": "Deleted successfully"})
	})

	r.Run()
}
