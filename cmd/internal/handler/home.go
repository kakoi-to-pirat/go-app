package handler

import (
	"net/http"

	"github.com/kakoi-to-pirat/go-app/cmd/internal/view/home"
)

type HomeHandler struct{}

func (h HomeHandler) HandleIndex(w http.ResponseWriter, r *http.Request) error {
	return home.Index().Render(r.Context(), w)
}

func (h HomeHandler) HandleAbout(w http.ResponseWriter, r *http.Request) error {
	return home.About().Render(r.Context(), w)
}
