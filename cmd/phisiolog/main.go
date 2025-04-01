package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/creative-snails/phisio-log-backend-go/config"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Load configuration
	config, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	log.SetReportCaller(true)
	r := chi.NewRouter()

	// Get port from evnironment or config
	port := config.Server.Port
	if envPort := os.Getenv("PORT"); envPort != "" {
		if p, err := strconv.Atoi(envPort); err == nil {
			port = p
		} else {
			log.Warnf("Invalid PORT environment variable: %s", envPort)
		}
	}
	
	host := config.Server.Host
	address := fmt.Sprintf("%s:%d", host, port)

	log.Infof("Server starting on %s...", address)
	if err:= http.ListenAndServe(address, r); err != nil {
		log.Fatal(err)
	}
}
