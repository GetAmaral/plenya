package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ACSMRiskLevel define os níveis de risco ACSM
type ACSMRiskLevel string

const (
	ACSMRiskLow      ACSMRiskLevel = "low"
	ACSMRiskModerate ACSMRiskLevel = "moderate"
	ACSMRiskHigh     ACSMRiskLevel = "high"
)

// SmokingStatus define o status de tabagismo
type SmokingStatus string

const (
	SmokingNever   SmokingStatus = "never"
	SmokingFormer  SmokingStatus = "former"
	SmokingCurrent SmokingStatus = "current"
)

// PhysicalActivityLevel define o nível de atividade física
type PhysicalActivityLevel string

const (
	ActivitySedentary    PhysicalActivityLevel = "sedentary"
	ActivityInsufficient PhysicalActivityLevel = "insufficient"
	ActivityModerate     PhysicalActivityLevel = "moderate"
	ActivityActive       PhysicalActivityLevel = "active"
)

// PhysicalAssessment armazena dados brutos de avaliação física + ACSM
// @Description Avaliação física com estratificação de risco ACSM
type PhysicalAssessment struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	PatientID   uuid.UUID `gorm:"type:uuid;not null;index" json:"patientId"`
	CreatedByID uuid.UUID `gorm:"type:uuid;not null" json:"createdById"`

	AssessmentDate time.Time `gorm:"type:date;not null" json:"assessmentDate"`

	// Antropometria
	Weight             *float64 `gorm:"type:numeric(6,2)" json:"weight"`
	Height             *float64 `gorm:"type:numeric(5,1)" json:"height"`
	WaistCircumference *float64 `gorm:"type:numeric(5,1)" json:"waistCircumference"`

	// Composição Corporal (calculados no Create)
	BMI            *float64 `gorm:"type:numeric(5,2)" json:"bmi"`
	BRI            *float64 `gorm:"type:numeric(5,2)" json:"bri"`
	BodyFatPercent *float64 `gorm:"type:numeric(5,2)" json:"bodyFatPercent"`
	LeanMass       *float64 `gorm:"type:numeric(6,2)" json:"leanMass"`

	// Cardiovascular
	SystolicBP       *int `gorm:"type:int" json:"systolicBp"`
	DiastolicBP      *int `gorm:"type:int" json:"diastolicBp"`
	RestingHeartRate *int `gorm:"type:int" json:"restingHeartRate"`

	// Laboratorial (preenchido automaticamente do último LabResult ou input manual)
	LDL              *float64 `gorm:"type:numeric(6,2)" json:"ldl"`
	HDL              *float64 `gorm:"type:numeric(6,2)" json:"hdl"`
	TotalCholesterol *float64 `gorm:"type:numeric(6,2)" json:"totalCholesterol"`
	Triglycerides    *float64 `gorm:"type:numeric(6,2)" json:"triglycerides"`
	FastingGlucose   *float64 `gorm:"type:numeric(6,2)" json:"fastingGlucose"`
	HbA1c            *float64 `gorm:"type:numeric(5,2)" json:"hbA1c"`

	// Histórico (ACSM risk factors)
	FamilyHistory         *bool                  `gorm:"type:bool" json:"familyHistory"`
	SmokingStatus         *SmokingStatus         `gorm:"type:varchar(10);check:smoking_status IN ('never','former','current')" json:"smokingStatus"`
	PhysicalActivityLevel *PhysicalActivityLevel `gorm:"type:varchar(15);check:physical_activity_level IN ('sedentary','insufficient','moderate','active')" json:"physicalActivityLevel"`

	// Condições
	CardiovascularDisease *bool   `gorm:"type:bool" json:"cardiovascularDisease"`
	DiabetesType          *string `gorm:"type:varchar(20)" json:"diabetesType"`
	Symptoms              *string `gorm:"type:text" json:"symptoms"`
	ClinicalAlert         *bool   `gorm:"type:bool" json:"clinicalAlert"`

	// ACSM Estratificação (calculada no Create)
	ACSMRiskLevel        *ACSMRiskLevel `gorm:"type:varchar(10);check:acsm_risk_level IN ('low','moderate','high')" json:"acsmRiskLevel"`
	ACSMRiskFactorsCount int            `gorm:"type:int;not null;default:0" json:"acsmRiskFactorsCount"`
	ACSMRiskFactors      []string       `gorm:"type:jsonb;serializer:json" json:"acsmRiskFactors"`
	ACSMProtectiveFactors []string      `gorm:"type:jsonb;serializer:json" json:"acsmProtectiveFactors"`
	ACSMRecommendation   *string        `gorm:"type:text" json:"acsmRecommendation"`
	ACSMTags             []string       `gorm:"type:jsonb;serializer:json" json:"acsmTags"`

	// Fotos
	FrontPhotoURL *string `gorm:"type:text" json:"frontPhotoUrl"`
	SidePhotoURL  *string `gorm:"type:text" json:"sidePhotoUrl"`

	// IA
	AIRecommendation *string `gorm:"type:text" json:"aiRecommendation"`

	// Título computado
	DisplayTitle string `gorm:"-" json:"displayTitle"`

	// Timestamps
	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relações
	Patient   Patient `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"patient,omitempty"`
	CreatedBy User    `gorm:"foreignKey:CreatedByID;constraint:OnDelete:RESTRICT" json:"createdBy,omitempty"`
}

func (PhysicalAssessment) TableName() string {
	return "physical_assessments"
}

func (pa *PhysicalAssessment) GetTitle() string {
	risk := "Não calculado"
	if pa.ACSMRiskLevel != nil {
		riskLabels := map[ACSMRiskLevel]string{
			ACSMRiskLow:      "Baixo Risco",
			ACSMRiskModerate: "Risco Moderado",
			ACSMRiskHigh:     "Alto Risco",
		}
		if label, ok := riskLabels[*pa.ACSMRiskLevel]; ok {
			risk = label
		}
	}
	return fmt.Sprintf("Avaliação Física - %s - %s", pa.AssessmentDate.Format("02/01/2006"), risk)
}

func (pa *PhysicalAssessment) AfterFind(tx *gorm.DB) error {
	pa.DisplayTitle = pa.GetTitle()
	return nil
}

func (pa *PhysicalAssessment) BeforeCreate(tx *gorm.DB) error {
	if pa.ID == uuid.Nil {
		pa.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
