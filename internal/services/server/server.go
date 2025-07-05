package server

import (
	"log"
	"net/http"

	"github.com/sangnt1552314/digimon-godesk/internal/config"
	"github.com/sangnt1552314/digimon-godesk/internal/logger"
)

type Server struct {
	Config *config.Config
}

func NewServer() *http.Server {
	// Load configuration first
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Setup logging with configuration
	if err := logger.SetupLogging(cfg); err != nil {
		panic(err)
	}

	// Create a new server instance with the loaded configuration
	NewServer := &Server{
		Config: cfg,
	}

	// Create a new HTTP server
	log.Printf("Starting server on port %s in %s environment", cfg.Port, cfg.Environment)
	server := &http.Server{
		Addr:    cfg.Port,
		Handler: NewServer.RegisterRoutes(),
	}

	return server
}
