package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// PatientPlanDossier — o prontuário compilado que o plano usou, congelado.
//
// O dossiê vivo é recalculado a cada chamada (~28 queries) e muda quando o prontuário muda. Os dois
// fatos juntos tornam impossível escrever um documento contra ele: no caso que motivou esta tabela,
// o escore da paciente mudou três vezes enquanto o deck estava sendo escrito. Congelar dá chão
// estável para o texto e resposta para "o que a máquina sabia quando esta frase foi escrita".
//
// Nunca é atualizado sozinho. Refrescar troca número debaixo do autor e invalida a base contra a
// qual o conteúdo foi conferido, então é sempre ato explícito.
type PatientPlanDossier struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	PlanID uuid.UUID `gorm:"type:uuid;not null;index" json:"planId"`
	Seq    int       `gorm:"not null" json:"seq"`

	// Payload — o dto.PlanDossierResponse inteiro. Guardado como JSON cru: quem lê desserializa
	// no DTO, e o model não precisa conhecer a forma.
	Payload datatypes.JSON `gorm:"type:jsonb;not null" json:"payload"`

	// Marcas d'água do prontuário no instante do congelamento. Comparar estas três com o estado
	// de agora custa UMA query e diz se o dossiê envelheceu — remontá-lo para descobrir custaria
	// as ~28 que o congelamento existe para evitar.
	SourceSnapshotID *uuid.UUID `gorm:"type:uuid" json:"sourceSnapshotId,omitempty"`
	LatestLabAt      *time.Time `gorm:"type:timestamptz" json:"latestLabAt,omitempty"`
	LatestVitalsAt   *time.Time `gorm:"type:timestamptz" json:"latestVitalsAt,omitempty"`
	LatestSnapshotAt *time.Time `gorm:"type:timestamptz" json:"latestSnapshotAt,omitempty"`

	BuiltAt   time.Time  `gorm:"not null" json:"builtAt"`
	BuiltByID *uuid.UUID `gorm:"type:uuid" json:"builtById,omitempty"`
}

func (PatientPlanDossier) TableName() string { return "patient_plan_dossiers" }

func (d *PatientPlanDossier) BeforeCreate(tx *gorm.DB) error {
	if d.ID == uuid.Nil {
		d.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
