package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/creative-snails/phisio-log-backend-go/internal/prompts"
	"github.com/creative-snails/phisio-log-backend-go/internal/services"
	"github.com/creative-snails/phisio-log-backend-go/internal/types"
	"github.com/robfig/cron/v3"
)

var c = cron.New()

func InitCron() {
	c.AddFunc("0 * * * *", func() {
		now := time.Now()
		maxConversationAge := 24 * time.Hour
		for id, conversation := range Conversations {
			if !conversation.LastAccessed.IsZero() && now.Sub(conversation.LastAccessed) > maxConversationAge {
				delete(Conversations, id)
			}
		}
	})
	c.Start()
}
type Handler struct {
	healthRecordService services.HealthRecordService
}

func NewHandler(healthRecordService services.HealthRecordService) *Handler {
	InitCron()

	return &Handler {
		healthRecordService: healthRecordService,
	}
}

func (h *Handler) GetHealthRecord(w http.ResponseWriter, r *http.Request) {
	healthRecordId := "4b5959c6-c88e-4c12-8115-e971669dbbe4"
	record, err := h.healthRecordService.GetHealthRecord(r.Context(), healthRecordId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	symptoms, err := h.healthRecordService.GetSymptoms(r.Context(), healthRecordId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	mappedSymptoms := make([]types.Symptom, len(symptoms))

	for i, symptom := range(symptoms) {
		affectedParts, err := h.healthRecordService.GetAffectedParts(r.Context(), symptom.ID.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		mappedAffectedParts := make([]types.AffectedPart, len(affectedParts))

		for j, affectedPart := range(affectedParts) {
			mappedAffectedParts[j] = types.AffectedPart{
				Key: string(affectedPart.BodyPart),
				State: strconv.Itoa(int(affectedPart.State)),
			}
		}

		mappedSymptoms[i] = types.Symptom{
			ID: symptom.ID.String(),
			Name: symptom.Name,
			StartDate: symptom.StartDate.Time.String(),
			AffectedParts: mappedAffectedParts,
		}
	}

	medicalConsultations, err := h.healthRecordService.GetMedicalConsultations(r.Context(), healthRecordId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	mappedMedicalConsultations := make([]types.MedicalConsultation, len(medicalConsultations))

	for i, medicalConsulation := range(medicalConsultations) {
		mappedMedicalConsultations[i] = types.MedicalConsultation{
			ID: medicalConsulation.ID.String(),
			Consultant: medicalConsulation.Consultant,
			Date: medicalConsulation.Date.String(),
			Diagnosis: medicalConsulation.Diagnosis,
			FollowUpActions: medicalConsulation.FollowUpActions,
		}
	}

	healthRecordPayload := types.HealthRecordPayload{
		ID: record.ID.String(),
		Description: record.Description,
		TreatmentsTried: record.TreatmentsTried,
		Status: types.Status{
			Stage: types.Stage(record.Stage),
			Severity: types.Severity(record.Severity),
			Progression: types.Progression(record.Progression),
		},
		Symptoms: mappedSymptoms,
		MedicalConsultations: mappedMedicalConsultations,
		CreatedAt: record.CreatedAt.Time.String(),
		UpdatedAt: record.UpdatedAt.Time.String(),
	}


	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(healthRecordPayload); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
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