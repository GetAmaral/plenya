-- Curadoria de escore (Histórico de doenças e correlatos), consolidando 4 deltas:
--   (1) Sufixos coerentes nos anamnese_item_code (substitui _2/_3 por sufixo de
--       contexto: _FAMILIAR, _MAE/_PAI, etapa de vida, gênero+faixa, _ATUAL/_HISTORICO).
--       NÃO toca genética nem maxscore de escala. ATIVIDADE_FISICA/EXERCICIO_FISICO
--       viram _VIDA_ADULTA (acoplado a clinical_codes.go — por isso vive numa migration
--       que roda junto do deploy do api).
--   (2) "Hipertensão arterial" pessoal em Doenças crônicas (6 níveis, 5 templates).
--   (3) Artrite reumatóide vira filha de DOENCAS_AUTO_IMUNES.
--   (4) Psoríase, Vitiligo e Tireoidite de Hashimoto como novas filhas de
--       DOENCAS_AUTO_IMUNES (6 níveis cada, nos mesmos 5 templates das irmãs).
-- Tudo idempotente (guards por código/padrão original); portável dev/prod (resolve
-- pai/subgrupo por código/nome, nunca por UUID de score_item; templates por UUID fixo).
-- A ordem importa: (1) limpa os códigos antes de (2)(3)(4) os referenciarem.

-- +goose Up

-- ███ (1) SUFIXOS COERENTES ██████████████████████████████████████████████████
-- A1) Doenças FAMILIARES -> _FAMILIAR (CV, auto-imunes, virais, HAS)
UPDATE score_items si SET anamnese_item_code = regexp_replace(si.anamnese_item_code,'_[0-9]+$','') || '_FAMILIAR', updated_at = now()
FROM score_subgroups sg JOIN score_groups g ON g.id = sg.group_id
WHERE si.subgroup_id = sg.id AND g.name = 'Histórico Familiar de Doenças' AND si.deleted_at IS NULL
  AND si.anamnese_item_code ~ '^(DOENCA_CARDIOVASCULAR|DOENCAS_AUTO_IMUNES|DOENCAS_VIRAIS_CRONICAS|HIPERTENSAO_ARTERIAL)(_[0-9]+)?$';

-- A2) Doenças PESSOAIS -> código limpo (tira o _N que sobrou)
UPDATE score_items si SET anamnese_item_code = regexp_replace(si.anamnese_item_code,'_[0-9]+$',''), updated_at = now()
FROM score_subgroups sg JOIN score_groups g ON g.id = sg.group_id
WHERE si.subgroup_id = sg.id AND g.name = 'Histórico de doenças' AND si.deleted_at IS NULL
  AND si.anamnese_item_code ~ '^(DOENCAS_AUTO_IMUNES|DOENCAS_VIRAIS_CRONICAS)_[0-9]+$';

-- A3) Álcool (hábitos/vícios) -> _HABITOS (álcool da Alimentação fica limpo)
UPDATE score_items SET anamnese_item_code = 'ALCOOL_HABITOS', updated_at = now()
WHERE anamnese_item_code = 'ALCOOL_2' AND deleted_at IS NULL;

-- B) Timeline familiar por PAI (Mãe/Pai)
UPDATE score_items si SET anamnese_item_code = regexp_replace(si.anamnese_item_code,'_[0-9]+$','')
    || '_' || CASE p.name WHEN 'Mãe' THEN 'MAE' WHEN 'Pai' THEN 'PAI' END, updated_at = now()
FROM score_items p
WHERE si.parent_item_id = p.id AND si.deleted_at IS NULL AND p.name IN ('Mãe','Pai')
  AND si.anamnese_item_code ~ '^(DURANTE_ADOLESCENCIA_DO_PACIENTE|DURANTE_INFANCIA_DO_PACIENTE|PRE_NATAL_PRE_CONCEPCAO)(_[0-9]+)?$';

-- C) Etapas de vida por GRUPO (Infância/Adolescência/Vida adulta/Pré-natal)
UPDATE score_items si SET anamnese_item_code = regexp_replace(si.anamnese_item_code,'_[0-9]+$','')
    || '_' || CASE g.name WHEN 'Alimentação' THEN 'ALIMENTACAO' WHEN 'Composição corporal' THEN 'COMPOSICAO'
                          WHEN 'Movimento e atividade física' THEN 'MOVIMENTO' WHEN 'Sono' THEN 'SONO' END, updated_at = now()
FROM score_subgroups sg JOIN score_groups g ON g.id = sg.group_id
WHERE si.subgroup_id = sg.id AND si.deleted_at IS NULL
  AND g.name IN ('Alimentação','Composição corporal','Movimento e atividade física','Sono')
  AND si.anamnese_item_code ~ '^(INFANCIA|ADOLESCENCIA|VIDA_ADULTA|PRE_NATAL)(_[0-9]+)?$';

-- D) Movimento histórico (atividade/exercício/ar livre/esporte) por etapa de vida (pai)
UPDATE score_items si SET anamnese_item_code = regexp_replace(si.anamnese_item_code,'_[0-9]+$','')
    || '_' || CASE p.name WHEN 'Infância' THEN 'INFANCIA' WHEN 'Adolescência' THEN 'ADOLESCENCIA' WHEN 'Vida adulta' THEN 'VIDA_ADULTA' END, updated_at = now()
FROM score_items p
WHERE si.parent_item_id = p.id AND si.deleted_at IS NULL AND p.name IN ('Infância','Adolescência','Vida adulta')
  AND si.anamnese_item_code ~ '^(ATIVIDADE_FISICA|EXERCICIO_FISICO|ATIVIDADES_AO_AR_LIVRE|ESPORTE_COMPETITIVO)(_[0-9]+)?$';

-- E) Testes práticos por GÊNERO + FAIXA ETÁRIA (colunas)
UPDATE score_items si SET anamnese_item_code = regexp_replace(si.anamnese_item_code,'_[0-9]+$','')
    || '_' || CASE si.gender WHEN 'male' THEN 'HOMENS' WHEN 'female' THEN 'MULHERES' END
    || COALESCE('_' || si.age_range_min || '_' || si.age_range_max, ''), updated_at = now()
WHERE si.deleted_at IS NULL AND si.gender IN ('male','female')
  AND si.anamnese_item_code ~ '^(RESISTENCIA_NEUROMUSCULAR_ABDOMINAL|RESISTENCIA_NEUROMUSCULAR_FLEXAO_DE_SOLO|RESISTENCIA_NEUROMUSCULAR_PRANCHA|RESISTENCIA_CARDIORRESPIRATORIA_BURPEE)_[0-9]+$';

-- F) Social -> _ATUAL / _HISTORICO (por subgrupo)
UPDATE score_items si SET anamnese_item_code = regexp_replace(si.anamnese_item_code,'_[0-9]+$','')
    || '_' || CASE sg.name WHEN 'Atual' THEN 'ATUAL' WHEN 'Histórico' THEN 'HISTORICO' END, updated_at = now()
FROM score_subgroups sg JOIN score_groups g ON g.id = sg.group_id
WHERE si.subgroup_id = sg.id AND g.name = 'Social' AND sg.name IN ('Atual','Histórico') AND si.deleted_at IS NULL
  AND si.anamnese_item_code ~ '^(CONDICOES_DE_MORADIA|PROFISSOES|RELIGIOSIDADE|SITUACAO_CONJUGAL|SITUACAO_DE_PETS|SITUACAO_FAMILIAR|SITUACAO_FINANCEIRA_RENDA_ATIVA_E_PASSIVA)(_[0-9]+)?$';

-- ███ (2) HIPERTENSÃO ARTERIAL (pessoal) em Histórico de doenças › Doenças crônicas ███
INSERT INTO score_items (id, name, anamnese_item_code, subgroup_id, "order", points, default_level5, gender, created_at, updated_at)
SELECT uuid_generate_v7(), 'Hipertensão arterial', 'HIPERTENSAO_ARTERIAL', sg.id, 7, 25, true, 'not_applicable', now(), now()
FROM score_subgroups sg JOIN score_groups g ON g.id = sg.group_id
WHERE g.name = 'Histórico de doenças' AND sg.name LIKE 'Doenças crônicas%'
  AND NOT EXISTS (SELECT 1 FROM score_items x WHERE x.anamnese_item_code = 'HIPERTENSAO_ARTERIAL' AND x.deleted_at IS NULL);

INSERT INTO score_levels (id, item_id, level, operator, name, site_legend, created_at, updated_at)
SELECT uuid_generate_v7(), si.id, v.level, '=', v.lname, v.legend, now(), now()
FROM score_items si
CROSS JOIN (VALUES
  (0, 'Tenho há mais de 5 anos, com controle inadequado', 'Hipertensão há mais de cinco anos e mal controlada, exige cuidado intensivo.'),
  (1, 'Tenho há menos de 5 anos, com controle inadequado', 'Hipertensão recente com a pressão ainda mal controlada, requer ajuste.'),
  (2, 'Tenho há mais de 5 anos, com controle adequado', 'Hipertensão de longa data, porém bem controlada.'),
  (3, 'Tenho há menos de 5 anos, com controle adequado', 'Hipertensão há menos de cinco anos e bem controlada.'),
  (4, 'Não tenho, mas tenho histórico familiar forte', 'Sem hipertensão, mas com histórico forte na família.'),
  (5, 'Não tenho', 'Não tem hipertensão, situação favorável.')
) AS v(level, lname, legend)
WHERE si.anamnese_item_code = 'HIPERTENSAO_ARTERIAL' AND si.deleted_at IS NULL
  AND NOT EXISTS (SELECT 1 FROM score_levels sl WHERE sl.item_id = si.id AND sl.level = v.level AND sl.deleted_at IS NULL);

INSERT INTO anamnesis_template_items (id, anamnesis_template_id, score_item_id, "order", created_at, updated_at)
SELECT uuid_generate_v7(), t.id, si.id, 0, now(), now()
FROM score_items si
JOIN anamnesis_templates t ON t.id IN (
  '11111111-1111-7111-8111-111111111101', '11111111-1111-7111-8111-111111111102',
  '11111111-1111-7111-8111-111111111104', '11111111-1111-7111-8111-111111111106',
  '11111111-1111-7111-8111-111111111107')
WHERE si.anamnese_item_code = 'HIPERTENSAO_ARTERIAL' AND si.deleted_at IS NULL
  AND NOT EXISTS (
    SELECT 1 FROM anamnesis_template_items x
    WHERE x.anamnesis_template_id = t.id AND x.score_item_id = si.id AND x.deleted_at IS NULL);

-- ███ (3) ARTRITE REUMATÓIDE vira filha de DOENCAS_AUTO_IMUNES ████████████████
UPDATE score_items si
SET parent_item_id = p.id, "order" = 23, updated_at = now()
FROM score_items p
JOIN score_subgroups psg ON psg.id = p.subgroup_id
JOIN score_groups   pg  ON pg.id  = psg.group_id
WHERE p.anamnese_item_code = 'DOENCAS_AUTO_IMUNES' AND pg.name = 'Histórico de doenças' AND p.deleted_at IS NULL
  AND si.anamnese_item_code = 'ARTRITE_REUMATOIDE' AND si.subgroup_id = p.subgroup_id AND si.deleted_at IS NULL;

-- ███ (4) PSORÍASE / VITILIGO / HASHIMOTO como filhas de DOENCAS_AUTO_IMUNES ███
WITH parent AS (
  SELECT p.id AS parent_id, p.subgroup_id
  FROM score_items p
  JOIN score_subgroups psg ON psg.id = p.subgroup_id
  JOIN score_groups   pg  ON pg.id  = psg.group_id
  WHERE p.anamnese_item_code = 'DOENCAS_AUTO_IMUNES' AND pg.name = 'Histórico de doenças' AND p.deleted_at IS NULL
  LIMIT 1
)
INSERT INTO score_items
  (id, name, anamnese_item_code, subgroup_id, parent_item_id, "order", points,
   default_level5, gender, site_render_type, site_question, site_explanation,
   patient_explanation, created_at, updated_at)
SELECT uuid_generate_v7(), v.name, v.code, parent.subgroup_id, parent.parent_id,
       v.ord, v.points, true, 'not_applicable', 'level_choice', v.site_q, v.lay,
       v.lay, now(), now()
FROM parent
CROSS JOIN (VALUES
  ('Psoríase', 'PSORIASE', 31, 14,
   'Você tem psoríase (doença autoimune da pele com placas avermelhadas e descamação)?',
   'A psoríase é uma doença inflamatória crônica de base imunológica em que a pele renova suas células rápido demais, formando placas avermelhadas com descamação esbranquiçada, em geral em cotovelos, joelhos, couro cabeludo e região lombar. Não é contagiosa. Em parte das pessoas vem acompanhada de dor e inchaço nas articulações (artrite psoriásica) e associa-se a maior risco cardiovascular e metabólico, por isso importa olhar além da pele. Com tratamento adequado e hábitos saudáveis, como controle do peso, redução do álcool, parar de fumar, manejo do estresse e sono de qualidade, a maioria das pessoas mantém a doença controlada e boa qualidade de vida. Registrar há quanto tempo tem o diagnóstico, a extensão das lesões e o tratamento em uso ajuda a equipe a cuidar de você de forma completa.'),
  ('Vitiligo', 'VITILIGO', 32, 10,
   'Você tem vitiligo (manchas brancas na pele pela perda de pigmento, de origem autoimune)?',
   'O vitiligo é uma condição autoimune em que o sistema de defesa do corpo ataca as células que produzem o pigmento da pele (melanócitos), criando manchas brancas bem delimitadas, que costumam aparecer nas mãos, no rosto, ao redor de aberturas naturais e em áreas de atrito. Não causa dor nem é contagioso, mas pode ter impacto emocional importante. Costuma andar junto de outras condições autoimunes, principalmente da tireoide, por isso vale investigar. Com tratamento e acompanhamento é possível estabilizar a progressão e, em muitos casos, recuperar parte da cor da pele. Proteger a pele do sol e cuidar do bem estar emocional fazem parte do cuidado. Registrar há quanto tempo tem, se está aumentando e o tratamento em uso ajuda a equipe a orientar melhor.'),
  ('Tireoidite de Hashimoto', 'HASHIMOTO', 33, 14,
   'Você tem tireoidite de Hashimoto (doença autoimune da tireoide)?',
   'A tireoidite de Hashimoto é uma doença autoimune em que o sistema de defesa do corpo ataca a tireoide, a glândula que regula o metabolismo. Com o tempo pode reduzir a produção de hormônios (hipotireoidismo), causando cansaço, ganho de peso, intestino preso, pele seca, queda de cabelo, sensação de frio e desânimo. É a causa mais comum de hipotireoidismo e costuma ser bem controlada com reposição hormonal quando necessária, acompanhada de hábitos saudáveis. Pode vir junto de outras condições autoimunes, então vale ficar atento. Registrar há quanto tempo tem o diagnóstico, se a função da tireoide está controlada e o tratamento em uso é importante para o seu cuidado.')
) AS v(name, code, ord, points, site_q, lay)
WHERE NOT EXISTS (SELECT 1 FROM score_items x WHERE x.anamnese_item_code = v.code AND x.deleted_at IS NULL);

UPDATE score_items SET clinical_relevance =
'A psoríase é uma doença inflamatória crônica imunomediada, com desregulação do eixo IL-23/Th17 e papel central de IL-17, TNF-alfa e IL-22, levando à hiperproliferação de queratinócitos e à inflamação sistêmica persistente. Acomete cerca de 2 a 3 por cento da população. A gravidade é estimada pela extensão (BSA), por índices como o PASI e pelo impacto na qualidade de vida (DLQI). É hoje entendida como doença sistêmica: associa-se a artrite psoriásica (presente em cerca de 20 a 30 por cento dos casos), síndrome metabólica, obesidade, esteatose hepática, doença cardiovascular aterosclerótica, diabetes, doença inflamatória intestinal, depressão e ansiedade. A inflamação crônica (PCR e citocinas elevadas) contribui para o risco cardiometabólico de forma independente dos fatores clássicos. Gatilhos e agravantes incluem infecção estreptocócica (psoríase gutata), tabagismo, álcool, obesidade, estresse psicológico, trauma cutâneo (fenômeno de Koebner) e fármacos (betabloqueadores, lítio, antimaláricos, retirada de corticoide sistêmico). Deficiência de vitamina D é frequente. A avaliação deve estratificar gravidade cutânea, comprometimento articular e carga de comorbidades cardiometabólicas.'
WHERE anamnese_item_code = 'PSORIASE' AND deleted_at IS NULL;

UPDATE score_items SET conduct =
'## Conduta Clínica: Psoríase

### 1. Avaliação
- Tempo de doença, extensão (BSA) e gravidade; impacto na qualidade de vida (DLQI).
- Rastreio de artrite psoriásica: dor e rigidez articular inflamatória, dactilite, entesite, dor lombar inflamatória.
- Comorbidades cardiometabólicas e saúde mental (depressão, ansiedade).
- Gatilhos: infecções, tabagismo, álcool, peso, estresse, fármacos.

### 2. Investigação
- PCR e VHS (carga inflamatória).
- Perfil lipídico, glicemia e insulina de jejum, HbA1c.
- Função hepática (esteatose), ácido úrico.
- Vitamina D; demais conforme contexto.

### 3. Encaminhamentos
- Dermatologia para todos os casos.
- Reumatologia se suspeita de artrite psoriásica.
- Cardiologia ou endocrinologia conforme risco cardiometabólico.

### 4. Intervenções
- Tópicos, fototerapia NB-UVB, sistêmicos e biológicos (anti-TNF, anti-IL-17, anti-IL-23) conforme dermatologia.
- Estilo de vida anti-inflamatório, controle de peso, cessação de tabagismo e álcool.
- Vitamina D, ômega-3, manejo de estresse e sono de qualidade.

### 5. Monitoramento
- Reavaliar gravidade cutânea, articular e perfil cardiometabólico periodicamente.'
WHERE anamnese_item_code = 'PSORIASE' AND deleted_at IS NULL;

UPDATE score_items SET clinical_relevance =
'O vitiligo é uma doença autoimune órgão específica caracterizada pela destruição de melanócitos mediada por linfócitos T CD8 citotóxicos, com papel central do eixo IFN-gama e da quimiocina CXCL10, resultando em máculas acrômicas bem delimitadas. Prevalência aproximada de 0,5 a 2 por cento. Apresenta forte associação com outras doenças autoimunes, em especial a tireoidite autoimune (Hashimoto e Graves), além de diabetes tipo 1, anemia perniciosa, doença de Addison, alopecia areata e artrite reumatoide, configurando com frequência quadro de poliautoimunidade. A atividade (estável ou em progressão) e a extensão corporal orientam conduta e prognóstico. Gatilhos incluem estresse oxidativo, trauma cutâneo (fenômeno de Koebner), estresse psicológico e exposição a fenóis. O impacto na qualidade de vida e o sofrimento psicológico são relevantes e muitas vezes subestimados. A avaliação deve rastrear ativamente disfunção tireoidiana e outras autoimunes associadas, mesmo na ausência de sintomas.'
WHERE anamnese_item_code = 'VITILIGO' AND deleted_at IS NULL;

UPDATE score_items SET conduct =
'## Conduta Clínica: Vitiligo

### 1. Avaliação
- Extensão (BSA) e atividade (estável ou em progressão); fenômeno de Koebner.
- Impacto psicológico e na qualidade de vida.
- História pessoal e familiar de autoimunidade.

### 2. Investigação
- TSH, T4 livre e anti-TPO (rastreio de tireoidopatia autoimune).
- Glicemia e HbA1c; hemograma e vitamina B12 (anemia perniciosa).
- Vitamina D; rastreio de outras autoimunes conforme clínica.

### 3. Encaminhamentos
- Dermatologia.
- Endocrinologia se disfunção tireoidiana.
- Psicologia pelo impacto na qualidade de vida.

### 4. Intervenções
- Corticoide e inibidores de calcineurina tópicos; inibidores de JAK tópicos.
- Fototerapia NB-UVB nas formas extensas ou em progressão.
- Fotoproteção rigorosa das áreas despigmentadas e suporte psicológico.
- Abordagem antioxidante e anti-inflamatória; correção de deficiências.

### 5. Monitoramento
- Reavaliar atividade e função tireoidiana periodicamente.'
WHERE anamnese_item_code = 'VITILIGO' AND deleted_at IS NULL;

UPDATE score_items SET clinical_relevance =
'A tireoidite de Hashimoto (tireoidite linfocítica crônica) é a causa mais comum de hipotireoidismo em regiões com suficiência de iodo. Caracteriza-se por infiltrado linfocítico da tireoide e por autoanticorpos anti-tireoperoxidase (anti-TPO) e anti-tireoglobulina (anti-Tg), com destruição progressiva do parênquima glandular. Predomina no sexo feminino e tem forte componente genético. Associa-se a outras doenças autoimunes (doença celíaca, vitiligo, artrite reumatoide, diabetes tipo 1, anemia perniciosa, doença de Addison), compondo síndromes poliglandulares autoimunes. O espectro funcional varia de eutireoidismo com anticorpos positivos a hipotireoidismo subclínico e franco. A disfunção tireoidiana repercute em perfil lipídico, peso, humor, fertilidade e risco cardiovascular. Deficiências de vitamina D, selênio, zinco e ferro são frequentes e relevantes; a suplementação de selênio reduz títulos de anti-TPO em parte dos pacientes. Há associação com doença celíaca e relato de benefício da retirada de glúten em subgrupos. O monitoramento baseia-se em TSH, T4 livre e anticorpos, com ultrassonografia quando há nódulo ou bócio.'
WHERE anamnese_item_code = 'HASHIMOTO' AND deleted_at IS NULL;

UPDATE score_items SET conduct =
'## Conduta Clínica: Tireoidite de Hashimoto

### 1. Avaliação
- Sintomas de hipotireoidismo (ou hipertireoidismo transitório inicial), bócio, história familiar.
- Rastreio de doenças autoimunes associadas.

### 2. Investigação
- TSH, T4 livre, T3, anti-TPO e anti-Tg.
- Perfil lipídico, glicemia e insulina, hemograma, ferritina.
- Vitamina D, vitamina B12, selênio e zinco.
- Rastreio de doença celíaca; ultrassonografia de tireoide se nódulo ou bócio.

### 3. Encaminhamentos
- Endocrinologia para diagnóstico e ajuste da reposição.
- Nutrição.

### 4. Intervenções
- Levotiroxina conforme TSH e clínica, com meta individualizada.
- Reposição de selênio quando indicado; correção de vitamina D, B12, ferro e zinco.
- Considerar exclusão de glúten em casos selecionados.
- Estilo de vida anti-inflamatório, sono e manejo de estresse.

### 5. Monitoramento
- TSH e T4 livre periódicos; anti-TPO conforme contexto.'
WHERE anamnese_item_code = 'HASHIMOTO' AND deleted_at IS NULL;

INSERT INTO score_levels (id, item_id, level, operator, name, site_legend, created_at, updated_at)
SELECT uuid_generate_v7(), si.id, v.level, '=', v.lname, v.legend, now(), now()
FROM score_items si
JOIN (VALUES
  ('PSORIASE', 0, 'Moderada a grave, sem tratamento',                'Psoríase moderada a grave sem tratamento, exige cuidado.'),
  ('PSORIASE', 1, 'Moderada a grave, de difícil controle',          'Psoríase ativa e difícil de controlar, requer acompanhamento próximo.'),
  ('PSORIASE', 2, 'Moderada a grave, controlada com tratamento',    'Psoríase de maior extensão, porém controlada com tratamento.'),
  ('PSORIASE', 3, 'Leve, controlada com tratamento tópico',         'Psoríase leve, controlada apenas com tratamento na pele.'),
  ('PSORIASE', 4, 'Não tenho, mas tenho histórico familiar forte',  'Sem psoríase, mas com histórico familiar forte, o que pede atenção.'),
  ('PSORIASE', 5, 'Não tenho',                                      'Sem psoríase, situação tranquila.'),
  ('VITILIGO', 0, 'Extenso e em progressão, sem tratamento',        'Vitiligo extenso e avançando sem tratamento, requer cuidado.'),
  ('VITILIGO', 1, 'Em progressão, de difícil controle',            'Vitiligo aumentando e difícil de estabilizar, requer acompanhamento.'),
  ('VITILIGO', 2, 'Extenso, porém estável ou controlado',          'Vitiligo de maior extensão, porém estável com acompanhamento.'),
  ('VITILIGO', 3, 'Localizado e estável',                          'Vitiligo em poucas áreas e estável.'),
  ('VITILIGO', 4, 'Não tenho, mas tenho histórico familiar forte',  'Sem vitiligo, mas com histórico familiar forte, o que pede atenção.'),
  ('VITILIGO', 5, 'Não tenho',                                      'Sem vitiligo, situação tranquila.'),
  ('HASHIMOTO', 0, 'Hipotireoidismo sem tratamento',               'Tireoide pouco ativa e sem tratamento, exige cuidado.'),
  ('HASHIMOTO', 1, 'Em tratamento, com controle inadequado',       'Hashimoto em tratamento, mas com a função da tireoide ainda fora da meta.'),
  ('HASHIMOTO', 2, 'Em tratamento, bem controlada',                'Hashimoto com a função da tireoide bem controlada pelo tratamento.'),
  ('HASHIMOTO', 3, 'Anticorpos positivos, sem alteração da função', 'Hashimoto sem afetar a função da tireoide, em observação.'),
  ('HASHIMOTO', 4, 'Não tenho, mas tenho histórico familiar forte', 'Sem Hashimoto, mas com histórico familiar forte, o que pede atenção.'),
  ('HASHIMOTO', 5, 'Não tenho',                                     'Sem doença da tireoide, situação tranquila.')
) AS v(code, level, lname, legend) ON v.code = si.anamnese_item_code
WHERE si.deleted_at IS NULL
  AND NOT EXISTS (SELECT 1 FROM score_levels sl WHERE sl.item_id = si.id AND sl.level = v.level AND sl.deleted_at IS NULL);

INSERT INTO anamnesis_template_items (id, anamnesis_template_id, score_item_id, "order", created_at, updated_at)
SELECT uuid_generate_v7(), t.id, si.id, 0, now(), now()
FROM score_items si
JOIN anamnesis_templates t ON t.id IN (
  '11111111-1111-7111-8111-111111111101', '11111111-1111-7111-8111-111111111102',
  '11111111-1111-7111-8111-111111111104', '11111111-1111-7111-8111-111111111106',
  '11111111-1111-7111-8111-111111111107')
WHERE si.anamnese_item_code IN ('PSORIASE','VITILIGO','HASHIMOTO') AND si.deleted_at IS NULL
  AND NOT EXISTS (
    SELECT 1 FROM anamnesis_template_items x
    WHERE x.anamnesis_template_id = t.id AND x.score_item_id = si.id AND x.deleted_at IS NULL);

-- Renumera "order" (ordem natural g/sg/i) nos 5 templates tocados por (2) e (4)
WITH o AS (
  SELECT ti.id, ROW_NUMBER() OVER (PARTITION BY ti.anamnesis_template_id
                                   ORDER BY g."order", sg."order", si."order") AS rn
  FROM anamnesis_template_items ti
  JOIN score_items si ON si.id = ti.score_item_id
  JOIN score_subgroups sg ON sg.id = si.subgroup_id
  JOIN score_groups g ON g.id = sg.group_id
  WHERE ti.anamnesis_template_id IN (
    '11111111-1111-7111-8111-111111111101','11111111-1111-7111-8111-111111111102',
    '11111111-1111-7111-8111-111111111104','11111111-1111-7111-8111-111111111106',
    '11111111-1111-7111-8111-111111111107') AND ti.deleted_at IS NULL
)
UPDATE anamnesis_template_items t SET "order" = o.rn, updated_at = now() FROM o WHERE t.id = o.id;

-- +goose Down
-- Remove o que é aditivo (HAS pessoal + as 3 novas autoimunes) e reverte o pai da Artrite.
-- Os RENAMES de sufixo (bloco 1) NÃO são revertidos automaticamente: os códigos numéricos
-- originais não foram snapshotados aqui e clinical_codes.go também precisaria reverter.
-- Para rollback completo, restaurar do backup pré-deploy.
DELETE FROM anamnesis_template_items WHERE score_item_id IN (
  SELECT id FROM score_items WHERE anamnese_item_code IN ('HIPERTENSAO_ARTERIAL','PSORIASE','VITILIGO','HASHIMOTO'));
DELETE FROM score_levels WHERE item_id IN (
  SELECT id FROM score_items WHERE anamnese_item_code IN ('HIPERTENSAO_ARTERIAL','PSORIASE','VITILIGO','HASHIMOTO'));
DELETE FROM score_items WHERE anamnese_item_code IN ('HIPERTENSAO_ARTERIAL','PSORIASE','VITILIGO','HASHIMOTO');
UPDATE score_items SET parent_item_id = NULL, updated_at = now() WHERE anamnese_item_code = 'ARTRITE_REUMATOIDE';
