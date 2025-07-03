package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sangnt1552314/digimon-godesk/internal/utils"
)

func main() {
	// Setup logging
	if err := utils.SetupLogging(); err != nil {
		panic(err)
	}

	// Load configuration
	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	router := http.NewServeMux()

	//Handles static files
	router.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	})

	server := &http.Server{
		Addr:    config.Port,
		Handler: router,
	}

	log.Printf("Starting server on port %s", config.Port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
