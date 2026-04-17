package main

import (
	"backend/internal/config"
	"backend/internal/logger"
	"backend/internal/server"
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

func gracefulShutdown(fiberServer *server.FiberServer, done chan bool) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	logger.Log.Info("shutting down gracefully, press Ctrl+C again to force")
	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := fiberServer.ShutdownWithContext(ctx); err != nil {
		logger.Log.Errorf("Server forced to shutdown with error: %v", err)
	}

	logger.Log.Info("Server exiting")

	done <- true
}

func main() {
	if err := logger.Init(); err != nil {
		fmt.Printf("failed to initialize logger: %v\n", err)
		return
	}
	defer logger.Sync()

	server := server.New()

	done := make(chan bool, 1)

	go func() {
		port := config.GetPort()
		err := server.Listen(fmt.Sprintf(":%d", port))
		if err != nil {
			logger.Log.Fatalf("http server error: %s", err)
		}
	}()

	go gracefulShutdown(server, done)

	<-done
	logger.Log.Info("Graceful shutdown complete.")
}
