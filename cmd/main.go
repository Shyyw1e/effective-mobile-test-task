// @title People API
// @version 1.0
// @description Тестовое задание: API для работы с обогащёнными людьми
// @host localhost:8080
// @BasePath /api
package main

import (
	"log/slog"
	"os"

	"github.com/Shyyw1e/effective-mobile-test-task/internal/client"
	"github.com/Shyyw1e/effective-mobile-test-task/internal/config"
	"github.com/Shyyw1e/effective-mobile-test-task/internal/handler"
	"github.com/Shyyw1e/effective-mobile-test-task/internal/repository"
	"github.com/Shyyw1e/effective-mobile-test-task/internal/service"
	"github.com/Shyyw1e/effective-mobile-test-task/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/Shyyw1e/effective-mobile-test-task/docs"


)

func main() {
	logger.Log = logger.New(slog.LevelDebug)
	cfg := config.LoadConfig()

	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		os.Exit(1)
	}


	personRepo := repository.NewPersonRepository(db)
	client := client.NewRealClient()
	enricher := service.NewEnrichService(personRepo, client, logger.Log)
	h := handler.NewHandler(personRepo, enricher)

	r := h.InitRoutes()


	logger.Log.Info("server started", "port", cfg.Port)

	if err := r.Run(":" + cfg.Port); err != nil {
		logger.Log.Error("failed to start server", "error", err)
		os.Exit(1)
	}
}
