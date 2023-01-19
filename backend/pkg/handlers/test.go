package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Test(mux chi.Router) {
	mux.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
}
