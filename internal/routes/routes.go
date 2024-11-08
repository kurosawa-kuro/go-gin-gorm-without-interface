package routes

import (
	"go-gin-gorm-without-interface/internal/controllers"
	"go-gin-gorm-without-interface/internal/models"
	"go-gin-gorm-without-interface/internal/services"

	_ "go-gin-gorm-without-interface/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// マイグレーション
	db.AutoMigrate(&models.Micropost{})

	// Swagger設定
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// サービスとコントローラの初期化
	micropostService := services.NewMicropostService(db)
	micropostController := controllers.NewMicropostController(micropostService)

	// APIルート
	setupMicropostRoutes(r, micropostController)
}

func setupMicropostRoutes(r *gin.Engine, mc *controllers.MicropostController) {
	r.POST("/microposts", mc.Create)
	r.GET("/microposts", mc.GetAll)
	r.GET("/microposts/:id", mc.GetByID)
	r.PUT("/microposts/:id", mc.Update)
	r.DELETE("/microposts/:id", mc.Delete)
}
