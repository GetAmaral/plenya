package dto

// ErrorResponse representa uma resposta de erro
type ErrorResponse struct {
	Error   string            `json:"error"`
	Message string            `json:"message,omitempty"`
	Details map[string]string `json:"details,omitempty"` // Para erros de validação
}
