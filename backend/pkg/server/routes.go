package server

import (
	"backend/pkg/handlers"

	"github.com/go-chi/chi/v5"
)

func (s Server) SetupRoutes() {

	s.Mux.Get("/health", handlers.Health)

	s.Mux.Mount("/search", GetEmails(s))

}

func GetEmails(s Server) chi.Router {
	r := chi.NewRouter()
	r.Get("/{term}", handlers.SearchEmail)
	return r
}
