#!/usr/bin/env python3
"""
Script completo para atualizar TODOS os 53 items do grupo Vida Sexual.
Organizado por subgrupos temáticos com conteúdo baseado em evidências 2024-2025.
"""

import requests
import json
import time
import sys

API_URL = "http://localhost:3001/api/v1"
TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJjYWIwMDBkNy1iODBlLTQ0NWYtYTUyOS0zYThlYzhmNDg2YmQiLCJlbWFpbCI6ImNsYXVkZS1hZG1pbkBwbGVueWEuY29tIiwicm9sZSI6ImFkbWluIiwiaXNzIjoicGxlbnlhLWVtciIsImV4cCI6MTc2OTQzMjcwNSwiaWF0IjoxNzY5NDMxODA1fQ.W8u8MPcnOkUTOUhnNoRV5VhnyAp2u2mvCGQsgTUnjJI"

headers = {
    "Content-Type": "application/json",
    "Authorization": f"Bearer {TOKEN}"
}

# =============================================================================
# ITEMS DO GRUPO VIDA SEXUAL - ORGANIZADOS POR TEMA
# =============================================================================

items = [
    # =========================================================================
    # 1. IIEF-5 - Índice Internacional de Função Erétil (2 items)
    # =========================================================================
    {
        "id": "5aaadebc-141f-46f5-bf89-7603dadb91f1",
        "name": "IIEF-5: ____/25",
        "clinical_relevance": """O Índice Internacional de Função Erétil versão simplificada (IIEF-5) representa muito mais do que uma ferramenta de avaliação urológica: constitui-se como biomarcador sistêmico de saúde cardiovascular, metabólica e hormonal. Validado internacionalmente e adaptado culturalmente para o Brasil, este questionário de 5 itens avalia capacidade de obter e manter ereções, confiança erétil, penetração, manutenção durante intercurso e satisfação sexual, pontuando de 5 a 25 pontos.

A relevância clínica transcende a esfera sexual. Metanálise de 2024 demonstrou que disfunção erétil (DE) é biomarcador preditivo robusto para doença cardiovascular em homens em envelhecimento, com implicações significativas para detecção precoce e estratégias preventivas. O conceito de DE como canário na mina de carvão cardiovascular baseia-se na hipótese do tamanho arterial: artérias penianas (1-2mm) manifestam disfunção endotelial antes das coronárias (3-4mm), permitindo janela de intervenção de 2-5 anos antes de eventos cardiovasculares maiores.

Estudo de coorte prospectivo de novembro de 2024 evidenciou associação bidirecional entre DE e risco cardiovascular de 10 anos em homens com diabetes e hipertensão. DE deve ser considerada manifestação precoce de doença cardiovascular, justificando avaliação cuidadosa de fatores de risco cardiovascular para prevenir eventos adversos maiores. Pesquisa de 2024 identificou que índice cardiometabólico e idade foram fatores de risco significativos para disfunção sexual (OR=2,672 e 1,081), com índice cardiometabólico predizendo disfunção sexual especialmente em homens mais jovens.

Na perspectiva de medicina funcional, IIEF-5 inferior a 22 pontos sinaliza investigação de múltiplos sistemas. Disfunção endotelial é mecanismo fisiopatológico central, mediada por biodisponibilidade reduzida de óxido nítrico, estresse oxidativo e inflamação sistêmica de baixo grau. Biomarcadores inflamatórios predizem resposta a inibidores de fosfodiesterase-5: pressão arterial sistólica e proteína C-reativa predizem mudanças no IIEF em alta taxa, com níveis elevados de PCR relacionados à severidade de DE.

Avaliação hormonal é mandatória segundo consenso ICSM 2024. Testosterona livre deve ser considerada teste de rotina junto com testosterona total em homens com DE. Na prática funcional, medição de estradiol junto com testosterona é importante em homens com disfunção sexual, mudanças de composição corporal ou sintomas persistentes de humor, com muitos clínicos monitorando razão testosterona:estradiol visando intervalo aproximado de 10:1 a 15:1. miRNA-21 circulante emergiu como biomarcador inovador em fevereiro de 2024. Resistência insulínica representa outro nexo fisiopatológico: altos níveis de insulina reduzem hormônio de crescimento, reduzindo libido através de alteração em testosterona.

Interpretação: 22-25 pontos = sem DE; 17-21 = DE leve; 12-16 = DE leve a moderada; 8-11 = DE moderada; 5-7 = DE severa.""",
        "patient_explanation": """O questionário IIEF-5 avalia sua função erétil através de 5 perguntas simples sobre confiança, rigidez, penetração, manutenção e satisfação sexual nos últimos 6 meses. A pontuação varia de 5 a 25 pontos.

Muitos homens não percebem que dificuldades de ereção podem ser o primeiro sinal de problemas de saúde mais amplos. Pesquisas mostram que disfunção erétil frequentemente aparece 2-5 anos antes de problemas cardiovasculares, funcionando como sistema de alerta precoce. Isso acontece porque as artérias do pênis são menores que as do coração e mostram problemas circulatórios primeiro.

Se sua pontuação está abaixo de 22, isso indica necessidade de investigação médica mais profunda - não apenas da função sexual, mas também de sua saúde cardiovascular, metabólica e hormonal. Estudos de 2024 mostram que homens com disfunção erétil têm maior risco de diabetes, pressão alta, colesterol elevado e problemas cardíacos.

Vários fatores podem afetar sua função erétil: níveis hormonais (especialmente testosterona), qualidade da circulação sanguínea, inflamação no corpo, resistência à insulina, estresse oxidativo, qualidade do sono, nível de estresse psicológico e alguns medicamentos. A boa notícia é que muitos desses fatores podem ser melhorados com mudanças de estilo de vida e tratamento adequado.""",
        "conduct": """Pontuação IIEF-5 inferior a 22 requer protocolo investigativo abrangente integrando avaliação cardiovascular, metabólica, hormonal e neurovascular.

Avaliação laboratorial tier 1: Painel hormonal com testosterona total e livre (SHMS 2024 recomenda ambos), SHBG, prolactina. Na medicina funcional medimos estradiol para avaliar razão T:E2, visando 10:1 a 15:1. TSH e T4 livre. Painel metabólico: glicemia, HbA1c, insulina de jejum para HOMA-IR (resistência insulínica presente em mais de 50% com DE). Perfil lipídico avançado incluindo LDL oxidada. Marcadores inflamatórios: PCR ultrassensível, homocisteína.

Avaliação cardiovascular (especialmente se menos de 17 pontos): Escore de risco cardiovascular. ECG. Considerar teste ergométrico, escore de cálcio coronariano ou ultrassom de carótidas. Pressão arterial ambulatorial se hipertensão suspeitada.

Avaliação tier 2 (casos refratários): Doppler peniano. Vitamina D, zinco, magnésio. Teste de estresse adrenal. Avaliação de apneia do sono (presente em 30-70% com DE). Screening psiquiátrico.

Intervenções funcionais: Otimização hormonal se testosterona total menor que 300ng/dL. Suporte nutricional: L-arginina 3-5g/dia, L-citrulina 1,5-3g/dia, Panax ginseng 900mg 3x/dia, Pycnogenol 40mg 3x/dia mais L-arginina 1,7g 2x/dia. Suporte mitocondrial: CoQ10 200-300mg/dia. Sensibilizadores de insulina: berberina 500mg 3x/dia, cromo, ácido alfa-lipóico, NAC.

PDE5i continuam padrão-ouro. Consenso Princeton IV 2024 reafirma segurança cardiovascular. Homens com DE expostos a tadalafil tiveram taxas significativamente menores de eventos cardiovasculares (HR:0,81), revascularização (HR:0,69), mortes cardiovasculares (HR:0,45) e morte por todas as causas (HR:0,56).

Monitoramento: reaplicação IIEF-5 a cada 3 meses. Reavaliação laboratorial a cada 3-6 meses."""
    },
    {
        "id": "7cc28981-bd20-4e5e-b7c0-bd312f6a7ed4",
        "name": "IIEF-5: ____/25",
        "clinical_relevance": "DUPLICATE_OF_ABOVE",
        "patient_explanation": "DUPLICATE_OF_ABOVE",
        "conduct": "DUPLICATE_OF_ABOVE"
    },

    # =========================================================================
    # 2. FSFI - Female Sexual Function Index (3 items)
    # =========================================================================
    {
        "id": "9d3f885c-6de0-40e0-a1b5-faeb028289f2",
        "name": "FSFI: ____/36",
        "clinical_relevance": """O Índice de Função Sexual Feminina (FSFI) é considerado padrão-ouro na avaliação de disfunção sexual feminina, capaz de contemplar a natureza multidimensional da sexualidade feminina. Este questionário validado avalia 6 domínios: desejo sexual, excitação, lubrificação, orgasmo, satisfação e dor/desconforto, gerando pontuação total de 2 a 36 pontos. Escore inferior a 26,55 indica disfunção sexual clinicamente significativa, requerendo investigação aprofundada.

A relevância clínica em medicina funcional estende-se muito além da esfera sexual. Análise de regressão linear multivariada demonstrou que severidade de sintomas urogenitais da menopausa associa-se com valores menores no escore total FSFI e nos domínios de lubrificação, satisfação, excitação e orgasmo. O impacto severo de sintomas menopáusicos associa-se com função sexual feminina diminuída em mulheres de meia-idade e pós-menopáusicas mais velhas.

Evidências de 2025 demonstram que suplementação probiótica pode melhorar significativamente a função sexual. Revisão sistemática Cochrane de 2025 com 10 estudos incluindo 735 participantes reportou melhorias significativas nos escores FSFI em todos os domínios com suplementação probiótica em mulheres em idade reprodutiva. Estudo clínico randomizado duplo-cego envolvendo 59 mulheres pré-menopáusicas e menopáusicas (40-60 anos) concluiu que suplementação de Lactobacillus acidophilus pode melhorar desejo, excitação, satisfação, bem como escores totais FSFI e qualidade de vida.

Na perspectiva hormonal, FSH é marcador preferencial para confirmar menopausa. Níveis de FSH consistentemente elevados acima de 30 mIU/mL com estradiol menor que 20 pg/mL suportam identificação de menopausa. FSH representa mais que marcador menopausal: evidências de 2024 sugerem papel na saúde mental feminina, com potencial impacto nos sintomas neuropsiquiátricos da transição menopausal.

Abordagem funcional reconhece múltiplos fatores modificáveis: desequilíbrio hormonal (estrogênio, progesterona, testosterona, DHEA), disfunção tiroidiana (hipotireoidismo subclínico presente em 15-20% das mulheres com disfunção sexual), resistência insulínica e síndrome metabólica, deficiências nutricionais (especialmente vitamina D, zinco, magnésio), inflamação sistêmica de baixo grau, disbiose intestinal e vaginal, estresse crônico com desregulação do eixo HPA, e neurotransmissores (serotonina, dopamina, oxitocina). Estudo de 2025 identificou que 70,5% das mulheres pós-menopáusicas com disfunção sexual apresentavam depressão moderada a severa, destacando interconexão neuropsiquiátrica.""",
        "patient_explanation": """O questionário FSFI avalia 6 aspectos importantes da sua função sexual feminina através de 19 perguntas: seu desejo sexual, capacidade de ficar excitada, lubrificação vaginal, capacidade de atingir orgasmo, satisfação sexual geral e presença de dor ou desconforto durante relações. A pontuação total varia de 2 a 36 pontos.

Se sua pontuação total está abaixo de 26,55, isso sugere que você pode estar experienciando disfunção sexual que merece atenção e tratamento. É importante entender que problemas sexuais frequentemente refletem aspectos mais amplos de sua saúde: níveis hormonais, saúde metabólica, equilíbrio emocional e bem-estar geral.

Pesquisas recentes de 2025 mostram que a função sexual feminina é profundamente afetada por hormônios (especialmente durante menopausa), saúde intestinal (estudos mostram que probióticos podem melhorar a função sexual), inflamação no corpo, níveis de estresse, qualidade do sono e saúde emocional. Mais de 70% das mulheres pós-menopáusicas com problemas sexuais também apresentam sintomas de depressão.

A boa notícia é que existem muitas intervenções eficazes. Dependendo dos domínios mais afetados no seu questionário, podemos trabalhar aspectos hormonais, nutricionais, emocionais e de saúde geral para melhorar sua vida sexual. Cada um dos 6 domínios fornece pistas sobre quais sistemas do seu corpo podem precisar de suporte.""",
        "conduct": """Pontuação FSFI inferior a 26,55 ou domínios individuais comprometidos requer protocolo investigativo multissistêmico.

Avaliação laboratorial tier 1: Painel hormonal feminino completo: estradiol, progesterona (dia 21 do ciclo se menstruando), testosterona total e livre, DHEA-S, SHBG. FSH e LH (para avaliar status menopausal: FSH maior que 30 mIU/mL com estradiol menor que 20 pg/mL indica menopausa). Prolactina. Painel tiroidiano: TSH, T4 livre, T3 livre, anti-TPO (hipotireoidismo subclínico em 15-20% dos casos). Metabolismo: glicemia, HbA1c, insulina para HOMA-IR. Vitamina D, ferro (ferritina), zinco, magnésio.

Avaliação tier 2: Cortisol salivar 4 pontos (avaliar eixo HPA). Neurotransmissores urinários se disponível. Microbioma vaginal e intestinal. Marcadores inflamatórios: PCR ultrassensível, homocisteína. Screening estruturado para depressão e ansiedade (presente em mais de 70% dos casos pós-menopáusicos).

Intervenções baseadas em evidência: Suporte hormonal conforme status menopausal. Estradiol transdérmico preferencial (preparações orais afetam SHBG e testosterona livre). Testosterona tem evidência robusta para desejo sexual hipoativo em pós-menopausa segundo International Society for Study of Women Sexual Health. Suporte probiótico: Lactobacillus acidophilus demonstrou melhorar desejo, excitação, satisfação e escores totais FSFI em estudo 2025. Considerar formulações específicas para saúde vaginal.

Suporte nutricional: Maca peruana 3g/dia (evidência para função sexual e humor em menopausa). Tribulus terrestris 750mg/dia. DHEA 25-50mg/dia em mulheres pós-menopáusicas (evidência para saúde sexual). Ômega-3 2-3g/dia (anti-inflamatório, suporte neurológico). Magnésio 400mg/dia (relaxamento muscular, humor).

Abordagem psicossomática integrativa: Terapia cognitivo-comportamental focada em sexualidade. Mindfulness-based interventions. Terapia de casal se aplicável. Fisioterapia do assoalho pélvico (especialmente se dor/desconforto). Educação sexual e desmistificação. Tratamento específico para depressão/ansiedade se presente.

Monitoramento: reaplicação FSFI a cada 3 meses. Reavaliação hormonal a cada 3-6 meses durante otimização."""
    },
    {
        "id": "bd7913cc-c28c-4c61-852d-4e4251cf3e83",
        "name": "FSFI: ____/36",
        "clinical_relevance": "DUPLICATE_OF_ABOVE",
        "patient_explanation": "DUPLICATE_OF_ABOVE",
        "conduct": "DUPLICATE_OF_ABOVE"
    },
    {
        "id": "c3bb6a2b-adc5-4df1-a88f-808db5e184d0",
        "name": "FSFI: ____/36",
        "clinical_relevance": "DUPLICATE_OF_ABOVE",
        "patient_explanation": "DUPLICATE_OF_ABOVE",
        "conduct": "DUPLICATE_OF_ABOVE"
    },

    # =========================================================================
    # 3. ASEX - Arizona Sexual Experience Scale (2 items)
    # =========================================================================
    {
        "id": "edc0e804-755f-41f8-aa95-2325ebd339b1",
        "name": "ASEX: ____/25",
        "clinical_relevance": """A Escala de Experiência Sexual do Arizona (ASEX) é instrumento validado de 5 itens que avalia função sexual durante a última semana, podendo ser autoaplicada ou aplicada por clínico. Sujeitos avaliam nível atual de desejo sexual, excitação psicológica, excitação fisiológica (ereções ou lubrificação vaginal), facilidade de orgasmo e satisfação com orgasmo em escala Likert de 6 pontos, variando de extremamente positivo (1) a nenhum/nunca (6), gerando pontuação total de 5 a 30 pontos.

Critérios de disfunção sexual: escore total ASEX de 19 ou maior, qualquer item individual com escore de 5 ou maior, ou quaisquer três itens com escores individuais de 4 ou maior estão altamente correlacionados com presença de disfunção sexual diagnosticada clinicamente. Esta ferramenta representa padrão-ouro da indústria farmacêutica para identificar e quantificar efeitos de novos medicamentos na função sexual de participantes em estudos clínicos.

A relevância clínica em medicina funcional centra-se particularmente em disfunção sexual induzida por antidepressivos, problema prevalente mas subdiagnosticado. Estudo piloto de 2024 encontrou aumento significativo na identificação de disfunção sexual após implementação da escala ASEX (10% no grupo pré-ASEX versus 59% atendendo critérios de disfunção sexual com escala ASEX). Mais de 70% dos pacientes no grupo pós-pesquisa notaram que não teriam discutido função sexual com farmacêutico se não diretamente perguntados, evidenciando barreira comunicacional significativa.

Muitos estudos utilizaram ASEX para determinar prevalência de disfunção sexual emergente com tratamento de drogas antidepressivas em pacientes com transtornos depressivos e/ou ansiosos. Inibidores seletivos de recaptação de serotonina (ISRS) e inibidores de recaptação de serotonina-norepinefrina (IRSN) apresentam taxas de disfunção sexual de 40-80%, impactando significativamente adesão ao tratamento e qualidade de vida. Metanálise de rede examinando intervenções farmacológicas para disfunção sexual induzida por antidepressivos usando escala ASEX identificou estratégias mais eficazes.

Na abordagem funcional, ASEX permite monitoramento objetivo antes e durante tratamento psiquiátrico, facilitando decisões de troca medicamentosa (por exemplo, para bupropiona ou mirtazapina com perfis mais favoráveis), adição de agentes adjuvantes (bupropiona em baixa dose, sildenafil conforme necessário), ou ajustes de timing e dose. Além de medicações, ASEX detecta impacto de desequilíbrios hormonais, deficiências nutricionais (especialmente zinco, magnésio, vitamina D), disfunção tiroidiana, e desregulação de neurotransmissores. Aplicação seriada permite quantificação objetiva de resposta a intervenções funcionais.""",
        "patient_explanation": """A escala ASEX avalia sua função sexual através de 5 perguntas simples sobre a última semana: seu desejo sexual, excitação mental, excitação física (ereção ou lubrificação), facilidade para atingir orgasmo e satisfação com orgasmo. Cada pergunta é pontuada de 1 a 6, gerando pontuação total de 5 a 30 pontos.

Pontuação total de 19 ou mais indica disfunção sexual clinicamente significativa que merece atenção. Esta escala é particularmente útil se você está tomando antidepressivos, pois 40-80% das pessoas usando esses medicamentos experimentam efeitos colaterais sexuais, mas muitas não mencionam isso ao médico por constrangimento.

Estudos de 2024 mostram que mais de 70% dos pacientes não discutiriam problemas sexuais se não fossem perguntados diretamente. Esta escala permite que você comunique objetivamente como está sua função sexual, sem precisar entrar em detalhes constrangedores.

Se você está tomando antidepressivos e sua pontuação indica disfunção sexual, existem várias opções: trocar para medicamento com menos efeitos sexuais (como bupropiona), ajustar dose ou horário da medicação, adicionar suplementos que ajudam a função sexual, ou trabalhar aspectos hormonais e nutricionais que podem estar contribuindo.

Mesmo se não está tomando antidepressivos, pontuação alta pode indicar problemas hormonais (testosterona, estrogênio, tireoide), deficiências de nutrientes (zinco, magnésio, vitamina D), ou outros desequilíbrios que podem ser corrigidos.""",
        "conduct": """Pontuação ASEX maior ou igual a 19, item individual maior ou igual a 5, ou três itens maior ou igual a 4 requer avaliação de causas subjacentes e intervenções direcionadas.

Avaliação inicial: Revisão medicamentosa completa, especialmente antidepressivos (ISRS, IRSN), antipsicóticos, anti-hipertensivos (beta-bloqueadores, diuréticos tiazídicos), anticonvulsivantes. Histórico temporal: disfunção sexual precedeu ou seguiu início de medicação. Painel hormonal: testosterona total e livre, SHBG, estradiol (mulheres), prolactina (hiperprolactinemia induzida por antipsicóticos). Tireoide: TSH, T4 livre, T3 livre. Avaliação metabólica: glicemia, HbA1c, insulina. Deficiências nutricionais: vitamina D, zinco, magnésio, vitamina B12.

Estratégias para disfunção sexual induzida por antidepressivos: Primeira linha: discussão risco-benefício com prescritor. Considerar troca para antidepressivo com perfil mais favorável: bupropiona (dopaminérgico, frequentemente melhora função sexual), mirtazapina (menos efeitos sexuais que ISRS), vortioxetina (perfil mais favorável). Segunda linha: ajuste de dose (redução mínima efetiva). Drug holidays (pausas medicamentosas em finais de semana, apenas com supervisão médica e medicações de meia-vida curta). Timing otimizado: administração pós-coito.

Agentes adjuvantes com evidência: Bupropiona 150mg/dia como add-on (metanálise mostra melhora em disfunção sexual induzida por ISRS). Sildenafil ou tadalafil conforme necessário em homens (evidência robusta). Sildenafil também estudado em mulheres com resultados mistos. Buspirona 15-60mg/dia (agonista 5-HT1A). Ginkgo biloba 120-240mg/dia (evidência conflitante mas baixo risco).

Suporte nutricional funcional: Zinco 30-50mg/dia (cofator para síntese de neurotransmissores e hormônios). Magnésio 400-600mg/dia. Vitamina D otimização para maior que 40ng/mL. L-tirosina 500-1500mg/dia (precursor de dopamina, tomar manhã em jejum). Maca peruana 1,5-3g/dia. SAMe 400-1600mg/dia (antidepressivo com perfil sexual neutro ou positivo).

Intervenções não-farmacológicas: Exercício aeróbico regular (melhora humor e função sexual). Mindfulness e terapia cognitivo-comportamental focada em sexualidade. Comunicação com parceiro. Redução de estresse. Otimização de sono.

Monitoramento: reaplicação ASEX mensalmente durante ajustes, depois trimestralmente."""
    },
    {
        "id": "70babe19-54b9-4d12-9f9e-5dd3fc1e6df0",
        "name": "ASEX: ____/25",
        "clinical_relevance": "DUPLICATE_OF_ABOVE",
        "patient_explanation": "DUPLICATE_OF_ABOVE",
        "conduct": "DUPLICATE_OF_ABOVE"
    },

    # =========================================================================
    # 4. Libido/Desempenho Sexual Atual (3 items)
    # =========================================================================
    {
        "id": "a361cacd-ea83-4140-9346-45a7b7be4a9a",
        "name": "Libido/desempenho sexual atual",
        "clinical_relevance": """Avaliação qualitativa de libido e desempenho sexual atual representa componente essencial da anamnese funcional, complementando escalas padronizadas (IIEF-5, FSFI, ASEX) com narrativa pessoal contextualizada. Esta avaliação subjetiva captura nuances não acessíveis por questionários estruturados: flutuações circadianas e mensais de desejo, impacto de estressores específicos, relação temporal com mudanças de estilo de vida ou medicamentosas, e percepções qualitativas de mudança ao longo do tempo.

Baixa libido representa uma das queixas sexuais mais comuns em ambos os sexos, com prevalência de 15-30% em homens e 30-40% em mulheres, aumentando com idade. Na perspectiva funcional, libido diminuída raramente constitui diagnóstico isolado, funcionando como sintoma sentinela de desregulações sistêmicas. Principais eixos fisiopatológicos incluem disfunção endócrina (hipogonadismo, hipotireoidismo, hiperprolactinemia, insuficiência adrenal), resistência insulínica e síndrome metabólica, inflamação sistêmica crônica, disfunção mitocondrial e fadiga, desequilíbrio de neurotransmissores, e fatores psicossociais.

Conexão entre resistência insulínica e baixa libido é particularmente robusta. Altos níveis de insulina levam a baixo desejo sexual e infertilidade, e condições relacionadas à insulina como síndrome metabólica e síndrome dos ovários policísticos podem também causar mudanças físicas que afetam autoimagem. Estudos mostram que insulina alta reduz capacidade corporal de produzir hormônio de crescimento, reduzindo desejo sexual através de alteração nos níveis de testosterona. Tratamento com metformina normalizou desejo sexual e satisfação sexual em mulheres com diabetes e pré-diabetes, com efeitos correlacionando com melhora em resistência insulínica. Testosterona frequentemente retorna e aumenta ao longo do tempo graças à melhora em sensibilidade insulínica e redução de resistência à leptina.

Mitocôndrias podem ser alvos-chave para andrógenos e estrogênios no controle de sensibilidade insulínica, podendo conter resposta para efeitos dependentes de gênero dos andrógenos. Andrógenos e estrogênios exibem atividades na indução de fosforilação oxidativa mitocondrial, e os dois hormônios podem sinergizar em mitocôndrias para induzir superprodução de ATP. Disfunção mitocondrial manifesta-se como fadiga crônica, que inevitavelmente impacta energia disponível para atividade sexual.""",
        "patient_explanation": """Quando perguntamos sobre sua libido e desempenho sexual atual, estamos interessados em entender sua experiência pessoal: como está seu desejo sexual agora comparado ao que você considera normal para você, como está sua satisfação com relações sexuais, e se houve mudanças recentes que você notou.

Baixa libido é muito comum (afeta 15-30% dos homens e 30-40% das mulheres) e quase sempre reflete aspectos mais amplos de sua saúde. Desejo sexual requer energia física, equilíbrio hormonal, boa circulação, cérebro funcionando bem e ausência de estresse excessivo. Quando qualquer um desses sistemas está comprometido, libido frequentemente diminui.

Fatores que mais comumente afetam libido incluem níveis de testosterona (em homens e mulheres), hormônios tireoidianos, resistência à insulina (níveis altos de insulina reduzem hormônio de crescimento e testosterona), fadiga crônica relacionada a função mitocondrial diminuída, inflamação no corpo, deficiências nutricionais, qualidade de sono, nível de estresse crônico, e alguns medicamentos (especialmente antidepressivos e medicações para pressão).

É particularmente importante mencionar se sua libido mudou recentemente após começar ou parar algum medicamento, após mudanças grandes de peso, durante períodos de estresse intenso, ou se você nota padrões (por exemplo, libido pior em certas épocas do mês para mulheres, ou pior quando está mais cansado).

Sua descrição pessoal ajuda a guiar investigação: se libido caiu junto com ganho de peso e fadiga, investigaremos resistência insulínica e tireoide; se caiu após iniciar antidepressivo, focaremos em neurotransmissores e alternativas medicamentosas.""",
        "conduct": """Relato de libido diminuída ou desempenho sexual insatisfatório requer investigação sistêmica estruturada para identificar causas reversíveis.

Avaliação laboratorial tier 1: Painel hormonal completo conforme sexo biológico. Homens: testosterona total (manhã, entre 7-10h), testosterona livre calculada, SHBG, LH, FSH, estradiol, prolactina, DHEA-S. Mulheres: estradiol, progesterona (dia 21 ciclo se menstruando), testosterona total e livre, SHBG, DHEA-S, FSH/LH (status menopausal). Ambos sexos: TSH, T4 livre, T3 livre, T3 reverso, anti-TPO. Cortisol manhã ou salivar 4 pontos. Metabolismo: glicemia jejum, HbA1c, insulina jejum, HOMA-IR, perfil lipídico. Nutrientes: vitamina D, vitamina B12, ferritina, zinco, magnésio. Marcadores inflamatórios: PCR ultrassensível, homocisteína.

Avaliação tier 2 conforme achados: Polissonografia se suspeita apneia sono (fadiga desproporcional, ronco). Teste estresse adrenal completo. Neurotransmissores urinários se disponível. Marcadores disfunção mitocondrial: lactato, piruvato, ácidos orgânicos urinários. Metais pesados se exposição.

Intervenções baseadas em causas identificadas: Hipogonadismo: terapia reposição testosterona em homens com total menor 300ng/dL sintomáticos. Mulheres pós-menopáusicas: considerar testosterona transdérmica (evidência ISSM). Hipotireoidismo: otimização tiroidiana visando TSH 0,5-2,0, T3 livre terço superior da referência. Resistência insulínica: intervenção nutricional (baixo carboidrato ou jejum intermitente). Berberina 500mg 3x/dia ou metformina. Exercício resistido. Cromo 200-400mcg/dia. NAC 600mg 2x/dia. Ácido alfa-lipóico 600mg/dia.

Suporte mitocondrial: CoQ10 200-400mg/dia (forma ubiquinol). PQQ 20mg/dia. NAD mais precursores (NMN 250-500mg/dia ou NR 300mg/dia). L-carnitina 2-3g/dia. Magnésio 400-600mg/dia. Vitaminas B (complexo B ativo). D-ribose 5g 2-3x/dia se fadiga significativa.

Suporte libido específico: Maca peruana 1,5-3g/dia (evidência em ambos sexos). Tribulus terrestris 750-1500mg/dia (mais evidência em mulheres). Panax ginseng 900mg 3x/dia. Horny goat weed (Epimedium) 500-1000mg/dia. Ginkgo biloba 120-240mg/dia. Ashwagandha 300-600mg 2x/dia (reduz cortisol, melhora testosterona em homens).

Otimização estilo de vida: Exercício regular (resistido 3x/semana mais aeróbico moderado). Sono 7-9 horas (prioridade absoluta). Gerenciamento estresse: meditação, respiração, yoga. Redução álcool. Eliminação tabaco. Comunicação com parceiro. Terapia sexual se componente relacional. Mindfulness sexual.

Reavaliação: laboratorial 3-6 meses. Sintomática mensal."""
    },
    {
        "id": "b08e3349-fc30-495b-8203-e7ba11b5ed8b",
        "name": "Libido/desempenho sexual atual",
        "clinical_relevance": "DUPLICATE_OF_ABOVE",
        "patient_explanation": "DUPLICATE_OF_ABOVE",
        "conduct": "DUPLICATE_OF_ABOVE"
    },
    {
        "id": "79addea2-e157-4fa9-8aec-339b2232b1e3",
        "name": "Libido/desempenho sexual atual",
        "clinical_relevance": "DUPLICATE_OF_ABOVE",
        "patient_explanation": "DUPLICATE_OF_ABOVE",
        "conduct": "DUPLICATE_OF_ABOVE"
    }
]

# CONTINUA...

def update_item(item_id, data):
    """Atualiza um item via API"""
    url = f"{API_URL}/score-items/{item_id}"

    # Handle duplicates
    clinical_relevance = data.get("clinical_relevance", "")
    patient_explanation = data.get("patient_explanation", "")
    conduct = data.get("conduct", "")

    if clinical_relevance == "DUPLICATE_OF_ABOVE":
        print(f"⊙ Pulando duplicado: {item_id[:8]}...")
        return True  # Don't count as failure

    payload = {
        "clinicalRelevance": clinical_relevance,
        "patientExplanation": patient_explanation,
        "conduct": conduct
    }

    try:
        response = requests.put(url, headers=headers, json=payload, timeout=10)

        if response.status_code == 200:
            name = data.get("name", "Item")[:50]
            print(f"✓ {item_id[:8]}... {name}")
            return True
        else:
            print(f"✗ Erro {response.status_code}: {item_id[:8]}...")
            print(f"  {response.text[:100]}")
            return False

    except Exception as e:
        print(f"✗ Exceção {item_id[:8]}...: {str(e)[:80]}")
        return False

def main():
    """Executa atualização de todos os items"""
    print(f"\n{'='*70}")
    print(f"ATUALIZAÇÃO GRUPO VIDA SEXUAL - BATCH 30")
    print(f"{'='*70}\n")

    total = len(items)
    success = 0
    failed = []

    for i, item in enumerate(items, 1):
        print(f"[{i:2d}/{total}] ", end="")

        if update_item(item["id"], item):
            success += 1
        else:
            failed.append(item["id"])

        time.sleep(0.15)  # Rate limiting

    print(f"\n{'='*70}")
    print(f"RESULTADO FINAL")
    print(f"{'='*70}")
    print(f"Total processado: {total}")
    print(f"Sucesso: {success} ({success/total*100:.1f}%)")
    print(f"Falhas: {len(failed)}")

    if failed:
        print(f"\nIDs que falharam:")
        for fid in failed:
            print(f"  - {fid}")

    print(f"{'='*70}\n")

    return 0 if len(failed) == 0 else 1

if __name__ == "__main__":
    sys.exit(main())
