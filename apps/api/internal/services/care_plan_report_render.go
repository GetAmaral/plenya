package services

import (
	"fmt"
	"time"

	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/pdfdoc"
)

// O relatório AGIR é UM MODO DO DECK, não um gerador próprio.
//
// Antes havia dois renderizadores de relatório: este, que derivava o conteúdo do escore, e o do
// plano, que renderiza slides escritos à mão. Dois caminhos para o mesmo documento significam duas
// papelarias para manter em sincronia e dois lugares para corrigir quando algo sai torto.
//
// Agora a derivação produz SLIDES, e o render é o mesmo `pdfdoc.RenderPlanReport` que o plano usa.
// O que muda entre um relatório derivado do escore e um escrito pelo médico é só quem escreveu o
// conteúdo.

// buildCarePlanDeck traduz escore + plano de cuidado nos slides do relatório.
func buildCarePlanDeck(
	patient *models.Patient,
	snapshot *models.PatientScoreSnapshot,
	items []models.CarePlanItem,
	now time.Time,
) []pdfdoc.DeckSlide {
	var slides []pdfdoc.DeckSlide

	// Abertura: o escore e quando foi calculado.
	slides = append(slides, pdfdoc.DeckSlide{
		Kind:    pdfdoc.DeckCover,
		Eyebrow: "Escore calculado em " + snapshot.CalculatedAt.In(saoPaulo()).Format("02/01/2006"),
		Title:   patient.Name,
		Lede: fmt.Sprintf("Escore Plenya de saúde global: %.0f%%.",
			snapshot.TotalScorePercentage),
	})

	// Biomarcadores, separados como no deck: o que pede atenção e o que está no ótimo.
	var atencao, otimo []pdfdoc.DeckTableRow
	for i := range snapshot.ItemResults {
		r := snapshot.ItemResults[i]
		if r.Status != models.EvaluationStatusEvaluated || r.LevelNumber == nil || r.Item == nil {
			continue
		}
		lvl := *r.LevelNumber
		linha := pdfdoc.DeckTableRow{Cells: []string{r.Item.Name, reportLevelLabel[lvl]}}
		switch {
		case lvl <= 2:
			atencao = append(atencao, linha)
		case lvl >= 4:
			otimo = append(otimo, linha)
		}
	}
	colunas := []pdfdoc.DeckTableCol{{Label: "Marcador"}, {Label: "", Style: pdfdoc.DeckColTag}}
	if len(atencao) > 0 {
		slides = append(slides, pdfdoc.DeckSlide{
			Kind: pdfdoc.DeckTableKind, Eyebrow: "Pontos de atenção",
			Title: "O que pede atenção",
			Table: &pdfdoc.DeckTable{Dense: true, Columns: colunas, Rows: atencao},
		})
	}
	if len(otimo) > 0 {
		slides = append(slides, pdfdoc.DeckSlide{
			Kind: pdfdoc.DeckTableKind, Eyebrow: "No ótimo",
			Title: "O que está no melhor nível",
			Table: &pdfdoc.DeckTable{Dense: true, Columns: colunas, Rows: otimo},
		})
	}

	// Plano de cuidado, um slide por pilar AGIR.
	for _, letra := range []string{"A", "G", "I", "R"} {
		var rows []pdfdoc.DeckTableRow
		for _, it := range items {
			if it.LetterCode != letra {
				continue
			}
			meta := ""
			if it.Target != nil {
				meta = *it.Target
			}
			rows = append(rows, pdfdoc.DeckTableRow{Cells: []string{it.Recommendation, meta}})
		}
		if len(rows) == 0 {
			continue
		}
		slides = append(slides, pdfdoc.DeckSlide{
			Kind: pdfdoc.DeckTableKind, Eyebrow: "Plano de cuidado AGIR",
			Title: agirNames[letra],
			Table: &pdfdoc.DeckTable{
				Columns: []pdfdoc.DeckTableCol{{Label: "Recomendação"}, {Label: "Meta", Style: pdfdoc.DeckColDose}},
				Rows:    rows,
			},
		})
	}
	return slides
}

// renderCarePlanReportBytes gera o relatório AGIR pelo MESMO render do plano.
func renderCarePlanReportBytes(
	patient *models.Patient,
	snapshot *models.PatientScoreSnapshot,
	items []models.CarePlanItem,
	doctor *models.User,
	validationURL string,
	hasDigital bool,
	now time.Time,
) ([]byte, error) {
	return pdfdoc.RenderPlanReport(pdfdoc.PlanReport{
		Title:     "Relatório de Saúde, Performance e Longevidade",
		Patient:   pdfdoc.Patient{Name: patient.Name},
		Slides:    buildCarePlanDeck(patient, snapshot, items, now),
		EmittedAt: now.In(saoPaulo()).Format("02/01/2006"),
		Doctor:    pdfdoc.Doctor{Name: doctor.Name, Credentials: doctorCredentials(doctor)},
		Signature: pdfdoc.Signature{
			Digital:     hasDigital,
			SignedAt:    signedAtPT(&now),
			ValidateURL: validationURL,
			PlaceDate:   placeDatePT(now),
		},
	})
}
