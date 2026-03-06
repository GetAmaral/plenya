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

// PhysicalAssessment representa uma avaliação física
// NÃO armazena dados que já existem em AnamnesisItems ou LabResults (ZERO duplicação).
// Referencia uma Anamnesis existente e calcula ACSM a partir dos dados referenciados.
// @Description Avaliação física com estratificação de risco ACSM
type PhysicalAssessment struct {
	// ID único da avaliação
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// ID do paciente
	PatientID uuid.UUID `gorm:"type:uuid;not null;index" json:"patientId"`

	// ID do profissional que criou
	CreatedByID uuid.UUID `gorm:"type:uuid;not null" json:"createdById"`

	// Data da avaliação
	AssessmentDate time.Time `gorm:"type:date;not null" json:"assessmentDate"`

	// Referência à Anamnesis que contém os dados físicos (AnamnesisItems -> ScoreItems)
	AnamnesisID uuid.UUID `gorm:"type:uuid;not null;index" json:"anamnesisId"`

	// ACSM Estratificação (CALCULADO a partir de LabResults + AnamnesisItems)
	ACSMRiskLevel         *ACSMRiskLevel `gorm:"type:varchar(10);check:acsm_risk_level IN ('low','moderate','high')" json:"acsmRiskLevel"`
	ACSMRiskFactorsCount  int            `gorm:"type:int;not null;default:0" json:"acsmRiskFactorsCount"`
	ACSMRiskFactors       []string       `gorm:"type:jsonb;serializer:json" json:"acsmRiskFactors"`
	ACSMProtectiveFactors []string       `gorm:"type:jsonb;serializer:json" json:"acsmProtectiveFactors"`
	ACSMRecommendation    *string        `gorm:"type:text" json:"acsmRecommendation"`
	ACSMTags              []string       `gorm:"type:jsonb;serializer:json" json:"acsmTags"`

	// Fotos (único dado próprio — não existe em outro lugar)
	FrontPhotoURL *string `gorm:"type:text" json:"frontPhotoUrl"`
	SidePhotoURL  *string `gorm:"type:text" json:"sidePhotoUrl"`

	// IA
	AIRecommendation *string `gorm:"type:text" json:"aiRecommendation"`

	// Título computado para exibição no frontend (não persistido)
	DisplayTitle string `gorm:"-" json:"displayTitle"`

	// Timestamps
	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relações
	Patient   Patient   `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"patient,omitempty"`
	CreatedBy User      `gorm:"foreignKey:CreatedByID;constraint:OnDelete:RESTRICT" json:"createdBy,omitempty"`
	Anamnesis Anamnesis `gorm:"foreignKey:AnamnesisID;constraint:OnDelete:RESTRICT" json:"anamnesis,omitempty"`
}

// TableName especifica o nome da tabela
func (PhysicalAssessment) TableName() string {
	return "physical_assessments"
}

// GetTitle retorna um título legível para a avaliação
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

// AfterFind popula DisplayTitle após carregar do banco
func (pa *PhysicalAssessment) AfterFind(tx *gorm.DB) error {
	pa.DisplayTitle = pa.GetTitle()
	return nil
}

// BeforeCreate hook to generate UUID v7
func (pa *PhysicalAssessment) BeforeCreate(tx *gorm.DB) error {
	if pa.ID == uuid.Nil {
		pa.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
