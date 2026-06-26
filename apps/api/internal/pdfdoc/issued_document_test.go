package pdfdoc

import (
	"os"
	"testing"
)

func TestRenderIssuedDocument(t *testing.T) {
	if !chromiumAvailable() {
		t.Skip("chromium ausente")
	}
	in := IssuedDoc{
		Title:   "Atestado Médico",
		Patient: Patient{Name: "Maria Helena Soares", BirthInfo: "12/03/1979 · 47 anos", CPFMasked: "***.456.789-**"},
		Body: "Atesto, para os devidos fins, que a paciente acima identificada esteve sob meus cuidados " +
			"médicos nesta data, necessitando de afastamento de suas atividades laborais por 3 (três) dias, " +
			"a contar de 10/06/2026.\n\nPermaneço à disposição para esclarecimentos.",
		CID:    "CID-10: N18.3",
		Doctor: Doctor{Name: "Dr. Getúlio José Mattos do Amaral Filho", Credentials: "CRM-PR 21.876 · RQE 16.038 · Nefrologia"},
		Signature: Signature{Digital: true, SignedAt: "10/06/2026, 14:32 (horário de Brasília)",
			ValidateURL: "https://app.plenyasaude.com.br/documents/validate/019eb4a2-7c10-7f3a-9c21-8d4e5b60a1b2"},
	}
	pdf, err := RenderIssuedDocument(in)
	if err != nil {
		t.Fatalf("render: %v", err)
	}
	if len(pdf) < 2000 || string(pdf[:5]) != "%PDF-" {
		t.Fatalf("não é PDF (len=%d)", len(pdf))
	}
	_ = os.WriteFile("/tmp/atestado.pdf", pdf, 0o644)
	t.Logf("OK: %d bytes", len(pdf))
}

// TestRenderIssuedDocumentLong — corpo longo deve paginar (cabeçalho/rodapé em toda página,
// assinatura nunca cortada). Regressão do bug "assinatura empurrada pra fora / rodapé cortado".
func TestRenderIssuedDocumentLong(t *testing.T) {
	if !chromiumAvailable() {
		t.Skip("chromium ausente")
	}
	var body string
	for i := 0; i < 14; i++ {
		body += "Orientações clínicas detalhadas para o acompanhamento renal do paciente, incluindo " +
			"metas de pressão arterial, controle glicêmico, restrição de sódio, hidratação alvo de " +
			"2.500 ml/dia (30 ml/kg), reavaliação laboratorial periódica e sinais de alerta que " +
			"exigem retorno imediato ao consultório. Parágrafo número " + string(rune('A'+i)) + ".\n\n"
	}
	in := IssuedDoc{
		Kind:    "ORIENTAÇÕES",
		Title:   "Orientações ao Paciente",
		Patient: Patient{Name: "João Fernandes", BirthInfo: "20/07/1968 · 57 anos", CPFMasked: "***.456.789-**"},
		Body:    body,
		Doctor:  Doctor{Name: "Dr. Getúlio José Mattos do Amaral Filho", Credentials: "CRM-PR 21.876 · RQE 16.038 · Nefrologia"},
		Signature: Signature{Digital: true, SignedAt: "20/07/2026, 09:51 (horário de Brasília)",
			ValidateURL: "https://app.plenyasaude.com.br/documents/validate/019eb4a2-7c10-7f3a-9c21-8d4e5b60a1b2"},
	}
	pdf, err := RenderIssuedDocument(in)
	if err != nil {
		t.Fatalf("render: %v", err)
	}
	if len(pdf) < 2000 || string(pdf[:5]) != "%PDF-" {
		t.Fatalf("não é PDF (len=%d)", len(pdf))
	}
	_ = os.WriteFile("/tmp/orientacoes-longo.pdf", pdf, 0o644)
	t.Logf("OK: %d bytes", len(pdf))
}

// TestRenderIssuedDocumentSparse — caso do João: pág. 1 cheia, pág. 2 com poucas linhas.
// Rodapé E assinatura devem ficar no PÉ da pág. 2 (não no alto).
func TestRenderIssuedDocumentSparse(t *testing.T) {
	if !chromiumAvailable() {
		t.Skip("chromium ausente")
	}
	bodyHTML := `
<h2>Pressão arterial</h2>
<ul>
<li>Medir em casa pela manhã e à noite, anotar os valores e trazer no retorno.</li>
<li>Alvo abaixo de 130/80 mmHg na maioria das medidas.</li>
</ul>
<h2>Alimentação</h2>
<ul>
<li>Reduzir o sal para menos de 5 g por dia. Evitar embutidos, enlatados e temperos prontos.</li>
<li>Priorizar comida de verdade, com mais vegetais, e diminuir os ultraprocessados.</li>
<li>Moderar proteína de origem animal conforme combinamos na consulta.</li>
</ul>
<h2>Hidratação</h2>
<ul>
<li>Beber cerca de 2,5 litros de água ao longo do dia, salvo orientação em contrário.</li>
</ul>
<h2>Medicações</h2>
<ul>
<li>Tomar o anti-hipertensivo todos os dias no mesmo horário, mesmo com a pressão normal.</li>
<li>Não interromper nenhum remédio por conta própria sem falar comigo antes.</li>
</ul>
<h2>Atividade física</h2>
<ul>
<li>Caminhada de 30 minutos, cinco vezes por semana, em ritmo confortável.</li>
<li>Aumentar a intensidade aos poucos, conforme a tolerância.</li>
</ul>
<h2>Exames de acompanhamento</h2>
<ul>
<li>Repetir creatinina, ureia, sódio, potássio e exame de urina em 90 dias.</li>
<li>Trazer os resultados impressos no retorno.</li>
</ul>
<h2>Sinais de alerta</h2>
<ul>
<li>Procurar atendimento se notar inchaço nas pernas, falta de ar, urina muito espumosa ou queda importante do volume de urina.</li>
</ul>
<h2>Retorno</h2>
<ul>
<li>Agendar a consulta de acompanhamento em três meses, ou antes se surgir qualquer sinal de alerta.</li>
</ul>`
	in := IssuedDoc{
		Kind:     "ORIENTAÇÕES",
		Title:    "Orientações ao Paciente",
		Patient:  Patient{Name: "João Fernandes de Andrade", BirthInfo: "20/07/1968 · 57 anos"},
		BodyHTML: bodyHTML,
		Doctor:   Doctor{Name: "Dr. Getúlio José Mattos do Amaral Filho", Credentials: "CRM-PR 21.876 · RQE 16.038 · Nefrologia"},
		Signature: Signature{Digital: true, SignedAt: "25/06/2026, 18:59 (horário de Brasília)",
			ValidateURL: "https://app.plenyasaude.com.br/documents/validate/019eb4a2-7c10-7f3a-9c21-8d4e5b60a1b2"},
	}
	pdf, err := RenderIssuedDocument(in)
	if err != nil {
		t.Fatalf("render: %v", err)
	}
	if len(pdf) < 2000 || string(pdf[:5]) != "%PDF-" {
		t.Fatalf("não é PDF (len=%d)", len(pdf))
	}
	_ = os.WriteFile("/tmp/orientacoes-sparse.pdf", pdf, 0o644)
	t.Logf("OK: %d bytes", len(pdf))
}
