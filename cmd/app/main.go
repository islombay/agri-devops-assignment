package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/islombay/agri-devops-assignment/internal/config"
	"github.com/islombay/agri-devops-assignment/internal/handler"
	"github.com/islombay/agri-devops-assignment/internal/logger"
	"github.com/islombay/agri-devops-assignment/internal/metrics"
	"github.com/islombay/agri-devops-assignment/internal/router"
	"github.com/islombay/agri-devops-assignment/internal/service"
	"github.com/islombay/agri-devops-assignment/pkg/health"
)

func main() {
	// Load config
	cfg := config.Load()

	// Init logger
	log := logger.New(logger.Config{
		Env:      cfg.Env,
		LogLevel: cfg.LogLevel,
	})

	log.Info("application starting",
		"env", cfg.Env,
		"port", cfg.Port,
	)

	// Init service layer
	agriService := service.NewAgricultureService()
	agriHandler := handler.NewAgricultureHandler(agriService)

	metrics.Register()

	// Router
	mux := http.NewServeMux()
	mux.Handle("/", router.NewRouter(agriHandler))
	mux.HandleFunc("/health/live", health.Liveness)
	mux.HandleFunc("/health/ready", health.Readiness)

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: mux,
	}

	// Start server
	go func() {
		log.Info("http server started", "addr", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("server error", "err", err)
		}
	}()

	// Graceful shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	<-shutdown
	log.Warn("shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Error("graceful shutdown failed", "err", err)
	}

	log.Info("server stopped gracefully")
}
