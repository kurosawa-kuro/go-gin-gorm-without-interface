package server

import (
	"go-gin-gorm-without-interface/internal/config"
	"go-gin-gorm-without-interface/internal/database"
	"go-gin-gorm-without-interface/internal/routes"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config *config.Config
	router *gin.Engine
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		config: cfg,
		router: gin.Default(),
	}
}

func (s *Server) Run() error {
	// データベース初期化
	db, err := database.InitDB()
	if err != nil {
		return err
	}

	// ルーティングの設定
	routes.SetupRoutes(s.router, db)

	// サーバー起動
	return s.router.Run()
}
