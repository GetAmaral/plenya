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

// verboDaVia — o verbo com que a frase da posologia começa. A via decide: "tomar" um colírio ou
// "tomar" uma injeção é instrução errada, e é o paciente que lê esta linha, não o farmacêutico.
func verboDaVia(rota string) string {
	switch strings.ToLower(strings.TrimSpace(rota)) {
	// Via EM BRANCO cai no default, de propósito. Ela só existe em registro antigo (a coluna é
	// not null e o DTO exige), e tratá-la como oral fazia o PDF inventar "Tomar 1 seringa" — uma
	// instrução de administração errada num documento assinado. "Usar" não afirma nada.
	case "oral", "sublingual":
		return "Tomar"
	// Tópicas e injetáveis: o mesmo verbo, por motivos diferentes — nenhuma delas se engole.
	case "tópica", "topica", "oftálmica", "oftalmica", "nasal",
		"intravenosa", "intramuscular", "subcutânea", "subcutanea":
		return "Aplicar"
	default:
		return "Usar"
	}
}

// medPosology — a segunda linha do item, escrita como frase para o paciente:
// "Tomar 1 comprimido uma vez ao dia". A duração e a quantidade NÃO entram aqui: elas vão para o
// campo à direita do nome, ligado pela guia pontilhada.
//
// A via só aparece quando não é oral. Em oral ela é o que o leitor já assume, e repetir empurra a
// informação que importa para o fim da linha; nas demais, omitir seria o erro de administração.
func medPosology(m models.PrescriptionMedication) string {
	frase := verboDaVia(m.Route)
	var p []string
	if m.Dosage != "" {
		p = append(p, m.Dosage)
	}
	if m.Frequency != "" {
		p = append(p, m.Frequency)
	}
	via := ""
	if r := strings.ToLower(strings.TrimSpace(m.Route)); r != "" && r != "oral" {
		via = "via " + strings.TrimSpace(m.Route)
	}
	if len(p) == 0 {
		// Registro antigo sem dosagem nem frequência: a frase perde o sentido, mas a VIA não pode
		// sumir do papel. Devolvê-la sozinha diz menos do que se queria e nada que seja falso.
		if via == "" {
			return ""
		}
		return strings.ToUpper(via[:1]) + via[1:]
	}
	frase += " " + strings.Join(p, " ")
	if via != "" {
		frase += ", " + via
	}
	// A duração só volta para a frase quando o campo da direita está ocupado pela QUANTIDADE. Sem
	// isto, "60 comprimidos, por 30 dias" perdia o prazo: o campo mostra um dos dois, e o que
	// sobrasse de fora sumia do papel. Quando a duração é quem ocupa o campo, repeti-la aqui seria
	// dizer duas vezes a mesma coisa.
	if m.Quantity > 0 {
		// O campo da direita está ocupado pela quantidade, então é AQUI que o prazo tem de sair —
		// e a ausência de prazo também. Sem este ramo, "Losartana ... 30 / Tomar 1 comprimido uma
		// vez ao dia" não dizia em lugar nenhum que o uso é contínuo, que é exatamente o silêncio
		// que esta mudança foi escrita para tirar da receita.
		if m.Duration > 0 {
			frase += ", por " + duracaoEmDias(m.Duration)
		} else {
			frase += ", uso contínuo"
		}
	}
	return frase
}

// medDispense — o campo à DIREITA do nome, no fim da guia pontilhada. É o que a farmácia lê para
// saber quanto entregar, e o que diz que não há prazo quando não há.
//
// Duração em branco é uso contínuo, e a receita precisa DIZER isso. Só omitir deixava o item sem
// prazo nenhum, e quem lê não sabe se o médico esqueceu ou se não há prazo — a farmácia e o
// paciente leem coisas diferentes desse silêncio.
func medDispense(m models.PrescriptionMedication) string {
	if m.Quantity > 0 {
		// Com o extenso, "60 (sessenta comprimidos)" se explica sozinho. Sem ele — e nada obriga o
		// extenso fora dos controlados — um "30" pelado no mesmo campo que às vezes traz
		// "por 30 dias" lê como prazo. O rótulo resolve a ambiguidade sem inventar unidade.
		if w := strings.TrimSpace(m.QuantityInWords); w != "" {
			return strconv.Itoa(m.Quantity) + " (" + w + ")"
		}
		return "Quantidade: " + strconv.Itoa(m.Quantity)
	}
	if m.Duration > 0 {
		return "por " + duracaoEmDias(m.Duration)
	}
	return "uso contínuo"
}

// duracaoEmDias — "1 dia" / "30 dias".
func duracaoEmDias(d int) string {
	if d == 1 {
		return "1 dia"
	}
	return fmt.Sprintf("%d dias", d)
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

// formulaDispenseLine — "60 (sessenta) cápsulas". Sem o "Aviar": no layout do receituário a
// palavra é a etiqueta da linha, impressa pelo pdfdoc, e viria duplicada se saísse daqui.
func formulaDispenseLine(f models.PrescriptionFormula) string {
	qty := formatQuantityPT(f.QuantityToDispense)
	if w := strings.TrimSpace(f.QuantityInWords); w != "" {
		qty += " (" + w + ")"
	}
	return strings.TrimSpace(qty + " " + f.QuantityUnit)
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
	// Mesma regra do industrializado: sem prazo é uso contínuo, e a receita diz.
	if f.Duration > 0 {
		p = append(p, "por "+duracaoEmDias(f.Duration))
	} else {
		p = append(p, "uso contínuo")
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
		instr := ""
		if m.Instructions != nil {
			instr = strings.TrimSpace(*m.Instructions)
		}
		meds = append(meds, pdfdoc.Med{
			Name:             m.MedicationName,
			Concentration:    m.Concentration,
			ActiveIngredient: m.ActiveIngredient,
			Posology:         medPosology(m),
			Quantity:         medDispense(m),
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
