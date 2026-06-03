# Receita médica eletrônica no Brasil — validade, SNCR/Anvisa e integração CFM (sem Memed)

> **Estudo jurídico-regulatório + leitura do código Plenya.** Data: 03/06/2026.
> Fontes primárias: Planalto, DOU/Datalegis (anvisalegis), gov.br/anvisa, sistemas.cfm.org.br,
> iti.gov.br, conselhos (CFF, CRF-SP, CRF-RJ, CRF-MS, ANFARMAG). Verificação adversarial dos 3
> pontos de maior consequência. Substitui a premissa "Memed/CFM" do P3 (ver
> `plano-ux-medico-consultorio.md`) pela estratégia "ICP-Brasil em nuvem própria + SNCR direto".

## TL;DR

| Pergunta | Resposta |
|---|---|
| Nossa receita assinada já vale? | **Sim** para receita comum, atestado, laudo, solicitação de exame. O mecanismo (PAdES + ICP-Brasil) é correto. Três ressalvas de implementação (§4). **Controlado: NÃO** (SNCR é stub fake). |
| Já é obrigatório registrar receita na Anvisa (SNCR)? | **Não.** Prorrogado de 01/06/2026 para **30/09/2026** (DICOL 27/05/2026). Hoje voluntário; talonário físico vale. |
| Integrar o CFM sem Memed? | **Não há API do CFM** (é portal web manual). Caminho: assinar ICP-Brasil no próprio backend via **e-CPF em nuvem** (Certillion/IntegraICP/gov.br); controlados via **API do SNCR da Anvisa**, não CFM. |

---

## 1. Validade da receita eletrônica assinada (não controlada)

Base jurídica tripla, todas vigentes em 2026:

- **MP 2.200-2/2001, art. 10 §1º** — documento eletrônico com certificado ICP-Brasil presume-se
  verdadeiro quanto ao signatário; equivale à assinatura manuscrita.
- **Lei 14.063/2020** — classifica assinatura em **simples / avançada / qualificada** (qualificada =
  ICP-Brasil). Art. 13: **controle especial e atestado exigem qualificada**. Art. 14: demais
  documentos de saúde valem com **avançada OU qualificada**. Art. 15 reescreve o art. 35 da
  **Lei 5.991/73** (receita eletrônica vale com avançada/qualificada; validade nacional; qualificada
  obrigatória só p/ controlado e atestado).
- **Resolução CFM 2.299/2021, art. 4º** — para o médico, **todo** documento médico eletrônico deve
  usar **ICP-Brasil NGS2** (mais exigente que o piso legal). Art. 2º lista os **dados obrigatórios**
  (nome+CRM+endereço do médico; RQE se houver; nome+documento do paciente; data/hora; assinatura
  digital). Art. 5º: **a plataforma de prescrição deve estar inscrita no CRM da sede e indicar
  Diretor Técnico médico** (item de compliance da Plenya).
- **Resolução CFM 2.314/2022 (telemedicina), art. 13** — prescrição/atestado à distância exige
  certificação ICP-Brasil (relevante para nossas teleconsultas).

**Conclusão:** assinar **sempre com ICP-Brasil (qualificada, PAdES no PDF)** satisfaz Anvisa + Lei
14.063 + CFM de uma vez. É o que o nosso `SignatureService` já faz. Para não controlado a lei
aceitaria até a avançada (gov.br), mas o CFM não a considera suficiente no plano ético; logo,
ICP-Brasil é o caminho seguro para tudo.

### Farmácia é obrigada a aceitar?
A farmácia **não pode exigir via física manuscrita** nem condicionar a uma plataforma proprietária;
recusar receita eletrônica válida pode ser infração sanitária. **Porém** o farmacêutico mantém o
direito de **recusa justificada** (Cód. Ética Farmacêutica, Res. CFF 596/2014), inclusive se não
conseguir verificar a autenticidade. Não é obrigação incondicional de dispensar. **Foto/scan de
receita em papel NÃO é receita eletrônica.** Validação pelo farmacêutico: `validar.iti.gov.br` /
`assinaturadigital.iti.gov.br` (confere integridade, autoria, registro e habilitação do prescritor).
O QR na receita é atalho para esse validador; o que dá validade é a assinatura, não o QR.

---

## 2. Registro na Anvisa (SNCR) — controlados

Mapa de normas (corrigido):

- **Portaria SVS/MS 344/98** (+ RDC 970/2025 atualizando o Anexo I) = base das listas de controlados.
  *A RDC 471/2021 é sobre antimicrobianos, não consolida a 344/98.*
- **RDC 873/2024** (DOU 03/06/2024) = institui o **SNCR**, mas só a camada de **distribuição de
  numeração** entre as Vigilâncias Sanitárias (obrigatória p/ a VISA desde 01/01/2025). **Não** trata
  da emissão eletrônica pelo médico.
- **RDC 1.000/2025** (DOU 15/12/2025, vigente desde 13/02/2026) = regulamenta a **emissão eletrônica**
  de Notificações de Receita e Receitas de Controle Especial integradas ao SNCR. Art. 4º: emissão só
  por serviço **integrado ao SNCR via API**. Art. 8º: Notificação e Controle Especial exigem
  assinatura **qualificada (ICP-Brasil)**. Art. 11: receita sujeita a retenção aceita avançada ou
  qualificada. Art. 16: prazo de disponibilização do SNCR.

**Estado em 03/06/2026:** a obrigatoriedade do SNCR para o prescritor (que valeria em **01/06/2026**)
foi **PRORROGADA para 30/09/2026** (decisão unânime da DICOL/Anvisa em 27/05/2026, formalizada pela
**RDC 1.028/2026** — número a confirmar no DOU). Logo:

- Emissão eletrônica de controlado via SNCR é **voluntária** hoje; **talonário físico continua válido**.
- **Notificação A (amarela)** e **B (azul, B1/B2)** ainda só em papel até liberação da fase do SNCR.
- **Receita de Controle Especial** já pode ser eletrônica em serviço integrado, mas por opção.
- O médico **não precisa** se cadastrar/acessar o SNCR diretamente nesta fase; numeração vem da VISA
  local ou da plataforma integrada.
- Lado da farmácia: **SNGPC** segue normal (não foi substituído; coexiste com o SNCR).
- A DICOL flexibilizou a assinatura qualificada nas etapas de **acesso e requisição de numeração**,
  mantendo-a obrigatória **só na emissão final** do receituário.

**Para o EMR integrar (quando obrigar, 30/09/2026):** consumir a **API do SNCR da Anvisa** (doc a
partir de jun/2026) para numeração nacional rastreável; assinar ICP-Brasil; tratar data da assinatura
= data de emissão (art. 12); garantir emissão pelo próprio prescritor (art. 6º); suportar validação
na dispensação (art. 13).

---

## 3. Integração CFM sem Memed

**Não existe API do CFM para EMR de terceiros.** A "Prescrição Eletrônica do CFM"
(`prescricaoeletronica.cfm.org.br`) é **portal web manual**, gratuito: o médico loga (Portal de
Serviços do CFM + certificado), preenche e assina. Emite receita simples, antimicrobiano, controle
especial (branca), atestado (com selo **Atesta CFM**, Res. 2.382/2024), relatório, laudo, parecer,
solicitação de exames. **Não** emite Notificação azul/amarela. Por baixo, a assinatura é operada pelo
**Certillion** (motor ICP-Brasil de terceiro), aceitando A1/A3/nuvem.

A única coisa "integrável programaticamente" é o **motor de assinatura**, não o CFM. Estratégia para
o EMR Go, sem Memed:

### a) Assinatura — migrar de A1-arquivo para e-CPF em nuvem (curto prazo)
Motores com API REST chamável do Go:
- **Certillion** (o mesmo do CFM) — PAdES, certificado local e nuvem (BirdID/VIDaaS/SafeID/etc.).
- **IntegraICP** — broker que agrega VIDaaS/BirdID/SafeID; credencial com lifetime configurável
  (até 168h), permite **assinar em lote** após uma autorização.
- **API de assinatura gov.br** — gratuita, OAuth2, hash SHA-256, aceita ICP-Brasil qualificada.

Fluxo: o médico autoriza **uma vez** (push/biometria/OTP), o backend recebe uma credencial e assina
em lote dentro da janela. **Não há assinatura ICP-Brasil 100% headless** por exigência legal (sempre
um ato de vontade do titular), mas elimina o token físico por documento. O **certificado em nuvem do
CFM (AR CFM)** é, na prática, um VIDaaS, então encaixa no mesmo fluxo.

### b) Controlados — integrar direto ao SNCR da Anvisa (médio prazo)
Trocar o `SNCRProductionProvider` stub pela integração real com a **API do SNCR** (Anvisa). Monitorar
a publicação da documentação (a partir de jun/2026; obrigatório 30/09/2026). **Não passa pelo CFM.**

### c) Compliance de plataforma
Res. CFM 2.299/2021 art. 5º: a Plenya, como plataforma de prescrição, deve estar **inscrita no CRM**
da sede e indicar **Diretor Técnico médico**. Acionar com o Getúlio.

### Padrão técnico de assinatura (Frente 4)
- **PAdES** para PDF (envelope embutido, validável no ITI). CAdES só p/ não-PDF, XAdES só p/ XML.
- **Carimbo de tempo** (AD-RT) e LTV (AD-RA) **não são exigidos** por lei, mas recomendados para
  longevidade probatória (sobretudo controlados). Hoje assinamos AD-RB (básico), legalmente válido.
- Certificado: **qualificado em NUVEM (NGS2)**. Descartar A3 (token, inviável server-side); evitar
  A1-arquivo no servidor (frágil). PSCs: VIDaaS/Valid, BirdID/Soluti, SafeID/Safeweb, SerproID.

---

## 4. Estado do código Plenya (jun/2026)

### Já temos (e está juridicamente correto p/ não controlado)
- `signature_service.go` — PAdES/ICP-Brasil A1, SHA-256, `SignPrescriptionPDF` / `SignDocumentPDF`
  (digitorus/pdfsign). `certificate_service.go` — upload/validação do .pfx A1 (validade, KeyUsage,
  issuer ICP-Brasil, OID CPF), armazenado AES-256-GCM em `models/user.go:119-142`.
- `prescription_service.go` — CRUD escopado por `SelectedPatient`, `calculateValidUntil` (validade
  mais restritiva por categoria), max 10 meds / max 3 C1. `prescription_pdf_service.go` gera PDF
  timbrado + QR, assina, persiste `SignedPDFPath/Hash/SignedAt/CertificateSerial/QRCodeData`.
- `issued_document_service.go` — atestado/declaração/laudo assinados ICP-Brasil, com **degradação
  graciosa** p/ assinatura manual, publicação como `PatientDocument`, CID com consentimento.
- `medication_definition.go` — catálogo com flags `RequiresDigitalSignature/RequiresSNCR/...`.

### Estado das ressalvas (atualizado 2026-06-03 — commits `30e0ef9e`, `0ef5ce98`)
1. **Controlado:** ✅ RESOLVIDO — controlado (C1/C5) gera receita **manual** (imprimir, carimbar,
   assinar à mão); o stub deixou de fabricar número (`BR-STUB`); `SNCR_ENABLED=false` por default.
   PENDENTE: integração real ao SNCR + modelos de Notificação de Receita A/B e Controle Especial em
   2 vias (ver §5b, alvo 30/09/2026).
2. **Assinatura em nuvem:** ✅ e-CPF em nuvem implementado (PSC/broker, gated off por `ICP_CLOUD_ENABLED`).
   PENDENTE: **A3/token (PKCS#11)** não suportado; verificação de cadeia ICP-Brasil ainda por
   **string-match de issuer**, não validação criptográfica contra a AC-Raiz.
3. **`VerifySignature()`:** ✅ RESOLVIDO — verificação real via `pdfsign/verify` (integridade
   criptográfica + signatário), usada no validador público da prescrição.
4. **Carimbo de tempo:** ✅ suporte a PAdES-T (RFC 3161) opcional via `ICP_TSA_URL` (config-gated;
   sem ACT configurada continua AD-RB, válido). LTV (AD-RA) fica para depois.
5. **Entrega do PDF:** ✅ RESOLVIDO — prescrição publicada como `PatientDocument` e baixada por
   **endpoint autenticado** (`GET /prescriptions/:id/download`); fim do `/uploads` estático.
6. **`SignAndGenerate`:** ✅ RESOLVIDO — degradação graciosa (sem certificado ativo → modo manual).
   PENDENTE: validação fina de posologia/quantidade/duração máxima por categoria controlada.

### Arquivos-chave
`signature_service.go`, `certificate_service.go`, `models/user.go:119-142`, `sncr_service.go`,
`prescription_service.go`, `prescription_pdf_service.go`, `handlers/prescription_handler.go`,
`models/prescription.go`, `models/medication_definition.go`, `issued_document_service.go`,
`document_pdf_service.go`, `config/config.go:212-217`, `cmd/server/main.go` (rotas 825-868).

---

## 5. Roadmap recomendado (substitui "Memed/CFM" do P3)

1. ✅ **Imediato (clínico):** receita comum + atestado/laudo já válidos; controlado sai por
   impressão/assinatura manual (sem número SNCR fictício).
2. ✅ **Hardening de assinatura:** carimbo de tempo (PAdES-T opcional), `VerifySignature` real,
   prescrição servida como `PatientDocument` autenticado, degradação graciosa no `SignAndGenerate`.
3. ✅ **e-CPF em nuvem:** abstração de PSC/broker (Certillion/IntegraICP) plugada no
   `certificate_service`/`signature_service`; gated off por `ICP_CLOUD_ENABLED` até haver credencial.
4. ⏳ **Compliance de plataforma:** inscrição no CRM + Diretor Técnico (CFM 2.299/2021 art. 5º). [PENDENTE]
5. ⏳ **Controlados via SNCR (alvo 30/09/2026):** implementar `SNCRProductionProvider` real contra a
   API da Anvisa quando a doc sair; modelar Notificação de Receita e Controle Especial. [PENDENTE — §5b]

## 5c. Pendências (o que falta — não bloqueante para o uso atual)

> Consolidação das pendências em aberto após as entregas de 2026-06-03. Nenhuma impede o uso clínico
> de receita comum, atestado e laudo (já válidos), nem a receita controlada por via física.

**Externas / compliance (dependem de ação humana, lentas):**
- **Inscrição da Plenya no CRM + Diretor Técnico médico** — exigência do art. 5º da Res. CFM 2.299/2021
  para operar uma plataforma de prescrição. Acionar com o Getúlio.
- **AFE + acesso gov.br ao SNCR** — pré-requisito para a integração de controlados (ver §5b).
- **Credencial/contrato de PSC** (VIDaaS/BirdID/SafeID via Certillion ou IntegraICP) para ligar o
  e-CPF em nuvem (`ICP_CLOUD_ENABLED=true`); confirmar caminhos/payloads REST na doc do PSC.

**Técnicas (código, quando houver demanda):**
- **A3 / token (PKCS#11)** — não suportado; hoje só A1 em arquivo ou nuvem. Avaliar só se algum médico
  usar token físico.
- **Validação criptográfica da cadeia ICP-Brasil** — hoje o reconhecimento do emissor é por
  string-match (lista de ACs); o ideal é validar a cadeia até a AC-Raiz da ICP-Brasil. A validação de
  confiança/revogação oficial para terceiros continua sendo o validador do ITI.
- **LTV (PAdES-A / AD-RA)** — embutir dados de revogação (OCSP/LCR) para verificabilidade de longo prazo
  além do carimbo de tempo; recomendável sobretudo para controlados.
- **Integração SNCR + Notificação de Receita A/B + Controle Especial em 2 vias** — quando a API abrir
  (§5b); inclui validação fina de posologia/quantidade/duração máxima por categoria controlada.

## 5b. Plano de retomada — integração ao SNCR (quando a Anvisa abrir a API)

> **Ler isto quando a Anvisa publicar a documentação da API do SNCR.** Gatilho de monitoramento:
> `https://www.gov.br/anvisa/pt-br/assuntos/medicamentos/controlados/sncr` (a fase de integração das
> plataformas de prescrição começa em **junho/2026**; obrigatoriedade em **30/09/2026**). Hoje
> (jun/2026) a API de integração para emissores **ainda não foi publicada** e o acesso será
> **vinculado à AFE** + login gov.br (posicionamento Anvisa 13/05/2026). Por isso a receita de
> controlado, por ora, é **impressa, carimbada e assinada à mão** (ver §4, estado do código).

**Pré-requisitos que travam, resolver em paralelo (são lentos):**
1. **AFE / acesso gov.br ao SNCR** para a Plenya como serviço de prescrição (confirmar qual o
   requisito exato para clínica/serviço de prescrição, não só farmácia).
2. **Inscrição da plataforma no CRM** + indicação de **Diretor Técnico médico** (CFM 2.299/2021
   art. 5º). Item de compliance independente do SNCR, mas exigido para operar a plataforma.

**Onde plugar no código (já preparado nesta entrega):**
- `apps/api/internal/services/sncr_service.go` → implementar `SNCRProductionProvider.RequestPrescriptionNumber`
  / `MarkAsUsed` / `GetStatus` contra os endpoints reais da Anvisa (hoje retornam erro "not available yet").
  A interface `SNCRProvider` e o switch stub/produção já existem.
- Ligar por env: `SNCR_PRODUCTION_MODE=true`, `SNCR_API_URL`, `SNCR_API_KEY` (config em `config.go`
  `SNCRConfig`). **`SNCR_ENABLED` está `false` por default** desde esta entrega (não gera número fictício).
- `apps/api/internal/services/prescription_pdf_service.go` → quando a integração existir, reverter o
  fluxo de controlado de "impressão + assinatura manual" para **emissão eletrônica assinada + numeração
  SNCR** (hoje controlado força `signature_mode = manual`). O ponto de decisão é `resolveSignatureMode`.
- Assinatura: Notificação de Receita e Receita de Controle Especial exigem **ICP-Brasil qualificada**
  (art. 8º). Já temos PAdES ICP-Brasil + **e-CPF em nuvem** (entrega desta sessão) + carimbo de tempo
  opcional. Nada a fazer do lado da assinatura além de ligar.

**Checklist técnico da emissão eletrônica de controlado (RDC 1.000/2025):**
- [ ] Numeração única do SNCR vinculada ao prescritor, rastreável e não-reutilizável (art. 4º/6º).
- [ ] Assinatura qualificada ICP-Brasil na emissão; data da assinatura = data de emissão (art. 12).
- [ ] Emissão pelo próprio prescritor (vedada por terceiros, art. 6º).
- [ ] Suporte à validação na dispensação (ITI + numeração no SNCR + registro de uso, art. 13).
- [ ] Notificação A (amarela) / B (azul) só quando a fase específica do SNCR liberar.

## Pontos a reconfirmar
- Número/data exatos no DOU da RDC de prorrogação (indício: **RDC 1.028/2026, 01/06/2026**).
- Publicação da documentação da **API do SNCR** pela Anvisa (a partir de jun/2026) e se haverá sandbox.
- Se o **Atesta CFM** admite credenciamento de sistemas de terceiros (ou se é selo só intra-plataforma).
- Limite de assinaturas por autorização no e-CPF nuvem (varia por PSC/contrato).
