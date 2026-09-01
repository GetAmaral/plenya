package services

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/pdfdoc"
)

// ErrPlanDossierNotFound — o plano ainda não tem dossiê congelado.
var ErrPlanDossierNotFound = errors.New("este plano ainda não tem dossiê congelado")

// marcasDoProntuario — o "quando foi a última vez" de cada fonte do dossiê.
//
// São as três coisas que fazem um dossiê envelhecer: exame novo, aferição nova de consulta e
// recálculo de escore. Comparar isto com o que estava gravado no congelamento custa UMA query e
// responde se vale a pena refrescar; descobrir remontando o dossiê custaria as ~28 que o
// congelamento existe para evitar.
type marcasDoProntuario struct {
	LatestLabAt      *time.Time
	LatestVitalsAt   *time.Time
	LatestSnapshotAt *time.Time
}

// Freeze monta o dossiê vivo e o congela como o próximo `seq` do plano, apontando
// `current_dossier_id`. Roda dentro da transação do chamador quando há uma.
func (s *PatientPlanDossierService) Freeze(tx *gorm.DB, planID, patientID uuid.UUID, byUserID *uuid.UUID) (*models.PatientPlanDossier, error) {
	if tx == nil {
		tx = s.db
	}
	vivo, err := s.Build(patientID)
	if err != nil {
		return nil, err
	}
	payload, err := json.Marshal(vivo)
	if err != nil {
		return nil, err
	}
	marcas, err := s.marcas(tx, patientID)
	if err != nil {
		return nil, err
	}

	var ultimo int
	if err := tx.Model(&models.PatientPlanDossier{}).
		Where("plan_id = ?", planID).
		Select("COALESCE(MAX(seq), 0)").Scan(&ultimo).Error; err != nil {
		return nil, err
	}

	congelado := models.PatientPlanDossier{
		PlanID:           planID,
		Seq:              ultimo + 1,
		Payload:          payload,
		LatestLabAt:      marcas.LatestLabAt,
		LatestVitalsAt:   marcas.LatestVitalsAt,
		LatestSnapshotAt: marcas.LatestSnapshotAt,
		BuiltAt:          time.Now(),
		BuiltByID:        byUserID,
	}
	if vivo.Snapshot != nil {
		if id, err := uuid.Parse(vivo.Snapshot.ID); err == nil {
			congelado.SourceSnapshotID = &id
		}
	}
	if err := tx.Create(&congelado).Error; err != nil {
		return nil, err
	}
	if err := tx.Model(&models.PatientPlan{}).Where("id = ?", planID).
		Update("current_dossier_id", congelado.ID).Error; err != nil {
		return nil, err
	}
	return &congelado, nil
}

// Current devolve o dossiê congelado em vigor, já desserializado.
//
// É ele que a tela de autoria lê. O dossiê vivo (`Build`) continua servindo quem quer o estado de
// agora, mas nunca a autoria: chão que se move enquanto se escreve é a origem do problema.
func (s *PatientPlanDossierService) Current(planID uuid.UUID) (*dto.PlanDossierResponse, *models.PatientPlanDossier, error) {
	var row models.PatientPlanDossier
	err := s.db.Where("plan_id = ?", planID).Order("seq DESC").First(&row).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil, ErrPlanDossierNotFound
	}
	if err != nil {
		return nil, nil, err
	}
	var payload dto.PlanDossierResponse
	if err := json.Unmarshal(row.Payload, &payload); err != nil {
		return nil, nil, err
	}
	return &payload, &row, nil
}

// Staleness diz se o prontuário andou desde o congelamento, e no quê.
//
// Não devolve "está velho" e pronto: devolve quais fontes mudaram, porque a decisão de refrescar
// depende disso. Exame novo num deck que fala de exame importa; aferição nova num deck que não cita
// pressão, não.
func (s *PatientPlanDossierService) Staleness(planID, patientID uuid.UUID) (*dto.PlanDossierStaleness, error) {
	var row models.PatientPlanDossier
	err := s.db.Where("plan_id = ?", planID).Order("seq DESC").First(&row).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &dto.PlanDossierStaleness{Stale: false}, nil
	}
	if err != nil {
		return nil, err
	}
	agora, err := s.marcas(s.db, patientID)
	if err != nil {
		return nil, err
	}

	out := dto.PlanDossierStaleness{
		DossierID: row.ID.String(),
		FrozenAt:  row.BuiltAt.In(saoPaulo()).Format(time.RFC3339),
	}
	if depois(agora.LatestSnapshotAt, row.LatestSnapshotAt) {
		out.Reasons = append(out.Reasons, "escore recalculado")
	}
	if depois(agora.LatestLabAt, row.LatestLabAt) {
		out.Reasons = append(out.Reasons, "exame novo")
	}
	if depois(agora.LatestVitalsAt, row.LatestVitalsAt) {
		out.Reasons = append(out.Reasons, "aferição nova na consulta")
	}
	out.Stale = len(out.Reasons) > 0
	return &out, nil
}

// RefreshDiff congela um dossiê novo e devolve o que mudou RESTRITO aos exames que o deck cita.
//
// A restrição é o ponto. Um dossiê tem 83 réguas; dizer "mudou" sobre o conjunto não ajuda ninguém.
// O que o médico precisa saber é: destes que você citou, estes três mudaram, e estão nestes slides.
func (s *PatientPlanDossierService) RefreshDiff(planID, patientID uuid.UUID, byUserID *uuid.UUID, slides []pdfdoc.DeckSlide) (*dto.PlanDossierRefreshResponse, error) {
	anterior, _, err := s.Current(planID)
	if err != nil && !errors.Is(err, ErrPlanDossierNotFound) {
		return nil, err
	}

	var novoRow *models.PatientPlanDossier
	if err := s.db.Transaction(func(tx *gorm.DB) error {
		r, err := s.Freeze(tx, planID, patientID, byUserID)
		novoRow = r
		return err
	}); err != nil {
		return nil, err
	}
	novo, _, err := s.Current(planID)
	if err != nil {
		return nil, err
	}

	out := dto.PlanDossierRefreshResponse{DossierID: novoRow.ID.String()}
	if anterior == nil {
		return &out, nil
	}

	citados := codigosCitados(slides)
	for code, ondeEsta := range citados {
		antes, tinha := anterior.Rulers[code]
		agora, tem := novo.Rulers[code]
		if !tinha || !tem {
			continue
		}
		va, ta := ultimoPontoDaRegua(antes)
		vd, td := ultimoPontoDaRegua(agora)
		if ta == td && va == vd {
			out.Unaffected++
			continue
		}
		out.Changed = append(out.Changed, dto.PlanDossierChange{
			Code: code, Name: agora.Name, Unit: agora.Unit,
			Was: ta, Now: td, CitedIn: ondeEsta,
		})
	}
	return &out, nil
}

// marcas lê as três marcas d'água numa consulta só.
func (s *PatientPlanDossierService) marcas(tx *gorm.DB, patientID uuid.UUID) (marcasDoProntuario, error) {
	var m marcasDoProntuario
	err := tx.Raw(`
		SELECT
		  (SELECT max(b.result_date) FROM lab_results r
		     JOIN lab_result_batches b ON b.id = r.lab_result_batch_id
		    WHERE b.patient_id = ? AND r.deleted_at IS NULL)      AS latest_lab_at,
		  (SELECT max(measured_at) FROM consultation_vitals
		    WHERE patient_id = ?)                                  AS latest_vitals_at,
		  (SELECT max(calculated_at) FROM patient_score_snapshots
		    WHERE patient_id = ?)                                  AS latest_snapshot_at
	`, patientID, patientID, patientID).Scan(&m).Error
	return m, err
}

func depois(a, b *time.Time) bool {
	if a == nil {
		return false
	}
	if b == nil {
		return true
	}
	return a.After(*b)
}

// codigosCitados varre os slides e devolve, por código de exame, onde ele aparece.
//
// Cobre régua de slide e régua embutida em linha de resumo, que são os dois lugares onde um número
// do dossiê chega ao deck com âncora.
func codigosCitados(slides []pdfdoc.DeckSlide) map[string][]dto.PlanDossierCitation {
	out := map[string][]dto.PlanDossierCitation{}
	anota := func(code string, c dto.PlanDossierCitation) {
		if code == "" {
			return
		}
		out[code] = append(out[code], c)
	}
	for i := range slides {
		s := slides[i]
		for j := range s.Rulers {
			anota(s.Rulers[j].Code, dto.PlanDossierCitation{
				SlideID: s.ID, Index: i + 1, Title: s.Title,
			})
		}
		if s.Summary == nil {
			continue
		}
		for ci := range s.Summary.Cards {
			for li := range s.Summary.Cards[ci].Lines {
				linha := s.Summary.Cards[ci].Lines[li]
				code := linha.Code
				if code == "" && linha.Ruler != nil {
					code = linha.Ruler.Code
				}
				anota(code, dto.PlanDossierCitation{
					SlideID: s.ID, Index: i + 1, Title: s.Title,
				})
			}
		}
	}
	return out
}

// ultimoPontoDaRegua devolve o valor e o texto do ponto mais recente do histórico.
func ultimoPontoDaRegua(r dto.PlanRuler) (float64, string) {
	if len(r.History) == 0 {
		return 0, ""
	}
	p := r.History[len(r.History)-1]
	return p.Value, p.Text
}
