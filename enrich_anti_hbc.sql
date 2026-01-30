-- ============================================================================
-- Enrichment for Score Item: Hepatite B - Anti-Hbc
-- Item ID: b3cb69e0-98fd-48ca-82ae-9eee62a8218e
-- Generated: 2026-01-29
-- ============================================================================

BEGIN;

-- Insert peer-reviewed articles
INSERT INTO articles (id, title, authors, journal, publish_date, doi, pm_id, article_type, original_link, abstract, language, created_at, updated_at)
VALUES
  -- Article 1: 2025 Clinical Guidelines
  (
    gen_random_uuid(),
    'Chronic hepatitis B in 2025: diagnosis, treatment and future directions',
    'Watson AG, Mulay AS, Gill US',
    'Clinical Medicine (London)',
    '2025-11-04',
    NULL,
    NULL,
    'review',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC12681808/',
    'Comprehensive 2025 review on hepatitis B diagnosis emphasizing anti-HBc role in serological pattern interpretation. Highlights common diagnostic pitfall of isolated anti-HBc positivity (typically reflects past exposure, requires management only if immunosuppression planned). Distinguishes anti-HBc IgM (acute infection) from IgG (chronic/resolved). Emphasizes reactivation risk in immunosuppressed anti-HBc-positive patients, particularly with B-cell depleting agents, warranting prophylactic therapy.',
    'en',
    NOW(),
    NOW()
  ),

  -- Article 2: Quantitative Anti-HBc
  (
    gen_random_uuid(),
    'Hepatitis B Core Antibody Level: A Surrogate Marker for Host Antiviral Immunity in Chronic Hepatitis B Virus Infections',
    'Multiple authors',
    'Viruses',
    '2023-05-03',
    '10.3390/v15051111',
    'PMC10221631',
    'research_article',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC10221631/',
    'Demonstrates quantitative anti-HBc (qAnti-HBc) as universal anti-HBV immune surrogate reflecting host immune response. Patients in immune-active phases show ~10-fold higher qAnti-HBc versus immune-tolerant phases. Correlates positively with ALT activity and histological inflammation severity. Cut-off ~4.5 log10 IU/mL diagnoses moderate-to-severe inflammation in normal-ALT patients. Baseline qAnti-HBc >4.0-4.5 log10 IU/mL independently predicts HBeAg seroconversion with antivirals.',
    'en',
    NOW(),
    NOW()
  ),

  -- Article 3: Clinical Significance and Standardization
  (
    gen_random_uuid(),
    'Clinical Significance and Remaining Issues of Anti-HBc Antibody and HBV Core-Related Antigen',
    'Multiple authors',
    'Diagnostics (Basel)',
    '2024-03-29',
    '10.3390/diagnostics14070728',
    NULL,
    'review',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC11011781/',
    'Anti-HBc appears early in HBV infection and persists lifelong, crucial for blood safety screening and identifying past exposure. Major challenges: lack of standardization across assay systems (ELISA, CMIA, CLIA, CLEIA) producing results in different units (IU/mL, INH%, S/O), complicating interpretation. Difficulty distinguishing low-titer true positives from false positives with highly sensitive assays. Quantitative anti-HBc predicts treatment response, carcinogenesis risk, reactivation, and HBsAg clearance during therapy.',
    'en',
    NOW(),
    NOW()
  ),

  -- Article 4: Occult Hepatitis B
  (
    gen_random_uuid(),
    'Occult Hepatitis B Virus Infection: An Update',
    'Saitta C, Pollicino T, Raimondo G',
    'Viruses',
    '2022-07-08',
    '10.3390/v14071504',
    NULL,
    'review',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC9318873/',
    'Occult HBV infection (OBI): HBsAg-negative with replication-competent HBV DNA in liver. Seropositive OBI (~80%): anti-HBc/anti-HBs detectable; seronegative OBI (~20%): all markers absent. Anti-HBc serves as surrogate marker for OBI screening in blood donors, organ recipients, immunosuppressed patients. Reactivation occurs in up to 40% with immunosuppression (highest risk: anti-CD20 therapies >10%; moderate: chemotherapy/TNFα inhibitors 1-10%). OBI accelerates cirrhosis in concurrent liver disease, maintains oncogenic potential (HBV integration in 75% OBI-HCC). Prophylactic antivirals recommended for high-risk immunosuppression in anti-HBc-positive patients.',
    'en',
    NOW(),
    NOW()
  );

-- Update score_items with comprehensive clinical content
UPDATE score_items
SET
  clinical_relevance = 'O anti-HBc (anticorpo contra antígeno core do vírus da hepatite B) é um marcador sorológico fundamental que persiste por toda a vida após exposição ao HBV, essencial para compreender o perfil imunológico do paciente. Diferencia-se em anti-HBc IgM (indica infecção aguda) e anti-HBc IgG (reflete infecção crônica ou resolvida). A presença de anti-HBc isolado (HBsAg negativo) representa um dos padrões mais desafiadores na interpretação sorológica: pode indicar infecção passada com perda de anti-HBs, infecção oculta (OBI), período de janela imunológica ou falso-positivo. A infecção oculta por HBV (OBI) ocorre em pacientes HBsAg-negativos com DNA viral detectável no fígado; aproximadamente 80% apresentam soropositividade para anti-HBc, tornando este marcador crucial para triagem de doadores de sangue, receptores de órgãos e candidatos a imunossupressão. Estudos recentes demonstram que a dosagem quantitativa de anti-HBc (qAnti-HBc) funciona como biomarcador da resposta imune do hospedeiro: pacientes em fase imuno-ativa apresentam níveis ~10 vezes maiores que em fase imuno-tolerante, correlacionando-se positivamente com atividade inflamatória hepática (ALT, histologia). Valores de corte de 4,0-4,5 log10 UI/mL predizem independentemente soroconversão HBeAg e inflamação moderada-grave. O anti-HBc também estratifica risco de reativação viral em imunossupressão: até 40% de reativação em terapias anti-CD20, moderado (1-10%) com quimioterapia/inibidores TNF-α, baixo (<1%) com antivirais diretos HCV. Diretrizes europeias recomendam profilaxia com análogos nucleos(t)ídeos para anti-HBc-positivos submetidos a alto risco de imunossupressão, mantendo terapia até 18 meses pós-tratamento. A OBI mantém potencial oncogênico (integração viral em 75% dos casos HCC-OBI), acelera progressão para cirrose em hepatopatias concomitantes (HCV, álcool, DHGNA) e transmite HBV por transfusão/transplante (dose infectante mínima ~3 UI/mL). Limitações incluem falta de padronização entre ensaios (ELISA, CMIA, CLIA produzem unidades diferentes: UI/mL, INH%, S/O) e dificuldade em distinguir baixos títulos verdadeiros de falso-positivos em testes ultrassensíveis.',

  patient_explanation = 'O anti-HBc é um anticorpo que seu organismo produz contra uma parte central do vírus da hepatite B, chamada "core" (núcleo). Diferente de outros marcadores que podem desaparecer com o tempo, o anti-HBc permanece detectável para sempre após o contato com o vírus, funcionando como uma "memória imunológica" permanente dessa exposição. Existem dois tipos principais: o anti-HBc IgM aparece durante infecções agudas (recentes) e desaparece em 6-9 meses, enquanto o anti-HBc IgG surge depois e persiste por toda a vida, indicando que você já teve contato com o vírus algum dia (pode ser infecção atual, passada ou curada). Quando o resultado mostra anti-HBc positivo isolado (sem outros marcadores positivos), significa que você teve exposição ao vírus no passado, mas existem várias possibilidades: pode ter se curado completamente (perdendo outros anticorpos com o tempo), estar em uma fase especial chamada "janela imunológica", ou ter uma infecção oculta com níveis muito baixos de vírus que não aparecem nos testes comuns. Em alguns casos raros, o resultado pode ser um falso-positivo. A dosagem quantitativa (medição da quantidade) de anti-HBc tem ganhado importância recente: níveis mais altos indicam que seu sistema imunológico está mais ativo no combate ao vírus, correlacionando-se com maior inflamação no fígado. Se você está planejando fazer tratamentos que diminuem a imunidade (quimioterapia, corticoides em doses altas, medicamentos para doenças autoimunes), ter anti-HBc positivo é uma informação crucial: existe risco do vírus "acordar" (reativar), especialmente com medicamentos como rituximabe (anti-CD20), onde esse risco chega a 40%. Nesses casos, seu médico pode recomendar medicamentos preventivos (antivirais profiláticos) para evitar reativação. É importante entender que ter anti-HBc positivo não significa necessariamente doença ativa ou risco de transmissão, mas requer interpretação cuidadosa junto com outros exames (HBsAg, anti-HBs, carga viral) para definir sua situação real e necessidade de acompanhamento ou tratamento.',

  conduct = 'A interpretação do anti-HBc exige análise integrada com demais marcadores sorológicos em painel completo ("triple panel": HBsAg, anti-HBs, anti-HBc), conforme diretrizes CDC 2025. PADRÕES SOROLÓGICOS PRINCIPAIS: (1) HBsAg+/anti-HBc IgM+/anti-HBs-: infecção aguda, solicitar HBV DNA quantitativo, ALT, AST, bilirrubinas, INR, encaminhar para hepatologia urgente se ALT >10x LSN ou INR elevado; (2) HBsAg+/anti-HBc IgG+/anti-HBs-/HBeAg±: infecção crônica, estratificar por HBV DNA, elastografia/fibroscan (F3-F4 indica cirrose), considerar biópsia se incerteza diagnóstica, avaliar critérios para tratamento (ALT elevado persistente, DNA >2000 UI/mL, fibrose significativa); (3) HBsAg-/anti-HBc+/anti-HBs+: infecção resolvida (imunidade natural), sem necessidade de seguimento exceto se imunossupressão planejada (ver protocolo abaixo); (4) HBsAg-/anti-HBc+/anti-HBs-: ANTI-HBC ISOLADO - padrão mais complexo, exige investigação adicional. PROTOCOLO PARA ANTI-HBC ISOLADO: (a) Repetir sorologia em 3-6 meses (pode ser janela imunológica com posterior surgimento de anti-HBs); (b) Solicitar anti-HBc IgM (se positivo, sugere infecção aguda na janela); (c) Considerar HBV DNA quantitativo se: paciente imunossuprimido, história epidemiológica significativa (origem de área endêmica, contato com portadores, PWID, HSH), ALT elevado sem causa aparente, coinfecção HIV/HCV; (d) Pacientes HIV+ com anti-HBc isolado: HBV DNA detectável em 1-10%, mas rotineiramente não recomendado testar; (e) Candidatos a transplante/quimioterapia: sempre solicitar HBV DNA ultrassensível para descartar OBI. ESTRATIFICAÇÃO DE RISCO PARA REATIVAÇÃO EM IMUNOSSUPRESSÃO (anti-HBc positivos/HBsAg negativos): ALTO RISCO (>10%): rituximabe/ofatumumabe/outros anti-CD20, transplante alogênico de células-tronco com condicionamento mieloablativo - CONDUTA: profilaxia obrigatória com entecavir 0,5mg/dia ou tenofovir 300mg/dia iniciando antes da imunossupressão, manter até 18 meses após término; monitorar HBsAg e ALT mensalmente nos primeiros 6 meses, depois trimestralmente. RISCO MODERADO (1-10%): quimioterapia citotóxica (antracíclicos, antraciclinas), corticoides ≥20mg/dia prednisona por ≥4 semanas, inibidores TNF-α (infliximabe, adalimumabe), transplante alogênico não-mieloablativo - CONDUTA: monitoramento pré-emptivo preferível (HBsAg + HBV DNA quantitativo mensais); se HBsAg positivar ou DNA >2000 UI/mL, iniciar tratamento imediato; considerar profilaxia se monitoramento frequente inviável. BAIXO RISCO (<1%): antivirais diretos HCV (DAAs), imunossupressores moderados (azatioprina, metotrexato, micofenolato), corticoides <20mg/dia - CONDUTA: monitoramento trimestral HBsAg e ALT durante tratamento e 6 meses pós. DOSAGEM QUANTITATIVA DE ANTI-HBC (quando disponível): valores >4,5 log10 UI/mL correlacionam-se com inflamação hepática moderada-grave, útil para decisão de biópsia em pacientes ALT-normal; >4,0-4,5 log10 UI/mL predizem melhor resposta a peginterferon/análogos nucleos(t)ídeos; acompanhar tendências (elevação sugere atividade imune aumentada). INVESTIGAÇÃO DE OBI (infecção oculta): indicada em anti-HBc isolado se: receptores de transplante hepático (doador ou receptor), doadores de sangue/órgãos, hepatopatia crônica de etiologia incerta (ALT persistentemente elevado), crioglobulinemia, HCC em não-cirrótico - solicitar HBV DNA ultrassensível (limite detecção <10 UI/mL) em múltiplas amostras (viremia intermitente), considerar biópsia hepática com PCR tecidual como padrão-ouro se alta suspeita clínica. SEGUIMENTO LONGITUDINAL: pacientes com infecção resolvida (anti-HBc+/anti-HBs+) não necessitam monitoramento exceto pré-imunossupressão; anti-HBc isolado sem fatores de risco: repetir sorologia anual por 2-3 anos, depois alta se estável; histórico de HBV com anti-HBs <10 mUI/mL: considerar revacinação (3 doses), reavaliar resposta sorológica. Documentar sempre status sorológico em prontuário eletrônico para consulta rápida em emergências ou pré-procedimentos.',

  last_review = NOW(),
  updated_at = NOW()
WHERE id = 'b3cb69e0-98fd-48ca-82ae-9eee62a8218e';

-- Link articles to score item (get article IDs from inserted records)
INSERT INTO article_score_items (score_item_id, article_id)
SELECT
  'b3cb69e0-98fd-48ca-82ae-9eee62a8218e',
  id
FROM articles
WHERE title IN (
  'Chronic hepatitis B in 2025: diagnosis, treatment and future directions',
  'Hepatitis B Core Antibody Level: A Surrogate Marker for Host Antiviral Immunity in Chronic Hepatitis B Virus Infections',
  'Clinical Significance and Remaining Issues of Anti-HBc Antibody and HBV Core-Related Antigen',
  'Occult Hepatitis B Virus Infection: An Update'
)
AND created_at >= NOW() - INTERVAL '1 minute';

COMMIT;

-- Verification query
SELECT
  si.name,
  ss.name as subgroup,
  sg.name as score_group,
  LENGTH(si.clinical_relevance) as clinical_relevance_chars,
  LENGTH(si.patient_explanation) as patient_explanation_chars,
  LENGTH(si.conduct) as conduct_chars,
  si.last_review,
  COUNT(DISTINCT asi.article_id) as linked_articles_count
FROM score_items si
JOIN score_subgroups ss ON si.subgroup_id = ss.id
JOIN score_groups sg ON ss.group_id = sg.id
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = 'b3cb69e0-98fd-48ca-82ae-9eee62a8218e'
GROUP BY si.id, si.name, ss.name, sg.name, si.clinical_relevance, si.patient_explanation, si.conduct, si.last_review;
