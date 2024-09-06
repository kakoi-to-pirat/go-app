package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi/v5"
	"github.com/kakoi-to-pirat/go-app/cmd/internal/config"
	"github.com/kakoi-to-pirat/go-app/cmd/internal/handler"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	cfg := config.NewConfig()
	router := chi.NewRouter()

	home := handler.HomeHandler{}

	router.Get("/", handler.NewHandler(home.HandleIndex))
	router.Get("/about", handler.NewHandler(home.HandleAbout))
	router.Handle("/assets/*", http.StripPrefix("/assets", http.FileServer(http.Dir(cfg.AssetsDir))))

	s := http.Server{
		Addr:    cfg.ServerAddr,
		Handler: router,
	}

	go func() {
		<-ctx.Done()
		slog.Info("shutting down server")
		s.Shutdown(ctx)
	}()

	slog.Info("Starting server!", slog.String("PORT", cfg.ServerAddr))

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
