-- Atualizar score_item com conteúdo clínico: Hematócrito - Mulheres
-- ID: 32037968-f263-4e7d-ab85-ea83efd61c7b
-- Data: 2026-01-28

UPDATE score_items
SET
  clinical_relevance = 'O hematócrito representa a porcentagem do volume sanguíneo ocupado por eritrócitos (glóbulos vermelhos), sendo fundamental para avaliar a capacidade de transporte de oxigênio. Em mulheres adultas, os valores de referência são 36-48%, aproximadamente 12% menores que em homens (40-54%), devido a múltiplos fatores fisiológicos. A perda menstrual crônica é o principal determinante dessa diferença, com sangramento superior a 80 mL por ciclo podendo levar à depleção de ferro e anemia ferropriva. Estudos demonstram que mulheres com menstruação intensa apresentam hematócrito, hemoglobina e ferritina significativamente menores (p<0.05). Além da menstruação, hormônios sexuais influenciam diretamente a eritropoiese: estrogênios causam vasodilatação na microvasculatura renal (<300 μm), reduzindo o hematócrito, enquanto a menor concentração de testosterona em mulheres resulta em menor estímulo à produção eritrocitária. Valores abaixo de 36% sugerem anemia (frequentemente ferropriva em mulheres em idade reprodutiva), enquanto valores acima de 48% podem indicar desidratação, policitemia vera ou hipóxia crônica. A interpretação deve considerar ferritina sérica, saturação de transferrina e índices eritrocitários (VCM, HCM) para diagnóstico diferencial adequado.',

  patient_explanation = 'O hematócrito mede o percentual do seu sangue ocupado pelos glóbulos vermelhos, as células responsáveis por levar oxigênio para todo o corpo. Em mulheres, o valor normal fica entre 36% e 48% - um pouco menor que nos homens devido à perda de sangue durante a menstruação. Quando o hematócrito está baixo (abaixo de 36%), pode indicar anemia, que é muito comum em mulheres por deficiência de ferro causada por menstruações intensas. Sintomas incluem cansaço constante, palidez, falta de ar ao fazer esforços, fraqueza, unhas quebradiças e queda de cabelo. Valores altos (acima de 48%) são menos comuns e geralmente indicam desidratação ou excesso de glóbulos vermelhos. É importante investigar hematócrito baixo com exames de ferro (ferritina) e vitamina B12, especialmente se você tem menstruações intensas ou prolongadas. Procure seu médico se sentir cansaço excessivo, falta de ar ou palidez.',

  conduct = '**Hematócrito baixo (<36%):**
• Investigar anemia ferropriva (principal causa em mulheres):
  - Ferritina sérica, ferro sérico, saturação de transferrina, capacidade total de ligação do ferro
  - Hemograma completo com índices (VCM, HCM, CHCM, RDW)
• Avaliar perdas menstruais:
  - Questionário de sangramento menstrual (> 80 mL/ciclo = menorragia)
  - Considerar avaliação ginecológica se fluxo intenso/prolongado
• Dosagem de vitamina B12 e ácido fólico (anemias megaloblásticas)
• Reticulócitos (avaliar resposta medular)
• Considerar suplementação de ferro oral (30-60 mg ferro elementar/dia) ou IV se má absorção
• Vitamina C auxilia absorção do ferro
• Referenciar hematologista se anemia severa (Hct <30%) ou não responsiva

**Hematócrito elevado (>48%):**
• Avaliar status de hidratação (primeira causa)
• Investigar policitemia:
  - Eritropoietina sérica
  - Gasometria arterial (hipóxia crônica)
  - História de tabagismo, apneia do sono
• Considerar policitemia vera se persistentemente elevado
• Referenciar hematologista se Hct >52% ou sintomas tromboembólicos

**Recomendações gerais:**
• Evitar coleta durante período menstrual (pode subestimar valores)
• Repetir exame em 2-4 semanas se alterado
• Correlacionar com hemoglobina (Hct ≈ 3x hemoglobina)',

  last_review = NOW()
WHERE id = '32037968-f263-4e7d-ab85-ea83efd61c7b';

-- Verificar atualização
SELECT
  id,
  name,
  LENGTH(clinical_relevance) as clinical_chars,
  LENGTH(patient_explanation) as patient_chars,
  LENGTH(conduct) as conduct_chars,
  last_review
FROM score_items
WHERE id = '32037968-f263-4e7d-ab85-ea83efd61c7b';
