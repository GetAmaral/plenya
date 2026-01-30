-- Enrichment for IGF-1 (31-50 anos) score item
-- Score Item ID: bfadf0e5-4df7-4e70-9ed7-772a237a03fd
-- Generated: 2026-01-29

BEGIN;

-- Insert scientific articles (ignore if already exist)
INSERT INTO articles (id, title, authors, journal, publish_date, doi, original_link, article_type, notes, created_at, updated_at)
VALUES
(
    gen_random_uuid(),
    'Serum Insulin-Like Growth Factor-1 (IGF-1) Age-Specific Reference Values for Healthy Adult Population of Serbia',
    'Stojanovic M, Popevic M, Pekic S, Doknic M, Miljic D, Medic-Stojanoska M, Topalov D, Stojanovic J, Milovanovic A, Petakov M, Damjanovic S, Popovic V',
    'Acta Endocrinologica (Bucharest)',
    '2021-01-01',
    '10.4183/aeb.2021.462',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC9206165/',
    'research_article',
    'Estudo normativo com 1.200 adultos saudáveis estabelecendo valores de referência específicos por idade para IGF-1. Demonstrou declínio progressivo mais proeminente na faixa de 21-50 anos, com mediana de 183 ng/mL (31-35 anos) reduzindo para 129 ng/mL (46-50 anos). Mulheres apresentaram valores médios significativamente maiores que homens na faixa de 31-46 anos. O método de transformação matemática LMS mostrou-se superior para padronização dessas medidas idade-dependentes.',
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'Long-Term IGF-1 Maintenance in the Upper-Normal Range Has Beneficial Effect on Low-Grade Inflammation Marker in Adults with Growth Hormone Deficiency',
    'Klinc A, Janež A, Jensterle M',
    'International Journal of Molecular Sciences',
    '2025-01-01',
    '10.3390/ijms26052010',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC11900236/',
    'research_article',
    'Estudo longitudinal de 5 anos com 31 adultos em reposição de GH demonstrando que manutenção de IGF-1 na faixa superior-normal associa-se com menores níveis de proteína C-reativa ultrassensível (0,8 vs 1,8 mg/L), indicando redução de inflamação sistêmica de baixo grau. Homens apresentaram maior probabilidade de atingir níveis superiores-normais de IGF-1 comparado a mulheres, destacando necessidade de titulação cuidadosa específica por sexo.',
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'The GH/IGF-1 axis in ageing and longevity',
    'Junnila RK, List EO, Berryman DE, Murrey JW, Kopchick JJ',
    'Nature Reviews Endocrinology',
    '2013-01-01',
    '10.1038/nrendo.2013.67',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC4074016/',
    'review',
    'Revisão abrangente demonstrando relação complexa entre eixo GH/IGF-1 e envelhecimento. Múltiplos modelos animais com atividade reduzida de GH/IGF-1 apresentam extensão significativa de longevidade (21-68%). Em humanos com resistência ao GH (síndrome de Laron), observa-se redução substancial de incidência de câncer. A somatopausa (declínio relacionado à idade no GH) pode representar processo protetor ao invés de patológico, sugerindo que a redução fisiológica de IGF-1 na meia-idade pode ter benefícios adaptativos.',
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'Reference Intervals for Insulin-like Growth Factor-1 (IGF-I) From Birth to Senescence: Results From a Multicenter Study Using a New Automated Chemiluminescence IGF-I Immunoassay',
    'Bidlingmaier M, Friedrich N, Emeny RT, Spranger J, Wolthers OD, Roswall J, Körner A, Obermayer-Pietsch B, Hübener C, Dahlgren J, Frystyk J, Pfeiffer AF, Döring A, Bielohuby M, Wallaschofski H, Arafat AM',
    'Journal of Clinical Endocrinology & Metabolism',
    '2014-05-01',
    '10.1210/jc.2013-3059',
    'https://academic.oup.com/jcem/article/99/5/1712/2537423',
    'research_article',
    'Estudo multicêntrico com 15.014 indivíduos (0-94 anos) estabelecendo intervalos de referência padronizados para IGF-1 usando imunoensaio automatizado conforme recomendações internacionais. Para adultos 31-50 anos, demonstrou declínio progressivo idade-dependente com valores médios reduzindo aproximadamente 50% desde pico pós-puberal. Estudo fundamental para padronização de interpretação clínica de IGF-1 em diferentes faixas etárias, enfatizando necessidade de ajuste por idade e sexo.',
    NOW(),
    NOW()
)
ON CONFLICT (doi) DO NOTHING;

-- Get the article IDs for linking (using the DOIs to identify them)
WITH article_ids AS (
    SELECT id, doi FROM articles
    WHERE doi IN (
        '10.4183/aeb.2021.462',
        '10.3390/ijms26052010',
        '10.1038/nrendo.2013.67',
        '10.1210/jc.2013-3059'
    )
)
-- Link articles to score item (ignore if already linked)
INSERT INTO article_score_items (score_item_id, article_id)
SELECT
    'bfadf0e5-4df7-4e70-9ed7-772a237a03fd'::uuid,
    id
FROM article_ids
ON CONFLICT (score_item_id, article_id) DO NOTHING;

-- Update score item with clinical content
UPDATE score_items
SET
    clinical_relevance = 'O IGF-1 (Fator de Crescimento Semelhante à Insulina tipo 1), também denominado somatomedina-C, é o principal mediador periférico dos efeitos do hormônio do crescimento (GH) e biomarcador fundamental para avaliação do eixo somatotrófico em adultos de meia-idade. Na faixa etária de 31-50 anos, observa-se o declínio fisiológico mais proeminente e progressivo dos níveis séricos de IGF-1, fenômeno denominado somatopausa, com redução de aproximadamente 50% desde os valores de pico pós-puberal. Os intervalos de referência específicos para esta faixa etária variam de 7,0-31,7 nmol/L (homens) e 6,4-31,0 nmol/L (mulheres), ou aproximadamente 90-280 ng/mL para 31-35 anos, reduzindo progressivamente para 81-263 ng/mL aos 46-50 anos. Este declínio relacionado à idade não deve ser confundido com deficiência patológica de GH (DGH), que cursa com níveis de IGF-1 reduzidos em mais de 90% da secreção esperada. A dosagem de IGF-1 apresenta importância clínica significativa como teste de triagem para DGH em adultos, embora valores isolados não sejam diagnósticos, requerendo testes de estímulo de GH para confirmação. Valores elevados sustentados podem indicar acromegalia, enquanto valores reduzidos associam-se não apenas a DGH, mas também a desnutrição, doença hepática, hipotireoidismo e resistência ao GH. A interpretação deve considerar fatores modificadores como sexo (mulheres na faixa 31-46 anos apresentam valores maiores), obesidade, uso de estrogênio oral e estado nutricional. Estudos recentes demonstram que manutenção de IGF-1 na faixa superior-normal durante reposição de GH associa-se com redução de marcadores inflamatórios (PCR-us), sugerindo benefícios cardiometabólicos, embora a somatopausa fisiológica possa ter efeitos protetores contra câncer e extensão de longevidade, destacando a complexidade da interpretação clínica deste biomarcador na meia-idade.',

    patient_explanation = 'O IGF-1 (Fator de Crescimento Semelhante à Insulina) é um hormônio produzido principalmente pelo fígado em resposta ao hormônio do crescimento liberado pela hipófise. Ele desempenha papel fundamental não apenas no crescimento durante a infância, mas também na manutenção da saúde metabólica, muscular e óssea na vida adulta. Entre 31 e 50 anos de idade, é completamente normal que os níveis de IGF-1 diminuam progressivamente - este é um processo natural chamado somatopausa, semelhante à menopausa, mas relacionado ao hormônio do crescimento. Esta redução faz parte do envelhecimento saudável e pode até ter efeitos protetores contra alguns tipos de câncer. Os valores normais nesta faixa etária variam aproximadamente de 90-280 ng/mL aos 31-35 anos, reduzindo para 81-263 ng/mL aos 46-50 anos, sendo que mulheres costumam ter valores um pouco maiores que homens até os 46 anos. Valores muito baixos podem indicar deficiência do hormônio do crescimento, problemas nutricionais, doenças do fígado ou tireoide, causando sintomas como fadiga, perda de massa muscular, ganho de gordura abdominal e redução da qualidade de vida. Valores muito altos, por outro lado, podem sugerir produção excessiva de hormônio do crescimento (acromegalia), que requer investigação adicional. É importante ressaltar que o resultado do IGF-1 deve sempre ser interpretado considerando sua idade, sexo, estado nutricional e outros exames, não sendo suficiente isoladamente para diagnóstico. Seu médico pode solicitar testes adicionais se os valores estiverem fora da faixa esperada para sua idade.',

    conduct = 'INTERPRETAÇÃO CLÍNICA: 1) Ajustar valores de IGF-1 por idade e sexo utilizando tabelas de referência específicas ou cálculo de escore-desvio-padrão (SDS) com fórmula: SDS = (valor medido - média para idade/sexo) / desvio-padrão. Valores entre -2 e +2 SDS são considerados normais. 2) IGF-1 reduzido (<-2 SDS ou abaixo do limite inferior para idade): considerar deficiência de GH em adultos (DGH) se contexto clínico apropriado (cirurgia hipofisária prévia, radioterapia craniana, tumor hipofisário, trauma cranioencefálico), mas IGF-1 isolado tem sensibilidade limitada (≈50%) para DGH. Valores baixos também podem indicar desnutrição proteico-calórica, cirrose hepática, hipotireoidismo não tratado, diabetes mellitus descompensado, resistência ao GH (síndrome de Laron) ou uso de estrogênio oral. 3) IGF-1 elevado (>+2 SDS ou acima do limite superior): investigar acromegalia através de dosagem de GH basal e teste de supressão com sobrecarga oral de glicose (TOTG), onde falha em suprimir GH <1 ng/mL confirma diagnóstico. Causas não patológicas de elevação incluem puberdade, gravidez e uso de estrogênio. CONDUTA PARA DGH SUSPEITA: Realizar teste de estímulo de GH (hipoglicemia insulínica, teste de glucagon ou GHRH+arginina) - pico de GH <3-5 ng/mL confirma DGH grave. Solicitar ressonância magnética de sela túrcica para avaliar patologia hipofisária. Avaliar outros eixos hipofisários (TSH, ACTH, LH/FSH, prolactina). MONITORAMENTO DE REPOSIÇÃO DE GH: Titular dose para atingir IGF-1 SDS entre 0 e +2 (faixa superior-normal), associado com melhores desfechos cardiometabólicos e menor inflamação sistêmica. Mulheres frequentemente requerem doses maiores que homens. Reavaliar IGF-1 a cada 3-6 meses durante titulação, posteriormente anualmente. CONSIDERAÇÕES ESPECIAIS PARA 31-50 ANOS: Declínio fisiológico de IGF-1 não requer intervenção. Reposição de GH em adultos sem DGH documentada não é recomendada devido riscos (intolerância à glicose, edema, artralgia) e ausência de benefícios comprovados. Valores na faixa inferior-normal podem ser apropriados e protetores nesta idade. INVESTIGAÇÃO COMPLEMENTAR: Dosagem de IGFBP-3 (proteína ligadora) pode auxiliar interpretação. Avaliar estado nutricional (albumina, pré-albumina), função hepática (AST, ALT, bilirrubinas) e função tireoidiana (TSH, T4 livre). Considerar contexto clínico completo incluindo composição corporal, força muscular, densidade óssea e qualidade de vida antes de iniciar investigação extensa ou terapia.',

    last_review = CURRENT_DATE,
    updated_at = NOW()
WHERE id = 'bfadf0e5-4df7-4e70-9ed7-772a237a03fd';

COMMIT;

-- Verification query
SELECT
    si.id,
    si.name,
    LENGTH(si.clinical_relevance) as clinical_relevance_chars,
    LENGTH(si.patient_explanation) as patient_explanation_chars,
    LENGTH(si.conduct) as conduct_chars,
    si.last_review,
    COUNT(asi.article_id) as linked_articles_count
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = 'bfadf0e5-4df7-4e70-9ed7-772a237a03fd'
GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct, si.last_review;
