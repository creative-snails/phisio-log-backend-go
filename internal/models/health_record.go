package models

import (
	"encoding/json"
	"reflect"

	"github.com/creative-snails/phisio-log-backend-go/internal/types"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type CreateHealthRecordRequest struct {
	// UserID			uuid.UUID			`json:"userId" validate:"required,uuid"`
	ParentRecordID	uuid.NullUUID		`json:"parentRecordId" validate:"omitempty,uuid"`
	Description		string				`json:"description" validate:"required,min=10,max=2000"`
	Progress		types.Stage		`json:"progress" validate:"omitempty,oneof=open closed in-progress"`
	Improvement		types.Progression	`json:"improvement" validate:"omitempty,oneof=improving stable worsening varying"`
	Severity		types.Severity		`json:"severity" validate:"omitempty,oneof=mild moderate severe variable"`
	TreatmentsTried []string			`json:"treatmentsTried" validate:"omitempty,dive,min=2,max=200"`
}

// UnmarshalJSON implements custon JSON unmarshalling to handle empty string for parentRecordId
func (r *CreateHealthRecordRequest) UnmarshalJSON(data []byte) error {
	// Create an alias to avoid infinite recursion when calling json.Unmarshal
	type Alias CreateHealthRecordRequest

	// Create tmeporary stuct with the same fields but using a string for parentRecordId
	temp := struct {
		ParentRecordID	string	`json:"parentRecordId"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	// Only set ParentRecordID if it's not an empty string
	if temp.ParentRecordID != "" {
		id, err := uuid.Parse(temp.ParentRecordID)
		if err != nil {
			return err
		}
		r.ParentRecordID = uuid.NullUUID{UUID: id, Valid: true}
	} else {
		r.ParentRecordID = uuid.NullUUID{Valid: false}
	}

	return nil
}

func (r *CreateHealthRecordRequest) Validate() error {
	validate := validator.New()

	validate.RegisterCustomTypeFunc(func(field reflect.Value) any {
		if field.Type() == reflect.TypeOf(uuid.UUID{}) {
			return field.Interface().(uuid.UUID).String()
		}
		return nil
	}, uuid.UUID{})

	return validate.Struct(r)
}

