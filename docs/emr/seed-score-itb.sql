-- ============================================================================
-- Seed: Índice Tornozelo-Braquial (ITB) — exame no catálogo + item do Escore.
-- Ancorado no escore de cálcio coronariano (aterosclerose subclínica): subgrupo
-- Imagem, pilares Cardiovascular + Metabólico, numeric_classifier.
-- Topologia em U: baixo (≤0,90 = DAP) e muito alto (≥1,41 = artéria
-- não-compressível/calcificação medial) são ruins; ótimo é 1,00–1,40.
-- Procedimento de consultório (Doppler de bolso + manguito), dirigido.
-- TUSS NULL: pesquisado online (Tabela 22/SBN, iClinic, operadoras) — o ITB
-- NÃO tem TUSS dedicado; é faturado sob código de Doppler que varia por
-- operadora. Não cravar código (regra: não chutar dado verificável). 12 pts.
-- Idempotente. APLICAR EM DEV; prod só sob ordem.
-- ============================================================================
-- IDs fixos: lab_def ...00a1 · score_item ...00c2 · levels ...00f0..f5

BEGIN;

-- 1) Catálogo de exames -------------------------------------------------------
INSERT INTO lab_test_definitions
  (id, code, name, short_name, tuss_code, category, is_requestable, result_type,
   unit, specimen_type, fasting_hours, sex_applicability, description,
   display_order, is_active, created_at, updated_at)
SELECT
  '0c9e7d10-0000-7000-8000-0000000000a1', 'PLNITBABI01',
  'Índice tornozelo-braquial (ITB)', 'ITB',
  NULL, 'imaging', true, 'numeric',
  NULL, NULL, NULL, 'all',
  'Razão entre a pressão sistólica do tornozelo e a do braço (Doppler). Rastreia doença arterial periférica e estratifica risco cardiovascular.',
  0, true, now(), now()
WHERE NOT EXISTS (SELECT 1 FROM lab_test_definitions WHERE code = 'PLNITBABI01');

-- 2) Score item ---------------------------------------------------------------
INSERT INTO score_items
  (id, name, unit, points, "order", subgroup_id, lab_test_code, gender,
   site_render_type, clinical_relevance, patient_explanation, conduct,
   site_explanation, created_at, updated_at)
SELECT
  '0c9e7d10-0000-7000-8000-0000000000c2',
  'Índice tornozelo-braquial (ITB)',
  NULL,
  12,
  COALESCE((SELECT "order" FROM score_items WHERE name = 'TC coração para escore de cálcio coronariano' AND deleted_at IS NULL LIMIT 1), 0) + 1,
  '019bf31d-2ef0-791d-b3c3-558b25e27fe5',  -- subgrupo Imagem
  'PLNITBABI01',
  'not_applicable',
  'numeric_classifier',
$cr$O índice tornozelo-braquial (ITB) é a razão entre a pressão arterial sistólica do tornozelo e a do braço, aferida com Doppler de onda contínua. É o teste não-invasivo de primeira linha para doença arterial periférica (DAP) e um marcador independente de risco cardiovascular global: um ITB baixo prediz mortalidade cardiovascular e por todas as causas mesmo em assintomáticos, refinando a estratificação além dos fatores clássicos.

**Técnica:** medir a pressão sistólica nas artérias braquiais (ambos os braços) e nas artérias pediosa dorsal e tibial posterior de cada tornozelo. O ITB de cada perna = maior pressão do tornozelo daquela perna ÷ maior das duas pressões braquiais. Repouso ≥10 min, decúbito dorsal.

**Interpretação (AHA/ACC):**
- ≥1,41 — artéria não-compressível (calcificação medial): ITB não confiável para DAP; comum em diabetes e doença renal crônica; ainda assim associa-se a risco cardiovascular elevado
- 1,00–1,40 — normal
- 0,91–0,99 — limítrofe
- ≤0,90 — diagnóstico de DAP
  - 0,41–0,90 — DAP leve a moderada
  - ≤0,40 — DAP grave (risco de isquemia crítica de membro)

**Quando o ITB de repouso é normal mas há suspeita clínica (claudicação típica):** considerar ITB pós-exercício (esteira), que desmascara DAP funcional.

**Valor de rastreio:** dirigido — maior rendimento em diabéticos, tabagistas/ex-tabagistas, idosos (>65 anos, ou >50 com fator de risco) e em qualquer sintoma de claudicação. Em assintomáticos de baixo risco da população geral, o rastreio universal não tem evidência (USPSTF "I").

**Caveat do valor alto:** ITB ≥1,41 não é "bom" — indica artérias rígidas/calcificadas e não permite avaliar DAP. Nesses casos, usar índice hálux-braquial (pressão de dedo) e investigar diabetes/doença renal.

**Marcadores complementares:** escore de cálcio coronariano, Doppler de carótidas (placa), perfil lipídico/ApoB, glicemia/HbA1c, PCR ultrassensível, função renal.$cr$,
$pe$O ITB compara a pressão do sangue no seu tornozelo com a do seu braço. Quando as artérias das pernas estão entupindo (placas de gordura), a pressão lá embaixo cai, e essa conta fica baixa.

O ideal é ficar entre 1,00 e 1,40. Valores abaixo de 0,90 indicam que o sangue está chegando com dificuldade às pernas (doença arterial periférica), o que também sinaliza maior risco para o coração. Valores muito altos (1,41 ou mais) não são bons: significam artérias endurecidas e calcificadas, comuns em quem tem diabetes ou problema nos rins.

É um exame simples, feito no consultório com um aparelhinho de ultrassom e o manguito de pressão, sem agulhas. Quase tudo que ele detecta cedo melhora com parar de fumar, atividade física, controle do colesterol, do açúcar e da pressão.$pe$,
$co$## Conduta Clínica — Índice Tornozelo-Braquial

Interpretar sempre junto do contexto (sintomas de claudicação, diabetes, tabagismo, função renal). Em diabético/renal com ITB ≥1,41, o índice é inválido para DAP — usar pressão de dedo (índice hálux-braquial).

### Estratificação por faixa

| ITB | Interpretação | Conduta |
|---|---|---|
| ≤0,40 | DAP grave / risco de isquemia crítica | Encaminhamento vascular urgente; avaliar revascularização; cuidado com lesões de pé |
| 0,41–0,90 | DAP leve a moderada | Antiagregante + estatina de alta potência; cessação de tabagismo; exercício supervisionado; tratar claudicação; controle pressórico e glicêmico |
| 0,91–0,99 | Limítrofe | Repetir/observar; otimizar fatores de risco; considerar ITB pós-exercício se sintoma |
| 1,00–1,40 | Normal | Manutenção; rastreio dirigido conforme risco |
| ≥1,41 | Não-compressível (calcificação medial) | ITB inválido p/ DAP → índice hálux-braquial; investigar diabetes/DRC; tratar como alto risco cardiovascular |

### Avaliação complementar
- Escore de cálcio coronariano e Doppler de carótidas (placa) para mapear a aterosclerose em outros leitos
- Perfil lipídico / ApoB, glicemia e HbA1c, PCR ultrassensível, função renal
- ITB pós-exercício quando claudicação típica com ITB de repouso normal

### Medidas (qualquer faixa alterada)
- Cessação de tabagismo (maior fator modificável da DAP)
- Estatina de alta potência (reduz eventos e melhora desfecho do membro)
- Antiagregação conforme indicação; exercício/caminhada supervisionada
- Controle rigoroso de pressão arterial, glicemia e peso

### Encaminhamento
- Cirurgia vascular: ITB ≤0,40, claudicação limitante, dor isquêmica de repouso ou lesão trófica$co$,
$se$O ITB compara a pressão do tornozelo com a do braço para ver se o sangue chega bem às pernas. O ideal fica entre 1,00 e 1,40; abaixo de 0,90 indica artérias das pernas obstruídas (e maior risco para o coração), e 1,41 ou mais indica artérias endurecidas/calcificadas. É um teste simples de consultório, sem agulhas.$se$,
  now(), now()
WHERE NOT EXISTS (SELECT 1 FROM score_items WHERE id = '0c9e7d10-0000-7000-8000-0000000000c2');

-- 3) Levels (0-5 por qualidade; U-shape; multiplicador = level/5) --------------
INSERT INTO score_levels (id, item_id, level, name, lower_limit, upper_limit, operator, site_legend, created_at, updated_at)
SELECT v.id, '0c9e7d10-0000-7000-8000-0000000000c2', v.level, v.name, v.lo, v.hi, v.op, v.legend, now(), now()
FROM (VALUES
  ('0c9e7d10-0000-7000-8000-0000000000f0'::uuid, 0, '0,40 ou menos', NULL,  '0.40', '<=',      'Pressão nas pernas muito baixa: artérias gravemente obstruídas, sinal de alerta urgente.'),
  ('0c9e7d10-0000-7000-8000-0000000000f1'::uuid, 1, '0,41 a 0,70',   '0.41', '0.70', 'between', 'Doença arterial nas pernas moderada: sangue chegando com dificuldade, exige tratamento.'),
  ('0c9e7d10-0000-7000-8000-0000000000f2'::uuid, 2, '0,71 a 0,90',   '0.71', '0.90', 'between', 'Doença arterial nas pernas leve: início de obstrução, requer atenção e cuidado dos fatores de risco.'),
  ('0c9e7d10-0000-7000-8000-0000000000f3'::uuid, 3, '1,41 ou mais',  '1.41', NULL,   '>=',      'Artérias endurecidas e calcificadas: o exame não avalia bem a obstrução e indica risco cardiovascular alto.'),
  ('0c9e7d10-0000-7000-8000-0000000000f4'::uuid, 4, '0,91 a 0,99',   '0.91', '0.99', 'between', 'Índice no limite inferior do normal: acompanhar e cuidar dos fatores de risco.'),
  ('0c9e7d10-0000-7000-8000-0000000000f5'::uuid, 5, '1,00 a 1,40',   '1.00', '1.40', 'between', 'Índice tornozelo-braquial normal: boa circulação nas pernas, resultado ótimo.')
) AS v(id, level, name, lo, hi, op, legend)
WHERE NOT EXISTS (SELECT 1 FROM score_levels WHERE item_id = '0c9e7d10-0000-7000-8000-0000000000c2' AND level = v.level AND deleted_at IS NULL);

-- 4) Pilares (mesmos do escore de cálcio: Cardiovascular + Metabólico) ---------
INSERT INTO score_item_method_pillars (score_item_id, method_pillar_id)
SELECT '0c9e7d10-0000-7000-8000-0000000000c2', p.method_pillar_id
FROM score_item_method_pillars p
WHERE p.score_item_id = (SELECT id FROM score_items WHERE name = 'TC coração para escore de cálcio coronariano' AND deleted_at IS NULL LIMIT 1)
ON CONFLICT (score_item_id, method_pillar_id) DO NOTHING;

COMMIT;
