package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/canter-tech/car-service/config"
	v1 "github.com/canter-tech/car-service/internal/transport/http/v1"
	"github.com/canter-tech/car-service/pkg/httpserver"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
)

func Run(cfg *config.Config) {
	logger := zerolog.New(os.Stdout)

	handler := chi.NewRouter()
	handler.Use(middleware.Logger)
	handler.Use(middleware.RedirectSlashes)
	handler.Use(middleware.Recoverer)
	handler.Use(middleware.Timeout(5 * time.Second))
	v1.NewRouter(handler)

	httpServer := httpserver.New(handler)
	httpServer.Start()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	var err error

	select {
	case s := <-interrupt:
		logger.Info().Msg("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		logger.Error().Err(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		logger.Error().Err(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
