package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	r := chi.NewRouter()
	port := 5000

	fmt.Printf("Server listening at port %v ...\n", port)
	
	address := fmt.Sprintf("localhost:%d", port)
	if err:= http.ListenAndServe(address, r); err != nil {
		log.Error(err)
	}
}
