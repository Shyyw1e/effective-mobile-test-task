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

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

)

func main() {
	cfg := config.LoadConfig()

	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		os.Exit(1)
	}

	logg := logger.New(slog.LevelDebug)

	personRepo := repository.NewPersonRepository(db)
	client := client.NewRealClient()
	enricher := service.NewEnrichService(personRepo, client, logg)
	h := handler.NewHandler(personRepo, enricher)

	r := h.InitRoutes()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	logg.Info("server started", "port", cfg.Port)

	if err := r.Run(":" + cfg.Port); err != nil {
		logg.Error("failed to start server", "error", err)
		os.Exit(1)
	}
}
