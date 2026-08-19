package cmed

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

// Record é uma linha da Lista de Preços já normalizada — antes de virar linha do catálogo.
type Record struct {
	GGREM                string
	Registro             string
	Substance            string
	Product              string
	Presentation         string
	Laboratory           string
	EAN13                string
	TherapeuticClass     string
	TherapeuticClassCode string
	ProductType          string
	Stripe               string
	PMCPrice             *float64
	HospitalOnly         bool
}

// Colunas que o import consome, pelo nome EXATO publicado pela ANVISA (comparado já
// normalizado: sem acento, maiúsculo, espaços colapsados).
const (
	colSubstance    = "SUBSTANCIA"
	colLaboratory   = "LABORATORIO"
	colGGREM        = "CODIGO GGREM"
	colRegistro     = "REGISTRO"
	colEAN1         = "EAN 1"
	colProduct      = "PRODUTO"
	colPresentation = "APRESENTACAO"
	colClass        = "CLASSE TERAPEUTICA"
	colProductType  = "TIPO DE PRODUTO (STATUS DO PRODUTO)"
	colPMC0         = "PMC 0 %"
	colHospital     = "RESTRICAO HOSPITALAR"
	colStripe       = "TARJA"
)

// ParseFile lê a planilha aberta e devolve as linhas já normalizadas, deduplicadas por GGREM
// (a última ocorrência vence, que é como a CMED publica correções dentro da mesma edição).
func ParseFile(f *excelize.File) ([]Record, error) {
	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return nil, fmt.Errorf("planilha sem abas")
	}
	sheet := sheets[0]

	headerRow, cols, err := findHeader(f, sheet)
	if err != nil {
		return nil, err
	}

	rows, err := f.Rows(sheet)
	if err != nil {
		return nil, fmt.Errorf("abrir linhas: %w", err)
	}
	defer func() { _ = rows.Close() }()

	// Ordem de chegada preservada: a dedup guarda o índice, para a saída ser determinística
	// (import roda todo mês; diff instável esconderia mudança de verdade).
	byGGREM := map[string]Record{}
	order := []string{}

	line := 0
	for rows.Next() {
		line++
		if line <= headerRow {
			continue
		}
		// RawCellValue: sem isso o preço vem FORMATADO pela máscara da célula e o parser
		// brasileiro ("1.234,56") destrói o número.
		cells, err := rows.Columns(excelize.Options{RawCellValue: true})
		if err != nil {
			return nil, fmt.Errorf("ler linha %d: %w", line, err)
		}

		ggrem := strings.TrimSpace(at(cells, cols[colGGREM]))
		if ggrem == "" {
			continue // linhas de rodapé/nota da planilha
		}

		classRaw := at(cells, cols[colClass])
		code, desc := SplitTherapeuticClass(classRaw)

		rec := Record{
			GGREM:                ggrem,
			Registro:             strings.TrimSpace(at(cells, cols[colRegistro])),
			Substance:            strings.TrimSpace(at(cells, cols[colSubstance])),
			Product:              strings.TrimSpace(at(cells, cols[colProduct])),
			Presentation:         strings.TrimSpace(at(cells, cols[colPresentation])),
			Laboratory:           strings.TrimSpace(at(cells, cols[colLaboratory])),
			EAN13:                NormalizeEAN(at(cells, cols[colEAN1])),
			TherapeuticClass:     desc,
			TherapeuticClassCode: code,
			ProductType:          strings.TrimSpace(at(cells, cols[colProductType])),
			Stripe:               NormalizeStripe(at(cells, cols[colStripe])),
			PMCPrice:             ParseBRNumber(at(cells, cols[colPMC0])),
			HospitalOnly:         strings.EqualFold(strings.TrimSpace(at(cells, cols[colHospital])), "Sim"),
		}

		if _, seen := byGGREM[ggrem]; !seen {
			order = append(order, ggrem)
		}
		byGGREM[ggrem] = rec
	}

	out := make([]Record, 0, len(order))
	for _, g := range order {
		out = append(out, byGGREM[g])
	}
	return out, nil
}

// findHeader localiza a linha de cabeçalho e mapeia coluna→índice POR NOME.
//
// A planilha da CMED tem ~40 linhas de preâmbulo com notas legais, e esse tamanho MUDA entre
// edições mensais. Ancorar em índice fixo quebra silenciosamente no mês seguinte — daí a
// varredura procurando a linha que tem "SUBSTÂNCIA" e alguma célula com "GGREM".
func findHeader(f *excelize.File, sheet string) (int, map[string]int, error) {
	rows, err := f.Rows(sheet)
	if err != nil {
		return 0, nil, fmt.Errorf("abrir linhas: %w", err)
	}
	defer func() { _ = rows.Close() }()

	line := 0
	for rows.Next() && line < 80 {
		line++
		cells, err := rows.Columns()
		if err != nil {
			continue
		}

		hasSubstance, hasGGREM := false, false
		cols := map[string]int{}
		for i, raw := range cells {
			name := NormalizeHeader(raw)
			if name == "" {
				continue
			}
			if _, dup := cols[name]; !dup {
				cols[name] = i
			}
			if name == colSubstance {
				hasSubstance = true
			}
			if strings.Contains(name, "GGREM") {
				hasGGREM = true
			}
		}

		if hasSubstance && hasGGREM {
			for _, required := range []string{colGGREM, colProduct, colPresentation} {
				if _, ok := cols[required]; !ok {
					return 0, nil, fmt.Errorf("cabeçalho na linha %d sem a coluna %q", line, required)
				}
			}
			return line, cols, nil
		}
	}
	return 0, nil, fmt.Errorf("cabeçalho não encontrado nas primeiras 80 linhas — a planilha mudou de formato")
}

// NormalizeHeader deixa o nome da coluna comparável: sem acento, maiúsculo, espaços colapsados.
// Trata o NBSP que a ANVISA deixa no fim de "DESTINAÇÃO COMERCIAL".
func NormalizeHeader(raw string) string {
	s := strings.ReplaceAll(raw, " ", " ")
	s = removeAccents(s)
	s = strings.ToUpper(strings.TrimSpace(s))
	return reSpaces.ReplaceAllString(s, " ")
}

// ParseBRNumber lê número no formato brasileiro. Devolve nil quando não há número.
//
// Com RawCellValue o valor vem do XML com ponto decimal ("1234.56"); em edições exportadas
// como texto vem "1.234,56". Os dois precisam funcionar, e trocar vírgula por ponto sem olhar
// transformaria "1234.56" em 123456.
func ParseBRNumber(raw string) *float64 {
	s := strings.TrimSpace(raw)
	if s == "" || s == "-" {
		return nil
	}
	if strings.Contains(s, ",") {
		// Vírgula presente ⇒ ela é o separador decimal e o ponto é milhar.
		s = strings.ReplaceAll(s, ".", "")
		s = strings.ReplaceAll(s, ",", ".")
	}
	s = strings.ReplaceAll(s, " ", "")
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil
	}
	return &v
}

// NormalizeEAN limpa o código de barras. A CMED escreve "    -     " quando não há EAN.
func NormalizeEAN(raw string) string {
	s := strings.TrimSpace(raw)
	s = strings.Trim(s, "-")
	s = strings.TrimSpace(s)
	for _, r := range s {
		if r < '0' || r > '9' {
			return ""
		}
	}
	if len(s) < 8 || len(s) > 14 {
		return ""
	}
	return s
}

// at lê a coluna pelo índice tolerando linha curta (célula vazia à direita some do slice).
func at(cells []string, idx int) string {
	if idx < 0 || idx >= len(cells) {
		return ""
	}
	return cells[idx]
}
