package dto

// CreateLabRequestRequest representa o payload de criação de pedido de exames
type CreateLabRequestRequest struct {
	Date                 string  `json:"date" validate:"required"`  // Formato: "2006-01-02"
	Exams                string  `json:"exams" validate:"required"`
	Notes                *string `json:"notes,omitempty"`
	LabRequestTemplateID *string `json:"labRequestTemplateId,omitempty"` // ID do template usado
}

// UpdateLabRequestRequest representa o payload de atualização de pedido de exames
type UpdateLabRequestRequest struct {
	Date                 *string `json:"date,omitempty"`
	Exams                *string `json:"exams,omitempty"`
	Notes                *string `json:"notes,omitempty"`
	LabRequestTemplateID *string `json:"labRequestTemplateId,omitempty"`
}
