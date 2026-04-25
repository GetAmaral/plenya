package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AppointmentResourceType enumera os tipos de recurso alocado num appointment.
//
// V1: só "doctor" é gravado. "room" e "equipment" ficam pra V2 — o schema é
// criado agora pra evitar migration breaking depois (campos nullable, índices
// preparados).
type AppointmentResourceType string

const (
	AppointmentResourceDoctor    AppointmentResourceType = "doctor"
	AppointmentResourceRoom      AppointmentResourceType = "room"
	AppointmentResourceEquipment AppointmentResourceType = "equipment"
)

// AppointmentResource representa um recurso (médico, sala, equipamento)
// alocado a um appointment.
//
// Schema preparado pra V2:
//   - V1: AppointmentService só cria 1 row (ResourceType="doctor", ResourceID=DoctorID).
//   - V2: UI permite alocar sala/equipamento; slot picker passa a considerar
//     conflito multi-recurso (não só do médico).
//
// @Description Recurso alocado a um appointment (V1: só doctor)
type AppointmentResource struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Appointment dono (FK appointments.id, OnDelete CASCADE).
	AppointmentID uuid.UUID `gorm:"type:uuid;not null;index" json:"appointmentId"`

	// Tipo do recurso. CHECK constraint enforce os 3 valores válidos.
	ResourceType string `gorm:"type:varchar(20);not null;check:resource_type IN ('doctor','room','equipment')" json:"resourceType" validate:"oneof=doctor room equipment"`

	// ID do recurso (users.id se doctor; rooms.id ou equipment.id em V2).
	// Não usamos FK polimórfica — validação é a nível de service.
	ResourceID uuid.UUID `gorm:"type:uuid;not null;index" json:"resourceId"`

	CreatedAt time.Time `gorm:"not null;autoCreateTime" json:"createdAt"`
}

// TableName especifica o nome da tabela
func (AppointmentResource) TableName() string {
	return "appointment_resources"
}

// BeforeCreate hook gera UUID v7
func (a *AppointmentResource) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
