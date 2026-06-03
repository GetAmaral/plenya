package services

import (
	"bytes"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/plenya/api/internal/models"
)

// testPrescription monta uma prescrição sintética (sem DB) para exercitar a geração de PDF.
func testPrescription(controlled bool) *models.Prescription {
	crm := "12345"
	uf := "PR"
	cpf := "12345678901"
	cat := models.MedCategorySimple
	if controlled {
		cat = models.MedCategoryC1
	}
	return &models.Prescription{
		ID:               uuid.Must(uuid.NewV7()),
		PrescriptionDate: time.Now(),
		ValidUntil:       time.Now().AddDate(0, 0, 30),
		Doctor: models.User{
			Name:  "Dr. Teste",
			CRM:   &crm,
			CRMUF: &uf,
		},
		Patient: models.Patient{
			Name: "Paciente Teste",
			CPF:  &cpf,
		},
		Medications: []models.PrescriptionMedication{{
			MedicationName:   "Medicamento X",
			ActiveIngredient: "substância x",
			Category:         cat,
			Concentration:    "20mg",
			Quantity:         30,
			QuantityInWords:  "trinta",
			Dosage:           "1 comprimido",
			Frequency:        "2x/dia",
			Route:            "VO",
			Duration:         30,
		}},
	}
}

// TestGeneratePrescriptionPDF verifica os dois modos de geração:
//   - manual (controlado): produz PDF, sem QR digital.
//   - digital (simples): produz PDF e seta QRCodeData de validação.
//
// Depende das fontes/timbrado disponíveis no container api; pula fora dele.
func TestGeneratePrescriptionPDF(t *testing.T) {
	if _, err := os.Stat("/usr/share/fonts/opensans/OpenSans-Regular.ttf"); err != nil {
		t.Skip("fontes OpenSans indisponíveis — rode dentro do container api")
	}

	svc := &PrescriptionPDFService{}

	cases := []struct {
		name       string
		controlled bool
		manual     bool
		wantQR     bool
	}{
		{"manual_controlado", true, true, false},
		{"digital_simples", false, false, true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := testPrescription(tc.controlled)
			b, err := svc.generatePDFContent(p, tc.manual)
			if err != nil {
				t.Fatalf("generatePDFContent erro: %v", err)
			}
			if len(b) < 1000 || !bytes.HasPrefix(b, []byte("%PDF")) {
				t.Fatalf("saída não parece um PDF válido (%d bytes)", len(b))
			}
			if tc.wantQR && p.QRCodeData == nil {
				t.Fatalf("modo digital deveria popular QRCodeData")
			}
			if !tc.wantQR && p.QRCodeData != nil {
				t.Fatalf("modo manual não deveria popular QRCodeData")
			}
		})
	}
}

// TestVerifySignatureUnsigned garante que VerifySignature lida com PDF não-assinado sem
// erro de servidor (reporta Signed=false), validando o binding com pdfsign/verify.
func TestVerifySignatureUnsigned(t *testing.T) {
	if _, err := os.Stat("/usr/share/fonts/opensans/OpenSans-Regular.ttf"); err != nil {
		t.Skip("fontes OpenSans indisponíveis — rode dentro do container api")
	}
	pdfBytes, err := (&PrescriptionPDFService{}).generatePDFContent(testPrescription(false), true)
	if err != nil {
		t.Fatalf("falha ao gerar PDF base: %v", err)
	}
	res, err := (&SignatureService{}).VerifySignature(pdfBytes)
	if err != nil {
		t.Fatalf("VerifySignature retornou erro inesperado: %v", err)
	}
	if res == nil {
		t.Fatal("resultado de verificação nil")
	}
	if res.Signed {
		t.Fatalf("PDF não-assinado não deveria reportar Signed=true (res=%+v)", res)
	}
}
