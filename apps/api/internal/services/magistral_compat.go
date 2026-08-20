package services

import (
	"fmt"
	"sort"
	"strings"

	"github.com/google/uuid"

	"github.com/plenya/api/internal/models"
)

// Compatibilidade e palatabilidade de fórmula magistral.
//
// Dois mecanismos, porque nenhum cobre sozinho:
//
//   1. REGRAS DERIVADAS DE FLAG — valem para qualquer par que tenha a característica (eutético
//      com eutético, oxidante com oxidável, higroscópico em cápsula gelatinosa). Escalam sem
//      cadastro nenhum.
//   2. PARES CURADOS — o que não sai de flag, com mecanismo escrito à mão.
//
// A apresentação é a mesma do alerta de alergia que já existe: AVISA, NÃO BLOQUEIA. Um conjunto
// pequeno que nunca dá falso positivo constrói confiança; um grande e barulhento vira ruído
// ignorado em duas semanas — por isso as regras aqui são poucas e literais.

// MagistralAlertLevel espelha a severidade das incompatibilidades curadas.
type MagistralAlertLevel string

const (
	AlertInfo  MagistralAlertLevel = "info"
	AlertWarn  MagistralAlertLevel = "warn"
	AlertAvoid MagistralAlertLevel = "avoid"
)

// MagistralAlert é um aviso sobre a fórmula.
type MagistralAlert struct {
	Level      MagistralAlertLevel `json:"level"`
	Kind       string              `json:"kind"` // eutectic, oxidation, hygroscopic, palatability, pair
	Substances []string            `json:"substances"`
	Message    string              `json:"message"`
}

// FormulaCheckInput — a fórmula como está na tela, já resolvida contra o catálogo.
type FormulaCheckInput struct {
	PharmaceuticalForm string
	Components         []FormulaCheckComponent
	// Quantas tomadas por dia a posologia manda. Zero é tratado como 1.
	DosesPerDay float64
	// Tetos do Anexo IV da IN 28, por nutriente. Vazio = a conferência não roda.
	In28 map[string]models.In28Limit
	// Veículo escrito na fórmula, para conferir contra as regras de base.
	Vehicle string
	// Regras de incompatibilidade com a base.
	BaseRules []models.MagistralBaseIncompatibility
}

// doseNaUnidadeDoCatalogo põe a quantidade escrita na receita na unidade em que o catálogo
// raciocina. Devolve false quando não dá para converter — aí a conferência é pulada, que é melhor
// do que comparar grandezas diferentes.
func doseNaUnidadeDoCatalogo(c FormulaCheckComponent) (float64, bool) {
	if c.Catalog == nil {
		return 0, false
	}
	if mesmaUnidade(c.Unit, c.Catalog.DefaultUnit) {
		return c.Quantity, true
	}
	return ConverteDose(c.Quantity, c.Unit, c.Catalog.DefaultUnit)
}

// nomeLimpoIN28 tira a letra de nota de rodapé que a norma cola no nome ("Cálciov",
// "Vitamina Ai"): ela é referência do rodapé do Anexo IV e não faz sentido na tela.
func nomeLimpoIN28(n string) string {
	for _, sufixo := range []string{" iv", " ix", "iii", "ii", "i", "v"} {
		if strings.HasSuffix(n, sufixo) && len(n) > len(sufixo)+3 {
			return strings.TrimSpace(strings.TrimSuffix(n, sufixo))
		}
	}
	return n
}

// componentesQueCasam devolve os componentes alcançados pela regra de base. Regra sem padrão de
// substância vale para a fórmula toda; regra com percentual mínimo só dispara acima dele, e só
// quando a quantidade está escrita EM PORCENTAGEM — comparar 10 mg com "10%" seria inventar.
func componentesQueCasam(comps []FormulaCheckComponent, r models.MagistralBaseIncompatibility) []string {
	var out []string
	for _, c := range comps {
		if r.SubstancePattern != nil {
			if !strings.Contains(normalizaTexto(c.Substance), normalizaTexto(*r.SubstancePattern)) {
				continue
			}
		}
		if r.MinPercent != nil {
			if strings.TrimSpace(c.Unit) != "%" || c.Quantity < *r.MinPercent {
				continue
			}
		}
		out = append(out, c.Substance)
	}
	return out
}

// normalizaTexto deixa em caixa baixa e sem acento, para casar "não iônico" com "nao ionico".
func normalizaTexto(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	trocas := strings.NewReplacer(
		"á", "a", "à", "a", "ã", "a", "â", "a",
		"é", "e", "ê", "e", "í", "i", "ó", "o", "ô", "o", "õ", "o",
		"ú", "u", "ü", "u", "ç", "c")
	return trocas.Replace(s)
}

func formaSublingualOuTransdermica(forma string) bool {
	f := strings.ToLower(strings.TrimSpace(forma))
	for _, v := range []string{"sublingual", "transdérmic", "transdermic", "adesivo"} {
		if strings.Contains(f, v) {
			return true
		}
	}
	return false
}

// porTomada explicita a conta quando a comparação somou mais de uma tomada: sem isso o médico lê
// "abaixo de 600 mg" olhando para uma cápsula de 400 e não entende que já foram contadas duas.
func porTomada(tomadas float64, base string) string {
	if base != "diária" || tomadas <= 1 {
		return ""
	}
	return fmt.Sprintf(", com %s tomadas ao dia", fmtML(tomadas))
}

// FormulaCheckComponent — um componente e o que o catálogo sabe dele (nil = não catalogado).
type FormulaCheckComponent struct {
	Substance string
	Quantity  float64
	Unit      string
	Catalog   *models.MagistralComponent
	// AsElemental: a dose escrita é do elemento, não do insumo.
	AsElemental bool
}

// isCapsule / isSachet — a forma farmacêutica muda quais regras se aplicam.
func isCapsule(form string) bool { return strings.Contains(strings.ToLower(form), "cápsula") }
func isSachet(form string) bool {
	f := strings.ToLower(form)
	return strings.Contains(f, "sachê") || strings.Contains(f, "sache")
}

// CheckFormula devolve os avisos da fórmula, do mais grave para o menos.
func CheckFormula(in FormulaCheckInput, pairs []models.MagistralIncompatibility) []MagistralAlert {
	var alerts []MagistralAlert

	byID := map[uuid.UUID]FormulaCheckComponent{}
	for _, c := range in.Components {
		if c.Catalog != nil {
			byID[c.Catalog.ID] = c
		}
	}

	// 1) Eutéticos: duas substâncias formadoras de mistura eutética viram líquido/pastoso ao
	// serem trituradas juntas (mentol + cânfora é o caso clássico).
	var eutectics []string
	for _, c := range in.Components {
		if c.Catalog != nil && c.Catalog.EutecticFormer {
			eutectics = append(eutectics, c.Substance)
		}
	}
	if len(eutectics) >= 2 {
		sort.Strings(eutectics)
		alerts = append(alerts, MagistralAlert{
			Level:      AlertWarn,
			Kind:       "eutectic",
			Substances: eutectics,
			Message: "Mistura eutética provável entre " + strings.Join(eutectics, " e ") +
				": juntas podem liquefazer na trituração. A farmácia precisa de adsorvente ou de separar as fases.",
		})
	}

	// 2) Oxidante junto de substância sensível à oxidação.
	var oxidizers, sensitive []string
	for _, c := range in.Components {
		if c.Catalog == nil {
			continue
		}
		if c.Catalog.Oxidizing {
			oxidizers = append(oxidizers, c.Substance)
		}
		if c.Catalog.OxidationSensitive {
			sensitive = append(sensitive, c.Substance)
		}
	}
	if len(oxidizers) > 0 && len(sensitive) > 0 {
		alerts = append(alerts, MagistralAlert{
			Level:      AlertWarn,
			Kind:       "oxidation",
			Substances: append(append([]string{}, oxidizers...), sensitive...),
			Message: "Componente oxidante (" + strings.Join(oxidizers, ", ") + ") junto de " +
				strings.Join(sensitive, ", ") + ": risco de degradação. Considere separar em fórmulas diferentes.",
		})
	}

	// 3) Higroscópico em cápsula gelatinosa: puxa umidade da própria cápsula e a amolece.
	if isCapsule(in.PharmaceuticalForm) {
		var hygroscopic []string
		for _, c := range in.Components {
			if c.Catalog != nil && c.Catalog.Hygroscopic {
				hygroscopic = append(hygroscopic, c.Substance)
			}
		}
		if len(hygroscopic) > 0 {
			alerts = append(alerts, MagistralAlert{
				Level:      AlertInfo,
				Kind:       "hygroscopic",
				Substances: hygroscopic,
				Message: strings.Join(hygroscopic, ", ") +
					" é higroscópico: em cápsula gelatinosa pode amolecer o invólucro. Cápsula vegetal ou sachê contornam.",
			})
		}
	}

	// 4) Palatabilidade: amargor alto em sachê é a queixa que faz o paciente abandonar a fórmula.
	if isSachet(in.PharmaceuticalForm) {
		var bitter []string
		for _, c := range in.Components {
			if c.Catalog == nil {
				continue
			}
			if (c.Catalog.Bitterness != nil && *c.Catalog.Bitterness >= 2) ||
				(c.Catalog.SachetOK != nil && !*c.Catalog.SachetOK) {
				bitter = append(bitter, c.Substance)
			}
		}
		if len(bitter) > 0 {
			alerts = append(alerts, MagistralAlert{
				Level:      AlertWarn,
				Kind:       "palatability",
				Substances: bitter,
				Message: strings.Join(bitter, ", ") +
					" tem sabor marcante: em sachê costuma ser recusado. Cápsula resolve.",
			})
		}
	}

	// 5) Pares curados.
	for _, p := range pairs {
		a, okA := byID[p.ComponentAID]
		b, okB := byID[p.ComponentBID]
		if !okA || !okB {
			continue
		}
		msg := fmt.Sprintf("%s com %s", a.Substance, b.Substance)
		if p.Mechanism != "" {
			msg += ": " + p.Mechanism
		}
		if p.Note != nil && strings.TrimSpace(*p.Note) != "" {
			msg += " " + strings.TrimSpace(*p.Note)
		}
		alerts = append(alerts, MagistralAlert{
			Level:      MagistralAlertLevel(p.Severity),
			Kind:       "pair",
			Substances: []string{a.Substance, b.Substance},
			Message:    msg,
		})
	}

	// 6) Fora da faixa cadastrada. Só fala quando HÁ faixa: catálogo sem dose não vira alerta.
	//
	// A comparação é feita na base em que a faixa foi escrita. Faixa diária contra dose de uma
	// cápsula produzia alerta em fórmula certa e silêncio em fórmula errada, dependendo só de
	// como aquela linha do catálogo tinha sido preenchida.
	tomadas := in.DosesPerDay
	if tomadas <= 0 {
		tomadas = 1
	}
	for _, c := range in.Components {
		if c.Catalog == nil {
			continue
		}
		// A dose vem para a unidade do catálogo antes de comparar: "metilcobalamina 0,5 mg"
		// contra faixa em mcg é a mesma dose, e antes era simplesmente ignorada.
		quantidade, conversivel := doseNaUnidadeDoCatalogo(c)
		if !conversivel {
			continue
		}
		// Faixa do catálogo é de via oral. Sublingual e transdérmico pulam a primeira passagem
		// hepática e usam dose menor de propósito — acusar "abaixo da faixa" ali é acusar a via,
		// não a dose. O lado de cima continua valendo: dose alta por qualquer via merece conferência.
		viaComAbsorcaoPropria := formaSublingualOuTransdermica(in.PharmaceuticalForm)

		dose, base := quantidade, "por tomada"
		if c.Catalog.DoseBasis != "por_dose" {
			dose, base = quantidade*tomadas, "diária"
		}
		if c.Catalog.MaxDose != nil && dose > *c.Catalog.MaxDose {
			alerts = append(alerts, MagistralAlert{
				Level:      AlertInfo,
				Kind:       "dose",
				Substances: []string{c.Substance},
				Message: fmt.Sprintf("%s acima da faixa %s cadastrada (%s %s%s). Confirme a intenção.",
					c.Substance, base, fmtML(*c.Catalog.MaxDose), c.Catalog.DefaultUnit,
					porTomada(tomadas, base)),
			})
		} else if c.Catalog.MinDose != nil && dose < *c.Catalog.MinDose && !viaComAbsorcaoPropria {
			alerts = append(alerts, MagistralAlert{
				Level:      AlertInfo,
				Kind:       "dose",
				Substances: []string{c.Substance},
				Message: fmt.Sprintf("%s abaixo da faixa %s cadastrada (%s %s%s).",
					c.Substance, base, fmtML(*c.Catalog.MinDose), c.Catalog.DefaultUnit,
					porTomada(tomadas, base)),
			})
		}
	}

	// 6b) Teto de suplemento alimentar da IN 28, somado POR NUTRIENTE.
	//
	// Somar por nutriente, e não por substância, é o ponto: foi assim que o formulário das
	// parceiras acumulou 107 mg de vitamina B6 somando piridoxal-5-fosfato com "vitamina B6" na
	// mesma cápsula, cada um parecendo comportado sozinho.
	//
	// O alerta é INFORMATIVO e sai como um só, listando tudo. Passar do teto da IN 28 é
	// exatamente o que separa um suplemento de uma prescrição — B12 de 1.000 mcg fica 100× acima
	// e é conduta corriqueira. Tratar isso como erro faria o médico parar de ler o painel.
	if len(in.In28) > 0 {
		porNutriente := map[string]float64{}
		for _, c := range in.Components {
			if c.Catalog == nil || c.Catalog.In28Nutrient == nil {
				continue
			}
			// O fator da IN 28 é calibrado na unidade do CATÁLOGO: a dose escrita na receita
			// precisa chegar nela antes de ser multiplicada.
			base, ok := doseNaUnidadeDoCatalogo(c)
			if !ok {
				continue
			}
			qtd := base * tomadas
			// Dose do insumo vira dose do elemento antes de comparar com a norma, que fala do
			// nutriente. Dose já declarada como do elemento entra como está.
			if !c.AsElemental && c.Catalog.ElementalPercent != nil && *c.Catalog.ElementalPercent > 0 {
				qtd = qtd * *c.Catalog.ElementalPercent / 100
			}
			fator := c.Catalog.In28Factor
			if fator <= 0 {
				fator = 1
			}
			porNutriente[*c.Catalog.In28Nutrient] += qtd * fator
		}

		var acima []string
		var nomes []string
		for nutriente, total := range porNutriente {
			limite, ok := in.In28[nutriente]
			if !ok || limite.MaxAdult == nil || total <= *limite.MaxAdult {
				continue
			}
			acima = append(acima, fmt.Sprintf("%s %s %s (teto %s)",
				nomeLimpoIN28(nutriente), fmtML(total), limite.Unit, fmtML(*limite.MaxAdult)))
			nomes = append(nomes, nomeLimpoIN28(nutriente))
		}
		if len(acima) > 0 {
			sort.Strings(acima)
			sort.Strings(nomes)
			alerts = append(alerts, MagistralAlert{
				Level:      AlertInfo,
				Kind:       "in28",
				Substances: nomes,
				Message: fmt.Sprintf("No dia, acima do teto de suplemento alimentar da IN 28: %s. "+
					"O teto é de suplemento, não de prescrição: em fórmula magistral isso é decisão sua, "+
					"e é o que separa um suplemento de uma receita.", strings.Join(acima, "; ")),
			})
		}
	}

	// 6c) Incompatibilidade com a BASE.
	//
	// É a maior fatia do problema real e a que faltava: no levantamento da farmácia-escola da
	// UFRJ, 63% dos erros farmacotécnicos são ativo × formulação e 23% ativo × base, contra 13%
	// ativo × ativo. Casa por texto porque o veículo é campo livre.
	veiculo := normalizaTexto(in.Vehicle)
	if veiculo != "" {
		for _, r := range in.BaseRules {
			if !r.IsActive || !strings.Contains(veiculo, normalizaTexto(r.BasePattern)) {
				continue
			}
			atingidos := componentesQueCasam(in.Components, r)
			if len(atingidos) == 0 {
				continue
			}
			msg := r.Mechanism
			if strings.TrimSpace(r.Recommendation) != "" {
				msg += " " + strings.TrimSpace(r.Recommendation)
			}
			alerts = append(alerts, MagistralAlert{
				Level:      MagistralAlertLevel(r.Severity),
				Kind:       "base",
				Substances: atingidos,
				Message:    fmt.Sprintf("%s com %s: %s", strings.Join(atingidos, ", "), strings.TrimSpace(in.Vehicle), msg),
			})
		}
	}

	// 6d) Substância que atrapalha exame.
	//
	// Biotina a partir de 5 mg/dia é o caso clássico: interfere em imunoensaio biotinilado e
	// devolve TSH, troponina e hormônio falsamente alto ou baixo conforme o formato do ensaio.
	// Num sistema que prescreve E lê exame, isso fecha um ciclo ruim — a regra de dose dinâmica
	// leria um número que não é do paciente.
	for _, c := range in.Components {
		if c.Catalog == nil || c.Catalog.AssayInterference == nil ||
			strings.TrimSpace(*c.Catalog.AssayInterference) == "" {
			continue
		}
		// O limiar está na unidade do catálogo; a receita pode estar em outra.
		base, ok := doseNaUnidadeDoCatalogo(c)
		if !ok {
			continue
		}
		dose := base * tomadas
		if c.Catalog.AssayInterferenceDose != nil && dose < *c.Catalog.AssayInterferenceDose {
			continue
		}
		alerts = append(alerts, MagistralAlert{
			Level:      AlertWarn,
			Kind:       "assay",
			Substances: []string{c.Substance},
			Message: fmt.Sprintf("%s %s %s ao dia atrapalha exame: %s",
				c.Substance, fmtML(dose), c.Catalog.DefaultUnit, strings.TrimSpace(*c.Catalog.AssayInterference)),
		})
	}

	// 7) Insumo diluído ou quelado sem dizer se a dose é do elemento. "Magnésio quelato 300 mg"
	// pode significar 300 mg de elemento (1 g de bisglicinato) ou 300 mg do quelato (90 mg de
	// elemento). A farmácia precisa da resposta, e a diferença é de mais de três vezes.
	for _, c := range in.Components {
		if c.Catalog == nil || c.Catalog.ElementalPercent == nil || c.AsElemental {
			continue
		}
		alerts = append(alerts, MagistralAlert{
			Level:      AlertInfo,
			Kind:       "correction",
			Substances: []string{c.Substance},
			Message: fmt.Sprintf(
				"%s tem %s%% de ativo no insumo. A dose está sendo entendida como do insumo; marque \"dose do elemento\" se a intenção for outra.",
				c.Substance, fmtML(*c.Catalog.ElementalPercent)),
		})
	}

	sort.SliceStable(alerts, func(i, j int) bool {
		return alertRank(alerts[i].Level) > alertRank(alerts[j].Level)
	})
	return alerts
}

func alertRank(l MagistralAlertLevel) int {
	switch l {
	case AlertAvoid:
		return 2
	case AlertWarn:
		return 1
	default:
		return 0
	}
}
