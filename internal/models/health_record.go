package models

import (
	t "github.com/creative-snails/phisio-log-backend-go/internal/types"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type CreateHealthRecordRequest struct {
	UserID			uuid.UUID		`validate:"required,uuid"`
	ParentRecordID	uuid.NullUUID	`validate:"omitempty,uuid"`
	Description		string			`validate:"required,min=10,max=2000"`
	Progress		t.Progress		`validate:"omitempty,oneof=open closed in-progress"`
	Improvement		t.Improvement	`validate:"omitempty,oneof=improving stable worsening varying"`
	Severity		t.Severity		`validate:"omitempty,oneof=mild moderate severe variable"`
	TreatmentsTried []string		`validate:"omitempty,dive,min=2,max=200"`
}

func (r *CreateHealthRecordRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

