-- Enriquecimento: Hematócrito - Homens
-- ID: b99d1e15-baa3-4bcb-956e-7314dbccfa82
-- Data: 2026-01-28
-- Baseado em evidências científicas: NCBI, NEJM, StatPearls

UPDATE score_items
SET
  clinical_relevance = 'O hematócrito representa a porcentagem do volume sanguíneo ocupado por eritrócitos, sendo um marcador fundamental da capacidade de transporte de oxigênio. Em homens adultos, os valores de referência variam entre 40-54%, podendo ser ligeiramente menores em idosos (NCBI NBK259, 2024). Valores baixos (<40%) indicam anemia, sugerindo perda sanguínea, deficiências nutricionais (ferro, B12, folato) ou distúrbios medulares. Valores elevados (>54%) podem refletir policitemia vera, eritrocitose secundária (DPOC, apneia do sono, altitude) ou desidratação. A relação entre hematócrito e viscosidade sanguínea é exponencial: cada aumento de 1% no hematócrito eleva a viscosidade em até 4%, com impacto significativo a partir de 60-70% (ScienceDirect Topics, 2024). Estudos demonstram que hematócrito >45% em policitemia vera aumenta significativamente o risco de eventos trombóticos (NEJM 2013: HR 3.91 para eventos cardiovasculares maiores). A interpretação deve considerar hemoglobina, VCM, contagem de leucócitos e marcadores de ferro para diagnóstico diferencial adequado.',

  patient_explanation = 'O hematócrito mede a porcentagem do seu sangue composta por glóbulos vermelhos, que são responsáveis por transportar oxigênio para todo o corpo. Em homens saudáveis, esse valor normalmente fica entre 40% e 54%. Quando está baixo, pode indicar anemia, causando cansaço, fraqueza e falta de ar. Valores muito altos podem deixar o sangue "grosso" demais, aumentando o risco de formação de coágulos, derrame ou infarto. Alterações podem ser causadas por desidratação, problemas nos pulmões, doenças do sangue ou até mesmo pela altitude onde você vive. Se seus resultados estiverem fora da faixa normal, seu médico pode solicitar outros exames para identificar a causa e orientar o tratamento adequado.',

  conduct = '**Hematócrito Baixo (<40%):**
- Investigar anemia: hemograma completo com índices (VCM, HCM, CHCM)
- Solicitar: ferritina, ferro sérico, capacidade de ligação, vitamina B12, ácido fólico
- Se anemia microcítica: investigar sangramento oculto (SOF)
- Se anemia macrocítica: considerar causas megaloblásticas ou mielodisplasia
- Encaminhar ao hematologista se Hb <10 g/dL ou refratário a tratamento

**Hematócrito Elevado (>54%):**
- Afastar desidratação: repetir exame após hidratação adequada
- Se persistente: solicitar gasometria arterial, saturação O2, eritropoetina sérica
- Investigar policitemia vera: JAK2 mutation, exame de medula óssea (se indicado)
- Avaliar causas secundárias: radiografia de tórax, polissonografia
- Meta terapêutica em PV: manter Ht <45% (reduz eventos trombóticos em 70%)
- Encaminhar ao hematologista se Ht >60% ou sintomas de hiperviscosidade

**Orientações Gerais:**
- Coleta ideal: jejum de 4h, hidratação normal (evitar desidratação/hemoconcentração)
- Correlacionar sempre com hemoglobina (Hb ≈ Ht/3)
- Monitorar fatores de risco cardiovascular em valores limítrofes',

  last_review = NOW()
WHERE id = 'b99d1e15-baa3-4bcb-956e-7314dbccfa82';

-- Verificação
SELECT
  id,
  name,
  LENGTH(clinical_relevance) as clinical_len,
  LENGTH(patient_explanation) as patient_len,
  LENGTH(conduct) as conduct_len,
  last_review
FROM score_items
WHERE id = 'b99d1e15-baa3-4bcb-956e-7314dbccfa82';
