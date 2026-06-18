package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// LabRequestTemplate representa um template de pedido de exames
// @Description Template pré-configurado com lista de exames para facilitar solicitações
type LabRequestTemplate struct {
	// ID único do template
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Nome do template
	// @example "Check-up Anual", "Avaliação Cardíaca", "Perfil Tireoidiano"
	// @required
	Name string `gorm:"type:varchar(200);not null;unique" json:"name" validate:"required,max=200"`

	// Descrição do template
	// @example "Exames de rotina para check-up anual geral"
	Description *string `gorm:"type:text" json:"description,omitempty"`

	// Ordem de exibição/prioridade
	DisplayOrder int `gorm:"type:integer;not null;default:0" json:"displayOrder"`

	// Status (ativo/inativo)
	IsActive bool `gorm:"type:boolean;not null;default:true;index" json:"isActive"`

	// Timestamps
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relacionamentos

	// Exames incluídos neste template (many-to-many)
	LabTests []LabTestDefinition `gorm:"many2many:lab_request_template_tests;constraint:OnDelete:CASCADE" json:"labTests,omitempty"`
}

// TableName especifica o nome da tabela
func (LabRequestTemplate) TableName() string {
	return "lab_request_templates"
}

// BeforeCreate hook to generate UUID v7
func (lrt *LabRequestTemplate) BeforeCreate(tx *gorm.DB) error {
	if lrt.ID == uuid.Nil {
		lrt.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// LabRequestTemplateTest representa a tabela intermediária many-to-many
// entre templates e exames laboratoriais
type LabRequestTemplateTest struct {
	LabRequestTemplateID uuid.UUID `gorm:"type:uuid;primaryKey" json:"labRequestTemplateId"`
	LabTestDefinitionID  uuid.UUID `gorm:"type:uuid;primaryKey" json:"labTestDefinitionId"`
	// DisplayOrder define a ordem do exame DENTRO do template (fonte da verdade do layout).
	DisplayOrder int `gorm:"type:integer;not null;default:0" json:"displayOrder"`
	// PageBreakBefore: quando true, insere linha em branco (nova página) ANTES deste exame ao
	// aplicar o template. É o único controle de paginação — nenhuma regra de agrupamento no código.
	PageBreakBefore    bool               `gorm:"type:boolean;not null;default:false" json:"pageBreakBefore"`
	CreatedAt          time.Time          `gorm:"autoCreateTime" json:"createdAt"`
	LabRequestTemplate LabRequestTemplate `gorm:"foreignKey:LabRequestTemplateID;constraint:OnDelete:CASCADE" json:"-"`
	LabTestDefinition  LabTestDefinition  `gorm:"foreignKey:LabTestDefinitionID;constraint:OnDelete:CASCADE" json:"-"`
}

// TableName especifica o nome da tabela intermediária
func (LabRequestTemplateTest) TableName() string {
	return "lab_request_template_tests"
}
