package routes

import (
	"net/http"

	"github.com/creative-snails/phisio-log-backend-go/internal/handlers"
	"github.com/go-chi/chi"
)

func HealthRecords(r chi.Router, handler *handlers.Handler) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Getting health records!"))
	})

	r.Post("/new-record", handler.CreateHealthRecord)
}