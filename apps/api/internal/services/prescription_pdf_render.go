package services

import (
	"fmt"
	"strings"

	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/pdfdoc"
)

// formatCPFFull — "12345678901" => "123.456.789-01" (CPF completo do prescritor; usado SÓ em
// receituário de controle, exigido pela RDC Anvisa 1.000/2025).
func formatCPFFull(cpf string) string {
	d := strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, cpf)
	if len(d) != 11 {
		return ""
	}
	return d[0:3] + "." + d[3:6] + "." + d[6:9] + "-" + d[9:11]
}

func medPosology(m models.PrescriptionMedication) string {
	var p []string
	if m.Dosage != "" {
		p = append(p, m.Dosage)
	}
	if m.Frequency != "" {
		p = append(p, m.Frequency)
	}
	if m.Route != "" {
		p = append(p, "via "+m.Route)
	}
	if m.Duration > 0 {
		unit := "dias"
		if m.Duration == 1 {
			unit = "dia"
		}
		p = append(p, fmt.Sprintf("por %d %s", m.Duration, unit))
	}
	return strings.Join(p, " · ")
}

// buildPrescription mapeia o model Prescription para o input do pacote pdfdoc.
func buildPrescription(p *models.Prescription, manual bool) pdfdoc.Prescription {
	var pat pdfdoc.Patient
	cpf := ""
	if p.Patient.CPF != nil {
		cpf = maskCPF(*p.Patient.CPF)
	}
	pat = pdfdoc.Patient{Name: p.Patient.Name, BirthInfo: patientBirthInfo(&p.Patient), CPFMasked: cpf}

	meds := make([]pdfdoc.Med, 0, len(p.Medications))
	for _, m := range p.Medications {
		qty := ""
		if m.Quantity > 0 {
			qty = fmt.Sprintf("Quantidade: %d (%s)", m.Quantity, m.QuantityInWords)
		}
		instr := ""
		if m.Instructions != nil {
			instr = strings.TrimSpace(*m.Instructions)
		}
		meds = append(meds, pdfdoc.Med{
			Name:             m.MedicationName,
			Concentration:    m.Concentration,
			ActiveIngredient: m.ActiveIngredient,
			Posology:         medPosology(m),
			Quantity:         qty,
			Instructions:     instr,
		})
	}

	cred := doctorCredentials(&p.Doctor)
	controlLabel := ""
	if prescriptionHasControlled(p) {
		controlLabel = "Receituário de Controle Especial"
		// RDC 1.000/2025: CPF do prescritor obrigatório em receita de controle.
		if p.Doctor.CPF != nil {
			if c := formatCPFFull(*p.Doctor.CPF); c != "" {
				cred = strings.TrimSpace(cred + " · CPF " + c)
			}
		}
	}

	general := ""
	if p.GeneralInstructions != nil {
		general = strings.TrimSpace(*p.GeneralInstructions)
	}
	validUntil := ""
	if !p.ValidUntil.IsZero() {
		validUntil = p.ValidUntil.In(saoPaulo()).Format("02/01/2006")
	}

	return pdfdoc.Prescription{
		ControlLabel:        controlLabel,
		Patient:             pat,
		Meds:                meds,
		GeneralInstructions: general,
		ValidUntil:          validUntil,
		Doctor:              pdfdoc.Doctor{Name: p.Doctor.Name, Credentials: cred},
		Signature: pdfdoc.Signature{
			Digital:     !manual,
			SignedAt:    signedAtPT(p.SignedAt),
			ValidateURL: fmt.Sprintf("https://app.plenyasaude.com.br/prescriptions/validate/%s", p.ID),
			PlaceDate:   placeDatePT(p.PrescriptionDate),
		},
	}
}

// renderPrescriptionBytes gera os bytes do PDF do receituário pelo pipeline pdfdoc (substitui
// generatePDFContent/gofpdf). Modo manual => carimbo manual; senão => selo ICP-Brasil.
func renderPrescriptionBytes(p *models.Prescription, manual bool) ([]byte, error) {
	return pdfdoc.RenderPrescription(buildPrescription(p, manual))
}
