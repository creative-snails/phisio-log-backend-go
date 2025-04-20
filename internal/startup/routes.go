package startup

import (
	h "github.com/creative-snails/phisio-log-backend-go/internal/handlers"
	"github.com/creative-snails/phisio-log-backend-go/internal/routes"
	"github.com/go-chi/chi"
)

func Routes(r *chi.Mux, handler *h.Handler) {
	r.Route("/api/health-records", func(r chi.Router) {
		routes.HealthRecords(r, handler)
	})
}