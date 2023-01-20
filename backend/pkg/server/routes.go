package server

import (
	"backend/pkg/handlers"

	"github.com/go-chi/chi/v5"
)

func (s Server) SetupRoutes() {

	s.Mux.Get("/health", handlers.Health)

	s.Mux.Route("/search", func(r chi.Router) {
		r.Get("/{temp}", handlers.SearchEmail)
	})

}
