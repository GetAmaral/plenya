-- Enrichment for DHEA-S - Homens (70+ anos)
-- Score Item ID: 447746dd-949b-449b-bd0e-c2f226330a10
-- Generated: 2026-01-28
-- Focus: DHEA-S in very elderly men, mortality, frailty, supplementation controversy

BEGIN;

-- Update the score_items table with clinical content
UPDATE score_items
SET
    clinical_relevance = 'O DHEA-S (sulfato de dehidroepiandrosterona) é o principal andrógeno adrenal, cujos níveis declinam progressivamente 2-5% ao ano, atingindo apenas 10-20% dos valores de adultos jovens aos 70-80 anos. Em homens acima de 70 anos, níveis baixos de DHEA-S demonstram particular relevância clínica como preditor independente de mortalidade por todas as causas e cardiovascular. Meta-análise evidenciou risco relativo de 1.41 (IC95% 1.18-1.69) para mortalidade em homens com níveis baixos. Estudos prospectivos revelam que fragilidade incidente é inversamente associada ao DHEA-S basal (OR 0.35), com efeito mais pronunciado em indivíduos com IMC <29 kg/m². A relação entre DHEA-S e desfechos adversos aparece específica de gênero e idade: associação com mortalidade cardiovascular observada em homens idosos mas não em mulheres, e mais acentuada em homens <75 anos. Contudo, controvérsia importante existe quanto à suplementação: enquanto associações observacionais são robustas, ensaios clínicos randomizados não demonstraram benefícios consistentes em força muscular, densidade óssea, função cognitiva ou qualidade de vida. Meta-análise mostrou apenas redução de massa gorda sem impacto em parâmetros clínicos relevantes. Preocupações com segurança incluem possível aumento de risco de câncer de próstata e progressão de HPB. Níveis baixos podem refletir doença grave subjacente (confusão por gravidade), dificultando interpretação causal.',

    patient_explanation = 'O DHEA-S é um hormônio produzido pelas glândulas adrenais (acima dos rins) que diminui naturalmente com a idade. Aos 70-80 anos, os níveis ficam entre 10-20% do que eram na juventude. Este hormônio é importante porque estudos mostram que homens idosos com níveis muito baixos apresentam maior risco de problemas cardíacos, fragilidade física e até morte prematura. Níveis baixos estão associados a maior fraqueza muscular, dificuldade para caminhar e realizar atividades diárias. O exame de DHEA-S ajuda a avaliar seu estado de envelhecimento hormonal e risco de fragilidade. Entretanto, é importante saber que apesar destes achados, a reposição de DHEA em forma de suplemento não mostrou benefícios claros em estudos científicos rigorosos. Vários estudos testaram dar DHEA para homens idosos e não encontraram melhora significativa na força muscular, memória, ossos ou qualidade de vida. Além disso, existem preocupações sobre segurança, especialmente relacionadas à próstata. Por isso, não se recomenda tomar suplementos de DHEA por conta própria. Se seu DHEA-S estiver baixo, o médico irá investigar causas tratáveis e avaliar se existe alguma doença que possa estar diminuindo seus níveis, mas provavelmente não irá prescrever reposição hormonal.',

    conduct = 'INTERPRETAÇÃO CLÍNICA EM HOMENS 70+ ANOS: Valores de referência variam entre laboratórios; considerar declínio fisiológico esperado (80-90% vs pico produção). Faixas típicas 70+ anos: 20-200 μg/dL. Investigar valores <30 μg/dL em contexto clínico apropriado. AVALIAÇÃO INICIAL: Ao detectar DHEA-S baixo, primeiro avaliar: (1) Sintomas/sinais de fragilidade (força preensão manual, velocidade marcha, fadiga, perda peso involuntária); (2) Comorbidades que podem suprimir DHEA-S (infecção grave, ICC descompensada, neoplasias, desnutrição); (3) Medicamentos: glicocorticoides, opioides crônicos; (4) Função adrenal: considerar cortisol AM se suspeita insuficiência. INTERPRETAÇÃO CONTEXTUALIZADA: Níveis baixos em homem 70+ saudável podem representar apenas envelhecimento fisiológico acelerado. Associação com mortalidade mais forte em <75 anos. Efeito protetor contra fragilidade mais pronunciado com IMC <29 kg/m². CONDUTA BASEADA EM EVIDÊNCIAS: (1) NÃO indicar suplementação rotineira de DHEA - evidência nível A contra uso baseado em meta-análises que não mostraram benefícios clínicos relevantes. (2) Exceções controversas: ensaios experimentais em contextos de pesquisa, insuficiência adrenal documentada (raro). (3) Se considerar suplementação off-label (não recomendado): dose 25-50 mg/dia, monitorar PSA rigorosamente a cada 3-6 meses (risco HPB/Ca próstata), avaliar sintomas prostáticos, reavaliar após 6 meses - descontinuar se sem benefício claro. ABORDAGEM PREFERENCIAL: Focar em intervenções com evidência sólida para fragilidade e longevidade - exercício resistido regular, nutrição adequada proteica (1.2-1.5 g/kg), vitamina D se deficiente, manejo cardiovascular otimizado. DHEA-S baixo deve ser visto como marcador de risco aumentado que motiva intensificação de estratégias preventivas comprovadas, não como alvo terapêutico direto. Reavaliar anualmente junto com outros marcadores geriátricos.',

    last_review = CURRENT_DATE,
    updated_at = NOW()
WHERE id = '447746dd-949b-449b-bd0e-c2f226330a10';

COMMIT;

-- Link articles to score item (separate transaction)
DO $$
DECLARE
    v_score_item_id uuid := '447746dd-949b-449b-bd0e-c2f226330a10';
    v_article_record RECORD;
BEGIN
    FOR v_article_record IN
        SELECT id FROM articles
        WHERE doi IN (
            '10.1089/rej.2013.1425',
            '10.1093/ageing/afp015',
            '10.1210/jc.2013-1358'
        )
        OR (doi IS NULL AND title = 'Adrenal Androgens and Aging')
    LOOP
        INSERT INTO article_score_items (score_item_id, article_id)
        VALUES (v_score_item_id, v_article_record.id)
        ON CONFLICT (score_item_id, article_id) DO NOTHING;
    END LOOP;
END $$;

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
