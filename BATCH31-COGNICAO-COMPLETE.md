# BATCH 31 - COGNIÇÃO - IMPLEMENTAÇÃO COMPLETA

## Status Geral
- **Grupo:** Cognição
- **Total de Items:** 61
- **Progresso Inicial:** 20/81 (24.7%)
- **Meta Final:** 81/81 (100%)
- **Items a Completar:** 61 items

## Estrutura de Implementação

### Fase 1: Testes Neuropsicológicos (12 items) ✅
**Arquivo:** `scripts/batch31_phase1_testes.sql`

**Items:**
- Dubois Imediato (3 items)
- Dubois Tardio (3 items)
- Span de Dígitos Direto (3 items)
- Span de Dígitos Inverso (3 items)

**Temas:**
- Memória episódica verbal
- Consolidação de memória de longo prazo
- Memória de trabalho e funções executivas
- Atenção sustentada

### Fase 2: Capacidades Cognitivas Subjetivas (18 items)
**Arquivo:** `scripts/batch31_phase2_capacidades.sql`

**Items:**
- Capacidade da memória percebida (3 items)
- Capacidade de foco/concentração/aprendizado (3 items)
- Disposição percebida subjetiva (3 items)
- Disposição/energia (3 items)
- Testes rápidos de memória (3 items)
- Uso atual de psicotrópicos para cognição (3 items)

**Temas:**
- Autopercepção cognitiva
- Queixas subjetivas de memória
- Energia mental e fadiga cognitiva
- Medicações e suplementos atuais

### Fase 3: Escalas Validadas (12 items)
**Arquivo:** `scripts/batch31_phase3_escalas.sql`

**Items:**
- Escala PHQ-9 (humor) (3 items)
- GAD-7 (ansiedade) (3 items)
- Escala de Sonolência de Epworth (3 items)
- Escala de Severidade de Fadiga (FSS) (3 items)

**Temas:**
- Sintomas depressivos
- Sintomas ansiosos
- Sonolência diurna excessiva
- Fadiga patológica

### Fase 4: Fatores Hormonais e Afetivos (6 items)
**Arquivo:** `scripts/batch31_phase4_hormonal.sql`

**Items:**
- Mulheres: percepção mudanças ciclo menstrual/menopausa (3 items)
- Presença de sintomas tristeza/depressão/ansiedade (3 items)

**Temas:**
- Neuroendocrinologia feminina
- Impacto hormonal na cognição
- Interface humor-cognição

### Fase 5: História Cognitiva Longitudinal (13 items)
**Arquivo:** `scripts/batch31_phase5_historico.sql`

**Items:**
- Capacidade da memória ao longo da vida (1 item)
- Desempenho escolar (1 item)
- Desenvolvimento neuropsicomotor infância (1 item)
- Disposição/energia ao longo da vida (1 item)
- Foco/concentração/aprendizado ao longo da vida (1 item)
- Histórico uso medicamentos/suplementos cognição (2 items)
- Sintomas ligados TEA ou TDAH (1 item)
- Situações de tristeza/depressão ao longo da vida (1 item)
- Socialização infância/adolescência (1 item)
- Épocas de melhor desempenho cognitivo (1 item)
- Épocas de pior desempenho cognitivo (3 items)

**Temas:**
- Trajetória cognitiva ao longo da vida
- Neurodesenvolvimento
- Fatores de risco e proteção históricos

## Abordagem de Medicina Funcional Integrativa

### Princípios Aplicados
1. **Psiquiatria Metabólica:** Conexão entre metabolismo e função cerebral
2. **Ritmo Circadiano:** Sincronização sono-vigília e consolidação de memória
3. **Cardiologia Funcional:** Perfusão cerebral e saúde vascular cognitiva
4. **Neuroinflamação:** Impacto sistêmico no sistema nervoso central
5. **Neuroplasticidade:** Capacidade de regeneração e adaptação cerebral

### Biomarcadores Priorizados
- Vitaminas B (B12, folato, B6, tiamina)
- Homocisteína
- Vitamina D
- Perfil tireoidiano completo
- Glicemia, insulina, HbA1c
- Ferritina e ferro
- Magnésio e zinco
- Ômega-3 (EPA/DHA)
- Cortisol salivar
- Marcadores inflamatórios (hs-CRP)

### Intervenções Multimodais
1. **Nutricionais:** Dieta MIND, estabilização glicêmica, jejum intermitente
2. **Suplementação:** Nootrópicos naturais, cofatores de neurotransmissores
3. **Sono:** Otimização das fases profundas (N3) e REM
4. **Exercício:** Aeróbico, HIIT, resistido, coordenação
5. **Cognitivas:** Treinamento de memória, aprendizado contínuo
6. **Mente-Corpo:** Mindfulness, manejo de estresse, HRV training

## Evidências Científicas Utilizadas
- Meta-análises 2023-2026
- Estudos de neuroimagem funcional (fMRI, PET)
- Protocolos de reversão cognitiva (ReCODE)
- Intervenções nutricionais validadas
- Pesquisas de neuroplasticidade

## Como Executar

### Método 1: SQL Direto (Recomendado)
```bash
cd /home/user/plenya

# Fase 1: Testes Neuropsicológicos
docker compose exec -T db psql -U plenya_user -d plenya_db < scripts/batch31_phase1_testes.sql

# Fase 2: Capacidades Cognitivas
docker compose exec -T db psql -U plenya_user -d plenya_db < scripts/batch31_phase2_capacidades.sql

# Fase 3: Escalas Validadas
docker compose exec -T db psql -U plenya_user -d plenya_db < scripts/batch31_phase3_escalas.sql

# Fase 4: Fatores Hormonais
docker compose exec -T db psql -U plenya_user -d plenya_db < scripts/batch31_phase4_hormonal.sql

# Fase 5: Histórico
docker compose exec -T db psql -U plenya_user -d plenya_db < scripts/batch31_phase5_historico.sql

# Verificação Final
docker compose exec -T db psql -U plenya_user -d plenya_db -c "SELECT COUNT(*) as completos FROM score_items si LEFT JOIN score_subgroups sg ON si.subgroup_id = sg.id LEFT JOIN score_groups g ON sg.group_id = g.id WHERE g.name = 'Cognição' AND si.clinical_relevance IS NOT NULL;"
```

### Método 2: Script Bash Unificado
```bash
chmod +x scripts/execute_batch31_all.sh
./scripts/execute_batch31_all.sh
```

## Validação

### Queries de Verificação
```sql
-- Total de items atualizados
SELECT COUNT(*) as completos
FROM score_items si
LEFT JOIN score_subgroups sg ON si.subgroup_id = sg.id
LEFT JOIN score_groups g ON sg.group_id = g.id
WHERE g.name = 'Cognição'
  AND si.clinical_relevance IS NOT NULL;

-- Por subgrupo
SELECT sg.name as subgrupo, COUNT(*) as completos
FROM score_items si
LEFT JOIN score_subgroups sg ON si.subgroup_id = sg.id
LEFT JOIN score_groups g ON sg.group_id = g.id
WHERE g.name = 'Cognição'
  AND si.clinical_relevance IS NOT NULL
GROUP BY sg.name
ORDER BY sg.name;

-- Items ainda pendentes
SELECT si.id, si.name, sg.name as subgrupo
FROM score_items si
LEFT JOIN score_subgroups sg ON si.subgroup_id = sg.id
LEFT JOIN score_groups g ON sg.group_id = g.id
WHERE g.name = 'Cognição'
  AND si.clinical_relevance IS NULL
ORDER BY sg.name, si.name;
```

## Cronograma Estimado
- **Fase 1:** 15 minutos
- **Fase 2:** 20 minutos
- **Fase 3:** 15 minutos
- **Fase 4:** 10 minutos
- **Fase 5:** 15 minutos
- **Validação:** 5 minutos
- **Total:** ~80 minutos

## Resultados Esperados
- 61 items completados
- 100% do grupo Cognição preenchido
- Progresso global: 377 → 438 items (18.9% do total)
- Conteúdos ricos em evidências MFI
- Orientações clínicas práticas e aplicáveis

## Próximos Passos Após Conclusão
1. Validar todos os items via queries SQL
2. Revisar amostra de 10% manualmente
3. Documentar lições aprendidas
4. Preparar próximo batch (Batch 32)
5. Atualizar dashboard de progresso

---

**Criado:** 2026-01-26
**Autor:** Claude Sonnet 4.5 + Medicina Funcional Integrativa
**Versão:** 1.0
