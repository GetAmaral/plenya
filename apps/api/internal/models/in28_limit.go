package models

import (
	"time"

	"github.com/google/uuid"
)

// In28Limit é uma linha do Anexo IV da IN 28/2018 (texto consolidado): o teto diário de um
// nutriente para SUPLEMENTO ALIMENTAR.
//
// A fronteira está no nome: é teto de suplemento, não de prescrição. Fórmula magistral prescrita
// passa dele legitimamente — B12 de 1.000 mcg e vitamina D de 7.000 UI são conduta comum e ficam
// 100× e 3,5× acima do Anexo IV. O que o sistema faz com isto é AVISAR que a fórmula saiu do
// território de suplemento, o que é informação real para o prescritor e para a farmácia.
type In28Limit struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Nutrient string    `gorm:"type:varchar(400);not null;uniqueIndex" json:"nutrient"`
	Unit     string    `gorm:"type:varchar(20);not null" json:"unit"`
	// Coluna "≥19 anos". Nulo quando a norma diz NE (não estabelecido) ou NA (não se aplica).
	MaxAdult *float64 `gorm:"type:numeric(14,4)" json:"maxAdult,omitempty"`
	Kind     string   `gorm:"type:varchar(10);not null;default:'valor'" json:"kind"`
	Source   string   `gorm:"type:varchar(200);not null" json:"source"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

func (In28Limit) TableName() string { return "in28_limits" }
