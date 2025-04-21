package handlers

import (
	"encoding/json"
	"net/http"

	m "github.com/creative-snails/phisio-log-backend-go/internal/models"
	s "github.com/creative-snails/phisio-log-backend-go/internal/services"
	t "github.com/creative-snails/phisio-log-backend-go/internal/types"
	"github.com/google/uuid"
)

type Handler struct {
	healthRecordService s.HealthRecordService
}

func NewHandler(healthRecordService *s.HealthRecordService) *Handler {
	return &Handler {
		healthRecordService: *healthRecordService,
	}
}

func (h *Handler) CreateHealthRecord(w http.ResponseWriter, r *http.Request) {
	var rawReq struct {
		UserID			string	`json:"userId"`
		ParentRecordID	string	`json:"parentRecordId,omitempty"`
		Description		string	`json:"description"`
		Progress		string	`json:"progress"`
		Improvement		string	`json:"improvement"`
		Severity		string	`json:"severity"`
		TreatmentsTried	[]string	`json:"treatmentsTried"`
	}

	if err := json.NewDecoder(r.Body).Decode(&rawReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	treatments := []string{}
	if rawReq.TreatmentsTried != nil {
		treatments = rawReq.TreatmentsTried
	}

	req := m.CreateHealthRecordRequest {
		Description: rawReq.Description,
		Progress: t.Progress(rawReq.Progress),
		Improvement: t.Improvement(rawReq.Improvement),
		Severity: t.Severity(rawReq.Severity),
		TreatmentsTried: treatments,
	}

	userID, err := uuid.Parse(rawReq.UserID)
	if err != nil {
		http.Error(w, "Invalid userId format", http.StatusBadRequest)
		return
	}
	req.UserID = userID

	if rawReq.ParentRecordID != "" {
		parentID, err := uuid.Parse(rawReq.ParentRecordID)
		if err != nil {
			http.Error(w, "Invalid parentId format", http.StatusBadRequest)
			return
		}

		req.ParentRecordID = uuid.NullUUID{
			UUID: parentID,
			Valid: true,
		}
	} else {
		req.ParentRecordID = uuid.NullUUID{
			Valid: false,
		}
	}

	record, err := h.healthRecordService.CreateHealthRecord(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(record); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}