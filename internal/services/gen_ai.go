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

func TextGen(messages []Message) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("OPENAI_API_KEY environment variable is not set")
	}

	for _, msg := range messages {
		if err := msg.Validate(); err != nil {
			return "", fmt.Errorf("message validation failed: %w", err)
		}
	}

	client := resty.New()
	
	requestBody := struct {
		Model		string 		`json:"model"`
		Messages	[]Message	`json:"messages"`
	}{
		Model:		"gpt-3.5-turbo",
		Messages:	messages,
	}

	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+apiKey).
		SetBody(requestBody).
		Post("https://api.openai.com/v1/chat/completions")

	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}

	var openAIResponse OpenAIResponse
	if err := json.Unmarshal(response.Body(), &openAIResponse); err != nil {
		return "", fmt.Errorf("failed to unmarshall response: %w", err)
	}

	return openAIResponse.Choices[0].Message.Content, nil
}