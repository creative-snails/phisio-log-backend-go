package types

import "time"

type Stage string
const (
	Open		Stage = "open"
	Closed		Stage = "closed"
	InProgress 	Stage = "in-progress"
)

type Severity string
const (
	Mild		Severity = "mild"
	Moderate	Severity = "moderate"
	Severe		Severity = "severe"
	Variable	Severity = "variable"
)

type Progression string
const (
	Improving	Progression = "improving"
	Stable		Progression = "stable"
	Worsening	Progression = "worsening"
	Varying		Progression = "varying"
)
type User struct {
	ID		string	`json:"id"`
	Name	string	`json:"name"`
} 

type Status struct {
	Stage 			Stage			`json:"stage"`
	Severity 		Severity		`json:"severity"`
	Progression		Progression		`json:"progression"`
}

type AffectedPart struct {
	Key		string	`json:"key"`
	State   string	`json:"state"`
}

type Symptom struct {
	ID				string			`json:"id"`
	Name			string			`json:"name"`
	StartDate		time.Time		`json:"startDate"`
	AffectedParts	[]AffectedPart	`json:"affectedParts"`
}

type MedicalConsultation struct {
	ID				string		`json:"id"`
	Consultant		string		`json:"consultant"`
	Date			time.Time	`json:"date"`
	Diagnosis		string		`json:"diagnosis"`
	FollowUpActions	[]string	`json:"followUpActions,omitempty"`
}
type HealthRecord struct {
	ID 						string 					`json:"id"`
	ParentRecordID			string					`json:"parentRecordId,omitempty"`
	Description 			string		 			`json:"description,omitempty"`
	Status					Status					`json:"status"`
	Symptoms				[]Symptom				`json:"symptoms,omitempty"`
	TreatmentsTried 		[]string 				`json:"treatmentsTried,omitempty"`
	MedicalConsultations	[]MedicalConsultation	`json:"medicalConsultations,omitempty"`
	CreatedAt 				time.Time				`json:"createdAt"`
	UpdatedAt 				time.Time				`json:"updatedAt"`
}

type HealthRecordPayload struct {
	ID 						string 					`json:"id"`
	ParentRecordID			string					`json:"parentRecordId,omitempty"`
	Description 			string		 			`json:"description"`
	Status					Status					`json:"status"`
	Symptoms				[]Symptom				`json:"symptoms"`
	TreatmentsTried 		[]string 				`json:"treatmentsTried,omitempty"`
	MedicalConsultations	[]MedicalConsultation	`json:"medicalConsultations,omitempty"`
	Updates					[]HealthRecord			`json:"updates,omitempty"`
	CreatedAt 				time.Time				`json:"createdAt"`
	UpdatedAt 				time.Time				`json:"updatedAt"`
}
