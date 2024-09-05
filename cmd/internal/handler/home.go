package handler

import (
	"io"
	"net/http"
)

type homeHandler struct{}

func (h homeHandler) handleIndex(w http.ResponseWriter, r *http.Request) error {
	if _, err := io.WriteString(w, "Hellom world!"); err != nil {
		return err
	}

	return nil
}
