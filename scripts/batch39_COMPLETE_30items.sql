-- BATCH 39 - COMPOSIÇÃO CORPORAL PARTE 1: ENRICHMENT COMPLETO
-- 30 items com conteúdo clínico profundo baseado em evidências 2025-2026
-- Data: 2026-01-27
-- Referências: EWGSOP2, AWGS, Meta-análises 2025 (creatina+HMB), DEXA vs BIA

BEGIN;

-- ========================================
-- CATEGORIA 1: % GORDURA CORPORAL - HOMEM (2 items)
-- IDs: 299e22bd, 9ff7a160
-- ========================================

UPDATE score_items
SET
    clinical_relevance = E'O percentual de gordura corporal (BF%) representa a proporção de massa adiposa em relação ao peso corporal total e é um marcador fundamental de composição corporal e risco cardiometabólico. Em homens, valores funcionais ideais variam conforme idade: 8-19% (20-39 anos), 11-21% (40-59 anos) e 13-24% (60-79 anos). Atletas frequentemente mantêm 6-13%. A mensuração por DEXA é considerada padrão-ouro com erro de precisão de 1.5%, enquanto bioimpedância (BIA) oferece correlações fortes (ρ=0.94) mas pode superestimar massa livre de gordura em 7-8kg em indivíduos com sobrepeso.\n\nFisiologicamente, o tecido adiposo não é mero depósito energético mas órgão endócrino ativo secretando adipocinas (leptina, adiponectina, resistina). BF% elevado associa-se a inflamação crônica de baixo grau (elevação de IL-6, TNF-α, PCR), resistência insulínica, disfunção endotelial e aumento de risco cardiovascular. A distribuição da gordura (visceral vs subcutânea) é crítica: gordura visceral tem maior atividade metabólica e correlação com síndrome metabólica.\n\nNa medicina funcional, valores "normais" podem ser subótimos. Homens acima de 40 anos com BF% >25% apresentam risco aumentado para diabetes tipo 2, doenças cardiovasculares e declínio cognitivo, mesmo com IMC normal. Obesidade sarcopênica (BF% elevado + baixa massa muscular) afeta 30% dos indivíduos >60 anos e associa-se a fragilidade, quedas e mortalidade aumentada.',

    patient_explanation = E'Seu percentual de gordura corporal indica quanto do seu peso total é composto por gordura. Valores ideais variam com a idade: homens jovens (20-39 anos) devem ter entre 8-19%, homens de meia-idade (40-59 anos) entre 11-21%, e homens mais velhos (60-79 anos) entre 13-24%. Atletas geralmente ficam entre 6-13%.\n\nTer gordura em excesso aumenta riscos de diabetes, doenças do coração, pressão alta e inflamação crônica no corpo. A gordura não é apenas "energia armazenada" - ela produz hormônios e substâncias que afetam todo seu metabolismo. Gordura na região abdominal (visceral) é especialmente prejudicial à saúde metabólica.\n\nMedir sua composição corporal regularmente ajuda a acompanhar se suas intervenções (dieta, exercícios, suplementos) estão funcionando. DEXA é o método mais preciso, mas bioimpedância é prática para acompanhamento. Lembre-se: o objetivo não é apenas ter "peso normal" no IMC, mas ter composição corporal saudável com massa muscular adequada e gordura controlada.',

    conduct = E'**VALORES FUNCIONAIS ÓTIMOS (HOMENS):**\n- 20-39 anos: 8-19% (ideal funcional: 10-15%)\n- 40-59 anos: 11-21% (ideal funcional: 12-18%)\n- 60-79 anos: 13-24% (ideal funcional: 15-22%)\n- Atletas: 6-13%\n\n**AVALIAÇÃO COMPLEMENTAR:**\n- BF% >25%: Solicitar DEXA para distribuição gordura visceral/subcutânea\n- BF% elevado + circunferência abdominal >102cm: Investigar síndrome metabólica (glicemia, insulina, perfil lipídico, HbA1c, HOMA-IR)\n- Suspeita obesidade sarcopênica: Avaliar ASMI e força de preensão manual\n\n**INTERVENÇÕES BASEADAS EM EVIDÊNCIAS:**\n\n*Nutricionais:*\n- Déficit calórico moderado (300-500 kcal/dia) preservando massa magra\n- Proteína elevada: 1.6-2.2 g/kg/dia distribuída em 25-30g por refeição\n- Redução carboidratos refinados, aumento fibras (>30g/dia)\n- Considerar jejum intermitente (16:8) se aderência adequada\n\n*Exercício Físico:*\n- **Prioridade:** Treinamento resistido 3-4x/semana (60-80% 1RM, 8-12 semanas)\n- HIIT 2-3x/semana (reduz gordura visceral)\n- Aeróbico moderado 150-300min/semana\n\n*Suplementação:*\n- Creatina monohidratada 3-5g/dia (preserva massa magra durante déficit calórico)\n- Proteína whey 20-30g pós-treino\n- Ômega-3 2-4g/dia (EPA+DHA, reduz inflamação)\n- Considerar HMB 3g/dia se >60 anos ou resistência a perda de peso\n\n**MONITORAMENTO:**\n- Reavaliação BF% a cada 8-12 semanas\n- Meta realista: -0.5 a 1% BF por mês\n- Evitar perda rápida (>1kg/semana) que compromete massa muscular\n- Acompanhar biomarcadores: glicemia, insulina, perfil lipídico, PCR',

    last_review = NOW()
WHERE id IN (
    '299e22bd-184e-4513-8f21-26f26d91d737',
    '9ff7a160-8ad1-4c2a-857a-53677355a631'
);

-- Verificação
SELECT '✓ Atualizados: % Gordura Homem (' || COUNT(*) || ' items)' as status
FROM score_items
WHERE id IN ('299e22bd-184e-4513-8f21-26f26d91d737', '9ff7a160-8ad1-4c2a-857a-53677355a631')
AND LENGTH(clinical_relevance) > 100;

-- ========================================
-- CATEGORIA 2: % GORDURA CORPORAL - MULHER (2 items)
-- IDs: 29ec8df2, 35f8405e
-- ========================================

UPDATE score_items
SET
    clinical_relevance = E'O percentual de gordura corporal feminino é naturalmente superior ao masculino devido a diferenças hormonais (estrogênio promove depósito adiposo subcutâneo) e reprodutivas (reservas energéticas para gestação/lactação). Valores funcionais ideais variam por idade: 22-33% (20-39 anos), 24-34% (40-59 anos) e 25-36% (60-79 anos). Atletas mantêm 14-20%. Valores abaixo de 12-14% comprometem função reprodutiva (amenorreia hipotalâmica) e saúde óssea (baixo estrogênio aumenta reabsorção óssea).\n\nA distribuição da gordura corporal em mulheres é influenciada por estágios hormonais. Na pré-menopausa, predomínio de gordura subcutânea ginóide (quadris/coxas) com menor risco metabólico. Na pós-menopausa, queda estrogênica promove redistribuição para padrão andróide (abdominal/visceral), aumentando risco cardiometabólico significativamente.\n\nObesidade sarcopênica pós-menopausa afeta até 40% das mulheres >65 anos, associando-se a fraturas por fragilidade (osteoporose + quedas), dependência funcional e mortalidade aumentada. BF% elevado + baixo ASMI (<5.0 kg/m²) caracterizam a condição. Circunferência abdominal >88cm indica acúmulo visceral crítico mesmo com BF% "normal".\n\nA avaliação por DEXA é especialmente importante em mulheres pós-menopausa para quantificar densidade mineral óssea simultaneamente. Bioimpedância subestima gordura em 4-5kg comparado a DEXA em mulheres com obesidade, limitando uso diagnóstico mas sendo útil para monitoramento de tendências.',

    patient_explanation = E'Mulheres naturalmente têm mais gordura corporal que homens devido a hormônios e funções reprodutivas. Valores ideais variam com idade: mulheres jovens (20-39 anos) devem ter 22-33%, meia-idade (40-59 anos) 24-34%, e acima de 60 anos 25-36%. Atletas geralmente ficam entre 14-20%. Ter gordura muito baixa (<12-14%) pode afetar menstruação e saúde dos ossos.\n\nA localização da gordura muda com a menopausa. Antes, a gordura fica mais nos quadris e coxas (padrão feminino) e é menos prejudicial. Depois da menopausa, a queda do estrogênio faz a gordura se acumular na barriga (padrão masculino), aumentando riscos de diabetes, doenças cardíacas e inflamação.\n\nApós a menopausa, é comum desenvolver "obesidade sarcopênica" - gordura alta com perda de músculos. Isso aumenta risco de quedas, fraturas e dependência. Por isso é crucial combinar controle do percentual de gordura com manutenção da massa muscular através de exercícios resistidos e proteína adequada na dieta.\n\nMedir regularmente ajuda a identificar mudanças precoces e ajustar seu plano. DEXA é ideal pois mede gordura E densidade óssea juntos - importante após menopausa. Bioimpedância é prática para acompanhamento mensal.',

    conduct = E'**VALORES FUNCIONAIS ÓTIMOS (MULHERES):**\n- 20-39 anos: 22-33% (ideal funcional: 22-28%)\n- 40-59 anos: 24-34% (ideal funcional: 24-30%)\n- 60-79 anos: 25-36% (ideal funcional: 25-32%)\n- Atletas: 14-20%\n- **ATENÇÃO:** <14% pode causar amenorreia e perda óssea\n\n**AVALIAÇÃO COMPLEMENTAR:**\n- BF% >33% + circunferência abdominal >88cm: Investigar síndrome metabólica completa\n- Pós-menopausa com BF% >30%: DEXA obrigatório (avaliar DMO + distribuição gordura)\n- Suspeita obesidade sarcopênica: ASMI + força preensão + velocidade marcha\n- BF% <18% com amenorreia: Avaliar eixo hipotálamo-hipófise-ovariano, densitometria óssea\n\n**INTERVENÇÕES BASEADAS EM EVIDÊNCIAS:**\n\n*Nutricionais:*\n- Déficit calórico moderado 300-500 kcal/dia (evitar <1200 kcal/dia)\n- Proteína elevada: 1.2-1.6 g/kg/dia (1.5-2.0 g/kg se pós-menopausa)\n- Cálcio 1200mg/dia + Vitamina D 2000-4000 UI/dia (saúde óssea)\n- Fitoestrogênios (soja, linhaça) podem auxiliar composição corporal pós-menopausa\n\n*Exercício Físico:*\n- **Essencial:** Treinamento resistido 3-4x/semana com progressão de carga\n- Exercícios de impacto moderado (caminhada, dança) para saúde óssea\n- HIIT 2x/semana (reduz gordura visceral, melhora sensibilidade insulínica)\n- Pilates/yoga complementar para core e equilíbrio\n\n*Suplementação (Pós-Menopausa):*\n- Creatina 3-5g/dia (massa magra + força + cognição)\n- Colágeno hidrolisado 10g/dia (ossos + articulações)\n- Ômega-3 2-3g/dia (inflamação + risco CV)\n- Considerar HMB 3g/dia se resistência a ganho muscular\n- Vitamina K2 100-200mcg/dia (direciona cálcio para ossos)\n\n**CONSIDERAÇÕES ESPECIAIS:**\n- Pré-menopausa: Acompanhar ciclo menstrual, evitar déficit severo\n- Pós-menopausa: Foco em preservar massa magra, não apenas reduzir BF%\n- TRH (reposição hormonal): Pode facilitar manutenção composição corporal, discutir riscos/benefícios\n\n**MONITORAMENTO:**\n- Reavaliação BF% + DEXA a cada 8-12 semanas\n- Meta: -0.5% BF/mês preservando massa magra\n- Biomarcadores: glicemia, insulina, estradiol (se pré-menopausa), TSH, vitamina D',

    last_review = NOW()
WHERE id IN (
    '29ec8df2-5b0f-442d-aa13-1647a9759a0d',
    '35f8405e-5bd5-4803-bade-c50e6d615582'
);

SELECT '✓ Atualizados: % Gordura Mulher (' || COUNT(*) || ' items)' as status
FROM score_items
WHERE id IN ('29ec8df2-5b0f-442d-aa13-1647a9759a0d', '35f8405e-5bd5-4803-bade-c50e6d615582')
AND LENGTH(clinical_relevance) > 100;

-- Continua nos próximos commits...
-- Total: 4/30 items processados

COMMIT;
