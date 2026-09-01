package services

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/google/uuid"

	"github.com/plenya/api/internal/models"
)

// signedPDF — o PDF que vai para o portal, com o que o prontuário precisa gravar sobre ele.
type signedPDF struct {
	Bytes      []byte
	Hash       string
	Digital    bool
	CertSerial *string
}

// signOrDegrade assina o PDF com o certificado ICP-Brasil do médico e, se a assinatura falhar,
// **re-renderiza o documento sem o selo** antes de publicar.
//
// O re-render é o ponto: o PDF traz impresso "documento assinado digitalmente" e o QR de validação.
// Publicar esse mesmo arquivo sem a assinatura embutida entregaria ao paciente um documento que
// afirma ser assinado e não é, e que falha na verificação do ITI. Melhor cair para a assinatura
// manual, que é honesta.
//
// `render(digital)` monta o PDF sabendo se o selo vai existir.
func signOrDegrade(
	sig *SignatureService,
	doctor *models.User,
	doctorID uuid.UUID,
	reason, location string,
	render func(digital bool) ([]byte, error),
) (*signedPDF, error) {
	// `tentaAssinar` é o que o PDF vai AFIRMAR sobre si mesmo. Se ele afirma e a assinatura não
	// acontece — por erro do certificado ou por não haver serviço de assinatura ligado — o
	// documento precisa ser refeito sem a afirmação.
	tentaAssinar := doctor != nil && doctor.CertificateActive && sig != nil

	pdfBytes, err := render(tentaAssinar)
	if err != nil {
		return nil, err
	}

	if tentaAssinar {
		signed, hash, sErr := sig.SignDocumentPDF(pdfBytes, doctorID, reason, location)
		if sErr == nil {
			return &signedPDF{Bytes: signed, Hash: hash, Digital: true, CertSerial: doctor.CertificateSerial}, nil
		}
		// Falhou: refaz sem o selo.
		unsigned, rErr := render(false)
		if rErr != nil {
			return nil, rErr
		}
		pdfBytes = unsigned
	}

	sum := sha256.Sum256(pdfBytes)
	return &signedPDF{Bytes: pdfBytes, Hash: hex.EncodeToString(sum[:]), Digital: false}, nil
}
