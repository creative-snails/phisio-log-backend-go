package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/creative-snails/phisio-log-backend-go/internal/prompts"
	"github.com/creative-snails/phisio-log-backend-go/internal/services"
)

type Handler struct {
	healthRecordService services.HealthRecordService
}

func NewHandler(healthRecordService services.HealthRecordService) *Handler {
	return &Handler {
		healthRecordService: healthRecordService,
	}
}

func (h *Handler) CreateHealthRecord(w http.ResponseWriter, r *http.Request) {

	var rawReq struct {
		Message 		string	`json:"message"`
		ConversationID	string	`json:"conversationId,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&rawReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	conversation := GetOrCreateConvesation(rawReq.ConversationID, prompts.NewPrompts().System.Init)
	conversation.History = append(conversation.History, services.Message{Role: "user", Content: rawReq.Message})

	fmt.Println(*Conversations[conversation.ID])


	// treatments := []string{}
	// if rawReq.TreatmentsTried != nil {
	// 	treatments = rawReq.TreatmentsTried
	// }

	// req := models.CreateHealthRecordRequest {
	// 	Description: rawReq.Description,
	// 	Progress: types.Progress(rawReq.Progress),
	// 	Improvement: types.Improvement(rawReq.Improvement),
	// 	Severity: types.Severity(rawReq.Severity),
	// 	TreatmentsTried: treatments,
	// }

	// userID, err := uuid.Parse(rawReq.UserID)
	// if err != nil {
	// 	http.Error(w, "Invalid userId format", http.StatusBadRequest)
	// 	return
	// }
	// req.UserID = userID

	// if rawReq.ParentRecordID != "" {
	// 	parentID, err := uuid.Parse(rawReq.ParentRecordID)
	// 	if err != nil {
	// 		http.Error(w, "Invalid parentId format", http.StatusBadRequest)
	// 		return
	// 	}

	// 	req.ParentRecordID = uuid.NullUUID{
	// 		UUID: parentID,
	// 		Valid: true,
	// 	}
	// } else {
	// 	req.ParentRecordID = uuid.NullUUID{
	// 		Valid: false,
	// 	}
	// }

	// record, err := h.healthRecordService.CreateHealthRecord(r.Context(), &req)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// if err := json.NewEncoder(w).Encode(record); err != nil {
	// 	http.Error(w, "Error encoding response", http.StatusInternalServerError)
	// 	return
	// }
}