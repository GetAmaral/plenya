package services

import "gorm.io/gorm"

// catalogoDeExames carrega, de uma vez, o que o motor de escore precisa saber sobre o catálogo
// de exames para explicar e para comparar unidades. Sem isto cada item faria a própria consulta.
type catalogoDeExames struct {
	// nomes: código do exame → nome, para o motivo de não-avaliação dizer "PSA total" em vez
	// de "PLN1BFE6CA3".
	nomes map[string]string
	// sinonimos: código do exame → pares de unidades que o catálogo registra como fator 1
	// para AQUELE exame. É o que permite aceitar `mEq/L` = `mmol/L` no sódio (monovalente) sem
	// aceitar o mesmo num divalente, onde o número seria o dobro.
	sinonimos map[string][][2]string
}

func (c catalogoDeExames) sinonimosDe(code *string) [][2]string {
	if c.sinonimos == nil || code == nil {
		return nil
	}
	return c.sinonimos[*code]
}

func (c catalogoDeExames) nomeDe(code string) string {
	if c.nomes == nil {
		return ""
	}
	return c.nomes[code]
}

// carregaCatalogoDeExames lê nomes e sinônimos de unidade numa consulta cada. Falhar aqui não
// impede o cálculo: o motivo sai com o código cru e a guarda usa só as equivalências mecânicas.
func carregaCatalogoDeExames(db *gorm.DB) catalogoDeExames {
	cat := catalogoDeExames{
		nomes:     map[string]string{},
		sinonimos: map[string][][2]string{},
	}
	if db == nil {
		return cat
	}

	var defs []struct {
		Code string
		Name string
	}
	if err := db.Table("lab_test_definitions").
		Select("code, name").
		Where("deleted_at IS NULL").
		Scan(&defs).Error; err == nil {
		for _, d := range defs {
			cat.nomes[d.Code] = d.Name
		}
	}

	// Só fator 1: são as grafias equivalentes ("mEq/L" e "mmol/L" no sódio). Fator diferente de
	// 1 é conversão de verdade, e a ingestão já converte o resultado antes de gravar.
	var conv []struct {
		Code          string
		MainUnit      string
		SecondaryUnit string
	}
	if err := db.Table("lab_test_unit_conversions AS c").
		Select("d.code AS code, c.main_unit AS main_unit, c.secondary_unit AS secondary_unit").
		Joins("JOIN lab_test_definitions d ON d.id = c.lab_test_definition_id").
		Where("c.deleted_at IS NULL AND c.conversion_factor = 1").
		Scan(&conv).Error; err == nil {
		for _, c := range conv {
			cat.sinonimos[c.Code] = append(cat.sinonimos[c.Code], [2]string{c.MainUnit, c.SecondaryUnit})
		}
	}

	return cat
}
