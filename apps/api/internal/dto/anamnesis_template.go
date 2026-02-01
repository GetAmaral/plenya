package dto

// CreateAnamnesisTemplateRequest represents the request to create an anamnesis template
type CreateAnamnesisTemplateRequest struct {
	Name string `json:"name" validate:"required,min=2,max=200"`
	Area string `json:"area" validate:"required,oneof=Medicina Nutricao Psicologia 'Educacao Fisica'"`
}

// UpdateAnamnesisTemplateRequest represents the request to update an anamnesis template
type UpdateAnamnesisTemplateRequest struct {
	Name *string `json:"name,omitempty" validate:"omitempty,min=2,max=200"`
	Area *string `json:"area,omitempty" validate:"omitempty,oneof=Medicina Nutricao Psicologia 'Educacao Fisica'"`
}

// UpdateAnamnesisTemplateItemsRequest represents the request to update template items
type UpdateAnamnesisTemplateItemsRequest struct {
	Items []AnamnesisTemplateItemData `json:"items" validate:"required,dive"`
}

// AnamnesisTemplateItemData represents an item in the update request
type AnamnesisTemplateItemData struct {
	ScoreItemID string `json:"scoreItemId" validate:"required,uuid"`
	Order       int    `json:"order" validate:"gte=0,lte=9999"`
}
