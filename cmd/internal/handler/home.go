package handler

import (
	"net/http"

	"github.com/kakoi-to-pirat/go-app/cmd/internal/view/home"
)

type homeHandler struct{}

func (h homeHandler) handleIndex(w http.ResponseWriter, r *http.Request) error {
	return home.Index().Render(r.Context(), w)
}
