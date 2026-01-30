-- Enrichment for Score Item: Escroto / epidídimos
-- Score Item ID: e010b2a7-6e9e-4b42-9d37-487e91411a18
-- Generated: 2026-01-28

BEGIN;

-- Insert scientific articles
INSERT INTO articles (id, title, authors, journal, publish_date, article_type, doi, pm_id, original_link, abstract, language, created_at, updated_at)
VALUES
(
    gen_random_uuid(),
    'Epididymitis',
    'Rupp TJ, Leslie SW',
    'StatPearls',
    '2023-07-17',
    'review',
    NULL,
    '30855799',
    'https://www.ncbi.nlm.nih.gov/books/NBK430814/',
    'Revisão abrangente sobre epididimite, a causa mais comum de dor escrotal aguda em adultos. Aborda epidemiologia, fisiopatologia, diagnóstico diferencial, achados do exame físico (sensibilidade à palpação do epidídimo ao longo do aspecto posterior e superior do testículo), ultrassonografia com avaliação de fluxo vascular, e manejo terapêutico. Enfatiza a importância de excluir torção testicular como emergência cirúrgica. Incidência anual superior a 600.000 casos nos EUA, com pico entre 20-39 anos.',
    'en',
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'Testicular Torsion',
    'Schick MA, Sternard BT',
    'StatPearls',
    '2023-06-12',
    'review',
    NULL,
    '29262062',
    'https://www.ncbi.nlm.nih.gov/books/NBK448199/',
    'Revisão sobre torção testicular, emergência urológica tempo-dependente. A viabilidade testicular diminui significativamente após 6 horas do início dos sintomas. Achados físicos incluem testículo em posição transversa ou elevada. Ultrassonografia com Doppler apresenta sensibilidade de 93% e especificidade de 100%. Taxa de salvamento aproxima-se de 100% quando a intervenção cirúrgica ocorre dentro de 6 horas, mas cai para menos de 50% após 12-24 horas. Mais comum em pacientes abaixo de 25 anos.',
    'en',
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'Scrotal Masses',
    'Langan RC, Puente ME',
    'American Family Physician',
    '2022-08-15',
    'review',
    NULL,
    NULL,
    'https://www.aafp.org/pubs/afp/issues/2022/0800/scrotal-masses.html',
    'Revisão clínica sobre massas escrotais, dividindo condições em dolorosas (torção testicular, torção de apêndice testicular, epididimite) e indolores (hidrocele, varicocele, câncer testicular). Sistema TWIST para estratificação de risco de torção (escore ≥5 indica necessidade de exploração cirúrgica urgente). Epididimite mais comum por Chlamydia trachomatis ou Neisseria gonorrhoeae em homens sexualmente ativos de 14-35 anos. Câncer testicular é tumor sólido mais comum em homens de 15-34 anos, apresentando-se como nódulo unilateral firme.',
    'en',
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'Standards for scrotal ultrasonography',
    'Tyloch JF, Wieczorek AP',
    'Journal of Ultrasonography',
    '2016-12-01',
    'research_article',
    '10.15557/JoU.2016.0040',
    '28138411',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC5269526/',
    'Artigo estabelecendo padrões para ultrassonografia escrotal. Transdutores lineares de alta frequência (≥7 MHz) são essenciais, com capacidades de Doppler colorido e espectral. Avaliação modo-B precisa incluindo volumetria testicular. Na avaliação epididimária, quando aumentado, documentar tamanho da cabeça no plano sagital, corpo e cauda. Manobra de Valsalva é crítica no diagnóstico de varicocele, produzindo aumento do sinal Doppler e reversão do fluxo sanguíneo. Medição volumétrica de hidrocele recomendada pois pode associar-se a tumores embrionários.',
    'en',
    NOW(),
    NOW()
);

-- Update score item with clinical content
UPDATE score_items
SET
    clinical_relevance = 'O exame do escroto e epidídimos é componente essencial da avaliação urológica masculina, permitindo detecção precoce de condições que variam desde processos inflamatórios benignos até emergências cirúrgicas e neoplasias. A epididimite representa a causa mais comum de dor escrotal aguda em adultos, afetando mais de 600.000 homens anualmente nos EUA, com pico entre 20-39 anos. O exame físico revela sensibilidade à palpação do epidídimo ao longo do aspecto posterior e superior do testículo, frequentemente acompanhada de edema escrotal, eritema cutâneo e possível hidrocele reativa. Diferenciar epididimite de torção testicular é crítico, pois esta última constitui emergência urológica com viabilidade testicular decrescente após 6 horas (taxa de salvamento ~90% em 6h, caindo para 50% em 12h e 10% em 24h). Achados sugestivos de torção incluem testículo em posição elevada ou transversa, dor súbita intensa, náuseas e vômitos. Massas escrotais dividem-se em dolorosas (torção, epididimite) e indolores (hidrocele, varicocele, câncer testicular). Varicocele acomete 15% dos homens adultos, apresentando-se como estruturas venosas serpiginosas palpáveis que aumentam com Valsalva. Hidrocele manifesta-se como coleção fluida entre camadas da túnica vaginal. Câncer testicular, tumor sólido mais comum em homens de 15-34 anos, tipicamente apresenta-se como nódulo unilateral firme. A ultrassonografia com Doppler colorido é modalidade diagnóstica de escolha, com sensibilidade de 93-100% e especificidade de 97-100% para torção testicular, além de permitir caracterização de massas, avaliação de fluxo vascular e diferenciação entre lesões císticas e sólidas.',

    patient_explanation = 'O exame do escroto (bolsa que contém os testículos) e dos epidídimos (pequenas estruturas em formato de vírgula localizadas atrás de cada testículo que armazenam e transportam espermatozoides) é fundamental para detectar problemas que podem afetar sua saúde reprodutiva e geral. Durante o exame, o médico palpa cuidadosamente o escroto procurando alterações como inchaço, nódulos, dor ou mudanças de temperatura. Várias condições podem ser identificadas: a epididimite é uma inflamação do epidídimo que causa dor e inchaço, geralmente tratada com antibióticos; a varicocele são veias dilatadas no escroto que podem parecer um "saco de minhocas" e ocasionalmente afetam a fertilidade; a hidrocele é acúmulo de líquido ao redor do testículo que causa aumento indolor do escroto; a torção testicular é emergência médica onde o testículo torce sobre si mesmo, cortando o suprimento sanguíneo e causando dor súbita intensa - requer cirurgia urgente para salvar o testículo. Nódulos ou massas podem indicar desde cistos benignos até câncer testicular, mais comum em homens jovens mas altamente tratável quando detectado precocemente. Sintomas que exigem avaliação urgente incluem dor escrotal súbita e intensa, inchaço rápido, febre com dor escrotal, ou descoberta de nódulo ou massa. O autoexame testicular mensal, realizado após banho quente quando o escroto está relaxado, ajuda a familiarizar-se com a anatomia normal e detectar mudanças precocemente. Mantenha comunicação aberta com seu médico sobre qualquer alteração notada.',

    conduct = 'AVALIAÇÃO INICIAL: Anamnese detalhada investigando início, duração e caráter da dor (súbita sugere torção, gradual sugere epididimite), sintomas urinários (disúria, urgência), corrimento uretral, febre, trauma, atividade sexual recente, histórico de DSTs. Examinar paciente em posição supina e ortostática em ambiente aquecido. TÉCNICA DE EXAME: Inspeção visual do escroto avaliando assimetria, eritema, edema, lesões cutâneas, posição testicular (testículo elevado/transverso sugere torção). Palpação bilateral comparativa começando pelo lado não afetado: avaliar cada testículo quanto a tamanho (4-5cm comprimento), consistência (firme-elástica normal), superfície (lisa), presença de massas ou nódulos. Palpar epidídimo (localizado póstero-superiormente ao testículo) avaliando sensibilidade, espessamento, nódulos. Avaliar cordão espermático quanto a torção, espessamento, massas. Testar reflexo cremastérico (ausente em 100% das torções em adultos, mas menos confiável em crianças). Realizar manobra de Valsalva para avaliar varicocele (aumento de estruturas venosas palpáveis). Transiluminação para diferenciar lesões císticas (hidrocele, espermatocele) de sólidas. DIAGNÓSTICO DIFERENCIAL: Sistema TWIST para estratificação de risco de torção: edema testicular (2 pontos), testículo endurecido (2 pontos), ausência de reflexo cremastérico (1 ponto), náusea/vômito (1 ponto), testículo elevado (1 ponto) - escore ≥5 indica alto risco, requerer exploração cirúrgica urgente; escore ≤2 permite imagem inicial. INVESTIGAÇÃO COMPLEMENTAR: Ultrassonografia com Doppler colorido é padrão-ouro (transdutor linear ≥7MHz): avaliar ecotextura testicular, fluxo sanguíneo (diminuído/ausente em torção, aumentado em epididimite), medidas epididimárias quando aumentado (cabeça >10mm, corpo >4mm, cauda >5mm), caracterização de massas (císticas vs sólidas), presença de hidrocele (medir volume se >5mL). Urina I e urocultura se suspeita infecciosa. Teste para Chlamydia/Gonorreia em pacientes sexualmente ativos 14-35 anos. MANEJO: Torção testicular: exploração cirúrgica urgente (idealmente <6h do início); Epididimite: antibioticoterapia empírica (ceftriaxona 500mg IM dose única + doxiciclina 100mg VO 12/12h por 10 dias se <35 anos e sexualmente ativo; levofloxacino 500mg/dia por 10 dias se >35 anos), AINEs, repouso, elevação escrotal, reavaliação em 48-72h; Massas/nódulos testiculares: encaminhar urologia para avaliação adicional, considerar marcadores tumorais (AFP, β-hCG, LDH) se suspeita neoplásica; Hidrocele/Varicocele: observação se assintomática, encaminhar urologia se sintomática ou associada a infertilidade.',

    last_review = CURRENT_DATE,
    updated_at = NOW()
WHERE id = 'e010b2a7-6e9e-4b42-9d37-487e91411a18';

-- Link articles to score item
INSERT INTO article_score_items (score_item_id, article_id)
SELECT
    'e010b2a7-6e9e-4b42-9d37-487e91411a18',
    a.id
FROM articles a
WHERE a.original_link IN (
    'https://www.ncbi.nlm.nih.gov/books/NBK430814/',
    'https://www.ncbi.nlm.nih.gov/books/NBK448199/',
    'https://www.aafp.org/pubs/afp/issues/2022/0800/scrotal-masses.html',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC5269526/'
)
AND NOT EXISTS (
    SELECT 1 FROM article_score_items asi
    WHERE asi.score_item_id = 'e010b2a7-6e9e-4b42-9d37-487e91411a18'
    AND asi.article_id = a.id
);

COMMIT;

-- Verification query
SELECT
    si.name,
    si.slug,
    LENGTH(si.clinical_relevance) as clinical_relevance_length,
    LENGTH(si.patient_explanation) as patient_explanation_length,
    LENGTH(si.conduct) as conduct_length,
    si.last_review,
    COUNT(DISTINCT asi.article_id) as linked_articles_count
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = 'e010b2a7-6e9e-4b42-9d37-487e91411a18'
GROUP BY si.id, si.name, si.slug, si.clinical_relevance, si.patient_explanation, si.conduct, si.last_review;
