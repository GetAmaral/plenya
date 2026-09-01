package dto

// Cobertura de exames do paciente: quando cada exame do catálogo foi feito pela última vez.
//
// Existe para acabar com um erro concreto: o painel do catálogo quase nunca tem resultado próprio
// (o laboratório reporta os analitos filhos), então cruzar o protocolo com o prontuário olhando só
// o painel diz "nunca feito" e manda repetir exame de quem acabou de fazer.

// LabCoverageVia — de onde saiu a data.
type LabCoverageVia string

const (
	// LabCoverageNever — o paciente nunca fez, nem o exame nem nenhum analito dele.
	LabCoverageNever LabCoverageVia = "never"
	// LabCoverageOwn — o próprio exame tem resultado.
	LabCoverageOwn LabCoverageVia = "own"
	// LabCoverageChildren — quem tem resultado são os analitos filhos. É o caso normal dos painéis.
	LabCoverageChildren LabCoverageVia = "children"
)

// LabCoverageEntry — um exame do catálogo e quando ele foi feito.
type LabCoverageEntry struct {
	Code string `json:"code"`
	Name string `json:"name"`

	// LastDoneAt — AAAA-MM-DD da coleta mais recente. Vazio quando nunca foi feito.
	LastDoneAt string `json:"lastDoneAt,omitempty"`
	// DaysAgo — há quantos dias. É o número que decide se vale repetir.
	DaysAgo *int           `json:"daysAgo,omitempty"`
	Via     LabCoverageVia `json:"via"`

	// Para painel, DOIS números, e a diferença entre eles é clínica:
	//
	//   ChildrenDone      — quantos analitos vieram NA COLETA de `lastDoneAt`.
	//   ChildrenDoneEver  — quantos já vieram alguma vez, em qualquer coleta.
	//
	// Medido neste banco, "Rotina de urina" tem 14 analitos já vistos mas só 8 na coleta mais
	// recente; os outros 6 são de três meses antes. Mostrar 14 ao lado de "há 204 dias" faz o
	// médico deixar de pedir analito que não foi colhido — exatamente o erro que este endpoint
	// existe para evitar, só que ao contrário.
	ChildrenDone     int `json:"childrenDone,omitempty"`
	ChildrenDoneEver int `json:"childrenDoneEver,omitempty"`
	ChildrenTotal    int `json:"childrenTotal,omitempty"`
}

// LabCoverageResponse — a cobertura inteira, ordenada por nome.
type LabCoverageResponse struct {
	PatientID   string             `json:"patientId"`
	Entries     []LabCoverageEntry `json:"entries"`
	GeneratedAt string             `json:"generatedAt"`
}
