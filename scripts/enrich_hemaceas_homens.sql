-- Enriquecimento: Hemácias - Homens
-- ID: c04e1997-0846-4557-8990-baffb1f7542a

BEGIN;

-- 1. Inserir artigos científicos
INSERT INTO articles (id, title, url, source, created_at, updated_at)
VALUES
    (gen_random_uuid(), 'Normal and Abnormal Complete Blood Count With Differential', 'https://www.ncbi.nlm.nih.gov/books/NBK604207/', 'StatPearls - NCBI Bookshelf', NOW(), NOW()),
    (gen_random_uuid(), 'Investigation and management of erythrocytosis', 'https://pmc.ncbi.nlm.nih.gov/articles/PMC7829024/', 'CMAJ - Canadian Medical Association Journal', NOW(), NOW()),
    (gen_random_uuid(), 'Polycythemia Vera - Clinical Features and Diagnosis', 'https://www.ncbi.nlm.nih.gov/books/NBK557660/', 'StatPearls - NCBI Bookshelf', NOW(), NOW()),
    (gen_random_uuid(), 'Red Cell Indices - Clinical Significance', 'https://www.ncbi.nlm.nih.gov/books/NBK260/', 'Clinical Methods - NCBI Bookshook', NOW(), NOW()),
    (gen_random_uuid(), 'Erythrocytosis: Diagnosis and investigation', 'https://onlinelibrary.wiley.com/doi/10.1111/ijlh.14298', 'International Journal of Laboratory Hematology', NOW(), NOW())
ON CONFLICT (url) DO UPDATE
SET title = EXCLUDED.title,
    source = EXCLUDED.source,
    updated_at = NOW();

-- 2. Criar relacionamentos com score_item
INSERT INTO article_score_items (score_item_id, article_id)
SELECT
    'c04e1997-0846-4557-8990-baffb1f7542a'::uuid,
    a.id
FROM articles a
WHERE a.url IN (
    'https://www.ncbi.nlm.nih.gov/books/NBK604207/',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC7829024/',
    'https://www.ncbi.nlm.nih.gov/books/NBK557660/',
    'https://www.ncbi.nlm.nih.gov/books/NBK260/',
    'https://onlinelibrary.wiley.com/doi/10.1111/ijlh.14298'
)
ON CONFLICT (score_item_id, article_id) DO NOTHING;

-- 3. Atualizar score_item com conteúdo clínico
UPDATE score_items
SET
    clinical_relevance = 'A contagem de hemácias (eritrócitos) é um dos parâmetros fundamentais do hemograma completo, refletindo a capacidade de transporte de oxigênio do sangue. Em homens adultos saudáveis, valores normais situam-se entre 4,5-5,9 milhões/µL, com variações individuais dependentes de fatores como genética, altitude de residência, nível de condicionamento físico e tabagismo.

Valores elevados (policitemia/eritrocitose) podem indicar condições primárias como policitemia vera - caracterizada pela mutação JAK2 V617F presente em >95% dos casos - ou secundárias como hipóxia crônica, tumores produtores de eritropoietina, ou desidratação relativa. Os critérios diagnósticos da OMS 2022 estabelecem hemoglobina >165 g/L (hematócrito >49%) como threshold para investigação de policitemia em homens ao nível do mar.

Valores reduzidos sugerem anemia, classificada pela análise integrada com VCM (volume corpuscular médio) em microcítica (deficiência de ferro), normocítica (doença crônica, hemólise) ou macrocítica (deficiência de B12/folato). A interpretação requer correlação com hemoglobina, hematócrito e índices eritrocitários, considerando que alterações no volume plasmático (gravidez, desidratação) afetam valores relativos sem refletir mudanças absolutas na massa eritrocitária.

Evidências recentes (2024-2025) destacam a importância das alterações na deformabilidade e tempo de vida das hemácias em condições como COVID-19, doença falciforme e lesões de estocagem, cada vez mais reconhecidas em decisões diagnósticas e terapêuticas.',

    patient_explanation = 'As hemácias (glóbulos vermelhos) são as células do sangue responsáveis por levar oxigênio para todos os órgãos e tecidos do corpo. Este exame conta quantas hemácias existem em uma pequena amostra do seu sangue.

Em homens adultos, valores normais ficam geralmente entre 4,5 a 5,9 milhões de hemácias por microlitro de sangue. Esse número pode variar um pouco dependendo de onde você mora (pessoas em lugares altos têm mais hemácias), seu condicionamento físico, genética e se você fuma.

Hemácias AUMENTADAS podem indicar:
- Policitemia vera: condição onde a medula óssea produz hemácias demais
- Problemas respiratórios crônicos ou viver em grandes altitudes
- Desidratação (que concentra o sangue)
- Alguns tumores raros

Hemácias DIMINUÍDAS indicam anemia, que pode ser causada por:
- Falta de ferro (anemia ferropriva)
- Deficiência de vitaminas B12 ou ácido fólico
- Doenças crônicas (rins, inflamações)
- Perda de sangue ou destruição aumentada de hemácias

Importante: este valor sempre deve ser analisado junto com outros dados do hemograma (hemoglobina, hematócrito, tamanho das hemácias) para um diagnóstico correto. Seu médico irá avaliar o conjunto de resultados considerando seus sintomas e histórico clínico.',

    conduct = 'VALORES NORMAIS (4,5-5,9 milhões/µL em homens):
- Manter acompanhamento de rotina conforme indicação clínica
- Reforçar hábitos saudáveis: hidratação adequada, dieta balanceada, atividade física regular
- Avaliar contexto individual (altitude, tabagismo, condicionamento físico)

ERITROCITOSE/POLICITEMIA (>5,9 milhões/µL):
Investigação inicial:
- Repetir hemograma completo com reticulócitos
- Gasometria arterial (avaliar saturação de O2)
- Dosagem de eritropoietina sérica
- Pesquisa de mutação JAK2 V617F (se Hb >165 g/L ou Hct >49%)
- Ultrassom abdominal (avaliar rins, baço, fígado)

Critérios OMS 2022 para Policitemia Vera:
Maiores: (1) Hb >165 g/L ou Hct >49%, (2) Biópsia de medula com hipercelularidade, (3) JAK2 V617F positivo
Menor: Eritropoietina sérica baixa/normal-baixa
Diagnóstico: 3 critérios maiores OU 2 maiores + 1 menor

Conduta conforme etiologia:
- Policitemia vera: encaminhar para hematologia (flebotomias, antiagregação, eventual hidroxiureia)
- Secundária: tratar causa base (DPOC, apneia do sono, tumor)
- Relativa: reidratação, investigar causas de hemoconcentração

ANEMIA (<4,5 milhões/µL):
Classificar por VCM:
- Microcítica (VCM <80 fL): investigar ferro sérico, ferritina, capacidade de ligação do ferro, talassemia
- Normocítica (VCM 80-100 fL): avaliar função renal, doenças crônicas, reticulócitos (hemólise?)
- Macrocítica (VCM >100 fL): dosar vitamina B12, ácido fólico, TSH, considerar mielograma se persistente

Conduta adicional:
- Sintomas (fadiga, dispneia, palidez): urgência na investigação
- Hemoglobina <70 g/L: considerar transfusão e internação
- Anemia crônica leve-moderada: investigação ambulatorial, reposição conforme deficiência identificada

MONITORAMENTO:
- Eritrocitose: hemograma a cada 2-3 meses até estabilização
- Anemia: repetir após 4-6 semanas de tratamento específico
- Avaliar sempre em conjunto: hemoglobina, hematócrito, VCM, HCM, RDW
- Atenção a sintomas novos: cefaleias, prurido aquagênico (policitemia), fadiga extrema, taquicardia (anemia)',

    last_review = NOW(),
    updated_at = NOW()
WHERE id = 'c04e1997-0846-4557-8990-baffb1f7542a';

-- 4. Verificar resultado
SELECT
    'Hemácias - Homens' as item_name,
    LENGTH(clinical_relevance) as clinical_length,
    LENGTH(patient_explanation) as patient_length,
    LENGTH(conduct) as conduct_length,
    last_review,
    (SELECT COUNT(*) FROM article_score_items WHERE score_item_id = 'c04e1997-0846-4557-8990-baffb1f7542a') as articles_count
FROM score_items
WHERE id = 'c04e1997-0846-4557-8990-baffb1f7542a';

COMMIT;

-- Mensagem de sucesso
\echo '✅ Enriquecimento completo: Hemácias - Homens'
