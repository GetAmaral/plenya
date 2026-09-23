package services

import (
	"fmt"

	"github.com/google/uuid"
)

// Endereços de validação pública impressos no QR dos documentos clínicos.
//
// Ficam num lugar só porque a string já tinha divergido: cinco geradores usavam
// app.plenyasaude.com.br e outros quatro ainda usavam plenya.com.br, um domínio que não
// responde. O QR do atestado/declaração/laudo assinado apontava para esse domínio morto, então
// conferir o documento pelo celular era impossível — o mesmo sintoma que a receita teve por outro
// motivo (a rota pública ficava atrás do Auth do grupo).
//
// O host é constante de propósito, e não variável de ambiente: isto vai IMPRESSO em papel que
// sobrevive ao deploy. Uma env não setada viraria um QR com "localhost" numa receita assinada,
// que é pior do que não ter QR.
const validationHost = "https://app.plenyasaude.com.br"

// PrescriptionValidationURL — receituário.
func PrescriptionValidationURL(id uuid.UUID) string {
	return fmt.Sprintf("%s/prescriptions/validate/%s", validationHost, id)
}

// LabRequestValidationURL — pedido de exames.
func LabRequestValidationURL(id uuid.UUID) string {
	return fmt.Sprintf("%s/lab-requests/validate/%s", validationHost, id)
}

// IssuedDocumentValidationURL — atestado, declaração, laudo e o relatório do plano de cuidado.
func IssuedDocumentValidationURL(id uuid.UUID) string {
	return fmt.Sprintf("%s/documentos/validar/%s", validationHost, id)
}
