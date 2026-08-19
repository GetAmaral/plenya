package cmed

import (
	"testing"

	"github.com/xuri/excelize/v2"
)

// Cabeçalho real da planilha da CMED, com as armadilhas: acentos, o NBSP no fim de
// "DESTINAÇÃO COMERCIAL" e a ordem original das colunas.
var cabecalhoCMED = []string{
	"SUBSTÂNCIA", "CNPJ", "LABORATÓRIO", "CÓDIGO GGREM", "REGISTRO",
	"EAN 1", "EAN 2", "EAN 3", "PRODUTO", "APRESENTAÇÃO",
	"CLASSE TERAPÊUTICA", "TIPO DE PRODUTO (STATUS DO PRODUTO)", "REGIME DE PREÇO",
	"PMC 0 %", "RESTRIÇÃO HOSPITALAR", "TARJA", "DESTINAÇÃO COMERCIAL ",
}

// planilha monta um XLSX com N linhas de preâmbulo antes do cabeçalho — a CMED muda esse
// número entre edições, e é por isso que o parser procura o cabeçalho em vez de fixar a linha.
func planilha(t *testing.T, preambulo int, header []string, linhas [][]string) *excelize.File {
	t.Helper()
	f := excelize.NewFile()
	sheet := f.GetSheetList()[0]

	for i := 0; i < preambulo; i++ {
		cell, _ := excelize.CoordinatesToCellName(1, i+1)
		if err := f.SetCellValue(sheet, cell, "Nota legal da ANVISA, linha de preâmbulo"); err != nil {
			t.Fatalf("preâmbulo: %v", err)
		}
	}

	write := func(row int, valores []string) {
		for col, v := range valores {
			cell, _ := excelize.CoordinatesToCellName(col+1, row)
			if err := f.SetCellValue(sheet, cell, v); err != nil {
				t.Fatalf("escrever célula: %v", err)
			}
		}
	}

	write(preambulo+1, header)
	for i, linha := range linhas {
		write(preambulo+2+i, linha)
	}
	return f
}

func linhaCMED(ggrem, substancia, produto, apresentacao, classe, tipo, pmc, tarja, ean string) []string {
	return []string{
		substancia, "18.459.628/0001-15", "LABORATÓRIO EXEMPLO S.A.", ggrem, "1705600230032",
		ean, "    -     ", "    -     ", produto, apresentacao,
		classe, tipo, "Regulado",
		pmc, "Não", tarja, "",
	}
}

func TestParseFile(t *testing.T) {
	f := planilha(t, 41, cabecalhoCMED, [][]string{
		linhaCMED("538912020009303", "CLORIDRATO DE CIPROFLOXACINO", "CIPRO",
			"500 MG COM REV CT 2 BL AL PLAS TRANS X 07",
			"J1G1 - FLUORQUINOLONAS ORAIS", "Genérico", "45.49", "Tarja Vermelha sob restrição",
			"7891106000956"),
		linhaCMED("538912020009999", "LOSARTANA POTÁSSICA", "LOSARTANA",
			"50 MG COM REV CT BL AL PLAS TRANS X 30",
			"C9C0 - ANTAGONISTAS DE ANGIOTENSINA II", "Similar", "1.234,56", "Tarja Vermelha",
			"    -     "),
	})

	recs, err := ParseFile(f)
	if err != nil {
		t.Fatalf("ParseFile: %v", err)
	}
	if len(recs) != 2 {
		t.Fatalf("esperava 2 linhas, veio %d", len(recs))
	}

	first := recs[0]
	if first.GGREM != "538912020009303" {
		t.Errorf("GGREM = %q", first.GGREM)
	}
	if first.Substance != "CLORIDRATO DE CIPROFLOXACINO" {
		t.Errorf("substância = %q", first.Substance)
	}
	if first.TherapeuticClassCode != "J1G1" || first.TherapeuticClass != "FLUORQUINOLONAS ORAIS" {
		t.Errorf("classe = (%q, %q)", first.TherapeuticClassCode, first.TherapeuticClass)
	}
	if first.Stripe != "vermelha_restrita" {
		t.Errorf("tarja = %q", first.Stripe)
	}
	if first.EAN13 != "7891106000956" {
		t.Errorf("EAN = %q", first.EAN13)
	}
	if first.PMCPrice == nil || *first.PMCPrice != 45.49 {
		t.Errorf("preço = %v, esperado 45.49", first.PMCPrice)
	}

	// "    -     " é como a CMED escreve "sem EAN"; não pode virar código.
	if recs[1].EAN13 != "" {
		t.Errorf("EAN ausente deveria ficar vazio, veio %q", recs[1].EAN13)
	}
	// Formato brasileiro com milhar: 1.234,56 e não 123456.
	if recs[1].PMCPrice == nil || *recs[1].PMCPrice != 1234.56 {
		t.Errorf("preço BR = %v, esperado 1234.56", recs[1].PMCPrice)
	}
}

// O tamanho do preâmbulo muda a cada edição mensal — o parser tem que achar o cabeçalho.
func TestParseFile_PreambuloVariavel(t *testing.T) {
	for _, preambulo := range []int{0, 10, 41, 55} {
		f := planilha(t, preambulo, cabecalhoCMED, [][]string{
			linhaCMED("1", "PARACETAMOL", "TYLENOL", "500 MG COM CT BL X 20",
				"N2B0 - ANALGÉSICOS", "Novo", "12,90", "Tarja Sem Tarja", "7891"),
		})
		recs, err := ParseFile(f)
		if err != nil {
			t.Fatalf("preâmbulo de %d linhas: %v", preambulo, err)
		}
		if len(recs) != 1 || recs[0].Product != "TYLENOL" {
			t.Fatalf("preâmbulo de %d linhas: leitura errada (%+v)", preambulo, recs)
		}
	}
}

// Colunas fora da ordem esperada continuam sendo achadas: o mapeamento é por NOME.
func TestParseFile_ColunasForaDeOrdem(t *testing.T) {
	header := []string{"CÓDIGO GGREM", "PRODUTO", "APRESENTAÇÃO", "SUBSTÂNCIA", "TARJA"}
	f := planilha(t, 5, header, [][]string{
		{"999", "DIPIRONA", "500 MG COM CT BL X 10", "DIPIRONA SÓDICA", "Tarja Vermelha"},
	})

	recs, err := ParseFile(f)
	if err != nil {
		t.Fatalf("ParseFile: %v", err)
	}
	if len(recs) != 1 || recs[0].Product != "DIPIRONA" || recs[0].Stripe != "vermelha" {
		t.Fatalf("mapeamento por nome falhou: %+v", recs)
	}
}

// GGREM repetido dentro da mesma edição é correção publicada: a última linha vence.
func TestParseFile_GGREMRepetido(t *testing.T) {
	f := planilha(t, 3, cabecalhoCMED, [][]string{
		linhaCMED("777", "SUBSTANCIA", "VERSÃO ANTIGA", "10 MG COM CT X 10", "C1A0 - X", "Novo", "10,00", "Tarja Vermelha", "1"),
		linhaCMED("777", "SUBSTANCIA", "VERSÃO CORRIGIDA", "10 MG COM CT X 10", "C1A0 - X", "Novo", "11,00", "Tarja Vermelha", "1"),
	})

	recs, err := ParseFile(f)
	if err != nil {
		t.Fatalf("ParseFile: %v", err)
	}
	if len(recs) != 1 {
		t.Fatalf("GGREM repetido deveria colapsar em 1 linha, veio %d", len(recs))
	}
	if recs[0].Product != "VERSÃO CORRIGIDA" {
		t.Fatalf("deveria vencer a última ocorrência, veio %q", recs[0].Product)
	}
}

// Planilha sem o cabeçalho esperado tem que falhar RUIDOSAMENTE — silêncio aqui viraria
// catálogo vazio em produção sem ninguém perceber.
func TestParseFile_SemCabecalho(t *testing.T) {
	f := planilha(t, 2, []string{"COLUNA A", "COLUNA B"}, [][]string{{"x", "y"}})
	if _, err := ParseFile(f); err == nil {
		t.Fatal("planilha sem cabeçalho da CMED deveria dar erro")
	}
}

func TestParseBRNumber(t *testing.T) {
	casos := map[string]*float64{
		"1.234,56": ptr(1234.56),
		"45.49":    ptr(45.49), // valor cru do XML, com ponto decimal
		"12,90":    ptr(12.90),
		"":         nil,
		"-":        nil,
		"abc":      nil,
	}
	for raw, want := range casos {
		got := ParseBRNumber(raw)
		switch {
		case want == nil && got != nil:
			t.Errorf("ParseBRNumber(%q) = %v, esperado nil", raw, *got)
		case want != nil && (got == nil || *got != *want):
			t.Errorf("ParseBRNumber(%q) = %v, esperado %v", raw, got, *want)
		}
	}
}

func ptr(v float64) *float64 { return &v }
