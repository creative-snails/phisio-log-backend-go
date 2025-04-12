package services

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