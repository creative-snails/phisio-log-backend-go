package models

import (
	"time"
)

type Progress string
const (
	Open		Progress = "open"
	Closed		Progress = "closed"
	InProgress 	Progress = "in-progress"
)

type Improvement string
const (
	Improving	Improvement = "improving"
	Stable		Improvement = "stable"
	Worsening	Improvement = "worsening"
	Varying		Improvement = "varying"
)

type Severity string
const (
	Mild		Severity = "mild"
	Moderate	Severity = "moderate"
	Severe		Severity = "severe"
	Variable	Severity = "variable"
)

// DB structs
type UserDB struct {
	ID		string	`sql:"id,omitempty"`
	Name	string	`sql:"name" validate:"required,min=2,max=100"`
}
type SymptomDB struct {
	ID				string		`sql:"id,omitempty"`
	HealthRecordID	string		`sql:"health_record_id" validate:"required,uuid"`
	Name			string		`sql:"name" validate:"required,min=2,max=200"`
	StartDate		time.Time	`sql:"start_date" validate:"omitempty"`
}
type MedicalConsultationDB struct {
	ID				string		`sql:"id,omitempty"`
	HealthRecordID	string		`sql:"health_record_id" validate:"required,uuid"`
	Consultant		string		`sql:"consultant" validate:"required,min=2,max=100"`
	Date			time.Time	`sql:"date" validate:"required"`
	Diagnosis		string		`sql:"diagnosis" validate:"required,min=2,max=1000"`
	FollowUpActions	[]string	`sql:"follow_up_actions" validate:"omitempty,dive,min=2,max=200"`
}
type HealthRecordDB struct {
	ID 				string		`sql:"id,omitempty"`
	UserID			string		`sql:"user_id" validate:"required,uuid"`
	ParentRecordID	string		`sql:"parent_record_id" validate:"omitempty,uuid"`
	Description		string		`sql:"description" validate:"required,min=10,max=2000"`
	Progress		Progress	`sql:"progress" validate:"omitempty,oneof=open closed in-progress"`
	Improvement		Improvement	`sql:"improvement" validate:"omitempty,oneof=improving stable worsening varying"`
	Severity 		Severity	`sql:"severity" validate:"omitempty,oneof=mild moderate severe variable"`
	TreatmentsTried []string 	`sql:"treatments_tried" validate:"omitempty,dive,min=2,max=200"`
	CreatedAt		time.Time	`sql:"created_at,omitempty"`
	UpdatedAt		time.Time	`sql:"updated_at,omitempty"`
}

// API structs
type User struct {
	ID		string	`json:"id"`
	Name	string	`json:"name"`
} 

type Symptom struct {
	ID			string		`json:"id"`
	Name		string		`json:"name"`
	StartDate	time.Time	`json:"startDate"`
}

type MedicalConsultation struct {
	ID				string		`json:"id"`
	Consultant		string		`json:"consultant"`
	Date			time.Time	`json:"date"`
	Diagnosis		string		`json:"diagnosis"`
	FollowUpActions	[]string	`json:"followUpActions"`
}
type HealthRecord struct {
	ID 						string 					`json:"id"`
	User 					User					`json:"user"`
	Description 			string		 			`json:"description"`
	Progress 				Progress				`json:"progress"`
	Improvement				Improvement				`json:"improvement"`
	Severity 				Severity				`json:"severity"`
	Symptoms				[]Symptom				`json:"symptoms"`
	TreatmentsTried 		[]string 				`json:"treatmentsTried"`
	MedicalConsultations	[]MedicalConsultation	`json:"medicalConsultations"`
	Updates					[]HealthRecord			`json:"updates"`
	CreatedAt 				time.Time				`json:"createdAt"`
	UpdatedAt 				time.Time				`json:"updatedAt"`
}