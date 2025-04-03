package startup

import (
	"github.com/creative-snails/phisio-log-backend-go/internal/routes"
	"github.com/go-chi/chi"
)

func Routes(r *chi.Mux) {
	r.Route("/api/health-records", routes.HealthRecords)
}