package app

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/reson-xu/english-platform/internal/adapters/httpapi"
	"github.com/reson-xu/english-platform/internal/adapters/httpapi/handlers"
	"github.com/reson-xu/english-platform/internal/adapters/storage/memory"
	user "github.com/reson-xu/english-platform/internal/modules/user"
	"github.com/reson-xu/english-platform/internal/platform/config"
	"github.com/reson-xu/english-platform/internal/platform/logger"
	"github.com/reson-xu/english-platform/internal/platform/security"
)

func RunAPI() error {
	cfg := config.Load()
	log := logger.New()
	slog.SetDefault(log)

	userStore := memory.NewUserStore()
	passwords := security.NewPasswordService()
	tokens := security.NewTokenService(cfg.JWTSecret, cfg.AccessTokenTTL)
	userService := user.NewService(userStore, passwords, tokens)
	userHandler := handlers.NewUserHandler(userService)

	server := &http.Server{
		Addr:              cfg.APIAddr,
		Handler:           httpapi.NewRouter(userHandler, log),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	return runHTTPServer(server, log)
}

func runHTTPServer(server *http.Server, log *slog.Logger) error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	errCh := make(chan error, 1)
	go func() {
		log.Info("api server listening", "addr", server.Addr)
		errCh <- server.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			return err
		}

		return nil
	case err := <-errCh:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return err
	}
}
