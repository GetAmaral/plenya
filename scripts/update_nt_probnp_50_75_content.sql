-- Update NT-proBNP (50-75 anos) with expanded clinical content
-- Expanding clinical_relevance to meet 1500-2000 character requirement

BEGIN;

UPDATE score_items
SET
  clinical_relevance = 'O NT-proBNP (peptídeo natriurético cerebral N-terminal) é o biomarcador mais validado para diagnóstico e estratificação de insuficiência cardíaca (IC) em adultos de 50-75 anos. Esta faixa etária requer pontos de corte ajustados pela idade devido ao aumento fisiológico dos peptídeos natriuréticos com o envelhecimento. As diretrizes ESC 2023 estabelecem valores de referência específicos: ≥250 pg/mL para rastreio ambulatorial (sensibilidade 88,5%, especificidade 67,8%) e ≥900 pg/mL para IC aguda no departamento de emergência. O NT-proBNP auxilia na diferenciação entre dispneia de origem cardíaca e não-cardíaca, sendo particularmente útil em pacientes com apresentação atípica. Valores <300 pg/mL têm valor preditivo negativo de 98%, excluindo IC com alta confiança. A zona cinzenta (300-900 pg/mL) requer investigação complementar com ecocardiograma. Condições que elevam falsamente o NT-proBNP incluem: fibrilação atrial (presente em 30% dos casos de IC de novo), disfunção renal (TFG <60 mL/min), idade avançada e sepse. Obesidade reduz os níveis (ajuste recomendado para IMC >30). O monitoramento seriado do NT-proBNP auxilia na titulação de terapia e prediz desfechos, com reduções >30% associadas à melhora prognóstica. Em pacientes com IC estabelecida, valores persistentemente elevados (>1000 pg/mL) indicam maior risco de hospitalização e mortalidade cardiovascular. Estudos recentes de 2024-2025 demonstram que os limiares ajustados por idade aumentam a especificidade diagnóstica de 50% para 67,8% em pacientes de 50-74 anos, reduzindo encaminhamentos desnecessários à cardiologia. A correlação entre NT-proBNP e fração de ejeção ventricular esquerda (FEVE) é moderada, sendo fundamental realizar ecocardiograma mesmo com valores limítrofes. Em contextos de emergência, a dosagem de NT-proBNP reduz o tempo de diagnóstico em 30% e melhora desfechos ao permitir início precoce de terapia dirigida.',

  updated_at = CURRENT_TIMESTAMP
WHERE id = '7998e69e-a0e0-488b-bcc3-9da32e59adfb';

COMMIT;

-- Verify character counts
SELECT
  name,
  LENGTH(clinical_relevance) as clinical_chars,
  LENGTH(patient_explanation) as patient_chars,
  LENGTH(conduct) as conduct_chars
FROM score_items
WHERE id = '7998e69e-a0e0-488b-bcc3-9da32e59adfb';
