package startup

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/creative-snails/phisio-log-backend-go/config"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)


func Server(r *chi.Mux, sc config.ServerConfig) {
	// Get port from evnironment or config
	port := sc.Port
	if envPort := os.Getenv("PORT"); envPort != "" {
		if p, err := strconv.Atoi(envPort); err == nil {
			port = p
		} else {
			log.Warnf("Invalid PORT environment variable: %s", envPort)
		}
	}
	
	host := sc.Host
	address := fmt.Sprintf("%s:%d", host, port)

	log.Infof("Server starting on %s...", address)
	if err := http.ListenAndServe(address, r); err != nil {
		log.Fatal(err)
	}
}
