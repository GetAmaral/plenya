-- ============================================================================
-- Escore — substitui sufixos numéricos (_2/_3/...) por sufixos COERENTES
-- ============================================================================
-- Colisões de anamnese (mesmo nome em contextos diferentes) eram desambiguadas
-- por número (ridículo). Aqui cada uma vira um sufixo coerente com grupo/subgrupo/
-- pai/gênero. NÃO toca genética (códigos rsID, risco de matching) nem maxscore de
-- escala (ASEX_25, GAD_7_21...). ATIVIDADE_FISICA / EXERCICIO_FISICO viram variantes
-- por etapa de vida; a constante clínica (clinical_codes.go) passa a apontar _VIDA_ADULTA.
-- Idempotente: WHERE ancorado no padrão original; re-rodar não reprocessa.
-- ============================================================================
BEGIN;

-- A1) Doenças FAMILIARES → _FAMILIAR (CV, auto-imunes, virais, HAS) ----------
UPDATE score_items si SET anamnese_item_code = regexp_replace(si.anamnese_item_code,'_[0-9]+$','') || '_FAMILIAR', updated_at = now()
FROM score_subgroups sg JOIN score_groups g ON g.id = sg.group_id
WHERE si.subgroup_id = sg.id AND g.name = 'Histórico Familiar de Doenças' AND si.deleted_at IS NULL
  AND si.anamnese_item_code ~ '^(DOENCA_CARDIOVASCULAR|DOENCAS_AUTO_IMUNES|DOENCAS_VIRAIS_CRONICAS|HIPERTENSAO_ARTERIAL)(_[0-9]+)?$';

-- A2) Doenças PESSOAIS → código limpo (tira o _N que sobrou) ------------------
UPDATE score_items si SET anamnese_item_code = regexp_replace(si.anamnese_item_code,'_[0-9]+$',''), updated_at = now()
FROM score_subgroups sg JOIN score_groups g ON g.id = sg.group_id
WHERE si.subgroup_id = sg.id AND g.name = 'Histórico de doenças' AND si.deleted_at IS NULL
  AND si.anamnese_item_code ~ '^(DOENCAS_AUTO_IMUNES|DOENCAS_VIRAIS_CRONICAS)_[0-9]+$';

-- A3) Álcool (hábitos/vícios) → _HABITOS (álcool da Alimentação fica limpo) ---
UPDATE score_items SET anamnese_item_code = 'ALCOOL_HABITOS', updated_at = now()
WHERE anamnese_item_code = 'ALCOOL_2' AND deleted_at IS NULL;

-- B) Timeline familiar por PAI (Mãe/Pai) -------------------------------------
UPDATE score_items si SET anamnese_item_code = regexp_replace(si.anamnese_item_code,'_[0-9]+$','')
    || '_' || CASE p.name WHEN 'Mãe' THEN 'MAE' WHEN 'Pai' THEN 'PAI' END, updated_at = now()
FROM score_items p
WHERE si.parent_item_id = p.id AND si.deleted_at IS NULL AND p.name IN ('Mãe','Pai')
  AND si.anamnese_item_code ~ '^(DURANTE_ADOLESCENCIA_DO_PACIENTE|DURANTE_INFANCIA_DO_PACIENTE|PRE_NATAL_PRE_CONCEPCAO)(_[0-9]+)?$';

-- C) Etapas de vida por GRUPO (Infância/Adolescência/Vida adulta/Pré-natal) ---
UPDATE score_items si SET anamnese_item_code = regexp_replace(si.anamnese_item_code,'_[0-9]+$','')
    || '_' || CASE g.name WHEN 'Alimentação' THEN 'ALIMENTACAO' WHEN 'Composição corporal' THEN 'COMPOSICAO'
                          WHEN 'Movimento e atividade física' THEN 'MOVIMENTO' WHEN 'Sono' THEN 'SONO' END, updated_at = now()
FROM score_subgroups sg JOIN score_groups g ON g.id = sg.group_id
WHERE si.subgroup_id = sg.id AND si.deleted_at IS NULL
  AND g.name IN ('Alimentação','Composição corporal','Movimento e atividade física','Sono')
  AND si.anamnese_item_code ~ '^(INFANCIA|ADOLESCENCIA|VIDA_ADULTA|PRE_NATAL)(_[0-9]+)?$';

-- D) Movimento histórico (atividade/exercício/ar livre/esporte) por etapa de vida (pai) -
UPDATE score_items si SET anamnese_item_code = regexp_replace(si.anamnese_item_code,'_[0-9]+$','')
    || '_' || CASE p.name WHEN 'Infância' THEN 'INFANCIA' WHEN 'Adolescência' THEN 'ADOLESCENCIA' WHEN 'Vida adulta' THEN 'VIDA_ADULTA' END, updated_at = now()
FROM score_items p
WHERE si.parent_item_id = p.id AND si.deleted_at IS NULL AND p.name IN ('Infância','Adolescência','Vida adulta')
  AND si.anamnese_item_code ~ '^(ATIVIDADE_FISICA|EXERCICIO_FISICO|ATIVIDADES_AO_AR_LIVRE|ESPORTE_COMPETITIVO)(_[0-9]+)?$';

-- E) Testes práticos por GÊNERO + FAIXA ETÁRIA (colunas) ----------------------
UPDATE score_items si SET anamnese_item_code = regexp_replace(si.anamnese_item_code,'_[0-9]+$','')
    || '_' || CASE si.gender WHEN 'male' THEN 'HOMENS' WHEN 'female' THEN 'MULHERES' END
    || COALESCE('_' || si.age_range_min || '_' || si.age_range_max, ''), updated_at = now()
WHERE si.deleted_at IS NULL AND si.gender IN ('male','female')
  AND si.anamnese_item_code ~ '^(RESISTENCIA_NEUROMUSCULAR_ABDOMINAL|RESISTENCIA_NEUROMUSCULAR_FLEXAO_DE_SOLO|RESISTENCIA_NEUROMUSCULAR_PRANCHA|RESISTENCIA_CARDIORRESPIRATORIA_BURPEE)_[0-9]+$';

-- F) Social → _ATUAL / _HISTORICO (por subgrupo) -----------------------------
UPDATE score_items si SET anamnese_item_code = regexp_replace(si.anamnese_item_code,'_[0-9]+$','')
    || '_' || CASE sg.name WHEN 'Atual' THEN 'ATUAL' WHEN 'Histórico' THEN 'HISTORICO' END, updated_at = now()
FROM score_subgroups sg JOIN score_groups g ON g.id = sg.group_id
WHERE si.subgroup_id = sg.id AND g.name = 'Social' AND sg.name IN ('Atual','Histórico') AND si.deleted_at IS NULL
  AND si.anamnese_item_code ~ '^(CONDICOES_DE_MORADIA|PROFISSOES|RELIGIOSIDADE|SITUACAO_CONJUGAL|SITUACAO_DE_PETS|SITUACAO_FAMILIAR|SITUACAO_FINANCEIRA_RENDA_ATIVA_E_PASSIVA)(_[0-9]+)?$';

-- Conferência: nenhum código duplicado fora da genética -----------------------
SELECT si.anamnese_item_code, COUNT(*) AS n
FROM score_items si JOIN score_subgroups sg ON sg.id=si.subgroup_id JOIN score_groups g ON g.id=sg.group_id
WHERE si.deleted_at IS NULL AND si.anamnese_item_code IS NOT NULL AND g.name <> 'Genética'
GROUP BY si.anamnese_item_code HAVING COUNT(*) > 1;

COMMIT;
