-- ============================================================================
-- Seed: Peptídeo C (jejum) — exame no catálogo + item pontuável do Escore Plenya
-- Espelha INSULINA 0 MIN (JEJUM): subgrupo Laboratoriais, 15 pts,
-- pilares Cardiovascular + Metabólico + Adiposidade e Risco Cardiometabólico,
-- 6 levels (0-5) por qualidade, multiplicador = level/5.
-- Faixas de medicina funcional integrativa (Lamkin Clinic / NHANES III).
-- Idempotente (re-execução segura). APLICAR EM DEV; prod só sob ordem.
-- ============================================================================

-- IDs fixos (idempotência)
-- lab_test_definition: 0c9e7d10-0000-7000-8000-0000000000a0
-- score_item:          0c9e7d10-0000-7000-8000-0000000000c1
-- levels L0..L5:       ...00d0 .. ...00d5

BEGIN;

-- 1) Catálogo de exames -------------------------------------------------------
INSERT INTO lab_test_definitions
  (id, code, name, short_name, tuss_code, category, is_requestable, result_type,
   unit, specimen_type, fasting_hours, sex_applicability, description,
   display_order, is_active, created_at, updated_at)
SELECT
  '0c9e7d10-0000-7000-8000-0000000000a0', 'PLNCPEPTD01', 'Peptídeo C', 'Pep-C',
  '40316394', 'hormones', true, 'numeric', 'ng/mL', 'Soro', 8, 'all',
  'Dosagem de peptídeo C de jejum — marcador mais fiel da secreção endógena de insulina e da reserva de célula β.',
  0, true, now(), now()
WHERE NOT EXISTS (SELECT 1 FROM lab_test_definitions WHERE code = 'PLNCPEPTD01');

-- 2) Score item ---------------------------------------------------------------
INSERT INTO score_items
  (id, name, unit, points, "order", subgroup_id, lab_test_code, gender,
   site_render_type, clinical_relevance, patient_explanation, conduct,
   site_explanation, created_at, updated_at)
SELECT
  '0c9e7d10-0000-7000-8000-0000000000c1',
  'Peptídeo C (jejum)',
  'ng/mL',
  15,
  COALESCE((SELECT "order" FROM score_items WHERE name = 'INSULINA 0 MIN (JEJUM)' AND deleted_at IS NULL LIMIT 1), 0),
  '019bf31d-2ef0-7f26-8df9-a0040192b8b8',  -- subgrupo Laboratoriais (grupo Exames)
  'PLNCPEPTD01',
  'not_applicable',
  'numeric_classifier',
$cr$O peptídeo C de jejum é o marcador mais fiel da secreção endógena de insulina. É clivado da proinsulina em quantidade equimolar à insulina (1:1) dentro da célula β, mas, ao contrário da insulina, não sofre extração hepática de primeira passagem (a insulina é extraída em ~50–60% pelo fígado), tem meia-vida de 30–35 min (vs. 4–5 min da insulina) e é depurado pelo rim. Por isso reflete melhor a taxa real de secreção de insulina, é mais estável e reprodutível, e permanece interpretável mesmo sob insulinoterapia (não há peptídeo C em insulina exógena).

**Valores de referência e interpretação funcional (jejum):**
- Ótimo (medicina funcional integrativa): 1,0–2,5 ng/mL — produção endógena adequada sem hiperinsulinemia
- Limite inferior: 0,6–1,0 ng/mL — reserva de célula β reduzida; avaliar LADA em contexto adequado
- Limite superior: 2,5–3,5 ng/mL — hiperinsulinemia limítrofe; investigar resistência insulínica
- Elevado: >3,5 ng/mL — hipersecreção por resistência insulínica (ou, raramente, insulinoma)
- Muito baixo: <0,6 ng/mL — falência de célula β (DM1, LADA, DM2 terminal)
- Faixa laboratorial convencional (~0,8–3,5 ng/mL) já engloba valores hiperinsulinêmicos do ponto de vista funcional

**Por que medir junto da insulina e da glicemia de jejum:**
Com glicemia e insulina de jejum no mesmo painel, o peptídeo C separa dois cenários que a insulina isolada não distingue: resistência insulínica com célula β preservada (peptídeo C normal-alto/alto) versus falência de célula β (peptídeo C baixo). Permite também acompanhar a reserva pancreática ao longo do tempo — queda seriada indica exaustão β progressiva; elevação após intervenção metabólica indica recuperação funcional por redução da glicotoxicidade.

**Valor prognóstico em não-diabéticos (público de longevidade):**
Coorte NHANES III (n≈9.211, sem diabetes): o maior quartil de peptídeo C de jejum vs. o menor associou-se a maior mortalidade geral (HR 1,80; IC95% 1,33–2,43), cardiovascular (HR 3,20; 2,07–4,93) e por doença arterial coronariana (HR 2,73; 1,55–4,82), prevendo desfechos melhor que HbA1c ou glicemia (Min & Min, CMAJ 2013;185:E402). Confirmado em Diabetes Care 2013;36:708.

**Caveat renal:** o peptídeo C é depurado pelo rim — na doença renal crônica acumula e superestima a secreção. Interpretar sempre junto da função renal (creatinina, cistatina C, RFG).

**Marcadores complementares:** insulina de jejum, glicemia de jejum, HOMA-IR, HbA1c, curva insulinêmica-glicêmica, triglicerídeos/HDL, PCR-us, função renal.$cr$,
$pe$O peptídeo C é uma "peça" que o pâncreas libera junto com a insulina, na mesma quantidade. Como dura mais tempo no sangue e não passa pelo filtro do fígado, ele mede melhor do que a própria insulina o quanto o seu pâncreas está produzindo.

Em jejum, o valor ideal fica entre 1,0 e 2,5. Acima disso costuma indicar que o corpo está fabricando insulina demais (resistência à insulina), o que ao longo dos anos aumenta o risco de diabetes, problemas no coração e fígado gorduroso. Valores muito baixos indicam o contrário: um pâncreas que já produz pouca insulina.

Ajuda a entender se o "motor" do açúcar está trabalhando demais, de menos, ou no ponto certo. Na maioria dos casos, os excessos são reversíveis com alimentação, exercício, sono de qualidade e controle do estresse.$pe$,
$co$## Conduta Clínica — Peptídeo C de Jejum

Interpretar sempre em conjunto com glicemia e insulina de jejum e com a função renal (o peptídeo C acumula quando o RFG cai).

### Estratificação por faixa

| Peptídeo C (jejum) | Interpretação | Conduta |
|---|---|---|
| <0,6 ng/mL | Falência de célula β | Rastrear DM1/LADA (anti-GAD65, anti-IA-2); evitar sulfonilureia; considerar insulinoterapia |
| 0,6–1,0 ng/mL | Reserva β reduzida | Avaliar LADA em adulto; proteção metabólica agressiva |
| 1,0–2,5 ng/mL | Ótimo | Manutenção |
| 2,5–3,5 ng/mL | Hiperinsulinemia limítrofe | Protocolo de resistência insulínica |
| 3,6–5,0 ng/mL | Hiperinsulinemia / RI | Protocolo completo de RI |
| >5,0 ng/mL | Hipersecreção acentuada | Protocolo + investigar insulinoma se houver hipoglicemia |

### Avaliação complementar
- Insulina e glicemia de jejum + HOMA-IR; curva insulinêmica-glicêmica quando RI provável
- HbA1c (ideal <5,5% funcional)
- **Função renal** (creatinina, cistatina C, RFG): obrigatória para interpretar o peptídeo C
- Anti-GAD65 e anti-IA-2 quando peptídeo C baixo com hiperglicemia (rastreio de LADA)
- Perfil lipídico, relação TG/HDL, PCR-us; composição corporal (DEXA) com foco em gordura visceral

### Peptídeo C alto (hiperinsulinemia / RI) — proteger a célula β
- Redução da carga glicêmica (low-carb ou mediterrâneo); jejum intermitente / alimentação com tempo restrito
- Exercício resistido + aeróbico (a contração muscular capta glicose independentemente de insulina)
- Perda de gordura visceral; sono regular; manejo do estresse
- Suporte: magnésio, zinco, vitamina D (alvo 60–80 ng/mL), berberina
- Considerar metformina e/ou agonista de GLP-1 conforme contexto (reduzem glicotoxicidade e a demanda secretória sobre a célula β)

### Peptídeo C baixo (reserva reduzida / falência)
- Confirmar com anticorpos (LADA); **evitar secretagogos** (sulfonilureia) que aceleram a exaustão β
- Considerar insulinoterapia precoce para "repouso" da célula β em DM2 com peptídeo C baixo

### Monitoramento
- Peptídeo C seriado (com insulina/glicemia de jejum, HbA1c e função renal) a cada 3–6 meses
- A tendência é mais informativa que o valor isolado: queda progressiva indica exaustão β; subida pós-intervenção indica recuperação funcional

### Encaminhamento
- Endocrinologista: suspeita de LADA/insulinoma, falência β, falha terapêutica
- Nutricionista: individualização do plano alimentar$co$,
$se$O peptídeo C é uma "peça" que o pâncreas libera junto com a insulina, na mesma quantidade. Como dura mais tempo no sangue e não passa pelo filtro do fígado, mede melhor do que a própria insulina o quanto o seu pâncreas está produzindo. O valor ideal em jejum fica entre 1,0 e 2,5; valores mais altos sugerem produção excessiva de insulina e valores muito baixos, produção insuficiente.$se$,
  now(), now()
WHERE NOT EXISTS (SELECT 1 FROM score_items WHERE id = '0c9e7d10-0000-7000-8000-0000000000c1');

-- 3) Levels (0-5, por qualidade; multiplicador = level/5) ----------------------
INSERT INTO score_levels (id, item_id, level, name, lower_limit, upper_limit, operator, site_legend, created_at, updated_at)
SELECT v.id, '0c9e7d10-0000-7000-8000-0000000000c1', v.level, v.name, v.lo, v.hi, v.op, v.legend, now(), now()
FROM (VALUES
  ('0c9e7d10-0000-7000-8000-0000000000d0'::uuid, 0, '5,0 ou mais',   '5.0', NULL,  '>=',      'Peptídeo C em jejum muito alto: hipersecreção importante de insulina, sinal de alerta.'),
  ('0c9e7d10-0000-7000-8000-0000000000d1'::uuid, 1, '3,6 a 4,9',     '3.6', '4.9', 'between', 'Peptídeo C em jejum elevado: excesso de insulina por resistência insulínica.'),
  ('0c9e7d10-0000-7000-8000-0000000000d2'::uuid, 2, 'menos que 0,6', NULL,  '0.6', '<',       'Peptídeo C em jejum muito baixo: pâncreas produzindo pouca insulina, sinal de alerta.'),
  ('0c9e7d10-0000-7000-8000-0000000000d3'::uuid, 3, '2,6 a 3,5',     '2.6', '3.5', 'between', 'Peptídeo C em jejum no limite superior: começa a sugerir excesso de insulina, atenção.'),
  ('0c9e7d10-0000-7000-8000-0000000000d4'::uuid, 4, '0,6 a 0,9',     '0.6', '0.9', 'between', 'Peptídeo C em jejum no limite inferior do normal: produção de insulina um pouco baixa, acompanhar.'),
  ('0c9e7d10-0000-7000-8000-0000000000d5'::uuid, 5, '1,0 a 2,5',     '1.0', '2.5', 'between', 'Peptídeo C em jejum em nível ótimo: produção de insulina adequada e equilibrada.')
) AS v(id, level, name, lo, hi, op, legend)
WHERE NOT EXISTS (SELECT 1 FROM score_levels WHERE item_id = '0c9e7d10-0000-7000-8000-0000000000c1' AND level = v.level AND deleted_at IS NULL);

-- 4) Pilares (mesmos da insulina de jejum) ------------------------------------
INSERT INTO score_item_method_pillars (score_item_id, method_pillar_id)
SELECT '0c9e7d10-0000-7000-8000-0000000000c1', p.method_pillar_id
FROM score_item_method_pillars p
WHERE p.score_item_id = (SELECT id FROM score_items WHERE name = 'INSULINA 0 MIN (JEJUM)' AND deleted_at IS NULL LIMIT 1)
ON CONFLICT (score_item_id, method_pillar_id) DO NOTHING;

COMMIT;
