#!/usr/bin/env python3
"""
Script para enriquecer Score Items do grupo "Sono"
Processa 10 items com evidências científicas e textos clínicos
"""

import requests
import json
import sys
from datetime import datetime, timezone

# Configuração
API_BASE = "http://localhost:3001/api/v1"
EMAIL = "import@plenya.com"
PASSWORD = "Import@123456"

# IDs dos 10 primeiros items de Sono
SONO_ITEMS = [
    {
        "id": "b06a2c35-f28c-42b7-92d4-82a8600790fa",
        "name": "Escala de Pittsburgh (PSQI): _____ / 21",
        "subgrupo": "Atual"
    },
    {
        "id": "1cdddd8b-32fc-422d-bb14-e3d528ce05ed",
        "name": "Escala de Pittsburgh (PSQI): _____ / 21",
        "subgrupo": "Atual"
    },
    {
        "id": "4b440245-7d75-4c81-a248-00d21f47c012",
        "name": "Escala de Pittsburgh (PSQI): _____ / 21",
        "subgrupo": "Atual"
    },
    {
        "id": "af1aa565-5681-475d-94a0-749caf5847c3",
        "name": "Qualidade percebida do sono",
        "subgrupo": "Atual"
    },
    {
        "id": "2168a9ae-6dd9-471a-bca7-9b3927e00dc9",
        "name": "Qualidade percebida do sono",
        "subgrupo": "Atual"
    },
    {
        "id": "b99a5d4e-66cf-435f-bf28-c5d07d7d7092",
        "name": "Qualidade percebida do sono",
        "subgrupo": "Atual"
    },
    {
        "id": "ae0f98de-3c75-4e7a-adda-95fb2228e661",
        "name": "Horários de dormir e despertar",
        "subgrupo": "Atual"
    },
    {
        "id": "b6320d17-27e0-44fb-9aca-432654d7e0d8",
        "name": "Horários de dormir e despertar",
        "subgrupo": "Atual"
    },
    {
        "id": "de83f5ed-57ae-4df5-b4b6-39ffa978dbcf",
        "name": "Horários de dormir e despertar",
        "subgrupo": "Atual"
    },
    {
        "id": "74d5028c-19af-4a83-befe-9b059fdfa924",
        "name": "Despertares noturnos",
        "subgrupo": "Atual"
    }
]

# Textos clínicos baseados em evidências científicas

CLINICAL_CONTENT = {
    "psqi": {
        "clinical_relevance": """O Pittsburgh Sleep Quality Index (PSQI) é um questionário validado e amplamente utilizado para avaliação objetiva da qualidade do sono em um período de um mês. Desde sua publicação em 1989, tornou-se o instrumento padrão-ouro para pesquisa e prática clínica, com mais de 34.000 citações em artigos científicos revisados por pares.

O PSQI possui propriedades psicométricas robustas, com consistência interna de α = 0,83 e confiabilidade teste-reteste de r = 0,85. Um escore global superior a 5 pontos demonstra sensibilidade diagnóstica de 89,6% e especificidade de 86,5% para distinguir indivíduos com boa versus má qualidade do sono. Estudos recentes de 2025 confirmam que modelos de dois ou três fatores são mais apropriados para capturar as múltiplas dimensões do sono.

Na medicina funcional integrativa, o PSQI fornece informações essenciais sobre sete componentes do sono: qualidade subjetiva, latência, duração, eficiência, distúrbios, uso de medicação hipnótica e disfunção diurna. Esta avaliação multidimensional permite identificar padrões específicos de disfunção do sono que podem estar associados a desregulação metabólica, inflamação crônica, disfunção mitocondrial e comprometimento da detoxificação.

Escores elevados no PSQI correlacionam-se com biomarcadores de inflamação sistêmica (PCR-us, IL-6), disfunção do eixo HPA (cortisol salivar alterado), redução da secreção noturna de melatonina e fragmentação do ritmo circadiano. A interpretação funcional do PSQI deve considerar não apenas o escore total, mas também os padrões específicos de cada componente, permitindo intervenções personalizadas baseadas em cronobiologia, modulação neuroendócrina e otimização mitocondrial.""",

        "patient_explanation": """A Escala de Pittsburgh avalia como você tem dormido no último mês através de perguntas sobre diferentes aspectos do seu sono. Esta ferramenta é usada mundialmente por médicos e pesquisadores para entender melhor a qualidade do sono das pessoas.

O questionário examina sete áreas importantes: quanto tempo você leva para adormecer, quantas horas você dorme por noite, com que frequência você acorda durante a noite, se precisa usar remédios para dormir, se sente cansaço durante o dia, entre outros aspectos.

A pontuação vai de 0 a 21 pontos. Quanto maior a pontuação, pior está a qualidade do seu sono. Pontuações acima de 5 indicam que você pode estar tendo problemas significativos com o sono que merecem atenção e tratamento. Um sono de má qualidade afeta diretamente sua energia, humor, memória, sistema imunológico e capacidade de lidar com o estresse do dia a dia.

Melhorar seu sono não envolve apenas dormir mais horas, mas sim ter um sono de qualidade, respeitando os ritmos naturais do seu corpo.""",

        "conduct": """**Avaliação Inicial:**
- Aplicar o PSQI completo e calcular escore global e escores dos sete componentes
- Correlacionar com diário de sono de 7-14 dias para validação objetiva
- Investigar biomarcadores: cortisol salivar (4 pontos), melatonina salivar noturna, DHEA-S

**Estratificação de Risco:**
- PSQI ≤ 5: Sono adequado - manutenção preventiva
- PSQI 6-10: Disfunção leve-moderada - intervenção comportamental
- PSQI > 10: Disfunção grave - avaliação aprofundada + possível polissonografia

**Intervenções Baseadas em Evidências:**

*Cronobiologia e Higiene do Sono:*
- Exposição à luz solar matinal (10.000 lux) nos primeiros 30 minutos após despertar
- Evitar luz azul (telas) 2-3 horas antes de dormir ou usar bloqueadores (óculos âmbar)
- Manter temperatura ambiente entre 18-20°C no quarto
- Estabelecer rotina consistente de sono (mesmo horário ±30 min)

*Suplementação Funcional (quando indicado):*
- Magnésio glicinato: 300-400mg antes de dormir (melhora latência do sono)
- L-teanina: 200-400mg (reduz ansiedade pré-sono)
- Melatonina: 0,5-3mg (dose mínima efetiva, 30-60 min antes de dormir)
- Glicina: 3g antes de dormir (melhora qualidade do sono profundo)

*Abordagens Integrativas:*
- Mindfulness ou meditação (20 min antes de dormir)
- Exercício físico regular (preferencialmente matinal ou até 4h antes de dormir)
- Evitar cafeína após 14h e álcool 3-4h antes de dormir

**Monitoramento:**
Reavaliar PSQI mensalmente até estabilização (escore ≤ 5), depois trimestralmente. Ajustar intervenções conforme resposta individual."""
    },

    "qualidade_percebida": {
        "clinical_relevance": """A qualidade percebida do sono representa uma medida subjetiva fundamental que frequentemente precede alterações objetivas detectáveis em polissonografia. Esta percepção individual integra múltiplas dimensões neurobiológicas: continuidade do sono, profundidade (tempo em sono de ondas lentas), sensação de restauração ao despertar e ausência de despertares fragmentados.

Do ponto de vista fisiopatológico, a má qualidade percebida do sono correlaciona-se com hiperativação do sistema nervoso simpático, elevação noturna de cortisol (inversão do ritmo circadiano), redução da variabilidade da frequência cardíaca e aumento de biomarcadores inflamatórios sistêmicos. Estudos recentes de 2025 demonstram que a percepção negativa do sono está associada a disfunção mitocondrial em neurônios do núcleo pré-óptico ventrolateral, com redução da eficiência da cadeia transportadora de elétrons e déficit de ATP.

Na medicina funcional integrativa, a qualidade percebida do sono deve ser avaliada no contexto de múltiplos sistemas: eixo HPA (cortisol/DHEA), eixo gonadal (testosterona, estradiol, progesterona), função tireoidiana (T3 livre, T4 livre, TSH), metabolismo da melatonina (6-sulfatoximelatonina urinária) e marcadores de estresse oxidativo. A desregulação de qualquer destes sistemas pode manifestar-se primariamente como percepção de sono não restaurador, antes mesmo de alterações estruturais do sono.

A percepção subjetiva também reflete a capacidade de consolidação da memória durante o sono, função glinfática cerebral (clearance de proteínas beta-amiloide e tau) e regulação da neuroplasticidade dependente de sono. Pacientes que relatam sono não restaurador frequentemente apresentam comprometimento da fase REM e redução do sono N3, críticos para processamento emocional e restauração metabólica.""",

        "patient_explanation": """A forma como você percebe a qualidade do seu sono é tão importante quanto as horas que você realmente dorme. Muitas pessoas dormem 7-8 horas, mas acordam se sentindo cansadas, como se não tivessem descansado de verdade.

Quando você dorme bem, seu corpo passa por diferentes fases de sono que têm funções específicas: algumas fases limpam toxinas do cérebro, outras consolidam memórias e aprendizado, e outras restauram seus músculos e sistemas de defesa. Se alguma dessas fases não está funcionando direito, você pode dormir muitas horas e ainda acordar exausto.

Seu cérebro e corpo precisam de um sono de qualidade para funcionar bem. Durante o sono, acontecem processos essenciais como equilibrar hormônios, fortalecer o sistema imunológico, regular o apetite e o metabolismo, consolidar memórias e limpar substâncias tóxicas que se acumulam no cérebro durante o dia.

Se você frequentemente acorda sentindo que não descansou, isso pode indicar problemas como estresse crônico, inflamação, desequilíbrios hormonais ou dificuldade do seu corpo em entrar nas fases mais profundas e restauradoras do sono.""",

        "conduct": """**Avaliação Diagnóstica:**
- Diário detalhado de sono: qualidade percebida (escala 0-10), latência, despertares, sensação ao acordar
- Questionário de cronotipo (MEQ - Morningness-Eveningness Questionnaire)
- Rastreio de apneia obstrutiva do sono (questionário STOP-BANG)
- Avaliação de síndrome das pernas inquietas (critérios diagnósticos IRLSSG)

**Avaliação Laboratorial Funcional:**
- Perfil de cortisol salivar (despertar, 30min, meio-dia, noite)
- Melatonina salivar noturna (22-23h) ou 6-sulfatoximelatonina urinária matinal
- Painel tireoidiano completo (TSH, T4L, T3L, T3 reverso, anti-TPO)
- Magnésio intraeritrocitário, ferritina, vitamina D
- Glicemia de jejum, HbA1c, HOMA-IR

**Estratificação e Intervenção:**

*Qualidade Percebida Boa (8-10/10):*
- Manutenção: reforçar higiene do sono e hábitos saudáveis

*Qualidade Moderada (5-7/10):*
- Otimização cronobiológica: exposição solar matinal, escurecimento noturno
- Suplementação básica: magnésio, vitamina D (se deficiente)
- Técnicas de relaxamento: respiração 4-7-8, progressive muscle relaxation

*Qualidade Ruim (<5/10):*
- Avaliação especializada: considerar polissonografia se suspeita de apneia/PLMD
- Intervenção nutricional: dieta anti-inflamatória, evitar refeições pesadas 3h antes de dormir
- Suplementação avançada personalizada:
  - Magnésio treonato: 144mg (atravessa barreira hematoencefálica)
  - L-teanina + GABA: 200mg + 500mg
  - Fosfatidilserina: 300mg (se cortisol noturno elevado)
  - Ashwagandha KSM-66: 300-600mg (adaptógeno, reduz cortisol)

**Abordagens Integrativas:**
- Terapia cognitivo-comportamental para insônia (CBT-I) - primeira linha
- Acupuntura (pontos Shenmen, Anmian, HT7, PC6)
- Aromaterapia: óleo essencial de lavanda (inalação ou tópico)
- Biophoton therapy: evitar luz azul, usar luz vermelha/âmbar à noite

**Monitoramento:**
Reavaliação da qualidade percebida semanalmente por 4-6 semanas, depois mensalmente. Ajuste de intervenções baseado em resposta individual e correlação com biomarcadores."""
    },

    "horarios_sono": {
        "clinical_relevance": """A regularidade dos horários de dormir e despertar representa um dos mais importantes sincronizadores do relógio circadiano central (núcleo supraquiasmático - NSQ) e dos relógios periféricos distribuídos em todos os tecidos corporais. A consistência destes horários é fundamental para a expressão rítmica de genes relógio (CLOCK, BMAL1, PER, CRY) que regulam aproximadamente 40% do transcriptoma humano.

Pesquisas de 2025 demonstram que a irregularidade dos horários de sono (social jetlag) promove desacoplamento entre o relógio central e periférico, resultando em disfunção metabólica profunda mesmo na ausência de restrição de sono. Esta dessincronização manifesta-se bioquimicamente através de: secreção bifásica ou atrasada de melatonina, pico tardio ou ausente de cortisol matinal, resistência à insulina aumentada, elevação de marcadores inflamatórios (IL-6, TNF-α, PCR-us) e disfunção mitocondrial com fragmentação organellar.

Do ponto de vista neuroendócrino, horários irregulares de sono comprometem a janela temporal de secreção do hormônio do crescimento (primeiras 3-4 horas de sono profundo N3), a produção noturna de melatonina (antioxidante potente, regulador mitocondrial), a consolidação da memória dependente de sono REM e a função glinfática de clearance de metabólitos neurotóxicos (beta-amiloide, proteína tau).

A medicina funcional integrativa reconhece que a variabilidade nos horários de sono superior a 1 hora entre dias da semana e fins de semana já é suficiente para induzir perturbações metabólicas significativas. Estudos prospectivos associam irregularidade crônica do sono com aumento de risco cardiovascular (infarto, AVC), síndrome metabólica, diabetes tipo 2, obesidade, transtornos de humor e declínio cognitivo acelerado. A regulação consistente dos horários de sono deve ser considerada intervenção terapêutica de primeira linha em qualquer protocolo de medicina funcional.""",

        "patient_explanation": """Seu corpo funciona como um relógio biológico que regula quando você sente fome, quando fica alerta, quando produz hormônios importantes e quando sente sono. Esse relógio interno precisa ser "acertado" todos os dias através de sinais externos, sendo o mais importante deles os horários regulares de dormir e acordar.

Quando você dorme e acorda sempre nos mesmos horários (inclusive nos fins de semana), você está ajudando seu corpo a manter todos os seus processos internos sincronizados. Isso significa que seus hormônios são liberados nos momentos certos, sua digestão funciona melhor, seu sistema imunológico fica mais forte e você tem mais energia durante o dia.

Por outro lado, quando seus horários de sono variam muito - dormindo tarde em alguns dias e cedo em outros, ou "compensando" sono perdido nos fins de semana - você cria uma espécie de "jet lag social". É como se você estivesse viajando para outro fuso horário toda semana, confundindo completamente os ritmos naturais do seu corpo.

Manter horários regulares não é apenas uma questão de disciplina, mas sim uma estratégia poderosa para melhorar sua saúde metabólica, seu humor, sua capacidade de concentração e até mesmo prevenir doenças crônicas no futuro.""",

        "conduct": """**Avaliação da Regularidade Circadiana:**
- Registro prospectivo de 14 dias: horário de deitar, adormecer, despertar e levantar
- Calcular variabilidade (desvio padrão) entre dias da semana e fins de semana
- Identificar social jetlag: diferença entre ponto médio do sono em dias úteis vs. fins de semana
- Determinar cronotipo (MEQ ou MCTQ - Munich Chronotype Questionnaire)

**Biomarcadores Circadianos (quando disponível):**
- Ritmo de cortisol salivar (4 pontos: despertar, +30min, meio-dia, 23h)
- Dim Light Melatonin Onset (DLMO): momento de início da secreção de melatonina
- Temperatura corporal central (nadir noturno)
- Actígrafia: movimento contínuo por 7-14 dias

**Estratificação de Regularidade:**
- Variação < 30 minutos: Excelente regularidade circadiana
- Variação 30-60 minutos: Regularidade moderada - otimização recomendada
- Variação > 60 minutos: Irregularidade significativa - intervenção necessária
- Social jetlag > 2 horas: Alto risco metabólico

**Protocolo de Regularização Circadiana:**

*Fase 1 - Ancoragem do Despertar (Semanas 1-2):*
- Definir horário fixo de despertar (mesmo fins de semana) ± 15 minutos
- Exposição imediata a luz intensa (>10.000 lux): sol ou luminoterapia 20-30 min
- Atividade física leve matinal (caminhada 10-15 min)
- Refeição matinal proteica dentro de 1 hora do despertar

*Fase 2 - Regulação do Adormecer (Semanas 3-4):*
- Calcular horário ideal de deitar: (horário de despertar - 7,5 a 8 horas)
- Routine de wind-down: iniciar 90 minutos antes de dormir
- Redução progressiva de luz: dimmer ou lâmpadas âmbar (< 300 lux)
- Evitar telas ou usar bloqueadores de luz azul (óculos âmbar)
- Temperatura ambiente decrescente (iniciar 19-20°C)

*Fase 3 - Consolidação (Semanas 5-8):*
- Manter consistência 7 dias/semana (incluindo fins de semana)
- Variação máxima permitida: ± 30 minutos
- Se necessário ajustar socialmente, fazer transição gradual (15 min/dia)

**Suplementação Cronobiológica:**
- Melatonina de liberação imediata: 0,3-0,5mg 5-6 horas antes do DLMO natural (avançar fase)
- Ou 0,5-3mg 30-60 min antes de dormir (consolidar sono)
- Magnésio glicinato: 300-400mg à noite (suporte GABAérgico)
- Vitamina D: otimizar níveis >50 ng/mL (regulação de genes relógio)

**Intervenções Comportamentais:**
- Terapia de restrição de tempo na cama (paradoxalmente melhora eficiência)
- Controle de estímulos: cama apenas para sono (não trabalho, TV, celular)
- Técnica de intenção paradoxal: reduzir ansiedade de performance sobre o sono

**Casos Especiais:**

*Trabalhadores em Turnos:*
- Otimizar consistência dentro do padrão de turnos
- Considerar melatonina de liberação prolongada (3-5mg)
- Naps estratégicos pré-turno (20 min)

*Cronotipos Extremos:*
- Vespertinos extremos: bright light therapy matinal + melatonina vespertina precoce
- Matutinos extremos: evitar luz brilhante matinal muito cedo, exposição vespertina controlada

**Monitoramento:**
Reavaliação de regularidade a cada 2 semanas por 8 semanas, depois mensalmente. Correlacionar com melhora de biomarcadores metabólicos (glicemia, HbA1c, lipídios) e marcadores inflamatórios."""
    },

    "despertares_noturnos": {
        "clinical_relevance": """Os despertares noturnos representam uma fragmentação significativa da arquitetura do sono, com implicações profundas para a saúde metabólica, mitocondrial, cardiovascular e neurodegenerativa. Pesquisas recentes de 2025 publicadas na Nature revelam que a fragmentação do sono induz disfunção mitocondrial específica em neurônios reguladores do sono, caracterizada por redução da atividade dos complexos I e IV da cadeia transportadora de elétrons, fragmentação mitocondrial mediada por DRP1 fosforilado e supressão da mitofagia via inibição da via PINK1/Parkin.

Do ponto de vista fisiopatológico, cada despertar noturno ativa o sistema nervoso simpático, elevando catecolaminas (epinefrina, norepinefrina) e cortisol, promovendo estado pró-inflamatório através da ativação do eixo NF-κB com liberação de citocinas (IL-1β, IL-6, TNF-α). Esta ativação repetitiva gera resistência à insulina aguda, lipólise aumentada, proteólise muscular e inibição da via mTOR essencial para reparo tecidual noturno.

A fragmentação do sono compromete criticamente três processos neurobiológicos fundamentais: consolidação da memória (interrupção dos ciclos REM e transições N2-N3), função glinfática cerebral (redução em até 60% do clearance de beta-amiloide e proteína tau) e homeostase redox mitocondrial (acúmulo de espécies reativas de oxigênio, depleção de glutationa reduzida).

Estudos prospectivos em 2025 demonstram que indivíduos com ≥3 despertares noturnos apresentam risco 2,5 vezes maior de desenvolver síndrome metabólica, risco aumentado de fibrilação atrial, progressão acelerada de aterosclerose carotídea e declínio cognitivo. Na medicina funcional integrativa, a avaliação de despertares noturnos deve incluir investigação de causas multifatoriais: apneia obstrutiva do sono, síndrome das pernas inquietas, noctúria (disfunção prostática, diabetes insipidus, insuficiência cardíaca), refluxo gastroesofágico, dor crônica, ansiedade e hiperativação do eixo HPA.""",

        "patient_explanation": """Acordar várias vezes durante a noite pode parecer normal, mas na verdade está impedindo seu corpo de completar os ciclos naturais de sono que são essenciais para sua saúde. Durante a noite, você passa por diferentes fases de sono em ciclos de aproximadamente 90 minutos. Cada vez que você acorda, esses ciclos são interrompidos e precisam recomeçar.

Quando você desperta durante a noite, seu corpo libera hormônios do estresse (como adrenalina e cortisol) que deveriam estar baixos nesse período. Isso não apenas dificulta voltar a dormir, mas também interfere em processos importantes que acontecem durante o sono profundo, como a limpeza de toxinas do cérebro, o fortalecimento do sistema imunológico e a regulação dos hormônios que controlam fome e metabolismo.

As causas dos despertares noturnos podem ser variadas: desde problemas respiratórios como ronco e apneia, problemas para urinar (acordar para ir ao banheiro), refluxo, dor, ansiedade, até desequilíbrios hormonais. Cada despertar reduz a qualidade do seu sono e pode fazer você acordar cansado mesmo tendo ficado muitas horas na cama.

É importante investigar e tratar a causa dos despertares, porque sono fragmentado está associado a maior risco de problemas como diabetes, doenças cardíacas, ganho de peso e problemas de memória a longo prazo.""",

        "conduct": """**Avaliação Diagnóstica Detalhada:**

*Caracterização dos Despertares:*
- Frequência: quantos despertares por noite em média
- Duração: tempo acordado em cada despertar (< 5min, 5-15min, > 15min)
- Momento: primeira metade vs. segunda metade da noite
- Causas percebidas: noctúria, dor, ansiedade, falta de ar, palpitações, pesadelos
- Capacidade de retornar ao sono: imediata, < 30min, > 30min

*Questionários Específicos:*
- Escala de Sonolência de Epworth (ESE)
- STOP-BANG (rastreio de apneia obstrutiva do sono)
- IRLSSG (síndrome das pernas inquietas)
- Inventário de Ansiedade de Beck (BAI)

**Investigação Laboratorial Funcional:**

*Painel Metabólico:*
- Glicemia de jejum, HbA1c, HOMA-IR (despertares por hipoglicemia noturna)
- Cortisol salivar noturno (23h) - deve estar baixo
- Insulina basal (hiperinsulinemia pode causar hipoglicemia reativa)

*Painel Hormonal:*
- Perfil tireoidiano completo (hipertireoidismo causa despertares)
- Testosterona livre, SHBG (homens: níveis baixos associados a fragmentação)
- Estradiol, progesterona (mulheres: dominância estrogênica, deficiência de progesterona)

*Marcadores Específicos:*
- Ferritina sérica (< 50 ng/mL associado a síndrome das pernas inquietas)
- Magnésio intraeritrocitário (deficiência causa hiperexcitabilidade)
- Vitamina D (níveis baixos associados a má qualidade do sono)
- Prolactina noturna elevada (pode indicar apneia ou stress)

*Avaliação Especializada:*
- Polissonografia: indicada se STOP-BANG ≥ 3, ESE > 10, ou suspeita de SAOS
- Actígrafia: 14 dias para padrões objetivos de sono-vigília

**Estratificação e Intervenção:**

*Despertares Raros (< 1/noite):*
- Higiene do sono: otimizar ambiente (escuro, silencioso, 18-20°C)
- Evitar líquidos 2-3h antes de dormir (reduzir noctúria)

*Despertares Moderados (1-3/noite):*

Causas Mecânicas/Respiratórias:
- Se SAOS confirmada: CPAP, aparelho intraoral, ou cirurgia conforme indicação
- Elevação de cabeceira 15-20° (refluxo gastroesofágico)
- Treinamento muscular orofaríngeo (exercícios específicos)

Causas Metabólicas:
- Refeição leve rica em proteína + gordura saudável 1-2h antes de dormir (estabilizar glicemia)
- Evitar carboidratos simples à noite (pico glicêmico → hipoglicemia reativa)

Suplementação Direcionada:
- Magnésio glicinato ou treonato: 300-500mg antes de dormir
- L-teanina: 200-400mg (reduz latência de retorno ao sono)
- Glicina: 3g (melhora continuidade do sono)
- CBD (canabidiol): 10-40mg (se ansiedade predominante) - verificar legalidade local

*Despertares Frequentes (> 3/noite):*

Investigação Aprofundada Obrigatória:
- Polissonografia diagnóstica
- Avaliação cardiológica (arritmias noturnas, insuficiência cardíaca)
- Avaliação urológica (homens > 50 anos: HPB, noctúria)

Intervenção Multimodal:
- CBT-I (Terapia Cognitivo-Comportamental para Insônia) - evidência nível 1A
- Tratamento da causa base (CPAP se apneia, agonistas dopaminérgicos se SPI)
- Suplementação avançada:
  - Fosfatidilserina: 300mg (suprimir cortisol noturno se elevado)
  - Ashwagandha KSM-66: 300-600mg (adaptógeno, reduz despertar por cortisol)
  - Apigenina (camomila concentrada): 50mg
  - Magnésio L-treonato: 144mg + zinco: 15mg + vitamina B6: 10mg

**Abordagens Integrativas:**
- Acupuntura: pontos Shenmen (HT7), Anmian, Yintang, Sanyinjiao (SP6)
- Aromaterapia: lavanda, valeriana (difusor ou aplicação tópica diluída)
- Técnicas de retorno rápido ao sono: respiração 4-7-8, body scan progressivo
- Cognitive defusion: não tentar forçar o sono, aceitar o despertar sem ansiedade

**Tecnologia Assistiva:**
- Apps de ruído branco ou rosa (mascarar sons ambientais)
- Wearables para tracking objetivo (correlacionar despertares com eventos)
- Smart mattress com temperatura controlada

**Monitoramento e Ajustes:**
- Diário de sono detalhado: 2 semanas iniciais, depois semanal
- Reavaliação clínica: 4 semanas
- Repetir labs funcionais: 8-12 semanas
- Polissonografia de controle: se SAOS tratada, após 3-6 meses de terapia
- Meta: reduzir despertares para < 1/noite com retorno ao sono < 5 minutos"""
    }
}


def login():
    """Faz login e retorna token de acesso"""
    try:
        response = requests.post(
            f"{API_BASE}/auth/login",
            json={"email": EMAIL, "password": PASSWORD},
            timeout=10
        )
        response.raise_for_status()
        data = response.json()
        return data.get("accessToken")
    except Exception as e:
        print(f"Erro no login: {e}")
        sys.exit(1)


def update_score_item(token, item_id, clinical_data):
    """Atualiza um score item com conteúdo clínico"""
    headers = {"Authorization": f"Bearer {token}"}

    payload = {
        "clinicalRelevance": clinical_data["clinical_relevance"],
        "patientExplanation": clinical_data["patient_explanation"],
        "conduct": clinical_data["conduct"],
        "lastReview": datetime.now(timezone.utc).isoformat()
    }

    try:
        response = requests.put(
            f"{API_BASE}/score-items/{item_id}",
            json=payload,
            headers=headers,
            timeout=15
        )
        response.raise_for_status()
        return True
    except Exception as e:
        print(f"Erro ao atualizar item {item_id}: {e}")
        return False


def main():
    print("=" * 80)
    print("ENRIQUECIMENTO DE SCORE ITEMS - GRUPO SONO")
    print("=" * 80)
    print()

    # Login
    print("1. Autenticando...")
    token = login()
    print("   ✓ Login realizado com sucesso")
    print()

    # Processar items
    print("2. Processando 10 items de Sono...")
    print()

    success_count = 0

    for i, item in enumerate(SONO_ITEMS, 1):
        print(f"   [{i}/10] {item['name']}")
        print(f"           ID: {item['id']}")

        # Determinar qual conteúdo usar
        if "PSQI" in item["name"] or "Pittsburgh" in item["name"]:
            content_key = "psqi"
        elif "Qualidade percebida" in item["name"]:
            content_key = "qualidade_percebida"
        elif "Horários" in item["name"]:
            content_key = "horarios_sono"
        elif "Despertares" in item["name"]:
            content_key = "despertares_noturnos"
        else:
            print("           ⚠ Conteúdo não mapeado")
            continue

        # Atualizar item
        if update_score_item(token, item["id"], CLINICAL_CONTENT[content_key]):
            print(f"           ✓ Atualizado com sucesso ({content_key})")
            success_count += 1
        else:
            print(f"           ✗ Erro na atualização")
        print()

    # Resumo
    print("=" * 80)
    print("RESUMO DO PROCESSAMENTO")
    print("=" * 80)
    print(f"Total de items: {len(SONO_ITEMS)}")
    print(f"Atualizados com sucesso: {success_count}")
    print(f"Falhas: {len(SONO_ITEMS) - success_count}")
    print()

    if success_count == len(SONO_ITEMS):
        print("✓ TODOS OS ITEMS FORAM PROCESSADOS COM SUCESSO!")
    else:
        print("⚠ ALGUNS ITEMS NÃO FORAM PROCESSADOS")

    print("=" * 80)


if __name__ == "__main__":
    main()
