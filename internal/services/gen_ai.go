package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
)

type OpenAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Error *OpenAIError `json:"error,omitempty"`
}

type OpenAIError struct {
	Message	string	`json:"message"`
	Type 	string	`json:"type"`
	Code 	string	`json:"code"`
}

type Role string
const (
	System		Role = "system"
	Assistant	Role = "assistant"
	User		Role = "user"
)

type Message struct {
	Role	Role	`json:"role"`
	Content	string	`json:"content"`
}

func (m *Message) Validate() error {
	switch m.Role {
	case System, Assistant, User:
		return nil
	default:
		return errors.New("invalide role")
	}
}

func GenAI(messages []Message, genType string) (string, error) {
	if genType != "text" && genType != "json" {
		return "", fmt.Errorf("invalid getnType: must be 'text' or 'json'")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("OPENAI_API_KEY environment variable is not set")
	}

	for _, msg := range messages {
		if err := msg.Validate(); err != nil {
			return "", fmt.Errorf("message validation failed: %w", err)
		}
	}

	type responseFormat struct {
		Type string `json:"type"`
	}

	type requestBody struct {
		Model			string			`json:"model"`
		Messages		[]Message		`json:"messages"`
		ResponseFormat	*responseFormat	`json:"response_format,omitempty"`
	}
	
	
	body := requestBody {
		Model:		"gpt-3.5-turbo",
		Messages:	messages,
	}

	if genType == "json" {
		body.ResponseFormat = &responseFormat{
			Type: "json_object",
		}
	}

	client := resty.New()

	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+apiKey).
		SetBody(body).
		Post("https://api.openai.com/v1/chat/completions")

	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}

	var openAIResponse OpenAIResponse
	if err := json.Unmarshal(response.Body(), &openAIResponse); err != nil {
		return "", fmt.Errorf("failed to unmarshall response: %w", err)
	}

	if openAIResponse.Error != nil {
		return "", fmt.Errorf("API error: %s (type: %s, code: %s)", 
		openAIResponse.Error.Message,
		openAIResponse.Error.Type,
		openAIResponse.Error.Code)
	}

	if len(openAIResponse.Choices) == 0 {
		fmt.Print(openAIResponse)
		return "", fmt.Errorf("there are no choices in the response body")
	}

	return openAIResponse.Choices[0].Message.Content, nil
}