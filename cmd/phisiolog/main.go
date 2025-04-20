package main

import (
	"github.com/creative-snails/phisio-log-backend-go/internal/handlers"
	"github.com/creative-snails/phisio-log-backend-go/internal/services"
	"github.com/creative-snails/phisio-log-backend-go/internal/startup"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	
	r := chi.NewRouter()
	
	// Connect to databse
	queries, err := startup.InitializeDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	
	// Register routes before initializing the server
	healthRecordService := services.NewHealthRecordService(queries)
	if healthRecordService == nil {
		log.Fatalf("Failed to create heald record service")
	}

	h := handlers.NewHandler(healthRecordService)
	if h == nil {
		log.Fatal("Failed to create handler")
	}
	startup.Routes(r, h)

	// Initialize server
	startup.Server(r)


}