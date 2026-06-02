package services

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"
)

var ErrNoSnapshotForReport = errors.New("paciente sem escore calculado — calcule o escore antes de gerar o relatório")

// CarePlanReportService gera o relatório longitudinal AGIR (Escore + biomarcadores
// Normal-vs-Ótimo + plano por pilar), converte HTML→PDF (go-rod), assina (ICP-Brasil ou
// degrada para assinatura manual) e publica no portal do paciente como IssuedDocument(report).
type CarePlanReportService struct {
	db           *gorm.DB
	snapshotRepo *repository.ScoreSnapshotRepository
	carePlan     *CarePlanService
	pdf          *ScorePDFService
	signature    *SignatureService
	documents    *PatientDocumentsService
}

func NewCarePlanReportService(
	db *gorm.DB,
	snapshotRepo *repository.ScoreSnapshotRepository,
	carePlan *CarePlanService,
	pdf *ScorePDFService,
	signature *SignatureService,
	documents *PatientDocumentsService,
) *CarePlanReportService {
	return &CarePlanReportService{db: db, snapshotRepo: snapshotRepo, carePlan: carePlan, pdf: pdf, signature: signature, documents: documents}
}

// GenerateAndPublish compõe, assina e publica o relatório. Retorna o IssuedDocument criado.
func (s *CarePlanReportService) GenerateAndPublish(patientID, doctorID uuid.UUID) (string, error) {
	var patient models.Patient
	if err := s.db.First(&patient, patientID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", ErrPatientNotFound
		}
		return "", err
	}

	snapshot, err := s.snapshotRepo.GetLatestByPatientID(patientID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", ErrNoSnapshotForReport
		}
		return "", err
	}

	items, err := s.carePlan.ListActiveModels(patientID)
	if err != nil {
		return "", err
	}

	var doctor models.User
	if err := s.db.First(&doctor, doctorID).Error; err != nil {
		return "", err
	}

	// 1. IssuedDocument(report) em draft → pega o ID p/ o QR de validação.
	doc := &models.IssuedDocument{
		PatientID:      patientID,
		DoctorID:       doctorID,
		Type:           models.IssuedDocReport,
		Title:          "Relatório de Saúde, Performance e Longevidade",
		Body:           "Relatório longitudinal AGIR.",
		Status:         models.IssuedDocDraft,
		IssuedByUserID: doctorID,
	}
	if err := s.db.Create(doc).Error; err != nil {
		return "", err
	}

	validationURL := fmt.Sprintf("https://plenya.com.br/documentos/validar/%s", doc.ID)
	hasDigital := doctor.CertificateActive

	render := func(digital bool) ([]byte, error) {
		htmlStr := s.buildReportHTML(&patient, snapshot, items, &doctor, validationURL, digital)
		return s.pdf.GeneratePDFFromHTML(htmlStr)
	}

	pdfBytes, err := render(hasDigital)
	if err != nil {
		return "", fmt.Errorf("erro ao gerar PDF do relatório: %w", err)
	}

	var finalPDF []byte
	var hash string
	var certSerial *string

	if hasDigital {
		signed, sigHash, sErr := s.signature.SignDocumentPDF(pdfBytes, doctorID,
			"Assinatura Digital de Relatório Médico", "Plenya EMR - Relatório AGIR")
		if sErr != nil {
			hasDigital = false
			unsigned, rErr := render(false)
			if rErr != nil {
				return "", rErr
			}
			finalPDF = unsigned
			sum := sha256.Sum256(unsigned)
			hash = hex.EncodeToString(sum[:])
		} else {
			finalPDF = signed
			hash = sigHash
			certSerial = doctor.CertificateSerial
		}
	} else {
		finalPDF = pdfBytes
		sum := sha256.Sum256(pdfBytes)
		hash = hex.EncodeToString(sum[:])
	}

	// 2. Publicar no portal (idempotente por doc.ID — retry após falha do Updates não duplica).
	uploadedBy := doctorID
	idemKey := "report:" + doc.ID.String()
	patientDoc, err := s.documents.CreateFromBytes(CreateFromBytesInput{
		PatientID:         patientID,
		Bytes:             finalPDF,
		Filename:          fmt.Sprintf("relatorio_agir_%s.pdf", doc.ID),
		Title:             doc.Title,
		Type:              models.DocumentTypeReport,
		Source:            models.DocumentSourceStaffUpload,
		UploadedBy:        &uploadedBy,
		OriginWAMessageID: &idemKey,
	})
	if err != nil {
		return "", fmt.Errorf("erro ao publicar relatório: %w", err)
	}

	// 3. Persistir metadados.
	now := time.Now()
	updates := map[string]interface{}{
		"status":                models.IssuedDocSigned,
		"has_digital_signature": hasDigital,
		"signed_at":             now,
		"signed_pdf_hash":       hash,
		"signed_pdf_path":       patientDoc.FilePath,
		"qr_code_data":          validationURL,
		"patient_document_id":   patientDoc.ID,
	}
	if certSerial != nil {
		updates["certificate_serial"] = *certSerial
	}
	if err := s.db.Model(doc).Updates(updates).Error; err != nil {
		return "", err
	}

	return doc.ID.String(), nil
}

var reportLevelLabel = map[int]string{
	0: "Crítico", 1: "Muito alterado", 2: "Subótimo", 3: "Limítrofe", 4: "Bom", 5: "Ótimo", 6: "N/A",
}
var reportLevelColor = map[int]string{
	0: "#dc2626", 1: "#ea580c", 2: "#ca8a04", 3: "#2563eb", 4: "#16a34a", 5: "#059669", 6: "#6b7280",
}
var agirNames = map[string]string{"A": "Atividade", "G": "Gestão", "I": "Integração", "R": "Ritmo"}

func (s *CarePlanReportService) buildReportHTML(
	patient *models.Patient,
	snapshot *models.PatientScoreSnapshot,
	items []models.CarePlanItem,
	doctor *models.User,
	validationURL string,
	hasDigital bool,
) string {
	var b strings.Builder
	esc := html.EscapeString

	// Biomarcadores por nível.
	type bm struct {
		name  string
		level int
	}
	var atencao, otimo []bm
	for i := range snapshot.ItemResults {
		r := snapshot.ItemResults[i]
		if r.Status != models.EvaluationStatusEvaluated || r.LevelNumber == nil || r.Item == nil {
			continue
		}
		lvl := *r.LevelNumber
		entry := bm{name: r.Item.Name, level: lvl}
		if lvl <= 2 {
			atencao = append(atencao, entry)
		} else if lvl == 4 || lvl == 5 {
			otimo = append(otimo, entry)
		}
	}

	b.WriteString(`<!DOCTYPE html><html lang="pt-BR"><head><meta charset="utf-8"><style>`)
	b.WriteString(`
* { box-sizing: border-box; }
body { font-family: -apple-system, "Segoe UI", Roboto, Arial, sans-serif; color: #1f2937; margin: 0; padding: 0; }
h1 { font-size: 22px; letter-spacing: 4px; font-weight: 600; color: #1f2937; margin: 0 0 2px; }
.tag { color: #6b7280; font-size: 12px; letter-spacing: 1px; }
.meta { margin: 18px 0 24px; font-size: 13px; color: #374151; }
.score-box { display: flex; align-items: baseline; gap: 12px; background: #f8fafc; border: 1px solid #e5e7eb; border-radius: 12px; padding: 18px 24px; margin-bottom: 24px; }
.score-num { font-size: 40px; font-weight: 700; color: #0f766e; }
.section-title { font-size: 14px; font-weight: 700; text-transform: uppercase; letter-spacing: 1px; color: #374151; border-bottom: 2px solid #e5e7eb; padding-bottom: 6px; margin: 22px 0 10px; }
table { width: 100%; border-collapse: collapse; font-size: 13px; }
td { padding: 5px 6px; border-bottom: 1px solid #f1f5f9; }
.lvl { display: inline-block; padding: 1px 8px; border-radius: 10px; color: #fff; font-size: 11px; }
.pillar { margin-bottom: 14px; }
.pillar-h { display: flex; align-items: center; gap: 8px; font-weight: 700; margin-bottom: 4px; }
.dot { width: 18px; height: 18px; border-radius: 50%; color: #fff; font-size: 11px; display: inline-flex; align-items: center; justify-content: center; font-weight: 700; }
.rec { font-size: 13px; padding: 4px 0 4px 26px; border-bottom: 1px solid #f1f5f9; }
.rec .target { color: #6b7280; font-size: 12px; }
.footer { margin-top: 36px; padding-top: 14px; border-top: 1px solid #e5e7eb; font-size: 11px; color: #6b7280; }
.footer .sig { color: #047857; font-weight: 600; }
@media print { body { padding: 0; } }
`)
	b.WriteString(`</style></head><body>`)

	b.WriteString(`<h1>PLENYA</h1><div class="tag">Saúde, Performance &amp; Longevidade</div>`)
	b.WriteString(`<div class="meta"><strong>` + esc(patient.Name) + `</strong><br>`)
	b.WriteString(`Relatório emitido em ` + time.Now().Format("02/01/2006") + ` · Escore calculado em ` + snapshot.CalculatedAt.Format("02/01/2006") + `</div>`)

	// Escore
	b.WriteString(`<div class="score-box"><span class="score-num">` + fmt.Sprintf("%.0f%%", snapshot.TotalScorePercentage) + `</span><span>Escore Plenya — saúde global</span></div>`)

	// Atenção
	dotColors := map[string]string{"A": "#d97706", "G": "#0d9488", "I": "#0284c7", "R": "#059669"}
	writeBM := func(title string, rows []bm) {
		if len(rows) == 0 {
			return
		}
		b.WriteString(`<div class="section-title">` + title + `</div><table>`)
		for _, r := range rows {
			b.WriteString(`<tr><td>` + esc(r.name) + `</td><td style="text-align:right"><span class="lvl" style="background:` + reportLevelColor[r.level] + `">` + reportLevelLabel[r.level] + `</span></td></tr>`)
		}
		b.WriteString(`</table>`)
	}
	writeBM("Pontos de atenção", atencao)
	writeBM("No ótimo", otimo)

	// Plano AGIR
	if len(items) > 0 {
		b.WriteString(`<div class="section-title">Plano de cuidado AGIR</div>`)
		for _, letter := range []string{"A", "G", "I", "R"} {
			var group []models.CarePlanItem
			for _, it := range items {
				if it.LetterCode == letter {
					group = append(group, it)
				}
			}
			if len(group) == 0 {
				continue
			}
			b.WriteString(`<div class="pillar"><div class="pillar-h"><span class="dot" style="background:` + dotColors[letter] + `">` + letter + `</span>` + agirNames[letter] + `</div>`)
			for _, it := range group {
				b.WriteString(`<div class="rec">` + esc(it.Recommendation))
				if it.Target != nil && *it.Target != "" {
					b.WriteString(` <span class="target">(meta: ` + esc(*it.Target) + `)</span>`)
				}
				b.WriteString(`</div>`)
			}
			b.WriteString(`</div>`)
		}
	}

	// Footer / assinatura
	b.WriteString(`<div class="footer">`)
	crm := ""
	if doctor.CRM != nil && doctor.CRMUF != nil {
		crm = "CRM-" + esc(*doctor.CRMUF) + " " + esc(*doctor.CRM)
	}
	if hasDigital {
		b.WriteString(`<span class="sig">Documento assinado digitalmente (ICP-Brasil)</span> · ` + esc(doctor.Name) + ` · ` + crm + `<br>`)
		b.WriteString(`Validar: ` + validationURL + ` · Verificar assinatura: https://validar.iti.gov.br`)
	} else {
		b.WriteString(esc(doctor.Name) + ` · ` + crm + `<br>Validar: ` + validationURL)
	}
	b.WriteString(`</div></body></html>`)
	return b.String()
}
