package dto

// AnamnesisItemRequest representa um item de anamnese
type AnamnesisItemRequest struct {
	ScoreItemID  string   `json:"scoreItemId" validate:"required,uuid"`
	TextValue    *string  `json:"textValue,omitempty"`
	NumericValue *float64 `json:"numericValue,omitempty"`
	Order        int      `json:"order" validate:"gte=0,lte=9999"`
}

// CreateAnamnesisRequest representa o payload de criação de anamnese
type CreateAnamnesisRequest struct {
	PatientID           string                 `json:"patientId,omitempty" validate:"omitempty,uuid"`                         // Optional - will use selectedPatient if not provided
	AnamnesisTemplateID *string                `json:"anamnesisTemplateId,omitempty" validate:"omitempty,uuid"`               // Template utilizado (opcional)
	ConsultationDate    string                 `json:"consultationDate" validate:"required"`                                  // formato: RFC3339
	Content             *string                `json:"content,omitempty"`                                                     // Conteúdo completo da anamnese
	Summary             *string                `json:"summary,omitempty"`                                                     // Resumo da anamnese
	Visibility          string                 `json:"visibility" validate:"required,oneof=all medicalOnly psychOnly"`        // all, medicalOnly, psychOnly
	Notes               *string                `json:"notes,omitempty"`                                                       // Observações gerais
	Items               []AnamnesisItemRequest `json:"items,omitempty"`                                                       // Items de anamnese vinculados a ScoreItems
}

// UpdateAnamnesisRequest representa o payload de atualização de anamnese
type UpdateAnamnesisRequest struct {
	AnamnesisTemplateID *string                `json:"anamnesisTemplateId,omitempty" validate:"omitempty,uuid"`               // Template utilizado (opcional)
	ConsultationDate    *string                `json:"consultationDate,omitempty"`                                            // formato: RFC3339
	Content             *string                `json:"content,omitempty"`                                                     // Conteúdo completo da anamnese
	Summary             *string                `json:"summary,omitempty"`                                                     // Resumo da anamnese
	Visibility          *string                `json:"visibility,omitempty" validate:"omitempty,oneof=all medicalOnly psychOnly"` // all, medicalOnly, psychOnly
	Notes               *string                `json:"notes,omitempty"`                                                       // Observações gerais
	Items               []AnamnesisItemRequest `json:"items,omitempty"`                                                       // Substitui todos os items existentes
}

// AnamnesisItemResponse representa um item de anamnese na resposta
type AnamnesisItemResponse struct {
	ID           string   `json:"id"`
	ScoreItemID  string   `json:"scoreItemId"`
	TextValue    *string  `json:"textValue,omitempty"`
	NumericValue *float64 `json:"numericValue,omitempty"`
	Order        int      `json:"order"`
	CreatedAt    string   `json:"createdAt"`
	UpdatedAt    string   `json:"updatedAt"`
}

// AnamnesisResponse representa uma anamnese na resposta
type AnamnesisResponse struct {
	ID                  string                  `json:"id"`
	PatientID           string                  `json:"patientId"`
	AuthorID            string                  `json:"authorId"`
	AnamnesisTemplateID *string                 `json:"anamnesisTemplateId,omitempty"`
	ConsultationDate    string                  `json:"consultationDate"`
	Content             *string                 `json:"content,omitempty"`
	Summary             *string                 `json:"summary,omitempty"`
	Visibility          string                  `json:"visibility"`
	Notes               *string                 `json:"notes,omitempty"`
	Items               []AnamnesisItemResponse `json:"items,omitempty"`
	CreatedAt           string                  `json:"createdAt"`
	UpdatedAt           string                  `json:"updatedAt"`
}
