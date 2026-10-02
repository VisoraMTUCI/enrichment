package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ppMTUCI/enrichment/internal/config"
	"github.com/ppMTUCI/enrichment/internal/handler"
)

type HandlersMap = map[string]http.HandlerFunc

func main() {
	// конфиг
	ctx := context.Background()

	cfg, err := config.InitConfig()
	if err != nil {
		slog.Error(fmt.Sprintf("Cannot init config: %s", err))
		os.Exit(1)
	}

	logger := slog.New(slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level: cfg.PrepareLogLevel(),
		},
	))

	logger.Info(fmt.Sprintf("App started with config:%+v", cfg))

	// кафка
	//
	// !FIXME
	//

	// сервера
	publicHandler := HandlersMap{
		"/event": handler.IndexHandler,
	}

	debugHandlers := HandlersMap{
		"/live":  handler.LiveHandler,
		"/ready": handler.ReadyHandler,
	}

	publicServer := handler.ServerFactory(cfg.PublicServer.Port, publicHandler)
	debugServer := handler.ServerFactory(cfg.DebugServer.Port, debugHandlers)

	servers := map[string]*http.Server{
		"public": publicServer,
		"debug":  debugServer,
	}

	for serverName, server := range servers {
		go func() {
			err := server.ListenAndServe()
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				logger.ErrorContext(ctx,
					"Error starting server",
					"server_name", serverName,
					"error", err.Error(),
				)
			}
		}()
	}

	// graceful shutdown
	stopChn := make(chan os.Signal, 1)
	signal.Notify(stopChn, syscall.SIGINT, syscall.SIGTERM)
	<-stopChn

	logger.InfoContext(ctx, "Shutting down by signal")

	shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	for serverName, server := range servers {
		if err := server.Shutdown(shutdownCtx); err != nil {
			logger.ErrorContext(
				ctx,
				"Error shutting down server",
				"server_name", serverName,
				"error", err.Error(),
			)
		}
	}

	logger.InfoContext(ctx, "Servers shut down gracefully")
}
