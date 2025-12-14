package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/unicdev/go-shop/internal/config"
	"github.com/unicdev/go-shop/internal/database"
	"github.com/unicdev/go-shop/internal/interfaces"
	"github.com/unicdev/go-shop/internal/logger"
	"github.com/unicdev/go-shop/internal/providers"
	"github.com/unicdev/go-shop/internal/server"
	"github.com/unicdev/go-shop/internal/services"
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

	defer func() {
		if err := mainDB.Close(); err != nil {
			log.Fatal().Err(err).Msg("failed to close main database")
		}
	}()

	gin.SetMode(cfg.Server.GinMode)

	authService := services.NewAuthService(db, cfg)
	productService := services.NewProductService(db)
	userService := services.NewUserService(db)

	var uploadProvider interfaces.UploadProvider
	if cfg.Upload.UploadProvider == "s3" {
		uploadProvider = providers.NewS3Provider(cfg)
	} else {
		uploadProvider = providers.NewLocalUploadProvider(cfg.Upload.Path)
	}

	uploadService := services.NewUploadService(uploadProvider)
	cartService := services.NewCartService(db)
	orderService := services.NewOrderService(db)

	srv := server.New(cfg,
		db,
		&log,
		authService,
		productService,
		userService,
		uploadService,
		cartService,
		orderService,
	)
	router := srv.SetupRoutes()

	httpServer := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	go func() {
		log.Info().Str("port", cfg.Server.Port).Msg("Starting HTTP server")
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("failed to start HTTP server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info().Msg("Shutting down API server")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("failed to shutdown HTTP server")
		return
	}
	log.Info().Msg("API server shutdown successfully")
}
