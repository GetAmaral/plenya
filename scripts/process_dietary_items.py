#!/usr/bin/env python3
"""
Script para processar Score Items de Alimentação com conteúdo clínico baseado em evidências
"""

import requests
import json

API_URL = "http://localhost:3001/api/v1"
EMAIL = "import@plenya.com"
PASSWORD = "Import@123456"

# Conteúdo clínico baseado nas evidências coletadas das lectures MFI
DIETARY_ITEMS = {
    "Estratégia macro atual": {
        "clinicalRelevance": """A avaliação da estratégia macronutriente atual do paciente é fundamental na medicina funcional integrativa, pois permite identificar desbalanços que podem estar contribuindo para processos inflamatórios, resistência insulínica e disfunções metabólicas. Cada indivíduo possui necessidades únicas de carboidratos, proteínas e gorduras, determinadas por fatores como genética (incluindo polimorfismos como COMT e genes relacionados aos receptores de dopamina), nível de atividade física, estado metabólico e objetivos terapêuticos.

A análise da composição macronutriente permite identificar padrões problemáticos comuns, como o consumo excessivo de carboidratos refinados (farinhas e açúcares), que na perspectiva evolutiva representam uma introdução extremamente recente na dieta humana. Estudos mostram que a agricultura surgiu há apenas cerca de 10.000 anos, enquanto a farinha refinada existe há menos de 200 anos - um período insuficiente para adaptação genética. O consumo repetido de alimentos ultraprocessados, especialmente aqueles ricos em amido, está associado a resistência insulínica mesmo em pacientes com glicemia de jejum normal, podendo ser detectado através da curva glicêmica e insulinêmica.

A estratégia macronutriente ideal deve considerar não apenas a proporção de nutrientes, mas também sua qualidade, timing e forma de combinação. Por exemplo, iniciar refeições com vegetais fibrosos e combinar carboidratos com proteínas ou gorduras pode modular significativamente a resposta glicêmica. A individualização é essencial, pois não existe uma proporção universal ideal - atletas podem manter cetose mesmo com 100g de carboidratos diários devido à alta demanda energética, enquanto pacientes com resistência insulínica severa podem necessitar de restrição mais rigorosa.""",

        "patientExplanation": """A forma como você distribui os tipos de alimentos no seu dia - carboidratos, proteínas e gorduras - tem um impacto muito grande na sua energia, disposição e saúde. Não existe uma fórmula única que funcione para todos, pois cada pessoa é diferente.

Avaliar sua alimentação atual ajuda a entender se você está comendo muitos carboidratos (como pães, massas e doces), pouca proteína (carnes, ovos, legumes) ou gorduras inadequadas. Muitas vezes, problemas como cansaço, dificuldade para emagrecer, fome constante ou desejo exagerado por doces podem estar relacionados ao desequilíbrio entre esses nutrientes.

O importante é encontrar o equilíbrio certo para você, considerando sua rotina, seu corpo e seus objetivos de saúde. Pequenos ajustes na proporção desses alimentos podem trazer grandes melhorias no seu bem-estar.""",

        "conduct": """1. Realizar anamnese alimentar detalhada, incluindo recordatório de 24h e questionário de frequência alimentar, para mapear a distribuição macronutriente habitual do paciente.

2. Avaliar indicadores metabólicos através de exames laboratoriais: glicemia de jejum, insulina basal (ideal <6 µU/mL), curva glicêmica e insulinêmica quando houver suspeita de resistência insulínica, perfil lipídico completo incluindo LDL pequeno e denso, ApoB/ApoA.

3. Considerar teste genético para polimorfismos relevantes (COMT, DRD2, receptores de leptina) quando indicado, especialmente em casos de obesidade severa ou resposta inadequada a intervenções.

4. Estabelecer estratégia inicial baseada no perfil individual:
   - Para emagrecimento: 20-50g carboidratos/dia, proteína 1-1,5g/kg/dia, gordura até saciedade
   - Para terapia metabólica (câncer, epilepsia, autoimune): <20g carboidratos, 0,8-1g/kg proteína, 70-90% calorias de gordura
   - Para manutenção/saúde geral: ajustar conforme resposta individual

5. Implementar estratégias de modulação glicêmica: iniciar refeições com vegetais fibrosos, combinar carboidratos com proteínas/gorduras, priorizar alimentos de baixo índice glicêmico.

6. Considerar ciclagem de estratégias dietéticas (low-carb → mediterrânea → jejum intermitente → cetogênica) conforme o momento de vida e objetivos do paciente, evitando estagnação.

7. Monitorar resposta através de parâmetros clínicos (peso, circunferência abdominal, energia, saciedade) e laboratoriais, ajustando a estratégia a cada 4-6 semanas.

8. Educar o paciente sobre a importância da qualidade dos alimentos (comida de verdade vs ultraprocessados) e evitar dogmatismo, respeitando a individualidade bioquímica."""
    },

    "Livre": {
        "clinicalRelevance": """A abordagem de alimentação livre ou intuitiva representa uma estratégia não restritiva que se baseia na reconexão do indivíduo com seus sinais internos de fome, saciedade e preferências alimentares, afastando-se de regras rígidas e contagens calóricas. Esta abordagem, quando bem aplicada, pode ser particularmente benéfica para pacientes com histórico de dietas restritivas, transtornos alimentares ou relação disfuncional com a comida.

Do ponto de vista neuroquímico, restrições alimentares excessivas podem gerar desregulação dos sistemas de recompensa dopaminérgico e serotoninérgico, aumentando o risco de compulsões alimentares. Genes associados ao receptor de dopamina (DRD2) e polimorfismos como COMT rápido (que degrada dopamina rapidamente) podem intensificar a busca por prazer alimentar, tornando restrições rígidas contraproducentes para certos indivíduos.

A alimentação intuitiva não significa ausência de orientação nutricional, mas sim uma mudança de paradigma: ao invés de impor restrições externas, o foco está em restaurar a capacidade natural do organismo de autorregulação. Evidências mostram que esta abordagem pode melhorar marcadores de saúde mental, reduzir comportamentos de compulsão alimentar e, paradoxalmente, melhorar a escolha alimentar a longo prazo quando comparada a dietas restritivas cíclicas.

No entanto, é essencial reconhecer que a alimentação intuitiva só é efetiva após um processo de educação nutricional. Pacientes com resistência insulínica severa, síndrome metabólica ou condições que exigem controle glicêmico rigoroso podem não se beneficiar desta abordagem inicialmente, necessitando primeiro de intervenções estruturadas para restaurar a sensibilidade metabólica. A seleção adequada do paciente e o timing da intervenção são cruciais para o sucesso.""",

        "patientExplanation": """Alimentação livre significa aprender a confiar nos sinais do seu próprio corpo, comendo quando sente fome de verdade e parando quando está satisfeito, sem seguir regras rígidas ou contar calorias. É uma forma de fazer as pazes com a comida, saindo do ciclo de dietas restritivas que muitas vezes levam a compulsões e culpa.

Essa abordagem não significa comer tudo o que quiser sem limites, mas sim reconectar-se com as necessidades reais do seu corpo, aprendendo a diferenciar fome física de fome emocional. Com o tempo, você desenvolve uma relação mais saudável com a comida, escolhendo naturalmente alimentos que fazem você se sentir bem.

É importante entender que essa liberdade vem após aprender sobre alimentação saudável. Para algumas pessoas com problemas metabólicos específicos, pode ser necessário primeiro um período de orientação mais estruturada antes de adotar essa abordagem.""",

        "conduct": """1. Avaliar adequação do paciente para alimentação intuitiva: histórico de dietas restritivas, presença de transtornos alimentares (incluindo ortorexia), estado metabólico atual e capacidade de autorregulação.

2. Critérios de exclusão ou necessidade de abordagem híbrida: resistência insulínica severa, síndrome metabólica descompensada, diabetes tipo 2 descontrolada, condições que exigem restrições específicas (ex: epilepsia refratária em dieta cetogênica).

3. Realizar educação nutricional prévia focada em:
   - Reconhecimento de sinais de fome e saciedade
   - Diferenciação entre fome física e emocional
   - Qualidade nutricional dos alimentos (comida de verdade vs ultraprocessados)
   - Impacto de diferentes alimentos no bem-estar individual

4. Implementar os princípios fundamentais da alimentação intuitiva:
   - Rejeitar a mentalidade de dieta e culturas restritivas
   - Honrar a fome e respeitar a saciedade
   - Fazer as pazes com todos os alimentos, eliminando categorias de "proibido"
   - Desafiar a "polícia alimentar" (vozes críticas internas)
   - Descobrir o fator de satisfação nas refeições

5. Monitorar sinais de alerta: aumento significativo de consumo de ultraprocessados, descontrole alimentar persistente, piora de parâmetros metabólicos, ganho de peso não intencional.

6. Avaliar necessidade de suporte psicológico concomitante, especialmente em casos com componente emocional forte ou histórico de transtornos alimentares.

7. Medir desfechos além do peso: qualidade de vida, satisfação alimentar, redução de episódios de compulsão, melhora da relação com o corpo, marcadores de estresse (cortisol).

8. Considerar abordagem híbrida quando necessário: princípios de alimentação intuitiva combinados com orientações estruturadas pontuais (ex: incluir proteína em todas as refeições, priorizar vegetais).

9. Reavaliar periodicamente (a cada 3-6 meses) se a abordagem permanece adequada ou se ajustes são necessários baseados na evolução do paciente."""
    },

    "Low carb": {
        "clinicalRelevance": """A estratégia low carb (baixo carboidrato) caracteriza-se pela redução do consumo de carboidratos para aproximadamente 20-50g/dia ou 30-35% das calorias totais, com aumento proporcional de proteínas (20-35%) e gorduras (35-50%). Esta abordagem tem sólida fundamentação metabólica e extensa base de evidências para múltiplas condições clínicas.

O racional metabólico baseia-se na modulação hormonal: a redução de carboidratos diminui a secreção de insulina e aumenta o glucagon, promovendo lipólise (quebra de gordura) e melhorando a sensibilidade insulínica. Estudos clínicos demonstram benefícios significativos em pacientes com resistência insulínica, síndrome metabólica, obesidade, diabetes tipo 2, síndrome dos ovários policísticos (SOP), esteatose hepática não alcoólica e enxaqueca.

Em termos cardiometabólicos, a dieta low carb promove redução de triglicerídeos, aumento de HDL, diminuição de LDL pequeno e denso (mais aterogênico), redução da relação ApoB/ApoA e melhora de marcadores inflamatórios. A perda de peso observada é acompanhada de preservação de massa magra quando a ingestão proteica é adequada (1-1,5g/kg/dia), diferentemente de dietas hipocalóricas tradicionais que frequentemente resultam em perda muscular significativa.

A abordagem low carb pode servir como porta de entrada para protocolos mais restritivos (cetogênica) ou como estratégia de manutenção a longo prazo. A individualização é fundamental: sedentários geralmente necessitam de restrição mais acentuada (~20g/dia) para atingir benefícios metabólicos, enquanto indivíduos ativos podem tolerar até 100g/dia mantendo resultados. A qualidade dos carboidratos consumidos (vegetais fibrosos, baixo índice glicêmico) é tão importante quanto a quantidade. Contraindicações relativas incluem insuficiência renal avançada, gravidez/lactação (requer adaptações) e uso de certos medicamentos (insulina, anti-hipertensivos) que demandam ajuste posológico.""",

        "patientExplanation": """Low carb significa diminuir bastante a quantidade de carboidratos (pães, massas, arroz, açúcar) que você come todos os dias, substituindo por mais proteínas (carnes, ovos, peixes) e gorduras saudáveis (abacate, castanhas, azeite, manteiga). Isso ajuda seu corpo a usar a gordura armazenada como energia ao invés de depender tanto de açúcar.

Essa forma de comer pode ajudar muito quem tem resistência à insulina, diabetes, dificuldade para emagrecer, fígado gorduroso ou enxaqueca. Muitas pessoas relatam mais energia, menos fome entre as refeições e melhora no controle do peso. Também ajuda a reduzir a gordura da barriga e melhora os exames de colesterol e triglicerídeos.

No início você pode sentir um pouco de cansaço ou dor de cabeça enquanto seu corpo se adapta (geralmente 1-2 semanas), mas isso passa. É importante beber bastante água e consumir sal suficiente. Não é uma dieta de passar fome - você come até ficar satisfeito, mas escolhe alimentos diferentes dos habituais.""",

        "conduct": """1. Avaliar adequação do paciente para low carb: presença de resistência insulínica (insulina basal >6, curva insulinêmica alterada), síndrome metabólica, diabetes tipo 2, obesidade, SOP, esteatose hepática, enxaqueca crônica.

2. Investigar contraindicações: insuficiência renal (clearance <60ml/min considerar com cautela), gestação/lactação (possível mas requer supervisão rigorosa), histórico de transtornos alimentares, uso de insulina ou hipoglicemiantes orais (requer ajuste de dose).

3. Realizar exames basais: glicemia, insulina, HbA1c, perfil lipídico completo (incluindo LDL fracionado se disponível), função renal, função hepática, TSH, eletrólitos (especialmente em hipertensos).

4. Implementar transição gradual (exceto em indicações terapêuticas urgentes):
   - Semana 1-2: Retirar ultraprocessados, açúcares adicionados e farinhas refinadas
   - Semana 3-4: Reduzir carboidratos para ~100g/dia, aumentar vegetais fibrosos
   - Semana 5+: Ajustar para meta de 20-50g/dia conforme objetivos individuais

5. Estruturar o plano alimentar:
   - Carboidratos: Priorizar vegetais folhosos e fibrosos, evitar amidos
   - Proteínas: 1-1,5g/kg/dia de fontes variadas (carnes, ovos, peixes, laticínios se tolerados)
   - Gorduras: Preferir monoinsaturadas e saturadas naturais, evitar óleos de sementes ricos em ômega-6

6. Prevenir e manejar "keto flu" (sintomas de adaptação):
   - Hidratação: Mínimo 2L água/dia
   - Eletrólitos: Aumentar sal (1 colher de chá dissolvida em água), considerar potássio (sal light) e magnésio
   - Monitorar PA em hipertensos, ajustar anti-hipertensivos se necessário

7. Ajustar medicações conforme necessário:
   - Insulina/antidiabéticos: Redução pode ser necessária em dias, monitorar glicemia capilar
   - Anti-hipertensivos: Efeito diurético da dieta pode requerer redução de dose
   - Diuréticos: Considerar redução devido à perda de água inicial

8. Monitoramento a curto prazo (primeiras 4-8 semanas):
   - Sintomas: Energia, sono, fome, saciedade
   - Parâmetros clínicos: Peso semanal, circunferência abdominal quinzenal
   - Laboratoriais (4-6 semanas): Glicemia, insulina, HbA1c, perfil lipídico

9. Monitoramento a longo prazo (3-6 meses):
   - Avaliação de sustentabilidade e adesão
   - Considerar ciclagem com outras estratégias (jejum intermitente, cetogênica semanal)
   - Reavaliar necessidade de manutenção estrita vs flexibilização

10. Educar sobre sinais de alerta: hipoglicemia persistente, fadiga extrema não resolvida, queda significativa de cabelo (avaliar tireoide e ferro), piora paradoxal de lipídios (raro, avaliar qualidade das gorduras consumidas)."""
    },

    "Vegetariana": {
        "clinicalRelevance": """A dieta vegetariana, quando adequadamente planejada, pode ser nutricionalmente completa e oferecer benefícios significativos para a saúde, incluindo redução do risco de doenças cardiometabólicas, certos tipos de câncer e obesidade. A Academy of Nutrition and Dietetics reafirmou em janeiro de 2025 que padrões alimentares vegetarianos e veganos apropriadamente planejados são nutricionalmente adequados e podem promover benefícios de saúde a longo prazo.

Do ponto de vista da medicina funcional integrativa, a dieta vegetariana bem estruturada fornece alta densidade de fitoquímicos anti-inflamatórios, antioxidantes, fibras prebióticas e compostos bioativos que modulam positivamente a microbiota intestinal. Estudos mostram melhora da relação Bacteroidetes/Firmicutes e aumento da diversidade microbiana em vegetarianos, associados a redução de marcadores inflamatórios sistêmicos.

No entanto, a qualidade da dieta vegetariana é mais importante que sua classificação. Dietas vegetarianas ricas em alimentos ultraprocessados (substitutos vegetais industrializados, grãos refinados, açúcares) não conferem os mesmos benefícios de dietas baseadas em alimentos integrais. A personalização é essencial, pois estudos observacionais mostram resultados mistos, e a resposta individual varia significativamente mesmo dentro de famílias.

Aspectos críticos na prescrição de dieta vegetariana incluem garantir adequação de: proteína completa (combinação de leguminosas, grãos, nozes e sementes), vitamina B12 (suplementação mandatória em veganos, recomendada em ovolactovegetarianos), ferro (fontes vegetais com vitamina C para melhorar absorção), zinco, cálcio, vitamina D, iodo e ácidos graxos ômega-3 de cadeia longa. Para ômega-3, vegetarianos dependem de conversão de ALA (linhaça, chia) em EPA/DHA, processo ineficiente (taxa de conversão <10%), justificando suplementação com óleo de algas.

Contraindicações relativas incluem gestação/lactação sem acompanhamento (risco de deficiências), crianças em crescimento (requer monitoramento rigoroso), anemia ferropriva grave não responsiva, e condições que demandem alta densidade proteica (sarcopenia, desnutrição, pós-operatório).""",

        "patientExplanation": """Alimentação vegetariana é aquela que exclui carnes de qualquer tipo (boi, porco, frango, peixe), mas pode incluir ovos e laticínios dependendo do tipo. Quando bem planejada, essa forma de comer pode trazer vários benefícios para a saúde, como melhor controle do colesterol, pressão arterial e peso, além de reduzir inflamação no corpo.

Os alimentos vegetais (frutas, verduras, legumes, grãos, castanhas, sementes) são ricos em fibras, vitaminas, minerais e substâncias protetoras que ajudam a prevenir doenças. Mas é muito importante fazer escolhas inteligentes: comer alimentos de verdade, e não apenas produtos industrializados "vegetarianos".

Para garantir que não falte nenhum nutriente importante, é essencial variar muito os alimentos, combinar bem as proteínas vegetais (feijão com arroz, por exemplo) e, em alguns casos, usar suplementos, especialmente vitamina B12, ômega-3 e ferro. O acompanhamento profissional ajuda a evitar deficiências e garantir que você está recebendo tudo que precisa.""",

        "conduct": """1. Avaliar motivação e experiência prévia com alimentação vegetariana: ética, saúde, ambiental, religiosa. Identificar tipo: ovolactovegetariana, lactovegetariana, ovovegetariana, ou pretensão de transição para vegana.

2. Realizar avaliação nutricional e laboratorial basal:
   - Hemograma completo (avaliar anemia)
   - Ferritina, ferro sérico, saturação de transferrina
   - Vitamina B12 sérica e, idealmente, metilmalônico e homocisteína
   - Vitamina D (25-OH)
   - Zinco, cálcio, magnésio
   - Perfil lipídico, glicemia
   - TSH (iodo pode ser limitante em vegetarianos)
   - Proteínas totais, albumina

3. Estruturar plano alimentar baseado em alimentos integrais:
   - Proteínas: Combinar leguminosas (feijões, lentilha, grão-de-bico) + grãos integrais + nozes/sementes. Meta: 1-1,2g/kg/dia mínimo.
   - Ferro: Fontes vegetais (lentilha, espinafre, tofu, quinoa) + vitamina C na mesma refeição. Evitar chá/café junto com refeições.
   - Ômega-3: Linhaça, chia, nozes. Considerar suplemento de algas (DHA/EPA) 250-500mg/dia.
   - Cálcio: Vegetais folhosos escuros (couve, brócolis), tofu fortificado, leite vegetal fortificado (se não ovolacto). Meta: 1000-1200mg/dia.

4. Implementar suplementação essencial:
   - Vitamina B12: OBRIGATÓRIO. Cianocobalamina 50-100mcg/dia ou metilcobalamina 1000mcg 2-3x/semana
   - Vitamina D3: 1000-2000 UI/dia, ajustar conforme níveis séricos
   - Iodo: Sal iodado ou alga nori ocasional. Evitar excesso.
   - Considerar: Ferro (se ferritina <30), zinco (15-30mg/dia se deficiente), DHA/EPA de algas

5. Educar sobre armadilhas comuns:
   - Evitar excesso de produtos ultraprocessados "vegetarianos" (hambúrgueres veganos industrializados, queijos veganos processados)
   - Não substituir proteína animal por carboidratos refinados
   - Garantir variedade e não depender excessivamente de soja

6. Monitoramento periódico:
   - Primeiros 3 meses: Avaliação clínica (energia, disposição, sintomas GI)
   - 6 meses: Repetir labs (B12, ferritina, hemograma, vitamina D)
   - Anual: Reavaliação completa incluindo zinco, homocisteína

7. Populações especiais:
   - Gestantes/lactantes: Suplementação mais rigorosa (B12, ferro, DHA, folato), monitoramento intensivo
   - Crianças/adolescentes: Acompanhamento de crescimento, desenvolvimento neurológico, suplementação profilática
   - Idosos: Atenção especial a proteína (risco sarcopenia), B12 (absorção reduzida), vitamina D

8. Avaliar necessidade de ajustes em condições específicas:
   - Anemia: Ferro + vitamina C, investigar causas
   - Hipotireoidismo: Garantir iodo adequado, evitar excesso de alimentos goitrogênicos crus (crucíferas)
   - Atletas: Proteína aumentada (1,4-2g/kg), considerar creatina, beta-alanina se performance comprometida

9. Considerar teste terapêutico de qualidade alimentar: Período de 3 semanas removendo processados vegetarianos e focando em alimentos integrais para avaliar impacto em sintomas e bem-estar.

10. Respeitar individualidade: Reconhecer que nem todos prosperam em dieta vegetariana. Se após 6-12 meses com plano otimizado e suplementação adequada houver fadiga persistente, deficiências recorrentes ou piora de saúde, considerar reintrodução de fontes animais seletivas ou transição para padrões mais flexíveis (ex: pescetarianismo)."""
    },

    "Vegana": {
        "clinicalRelevance": """A dieta vegana (plant-based estrita) exclui todos os produtos de origem animal - carnes, laticínios, ovos, mel. Quando rigorosamente planejada e suplementada, pode ser nutricionalmente adequada e oferecer benefícios cardiometabólicos significativos, incluindo redução do risco de hipertensão, diabetes tipo 2, obesidade e certos cânceres. A posição atualizada da Academy of Nutrition and Dietetics (2025) confirma que dietas veganas apropriadamente planejadas são nutricionalmente adequadas para todos os estágios da vida.

Do ponto de vista fisiopatológico, dietas veganas baseadas em alimentos integrais promovem redução de inflamação sistêmica (marcadores como PCR, IL-6), melhora do perfil lipídico (redução de LDL e triglicerídeos), aumento da sensibilidade insulínica e modulação positiva da microbiota intestinal. A alta densidade de fitoquímicos (polifenóis, carotenoides, isoflavonas), fibras prebióticas e baixa carga de ácidos graxos saturados contribuem para esses efeitos.

Entretanto, a dieta vegana apresenta desafios nutricionais significativos que exigem conhecimento técnico aprofundado e adesão rigorosa à suplementação. Deficiências críticas incluem: vitamina B12 (ausente em fontes vegetais), vitamina D (limitada, especialmente D3), ômega-3 de cadeia longa DHA/EPA (conversão de ALA é <5%), ferro (biodisponibilidade reduzida da forma não-heme), zinco (fitatos reduzem absorção), cálcio (biodisponibilidade variável), iodo (especialmente em dietas sem sal iodado), e proteína completa (requer combinações estratégicas).

A qualidade da dieta vegana é determinante: versões baseadas em ultraprocessados (carnes vegetais industrializadas, queijos veganos, doces veganos) não conferem benefícios e podem ser deletérias. Estudos mostram grande heterogeneidade de resultados, com benefícios observados principalmente em dietas whole-food plant-based (WFPB).

Contraindicações relativas incluem: gestação/lactação sem supervisão especializada (risco de deficiências fetais/neonatais, especialmente B12 e DHA), crianças sem monitoramento rigoroso (risco de atraso de crescimento e desenvolvimento), anemia megaloblástica severa, condições que demandem alta densidade proteica (sarcopenia, caquexia, recuperação pós-operatória), e história de transtornos alimentares (ortorexia pode ser exacerbada). A medicina funcional integrativa reconhece que a individualidade bioquímica é determinante - alguns indivíduos prosperam em dietas veganas, enquanto outros desenvolvem fadiga crônica e deficiências apesar da suplementação adequada, possivelmente relacionado a polimorfismos genéticos que afetam conversão de nutrientes (ex: BCMO1 para conversão de betacaroteno em vitamina A, FADS1/FADS2 para conversão de ALA em DHA/EPA).""",

        "patientExplanation": """Dieta vegana é aquela que não inclui nenhum alimento de origem animal - nem carnes, nem ovos, nem leite, nem mel. Muitas pessoas escolhem esse tipo de alimentação por questões éticas, ambientais ou de saúde, e quando bem planejada, pode trazer benefícios como melhora do colesterol, pressão arterial e controle de peso.

Essa alimentação se baseia em frutas, verduras, legumes, grãos, castanhas, sementes e leguminosas (feijão, lentilha, grão-de-bico). Esses alimentos são muito ricos em fibras, vitaminas e substâncias protetoras que fazem bem ao corpo. Porém, alguns nutrientes essenciais não existem ou são muito difíceis de obter em quantidades adequadas apenas de fontes vegetais.

Por isso, é FUNDAMENTAL tomar suplementos, especialmente vitamina B12, e provavelmente ômega-3, vitamina D e ferro. Também é muito importante comer alimentos de verdade e não apenas produtos industrializados "veganos". O acompanhamento com nutricionista especializado e exames regulares são essenciais para garantir que você está recebendo todos os nutrientes necessários e que sua saúde está sendo protegida.""",

        "conduct": """1. Avaliação inicial extensiva - A dieta vegana requer avaliação e acompanhamento mais rigoroso que outras estratégias alimentares:
   - Motivação: Ética, saúde, ambiental. Avaliar firmeza da decisão e disposição para suplementação rigorosa.
   - Experiência prévia: Transição gradual (vegetariana → vegana) ou abrupta. Tempo de dieta vegana.
   - Conhecimento nutricional: Avaliar se paciente compreende necessidade de suplementação e combinações alimentares.

2. Bateria laboratorial basal COMPLETA (não negociável):
   - Hemograma completo com índices
   - Vitamina B12 sérica + Metilmalônico + Homocisteína (fundamental - B12 pode estar falsamente normal)
   - Ferritina + Ferro sérico + Saturação de transferrina
   - Vitamina D (25-OH)
   - Zinco
   - Cálcio + Magnésio
   - TSH, T4 livre (avaliar iodo indiretamente)
   - Proteínas totais, albumina, pré-albumina
   - Perfil lipídico completo
   - Glicemia, HbA1c
   - Opcional mas recomendado: Ômega-3 index, teste genético FADS1/FADS2 e BCMO1 se disponível

3. Estruturação do plano alimentar WFPB (Whole Food Plant-Based):
   - Proteínas (1,2-1,4g/kg/dia, maior que em omnívoros devido a menor biodisponibilidade):
     * Combinações estratégicas: Leguminosas + grãos integrais (feijão + arroz, lentilha + quinoa)
     * Fontes concentradas: Tofu, tempeh, edamame, seitan (glúten puro)
     * Complementares: Nozes, sementes (hemp, chia), manteigas de castanhas

   - Ferro (dobro da RDA - 16-18mg/dia para mulheres, 8-10mg para homens):
     * Fontes: Lentilhas, espinafre, tofu, quinoa, amaranto, sementes de abóbora
     * SEMPRE combinar com vitamina C (pimentão, brócolis, morango, kiwi) na mesma refeição
     * EVITAR chá preto/verde, café, cálcio em excesso junto com refeições ricas em ferro

   - Ômega-3:
     * ALA (precursor): Linhaça moída (2 colheres sopa/dia), chia, nozes
     * Suplemento DHA/EPA de algas: OBRIGATÓRIO 250-500mg DHA+EPA/dia

   - Cálcio (1200mg/dia):
     * Tofu com sulfato de cálcio, vegetais folhosos escuros (couve, brócolis, bok choy)
     * Leites vegetais fortificados (amêndoa, soja, aveia - verificar rótulo)
     * Sementes de gergelim, tahini, amêndoas

   - Iodo (150mcg/dia, cuidado com excesso):
     * Sal iodado (1/2 colher chá) OU
     * Alga nori ocasional (1-2x/semana, não diariamente - risco de excesso)

4. Protocolo de suplementação MANDATÓRIA:
   - Vitamina B12: INEGOCIÁVEL. Opções:
     * Cianocobalamina 50-100mcg/dia OU
     * Metilcobalamina 1000mcg 2-3x/semana OU
     * Injeção 1000mcg mensal (se má absorção oral)

   - Vitamina D3 (vegana derivada de líquen):
     * 1000-2000 UI/dia, ajustar conforme níveis séricos (alvo: 40-60 ng/mL)

   - DHA/EPA de algas:
     * 250-500mg combinados/dia (não ALA - conversão <5%)

   - Ferro (se ferritina <30 ng/mL):
     * Bisglicinato ferroso 25-50mg/dia em jejum com vitamina C, longe de cálcio

   - Zinco (se níveis <90 mcg/dL):
     * Picolinato ou bisglicinato de zinco 15-30mg/dia

   - Iodo (se TSH elevado ou dieta sem sal iodado):
     * Iodeto de potássio 150mcg/dia (não exceder)

   - Considerar:
     * Creatina 3-5g/dia (estoques musculares reduzidos em veganos)
     * Taurina 500mg/dia (condicionalmente essencial em veganos)
     * Carnosina/Beta-alanina se atleta

5. Educação intensiva sobre armadilhas:
   - "Junk food vegana": Oreos, salgadinhos, refrigerantes são veganos mas não saudáveis
   - Substituição 1:1: Não trocar carne por carboidrato refinado
   - Ultraprocessados vegetarianos: Hambúrgueres veganos, nuggets, queijos veganos industrializados são nutricionalmente pobres
   - Excesso de soja: Variar fontes proteicas, não depender exclusivamente de soja

6. Monitoramento rigoroso e frequente:
   - Mês 1-2: Avaliação clínica (energia, digestão, sintomas neurológicos)
   - Mês 3: Repetir B12 + metilmalônico + homocisteína (crucial)
   - Mês 6: Labs completos (hemograma, ferritina, B12, D, zinco)
   - Anual: Reavaliação completa + ômega-3 index + homocisteína

7. Populações de ALTÍSSIMO RISCO - requerem acompanhamento especializado intensivo:
   - Gestantes/lactantes:
     * B12: 200mcg/dia ou 2000mcg 2x/semana
     * DHA: 300-600mg/dia (crítico para desenvolvimento fetal/neonatal)
     * Ferro: Suplementação profilática frequentemente necessária
     * Folato: 400-600mcg/dia
     * Monitoramento mensal de B12 e homocisteína
     * Risco de defeitos de tubo neural e atraso desenvolvimento se deficiências não corrigidas

   - Lactentes e crianças:
     * NUNCA sem supervisão de pediatra e nutricionista especializado
     * Risco de atraso de crescimento, desenvolvimento neurológico comprometido
     * Suplementação desde nascimento (B12 via leite materno ou fórmula)
     * Monitoramento trimestral de crescimento e desenvolvimento

   - Idosos:
     * Absorção de B12 pode estar comprometida (considerar injetável)
     * Risco aumentado de sarcopenia - proteína aumentada (1,4-1,6g/kg)
     * Suplementação de creatina pode ser particularmente benéfica

8. Condições clínicas que requerem atenção especial:
   - Anemia: Investigar causa (B12 vs ferro vs folato). Tratar agressivamente.
   - Hipotireoidismo: Garantir iodo adequado mas evitar excesso de crucíferas cruas (goitrogênicas). Levotiroxina pode precisar ajuste.
   - Atletas: Proteína 1,6-2g/kg, creatina, beta-alanina, considerar BCAA se recuperação comprometida
   - Gestação/lactação: Ver item 7
   - Histórico de transtornos alimentares: Alto risco de ortorexia, avaliar motivação real

9. Sinais de alerta para intervenção imediata:
   - Fadiga progressiva não responsiva a suplementação
   - Sintomas neurológicos (formigamento, dormência, alterações cognitivas) - B12
   - Anemia refratária
   - Perda de cabelo significativa - ferro, zinco, proteína
   - Amenorreia em mulheres em idade reprodutiva
   - Perda de massa muscular - proteína, creatina
   - Infecções recorrentes - zinco, proteína

10. Considerar transição ou modificação se:
    - Deficiências recorrentes apesar de suplementação otimizada
    - Fadiga crônica não resolvida após 12 meses
    - Gestação/lactação com dificuldade de manter níveis adequados
    - Condições que demandem alta densidade nutricional (caquexia, pós-operatório, sarcopenia)
    - Testes genéticos revelam polimorfismos desfavoráveis (FADS1/2, BCMO1) com sintomas clínicos correlatos

11. Estratégia de "reset" se plano não funcionar:
    - Reavaliação completa das fontes alimentares e suplementos usados
    - Investigar má absorção (doença celíaca, SIBO, disbiose)
    - Avaliar compliance real (muitos pacientes não tomam suplementos conforme prescrito)
    - Considerar abordagem híbrida: "Plant-based+" com inclusão seletiva de fontes animais (ex: ovos, peixes pequenos, moluscos) se motivação permitir
    - Se motivação ética for absoluta e saúde comprometer, otimizar ao máximo e monitorar intensivamente, mas informar riscos"""
    }
}

def login():
    """Faz login e retorna o token"""
    response = requests.post(
        f"{API_URL}/auth/login",
        json={"email": EMAIL, "password": PASSWORD}
    )
    response.raise_for_status()
    return response.json()["accessToken"]

def get_score_items(token):
    """Busca os score items de Alimentação/Atual"""
    headers = {"Authorization": f"Bearer {token}"}
    # Buscar árvore completa
    response = requests.get(
        f"{API_URL}/score-groups/tree",
        headers=headers
    )
    response.raise_for_status()
    groups = response.json()

    # Encontrar grupo Alimentação
    alimentacao = None
    for group in groups:
        if group.get("name") == "Alimentação":
            alimentacao = group
            break

    if not alimentacao:
        raise Exception("Grupo 'Alimentação' não encontrado")

    # Encontrar subgrupo Atual
    atual = None
    for subgroup in alimentacao.get("scoreSubgroups", []):
        if subgroup.get("name") == "Atual":
            atual = subgroup
            break

    if not atual:
        raise Exception("Subgrupo 'Atual' não encontrado")

    return atual.get("scoreItems", [])

def update_score_item(token, item_id, content):
    """Atualiza um score item com conteúdo clínico"""
    headers = {
        "Authorization": f"Bearer {token}",
        "Content-Type": "application/json"
    }
    response = requests.put(
        f"{API_URL}/score-items/{item_id}",
        headers=headers,
        json=content
    )
    response.raise_for_status()
    return response.json()

def link_articles(token, article_ids, item_id):
    """Linka articles a um score item"""
    headers = {
        "Authorization": f"Bearer {token}",
        "Content-Type": "application/json"
    }
    results = []
    for article_id in article_ids:
        try:
            response = requests.post(
                f"{API_URL}/articles/{article_id}/score-items",
                headers=headers,
                json={"scoreItemIds": [item_id]}
            )
            response.raise_for_status()
            results.append({"article_id": article_id, "success": True})
        except Exception as e:
            results.append({"article_id": article_id, "success": False, "error": str(e)})
    return results

# Articles relevantes encontrados via grep
RELEVANT_ARTICLES = {
    "Estratégia macro atual": [
        "Emagrecimento_XVII",
        "Dieta_Cetogênica_-_Parte_I",
        "Carboidratos_IV",
        "Proteínas",
        "Ácidos_Graxos",
        "Introdução_a_Nutrição_Aplicada_a_Prática_Clínica"
    ],
    "Low carb": [
        "Dieta_Cetogênica_-_Parte_I",
        "Dieta_Cetogênica_-_Parte_II",
        "Emagrecimento_XVII",
        "Emagrecimento_XIII",
        "Emagrecimento_XV",
        "Emagrecimento_-_Parte_III",
        "Emagrecimento_-_Parte_IV",
        "Emagrecimento_-_Parte_VII",
        "Resistência_Insulínica",
        "Carboidratos_IV",
        "Ácidos_Graxos_Saturados_de_Cadeia_Longa_I",
        "Cardiologia_II",
        "Cardiologia_VIII"
    ],
    "Vegetariana": [
        "Entrevista_-_Dra._Uma_Naidoo",
        "Microbioma_Intestinal_II",
        "Microbioma_Intestinal_V",
        "Ácidos_Graxos_Poliinsaturados",
        "Proteínas",
        "Genética_e_Epigenética_I",
        "Genética_e_Epigenética_II"
    ],
    "Vegana": [
        "Entrevista_-_Dra._Uma_Naidoo",
        "Microbioma_Intestinal_II",
        "Suplementação_IV",
        "Ácidos_Graxos_Monoinsaturados_II",
        "Proteínas",
        "Bases_Metabólicas_das_Doenças_Crônicas_e_Gerenciamento_-_Submetilação"
    ],
    "Livre": [
        "Emagrecimento_XVII",
        "Psiquiatria_Metabólica_Funcional_Integrativa_AULA_17",
        "Entrevista_-_Dra._Uma_Naidoo"
    ]
}

def main():
    print("=== Processamento de Score Items - Alimentação ===\n")

    # Login
    print("1. Fazendo login...")
    token = login()
    print("   ✓ Login realizado com sucesso\n")

    # Buscar items
    print("2. Buscando score items...")
    items = get_score_items(token)
    print(f"   ✓ Encontrados {len(items)} items\n")

    # Processar cada item
    results = []
    for item in items:
        item_name = item["name"]
        item_id = item["id"]

        if item_name not in DIETARY_ITEMS:
            continue

        print(f"3. Processando '{item_name}'...")

        # Update content
        try:
            content = DIETARY_ITEMS[item_name]
            update_result = update_score_item(token, item_id, content)
            print(f"   ✓ Conteúdo atualizado")

            # Link articles (simulado - precisaríamos dos IDs reais dos articles)
            # article_links = link_articles(token, RELEVANT_ARTICLES.get(item_name, []), item_id)
            # print(f"   ✓ {len([a for a in article_links if a['success']])} articles linkados")

            results.append({
                "item": item_name,
                "id": item_id,
                "status": "success",
                "content_updated": True,
                "relevant_articles": len(RELEVANT_ARTICLES.get(item_name, []))
            })
            print()
        except Exception as e:
            print(f"   ✗ Erro: {str(e)}\n")
            results.append({
                "item": item_name,
                "id": item_id,
                "status": "error",
                "error": str(e)
            })

    # Relatório final
    print("\n=== RELATÓRIO FINAL ===\n")
    for result in results:
        print(f"Item: {result['item']}")
        print(f"  ID: {result['id']}")
        print(f"  Status: {result['status']}")
        if result['status'] == 'success':
            print(f"  Conteúdo atualizado: {result['content_updated']}")
            print(f"  Articles relevantes: {result['relevant_articles']}")
        else:
            print(f"  Erro: {result.get('error', 'Desconhecido')}")
        print()

if __name__ == "__main__":
    main()
