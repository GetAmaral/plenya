// Package services — MedicalRecordAggregateService faz query UNION sobre 8
// modelos de "evento" do paciente (anamneses, exames, scores, avaliações,
// prescrições, consultas) pra alimentar a tab /patients/:id/prontuario.
//
// Filtros: tipo + autor (user que criou). Paginação obrigatória.
package services

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MedicalRecordAggregateService struct {
	db *gorm.DB
}

func NewMedicalRecordAggregateService(db *gorm.DB) *MedicalRecordAggregateService {
	return &MedicalRecordAggregateService{db: db}
}

// MedicalRecordEntry é a projeção uniforme de um "evento" do paciente.
//
// Campos opcionais (Status, Specialty) só vêm preenchidos quando relevantes
// pra aquele tipo. Frontend usa Type pra decidir como renderizar/linkar.
type MedicalRecordEntry struct {
	ID          uuid.UUID `json:"id"`
	Type        string    `json:"type"` // anamnesis|lab_batch|score_snapshot|physical_assessment|fitness_test|postural_assessment|prescription|appointment
	EventDate   time.Time `json:"eventDate"`
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle,omitempty"`
	Status      string    `json:"status,omitempty"`
	HasNote     bool      `json:"hasNote"` // só relevante p/ type=appointment: existe ClinicalNote vinculada
	AuthorID    uuid.UUID `json:"authorId"`
	AuthorName  string    `json:"authorName,omitempty"`
	AuthorRoles string    `json:"authorRoles,omitempty"` // CSV pra filtro por especialidade no front
	PatientID   uuid.UUID `json:"patientId"`
	CreatedAt   time.Time `json:"createdAt"`
}

// AggregateFilter — filtros pro endpoint.
type AggregateFilter struct {
	Types     []string  // se vazio, traz todos os 8
	AuthorIDs []uuid.UUID
	From      *time.Time
	To        *time.Time
	Limit     int
	Offset    int
}

// allTypes é a lista canônica suportada. Usada quando filter.Types é vazio.
var allTypes = []string{
	"anamnesis",
	"lab_batch",
	"score_snapshot",
	"physical_assessment",
	"fitness_test",
	"postural_assessment",
	"prescription",
	"appointment",
}

// List monta UNION ALL das 8 tabelas filtradas por patient_id, depois agrega
// users.name/roles via outer join (pra autor humano-amigável + filtro por
// role no frontend). Paginação aplicada no result combinado.
//
// Performance: cada subquery filtra por patient_id (índice). UNION ALL é
// barato. JOIN final em users também usa PK. Pra paciente típico (centenas
// de eventos / anos), latência <50ms em cargas razoáveis.
func (s *MedicalRecordAggregateService) List(patientID uuid.UUID, filter AggregateFilter) ([]MedicalRecordEntry, error) {
	if filter.Limit <= 0 || filter.Limit > 200 {
		filter.Limit = 50
	}

	types := filter.Types
	if len(types) == 0 {
		types = allTypes
	}
	typeSet := make(map[string]bool, len(types))
	for _, t := range types {
		typeSet[t] = true
	}

	// Cada subquery projeta 10 colunas na MESMA ordem (a última, has_note, é só
	// relevante para appointment; nas demais é FALSE). O alias AS has_note vai em
	// TODAS as partes porque o nome de coluna do UNION vem da primeira incluída —
	// e qualquer tipo pode ser o primeiro dependendo do filtro.
	parts := []string{}
	if typeSet["anamnesis"] {
		parts = append(parts, `
SELECT id, 'anamnesis'::text AS type, consultation_date AS event_date,
       COALESCE(NULLIF(notes,''), 'Anamnese') AS title,
       ''::text AS subtitle, ''::text AS status,
       author_id AS author_id, patient_id, created_at, FALSE AS has_note
  FROM anamnesis WHERE patient_id = ? AND deleted_at IS NULL`)
	}
	if typeSet["lab_batch"] {
		parts = append(parts, `
SELECT id, 'lab_batch'::text, COALESCE(result_date, collection_date),
       laboratory_name,
       COALESCE(observations,'')::text,
       status::text,
       COALESCE(requesting_doctor_id, '00000000-0000-0000-0000-000000000000'::uuid),
       patient_id, created_at, FALSE AS has_note
  FROM lab_result_batches WHERE patient_id = ? AND deleted_at IS NULL`)
	}
	if typeSet["score_snapshot"] {
		parts = append(parts, `
SELECT id, 'score_snapshot'::text, calculated_at,
       'Escore Plenya — ' || ROUND(total_score_percentage::numeric, 1) || '%',
       'Itens avaliados: ' || items_evaluated_count,
       ''::text,
       calculated_by_user_id, patient_id, created_at, FALSE AS has_note
  FROM patient_score_snapshots WHERE patient_id = ? AND deleted_at IS NULL`)
	}
	if typeSet["physical_assessment"] {
		parts = append(parts, `
SELECT id, 'physical_assessment'::text, assessment_date::timestamp,
       'Avaliação Física',
       COALESCE(acsm_risk_level::text, ''),
       ''::text,
       created_by_id, patient_id, created_at, FALSE AS has_note
  FROM physical_assessments WHERE patient_id = ? AND deleted_at IS NULL`)
	}
	if typeSet["fitness_test"] {
		parts = append(parts, `
SELECT id, 'fitness_test'::text, assessment_date::timestamp,
       'Teste de Fitness — ' || overall_classification,
       'Pontos: ' || overall_score,
       ''::text,
       created_by_id, patient_id, created_at, FALSE AS has_note
  FROM fitness_test_results WHERE patient_id = ? AND deleted_at IS NULL`)
	}
	if typeSet["postural_assessment"] {
		parts = append(parts, `
SELECT id, 'postural_assessment'::text, assessment_date::timestamp,
       'Avaliação Postural — ' || postural_classification,
       'Pontos: ' || postural_score,
       ''::text,
       created_by_id, patient_id, created_at, FALSE AS has_note
  FROM postural_assessments WHERE patient_id = ? AND deleted_at IS NULL`)
	}
	if typeSet["prescription"] {
		parts = append(parts, `
SELECT id, 'prescription'::text, prescription_date,
       'Prescrição',
       COALESCE(general_instructions,''),
       status::text,
       doctor_id, patient_id, created_at, FALSE AS has_note
  FROM prescriptions WHERE patient_id = ? AND deleted_at IS NULL`)
	}
	if typeSet["appointment"] {
		parts = append(parts, `
SELECT id, 'appointment'::text, scheduled_at,
       reason,
       type::text,
       status::text,
       doctor_id, patient_id, created_at,
       EXISTS(SELECT 1 FROM clinical_notes cn WHERE cn.appointment_id = appointments.id AND cn.deleted_at IS NULL) AS has_note
  FROM appointments WHERE patient_id = ? AND deleted_at IS NULL`)
	}

	if len(parts) == 0 {
		return []MedicalRecordEntry{}, nil
	}

	// Args = patient_id N vezes (uma por subquery), depois limit/offset.
	args := make([]any, 0, len(parts)+2)
	for range parts {
		args = append(args, patientID)
	}

	// Wrap UNION ALL como CTE, faz join com users pra trazer nome+roles, filtra
	// por author/data, ordena, pagina.
	authorClause := ""
	if len(filter.AuthorIDs) > 0 {
		placeholders := make([]string, len(filter.AuthorIDs))
		for i, a := range filter.AuthorIDs {
			placeholders[i] = "?"
			args = append(args, a)
		}
		authorClause = " AND e.author_id IN (" + strings.Join(placeholders, ",") + ")"
	}
	dateClause := ""
	if filter.From != nil {
		dateClause += " AND e.event_date >= ?"
		args = append(args, *filter.From)
	}
	if filter.To != nil {
		dateClause += " AND e.event_date < ?"
		args = append(args, *filter.To)
	}
	args = append(args, filter.Limit, filter.Offset)

	sql := `
WITH events AS (` + strings.Join(parts, "\nUNION ALL\n") + `
)
SELECT e.id, e.type, e.event_date, e.title, e.subtitle, e.status, e.has_note,
       e.author_id, COALESCE(u.name, '') AS author_name,
       COALESCE(u.roles::text, '') AS author_roles,
       e.patient_id, e.created_at
  FROM events e
  LEFT JOIN users u ON u.id = e.author_id
 WHERE 1=1` + authorClause + dateClause + `
 ORDER BY e.event_date DESC
 LIMIT ? OFFSET ?`

	rows := []MedicalRecordEntry{}
	if err := s.db.Raw(sql, args...).Scan(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}
