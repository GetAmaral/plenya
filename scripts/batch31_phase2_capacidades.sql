-- ============================================================================
-- BATCH 31 - COGNIÇÃO - FASE 2: CAPACIDADES COGNITIVAS SUBJETIVAS (18 items)
-- ============================================================================

-- CAPACIDADE DA MEMÓRIA PERCEBIDA (3 items)
UPDATE score_items
SET
    clinical_relevance = 'A autopercepção da capacidade de memória representa um indicador clínico significativo que frequentemente precede alterações objetivas em testes neuropsicológicos. Queixas subjetivas de memória (QSM) constituem um fenômeno complexo na interface entre função cognitiva objetiva, estado de humor, metacognição e insight preservado sobre mudanças cognitivas incipientes.

Estudos longitudinais prospectivos (2023-2024) demonstram que QSM consistentes ao longo do tempo, especialmente quando corroboradas por informantes próximos, apresentam valor preditivo significativo para conversão de cognição normal para comprometimento cognitivo leve (CCL) em 3-5 anos (hazard ratio 2,1-2,8). No entanto, a relação entre QSM e performance objetiva é modulada por múltiplos fatores: sintomas depressivos (frequentemente amplificam percepção de déficits), ansiedade generalizada (interfere com consolidação por sobrecarga de memória de trabalho), perfeccionismo e metacognição excessiva, e contexto de alta demanda cognitiva ocupacional.

Na perspectiva da Medicina Funcional Integrativa, QSM podem refletir estados de fadiga cognitiva por depleção de neurotransmissores (dopamina, norepinefrina), privação crônica de sono profundo (consolidação inadequada), sobrecarga de cortisol crônico (atrofia hipocampal reversível), déficits nutricionais subclínicos afetando função cerebral antes de manifestação laboratorial óbvia, neuroinflamação sistêmica de baixo grau, ou resistência insulínica cerebral inicial.',
    patient_explanation = 'Como você percebe sua própria memória é tão importante quanto os testes objetivos. Muitas pessoas que desenvolvem problemas de memória percebem mudanças sutis antes que apareçam em exames. Por outro lado, ansiedade e depressão podem fazer você achar que sua memória está pior do que realmente está.

Se você nota consistentemente dificuldades que outras pessoas também percebem, isso merece investigação. Mas se apenas você nota e está muito ansioso ou estressado, pode ser que a sobrecarga emocional esteja interferindo temporariamente com sua capacidade de prestar atenção e formar memórias.

Fatores tratáveis que afetam percepção de memória: dormir mal (cérebro não consolida bem), muito estresse (ocupa espaço mental), deficiências nutricionais (vitaminas B, ferro), excesso de multitarefa (divide atenção), e ansiedade/depressão (fazem você se cobrar demais).',
    conduct = 'AVALIAÇÃO DIFERENCIADA: Aplicar escalas de humor (PHQ-9) e ansiedade (GAD-7) para diferenciar QSM primária de amplificação por transtornos de humor. Investigar corroboração por informante próximo (cônjuge, familiar). Quantificar impacto funcional real (esquecimentos consequentes vs. preocupação sem impacto).

TESTES OBJETIVOS: Aplicar testes neuropsicológicos breves (5 palavras Dubois, span de dígitos, fluência verbal) para correlacionar percepção subjetiva com performance objetiva. Discrepância grande sugere componente afetivo predominante.

INVESTIGAÇÃO CAUSAL: Avaliar qualidade do sono (diário de 7 dias, escala de Epworth), nível de estresse crônico (cortisol salivar, questionário de estresse percebido), carga de multitarefa e demanda cognitiva ocupacional.

LABORATÓRIO BÁSICO: Rastreio de causas metabólicas - TSH, B12, folato, vitamina D, ferritina, glicemia, hemograma.

INTERVENÇÕES:
- Se componente afetivo predominante: tratar depressão/ansiedade subjacente prioritariamente
- Se privação de sono: protocolo rigoroso de higiene do sono
- Se sobrecarga cognitiva: técnicas de gestão atencional, mindfulness, redução de multitarefa
- Suplementação se déficits: ômega-3 2g/dia, complexo B, magnésio, vitamina D
- Exercício aeróbico regular 150 min/semana (reduz ansiedade, melhora sono, neuroproteção)

REAVALIAÇÃO: Após 8-12 semanas, reavaliar percepção subjetiva e retestes objetivos. Se QSM persistente com declínio objetivo documentado, encaminhar para avaliação neuropsicológica formal completa.',
    updated_at = NOW()
WHERE id IN (
    'f9f0c524-bc41-45dc-b33b-46b7f2234fe9',
    '6437b500-5ab6-4e7a-8f62-981cedc96b40',
    'f703185e-6f49-4bec-942e-c1eed243a6a0'
);

-- CAPACIDADE DE FOCO/CONCENTRAÇÃO/APRENDIZADO (3 items)
UPDATE score_items
SET
    clinical_relevance = 'A capacidade de sustentar atenção, concentrar-se em tarefas complexas e adquirir novos conhecimentos depende da integridade de redes neurais distribuídas envolvendo córtex pré-frontal dorsolateral (atenção executiva), córtex parietal posterior (atenção espacial), e sistemas de alerta subcorticais (locus coeruleus noradrenérgico, núcleo basal de Meynert colinérgico).

Déficits atencionais e de aprendizado podem refletir: (1) TDAH não diagnosticado (persistente desde infância vs. início adulto), (2) fadiga atencional por demanda cognitiva excessiva crônica sem períodos de recuperação, (3) fragmentação do sono com redução de sono profundo (consolidação inadequada do aprendizado), (4) síndrome de burnout com depleção de catecolaminas, (5) depressão com comprometimento executivo ("pseudodemência"), (6) efeitos residuais de medicações (benzodiazepínicos, anticonvulsivantes, alguns antidepressivos), (7) déficits nutricionais de cofatores de neurotransmissores, ou (8) processos neurodegenerativos iniciais.

Estudos de neuroimagem funcional (2024) em indivíduos com queixas de atenção demonstram padrões de hipoativação em rede de controle executivo durante tarefas demandantes, correlacionando-se com níveis de dopamina e norepinefrina pré-frontais. Intervenções que modulam estes sistemas (exercício físico regular, suplementação de tirosina, treinamento cognitivo, otimização do sono) demonstram capacidade de restaurar padrões de ativação neural e melhorar performance atencional.',
    patient_explanation = 'Sua capacidade de manter foco, concentrar-se profundamente e aprender coisas novas depende de sistemas cerebrais específicos que precisam estar bem alimentados, descansados e equilibrados. Problemas de concentração são extremamente comuns no mundo moderno com excesso de distrações, multitarefa constante, sono inadequado e estresse crônico.

Se você sempre teve dificuldades de atenção desde criança, pode ser TDAH não diagnosticado que merece avaliação. Se as dificuldades são mais recentes, geralmente há causas tratáveis: sono fragmentado (seu cérebro não tem energia), excesso de estresse (esgota neurotransmissores importantes), dieta pobre em nutrientes essenciais, sedentarismo (falta estimulação cerebral), ou excesso de multitarefa digital (treina seu cérebro a não se concentrar).

A boa notícia é que atenção e capacidade de aprendizado podem ser treinadas e recuperadas. Práticas comprovadas incluem: períodos de trabalho focado sem distrações, exercícios físicos regulares, meditação mindfulness, sono de qualidade, alimentação estabilizadora de energia cerebral, e períodos regulares de descanso cognitivo.',
    conduct = 'RASTREAMENTO TDAH: Se padrão sugestivo de déficit atencional persistente desde infância, aplicar escalas validadas (ASRS-18, DIVA 2.0). Investigar sintomas nucleares (desatenção, hiperatividade, impulsividade) e prejuízo funcional em múltiplos contextos. Considerar encaminhamento para psiquiatra se positivo.

AVALIAÇÃO DO CONTEXTO: Quantificar demanda cognitiva atual (horas de trabalho intelectual, nível de multitarefa, tempo de tela diário), qualidade e quantidade de sono, nível de atividade física, padrão alimentar.

TESTES OBJETIVOS: Span de dígitos inverso (função executiva), teste de atenção sustentada (d2, CPT), fluência verbal (acesso lexical executivo), trail making test parte B (flexibilidade cognitiva).

LABORATÓRIO FOCADO EM NEUROTRANSMISSÃO:
- Ferritina >75 ng/mL homens, >50 ng/mL mulheres (cofator tirosina hidroxilase para dopamina)
- B12 >500 pg/mL, folato >12 ng/mL (síntese de monoaminas)
- Vitamina D >40 ng/mL (expressão genes neurotransmissores)
- TSH <3,0 mU/L (hipotireoidismo subclínico afeta cognição)
- Glicemia, insulina, HbA1c (hipoglicemias reativas comprometem atenção)
- Cortisol salivar matinal e noturno (padrão circadiano)

INTERVENÇÕES COMPORTAMENTAIS:
- Técnicas de foco profundo: blocos de 25-50 min trabalho focado + pausas (Pomodoro)
- Eliminação de multitarefa: monotarefa deliberada
- Redução de distrações: ambientes silenciosos, modo avião em dispositivos
- Mindfulness: 20 min/dia de meditação focada na respiração (melhora sustentação atencional em 8 semanas)

OTIMIZAÇÃO DO SONO: 7-9h/noite com >90 min sono profundo. Protocolo rigoroso de higiene do sono.

NUTRIÇÃO PARA NEUROTRANSMISSORES:
- Proteína em todas refeições (1,6-2g/kg/dia) para aminoácidos precursores
- Café da manhã com 30g proteína dentro de 1h após acordar (ativa síntese de catecolaminas)
- Carboidratos complexos para liberação sustentada de glicose
- Peixes gordos, ovos, oleaginosas (cofatores)
- Hidratação adequada (desidratação de 2% compromete atenção)

SUPLEMENTAÇÃO:
Base: Complexo B metilado, magnésio L-treonato 2g/dia, ferro se deficiente, vitamina D3 5000 UI, ômega-3 EPA/DHA 2-3g/dia

Atenção específica: L-tirosina 1-2g manhã em jejum (precursor catecolaminas), creatina 5g/dia (energia cerebral), Rhodiola rosea 200-400mg manhã (adaptógeno anti-fadiga), L-teanina 200mg 2-3x/dia (foco calmo), Bacopa monnieri 300mg 2x/dia (atenção sustentada)

EXERCÍCIO: Mínimo 150 min/semana moderado-vigoroso. HIIT 2-3x/semana (máximo aumento BDNF e catecolaminas). Exercício matinal preferencial.

TREINAMENTO COGNITIVO: Apps validados focados em atenção (Cogmed, dual n-back, Lumosity) 25 min/dia 5x/semana.

REAVALIAÇÃO: 8-12 semanas. Retestes objetivos de atenção e percepção subjetiva de melhora funcional. Se insuficiente, considerar avaliação neuropsicológica formal completa e reavaliação de TDAH.',
    updated_at = NOW()
WHERE id IN (
    'c31657f0-f591-4cf3-b054-e146504842bd',
    '28a6710a-f04c-46fe-81df-b10419e8d167',
    'ded705c3-5ecc-4ccd-8dfd-cff55da653fd'
);

-- DISPOSIÇÃO PERCEBIDA SUBJETIVA (3 items) + DISPOSIÇÃO/ENERGIA (3 items) = 6 items
UPDATE score_items
SET
    clinical_relevance = 'Fadiga cognitiva e redução de energia mental representam queixas extremamente prevalentes na prática clínica contemporânea, refletindo a interseção complexa entre demanda cognitiva elevada crônica, privação de sono, disfunções metabólicas, desequilíbrios nutricionais e desregulação neuroendócrina.

A sensação subjetiva de disposição e energia mental correlaciona-se com múltiplos sistemas neurobiológicos: (1) disponibilidade de ATP mitocondrial neuronal (produção energética celular), (2) níveis de neurotransmissores monoaminérgicos (dopamina, norepinefrina, serotonina), (3) regulação do eixo hipotálamo-hipófise-adrenal (cortisol e resposta ao estresse), (4) qualidade do sono e arquitetura de fases (especialmente sono profundo restaurador), (5) inflamação sistêmica de baixo grau (citocinas pró-inflamatórias como IL-6 induzem fadiga), (6) função tireoidiana (metabolismo energético global), e (7) status de ferro e outros micronutrientes essenciais.

Estudos recentes em fadiga crônica e fadiga cognitiva (2023-2025) identificam fenótipos distintos: fadiga periférica (muscular, relacionada a disfunção mitocondrial), fadiga central (cerebral, relacionada a depleção de neurotransmissores e neuroinflamação), e fadiga comportamental (relacionada a sobrecarga cognitiva sem períodos adequados de recuperação). A fadiga cognitiva específica manifesta-se como redução progressiva de performance em tarefas demandantes ao longo do dia, lentificação do processamento mental, dificuldade de sustentar atenção, e sensação de "névoa cerebral".

Biomarcadores funcionais associados a baixa energia mental incluem: ferritina <50 ng/mL (mulheres) ou <75 ng/mL (homens), TSH >2,5 mU/L mesmo com T4 normal (hipotireoidismo subclínico), vitamina D <30 ng/mL, vitamina B12 <500 pg/mL, magnésio eritrocitário baixo, coenzima Q10 reduzida, cortisol salivar com padrão achatado (fadiga adrenal), DHEA-S baixo para idade, testosterona livre baixa (homens e mulheres), e marcadores inflamatórios elevados.',
    patient_explanation = 'A sensação de ter energia mental e disposição para enfrentar o dia depende de múltiplos fatores trabalhando em harmonia. Seu cérebro precisa de combustível constante (glicose ou cetonas), neurotransmissores equilibrados (dopamina para motivação, serotonina para bem-estar), sono restaurador (para "recarregar as baterias"), e ausência de inflamação que drena sua energia.

Fadiga mental persistente raramente é "apenas preguiça" ou "falta de vontade". Geralmente há causas biológicas tratáveis: sono de má qualidade ou insuficiente (seu corpo não se recupera), deficiências nutricionais (especialmente ferro, vitaminas B, vitamina D, magnésio), tireoide funcionando abaixo do ideal, níveis de açúcar instáveis no sangue, estresse crônico esgotando suas reservas, inflamação no corpo afetando o cérebro, ou sobrecarga cognitiva sem períodos de descanso adequados.

A recuperação da energia mental geralmente requer abordagem em múltiplas frentes: otimizar sono profundamente, corrigir deficiências nutricionais, estabilizar açúcar no sangue, reduzir inflamação através da dieta, exercitar-se regularmente (parece contraditório mas funciona), praticar técnicas de recuperação de estresse, e estabelecer ritmos de trabalho-descanso sustentáveis.',
    conduct = 'AVALIAÇÃO CLÍNICA ABRANGENTE:
Diferenciar fadiga física vs. cognitiva vs. emocional. Quantificar cronologia (recente vs. crônica), padrão temporal (pior manhã vs. tarde vs. constante), fatores moduladores (melhora com descanso? melhora com exercício?), impacto funcional (limita atividades?).

Rastrear depressão (PHQ-9), ansiedade (GAD-7), e burnout (Maslach Burnout Inventory). Avaliar qualidade do sono (diário 7 dias, Epworth, PSQI). Investigar apneia do sono (questionário de Berlim).

LABORATÓRIO METABÓLICO COMPLETO:
- Hemograma (anemia frequente em fadiga)
- Ferritina, ferro sérico, saturação transferrina, TIBC
- Perfil tireoidiano completo: TSH, T4 livre, T3 livre, T3 reverso, anti-TPO
- Vitamina B12 >500 pg/mL, folato >12 ng/mL
- Vitamina D 25-OH (alvo 50-80 ng/mL)
- Magnésio eritrocitário (não sérico)
- Glicemia jejum, insulina, HbA1c, TOTG se suspeita hipoglicemia reativa
- Cortisol salivar ao acordar, meio-dia, tarde, noite (padrão circadiano completo)
- DHEA-S
- Perfil hormonal sexual (testosterona total/livre, estradiol se mulher)
- Função hepática (detoxificação)
- hs-CRP, VHS, ferritina (inflamação)
- Coenzima Q10, carnitina se suspeita disfunção mitocondrial

Se fadiga severa persistente: considerar autoimunidade (FAN, anti-tiroidiano, celíaca), infecções crônicas (EBV, CMV, Lyme), metais pesados, micotoxinas, permeabilidade intestinal.

POLISSONOGRAFIA: Se Epworth >10, roncos, apneias testemunhadas. Apneia do sono é causa tratável de fadiga severa.

INTERVENÇÕES PRIORITÁRIAS:

1. OTIMIZAÇÃO DO SONO (não-negociável):
- 7-9h/noite com mínimo 90 min sono profundo (N3)
- Horários consistentes ±30 min
- Exposição solar matinal 10-30 min (reset circadiano)
- Escurecimento total quarto, temperatura 17-19°C
- Sem telas 2h antes dormir
- Rotina relaxante pré-sono 45-60 min

2. NUTRIÇÃO ENERGÉTICA:
- Proteína adequada 1,6-2g/kg/dia (aminoácidos precursores neurotransmissores)
- Café da manhã com 30g proteína dentro 1h acordar
- Carboidratos complexos baixo IG para liberação sustentada glicose
- Gorduras saudáveis (ômega-3, azeite, abacate, oleaginosas)
- Eliminar açúcares refinados (montanha-russa glicêmica)
- Hidratação 35-40 ml/kg/dia
- Considerar jejum intermitente 16:8 (aumenta energia via cetonas)

3. SUPLEMENTAÇÃO ANTI-FADIGA (baseada em déficits identificados):
Base: Complexo B metilado completo, ferro bisglicinato 25-50mg se ferritina baixa, magnésio 400-600mg (glicinato + treonato), vitamina D3 5000-10.000 UI se <40, ômega-3 EPA/DHA 2-3g/dia

Mitocondrial: Coenzima Q10 ubiquinol 200-400mg, PQQ 20mg, acetil-L-carnitina 1500-2000mg, ácido R-alfa-lipóico 300-600mg, creatina 5g/dia

Adaptógenos anti-fadiga: Rhodiola rosea 200-400mg manhã, Ashwagandha 500mg 2x/dia (se cortisol alto), Cordyceps 1000mg manhã, Ginseng Panax 200mg manhã

Neurotransmissores: L-tirosina 1-2g manhã (dopamina/norepinefrina), 5-HTP 50-100mg se serotonina baixa

4. EXERCÍCIO PARADOXAL (fadiga melhora com movimento):
- Exercício aeróbico moderado 30 min/dia 5x/semana (aumenta mitocôndrias, melhora sono)
- HIIT 2-3x/semana (máximo aumento energia e BDNF)
- Evitar overtraining (fadiga por excesso)
- Exercício matinal preferencialmente (sincronização circadiana)

5. MANEJO DO ESTRESSE:
- Técnicas diárias: respiração diafragmática, meditação, yoga, tai chi
- Biofeedback HRV se disponível
- Estabelecer limites trabalho-vida
- Períodos regulares de desconexão digital
- Contato com natureza 120 min/semana

6. RITMO CIRCADIANO:
- Rotina consistente acordar/dormir
- Exposição luz natural manhã e tarde
- Redução luz azul à noite
- Refeições em horários regulares
- Evitar cafeína após 14h

REAVALIAÇÃO: 8-12 semanas com relaboratorial de deficiências corrigidas. Escalas de fadiga (FSS, MFIS). Se fadiga severa persistente apesar de otimizações, encaminhar para especialista em fadiga crônica/medicina integrativa.',
    updated_at = NOW()
WHERE id IN (
    '4c8229d8-3376-43be-a120-d777a7e5cc1b',
    'ee16e133-05c2-4ba6-ab4c-a4b6f7de1229',
    '66a04c19-a77d-4af1-b818-19ae0c1065b9',
    'c3689015-b4ed-4af8-86c0-831dec1f4f6f',
    '5bf586c3-4085-4532-a9d4-5fd9fb2c946e',
    '4fdd5c7f-dfd5-4ad9-bf6e-c4266e3d5df5'
);

-- TESTES RÁPIDOS DE MEMÓRIA (3 items)
UPDATE score_items
SET
    clinical_relevance = 'Testes rápidos de memória aplicados no contexto clínico ambulatorial representam ferramentas de rastreamento cognitivo essenciais para detecção precoce de declínio cognitivo. Instrumentos breves validados incluem: Mini-Exame do Estado Mental (MEEM), Montreal Cognitive Assessment (MoCA), Teste do Desenho do Relógio, Fluência Verbal (animais em 1 minuto, palavras com F/A/S), Teste das 5 Palavras de Dubois, e Span de Dígitos.

A vantagem destes testes rápidos é a aplicabilidade em consultório com tempo limitado (5-10 minutos), sensibilidade razoável para detecção de comprometimento cognitivo leve (CCL) e demência inicial, e capacidade de monitoramento longitudinal evolutivo. O MoCA demonstra sensibilidade superior ao MEEM para CCL (90% vs. 18% em meta-análise de 2024), sendo preferível quando objetivo é detecção precoce ao invés de rastreamento de demência estabelecida.

Na perspectiva funcional, performance inferior ao esperado para idade e escolaridade sugere necessidade de investigação aprofundada de causas potencialmente reversíveis antes de assumir processo neurodegenerativo: deficiências de B12/folato (ácido metilmalônico, homocisteína), hipotireoidismo (TSH, T4 livre, T3 livre), privação de sono crônica, depressão pseudodemência, efeitos medicamentosos anticolinérgicos, deficiências de vitamina D, ferro, magnésio, ou resistência insulínica cerebral.',
    patient_explanation = 'Testes rápidos de memória são como uma "triagem" inicial para avaliar como seu cérebro está funcionando em várias áreas: memória, atenção, linguagem, capacidade de planejar e organizar. São testes simples que levam poucos minutos mas fornecem informações valiosas.

Exemplos incluem: lembrar 3-5 palavras depois de alguns minutos, desenhar um relógio mostrando hora específica, nomear o máximo de animais em 1 minuto, repetir números de trás para frente, copiar desenhos geométricos. Cada tarefa avalia uma parte diferente do cérebro.

Se você tem dificuldade nesses testes, não significa automaticamente algo grave. Muitas causas tratáveis podem afetar performance: ansiedade durante o teste, sono ruim, deficiências vitamínicas, depressão, efeito de medicações, ou simplesmente não estar familiarizado com esse tipo de teste. Por isso, resultados sempre devem ser interpretados junto com sua história, outros exames, e como você funciona no dia-a-dia.',
    conduct = 'SELEÇÃO DO INSTRUMENTO: MoCA para rastreamento sensível de CCL, MEEM se rastreamento de demência estabelecida, Teste das 5 Palavras de Dubois se foco em memória episódica.

APLICAÇÃO PADRONIZADA: Seguir protocolo validado rigorosamente. Considerar ajuste por escolaridade (MoCA adiciona 1 ponto se ≤12 anos estudo).

INTERPRETAÇÃO CONTEXTUALIZADA:
- Considerar ansiedade de performance, familiaridade com testes
- Correlacionar com percepção do paciente e informante
- Avaliar impacto funcional real nas AVDs
- Repetir em condições otimizadas se fatores agudos presentes

PONTOS DE CORTE:
- MoCA: <26/30 sugere CCL, <18 sugere demência
- MEEM: <24/30 ajustado por escolaridade sugere demência
- Fluência verbal: <12 animais/minuto preocupante

SE ESCORE ALTERADO:
1. Investigação de causas reversíveis (laboratorial completo cognitivo)
2. Otimização de sono, nutrição, medicações
3. Repetição em 8-12 semanas após otimizações
4. Se persiste alterado, encaminhar para neuropsicologia formal completa
5. Considerar neuroimagem (RM encéfalo) se progressão documentada

MONITORAMENTO: Repetir mesmo teste a cada 6-12 meses para avaliar evolução temporal. Declínio progressivo requer investigação mais agressiva.',
    updated_at = NOW()
WHERE id IN (
    'bad56a9b-1950-447a-b904-b8817754ba77',
    'e8416da2-f19b-4929-813b-9ead82399476',
    '5f817e73-71cb-46f9-a3a6-2fc1258b01f7'
);

-- USO ATUAL DE PSICOTRÓPICOS PARA COGNIÇÃO (3 items)
UPDATE score_items
SET
    clinical_relevance = 'O uso de psicotrópicos e medicações com efeitos sobre cognição é extremamente prevalente e frequentemente representa fator modificável importante em queixas cognitivas. Múltiplas classes medicamentosas afetam cognição através de mecanismos anticolinérgicos, gabaérgicos, ou dopaminérgicos: benzodiazepínicos (comprometimento de memória e atenção), antidepressivos tricíclicos (efeitos anticolinérgicos), anti-histamínicos de primeira geração, anticonvulsivantes, antipsicóticos, relaxantes musculares, antiespasmódicos, alguns anti-hipertensivos, e opioides.

Estudos recentes (2023-2024) demonstram carga anticolinérgica cumulativa como fator de risco independente para declínio cognitivo e demência em idosos. Meta-análise de 2024 identificou aumento de 50% no risco de demência associado a uso crônico de múltiplas medicações anticolinérgicas. Benzodiazepínicos, embora amplamente prescritos, demonstram efeitos deletérios na memória episódica e aumento de risco de acidentes, quedas e declínio cognitivo com uso prolongado.

Por outro lado, alguns psicotrópicos podem beneficiar cognição quando apropriadamente indicados: inibidores seletivos de recaptação de serotonina (ISRS) melhoram cognição secundariamente via tratamento de depressão; estimulantes (metilfenidato, lisdexanfetamina) melhoram atenção e função executiva em TDAH; alguns anticonvulsivantes estabilizadores de humor beneficiam função executiva em transtorno bipolar; e inibidores de acetilcolinesterase (donepezila, rivastigmina) estão indicados em demências.

A abordagem funcional prioriza: (1) revisão crítica de indicação atual de cada medicação, (2) identificação de medicações potencialmente prejudiciais à cognição, (3) discussão de alternativas menos prejudiciais quando possível, (4) descontinuação gradual de medicações desnecessárias, (5) otimização de doses mínimas efetivas.',
    patient_explanation = 'Muitas medicações comumente prescritas podem afetar sua memória, atenção e clareza mental. Calmantes como benzodiazepínicos (Rivotril, Diazepam, Alprazolam), alguns antidepressivos antigos, remédios para alergia que causam sono, alguns relaxantes musculares, e vários outros podem causar "névoa cerebral", dificuldades de memória, e lentificação do pensamento.

Isso não significa que você deve parar suas medicações sozinho - nunca faça isso sem orientação médica, pois pode ser perigoso. Mas é importante revisar com seu médico se cada medicação ainda é realmente necessária, se existe alternativa que afete menos a cognição, e se a dose pode ser reduzida para o mínimo efetivo.

Por outro lado, algumas medicações podem ajudar sua cognição quando corretamente indicadas: antidepressivos se você tem depressão (a depressão em si prejudica memória), estimulantes se você tem TDAH diagnosticado, ou medicações específicas para demência se for o caso. O importante é sempre balancear benefícios e riscos com acompanhamento médico.',
    conduct = 'REVISÃO MEDICAMENTOSA COMPLETA:
Listar todas medicações em uso (prescritas e automedicação), incluindo dose, frequência, tempo de uso, e indicação original.

IDENTIFICAR MEDICAÇÕES POTENCIALMENTE PREJUDICIAIS:
- Benzodiazepínicos: considerar descontinuação gradual (redução 10-25% a cada 2 semanas), substituir por alternativas (trazodona baixa dose para sono, buspirona para ansiedade, TCC)
- Anticolinérgicos: revisar necessidade, substituir por alternativas (antidepressivos ISRS ao invés de tricíclicos, anti-histamínicos segunda geração ao invés de primeira)
- Z-drugs (zolpidem, zopiclone): similar a benzodiazepínicos, descontinuação gradual preferível
- Gabapentinoides em altas doses: considerar redução dose
- Anticonvulsivantes sedativos: otimizar dose mínima
- Opioides crônicos: protocolo de redução se apropriado

CALCULAR CARGA ANTICOLINÉRGICA TOTAL: Utilizar escala ACB (Anticholinergic Cognitive Burden). Escore ≥3 associado a risco aumentado de declínio cognitivo.

ALTERNATIVAS MENOS PREJUDICIAIS:
- Insônia: higiene do sono + melatonina + magnésio + trazodona baixa dose ao invés de benzodiazepínicos
- Ansiedade: ISRS (sertralina, escitalopram) + TCC ao invés de benzodiazepínicos
- Dor crônica: abordagens multimodais (fisioterapia, acupuntura, anti-inflamatórios, duloxetina) ao invés de opioides

DESMAME SEGURO:
- Nunca abrupto (risco de abstinência e convulsões)
- Benzodiazepínicos: redução 10-25% a cada 1-2 semanas, pode levar meses
- Monitorar sintomas de abstinência (ansiedade rebote, insônia, tremor, sintomas autonômicos)
- Suporte com magnésio, taurina, valerian root durante desmame

MONITORAMENTO: Reavaliar cognição 8-12 semanas após ajustes medicamentosos. Documentar melhora ou ausência de mudança. Se cognição melhora significativamente, confirma contribuição medicamentosa ao déficit.',
    updated_at = NOW()
WHERE id IN (
    'ac386243-fea2-461a-9428-fa18dffe6a31',
    '36d39438-5337-4e06-b67d-edfaa498ba25',
    'c5cec641-d523-4021-b96c-6fab4713d55e'
);

-- Verificação
SELECT COUNT(*) as updated FROM score_items WHERE id IN (
    'f9f0c524-bc41-45dc-b33b-46b7f2234fe9', '6437b500-5ab6-4e7a-8f62-981cedc96b40', 'f703185e-6f49-4bec-942e-c1eed243a6a0',
    'c31657f0-f591-4cf3-b054-e146504842bd', '28a6710a-f04c-46fe-81df-b10419e8d167', 'ded705c3-5ecc-4ccd-8dfd-cff55da653fd',
    '4c8229d8-3376-43be-a120-d777a7e5cc1b', 'ee16e133-05c2-4ba6-ab4c-a4b6f7de1229', '66a04c19-a77d-4af1-b818-19ae0c1065b9',
    'c3689015-b4ed-4af8-86c0-831dec1f4f6f', '5bf586c3-4085-4532-a9d4-5fd9fb2c946e', '4fdd5c7f-dfd5-4ad9-bf6e-c4266e3d5df5',
    'bad56a9b-1950-447a-b904-b8817754ba77', 'e8416da2-f19b-4929-813b-9ead82399476', '5f817e73-71cb-46f9-a3a6-2fc1258b01f7',
    'ac386243-fea2-461a-9428-fa18dffe6a31', '36d39438-5337-4e06-b67d-edfaa498ba25', 'c5cec641-d523-4021-b96c-6fab4713d55e'
) AND clinical_relevance IS NOT NULL;
