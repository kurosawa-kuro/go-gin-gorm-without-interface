package main

import (
	"log"

	"go-gin-gorm-without-interface/internal/config"
	"go-gin-gorm-without-interface/internal/server"
)

func main() {
	// 設定の読み込み
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// サーバーの初期化と起動
	app := server.NewServer(cfg)
	if err := app.Run(); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
