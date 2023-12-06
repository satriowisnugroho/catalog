package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/satriowisnugroho/catalog/internal/config"
	"github.com/satriowisnugroho/catalog/internal/handler/http/middleware"
	httpv1 "github.com/satriowisnugroho/catalog/internal/handler/http/v1"
	"github.com/satriowisnugroho/catalog/internal/parser"
	"github.com/satriowisnugroho/catalog/internal/repository/postgres"
	"github.com/satriowisnugroho/catalog/internal/usecase"
	"github.com/satriowisnugroho/catalog/pkg/httpserver"
	"github.com/satriowisnugroho/catalog/pkg/logger"
	pkgpostgres "github.com/satriowisnugroho/catalog/pkg/postgres"
)

func main() {
	cfg := config.NewConfig()

	l := logger.New(cfg.LogLevel)

	// Initialize postgres
	postgresDb, err := pkgpostgres.NewPostgres(&cfg.DatabaseConfig)
	if err != nil {
		l.Fatal(fmt.Errorf("app - api - postgres.NewPostgres: %w", err))
	}
	defer postgresDb.Db.Close()

	// Initialize repositories
	dbTransactionRepo := postgres.NewPostgresTransactionRepository(postgresDb.Db)
	productRepo := postgres.NewProductRepository(postgresDb.Db)

	// Initialize usecases
	productUsecase := usecase.NewProductUsecase(productRepo, dbTransactionRepo)

	// Initialize parsers
	productParser := parser.NewProductParser()

	// HTTP Server
	handler := gin.New()

	// Set middleware
	handler.Use(middleware.Tenant())

	// Set router
	httpv1.NewRouter(handler, l, productParser, productUsecase)
	httpServer := httpserver.New(handler, httpserver.Port(fmt.Sprint(cfg.Port)))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - api - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - api - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - api - httpServer.Shutdown: %w", err))
	}
}
