-- ========================================
-- ENRICHMENT: Hepatite C - Anti-HCV
-- Score Item ID: 73411e66-c180-46ad-8f5b-7d67b26ef877
-- Generated: 2026-01-29
-- ========================================

BEGIN;

-- ========================================
-- PART 1: Insert Scientific Articles
-- ========================================

-- Article 1: Global HCV Elimination Progress 2024
INSERT INTO scientific_articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    doi,
    pmid,
    url,
    article_type,
    key_findings,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'A 2024 global report on national policies, programmes, and progress towards hepatitis C elimination: findings from 33 hepatitis elimination profiles',
    'Hiebert-Suwondo L, Manning J, Tohme RA, Buti M, Kondili LA, Spearman CW, Hajarizadeh B, Turnier V, Lazarus JV, Grebely J, Dore GJ, Waked I, Ward JW',
    'Lancet Gastroenterol Hepatol',
    '2025-05-20',
    '10.1016/S2468-1253(25)00068-8',
    NULL,
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC12308979/',
    'review',
    'Análise de 33 países representando 73% das pessoas vivendo com hepatite C globalmente. Demonstrou que apenas 10 países (30%) atingiram a meta da OMS de 60% de cobertura diagnóstica até 2025, e apenas 5 países (15%) atingiram a meta de 50% de tratamento. Países de alta renda apresentaram escores médios de políticas de 20.1 versus 14.7 em países de renda baixa/média-baixa. Destaca que 12 países (36%) ainda mantêm restrições parciais ou totais na autoridade de prescrição para não-especialistas, limitando o acesso ao tratamento.',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
) RETURNING id AS article_1_id \gset

-- Article 2: HCV Epidemiology and Elimination Strategies
INSERT INTO scientific_articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    doi,
    pmid,
    url,
    article_type,
    key_findings,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Hepatitis C Virus: Epidemiological Challenges and Global Strategies for Elimination',
    'Toma D, Anghel L, Patraș D, Ciubără A',
    'Viruses',
    '2025-07-31',
    '10.3390/v17081069',
    NULL,
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC12390683/',
    'review',
    'Revisão sistemática demonstrando que aproximadamente 1 milhão de novas infecções por HCV ocorrem anualmente, com 242.000 mortes relacionadas em 2022. Os antivirais de ação direta (DAAs) revolucionaram o tratamento, oferecendo taxas de cura superiores a 95% com excelente perfil de segurança, permitindo clearance viral em 8-12 semanas. Identifica uso de drogas injetáveis como principal via de transmissão atual. Prevalência global de HCV em populações carcerárias estimada em 17.7%. Programas bem-sucedidos no Egito, Espanha e Austrália demonstram viabilidade de atingir as metas da OMS para 2030.',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
) RETURNING id AS article_2_id \gset

-- Article 3: SASLT Guidelines Update 2024
INSERT INTO scientific_articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    doi,
    pmid,
    url,
    article_type,
    key_findings,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'SASLT guidelines: Update in treatment of hepatitis C virus infection, 2024',
    'Alghamdi AS, Alghamdi H, Alserehi HA, Babatin MA, Alswat KA, et al.',
    'Saudi J Gastroenterol',
    '2024-01-03',
    '10.4103/sjg.sjg_333_23',
    NULL,
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC10856511/',
    'clinical_guideline',
    'Diretrizes atualizadas da Associação Saudita para Estudo do Fígado enfatizando que regimes pangenotípicos de DAAs são recomendados como primeira opção de tratamento para todos os indivíduos com infecção crônica por HCV, simplificando protocolos terapêuticos. Todos os pacientes com infecção ativa por HCV devem receber tratamento sem demora, com urgência particular para aqueles com fibrose avançada ou cirrose. Pacientes que alcançam resposta virológica sustentada com cirrose requerem vigilância contínua para carcinoma hepatocelular a cada seis meses. Avaliação pré-tratamento abrangente incluindo estadiamento de fibrose, testes para HBV/HIV e triagem de interações medicamentosas é essencial.',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
) RETURNING id AS article_3_id \gset

-- Article 4: Cure of chronic HCV after DAA treatment
INSERT INTO scientific_articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    doi,
    pmid,
    url,
    article_type,
    key_findings,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Cure of chronic hepatitis C virus infection after DAA treatment only partially restores the functional capacity of exhausted T cell subsets: a systematic review',
    'Authors from Frontiers in Immunology',
    'Front Immunol',
    '2025-01-15',
    '10.3389/fimmu.2025.1546915',
    NULL,
    'https://www.frontiersin.org/journals/immunology/articles/10.3389/fimmu.2025.1546915/full',
    'systematic_review',
    'Revisão sistemática demonstrando que a cura da infecção crônica pelo HCV após tratamento com DAAs resulta em taxas de cura superiores a 95%. No entanto, a restauração da capacidade funcional de subconjuntos de células T exaustas é apenas parcial após o clearance viral. Mesmo após alcançar resposta virológica sustentada (SVR), alterações imunológicas persistem, com implicações para compreensão da recuperação imune pós-tratamento. Destaca a importância do tratamento precoce antes do estabelecimento de dano imunológico irreversível.',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
) RETURNING id AS article_4_id \gset

-- ========================================
-- PART 2: Link Articles to Score Item
-- ========================================

INSERT INTO score_item_articles (score_item_id, article_id, created_at, updated_at)
VALUES
    ('73411e66-c180-46ad-8f5b-7d67b26ef877', :'article_1_id', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('73411e66-c180-46ad-8f5b-7d67b26ef877', :'article_2_id', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('73411e66-c180-46ad-8f5b-7d67b26ef877', :'article_3_id', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('73411e66-c180-46ad-8f5b-7d67b26ef877', :'article_4_id', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- ========================================
-- PART 3: Update Score Item Clinical Content
-- ========================================

UPDATE score_items
SET
    clinical_relevance = 'O teste anti-HCV é a ferramenta de triagem essencial para detecção de infecção pelo vírus da hepatite C, detectando anticorpos produzidos em resposta à exposição ao HCV. A hepatite C crônica representa um problema global de saúde pública, com aproximadamente 58 milhões de pessoas vivendo com infecção crônica e cerca de 1 milhão de novas infecções anuais, resultando em 242.000 mortes em 2022. A importância clínica da detecção e tratamento está fundamentada na revolução terapêutica proporcionada pelos antivirais de ação direta (DAAs), que alcançam taxas de cura superiores a 95% em 8-12 semanas de tratamento, comparado aos antigos esquemas à base de interferon. O anti-HCV positivo indica exposição ao vírus, mas não diferencia infecção ativa de infecção resolvida, necessitando confirmação com HCV RNA. A OMS estabeleceu meta de eliminar a hepatite C como ameaça à saúde pública até 2030, visando diagnosticar 90% dos infectados e tratar 80% dos diagnosticados. Populações de risco incluem usuários de drogas injetáveis (principal via de transmissão atual), pessoas em situação de cárcere (prevalência de 17.7%), profissionais de saúde expostos a material biológico, receptores de hemoderivados antes de 1992, e pacientes submetidos a procedimentos invasivos sem adequada esterilização. A triagem universal em adultos está recomendada em diversos países. O tratamento precoce previne progressão para cirrose, carcinoma hepatocelular e insuficiência hepática, além de interromper a cadeia de transmissão. Regimes pangenotípicos simplificaram dramaticamente o manejo, permitindo que provedores não-especialistas prescrevam tratamento, embora 36% dos países ainda mantenham restrições. Pacientes com cirrose que alcançam resposta virológica sustentada (SVR) requerem vigilância contínua para hepatocarcinoma a cada seis meses, pois o risco persiste mesmo após cura virológica.',

    patient_explanation = 'O exame anti-HCV pesquisa no sangue anticorpos contra o vírus da hepatite C, uma infecção que afeta o fígado. Quando você entra em contato com o vírus da hepatite C, seu corpo produz anticorpos (proteínas de defesa) para tentar combatê-lo, e este exame detecta esses anticorpos. Um resultado positivo significa que você teve contato com o vírus em algum momento da vida, mas não indica se a infecção está ativa no momento ou se já foi eliminada naturalmente pelo organismo. Por isso, quando o anti-HCV dá positivo, é necessário fazer um segundo exame chamado HCV RNA (carga viral) para confirmar se o vírus ainda está presente no corpo. A hepatite C é transmitida principalmente através do contato com sangue contaminado, como compartilhamento de agulhas, seringas ou outros materiais perfurocortantes, procedimentos médicos ou odontológicos sem esterilização adequada, tatuagens ou piercings feitos sem material descartável, e raramente por relações sexuais ou de mãe para filho. A boa notícia é que a hepatite C tem cura! Os medicamentos modernos, chamados antivirais de ação direta (DAAs), são tomados por via oral durante 8 a 12 semanas e curam mais de 95% dos casos, com poucos efeitos colaterais. Estes medicamentos revolucionaram o tratamento, eliminando a necessidade das antigas injeções de interferon que causavam muitos efeitos adversos. Se não tratada, a hepatite C pode evoluir silenciosamente durante anos ou décadas, causando cirrose (cicatrizes no fígado), câncer de fígado e insuficiência hepática. Por isso a detecção precoce e o tratamento adequado são fundamentais para prevenir complicações graves e interromper a transmissão para outras pessoas.',

    conduct = 'INTERPRETAÇÃO DO ANTI-HCV: Anti-HCV negativo geralmente exclui infecção, exceto em período de janela imunológica (até 12 semanas após exposição) ou em imunossuprimidos graves. Anti-HCV positivo requer confirmação obrigatória com HCV RNA qualitativo ou quantitativo, pois aproximadamente 15-25% das infecções agudas são resolvidas espontaneamente, deixando anticorpos detectáveis sem viremia ativa. CONDUTA ANTI-HCV POSITIVO: 1) Solicitar imediatamente HCV RNA quantitativo (carga viral) para confirmar infecção ativa; 2) Avaliar função hepática completa (AST, ALT, GGT, FA, bilirrubinas, albumina, TAP/INR); 3) Hemograma completo e creatinina; 4) Sorologias para hepatite B (HBsAg, anti-HBs, anti-HBc) e HIV, dada possibilidade de coinfecções; 5) Genotipagem do HCV (embora regimes pangenotípicos sejam preferidos, pode ser útil em situações específicas); 6) Estadiamento de fibrose hepática por elastografia transitória (Fibroscan) ou scores não-invasivos (FIB-4, APRI) - biópsia hepática raramente necessária. HCV RNA POSITIVO (INFECÇÃO ATIVA): Encaminhar para tratamento sem demora, priorizando pacientes com fibrose avançada (F3-F4), cirrose, manifestações extra-hepáticas, coinfecção HIV/HBV, transplante de órgãos, ou risco de transmissão. TRATAMENTO: Regimes pangenotípicos de DAAs são primeira linha para todos os pacientes: sofosbuvir/velpatasvir por 12 semanas (sem cirrose) ou glecaprevir/pibrentasvir por 8-12 semanas são opções padrão. Avaliar interações medicamentosas rigorosamente antes de iniciar (anticonvulsivantes, rifampicina, erva-de-são-joão, inibidores de bomba de prótons em altas doses podem reduzir eficácia). MONITORAMENTO EM TRATAMENTO: HCV RNA na semana 4 (opcional) e obrigatoriamente 12 semanas após término do tratamento (SVR12) para confirmar cura. Função hepática e hemograma em intervalos regulares. RESPOSTA VIROLÓGICA SUSTENTADA (SVR12): HCV RNA indetectável 12 semanas após fim do tratamento indica cura em >99% dos casos. Pacientes com cirrose mantêm risco de hepatocarcinoma e requerem vigilância com ultrassom abdominal + alfa-fetoproteína a cada 6 meses indefinidamente. Pacientes sem cirrose podem receber alta da hepatologia após confirmação de SVR, mantendo seguimento clínico para fatores de risco metabólicos. HCV RNA NEGATIVO (INFECÇÃO RESOLVIDA): Documentar resolução espontânea ou pós-tratamento. Orientar sobre risco de reinfecção se exposição persistir. JANELA IMUNOLÓGICA: Em caso de exposição recente de alto risco com anti-HCV negativo, considerar HCV RNA direto ou repetir anti-HCV após 12 semanas.',

    last_review = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE id = '73411e66-c180-46ad-8f5b-7d67b26ef877';

COMMIT;

-- ========================================
-- VERIFICATION QUERY
-- ========================================
SELECT
    si.id,
    si.name,
    si.score_level_id,
    LENGTH(si.clinical_relevance) as clinical_relevance_chars,
    LENGTH(si.patient_explanation) as patient_explanation_chars,
    LENGTH(si.conduct) as conduct_chars,
    si.last_review,
    COUNT(sia.article_id) as linked_articles_count
FROM score_items si
LEFT JOIN score_item_articles sia ON si.id = sia.score_item_id
WHERE si.id = '73411e66-c180-46ad-8f5b-7d67b26ef877'
GROUP BY si.id, si.name, si.score_level_id, si.clinical_relevance, si.patient_explanation, si.conduct, si.last_review;
