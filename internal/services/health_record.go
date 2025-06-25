package services

import (
	"context"
	"fmt"

	"github.com/creative-snails/phisio-log-backend-go/internal/db"
	"github.com/creative-snails/phisio-log-backend-go/internal/models"
	"github.com/google/uuid"
)


type HealthRecordServiceImpl struct {
	queries *db.Queries
}

func NewHealthRecordService(queries *db.Queries) HealthRecordService {
	return &HealthRecordServiceImpl{
		queries: queries,
	}
}

func (s *HealthRecordServiceImpl) GetHealthRecord(ctx context.Context, healthRecordId string) (db.HealthRecord, error) {
	id, err := uuid.Parse(healthRecordId)
	if err != nil {
		return db.HealthRecord{}, fmt.Errorf("invalid UUID: %w", err)
	}
	
	return s.queries.GetHealthRecord(ctx, id)
}

func (s *HealthRecordServiceImpl) GetSymptoms(ctx context.Context, healthRecordId string) ([]db.Symptom, error) {
	id, err := uuid.Parse(healthRecordId)
	if err != nil {
		return []db.Symptom{}, fmt.Errorf("invalid UUID: %w", err)
	}

	return s.queries.GetSymptoms(ctx, id)
}

func (s*HealthRecordServiceImpl) GetAffectedParts(ctx context.Context, symptomId string)([]db.AffectedPart, error) {
	id, err := uuid.Parse(symptomId)
	if err != nil {
		return []db.AffectedPart{}, fmt.Errorf("invalid UUID: %w", err)
	}

	return s.queries.GetAffectedParts(ctx, id)
}

func (s*HealthRecordServiceImpl) GetMedicalConsultations(ctx context.Context, healthRecordId string) ([]db.MedicalConsultation, error) {
	id, err := uuid.Parse(healthRecordId)
	if err != nil {
		return []db.MedicalConsultation{}, fmt.Errorf("invalid UUID: %w", err)
	}

	return s.queries.GetMedicalConsultations(ctx, id)
}

func (s *HealthRecordServiceImpl) CreateHealthRecord(ctx context.Context, req *models.CreateHealthRecordRequest) (db.HealthRecord, error) {
	if err := req.Validate(); err != nil {
		return db.HealthRecord{}, fmt.Errorf("validtion failed: %w", err)
	}

	params := db.CreateHealthRecordParams{
		// UserID:				req.UserID,
		ParentRecordID: 	req.ParentRecordID,
		Description: 		req.Description,
		Stage: 				db.StageEnum(req.Improvement),
		Severity: 			db.SeverityEnum(req.Severity),
		Progression: 		db.ProgressionEnum(req.Progress),	
		TreatmentsTried: 	req.TreatmentsTried,
	}

	return s.queries.CreateHealthRecord(ctx, params)
}