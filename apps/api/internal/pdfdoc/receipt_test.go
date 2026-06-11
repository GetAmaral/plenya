package pdfdoc

import (
	"os"
	"testing"
)

func TestRenderReceipt(t *testing.T) {
	if !chromiumAvailable() {
		t.Skip("chromium ausente")
	}
	in := Receipt{
		Number: "2026-000123", PayerName: "Maria Helena Soares",
		AmountBRL: "R$ 850,00", Description: "consulta médica", Method: "PIX",
		PaidAt: "10/06/2026", PlaceDate: "Londrina, 10 de junho de 2026",
		Clinic: Clinic{
			LegalName: "Plenya Serviços de Saúde Ltda.", CNPJ: "66.991.259/0001-50",
			AddressLine: "Av. Gil de Abreu e Souza, 2335, Casa 634",
			AddressCont: "Bairro Esperança · Londrina/PR · 86058-100",
			Phone: "(43) 99974-8899", Email: "contato@plenyasaude.com.br", Site: "plenyasaude.com.br",
		},
	}
	pdf, err := RenderReceipt(in)
	if err != nil {
		t.Fatalf("render: %v", err)
	}
	if len(pdf) < 2000 || string(pdf[:5]) != "%PDF-" {
		t.Fatalf("não é PDF (len=%d)", len(pdf))
	}
	_ = os.WriteFile("/tmp/recibo.pdf", pdf, 0o644)
	t.Logf("OK: %d bytes", len(pdf))
}
