package app

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"studio/internal/config"
)

func Run(ctx context.Context) error {
	cfg := config.NewConfig()

	// TODO: to fiber
	s := http.Server{
		Addr: cfg.ServerAddr,
	}

	go func() {
		<-ctx.Done()
		slog.Info("Shutting down server")
		_ = s.Shutdown(ctx)
	}()

	slog.Info("Starting server", slog.String("addr", cfg.ServerAddr))
	if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
