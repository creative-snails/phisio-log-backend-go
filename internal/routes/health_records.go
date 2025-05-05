package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/creative-snails/phisio-log-backend-go/internal/handlers"
	"github.com/creative-snails/phisio-log-backend-go/internal/services"
	"github.com/go-chi/chi"
)

func HealthRecords(r chi.Router, handler *handlers.Handler) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		messages := []services.Message{
    {
        Role:    "system",
        Content: "You are a helpful assistant.",
    },
    {
        Role:    "user",
        Content: "Hello!",
    },
}
		message, err := services.GenAI(messages, "text")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"error": err.Error(),
			})
			return
		}

		fmt.Printf("Message: %s",message)

		w.Write([]byte(message))
	})

	r.Post("/new-record", handler.CreateHealthRecord)
}