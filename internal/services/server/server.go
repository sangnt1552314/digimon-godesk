package server

import (
	"log"
	"net/http"

	"github.com/sangnt1552314/digimon-godesk/internal/utils"
)

type Server struct {
	Port string
}

func NewServer() *http.Server {
	// Setup logging
	if err := utils.SetupLogging(); err != nil {
		panic(err)
	}

	// Load configuration
	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Create a new server instance with the loaded configuration
	NewServer := &Server{
		Port: config.Port,
	}

	// Create a new HTTP server
	log.Printf("Starting server on port %s", NewServer.Port)
	server := &http.Server{
		Addr:    NewServer.Port,
		Handler: NewServer.RegisterRoutes(),
	}

	return server
}
