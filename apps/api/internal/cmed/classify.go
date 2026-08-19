package cmed

import (
	"strings"

	"github.com/plenya/api/internal/models"
)

// Classification é a categoria regulatória atribuída a uma linha da CMED — e, junto com ela,
// o quanto essa atribuição é confiável.
type Classification struct {
	Category       models.MedicationCategory
	CategorySource string // manual | cmed_derived | cmed_fallback
	NeedsReview    bool
	IsPrescribable bool
}

// glp1Substances — GLP-1 precisa ser identificado por SUBSTÂNCIA, não por classe terapêutica.
// A CMED põe Saxenda e Wegovy em "preparações antiobesidade" e os combos com insulina em
// "outras insulinas"; classificar pela classe perderia justamente os mais prescritos.
var glp1Substances = []string{
	"SEMAGLUTIDA", "LIRAGLUTIDA", "TIRZEPATIDA", "DULAGLUTIDA", "EXENATIDA", "LIXISENATIDA",
}

// anabolicSubstances — lista curada de andrógenos/anabolizantes (Lista C5 da Portaria 344/98).
// É lista pequena e verificável um a um, ao contrário de qualquer tentativa de derivar C5 da
// tarja (que não distingue C5 de C1).
var anabolicSubstances = []string{
	"TESTOSTERONA", "NANDROLONA", "OXANDROLONA", "ESTANOZOLOL", "MESTEROLONA", "DANAZOL",
	"OXIMETOLONA", "METANDIENONA",
}

// Classify decide a categoria regulatória a partir do que a CMED oferece: substância, classe
// terapêutica e tarja. A primeira regra que casa vence.
//
// O ponto central é a HONESTIDADE: a CMED não publica as listas da Portaria 344/98, então
// várias linhas só podem receber um palpite conservador. Essas saem com CategorySource
// 'cmed_fallback' e NeedsReview=true — o catálogo sabendo onde não sabe, para a tela poder
// avisar antes de o médico prescrever.
func Classify(substance, therapeuticClassCode, stripe string) Classification {
	sub := strings.ToUpper(removeAccents(substance))
	class := strings.ToUpper(strings.TrimSpace(therapeuticClassCode))

	derived := func(cat models.MedicationCategory, review bool) Classification {
		return Classification{
			Category:       cat,
			CategorySource: models.MedCategorySourceDerived,
			NeedsReview:    review,
			IsPrescribable: true,
		}
	}

	// 1. GLP-1 por substância.
	if containsAny(sub, glp1Substances) {
		return derived(models.MedCategoryGLP1, false)
	}

	// 2. Antibacterianos pela classe terapêutica (J1*). Mais confiável que a tarja: existem
	//    antibacterianos publicados sem tarja nenhuma. J2 (antifúngicos) e J5 (antivirais)
	//    NÃO entram — a RDC 471 não é "tudo que começa com J", e chutar aqui daria validade
	//    de 10 dias a antirretroviral de uso contínuo.
	if strings.HasPrefix(class, "J1") {
		return derived(models.MedCategoryAntibiotic, false)
	}

	// 3. Anabolizantes (C5) por lista curada.
	if containsAny(sub, anabolicSubstances) {
		return derived(models.MedCategoryC5, true)
	}

	// 4. Tarja preta = Notificação de Receita A/B, que este EMR não emite. Entra no catálogo
	//    (serve para reconciliar o que o paciente usa) mas fora do autocomplete de receita.
	if stripe == models.MedStripeBlack {
		return Classification{
			Category:       models.MedCategoryAB,
			CategorySource: models.MedCategorySourceDerived,
			NeedsReview:    true,
			IsPrescribable: false,
		}
	}

	// 5. Tarja vermelha sob restrição + classe do sistema nervoso: é o desenho clássico do
	//    receituário de controle especial (C1).
	if stripe == models.MedStripeRedRestricted {
		if strings.HasPrefix(class, "N") {
			return derived(models.MedCategoryC1, false)
		}
		// Fora do SNC a retenção pode ser por outro motivo (isotretinoína, talidomida,
		// antimicrobiano). Classifica no mais restritivo e pede revisão.
		return Classification{
			Category:       models.MedCategoryC1,
			CategorySource: models.MedCategorySourceFallback,
			NeedsReview:    true,
			IsPrescribable: true,
		}
	}

	// 6. Tarja vermelha comum ou isento de prescrição: receita simples.
	if stripe == models.MedStripeRed || stripe == models.MedStripeNone {
		return derived(models.MedCategorySimple, false)
	}

	// 7. Sem tarja publicada (a CMED imprime "- (*)" em ~18% das linhas): assume simples e
	//    admite que é palpite.
	return Classification{
		Category:       models.MedCategorySimple,
		CategorySource: models.MedCategorySourceFallback,
		NeedsReview:    true,
		IsPrescribable: true,
	}
}

func containsAny(haystack string, needles []string) bool {
	for _, n := range needles {
		if strings.Contains(haystack, n) {
			return true
		}
	}
	return false
}

// NormalizeStripe traduz o texto da coluna TARJA da CMED para o vocabulário interno.
// "- (*)" e qualquer coisa não reconhecida viram "" — tarja desconhecida não se afirma.
func NormalizeStripe(raw string) string {
	switch strings.ToLower(strings.TrimSpace(removeAccents(raw))) {
	case "tarja vermelha":
		return models.MedStripeRed
	case "tarja vermelha sob restricao":
		return models.MedStripeRedRestricted
	case "tarja preta":
		return models.MedStripeBlack
	case "tarja sem tarja", "sem tarja":
		return models.MedStripeNone
	default:
		return ""
	}
}

// SplitTherapeuticClass separa "J1G1 - FLUORQUINOLONAS ORAIS" em código e descrição.
func SplitTherapeuticClass(raw string) (code, description string) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return "", ""
	}
	if i := strings.Index(raw, " - "); i > 0 {
		return strings.TrimSpace(raw[:i]), strings.TrimSpace(raw[i+3:])
	}
	return "", raw
}
