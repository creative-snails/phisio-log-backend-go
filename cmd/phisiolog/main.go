package main

import (
	"github.com/creative-snails/phisio-log-backend-go/internal/startup"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	startup.Server(r)
}