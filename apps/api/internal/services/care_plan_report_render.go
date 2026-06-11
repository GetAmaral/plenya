package services

import (
	"fmt"
	"time"

	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/pdfdoc"
)

func reportLevelKind(lvl int) string {
	switch {
	case lvl <= 1:
		return "critical"
	case lvl == 2:
		return "attention"
	case lvl == 4:
		return "good"
	case lvl == 5:
		return "optimal"
	default:
		return "attention"
	}
}

// buildCarePlanReport mapeia os dados do relatório longitudinal AGIR para o input do pacote pdfdoc.
func buildCarePlanReport(
	patient *models.Patient,
	snapshot *models.PatientScoreSnapshot,
	items []models.CarePlanItem,
	doctor *models.User,
	validationURL string,
	hasDigital bool,
	now time.Time,
) pdfdoc.CarePlanReport {
	var attention, optimal []pdfdoc.ReportBiomarker
	for i := range snapshot.ItemResults {
		r := snapshot.ItemResults[i]
		if r.Status != models.EvaluationStatusEvaluated || r.LevelNumber == nil || r.Item == nil {
			continue
		}
		lvl := *r.LevelNumber
		entry := pdfdoc.ReportBiomarker{Name: r.Item.Name, Label: reportLevelLabel[lvl], Kind: reportLevelKind(lvl)}
		if lvl <= 2 {
			attention = append(attention, entry)
		} else if lvl == 4 || lvl == 5 {
			optimal = append(optimal, entry)
		}
	}

	var pillars []pdfdoc.ReportPillar
	for _, letter := range []string{"A", "G", "I", "R"} {
		var recs []pdfdoc.ReportRec
		for _, it := range items {
			if it.LetterCode != letter {
				continue
			}
			rec := pdfdoc.ReportRec{Text: it.Recommendation}
			if it.Target != nil {
				rec.Target = *it.Target
			}
			recs = append(recs, rec)
		}
		if len(recs) > 0 {
			pillars = append(pillars, pdfdoc.ReportPillar{Letter: letter, Name: agirNames[letter], Recs: recs})
		}
	}

	return pdfdoc.CarePlanReport{
		Patient:     pdfdoc.Patient{Name: patient.Name},
		EmittedAt:   now.In(saoPaulo()).Format("02/01/2006"),
		ScoreCalcAt: snapshot.CalculatedAt.In(saoPaulo()).Format("02/01/2006"),
		ScorePct:    fmt.Sprintf("%.0f%%", snapshot.TotalScorePercentage),
		Attention:   attention,
		Optimal:     optimal,
		Pillars:     pillars,
		Doctor:      pdfdoc.Doctor{Name: doctor.Name, Credentials: doctorCredentials(doctor)},
		Signature: pdfdoc.Signature{
			Digital:     hasDigital,
			SignedAt:    signedAtPT(&now),
			ValidateURL: validationURL,
			PlaceDate:   placeDatePT(now),
		},
	}
}

// renderCarePlanReportBytes gera os bytes do relatório AGIR pelo pipeline pdfdoc (papelaria de marca).
func renderCarePlanReportBytes(
	patient *models.Patient,
	snapshot *models.PatientScoreSnapshot,
	items []models.CarePlanItem,
	doctor *models.User,
	validationURL string,
	hasDigital bool,
	now time.Time,
) ([]byte, error) {
	return pdfdoc.RenderCarePlanReport(buildCarePlanReport(patient, snapshot, items, doctor, validationURL, hasDigital, now))
}
