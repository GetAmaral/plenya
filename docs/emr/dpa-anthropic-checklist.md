# DPA Anthropic + transferência internacional (LGPD) — checklist operacional

> Como habilitar a IA da Anthropic (Claude) sobre dado de paciente real de forma defensável
> sob a LGPD. Itens marcados **[Sales/jurídico]** dependem de ação humana, não são self-serve.
> Pesquisa: workflow `dpa-anthropic-como-fazer` (2026-06-03). **Não processar dado real antes de
> fechar o checklist da seção 5.**

## 1. DPA — já está (quase) feito
O **DPA (Data Processing Addendum)** formaliza Plenya = controladora, Anthropic = operadora (LGPD
art. 39). O da Anthropic é **padrão e incorporado automaticamente** ao aceitar os Commercial Terms
numa **conta comercial** (Console com billing de empresa — que já temos). Não assina à parte.
Já traz SCCs da UE + proibição contratual de treinar com o conteúdo do cliente.
- Arquivar PDF no dossiê: https://www.anthropic.com/legal/data-processing-addendum +
  https://www.anthropic.com/legal/commercial-terms
- Confirmar que a chave do backend está em **org comercial**, não conta consumidor.

## 2. ZDR + BAA — um único ticket em https://claude.com/contact-sales
- **ZDR (Zero Data Retention):** sales-assisted (não é checkbox). Padrão da API retém input/output
  por até 30 dias (até 2 anos se sinalizado por violação de policy); ZDR = nada em repouso após a
  resposta. Habilitado **por organização**. Pedir como requisito (clínica/dado sensível). Doc:
  https://platform.claude.com/docs/en/manage-claude/api-and-data-retention
- **BAA (HIPAA):** opcional, reforço (não substitui LGPD). Provisiona org HIPAA dedicada.
- Pedir **ZDR + BAA + adendo de transferência ANPD (seção 3) no MESMO ticket**.

## 3. Transferência internacional (ANPD) — o item que dá trabalho [Sales/jurídico]
DPA **não basta** pra mandar dado pros EUA (LGPD art. 33). EUA **não têm adequação** da ANPD (só
UE/EEE, Res. 32/2026). As **Cláusulas-Padrão Contratuais (CPC)** da ANPD (Anexo II da **Res. CD/ANPD
nº 19/2024**) **já estão em vigor** — período de graça encerrou **23/08/2025**, então é exigível hoje.
- As SCCs da UE do DPA da Anthropic **não satisfazem automaticamente a ANPD**.
- **Ação:** pedir à Anthropic um adendo de transferência incorporando as **CPC da ANPD**
  (Plenya = exportadora/controladora, Anthropic = importadora/operadora). **[confirmar se a Anthropic
  aceita o adendo Brasil]**
- Plano B se recusarem: anexar as CPC como cláusulas equivalentes + documentar a equivalência no
  RIPD (mais frágil; "cláusulas específicas" do art. 33,II exigem aprovação prévia da ANPD = meses).
- Página ANPD: https://www.gov.br/anpd/pt-br/assuntos/assuntos-internacionais/transferencia-internacional-de-dados
- Perguntar sobre **data residency / `inference_geo`** (roteamento fora dos EUA reduz exposição).

## 4. Documentar internamente
- **Base legal art. 11, II, "f" (tutela da saúde)** — dispensa consentimento; uso só instrumental ao
  atendimento daquele paciente. Registrar no registro de operações (art. 37). (Camada distinta do
  mecanismo de transferência da seção 3 — são cumulativos.)
- **RIPD/DPIA** — obrigatório (dado sensível + transferência internacional + IA). Riscos residuais:
  retenção até 2 anos se violação de policy; acesso por autoridade dos EUA. Mitigações: ZDR, BAA,
  Messages API inline, minimização.
- **Monitor da página de subprocessadores** (https://www.anthropic.com/subprocessors) — DPA dá 15
  dias de aviso + objeção. Cadastrar num firecrawl_monitor.
- Guardar **confirmação escrita do ZDR**; certificações (SOC2, ISO 27001, ISO 42001) em
  https://trust.anthropic.com.
- ✅ Termo de teleconsulta já atualizado (cláusulas 8 IA-de-apoio + 9 transferência internacional).

## 5. Checklist "podemos ligar IA com paciente real?"
- [ ] API em org comercial (não consumidor).
- [ ] DPA + Commercial Terms arquivados.
- [ ] **ZDR confirmado por escrito** na org do backend.
- [ ] **Adendo CPC da ANPD assinado** (ou plano B documentado + datado).
- [ ] BAA + org HIPAA (se seguir esse caminho).
- [ ] Código usa **só Messages API inline** pra PHI (sem Files/Batch/Code Execution); chamadas só do
      backend Go. ✅ (nossa implementação já é assim — transcript vai inline no `content`).
- [ ] **Nenhum PHI** em nome de campo/enum/JSON schema. ✅ (schema usa chaves genéricas).
- [ ] RIPD/DPIA concluído.
- [ ] Base legal art. 11 registrada.
- [ ] Política de Privacidade atualizada.
- [ ] Monitor de subprocessadores ativo.

## 6. Riscos honestos
- DPA da Anthropic não cita LGPD/Brasil — fechar o gap **depende de Sales** (pode não ter adendo BR pronto).
- Consentimento do paciente **não** é base segura da transferência (revogável) — é só reforço; a base é a CPC.
- Retenção até 2 anos / acesso por autoridade US = riscos residuais não-elimináveis (registrar no RIPD).
- **Regra de código:** mandar PHI via Files API/Batch quebra o ZDR. Virou regra de revisão no `apps/api`.

**Único item que NÃO é self-serve:** o adendo de transferência com as CPC da ANPD — exige
jurídico da Plenya + negociação com a Anthropic. O resto (DPA, ZDR, BAA) sai pelo contact-sales.
