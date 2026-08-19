package cmed

import "github.com/plenya/api/internal/models"

// Rules são as regras de prescrição que decorrem da categoria regulatória: prazo de validade
// da receita, quantos itens cabem, por quanto tempo se pode tratar, e se exige assinatura
// digital / registro no SNCR.
//
// Estes campos existem em medication_definitions desde sempre e nunca foram preenchidos — o
// cálculo de validade vive num mapa hardcoded no prescription_service. Preenchê-los no import
// é o que permite, depois, o serviço ler a regra do catálogo em vez de adivinhar.
type Rules struct {
	ValidityDays             int
	MaxPerPrescription       int
	MaxTreatmentDays         int
	RequiresDigitalSignature bool
	RequiresSNCR             bool
}

var rulesByCategory = map[models.MedicationCategory]Rules{
	models.MedCategorySimple:     {ValidityDays: 30, MaxPerPrescription: 10, MaxTreatmentDays: 60, RequiresDigitalSignature: false, RequiresSNCR: false},
	models.MedCategoryAntibiotic: {ValidityDays: 10, MaxPerPrescription: 10, MaxTreatmentDays: 30, RequiresDigitalSignature: false, RequiresSNCR: false},
	models.MedCategoryC1:         {ValidityDays: 30, MaxPerPrescription: 3, MaxTreatmentDays: 60, RequiresDigitalSignature: true, RequiresSNCR: true},
	models.MedCategoryC5:         {ValidityDays: 30, MaxPerPrescription: 3, MaxTreatmentDays: 60, RequiresDigitalSignature: true, RequiresSNCR: true},
	models.MedCategoryGLP1:       {ValidityDays: 90, MaxPerPrescription: 10, MaxTreatmentDays: 180, RequiresDigitalSignature: false, RequiresSNCR: false},
	// Notificação de Receita A/B: o EMR não emite, mas a linha existe no catálogo. As regras
	// ficam no patamar mais restritivo para nunca subestimarem o controle.
	models.MedCategoryAB: {ValidityDays: 30, MaxPerPrescription: 3, MaxTreatmentDays: 60, RequiresDigitalSignature: true, RequiresSNCR: true},
}

// RulesFor devolve as regras da categoria. Categoria desconhecida cai no perfil simples.
func RulesFor(category models.MedicationCategory) Rules {
	if r, ok := rulesByCategory[category]; ok {
		return r
	}
	return rulesByCategory[models.MedCategorySimple]
}
