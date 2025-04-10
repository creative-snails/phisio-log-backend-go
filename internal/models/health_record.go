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

type UserDB struct {
	ID		string	`sql:"id"`
	Name	string	`sql:"name"`
} 
type User struct {
	ID		string	`json:"id"`
	Name	string	`json:"name"`
} 

type SymptomDB struct {
	ID				string		`sql:"id"`
	HealthRecordID	string		`sql:"health_record_id"`
	Name			string		`sql:"name"`
	StartDate		time.Time	`sql:"start_date"`
}
type Symptom struct {
	ID			string		`json:"id"`
	Name		string		`json:"name"`
	StartDate	time.Time	`json:"startDate"`
}

type MedicalConsultationDB struct {
	ID				string		`sql:"id"`
	HealthRecordID	string		`sql:"health_record_id"`
	Consultant		string		`sql:"consultant"`
	Date			time.Time	`sql:"date"`
	Diagnosis		string		`sql:"diagnosis"`
	FollowUpActions	[]string	`sql:"follow_up_actions"`
}
type MedicalConsultation struct {
	ID				string		`json:"id"`
	Consultant		string		`json:"consultant"`
	Date			time.Time	`json:"date"`
	Diagnosis		string		`json:"diagnosis"`
	FollowUpActions	[]string	`json:"followUpActions"`
}

type HealthRecordDB struct {
	ID 				string		`sql:"id"`
	UserID			string		`sql:"user_id"`
	ParentRecordID	string		`sql:"parent_record_id"`
	Description		string		`sql:"description"`
	Progress		Progress	`sql:"progress"`	
	Improvement		Improvement	`sql:"improvement"`
	Severity 		Severity	`sql:"severity"`
	TreatmentsTried []string 	`sql:"treatments_tried"`
	CreatedAt		time.Time	`sql:"created_at"`
	UpdatedAt		time.Time	`sql:"updated_at"`
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