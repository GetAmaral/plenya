package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PosturalAssessment stores goniometer measurements + computed scores
// @Description Avaliacao postural com medicoes e scores
type PosturalAssessment struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	PatientID   uuid.UUID `gorm:"type:uuid;not null;index" json:"patientId"`
	CreatedByID uuid.UUID `gorm:"type:uuid;not null" json:"createdById"`
	AssessmentDate time.Time `gorm:"type:date;not null" json:"assessmentDate"`
	PhysicalAssessmentID *uuid.UUID `gorm:"type:uuid;index" json:"physicalAssessmentId"`

	ViewType string `gorm:"type:varchar(20);not null;default:'front'" json:"viewType"` // front|side_left|side_right|back

	// Raw measurements (degrees)
	ShoulderDeviation    *float64 `gorm:"type:numeric(5,2)" json:"shoulderDeviation"`
	HipDeviation         *float64 `gorm:"type:numeric(5,2)" json:"hipDeviation"`
	HeadLateralDeviation *float64 `gorm:"type:numeric(5,2)" json:"headLateralDeviation"`
	FHP                  *float64 `gorm:"type:numeric(5,2)" json:"fhp"`              // Forward Head Posture
	ThoracicKyphosis     *float64 `gorm:"type:numeric(5,2)" json:"thoracicKyphosis"`
	LumbarLordosis       *float64 `gorm:"type:numeric(5,2)" json:"lumbarLordosis"`
	KneeAngle            *float64 `gorm:"type:numeric(5,2)" json:"kneeAngle"`
	PhotoURL             *string  `gorm:"type:text" json:"photoUrl"`

	// Computed
	PosturalScore          int    `gorm:"type:int;not null;default:0" json:"posturalScore"`
	PosturalClassification string `gorm:"type:varchar(30);not null;default:''" json:"posturalClassification"`
	SevereDeviations       int    `gorm:"type:int;not null;default:0" json:"severeDeviations"`

	Notes *string `gorm:"type:text" json:"notes"`

	// Timestamps
	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	Patient   Patient `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"patient,omitempty"`
	CreatedBy User    `gorm:"foreignKey:CreatedByID;constraint:OnDelete:RESTRICT" json:"createdBy,omitempty"`
}

func (PosturalAssessment) TableName() string {
	return "postural_assessments"
}

func (pa *PosturalAssessment) GetTitle() string {
	return fmt.Sprintf("Avaliacao Postural - %s - %s", pa.ViewType, pa.AssessmentDate.Format("02/01/2006"))
}

func (pa *PosturalAssessment) BeforeCreate(tx *gorm.DB) error {
	if pa.ID == uuid.Nil {
		pa.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
