package main

import (
	"fmt"
	"net/http"

	"github.com/creative-snails/phisio-log-backend-go/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	// var r *chi.Mux = chi.NewRouter()
	r := chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API Service...")

	// err:= http.ListenAndServe("localhost:8000", r)
	// if err != nil {
	// 	log.Error(err)
	// }
	
	if err:= http.ListenAndServe("localhost:8000", r); err != nil {
		log.Error(err)
	}
}
