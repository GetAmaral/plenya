package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// PatientScoresService unifica histórico de Escore Light (autoavaliação) e
// Escore Completo (avaliado pela clínica) — paciente vê tudo numa timeline.
type PatientScoresService struct {
	db *gorm.DB
}

func NewPatientScoresService(db *gorm.DB) *PatientScoresService {
	return &PatientScoresService{db: db}
}

// ScoreEntryView é cada item da timeline.
type ScoreEntryView struct {
	ID                   uuid.UUID  `json:"id"`
	Source               string     `json:"source"` // "light" | "complete"
	CreatedAt            time.Time  `json:"createdAt"`
	TotalScorePercentage float64    `json:"totalScorePercentage"`
	ItemsEvaluatedCount  int        `json:"itemsEvaluatedCount"`
	PublicCode           *string    `json:"publicCode,omitempty"` // só Light
	GroupResults         []ScoreGroupResultView `json:"groupResults"`
}

type ScoreGroupResultView struct {
	GroupID         uuid.UUID `json:"groupId"`
	GroupName       string    `json:"groupName"`
	ScorePercentage float64   `json:"scorePercentage"`
}

// ListAll retorna todos os escores ordenados DESC.
// Inclui groupResults pra UI poder render radar/timeline por pilar sem N+1.
func (s *PatientScoresService) ListAll(patientID uuid.UUID) ([]ScoreEntryView, error) {
	out := []ScoreEntryView{}

	// Completos
	var completes []models.PatientScoreSnapshot
	err := s.db.Where("patient_id = ?", patientID).
		Preload("GroupResults.Group").
		Order("created_at DESC").
		Find(&completes).Error
	if err != nil {
		return nil, err
	}
	for i := range completes {
		c := &completes[i]
		entry := ScoreEntryView{
			ID:                   c.ID,
			Source:               "complete",
			CreatedAt:            c.CreatedAt,
			TotalScorePercentage: c.TotalScorePercentage,
			ItemsEvaluatedCount:  c.ItemsEvaluatedCount,
			GroupResults:         []ScoreGroupResultView{},
		}
		for j := range c.GroupResults {
			gr := &c.GroupResults[j]
			view := ScoreGroupResultView{
				GroupID:         gr.GroupID,
				ScorePercentage: gr.ScorePercentage,
			}
			if gr.Group != nil {
				view.GroupName = gr.Group.Name
			}
			entry.GroupResults = append(entry.GroupResults, view)
		}
		out = append(out, entry)
	}

	// Sessões já importadas pro prontuário (viram snapshot "complete") não devem aparecer
	// TAMBÉM como "light" — evita o mesmo escore duplicado no histórico.
	importedSessions := make(map[uuid.UUID]bool)
	for i := range completes {
		if completes[i].SourceSessionID != nil {
			importedSessions[*completes[i].SourceSessionID] = true
		}
	}

	// Light (autoavaliações claimed pelo paciente)
	var lights []models.AnonymousScoreSession
	err = s.db.
		Preload("Snapshot.GroupResults.Group").
		Where("claimed_by_patient_id = ?", patientID).
		Order("created_at DESC").
		Find(&lights).Error
	if err != nil {
		return nil, err
	}
	for i := range lights {
		sess := &lights[i]
		if sess.Snapshot == nil || importedSessions[sess.ID] {
			continue
		}
		entry := ScoreEntryView{
			ID:                   sess.Snapshot.ID,
			Source:               "light",
			CreatedAt:            sess.Snapshot.CreatedAt,
			TotalScorePercentage: sess.Snapshot.TotalScorePercentage,
			ItemsEvaluatedCount:  sess.Snapshot.ItemsEvaluatedCount,
			GroupResults:         []ScoreGroupResultView{},
		}
		code := sess.PublicCode
		entry.PublicCode = &code
		for j := range sess.Snapshot.GroupResults {
			gr := &sess.Snapshot.GroupResults[j]
			view := ScoreGroupResultView{
				GroupID:         gr.GroupID,
				ScorePercentage: gr.ScorePercentage,
			}
			if gr.Group != nil {
				view.GroupName = gr.Group.Name
			}
			entry.GroupResults = append(entry.GroupResults, view)
		}
		out = append(out, entry)
	}

	// Sort merged DESC por createdAt
	sortByCreatedAtDesc(out)
	return out, nil
}

func sortByCreatedAtDesc(rows []ScoreEntryView) {
	// Insertion sort — N pequeno (10s, max 100s)
	for i := 1; i < len(rows); i++ {
		for j := i; j > 0 && rows[j].CreatedAt.After(rows[j-1].CreatedAt); j-- {
			rows[j], rows[j-1] = rows[j-1], rows[j]
		}
	}
}

// GetCompleteSnapshot retorna detalhe de um escore completo (com itens).
// Filtra por patientID — paciente nunca vê snapshots de outros.
func (s *PatientScoresService) GetCompleteSnapshot(patientID, snapshotID uuid.UUID) (*models.PatientScoreSnapshot, error) {
	var snap models.PatientScoreSnapshot
	err := s.db.
		Where("id = ? AND patient_id = ?", snapshotID, patientID).
		Preload("GroupResults.Group").
		First(&snap).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("escore não encontrado")
		}
		return nil, err
	}
	return &snap, nil
}

// GetLatestCompleteSnapshot retorna o escore COMPLETO mais recente do paciente,
// com itens + pilares AGIR pré-carregados, para o radar do portal usar a MESMA
// fonte (buildAgir) do EMR staff. Retorna (nil, nil) se o paciente não tiver
// nenhum escore completo (pode ter só Light, que é por grupo e sem pilares).
func (s *PatientScoresService) GetLatestCompleteSnapshot(patientID uuid.UUID) (*models.PatientScoreSnapshot, error) {
	var snap models.PatientScoreSnapshot
	err := s.db.
		Where("patient_id = ?", patientID).
		Preload("GroupResults", func(db *gorm.DB) *gorm.DB {
			return db.Order("score_percentage DESC")
		}).
		Preload("GroupResults.Group", "deleted_at IS NULL").
		Preload("ItemResults", func(db *gorm.DB) *gorm.DB {
			return db.Order("actual_points DESC")
		}).
		Preload("ItemResults.Item", "deleted_at IS NULL").
		Preload("ItemResults.Item.MethodPillars", func(db *gorm.DB) *gorm.DB {
			return db.Where("deleted_at IS NULL").Order("\"order\" ASC")
		}).
		Preload("ItemResults.Item.MethodPillars.Letter", "deleted_at IS NULL").
		Preload("ItemResults.Item.MethodPillars.Letter.Method", "deleted_at IS NULL").
		Order("calculated_at DESC").
		First(&snap).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &snap, nil
}
