package handler

import (
	"log/slog"
	"net/http"
)

type handlerFunc func(http.ResponseWriter, *http.Request) error

func NewHandler(h handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			handleError(w, r, err)
		}
	}
}

func handleError(w http.ResponseWriter, r *http.Request, err error) {
	slog.Error("error during request", slog.String("err", err.Error()))
}
