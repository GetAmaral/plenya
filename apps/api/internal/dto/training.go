package dto

import "github.com/plenya/api/internal/models"

// === Exercise DTOs ===

type ExerciseResponse struct {
	ID               string   `json:"id"`
	ExternalId       string   `json:"externalId"`
	Name             string   `json:"name"`
	NamePt           string   `json:"namePt"`
	BodyParts        []string `json:"bodyParts"`
	BodyPartsPt      []string `json:"bodyPartsPt"`
	TargetMuscles    []string `json:"targetMuscles"`
	TargetMusclesPt  []string `json:"targetMusclesPt"`
	Equipments       []string `json:"equipments"`
	EquipmentsPt     []string `json:"equipmentsPt"`
	SecondaryMuscles []string `json:"secondaryMuscles"`
	Instructions     []string `json:"instructions"`
	InstructionsPt   []string `json:"instructionsPt"`
	GifUrl           string   `json:"gifUrl"`
	GifUrlFallback   string   `json:"gifUrlFallback"`
	IsActive         bool     `json:"isActive"`
}

type ExerciseListResponse struct {
	Exercises []ExerciseResponse `json:"exercises"`
	Total     int64              `json:"total"`
}

// === Workout Plan DTOs ===

type CreateWorkoutPlanRequest struct {
	PatientID       string                   `json:"patientId,omitempty"`
	Name            string                   `json:"name" validate:"required,min=2,max=200"`
	Objective       models.WorkoutObjective  `json:"objective" validate:"required,oneof=hypertrophy strength endurance weight_loss conditioning rehabilitation"`
	Intensity       models.WorkoutIntensity  `json:"intensity" validate:"required,oneof=very_light light moderate high very_high"`
	DurationMinutes int                      `json:"durationMinutes" validate:"required,min=10,max=300"`
	WeeklyFrequency int                      `json:"weeklyFrequency" validate:"required,min=1,max=7"`
	Sessions        []CreateSessionRequest   `json:"sessions,omitempty"`
}

type CreateSessionRequest struct {
	Name      string                        `json:"name" validate:"required,min=1,max=100"`
	Order     int                           `json:"order"`
	Exercises []CreateSessionExerciseRequest `json:"exercises,omitempty"`
}

type CreateSessionExerciseRequest struct {
	ExerciseID              string                  `json:"exerciseId" validate:"required,uuid"`
	Phase                   models.ExercisePhase    `json:"phase" validate:"required,oneof=warmup main cooldown"`
	Order                   int                     `json:"order"`
	Sets                    int                     `json:"sets" validate:"required,min=1,max=20"`
	Reps                    string                  `json:"reps" validate:"required"`
	Cadence                 models.ExerciseCadence  `json:"cadence" validate:"required,oneof=normal slow paused explosive free"`
	RestBetweenSetsSec      int                     `json:"restBetweenSetsSec" validate:"min=0,max=600"`
	RestBetweenExercisesSec int                     `json:"restBetweenExercisesSec" validate:"min=0,max=600"`
	Notes                   *string                 `json:"notes,omitempty"`
}

type UpdateWorkoutPlanRequest struct {
	Name            *string                  `json:"name,omitempty" validate:"omitempty,min=2,max=200"`
	Objective       *models.WorkoutObjective `json:"objective,omitempty" validate:"omitempty,oneof=hypertrophy strength endurance weight_loss conditioning rehabilitation"`
	Intensity       *models.WorkoutIntensity `json:"intensity,omitempty" validate:"omitempty,oneof=very_light light moderate high very_high"`
	DurationMinutes *int                     `json:"durationMinutes,omitempty" validate:"omitempty,min=10,max=300"`
	WeeklyFrequency *int                     `json:"weeklyFrequency,omitempty" validate:"omitempty,min=1,max=7"`
	IsActive        *bool                    `json:"isActive,omitempty"`
}

type WorkoutPlanResponse struct {
	ID              string                       `json:"id"`
	PatientID       string                       `json:"patientId"`
	CreatedByID     string                       `json:"createdById"`
	Name            string                       `json:"name"`
	Objective       models.WorkoutObjective      `json:"objective"`
	Intensity       models.WorkoutIntensity      `json:"intensity"`
	DurationMinutes int                          `json:"durationMinutes"`
	WeeklyFrequency int                          `json:"weeklyFrequency"`
	PublicCode      string                       `json:"publicCode"`
	HtmlContent     *string                      `json:"htmlContent,omitempty"`
	IsActive        bool                         `json:"isActive"`
	DisplayTitle    string                       `json:"displayTitle"`
	Sessions        []WorkoutPlanSessionResponse `json:"sessions,omitempty"`
	CreatedAt       string                       `json:"createdAt"`
	UpdatedAt       string                       `json:"updatedAt"`
}

type WorkoutPlanSessionResponse struct {
	ID        string                           `json:"id"`
	PlanID    string                           `json:"planId"`
	Name      string                           `json:"name"`
	Order     int                              `json:"order"`
	Exercises []WorkoutSessionExerciseResponse `json:"exercises,omitempty"`
}

type WorkoutSessionExerciseResponse struct {
	ID                      string                  `json:"id"`
	SessionID               string                  `json:"sessionId"`
	ExerciseID              string                  `json:"exerciseId"`
	Phase                   models.ExercisePhase    `json:"phase"`
	Order                   int                     `json:"order"`
	Sets                    int                     `json:"sets"`
	Reps                    string                  `json:"reps"`
	Cadence                 models.ExerciseCadence  `json:"cadence"`
	RestBetweenSetsSec      int                     `json:"restBetweenSetsSec"`
	RestBetweenExercisesSec int                     `json:"restBetweenExercisesSec"`
	Notes                   *string                 `json:"notes,omitempty"`
	Exercise                *ExerciseResponse       `json:"exercise,omitempty"`
}

// === Physical Assessment DTOs ===

type CreatePhysicalAssessmentRequest struct {
	PatientID      string `json:"patientId,omitempty"`
	AssessmentDate string `json:"assessmentDate" validate:"required"` // format: 2006-01-02

	// Antropometria
	Weight             *float64 `json:"weight"`
	Height             *float64 `json:"height"`
	WaistCircumference *float64 `json:"waistCircumference"`

	// Composição corporal (opcional, calculado se não fornecido)
	BodyFatPercent *float64 `json:"bodyFatPercent"`
	LeanMass       *float64 `json:"leanMass"`

	// Cardiovascular
	SystolicBP       *int `json:"systolicBp"`
	DiastolicBP      *int `json:"diastolicBp"`
	RestingHeartRate *int `json:"restingHeartRate"`

	// Laboratorial (preenchido automaticamente do último LabResult se não fornecido)
	LDL              *float64 `json:"ldl"`
	HDL              *float64 `json:"hdl"`
	TotalCholesterol *float64 `json:"totalCholesterol"`
	Triglycerides    *float64 `json:"triglycerides"`
	FastingGlucose   *float64 `json:"fastingGlucose"`
	HbA1c            *float64 `json:"hbA1c"`

	// Histórico
	FamilyHistory         *bool   `json:"familyHistory"`
	SmokingStatus         *string `json:"smokingStatus" validate:"omitempty,oneof=never former current"`
	PhysicalActivityLevel *string `json:"physicalActivityLevel" validate:"omitempty,oneof=sedentary insufficient moderate active"`

	// Condições
	CardiovascularDisease *bool   `json:"cardiovascularDisease"`
	DiabetesType          *string `json:"diabetesType"`
	Symptoms              *string `json:"symptoms"`
	ClinicalAlert         *bool   `json:"clinicalAlert"`

	// Fotos
	FrontPhotoURL *string `json:"frontPhotoUrl"`
	SidePhotoURL  *string `json:"sidePhotoUrl"`
}

type PhysicalAssessmentResponse struct {
	ID             string `json:"id"`
	PatientID      string `json:"patientId"`
	CreatedByID    string `json:"createdById"`
	AssessmentDate string `json:"assessmentDate"`

	// Antropometria
	Weight             *float64 `json:"weight"`
	Height             *float64 `json:"height"`
	WaistCircumference *float64 `json:"waistCircumference"`

	// Composição corporal
	BMI            *float64 `json:"bmi"`
	BRI            *float64 `json:"bri"`
	BodyFatPercent *float64 `json:"bodyFatPercent"`
	LeanMass       *float64 `json:"leanMass"`

	// Cardiovascular
	SystolicBP       *int `json:"systolicBp"`
	DiastolicBP      *int `json:"diastolicBp"`
	RestingHeartRate *int `json:"restingHeartRate"`

	// Laboratorial
	LDL              *float64 `json:"ldl"`
	HDL              *float64 `json:"hdl"`
	TotalCholesterol *float64 `json:"totalCholesterol"`
	Triglycerides    *float64 `json:"triglycerides"`
	FastingGlucose   *float64 `json:"fastingGlucose"`
	HbA1c            *float64 `json:"hbA1c"`

	// Histórico
	FamilyHistory         *bool   `json:"familyHistory"`
	SmokingStatus         *string `json:"smokingStatus"`
	PhysicalActivityLevel *string `json:"physicalActivityLevel"`

	// Condições
	CardiovascularDisease *bool   `json:"cardiovascularDisease"`
	DiabetesType          *string `json:"diabetesType"`
	Symptoms              *string `json:"symptoms"`
	ClinicalAlert         *bool   `json:"clinicalAlert"`

	// ACSM
	ACSMRiskLevel         *string  `json:"acsmRiskLevel"`
	ACSMRiskFactorsCount  int      `json:"acsmRiskFactorsCount"`
	ACSMRiskFactors       []string `json:"acsmRiskFactors"`
	ACSMProtectiveFactors []string `json:"acsmProtectiveFactors"`
	ACSMRecommendation    *string  `json:"acsmRecommendation"`
	ACSMTags              []string `json:"acsmTags"`

	// Fotos
	FrontPhotoURL *string `json:"frontPhotoUrl"`
	SidePhotoURL  *string `json:"sidePhotoUrl"`

	// IA
	AIRecommendation *string `json:"aiRecommendation"`

	// Zonas FC (calculado, não persistido)
	MaxHR   *int     `json:"maxHr"`
	HRZones []HRZone `json:"hrZones,omitempty"`

	DisplayTitle string `json:"displayTitle"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

type HRZone struct {
	Zone    int    `json:"zone"`
	Name    string `json:"name"`
	MinBPM  int    `json:"minBpm"`
	MaxBPM  int    `json:"maxBpm"`
	Percent string `json:"percent"`
}

// === Periodization DTOs ===

type CreatePeriodizationRequest struct {
	PatientID  string                         `json:"patientId,omitempty"`
	Framework  models.PeriodizationFramework  `json:"framework" validate:"required,oneof=bompa linear undulating block"`
	StartDate  string                         `json:"startDate" validate:"required"`
	TotalWeeks int                            `json:"totalWeeks" validate:"required,min=1,max=104"`
	Objective  string                         `json:"objective" validate:"required,min=2,max=200"`
}

type PeriodizationResponse struct {
	ID                      string                         `json:"id"`
	PatientID               string                         `json:"patientId"`
	CreatedByID             string                         `json:"createdById"`
	Framework               models.PeriodizationFramework  `json:"framework"`
	StartDate               string                         `json:"startDate"`
	TotalWeeks              int                            `json:"totalWeeks"`
	Objective               string                         `json:"objective"`
	ScientificJustification *string                        `json:"scientificJustification"`
	DisplayTitle            string                         `json:"displayTitle"`
	Mesocycles              []MesocycleResponse            `json:"mesocycles,omitempty"`
	CreatedAt               string                         `json:"createdAt"`
	UpdatedAt               string                         `json:"updatedAt"`
}

type MesocycleResponse struct {
	ID                 string                 `json:"id"`
	PeriodizationID    string                 `json:"periodizationId"`
	Order              int                    `json:"order"`
	Phase              models.MesocyclePhase  `json:"phase"`
	DurationWeeks      int                    `json:"durationWeeks"`
	VolumePercent      int                    `json:"volumePercent"`
	IntensityPercent   int                    `json:"intensityPercent"`
	PhysiologicalFocus string                 `json:"physiologicalFocus"`
	StartDate          string                 `json:"startDate"`
	EndDate            string                 `json:"endDate"`
}
