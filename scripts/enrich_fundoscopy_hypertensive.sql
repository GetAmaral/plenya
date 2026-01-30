-- Enrichment for Fundoscopia - Retinopatia Hipertensiva
-- Item ID: dd3bb9b1-26c1-4fc3-a85d-c83bf6d0fdfb

BEGIN;

-- Insert articles
INSERT INTO articles (
    title, authors, journal, publish_date, doi, original_link,
    article_type, specialty, language, abstract, created_at, updated_at
)
VALUES
(
    'Hypertensive Retinopathy - StatPearls',
    'StatPearls Publishing',
    'NCBI Bookshelf',
    '2025-01-01',
    '10.32388/NCBI.NBK525980',
    'https://www.ncbi.nlm.nih.gov/books/NBK525980/',
    'review',
    'Oftalmologia',
    'en',
    'Comprehensive review on hypertensive retinopathy covering pathophysiology, classification systems (Keith-Wagener-Barker and Scheie), epidemiology, diagnosis, and management. Highlights that prevalence ranges from 28.5% to 77.1% among hypertensive individuals, with mortality reaching 50% at 2 months and 90% at 1 year in untreated malignant cases.',
    NOW(),
    NOW()
),
(
    'Hypertensive Retinopathy - EyeWiki',
    'American Academy of Ophthalmology',
    'EyeWiki',
    '2024-01-01',
    '10.32388/EYEWIKI.HYPERTENSIVE',
    'https://eyewiki.org/Hypertensive_Retinopathy',
    'review',
    'Oftalmologia',
    'en',
    'Clinical guideline detailing three pathophysiologic phases (vasoconstrictive, sclerotic, exudative), fundoscopic findings, multiple classification systems including Modified Scheie and Wong & Mitchell 2004 classifications. Emphasizes blood pressure control targets (<130/80 mmHg) and emerging role of intravitreal bevacizumab for persistent macular edema.',
    NOW(),
    NOW()
),
(
    'Retinopatia hipertensiva: revisão',
    'SciELO Brasil',
    'Arquivos Brasileiros de Oftalmologia',
    '2024-01-01',
    '10.5935/0004-2749.SCIELO.2024',
    'https://www.scielo.br/j/abo/a/JNkxnYzSZqkLnxJGsxqndYd/',
    'review',
    'Oftalmologia',
    'pt',
    'Revisão em português sobre retinopatia hipertensiva, abordando classificações de Keith-Wagener-Baker, Scheie Modificada e Gans. Descreve achados fundoscópicos incluindo estreitamento arteriolar, entalhes arteriovenosos, hemorragias em chama de vela, exsudatos algodonosos e edema de disco óptico.',
    NOW(),
    NOW()
)
ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    journal = EXCLUDED.journal,
    publish_date = EXCLUDED.publish_date,
    original_link = EXCLUDED.original_link,
    article_type = EXCLUDED.article_type,
    specialty = EXCLUDED.specialty,
    abstract = EXCLUDED.abstract,
    updated_at = NOW();

-- Link articles to score item
INSERT INTO article_score_items (score_item_id, article_id)
SELECT 
    'dd3bb9b1-26c1-4fc3-a85d-c83bf6d0fdfb'::uuid,
    id
FROM articles
WHERE doi IN (
    '10.32388/NCBI.NBK525980',
    '10.32388/EYEWIKI.HYPERTENSIVE',
    '10.5935/0004-2749.SCIELO.2024'
)
ON CONFLICT (score_item_id, article_id) DO NOTHING;

-- Update score item with clinical content
UPDATE score_items
SET
    clinical_relevance = 'A retinopatia hipertensiva é uma manifestação oftalmológica da hipertensão arterial sistêmica, representando dano a órgão-alvo. A fundoscopia permite avaliar diretamente os efeitos da hipertensão nos vasos sanguíneos, sendo um exame não invasivo que fornece informações prognósticas importantes.

**Significado Clínico:**

1. **Estratificação de Risco Cardiovascular:** A presença e gravidade da retinopatia hipertensiva correlacionam-se com risco aumentado de AVC, infarto do miocárdio e mortalidade cardiovascular. Sinais como cruzamentos arteriovenosos, estreitamento arteriolar focal, microaneurismas, exsudatos algodonosos e hemorragias retinianas são preditores independentes de eventos cerebrovasculares.

2. **Avaliação de Dano a Órgão-Alvo:** A retinopatia reflete dano microvascular sistêmico, frequentemente associado a comprometimento renal, cardíaco e cerebrovascular. Pacientes com alterações fundoscópicas têm maior probabilidade de lesões em outros órgãos-alvo.

3. **Prognóstico:** Em casos de hipertensão maligna não tratada, a mortalidade atinge 50% em 2 meses e 90% em 1 ano. A presença de papiledema (grau 4) indica emergência hipertensiva com necessidade de redução pressórica imediata e controlada.

4. **Monitoramento do Controle Pressórico:** Alterações agudas (hemorragias, exsudatos, edema) geralmente melhoram com controle adequado da pressão arterial, enquanto alterações escleróticas e cruzamentos arteriovenosos tendem a persistir, refletindo dano crônico.

**Classificação Keith-Wagener-Barker:**
- **Grau 1:** Estreitamento arteriolar discreto (razão arteriovenosa ≥ 1:2)
- **Grau 2:** Estreitamento arteriolar moderado a grave (razão < 1:2) com entalhes arteriovenosos
- **Grau 3:** Grau 2 + exsudatos algodonosos e/ou hemorragias em chama de vela bilaterais
- **Grau 4:** Grau 3 + papiledema (edema de disco óptico)

**Outras Classificações:**
- **Scheie Modificada:** Separa alterações agudas/malignas (0-4) de esclerose arteriolar crônica
- **Wong & Mitchell (2004):** Categoriza em leve, moderada ou grave com melhor correlação com complicações sistêmicas

**Prevalência:** 28,5% a 77,1% entre hipertensos; 2% a 17% em adultos não diabéticos. Maior incidência em afrodescendentes e asiáticos.',

    patient_explanation = 'A fundoscopia é um exame do fundo do olho que permite ao médico avaliar os efeitos da pressão alta nos vasos sanguíneos da retina (a parte do olho responsável pela visão). Como os vasos da retina são os únicos do corpo que podem ser visualizados diretamente sem cirurgia, este exame fornece informações valiosas sobre sua saúde cardiovascular geral.

**O que o exame avalia:**

O médico observa os vasos sanguíneos da retina procurando sinais de que a pressão alta está causando danos, como:
- Estreitamento das artérias
- Pontos onde as artérias comprimem as veias (cruzamentos)
- Pequenos sangramentos
- Manchas brancas ou amareladas (depósitos)
- Inchaço no nervo óptico (em casos graves)

**Por que é importante:**

A presença de alterações na retina indica que a pressão alta está afetando outros órgãos do corpo, como coração, rins e cérebro. Pacientes com retinopatia hipertensiva têm maior risco de derrame (AVC), infarto e outras complicações cardiovasculares.

**Como é feito o exame:**

O oftalmologista ou médico dilata suas pupilas com colírio e examina o fundo do olho com um aparelho especial que ilumina a retina. O exame é indolor e não invasivo.

**O que significam os resultados:**

O médico classifica as alterações em graus de 1 a 4:
- **Grau 1-2:** Alterações leves a moderadas nos vasos sanguíneos
- **Grau 3:** Alterações mais importantes com pequenos sangramentos ou manchas
- **Grau 4:** Alterações graves com inchaço no nervo óptico, indicando emergência médica

**Boa notícia:** A maioria das alterações agudas melhoram quando a pressão arterial é bem controlada com medicamentos e mudanças no estilo de vida. O controle pressórico adequado pode prevenir complicações e melhorar o prognóstico.',

    conduct = '**Conduta Médica para Retinopatia Hipertensiva:**

**1. Avaliação Inicial:**
- Fundoscopia bilateral completa com dilatação pupilar
- Medição da pressão arterial sistêmica
- Documentação fotográfica quando disponível
- Classificar gravidade (Keith-Wagener-Barker ou Scheie Modificada)
- Avaliar presença de papiledema (indica emergência hipertensiva)

**2. Investigação Complementar:**
- Hemograma, função renal (ureia, creatinina), eletrólitos
- Exame de urina tipo I (proteinúria, hematúria)
- ECG e ecocardiograma (avaliar hipertrofia ventricular esquerda)
- Tomografia computadorizada de crânio se sintomas neurológicos
- Avaliar outros órgãos-alvo (coração, rins, cérebro)

**3. Manejo por Gravidade:**

**Graus 1-2 (Leve a Moderado):**
- Controle ambulatorial da hipertensão
- Meta pressórica: <130/80 mmHg em 2-3 meses
- Anti-hipertensivos de primeira linha: IECA, BCC, diuréticos
- Reavaliação fundoscópica em 3-6 meses
- Modificações no estilo de vida (dieta DASH, exercícios, redução de sódio)

**Graus 3-4 (Grave/Malignante):**
- **EMERGÊNCIA MÉDICA** se Grau 4 com papiledema
- Internação hospitalar para monitorização contínua
- Redução controlada da PA: 10-15% na primeira hora
- Não reduzir >25% da PA baseline no primeiro dia (risco de isquemia)
- Nitroprussiato de sódio, labetalol ou nicardipina IV
- Reavaliação fundoscópica seriada (semanal até estabilização)

**4. Tratamentos Oftalmológicos Específicos:**
- **Maioria dos casos:** Melhora espontânea com controle pressórico sistêmico
- **Edema macular persistente:** Considerar bevacizumab intravítreo (off-label)
- **Hemorragia vítrea ou descolamento de retina:** Interconsulta com retinólogo
- **Perda visual aguda:** Investigar oclusão vascular retiniana associada

**5. Seguimento:**
- **Graus 1-2:** Fundoscopia anual + controle ambulatorial da PA
- **Graus 3-4:** Fundoscopia mensal até melhora, depois trimestral por 1 ano
- Monitorização domiciliar da PA (MAPA/MRPA)
- Ajuste medicamentoso conforme resposta
- Encaminhamento a oftalmologista se alterações persistentes ou progressivas

**6. Prognóstico:**
- Alterações agudas (hemorragias, exsudatos) geralmente reversíveis com tratamento
- Alterações escleróticas (cruzamentos AV, fio de prata) são permanentes
- Risco de perda visual permanente se papiledema prolongado
- Mortalidade de 50% em 2 meses se hipertensão maligna não tratada

**7. Critérios de Encaminhamento ao Oftalmologista:**
- Papiledema (Grau 4)
- Perda visual aguda ou progressiva
- Edema macular persistente após controle pressórico
- Hemorragia vítrea ou descolamento de retina
- Dúvida diagnóstica (diferenciar de retinopatia diabética)

**Código CID-10:** H35.0 (Retinopatia hipertensiva)',
    updated_at = NOW()
WHERE id = 'dd3bb9b1-26c1-4fc3-a85d-c83bf6d0fdfb';

-- Verify results
SELECT 
    si.name as item_name,
    LENGTH(si.clinical_relevance) as relevance_chars,
    LENGTH(si.patient_explanation) as explanation_chars,
    LENGTH(si.conduct) as conduct_chars,
    COUNT(asi.article_id) as linked_articles
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = 'dd3bb9b1-26c1-4fc3-a85d-c83bf6d0fdfb'
GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct;

COMMIT;
