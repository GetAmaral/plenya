package dto

// CreateLabTestUnitConversionRequest representa a requisição para criar uma conversão de unidade
type CreateLabTestUnitConversionRequest struct {
	LabTestDefinitionID string  `json:"labTestDefinitionId" validate:"required,uuid"`
	MainUnit            string  `json:"mainUnit" validate:"required,max=50"`
	SecondaryUnit       string  `json:"secondaryUnit" validate:"required,max=50"`
	ConversionFactor    float64 `json:"conversionFactor" validate:"required,gt=0,lte=1000000"`
}

// UpdateLabTestUnitConversionRequest representa a requisição para atualizar uma conversão
type UpdateLabTestUnitConversionRequest struct {
	MainUnit         *string  `json:"mainUnit,omitempty" validate:"omitempty,max=50"`
	SecondaryUnit    *string  `json:"secondaryUnit,omitempty" validate:"omitempty,max=50"`
	ConversionFactor *float64 `json:"conversionFactor,omitempty" validate:"omitempty,gt=0,lte=1000000"`
}

// LabTestUnitConversionResponse representa a resposta de uma conversão
type LabTestUnitConversionResponse struct {
	ID                  string  `json:"id"`
	LabTestDefinitionID string  `json:"labTestDefinitionId"`
	MainUnit            string  `json:"mainUnit"`
	SecondaryUnit       string  `json:"secondaryUnit"`
	ConversionFactor    float64 `json:"conversionFactor"`
	CreatedAt           string  `json:"createdAt"`
	UpdatedAt           string  `json:"updatedAt"`
}
