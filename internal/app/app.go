package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/canter-tech/car-service/config"
	"github.com/canter-tech/car-service/internal/metrics"
	middleware "github.com/canter-tech/car-service/internal/middleware/http"
	"github.com/canter-tech/car-service/internal/repository/inmem"
	"github.com/canter-tech/car-service/internal/services"
	v1 "github.com/canter-tech/car-service/internal/transport/http/v1"
	"github.com/canter-tech/car-service/pkg/httpserver"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
)

func Run(cfg *config.Config) {
	logger := zerolog.New(os.Stdout)

	repo := inmem.NewCarStore()
	service := services.NewCarService(repo)

	handler := chi.NewRouter()
	handler.Use(middleware.HTTPMetrics)
	handler.Use(chiMiddleware.Recoverer)
	handler.Use(chiMiddleware.Timeout(5 * time.Second))

	v1.NewRouter(handler, service)

	httpServer := httpserver.New(handler, httpserver.Addr(cfg.HttpConfig.Address))
	httpServer.Start()

	metricsServer := httpserver.New(metrics.New(),
		httpserver.Addr(cfg.MetricsConfig.Address),
		httpserver.WriteTimeout(time.Minute),
	)
	metricsServer.Start()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	var err error

	select {
	case s := <-interrupt:
		logger.Info().Msg("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		logger.Error().Err(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	case err = <-metricsServer.Notify():
		logger.Error().Err(fmt.Errorf("app - Run - metricsServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		logger.Error().Err(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	err = metricsServer.Shutdown()
	if err != nil {
		logger.Error().Err(fmt.Errorf("app - Run - metricsServer.Shutdown: %w", err))
	}
}
