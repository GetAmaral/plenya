package dto

// PDFExtractionExam - Dados extraídos de um exame
type PDFExtractionExam struct {
	NomeExame       string  `json:"nomeExame"`                 // Obrigatório
	Resultado       string  `json:"resultado"`                 // Obrigatório
	Unidade         *string `json:"unidade,omitempty"`         // Opcional - omitido se não encontrado
	Material        *string `json:"material,omitempty"`        // Opcional - espécime (Sangue, Urina, Soro...)
	Metodo          *string `json:"metodo,omitempty"`          // Opcional - método analítico
	ValorReferencia *string `json:"valorReferencia,omitempty"` // Opcional - faixa de ref impressa do lab
	DataColetaExame *string `json:"dataColetaExame,omitempty"` // Opcional - data de coleta DESTE exame (YYYY-MM-DD)
}

// PDFExtractionResponse - Resposta da IA com exames extraídos + metadados do laudo
type PDFExtractionResponse struct {
	Exames      []PDFExtractionExam `json:"exames"`
	Laboratorio *string             `json:"laboratorio,omitempty"` // Nome do laboratório emissor (opcional)
	DataColeta  *string             `json:"dataColeta,omitempty"`  // Data da coleta YYYY-MM-DD (opcional)
}
