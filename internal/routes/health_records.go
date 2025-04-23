package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

func HealthRecords(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Getting health records!"))
	})
}