package services

// Templates de nota clínica gerada por IA a partir do transcript da teleconsulta.
// Fonte única das seções: o AIService usa pra montar o schema (tool_use) + prompt,
// e o TelemedRecordingService usa pra parsear a saída em seções ordenadas + mapear
// pro campo SOAP correspondente. NÃO duplicar as chaves em outro lugar.

const (
	NoteFormatAnamnese = "anamnese"
	NoteFormatSOAP     = "soap"
)

// noteSectionDef — uma seção da nota. SoapTarget liga a seção ao campo da
// ClinicalNote (subjective|objective|assessment|plan) no "Inserir na nota".
type noteSectionDef struct {
	Key        string
	Titulo     string
	SoapTarget string
	Hint       string // instrução pro LLM do que entra nessa seção
}

var anamneseTemplate = []noteSectionDef{
	{"queixa_principal", "Queixa principal", "subjective", "Motivo principal da consulta, nas palavras do paciente."},
	{"hda", "História da doença atual", "subjective", "Evolução cronológica da queixa (início, duração, fatores, sintomas associados). Só o relatado."},
	{"antecedentes_pessoais", "Antecedentes pessoais", "subjective", "Doenças prévias, cirurgias, internações — atenção a renal/cardiovascular/metabólico. Só o mencionado."},
	{"antecedentes_familiares", "Antecedentes familiares", "subjective", "Doenças em familiares, se mencionadas."},
	{"medicacoes_em_uso", "Medicações em uso", "subjective", "Medicamentos citados (nome e dose SÓ se ditos). Nunca inventar dose."},
	{"alergias", "Alergias", "subjective", "Alergias mencionadas. Se o paciente negou, registrar a negação."},
	{"habitos_historia_social", "Hábitos e história social", "subjective", "Tabagismo, álcool, atividade física, sono, alimentação — só o mencionado."},
	{"revisao_sistemas", "Revisão de sistemas", "subjective", "Outros sintomas por sistema, só os abordados. Preservar negações."},
	{"objetivo_relatado", "Objetivo (relatado pelo paciente)", "objective", "APENAS dados que o paciente leu/relatou (PA aferida em casa, peso, exames trazidos), sempre como 'refere'. Default: 'não avaliado nesta teleconsulta'."},
	{"avaliacao", "Avaliação e hipóteses", "assessment", "Impressão e hipóteses ditas PELO MÉDICO. CID só se o médico citou. Não inventar diagnóstico."},
	{"plano", "Plano e conduta", "plan", "Conduta, prescrição, exames, orientações e retorno ditos PELO MÉDICO."},
}

var soapTemplate = []noteSectionDef{
	{"subjetivo", "Subjetivo", "subjective", "Relato do paciente: queixa, história, sintomas — só o dito."},
	{"objetivo", "Objetivo", "objective", "Só o relatado pelo paciente (teleconsulta, sem exame físico presencial). Default 'não avaliado nesta teleconsulta'."},
	{"avaliacao", "Avaliação", "assessment", "Impressão e hipóteses do médico. Não inventar."},
	{"plano", "Plano", "plan", "Conduta, prescrição, orientações e retorno ditos pelo médico."},
}

// noteTemplate retorna o template do formato (anamnese é o default).
func noteTemplate(format string) ([]noteSectionDef, bool) {
	switch format {
	case NoteFormatSOAP:
		return soapTemplate, true
	case NoteFormatAnamnese, "":
		return anamneseTemplate, true
	default:
		return nil, false
	}
}
