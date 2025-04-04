package main

import (
	"github.com/creative-snails/phisio-log-backend-go/internal/startup"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	// Register routes before initializing the server
	startup.Routes(r)

	// Initialize server
	startup.Server(r)
}