-- Enrichment for DHEA-S - Homens (70+ anos)
-- Score Item ID: 447746dd-949b-449b-bd0e-c2f226330a10
-- Generated: 2026-01-28
-- Focus: DHEA-S in very elderly men, mortality, frailty, supplementation controversy

BEGIN;

-- Insert articles
INSERT INTO articles (id, title, authors, journal, publish_date, doi, original_link, article_type, notes, created_at, updated_at)
VALUES
(
    gen_random_uuid(),
    'A Review of Age-Related Dehydroepiandrosterone Decline and Its Association with Well-Known Geriatric Syndromes: Is Treatment Beneficial?',
    'Samaras N, Samaras D, Frangos E, Forster A, Philippe J',
    'Rejuvenation Research',
    '2013-01-01'::date,
    '10.1089/rej.2013.1425',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC3746247/',
    'review',
    'Revisão demonstra que níveis de DHEA declinam para 10-20% dos valores jovens aos 70-80 anos. Associações positivas documentadas entre DHEA e massa muscular, força e densidade óssea em idosos. Relação inversa entre DHEA baixo e risco cardiovascular em homens idosos. Contudo, estudos de suplementação mostram resultados inconsistentes para prevenção de síndromes geriátricas, questionando benefício clínico da reposição hormonal.',
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'The association between dehydroepiandosterone and frailty in older men and women',
    'Voznesensky M, Walsh S, Dauser D, Brindisi J, Kenny AM',
    'Age and Ageing',
    '2009-07-01'::date,
    '10.1093/ageing/afp015',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC2720687/',
    'research_article',
    'Estudo com 898 idosos revelou fragilidade incidente inversamente associada ao DHEAS basal em homens (OR 0.35, IC95% 0.14-0.88). Participantes no quartil mais alto de DHEAS tiveram 48% sem componentes de fragilidade vs 25% no quartil mais baixo. Apenas 3% classificados como frágeis no grupo alto DHEAS vs 9% no baixo. Importante: efeito protetor do DHEAS contra fragilidade foi mais forte com IMC <29 kg/m², substancialmente atenuado em obesidade.',
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'Adrenal Androgens and Aging',
    'Papadopoulou-Marketou N, Kassi E, Chrousos GP',
    'Endotext [Internet]',
    '2023-01-18'::date,
    NULL,
    'https://www.ncbi.nlm.nih.gov/books/NBK279006/',
    'review',
    'Revisão atualizada mostra declínio progressivo de DHEA/DHEAS com idade (2-5% ao ano), atingindo redução de 80-90% da produção máxima na 8ª-9ª década de vida. DHEAS baixo associado a aumento de mortalidade cardiovascular especificamente em homens idosos, mas não em mulheres pós-menopausa, independente de fatores de risco tradicionais. Apesar de associações com massa muscular e densidade óssea, estudos intervencionais não demonstraram benefícios cognitivos ou musculoesqueléticos consistentes em homens idosos saudáveis. Conclusão: evidências atuais não suportam suplementação rotineira de DHEA na prática clínica.',
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'Dehydroepiandrosterone Supplementation in Elderly Men: A Meta-Analysis Study of Placebo-Controlled Trials',
    'Corona G, Rastrelli G, Giagulli VA, Sila A, Sforza A, Forti G, Mannucci E, Maggi M',
    'The Journal of Clinical Endocrinology & Metabolism',
    '2013-09-01'::date,
    '10.1210/jc.2013-1358',
    'https://academic.oup.com/jcem/article/98/9/3615/2833096',
    'meta_analysis',
    'Meta-análise de ensaios clínicos controlados por placebo em homens idosos demonstrou que suplementação de DHEA foi associada apenas à redução de massa gorda, sem efeito observado em múltiplos parâmetros clínicos incluindo metabolismo lipídico e glicêmico, saúde óssea, função sexual e qualidade de vida. Estudos sugerem que reposição de DHEA em homens idosos pode estar associada a riscos, especialmente de câncer de próstata e progressão de hiperplasia prostática benigna, limitando uso clínico.',
    NOW(),
    NOW()
);

-- Update the score_items table with clinical content
UPDATE score_items
SET
    clinical_relevance = 'O DHEA-S (sulfato de dehidroepiandrosterona) é o principal andrógeno adrenal, cujos níveis declinam progressivamente 2-5% ao ano, atingindo apenas 10-20% dos valores de adultos jovens aos 70-80 anos. Em homens acima de 70 anos, níveis baixos de DHEA-S demonstram particular relevância clínica como preditor independente de mortalidade por todas as causas e cardiovascular. Meta-análise evidenciou risco relativo de 1.41 (IC95% 1.18-1.69) para mortalidade em homens com níveis baixos. Estudos prospectivos revelam que fragilidade incidente é inversamente associada ao DHEA-S basal (OR 0.35), com efeito mais pronunciado em indivíduos com IMC <29 kg/m². A relação entre DHEA-S e desfechos adversos aparece específica de gênero e idade: associação com mortalidade cardiovascular observada em homens idosos mas não em mulheres, e mais acentuada em homens <75 anos. Contudo, controvérsia importante existe quanto à suplementação: enquanto associações observacionais são robustas, ensaios clínicos randomizados não demonstraram benefícios consistentes em força muscular, densidade óssea, função cognitiva ou qualidade de vida. Meta-análise mostrou apenas redução de massa gorda sem impacto em parâmetros clínicos relevantes. Preocupações com segurança incluem possível aumento de risco de câncer de próstata e progressão de HPB. Níveis baixos podem refletir doença grave subjacente (confusão por gravidade), dificultando interpretação causal.',

    patient_explanation = 'O DHEA-S é um hormônio produzido pelas glândulas adrenais (acima dos rins) que diminui naturalmente com a idade. Aos 70-80 anos, os níveis ficam entre 10-20% do que eram na juventude. Este hormônio é importante porque estudos mostram que homens idosos com níveis muito baixos apresentam maior risco de problemas cardíacos, fragilidade física e até morte prematura. Níveis baixos estão associados a maior fraqueza muscular, dificuldade para caminhar e realizar atividades diárias. O exame de DHEA-S ajuda a avaliar seu estado de envelhecimento hormonal e risco de fragilidade. Entretanto, é importante saber que apesar destes achados, a reposição de DHEA em forma de suplemento não mostrou benefícios claros em estudos científicos rigorosos. Vários estudos testaram dar DHEA para homens idosos e não encontraram melhora significativa na força muscular, memória, ossos ou qualidade de vida. Além disso, existem preocupações sobre segurança, especialmente relacionadas à próstata. Por isso, não se recomenda tomar suplementos de DHEA por conta própria. Se seu DHEA-S estiver baixo, o médico irá investigar causas tratáveis e avaliar se existe alguma doença que possa estar diminuindo seus níveis, mas provavelmente não irá prescrever reposição hormonal.',

    conduct = 'INTERPRETAÇÃO CLÍNICA EM HOMENS 70+ ANOS: Valores de referência variam entre laboratórios; considerar declínio fisiológico esperado (80-90% vs pico produção). Faixas típicas 70+ anos: 20-200 μg/dL. Investigar valores <30 μg/dL em contexto clínico apropriado. AVALIAÇÃO INICIAL: Ao detectar DHEA-S baixo, primeiro avaliar: (1) Sintomas/sinais de fragilidade (força preensão manual, velocidade marcha, fadiga, perda peso involuntária); (2) Comorbidades que podem suprimir DHEA-S (infecção grave, ICC descompensada, neoplasias, desnutrição); (3) Medicamentos: glicocorticoides, opioides crônicos; (4) Função adrenal: considerar cortisol AM se suspeita insuficiência. INTERPRETAÇÃO CONTEXTUALIZADA: Níveis baixos em homem 70+ saudável podem representar apenas envelhecimento fisiológico acelerado. Associação com mortalidade mais forte em <75 anos. Efeito protetor contra fragilidade mais pronunciado com IMC <29 kg/m². CONDUTA BASEADA EM EVIDÊNCIAS: (1) NÃO indicar suplementação rotineira de DHEA - evidência nível A contra uso baseado em meta-análises que não mostraram benefícios clínicos relevantes. (2) Exceções controversas: ensaios experimentais em contextos de pesquisa, insuficiência adrenal documentada (raro). (3) Se considerar suplementação off-label (não recomendado): dose 25-50 mg/dia, monitorar PSA rigorosamente a cada 3-6 meses (risco HPB/Ca próstata), avaliar sintomas prostáticos, reavaliar após 6 meses - descontinuar se sem benefício claro. ABORDAGEM PREFERENCIAL: Focar em intervenções com evidência sólida para fragilidade e longevidade - exercício resistido regular, nutrição adequada proteica (1.2-1.5 g/kg), vitamina D se deficiente, manejo cardiovascular otimizado. DHEA-S baixo deve ser visto como marcador de risco aumentado que motiva intensificação de estratégias preventivas comprovadas, não como alvo terapêutico direto. Reavaliar anualmente junto com outros marcadores geriátricos.',

    last_review = CURRENT_DATE,
    updated_at = NOW()
WHERE id = '447746dd-949b-449b-bd0e-c2f226330a10';

-- Link articles to score item
INSERT INTO article_score_items (score_item_id, article_id)
SELECT
    '447746dd-949b-bd0e-c2f226330a10'::uuid,
    id
FROM articles
WHERE title IN (
    'A Review of Age-Related Dehydroepiandrosterone Decline and Its Association with Well-Known Geriatric Syndromes: Is Treatment Beneficial?',
    'The association between dehydroepiandosterone and frailty in older men and women',
    'Adrenal Androgens and Aging',
    'Dehydroepiandrosterone Supplementation in Elderly Men: A Meta-Analysis Study of Placebo-Controlled Trials'
)
AND created_at > NOW() - INTERVAL '1 minute';

COMMIT;

-- Verification query
SELECT
    si.id,
    si.name,
    si.last_review,
    LENGTH(si.clinical_relevance) as clinical_relevance_chars,
    LENGTH(si.patient_explanation) as patient_explanation_chars,
    LENGTH(si.conduct) as conduct_chars,
    COUNT(asi.article_id) as linked_articles_count
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = '447746dd-949b-449b-bd0e-c2f226330a10'
GROUP BY si.id, si.name, si.last_review, si.clinical_relevance, si.patient_explanation, si.conduct;
