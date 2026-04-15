package services

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

type AssessmentHTMLService struct {
	db *gorm.DB
}

func NewAssessmentHTMLService(db *gorm.DB) *AssessmentHTMLService {
	return &AssessmentHTMLService{db: db}
}

// GenerateHTML generates an HTML report for a physical assessment and saves it.
func (s *AssessmentHTMLService) GenerateHTML(assessmentID uuid.UUID) (string, error) {
	var assessment models.PhysicalAssessment
	if err := s.db.First(&assessment, assessmentID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", ErrPhysicalAssessmentNotFound
		}
		return "", err
	}

	var patient models.Patient
	if err := s.db.First(&patient, assessment.PatientID).Error; err != nil {
		return "", err
	}
	patient.CalculateAge()

	html := s.buildHTML(&assessment, &patient)

	// Save to assessment
	assessment.HtmlContent = &html
	if err := s.db.Model(&assessment).Update("html_content", html).Error; err != nil {
		return "", err
	}

	return html, nil
}

func (s *AssessmentHTMLService) buildHTML(a *models.PhysicalAssessment, p *models.Patient) string {
	var b strings.Builder

	// Helper to format float pointer
	ff := func(v *float64, decimals int) string {
		if v == nil {
			return "-"
		}
		return fmt.Sprintf("%.*f", decimals, *v)
	}
	// Helper to format int pointer
	fi := func(v *int) string {
		if v == nil {
			return "-"
		}
		return fmt.Sprintf("%d", *v)
	}

	// Risk level styling
	riskColor := "#4caf50"
	riskLabel := "Baixo Risco"
	if a.ACSMRiskLevel != nil {
		switch *a.ACSMRiskLevel {
		case models.ACSMRiskModerate:
			riskColor = "#ff9800"
			riskLabel = "Risco Moderado"
		case models.ACSMRiskHigh:
			riskColor = "#f44336"
			riskLabel = "Alto Risco"
		}
	}

	// BMI category
	bmiCategory := ""
	bmiCategoryIndex := -1
	if a.BMI != nil {
		bmi := *a.BMI
		switch {
		case bmi < 18.5:
			bmiCategory = "Baixo Peso"
			bmiCategoryIndex = 0
		case bmi < 25:
			bmiCategory = "Normal"
			bmiCategoryIndex = 1
		case bmi < 30:
			bmiCategory = "Sobrepeso"
			bmiCategoryIndex = 2
		case bmi < 35:
			bmiCategory = "Obesidade I"
			bmiCategoryIndex = 3
		case bmi < 40:
			bmiCategory = "Obesidade II"
			bmiCategoryIndex = 4
		default:
			bmiCategory = "Obesidade III"
			bmiCategoryIndex = 5
		}
	}

	// HR Zones
	maxHR := CalculateMaxHR(p.Age)
	restHR := 70
	if a.RestingHeartRate != nil {
		restHR = *a.RestingHeartRate
	}
	zones := CalculateKarvonenZones(maxHR, restHR)

	// --- Build HTML document ---

	b.WriteString(`<!DOCTYPE html><html lang="pt-BR"><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width, initial-scale=1.0"><title>Avaliacao Fisica</title>`)

	// CSS styles
	b.WriteString(`<style>
* { margin: 0; padding: 0; box-sizing: border-box; }
body { font-family: system-ui, -apple-system, sans-serif; color: #1f2937; background: #f9fafb; padding: 2rem; }
.container { max-width: 800px; margin: 0 auto; }
.header { text-align: center; margin-bottom: 2rem; border-bottom: 3px solid #1a56db; padding-bottom: 1rem; }
.header h1 { color: #1a56db; font-size: 1.5rem; }
.header .patient-name { font-size: 1.2rem; margin-top: 0.5rem; }
.header .date { color: #6b7280; font-size: 0.9rem; }
.card { background: white; border: 1px solid #e5e7eb; border-radius: 8px; padding: 1.5rem; margin-bottom: 1rem; box-shadow: 0 1px 3px rgba(0,0,0,0.1); }
.card h2 { font-size: 1rem; color: #374151; margin-bottom: 1rem; border-bottom: 1px solid #e5e7eb; padding-bottom: 0.5rem; }
.metrics { display: grid; grid-template-columns: repeat(auto-fit, minmax(150px, 1fr)); gap: 1rem; }
.metric { text-align: center; }
.metric .value { font-size: 1.5rem; font-weight: 700; color: #1a56db; }
.metric .label { font-size: 0.75rem; color: #6b7280; text-transform: uppercase; }
.metric .unit { font-size: 0.8rem; color: #9ca3af; }
.risk-badge { display: inline-block; padding: 0.5rem 1.5rem; border-radius: 20px; color: white; font-weight: 700; font-size: 1.1rem; }
.risk-section { text-align: center; margin-bottom: 1rem; }
.factor-list { display: flex; flex-wrap: wrap; gap: 0.5rem; margin-top: 0.5rem; }
.factor { padding: 0.25rem 0.75rem; border-radius: 12px; font-size: 0.8rem; }
.factor-risk { background: #fef2f2; color: #991b1b; border: 1px solid #fecaca; }
.factor-protective { background: #f0fdf4; color: #166534; border: 1px solid #bbf7d0; }
.bmi-table { width: 100%%; border-collapse: collapse; margin-top: 0.5rem; }
.bmi-table th, .bmi-table td { padding: 0.5rem; text-align: center; border: 1px solid #e5e7eb; font-size: 0.85rem; }
.bmi-table .highlight { background: #dbeafe; font-weight: 700; }
.zone-row { display: flex; align-items: center; padding: 0.5rem; border-radius: 4px; margin-bottom: 0.25rem; }
.zone-row:nth-child(odd) { background: #f9fafb; }
.zone-badge { width: 40px; text-align: center; font-weight: 700; font-size: 0.8rem; padding: 0.25rem; border-radius: 4px; color: white; margin-right: 0.75rem; }
.zone-name { flex: 1; font-size: 0.85rem; }
.zone-bpm { font-family: monospace; font-size: 0.85rem; }
.tags { display: flex; flex-wrap: wrap; gap: 0.5rem; }
.tag { padding: 0.25rem 0.75rem; border-radius: 12px; font-size: 0.8rem; background: #f3f4f6; border: 1px solid #d1d5db; }
.footer { text-align: center; margin-top: 2rem; padding-top: 1rem; border-top: 1px solid #e5e7eb; font-size: 0.75rem; color: #9ca3af; }
@media print {
	body { padding: 0; background: white; }
	.card { box-shadow: none; break-inside: avoid; }
}
</style>`)

	b.WriteString(`</head><body><div class="container">`)

	// Header
	b.WriteString(fmt.Sprintf(`<div class="header">
<h1>AVALIACAO FISICA</h1>
<div class="patient-name">%s</div>
<div class="date">%s | %d anos | %s</div>
</div>`, p.Name, a.AssessmentDate.Format("02/01/2006"), p.Age, genderLabel(p.Gender)))

	// ACSM Risk Card
	b.WriteString(fmt.Sprintf(`<div class="card">
<div class="risk-section">
<span class="risk-badge" style="background:%s">%s</span>
<p style="margin-top:0.5rem;color:#6b7280">%d fator(es) de risco cardiovascular</p>
</div>`, riskColor, riskLabel, a.ACSMRiskFactorsCount))

	if len(a.ACSMRiskFactors) > 0 {
		b.WriteString(`<div><strong style="font-size:0.85rem">Fatores de Risco:</strong><div class="factor-list">`)
		for _, f := range a.ACSMRiskFactors {
			b.WriteString(fmt.Sprintf(`<span class="factor factor-risk">%s</span>`, f))
		}
		b.WriteString(`</div></div>`)
	}
	if len(a.ACSMProtectiveFactors) > 0 {
		b.WriteString(`<div style="margin-top:0.5rem"><strong style="font-size:0.85rem">Fatores Protetores:</strong><div class="factor-list">`)
		for _, f := range a.ACSMProtectiveFactors {
			b.WriteString(fmt.Sprintf(`<span class="factor factor-protective">%s</span>`, f))
		}
		b.WriteString(`</div></div>`)
	}
	b.WriteString(`</div>`)

	// Anthropometry
	b.WriteString(`<div class="card"><h2>Antropometria</h2><div class="metrics">`)
	writeMetric(&b, "Peso", ff(a.Weight, 1), "kg")
	writeMetric(&b, "Altura", ff(a.Height, 1), "cm")
	writeMetric(&b, "Cintura", ff(a.WaistCircumference, 1), "cm")
	writeMetric(&b, "IMC", ff(a.BMI, 1), "kg/m2")
	writeMetric(&b, "BRI", ff(a.BRI, 1), "")
	b.WriteString(`</div></div>`)

	// BMI Table
	if a.BMI != nil {
		bmiCategories := []struct{ name, rangeStr string }{
			{"Baixo Peso", "<18.5"},
			{"Normal", "18.5-24.9"},
			{"Sobrepeso", "25-29.9"},
			{"Obesidade I", "30-34.9"},
			{"Obesidade II", "35-39.9"},
			{"Obesidade III", ">=40"},
		}
		b.WriteString(`<div class="card"><h2>Classificacao IMC</h2>`)
		b.WriteString(fmt.Sprintf(`<p style="text-align:center;margin-bottom:0.5rem">IMC: <strong>%s</strong> - <strong>%s</strong></p>`, ff(a.BMI, 1), bmiCategory))
		b.WriteString(`<table class="bmi-table"><tr><th>Classificacao</th><th>Faixa</th></tr>`)
		for i, cat := range bmiCategories {
			cls := ""
			if i == bmiCategoryIndex {
				cls = ` class="highlight"`
			}
			b.WriteString(fmt.Sprintf(`<tr%s><td>%s</td><td>%s</td></tr>`, cls, cat.name, cat.rangeStr))
		}
		b.WriteString(`</table></div>`)
	}

	// Body Composition
	if a.BodyFatPercent != nil || a.LeanMass != nil {
		b.WriteString(`<div class="card"><h2>Composicao Corporal</h2><div class="metrics">`)
		writeMetric(&b, "% Gordura", ff(a.BodyFatPercent, 1), "%")
		writeMetric(&b, "Massa Magra", ff(a.LeanMass, 1), "kg")
		if a.BodyFatPercent != nil && a.Weight != nil {
			fatMass := *a.Weight * *a.BodyFatPercent / 100
			writeMetric(&b, "Massa Gorda", fmt.Sprintf("%.1f", fatMass), "kg")
		}
		b.WriteString(`</div></div>`)
	}

	// Cardiovascular
	b.WriteString(`<div class="card"><h2>Cardiovascular</h2><div class="metrics">`)
	if a.SystolicBP != nil && a.DiastolicBP != nil {
		writeMetric(&b, "PA", fmt.Sprintf("%d/%d", *a.SystolicBP, *a.DiastolicBP), "mmHg")
	}
	writeMetric(&b, "FC Repouso", fi(a.RestingHeartRate), "bpm")
	writeMetric(&b, "FC Maxima", fmt.Sprintf("%d", maxHR), "bpm")
	b.WriteString(`</div></div>`)

	// HR Zones
	zoneColors := []string{"#22c55e", "#84cc16", "#eab308", "#f97316", "#ef4444"}
	b.WriteString(`<div class="card"><h2>Zonas de Frequencia Cardiaca (Karvonen)</h2>`)
	for i, z := range zones {
		color := zoneColors[i]
		b.WriteString(fmt.Sprintf(`<div class="zone-row">
<span class="zone-badge" style="background:%s">Z%d</span>
<span class="zone-name">%s (%s)</span>
<span class="zone-bpm">%d-%d bpm</span>
</div>`, color, z.Zone, z.Name, z.Percent, z.MinBPM, z.MaxBPM))
	}
	b.WriteString(`</div>`)

	// Lab Results
	if a.LDL != nil || a.HDL != nil || a.TotalCholesterol != nil || a.FastingGlucose != nil {
		b.WriteString(`<div class="card"><h2>Laboratorial</h2><div class="metrics">`)
		writeMetric(&b, "LDL", ff(a.LDL, 0), "mg/dL")
		writeMetric(&b, "HDL", ff(a.HDL, 0), "mg/dL")
		writeMetric(&b, "Col. Total", ff(a.TotalCholesterol, 0), "mg/dL")
		writeMetric(&b, "Triglicerideos", ff(a.Triglycerides, 0), "mg/dL")
		writeMetric(&b, "Glicemia", ff(a.FastingGlucose, 0), "mg/dL")
		writeMetric(&b, "HbA1c", ff(a.HbA1c, 1), "%")
		b.WriteString(`</div></div>`)
	}

	// ACSM Tags
	if len(a.ACSMTags) > 0 {
		b.WriteString(`<div class="card"><h2>Tags ACSM</h2><div class="tags">`)
		for _, tag := range a.ACSMTags {
			b.WriteString(fmt.Sprintf(`<span class="tag">%s</span>`, tag))
		}
		b.WriteString(`</div></div>`)
	}

	// Recommendation
	if a.ACSMRecommendation != nil {
		b.WriteString(fmt.Sprintf(`<div class="card"><h2>Recomendacao</h2><p style="font-size:0.9rem">%s</p></div>`, *a.ACSMRecommendation))
	}

	// Footer
	b.WriteString(`<div class="footer">
<p>Gerado pelo sistema Plenya</p>
<p>Referencias: ACSM's Guidelines for Exercise Testing and Prescription, 11th ed. | Pollock &amp; Wilmore, 1993</p>
</div>`)

	b.WriteString(`</div></body></html>`)

	return b.String()
}

func writeMetric(b *strings.Builder, label, value, unit string) {
	b.WriteString(fmt.Sprintf(`<div class="metric"><div class="value">%s<span class="unit"> %s</span></div><div class="label">%s</div></div>`, value, unit, label))
}

func genderLabel(g models.Gender) string {
	switch g {
	case models.GenderMale:
		return "Masculino"
	case models.GenderFemale:
		return "Feminino"
	default:
		return ""
	}
}
