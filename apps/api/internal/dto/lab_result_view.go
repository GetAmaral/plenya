package dto

// CreateLabResultViewRequest representa o payload para criar uma view
type CreateLabResultViewRequest struct {
	Name         string  `json:"name" validate:"required,min=2,max=200"`
	Description  *string `json:"description,omitempty"`
	DisplayOrder int     `json:"displayOrder" validate:"gte=0,lte=9999"`
}

// UpdateLabResultViewRequest representa o payload para atualizar uma view
type UpdateLabResultViewRequest struct {
	Name         *string `json:"name,omitempty" validate:"omitempty,min=2,max=200"`
	Description  *string `json:"description,omitempty"`
	IsActive     *bool   `json:"isActive,omitempty"`
	DisplayOrder *int    `json:"displayOrder,omitempty" validate:"omitempty,gte=0,lte=9999"`
}

// UpdateLabResultViewItemsRequest representa o payload para atualizar items de uma view
type UpdateLabResultViewItemsRequest struct {
	Items []LabResultViewItemData `json:"items" validate:"required,dive"`
}

// LabResultViewItemData representa um item de view no payload
type LabResultViewItemData struct {
	LabTestDefinitionID string `json:"labTestDefinitionId" validate:"required,uuid"`
	Order               int    `json:"order" validate:"gte=0,lte=9999"`
}
