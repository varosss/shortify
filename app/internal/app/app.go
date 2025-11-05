package app

import (
	"log"
	"shortify/internal/auth"
	"shortify/internal/config"
	"shortify/internal/db"
	"shortify/internal/links"
	"shortify/internal/server"
	"time"
)

func Run() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	conn, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close(conn)

	authHandler := auth.NewAuthHandler(conn)
	linksHandler := links.NewLinksHandler(conn, cfg.RedisAddr)

	router := server.NewRouter(authHandler, linksHandler, auth.NewJWTManager(time.Hour))

	log.Printf("Server running on port %s", cfg.ServerPort)

	return router.Run(":" + cfg.ServerPort)
}
