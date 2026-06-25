package services

import (
	"bytes"
	"crypto"
	"crypto/x509"
	"strings"
	"testing"

	"github.com/digitorus/pdf"
	"github.com/digitorus/pdfsign/sign"
	"github.com/digitorus/pdfsign/verify"

	"github.com/plenya/api/internal/pdfdoc"
)

// TestFreeTSADigiCert prova que um TSA RFC 3161 GRÁTIS (DigiCert) embute carimbo de tempo
// PAdES-T com nossa stack (digitorus). É integração: pula se o container não tiver egress.
func TestFreeTSADigiCert(t *testing.T) {
	if testing.Short() {
		t.Skip("integração de rede (TSA externo); pulado em -short")
	}
	if !chromiumPresent() {
		t.Skip("chromium ausente")
	}
	root, rootKey := mkCert(t, "AC Raiz Teste", true, nil, nil)
	inter, interKey := mkCert(t, "AC Intermediaria Teste", true, root, rootKey)
	leaf, leafKey := mkCert(t, "Dr. Teste:00000000000", false, inter, interKey)

	pdfBytes, err := pdfdoc.RenderIssuedDocument(pdfdoc.IssuedDoc{
		Title: "Atestado", Patient: pdfdoc.Patient{Name: "Paciente"},
		Body: "Corpo.", Doctor: pdfdoc.Doctor{Name: "Dr. Teste", Credentials: "CRM"},
		Signature: pdfdoc.Signature{Digital: true, SignedAt: "agora", ValidateURL: "https://x"},
	})
	if err != nil {
		t.Fatalf("render: %v", err)
	}

	sd := sign.SignData{
		Signature: sign.SignDataSignature{
			Info:       sign.SignDataSignatureInfo{Name: "Plenya", Location: "Brasil", Reason: "teste"},
			CertType:   sign.CertificationSignature,
			DocMDPPerm: sign.AllowFillingExistingFormFieldsAndSignaturesPerms,
		},
		DigestAlgorithm:   crypto.SHA256,
		Certificate:       leaf,
		Signer:            leafKey,
		CertificateChains: [][]*x509.Certificate{{leaf, inter, root}},
		TSA:               sign.TSA{URL: "http://timestamp.digicert.com"},
	}
	r := bytes.NewReader(pdfBytes)
	pr, err := pdf.NewReader(r, int64(len(pdfBytes)))
	if err != nil {
		t.Fatal(err)
	}
	var out bytes.Buffer
	if err := sign.Sign(r, &out, pr, int64(len(pdfBytes)), sd); err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "timestamp") ||
			strings.Contains(strings.ToLower(err.Error()), "dial") ||
			strings.Contains(strings.ToLower(err.Error()), "no such host") ||
			strings.Contains(strings.ToLower(err.Error()), "timeout") {
			t.Skipf("sem egress p/ TSA (esperado offline): %v", err)
		}
		t.Fatalf("sign c/ TSA: %v", err)
	}

	v, err := verify.Verify(bytes.NewReader(out.Bytes()), int64(out.Len()))
	if err != nil {
		t.Fatalf("verify: %v", err)
	}
	hasTS := false
	for _, sg := range v.Signers {
		if sg.TimeStamp != nil {
			hasTS = true
		}
	}
	if !hasTS {
		t.Fatalf("PDF assinado SEM carimbo de tempo (esperava PAdES-T)")
	}
	t.Logf("OK: carimbo de tempo DigiCert (grátis) presente — %d bytes", out.Len())
}
