package services

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/plenya/api/internal/dto"
)

// Números do deck, e de onde eles vieram.
//
// Todo número que chega ao documento que o paciente lê sai de um conjunto fechado: o dossiê
// congelado. Isso permite VERIFICAR, em vez de pedir ao modelo que não invente. Pedir é o máximo
// que se consegue quando a entrada é uma transcrição livre (é o que `ai_note_service` faz); aqui a
// entrada é um catálogo, então a garantia pode ser de código.
//
// O limite honesto, e ele precisa estar escrito aqui e visível na tela: isto prova que o número
// EXISTE no dossiê, nunca que ele significa o que a frase diz. "Sua ferritina está em 96" e "seu
// colesterol está em 96" passam identicamente se 96 existir em algum lugar. O que a verificação
// entrega é superfície de julgamento — a operação vira sugestão com a origem candidata anexada, e
// o médico julga uma afirmação específica em vez de reler um parágrafo.

// numeralRe casa um número em grafia PT-BR ou inglesa, com o sufixo colado quando houver.
//
// Cobre `139`, `13,5`, `1.023`, `63,10`, `112 g`, `12%`, `-2,5`. O sufixo é capturado porque
// "112 g" e "112 mg" são coisas diferentes e a distinção às vezes é tudo o que separa uma dose
// certa de uma errada.
var numeralRe = regexp.MustCompile(`(-?\d[\d.,]*)\s*(%|[a-zA-ZµμÁ-ú/³²]+(?:/[a-zA-ZµμÁ-ú³²]+)?)?`)

// Numeral — um número como ele aparece no texto, com as leituras possíveis.
//
// `Values` tem mais de um elemento quando a grafia é ambígua: `1.023` é mil e vinte e três em
// PT-BR e é 1,023 em inglês, e a densidade urinária de verdade vale 1,023. Emitir as duas leituras
// e aceitar qualquer uma é o certo aqui, porque a pergunta é "este número existe no dossiê", e não
// "quanto ele vale".
type Numeral struct {
	Raw      string
	Values   []float64
	Unit     string
	Decimals int // casas decimais escritas, para calibrar a tolerância de arredondamento
}

// mesesPT — para reconhecer data por extenso e não tratar o dia como número clínico.
var mesesPT = regexp.MustCompile(`(?i)^\s*de\s+(janeiro|fevereiro|mar[çc]o|abril|maio|junho|julho|agosto|setembro|outubro|novembro|dezembro)`)

// anoRe — 1900 a 2099 solto no texto é ano, não medida.
var anoRe = regexp.MustCompile(`^(19|20)\d\d$`)

// ExtractNumerals tira de um texto todos os números com as leituras possíveis.
//
// Ignora data escrita por extenso e ano. Não é preciosismo: no primeiro teste com o modelo, "7 de
// fevereiro de 2026" produziu duas falsas provas — o 7 casou com a borda de faixa de um exame sem
// relação, e o 2026 virou "número sem origem no dossiê". Alarme falso na tela de aceite é pior que
// inútil: ensina a clicar sem ler, que é exatamente a falha que esta verificação existe para evitar.
func ExtractNumerals(s string) []Numeral {
	var out []Numeral
	for _, m := range numeralRe.FindAllStringSubmatchIndex(s, -1) {
		bruto := strings.TrimRight(s[m[2]:m[3]], ".,;:")
		if bruto == "" {
			continue
		}
		// "7 de fevereiro": o 7 é dia, não medida. Confere a partir do fim do NÚMERO, e não do
		// fim do match inteiro — o "de" é capturado como se fosse unidade, e olhar depois dele
		// perderia justamente o mês que identifica a data.
		if mesesPT.MatchString(s[m[3]:]) {
			continue
		}
		if anoRe.MatchString(bruto) {
			continue
		}
		vals, casas := leituras(bruto)
		if len(vals) == 0 {
			continue
		}
		unidade := ""
		if m[4] >= 0 {
			unidade = strings.TrimSpace(s[m[4]:m[5]])
		}
		out = append(out, Numeral{
			Raw:      bruto,
			Values:   vals,
			Unit:     unidade,
			Decimals: casas,
		})
	}
	return out
}

// leituras interpreta a grafia. Devolve as leituras possíveis e quantas casas decimais foram
// escritas na mais provável.
func leituras(bruto string) ([]float64, int) {
	s := strings.TrimSpace(bruto)
	if s == "" {
		return nil, 0
	}
	temVirgula := strings.Contains(s, ",")
	temPonto := strings.Contains(s, ".")

	switch {
	case temVirgula && temPonto:
		// "1.234,5": ponto é milhar, vírgula é decimal. Sem ambiguidade.
		limpo := strings.ReplaceAll(s, ".", "")
		limpo = strings.Replace(limpo, ",", ".", 1)
		if v, err := strconv.ParseFloat(limpo, 64); err == nil {
			return []float64{v}, casasDepoisDe(s, ",")
		}
	case temVirgula:
		// "13,5": decimal PT-BR.
		if v, err := strconv.ParseFloat(strings.Replace(s, ",", ".", 1), 64); err == nil {
			return []float64{v}, casasDepoisDe(s, ",")
		}
	case temPonto:
		// "1.023" é ambíguo: milhar em PT-BR, decimal em inglês. As duas leituras entram.
		var vals []float64
		if v, err := strconv.ParseFloat(s, 64); err == nil {
			vals = append(vals, v)
		}
		if semPonto := strings.ReplaceAll(s, ".", ""); semPonto != s {
			if v, err := strconv.ParseFloat(semPonto, 64); err == nil && (len(vals) == 0 || v != vals[0]) {
				vals = append(vals, v)
			}
		}
		return vals, casasDepoisDe(s, ".")
	default:
		if v, err := strconv.ParseFloat(s, 64); err == nil {
			return []float64{v}, 0
		}
	}
	return nil, 0
}

func casasDepoisDe(s, sep string) int {
	i := strings.LastIndex(s, sep)
	if i < 0 {
		return 0
	}
	return len(s) - i - 1
}

// NumeralFact — uma ocorrência de número no dossiê, com de onde ela veio.
//
// `Source` é opaco de propósito e legível por gente: é o que vai para a tela ao lado do botão de
// aceitar, para o médico julgar a origem em vez de aceitar prosa.
type NumeralFact struct {
	Value  float64 `json:"value"`
	Unit   string  `json:"unit,omitempty"`
	Source string  `json:"source"`
	Label  string  `json:"label"`
}

// NumericIndex — todos os números do dossiê, indexados por valor arredondado.
//
// A chave é o valor com quatro casas, que é a precisão do banco (`numeric(12,4)`). A busca aceita
// tolerância maior, então a chave só serve para reduzir o espaço de comparação.
type NumericIndex struct {
	porValor map[string][]NumeralFact
	todos    []NumeralFact
}

func chaveDeValor(v float64) string { return strconv.FormatFloat(math.Round(v*1e4)/1e4, 'f', -1, 64) }

// BuildNumericIndex varre o dossiê congelado e cataloga todo número que ele contém.
//
// Cobre régua (histórico, eixo, bordas e faixas), achados, escore, vitais e a idade. O que NÃO
// entra é o que o dossiê não expõe em número: o conteúdo das prescrições chega só com id, data e
// status, então dose de fórmula não é verificável por aqui.
func BuildNumericIndex(d *dto.PlanDossierResponse) *NumericIndex {
	ix := &NumericIndex{porValor: map[string][]NumeralFact{}}
	if d == nil {
		return ix
	}
	add := func(v float64, unit, source, label string) {
		f := NumeralFact{Value: v, Unit: unit, Source: source, Label: label}
		ix.porValor[chaveDeValor(v)] = append(ix.porValor[chaveDeValor(v)], f)
		ix.todos = append(ix.todos, f)
	}

	for code, r := range d.Rulers {
		nome := r.Name
		for _, p := range r.History {
			add(p.Value, r.Unit,
				fmt.Sprintf("ruler:%s:history:%s", code, p.Date),
				fmt.Sprintf("%s, %s", nome, p.Date))
		}
		for i, a := range r.Axis {
			add(a, r.Unit, fmt.Sprintf("ruler:%s:axis:%d", code, i), nome+", limite do eixo")
		}
		for _, e := range r.Edges {
			add(e, r.Unit, fmt.Sprintf("ruler:%s:edge", code), nome+", borda de faixa")
		}
		for _, sg := range r.Segments {
			add(sg.A, r.Unit, fmt.Sprintf("ruler:%s:segment:%d:a", code, sg.Level), nome+", início da faixa")
			add(sg.B, r.Unit, fmt.Sprintf("ruler:%s:segment:%d:b", code, sg.Level), nome+", fim da faixa")
		}
		if r.Points > 0 {
			add(r.Points, "pontos", "ruler:"+code+":points", nome+", peso no escore")
		}
	}

	achados := func(lista []dto.PlanFinding, tipo string) {
		for _, f := range lista {
			base := fmt.Sprintf("finding:%s:%s", tipo, f.Code)
			add(f.Value, f.Unit, base+":value", f.Name)
			if f.PointsLost > 0 {
				add(f.PointsLost, "pontos", base+":pointsLost", f.Name+", pontos perdidos")
			}
			if f.Points > 0 {
				add(f.Points, "pontos", base+":points", f.Name+", peso")
			}
		}
	}
	achados(d.Strong, "strong")
	achados(d.Moving, "moving")

	if d.Snapshot != nil {
		add(d.Snapshot.TotalPercentage, "%", "snapshot:total", "escore total")
	}
	for i, v := range d.Vitals {
		p := fmt.Sprintf("vitals:%d", i)
		intOpcional(add, v.SystolicBP, "mmHg", p+":sistolica", "pressão sistólica")
		intOpcional(add, v.DiastolicBP, "mmHg", p+":diastolica", "pressão diastólica")
		intOpcional(add, v.HeartRate, "bpm", p+":fc", "frequência cardíaca")
		numOpcional(add, v.Weight, "kg", p+":peso", "peso")
		numOpcional(add, v.Height, "cm", p+":altura", "altura")
		numOpcional(add, v.Waist, "cm", p+":cintura", "cintura")
		numOpcional(add, v.BMI, "kg/m²", p+":imc", "IMC")
	}
	if d.Patient.Age > 0 {
		add(float64(d.Patient.Age), "anos", "patient:age", "idade do paciente")
	}
	return ix
}

func numOpcional(add func(float64, string, string, string), v *float64, unit, source, label string) {
	if v != nil {
		add(*v, unit, source, label)
	}
}

func intOpcional(add func(float64, string, string, string), v *int, unit, source, label string) {
	if v != nil {
		add(float64(*v), unit, source, label)
	}
}

// Match procura o numeral no dossiê e devolve as origens candidatas.
//
// A tolerância é a do arredondamento escrito: "97" casa com 96,8 porque quem escreve 97 está
// arredondando; "96,8" não casa com 97, porque quem escreveu duas casas afirmou duas casas.
func (ix *NumericIndex) Match(n Numeral) []NumeralFact {
	if ix == nil {
		return nil
	}
	var out []NumeralFact
	vistos := map[string]bool{}
	for _, v := range n.Values {
		for _, f := range ix.todos {
			if !mesmoNumero(v, f.Value, n.Decimals) {
				continue
			}
			if vistos[f.Source] {
				continue
			}
			vistos[f.Source] = true
			out = append(out, f)
		}
	}
	// Ordena por quão plausível é que a frase esteja falando DAQUILO. A tela mostra a primeira
	// origem como a candidata, e mostrar a errada é pior que mostrar várias: parece autoritativo.
	//
	// Veio de um caso real: "sua lipase estava em 27 U/L" casou primeiro com o limite do eixo do
	// cortisol, que por acaso também vale 27. Medida do paciente vem antes de estrutura de escala,
	// e unidade igual desempata.
	sort.SliceStable(out, func(i, j int) bool {
		pi, pj := relevanciaDaOrigem(out[i], n), relevanciaDaOrigem(out[j], n)
		return pi > pj
	})
	return out
}

// relevanciaDaOrigem pontua o quanto uma origem provavelmente é a que a frase cita.
//
// Uma medida do paciente (histórico de exame, valor de achado) é quase sempre o que se está
// citando; borda de faixa, limite de eixo e peso de item existem no dossiê mas raramente aparecem
// numa frase para o paciente.
func relevanciaDaOrigem(f NumeralFact, n Numeral) int {
	p := 0
	switch {
	case strings.Contains(f.Source, ":history:"):
		p += 100 // medida do paciente, com data
	case strings.Contains(f.Source, "finding:"):
		p += 90
	case strings.HasPrefix(f.Source, "vitals:"):
		p += 80
	case strings.HasPrefix(f.Source, "snapshot:"), strings.HasPrefix(f.Source, "patient:"):
		p += 60
	case strings.Contains(f.Source, ":segment:"), strings.Contains(f.Source, ":edge"):
		p += 20 // estrutura da escala
	case strings.Contains(f.Source, ":axis:"):
		p += 10
	case strings.Contains(f.Source, ":points"):
		p += 5
	}
	// Unidade escrita igual à da origem é evidência forte de que é a mesma grandeza.
	if n.Unit != "" && f.Unit != "" && strings.EqualFold(strings.TrimSpace(n.Unit), strings.TrimSpace(f.Unit)) {
		p += 50
	}
	return p
}

func mesmoNumero(escrito, doDossie float64, casas int) bool {
	p := math.Pow(10, float64(casas))
	return math.Abs(math.Round(doDossie*p)/p-escrito) < 1e-9
}

// NumeralDelta devolve os numerais que existem em `depois` e não existiam em `antes`.
//
// É o coração da classificação: uma edição de texto que não introduz número novo pode ser aplicada
// direto ("encurta o título", "tira o jargão"); uma que introduz precisa de aceite, mesmo que o
// campo seja `title`. Classificar pelo CAMPO não funciona, porque `punch`, `title` e célula de
// tabela carregam número o tempo todo.
func NumeralDelta(antes, depois string) []Numeral {
	tinha := map[string]bool{}
	for _, n := range ExtractNumerals(antes) {
		tinha[chaveDeNumeral(n)] = true
	}
	var novos []Numeral
	for _, n := range ExtractNumerals(depois) {
		if !tinha[chaveDeNumeral(n)] {
			novos = append(novos, n)
		}
	}
	return novos
}

// chaveDeNumeral identifica o numeral pela grafia normalizada mais a unidade colada: trocar "112 g"
// por "112 mg" é mudança, ainda que o número seja o mesmo.
func chaveDeNumeral(n Numeral) string {
	return strings.ReplaceAll(n.Raw, ",", ".") + "|" + strings.ToLower(n.Unit)
}
