package prompts

import (
	"fmt"
	"time"

	"github.com/creative-snails/phisio-log-backend-go/internal/models"
)

type Prompts struct {
	System		SystemPrompts
	Assistant	AssistantPrompts
}

type SystemPrompts struct {
	Init		string
	Treatments	func(currentRecord models.CreateHealthRecordRequest) string
	Validation	string
}

type AssistantPrompts struct {
	Treatments string
}

func NewPrompts() *Prompts {
	p := &Prompts{}

	p.System.Init = fmt.Sprintf(`
		Based on the user description, generate a JSON object that accurately matches the Go struct definition and validation rules.

		- For "description": summarize, clean up, and correct mistakes before adding it to the JSON. Do not include any placeholder text, meaningless phrases, or unrelated information.
		- For fields "progress", "improvement", and "severity", interpret the description and select a value from their respective accepted options. These fields are optional and have default values, so if the information is not clearly present, use the default values.
		- For "treatmentsTried": extract all clearly mentioned treatments. Do not infer or assume treatments. If none are mentioned, leave the array empty.
		- Extract only clear, medically relevant details from the user's input. Disregard any vague, unrelated, non-medical, or ambiguous information.
		- If any required fields are missing and have no default, leave those fields empty.
		- Be aware that today's date is %s. Ensure no future dates are assigned to any date field.
		- When multiple "user" messages exist in the conversation history, integrate the medically relevant information from all messages according to the extraction and formatting guidelines to generate a unified JSON object.
		- Summarize and clean up all extracted data, correcting typos and inconsistent phrasing before adding it to the final JSON.

		Go Struct Definition:
		type CreateHealthRecordRequest struct {
			parentRecordId	 string	 // optional, UUID
			description		 string	 // required, min 10, max 2000 characters
			progress		 string	 // optional, one of: "open", "closed", "in-progress" (default: "open")
			improvement		 string  // optional, one of: "improving", "stable", "worsening", "varying" (default: "stable")
			severity		 string  // optional, one of: "mild", "moderate", "severe", "variable" (default: "variable")
			treatmentsTried	 []string // optional, each string min 2, max 200 characters
		}

		Expected JSON Output structure:
		{
			"parentRecordId": "",
			"description": "",
			"progress": "",
			"improvement": "",
			"severity": "",
			"treatmentsTried": []
		}
	`, time.Now().Format("2006-01-02"))

	p.System.Validation =  `
      Generate a user-friendly message using the error messages resulting from the validation of the previous input. Start with the following prompt and ensure the message is clear and helpful for the user:
      'Please provide the following information to complete the health record:'
      Use the validation errors to guide the user on what specific information is missing or incorrect. Ensure the message is polite, clear, and supportive.
    `
	p.Assistant.Treatments = "Have you tried any treatments on your own to manage your condition? If yes, please share the details."

	return p
}