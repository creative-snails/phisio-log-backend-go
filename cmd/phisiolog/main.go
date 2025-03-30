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
	config, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	log.SetReportCaller(true)
	r := chi.NewRouter()

	port := config.Server.Port
	if envPort := os.Getenv("PORT"); envPort != "" {
		if p, err := strconv.Atoi(envPort); err == nil {
			port = p
		}
	}

	fmt.Printf("Server listening at port %d...\n", port)
	
	address := fmt.Sprintf("localhost:%d", port)
	if err:= http.ListenAndServe(address, r); err != nil {
		log.Error(err)
	}
}
