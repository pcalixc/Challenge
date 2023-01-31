package server

import (
	"backend/pkg/handlers"

	"github.com/go-chi/chi/v5"
)

func (s Server) SetupRoutes() {

	s.Mux.Get("/", handlers.Health)

	s.Mux.Route("/search", func(r chi.Router) {
		r.Get("/", handlers.SearchEmail)
		r.Get("/{temp}", handlers.SearchEmail)
	})

}
