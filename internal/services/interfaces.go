package services

import (
	"context"

	"github.com/creative-snails/phisio-log-backend-go/internal/db"
	"github.com/creative-snails/phisio-log-backend-go/internal/models"
)

type HealthRecordService interface {
	CreateHealthRecord(ctx context.Context, req *models.CreateHealthRecordRequest) (db.HealthRecord, error)
}