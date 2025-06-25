package services

import (
	"context"

	"github.com/creative-snails/phisio-log-backend-go/internal/db"
	"github.com/creative-snails/phisio-log-backend-go/internal/models"
)

type HealthRecordService interface {
	GetHealthRecord(ctx context.Context, healthRecordId string) (db.HealthRecord, error)
	GetSymptoms(ctx context.Context, healthRecordId string) ([]db.Symptom, error)
	GetAffectedParts(ctx context.Context, symptomId string) ([]db.AffectedPart, error)
	CreateHealthRecord(ctx context.Context, req *models.CreateHealthRecordRequest) (db.HealthRecord, error)
}