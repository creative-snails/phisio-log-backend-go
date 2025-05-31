package routes

import (
	"encoding/json"
	"net/http"

	"github.com/creative-snails/phisio-log-backend-go/internal/handlers"
	"github.com/creative-snails/phisio-log-backend-go/internal/models"
	"github.com/creative-snails/phisio-log-backend-go/internal/prompts"
	"github.com/creative-snails/phisio-log-backend-go/internal/services"
	"github.com/go-chi/chi"
)

func HealthRecords(r chi.Router, handler *handlers.Handler) {
	r.Get("/queries/", func(w http.ResponseWriter, r *http.Request) {
		handler.GetHealthRecord(w, r)

	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		messages := []services.Message{
			{
				Role:    "system",
				Content: prompts.NewPrompts().System.Init,
			},
			{
				Role:    "user",
				Content: "I went for a run and strained my ankle, I tried putting ice on it",
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
		if validationResult.AssistantPrompt != "" {
			message, err := json.Marshal(struct{Message string}{Message: validationResult.AssistantPrompt})
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{
					"error": err.Error(),
				})
				return
			}
			w.Write(message)
			return
		}
		
		

		w.Write([]byte(message))
	})

	r.Post("/new-record", handler.CreateHealthRecord)
}