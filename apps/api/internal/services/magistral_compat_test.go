package services

import (
	"strings"
	"testing"

	"github.com/google/uuid"

	"github.com/plenya/api/internal/models"
)

func cat(name string, mut func(*models.MagistralComponent)) *models.MagistralComponent {
	c := &models.MagistralComponent{ID: uuid.Must(uuid.NewV7()), Name: name, DefaultUnit: "mg"}
	if mut != nil {
		mut(c)
	}
	return c
}

func alertKinds(alerts []MagistralAlert) []string {
	out := make([]string, 0, len(alerts))
	for _, a := range alerts {
		out = append(out, a.Kind)
	}
	return out
}

func TestCheckFormulaEutectic(t *testing.T) {
	mentol := cat("Mentol", func(c *models.MagistralComponent) { c.EutecticFormer = true })
	canfora := cat("Cânfora", func(c *models.MagistralComponent) { c.EutecticFormer = true })

	alerts := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "creme",
		Components: []FormulaCheckComponent{
			{Substance: "Mentol", Quantity: 2, Unit: "%", Catalog: mentol},
			{Substance: "Cânfora", Quantity: 2, Unit: "%", Catalog: canfora},
		},
	}, nil)

	if len(alerts) != 1 || alerts[0].Kind != "eutectic" {
		t.Fatalf("esperava 1 alerta de eutético, veio %v", alertKinds(alerts))
	}

	// Um formador sozinho não é alerta — falso positivo é o que faz o médico parar de ler.
	solo := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "creme",
		Components:         []FormulaCheckComponent{{Substance: "Mentol", Catalog: mentol}},
	}, nil)
	if len(solo) != 0 {
		t.Errorf("um formador de eutético sozinho não deveria alertar: %v", alertKinds(solo))
	}
}

func TestCheckFormulaSemCatalogoNaoInventa(t *testing.T) {
	alerts := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "cápsula",
		Components: []FormulaCheckComponent{
			{Substance: "Substância desconhecida", Quantity: 100, Unit: "mg"},
			{Substance: "Outra desconhecida", Quantity: 50, Unit: "mg"},
		},
	}, nil)
	if len(alerts) != 0 {
		t.Errorf("sem catálogo não há o que afirmar, veio %v", alertKinds(alerts))
	}
}

func TestCheckFormulaHigroscopicoEPalatabilidade(t *testing.T) {
	bitterness := 3
	sachetNo := false
	mg := cat("Cloreto de magnésio", func(c *models.MagistralComponent) {
		c.Hygroscopic = true
		c.Bitterness = &bitterness
		c.SachetOK = &sachetNo
	})

	emCapsula := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "cápsula",
		Components:         []FormulaCheckComponent{{Substance: mg.Name, Catalog: mg}},
	}, nil)
	if len(emCapsula) != 1 || emCapsula[0].Kind != "hygroscopic" {
		t.Errorf("higroscópico em cápsula deveria alertar, veio %v", alertKinds(emCapsula))
	}

	emSache := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "sachê",
		Components:         []FormulaCheckComponent{{Substance: mg.Name, Catalog: mg}},
	}, nil)
	if len(emSache) != 1 || emSache[0].Kind != "palatability" {
		t.Errorf("amargor alto em sachê deveria alertar, veio %v", alertKinds(emSache))
	}
}

func TestCheckFormulaParCurado(t *testing.T) {
	a := cat("Ácido ascórbico", nil)
	b := cat("Cianocobalamina", nil)
	note := "Separar em fórmulas distintas."
	pair := models.MagistralIncompatibility{
		ComponentAID: a.ID, ComponentBID: b.ID,
		Severity:  models.IncompatAvoid,
		Mechanism: "o ácido ascórbico degrada a cianocobalamina",
		Note:      &note,
	}

	alerts := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "cápsula",
		Components: []FormulaCheckComponent{
			{Substance: a.Name, Catalog: a},
			{Substance: b.Name, Catalog: b},
		},
	}, []models.MagistralIncompatibility{pair})

	if len(alerts) != 1 || alerts[0].Level != AlertAvoid {
		t.Fatalf("par curado deveria alertar como avoid, veio %v", alerts)
	}
	if !strings.Contains(alerts[0].Message, "degrada") || !strings.Contains(alerts[0].Message, "Separar") {
		t.Errorf("mecanismo e nota precisam aparecer: %q", alerts[0].Message)
	}

	// O mesmo par com só um dos dois na fórmula não alerta.
	solo := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "cápsula",
		Components:         []FormulaCheckComponent{{Substance: a.Name, Catalog: a}},
	}, []models.MagistralIncompatibility{pair})
	if len(solo) != 0 {
		t.Errorf("par incompleto não deveria alertar: %v", alertKinds(solo))
	}
}

func TestCheckFormulaFaixaDeDose(t *testing.T) {
	max := 10.0
	min := 1.0
	mel := cat("Melatonina", func(c *models.MagistralComponent) {
		c.MinDose = &min
		c.MaxDose = &max
	})

	acima := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "cápsula",
		Components:         []FormulaCheckComponent{{Substance: mel.Name, Quantity: 20, Unit: "mg", Catalog: mel}},
	}, nil)
	if len(acima) != 1 || acima[0].Kind != "dose" {
		t.Errorf("dose acima da faixa deveria avisar, veio %v", alertKinds(acima))
	}

	dentro := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "cápsula",
		Components:         []FormulaCheckComponent{{Substance: mel.Name, Quantity: 3, Unit: "mg", Catalog: mel}},
	}, nil)
	if len(dentro) != 0 {
		t.Errorf("dose dentro da faixa não alerta: %v", alertKinds(dentro))
	}

	// Unidade diferente da cadastrada não vira comparação numérica errada.
	outraUnidade := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "cápsula",
		Components:         []FormulaCheckComponent{{Substance: mel.Name, Quantity: 3000, Unit: "mcg", Catalog: mel}},
	}, nil)
	if len(outraUnidade) != 0 {
		t.Errorf("3000 mcg = 3 mg não pode virar alerta de dose alta: %v", outraUnidade)
	}
}

// Insumo diluído ou quelado prescrito sem dizer se a dose é do elemento: a farmácia precisa da
// resposta, e a diferença chega a ser de mais de três vezes.
func TestCheckFormulaFatorDeCorrecao(t *testing.T) {
	pct := 30.0
	mg := cat("Magnésio quelato", func(c *models.MagistralComponent) { c.ElementalPercent = &pct })

	semDizer := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "cápsula",
		Components:         []FormulaCheckComponent{{Substance: mg.Name, Quantity: 300, Unit: "mg", Catalog: mg}},
	}, nil)
	if len(semDizer) != 1 || semDizer[0].Kind != "correction" {
		t.Fatalf("deveria pedir para esclarecer elemento x insumo, veio %v", alertKinds(semDizer))
	}
	if !strings.Contains(semDizer[0].Message, "30") {
		t.Errorf("a mensagem precisa trazer o percentual: %q", semDizer[0].Message)
	}

	declarado := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "cápsula",
		Components: []FormulaCheckComponent{
			{Substance: mg.Name, Quantity: 300, Unit: "mg", Catalog: mg, AsElemental: true},
		},
	}, nil)
	if len(declarado) != 0 {
		t.Errorf("com a dose declarada como do elemento, não há o que apontar: %v", alertKinds(declarado))
	}

	// substância sem percentual cadastrado não vira alerta
	simples := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "cápsula",
		Components:         []FormulaCheckComponent{{Substance: "Melatonina", Quantity: 3, Unit: "mg", Catalog: cat("Melatonina", nil)}},
	}, nil)
	if len(simples) != 0 {
		t.Errorf("sem percentual cadastrado não há alerta: %v", alertKinds(simples))
	}
}

// A faixa do catálogo é DIÁRIA por padrão. Comparar a dose de uma cápsula contra faixa diária
// sem contar as tomadas acusava fórmula certa e absolvia fórmula errada — foi o que aconteceu com
// as 7 substâncias cuja faixa numérica tinha sido semeada da dose por cápsula das parceiras.
func TestCheckFormulaFaixaDiariaContaAsTomadas(t *testing.T) {
	min, max := 600.0, 1800.0
	nac := cat("N-acetilcisteína", func(c *models.MagistralComponent) {
		c.MinDose = &min
		c.MaxDose = &max
		c.DoseBasis = "por_dia"
	})

	// 400 mg por cápsula, 2 vezes ao dia = 800 mg/dia: dentro da faixa, não alerta.
	duas := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "cápsula",
		DosesPerDay:        2,
		Components:         []FormulaCheckComponent{{Substance: nac.Name, Quantity: 400, Unit: "mg", Catalog: nac}},
	}, nil)
	if len(duas) != 0 {
		t.Errorf("800 mg/dia está dentro de 600 a 1800: não devia alertar, veio %v", duas)
	}

	// A mesma cápsula uma vez ao dia são 400 mg/dia: aí sim está abaixo, e a frase precisa dizer
	// que a faixa é diária.
	uma := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "cápsula",
		DosesPerDay:        1,
		Components:         []FormulaCheckComponent{{Substance: nac.Name, Quantity: 400, Unit: "mg", Catalog: nac}},
	}, nil)
	if len(uma) != 1 || uma[0].Kind != "dose" {
		t.Fatalf("400 mg/dia está abaixo de 600: devia alertar, veio %v", alertKinds(uma))
	}
	if !strings.Contains(uma[0].Message, "diária") {
		t.Errorf("a frase precisa dizer que a faixa é diária, veio %q", uma[0].Message)
	}

	// Quando somou mais de uma tomada, a conta aparece: senão o médico lê "abaixo de 600" olhando
	// para uma cápsula de 400 e não entende que já foram contadas duas.
	acima := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "cápsula",
		DosesPerDay:        3,
		Components:         []FormulaCheckComponent{{Substance: nac.Name, Quantity: 700, Unit: "mg", Catalog: nac}},
	}, nil)
	if len(acima) != 1 || !strings.Contains(acima[0].Message, "3 tomadas ao dia") {
		t.Errorf("a frase precisa mostrar as tomadas contadas, veio %v", acima)
	}
}

// Substância cuja faixa É por tomada não pode ser multiplicada pelas tomadas do dia.
func TestCheckFormulaFaixaPorTomadaNaoMultiplica(t *testing.T) {
	min, max := 1.0, 10.0
	mel := cat("Melatonina", func(c *models.MagistralComponent) {
		c.MinDose = &min
		c.MaxDose = &max
		c.DoseBasis = "por_dose"
	})
	got := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "cápsula",
		DosesPerDay:        3,
		Components:         []FormulaCheckComponent{{Substance: mel.Name, Quantity: 5, Unit: "mg", Catalog: mel}},
	}, nil)
	if len(got) != 0 {
		t.Errorf("5 mg por tomada está dentro de 1 a 10 por tomada: não devia alertar, veio %v", got)
	}
}

// A faixa do catálogo é de via oral. Sublingual usa dose menor de propósito — o alerta de dose
// baixa ali acusa a via, não a dose. O lado de cima continua valendo.
func TestCheckFormulaSublingualNaoAcusaDoseBaixa(t *testing.T) {
	min, max := 50.0, 300.0
	htp := cat("5-HTP", func(c *models.MagistralComponent) {
		c.MinDose = &min
		c.MaxDose = &max
	})

	sub := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "sublingual",
		Components:         []FormulaCheckComponent{{Substance: htp.Name, Quantity: 25, Unit: "mg", Catalog: htp}},
	}, nil)
	if len(sub) != 0 {
		t.Errorf("25 mg sublingual é dose de sublingual, não devia alertar: %v", sub)
	}

	oral := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "cápsula",
		Components:         []FormulaCheckComponent{{Substance: htp.Name, Quantity: 25, Unit: "mg", Catalog: htp}},
	}, nil)
	if len(oral) != 1 || oral[0].Kind != "dose" {
		t.Errorf("os mesmos 25 mg em cápsula estão abaixo da faixa oral: %v", alertKinds(oral))
	}

	alto := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "sublingual",
		Components:         []FormulaCheckComponent{{Substance: htp.Name, Quantity: 400, Unit: "mg", Catalog: htp}},
	}, nil)
	if len(alto) != 1 {
		t.Errorf("dose alta por qualquer via merece conferência: %v", alertKinds(alto))
	}
}

// O teto da IN 28 se confere por NUTRIENTE, não por substância. Foi somando piridoxal-5-fosfato
// com "vitamina B6" que o formulário das parceiras chegou a 107 mg/dia de B6 com cada linha
// parecendo comportada sozinha.
func TestCheckFormulaTetoIN28SomaPorNutriente(t *testing.T) {
	b6 := "Vitamina B6"
	teto := 98.6
	limites := map[string]models.In28Limit{b6: {Nutrient: b6, Unit: "mg", MaxAdult: &teto}}

	p5p := cat("Piridoxal-5-fosfato", func(c *models.MagistralComponent) {
		c.In28Nutrient = &b6
		c.In28Factor = 1
	})
	piridoxina := cat("Vitamina B6", func(c *models.MagistralComponent) {
		c.In28Nutrient = &b6
		c.In28Factor = 1
	})

	juntos := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "cápsula",
		DosesPerDay:        1,
		In28:               limites,
		Components: []FormulaCheckComponent{
			{Substance: p5p.Name, Quantity: 7, Unit: "mg", Catalog: p5p},
			{Substance: piridoxina.Name, Quantity: 100, Unit: "mg", Catalog: piridoxina},
		},
	}, nil)
	var achou *MagistralAlert
	for i := range juntos {
		if juntos[i].Kind == "in28" {
			achou = &juntos[i]
		}
	}
	if achou == nil {
		t.Fatalf("7 + 100 mg de B6 passam do teto de 98,6: devia avisar, veio %v", alertKinds(juntos))
	}
	if !strings.Contains(achou.Message, "107") {
		t.Errorf("a soma precisa aparecer na frase, veio %q", achou.Message)
	}
	// É aviso, não impedimento: passar do teto é o que diferencia prescrição de suplemento.
	if achou.Level != AlertInfo {
		t.Errorf("o teto de suplemento não bloqueia prescrição, veio nível %v", achou.Level)
	}

	// Cada um sozinho fica abaixo e não alerta.
	sozinho := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "cápsula", DosesPerDay: 1, In28: limites,
		Components: []FormulaCheckComponent{{Substance: p5p.Name, Quantity: 7, Unit: "mg", Catalog: p5p}},
	}, nil)
	for _, a := range sozinho {
		if a.Kind == "in28" {
			t.Errorf("7 mg de B6 estão muito abaixo do teto: %q", a.Message)
		}
	}
}

// Mineral prescrito como insumo vira elemento antes de encostar na norma: 1 g de bisglicinato a
// 30% são 300 mg de magnésio, abaixo do teto de 350 — comparar o peso do insumo acusaria errado.
func TestCheckFormulaTetoIN28ConverteParaElemento(t *testing.T) {
	mg := "Magnésio"
	teto := 350.0
	limites := map[string]models.In28Limit{mg: {Nutrient: mg, Unit: "mg", MaxAdult: &teto}}
	pct := 30.0
	quelato := cat("Magnésio quelato", func(c *models.MagistralComponent) {
		c.In28Nutrient = &mg
		c.In28Factor = 1
		c.ElementalPercent = &pct
	})

	got := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "cápsula", DosesPerDay: 1, In28: limites,
		Components: []FormulaCheckComponent{{Substance: quelato.Name, Quantity: 1000, Unit: "mg", Catalog: quelato}},
	}, nil)
	for _, a := range got {
		if a.Kind == "in28" {
			t.Errorf("1 g de bisglicinato são 300 mg de magnésio, abaixo do teto: %q", a.Message)
		}
	}
}

// Incompatibilidade com a BASE é a fatia maior do problema real (63% ativo × formulação e 23%
// ativo × base, contra 13% ativo × ativo no levantamento da farmácia-escola da UFRJ) e era a que
// o modelo não representava.
func TestCheckFormulaIncompatibilidadeDeBase(t *testing.T) {
	pct := 10.0
	pat := "ácido"
	regra := models.MagistralBaseIncompatibility{
		BasePattern: "lanette", SubstancePattern: &pat, MinPercent: &pct,
		Severity: models.IncompatWarn, IsActive: true,
		Mechanism:      "o creme Lanette é aniônico e o ácido neutraliza o tensoativo.",
		Recommendation: "Creme não iônico.",
	}
	acido := cat("Ácido glicólico", nil)

	// 15% em Lanette: dispara.
	alto := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "creme", Vehicle: "Creme Lanette qsp 100 g",
		BaseRules:  []models.MagistralBaseIncompatibility{regra},
		Components: []FormulaCheckComponent{{Substance: acido.Name, Quantity: 15, Unit: "%", Catalog: acido}},
	}, nil)
	if len(alto) != 1 || alto[0].Kind != "base" {
		t.Fatalf("ácido a 15%% em Lanette devia avisar, veio %v", alertKinds(alto))
	}
	if !strings.Contains(alto[0].Message, "não iônico") {
		t.Errorf("a saída precisa trazer a recomendação, veio %q", alto[0].Message)
	}

	// 5%: a regra é de concentração, não de presença.
	baixo := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "creme", Vehicle: "Creme Lanette qsp 100 g",
		BaseRules:  []models.MagistralBaseIncompatibility{regra},
		Components: []FormulaCheckComponent{{Substance: acido.Name, Quantity: 5, Unit: "%", Catalog: acido}},
	}, nil)
	if len(baixo) != 0 {
		t.Errorf("ácido a 5%% convive com a base: %v", baixo)
	}

	// Outra base: a mesma fórmula não alerta.
	outra := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "creme", Vehicle: "Creme não iônico qsp 100 g",
		BaseRules:  []models.MagistralBaseIncompatibility{regra},
		Components: []FormulaCheckComponent{{Substance: acido.Name, Quantity: 15, Unit: "%", Catalog: acido}},
	}, nil)
	if len(outra) != 0 {
		t.Errorf("a regra é do Lanette, não de creme em geral: %v", outra)
	}

	// Dose em mg não vira comparação com percentual.
	emMg := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "cápsula", Vehicle: "Creme Lanette",
		BaseRules:  []models.MagistralBaseIncompatibility{regra},
		Components: []FormulaCheckComponent{{Substance: acido.Name, Quantity: 15, Unit: "mg", Catalog: acido}},
	}, nil)
	if len(emMg) != 0 {
		t.Errorf("15 mg não são 15%%: %v", emMg)
	}
}

// Substância que atrapalha exame: a biotina acima de 5 mg/dia derruba imunoensaio biotinilado, e
// num sistema que prescreve e lê exame o estrago volta pela porta da frente.
func TestCheckFormulaInterferenciaEmExame(t *testing.T) {
	texto := "acima de 5 mg/dia interfere em imunoensaio biotinilado; suspender 3 dias antes da coleta"
	limiar := 5.0
	biotina := cat("Biotina", func(c *models.MagistralComponent) {
		c.AssayInterference = &texto
		c.AssayInterferenceDose = &limiar
	})

	alto := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "cápsula", DosesPerDay: 1,
		Components: []FormulaCheckComponent{{Substance: biotina.Name, Quantity: 10, Unit: "mg", Catalog: biotina}},
	}, nil)
	achou := false
	for _, a := range alto {
		if a.Kind == "assay" {
			achou = true
			if a.Level != AlertWarn {
				t.Errorf("interferência em exame é aviso de nível warn, veio %v", a.Level)
			}
			if !strings.Contains(a.Message, "suspender") {
				t.Errorf("a frase precisa dizer o que fazer antes da coleta, veio %q", a.Message)
			}
		}
	}
	if !achou {
		t.Fatalf("10 mg de biotina deviam avisar, veio %v", alertKinds(alto))
	}

	// 2 mg ficam abaixo do limiar descrito e não viram alarme.
	baixo := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "cápsula", DosesPerDay: 1,
		Components: []FormulaCheckComponent{{Substance: biotina.Name, Quantity: 2, Unit: "mg", Catalog: biotina}},
	}, nil)
	for _, a := range baixo {
		if a.Kind == "assay" {
			t.Errorf("2 mg estão abaixo do limiar: %q", a.Message)
		}
	}

	// Duas tomadas de 3 mg somam 6 mg e passam do limiar.
	duas := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "cápsula", DosesPerDay: 2,
		Components: []FormulaCheckComponent{{Substance: biotina.Name, Quantity: 3, Unit: "mg", Catalog: biotina}},
	}, nil)
	achou = false
	for _, a := range duas {
		if a.Kind == "assay" {
			achou = true
		}
	}
	if !achou {
		t.Error("3 mg duas vezes ao dia somam 6 mg: o limiar é diário")
	}
}

// As três conferências comparam número da receita com número do catálogo, e a receita pode vir em
// outra unidade da mesma família. Sem converter, "biotina 500 mcg" contra limiar de 5 mg disparava
// (500 > 5) numa dose de rotina, e "biotina 10 mg" contra teto de 45 mcg passava batido — os dois
// erros no mesmo lugar, em direções opostas.
func TestCheckFormulaConverteUnidadeAntesDeComparar(t *testing.T) {
	texto := "acima de 5 mg/dia interfere em imunoensaio"
	limiar := 5000.0 // o limiar mora na unidade do catálogo, que aqui é mcg
	biotina := cat("Biotina", func(c *models.MagistralComponent) {
		c.DefaultUnit = "mcg"
		c.AssayInterference = &texto
		c.AssayInterferenceDose = &limiar
	})

	rotina := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "cápsula", DosesPerDay: 1,
		Components: []FormulaCheckComponent{{Substance: biotina.Name, Quantity: 500, Unit: "mcg", Catalog: biotina}},
	}, nil)
	for _, a := range rotina {
		if a.Kind == "assay" {
			t.Errorf("500 mcg é dose de rotina e não pode alertar: %q", a.Message)
		}
	}

	alta := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "cápsula", DosesPerDay: 1,
		Components: []FormulaCheckComponent{{Substance: biotina.Name, Quantity: 10, Unit: "mg", Catalog: biotina}},
	}, nil)
	achou := false
	for _, a := range alta {
		if a.Kind == "assay" {
			achou = true
		}
	}
	if !achou {
		t.Error("10 mg são 10.000 mcg e passam do limiar: precisa alertar")
	}
}

// Faixa de dose também converte: antes, componente escrito em unidade diferente da do catálogo era
// simplesmente ignorado pela conferência.
func TestCheckFormulaFaixaConverteUnidade(t *testing.T) {
	min, max := 100.0, 2000.0
	b12 := cat("Metilcobalamina", func(c *models.MagistralComponent) {
		c.DefaultUnit = "mcg"
		c.MinDose = &min
		c.MaxDose = &max
	})
	// 0,5 mg = 500 mcg: dentro da faixa, não alerta.
	dentro := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "cápsula", DosesPerDay: 1,
		Components: []FormulaCheckComponent{{Substance: b12.Name, Quantity: 0.5, Unit: "mg", Catalog: b12}},
	}, nil)
	if len(dentro) != 0 {
		t.Errorf("0,5 mg = 500 mcg está dentro de 100 a 2000: %v", dentro)
	}
	// 0,05 mg = 50 mcg: abaixo.
	abaixo := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "cápsula", DosesPerDay: 1,
		Components: []FormulaCheckComponent{{Substance: b12.Name, Quantity: 0.05, Unit: "mg", Catalog: b12}},
	}, nil)
	if len(abaixo) != 1 || abaixo[0].Kind != "dose" {
		t.Errorf("50 mcg está abaixo da faixa: %v", alertKinds(abaixo))
	}
	// Unidade sem massa não vira comparação inventada.
	semMassa := CheckFormula(FormulaCheckInput{
		PharmaceuticalForm: "cápsula", DosesPerDay: 1,
		Components: []FormulaCheckComponent{{Substance: b12.Name, Quantity: 5, Unit: "UI", Catalog: b12}},
	}, nil)
	if len(semMassa) != 0 {
		t.Errorf("UI não converte para mcg: %v", semMassa)
	}
}
