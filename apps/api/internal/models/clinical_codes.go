package models

// AnamneseItemCode — fonte única de verdade para score_items.anamnese_item_code
// Usado para buscar dados da anamnese do paciente via ClinicalDataService
type AnamneseItemCode = string

const (
	// Composição corporal
	AnamCodePeso            AnamneseItemCode = "PESO"
	AnamCodeAltura          AnamneseItemCode = "ALTURA"
	AnamCodeIMC             AnamneseItemCode = "IMC"
	AnamCodeBRI             AnamneseItemCode = "BRI"
	AnamCodeAbdominalHomem  AnamneseItemCode = "ABDOMINAL_HOMEM"
	AnamCodeAbdominalMulher AnamneseItemCode = "ABDOMINAL_MULHER"
	AnamCodeGorduraHomem    AnamneseItemCode = "GORDURA_CORPORAL_HOMEM"
	AnamCodeGorduraMulher   AnamneseItemCode = "GORDURA_CORPORAL_MULHER"
	AnamCodeMassaMuscular   AnamneseItemCode = "MASSA_MUSCULAR_ESQUELETICA"
	AnamCodeTaxaMetabolica  AnamneseItemCode = "TAXA_METABOLICA_BASAL"
	AnamCodeMassaGorda      AnamneseItemCode = "MASSA_GORDA_TOTAL"
	AnamCodeQuadril         AnamneseItemCode = "QUADRIL"

	// Hábitos e fatores de risco
	AnamCodeTabaco           AnamneseItemCode = "TABACO"
	AnamCodeAtividadeFisica  AnamneseItemCode = "ATIVIDADE_FISICA"
	AnamCodeExercicioFisico  AnamneseItemCode = "EXERCICIO_FISICO"

	// Doenças
	AnamCodeDoencaCV         AnamneseItemCode = "DOENCA_CARDIOVASCULAR"
	AnamCodeDoencaCVFamiliar AnamneseItemCode = "DOENCA_CARDIOVASCULAR_2"
	AnamCodeSintomas         AnamneseItemCode = "OUTROS_SINTOMAS"
)

// LabTestCode — fonte única de verdade para lab_test_definitions.code
// Usado para buscar resultados laboratoriais via ClinicalDataService
type LabTestCode = string

const (
	LabCodeLDL             LabTestCode = "PLNACA1492A"
	LabCodeHDL             LabTestCode = "PLN53A449CA"
	LabCodeColesterolTotal LabTestCode = "PLN919303A4"
	LabCodeTriglicerideos  LabTestCode = "PLNA0C5545F"
	LabCodeGlicemiaJejum   LabTestCode = "PLN9AF0BCF5"
	LabCodeHbA1c           LabTestCode = "PLN3FC5EDA6"
	LabCodeColesterolNHDL  LabTestCode = "PLN6C0B2525"
	LabCodeVLDL            LabTestCode = "PLN987B120E"
	LabCodeLDLOxidada      LabTestCode = "PLNDC7B5612"
	LabCodeRelColHDL       LabTestCode = "PLN09DBBB62"
	LabCodeRelTrigHDL      LabTestCode = "PLN73ED669D"
)
