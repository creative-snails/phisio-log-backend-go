package models

import (
	"reflect"

	t "github.com/creative-snails/phisio-log-backend-go/internal/types"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type CreateHealthRecordRequest struct {
	UserID			uuid.UUID		`json:"userId" validate:"required,uuid"`
	ParentRecordID	uuid.NullUUID	`json:"parentRecordId" validate:"omitempty,uuid"`
	Description		string			`json:"description" validate:"required,min=10,max=2000"`
	Progress		t.Progress		`json:"progress" validate:"omitempty,oneof=open closed in-progress"`
	Improvement		t.Improvement	`json:"improvement" validate:"omitempty,oneof=improving stable worsening varying"`
	Severity		t.Severity		`json:"severity" validate:"omitempty,oneof=mild moderate severe variable"`
	TreatmentsTried []string		`json:"treatmentsTried" validate:"omitempty,dive,min=2,max=200"`
}

func (r *CreateHealthRecordRequest) Validate() error {
	validate := validator.New()

	validate.RegisterCustomTypeFunc(func(field reflect.Value) interface{} {
		if field.Type() == reflect.TypeOf(uuid.UUID{}) {
			return field.Interface().(uuid.UUID).String()
		}
		return nil
	}, uuid.UUID{})

	return validate.Struct(r)
}

