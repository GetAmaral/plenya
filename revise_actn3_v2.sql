-- ================================================================
-- REVISÃO COMPLETA: ACTN3 rs1815739 R577X (Performance)
-- ID: 019c1a2b-a36f-77ac-984e-0dcda6b517f0
-- ================================================================

BEGIN;

-- 0. Criar snapshot do estado atual
DO $$
DECLARE
    v_before jsonb;
BEGIN
    SELECT jsonb_build_object(
        'clinical_relevance', clinical_relevance,
        'patient_explanation', patient_explanation,
        'conduct', conduct,
        'points', points,
        'last_review', last_review
    ) INTO v_before
    FROM score_items
    WHERE id = '019c1a2b-a36f-77ac-984e-0dcda6b517f0';

    -- Guardar em temporary table
    CREATE TEMP TABLE before_state AS
    SELECT v_before as snapshot;
END $$;

-- 1. Atualizar campos clínicos do ScoreItem
UPDATE score_items
SET
    clinical_relevance = 'Gene ACTN3 codifica alpha-actinina-3, proteína estrutural Z-disc expressa exclusivamente em fibras musculares de contração rápida (tipo II/IIx). SNP rs1815739 (R577X, c.1729C>T) resulta em substituição de arginina por códon de parada prematuro (nonsense mutation), gerando alelo nulo X com perda completa de função da proteína. Aproximadamente 18% da população mundial é homozigota XX (null genotype), com ausência total de alpha-actinina-3 funcional, percentual variando significativamente por ancestralidade: ~25% europeus, ~1% africanos subsaarianos, ~20% asiáticos orientais. Genótipo XX associado a desempenho reduzido em atividades de potência explosiva, sprint e força máxima, mas vantagem potencial em resistência aeróbica prolongada, possivelmente via compensação por alpha-actinina-2 (ACTN2) e metabolismo oxidativo aumentado. Estudos GWAS e metanálises demonstram: (1) Genótipo RR e RX superrepresentados em atletas elite de potência/sprint (76-95% vs 50% população geral); (2) Genótipo XX mais frequente em atletas de ultra-resistência; (3) Modulação de resposta hipertrófica ao treinamento resistido; (4) Influência em recuperação muscular pós-exercício; (5) Interação com metabolismo de cálcio muscular e contratilidade. Polimorfismo também investigado em contextos não-atléticos: função muscular no envelhecimento, sarcopenia, fragilidade, miopatias metabólicas, obesidade e metabolismo glicêmico (achados inconsistentes). Evidências emergentes sugerem papel em termorregulação durante exercício em calor. Interpretação clínica requer integração com: (1) Outros genes de performance (ACE I/D, PPARGC1A, PPARA, ADRB2); (2) Fibrotipagem muscular funcional; (3) VO2max e limiar anaeróbico; (4) Avaliação de potência/força explosiva; (5) Histórico de resposta a treinamento. Frequências alélicas: alelo R ~55%, alelo X ~45% (população europeia). Penetrância variável, expressividade dependente de treinamento, nutrição, sono e modulação epigenética.',

    patient_explanation = 'O gene ACTN3 produz uma proteína chamada alpha-actinina-3, encontrada exclusivamente nas fibras musculares de contração rápida - aquelas responsáveis por explosões de força, velocidade e potência (como sprints, saltos, levantamento de peso). Você possui uma variante genética (rs1815739 R577X) que determina se seu corpo produz ou não essa proteína. Existem três possibilidades: (1) Genótipo RR: você produz alpha-actinina-3 normalmente em ambas as cópias do gene - vantagem natural para atividades de potência, sprint e força explosiva; (2) Genótipo RX: você produz a proteína em quantidade reduzida (uma cópia funcional) - perfil intermediário; (3) Genótipo XX: você não produz nenhuma alpha-actinina-3 funcional - aproximadamente 1 em cada 5 pessoas tem este perfil, com vantagem potencial para resistência aeróbica prolongada, mas desvantagem para potência explosiva. IMPORTANTE: ausência de ACTN3 (genótipo XX) NÃO é doença ou defeito genético! É uma variação completamente normal e funcional encontrada em bilhões de pessoas saudáveis. Seu corpo compensa perfeitamente usando outra proteína similar (alpha-actinina-2). O genótipo XX foi mantido na evolução humana provavelmente porque oferece vantagens metabólicas específicas. O que isso significa na prática: seu genótipo ACTN3 NÃO determina se você pode ou não ser atlético, forte ou saudável, mas sugere qual tipo de treinamento físico e esporte seu corpo pode responder melhor naturalmente. Com treinamento adequado, nutrição otimizada e consistência, pessoas com qualquer genótipo ACTN3 alcançam excelentes resultados - a genética sugere tendências, não limites absolutos. Conhecer seu genótipo permite personalizar tipo de exercício, volume de treino, periodização e recuperação para maximizar resultados individuais.',

    conduct = 'Testar quando: (1) Otimização de prescrição de exercício físico personalizado; (2) Planejamento de performance atlética e seleção de modalidade esportiva baseada em potencial genético; (3) Avaliação de risco de sarcopenia e fragilidade em envelhecimento; (4) Investigação de resposta subótima a treinamento resistido ou aeróbico; (5) Medicina esportiva preventiva e preditiva. Método: genotipagem por PCR em tempo real (TaqMan) para SNP rs1815739 (C>T), microarray genômico (Illumina, Affymetrix), ou NGS se painel genético amplo. Interpretar sempre com frequências alélicas populacionais específicas por ancestralidade. Exames complementares direcionados: (1) Composição corporal por DXA (massa muscular total, distribuição regional, massa muscular apendicular); (2) Fibrotipagem muscular não-invasiva (estimativa via bioimpedância de fase ou EMG de superfície); (3) Testes funcionais: força máxima (1RM agachamento, supino), potência explosiva (salto vertical, Wingate test), sprint 30-60m, VO2max e limiar ventilatório; (4) Marcadores de turnover muscular: CPK basal e pós-exercício, mioglobina, LDH; (5) Painel metabólico: glicose, insulina, HOMA-IR (possível associação com metabolismo glicêmico em genótipo XX); (6) Inflamação sistêmica: PCR-us, IL-6; (7) Status de micronutrientes críticos para função muscular: vitamina D (25-OH), magnésio eritrocitário, selênio, zinco. Conduta personalizada rigorosamente por genótipo: **RR (produtor normal alpha-actinina-3):** Vantagem natural para potência/força explosiva - priorizar treinamento resistido com ênfase em força máxima e potência (cargas 85-95% 1RM, 1-5 repetições, explosividade), pliometria de alto impacto, sprints curtos (10-60m), recuperação adequada entre séries intensas (3-5 min). Nutrição: ingestão proteica 1.8-2.2g/kg/dia, timing periworkout otimizado (proteína + carboidrato 30-60min pré e pós-treino), creatina monohidratada 3-5g/dia (resposta potencialmente superior), beta-alanina 3-6g/dia para capacidade anaeróbica. **RX (produtor parcial):** Perfil versátil - excelente resposta a treinamento concorrente (força + resistência), periodização polarizada ou ondulada, combinar treinos de força (3x/semana) com aeróbico moderado-intenso (2-3x/semana). Flexibilidade estratégica conforme objetivos específicos. **XX (deficiência completa alpha-actinina-3):** Vantagem potencial para resistência aeróbica - priorizar treinos aeróbicos prolongados, HIIT com intervalos mais longos (>2min), treinamento resistido com volume alto e carga moderada (10-15 repetições, 65-75% 1RM, hipertrofia metabólica), menor tempo de recuperação entre séries. Atenção especial para: (1) Periodização cuidadosa em treinamento de força explosiva (requer mais tempo de adaptação neuromuscular); (2) Aquecimento mais prolongado antes de atividades de potência; (3) Suplementação estratégica: beta-alanina para compensar limitação em buffer anaeróbico, HMB (3g/dia) pode auxiliar preservação muscular, taurina para função contrátil. Vitamina D >40ng/mL essencial. **Todos os genótipos:** Sono 7-9h/noite (recuperação muscular), controle de estresse crônico (cortisol catabólico), hidratação adequada, anti-inflamatórios naturais (ômega-3 EPA+DHA 2-3g/dia, curcumina lipossomal, polifenóis). Monitoramento longitudinal: reavaliação funcional (força, potência, VO2max) a cada 12 semanas, composição corporal DXA semestral, ajuste de protocolo conforme resposta individual objetiva. Integrar sempre com outros genes de performance para estratégia verdadeiramente personalizada (ACE, PPARGC1A, PPARA, ADRB2, VDR). Considerar avaliação epigenética (metilação DNA muscular) se resposta anômala ao treinamento prescrito.',

    points = 0, -- Gene de performance não contribui para score de risco

    last_review = NOW(),
    updated_at = NOW()

WHERE id = '019c1a2b-a36f-77ac-984e-0dcda6b517f0';

-- 2. Criar snapshot do estado atualizado
DO $$
DECLARE
    v_before jsonb;
    v_after jsonb;
BEGIN
    SELECT snapshot INTO v_before FROM before_state;

    SELECT jsonb_build_object(
        'clinical_relevance', clinical_relevance,
        'patient_explanation', patient_explanation,
        'conduct', conduct,
        'points', points,
        'last_review', last_review
    ) INTO v_after
    FROM score_items
    WHERE id = '019c1a2b-a36f-77ac-984e-0dcda6b517f0';

    -- Registrar no histórico de revisão
    INSERT INTO score_item_review_history (
        score_item_id,
        reviewed_at,
        review_type,
        before_snapshot,
        after_snapshot,
        tier,
        confidence_score,
        model_used
    ) VALUES (
        '019c1a2b-a36f-77ac-984e-0dcda6b517f0',
        NOW(),
        'manual_comprehensive_revision',
        v_before,
        v_after,
        'manual',
        1.0,
        'human_expert_review'
    );
END $$;

-- 3. Verificar resultado
SELECT
    '=== RESULTADO DA ATUALIZAÇÃO ===' as info;

SELECT
    id,
    name,
    LENGTH(clinical_relevance) as clinical_relevance_chars,
    LENGTH(patient_explanation) as patient_explanation_chars,
    LENGTH(conduct) as conduct_chars,
    points as max_points,
    last_review,
    updated_at
FROM score_items
WHERE id = '019c1a2b-a36f-77ac-984e-0dcda6b517f0';

-- 4. Verificar artigos linkados
SELECT
    '=== ARTIGOS LINKADOS ===' as info;

SELECT
    COUNT(*) as total_artigos_linkados
FROM article_score_items
WHERE score_item_id = '019c1a2b-a36f-77ac-984e-0dcda6b517f0';

SELECT
    asi.confidence_score,
    asi.auto_linked,
    asi.linked_at,
    LEFT(a.title, 80) as title_preview
FROM article_score_items asi
JOIN articles a ON asi.article_id = a.id
WHERE asi.score_item_id = '019c1a2b-a36f-77ac-984e-0dcda6b517f0'
ORDER BY asi.confidence_score DESC;

COMMIT;

-- ================================================================
-- SUCESSO! Campos clínicos atualizados com audit trail
-- ================================================================
