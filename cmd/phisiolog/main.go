package main

import (
	"github.com/creative-snails/phisio-log-backend-go/internal/startup"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	
	r := chi.NewRouter()

	// Connect to databse
	startup.InitializeDB()

	// Register routes before initializing the server
	startup.Routes(r)

	// Initialize server
	startup.Server(r)


}