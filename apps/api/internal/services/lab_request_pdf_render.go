package services

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/pdfdoc"
)

var monthsPT = []string{"", "janeiro", "fevereiro", "março", "abril", "maio", "junho",
	"julho", "agosto", "setembro", "outubro", "novembro", "dezembro"}

func saoPaulo() *time.Location {
	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		return time.UTC
	}
	return loc
}

// maskCPF — "12345678901" => "***.456.789-**" (esconde 1º grupo e dígitos verificadores; LGPD).
func maskCPF(cpf string) string {
	d := strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, cpf)
	if len(d) != 11 {
		return ""
	}
	return "***." + d[3:6] + "." + d[6:9] + "-**"
}

func patientBirthInfo(p *models.Patient) string {
	if p == nil || p.BirthDate.IsZero() {
		return ""
	}
	loc := saoPaulo()
	b := p.BirthDate.In(loc)
	now := time.Now().In(loc)
	age := now.Year() - b.Year()
	if now.YearDay() < b.YearDay() {
		age--
	}
	return fmt.Sprintf("%s · %d anos", b.Format("02/01/2006"), age)
}

// doctorCredentials — "CRM-PR 21.876 · RQE 16.038 · Nefrologia" (campos opcionais).
func doctorCredentials(d *models.User) string {
	if d == nil {
		return ""
	}
	var parts []string
	if d.CRM != nil && *d.CRM != "" {
		uf := ""
		if d.CRMUF != nil {
			uf = *d.CRMUF
		}
		parts = append(parts, strings.TrimSpace("CRM-"+uf+" "+*d.CRM))
	}
	if d.RQE != nil && *d.RQE != "" {
		parts = append(parts, "RQE "+strings.TrimPrefix(*d.RQE, "RQE-"))
	}
	if d.Specialty != nil && *d.Specialty != "" {
		parts = append(parts, *d.Specialty)
	}
	return strings.Join(parts, " · ")
}

func placeDatePT(t time.Time) string {
	t = t.In(saoPaulo())
	return fmt.Sprintf("Londrina, %d de %s de %d", t.Day(), monthsPT[int(t.Month())], t.Year())
}

func signedAtPT(t *time.Time) string {
	if t == nil {
		return ""
	}
	l := t.In(saoPaulo())
	return l.Format("02/01/2006, 15:04") + " (horário de Brasília)"
}

// buildExamRequest mapeia o model LabRequest para o input do pacote pdfdoc (sistema-base).
func buildExamRequest(lr *models.LabRequest) pdfdoc.ExamRequest {
	var p pdfdoc.Patient
	if lr.Patient != nil {
		cpf := ""
		if lr.Patient.CPF != nil {
			cpf = maskCPF(*lr.Patient.CPF)
		}
		p = pdfdoc.Patient{Name: lr.Patient.Name, BirthInfo: patientBirthInfo(lr.Patient), CPFMasked: cpf}
	}
	doc := pdfdoc.Doctor{}
	if lr.Doctor != nil {
		doc = pdfdoc.Doctor{Name: lr.Doctor.Name, Credentials: doctorCredentials(lr.Doctor)}
	}
	indic := ""
	if lr.Notes != nil {
		indic = strings.TrimSpace(*lr.Notes)
	}
	sig := pdfdoc.Signature{
		Digital:     lr.Doctor != nil && lr.Doctor.CertificateActive && lr.SignedAt != nil,
		SignedAt:    signedAtPT(lr.SignedAt),
		ValidateURL: fmt.Sprintf("https://app.plenyasaude.com.br/lab-requests/validate/%s", lr.ID),
		PlaceDate:   placeDatePT(lr.Date),
	}
	return pdfdoc.ExamRequest{
		Patient:    p,
		Indication: indic,
		Exams:      lr.Exams,
		Doctor:     doc,
		Signature:  sig,
	}
}

// renderAndSaveLabRequestPDF gera o PDF vetorial (pdfdoc) e grava no volume de uploads,
// mantendo o mesmo esquema de nome usado pelo fluxo de assinatura. Função de pacote para
// ser reusada pelos dois fluxos (LabRequestService ativo + LabRequestPDFService de assinatura).
func renderAndSaveLabRequestPDF(lr *models.LabRequest) (string, error) {
	pdfBytes, err := pdfdoc.RenderExamRequest(buildExamRequest(lr))
	if err != nil {
		return "", fmt.Errorf("render pdf: %w", err)
	}
	patientName := "Paciente"
	if lr.Patient != nil && lr.Patient.Name != "" {
		patientName = lr.Patient.Name
	}
	filename := fmt.Sprintf("%s_labRequest_%s_%s.pdf",
		getInitials(patientName), lr.Date.Format("2006-01-02"), lr.ID.String()[:8])
	if err := os.MkdirAll("/app/uploads/lab-requests", 0o755); err != nil {
		return "", fmt.Errorf("mkdir uploads: %w", err)
	}
	if err := os.WriteFile("/app/uploads/lab-requests/"+filename, pdfBytes, 0o644); err != nil {
		return "", fmt.Errorf("save pdf: %w", err)
	}
	return "/uploads/lab-requests/" + filename, nil
}
