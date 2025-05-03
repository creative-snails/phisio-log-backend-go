package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/creative-snails/phisio-log-backend-go/internal/handlers"
	"github.com/go-chi/chi"
	"github.com/go-resty/resty/v2"
)


type OpenAIResponse struct {
 	Choices []struct {
 		Message struct {
 			Content string `json:"content"`
 		} `json:"message"`
 	} `json:"choices"`
}



func HealthRecords(r chi.Router, handler *handlers.Handler) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		apiKey := os.Getenv("OPENAI_API_KEY")

		if apiKey == "" {
			http.Error(w, "OPENAI_API_KEY environment variable is no set", http.StatusInternalServerError)
			return
		}

		client := resty.New()

		response, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetHeader("Authorization", "Bearer "+apiKey).
			SetBody(`{
				"model": "gpt-3.5-turbo",
				"messages": [{"role": "user", "content": "Say this is a test"}]
			}`).
			Post("https://api.openai.com/v1/chat/completions")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var openAIResponse OpenAIResponse
		if err := json.Unmarshal(response.Body(), &openAIResponse); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Response: %s", openAIResponse.Choices[0].Message.Content)

		w.Write([]byte("Getting health records!"))
	})

	r.Post("/new-record", handler.CreateHealthRecord)
}