# Plano — Papelaria médica world-class (sistema-base de PDFs do paciente)

> **Status:** direção e design aprovados (Solicitação de Exames). Implementação do sistema-base em curso.
> **Origem:** descarte dos PDFs beta (gofpdf, pré-branding). Mockup de referência vetorial em
> `docs/branding/papelaria/render-final.py` (HTML/CSS → Chromium → PDF A4 vetorial).

## 1. Objetivo
Papelaria única, elegante e **100% vetorial** para todos os PDFs médicos que vão ao paciente,
fundamentada no brandbook Plenya. Um **sistema-base compartilhado** (cabeçalho, marca d'água,
selo de assinatura, rodapé, fontes e tokens) aplicado a cada documento.

## 2. Direção visual aprovada (direção A — "papel editorial")
- **Papel:** cream oficial `#eae7da`.
- **Cores:** exclusivamente tokens de marca — petrol `#063b4f` (texto/labels), gold `#b38645`
  (filete/bullet/acentos), cream (papel/selo), ocean `#417e8e` e sage `#92b8b4` reservados.
  Secundário = petrol com opacidade (mesmo matiz, contraste AA garantido).
- **Tipografia:** título em **Cormorant Garamond** (livre, OFL — substituto da Nalieta, que é
  **paga**; mesmo fallback do site). Corpo em **Inter** (Regular/Medium/SemiBold/Bold, todos
  embutidos). **Nunca embutir Nalieta** em PDF distribuído (licença).
- **Marca d'água:** padrão P∞ dourado da papelaria a **4%** (efeito papel timbrado).
- **Cabeçalho:** símbolo P∞ + lockup PLENYA/tagline, **filete dourado único** (sem linha dupla).
- **Selo de assinatura digital:** chip **petrol** com selo ICP-Brasil em **creme** (vetorizado),
  padding petrol; nome + CRM/RQE/especialidade; carimbo temporal; base legal; QR + validador ITI.

## 3. Regras de conteúdo (travadas)
- **Lista de exames:** uma lista na **ordem registrada** na definição (sem agrupar por sistema).
  Bullet losango dourado, sem linha de tabela entre itens.
- **Colunas dinâmicas:** ≤20 exames = 1 coluna (largura quase total); 21–40 = 2 colunas com
  **20 por coluna** (preenche a 1ª até 20, depois a 2ª); **>40 ou linha em branco = nova página**.
- **Espaçamento dinâmico:** generoso com poucos exames, comprime até o mínimo (1.5px) com 20+ na
  coluna mais cheia.
- **Linha em branco no texto do pedido = quebra de página.** Cada página é um documento
  **completo e independente** (cabeçalho + dados do paciente + indicação + rodapé + assinatura
  repetidos) — folhas podem ser destacadas (ex.: laboratório vs imagem em locais diferentes).
- **Indicação clínica opcional:** se vazia, não renderiza (não ocupa espaço).
- **Título do PDF por documento** pode diferir do menu do EMR: ex. PDF = "Solicitação de Exames",
  EMR = "Pedido de Exames".
- **Fontes:** itens 10pt; selo digital 9pt (legais 8pt); manual 10pt; sem fonte abaixo disso em
  conteúdo legível.

## 4. Assinatura e validação
- **Modo digital** (`Doctor.CertificateActive`): selo ICP-Brasil (PAdES), identificação do
  signatário (nome + CRM/RQE), carimbo temporal com fuso, base legal **MP 2.200-2/2001 + Lei
  14.063/2020**, "se impresso, autenticar na forma da lei", **QR** apontando para a página real de
  validação **por ID** (`app.plenyasaude.com.br/<doc>/validate/:id`) e menção ao **validador
  oficial do ITI** (`validar.iti.gov.br`, valida por upload do PDF).
- **Sem códigos impressos** (UUID/hash/Nº): o QR carrega o link e o ITI valida pelo PDF.
- **Modo manual** (sem certificado): data + local à esquerda; nome/CRM centralizados; **nunca**
  alega assinatura digital nem mostra QR.
- **"Emitido por (AC)"**: omitido no mockup; em produção pode vir do certificado.

## 5. 🚨 CPF do médico — regra por tipo de documento (RDC Anvisa 1.000/2025)
A representação visual **não é exigida por lei** (validade é criptográfica). O identificador
obrigatório do médico em documento é **nome + CRM** (+ RQE). **CPF do médico é omitido por
padrão** (LGPD/minimização) — o CPF completo fica no certificado ICP-Brasil e no validador do ITI.

**Exceção (RDC 1.000/2025, vigente 13/02/2026):** o **CPF do prescritor é obrigatório** em
receituários de controle:

| Documento | CPF do médico | Observação |
|---|---|---|
| Solicitação de Exames, receita comum, atestado, declaração, laudo | **Não** | identificação = nome + CRM/RQE |
| Notificação de Receita "A" (amarela) | **Sim** | entorpecentes/psicotrópicos lista A |
| Notificação de Receita "B" (azul) | **Sim** | psicotrópicos lista B (zolpidem, benzodiazepínicos) |
| Notificação de Receita Especial (branca) | **Sim** | retinoides, talidomida, imunossupressores |
| Anabolizantes | **Sim** + **CID** | esteroides anabolizantes |
| Receita de Controle Especial (2 vias) | reforço de identificação¹ | foco das fontes é CPF do paciente/comprador |
| Sujeitos só à retenção (RDC 471/2021: antimicrobianos, GLP-1) | **Não** | obrigatório virou o CPF do **paciente** |

¹ A obrigatoriedade explícita do CPF do prescritor é clara nas **Notificações**; na Receita de
Controle Especial o destaque das fontes é o CPF do paciente.

→ **O template de receituário deve incluir CPF do prescritor (e CID p/ anabolizantes)
condicionalmente por tipo.** **Antes de codificar o receituário, conferir o anexo de modelos da
própria RDC 1.000/2025** para fixar os campos por tipo (não chutar o detalhe fino). Ver também
[[receita_cfm_anvisa_sncr]] (SNCR/Anvisa, controlados, prorrogação 30/09/2026).

## 6. Pipeline técnico
- **Render:** HTML/CSS → **go-rod (Chromium headless) → PDF vetorial**, reaproveitando o padrão de
  `score_pdf_service.generatePDFFromHTML(html, a4PDFOptions())`. Extrair um renderer compartilhado.
- **Assets vetoriais** (SVG) embutidos via `file://` ou data-URI: símbolo P∞, lockup, padrão P∞,
  selo ICP (`icp-onbrand.svg`: campo petrol + marca creme), QR gerado por documento.
- **Fontes** vendoradas no `apps/api`: Cormorant Garamond (woff2 OFL) + Inter (ttf).
- **Substituir** os serviços gofpdf: lab_request_pdf, pdf_service (lab), prescription_pdf,
  document_pdf, care_plan_report, anonymous_lab_pdf, payment (recibo). Manter rotas/handlers atuais.
- **🚨 Prod:** o `Dockerfile` de produção **não tem Chromium** (só o `.dev`). Adicionar chromium +
  copiar fontes/SVGs na imagem de prod antes de qualquer deploy.

## 7. Documentos do sistema-base
1. **Solicitação de Exames** — design fechado (mockup pronto).
2. **Receituário** — comum + controlado (CPF prescritor/CID condicional, §5).
3. **Atestado / Declaração / Laudo** (`IssuedDocument`).
4. **Plano AGIR / relatório longitudinal** (`care_plan_report`).
5. **Recibo** (fiscal: razão social + CNPJ + endereço fiscal; demais docs usam endereço de atendimento).

## 8. NAP (rodapé) — dados reais
- **Plenya Serviços de Saúde Ltda.** · CNPJ 66.991.259/0001-50
- Atendimento: Av. Ayrton Senna da Silva, 500, Ed. Torre Pietra, sala 1402, Gleba Palhano,
  Londrina/PR · (43) 99974-8899 · contato@plenyasaude.com.br · plenyasaude.com.br
- Recibo usa o **endereço fiscal** (Av. Gil de Abreu e Souza, 2335 — Esperança).

## 9. Verificação (dev) e qualidade
- Vetorial (zero raster via `pdfimages`), fontes embutidas/subset (`pdffonts`), texto selecionável,
  A4 exato, contraste AA (labels petrol = 10.6:1).
- Dev primeiro (paridade dev≡prod). **Sem deploy sem ordem explícita.**

## 10. Riscos / pendências
- Chromium na imagem de prod (bloqueante p/ go-rod em prod).
- Conferir anexo RDC 1.000/2025 antes do receituário controlado.
- `Doctor.CertificateActive`/dados do certificado (CPF, AC, série) para o selo real.
- QR real por documento (hoje sample).
- Fidelidade do `pnpm generate` (não afeta PDF, mas afeta tipos).
