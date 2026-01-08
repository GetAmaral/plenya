package models

import (
	"time"

	"github.com/google/uuid"
)

// AuditAction define as ações rastreáveis
type AuditAction string

const (
	ActionView   AuditAction = "view"
	ActionCreate AuditAction = "create"
	ActionUpdate AuditAction = "update"
	ActionDelete AuditAction = "delete"
)

// AuditResource define os recursos rastreáveis
type AuditResource string

const (
	ResourcePatients      AuditResource = "patients"
	ResourceAnamnesis     AuditResource = "anamnesis"
	ResourceAppointments  AuditResource = "appointments"
	ResourcePrescriptions AuditResource = "prescriptions"
	ResourceLabResults    AuditResource = "lab_results"
	ResourceHealthScores  AuditResource = "health_scores"
	ResourceUsers         AuditResource = "users"
)

// AuditLog representa um log de auditoria no sistema
// @Description Log de auditoria para rastreabilidade (LGPD obrigatório)
type AuditLog struct {
	// ID único do log
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`

	// ID do usuário que executou a ação
	UserID uuid.UUID `gorm:"type:uuid;not null;index" json:"userId"`

	// Ação executada
	// @enum view,create,update,delete
	Action AuditAction `gorm:"type:varchar(20);not null;check:action IN ('view','create','update','delete')" json:"action"`

	// Recurso acessado
	// @enum patients,anamnesis,lab_results,health_scores,users
	Resource AuditResource `gorm:"type:varchar(50);not null" json:"resource"`

	// ID do recurso acessado (opcional)
	ResourceID *uuid.UUID `gorm:"type:uuid;index" json:"resourceId,omitempty"`

	// Endereço IP do usuário
	IPAddress string `gorm:"type:varchar(45);not null" json:"ipAddress"`

	// User-Agent do navegador/app
	UserAgent *string `gorm:"type:text" json:"userAgent,omitempty"`

	// Se a operação foi bem sucedida
	Success bool `gorm:"type:boolean;default:true;not null" json:"success"`

	// Mensagem de erro (se houver)
	ErrorMessage *string `gorm:"type:text" json:"errorMessage,omitempty"`

	// Data de criação (IMUTÁVEL - nunca atualiza)
	CreatedAt time.Time `gorm:"not null;autoCreateTime;index" json:"createdAt"`

	// Relações
	User User `gorm:"foreignKey:UserID;constraint:OnDelete:RESTRICT" json:"user,omitempty"`
}

// TableName especifica o nome da tabela
func (AuditLog) TableName() string {
	return "audit_logs"
}
