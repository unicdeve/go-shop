package main

import (
	"github.com/gin-gonic/gin"
	"github.com/unicdev/go-shop/internal/config"
	"github.com/unicdev/go-shop/internal/database"
	"github.com/unicdev/go-shop/internal/logger"
)

func main() {
	log := logger.New()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}

	db, err := database.New(&cfg.Database)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
	}

	mainDB, err := db.DB()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get main database")
	}
	defer mainDB.Close()
	gin.SetMode(cfg.Server.GinMode)

	log.Info().Msg("Staring API server")
}
