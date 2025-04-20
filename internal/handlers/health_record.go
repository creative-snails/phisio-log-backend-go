package handlers

import (
	"encoding/json"
	"net/http"

	m "github.com/creative-snails/phisio-log-backend-go/internal/models"
	s "github.com/creative-snails/phisio-log-backend-go/internal/services"
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
	var req m.CreateHealthRecordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	record, err := h.healthRecordService.CreateHealthRecord(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(record); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	
}