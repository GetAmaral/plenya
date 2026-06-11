package services

import (
	"time"

	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/pdfdoc"
)

// buildIssuedDoc mapeia o IssuedDocument para o input do pacote pdfdoc.
func buildIssuedDoc(doc *models.IssuedDocument, patient *models.Patient, doctor *models.User, hasDigital bool, validationURL string) pdfdoc.IssuedDoc {
	cpf := ""
	if patient.CPF != nil {
		cpf = maskCPF(*patient.CPF)
	}
	cid := ""
	if doc.IncludesCID && doc.CIDConsent && doc.CIDCode != nil && *doc.CIDCode != "" {
		cid = "CID-10: " + *doc.CIDCode
	}
	// No modo digital, o selo ICP exibe o carimbo temporal (usa SignedAt ou o instante atual).
	sat := doc.SignedAt
	if sat == nil && hasDigital {
		n := time.Now()
		sat = &n
	}
	return pdfdoc.IssuedDoc{
		Title:   docTypeTitle(doc.Type),
		Patient: pdfdoc.Patient{Name: patient.Name, BirthInfo: patientBirthInfo(patient), CPFMasked: cpf},
		Body:    doc.Body,
		CID:     cid,
		Doctor:  pdfdoc.Doctor{Name: doctor.Name, Credentials: doctorCredentials(doctor)},
		Signature: pdfdoc.Signature{
			Digital:     hasDigital,
			SignedAt:    signedAtPT(sat),
			ValidateURL: validationURL,
			PlaceDate:   placeDatePT(doc.IssuedAt),
		},
	}
}

// RenderDocumentPDF gera os bytes do PDF de atestado/declaração/laudo pelo pipeline pdfdoc
// (substitui BuildDocumentPDF/gofpdf). Assinatura e publicação seguem em issued_document_service.
func (s *DocumentPDFService) RenderDocumentPDF(doc *models.IssuedDocument, patient *models.Patient, doctor *models.User, hasDigitalSignature bool, validationURL string) ([]byte, error) {
	return pdfdoc.RenderIssuedDocument(buildIssuedDoc(doc, patient, doctor, hasDigitalSignature, validationURL))
}
