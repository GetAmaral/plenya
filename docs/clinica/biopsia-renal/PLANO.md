# Biópsia renal no serviço de urologia parceiro: ofício + anexo + card do paciente

Documento para a Secretaria Municipal de Saúde com as orientações que a nefrologia sugere ao
serviço de urologia do hospital parceiro, que passa a realizar a biópsia renal. Plano vivo:
decisões tomadas na conversa entram aqui na hora.

## Decisões tomadas

- **Escopo:** só rim nativo (enxerto fora).
- **Fluxo:** nefrologia encaminha ao município, que agenda no hospital. Paciente não vai direto,
  então **a urologia providencia os exames pré-operatórios** (o ofício sugere a lista).
- **Tom:** o que a nefrologia se compromete a mandar + "sugerimos, respeitando os protocolos do
  serviço". Referências numeradas no fim, nunca no corpo.
- **Laboratório de patologia:** desconhecido. Tabela com ideal + alternativa e uma frase sugerindo
  que o serviço confirme com o laboratório antes da primeira biópsia (precisa fazer IF).
- **Fragmentos:** MO é rotina; IF e ME são necessárias em casos selecionados, mas coleta-se material
  para ambas em TODAS as biópsias, para evitar novo procedimento. 4 fragmentos, inteiros por estudo, sem dividir (prática da equipe, com respaldo):
  - 2 em formol tamponado 10% (MO);
  - 1 em meio de Michel; na falta, gaze umedecida em SF, entrega no mesmo dia (IF; nunca formol);
  - 1 em glutaraldeído 2-3% refrigerado; na falta, formol em frasco separado e identificado (ME).
  - Inspeção visual de cada fragmento para confirmar cortical.
- **Desmopressina:** fora das sugestões (evidência mista; único ECR positivo em TFG preservada).
- **Card do paciente:** visual. gpt-image-2.5 só para ilustração; texto e dias em HTML -> PDF/PNG
  (padrão `scripts/deck-builder/hamstring-guia/`), porque número dentro de imagem de IA sai errado.
  Só nome genérico de medicamento. **Sem AREM no card** (cabeçalho e rodapé neutros): o card
  passa a ser documento do DRAS, que põe a identidade que quiser. Aviso em destaque: não parar nada por conta própria, falar com
  o médico que receitou.
- **Ordem no card:** os mais comuns primeiro. Grupo de 5 dias: AAS, clopidogrel, ticagrelor,
  varfarina e, por último, **prasugrel com a marca "(7 dias)"**, coerente com a SIR.

- **Card APROVADO (2026-10-29), commit b86d210e.** Fonte: `scripts/deck-builder/biopsia-renal/card.html`
  (`node render.js --png --pdf`); PDF em `docs/clinica/biopsia-renal/card-paciente-biopsia-renal.pdf`.
  Anti-inflamatório = "Evite sempre" (não prazo).

- **Exames pré-op sugeridos à urologia (decidido):** hemograma, TP/INR, TTPa, creatinina, ureia,
  urina tipo 1 e ultrassom renal **sem exigência de data** (para não travar o agendamento).
  Fora: beta-hCG e tipagem ABO/Rh.

- **Critérios para adiar: NÃO entram no ofício** (decidido). Avaliação fica com o serviço.

- **Observação/repouso pós-biópsia (decidido, revisto):** entra como "a definição cabe ao serviço;
  a título de contribuição..." citando Whittier & Korbet 2004 (até 24 h) e Lin 2021 (alta no mesmo dia
  com sangramento semelhante em selecionados). Sugestão: **pelo menos 8 h, estendida até 24 h nos de
  maior risco** (internados, LRA, PA de difícil controle, anemia).
- **Laudo (decidido):** "Solicitamos, se possível, que o laudo anatomopatológico seja entregue ao
  paciente ou encaminhado ao nefrologista solicitante, para a continuidade do tratamento."
- **Critérios para adiar** continuam fora.

## Prazos de suspensão no card

| Grupo | Medicamentos | SIR 2019 (alto risco) |
|---|---|---|
| 5 dias antes | AAS, clopidogrel, ticagrelor, varfarina, prasugrel **(7 dias)** | AAS 3-5 d · clopidogrel 5 d · ticagrelor 5 d · **prasugrel 7 d** · varfarina 5 d (INR ≤ 1,8) |
| 2 dias antes | rivaroxabana, apixabana, dabigatrana, edoxabana | apixabana/dabigatrana 4 doses (ClCr ≥ 50), mais se menor; rivaroxabana/edoxabana 2 doses |
| 1 dia antes | enoxaparina | dose plena 24 h (2 doses); profilática 1 dose |
| evite sempre | anti-inflamatórios | SIR sem recomendação; decisão da equipe (paciente renal) |

Prasugrel fica no bloco de 5 dias (card simples), por último e marcado "(7 dias)" (SIR 2019).

Fonte: Patel IJ et al. SIR Consensus Guidelines Part II. J Vasc Interv Radiol 2019;30(8):1168-1184,
Tabela 6 (https://www.jvir.org/article/S1051-0443(19)30407-5/pdf). Lida por extração de PDF com
linhas desalinhadas; **conferir visualmente a página antes de fechar o card.**
Limites laboratoriais (mesma diretriz): INR ≤ 1,5-1,8; plaquetas > 50 mil.

## Pendências (uma por vez)

1. Remetente: assinam os três nefrologistas, sócios da Nefroclínica Londrina (decidido):
   - Dr. Danilo Ramos Cunha · CRM-PR 21.875 · RQE 16.035
   - Dr. Getúlio Amaral Filho · CRM-PR 21.876 · RQE 16.038
   - Dr. Michel Lima Moro Alves · CRM-PR 36.888 · RQE 32.897
   Fontes: nefroclinica.com (Danilo, Getúlio) e apps/site/content/doctors/*.mdx (os três).
   - **Cabeçalho do ofício:** sem AREM e sem cabeçalho institucional (decidido, revisto).
   - **Destinatário:** Diretoria de Regulação da Atenção à Saúde (DRAS), Autarquia Municipal de
     Saúde (AMS) de Londrina. Em Londrina a saúde é autarquia, não "Secretaria". Dirigir ao cargo;
     a diretora citada no Planejamento Estratégico 2025 (Evilin Santos Silva Gorckis) pode ter mudado.
2. **Ofício v1 em revisão** (commit 1bffcabd): `scripts/deck-builder/biopsia-renal/oficio.html`,
   `node render.js --png` -> `docs/clinica/biopsia-renal/oficio-biopsia-renal.pdf` (2 pág. + card).
   Número do ofício e data ficam em branco para preencher.

## Referências

- Patel IJ et al. SIR Part II. JVIR 2019;30(8):1168-1184.
- Walker PD et al. (Renal Pathology Society). Practice guidelines for the renal biopsy. Mod Pathol 2004;17:1555-63.
- Mayo Clinic Laboratories. Renal Biopsy Preparation Instructions (MC1235-44): 3 fragmentos, 2 formol (MO+ME), 1 Michel/Zeus (IF, ≥ 0,6 cm, ~5 glomérulos).
- Corapi KM et al. Bleeding complications of native kidney biopsy. AJKD 2012 (nº de passagens não prediz sangramento).
- Poggio ED et al. Native kidney biopsy complications. CJASN 2020;15(11):1595-1602.
- Whittier WL, Korbet SM. Timing of complications in percutaneous renal biopsy. JASN 2004;15(1):142-7.
- Hogan JJ, Mocanu M, Berns JS. The Native Kidney Biopsy. CJASN 2016;11(2):354-62.
- Luciano RL, Moeckel GW. Update on the Native Kidney Biopsy: Core Curriculum 2019. AJKD 2019;73(3):404-415.
- Lin SY et al. Complications of Outpatient and Inpatient Renal Biopsy: SR/MA. Diagnostics 2021;11(4):651.
- Manno C et al. Desmopressin in kidney biopsy. AJKD 2011;57(6):850-5 (motivo de não sugerir DDAVP: evidência mista).
- Nasr SH et al. Paraffin IF com pronase, técnica de resgate. Kidney Int 2006.
