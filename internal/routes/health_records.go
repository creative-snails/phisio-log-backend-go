package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/creative-snails/phisio-log-backend-go/internal/handlers"
	"github.com/creative-snails/phisio-log-backend-go/internal/models"
	"github.com/creative-snails/phisio-log-backend-go/internal/prompts"
	"github.com/creative-snails/phisio-log-backend-go/internal/services"
	"github.com/go-chi/chi"
)

func HealthRecords(r chi.Router, handler *handlers.Handler) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		messages := []services.Message{
			{
				Role:    "system",
				Content: prompts.NewPrompts().System.Init,
			},
			{
				Role:    "user",
				Content: "",
			},
		}
		message, err := services.GenAI(messages, "json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"error": err.Error(),
			})
			return
		}
	
		healthRecord := models.CreateHealthRecordRequest{}
		if err = healthRecord.UnmarshalJSON([]byte(message)); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"error": err.Error(),
			})
			return
		}

		validationResult, err := services.ValidateHealthRecord(&healthRecord)
		fmt.Println("Validation result: ", validationResult.AssistantPrompt)
		

		w.Write([]byte(message))
	})

	r.Post("/new-record", handler.CreateHealthRecord)
}