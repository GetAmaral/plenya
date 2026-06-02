package dto

// CreateCarePlanItemRequest — nova recomendação do plano AGIR. patientId vem do path.
type CreateCarePlanItemRequest struct {
	LetterCode       string  `json:"letterCode" validate:"required,oneof=A G I R"`
	Recommendation   string  `json:"recommendation" validate:"required,max=2000"`
	Rationale        *string `json:"rationale,omitempty" validate:"omitempty,max=2000"`
	Target           *string `json:"target,omitempty" validate:"omitempty,max=300"`
	LabTestCode      *string `json:"labTestCode,omitempty" validate:"omitempty,max=20"`
	ScoreItemID      *string `json:"scoreItemId,omitempty" validate:"omitempty,uuid"`
	SourceSnapshotID *string `json:"sourceSnapshotId,omitempty" validate:"omitempty,uuid"`
	Priority         string  `json:"priority,omitempty" validate:"omitempty,oneof=high medium low"`
}

// UpdateCarePlanItemRequest — campos opcionais (PATCH-like).
type UpdateCarePlanItemRequest struct {
	LetterCode     *string `json:"letterCode,omitempty" validate:"omitempty,oneof=A G I R"`
	Recommendation *string `json:"recommendation,omitempty" validate:"omitempty,max=2000"`
	Rationale      *string `json:"rationale,omitempty" validate:"omitempty,max=2000"`
	Target         *string `json:"target,omitempty" validate:"omitempty,max=300"`
	LabTestCode    *string `json:"labTestCode,omitempty" validate:"omitempty,max=20"`
	Priority       *string `json:"priority,omitempty" validate:"omitempty,oneof=high medium low"`
	Status         *string `json:"status,omitempty" validate:"omitempty,oneof=active achieved suspended"`
}

// CarePlanItemResponse — recomendação do plano.
type CarePlanItemResponse struct {
	ID               string  `json:"id"`
	PatientID        string  `json:"patientId"`
	LetterCode       string  `json:"letterCode"`
	Recommendation   string  `json:"recommendation"`
	Rationale        *string `json:"rationale,omitempty"`
	Target           *string `json:"target,omitempty"`
	LabTestCode      *string `json:"labTestCode,omitempty"`
	ScoreItemID      *string `json:"scoreItemId,omitempty"`
	SourceSnapshotID *string `json:"sourceSnapshotId,omitempty"`
	Priority         string  `json:"priority"`
	Status           string  `json:"status"`
	AuthorUserID     string  `json:"authorUserId"`
	AuthorRole       string  `json:"authorRole"`
	AuthorName       string  `json:"authorName,omitempty"`
	CreatedAt        string  `json:"createdAt"`
	UpdatedAt        string  `json:"updatedAt"`
}
