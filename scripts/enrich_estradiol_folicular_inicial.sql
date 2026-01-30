-- ============================================================================
-- ENRICHMENT: Estradiol - Mulheres Fase Folicular Inicial (Dias 1-7)
-- Score Item ID: 2208d693-f4b8-4332-8030-53e795aa5ef4
-- Generated: 2026-01-29
-- ============================================================================

BEGIN;

-- ============================================================================
-- STEP 1: Insert Articles (with duplicate prevention)
-- ============================================================================

-- Article 1: Estradiol Reference Intervals (LC-MS/MS Method)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    doi,
    article_type,
    abstract,
    language,
    specialty,
    created_at,
    updated_at
)
SELECT
    gen_random_uuid(),
    'Estradiol reference intervals in women during the menstrual cycle, postmenopausal women and men using an LC-MS/MS method',
    'Verdonk SJE, Vesper HW, Martens F, Sluss PM, Hillebrand JJ, Heijboer AC',
    'Clinica Chimica Acta',
    '2019-08-01',
    '10.1016/j.cca.2019.04.062',
    'research_article',
    'Estudo estabelece intervalos de referência harmonizados para estradiol usando método LC-MS/MS de alta precisão. Analisa 30 mulheres pré-menopáusicas com monitoramento diário durante ciclos menstruais completos. Demonstra variações cíclicas com níveis mais baixos na fase folicular inicial e picos durante o surto de LH (275-2864 pmol/L). O método mostra rastreabilidade aos padrões de referência do CDC e é adequado para tomada de decisões clínicas.',
    'en',
    'Endocrinologia',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
WHERE NOT EXISTS (
    SELECT 1 FROM articles WHERE doi = '10.1016/j.cca.2019.04.062'
);

-- Article 2: Physiology of Menstrual Cycle
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    article_type,
    abstract,
    language,
    specialty,
    pm_id,
    created_at,
    updated_at
)
SELECT
    gen_random_uuid(),
    'Physiology, Menstrual Cycle',
    'Thiyagarajan DK, Basit H, Jeanmonod R',
    'StatPearls Publishing',
    '2024-09-27',
    'review',
    'Revisão abrangente da fisiologia do ciclo menstrual publicada no StatPearls. Cobre regulação hormonal, fases do ciclo e significância clínica. Descreve detalhadamente a fase folicular inicial, incluindo níveis basais de estradiol (25-75 pg/mL no dia 3), recrutamento folicular mediado por FSH e produção gradual de estradiol pelos folículos em desenvolvimento. Recurso essencial para profissionais de saúde.',
    'en',
    'Ginecologia',
    '29763196',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
WHERE NOT EXISTS (
    SELECT 1 FROM articles WHERE pm_id = '29763196'
);

-- Article 3: Proliferative and Follicular Phases
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    article_type,
    abstract,
    language,
    specialty,
    pm_id,
    created_at,
    updated_at
)
SELECT
    gen_random_uuid(),
    'Proliferative and Follicular Phases of the Menstrual Cycle',
    'Monis CN, Tetrokalashvili M',
    'StatPearls Publishing',
    '2022-09-12',
    'review',
    'Artigo descritivo sobre a fase folicular e proliferativa do ciclo menstrual feminino. Explica o processo de maturação dos folículos ovarianos durante a fase folicular, preparando-os para liberação na ovulação. Aborda mudanças endometriais concomitantes e a importância clínica da fase folicular inicial para avaliação de reserva ovariana e função reprodutiva. Inclui correlações com níveis hormonais e aplicações diagnósticas.',
    'en',
    'Ginecologia',
    '31194386',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
WHERE NOT EXISTS (
    SELECT 1 FROM articles WHERE pm_id = '31194386'
);

-- Article 4: Fertility Indicator Using Estradiol
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    doi,
    article_type,
    abstract,
    language,
    specialty,
    created_at,
    updated_at
)
SELECT
    gen_random_uuid(),
    'A Novel Fertility Indicator Equation Using Estradiol Levels for Assessment of Phase of the Menstrual Cycle',
    'Usala SJ, Trindade AA',
    'Medicina (Kaunas)',
    '2020-10-22',
    '10.3390/medicina56110555',
    'research_article',
    'Pesquisa inovadora propondo equação indicadora de fertilidade (FIE) baseada na taxa de mudança dos níveis de estradiol sérico ao longo do ciclo. Demonstra que estradiol diário serve como melhor analito para determinar a fase do ciclo ovulatório e o dia de início da fertilidade. Estabelece que níveis de E2 ~411 pmol/L (112 pg/mL) indicam transição para fase fértil. Aplicável ao monitoramento de ciclos naturais em tratamento reprodutivo assistido.',
    'en',
    'Medicina Reprodutiva',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
WHERE NOT EXISTS (
    SELECT 1 FROM articles WHERE doi = '10.3390/medicina56110555'
);

-- ============================================================================
-- STEP 2: Link Articles to Score Item
-- ============================================================================

INSERT INTO article_score_items (score_item_id, article_id)
SELECT
    '2208d693-f4b8-4332-8030-53e795aa5ef4'::uuid,
    id
FROM articles
WHERE (
    doi IN (
        '10.1016/j.cca.2019.04.062',
        '10.3390/medicina56110555'
    )
    OR pm_id IN (
        '29763196',
        '31194386'
    )
)
AND NOT EXISTS (
    SELECT 1 FROM article_score_items asi
    WHERE asi.score_item_id = '2208d693-f4b8-4332-8030-53e795aa5ef4'::uuid
    AND asi.article_id = articles.id
);

-- ============================================================================
-- STEP 3: Update Score Item with Clinical Content
-- ============================================================================

UPDATE score_items
SET
    clinical_relevance = 'O estradiol (17-β estradiol) é o estrogênio mais potente e biologicamente ativo produzido pelos ovários. Durante a fase folicular inicial (dias 1-7 do ciclo menstrual), os níveis de estradiol encontram-se em seu ponto mais baixo, refletindo o estado de repouso ovariano antes do recrutamento e desenvolvimento folicular. Esta fase é crucial para avaliação da reserva ovariana e função reprodutiva basal. Os valores de referência para a fase folicular inicial variam entre 25-75 pg/mL (aproximadamente 92-275 pmol/L), com níveis ideais próximos à extremidade inferior desta faixa. No dia 3 do ciclo, o estradiol deve estar próximo ao seu nadir, uma vez que nenhum folículo se tornou dominante ainda, confirmando que os ovários estão verdadeiramente em repouso. A dosagem de estradiol no dia 3, em conjunto com FSH, constitui um dos testes mais importantes para avaliação de fertilidade e reserva ovariana. Níveis elevados de estradiol (>60-80 pg/mL) no dia 3 podem indicar presença de cisto ovariano funcional ou diminuição da reserva ovariana, podendo levar a valores artificialmente normais de FSH devido ao feedback negativo. Níveis baixos (~30 pg/mL) são críticos antes da estimulação ovariana em ciclos de reprodução assistida, pois valores mais altos estão associados a ciclos de estimulação inadequados. A mensuração precisa do estradiol na fase folicular inicial permite identificar disfunções ovulatórias, avaliar resposta a tratamentos de fertilidade e diagnosticar condições como síndrome dos ovários policísticos, insuficiência ovariana prematura e hipogonadismo hipogonadotrófico. A interpretação deve sempre considerar o contexto clínico, idade da paciente, regularidade menstrual e método analítico utilizado (preferencialmente LC-MS/MS para maior acurácia).',

    patient_explanation = 'O estradiol é o principal hormônio feminino produzido pelos ovários, responsável por regular seu ciclo menstrual e manter a saúde reprodutiva. Durante os primeiros dias da menstruação (dias 1 a 7), chamados de fase folicular inicial, os níveis de estradiol estão naturalmente baixos. Isso acontece porque neste período seus ovários estão "em repouso", antes de começarem a preparar um óvulo para ser liberado. Imagine o ciclo menstrual como uma jornada que se reinicia a cada mês: no início desta jornada (durante e logo após a menstruação), o estradiol está em seus níveis mais baixos, geralmente entre 25 e 75 pg/mL. Este é o momento ideal para avaliar como seus ovários estão funcionando em seu estado basal, sem influência de hormônios elevados. A dosagem do estradiol no terceiro dia do ciclo (dia 3) é frequentemente solicitada junto com outros hormônios para avaliar sua fertilidade e reserva ovariana - ou seja, a quantidade e qualidade de óvulos que seus ovários ainda podem produzir. Se o estradiol estiver muito alto nesta fase inicial, pode ser um sinal de alerta indicando que algo não está funcionando corretamente, como a presença de um cisto no ovário ou diminuição da reserva ovariana. Por outro lado, níveis adequadamente baixos nesta fase confirmam que seus ovários estão saudáveis e prontos para iniciar um novo ciclo. Este exame é especialmente importante para mulheres que estão investigando dificuldades para engravidar, planejando tratamentos de fertilidade ou querendo entender melhor como seus hormônios estão funcionando ao longo do ciclo menstrual.',

    conduct = 'PROTOCOLO CLÍNICO PARA AVALIAÇÃO DE ESTRADIOL NA FASE FOLICULAR INICIAL:

1. TIMING DA COLETA (CRÍTICO):
   - Realizar coleta entre dias 1-7 do ciclo menstrual, preferencialmente no dia 3
   - Dia 1 = primeiro dia de fluxo menstrual completo (não spotting)
   - Para avaliação de fertilidade/reserva ovariana: coletar especificamente no dia 3
   - Pacientes com ciclos irregulares: documentar fase do ciclo e considerar múltiplas coletas
   - Coleta em qualquer horário do dia (não há variação circadiana significativa)

2. INTERPRETAÇÃO DOS RESULTADOS - DIA 3:

   VALORES IDEAIS (25-50 pg/mL):
   - Confirmam ovários em repouso adequado
   - Reserva ovariana provavelmente normal
   - Resposta adequada esperada à estimulação ovariana
   - Prosseguir com avaliação complementar (FSH, AMH, contagem de folículos antrais)

   VALORES MODERADOS (51-75 pg/mL):
   - Ainda dentro da normalidade, porém no limite superior
   - Avaliar em conjunto com FSH (razão estradiol/FSH)
   - Considerar repetir em próximo ciclo para confirmação
   - Investigar causas de elevação discreta

   VALORES ELEVADOS (>75-80 pg/mL):
   - Sugestivo de cisto ovariano funcional → realizar ultrassom pélvico
   - Possível diminuição de reserva ovariana → solicitar AMH, contagem folicular
   - Pode mascarar FSH elevado (feedback negativo) → interpretar com cautela
   - Contraindicação relativa para início de estimulação ovariana
   - Repetir dosagem em ciclo subsequente após resolução de cisto

   VALORES MUITO BAIXOS (<25 pg/mL):
   - Considerar erro de timing (confirmar fase do ciclo)
   - Avaliar função ovariana global (FSH, LH, AMH)
   - Em contexto de amenorreia: investigar hipogonadismo hipogonadotrófico

3. CORRELAÇÃO COM OUTROS MARCADORES:
   - FSH dia 3: valores combinados melhoram acurácia diagnóstica
   - Razão FSH/LH: útil em suspeita de SOP
   - AMH: marcador independente do ciclo, complementa avaliação
   - Ultrassom com contagem de folículos antrais: avaliação morfológica
   - Inibina B: marcador adicional de reserva ovariana

4. MÉTODO LABORATORIAL:
   - Preferir ensaios LC-MS/MS (maior especificidade e acurácia)
   - Evitar imunoensaios em níveis muito baixos (menor sensibilidade)
   - Verificar rastreabilidade do método aos padrões CDC
   - Conhecer valores de referência específicos do laboratório

5. CONDUTAS ESPECÍFICAS POR CONTEXTO CLÍNICO:

   INVESTIGAÇÃO DE INFERTILIDADE:
   - Solicitar painel completo dia 3: estradiol, FSH, LH, TSH, prolactina
   - Programar ultrassom transvaginal no mesmo período
   - Repetir em 2-3 ciclos se resultados limítrofes

   PRÉ-FERTILIZAÇÃO IN VITRO:
   - Estradiol <30-40 pg/mL necessário para início de estimulação
   - Se >50 pg/mL: adiar ciclo e investigar causa
   - Monitorar evolução durante estimulação ovariana

   AMENORREIA/OLIGOMENORREIA:
   - Estradiol baixo + FSH alto: insuficiência ovariana
   - Estradiol baixo + FSH baixo: hipogonadismo central
   - Correlacionar com quadro clínico e exame físico

6. SEGUIMENTO:
   - Resultados normais + clínica compatível: prosseguir investigação conforme protocolo
   - Resultados alterados: repetir em 1-2 ciclos subsequentes
   - Encaminhar para especialista em reprodução se: estradiol persistentemente >80 pg/mL, suspeita de insuficiência ovariana, planejamento de TRA',

    updated_at = CURRENT_TIMESTAMP
WHERE id = '2208d693-f4b8-4332-8030-53e795aa5ef4'::uuid;

-- ============================================================================
-- VERIFICATION QUERY
-- ============================================================================

SELECT
    si.id,
    si.name,
    LENGTH(si.clinical_relevance) as clinical_relevance_chars,
    LENGTH(si.patient_explanation) as patient_explanation_chars,
    LENGTH(si.conduct) as conduct_chars,
    COUNT(DISTINCT asi.article_id) as linked_articles_count
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = '2208d693-f4b8-4332-8030-53e795aa5ef4'::uuid
GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct;

COMMIT;

-- ============================================================================
-- END OF ENRICHMENT
-- ============================================================================
