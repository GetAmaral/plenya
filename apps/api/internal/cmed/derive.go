// Package cmed lê a Lista de Preços de Medicamentos (PMC) publicada pela CMED/ANVISA e a
// traduz para o catálogo do EMR.
//
// A planilha oficial descreve o medicamento em UM campo de texto ("500 MG COM REV CT 2 BL AL
// PLAS TRANS X 07"): concentração, forma farmacêutica e embalagem misturadas com abreviações.
// Este arquivo desmancha esse texto. Nada aqui toca banco nem rede — é tudo função pura, para
// poder ser testado com as linhas reais da planilha.
package cmed

import (
	"regexp"
	"strconv"
	"strings"
)

// Derived é o que se consegue extrair do texto de apresentação.
type Derived struct {
	Concentration string // "500 MG", "30 MG/ML"
	Form          string // canônico: comprimido_revestido, capsula_dura, solucao_injetavel...
	Route         string // oral, intravenosa, subcutanea, topica...
	PackageQty    *int   // unidades na caixa, quando o texto termina em "X <inteiro>"
	Confidence    string // high | medium | none
}

// maxConcentrationLen limita o texto de concentração de uma associação. Acima disso o
// trecho deixa de ser dose e vira descrição — e a coluna do catálogo é varchar(120).
const maxConcentrationLen = 80

// Níveis de confiança da derivação. Vão gravados na linha para a UI poder avisar.
const (
	ConfidenceHigh   = "high"   // concentração E forma derivadas
	ConfidenceMedium = "medium" // só uma das duas
	ConfidenceNone   = "none"   // nenhuma
)

// dosePatterns — ordem de prioridade (concentração → peso → unidade). O primeiro padrão que
// casa vence e devolve a ÚLTIMA ocorrência dele no texto.
//
// Portado verbatim do nai2 (src/server/clinico/dose-apresentacao.ts), que por sua vez
// transcreveu o extrator do sistema legado. Medido sobre as 25.570 linhas da edição de
// julho/2026: só ~2% ficam sem dose, e são apresentações que realmente não trazem concentração
// no texto (vacinas, biológicos em seringa preenchida).
var dosePatterns = []*regexp.Regexp{
	regexp.MustCompile(`(?i)(\d+(?:[,.]\d+)?\s*μg/ml)`),
	regexp.MustCompile(`(?i)(\d+(?:[,.]\d+)?\s*mcg/ml)`),
	regexp.MustCompile(`(?i)(\d+(?:[,.]\d+)?\s*mg/ml)`),
	regexp.MustCompile(`(?i)(\d+(?:[,.]\d+)?\s*ui/ml)`),
	regexp.MustCompile(`(?i)(\d+(?:[,.]\d+)?\s*μg)`),
	regexp.MustCompile(`(?i)(\d+(?:[,.]\d+)?\s*µg)`),
	regexp.MustCompile(`(?i)(\d+(?:[,.]\d+)?\s*mcg)`),
	regexp.MustCompile(`(?i)(\d+(?:[,.]\d+)?\s*mg/g)`),
	regexp.MustCompile(`(?i)(\d+(?:[,.]\d+)?\s*mg)`),
	// "g" só quando não é multiplicador de embalagem ("X 30 G" é peso, "30 G X 2" não).
	regexp.MustCompile(`(?i)(\d+(?:[,.]\d+)?\s*g)(?:\s*[xX])?`),
	regexp.MustCompile(`(?i)(\d+(?:[,.]\d+)?\s*MUI)`),
	regexp.MustCompile(`(?i)(\d+(?:[,.]\d+)?\s*KUI)`),
	regexp.MustCompile(`(?i)(\d+(?:[,.]\d+)?\s*ui)`),
	regexp.MustCompile(`(?i)(\d+(?:[,.]\d+)?\s*u)(?:[^i]|$)`),
	regexp.MustCompile(`(?i)(\d+\s*MILHÕES/ML)`),
	regexp.MustCompile(`(?i)(\d+\s*MILHÕES)`),
	regexp.MustCompile(`(?i)(\d+(?:[,.]\d+)?\s*CFU)`),
	regexp.MustCompile(`(?i)(\d+(?:[,.]\d+)?\s*GV)`),
}

var (
	reLeadingParen = regexp.MustCompile(`^\s*\(([^)]+)\)`)
	// Case-SENSITIVE de propósito (verbatim do original): "50MG +" não casa aqui e cai no
	// ladder, que é o comportamento observado no dado.
	reCombo      = regexp.MustCompile(`(\d+(?:[,.]\d+)?\s*(?:mg|mcg|μg|µg|ui|u|g)(?:/(?:ml|g))?)\s*\+`)
	reComboParen = regexp.MustCompile(`(?i)\((\d+(?:[,.]\d+)?)\s*\+\s*(\d+(?:[,.]\d+)?)\)\s*(ui|u|mui|kui)`)
	rePackageQty = regexp.MustCompile(`(?i)\sX\s*(\d{1,4})\s*$`)
	reSpaces     = regexp.MustCompile(`\s+`)
)

// ExtractConcentration deriva a concentração do texto de apresentação. "" quando não deriva.
func ExtractConcentration(presentation string) string {
	if strings.TrimSpace(presentation) == "" {
		return ""
	}

	// (1) Parênteses logo no início descrevem a associação inteira: devolve o conteúdo.
	if m := reLeadingParen.FindStringSubmatch(presentation); m != nil {
		return strings.TrimSpace(m[1])
	}
	// (2) Associação com "+": a primeira dose representa.
	if m := reCombo.FindStringSubmatch(presentation); m != nil {
		return strings.TrimSpace(m[1])
	}
	// (3) Associação entre parênteses com unidade fora: "(100 + 100) UI".
	if m := reComboParen.FindStringSubmatch(presentation); m != nil {
		return strings.TrimSpace(m[1] + " " + m[3])
	}
	// (4) Ladder por prioridade. Dentro do padrão vencedor, o original devolvia a ÚLTIMA
	// ocorrência — o que, numa associação escrita em maiúsculas ("10 MG/G + 0,443 MG/G",
	// que a regra (2) não pega por ser case-sensitive), faria a concentração do creme virar
	// "0,443 MG/G": o ativo secundário passando por dose do produto. Numa receita isso é pior
	// que campo vazio, então associação devolve o trecho INTEIRO, do primeiro ao último match.
	for _, p := range dosePatterns {
		idx := p.FindAllStringSubmatchIndex(presentation, -1)
		if len(idx) == 0 {
			continue
		}
		last := idx[len(idx)-1]
		if len(idx) > 1 {
			span := strings.TrimSpace(presentation[idx[0][2]:last[3]])
			// Associação de 2-3 ativos cabe e informa; acima disso o "span" vira a
			// apresentação inteira e deixa de ser concentração. Nesse caso vale mais o
			// ativo principal (o primeiro) do que um texto que ninguém lê na receita.
			if strings.Contains(span, "+") && len(span) <= maxConcentrationLen {
				return span
			}
			return strings.TrimSpace(presentation[idx[0][2]:idx[0][3]])
		}
		return strings.TrimSpace(presentation[last[2]:last[3]])
	}
	return ""
}

// packagingTokens marcam o início da descrição de EMBALAGEM. A forma farmacêutica está entre
// a concentração e o primeiro destes.
var packagingTokens = map[string]bool{
	"CT": true, "CX": true, "FR": true, "BG": true, "EMB": true, "ENV": true, "POT": true,
	"TB": true, "AMP": true, "FA": true, "SER": true, "SIS": true, "BL": true, "CAR": true,
	"CAM": true, "CP": true, "SG": true, "ESTOJO": true,
}

// formCanonical mapeia a janela de tokens da ANVISA para a forma canônica. Ordem de checagem:
// as chaves mais longas primeiro (ver ExtractForm), senão "COM" engoliria "COM REV LIB PROL".
var formCanonical = map[string]string{
	"COM REV LIB PROL": "comprimido_revestido_liberacao_prolongada",
	"COM LIB PROL":     "comprimido_liberacao_prolongada",
	"COM REV":          "comprimido_revestido",
	"COM EFEV":         "comprimido_efervescente",
	"COM MAST":         "comprimido_mastigavel",
	"COM SUBL":         "comprimido_sublingual",
	"COM ORODISP":      "comprimido_orodispersivel",
	"COM DISP":         "comprimido_dispersivel",
	"COM":              "comprimido",
	"DRG":              "drageas",
	"CAP GEL DURA":     "capsula_dura",
	"CAP DURA":         "capsula_dura",
	"CAP GEL MOLE":     "capsula_mole",
	"CAP MOLE":         "capsula_mole",
	"CAP LIB PROL":     "capsula_liberacao_prolongada",
	"CAP":              "capsula",
	"PO LIOF SOL INJ":  "po_liofilizado_injetavel",
	"PO LIOF":          "po_liofilizado",
	"PO SOL OR":        "po_solucao_oral",
	"PO SUS OR":        "po_suspensao_oral",
	"PO":               "po",
	"SOL INJ IV":       "solucao_injetavel_iv",
	"SOL INJ IM":       "solucao_injetavel_im",
	"SOL INJ SC":       "solucao_injetavel_sc",
	"SOL INJ":          "solucao_injetavel",
	"SUS INJ":          "suspensao_injetavel",
	"EMU INJ":          "emulsao_injetavel",
	"SOL OR":           "solucao_oral",
	"SOL":              "solucao",
	"SUS OR":           "suspensao_oral",
	"SUS":              "suspensao",
	"XPE":              "xarope",
	"GTS":              "gotas",
	"SOL OFT":          "solucao_oftalmica",
	"SUS OFT":          "suspensao_oftalmica",
	"POM OFT":          "pomada_oftalmica",
	"SOL NAS":          "solucao_nasal",
	"SPR NAS":          "spray_nasal",
	"SOL OT":           "solucao_otologica",
	"CREM DERM":        "creme_dermatologico",
	"CREM VAG":         "creme_vaginal",
	"CREM":             "creme",
	"GEL DERM":         "gel_dermatologico",
	"GEL":              "gel",
	"POM DERM":         "pomada_dermatologica",
	"POM":              "pomada",
	"LOC":              "locao",
	"SOL TOP":          "solucao_topica",
	"AER":              "aerossol",
	"SPR":              "spray",
	"SUP":              "supositorio",
	"OVU":              "ovulo",
	"ADES TRANSD":      "adesivo_transdermico",
	"ADES":             "adesivo",
	"SHAMPOO":          "shampoo",
	"PAST":             "pastilha",
	"GRAN":             "granulado",
	"SACHE":            "sache",
}

// formRoute — via de administração DERIVADA da forma. Extrair a via do texto não funciona:
// comprimido oral simplesmente não escreve "oral", e ~76% das linhas ficariam indefinidas.
var formRoute = map[string]string{
	"comprimido": "oral", "comprimido_revestido": "oral", "comprimido_efervescente": "oral",
	"comprimido_mastigavel": "oral", "comprimido_orodispersivel": "oral",
	"comprimido_dispersivel": "oral", "comprimido_liberacao_prolongada": "oral",
	"comprimido_revestido_liberacao_prolongada": "oral", "comprimido_sublingual": "sublingual",
	"drageas": "oral", "capsula": "oral", "capsula_dura": "oral", "capsula_mole": "oral",
	"capsula_liberacao_prolongada": "oral", "po_solucao_oral": "oral", "po_suspensao_oral": "oral",
	"solucao_oral": "oral", "suspensao_oral": "oral", "xarope": "oral", "gotas": "oral",
	"granulado": "oral", "sache": "oral", "pastilha": "oral",
	"solucao_injetavel": "injetavel", "solucao_injetavel_iv": "intravenosa",
	"solucao_injetavel_im": "intramuscular", "solucao_injetavel_sc": "subcutanea",
	"suspensao_injetavel": "injetavel", "emulsao_injetavel": "injetavel",
	"po_liofilizado_injetavel": "injetavel",
	"solucao_oftalmica":        "oftalmica", "suspensao_oftalmica": "oftalmica",
	"pomada_oftalmica": "oftalmica", "solucao_otologica": "otologica",
	"solucao_nasal": "nasal", "spray_nasal": "nasal",
	"creme_dermatologico": "topica", "gel_dermatologico": "topica",
	"pomada_dermatologica": "topica", "solucao_topica": "topica", "creme": "topica",
	"gel": "topica", "pomada": "topica", "locao": "topica", "shampoo": "topica",
	"adesivo": "transdermica", "adesivo_transdermico": "transdermica",
	"creme_vaginal": "vaginal", "ovulo": "vaginal", "supositorio": "retal",
	"aerossol": "inalatoria", "spray": "topica",
}

// formKeysByLength são as chaves de formCanonical em ordem decrescente de nº de palavras,
// para casar sempre a mais específica.
var formKeysByLength = sortedFormKeys()

func sortedFormKeys() []string {
	keys := make([]string, 0, len(formCanonical))
	for k := range formCanonical {
		keys = append(keys, k)
	}
	// Ordena por nº de tokens (desc) e, empatado, por tamanho (desc): determinístico.
	for i := 0; i < len(keys); i++ {
		for j := i + 1; j < len(keys); j++ {
			ai, aj := len(strings.Fields(keys[i])), len(strings.Fields(keys[j]))
			if aj > ai || (aj == ai && len(keys[j]) > len(keys[i])) {
				keys[i], keys[j] = keys[j], keys[i]
			}
		}
	}
	return keys
}

// ExtractForm deriva a forma farmacêutica canônica. "" quando não reconhece.
func ExtractForm(presentation string) string {
	up := strings.ToUpper(removeAccents(presentation))
	// Corta em "/" de dose (ex.: "MG/ML") para não confundir com abreviação de forma.
	if i := strings.Index(up, " X "); i > 0 {
		// mantém o trecho antes do multiplicador de embalagem
		up = up[:i]
	}
	tokens := strings.Fields(up)

	// Descarta o prefixo de concentração (números e unidades) e para no 1º token de embalagem.
	start := 0
	for start < len(tokens) && looksNumericOrUnit(tokens[start]) {
		start++
	}
	end := start
	for end < len(tokens) && !packagingTokens[tokens[end]] {
		end++
	}
	window := strings.Join(tokens[start:end], " ")
	if window == "" {
		window = strings.Join(tokens, " ")
	}

	for _, key := range formKeysByLength {
		if window == key || strings.HasPrefix(window, key+" ") || strings.Contains(" "+window+" ", " "+key+" ") {
			return formCanonical[key]
		}
	}
	return ""
}

var reNumericUnit = regexp.MustCompile(`^[\d,.+()%/-]*(MG|MCG|G|ML|UI|U|KUI|MUI|CFU|GV|UG|%|MG/ML|MG/G|UI/ML)?[\d,.)%]*$`)

func looksNumericOrUnit(tok string) bool {
	if tok == "+" || tok == "-" {
		return true
	}
	return reNumericUnit.MatchString(tok) && strings.ContainsAny(tok, "0123456789")
}

// ExtractPackageQuantity devolve a quantidade da embalagem quando o texto termina em
// "X <inteiro>" sem unidade. Quando termina em volume ("X 100 ML") devolve nil: ali o número
// é volume, não contagem, e chutar viraria quantidade errada na receita.
func ExtractPackageQuantity(presentation string) *int {
	m := rePackageQty.FindStringSubmatch(strings.TrimSpace(presentation))
	if m == nil {
		return nil
	}
	n, err := strconv.Atoi(m[1])
	if err != nil || n <= 0 {
		return nil
	}
	return &n
}

// DeriveFromPresentation roda as três extrações e resume a confiança.
func DeriveFromPresentation(presentation string) Derived {
	d := Derived{
		Concentration: ExtractConcentration(presentation),
		Form:          ExtractForm(presentation),
		PackageQty:    ExtractPackageQuantity(presentation),
	}
	d.Route = formRoute[d.Form]

	switch {
	case d.Concentration != "" && d.Form != "":
		d.Confidence = ConfidenceHigh
	case d.Concentration != "" || d.Form != "":
		d.Confidence = ConfidenceMedium
	default:
		d.Confidence = ConfidenceNone
	}
	return d
}

// removeAccents deixa o texto comparável com as chaves ASCII do dicionário de formas
// ("PÓ LIOF" → "PO LIOF").
func removeAccents(s string) string {
	repl := strings.NewReplacer(
		"á", "a", "à", "a", "â", "a", "ã", "a", "ä", "a",
		"é", "e", "ê", "e", "è", "e", "ë", "e",
		"í", "i", "î", "i", "ì", "i", "ï", "i",
		"ó", "o", "ô", "o", "õ", "o", "ò", "o", "ö", "o",
		"ú", "u", "û", "u", "ù", "u", "ü", "u",
		"ç", "c", "ñ", "n",
		"Á", "A", "À", "A", "Â", "A", "Ã", "A", "Ä", "A",
		"É", "E", "Ê", "E", "È", "E", "Ë", "E",
		"Í", "I", "Î", "I", "Ì", "I", "Ï", "I",
		"Ó", "O", "Ô", "O", "Õ", "O", "Ò", "O", "Ö", "O",
		"Ú", "U", "Û", "U", "Ù", "U", "Ü", "U",
		"Ç", "C", "Ñ", "N",
	)
	return reSpaces.ReplaceAllString(repl.Replace(s), " ")
}
