package services

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/digitorus/pdf"
	"github.com/digitorus/pdfsign/sign"

	"github.com/plenya/api/internal/pdfdoc"
)

func chromiumPresent() bool {
	for _, p := range []string{"/usr/bin/chromium-browser", "/usr/bin/chromium"} {
		if _, err := os.Stat(p); err == nil {
			return true
		}
	}
	return false
}

// mkCert cria um certificado (auto-assinado se parent==nil) com sua chave RSA.
func mkCert(t *testing.T, cn string, isCA bool, parent *x509.Certificate, parentKey *rsa.PrivateKey) (*x509.Certificate, *rsa.PrivateKey) {
	t.Helper()
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	tmpl := &x509.Certificate{
		SerialNumber:          big.NewInt(time.Now().UnixNano()),
		Subject:               pkix.Name{CommonName: cn},
		NotBefore:             time.Now().Add(-time.Hour),
		NotAfter:              time.Now().Add(24 * time.Hour),
		KeyUsage:              x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
		IsCA:                  isCA,
	}
	if isCA {
		tmpl.KeyUsage |= x509.KeyUsageCertSign
	}
	signerCert, signerKey := tmpl, key
	if parent != nil {
		signerCert, signerKey = parent, parentKey
	}
	der, err := x509.CreateCertificate(rand.Reader, tmpl, signerCert, &key.PublicKey, signerKey)
	if err != nil {
		t.Fatal(err)
	}
	c, err := x509.ParseCertificate(der)
	if err != nil {
		t.Fatal(err)
	}
	return c, key
}

func signPDFWith(t *testing.T, pdfBytes []byte, cert *x509.Certificate, signer crypto.Signer, chain []*x509.Certificate) []byte {
	t.Helper()
	sd := sign.SignData{
		Signature: sign.SignDataSignature{
			Info:       sign.SignDataSignatureInfo{Name: "Plenya", Location: "Brasil", Reason: "teste", ContactInfo: "x"},
			CertType:   sign.CertificationSignature,
			DocMDPPerm: sign.AllowFillingExistingFormFieldsAndSignaturesPerms,
		},
		DigestAlgorithm: crypto.SHA256,
		Certificate:     cert,
		Signer:          signer,
	}
	if len(chain) > 0 {
		sd.CertificateChains = [][]*x509.Certificate{append([]*x509.Certificate{cert}, chain...)}
	}
	r := bytes.NewReader(pdfBytes)
	pr, err := pdf.NewReader(r, int64(len(pdfBytes)))
	if err != nil {
		t.Fatal(err)
	}
	var out bytes.Buffer
	if err := sign.Sign(r, &out, pr, int64(len(pdfBytes)), sd); err != nil {
		t.Fatalf("sign: %v", err)
	}
	return out.Bytes()
}

// TestPAdESEmbedsCertificateChain prova que setar CertificateChains embute as ACs
// (intermediária + raiz) além da folha — o que faz o leitor montar o caminho de confiança.
// Assina o MESMO PDF só-folha vs. com-cadeia e compara: a versão com cadeia é maior.
func TestPAdESEmbedsCertificateChain(t *testing.T) {
	if !chromiumPresent() {
		t.Skip("chromium ausente")
	}
	root, rootKey := mkCert(t, "AC Raiz Teste", true, nil, nil)
	inter, interKey := mkCert(t, "AC Intermediaria Teste", true, root, rootKey)
	leaf, leafKey := mkCert(t, "Dr. Teste:00000000000", false, inter, interKey)

	pdfBytes, err := pdfdoc.RenderIssuedDocument(pdfdoc.IssuedDoc{
		Title:     "Atestado",
		Patient:   pdfdoc.Patient{Name: "Paciente Teste"},
		Body:      "Corpo do documento.",
		Doctor:    pdfdoc.Doctor{Name: "Dr. Teste", Credentials: "CRM"},
		Signature: pdfdoc.Signature{Digital: true, SignedAt: "agora", ValidateURL: "https://x"},
	})
	if err != nil {
		t.Fatalf("render: %v", err)
	}

	leafOnly := signPDFWith(t, pdfBytes, leaf, leafKey, nil)
	withChain := signPDFWith(t, pdfBytes, leaf, leafKey, []*x509.Certificate{inter, root})

	// Cada cert RSA-2048 ~ 0.9 KB DER; duas ACs a mais devem somar > 1 KB.
	if len(withChain) <= len(leafOnly)+1000 {
		t.Fatalf("cadeia não embutida: só-folha=%d, com-cadeia=%d (diff=%d)",
			len(leafOnly), len(withChain), len(withChain)-len(leafOnly))
	}
	t.Logf("OK: só-folha=%d bytes, com-cadeia=%d bytes (+%d)", len(leafOnly), len(withChain), len(withChain)-len(leafOnly))
}
