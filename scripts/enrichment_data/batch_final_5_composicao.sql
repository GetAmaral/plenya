-- ============================================================================
-- BATCH FINAL: Enriquecimento de 3 Items de Composição Corporal
-- Grupo: Composição Corporal > Atual
-- Padrão MFI: Adiposidade visceral, sarcopenia, metabolismo adaptativo
-- ============================================================================

BEGIN;

-- Item 1: Quadril (em cm)
UPDATE score_items
SET
    clinical_relevance = 'A circunferência do quadril é um marcador fundamental para avaliação da distribuição de gordura corporal e composição músculo-esquelética. A medida reflete a massa muscular glútea e femoral, tecido adiposo subcutâneo periférico e estrutura óssea pélvica. Na medicina funcional integrativa, o quadril não é apenas uma medida antropométrica isolada, mas um componente essencial para o cálculo da razão cintura/quadril (RCQ), que estratifica o padrão de distribuição de adiposidade (andróide vs. ginóide). Valores adequados de circunferência do quadril, quando combinados com circunferência de cintura baixa, indicam perfil metabólico favorável com menor risco cardiometabólico. A relação quadril/coxa também é utilizada em protocolos avançados de avaliação de sarcopenia e fragilidade. Redução progressiva da circunferência do quadril em idosos pode sinalizar perda de massa muscular glútea (sarcopenia glútea), associada a quedas, fraturas de quadril e declínio funcional. A manutenção ou ganho de massa muscular no quadril através de treinamento de força (agachamentos, levantamento terra, exercícios unilaterais) é estratégia central em protocolos de envelhecimento saudável e prevenção de fragilidade.',

    patient_explanation = 'A medida do quadril mostra o tamanho da região do seu quadril e glúteos. Essa medida é importante porque, combinada com a medida da cintura, ajuda a entender como a gordura está distribuída no seu corpo. Pessoas que acumulam mais gordura no quadril e coxas (formato "pera") geralmente têm menor risco de problemas cardíacos e diabetes do que pessoas que acumulam gordura na barriga (formato "maçã"). Além disso, a medida do quadril reflete a quantidade de músculo nos glúteos e coxas, que são músculos essenciais para caminhar, subir escadas e prevenir quedas, especialmente com o envelhecimento. Manter ou aumentar a massa muscular do quadril através de exercícios de força (como agachamentos) é fundamental para sua saúde e independência a longo prazo.',

    conduct = 'AVALIAÇÃO INICIAL: Medir circunferência do quadril com fita métrica no ponto de maior protuberância glútea, paciente em posição ortostática relaxada. Realizar medida 3 vezes e calcular média. Registrar resultado e calcular razão cintura/quadril (RCQ = circunferência cintura / circunferência quadril). Avaliar em conjunto com bioimpedância, dobras cutâneas e força muscular.

INTERPRETAÇÃO CONTEXTUALIZADA: Valores isolados têm pouca relevância clínica. Avaliar sempre em contexto: RCQ, percentual de gordura, massa muscular apendicular, força de membros inferiores. Redução progressiva do quadril sem perda intencional de peso sugere sarcopenia. Aumento do quadril com ganho de força e desempenho físico indica hipertrofia muscular saudável.

INTERVENÇÕES NUTRICIONAIS: Proteína 1.6-2.2 g/kg/dia para síntese muscular glútea e femoral. Leucina 3-4 g por refeição para maximizar estímulo anabólico. Creatina monohidratada 5 g/dia para ganho de massa magra e força. Vitamina D >40 ng/mL e ômega-3 EPA+DHA 2-3 g/dia para modulação inflamatória e suporte à hipertrofia. Colágeno tipo I e II 10-15 g/dia para saúde articular do quadril.

PROTOCOLO DE TREINAMENTO: Foco em exercícios compostos para quadril e membros inferiores: agachamento livre, agachamento búlgaro, levantamento terra, hip thrust, stiff. Frequência mínima 2-3x/semana, progressão de carga 5-10% a cada 2 semanas. Incluir exercícios unilaterais (afundo, step-up) para correção de assimetrias. Volume: 3-5 séries de 6-12 repetições com carga moderada-alta. Idosos: iniciar com cadeira, elástico e peso corporal, progressão gradual.

MONITORAMENTO: Reavaliar circunferência do quadril a cada 8-12 semanas. Correlacionar com força (teste de sentar-levantar), potência (salto vertical), funcionalidade (TUG, velocidade de marcha). Aumento de 2-4 cm com ganho de força indica hipertrofia muscular positiva. Redução >3 cm sem intervenção intencional exige investigação de sarcopenia, caquexia ou doença sistêmica.',

    last_review = CURRENT_TIMESTAMP

WHERE id = '1a9be52d-aa0b-4a53-a71a-4eff3e0cc6ba';


-- Item 2: Razão cintura/quadril - homem
UPDATE score_items
SET
    clinical_relevance = 'A razão cintura/quadril (RCQ) em homens é um dos marcadores antropométricos mais robustos para estratificação de risco cardiometabólico e mortalidade. A RCQ reflete o padrão de distribuição de adiposidade: valores elevados indicam padrão andróide (acúmulo central/visceral), enquanto valores baixos indicam padrão ginóide (acúmulo periférico/subcutâneo). A adiposidade visceral, quantificada indiretamente pela RCQ elevada, está causalmente associada a resistência insulínica, dislipidemia aterogênica, inflamação crônica de baixo grau, esteatose hepática, síndrome metabólica, doença cardiovascular e mortalidade por todas as causas. Na medicina funcional integrativa, a RCQ é utilizada como marcador dinâmico de resposta a intervenções: protocolos de jejum intermitente, dieta cetogênica, treinamento intervalado de alta intensidade (HIIT) e treinamento de força demonstram redução significativa da RCQ através da mobilização preferencial de gordura visceral. Valores ótimos em homens: <0.90 (baixo risco). Valores 0.90-0.95 indicam risco moderado. Valores >0.95 indicam risco elevado e necessidade de intervenção intensiva multimodal. A RCQ tem valor preditivo superior ao IMC para desfechos cardiovasculares e metabólicos em múltiplos estudos epidemiológicos.',

    patient_explanation = 'A razão cintura/quadril é um cálculo simples (cintura dividido pelo quadril) que mostra se você acumula mais gordura na barriga ou nas pernas. Para homens, o ideal é ter uma razão menor que 0.90, pois isso significa que você não tem muita gordura visceral (a gordura perigosa que fica ao redor dos órgãos internos). Quando essa razão está acima de 0.95, significa que você tem muito acúmulo de gordura na região abdominal, o que aumenta significativamente o risco de diabetes, problemas cardíacos, pressão alta e gordura no fígado. A boa notícia é que a gordura da barriga responde muito bem a mudanças no estilo de vida: exercícios de força combinados com jejum intermitente ou dieta low-carb podem reduzir essa gordura de forma rápida e melhorar sua saúde metabólica. Esse número é mais importante do que o peso na balança para prever sua saúde futura.',

    conduct = 'AVALIAÇÃO INICIAL: Medir circunferência da cintura (ponto médio entre última costela e crista ilíaca) e circunferência do quadril (maior protuberância glútea). Calcular RCQ = cintura/quadril. Realizar medidas em jejum, bexiga vazia, sem roupa compressiva. Complementar com bioimpedância (gordura visceral), ultrassom de gordura abdominal ou tomografia para quantificação precisa de adiposidade visceral quando RCQ >0.95.

ESTRATIFICAÇÃO DE RISCO:
- RCQ <0.90: Baixo risco cardiometabólico, manutenção do padrão atual
- RCQ 0.90-0.95: Risco moderado, intervenção preventiva precoce
- RCQ 0.95-1.00: Risco elevado, intervenção intensiva necessária
- RCQ >1.00: Risco muito elevado, investigar síndrome metabólica, esteatose, resistência insulínica

PROTOCOLO NUTRICIONAL: Dieta com restrição de carboidratos refinados e açúcares (<50-100 g/dia) para mobilização de gordura visceral. Jejum intermitente 16:8 ou 18:6 aumenta lipólise abdominal. Proteína alta 1.8-2.2 g/kg para preservação de massa magra durante déficit calórico. Fibras solúveis 30-40 g/dia (psyllium, glucomanano) para modulação de microbiota e saciedade. Polifenóis (chá verde, resveratrol, curcumina) para inibição de adipogênese visceral.

TREINAMENTO ESPECÍFICO: HIIT 3x/semana (sprints, bike, remo) para maximizar lipólise visceral via catecolaminas. Treinamento de força 3-4x/semana para aumento de massa magra e gasto energético de repouso. Caminhar 10.000 passos/dia para mobilização contínua de ácidos graxos. Evitar excesso de cardio steady-state que pode aumentar cortisol e preservar gordura visceral.

SUPLEMENTAÇÃO ESTRATÉGICA: Ômega-3 EPA+DHA 2-4 g/dia para redução de gordura visceral e inflamação. Berberina 500 mg 3x/dia ou metformina para sensibilidade insulínica. Magnésio 400-600 mg/dia para metabolismo glicêmico. Probióticos (Lactobacillus, Bifidobacterium) para modulação de microbiota e redução de endotoxemia metabólica.

MONITORAMENTO: Reavaliar RCQ a cada 8 semanas. Redução de 0.02-0.05 pontos por mês indica resposta adequada. Avaliar marcadores associados: glicemia de jejum, HOMA-IR, triglicerídeos, HDL, PCR-us, ALT/AST (esteatose). Manter intervenções até RCQ <0.90 e estabilização por 6+ meses.',

    last_review = CURRENT_TIMESTAMP

WHERE id = 'c9348fbd-df44-4903-ac22-1d8b5b47179a';


-- Item 3: Razão cintura/quadril - mulher
UPDATE score_items
SET
    clinical_relevance = 'A razão cintura/quadril (RCQ) em mulheres é um marcador antropométrico essencial para avaliação de risco cardiometabólico, perfil hormonal e padrão de distribuição adiposa. Mulheres naturalmente apresentam distribuição ginóide (quadris e coxas) devido à ação estrogênica, com valores de RCQ tipicamente mais baixos que homens. A transição para padrão andróide (aumento de RCQ) ocorre principalmente na menopausa, devido à queda estrogênica e mudança no perfil de adipocinas, resultando em acúmulo visceral. Na medicina funcional integrativa, a RCQ em mulheres é avaliada contextualmente com fase da vida, status hormonal e sintomas associados. Valores ótimos: <0.80 (baixo risco). RCQ 0.80-0.85 indica risco moderado. RCQ >0.85 indica risco elevado, especialmente em pós-menopausa. Elevação da RCQ correlaciona-se com resistência insulínica, síndrome dos ovários policísticos (SOP), hiperandrogenismo, inflamação sistêmica, dislipidemia, doença cardiovascular e câncer de mama pós-menopausa. A RCQ é mais sensível que o IMC para detectar risco metabólico em mulheres com peso normal mas adiposidade visceral elevada (fenótipo "magra metabólicamente obesa"). Protocolos de modulação hormonal bioidêntica, treinamento de força e nutrição anti-inflamatória demonstram eficácia na redução da RCQ e reversão de disfunções metabólicas em mulheres peri e pós-menopáusicas.',

    patient_explanation = 'A razão cintura/quadril mostra como a gordura está distribuída no seu corpo. Para mulheres, o ideal é ter uma razão menor que 0.80, pois isso significa que você mantém o padrão feminino saudável de acumular gordura nos quadris e coxas, ao invés da barriga. Esse padrão é influenciado pelo estrogênio, o hormônio feminino. Quando a razão está acima de 0.85, significa que você está acumulando mais gordura na região abdominal (gordura visceral), o que aumenta o risco de diabetes, problemas cardíacos e desequilíbrios hormonais. Isso é especialmente comum após a menopausa, quando os níveis de estrogênio caem. A gordura visceral também está associada à síndrome dos ovários policísticos (SOP) e pode piorar sintomas como irregularidade menstrual e resistência à insulina. A boa notícia é que exercícios de força, alimentação anti-inflamatória e, em alguns casos, reposição hormonal bioidêntica podem ajudar a reverter esse padrão e melhorar sua saúde metabólica e hormonal.',

    conduct = 'AVALIAÇÃO INICIAL: Medir circunferência da cintura (ponto médio entre última costela e crista ilíaca) e circunferência do quadril (maior protuberância glútea). Calcular RCQ = cintura/quadril. Investigar contexto hormonal: idade, status menstrual, sintomas de SOP (acne, hirsutismo, irregularidade menstrual), uso de contraceptivos. Complementar com bioimpedância, perfil hormonal (estradiol, progesterona, testosterona, DHEA-S) e marcadores metabólicos (glicemia, insulina, HOMA-IR).

ESTRATIFICAÇÃO DE RISCO:
- RCQ <0.80: Baixo risco, padrão ginóide saudável
- RCQ 0.80-0.85: Risco moderado, transição para padrão andróide
- RCQ >0.85: Risco elevado, investigar SOP, resistência insulínica, pré-menopausa/menopausa
- RCQ >0.90: Risco muito elevado, síndrome metabólica provável

PROTOCOLO NUTRICIONAL: Dieta anti-inflamatória mediterrânea ou low-carb (100-150 g/dia) para sensibilidade insulínica. Proteína 1.6-2.0 g/kg/dia para preservação muscular e controle de apetite. Gorduras saudáveis (azeite, abacate, nozes) 30-35% das calorias para suporte hormonal. Fitoestrógenos (linhaça, soja orgânica) para modulação estrogênica leve em pós-menopausa. Fibras 30-40 g/dia para detoxificação de estrogênios e saúde intestinal. Evitar açúcares, álcool e alimentos ultra-processados que pioram resistência insulínica.

MODULAÇÃO HORMONAL: Avaliar reposição hormonal bioidêntica (estradiol transdérmico + progesterona micronizada) em pós-menopausa sintomática com RCQ >0.85. DIM (di-indolilmetano) 200-400 mg/dia para metabolização saudável de estrogênios. Vitex agnus-castus para suporte progesterônico em pré-menopausa. Myo-inositol + D-chiro-inositol para SOP com resistência insulínica.

TREINAMENTO ESTRATÉGICO: Treinamento de força 3-4x/semana para aumento de massa magra e gasto metabólico. Foco em grandes grupos musculares (agachamento, levantamento terra, remada). HIIT 2-3x/semana para mobilização de gordura visceral. Exercícios de core e pelve para fortalecimento funcional. Evitar excesso de cardio prolongado que pode elevar cortisol.

SUPLEMENTAÇÃO ADJUVANTE: Ômega-3 EPA+DHA 2-3 g/dia para modulação inflamatória. Magnésio 400 mg/dia para sensibilidade insulínica. Vitamina D >40 ng/mL para função imune e hormonal. Probióticos para estroboloma (metabolização intestinal de estrogênios). Berberina 500 mg 3x/dia para resistência insulínica em SOP.

MONITORAMENTO: Reavaliar RCQ a cada 8-12 semanas. Correlacionar com sintomas hormonais (ciclo menstrual, sintomas de menopausa), marcadores metabólicos (glicemia, insulina, lipídios) e composição corporal (bioimpedância). Redução de 0.02-0.04 pontos por mês indica resposta adequada. Manter acompanhamento até RCQ <0.80 e estabilização clínica.',

    last_review = CURRENT_TIMESTAMP

WHERE id = 'b2414f5e-8ad7-4b9c-846e-edd6a3c8277f';

COMMIT;

-- ============================================================================
-- VERIFICAÇÃO PÓS-EXECUÇÃO
-- ============================================================================
SELECT
    id,
    name,
    CASE
        WHEN clinical_relevance IS NOT NULL AND LENGTH(clinical_relevance) > 500 THEN '✓ ENRIQUECIDO'
        ELSE '✗ PENDENTE'
    END as status,
    LENGTH(clinical_relevance) as relevance_chars,
    LENGTH(patient_explanation) as explanation_chars,
    LENGTH(conduct) as conduct_chars,
    last_review
FROM score_items
WHERE id IN (
    '1a9be52d-aa0b-4a53-a71a-4eff3e0cc6ba',
    'c9348fbd-df44-4903-ac22-1d8b5b47179a',
    'b2414f5e-8ad7-4b9c-846e-edd6a3c8277f'
)
ORDER BY name;
