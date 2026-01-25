# CAC Score & CCTA - EMR Implementation Guide

## Overview

This guide provides step-by-step instructions for implementing CAC Score and CCTA functionality in the Plenya EMR system, following the established architecture patterns.

---

## Phase 1: Database Models (Go)

### Step 1.1: Create CAC Score Model

**File:** `/home/user/plenya/apps/api/internal/models/cac_score.go`

```go
package models

import (
    "time"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

// CACScore represents Coronary Artery Calcium Score (Agatston)
// @Description CAC Score for cardiovascular risk stratification
type CACScore struct {
    ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`

    // @required
    PatientID uuid.UUID `gorm:"type:uuid;not null;index" json:"patientId" validate:"required" example:"550e8400-e29b-41d4-a716-446655440001"`

    // Exam Information
    // @required
    ExamDate time.Time `gorm:"not null" json:"examDate" validate:"required" example:"2024-01-15T10:00:00Z"`

    // @maxLength 200
    Institution string `gorm:"type:varchar(200)" json:"institution,omitempty" example:"Hospital São Paulo - Tomografia"`

    // CAC Score (Agatston)
    // @required
    // @minimum 0
    AgatstonScore int `gorm:"not null;check:agatston_score >= 0" json:"agatstonScore" validate:"required,gte=0" example:"150"`

    // @minimum 0
    // @maximum 5
    // @description Risk level: 0=Very High (>1000), 1=High (401-1000), 2=Moderate-High (101-400), 3=Moderate (11-100), 4=Low (1-10), 5=Very Low (0)
    AgatstonRiskLevel int `gorm:"not null;check:agatston_risk_level BETWEEN 0 AND 5" json:"agatstonRiskLevel" example:"2"`

    // Percentile (age/sex/race-specific from MESA)
    // @minimum 0
    // @maximum 100
    Percentile *int `gorm:"check:percentile BETWEEN 0 AND 100" json:"percentile,omitempty" example:"75"`

    // @minimum 0
    // @maximum 4
    // @description 0=≥90th percentile, 1=75-89th, 2=50-74th, 3=25-49th, 4=<25th
    PercentileRiskLevel *int `gorm:"check:percentile_risk_level BETWEEN 0 AND 4" json:"percentileRiskLevel,omitempty" example:"1"`

    // CAC Density
    // @enum low,mixed,moderate,high
    CACDensity *string `gorm:"type:varchar(50);check:cac_density IN ('low','mixed','moderate','high')" json:"cacDensity,omitempty" example:"moderate"`

    // @minimum 0
    // @maximum 3
    CACDensityRiskLevel *int `gorm:"check:cac_density_risk_level BETWEEN 0 AND 3" json:"cacDensityRiskLevel,omitempty" example:"2"`

    // Vessel-specific scores
    // @minimum 0
    LMScore *int `gorm:"check:lm_score >= 0" json:"lmScore,omitempty" example:"0"` // Left Main

    // @minimum 0
    LADScore *int `gorm:"check:lad_score >= 0" json:"ladScore,omitempty" example:"85"` // Left Anterior Descending

    // @minimum 0
    LCXScore *int `gorm:"check:lcx_score >= 0" json:"lcxScore,omitempty" example:"35"` // Left Circumflex

    // @minimum 0
    RCAScore *int `gorm:"check:rca_score >= 0" json:"rcaScore,omitempty" example:"30"` // Right Coronary Artery

    // Progression (if repeat scan)
    PreviousScanID *uuid.UUID `gorm:"type:uuid" json:"previousScanId,omitempty" example:"550e8400-e29b-41d4-a716-446655440002"`

    // @description Annual percentage change in CAC score
    ProgressionRate *float64 `json:"progressionRate,omitempty" example:"18.5"`

    // @minimum 0
    // @maximum 3
    // @description 0=>30%/year, 1=20-30%/year, 2=15-20%/year, 3=<15%/year
    ProgressionRiskLevel *int `gorm:"check:progression_risk_level BETWEEN 0 AND 3" json:"progressionRiskLevel,omitempty" example:"2"`

    // Clinical recommendations
    StatinRecommended bool `gorm:"default:false" json:"statinRecommended" example:"true"`

    // @minimum 1
    // @maximum 15
    // @description Years until next CAC scan recommended (only if CAC=0)
    WarrantyPeriodYears *int `gorm:"check:warranty_period_years BETWEEN 1 AND 15" json:"warrantyPeriodYears,omitempty" example:"5"`

    // @maxLength 2000
    ClinicalNotes string `gorm:"type:text" json:"clinicalNotes,omitempty" example:"Patient counseled on lifestyle modifications"`

    // Audit
    CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
    UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

    // Relations
    Patient       Patient   `gorm:"foreignKey:PatientID" json:"-"`
    PreviousScan  *CACScore `gorm:"foreignKey:PreviousScanID" json:"previousScan,omitempty"`
}

// TableName overrides the table name
func (CACScore) TableName() string {
    return "cac_scores"
}

// BeforeSave hooks
func (c *CACScore) BeforeSave(tx *gorm.DB) error {
    c.CalculateRiskLevel()
    return nil
}

// CalculateRiskLevel based on Agatston score
func (c *CACScore) CalculateRiskLevel() {
    switch {
    case c.AgatstonScore == 0:
        c.AgatstonRiskLevel = 5 // Very Low
    case c.AgatstonScore <= 10:
        c.AgatstonRiskLevel = 4 // Low
    case c.AgatstonScore <= 100:
        c.AgatstonRiskLevel = 3 // Moderate
    case c.AgatstonScore <= 400:
        c.AgatstonRiskLevel = 2 // Moderate-High
    case c.AgatstonScore <= 1000:
        c.AgatstonRiskLevel = 1 // High
    default:
        c.AgatstonRiskLevel = 0 // Very High
    }

    // Calculate density risk level
    if c.CACDensity != nil {
        switch *c.CACDensity {
        case "low":
            level := 0
            c.CACDensityRiskLevel = &level
        case "mixed":
            level := 1
            c.CACDensityRiskLevel = &level
        case "moderate":
            level := 2
            c.CACDensityRiskLevel = &level
        case "high":
            level := 3
            c.CACDensityRiskLevel = &level
        }
    }

    // Statin recommendation: CAC ≥100
    c.StatinRecommended = c.AgatstonScore >= 100
}

// CalculateWarrantyPeriod based on patient demographics and risk factors
func (c *CACScore) CalculateWarrantyPeriod(birthDate time.Time, gender string, lpa *float64, diabetes bool) {
    // Only calculate if CAC = 0
    if c.AgatstonScore > 0 {
        c.WarrantyPeriodYears = nil
        return
    }

    age := time.Now().Year() - birthDate.Year()
    hasHighLpa := lpa != nil && *lpa > 50
    isMale := gender == "male"

    years := 10 // Default

    // Diabetes shortens warranty period significantly
    if diabetes {
        years = 3
        c.WarrantyPeriodYears = &years
        return
    }

    // Age and Lp(a) based calculation
    if isMale {
        if age < 45 {
            if hasHighLpa {
                years = 5
            } else {
                years = 12
            }
        } else if age < 55 {
            if hasHighLpa {
                years = 5
            } else {
                years = 7
            }
        } else {
            if hasHighLpa {
                years = 3
            } else {
                years = 5
            }
        }
    } else { // Female
        if age < 55 {
            if hasHighLpa {
                years = 7
            } else {
                years = 12
            }
        } else {
            if hasHighLpa {
                years = 5
            } else {
                years = 7
            }
        }
    }

    c.WarrantyPeriodYears = &years
}

// CalculateProgression if previous scan exists
func (c *CACScore) CalculateProgression(previousScore int, previousDate time.Time) {
    if c.AgatstonScore == 0 || previousScore == 0 {
        return // Cannot calculate progression from/to zero
    }

    years := c.ExamDate.Sub(previousDate).Hours() / 24 / 365
    if years <= 0 {
        return
    }

    // Annual percentage change
    progression := float64(c.AgatstonScore-previousScore) / float64(previousScore) / years * 100
    c.ProgressionRate = &progression

    // Determine risk level
    switch {
    case progression < 15:
        level := 3 // Normal
        c.ProgressionRiskLevel = &level
    case progression < 20:
        level := 2 // Typical
        c.ProgressionRiskLevel = &level
    case progression < 30:
        level := 1 // Accelerated
        c.ProgressionRiskLevel = &level
    default:
        level := 0 // Very Accelerated
        c.ProgressionRiskLevel = &level
    }
}

// GetRiskLabel returns human-readable risk label
func (c *CACScore) GetRiskLabel() string {
    labels := []string{
        "Risco Muito Alto", // 0
        "Risco Alto",       // 1
        "Risco Moderado-Alto", // 2
        "Risco Moderado",   // 3
        "Risco Baixo",      // 4
        "Risco Muito Baixo", // 5
    }
    if c.AgatstonRiskLevel >= 0 && c.AgatstonRiskLevel < len(labels) {
        return labels[c.AgatstonRiskLevel]
    }
    return "Desconhecido"
}

// GetAnnualEventRate returns estimated annual cardiovascular event rate
func (c *CACScore) GetAnnualEventRate() string {
    rates := []string{
        "3-4%",   // 0: >1000
        "2.1%",   // 1: 401-1000
        "1.4%",   // 2: 101-400
        "0.7%",   // 3: 11-100
        "0.3%",   // 4: 1-10
        "0.11%",  // 5: 0
    }
    if c.AgatstonRiskLevel >= 0 && c.AgatstonRiskLevel < len(rates) {
        return rates[c.AgatstonRiskLevel]
    }
    return "N/A"
}
```

### Step 1.2: Create CCTA Model

**File:** `/home/user/plenya/apps/api/internal/models/ccta.go`

```go
package models

import (
    "time"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

// CCTA represents Coronary CT Angiography results
// @Description CCTA with CAD-RADS classification and high-risk plaque features
type CCTA struct {
    ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`

    // @required
    PatientID uuid.UUID `gorm:"type:uuid;not null;index" json:"patientId" validate:"required" example:"550e8400-e29b-41d4-a716-446655440001"`

    // Exam Information
    // @required
    ExamDate time.Time `gorm:"not null" json:"examDate" validate:"required" example:"2024-01-15T10:00:00Z"`

    // @maxLength 200
    Institution string `gorm:"type:varchar(200)" json:"institution,omitempty" example:"Hospital São Paulo - Tomografia"`

    // CAD-RADS Classification
    // @required
    // @enum 0,1,2,3,4A,4B,5,N,S,G
    CADRADSScore string `gorm:"type:varchar(10);not null;check:cad_rads_score IN ('0','1','2','3','4A','4B','5','N','S','G')" json:"cadRadsScore" validate:"required" example:"3"`

    // @minimum 0
    // @maximum 6
    // @description Risk level: 0=CAD-RADS 5, 1=4B, 2=4A, 3=3, 4=2, 5=1, 6=0
    CADRADSRiskLevel int `gorm:"not null;check:cad_rads_risk_level BETWEEN 0 AND 6" json:"cadRadsRiskLevel" example:"3"`

    // Plaque Burden Score (P0-P4)
    // @enum P0,P1,P2,P3,P4
    PlaqueBurdenScore *string `gorm:"type:varchar(10);check:plaque_burden_score IN ('P0','P1','P2','P3','P4')" json:"plaqueBurdenScore,omitempty" example:"P2"`

    // @minimum 0
    // @maximum 4
    PlaqueBurdenRiskLevel *int `gorm:"check:plaque_burden_risk_level BETWEEN 0 AND 4" json:"plaqueBurdenRiskLevel,omitempty" example:"2"`

    // Segment Involvement Score (0-16)
    // @minimum 0
    // @maximum 16
    SegmentInvolvementScore *int `gorm:"check:segment_involvement_score BETWEEN 0 AND 16" json:"segmentInvolvementScore,omitempty" example:"6"`

    // High-Risk Plaque Features
    PositiveRemodeling bool `gorm:"default:false" json:"positiveRemodeling" example:"true"`
    LowAttenuationPlaque bool `gorm:"default:false" json:"lowAttenuationPlaque" example:"true"`
    NapkinRingSign bool `gorm:"default:false" json:"napkinRingSign" example:"false"`
    SpottyCalcification bool `gorm:"default:false" json:"spottyCalcification" example:"true"`

    // @minimum 0
    // @maximum 4
    HRPFeatureCount int `gorm:"check:hrp_feature_count BETWEEN 0 AND 4" json:"hrpFeatureCount" example:"3"`

    // @minimum 0
    // @maximum 4
    // @description HRP risk: 4=0 features, 3=1 feature, 2=2 features, 1=3 features, 0=4 features
    HRPRiskLevel int `gorm:"check:hrp_risk_level BETWEEN 0 AND 4" json:"hrpRiskLevel" example:"1"`

    // Stenosis details (worst lesion)
    // @minimum 0
    // @maximum 100
    WorstStenosisPercent *int `gorm:"check:worst_stenosis_percent BETWEEN 0 AND 100" json:"worstStenosisPercent,omitempty" example:"65"`

    // @maxLength 100
    // @enum LM,LAD_proximal,LAD_mid,LAD_distal,LCX_proximal,LCX_distal,RCA_proximal,RCA_mid,RCA_distal
    WorstStenosisLocation *string `gorm:"type:varchar(100)" json:"worstStenosisLocation,omitempty" example:"LAD_proximal"`

    // FFR-CT (Fractional Flow Reserve by CT)
    // @minimum 0
    // @maximum 1
    FFRCT *float64 `gorm:"check:ffrct BETWEEN 0 AND 1" json:"ffrct,omitempty" example:"0.78"`

    FFRCTPositive *bool `json:"ffrctPositive,omitempty" example:"true"` // <0.80 = positive

    // Clinical recommendations (auto-generated)
    InvasiveAngiography bool `gorm:"default:false" json:"invasiveAngiography" example:"false"`
    StressTestRecommended bool `gorm:"default:false" json:"stressTestRecommended" example:"true"`

    // @enum lifestyle,moderate,intensive,maximal
    MedicalTherapyIntensity string `gorm:"type:varchar(50);check:medical_therapy_intensity IN ('lifestyle','moderate','intensive','maximal')" json:"medicalTherapyIntensity" example:"intensive"`

    // Associated CAC Score
    CACScoreID *uuid.UUID `gorm:"type:uuid" json:"cacScoreId,omitempty" example:"550e8400-e29b-41d4-a716-446655440002"`

    // @maxLength 2000
    ClinicalNotes string `gorm:"type:text" json:"clinicalNotes,omitempty" example:"High-risk plaque features present"`

    // Audit
    CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
    UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

    // Relations
    Patient  Patient   `gorm:"foreignKey:PatientID" json:"-"`
    CACScore *CACScore `gorm:"foreignKey:CACScoreID" json:"cacScore,omitempty"`
}

// TableName overrides the table name
func (CCTA) TableName() string {
    return "cctas"
}

// BeforeSave hooks
func (c *CCTA) BeforeSave(tx *gorm.DB) error {
    c.CalculateRiskLevels()
    c.DetermineRecommendations()
    return nil
}

// CalculateRiskLevels based on CAD-RADS and HRP features
func (c *CCTA) CalculateRiskLevels() {
    // CAD-RADS risk level
    switch c.CADRADSScore {
    case "0":
        c.CADRADSRiskLevel = 6
    case "1":
        c.CADRADSRiskLevel = 5
    case "2":
        c.CADRADSRiskLevel = 4
    case "3":
        c.CADRADSRiskLevel = 3
    case "4A":
        c.CADRADSRiskLevel = 2
    case "4B":
        c.CADRADSRiskLevel = 1
    case "5":
        c.CADRADSRiskLevel = 0
    default:
        c.CADRADSRiskLevel = 6 // N, S, G default to lowest risk
    }

    // HRP feature count
    count := 0
    if c.PositiveRemodeling {
        count++
    }
    if c.LowAttenuationPlaque {
        count++
    }
    if c.NapkinRingSign {
        count++
    }
    if c.SpottyCalcification {
        count++
    }
    c.HRPFeatureCount = count
    c.HRPRiskLevel = 4 - count // Inverse: more features = higher risk (lower level number)

    // Plaque burden risk level
    if c.PlaqueBurdenScore != nil {
        switch *c.PlaqueBurdenScore {
        case "P0":
            level := 4
            c.PlaqueBurdenRiskLevel = &level
        case "P1":
            level := 3
            c.PlaqueBurdenRiskLevel = &level
        case "P2":
            level := 2
            c.PlaqueBurdenRiskLevel = &level
        case "P3":
            level := 1
            c.PlaqueBurdenRiskLevel = &level
        case "P4":
            level := 0
            c.PlaqueBurdenRiskLevel = &level
        }
    }

    // FFR-CT positive if <0.80
    if c.FFRCT != nil {
        positive := *c.FFRCT < 0.80
        c.FFRCTPositive = &positive
    }
}

// DetermineRecommendations based on CAD-RADS and HRP
func (c *CCTA) DetermineRecommendations() {
    switch c.CADRADSScore {
    case "0", "1":
        c.InvasiveAngiography = false
        c.StressTestRecommended = false
        c.MedicalTherapyIntensity = "lifestyle"
    case "2":
        c.InvasiveAngiography = false
        c.StressTestRecommended = false
        c.MedicalTherapyIntensity = "moderate"
    case "3":
        c.InvasiveAngiography = false
        c.StressTestRecommended = true
        c.MedicalTherapyIntensity = "intensive"
    case "4A":
        c.InvasiveAngiography = true // Likely
        c.StressTestRecommended = false
        c.MedicalTherapyIntensity = "maximal"
    case "4B", "5":
        c.InvasiveAngiography = true // Urgent
        c.StressTestRecommended = false
        c.MedicalTherapyIntensity = "maximal"
    }

    // Upgrade intensity if HRP features present
    if c.HRPFeatureCount >= 2 && c.MedicalTherapyIntensity != "maximal" {
        c.MedicalTherapyIntensity = "intensive"
    }

    // FFR-CT positive upgrades recommendations
    if c.FFRCTPositive != nil && *c.FFRCTPositive {
        c.InvasiveAngiography = true
        c.MedicalTherapyIntensity = "maximal"
    }
}

// GetRiskLabel returns human-readable risk label
func (c *CCTA) GetRiskLabel() string {
    labels := []string{
        "Oclusão Total",           // 0: CAD-RADS 5
        "Estenose Severa 3v/LM",   // 1: CAD-RADS 4B
        "Estenose Severa 1-2v",    // 2: CAD-RADS 4A
        "Estenose Moderada",       // 3: CAD-RADS 3
        "Estenose Leve",           // 4: CAD-RADS 2
        "Estenose Mínima",         // 5: CAD-RADS 1
        "Sem Placa",               // 6: CAD-RADS 0
    }
    if c.CADRADSRiskLevel >= 0 && c.CADRADSRiskLevel < len(labels) {
        return labels[c.CADRADSRiskLevel]
    }
    return "Não Classificado"
}

// Get5YearEventFreeRate returns 5-year event-free survival rate
func (c *CCTA) Get5YearEventFreeRate() string {
    rates := []string{
        "69.3%",  // 0: CAD-RADS 5
        "76.7%",  // 1: CAD-RADS 4B
        "76.7%",  // 2: CAD-RADS 4A
        "84.5%",  // 3: CAD-RADS 3
        "88.7%",  // 4: CAD-RADS 2
        "92.9%",  // 5: CAD-RADS 1
        "95.2%",  // 6: CAD-RADS 0
    }
    if c.CADRADSRiskLevel >= 0 && c.CADRADSRiskLevel < len(rates) {
        return rates[c.CADRADSRiskLevel]
    }
    return "N/A"
}
```

---

## Phase 2: Generate Migrations and Types

### Step 2.1: Generate Database Migrations

```bash
cd /home/user/plenya

# Generate migrations from Go models
pnpm generate:migrations

# Apply migrations to database
cd apps/api
atlas migrate apply --env dev
```

### Step 2.2: Generate TypeScript Types and Zod Schemas

```bash
cd /home/user/plenya

# Generate OpenAPI spec, TypeScript types, and Zod schemas
pnpm generate
```

---

## Phase 3: Backend API Implementation

### Step 3.1: Create Repository

**File:** `/home/user/plenya/apps/api/internal/repository/cac_score_repository.go`

```go
package repository

import (
    "github.com/google/uuid"
    "gorm.io/gorm"
    "plenya/internal/models"
)

type CACScoreRepository struct {
    db *gorm.DB
}

func NewCACScoreRepository(db *gorm.DB) *CACScoreRepository {
    return &CACScoreRepository{db: db}
}

func (r *CACScoreRepository) Create(cac *models.CACScore) error {
    return r.db.Create(cac).Error
}

func (r *CACScoreRepository) FindByID(id string) (*models.CACScore, error) {
    var cac models.CACScore
    err := r.db.Preload("PreviousScan").First(&cac, "id = ?", id).Error
    return &cac, err
}

func (r *CACScoreRepository) FindByPatientID(patientID string) ([]models.CACScore, error) {
    var cacs []models.CACScore
    err := r.db.Where("patient_id = ?", patientID).
        Order("exam_date DESC").
        Find(&cacs).Error
    return cacs, err
}

func (r *CACScoreRepository) Update(cac *models.CACScore) error {
    return r.db.Save(cac).Error
}

func (r *CACScoreRepository) Delete(id string) error {
    return r.db.Delete(&models.CACScore{}, "id = ?", id).Error
}

func (r *CACScoreRepository) GetLatestByPatient(patientID string) (*models.CACScore, error) {
    var cac models.CACScore
    err := r.db.Where("patient_id = ?", patientID).
        Order("exam_date DESC").
        First(&cac).Error
    return &cac, err
}
```

### Step 3.2: Create Handlers

**File:** `/home/user/plenya/apps/api/internal/handlers/cac_score_handler.go`

```go
package handlers

import (
    "time"
    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
    "plenya/internal/models"
    "plenya/internal/repository"
)

type CACScoreHandler struct {
    cacRepo     *repository.CACScoreRepository
    patientRepo *repository.PatientRepository
    labRepo     *repository.LabResultRepository
}

func NewCACScoreHandler(
    cacRepo *repository.CACScoreRepository,
    patientRepo *repository.PatientRepository,
    labRepo *repository.LabResultRepository,
) *CACScoreHandler {
    return &CACScoreHandler{
        cacRepo:     cacRepo,
        patientRepo: patientRepo,
        labRepo:     labRepo,
    }
}

// CreateCACScore godoc
// @Summary Create new CAC Score
// @Description Create CAC Score with automatic risk calculation
// @Tags cac-scores
// @Accept json
// @Produce json
// @Param id path string true "Patient ID"
// @Param body body CreateCACScoreRequest true "CAC Score data"
// @Success 201 {object} models.CACScore
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/patients/{id}/cac-scores [post]
func (h *CACScoreHandler) Create(c *fiber.Ctx) error {
    var req CreateCACScoreRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
    }

    patientID := c.Params("id")

    // Get patient
    patient, err := h.patientRepo.FindByID(patientID)
    if err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Patient not found"})
    }

    // Get latest Lp(a) if available
    var lpaValue *float64
    labs, _ := h.labRepo.GetLatestByPatientAndExam(patientID, "lpa")
    if len(labs) > 0 {
        val := labs[0].Value
        lpaValue = &val
    }

    // Check diabetes (simplified - should check diagnoses table)
    hasDiabetes := false // TODO: Check patient diagnoses

    // Create CAC score
    cac := &models.CACScore{
        PatientID:      uuid.MustParse(patientID),
        ExamDate:       req.ExamDate,
        Institution:    req.Institution,
        AgatstonScore:  req.AgatstonScore,
        Percentile:     req.Percentile,
        CACDensity:     req.CACDensity,
        LMScore:        req.LMScore,
        LADScore:       req.LADScore,
        LCXScore:       req.LCXScore,
        RCAScore:       req.RCAScore,
        PreviousScanID: req.PreviousScanID,
        ClinicalNotes:  req.ClinicalNotes,
    }

    // Calculate risk levels (done in BeforeSave hook)
    cac.CalculateWarrantyPeriod(patient.BirthDate, patient.Gender, lpaValue, hasDiabetes)

    // Calculate progression if previous scan exists
    if req.PreviousScanID != nil {
        prevScan, err := h.cacRepo.FindByID(req.PreviousScanID.String())
        if err == nil {
            cac.CalculateProgression(prevScan.AgatstonScore, prevScan.ExamDate)
        }
    }

    // Save
    if err := h.cacRepo.Create(cac); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to create CAC score"})
    }

    return c.Status(201).JSON(cac)
}

// GetByPatient godoc
// @Summary Get CAC Scores by patient
// @Description Get all CAC Scores for a patient
// @Tags cac-scores
// @Produce json
// @Param id path string true "Patient ID"
// @Success 200 {array} models.CACScore
// @Router /api/v1/patients/{id}/cac-scores [get]
func (h *CACScoreHandler) GetByPatient(c *fiber.Ctx) error {
    patientID := c.Params("id")
    cacs, err := h.cacRepo.FindByPatientID(patientID)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch CAC scores"})
    }
    return c.JSON(cacs)
}

// Request DTOs
type CreateCACScoreRequest struct {
    ExamDate       time.Time  `json:"examDate" validate:"required"`
    Institution    string     `json:"institution"`
    AgatstonScore  int        `json:"agatstonScore" validate:"required,gte=0"`
    Percentile     *int       `json:"percentile" validate:"omitempty,gte=0,lte=100"`
    CACDensity     *string    `json:"cacDensity" validate:"omitempty,oneof=low mixed moderate high"`
    LMScore        *int       `json:"lmScore" validate:"omitempty,gte=0"`
    LADScore       *int       `json:"ladScore" validate:"omitempty,gte=0"`
    LCXScore       *int       `json:"lcxScore" validate:"omitempty,gte=0"`
    RCAScore       *int       `json:"rcaScore" validate:"omitempty,gte=0"`
    PreviousScanID *uuid.UUID `json:"previousScanId"`
    ClinicalNotes  string     `json:"clinicalNotes"`
}
```

### Step 3.3: Register Routes

**File:** `/home/user/plenya/apps/api/cmd/server/routes.go`

```go
// Register CAC Score routes
cacRepo := repository.NewCACScoreRepository(db)
cacHandler := handlers.NewCACScoreHandler(cacRepo, patientRepo, labRepo)

api.Post("/patients/:id/cac-scores", middleware.Auth(), cacHandler.Create)
api.Get("/patients/:id/cac-scores", middleware.Auth(), cacHandler.GetByPatient)
api.Get("/cac-scores/:id", middleware.Auth(), cacHandler.GetByID)
api.Put("/cac-scores/:id", middleware.Auth(), cacHandler.Update)
api.Delete("/cac-scores/:id", middleware.Auth(), cacHandler.Delete)

// CCTA routes (similar pattern)
cctaRepo := repository.NewCCTARepository(db)
cctaHandler := handlers.NewCCTAHandler(cctaRepo, patientRepo)

api.Post("/patients/:id/cctas", middleware.Auth(), cctaHandler.Create)
api.Get("/patients/:id/cctas", middleware.Auth(), cctaHandler.GetByPatient)
```

---

## Phase 4: Frontend Implementation

### Step 4.1: Create React Components

**File:** `/home/user/plenya/apps/web/components/cac-score-card.tsx`

(See comprehensive documentation for full React component code)

### Step 4.2: Create Form for CAC Score Input

**File:** `/home/user/plenya/apps/web/components/forms/cac-score-form.tsx`

```typescript
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { z } from 'zod';
import { useMutation, useQueryClient } from '@tanstack/react-query';
import { apiClient } from '@/lib/api-client';

const cacScoreSchema = z.object({
  examDate: z.string().datetime(),
  institution: z.string().optional(),
  agatstonScore: z.number().int().min(0),
  percentile: z.number().int().min(0).max(100).optional(),
  cacDensity: z.enum(['low', 'mixed', 'moderate', 'high']).optional(),
  lmScore: z.number().int().min(0).optional(),
  ladScore: z.number().int().min(0).optional(),
  lcxScore: z.number().int().min(0).optional(),
  rcaScore: z.number().int().min(0).optional(),
  previousScanId: z.string().uuid().optional(),
  clinicalNotes: z.string().optional(),
});

type CACScoreFormData = z.infer<typeof cacScoreSchema>;

export function CACScoreForm({ patientId }: { patientId: string }) {
  const queryClient = useQueryClient();
  const form = useForm<CACScoreFormData>({
    resolver: zodResolver(cacScoreSchema),
  });

  const mutation = useMutation({
    mutationFn: (data: CACScoreFormData) =>
      apiClient.post(`/patients/${patientId}/cac-scores`, data),
    onSuccess: () => {
      queryClient.invalidateQueries(['cac-scores', patientId]);
      form.reset();
    },
  });

  return (
    <form onSubmit={form.handleSubmit((data) => mutation.mutate(data))}>
      {/* Form fields */}
    </form>
  );
}
```

---

## Phase 5: Testing

### Step 5.1: Unit Tests (Go)

```go
// apps/api/internal/models/cac_score_test.go

func TestCACScore_CalculateRiskLevel(t *testing.T) {
    tests := []struct {
        score    int
        expected int
    }{
        {0, 5},
        {5, 4},
        {50, 3},
        {200, 2},
        {600, 1},
        {1500, 0},
    }

    for _, tt := range tests {
        cac := &models.CACScore{AgatstonScore: tt.score}
        cac.CalculateRiskLevel()
        if cac.AgatstonRiskLevel != tt.expected {
            t.Errorf("Score %d: expected risk level %d, got %d",
                tt.score, tt.expected, cac.AgatstonRiskLevel)
        }
    }
}
```

### Step 5.2: Integration Tests

```bash
# Test API endpoints
curl -X POST http://localhost:3001/api/v1/patients/{id}/cac-scores \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer {token}" \
  -d '{
    "examDate": "2024-01-15T10:00:00Z",
    "agatstonScore": 150,
    "percentile": 75
  }'
```

---

## Phase 6: Documentation and Deployment

### Step 6.1: Update API Documentation

```bash
cd apps/api
swag init -g cmd/server/main.go
```

### Step 6.2: Deploy to Production

```bash
# Build backend
cd apps/api
go build -o bin/server cmd/server/main.go

# Build frontend
cd apps/web
pnpm build

# Deploy with Coolify (automated)
git push origin main
```

---

## Summary Checklist

- [ ] Create Go models (cac_score.go, ccta.go)
- [ ] Run `pnpm generate:migrations`
- [ ] Apply migrations with Atlas
- [ ] Run `pnpm generate` for TypeScript types
- [ ] Create repositories
- [ ] Create handlers
- [ ] Register routes
- [ ] Create React components
- [ ] Create forms
- [ ] Write tests
- [ ] Update Swagger docs
- [ ] Deploy to production

---

**Implementation Time Estimate:** 8-12 hours for complete CAC + CCTA functionality

**Priority:** HIGH (critical for cardiovascular risk assessment in functional medicine)

**Dependencies:**
- Phase 1-4 must be completed (monorepo, backend, frontend foundation)
- Patient model must exist
- Auth system operational

**Next Steps After Implementation:**
1. Integrate with advanced lipid panels (ApoB, Lp(a), LDL-P)
2. Create automated reports (PDF generation)
3. Build risk calculator dashboard
4. Implement warranty period reminders
5. Add progression tracking charts
