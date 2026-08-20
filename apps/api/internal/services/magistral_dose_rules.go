package services

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/plenya/api/internal/models"
)

// Motor de sugestão de dose das fórmulas-base.
//
// TRÊS FRONTEIRAS ESTRUTURAIS, e elas valem mais que a matemática:
//
//  1. isto devolve SUGESTÃO, não prescrição. O payload de criação de receita não tem noção de
//     regra: mesmo que este arquivo esteja inteiro errado, ele é INCAPAZ de produzir receita
//     assinada — só preenche um campo que o médico ainda vai olhar;
//  2. toda regra tem piso e teto obrigatórios. Peso errado no prontuário ou exame em unidade
//     trocada não viram dose absurda: a trava corta e a resposta DIZ que cortou;
//  3. dado ausente ou velho não vira sugestão. Devolve o motivo, em português, para a tela
//     mostrar por que não sugeriu.
//
// Funções puras, sem banco: quem busca peso e exame é o service, que passa os valores aqui.

// MeasurementInput é um dado do prontuário com procedência. A data importa tanto quanto o valor:
// dose calculada sobre peso de três anos atrás é risco, não conveniência.
type MeasurementInput struct {
	Value  float64
	Date   *time.Time
	Source string // "consulta", "avaliação física", "cadastro", nome do exame
	Unit   string
}

// DoseSuggestion é o que a tela mostra ao lado do campo de dose.
type DoseSuggestion struct {
	Substance string  `json:"substance"`
	Unit      string  `json:"unit"`
	BaseDose  float64 `json:"baseDose"`

	// Suggested só vem preenchido quando há regra E dado para aplicá-la.
	Suggested *float64 `json:"suggested,omitempty"`
	// Basis explica de onde saiu, sempre presente quando há sugestão.
	Basis string `json:"basis,omitempty"`
	// Clamped diz que a trava agiu — a tela precisa mostrar isso, não esconder.
	Clamped bool `json:"clamped"`
	// Reason explica por que NÃO houve sugestão.
	Reason string `json:"reason,omitempty"`

	// Bands é a escada inteira da regra, com a faixa do paciente marcada. Ver só a dose
	// escolhida esconde a conduta: o médico decide melhor vendo a régua e onde ele caiu.
	Bands []DoseBandView `json:"bands,omitempty"`
}

// DoseBandView é uma faixa como a tela mostra.
type DoseBandView struct {
	LowerBound *float64 `json:"lowerBound,omitempty"`
	UpperBound *float64 `json:"upperBound,omitempty"`
	Dose       float64  `json:"dose"`
	Label      string   `json:"label,omitempty"`
	Active     bool     `json:"active"`
}

// SuggestInput reúne o que o motor precisa para um componente.
type SuggestInput struct {
	Component models.MagistralFormulaTemplateComponent
	Rule      *models.MagistralFormulaTemplateRule
	Weight    *MeasurementInput
	Lab       *MeasurementInput
	Now       time.Time
	// DosesPerDay — quantas tomadas o dia tem, lidas da posologia da fórmula-base. Toda regra é
	// escrita em dose DIÁRIA; o que a tela preenche é a dose de UMA cápsula. Zero é lido como 1.
	DosesPerDay float64
}

// SuggestDose aplica a regra de um componente.
func SuggestDose(in SuggestInput) DoseSuggestion {
	out := DoseSuggestion{
		Substance: in.Component.Substance,
		Unit:      in.Component.Unit,
		BaseDose:  in.Component.Quantity,
	}
	if in.Rule == nil {
		out.Reason = "sem regra de dose: usa a dose da fórmula-base"
		return out
	}

	rule := in.Rule
	var raw float64
	var basis string

	switch rule.Kind {
	case models.DoseRuleFixed:
		if rule.FixedDose == nil {
			out.Reason = "regra de dose fixa sem valor cadastrado"
			return out
		}
		raw = *rule.FixedDose
		basis = fmt.Sprintf("dose fixa da regra: %s %s", fmtDose(raw), in.Component.Unit)

	case models.DoseRulePerKg:
		if rule.PerKg == nil {
			out.Reason = "regra por peso sem valor por kg cadastrado"
			return out
		}
		if in.Weight == nil {
			out.Reason = "sem peso registrado para este paciente"
			return out
		}
		if stale, age := isStale(in.Weight.Date, rule.MaxDataAgeDays, in.Now); stale {
			out.Reason = fmt.Sprintf("peso registrado há %s, acima do limite de %d dias da regra",
				age, rule.MaxDataAgeDays)
			return out
		}
		raw = *rule.PerKg * in.Weight.Value
		basis = fmt.Sprintf("peso %s kg (%s) × %s %s/kg",
			fmtDose(in.Weight.Value), describeSource(in.Weight), fmtDose(*rule.PerKg), in.Component.Unit)

	case models.DoseRuleLabThreshold:
		if rule.LabCode == nil || rule.LabOperator == nil || rule.LabThreshold == nil || rule.DoseIfTrue == nil {
			out.Reason = "regra por exame incompleta"
			return out
		}
		if in.Lab == nil {
			out.Reason = "sem resultado deste exame para o paciente"
			return out
		}
		if msg := unitMismatch(rule.LabUnit, in.Lab.Unit); msg != "" {
			out.Reason = msg
			return out
		}
		if stale, age := isStale(in.Lab.Date, rule.MaxDataAgeDays, in.Now); stale {
			out.Reason = fmt.Sprintf("exame colhido há %s, acima do limite de %d dias da regra",
				age, rule.MaxDataAgeDays)
			return out
		}

		hit := compare(in.Lab.Value, *rule.LabOperator, *rule.LabThreshold)
		condicao := fmt.Sprintf("%s %s", operatorLabel(*rule.LabOperator), fmtDose(*rule.LabThreshold))
		if hit {
			raw = *rule.DoseIfTrue
		} else {
			if rule.DoseIfFalse == nil {
				out.Reason = fmt.Sprintf("%s %s %s não atinge o limiar da regra e não há dose alternativa cadastrada",
					describeSource(in.Lab), fmtDose(in.Lab.Value), in.Lab.Unit)
				return out
			}
			raw = *rule.DoseIfFalse
			// A frase precisa dizer QUAL ramo disparou: "regra: < 30" com valor 37 sugere que a
			// condição foi atendida, e ela não foi.
			condicao = "não atinge " + condicao
		}
		basis = fmt.Sprintf("%s = %s %s%s · regra: %s ⇒ %s %s",
			describeSource(in.Lab), fmtDose(in.Lab.Value), in.Lab.Unit, dateSuffix(in.Lab.Date),
			condicao, fmtDose(raw), in.Component.Unit)

	case models.DoseRuleLabBand:
		if rule.LabCode == nil || len(rule.Bands) == 0 {
			out.Reason = "regra por faixa sem exame ou sem faixas cadastradas"
			return out
		}
		out.Bands = bandViews(rule.Bands, nil)
		if in.Lab == nil {
			out.Reason = "sem resultado deste exame para o paciente"
			return out
		}
		if msg := unitMismatch(rule.LabUnit, in.Lab.Unit); msg != "" {
			out.Reason = msg
			return out
		}
		if stale, age := isStale(in.Lab.Date, rule.MaxDataAgeDays, in.Now); stale {
			out.Reason = fmt.Sprintf("exame colhido há %s, acima do limite de %d dias da regra",
				age, rule.MaxDataAgeDays)
			return out
		}

		band := matchBand(in.Lab.Value, rule.Bands)
		out.Bands = bandViews(rule.Bands, band)
		if band == nil {
			// Faixas que não cobrem todo o eixo são erro de cadastro, e o silêncio aqui é o
			// comportamento certo: inventar a faixa mais próxima seria adivinhar conduta.
			out.Reason = fmt.Sprintf("%s = %s %s não cai em nenhuma faixa cadastrada da regra",
				describeSource(in.Lab), fmtDose(in.Lab.Value), in.Lab.Unit)
			return out
		}
		raw = band.Dose
		basis = fmt.Sprintf("%s = %s %s%s · faixa %s%s ⇒ %s %s",
			describeSource(in.Lab), fmtDose(in.Lab.Value), in.Lab.Unit, dateSuffix(in.Lab.Date),
			bandLabel(*band), bandName(*band), fmtDose(raw), in.Component.Unit)

	default:
		out.Reason = "tipo de regra desconhecido"
		return out
	}

	// A trava é DIÁRIA, como a regra: cortar antes de dividir pelas tomadas.
	final := raw
	if final < rule.MinDose {
		final = rule.MinDose
		out.Clamped = true
	}
	if final > rule.MaxDose {
		final = rule.MaxDose
		out.Clamped = true
	}
	if out.Clamped {
		basis += fmt.Sprintf(" · limitada pela trava da regra (%s a %s %s ao dia)",
			fmtDose(rule.MinDose), fmtDose(rule.MaxDose), in.Component.Unit)
	}

	// A regra fala do DIA; o campo que a tela preenche é a dose de UMA cápsula. Sem esta divisão,
	// uma regra de 5.000 UI/dia numa fórmula tomada duas vezes ao dia entregaria 10.000.
	tomadas := in.DosesPerDay
	if tomadas <= 0 {
		tomadas = 1
	}
	if tomadas > 1 {
		basis += fmt.Sprintf(" · %s %s ao dia ÷ %s tomadas",
			fmtDose(final), in.Component.Unit, fmtDose(tomadas))
		final = final / tomadas
	}

	// Arredondar por último: o que a farmácia pesa é a dose da cápsula, não a do dia.
	if rule.RoundTo != nil && *rule.RoundTo > 0 {
		rounded := math.Round(final/(*rule.RoundTo)) * *rule.RoundTo
		if rounded > 0 && math.Abs(rounded-final) > 1e-9 {
			basis += fmt.Sprintf(" · arredondada para %s %s", fmtDose(rounded), in.Component.Unit)
			final = rounded
		}
	}
	if strings.TrimSpace(rule.Note) != "" {
		basis += " · " + strings.TrimSpace(rule.Note)
	}

	out.Suggested = &final
	out.Basis = basis
	return out
}

// matchBand — intervalo MEIO-ABERTO (lower, upper], igual às faixas do escore. nil em lower é
// -infinito, nil em upper é +infinito.
func matchBand(value float64, bands []models.MagistralFormulaTemplateRuleBand) *models.MagistralFormulaTemplateRuleBand {
	for i := range bands {
		b := bands[i]
		if b.LowerBound != nil && value <= *b.LowerBound {
			continue
		}
		if b.UpperBound != nil && value > *b.UpperBound {
			continue
		}
		return &b
	}
	return nil
}

func bandViews(bands []models.MagistralFormulaTemplateRuleBand, active *models.MagistralFormulaTemplateRuleBand) []DoseBandView {
	out := make([]DoseBandView, 0, len(bands))
	for _, b := range bands {
		out = append(out, DoseBandView{
			LowerBound: b.LowerBound,
			UpperBound: b.UpperBound,
			Dose:       b.Dose,
			Label:      b.Label,
			Active:     active != nil && active.ID == b.ID,
		})
	}
	return out
}

func bandLabel(b models.MagistralFormulaTemplateRuleBand) string {
	switch {
	case b.LowerBound == nil && b.UpperBound != nil:
		return "≤ " + fmtDose(*b.UpperBound)
	case b.LowerBound != nil && b.UpperBound == nil:
		return "> " + fmtDose(*b.LowerBound)
	case b.LowerBound != nil && b.UpperBound != nil:
		return fmt.Sprintf("%s a %s", fmtDose(*b.LowerBound), fmtDose(*b.UpperBound))
	}
	return "todas"
}

func bandName(b models.MagistralFormulaTemplateRuleBand) string {
	if strings.TrimSpace(b.Label) == "" {
		return ""
	}
	return " (" + strings.TrimSpace(b.Label) + ")"
}

// unitMismatch — devolve o motivo quando a unidade do resultado não é a unidade em que a regra
// foi escrita, e string vazia quando pode seguir.
//
// Não é preciosismo: 390 dos 1.243 resultados numéricos do banco estão gravados numa unidade
// diferente da definição do exame. Há cortisol em nmol/L sobre definição em µg/dL (fator 27,6) e
// uma 25-OH-vitamina D gravada em pg/mL. Comparar o número cru com o limiar da regra nesses casos
// é errar a conduta com cara de acerto.
//
// A guarda só vale se não for barulhenta — por isso as variantes cosméticas do MESMO valor
// (mcg/L, µg/L, ug/L · µUI/mL, µU/mL, mUI/L) normalizam para igual e passam sem alarde.
func unitMismatch(ruleUnit *string, resultUnit string) string {
	// Regra sem unidade cadastrada: sem base para comparar, segue como antes. A tela de cadastro
	// preenche a unidade a partir do exame escolhido, então isto só alcança regra antiga.
	if ruleUnit == nil || strings.TrimSpace(*ruleUnit) == "" {
		return ""
	}
	want, got := canonicalUnit(*ruleUnit), canonicalUnit(resultUnit)
	if want == got {
		return ""
	}
	if got == "" {
		return fmt.Sprintf("resultado sem unidade registrada e a regra foi escrita em %s: confira o exame antes de dosar",
			strings.TrimSpace(*ruleUnit))
	}
	return fmt.Sprintf("resultado em %s e a regra foi escrita em %s: sem conversão, não sugere dose",
		strings.TrimSpace(resultUnit), strings.TrimSpace(*ruleUnit))
}

// canonicalUnit reduz uma unidade à forma canônica do seu VALOR: duas unidades numericamente
// iguais (ng/mL e µg/L) devolvem a mesma string; ng/mL e µg/dL, que diferem 10×, não.
func canonicalUnit(u string) string {
	s := strings.ToLower(strings.TrimSpace(u))
	s = strings.ReplaceAll(s, " ", "")
	// µ (micro), μ (mu grego) e o prefixo "mc" são a mesma coisa escrita de três jeitos.
	s = strings.NewReplacer("µ", "u", "μ", "u", "mcg", "ug", "mcmol", "umol", "mcl", "ul").Replace(s)
	// International Unit e Unidade Internacional.
	s = strings.ReplaceAll(s, "iu", "ui")
	if canon, ok := unitEquivalents[s]; ok {
		return canon
	}
	return s
}

// Equivalências exatas — só entram pares cujo valor numérico é idêntico.
var unitEquivalents = map[string]string{
	"ng/ml":  "ug/l", // 1 ng/mL = 1 µg/L
	"ug/l":   "ug/l",
	"ug/ml":  "mg/l", // 1 µg/mL = 1 mg/L
	"mg/l":   "mg/l",
	"mg/ml":  "g/l",
	"g/l":    "g/l",
	"uui/ml": "mui/l", // 1 µUI/mL = 1 mUI/L
	"mui/l":  "mui/l",
	"uu/ml":  "mui/l", // insulina: µU/mL é como µUI/mL
}

func compare(value float64, operator string, threshold float64) bool {
	switch operator {
	case "lt":
		return value < threshold
	case "lte":
		return value <= threshold
	case "gt":
		return value > threshold
	case "gte":
		return value >= threshold
	}
	return false
}

func operatorLabel(operator string) string {
	switch operator {
	case "lt":
		return "<"
	case "lte":
		return "≤"
	case "gt":
		return ">"
	case "gte":
		return "≥"
	}
	return operator
}

// isStale — sem data, o dado é tratado como velho: não dá para afirmar que é recente.
func isStale(date *time.Time, maxAgeDays int, now time.Time) (bool, string) {
	if date == nil {
		return true, "data desconhecida"
	}
	if maxAgeDays <= 0 {
		return false, ""
	}
	age := int(now.Sub(*date).Hours() / 24)
	if age > maxAgeDays {
		return true, fmt.Sprintf("%d dias", age)
	}
	return false, ""
}

func describeSource(m *MeasurementInput) string {
	if m == nil || strings.TrimSpace(m.Source) == "" {
		return "dado do prontuário"
	}
	return m.Source
}

func dateSuffix(date *time.Time) string {
	if date == nil {
		return ""
	}
	return " de " + date.Format("02/01")
}

// fmtDose — pt-BR, sem zeros à toa: 5000 → "5.000"; 0,25 → "0,25".
func fmtDose(v float64) string {
	if v == math.Trunc(v) && math.Abs(v) >= 1000 {
		s := fmt.Sprintf("%.0f", v)
		var parts []string
		for len(s) > 3 {
			parts = append([]string{s[len(s)-3:]}, parts...)
			s = s[:len(s)-3]
		}
		parts = append([]string{s}, parts...)
		return strings.Join(parts, ".")
	}
	s := strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.4f", v), "0"), ".")
	if s == "" {
		s = "0"
	}
	return strings.Replace(s, ".", ",", 1)
}
