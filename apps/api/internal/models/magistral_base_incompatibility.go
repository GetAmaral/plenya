package models

import (
	"time"

	"github.com/google/uuid"
)

// MagistralBaseIncompatibility é uma incompatibilidade entre um ativo e a BASE (veículo).
//
// A tabela de pares cobre ativo × ativo, que é a minoria do problema real: no levantamento de 400
// prescrições magistrais da farmácia-escola da UFRJ, 63% dos erros farmacotécnicos são ativo ×
// formulação e 23% ativo × base, contra 13% ativo × ativo.
//
// Casa por texto porque o veículo é campo livre na fórmula ("creme Lanette", "creme não iônico").
// MinPercent existe porque várias dessas regras são de concentração e não de presença: ácido
// acima de 10% derruba a emulsão aniônica, abaixo disso convive.
type MagistralBaseIncompatibility struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	BasePattern string    `gorm:"type:varchar(120);not null" json:"basePattern"`
	// Nulo = vale para qualquer componente que satisfaça as demais condições.
	SubstancePattern *string  `gorm:"type:varchar(120)" json:"substancePattern,omitempty"`
	MinPercent       *float64 `gorm:"type:numeric(6,2)" json:"minPercent,omitempty"`

	Severity       IncompatibilitySeverity `gorm:"type:varchar(10);not null;default:'warn'" json:"severity"`
	Mechanism      string                  `gorm:"type:text;not null" json:"mechanism"`
	Recommendation string                  `gorm:"type:text;not null;default:''" json:"recommendation"`
	Source         string                  `gorm:"type:varchar(300);not null;default:''" json:"source"`
	IsActive       bool                    `gorm:"not null;default:true" json:"isActive"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

func (MagistralBaseIncompatibility) TableName() string {
	return "magistral_base_incompatibilities"
}
