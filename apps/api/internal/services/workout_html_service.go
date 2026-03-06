package services

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

type WorkoutHtmlService struct {
	db         *gorm.DB
	uploadsDir string
}

func NewWorkoutHtmlService(db *gorm.DB, uploadsDir string) *WorkoutHtmlService {
	return &WorkoutHtmlService{db: db, uploadsDir: uploadsDir}
}

// GenerateHTML gera HTML completo para um plano de treino com GIFs base64 inline
func (s *WorkoutHtmlService) GenerateHTML(planID uuid.UUID) (string, error) {
	var plan models.WorkoutPlan
	if err := s.db.Where("id = ?", planID).
		Preload("Patient").
		Preload("CreatedBy").
		Preload("Sessions", func(db *gorm.DB) *gorm.DB { return db.Order("\"order\" ASC") }).
		Preload("Sessions.Exercises", func(db *gorm.DB) *gorm.DB { return db.Order("phase ASC, \"order\" ASC") }).
		Preload("Sessions.Exercises.Exercise").
		First(&plan).Error; err != nil {
		return "", fmt.Errorf("plano não encontrado: %w", err)
	}

	// Build template data
	data := htmlTemplateData{
		PlanName:        plan.Name,
		PatientName:     plan.Patient.Name,
		Objective:       objectiveLabel(plan.Objective),
		Intensity:       intensityLabel(plan.Intensity),
		WeeklyFrequency: plan.WeeklyFrequency,
		DurationMinutes: plan.DurationMinutes,
		CreatedByName:   plan.CreatedBy.Name,
		Date:            time.Now().Format("02/01/2006"),
		Sessions:        []htmlSessionData{},
	}

	for _, sess := range plan.Sessions {
		sessData := htmlSessionData{
			Name:     sess.Name,
			Warmup:   []htmlExerciseData{},
			Main:     []htmlExerciseData{},
			Cooldown: []htmlExerciseData{},
		}

		for _, ex := range sess.Exercises {
			exData := htmlExerciseData{
				Name:    ex.Exercise.NamePt,
				Sets:    ex.Sets,
				Reps:    ex.Reps,
				Cadence: cadenceLabel(ex.Cadence),
				RestSec: ex.RestBetweenSetsSec,
				Notes:   ptrStr(ex.Notes),
			}
			if exData.Name == "" {
				exData.Name = ex.Exercise.Name
			}

			// Try to load GIF and encode as base64
			exData.GifBase64 = s.loadGifBase64(ex.Exercise.ExternalId)

			switch ex.Phase {
			case models.ExercisePhaseWarmup:
				sessData.Warmup = append(sessData.Warmup, exData)
			case models.ExercisePhaseMain:
				sessData.Main = append(sessData.Main, exData)
			case models.ExercisePhaseCooldown:
				sessData.Cooldown = append(sessData.Cooldown, exData)
			}
		}

		data.Sessions = append(data.Sessions, sessData)
	}

	// Render template
	tmpl, err := template.New("workout").Parse(workoutHTMLTemplate)
	if err != nil {
		return "", fmt.Errorf("erro no template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("erro ao renderizar: %w", err)
	}

	html := buf.String()

	// Save to plan
	s.db.Model(&models.WorkoutPlan{}).Where("id = ?", planID).Update("html_content", html)

	return html, nil
}

func (s *WorkoutHtmlService) loadGifBase64(externalID string) string {
	gifPath := filepath.Join(s.uploadsDir, "exercises", externalID+".gif")
	data, err := os.ReadFile(gifPath)
	if err != nil {
		return ""
	}
	return "data:image/gif;base64," + base64.StdEncoding.EncodeToString(data)
}

// Template data types

type htmlTemplateData struct {
	PlanName        string
	PatientName     string
	Objective       string
	Intensity       string
	WeeklyFrequency int
	DurationMinutes int
	CreatedByName   string
	Date            string
	Sessions        []htmlSessionData
}

type htmlSessionData struct {
	Name     string
	Warmup   []htmlExerciseData
	Main     []htmlExerciseData
	Cooldown []htmlExerciseData
}

type htmlExerciseData struct {
	Name      string
	Sets      int
	Reps      string
	Cadence   string
	RestSec   int
	Notes     string
	GifBase64 string
}

func objectiveLabel(o models.WorkoutObjective) string {
	m := map[models.WorkoutObjective]string{
		models.WorkoutObjectiveHypertrophy:    "Hipertrofia",
		models.WorkoutObjectiveStrength:       "Força",
		models.WorkoutObjectiveEndurance:      "Resistência",
		models.WorkoutObjectiveWeightLoss:     "Emagrecimento",
		models.WorkoutObjectiveConditioning:   "Condicionamento",
		models.WorkoutObjectiveRehabilitation: "Reabilitação",
	}
	if l, ok := m[o]; ok {
		return l
	}
	return string(o)
}

func intensityLabel(i models.WorkoutIntensity) string {
	m := map[models.WorkoutIntensity]string{
		models.WorkoutIntensityVeryLight: "Muito Leve",
		models.WorkoutIntensityLight:     "Leve",
		models.WorkoutIntensityModerate:  "Moderado",
		models.WorkoutIntensityHigh:      "Alto",
		models.WorkoutIntensityVeryHigh:  "Muito Alto",
	}
	if l, ok := m[i]; ok {
		return l
	}
	return string(i)
}

func cadenceLabel(c models.ExerciseCadence) string {
	m := map[models.ExerciseCadence]string{
		models.ExerciseCadenceNormal:    "Normal",
		models.ExerciseCadenceSlow:      "Lento",
		models.ExerciseCadencePaused:    "Pausado",
		models.ExerciseCadenceExplosive: "Explosivo",
		models.ExerciseCadenceFree:      "Livre",
	}
	if l, ok := m[c]; ok {
		return l
	}
	return string(c)
}

func ptrStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// HTML Template (replicates exportador_treino.py style)
const workoutHTMLTemplate = `<!DOCTYPE html>
<html lang="pt-BR">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>{{.PlanName}} - {{.PatientName}}</title>
<style>
* { margin: 0; padding: 0; box-sizing: border-box; }
body { font-family: 'Segoe UI', Tahoma, sans-serif; background: #f5f5f5; color: #333; }
.container { max-width: 800px; margin: 0 auto; padding: 20px; }
.header { background: #2c2c2c; color: white; padding: 24px; border-radius: 12px; margin-bottom: 24px; text-align: center; }
.header h1 { font-size: 24px; color: #FFC107; margin-bottom: 8px; }
.header .patient { font-size: 18px; opacity: 0.9; }
.header .meta { font-size: 14px; opacity: 0.7; margin-top: 8px; }
.info-grid { display: grid; grid-template-columns: repeat(4, 1fr); gap: 12px; margin-bottom: 24px; }
.info-card { background: white; padding: 16px; border-radius: 8px; text-align: center; box-shadow: 0 1px 3px rgba(0,0,0,0.1); }
.info-card .label { font-size: 11px; text-transform: uppercase; color: #888; letter-spacing: 1px; }
.info-card .value { font-size: 18px; font-weight: bold; color: #2c2c2c; margin-top: 4px; }
.session { background: white; border-radius: 12px; margin-bottom: 24px; overflow: hidden; box-shadow: 0 2px 8px rgba(0,0,0,0.08); }
.session-header { background: #FFC107; color: #2c2c2c; padding: 16px 20px; font-size: 18px; font-weight: bold; }
.phase { padding: 16px 20px; }
.phase-title { font-size: 14px; font-weight: bold; text-transform: uppercase; color: #FFC107; border-bottom: 2px solid #FFC107; padding-bottom: 8px; margin-bottom: 12px; }
.exercise { display: flex; align-items: center; gap: 16px; padding: 12px 0; border-bottom: 1px solid #eee; }
.exercise:last-child { border-bottom: none; }
.exercise img { width: 80px; height: 80px; border-radius: 8px; object-fit: cover; background: #f0f0f0; }
.exercise .no-gif { width: 80px; height: 80px; border-radius: 8px; background: #e0e0e0; display: flex; align-items: center; justify-content: center; color: #999; font-size: 12px; }
.exercise-info { flex: 1; }
.exercise-name { font-weight: 600; font-size: 15px; margin-bottom: 4px; }
.exercise-details { display: flex; gap: 12px; flex-wrap: wrap; }
.exercise-detail { background: #f5f5f5; padding: 3px 10px; border-radius: 4px; font-size: 13px; color: #555; }
.exercise-notes { font-size: 12px; color: #888; font-style: italic; margin-top: 4px; }
.footer { text-align: center; padding: 20px; color: #888; font-size: 12px; border-top: 1px solid #ddd; }
.footer .professional { font-weight: 600; color: #555; }
@media print {
  body { background: white; }
  .container { padding: 0; max-width: 100%; }
  .session { break-inside: avoid; }
}
@media (max-width: 600px) {
  .info-grid { grid-template-columns: repeat(2, 1fr); }
  .exercise img, .exercise .no-gif { width: 60px; height: 60px; }
}
</style>
</head>
<body>
<div class="container">
  <div class="header">
    <h1>{{.PlanName}}</h1>
    <div class="patient">{{.PatientName}}</div>
    <div class="meta">Gerado em {{.Date}}</div>
  </div>

  <div class="info-grid">
    <div class="info-card">
      <div class="label">Objetivo</div>
      <div class="value">{{.Objective}}</div>
    </div>
    <div class="info-card">
      <div class="label">Intensidade</div>
      <div class="value">{{.Intensity}}</div>
    </div>
    <div class="info-card">
      <div class="label">Frequência</div>
      <div class="value">{{.WeeklyFrequency}}x/sem</div>
    </div>
    <div class="info-card">
      <div class="label">Duração</div>
      <div class="value">{{.DurationMinutes}} min</div>
    </div>
  </div>

  {{range .Sessions}}
  <div class="session">
    <div class="session-header">{{.Name}}</div>

    {{if .Warmup}}
    <div class="phase">
      <div class="phase-title">🔥 Aquecimento</div>
      {{range .Warmup}}
      <div class="exercise">
        {{if .GifBase64}}<img src="{{.GifBase64}}" alt="{{.Name}}">{{else}}<div class="no-gif">GIF</div>{{end}}
        <div class="exercise-info">
          <div class="exercise-name">{{.Name}}</div>
          <div class="exercise-details">
            <span class="exercise-detail">{{.Sets}}x{{.Reps}}</span>
            <span class="exercise-detail">{{.Cadence}}</span>
            <span class="exercise-detail">Desc: {{.RestSec}}s</span>
          </div>
          {{if .Notes}}<div class="exercise-notes">{{.Notes}}</div>{{end}}
        </div>
      </div>
      {{end}}
    </div>
    {{end}}

    {{if .Main}}
    <div class="phase">
      <div class="phase-title">💪 Principal</div>
      {{range .Main}}
      <div class="exercise">
        {{if .GifBase64}}<img src="{{.GifBase64}}" alt="{{.Name}}">{{else}}<div class="no-gif">GIF</div>{{end}}
        <div class="exercise-info">
          <div class="exercise-name">{{.Name}}</div>
          <div class="exercise-details">
            <span class="exercise-detail">{{.Sets}}x{{.Reps}}</span>
            <span class="exercise-detail">{{.Cadence}}</span>
            <span class="exercise-detail">Desc: {{.RestSec}}s</span>
          </div>
          {{if .Notes}}<div class="exercise-notes">{{.Notes}}</div>{{end}}
        </div>
      </div>
      {{end}}
    </div>
    {{end}}

    {{if .Cooldown}}
    <div class="phase">
      <div class="phase-title">🧘 Finalização</div>
      {{range .Cooldown}}
      <div class="exercise">
        {{if .GifBase64}}<img src="{{.GifBase64}}" alt="{{.Name}}">{{else}}<div class="no-gif">GIF</div>{{end}}
        <div class="exercise-info">
          <div class="exercise-name">{{.Name}}</div>
          <div class="exercise-details">
            <span class="exercise-detail">{{.Sets}}x{{.Reps}}</span>
            <span class="exercise-detail">{{.Cadence}}</span>
            <span class="exercise-detail">Desc: {{.RestSec}}s</span>
          </div>
          {{if .Notes}}<div class="exercise-notes">{{.Notes}}</div>{{end}}
        </div>
      </div>
      {{end}}
    </div>
    {{end}}

  </div>
  {{end}}

  <div class="footer">
    <div class="professional">Profissional: {{.CreatedByName}}</div>
    <p>Este programa de exercícios foi elaborado de forma individualizada. Não compartilhe sem autorização.</p>
    <p style="margin-top:8px">Gerado por Plenya</p>
  </div>
</div>
</body>
</html>`
