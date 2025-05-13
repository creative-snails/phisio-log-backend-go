package services

import (
	"fmt"

	"github.com/creative-snails/phisio-log-backend-go/internal/models"
)

type ValidateHealthRecordReturn struct {
	Success			bool	`json:"success"`
	SystemPrompt	string	`json:"systemPrompt,omitempty"`
	AssistantPrompt	string	`json:"assistantPrompt,omitempty"`
}

func ValidateHealthRecord(healthRecord *models.CreateHealthRecordRequest) (ValidateHealthRecordReturn, error) {
	result := ValidateHealthRecordReturn{}

	if err := healthRecord.Validate(); err != nil {
		message, genErr := GenAI([]Message{{Role: "user", Content: err.Error()}}, "text")
		if genErr != nil {
			return result, fmt.Errorf("validtion failed: %w", genErr)
		}
		result.AssistantPrompt = message
		result.Success = false
		return result, nil
	}

	return result, nil
}