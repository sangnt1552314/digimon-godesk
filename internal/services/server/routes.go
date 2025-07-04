package server

import (
	"net/http"

	"github.com/sangnt1552314/digimon-godesk/internal/handler/api"
	"github.com/sangnt1552314/digimon-godesk/internal/handler/web"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	// Register static file handler
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// API routes
	mux.HandleFunc("GET /health", api.HealthCheckHandler)

	// Web routes
	mux.HandleFunc("/", web.IndexHandler)
	mux.HandleFunc("POST /digimon", web.AddDigimonHandler)

	return mux
}
