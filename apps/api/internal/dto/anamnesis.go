package dto

// CreateAnamnesisRequest representa o payload de criação de anamnese
type CreateAnamnesisRequest struct {
	PatientID               string  `json:"patientId" validate:"required,uuid"`
	ChiefComplaint          string  `json:"chiefComplaint" validate:"required"`
	HistoryOfPresentIllness string  `json:"historyOfPresentIllness,omitempty"`
	PastMedicalHistory      *string `json:"pastMedicalHistory,omitempty"`
	CurrentMedications      *string `json:"currentMedications,omitempty"`
	Allergies               *string `json:"allergies,omitempty"`
	FamilyHistory           *string `json:"familyHistory,omitempty"`
	SocialHistory           *string `json:"socialHistory,omitempty"`
	ReviewOfSystems         *string `json:"reviewOfSystems,omitempty"`
	PhysicalExamination     *string `json:"physicalExamination,omitempty"`
	Assessment              *string `json:"assessment,omitempty"`
	Plan                    *string `json:"plan,omitempty"`
	ConsultationDate        string  `json:"consultationDate" validate:"required"` // formato: RFC3339
	Notes                   *string `json:"notes,omitempty"`
}

// UpdateAnamnesisRequest representa o payload de atualização de anamnese
type UpdateAnamnesisRequest struct {
	ChiefComplaint          *string `json:"chiefComplaint,omitempty"`
	HistoryOfPresentIllness *string `json:"historyOfPresentIllness,omitempty"`
	PastMedicalHistory      *string `json:"pastMedicalHistory,omitempty"`
	CurrentMedications      *string `json:"currentMedications,omitempty"`
	Allergies               *string `json:"allergies,omitempty"`
	FamilyHistory           *string `json:"familyHistory,omitempty"`
	SocialHistory           *string `json:"socialHistory,omitempty"`
	ReviewOfSystems         *string `json:"reviewOfSystems,omitempty"`
	PhysicalExamination     *string `json:"physicalExamination,omitempty"`
	Assessment              *string `json:"assessment,omitempty"`
	Plan                    *string `json:"plan,omitempty"`
	ConsultationDate        *string `json:"consultationDate,omitempty"` // formato: RFC3339
	Notes                   *string `json:"notes,omitempty"`
}

// AnamnesisResponse representa uma anamnese na resposta
type AnamnesisResponse struct {
	ID                      string  `json:"id"`
	PatientID               string  `json:"patientId"`
	DoctorID                string  `json:"doctorId"`
	ChiefComplaint          string  `json:"chiefComplaint"`
	HistoryOfPresentIllness string  `json:"historyOfPresentIllness,omitempty"`
	PastMedicalHistory      *string `json:"pastMedicalHistory,omitempty"`
	CurrentMedications      *string `json:"currentMedications,omitempty"`
	Allergies               *string `json:"allergies,omitempty"`
	FamilyHistory           *string `json:"familyHistory,omitempty"`
	SocialHistory           *string `json:"socialHistory,omitempty"`
	ReviewOfSystems         *string `json:"reviewOfSystems,omitempty"`
	PhysicalExamination     *string `json:"physicalExamination,omitempty"`
	Assessment              *string `json:"assessment,omitempty"`
	Plan                    *string `json:"plan,omitempty"`
	ConsultationDate        string  `json:"consultationDate"`
	Notes                   *string `json:"notes,omitempty"`
	CreatedAt               string  `json:"createdAt"`
	UpdatedAt               string  `json:"updatedAt"`
}
