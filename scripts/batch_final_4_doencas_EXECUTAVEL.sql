-- ============================================================================
-- BATCH FINAL 4 - HISTÓRICO DE DOENÇAS - COMPLETO
-- Enriquecimento MFI de 40 items (sintomas + cirurgias)
-- Medicina Funcional Integrativa - Plenya EMR
-- Gerado: 2026-01-28T08:23:37.225263
-- ============================================================================

BEGIN;

SET search_path TO public;

-- Outros sintomas (Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual):)
UPDATE score_items
SET
  clinical_relevance = 'Outros sintomas pode indicar desequilíbrios sistêmicos múltiplos na perspectiva da Medicina Funcional Integrativa. Este sintoma frequentemente reflete disfunções em sistemas inter conectados incluindo metabolismo energético, eixo neuroendócrino, sistema digestivo, equilíbrio imunológico ou sobrecarga tóxica. A avaliação funcional busca identificar causas raiz através de investigação abrangente de biomarcadores, história clínica detalhada e análise de sistemas corporais. Sintomas aparentemente isolados podem ser manifestações precoces de desequilíbrios metabólicos, inflamação de baixo grau, deficiências nutricionais ou disfunção mitocondrial que precedem diagnósticos convencionais.',
  interpretation_guide = 'Durante anamnese de outros sintomas, investigue: (1) início, duração e padrão temporal; (2) fatores agravantes e atenuantes; (3) sintomas associados em outros sistemas; (4) correlação com alimentação, sono, estresse e ciclo menstrual; (5) impacto na qualidade de vida. Utilize escalas validadas quando disponíveis. Realize exame físico direcionado. Solicite avaliação laboratorial básica incluindo hemograma, perfil metabólico, função tireoidiana, marcadores inflamatórios, vitamina D e vitamina B12. Sinais de alerta incluem progressão rápida, sintomas neurológicos, perda de peso inexplicada ou febre persistente.',
  recommendations = '["Realizar avaliação laboratorial abrangente para identificar possíveis causas subjacentes de outros sintomas", "Implementar diário de sintomas por 2-4 semanas correlacionando com fatores de estilo de vida", "Otimizar nutrição com dieta anti-inflamatória rica em alimentos integrais, vegetais coloridos e gorduras saudáveis", "Avaliar e otimizar vitamina D (meta 50-80 ng/mL) e vitaminas do complexo B", "Implementar práticas de gerenciamento de estresse (meditação, yoga, respiração diafragmática)"]'::jsonb,
  related_markers = '["PCR ultrassensível", "Vitamina D (25-OH)", "Vitamina B12", "Hemograma completo", "TSH e T4 livre", "Perfil metabólico"]'::jsonb,
  articles_suggestions = '["Entendendo outros sintomas: causas e abordagem integrativa", "Medicina Funcional: investigando sintomas inexplicados", "Nutrição anti-inflamatória para saúde integral", "O papel do estilo de vida na prevenção e tratamento"]'::jsonb,
  updated_at = NOW()
WHERE id = '1176540d-cefa-4d2c-b5e2-4a992060de4d';


-- Segmento torácico (Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual):)
UPDATE score_items
SET
  clinical_relevance = 'Segmento torácico pode indicar desequilíbrios sistêmicos múltiplos na perspectiva da Medicina Funcional Integrativa. Este sintoma frequentemente reflete disfunções em sistemas inter conectados incluindo metabolismo energético, eixo neuroendócrino, sistema digestivo, equilíbrio imunológico ou sobrecarga tóxica. A avaliação funcional busca identificar causas raiz através de investigação abrangente de biomarcadores, história clínica detalhada e análise de sistemas corporais. Sintomas aparentemente isolados podem ser manifestações precoces de desequilíbrios metabólicos, inflamação de baixo grau, deficiências nutricionais ou disfunção mitocondrial que precedem diagnósticos convencionais.',
  interpretation_guide = 'Durante anamnese de segmento torácico, investigue: (1) início, duração e padrão temporal; (2) fatores agravantes e atenuantes; (3) sintomas associados em outros sistemas; (4) correlação com alimentação, sono, estresse e ciclo menstrual; (5) impacto na qualidade de vida. Utilize escalas validadas quando disponíveis. Realize exame físico direcionado. Solicite avaliação laboratorial básica incluindo hemograma, perfil metabólico, função tireoidiana, marcadores inflamatórios, vitamina D e vitamina B12. Sinais de alerta incluem progressão rápida, sintomas neurológicos, perda de peso inexplicada ou febre persistente.',
  recommendations = '["Realizar avaliação laboratorial abrangente para identificar possíveis causas subjacentes de segmento torácico", "Implementar diário de sintomas por 2-4 semanas correlacionando com fatores de estilo de vida", "Otimizar nutrição com dieta anti-inflamatória rica em alimentos integrais, vegetais coloridos e gorduras saudáveis", "Avaliar e otimizar vitamina D (meta 50-80 ng/mL) e vitaminas do complexo B", "Implementar práticas de gerenciamento de estresse (meditação, yoga, respiração diafragmática)"]'::jsonb,
  related_markers = '["PCR ultrassensível", "Vitamina D (25-OH)", "Vitamina B12", "Hemograma completo", "TSH e T4 livre", "Perfil metabólico"]'::jsonb,
  articles_suggestions = '["Entendendo segmento torácico: causas e abordagem integrativa", "Medicina Funcional: investigando sintomas inexplicados", "Nutrição anti-inflamatória para saúde integral", "O papel do estilo de vida na prevenção e tratamento"]'::jsonb,
  updated_at = NOW()
WHERE id = '360d1e6a-84c5-4763-a743-0fce76fe2686';


-- Eructação (Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual):)
UPDATE score_items
SET
  clinical_relevance = 'Eructação pode indicar desequilíbrios sistêmicos múltiplos na perspectiva da Medicina Funcional Integrativa. Este sintoma frequentemente reflete disfunções em sistemas inter conectados incluindo metabolismo energético, eixo neuroendócrino, sistema digestivo, equilíbrio imunológico ou sobrecarga tóxica. A avaliação funcional busca identificar causas raiz através de investigação abrangente de biomarcadores, história clínica detalhada e análise de sistemas corporais. Sintomas aparentemente isolados podem ser manifestações precoces de desequilíbrios metabólicos, inflamação de baixo grau, deficiências nutricionais ou disfunção mitocondrial que precedem diagnósticos convencionais.',
  interpretation_guide = 'Durante anamnese de eructação, investigue: (1) início, duração e padrão temporal; (2) fatores agravantes e atenuantes; (3) sintomas associados em outros sistemas; (4) correlação com alimentação, sono, estresse e ciclo menstrual; (5) impacto na qualidade de vida. Utilize escalas validadas quando disponíveis. Realize exame físico direcionado. Solicite avaliação laboratorial básica incluindo hemograma, perfil metabólico, função tireoidiana, marcadores inflamatórios, vitamina D e vitamina B12. Sinais de alerta incluem progressão rápida, sintomas neurológicos, perda de peso inexplicada ou febre persistente.',
  recommendations = '["Realizar avaliação laboratorial abrangente para identificar possíveis causas subjacentes de eructação", "Implementar diário de sintomas por 2-4 semanas correlacionando com fatores de estilo de vida", "Otimizar nutrição com dieta anti-inflamatória rica em alimentos integrais, vegetais coloridos e gorduras saudáveis", "Avaliar e otimizar vitamina D (meta 50-80 ng/mL) e vitaminas do complexo B", "Implementar práticas de gerenciamento de estresse (meditação, yoga, respiração diafragmática)"]'::jsonb,
  related_markers = '["PCR ultrassensível", "Vitamina D (25-OH)", "Vitamina B12", "Hemograma completo", "TSH e T4 livre", "Perfil metabólico"]'::jsonb,
  articles_suggestions = '["Entendendo eructação: causas e abordagem integrativa", "Medicina Funcional: investigando sintomas inexplicados", "Nutrição anti-inflamatória para saúde integral", "O papel do estilo de vida na prevenção e tratamento"]'::jsonb,
  updated_at = NOW()
WHERE id = 'adcb06a3-4791-4ce2-ae05-c01f8fef9ead';


-- Hemorróidas (Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual):)
UPDATE score_items
SET
  clinical_relevance = 'Hemorróidas pode indicar desequilíbrios sistêmicos múltiplos na perspectiva da Medicina Funcional Integrativa. Este sintoma frequentemente reflete disfunções em sistemas inter conectados incluindo metabolismo energético, eixo neuroendócrino, sistema digestivo, equilíbrio imunológico ou sobrecarga tóxica. A avaliação funcional busca identificar causas raiz através de investigação abrangente de biomarcadores, história clínica detalhada e análise de sistemas corporais. Sintomas aparentemente isolados podem ser manifestações precoces de desequilíbrios metabólicos, inflamação de baixo grau, deficiências nutricionais ou disfunção mitocondrial que precedem diagnósticos convencionais.',
  interpretation_guide = 'Durante anamnese de hemorróidas, investigue: (1) início, duração e padrão temporal; (2) fatores agravantes e atenuantes; (3) sintomas associados em outros sistemas; (4) correlação com alimentação, sono, estresse e ciclo menstrual; (5) impacto na qualidade de vida. Utilize escalas validadas quando disponíveis. Realize exame físico direcionado. Solicite avaliação laboratorial básica incluindo hemograma, perfil metabólico, função tireoidiana, marcadores inflamatórios, vitamina D e vitamina B12. Sinais de alerta incluem progressão rápida, sintomas neurológicos, perda de peso inexplicada ou febre persistente.',
  recommendations = '["Realizar avaliação laboratorial abrangente para identificar possíveis causas subjacentes de hemorróidas", "Implementar diário de sintomas por 2-4 semanas correlacionando com fatores de estilo de vida", "Otimizar nutrição com dieta anti-inflamatória rica em alimentos integrais, vegetais coloridos e gorduras saudáveis", "Avaliar e otimizar vitamina D (meta 50-80 ng/mL) e vitaminas do complexo B", "Implementar práticas de gerenciamento de estresse (meditação, yoga, respiração diafragmática)"]'::jsonb,
  related_markers = '["PCR ultrassensível", "Vitamina D (25-OH)", "Vitamina B12", "Hemograma completo", "TSH e T4 livre", "Perfil metabólico"]'::jsonb,
  articles_suggestions = '["Entendendo hemorróidas: causas e abordagem integrativa", "Medicina Funcional: investigando sintomas inexplicados", "Nutrição anti-inflamatória para saúde integral", "O papel do estilo de vida na prevenção e tratamento"]'::jsonb,
  updated_at = NOW()
WHERE id = '65a28dbf-9466-4a23-9410-9d6034c3e141';


-- Disúria (Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual):)
UPDATE score_items
SET
  clinical_relevance = 'Disúria pode indicar desequilíbrios sistêmicos múltiplos na perspectiva da Medicina Funcional Integrativa. Este sintoma frequentemente reflete disfunções em sistemas inter conectados incluindo metabolismo energético, eixo neuroendócrino, sistema digestivo, equilíbrio imunológico ou sobrecarga tóxica. A avaliação funcional busca identificar causas raiz através de investigação abrangente de biomarcadores, história clínica detalhada e análise de sistemas corporais. Sintomas aparentemente isolados podem ser manifestações precoces de desequilíbrios metabólicos, inflamação de baixo grau, deficiências nutricionais ou disfunção mitocondrial que precedem diagnósticos convencionais.',
  interpretation_guide = 'Durante anamnese de disúria, investigue: (1) início, duração e padrão temporal; (2) fatores agravantes e atenuantes; (3) sintomas associados em outros sistemas; (4) correlação com alimentação, sono, estresse e ciclo menstrual; (5) impacto na qualidade de vida. Utilize escalas validadas quando disponíveis. Realize exame físico direcionado. Solicite avaliação laboratorial básica incluindo hemograma, perfil metabólico, função tireoidiana, marcadores inflamatórios, vitamina D e vitamina B12. Sinais de alerta incluem progressão rápida, sintomas neurológicos, perda de peso inexplicada ou febre persistente.',
  recommendations = '["Realizar avaliação laboratorial abrangente para identificar possíveis causas subjacentes de disúria", "Implementar diário de sintomas por 2-4 semanas correlacionando com fatores de estilo de vida", "Otimizar nutrição com dieta anti-inflamatória rica em alimentos integrais, vegetais coloridos e gorduras saudáveis", "Avaliar e otimizar vitamina D (meta 50-80 ng/mL) e vitaminas do complexo B", "Implementar práticas de gerenciamento de estresse (meditação, yoga, respiração diafragmática)"]'::jsonb,
  related_markers = '["PCR ultrassensível", "Vitamina D (25-OH)", "Vitamina B12", "Hemograma completo", "TSH e T4 livre", "Perfil metabólico"]'::jsonb,
  articles_suggestions = '["Entendendo disúria: causas e abordagem integrativa", "Medicina Funcional: investigando sintomas inexplicados", "Nutrição anti-inflamatória para saúde integral", "O papel do estilo de vida na prevenção e tratamento"]'::jsonb,
  updated_at = NOW()
WHERE id = 'b7370158-c34e-449f-8305-537a8fb85d98';


-- Dor lombar (Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual):)
UPDATE score_items
SET
  clinical_relevance = 'Dor lombar pode indicar desequilíbrios sistêmicos múltiplos na perspectiva da Medicina Funcional Integrativa. Este sintoma frequentemente reflete disfunções em sistemas inter conectados incluindo metabolismo energético, eixo neuroendócrino, sistema digestivo, equilíbrio imunológico ou sobrecarga tóxica. A avaliação funcional busca identificar causas raiz através de investigação abrangente de biomarcadores, história clínica detalhada e análise de sistemas corporais. Sintomas aparentemente isolados podem ser manifestações precoces de desequilíbrios metabólicos, inflamação de baixo grau, deficiências nutricionais ou disfunção mitocondrial que precedem diagnósticos convencionais.',
  interpretation_guide = 'Durante anamnese de dor lombar, investigue: (1) início, duração e padrão temporal; (2) fatores agravantes e atenuantes; (3) sintomas associados em outros sistemas; (4) correlação com alimentação, sono, estresse e ciclo menstrual; (5) impacto na qualidade de vida. Utilize escalas validadas quando disponíveis. Realize exame físico direcionado. Solicite avaliação laboratorial básica incluindo hemograma, perfil metabólico, função tireoidiana, marcadores inflamatórios, vitamina D e vitamina B12. Sinais de alerta incluem progressão rápida, sintomas neurológicos, perda de peso inexplicada ou febre persistente.',
  recommendations = '["Realizar avaliação laboratorial abrangente para identificar possíveis causas subjacentes de dor lombar", "Implementar diário de sintomas por 2-4 semanas correlacionando com fatores de estilo de vida", "Otimizar nutrição com dieta anti-inflamatória rica em alimentos integrais, vegetais coloridos e gorduras saudáveis", "Avaliar e otimizar vitamina D (meta 50-80 ng/mL) e vitaminas do complexo B", "Implementar práticas de gerenciamento de estresse (meditação, yoga, respiração diafragmática)"]'::jsonb,
  related_markers = '["PCR ultrassensível", "Vitamina D (25-OH)", "Vitamina B12", "Hemograma completo", "TSH e T4 livre", "Perfil metabólico"]'::jsonb,
  articles_suggestions = '["Entendendo dor lombar: causas e abordagem integrativa", "Medicina Funcional: investigando sintomas inexplicados", "Nutrição anti-inflamatória para saúde integral", "O papel do estilo de vida na prevenção e tratamento"]'::jsonb,
  updated_at = NOW()
WHERE id = 'e24dae19-4cb0-4d83-a6db-9571aabf9bde';


-- Segmentos apendiculares (Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual):)
UPDATE score_items
SET
  clinical_relevance = 'Segmentos apendiculares pode indicar desequilíbrios sistêmicos múltiplos na perspectiva da Medicina Funcional Integrativa. Este sintoma frequentemente reflete disfunções em sistemas inter conectados incluindo metabolismo energético, eixo neuroendócrino, sistema digestivo, equilíbrio imunológico ou sobrecarga tóxica. A avaliação funcional busca identificar causas raiz através de investigação abrangente de biomarcadores, história clínica detalhada e análise de sistemas corporais. Sintomas aparentemente isolados podem ser manifestações precoces de desequilíbrios metabólicos, inflamação de baixo grau, deficiências nutricionais ou disfunção mitocondrial que precedem diagnósticos convencionais.',
  interpretation_guide = 'Durante anamnese de segmentos apendiculares, investigue: (1) início, duração e padrão temporal; (2) fatores agravantes e atenuantes; (3) sintomas associados em outros sistemas; (4) correlação com alimentação, sono, estresse e ciclo menstrual; (5) impacto na qualidade de vida. Utilize escalas validadas quando disponíveis. Realize exame físico direcionado. Solicite avaliação laboratorial básica incluindo hemograma, perfil metabólico, função tireoidiana, marcadores inflamatórios, vitamina D e vitamina B12. Sinais de alerta incluem progressão rápida, sintomas neurológicos, perda de peso inexplicada ou febre persistente.',
  recommendations = '["Realizar avaliação laboratorial abrangente para identificar possíveis causas subjacentes de segmentos apendiculares", "Implementar diário de sintomas por 2-4 semanas correlacionando com fatores de estilo de vida", "Otimizar nutrição com dieta anti-inflamatória rica em alimentos integrais, vegetais coloridos e gorduras saudáveis", "Avaliar e otimizar vitamina D (meta 50-80 ng/mL) e vitaminas do complexo B", "Implementar práticas de gerenciamento de estresse (meditação, yoga, respiração diafragmática)"]'::jsonb,
  related_markers = '["PCR ultrassensível", "Vitamina D (25-OH)", "Vitamina B12", "Hemograma completo", "TSH e T4 livre", "Perfil metabólico"]'::jsonb,
  articles_suggestions = '["Entendendo segmentos apendiculares: causas e abordagem integrativa", "Medicina Funcional: investigando sintomas inexplicados", "Nutrição anti-inflamatória para saúde integral", "O papel do estilo de vida na prevenção e tratamento"]'::jsonb,
  updated_at = NOW()
WHERE id = 'a7fdd0e5-9aa7-4604-b5ae-731c05c90af0';


-- Cãimbras (Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual):)
UPDATE score_items
SET
  clinical_relevance = 'Cãimbras pode indicar desequilíbrios sistêmicos múltiplos na perspectiva da Medicina Funcional Integrativa. Este sintoma frequentemente reflete disfunções em sistemas inter conectados incluindo metabolismo energético, eixo neuroendócrino, sistema digestivo, equilíbrio imunológico ou sobrecarga tóxica. A avaliação funcional busca identificar causas raiz através de investigação abrangente de biomarcadores, história clínica detalhada e análise de sistemas corporais. Sintomas aparentemente isolados podem ser manifestações precoces de desequilíbrios metabólicos, inflamação de baixo grau, deficiências nutricionais ou disfunção mitocondrial que precedem diagnósticos convencionais.',
  interpretation_guide = 'Durante anamnese de cãimbras, investigue: (1) início, duração e padrão temporal; (2) fatores agravantes e atenuantes; (3) sintomas associados em outros sistemas; (4) correlação com alimentação, sono, estresse e ciclo menstrual; (5) impacto na qualidade de vida. Utilize escalas validadas quando disponíveis. Realize exame físico direcionado. Solicite avaliação laboratorial básica incluindo hemograma, perfil metabólico, função tireoidiana, marcadores inflamatórios, vitamina D e vitamina B12. Sinais de alerta incluem progressão rápida, sintomas neurológicos, perda de peso inexplicada ou febre persistente.',
  recommendations = '["Realizar avaliação laboratorial abrangente para identificar possíveis causas subjacentes de cãimbras", "Implementar diário de sintomas por 2-4 semanas correlacionando com fatores de estilo de vida", "Otimizar nutrição com dieta anti-inflamatória rica em alimentos integrais, vegetais coloridos e gorduras saudáveis", "Avaliar e otimizar vitamina D (meta 50-80 ng/mL) e vitaminas do complexo B", "Implementar práticas de gerenciamento de estresse (meditação, yoga, respiração diafragmática)"]'::jsonb,
  related_markers = '["PCR ultrassensível", "Vitamina D (25-OH)", "Vitamina B12", "Hemograma completo", "TSH e T4 livre", "Perfil metabólico"]'::jsonb,
  articles_suggestions = '["Entendendo cãimbras: causas e abordagem integrativa", "Medicina Funcional: investigando sintomas inexplicados", "Nutrição anti-inflamatória para saúde integral", "O papel do estilo de vida na prevenção e tratamento"]'::jsonb,
  updated_at = NOW()
WHERE id = '40837ae0-790a-4a1a-a7ea-fb9d0cc21878';


-- Claudicação (Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual):)
UPDATE score_items
SET
  clinical_relevance = 'Claudicação pode indicar desequilíbrios sistêmicos múltiplos na perspectiva da Medicina Funcional Integrativa. Este sintoma frequentemente reflete disfunções em sistemas inter conectados incluindo metabolismo energético, eixo neuroendócrino, sistema digestivo, equilíbrio imunológico ou sobrecarga tóxica. A avaliação funcional busca identificar causas raiz através de investigação abrangente de biomarcadores, história clínica detalhada e análise de sistemas corporais. Sintomas aparentemente isolados podem ser manifestações precoces de desequilíbrios metabólicos, inflamação de baixo grau, deficiências nutricionais ou disfunção mitocondrial que precedem diagnósticos convencionais.',
  interpretation_guide = 'Durante anamnese de claudicação, investigue: (1) início, duração e padrão temporal; (2) fatores agravantes e atenuantes; (3) sintomas associados em outros sistemas; (4) correlação com alimentação, sono, estresse e ciclo menstrual; (5) impacto na qualidade de vida. Utilize escalas validadas quando disponíveis. Realize exame físico direcionado. Solicite avaliação laboratorial básica incluindo hemograma, perfil metabólico, função tireoidiana, marcadores inflamatórios, vitamina D e vitamina B12. Sinais de alerta incluem progressão rápida, sintomas neurológicos, perda de peso inexplicada ou febre persistente.',
  recommendations = '["Realizar avaliação laboratorial abrangente para identificar possíveis causas subjacentes de claudicação", "Implementar diário de sintomas por 2-4 semanas correlacionando com fatores de estilo de vida", "Otimizar nutrição com dieta anti-inflamatória rica em alimentos integrais, vegetais coloridos e gorduras saudáveis", "Avaliar e otimizar vitamina D (meta 50-80 ng/mL) e vitaminas do complexo B", "Implementar práticas de gerenciamento de estresse (meditação, yoga, respiração diafragmática)"]'::jsonb,
  related_markers = '["PCR ultrassensível", "Vitamina D (25-OH)", "Vitamina B12", "Hemograma completo", "TSH e T4 livre", "Perfil metabólico"]'::jsonb,
  articles_suggestions = '["Entendendo claudicação: causas e abordagem integrativa", "Medicina Funcional: investigando sintomas inexplicados", "Nutrição anti-inflamatória para saúde integral", "O papel do estilo de vida na prevenção e tratamento"]'::jsonb,
  updated_at = NOW()
WHERE id = 'b4cda178-b5a9-4193-b496-afe7ce51ed91';


-- Dores articulares (Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual):)
UPDATE score_items
SET
  clinical_relevance = 'Dores articulares pode indicar desequilíbrios sistêmicos múltiplos na perspectiva da Medicina Funcional Integrativa. Este sintoma frequentemente reflete disfunções em sistemas inter conectados incluindo metabolismo energético, eixo neuroendócrino, sistema digestivo, equilíbrio imunológico ou sobrecarga tóxica. A avaliação funcional busca identificar causas raiz através de investigação abrangente de biomarcadores, história clínica detalhada e análise de sistemas corporais. Sintomas aparentemente isolados podem ser manifestações precoces de desequilíbrios metabólicos, inflamação de baixo grau, deficiências nutricionais ou disfunção mitocondrial que precedem diagnósticos convencionais.',
  interpretation_guide = 'Durante anamnese de dores articulares, investigue: (1) início, duração e padrão temporal; (2) fatores agravantes e atenuantes; (3) sintomas associados em outros sistemas; (4) correlação com alimentação, sono, estresse e ciclo menstrual; (5) impacto na qualidade de vida. Utilize escalas validadas quando disponíveis. Realize exame físico direcionado. Solicite avaliação laboratorial básica incluindo hemograma, perfil metabólico, função tireoidiana, marcadores inflamatórios, vitamina D e vitamina B12. Sinais de alerta incluem progressão rápida, sintomas neurológicos, perda de peso inexplicada ou febre persistente.',
  recommendations = '["Realizar avaliação laboratorial abrangente para identificar possíveis causas subjacentes de dores articulares", "Implementar diário de sintomas por 2-4 semanas correlacionando com fatores de estilo de vida", "Otimizar nutrição com dieta anti-inflamatória rica em alimentos integrais, vegetais coloridos e gorduras saudáveis", "Avaliar e otimizar vitamina D (meta 50-80 ng/mL) e vitaminas do complexo B", "Implementar práticas de gerenciamento de estresse (meditação, yoga, respiração diafragmática)"]'::jsonb,
  related_markers = '["PCR ultrassensível", "Vitamina D (25-OH)", "Vitamina B12", "Hemograma completo", "TSH e T4 livre", "Perfil metabólico"]'::jsonb,
  articles_suggestions = '["Entendendo dores articulares: causas e abordagem integrativa", "Medicina Funcional: investigando sintomas inexplicados", "Nutrição anti-inflamatória para saúde integral", "O papel do estilo de vida na prevenção e tratamento"]'::jsonb,
  updated_at = NOW()
WHERE id = 'deb8e2ee-27a7-4ac0-a743-d6cb400f4b27';


-- Lesão muscular (Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual):)
UPDATE score_items
SET
  clinical_relevance = 'Lesão muscular pode indicar desequilíbrios sistêmicos múltiplos na perspectiva da Medicina Funcional Integrativa. Este sintoma frequentemente reflete disfunções em sistemas inter conectados incluindo metabolismo energético, eixo neuroendócrino, sistema digestivo, equilíbrio imunológico ou sobrecarga tóxica. A avaliação funcional busca identificar causas raiz através de investigação abrangente de biomarcadores, história clínica detalhada e análise de sistemas corporais. Sintomas aparentemente isolados podem ser manifestações precoces de desequilíbrios metabólicos, inflamação de baixo grau, deficiências nutricionais ou disfunção mitocondrial que precedem diagnósticos convencionais.',
  interpretation_guide = 'Durante anamnese de lesão muscular, investigue: (1) início, duração e padrão temporal; (2) fatores agravantes e atenuantes; (3) sintomas associados em outros sistemas; (4) correlação com alimentação, sono, estresse e ciclo menstrual; (5) impacto na qualidade de vida. Utilize escalas validadas quando disponíveis. Realize exame físico direcionado. Solicite avaliação laboratorial básica incluindo hemograma, perfil metabólico, função tireoidiana, marcadores inflamatórios, vitamina D e vitamina B12. Sinais de alerta incluem progressão rápida, sintomas neurológicos, perda de peso inexplicada ou febre persistente.',
  recommendations = '["Realizar avaliação laboratorial abrangente para identificar possíveis causas subjacentes de lesão muscular", "Implementar diário de sintomas por 2-4 semanas correlacionando com fatores de estilo de vida", "Otimizar nutrição com dieta anti-inflamatória rica em alimentos integrais, vegetais coloridos e gorduras saudáveis", "Avaliar e otimizar vitamina D (meta 50-80 ng/mL) e vitaminas do complexo B", "Implementar práticas de gerenciamento de estresse (meditação, yoga, respiração diafragmática)"]'::jsonb,
  related_markers = '["PCR ultrassensível", "Vitamina D (25-OH)", "Vitamina B12", "Hemograma completo", "TSH e T4 livre", "Perfil metabólico"]'::jsonb,
  articles_suggestions = '["Entendendo lesão muscular: causas e abordagem integrativa", "Medicina Funcional: investigando sintomas inexplicados", "Nutrição anti-inflamatória para saúde integral", "O papel do estilo de vida na prevenção e tratamento"]'::jsonb,
  updated_at = NOW()
WHERE id = 'fdb0c742-1db1-48ba-95bc-91e454256d84';


-- Lesão ligamentar/tendínea (Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual):)
UPDATE score_items
SET
  clinical_relevance = 'Lesão ligamentar/tendínea pode indicar desequilíbrios sistêmicos múltiplos na perspectiva da Medicina Funcional Integrativa. Este sintoma frequentemente reflete disfunções em sistemas inter conectados incluindo metabolismo energético, eixo neuroendócrino, sistema digestivo, equilíbrio imunológico ou sobrecarga tóxica. A avaliação funcional busca identificar causas raiz através de investigação abrangente de biomarcadores, história clínica detalhada e análise de sistemas corporais. Sintomas aparentemente isolados podem ser manifestações precoces de desequilíbrios metabólicos, inflamação de baixo grau, deficiências nutricionais ou disfunção mitocondrial que precedem diagnósticos convencionais.',
  interpretation_guide = 'Durante anamnese de lesão ligamentar/tendínea, investigue: (1) início, duração e padrão temporal; (2) fatores agravantes e atenuantes; (3) sintomas associados em outros sistemas; (4) correlação com alimentação, sono, estresse e ciclo menstrual; (5) impacto na qualidade de vida. Utilize escalas validadas quando disponíveis. Realize exame físico direcionado. Solicite avaliação laboratorial básica incluindo hemograma, perfil metabólico, função tireoidiana, marcadores inflamatórios, vitamina D e vitamina B12. Sinais de alerta incluem progressão rápida, sintomas neurológicos, perda de peso inexplicada ou febre persistente.',
  recommendations = '["Realizar avaliação laboratorial abrangente para identificar possíveis causas subjacentes de lesão ligamentar/tendínea", "Implementar diário de sintomas por 2-4 semanas correlacionando com fatores de estilo de vida", "Otimizar nutrição com dieta anti-inflamatória rica em alimentos integrais, vegetais coloridos e gorduras saudáveis", "Avaliar e otimizar vitamina D (meta 50-80 ng/mL) e vitaminas do complexo B", "Implementar práticas de gerenciamento de estresse (meditação, yoga, respiração diafragmática)"]'::jsonb,
  related_markers = '["PCR ultrassensível", "Vitamina D (25-OH)", "Vitamina B12", "Hemograma completo", "TSH e T4 livre", "Perfil metabólico"]'::jsonb,
  articles_suggestions = '["Entendendo lesão ligamentar/tendínea: causas e abordagem integrativa", "Medicina Funcional: investigando sintomas inexplicados", "Nutrição anti-inflamatória para saúde integral", "O papel do estilo de vida na prevenção e tratamento"]'::jsonb,
  updated_at = NOW()
WHERE id = 'e375191b-d8be-4044-9d59-b71fa222cbaa';


-- Fraturas (Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual):)
UPDATE score_items
SET
  clinical_relevance = 'Fraturas pode indicar desequilíbrios sistêmicos múltiplos na perspectiva da Medicina Funcional Integrativa. Este sintoma frequentemente reflete disfunções em sistemas inter conectados incluindo metabolismo energético, eixo neuroendócrino, sistema digestivo, equilíbrio imunológico ou sobrecarga tóxica. A avaliação funcional busca identificar causas raiz através de investigação abrangente de biomarcadores, história clínica detalhada e análise de sistemas corporais. Sintomas aparentemente isolados podem ser manifestações precoces de desequilíbrios metabólicos, inflamação de baixo grau, deficiências nutricionais ou disfunção mitocondrial que precedem diagnósticos convencionais.',
  interpretation_guide = 'Durante anamnese de fraturas, investigue: (1) início, duração e padrão temporal; (2) fatores agravantes e atenuantes; (3) sintomas associados em outros sistemas; (4) correlação com alimentação, sono, estresse e ciclo menstrual; (5) impacto na qualidade de vida. Utilize escalas validadas quando disponíveis. Realize exame físico direcionado. Solicite avaliação laboratorial básica incluindo hemograma, perfil metabólico, função tireoidiana, marcadores inflamatórios, vitamina D e vitamina B12. Sinais de alerta incluem progressão rápida, sintomas neurológicos, perda de peso inexplicada ou febre persistente.',
  recommendations = '["Realizar avaliação laboratorial abrangente para identificar possíveis causas subjacentes de fraturas", "Implementar diário de sintomas por 2-4 semanas correlacionando com fatores de estilo de vida", "Otimizar nutrição com dieta anti-inflamatória rica em alimentos integrais, vegetais coloridos e gorduras saudáveis", "Avaliar e otimizar vitamina D (meta 50-80 ng/mL) e vitaminas do complexo B", "Implementar práticas de gerenciamento de estresse (meditação, yoga, respiração diafragmática)"]'::jsonb,
  related_markers = '["PCR ultrassensível", "Vitamina D (25-OH)", "Vitamina B12", "Hemograma completo", "TSH e T4 livre", "Perfil metabólico"]'::jsonb,
  articles_suggestions = '["Entendendo fraturas: causas e abordagem integrativa", "Medicina Funcional: investigando sintomas inexplicados", "Nutrição anti-inflamatória para saúde integral", "O papel do estilo de vida na prevenção e tratamento"]'::jsonb,
  updated_at = NOW()
WHERE id = '44f501b3-6c1d-4a31-b8c9-0f54b201dbcb';


-- Edema (Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual):)
UPDATE score_items
SET
  clinical_relevance = 'Edema pode indicar desequilíbrios sistêmicos múltiplos na perspectiva da Medicina Funcional Integrativa. Este sintoma frequentemente reflete disfunções em sistemas inter conectados incluindo metabolismo energético, eixo neuroendócrino, sistema digestivo, equilíbrio imunológico ou sobrecarga tóxica. A avaliação funcional busca identificar causas raiz através de investigação abrangente de biomarcadores, história clínica detalhada e análise de sistemas corporais. Sintomas aparentemente isolados podem ser manifestações precoces de desequilíbrios metabólicos, inflamação de baixo grau, deficiências nutricionais ou disfunção mitocondrial que precedem diagnósticos convencionais.',
  interpretation_guide = 'Durante anamnese de edema, investigue: (1) início, duração e padrão temporal; (2) fatores agravantes e atenuantes; (3) sintomas associados em outros sistemas; (4) correlação com alimentação, sono, estresse e ciclo menstrual; (5) impacto na qualidade de vida. Utilize escalas validadas quando disponíveis. Realize exame físico direcionado. Solicite avaliação laboratorial básica incluindo hemograma, perfil metabólico, função tireoidiana, marcadores inflamatórios, vitamina D e vitamina B12. Sinais de alerta incluem progressão rápida, sintomas neurológicos, perda de peso inexplicada ou febre persistente.',
  recommendations = '["Realizar avaliação laboratorial abrangente para identificar possíveis causas subjacentes de edema", "Implementar diário de sintomas por 2-4 semanas correlacionando com fatores de estilo de vida", "Otimizar nutrição com dieta anti-inflamatória rica em alimentos integrais, vegetais coloridos e gorduras saudáveis", "Avaliar e otimizar vitamina D (meta 50-80 ng/mL) e vitaminas do complexo B", "Implementar práticas de gerenciamento de estresse (meditação, yoga, respiração diafragmática)"]'::jsonb,
  related_markers = '["PCR ultrassensível", "Vitamina D (25-OH)", "Vitamina B12", "Hemograma completo", "TSH e T4 livre", "Perfil metabólico"]'::jsonb,
  articles_suggestions = '["Entendendo edema: causas e abordagem integrativa", "Medicina Funcional: investigando sintomas inexplicados", "Nutrição anti-inflamatória para saúde integral", "O papel do estilo de vida na prevenção e tratamento"]'::jsonb,
  updated_at = NOW()
WHERE id = '14881a00-579e-4297-bd6d-4c83d8081d2d';


-- Pele e tegumento (Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual):)
UPDATE score_items
SET
  clinical_relevance = 'Pele e tegumento pode indicar desequilíbrios sistêmicos múltiplos na perspectiva da Medicina Funcional Integrativa. Este sintoma frequentemente reflete disfunções em sistemas inter conectados incluindo metabolismo energético, eixo neuroendócrino, sistema digestivo, equilíbrio imunológico ou sobrecarga tóxica. A avaliação funcional busca identificar causas raiz através de investigação abrangente de biomarcadores, história clínica detalhada e análise de sistemas corporais. Sintomas aparentemente isolados podem ser manifestações precoces de desequilíbrios metabólicos, inflamação de baixo grau, deficiências nutricionais ou disfunção mitocondrial que precedem diagnósticos convencionais.',
  interpretation_guide = 'Durante anamnese de pele e tegumento, investigue: (1) início, duração e padrão temporal; (2) fatores agravantes e atenuantes; (3) sintomas associados em outros sistemas; (4) correlação com alimentação, sono, estresse e ciclo menstrual; (5) impacto na qualidade de vida. Utilize escalas validadas quando disponíveis. Realize exame físico direcionado. Solicite avaliação laboratorial básica incluindo hemograma, perfil metabólico, função tireoidiana, marcadores inflamatórios, vitamina D e vitamina B12. Sinais de alerta incluem progressão rápida, sintomas neurológicos, perda de peso inexplicada ou febre persistente.',
  recommendations = '["Realizar avaliação laboratorial abrangente para identificar possíveis causas subjacentes de pele e tegumento", "Implementar diário de sintomas por 2-4 semanas correlacionando com fatores de estilo de vida", "Otimizar nutrição com dieta anti-inflamatória rica em alimentos integrais, vegetais coloridos e gorduras saudáveis", "Avaliar e otimizar vitamina D (meta 50-80 ng/mL) e vitaminas do complexo B", "Implementar práticas de gerenciamento de estresse (meditação, yoga, respiração diafragmática)"]'::jsonb,
  related_markers = '["PCR ultrassensível", "Vitamina D (25-OH)", "Vitamina B12", "Hemograma completo", "TSH e T4 livre", "Perfil metabólico"]'::jsonb,
  articles_suggestions = '["Entendendo pele e tegumento: causas e abordagem integrativa", "Medicina Funcional: investigando sintomas inexplicados", "Nutrição anti-inflamatória para saúde integral", "O papel do estilo de vida na prevenção e tratamento"]'::jsonb,
  updated_at = NOW()
WHERE id = 'a7698447-ddbf-4bee-92b4-81a1559562ad';


-- Enfraquecimento capilar (Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual):)
UPDATE score_items
SET
  clinical_relevance = 'Enfraquecimento capilar pode indicar desequilíbrios sistêmicos múltiplos na perspectiva da Medicina Funcional Integrativa. Este sintoma frequentemente reflete disfunções em sistemas inter conectados incluindo metabolismo energético, eixo neuroendócrino, sistema digestivo, equilíbrio imunológico ou sobrecarga tóxica. A avaliação funcional busca identificar causas raiz através de investigação abrangente de biomarcadores, história clínica detalhada e análise de sistemas corporais. Sintomas aparentemente isolados podem ser manifestações precoces de desequilíbrios metabólicos, inflamação de baixo grau, deficiências nutricionais ou disfunção mitocondrial que precedem diagnósticos convencionais.',
  interpretation_guide = 'Durante anamnese de enfraquecimento capilar, investigue: (1) início, duração e padrão temporal; (2) fatores agravantes e atenuantes; (3) sintomas associados em outros sistemas; (4) correlação com alimentação, sono, estresse e ciclo menstrual; (5) impacto na qualidade de vida. Utilize escalas validadas quando disponíveis. Realize exame físico direcionado. Solicite avaliação laboratorial básica incluindo hemograma, perfil metabólico, função tireoidiana, marcadores inflamatórios, vitamina D e vitamina B12. Sinais de alerta incluem progressão rápida, sintomas neurológicos, perda de peso inexplicada ou febre persistente.',
  recommendations = '["Realizar avaliação laboratorial abrangente para identificar possíveis causas subjacentes de enfraquecimento capilar", "Implementar diário de sintomas por 2-4 semanas correlacionando com fatores de estilo de vida", "Otimizar nutrição com dieta anti-inflamatória rica em alimentos integrais, vegetais coloridos e gorduras saudáveis", "Avaliar e otimizar vitamina D (meta 50-80 ng/mL) e vitaminas do complexo B", "Implementar práticas de gerenciamento de estresse (meditação, yoga, respiração diafragmática)"]'::jsonb,
  related_markers = '["PCR ultrassensível", "Vitamina D (25-OH)", "Vitamina B12", "Hemograma completo", "TSH e T4 livre", "Perfil metabólico"]'::jsonb,
  articles_suggestions = '["Entendendo enfraquecimento capilar: causas e abordagem integrativa", "Medicina Funcional: investigando sintomas inexplicados", "Nutrição anti-inflamatória para saúde integral", "O papel do estilo de vida na prevenção e tratamento"]'::jsonb,
  updated_at = NOW()
WHERE id = '76578167-9f69-4981-b73d-46aa7c71ec6b';


-- Queda capilar (Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual):)
UPDATE score_items
SET
  clinical_relevance = 'Queda capilar pode indicar desequilíbrios sistêmicos múltiplos na perspectiva da Medicina Funcional Integrativa. Este sintoma frequentemente reflete disfunções em sistemas inter conectados incluindo metabolismo energético, eixo neuroendócrino, sistema digestivo, equilíbrio imunológico ou sobrecarga tóxica. A avaliação funcional busca identificar causas raiz através de investigação abrangente de biomarcadores, história clínica detalhada e análise de sistemas corporais. Sintomas aparentemente isolados podem ser manifestações precoces de desequilíbrios metabólicos, inflamação de baixo grau, deficiências nutricionais ou disfunção mitocondrial que precedem diagnósticos convencionais.',
  interpretation_guide = 'Durante anamnese de queda capilar, investigue: (1) início, duração e padrão temporal; (2) fatores agravantes e atenuantes; (3) sintomas associados em outros sistemas; (4) correlação com alimentação, sono, estresse e ciclo menstrual; (5) impacto na qualidade de vida. Utilize escalas validadas quando disponíveis. Realize exame físico direcionado. Solicite avaliação laboratorial básica incluindo hemograma, perfil metabólico, função tireoidiana, marcadores inflamatórios, vitamina D e vitamina B12. Sinais de alerta incluem progressão rápida, sintomas neurológicos, perda de peso inexplicada ou febre persistente.',
  recommendations = '["Realizar avaliação laboratorial abrangente para identificar possíveis causas subjacentes de queda capilar", "Implementar diário de sintomas por 2-4 semanas correlacionando com fatores de estilo de vida", "Otimizar nutrição com dieta anti-inflamatória rica em alimentos integrais, vegetais coloridos e gorduras saudáveis", "Avaliar e otimizar vitamina D (meta 50-80 ng/mL) e vitaminas do complexo B", "Implementar práticas de gerenciamento de estresse (meditação, yoga, respiração diafragmática)"]'::jsonb,
  related_markers = '["PCR ultrassensível", "Vitamina D (25-OH)", "Vitamina B12", "Hemograma completo", "TSH e T4 livre", "Perfil metabólico"]'::jsonb,
  articles_suggestions = '["Entendendo queda capilar: causas e abordagem integrativa", "Medicina Funcional: investigando sintomas inexplicados", "Nutrição anti-inflamatória para saúde integral", "O papel do estilo de vida na prevenção e tratamento"]'::jsonb,
  updated_at = NOW()
WHERE id = '23b90c6c-6f44-481a-8255-108158a64239';


-- Enfraquecimento ungueal (Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual):)
UPDATE score_items
SET
  clinical_relevance = 'Enfraquecimento ungueal pode indicar desequilíbrios sistêmicos múltiplos na perspectiva da Medicina Funcional Integrativa. Este sintoma frequentemente reflete disfunções em sistemas inter conectados incluindo metabolismo energético, eixo neuroendócrino, sistema digestivo, equilíbrio imunológico ou sobrecarga tóxica. A avaliação funcional busca identificar causas raiz através de investigação abrangente de biomarcadores, história clínica detalhada e análise de sistemas corporais. Sintomas aparentemente isolados podem ser manifestações precoces de desequilíbrios metabólicos, inflamação de baixo grau, deficiências nutricionais ou disfunção mitocondrial que precedem diagnósticos convencionais.',
  interpretation_guide = 'Durante anamnese de enfraquecimento ungueal, investigue: (1) início, duração e padrão temporal; (2) fatores agravantes e atenuantes; (3) sintomas associados em outros sistemas; (4) correlação com alimentação, sono, estresse e ciclo menstrual; (5) impacto na qualidade de vida. Utilize escalas validadas quando disponíveis. Realize exame físico direcionado. Solicite avaliação laboratorial básica incluindo hemograma, perfil metabólico, função tireoidiana, marcadores inflamatórios, vitamina D e vitamina B12. Sinais de alerta incluem progressão rápida, sintomas neurológicos, perda de peso inexplicada ou febre persistente.',
  recommendations = '["Realizar avaliação laboratorial abrangente para identificar possíveis causas subjacentes de enfraquecimento ungueal", "Implementar diário de sintomas por 2-4 semanas correlacionando com fatores de estilo de vida", "Otimizar nutrição com dieta anti-inflamatória rica em alimentos integrais, vegetais coloridos e gorduras saudáveis", "Avaliar e otimizar vitamina D (meta 50-80 ng/mL) e vitaminas do complexo B", "Implementar práticas de gerenciamento de estresse (meditação, yoga, respiração diafragmática)"]'::jsonb,
  related_markers = '["PCR ultrassensível", "Vitamina D (25-OH)", "Vitamina B12", "Hemograma completo", "TSH e T4 livre", "Perfil metabólico"]'::jsonb,
  articles_suggestions = '["Entendendo enfraquecimento ungueal: causas e abordagem integrativa", "Medicina Funcional: investigando sintomas inexplicados", "Nutrição anti-inflamatória para saúde integral", "O papel do estilo de vida na prevenção e tratamento"]'::jsonb,
  updated_at = NOW()
WHERE id = '4d1372ee-5bee-4ce8-845a-a4e176109ad2';


-- Genitália masculina (Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual):)
UPDATE score_items
SET
  clinical_relevance = 'Genitália masculina pode indicar desequilíbrios sistêmicos múltiplos na perspectiva da Medicina Funcional Integrativa. Este sintoma frequentemente reflete disfunções em sistemas inter conectados incluindo metabolismo energético, eixo neuroendócrino, sistema digestivo, equilíbrio imunológico ou sobrecarga tóxica. A avaliação funcional busca identificar causas raiz através de investigação abrangente de biomarcadores, história clínica detalhada e análise de sistemas corporais. Sintomas aparentemente isolados podem ser manifestações precoces de desequilíbrios metabólicos, inflamação de baixo grau, deficiências nutricionais ou disfunção mitocondrial que precedem diagnósticos convencionais.',
  interpretation_guide = 'Durante anamnese de genitália masculina, investigue: (1) início, duração e padrão temporal; (2) fatores agravantes e atenuantes; (3) sintomas associados em outros sistemas; (4) correlação com alimentação, sono, estresse e ciclo menstrual; (5) impacto na qualidade de vida. Utilize escalas validadas quando disponíveis. Realize exame físico direcionado. Solicite avaliação laboratorial básica incluindo hemograma, perfil metabólico, função tireoidiana, marcadores inflamatórios, vitamina D e vitamina B12. Sinais de alerta incluem progressão rápida, sintomas neurológicos, perda de peso inexplicada ou febre persistente.',
  recommendations = '["Realizar avaliação laboratorial abrangente para identificar possíveis causas subjacentes de genitália masculina", "Implementar diário de sintomas por 2-4 semanas correlacionando com fatores de estilo de vida", "Otimizar nutrição com dieta anti-inflamatória rica em alimentos integrais, vegetais coloridos e gorduras saudáveis", "Avaliar e otimizar vitamina D (meta 50-80 ng/mL) e vitaminas do complexo B", "Implementar práticas de gerenciamento de estresse (meditação, yoga, respiração diafragmática)"]'::jsonb,
  related_markers = '["PCR ultrassensível", "Vitamina D (25-OH)", "Vitamina B12", "Hemograma completo", "TSH e T4 livre", "Perfil metabólico"]'::jsonb,
  articles_suggestions = '["Entendendo genitália masculina: causas e abordagem integrativa", "Medicina Funcional: investigando sintomas inexplicados", "Nutrição anti-inflamatória para saúde integral", "O papel do estilo de vida na prevenção e tratamento"]'::jsonb,
  updated_at = NOW()
WHERE id = 'fe938ed9-3cf9-4ddf-a1ae-2f67c03e54a4';


-- Prepúcio / glande (Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual):)
UPDATE score_items
SET
  clinical_relevance = 'Prepúcio / glande pode indicar desequilíbrios sistêmicos múltiplos na perspectiva da Medicina Funcional Integrativa. Este sintoma frequentemente reflete disfunções em sistemas inter conectados incluindo metabolismo energético, eixo neuroendócrino, sistema digestivo, equilíbrio imunológico ou sobrecarga tóxica. A avaliação funcional busca identificar causas raiz através de investigação abrangente de biomarcadores, história clínica detalhada e análise de sistemas corporais. Sintomas aparentemente isolados podem ser manifestações precoces de desequilíbrios metabólicos, inflamação de baixo grau, deficiências nutricionais ou disfunção mitocondrial que precedem diagnósticos convencionais.',
  interpretation_guide = 'Durante anamnese de prepúcio / glande, investigue: (1) início, duração e padrão temporal; (2) fatores agravantes e atenuantes; (3) sintomas associados em outros sistemas; (4) correlação com alimentação, sono, estresse e ciclo menstrual; (5) impacto na qualidade de vida. Utilize escalas validadas quando disponíveis. Realize exame físico direcionado. Solicite avaliação laboratorial básica incluindo hemograma, perfil metabólico, função tireoidiana, marcadores inflamatórios, vitamina D e vitamina B12. Sinais de alerta incluem progressão rápida, sintomas neurológicos, perda de peso inexplicada ou febre persistente.',
  recommendations = '["Realizar avaliação laboratorial abrangente para identificar possíveis causas subjacentes de prepúcio / glande", "Implementar diário de sintomas por 2-4 semanas correlacionando com fatores de estilo de vida", "Otimizar nutrição com dieta anti-inflamatória rica em alimentos integrais, vegetais coloridos e gorduras saudáveis", "Avaliar e otimizar vitamina D (meta 50-80 ng/mL) e vitaminas do complexo B", "Implementar práticas de gerenciamento de estresse (meditação, yoga, respiração diafragmática)"]'::jsonb,
  related_markers = '["PCR ultrassensível", "Vitamina D (25-OH)", "Vitamina B12", "Hemograma completo", "TSH e T4 livre", "Perfil metabólico"]'::jsonb,
  articles_suggestions = '["Entendendo prepúcio / glande: causas e abordagem integrativa", "Medicina Funcional: investigando sintomas inexplicados", "Nutrição anti-inflamatória para saúde integral", "O papel do estilo de vida na prevenção e tratamento"]'::jsonb,
  updated_at = NOW()
WHERE id = '5ddc4af9-dd19-4ea1-8117-3d3e30377dab';


-- Escroto / epidídimos (Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual):)
UPDATE score_items
SET
  clinical_relevance = 'Escroto / epidídimos pode indicar desequilíbrios sistêmicos múltiplos na perspectiva da Medicina Funcional Integrativa. Este sintoma frequentemente reflete disfunções em sistemas inter conectados incluindo metabolismo energético, eixo neuroendócrino, sistema digestivo, equilíbrio imunológico ou sobrecarga tóxica. A avaliação funcional busca identificar causas raiz através de investigação abrangente de biomarcadores, história clínica detalhada e análise de sistemas corporais. Sintomas aparentemente isolados podem ser manifestações precoces de desequilíbrios metabólicos, inflamação de baixo grau, deficiências nutricionais ou disfunção mitocondrial que precedem diagnósticos convencionais.',
  interpretation_guide = 'Durante anamnese de escroto / epidídimos, investigue: (1) início, duração e padrão temporal; (2) fatores agravantes e atenuantes; (3) sintomas associados em outros sistemas; (4) correlação com alimentação, sono, estresse e ciclo menstrual; (5) impacto na qualidade de vida. Utilize escalas validadas quando disponíveis. Realize exame físico direcionado. Solicite avaliação laboratorial básica incluindo hemograma, perfil metabólico, função tireoidiana, marcadores inflamatórios, vitamina D e vitamina B12. Sinais de alerta incluem progressão rápida, sintomas neurológicos, perda de peso inexplicada ou febre persistente.',
  recommendations = '["Realizar avaliação laboratorial abrangente para identificar possíveis causas subjacentes de escroto / epidídimos", "Implementar diário de sintomas por 2-4 semanas correlacionando com fatores de estilo de vida", "Otimizar nutrição com dieta anti-inflamatória rica em alimentos integrais, vegetais coloridos e gorduras saudáveis", "Avaliar e otimizar vitamina D (meta 50-80 ng/mL) e vitaminas do complexo B", "Implementar práticas de gerenciamento de estresse (meditação, yoga, respiração diafragmática)"]'::jsonb,
  related_markers = '["PCR ultrassensível", "Vitamina D (25-OH)", "Vitamina B12", "Hemograma completo", "TSH e T4 livre", "Perfil metabólico"]'::jsonb,
  articles_suggestions = '["Entendendo escroto / epidídimos: causas e abordagem integrativa", "Medicina Funcional: investigando sintomas inexplicados", "Nutrição anti-inflamatória para saúde integral", "O papel do estilo de vida na prevenção e tratamento"]'::jsonb,
  updated_at = NOW()
WHERE id = 'e010b2a7-6e9e-4b42-9d37-487e91411a18';


-- Testículos (Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual):)
UPDATE score_items
SET
  clinical_relevance = 'Testículos pode indicar desequilíbrios sistêmicos múltiplos na perspectiva da Medicina Funcional Integrativa. Este sintoma frequentemente reflete disfunções em sistemas inter conectados incluindo metabolismo energético, eixo neuroendócrino, sistema digestivo, equilíbrio imunológico ou sobrecarga tóxica. A avaliação funcional busca identificar causas raiz através de investigação abrangente de biomarcadores, história clínica detalhada e análise de sistemas corporais. Sintomas aparentemente isolados podem ser manifestações precoces de desequilíbrios metabólicos, inflamação de baixo grau, deficiências nutricionais ou disfunção mitocondrial que precedem diagnósticos convencionais.',
  interpretation_guide = 'Durante anamnese de testículos, investigue: (1) início, duração e padrão temporal; (2) fatores agravantes e atenuantes; (3) sintomas associados em outros sistemas; (4) correlação com alimentação, sono, estresse e ciclo menstrual; (5) impacto na qualidade de vida. Utilize escalas validadas quando disponíveis. Realize exame físico direcionado. Solicite avaliação laboratorial básica incluindo hemograma, perfil metabólico, função tireoidiana, marcadores inflamatórios, vitamina D e vitamina B12. Sinais de alerta incluem progressão rápida, sintomas neurológicos, perda de peso inexplicada ou febre persistente.',
  recommendations = '["Realizar avaliação laboratorial abrangente para identificar possíveis causas subjacentes de testículos", "Implementar diário de sintomas por 2-4 semanas correlacionando com fatores de estilo de vida", "Otimizar nutrição com dieta anti-inflamatória rica em alimentos integrais, vegetais coloridos e gorduras saudáveis", "Avaliar e otimizar vitamina D (meta 50-80 ng/mL) e vitaminas do complexo B", "Implementar práticas de gerenciamento de estresse (meditação, yoga, respiração diafragmática)"]'::jsonb,
  related_markers = '["PCR ultrassensível", "Vitamina D (25-OH)", "Vitamina B12", "Hemograma completo", "TSH e T4 livre", "Perfil metabólico"]'::jsonb,
  articles_suggestions = '["Entendendo testículos: causas e abordagem integrativa", "Medicina Funcional: investigando sintomas inexplicados", "Nutrição anti-inflamatória para saúde integral", "O papel do estilo de vida na prevenção e tratamento"]'::jsonb,
  updated_at = NOW()
WHERE id = '933b7816-fe10-4e35-94d7-0d232cc258ce';


-- Genitália feminina (Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual):)
UPDATE score_items
SET
  clinical_relevance = 'Genitália feminina pode indicar desequilíbrios sistêmicos múltiplos na perspectiva da Medicina Funcional Integrativa. Este sintoma frequentemente reflete disfunções em sistemas inter conectados incluindo metabolismo energético, eixo neuroendócrino, sistema digestivo, equilíbrio imunológico ou sobrecarga tóxica. A avaliação funcional busca identificar causas raiz através de investigação abrangente de biomarcadores, história clínica detalhada e análise de sistemas corporais. Sintomas aparentemente isolados podem ser manifestações precoces de desequilíbrios metabólicos, inflamação de baixo grau, deficiências nutricionais ou disfunção mitocondrial que precedem diagnósticos convencionais.',
  interpretation_guide = 'Durante anamnese de genitália feminina, investigue: (1) início, duração e padrão temporal; (2) fatores agravantes e atenuantes; (3) sintomas associados em outros sistemas; (4) correlação com alimentação, sono, estresse e ciclo menstrual; (5) impacto na qualidade de vida. Utilize escalas validadas quando disponíveis. Realize exame físico direcionado. Solicite avaliação laboratorial básica incluindo hemograma, perfil metabólico, função tireoidiana, marcadores inflamatórios, vitamina D e vitamina B12. Sinais de alerta incluem progressão rápida, sintomas neurológicos, perda de peso inexplicada ou febre persistente.',
  recommendations = '["Realizar avaliação laboratorial abrangente para identificar possíveis causas subjacentes de genitália feminina", "Implementar diário de sintomas por 2-4 semanas correlacionando com fatores de estilo de vida", "Otimizar nutrição com dieta anti-inflamatória rica em alimentos integrais, vegetais coloridos e gorduras saudáveis", "Avaliar e otimizar vitamina D (meta 50-80 ng/mL) e vitaminas do complexo B", "Implementar práticas de gerenciamento de estresse (meditação, yoga, respiração diafragmática)"]'::jsonb,
  related_markers = '["PCR ultrassensível", "Vitamina D (25-OH)", "Vitamina B12", "Hemograma completo", "TSH e T4 livre", "Perfil metabólico"]'::jsonb,
  articles_suggestions = '["Entendendo genitália feminina: causas e abordagem integrativa", "Medicina Funcional: investigando sintomas inexplicados", "Nutrição anti-inflamatória para saúde integral", "O papel do estilo de vida na prevenção e tratamento"]'::jsonb,
  updated_at = NOW()
WHERE id = '2b01e2f0-76c7-4616-8f7b-33b1c9d11279';


-- Trofismo Urogenital (Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual):)
UPDATE score_items
SET
  clinical_relevance = 'Trofismo Urogenital pode indicar desequilíbrios sistêmicos múltiplos na perspectiva da Medicina Funcional Integrativa. Este sintoma frequentemente reflete disfunções em sistemas inter conectados incluindo metabolismo energético, eixo neuroendócrino, sistema digestivo, equilíbrio imunológico ou sobrecarga tóxica. A avaliação funcional busca identificar causas raiz através de investigação abrangente de biomarcadores, história clínica detalhada e análise de sistemas corporais. Sintomas aparentemente isolados podem ser manifestações precoces de desequilíbrios metabólicos, inflamação de baixo grau, deficiências nutricionais ou disfunção mitocondrial que precedem diagnósticos convencionais.',
  interpretation_guide = 'Durante anamnese de trofismo urogenital, investigue: (1) início, duração e padrão temporal; (2) fatores agravantes e atenuantes; (3) sintomas associados em outros sistemas; (4) correlação com alimentação, sono, estresse e ciclo menstrual; (5) impacto na qualidade de vida. Utilize escalas validadas quando disponíveis. Realize exame físico direcionado. Solicite avaliação laboratorial básica incluindo hemograma, perfil metabólico, função tireoidiana, marcadores inflamatórios, vitamina D e vitamina B12. Sinais de alerta incluem progressão rápida, sintomas neurológicos, perda de peso inexplicada ou febre persistente.',
  recommendations = '["Realizar avaliação laboratorial abrangente para identificar possíveis causas subjacentes de trofismo urogenital", "Implementar diário de sintomas por 2-4 semanas correlacionando com fatores de estilo de vida", "Otimizar nutrição com dieta anti-inflamatória rica em alimentos integrais, vegetais coloridos e gorduras saudáveis", "Avaliar e otimizar vitamina D (meta 50-80 ng/mL) e vitaminas do complexo B", "Implementar práticas de gerenciamento de estresse (meditação, yoga, respiração diafragmática)"]'::jsonb,
  related_markers = '["PCR ultrassensível", "Vitamina D (25-OH)", "Vitamina B12", "Hemograma completo", "TSH e T4 livre", "Perfil metabólico"]'::jsonb,
  articles_suggestions = '["Entendendo trofismo urogenital: causas e abordagem integrativa", "Medicina Funcional: investigando sintomas inexplicados", "Nutrição anti-inflamatória para saúde integral", "O papel do estilo de vida na prevenção e tratamento"]'::jsonb,
  updated_at = NOW()
WHERE id = '33ffce34-e1fb-4e31-b393-f2783549d874';


-- Suporte Pélvico (Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual):)
UPDATE score_items
SET
  clinical_relevance = 'Suporte Pélvico pode indicar desequilíbrios sistêmicos múltiplos na perspectiva da Medicina Funcional Integrativa. Este sintoma frequentemente reflete disfunções em sistemas inter conectados incluindo metabolismo energético, eixo neuroendócrino, sistema digestivo, equilíbrio imunológico ou sobrecarga tóxica. A avaliação funcional busca identificar causas raiz através de investigação abrangente de biomarcadores, história clínica detalhada e análise de sistemas corporais. Sintomas aparentemente isolados podem ser manifestações precoces de desequilíbrios metabólicos, inflamação de baixo grau, deficiências nutricionais ou disfunção mitocondrial que precedem diagnósticos convencionais.',
  interpretation_guide = 'Durante anamnese de suporte pélvico, investigue: (1) início, duração e padrão temporal; (2) fatores agravantes e atenuantes; (3) sintomas associados em outros sistemas; (4) correlação com alimentação, sono, estresse e ciclo menstrual; (5) impacto na qualidade de vida. Utilize escalas validadas quando disponíveis. Realize exame físico direcionado. Solicite avaliação laboratorial básica incluindo hemograma, perfil metabólico, função tireoidiana, marcadores inflamatórios, vitamina D e vitamina B12. Sinais de alerta incluem progressão rápida, sintomas neurológicos, perda de peso inexplicada ou febre persistente.',
  recommendations = '["Realizar avaliação laboratorial abrangente para identificar possíveis causas subjacentes de suporte pélvico", "Implementar diário de sintomas por 2-4 semanas correlacionando com fatores de estilo de vida", "Otimizar nutrição com dieta anti-inflamatória rica em alimentos integrais, vegetais coloridos e gorduras saudáveis", "Avaliar e otimizar vitamina D (meta 50-80 ng/mL) e vitaminas do complexo B", "Implementar práticas de gerenciamento de estresse (meditação, yoga, respiração diafragmática)"]'::jsonb,
  related_markers = '["PCR ultrassensível", "Vitamina D (25-OH)", "Vitamina B12", "Hemograma completo", "TSH e T4 livre", "Perfil metabólico"]'::jsonb,
  articles_suggestions = '["Entendendo suporte pélvico: causas e abordagem integrativa", "Medicina Funcional: investigando sintomas inexplicados", "Nutrição anti-inflamatória para saúde integral", "O papel do estilo de vida na prevenção e tratamento"]'::jsonb,
  updated_at = NOW()
WHERE id = 'f54dbfaa-599a-40ee-907a-28431ce4858a';


-- Vulva e Estruturas Externas (Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual):)
UPDATE score_items
SET
  clinical_relevance = 'Vulva e Estruturas Externas pode indicar desequilíbrios sistêmicos múltiplos na perspectiva da Medicina Funcional Integrativa. Este sintoma frequentemente reflete disfunções em sistemas inter conectados incluindo metabolismo energético, eixo neuroendócrino, sistema digestivo, equilíbrio imunológico ou sobrecarga tóxica. A avaliação funcional busca identificar causas raiz através de investigação abrangente de biomarcadores, história clínica detalhada e análise de sistemas corporais. Sintomas aparentemente isolados podem ser manifestações precoces de desequilíbrios metabólicos, inflamação de baixo grau, deficiências nutricionais ou disfunção mitocondrial que precedem diagnósticos convencionais.',
  interpretation_guide = 'Durante anamnese de vulva e estruturas externas, investigue: (1) início, duração e padrão temporal; (2) fatores agravantes e atenuantes; (3) sintomas associados em outros sistemas; (4) correlação com alimentação, sono, estresse e ciclo menstrual; (5) impacto na qualidade de vida. Utilize escalas validadas quando disponíveis. Realize exame físico direcionado. Solicite avaliação laboratorial básica incluindo hemograma, perfil metabólico, função tireoidiana, marcadores inflamatórios, vitamina D e vitamina B12. Sinais de alerta incluem progressão rápida, sintomas neurológicos, perda de peso inexplicada ou febre persistente.',
  recommendations = '["Realizar avaliação laboratorial abrangente para identificar possíveis causas subjacentes de vulva e estruturas externas", "Implementar diário de sintomas por 2-4 semanas correlacionando com fatores de estilo de vida", "Otimizar nutrição com dieta anti-inflamatória rica em alimentos integrais, vegetais coloridos e gorduras saudáveis", "Avaliar e otimizar vitamina D (meta 50-80 ng/mL) e vitaminas do complexo B", "Implementar práticas de gerenciamento de estresse (meditação, yoga, respiração diafragmática)"]'::jsonb,
  related_markers = '["PCR ultrassensível", "Vitamina D (25-OH)", "Vitamina B12", "Hemograma completo", "TSH e T4 livre", "Perfil metabólico"]'::jsonb,
  articles_suggestions = '["Entendendo vulva e estruturas externas: causas e abordagem integrativa", "Medicina Funcional: investigando sintomas inexplicados", "Nutrição anti-inflamatória para saúde integral", "O papel do estilo de vida na prevenção e tratamento"]'::jsonb,
  updated_at = NOW()
WHERE id = '30a7764c-4c11-4078-b942-860ee56136a4';


-- Vagina e Colo Uterino (Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual):)
UPDATE score_items
SET
  clinical_relevance = 'Vagina e Colo Uterino pode indicar desequilíbrios sistêmicos múltiplos na perspectiva da Medicina Funcional Integrativa. Este sintoma frequentemente reflete disfunções em sistemas inter conectados incluindo metabolismo energético, eixo neuroendócrino, sistema digestivo, equilíbrio imunológico ou sobrecarga tóxica. A avaliação funcional busca identificar causas raiz através de investigação abrangente de biomarcadores, história clínica detalhada e análise de sistemas corporais. Sintomas aparentemente isolados podem ser manifestações precoces de desequilíbrios metabólicos, inflamação de baixo grau, deficiências nutricionais ou disfunção mitocondrial que precedem diagnósticos convencionais.',
  interpretation_guide = 'Durante anamnese de vagina e colo uterino, investigue: (1) início, duração e padrão temporal; (2) fatores agravantes e atenuantes; (3) sintomas associados em outros sistemas; (4) correlação com alimentação, sono, estresse e ciclo menstrual; (5) impacto na qualidade de vida. Utilize escalas validadas quando disponíveis. Realize exame físico direcionado. Solicite avaliação laboratorial básica incluindo hemograma, perfil metabólico, função tireoidiana, marcadores inflamatórios, vitamina D e vitamina B12. Sinais de alerta incluem progressão rápida, sintomas neurológicos, perda de peso inexplicada ou febre persistente.',
  recommendations = '["Realizar avaliação laboratorial abrangente para identificar possíveis causas subjacentes de vagina e colo uterino", "Implementar diário de sintomas por 2-4 semanas correlacionando com fatores de estilo de vida", "Otimizar nutrição com dieta anti-inflamatória rica em alimentos integrais, vegetais coloridos e gorduras saudáveis", "Avaliar e otimizar vitamina D (meta 50-80 ng/mL) e vitaminas do complexo B", "Implementar práticas de gerenciamento de estresse (meditação, yoga, respiração diafragmática)"]'::jsonb,
  related_markers = '["PCR ultrassensível", "Vitamina D (25-OH)", "Vitamina B12", "Hemograma completo", "TSH e T4 livre", "Perfil metabólico"]'::jsonb,
  articles_suggestions = '["Entendendo vagina e colo uterino: causas e abordagem integrativa", "Medicina Funcional: investigando sintomas inexplicados", "Nutrição anti-inflamatória para saúde integral", "O papel do estilo de vida na prevenção e tratamento"]'::jsonb,
  updated_at = NOW()
WHERE id = '1ff1e7f8-5b15-4f0c-b76b-806c1d6ff1fd';


-- Registrar quaisquer cirurgias realizadas (Cirurgias já realizadas)
UPDATE score_items
SET
  clinical_relevance = 'O histórico de registrar quaisquer cirurgias realizadas tem implicações importantes para avaliação funcional e estratégias terapêuticas. Cirurgias que removem órgãos ou alteram anatomia podem impactar função metabólica, produção hormonal, absorção nutricional, microbioma e sistemas de detoxificação. A Medicina Funcional Integrativa reconhece que mesmo após recuperação cirúrgica bem-sucedida, podem persistir desequilíbrios fisiológicos que requerem suporte específico. A remoção de tecidos endócrinos (tireoide, ovários, próstata) necessita reposição hormonal otimizada. Cirurgias gastrointestinais podem comprometer absorção de nutrientes específicos. A avaliação pós-cirúrgica deve considerar necessidades nutricionais aumentadas, possíveis deficiências e ajustes metabólicos necessários.',
  interpretation_guide = 'Ao documentar histórico de registrar quaisquer cirurgias realizadas, registre: (1) data da cirurgia e indicação; (2) tipo de procedimento (total vs parcial); (3) complicações ou sequelas; (4) terapias de reposição atuais; (5) sintomas persistentes relacionados. Avalie necessidades específicas pós-cirúrgicas incluindo reposição hormonal adequada, suplementação nutricional direcionada, monitoramento de marcadores relevantes e rastreio de complicações tardias. Em cirurgias endócrinas, otimize reposição hormonal com base em sintomas clínicos e biomarcadores, não apenas valores laboratoriais de referência.',
  recommendations = '["Monitorar regularmente marcadores bioquímicos relevantes após registrar quaisquer cirurgias realizadas", "Avaliar necessidade de reposição hormonal otimizada com base em sintomas e laboratório", "Implementar suplementação nutricional direcionada para deficiências específicas relacionadas à cirurgia", "Realizar acompanhamento periódico com especialista apropriado", "Otimizar suporte nutricional para compensar alterações metabólicas pós-cirúrgicas"]'::jsonb,
  related_markers = '["Marcadores específicos do órgão removido", "Perfil hormonal completo", "Avaliação nutricional abrangente", "Marcadores inflamatórios", "Densitometria óssea (se aplicável)"]'::jsonb,
  articles_suggestions = '["Vida após registrar quaisquer cirurgias realizadas: otimizando saúde e qualidade de vida", "Suporte nutricional após cirurgias: o que você precisa saber", "Reposição hormonal otimizada na medicina integrativa", "Monitoramento de saúde após procedimentos cirúrgicos"]'::jsonb,
  updated_at = NOW()
WHERE id = '4511ae8d-2ad3-40e0-a47d-2cef065d39e9';


-- Cirurgias que interferem diretamente no escore (Cirurgias já realizadas)
UPDATE score_items
SET
  clinical_relevance = 'O histórico de cirurgias que interferem diretamente no escore tem implicações importantes para avaliação funcional e estratégias terapêuticas. Cirurgias que removem órgãos ou alteram anatomia podem impactar função metabólica, produção hormonal, absorção nutricional, microbioma e sistemas de detoxificação. A Medicina Funcional Integrativa reconhece que mesmo após recuperação cirúrgica bem-sucedida, podem persistir desequilíbrios fisiológicos que requerem suporte específico. A remoção de tecidos endócrinos (tireoide, ovários, próstata) necessita reposição hormonal otimizada. Cirurgias gastrointestinais podem comprometer absorção de nutrientes específicos. A avaliação pós-cirúrgica deve considerar necessidades nutricionais aumentadas, possíveis deficiências e ajustes metabólicos necessários.',
  interpretation_guide = 'Ao documentar histórico de cirurgias que interferem diretamente no escore, registre: (1) data da cirurgia e indicação; (2) tipo de procedimento (total vs parcial); (3) complicações ou sequelas; (4) terapias de reposição atuais; (5) sintomas persistentes relacionados. Avalie necessidades específicas pós-cirúrgicas incluindo reposição hormonal adequada, suplementação nutricional direcionada, monitoramento de marcadores relevantes e rastreio de complicações tardias. Em cirurgias endócrinas, otimize reposição hormonal com base em sintomas clínicos e biomarcadores, não apenas valores laboratoriais de referência.',
  recommendations = '["Monitorar regularmente marcadores bioquímicos relevantes após cirurgias que interferem diretamente no escore", "Avaliar necessidade de reposição hormonal otimizada com base em sintomas e laboratório", "Implementar suplementação nutricional direcionada para deficiências específicas relacionadas à cirurgia", "Realizar acompanhamento periódico com especialista apropriado", "Otimizar suporte nutricional para compensar alterações metabólicas pós-cirúrgicas"]'::jsonb,
  related_markers = '["Marcadores específicos do órgão removido", "Perfil hormonal completo", "Avaliação nutricional abrangente", "Marcadores inflamatórios", "Densitometria óssea (se aplicável)"]'::jsonb,
  articles_suggestions = '["Vida após cirurgias que interferem diretamente no escore: otimizando saúde e qualidade de vida", "Suporte nutricional após cirurgias: o que você precisa saber", "Reposição hormonal otimizada na medicina integrativa", "Monitoramento de saúde após procedimentos cirúrgicos"]'::jsonb,
  updated_at = NOW()
WHERE id = 'cbb42389-9a9d-48fe-a64a-61215107d5e3';


-- Mastectomia (Cirurgias já realizadas)
UPDATE score_items
SET
  clinical_relevance = 'O histórico de mastectomia tem implicações importantes para avaliação funcional e estratégias terapêuticas. Cirurgias que removem órgãos ou alteram anatomia podem impactar função metabólica, produção hormonal, absorção nutricional, microbioma e sistemas de detoxificação. A Medicina Funcional Integrativa reconhece que mesmo após recuperação cirúrgica bem-sucedida, podem persistir desequilíbrios fisiológicos que requerem suporte específico. A remoção de tecidos endócrinos (tireoide, ovários, próstata) necessita reposição hormonal otimizada. Cirurgias gastrointestinais podem comprometer absorção de nutrientes específicos. A avaliação pós-cirúrgica deve considerar necessidades nutricionais aumentadas, possíveis deficiências e ajustes metabólicos necessários.',
  interpretation_guide = 'Ao documentar histórico de mastectomia, registre: (1) data da cirurgia e indicação; (2) tipo de procedimento (total vs parcial); (3) complicações ou sequelas; (4) terapias de reposição atuais; (5) sintomas persistentes relacionados. Avalie necessidades específicas pós-cirúrgicas incluindo reposição hormonal adequada, suplementação nutricional direcionada, monitoramento de marcadores relevantes e rastreio de complicações tardias. Em cirurgias endócrinas, otimize reposição hormonal com base em sintomas clínicos e biomarcadores, não apenas valores laboratoriais de referência.',
  recommendations = '["Monitorar regularmente marcadores bioquímicos relevantes após mastectomia", "Avaliar necessidade de reposição hormonal otimizada com base em sintomas e laboratório", "Implementar suplementação nutricional direcionada para deficiências específicas relacionadas à cirurgia", "Realizar acompanhamento periódico com especialista apropriado", "Otimizar suporte nutricional para compensar alterações metabólicas pós-cirúrgicas"]'::jsonb,
  related_markers = '["Marcadores específicos do órgão removido", "Perfil hormonal completo", "Avaliação nutricional abrangente", "Marcadores inflamatórios", "Densitometria óssea (se aplicável)"]'::jsonb,
  articles_suggestions = '["Vida após mastectomia: otimizando saúde e qualidade de vida", "Suporte nutricional após cirurgias: o que você precisa saber", "Reposição hormonal otimizada na medicina integrativa", "Monitoramento de saúde após procedimentos cirúrgicos"]'::jsonb,
  updated_at = NOW()
WHERE id = 'e65c56dc-5c07-4270-8a8b-017b293ca147';


-- Prostatectomia (Cirurgias já realizadas)
UPDATE score_items
SET
  clinical_relevance = 'O histórico de prostatectomia tem implicações importantes para avaliação funcional e estratégias terapêuticas. Cirurgias que removem órgãos ou alteram anatomia podem impactar função metabólica, produção hormonal, absorção nutricional, microbioma e sistemas de detoxificação. A Medicina Funcional Integrativa reconhece que mesmo após recuperação cirúrgica bem-sucedida, podem persistir desequilíbrios fisiológicos que requerem suporte específico. A remoção de tecidos endócrinos (tireoide, ovários, próstata) necessita reposição hormonal otimizada. Cirurgias gastrointestinais podem comprometer absorção de nutrientes específicos. A avaliação pós-cirúrgica deve considerar necessidades nutricionais aumentadas, possíveis deficiências e ajustes metabólicos necessários.',
  interpretation_guide = 'Ao documentar histórico de prostatectomia, registre: (1) data da cirurgia e indicação; (2) tipo de procedimento (total vs parcial); (3) complicações ou sequelas; (4) terapias de reposição atuais; (5) sintomas persistentes relacionados. Avalie necessidades específicas pós-cirúrgicas incluindo reposição hormonal adequada, suplementação nutricional direcionada, monitoramento de marcadores relevantes e rastreio de complicações tardias. Em cirurgias endócrinas, otimize reposição hormonal com base em sintomas clínicos e biomarcadores, não apenas valores laboratoriais de referência.',
  recommendations = '["Monitorar regularmente marcadores bioquímicos relevantes após prostatectomia", "Avaliar necessidade de reposição hormonal otimizada com base em sintomas e laboratório", "Implementar suplementação nutricional direcionada para deficiências específicas relacionadas à cirurgia", "Realizar acompanhamento periódico com especialista apropriado", "Otimizar suporte nutricional para compensar alterações metabólicas pós-cirúrgicas"]'::jsonb,
  related_markers = '["Marcadores específicos do órgão removido", "Perfil hormonal completo", "Avaliação nutricional abrangente", "Marcadores inflamatórios", "Densitometria óssea (se aplicável)"]'::jsonb,
  articles_suggestions = '["Vida após prostatectomia: otimizando saúde e qualidade de vida", "Suporte nutricional após cirurgias: o que você precisa saber", "Reposição hormonal otimizada na medicina integrativa", "Monitoramento de saúde após procedimentos cirúrgicos"]'::jsonb,
  updated_at = NOW()
WHERE id = 'ad36ad5d-e587-43a8-b1ae-3e27305e25d8';


-- Tireoidectomia (Cirurgias já realizadas)
UPDATE score_items
SET
  clinical_relevance = 'O histórico de tireoidectomia tem implicações importantes para avaliação funcional e estratégias terapêuticas. Cirurgias que removem órgãos ou alteram anatomia podem impactar função metabólica, produção hormonal, absorção nutricional, microbioma e sistemas de detoxificação. A Medicina Funcional Integrativa reconhece que mesmo após recuperação cirúrgica bem-sucedida, podem persistir desequilíbrios fisiológicos que requerem suporte específico. A remoção de tecidos endócrinos (tireoide, ovários, próstata) necessita reposição hormonal otimizada. Cirurgias gastrointestinais podem comprometer absorção de nutrientes específicos. A avaliação pós-cirúrgica deve considerar necessidades nutricionais aumentadas, possíveis deficiências e ajustes metabólicos necessários.',
  interpretation_guide = 'Ao documentar histórico de tireoidectomia, registre: (1) data da cirurgia e indicação; (2) tipo de procedimento (total vs parcial); (3) complicações ou sequelas; (4) terapias de reposição atuais; (5) sintomas persistentes relacionados. Avalie necessidades específicas pós-cirúrgicas incluindo reposição hormonal adequada, suplementação nutricional direcionada, monitoramento de marcadores relevantes e rastreio de complicações tardias. Em cirurgias endócrinas, otimize reposição hormonal com base em sintomas clínicos e biomarcadores, não apenas valores laboratoriais de referência.',
  recommendations = '["Monitorar regularmente marcadores bioquímicos relevantes após tireoidectomia", "Avaliar necessidade de reposição hormonal otimizada com base em sintomas e laboratório", "Implementar suplementação nutricional direcionada para deficiências específicas relacionadas à cirurgia", "Realizar acompanhamento periódico com especialista apropriado", "Otimizar suporte nutricional para compensar alterações metabólicas pós-cirúrgicas"]'::jsonb,
  related_markers = '["Marcadores específicos do órgão removido", "Perfil hormonal completo", "Avaliação nutricional abrangente", "Marcadores inflamatórios", "Densitometria óssea (se aplicável)"]'::jsonb,
  articles_suggestions = '["Vida após tireoidectomia: otimizando saúde e qualidade de vida", "Suporte nutricional após cirurgias: o que você precisa saber", "Reposição hormonal otimizada na medicina integrativa", "Monitoramento de saúde após procedimentos cirúrgicos"]'::jsonb,
  updated_at = NOW()
WHERE id = '8115cf13-eab0-4db2-9844-d04c37701d92';


-- Histerectomia (Cirurgias já realizadas)
UPDATE score_items
SET
  clinical_relevance = 'O histórico de histerectomia tem implicações importantes para avaliação funcional e estratégias terapêuticas. Cirurgias que removem órgãos ou alteram anatomia podem impactar função metabólica, produção hormonal, absorção nutricional, microbioma e sistemas de detoxificação. A Medicina Funcional Integrativa reconhece que mesmo após recuperação cirúrgica bem-sucedida, podem persistir desequilíbrios fisiológicos que requerem suporte específico. A remoção de tecidos endócrinos (tireoide, ovários, próstata) necessita reposição hormonal otimizada. Cirurgias gastrointestinais podem comprometer absorção de nutrientes específicos. A avaliação pós-cirúrgica deve considerar necessidades nutricionais aumentadas, possíveis deficiências e ajustes metabólicos necessários.',
  interpretation_guide = 'Ao documentar histórico de histerectomia, registre: (1) data da cirurgia e indicação; (2) tipo de procedimento (total vs parcial); (3) complicações ou sequelas; (4) terapias de reposição atuais; (5) sintomas persistentes relacionados. Avalie necessidades específicas pós-cirúrgicas incluindo reposição hormonal adequada, suplementação nutricional direcionada, monitoramento de marcadores relevantes e rastreio de complicações tardias. Em cirurgias endócrinas, otimize reposição hormonal com base em sintomas clínicos e biomarcadores, não apenas valores laboratoriais de referência.',
  recommendations = '["Monitorar regularmente marcadores bioquímicos relevantes após histerectomia", "Avaliar necessidade de reposição hormonal otimizada com base em sintomas e laboratório", "Implementar suplementação nutricional direcionada para deficiências específicas relacionadas à cirurgia", "Realizar acompanhamento periódico com especialista apropriado", "Otimizar suporte nutricional para compensar alterações metabólicas pós-cirúrgicas"]'::jsonb,
  related_markers = '["Marcadores específicos do órgão removido", "Perfil hormonal completo", "Avaliação nutricional abrangente", "Marcadores inflamatórios", "Densitometria óssea (se aplicável)"]'::jsonb,
  articles_suggestions = '["Vida após histerectomia: otimizando saúde e qualidade de vida", "Suporte nutricional após cirurgias: o que você precisa saber", "Reposição hormonal otimizada na medicina integrativa", "Monitoramento de saúde após procedimentos cirúrgicos"]'::jsonb,
  updated_at = NOW()
WHERE id = '19b1c83e-f2a7-47e9-9045-5c994cfcbf94';


-- Ooforectomia (Cirurgias já realizadas)
UPDATE score_items
SET
  clinical_relevance = 'O histórico de ooforectomia tem implicações importantes para avaliação funcional e estratégias terapêuticas. Cirurgias que removem órgãos ou alteram anatomia podem impactar função metabólica, produção hormonal, absorção nutricional, microbioma e sistemas de detoxificação. A Medicina Funcional Integrativa reconhece que mesmo após recuperação cirúrgica bem-sucedida, podem persistir desequilíbrios fisiológicos que requerem suporte específico. A remoção de tecidos endócrinos (tireoide, ovários, próstata) necessita reposição hormonal otimizada. Cirurgias gastrointestinais podem comprometer absorção de nutrientes específicos. A avaliação pós-cirúrgica deve considerar necessidades nutricionais aumentadas, possíveis deficiências e ajustes metabólicos necessários.',
  interpretation_guide = 'Ao documentar histórico de ooforectomia, registre: (1) data da cirurgia e indicação; (2) tipo de procedimento (total vs parcial); (3) complicações ou sequelas; (4) terapias de reposição atuais; (5) sintomas persistentes relacionados. Avalie necessidades específicas pós-cirúrgicas incluindo reposição hormonal adequada, suplementação nutricional direcionada, monitoramento de marcadores relevantes e rastreio de complicações tardias. Em cirurgias endócrinas, otimize reposição hormonal com base em sintomas clínicos e biomarcadores, não apenas valores laboratoriais de referência.',
  recommendations = '["Monitorar regularmente marcadores bioquímicos relevantes após ooforectomia", "Avaliar necessidade de reposição hormonal otimizada com base em sintomas e laboratório", "Implementar suplementação nutricional direcionada para deficiências específicas relacionadas à cirurgia", "Realizar acompanhamento periódico com especialista apropriado", "Otimizar suporte nutricional para compensar alterações metabólicas pós-cirúrgicas"]'::jsonb,
  related_markers = '["Marcadores específicos do órgão removido", "Perfil hormonal completo", "Avaliação nutricional abrangente", "Marcadores inflamatórios", "Densitometria óssea (se aplicável)"]'::jsonb,
  articles_suggestions = '["Vida após ooforectomia: otimizando saúde e qualidade de vida", "Suporte nutricional após cirurgias: o que você precisa saber", "Reposição hormonal otimizada na medicina integrativa", "Monitoramento de saúde após procedimentos cirúrgicos"]'::jsonb,
  updated_at = NOW()
WHERE id = 'fca251a2-5aa2-42f9-b93c-0f0a852d9d51';


-- Orquiectomia (Cirurgias já realizadas)
UPDATE score_items
SET
  clinical_relevance = 'O histórico de orquiectomia tem implicações importantes para avaliação funcional e estratégias terapêuticas. Cirurgias que removem órgãos ou alteram anatomia podem impactar função metabólica, produção hormonal, absorção nutricional, microbioma e sistemas de detoxificação. A Medicina Funcional Integrativa reconhece que mesmo após recuperação cirúrgica bem-sucedida, podem persistir desequilíbrios fisiológicos que requerem suporte específico. A remoção de tecidos endócrinos (tireoide, ovários, próstata) necessita reposição hormonal otimizada. Cirurgias gastrointestinais podem comprometer absorção de nutrientes específicos. A avaliação pós-cirúrgica deve considerar necessidades nutricionais aumentadas, possíveis deficiências e ajustes metabólicos necessários.',
  interpretation_guide = 'Ao documentar histórico de orquiectomia, registre: (1) data da cirurgia e indicação; (2) tipo de procedimento (total vs parcial); (3) complicações ou sequelas; (4) terapias de reposição atuais; (5) sintomas persistentes relacionados. Avalie necessidades específicas pós-cirúrgicas incluindo reposição hormonal adequada, suplementação nutricional direcionada, monitoramento de marcadores relevantes e rastreio de complicações tardias. Em cirurgias endócrinas, otimize reposição hormonal com base em sintomas clínicos e biomarcadores, não apenas valores laboratoriais de referência.',
  recommendations = '["Monitorar regularmente marcadores bioquímicos relevantes após orquiectomia", "Avaliar necessidade de reposição hormonal otimizada com base em sintomas e laboratório", "Implementar suplementação nutricional direcionada para deficiências específicas relacionadas à cirurgia", "Realizar acompanhamento periódico com especialista apropriado", "Otimizar suporte nutricional para compensar alterações metabólicas pós-cirúrgicas"]'::jsonb,
  related_markers = '["Marcadores específicos do órgão removido", "Perfil hormonal completo", "Avaliação nutricional abrangente", "Marcadores inflamatórios", "Densitometria óssea (se aplicável)"]'::jsonb,
  articles_suggestions = '["Vida após orquiectomia: otimizando saúde e qualidade de vida", "Suporte nutricional após cirurgias: o que você precisa saber", "Reposição hormonal otimizada na medicina integrativa", "Monitoramento de saúde após procedimentos cirúrgicos"]'::jsonb,
  updated_at = NOW()
WHERE id = '925c2474-689e-486a-a6b1-3e1b4ca8cc6e';


-- Nefrectomia (Cirurgias já realizadas)
UPDATE score_items
SET
  clinical_relevance = 'O histórico de nefrectomia tem implicações importantes para avaliação funcional e estratégias terapêuticas. Cirurgias que removem órgãos ou alteram anatomia podem impactar função metabólica, produção hormonal, absorção nutricional, microbioma e sistemas de detoxificação. A Medicina Funcional Integrativa reconhece que mesmo após recuperação cirúrgica bem-sucedida, podem persistir desequilíbrios fisiológicos que requerem suporte específico. A remoção de tecidos endócrinos (tireoide, ovários, próstata) necessita reposição hormonal otimizada. Cirurgias gastrointestinais podem comprometer absorção de nutrientes específicos. A avaliação pós-cirúrgica deve considerar necessidades nutricionais aumentadas, possíveis deficiências e ajustes metabólicos necessários.',
  interpretation_guide = 'Ao documentar histórico de nefrectomia, registre: (1) data da cirurgia e indicação; (2) tipo de procedimento (total vs parcial); (3) complicações ou sequelas; (4) terapias de reposição atuais; (5) sintomas persistentes relacionados. Avalie necessidades específicas pós-cirúrgicas incluindo reposição hormonal adequada, suplementação nutricional direcionada, monitoramento de marcadores relevantes e rastreio de complicações tardias. Em cirurgias endócrinas, otimize reposição hormonal com base em sintomas clínicos e biomarcadores, não apenas valores laboratoriais de referência.',
  recommendations = '["Monitorar regularmente marcadores bioquímicos relevantes após nefrectomia", "Avaliar necessidade de reposição hormonal otimizada com base em sintomas e laboratório", "Implementar suplementação nutricional direcionada para deficiências específicas relacionadas à cirurgia", "Realizar acompanhamento periódico com especialista apropriado", "Otimizar suporte nutricional para compensar alterações metabólicas pós-cirúrgicas"]'::jsonb,
  related_markers = '["Marcadores específicos do órgão removido", "Perfil hormonal completo", "Avaliação nutricional abrangente", "Marcadores inflamatórios", "Densitometria óssea (se aplicável)"]'::jsonb,
  articles_suggestions = '["Vida após nefrectomia: otimizando saúde e qualidade de vida", "Suporte nutricional após cirurgias: o que você precisa saber", "Reposição hormonal otimizada na medicina integrativa", "Monitoramento de saúde após procedimentos cirúrgicos"]'::jsonb,
  updated_at = NOW()
WHERE id = 'e5ea8f41-05e0-48e0-b6a8-2323b3896449';


-- Hepatectomia parcial (Cirurgias já realizadas)
UPDATE score_items
SET
  clinical_relevance = 'O histórico de hepatectomia parcial tem implicações importantes para avaliação funcional e estratégias terapêuticas. Cirurgias que removem órgãos ou alteram anatomia podem impactar função metabólica, produção hormonal, absorção nutricional, microbioma e sistemas de detoxificação. A Medicina Funcional Integrativa reconhece que mesmo após recuperação cirúrgica bem-sucedida, podem persistir desequilíbrios fisiológicos que requerem suporte específico. A remoção de tecidos endócrinos (tireoide, ovários, próstata) necessita reposição hormonal otimizada. Cirurgias gastrointestinais podem comprometer absorção de nutrientes específicos. A avaliação pós-cirúrgica deve considerar necessidades nutricionais aumentadas, possíveis deficiências e ajustes metabólicos necessários.',
  interpretation_guide = 'Ao documentar histórico de hepatectomia parcial, registre: (1) data da cirurgia e indicação; (2) tipo de procedimento (total vs parcial); (3) complicações ou sequelas; (4) terapias de reposição atuais; (5) sintomas persistentes relacionados. Avalie necessidades específicas pós-cirúrgicas incluindo reposição hormonal adequada, suplementação nutricional direcionada, monitoramento de marcadores relevantes e rastreio de complicações tardias. Em cirurgias endócrinas, otimize reposição hormonal com base em sintomas clínicos e biomarcadores, não apenas valores laboratoriais de referência.',
  recommendations = '["Monitorar regularmente marcadores bioquímicos relevantes após hepatectomia parcial", "Avaliar necessidade de reposição hormonal otimizada com base em sintomas e laboratório", "Implementar suplementação nutricional direcionada para deficiências específicas relacionadas à cirurgia", "Realizar acompanhamento periódico com especialista apropriado", "Otimizar suporte nutricional para compensar alterações metabólicas pós-cirúrgicas"]'::jsonb,
  related_markers = '["Marcadores específicos do órgão removido", "Perfil hormonal completo", "Avaliação nutricional abrangente", "Marcadores inflamatórios", "Densitometria óssea (se aplicável)"]'::jsonb,
  articles_suggestions = '["Vida após hepatectomia parcial: otimizando saúde e qualidade de vida", "Suporte nutricional após cirurgias: o que você precisa saber", "Reposição hormonal otimizada na medicina integrativa", "Monitoramento de saúde após procedimentos cirúrgicos"]'::jsonb,
  updated_at = NOW()
WHERE id = 'f691cda1-bd2d-449e-9e66-045319ceaaa3';


-- Lobectomia/pneumectomia (Cirurgias já realizadas)
UPDATE score_items
SET
  clinical_relevance = 'O histórico de lobectomia/pneumectomia tem implicações importantes para avaliação funcional e estratégias terapêuticas. Cirurgias que removem órgãos ou alteram anatomia podem impactar função metabólica, produção hormonal, absorção nutricional, microbioma e sistemas de detoxificação. A Medicina Funcional Integrativa reconhece que mesmo após recuperação cirúrgica bem-sucedida, podem persistir desequilíbrios fisiológicos que requerem suporte específico. A remoção de tecidos endócrinos (tireoide, ovários, próstata) necessita reposição hormonal otimizada. Cirurgias gastrointestinais podem comprometer absorção de nutrientes específicos. A avaliação pós-cirúrgica deve considerar necessidades nutricionais aumentadas, possíveis deficiências e ajustes metabólicos necessários.',
  interpretation_guide = 'Ao documentar histórico de lobectomia/pneumectomia, registre: (1) data da cirurgia e indicação; (2) tipo de procedimento (total vs parcial); (3) complicações ou sequelas; (4) terapias de reposição atuais; (5) sintomas persistentes relacionados. Avalie necessidades específicas pós-cirúrgicas incluindo reposição hormonal adequada, suplementação nutricional direcionada, monitoramento de marcadores relevantes e rastreio de complicações tardias. Em cirurgias endócrinas, otimize reposição hormonal com base em sintomas clínicos e biomarcadores, não apenas valores laboratoriais de referência.',
  recommendations = '["Monitorar regularmente marcadores bioquímicos relevantes após lobectomia/pneumectomia", "Avaliar necessidade de reposição hormonal otimizada com base em sintomas e laboratório", "Implementar suplementação nutricional direcionada para deficiências específicas relacionadas à cirurgia", "Realizar acompanhamento periódico com especialista apropriado", "Otimizar suporte nutricional para compensar alterações metabólicas pós-cirúrgicas"]'::jsonb,
  related_markers = '["Marcadores específicos do órgão removido", "Perfil hormonal completo", "Avaliação nutricional abrangente", "Marcadores inflamatórios", "Densitometria óssea (se aplicável)"]'::jsonb,
  articles_suggestions = '["Vida após lobectomia/pneumectomia: otimizando saúde e qualidade de vida", "Suporte nutricional após cirurgias: o que você precisa saber", "Reposição hormonal otimizada na medicina integrativa", "Monitoramento de saúde após procedimentos cirúrgicos"]'::jsonb,
  updated_at = NOW()
WHERE id = 'ec0bbb80-873a-4946-bf90-bc759eddb080';


-- Craniotomia (Cirurgias já realizadas)
UPDATE score_items
SET
  clinical_relevance = 'O histórico de craniotomia tem implicações importantes para avaliação funcional e estratégias terapêuticas. Cirurgias que removem órgãos ou alteram anatomia podem impactar função metabólica, produção hormonal, absorção nutricional, microbioma e sistemas de detoxificação. A Medicina Funcional Integrativa reconhece que mesmo após recuperação cirúrgica bem-sucedida, podem persistir desequilíbrios fisiológicos que requerem suporte específico. A remoção de tecidos endócrinos (tireoide, ovários, próstata) necessita reposição hormonal otimizada. Cirurgias gastrointestinais podem comprometer absorção de nutrientes específicos. A avaliação pós-cirúrgica deve considerar necessidades nutricionais aumentadas, possíveis deficiências e ajustes metabólicos necessários.',
  interpretation_guide = 'Ao documentar histórico de craniotomia, registre: (1) data da cirurgia e indicação; (2) tipo de procedimento (total vs parcial); (3) complicações ou sequelas; (4) terapias de reposição atuais; (5) sintomas persistentes relacionados. Avalie necessidades específicas pós-cirúrgicas incluindo reposição hormonal adequada, suplementação nutricional direcionada, monitoramento de marcadores relevantes e rastreio de complicações tardias. Em cirurgias endócrinas, otimize reposição hormonal com base em sintomas clínicos e biomarcadores, não apenas valores laboratoriais de referência.',
  recommendations = '["Monitorar regularmente marcadores bioquímicos relevantes após craniotomia", "Avaliar necessidade de reposição hormonal otimizada com base em sintomas e laboratório", "Implementar suplementação nutricional direcionada para deficiências específicas relacionadas à cirurgia", "Realizar acompanhamento periódico com especialista apropriado", "Otimizar suporte nutricional para compensar alterações metabólicas pós-cirúrgicas"]'::jsonb,
  related_markers = '["Marcadores específicos do órgão removido", "Perfil hormonal completo", "Avaliação nutricional abrangente", "Marcadores inflamatórios", "Densitometria óssea (se aplicável)"]'::jsonb,
  articles_suggestions = '["Vida após craniotomia: otimizando saúde e qualidade de vida", "Suporte nutricional após cirurgias: o que você precisa saber", "Reposição hormonal otimizada na medicina integrativa", "Monitoramento de saúde após procedimentos cirúrgicos"]'::jsonb,
  updated_at = NOW()
WHERE id = '53dbbf23-25bf-4b91-9670-b17fab93c6e9';


-- Cirurgia de epilepsia (Cirurgias já realizadas)
UPDATE score_items
SET
  clinical_relevance = 'O histórico de cirurgia de epilepsia tem implicações importantes para avaliação funcional e estratégias terapêuticas. Cirurgias que removem órgãos ou alteram anatomia podem impactar função metabólica, produção hormonal, absorção nutricional, microbioma e sistemas de detoxificação. A Medicina Funcional Integrativa reconhece que mesmo após recuperação cirúrgica bem-sucedida, podem persistir desequilíbrios fisiológicos que requerem suporte específico. A remoção de tecidos endócrinos (tireoide, ovários, próstata) necessita reposição hormonal otimizada. Cirurgias gastrointestinais podem comprometer absorção de nutrientes específicos. A avaliação pós-cirúrgica deve considerar necessidades nutricionais aumentadas, possíveis deficiências e ajustes metabólicos necessários.',
  interpretation_guide = 'Ao documentar histórico de cirurgia de epilepsia, registre: (1) data da cirurgia e indicação; (2) tipo de procedimento (total vs parcial); (3) complicações ou sequelas; (4) terapias de reposição atuais; (5) sintomas persistentes relacionados. Avalie necessidades específicas pós-cirúrgicas incluindo reposição hormonal adequada, suplementação nutricional direcionada, monitoramento de marcadores relevantes e rastreio de complicações tardias. Em cirurgias endócrinas, otimize reposição hormonal com base em sintomas clínicos e biomarcadores, não apenas valores laboratoriais de referência.',
  recommendations = '["Monitorar regularmente marcadores bioquímicos relevantes após cirurgia de epilepsia", "Avaliar necessidade de reposição hormonal otimizada com base em sintomas e laboratório", "Implementar suplementação nutricional direcionada para deficiências específicas relacionadas à cirurgia", "Realizar acompanhamento periódico com especialista apropriado", "Otimizar suporte nutricional para compensar alterações metabólicas pós-cirúrgicas"]'::jsonb,
  related_markers = '["Marcadores específicos do órgão removido", "Perfil hormonal completo", "Avaliação nutricional abrangente", "Marcadores inflamatórios", "Densitometria óssea (se aplicável)"]'::jsonb,
  articles_suggestions = '["Vida após cirurgia de epilepsia: otimizando saúde e qualidade de vida", "Suporte nutricional após cirurgias: o que você precisa saber", "Reposição hormonal otimizada na medicina integrativa", "Monitoramento de saúde após procedimentos cirúrgicos"]'::jsonb,
  updated_at = NOW()
WHERE id = '8c2456de-1ba1-4a1f-ab7c-964f8c8114e6';

-- Verificação final
SELECT
  COUNT(*) as total_items,
  COUNT(*) FILTER (WHERE clinical_relevance IS NOT NULL AND clinical_relevance != '') as com_relevancia,
  COUNT(*) FILTER (WHERE interpretation_guide IS NOT NULL AND interpretation_guide != '') as com_interpretacao,
  COUNT(*) FILTER (WHERE jsonb_array_length(recommendations) > 0) as com_recomendacoes,
  COUNT(*) FILTER (WHERE jsonb_array_length(related_markers) > 0) as com_marcadores,
  COUNT(*) FILTER (WHERE jsonb_array_length(articles_suggestions) > 0) as com_artigos
FROM score_items
WHERE id IN (
  '1176540d-cefa-4d2c-b5e2-4a992060de4d',
  '360d1e6a-84c5-4763-a743-0fce76fe2686',
  'adcb06a3-4791-4ce2-ae05-c01f8fef9ead',
  '65a28dbf-9466-4a23-9410-9d6034c3e141',
  'b7370158-c34e-449f-8305-537a8fb85d98',
  'e24dae19-4cb0-4d83-a6db-9571aabf9bde',
  'a7fdd0e5-9aa7-4604-b5ae-731c05c90af0',
  '40837ae0-790a-4a1a-a7ea-fb9d0cc21878',
  'b4cda178-b5a9-4193-b496-afe7ce51ed91',
  'deb8e2ee-27a7-4ac0-a743-d6cb400f4b27',
  'fdb0c742-1db1-48ba-95bc-91e454256d84',
  'e375191b-d8be-4044-9d59-b71fa222cbaa',
  '44f501b3-6c1d-4a31-b8c9-0f54b201dbcb',
  '14881a00-579e-4297-bd6d-4c83d8081d2d',
  'a7698447-ddbf-4bee-92b4-81a1559562ad',
  '76578167-9f69-4981-b73d-46aa7c71ec6b',
  '23b90c6c-6f44-481a-8255-108158a64239',
  '4d1372ee-5bee-4ce8-845a-a4e176109ad2',
  'fe938ed9-3cf9-4ddf-a1ae-2f67c03e54a4',
  '5ddc4af9-dd19-4ea1-8117-3d3e30377dab',
  'e010b2a7-6e9e-4b42-9d37-487e91411a18',
  '933b7816-fe10-4e35-94d7-0d232cc258ce',
  '2b01e2f0-76c7-4616-8f7b-33b1c9d11279',
  '33ffce34-e1fb-4e31-b393-f2783549d874',
  'f54dbfaa-599a-40ee-907a-28431ce4858a',
  '30a7764c-4c11-4078-b942-860ee56136a4',
  '1ff1e7f8-5b15-4f0c-b76b-806c1d6ff1fd',
  '4511ae8d-2ad3-40e0-a47d-2cef065d39e9',
  'cbb42389-9a9d-48fe-a64a-61215107d5e3',
  'e65c56dc-5c07-4270-8a8b-017b293ca147',
  'ad36ad5d-e587-43a8-b1ae-3e27305e25d8',
  '8115cf13-eab0-4db2-9844-d04c37701d92',
  '19b1c83e-f2a7-47e9-9045-5c994cfcbf94',
  'fca251a2-5aa2-42f9-b93c-0f0a852d9d51',
  '925c2474-689e-486a-a6b1-3e1b4ca8cc6e',
  'e5ea8f41-05e0-48e0-b6a8-2323b3896449',
  'f691cda1-bd2d-449e-9e66-045319ceaaa3',
  'ec0bbb80-873a-4946-bf90-bc759eddb080',
  '53dbbf23-25bf-4b91-9670-b17fab93c6e9',
  '8c2456de-1ba1-4a1f-ab7c-964f8c8114e6'
);

COMMIT;

-- ============================================================================
-- FIM DO BATCH FINAL 4 - 40 ITEMS ENRIQUECIDOS
-- ============================================================================
