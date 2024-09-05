package app

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kakoi-to-pirat/go-app/cmd/internal/config"
	"github.com/kakoi-to-pirat/go-app/cmd/internal/handler"
)

func Run(ctx context.Context) error {
	cfg := config.NewConfig()

	r := chi.NewRouter()
	handler.RegisterRoutes(r)

	s := http.Server{
		Addr:    cfg.ServerAddr,
		Handler: r,
	}

	go func() {
		<-ctx.Done()
		slog.Info("shutting down server")
		s.Shutdown(ctx)
	}()

	slog.Info("starting server!", slog.String("PORT", cfg.ServerAddr))

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
