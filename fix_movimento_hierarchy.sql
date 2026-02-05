-- ============================================================================
-- Script de Correção: Hierarquia do Grupo "Movimento e atividade física"
-- Data: 2026-02-02
-- Autor: Claude Code
--
-- OBJETIVO:
--   - Adicionar items faltantes: Adolescência e Vida adulta (+ filhos)
--   - Corrigir hierarquia dos items órfãos de Infância
--   - NÃO remover ou alterar dados existentes
--
-- BACKUP OBRIGATÓRIO ANTES DE EXECUTAR:
--   pg_dump -U plenya_user plenya_db > backup_$(date +%Y%m%d_%H%M%S).sql
--
-- EXECUÇÃO:
--   psql -U plenya_user -d plenya_db -f fix_movimento_hierarchy.sql
-- ============================================================================

BEGIN;

-- ============================================================================
-- VALIDAÇÕES PRÉ-EXECUÇÃO
-- ============================================================================

DO $$
DECLARE
  v_grupo_id UUID;
  v_subgrupo_historico_id UUID;
  v_infancia_id UUID;
  v_count_items INTEGER;
BEGIN
  -- Verificar se grupo "Movimento e atividade física" existe
  SELECT id INTO v_grupo_id
  FROM score_groups
  WHERE name = 'Movimento e atividade física'
  AND deleted_at IS NULL;

  IF v_grupo_id IS NULL THEN
    RAISE EXCEPTION 'Grupo "Movimento e atividade física" não encontrado!';
  END IF;

  -- Verificar subgrupo "Histórico"
  SELECT id INTO v_subgrupo_historico_id
  FROM score_subgroups
  WHERE name = 'Histórico'
  AND group_id = v_grupo_id
  AND deleted_at IS NULL;

  IF v_subgrupo_historico_id IS NULL THEN
    RAISE EXCEPTION 'Subgrupo "Histórico" não encontrado!';
  END IF;

  -- Verificar item "Infância"
  SELECT id INTO v_infancia_id
  FROM score_items
  WHERE name = 'Infância'
  AND subgroup_id = v_subgrupo_historico_id
  AND deleted_at IS NULL;

  IF v_infancia_id IS NULL THEN
    RAISE EXCEPTION 'Item "Infância" não encontrado!';
  END IF;

  -- Verificar que Adolescência e Vida adulta NÃO existem ainda
  SELECT COUNT(*) INTO v_count_items
  FROM score_items si
  JOIN score_subgroups sg ON si.subgroup_id = sg.id
  WHERE sg.group_id = v_grupo_id
  AND si.name IN ('Adolescência', 'Vida adulta')
  AND si.deleted_at IS NULL;

  IF v_count_items > 0 THEN
    RAISE EXCEPTION 'Items "Adolescência" ou "Vida adulta" JÁ EXISTEM no grupo Movimento!';
  END IF;

  RAISE NOTICE '✅ Validações pré-execução OK';
END $$;

-- ============================================================================
-- PARTE 1: CORRIGIR HIERARQUIA DOS ITEMS ÓRFÃOS DE INFÂNCIA
-- ============================================================================

-- Buscar ID de Infância e do subgrupo Histórico
DO $$
DECLARE
  v_infancia_id UUID;
  v_subgrupo_historico_id UUID;
  v_updated_count INTEGER;
BEGIN
  -- Buscar subgrupo Histórico no grupo Movimento
  SELECT sg.id INTO v_subgrupo_historico_id
  FROM score_subgroups sg
  JOIN score_groups g ON sg.group_id = g.id
  WHERE g.name = 'Movimento e atividade física'
  AND sg.name = 'Histórico'
  AND sg.deleted_at IS NULL
  AND g.deleted_at IS NULL;

  -- Buscar item Infância
  SELECT id INTO v_infancia_id
  FROM score_items
  WHERE name = 'Infância'
  AND subgroup_id = v_subgrupo_historico_id
  AND deleted_at IS NULL;

  -- Corrigir parent_item_id dos filhos de Infância
  UPDATE score_items
  SET parent_item_id = v_infancia_id, updated_at = NOW()
  WHERE subgroup_id = v_subgrupo_historico_id
  AND name IN (
    'Atividades ao ar livre',
    'Esportes praticados (frequência e intensidade)'
  )
  AND parent_item_id IS NULL
  AND deleted_at IS NULL;

  GET DIAGNOSTICS v_updated_count = ROW_COUNT;
  RAISE NOTICE '✅ Parte 1: % items órfãos de Infância corrigidos', v_updated_count;
END $$;

-- ============================================================================
-- PARTE 2: CRIAR ITEM "ADOLESCÊNCIA" E SEUS FILHOS
-- ============================================================================

DO $$
DECLARE
  v_subgrupo_historico_id UUID;
  v_adolescencia_id UUID;
  v_esportes_adolescencia_id UUID;
  v_atividades_adolescencia_id UUID;
  v_next_order INTEGER;
BEGIN
  -- Buscar subgrupo Histórico
  SELECT sg.id INTO v_subgrupo_historico_id
  FROM score_subgroups sg
  JOIN score_groups g ON sg.group_id = g.id
  WHERE g.name = 'Movimento e atividade física'
  AND sg.name = 'Histórico'
  AND sg.deleted_at IS NULL
  AND g.deleted_at IS NULL;

  -- Buscar próximo order disponível
  SELECT COALESCE(MAX("order"), 0) + 1 INTO v_next_order
  FROM score_items
  WHERE subgroup_id = v_subgrupo_historico_id;

  -- 1. Criar item "Adolescência" (pai)
  INSERT INTO score_items (
    id, name, points, "order", subgroup_id, parent_item_id, created_at, updated_at
  ) VALUES (
    uuid_generate_v7(),
    'Adolescência',
    0,
    v_next_order,
    v_subgrupo_historico_id,
    NULL,
    NOW(),
    NOW()
  ) RETURNING id INTO v_adolescencia_id;

  RAISE NOTICE '✅ Item "Adolescência" criado: %', v_adolescencia_id;

  -- 2. Criar filho: Esportes praticados (frequência e intensidade)
  v_next_order := v_next_order + 1;
  INSERT INTO score_items (
    id, name, points, "order", subgroup_id, parent_item_id, created_at, updated_at
  ) VALUES (
    uuid_generate_v7(),
    'Esportes praticados (frequência e intensidade)',
    5,
    v_next_order,
    v_subgrupo_historico_id,
    v_adolescencia_id,
    NOW(),
    NOW()
  ) RETURNING id INTO v_esportes_adolescencia_id;

  RAISE NOTICE '✅ Filho "Esportes praticados" (Adolescência) criado: %', v_esportes_adolescencia_id;

  -- 3. Criar níveis para Esportes praticados - Adolescência
  INSERT INTO score_levels (level, name, operator, lower_limit, upper_limit, item_id) VALUES
    (0, 'Nenhum esporte na adolescência', '=', NULL, NULL, v_esportes_adolescencia_id),
    (1, 'Atividades irregulares por menos de 5 anos', '=', NULL, NULL, v_esportes_adolescencia_id),
    (2, '1-2h de atividade por semana por 5+ anos', '=', NULL, NULL, v_esportes_adolescencia_id),
    (3, '2-3h de atividade por semana por 5+ anos', '=', NULL, NULL, v_esportes_adolescencia_id),
    (4, '3-5h deatividade por semana por 5+ anos', '=', NULL, NULL, v_esportes_adolescencia_id),
    (5, 'Mais de 5h de atividades por semana por 5+ anos', '=', NULL, NULL, v_esportes_adolescencia_id);

  RAISE NOTICE '✅ 6 níveis criados para Esportes praticados (Adolescência)';

  -- 4. Criar filho: Atividades ao ar livre
  v_next_order := v_next_order + 1;
  INSERT INTO score_items (
    id, name, points, "order", subgroup_id, parent_item_id, created_at, updated_at
  ) VALUES (
    uuid_generate_v7(),
    'Atividades ao ar livre',
    0,
    v_next_order,
    v_subgrupo_historico_id,
    v_adolescencia_id,
    NOW(),
    NOW()
  ) RETURNING id INTO v_atividades_adolescencia_id;

  RAISE NOTICE '✅ Filho "Atividades ao ar livre" (Adolescência) criado: %', v_atividades_adolescencia_id;

  RAISE NOTICE '✅ Parte 2: Item "Adolescência" e filhos criados com sucesso';
END $$;

-- ============================================================================
-- PARTE 3: CRIAR ITEM "VIDA ADULTA" E SEUS FILHOS
-- ============================================================================

DO $$
DECLARE
  v_subgrupo_historico_id UUID;
  v_vida_adulta_id UUID;
  v_esportes_vida_adulta_id UUID;
  v_atividades_vida_adulta_id UUID;
  v_next_order INTEGER;
BEGIN
  -- Buscar subgrupo Histórico
  SELECT sg.id INTO v_subgrupo_historico_id
  FROM score_subgroups sg
  JOIN score_groups g ON sg.group_id = g.id
  WHERE g.name = 'Movimento e atividade física'
  AND sg.name = 'Histórico'
  AND sg.deleted_at IS NULL
  AND g.deleted_at IS NULL;

  -- Buscar próximo order disponível
  SELECT COALESCE(MAX("order"), 0) + 1 INTO v_next_order
  FROM score_items
  WHERE subgroup_id = v_subgrupo_historico_id;

  -- 1. Criar item "Vida adulta" (pai)
  INSERT INTO score_items (
    id, name, points, "order", subgroup_id, parent_item_id, created_at, updated_at
  ) VALUES (
    uuid_generate_v7(),
    'Vida adulta',
    0,
    v_next_order,
    v_subgrupo_historico_id,
    NULL,
    NOW(),
    NOW()
  ) RETURNING id INTO v_vida_adulta_id;

  RAISE NOTICE '✅ Item "Vida adulta" criado: %', v_vida_adulta_id;

  -- 2. Criar filho: Esportes praticados (frequência e intensidade)
  v_next_order := v_next_order + 1;
  INSERT INTO score_items (
    id, name, points, "order", subgroup_id, parent_item_id, created_at, updated_at
  ) VALUES (
    uuid_generate_v7(),
    'Esportes praticados (frequência e intensidade)',
    5,
    v_next_order,
    v_subgrupo_historico_id,
    v_vida_adulta_id,
    NOW(),
    NOW()
  ) RETURNING id INTO v_esportes_vida_adulta_id;

  RAISE NOTICE '✅ Filho "Esportes praticados" (Vida adulta) criado: %', v_esportes_vida_adulta_id;

  -- 3. Criar níveis para Esportes praticados - Vida adulta
  INSERT INTO score_levels (level, name, operator, lower_limit, upper_limit, item_id) VALUES
    (0, 'Nenhum esporte na vida adulta', '=', NULL, NULL, v_esportes_vida_adulta_id),
    (1, 'Atividades irregulares por menos de 5 anos', '=', NULL, NULL, v_esportes_vida_adulta_id),
    (2, '1-2h de atividade por semana por 5+ anos', '=', NULL, NULL, v_esportes_vida_adulta_id),
    (3, '2-3h de atividade por semana por 10+ anos', '=', NULL, NULL, v_esportes_vida_adulta_id),
    (4, '3-5h deatividade por semana por 10+ anos', '=', NULL, NULL, v_esportes_vida_adulta_id),
    (5, 'Mais de 5h de atividades por semana por 10+ anos', '=', NULL, NULL, v_esportes_vida_adulta_id);

  RAISE NOTICE '✅ 6 níveis criados para Esportes praticados (Vida adulta)';

  -- 4. Criar filho: Atividades ao ar livre
  v_next_order := v_next_order + 1;
  INSERT INTO score_items (
    id, name, points, "order", subgroup_id, parent_item_id, created_at, updated_at
  ) VALUES (
    uuid_generate_v7(),
    'Atividades ao ar livre',
    0,
    v_next_order,
    v_subgrupo_historico_id,
    v_vida_adulta_id,
    NOW(),
    NOW()
  ) RETURNING id INTO v_atividades_vida_adulta_id;

  RAISE NOTICE '✅ Filho "Atividades ao ar livre" (Vida adulta) criado: %', v_atividades_vida_adulta_id;

  RAISE NOTICE '✅ Parte 3: Item "Vida adulta" e filhos criados com sucesso';
END $$;

-- ============================================================================
-- VALIDAÇÕES PÓS-EXECUÇÃO
-- ============================================================================

DO $$
DECLARE
  v_grupo_id UUID;
  v_total_items INTEGER;
  v_items_com_parent INTEGER;
  v_items_adolescencia INTEGER;
  v_items_vida_adulta INTEGER;
  v_levels_esportes INTEGER;
BEGIN
  -- Buscar grupo
  SELECT id INTO v_grupo_id
  FROM score_groups
  WHERE name = 'Movimento e atividade física'
  AND deleted_at IS NULL;

  -- Contar total de items no grupo
  SELECT COUNT(*) INTO v_total_items
  FROM score_items si
  JOIN score_subgroups sg ON si.subgroup_id = sg.id
  WHERE sg.group_id = v_grupo_id
  AND si.deleted_at IS NULL;

  -- Contar items com parent
  SELECT COUNT(*) INTO v_items_com_parent
  FROM score_items si
  JOIN score_subgroups sg ON si.subgroup_id = sg.id
  WHERE sg.group_id = v_grupo_id
  AND si.parent_item_id IS NOT NULL
  AND si.deleted_at IS NULL;

  -- Verificar Adolescência existe
  SELECT COUNT(*) INTO v_items_adolescencia
  FROM score_items si
  JOIN score_subgroups sg ON si.subgroup_id = sg.id
  WHERE sg.group_id = v_grupo_id
  AND si.name = 'Adolescência'
  AND si.deleted_at IS NULL;

  -- Verificar Vida adulta existe
  SELECT COUNT(*) INTO v_items_vida_adulta
  FROM score_items si
  JOIN score_subgroups sg ON si.subgroup_id = sg.id
  WHERE sg.group_id = v_grupo_id
  AND si.name = 'Vida adulta'
  AND si.deleted_at IS NULL;

  -- Verificar níveis de Esportes praticados
  SELECT COUNT(*) INTO v_levels_esportes
  FROM score_levels sl
  JOIN score_items si ON sl.item_id = si.id
  JOIN score_subgroups sg ON si.subgroup_id = sg.id
  WHERE sg.group_id = v_grupo_id
  AND si.name = 'Esportes praticados (frequência e intensidade)'
  AND si.points = 5;

  RAISE NOTICE '';
  RAISE NOTICE '=== VALIDAÇÕES PÓS-EXECUÇÃO ===';
  RAISE NOTICE 'Total de items no grupo Movimento: % (esperado: 25)', v_total_items;
  RAISE NOTICE 'Items com parent_item_id: % (esperado: 6)', v_items_com_parent;
  RAISE NOTICE 'Item "Adolescência" existe: % (esperado: 1)', v_items_adolescencia;
  RAISE NOTICE 'Item "Vida adulta" existe: % (esperado: 1)', v_items_vida_adulta;
  RAISE NOTICE 'Níveis de Esportes praticados: % (esperado: 18 = 3×6)', v_levels_esportes;

  -- Validar valores esperados
  IF v_total_items <> 25 THEN
    RAISE EXCEPTION '❌ Total de items incorreto: % (esperado: 25)', v_total_items;
  END IF;

  IF v_items_com_parent <> 6 THEN
    RAISE EXCEPTION '❌ Items com parent incorreto: % (esperado: 6)', v_items_com_parent;
  END IF;

  IF v_items_adolescencia <> 1 THEN
    RAISE EXCEPTION '❌ Item Adolescência não encontrado!';
  END IF;

  IF v_items_vida_adulta <> 1 THEN
    RAISE EXCEPTION '❌ Item Vida adulta não encontrado!';
  END IF;

  IF v_levels_esportes <> 18 THEN
    RAISE EXCEPTION '❌ Níveis de Esportes praticados incorreto: % (esperado: 18)', v_levels_esportes;
  END IF;

  RAISE NOTICE '';
  RAISE NOTICE '✅ TODAS AS VALIDAÇÕES PASSARAM!';
  RAISE NOTICE '';
END $$;

-- ============================================================================
-- RESUMO FINAL
-- ============================================================================

SELECT
  'Correção concluída com sucesso!' as status,
  (SELECT COUNT(*) FROM score_items si
   JOIN score_subgroups sg ON si.subgroup_id = sg.id
   JOIN score_groups g ON sg.group_id = g.id
   WHERE g.name = 'Movimento e atividade física'
   AND si.deleted_at IS NULL) as total_items_movimento,
  (SELECT COUNT(*) FROM score_items si
   JOIN score_subgroups sg ON si.subgroup_id = sg.id
   JOIN score_groups g ON sg.group_id = g.id
   WHERE g.name = 'Movimento e atividade física'
   AND si.parent_item_id IS NOT NULL
   AND si.deleted_at IS NULL) as items_com_parent;

-- COMMIT MANUAL NECESSÁRIO!
-- Revisar output acima e executar: COMMIT; ou ROLLBACK;

-- Para aplicar as mudanças, execute no psql:
-- COMMIT;

-- Para reverter (caso algo dê errado), execute:
-- ROLLBACK;
