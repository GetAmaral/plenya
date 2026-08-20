package services

import (
	"fmt"
	"strconv"
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

// formatQuantityPT — número em pt-BR, sem zeros à toa: 3 → "3", 0.25 → "0,25", 12.5 → "12,5".
// Manipulado usa fração de miligrama o tempo todo; imprimir "3.0000 mg" numa receita é ruído.
func formatQuantityPT(v float64) string {
	s := strconv.FormatFloat(v, 'f', -1, 64)
	return strings.Replace(s, ".", ",", 1)
}

// componentQuantity — "300 mg", "0,25 mg", "10%". Porcentagem cola no número (convenção
// brasileira); as demais unidades levam espaço.
func componentQuantity(qty float64, unit string) string {
	u := strings.TrimSpace(unit)
	n := formatQuantityPT(qty)
	if u == "%" {
		return n + u
	}
	return strings.TrimSpace(n + " " + u)
}

// usageLabel — o destaque que evita o erro de administração mais grave do manipulado.
func usageLabel(u models.FormulaUsage) string {
	if u == models.FormulaUsageExternal {
		return "Uso externo"
	}
	return "Uso interno"
}

// formulaDispenseLine — "Aviar 60 (sessenta) cápsulas".
func formulaDispenseLine(f models.PrescriptionFormula) string {
	qty := formatQuantityPT(f.QuantityToDispense)
	if w := strings.TrimSpace(f.QuantityInWords); w != "" {
		qty += " (" + w + ")"
	}
	return strings.TrimSpace("Aviar " + qty + " " + f.QuantityUnit)
}

// formulaPosologyLine — posologia + duração, no mesmo formato do industrializado.
func formulaPosologyLine(f models.PrescriptionFormula) string {
	var p []string
	if s := strings.TrimSpace(f.Posology); s != "" {
		p = append(p, s)
	}
	if r := strings.TrimSpace(f.Route); r != "" {
		p = append(p, "via "+r)
	}
	if f.Duration > 0 {
		unit := "dias"
		if f.Duration == 1 {
			unit = "dia"
		}
		p = append(p, fmt.Sprintf("por %d %s", f.Duration, unit))
	}
	return strings.Join(p, " · ")
}

// buildFormulasPDF mapeia as fórmulas do model para o input do pdfdoc.
func buildFormulasPDF(formulas []models.PrescriptionFormula) []pdfdoc.Formula {
	out := make([]pdfdoc.Formula, 0, len(formulas))
	for _, f := range formulas {
		comps := make([]pdfdoc.FormulaComponent, 0, len(f.Components))
		for _, c := range f.Components {
			comps = append(comps, pdfdoc.FormulaComponent{
				Substance: c.Substance,
				Quantity:  componentQuantity(c.Quantity, c.Unit),
				Note:      c.Note,
				// Sem isto, a conversão elemento/insumo que a tela calcula morre antes da
				// farmácia: o PDF é o único documento que ela lê.
				AsElemental: c.AsElemental,
			})
		}
		instr := ""
		if f.Instructions != nil {
			instr = strings.TrimSpace(*f.Instructions)
		}
		out = append(out, pdfdoc.Formula{
			Name:         f.Name,
			Form:         f.PharmaceuticalForm,
			UsageLabel:   usageLabel(f.UsageType),
			Components:   comps,
			Vehicle:      strings.TrimSpace(f.Vehicle),
			Dispense:     formulaDispenseLine(f),
			Posology:     formulaPosologyLine(f),
			Instructions: instr,
		})
	}
	return out
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
		// valid_until é DATE pura: o driver devolve meia-noite UTC. Converter para São Paulo
		// (UTC-3) recuava um dia e a receita saía impressa vencendo 24h antes do que vale.
		validUntil = p.ValidUntil.UTC().Format("02/01/2006")
	}

	title := ""
	if p.Type == models.PrescriptionCompounded {
		title = "Receituário magistral"
	}

	return pdfdoc.Prescription{
		// O tipo da receita é que decide o layout: sem isto, o renderizador escolhia por "tem
		// fórmula" e uma receita com as duas listas descartava os industrializados em silêncio.
		Compounded:          p.Type == models.PrescriptionCompounded,
		Title:               title,
		ControlLabel:        controlLabel,
		Patient:             pat,
		Meds:                meds,
		Formulas:            buildFormulasPDF(p.Formulas),
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
