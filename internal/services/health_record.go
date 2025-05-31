package services

import (
	"context"
	"fmt"

	"github.com/creative-snails/phisio-log-backend-go/internal/db"
	"github.com/creative-snails/phisio-log-backend-go/internal/models"
)


type HealthRecordServiceImpl struct {
	queries *db.Queries
}

func NewHealthRecordService(queries *db.Queries) HealthRecordService {
	return &HealthRecordServiceImpl{
		queries: queries,
	}
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