package db

import (
	"context"
	"database/sql"
	"time"

	log "github.com/sirupsen/logrus"
)

type HealthRecordService struct {
	DB *sql.DB
}

type HealthRecord struct {
    ID    int
    Description  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewHealthRecordService(db *sql.DB) *HealthRecordService {
	return &HealthRecordService{DB: db}
}


func (s *HealthRecordService) Migrate()  {
	query := `
		CREATE TABLE IF NOT EXISTS healthRecords (
			id SERIAL PRIMARY KEY,
			description VARCHAR(500) NOT NULL,
			createdAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			updatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		);
	`

	_, err := s.DB.Exec(query)
	if err != nil {
		log.Fatalf("error createing healthRecords table: %v", err)
	}

	log.Info("Database migration completed successfully")
}

func (s *HealthRecordService) Create( ctx context.Context, healthRecord *HealthRecord)  {
	query := `
		INSERT INTO healthRecords (description, createdAt, updatedAt)
		VALUES ($1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING id
	`
	err := s.DB.QueryRowContext(ctx, query, healthRecord.Description).
		Scan(&healthRecord.ID)
	if err != nil {
		log.Fatalf("error creating healthRecord: %v", err)
	}
}