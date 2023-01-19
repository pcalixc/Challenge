package server

import "backend/pkg/handlers"

func (s *Server) SetupRoutes() {

	handlers.ZincSearch(s.Mux)
	handlers.Test(s.Mux)

}
