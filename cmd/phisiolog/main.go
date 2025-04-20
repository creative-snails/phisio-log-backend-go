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
	services.NewHealthRecordService(queries)
	h := handlers.NewHandler(&services.HealthRecordService{})
	startup.Routes(r, h)

	// Initialize server
	startup.Server(r)


}