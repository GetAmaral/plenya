-- ============================================================================
-- ENRICHMENT: Genitália Feminina Score Item
-- ============================================================================
-- Score Item ID: 2b01e2f0-76c7-4616-8f7b-33b1c9d11279
-- Category: Vida Sexual
-- Date: 2026-01-29
-- Articles: 4 peer-reviewed sources (2024-2026)
-- ============================================================================

BEGIN;

-- ----------------------------------------------------------------------------
-- 1. INSERT SCIENTIFIC ARTICLES
-- ----------------------------------------------------------------------------

-- Article 1: Gynecologic Pelvic Examination (StatPearls 2024)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    doi,
    original_link,
    article_type,
    abstract,
    language,
    specialty,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Gynecologic Pelvic Examination',
    'Mikes BA, Wray AA',
    'StatPearls Publishing',
    '2024-02-25',
    NULL,
    'https://www.ncbi.nlm.nih.gov/books/NBK534223/',
    'review',
    'O exame pélvico ginecológico é um componente crítico do cuidado abrangente à saúde da mulher. Envolve inspeção dos genitais externos, exame especular e palpação bimanual para avaliar estruturas reprodutivas internas. Serve para diagnosticar condições incluindo sangramento anormal, corrimento, dor pélvica, infecções e tumores; rastreamento de malignidades; e monitoramento da saúde reprodutiva. A única contraindicação absoluta é a falta de consentimento informado. Diretrizes de organizações como ACOG e ACP variam quanto a exames de rotina em mulheres assintomáticas, com evidências recentes sugerindo benefício limitado para detecção de câncer ovariano, porém o exame pode detectar precocemente cânceres vulvares ou vaginais tratáveis. Complicações incluem desconforto, ansiedade ou trauma mucoso, especialmente em pacientes com vaginite atrófica ou histórico de trauma.',
    'pt-BR',
    'Ginecologia',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
);

-- Article 2: Self-Collected HPV Testing for Cervical Cancer Screening (CA Cancer J 2026)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    doi,
    original_link,
    article_type,
    abstract,
    language,
    specialty,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Self-collected vaginal specimens for human papillomavirus testing and guidance on screening exit: An update to the American Cancer Society cervical cancer screening guideline',
    'Perkins RB, Wolf AMD, Church TR, et al.',
    'CA: A Cancer Journal for Clinicians',
    '2026-01-01',
    '10.3322/caac.70041',
    'https://acsjournals.onlinelibrary.wiley.com/doi/10.3322/caac.70041',
    'research_article',
    'Atualização da diretriz de rastreamento de câncer cervical da American Cancer Society para incluir autocoleta de espécimes vaginais para teste de HPV. Para mulheres de risco médio, recomenda-se iniciar rastreamento aos 25 anos com teste primário de HPV a cada 5 anos até 65 anos. Para teste primário de HPV, espécimes cervicais coletados por clínico são preferidos, mas espécimes vaginais autocoletados são aceitáveis. Quando espécimes vaginais autocoletados são HPV negativos no rastreamento, recomenda-se repetir teste em 3 anos. Para descontinuar rastreamento, recomenda-se que mulheres de risco médio tenham testes primários de HPV negativos ou coteste negativo aos 60 e 65 anos. Esta mudança representa avanço significativo ao aumentar acesso ao rastreamento, particularmente para populações com barreiras ao cuidado tradicional.',
    'pt-BR',
    'Ginecologia Oncológica',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
);

-- Article 3: Vulvovaginal Candidiasis and Co-infections (Frontiers Microbiology 2025)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    doi,
    original_link,
    article_type,
    abstract,
    language,
    specialty,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'New insights toward personalized therapies for vulvovaginal candidiasis and vaginal co-infections',
    'De Seta F, Warris A, Roselletti E',
    'Frontiers in Microbiology',
    '2025-08-08',
    '10.3389/fmicb.2025.1625952',
    'https://www.frontiersin.org/journals/microbiology/articles/10.3389/fmicb.2025.1625952/full',
    'review',
    'Candidíase vulvovaginal (CVV) e vaginose bacteriana (VB) são as duas infecções vaginais mais prevalentes globalmente, particularmente entre mulheres em idade reprodutiva. Estas condições não apenas causam sintomas clínicos significativos, mas também impactam severamente a qualidade de vida e saúde mental das mulheres. Para infecções recorrentes, regimes estendidos com metronidazol ou ácido bórico são opções para VB, enquanto para CVV, tanto Candida albicans quanto espécies não-albicans contribuem para recorrências, frequentemente requerendo regimes antifúngicos estendidos. Mulheres com histórico de CVV recorrente têm 78% maior risco de serem posteriormente diagnosticadas com diabetes tipo 2, com risco especialmente pronunciado entre mulheres acima de 55 anos. O artigo explora tratamentos convencionais e emergentes, incluindo zinco, Lactobacillus spp. e ácido láctico.',
    'pt-BR',
    'Ginecologia',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
);

-- Article 4: Pelvic Exam in Pelvic Organ Prolapse (IJGO 2024)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    doi,
    original_link,
    article_type,
    abstract,
    language,
    specialty,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Value of pelvic examination in women with pelvic organ prolapse: A systematic review',
    'Pizzoferrato AC, Sallée C, Thubert T, Fauconnier A, Deffieux X',
    'International Journal of Gynecology & Obstetrics',
    '2024-05-22',
    '10.1002/ijgo.15697',
    'https://obgyn.onlinelibrary.wiley.com/doi/abs/10.1002/ijgo.15697',
    'meta_analysis',
    'Revisão sistemática de 67 estudos sobre valor do exame pélvico em mulheres com prolapso de órgãos pélvicos (POP). Sintomas associados ao prolapso são pobremente correlacionados com diagnóstico de POP, sendo abaulamento vaginal o sintoma com melhor correlação (moderada a boa) com estágio de POP. Os fatores mais fortemente associados com risco de recorrência após cirurgia ou falha de pessário são clínicos, incluindo estágio mais alto de POP antes da cirurgia, avulsão do músculo levantador do ânus, e medidas vaginais e genitais. Conclui-se que exame pélvico (especular vaginal e exame digital vaginal) pode confirmar presença de POP e identificar fatores de risco para falha de tratamento ou recorrência. Evidência robusta sobre valor diagnóstico e prognóstico do exame físico na avaliação de POP.',
    'pt-BR',
    'Uroginecologia',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
);

-- ----------------------------------------------------------------------------
-- 2. LINK ARTICLES TO SCORE ITEM
-- ----------------------------------------------------------------------------

INSERT INTO article_score_items (score_item_id, article_id)
SELECT
    '2b01e2f0-76c7-4616-8f7b-33b1c9d11279'::uuid,
    id
FROM articles
WHERE title IN (
    'Gynecologic Pelvic Examination',
    'Self-collected vaginal specimens for human papillomavirus testing and guidance on screening exit: An update to the American Cancer Society cervical cancer screening guideline',
    'New insights toward personalized therapies for vulvovaginal candidiasis and vaginal co-infections',
    'Value of pelvic examination in women with pelvic organ prolapse: A systematic review'
)
AND created_at >= CURRENT_TIMESTAMP - INTERVAL '1 minute';

-- ----------------------------------------------------------------------------
-- 3. UPDATE SCORE ITEM WITH CLINICAL CONTENT
-- ----------------------------------------------------------------------------

UPDATE score_items
SET
    clinical_relevance = 'A avaliação da genitália feminina é componente essencial do cuidado ginecológico integral, permitindo diagnóstico precoce de condições benignas e malignas, rastreamento oncológico e monitoramento da saúde reprodutiva. O exame pélvico ginecológico completo inclui inspeção visual da genitália externa (vulva, períneo, ânus), exame especular da vagina e colo uterino, e palpação bimanual das estruturas internas. Condições detectáveis incluem: cânceres cervical, endometrial, ovariano, vaginal e vulvar; infecções (vaginose bacteriana, candidíase, clamídia, gonorreia, verrugas genitais, herpes genital, doença inflamatória pélvica, tricomoníase); e condições benignas (vaginite atrófica, pólipos cervicais, endometriose, cistos ovarianos, prolapso de órgãos pélvicos, miomas uterinos, líquen escleroso vulvar). O rastreamento de câncer cervical evoluiu significativamente com evidências de 2024-2026 recomendando teste primário de HPV a cada 5 anos iniciando aos 25 anos, com opção de autocoleta vaginal aumentando acessibilidade. Sintomas que justificam avaliação urgente incluem dor vulvovaginal persistente, sangramento anormal, corrimento com odor fétido, prurido intenso, lesões visíveis ou abaulamento vaginal. A única contraindicação absoluta ao exame é ausência de consentimento informado, sendo fundamental abordagem centrada na paciente com comunicação empática, especialmente em indivíduos com histórico de trauma.',

    patient_explanation = 'A genitália feminina refere-se aos órgãos reprodutivos externos (vulva, clitóris, lábios) e internos (vagina, colo do útero, útero, ovários). O exame ginecológico é importante para manter sua saúde sexual e reprodutiva em dia, detectar problemas precocemente e prevenir doenças graves como câncer. Durante a consulta, o médico primeiro conversa sobre seus sintomas, histórico menstrual e vida sexual. O exame físico começa com inspeção visual da região genital externa, seguido por exame com espéculo (instrumento que permite visualizar o interior da vagina e colo do útero para colher exames como Papanicolau ou teste de HPV) e palpação suave do abdômen e região pélvica para avaliar útero e ovários. Novidades importantes: desde 2024, você pode fazer autocoleta vaginal para teste de HPV em casa, facilitando o rastreamento de câncer cervical. Sinais de alerta que exigem avaliação incluem sangramento fora do período menstrual, corrimento com mau cheiro, coceira intensa, dor durante relação sexual, feridas ou caroços na região genital, ou sensação de "bola" saindo pela vagina. Problemas comuns incluem infecções por fungos (candidíase), vaginose bacteriana, infecções sexualmente transmissíveis e prolapso (descida) de órgãos pélvicos. O exame pode ser desconfortável, mas não deve doer; comunique qualquer desconforto ao profissional. Você tem direito a solicitar acompanhante durante o exame e pode recusar procedimentos - seu consentimento é obrigatório.',

    conduct = 'AVALIAÇÃO INICIAL: Anamnese detalhada incluindo ciclo menstrual, última menstruação, atividade sexual, métodos contraceptivos, histórico de infecções, sintomas atuais (sangramento anormal, corrimento, dor, prurido, dispareunia, lesões), histórico obstétrico e cirúrgico ginecológico. Avaliar fatores de risco para ISTs, câncer cervical (HPV, tabagismo, imunossupressão) e comorbidades (diabetes aumenta risco de candidíase recorrente em 78%). EXAME FÍSICO: 1) Inspeção externa: avaliar pele vulvar, distribuição pilosa, lesões, ulcerações, massas, corrimento, sinais de trauma, anatomia dos lábios maiores/menores, clitóris, períneo e ânus; 2) Exame especular: visualizar paredes vaginais (atrofia, inflamação, lesões), colo uterino (ectopia, pólipos, lesões suspeitas), características do corrimento (aquoso, espesso, cor, odor); colher material para citologia/HPV conforme protocolo de rastreamento; 3) Palpação bimanual: avaliar colo (consistência, mobilidade, dor à mobilização), útero (tamanho, formato, posição, mobilidade, massas), anexos (massas, sensibilidade), fundo de saco (abaulamentos, dor). RASTREAMENTO ONCOLÓGICO (diretrizes ACS 2026): Iniciar aos 25 anos com teste primário HPV a cada 5 anos até 65 anos; autocoleta vaginal aceitável se HPV negativo, repetir em 3 anos; descontinuar após testes negativos aos 60 e 65 anos em pacientes de risco médio. MANEJO DE CONDIÇÕES COMUNS: Candidíase vulvovaginal - fluconazol 150mg dose única ou clotrimazol tópico; se recorrente (≥4 episódios/ano), regime estendido e investigar diabetes. Vaginose bacteriana - metronidazol 500mg 2x/dia 7 dias ou gel vaginal; recorrências requerem regime estendido. ISTs - tratamento conforme agente (azitromicina para clamídia, ceftriaxona para gonorreia) e notificação de parceiros. Vaginite atrófica - lubrificantes vaginais ou estrogênio tópico de baixa dose. CRITÉRIOS DE ENCAMINHAMENTO: Lesões vulvares/vaginais/cervicais suspeitas (biópsia); sangramento pós-menopausa; massa pélvica palpável; prolapso sintomático estágio III-IV; dor pélvica crônica não responsiva; candidíase recorrente refratária; citologia anormal ou HPV persistente (colposcopia); suspeita de malignidade. CONSIDERAÇÕES ESPECIAIS: Obter consentimento informado; oferecer acompanhante; modificar técnica em pacientes com histórico de trauma sexual; documentar achados detalhadamente; registrar em audit log conforme LGPD. Pacientes com sintomas agudos (dor intensa, sangramento volumoso, febre com corrimento) requerem avaliação urgente para excluir doença inflamatória pélvica, gravidez ectópica ou torção anexial.',

    updated_at = CURRENT_TIMESTAMP,
    last_review = CURRENT_DATE
WHERE id = '2b01e2f0-76c7-4616-8f7b-33b1c9d11279';

-- ----------------------------------------------------------------------------
-- 4. VERIFICATION QUERY
-- ----------------------------------------------------------------------------

-- Display updated score item with character counts
SELECT
    id,
    name,
    LENGTH(clinical_relevance) as clinical_relevance_chars,
    LENGTH(patient_explanation) as patient_explanation_chars,
    LENGTH(conduct) as conduct_chars,
    last_review,
    (SELECT COUNT(*)
     FROM article_score_items asi
     WHERE asi.score_item_id = si.id) as linked_articles_count
FROM score_items si
WHERE id = '2b01e2f0-76c7-4616-8f7b-33b1c9d11279';

-- Display linked articles
SELECT
    a.title,
    a.authors,
    a.journal,
    a.publish_date,
    a.doi,
    a.article_type,
    LENGTH(a.abstract) as abstract_chars
FROM articles a
JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = '2b01e2f0-76c7-4616-8f7b-33b1c9d11279'
ORDER BY a.publish_date DESC;

COMMIT;

-- ============================================================================
-- END OF ENRICHMENT
-- ============================================================================
