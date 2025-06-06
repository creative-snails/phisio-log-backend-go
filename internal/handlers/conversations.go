package handlers

import (
	"time"

	"github.com/creative-snails/phisio-log-backend-go/internal/services"
	"github.com/google/uuid"
)

type Conversation struct {
	ID				uuid.UUID
	History			[]services.Message
	LastAccessed	time.Time
}

func NewConversation (systemPrompt string) *Conversation {
	conversation := &Conversation{}

	conversation.ID = uuid.New()
	conversation.History = []services.Message{{Role: "system", Content: systemPrompt}}
	conversation.LastAccessed = time.Now()

	Conversations[conversation.ID] = conversation
	return conversation
}

var Conversations = make(map[uuid.UUID]*Conversation)

func GetOrCreateConvesation(conversationID string, systemPrompt string) *Conversation {
	if conversationID == "" {
		return NewConversation(systemPrompt)
	}

	convID, err := uuid.Parse(conversationID) 
	if err != nil {
		return NewConversation(systemPrompt)
	}

	conv, exists := Conversations[convID]
	if !exists {
		return NewConversation(systemPrompt)
	}

	return conv
}