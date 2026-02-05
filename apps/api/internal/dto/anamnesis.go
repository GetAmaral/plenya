package dto

// ScoreItemBrief representa um ScoreItem com dados completos para exibição
type ScoreItemBrief struct {
	ID                 string              `json:"id"`
	Name               string              `json:"name"`
	Unit               *string             `json:"unit,omitempty"`
	UnitConversion     *string             `json:"unitConversion,omitempty"`
	ClinicalRelevance  *string             `json:"clinicalRelevance,omitempty"`
	PatientExplanation *string             `json:"patientExplanation,omitempty"`
	Conduct            *string             `json:"conduct,omitempty"`
	LastReview         *string             `json:"lastReview,omitempty"`
	Points             *float64            `json:"points,omitempty"`
	Order              int                 `json:"order"`
	SubgroupID         string              `json:"subgroupId"`
	ParentItemID       *string             `json:"parentItemId,omitempty"`
	Subgroup           *ScoreSubgroupBrief `json:"subgroup,omitempty"`
	Levels             []ScoreLevelBrief   `json:"levels,omitempty"`
	CreatedAt          string              `json:"createdAt"`
	UpdatedAt          string              `json:"updatedAt"`
}

// ScoreSubgroupBrief representa um Subgroup com Group
type ScoreSubgroupBrief struct {
	ID        string          `json:"id"`
	Name      string          `json:"name"`
	Order     int             `json:"order"`
	GroupID   string          `json:"groupId"`
	Group     *ScoreGroupBrief `json:"group,omitempty"`
	CreatedAt string          `json:"createdAt"`
	UpdatedAt string          `json:"updatedAt"`
}

// ScoreGroupBrief representa um Group
type ScoreGroupBrief struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Order     int    `json:"order"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// ScoreLevelBrief representa um Level
type ScoreLevelBrief struct {
	ID                 string  `json:"id"`
	Level              int     `json:"level"`
	Name               string  `json:"name"`
	LowerLimit         *string `json:"lowerLimit,omitempty"`
	UpperLimit         *string `json:"upperLimit,omitempty"`
	Operator           string  `json:"operator"`
	ClinicalRelevance  *string `json:"clinicalRelevance,omitempty"`
	PatientExplanation *string `json:"patientExplanation,omitempty"`
	Conduct            *string `json:"conduct,omitempty"`
	LastReview         *string `json:"lastReview,omitempty"`
	ItemID             string  `json:"itemId"`
	CreatedAt          string  `json:"createdAt"`
	UpdatedAt          string  `json:"updatedAt"`
}

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
	Content             *string                `json:"content,omitempty"`                                                     // Conteúdo completo da anamnese (texto plano)
	ContentHtml         *string                `json:"contentHtml,omitempty"`                                                 // Conteúdo completo da anamnese (HTML)
	Summary             *string                `json:"summary,omitempty"`                                                     // Resumo da anamnese (texto plano)
	SummaryHtml         *string                `json:"summaryHtml,omitempty"`                                                 // Resumo da anamnese (HTML)
	Visibility          string                 `json:"visibility" validate:"required,oneof=all medicalOnly psychOnly authorOnly"`        // all, medicalOnly, psychOnly, authorOnly
	Notes               *string                `json:"notes,omitempty"`                                                       // Observações gerais
	Items               []AnamnesisItemRequest `json:"items,omitempty"`                                                       // Items de anamnese vinculados a ScoreItems
}

// UpdateAnamnesisRequest representa o payload de atualização de anamnese
type UpdateAnamnesisRequest struct {
	AnamnesisTemplateID *string                `json:"anamnesisTemplateId,omitempty" validate:"omitempty,uuid"`               // Template utilizado (opcional)
	ConsultationDate    *string                `json:"consultationDate,omitempty"`                                            // formato: RFC3339
	Content             *string                `json:"content,omitempty"`                                                     // Conteúdo completo da anamnese (texto plano)
	ContentHtml         *string                `json:"contentHtml,omitempty"`                                                 // Conteúdo completo da anamnese (HTML)
	Summary             *string                `json:"summary,omitempty"`                                                     // Resumo da anamnese (texto plano)
	SummaryHtml         *string                `json:"summaryHtml,omitempty"`                                                 // Resumo da anamnese (HTML)
	Visibility          *string                `json:"visibility,omitempty" validate:"omitempty,oneof=all medicalOnly psychOnly authorOnly"` // all, medicalOnly, psychOnly, authorOnly
	Notes               *string                `json:"notes,omitempty"`                                                       // Observações gerais
	Items               []AnamnesisItemRequest `json:"items,omitempty"`                                                       // Substitui todos os items existentes
}

// AnamnesisItemResponse representa um item de anamnese na resposta
type AnamnesisItemResponse struct {
	ID           string          `json:"id"`
	ScoreItemID  string          `json:"scoreItemId"`
	ScoreItem    *ScoreItemBrief `json:"scoreItem,omitempty"`
	TextValue    *string         `json:"textValue,omitempty"`
	NumericValue *float64        `json:"numericValue,omitempty"`
	Order        int             `json:"order"`
	CreatedAt    string          `json:"createdAt"`
	UpdatedAt    string          `json:"updatedAt"`
}

// AnamnesisTemplateBrief representa um AnamnesisTemplate na resposta
type AnamnesisTemplateBrief struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Area string `json:"area"`
}

// AuthorBrief representa um Author (User) na resposta
type AuthorBrief struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

// AnamnesisResponse representa uma anamnese na resposta
type AnamnesisResponse struct {
	ID                  string                   `json:"id"`
	PatientID           string                   `json:"patientId"`
	AuthorID            string                   `json:"authorId"`
	AnamnesisTemplateID *string                  `json:"anamnesisTemplateId,omitempty"`
	ConsultationDate    string                   `json:"consultationDate"`
	Content             *string                  `json:"content,omitempty"`
	ContentHtml         *string                  `json:"contentHtml,omitempty"`
	Summary             *string                  `json:"summary,omitempty"`
	SummaryHtml         *string                  `json:"summaryHtml,omitempty"`
	Visibility          string                   `json:"visibility"`
	Notes               *string                  `json:"notes,omitempty"`
	Items               []AnamnesisItemResponse  `json:"items,omitempty"`
	Author              *AuthorBrief             `json:"author,omitempty"`
	AnamnesisTemplate   *AnamnesisTemplateBrief  `json:"anamnesisTemplate,omitempty"`
	CreatedAt           string                   `json:"createdAt"`
	UpdatedAt           string                   `json:"updatedAt"`
}
