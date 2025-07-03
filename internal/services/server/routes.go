package server

import (
	"net/http"

	"github.com/sangnt1552314/digimon-godesk/internal/services/handlers/api"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	// Register static file handler
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// API routes
	mux.HandleFunc("/health", api.HealthCheckHandler)

	return mux
}
