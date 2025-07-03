package main

import (
	"fmt"
	"net/http"

	"github.com/sangnt1552314/digimon-godesk/internal/services/server"
)

func main() {

	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}
