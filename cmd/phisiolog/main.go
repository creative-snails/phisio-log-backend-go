package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	r := chi.NewRouter()

	port := 5000
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
