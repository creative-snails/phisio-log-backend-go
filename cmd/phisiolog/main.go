package main

import (
	"github.com/creative-snails/phisio-log-backend-go/config"
	"github.com/creative-snails/phisio-log-backend-go/internal/handlers"
	"github.com/creative-snails/phisio-log-backend-go/internal/services"
	"github.com/creative-snails/phisio-log-backend-go/internal/startup"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Load configuration
	config, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	
	log.SetReportCaller(true)
	
	// Connect to database
	queries, err := startup.InitializeDB(config.Database)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	r := chi.NewRouter()
	// Register routes before initializing the server
	healthRecordService := services.NewHealthRecordService(queries)
	if healthRecordService == nil {
		log.Fatalf("Failed to create health record service")
	}

	h := handlers.NewHandler(healthRecordService)
	if h == nil {
		log.Fatal("Failed to create handler")
	}
	startup.Routes(r, h)

	// Initialize server
	startup.Server(r, config.Server)
}