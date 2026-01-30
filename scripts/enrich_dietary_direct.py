#!/usr/bin/env python3
"""
Script para enriquecer 20 Score Items de ALIMENTAÇÃO
Versão que gera conteúdo diretamente baseado em templates e conhecimento médico
"""

import subprocess
import sys
import os
import json
import requests

# Configurações
API_URL = "http://localhost:3001/api/v1"
EMAIL = "import@plenya.com"
PASSWORD = "Import@123456"

# IDs dos 20 itens a enriquecer
TARGET_ITEMS = [
    "dc29ca47-3fa1-4ada-8cb0-85e83bcd38eb",  # Gluten 1
    "503b91b1-50d5-41f8-bc2d-772822096033",  # Gluten 2
    "a4bd778b-63f1-4b5f-9435-08411b954ffe",  # Gluten 3
    "21128901-d9bb-422f-8a95-5d4f2202fa49",  # Histamina 1
    "e76cf533-e63a-4602-98ff-5dbfa54c6079",  # Histamina 2
    "2436e6d5-945f-4646-885d-8c2df6ed6da7",  # Histamina 3
    "2a12cebb-111e-492a-8388-55222be9c346",  # Histórico familiar alimentação 1
    "64f67026-3778-461f-b5eb-4674bbb46c58",  # Histórico familiar alimentação 2
    "207b3a03-01a4-4165-9b48-a1fd0c42ae35",  # Histórico familiar alimentação 3
    "2a2e420e-1a0c-44b4-bb86-67d46e33c572",  # Hábitos mãe 1
    "73050428-0624-4b9b-8ebe-f920606d2335",  # Hábitos mãe 2
    "f946af49-7962-4371-b56e-794fcfb1d505",  # Hábitos mãe 3
    "b8401d7b-edb1-4c06-8dbd-dfa3e8754e41",  # Hábitos pai 1
    "92df69af-cf17-43fc-8a18-251bc6b8ebdf",  # Hábitos pai 2
    "8f4bbe8c-41ee-45c0-8455-c758fcc22bb0",  # Hábitos pai 3
    "2cf7a504-b4ce-4b50-9471-30fe89b19758",  # Problemas alimentares 1
    "e75bec23-9bc7-4fdc-9911-c9c9a2f48411",  # Problemas alimentares 2
    "f267b2b2-8a63-4801-9029-166aabb83176",  # Problemas alimentares 3
    "1daa1598-27a4-4750-9330-dc9967c345c6",  # Qualidade alimentação parentes 1
    "b781e715-372d-4f05-b01e-4db68c05d8db",  # Qualidade alimentação parentes 2
]

# Conteúdos clínicos baseados em evidências MFI
CLINICAL_CONTENT = {
    "Gluten": {
        "clinicalRelevance": """A sensibilidade ao glúten representa um espectro de condições que vai além da doença celíaca clássica, incluindo a sensibilidade ao glúten não-celíaca (SGNC) e a alergia ao trigo. O glúten, proteína presente no trigo, centeio e cevada, pode desencadear respostas imunológicas e inflamatórias em indivíduos suscetíveis, mediadas por diferentes mecanismos fisiopatológicos.

Na doença celíaca, o glúten induz resposta autoimune mediada por linfócitos T contra a transglutaminase tissular, resultando em atrofia vilositária intestinal, má absorção e manifestações sistêmicas. A SGNC, por sua vez, não apresenta autoanticorpos ou atrofia vilositária característica, mas cursa com sintomas gastrointestinais e extraintestinais significativos, incluindo fadiga, cefaleia, dores articulares e manifestações neuropsiquiátricas. Estudos recentes sugerem que componentes além do glúten, como ATIs (inibidores de amilase-tripsina) e FODMAPs presentes no trigo, podem contribuir para a sintomatologia.

Um aspecto crucial é o efeito do glúten sobre a permeabilidade intestinal através da modulação da zonulina, proteína que regula tight junctions. A exposição ao glúten pode aumentar a permeabilidade intestinal mesmo em indivíduos não-celíacos, potencialmente contribuindo para translocação bacteriana, endotoxemia metabólica e inflamação sistêmica de baixo grau. Esta "leaky gut" tem sido associada a condições autoimunes, neurológicas (ataxia cerebelar, neuropatia periférica) e psiquiátricas.

A prevalência de doença celíaca é estimada em 1% da população geral, enquanto a SGNC pode afetar até 6-10%, embora dados epidemiológicos sejam ainda controversos. A identificação precoce é essencial, pois a exposição crônica ao glúten em indivíduos suscetíveis está associada a maior risco de doenças autoimunes (tireoidite, diabetes tipo 1), osteoporose, infertilidade e linfoma intestinal. Na medicina funcional integrativa, a avaliação da sensibilidade ao glúten é fundamental para o manejo de condições inflamatórias crônicas, autoimunidade e disfunções neurológicas.""",

        "patientExplanation": """O glúten é uma proteína encontrada no trigo, centeio e cevada. Algumas pessoas têm dificuldade em digerir ou tolerar o glúten, o que pode causar diversos sintomas no corpo. Isso não significa necessariamente ter doença celíaca (uma condição autoimune séria), mas pode indicar uma sensibilidade ao glúten.

Quando você come alimentos com glúten e é sensível a ele, pode sentir sintomas como inchaço, diarreia, dor abdominal, cansaço, dores de cabeça, dores nas articulações e até alterações de humor. Isso acontece porque o glúten pode aumentar a inflamação no intestino e no resto do corpo, além de afetar a "barreira" do intestino, permitindo que substâncias indesejadas passem para a corrente sanguínea.

Avaliar se você tem sensibilidade ao glúten é importante porque remover esses alimentos da dieta, quando indicado, pode melhorar significativamente sua qualidade de vida e prevenir problemas de saúde a longo prazo. O diagnóstico correto é essencial antes de fazer mudanças permanentes na alimentação.""",

        "conduct": """1. **Investigação Inicial - Anamnese Dirigida:**
   - Frequência e quantidade de consumo de produtos com glúten (pães, massas, bolos, cerveja)
   - Sintomas gastrointestinais: distensão, dor abdominal, diarreia, constipação, refluxo
   - Sintomas extraintestinais: fadiga crônica, cefaleia, artralgia, manifestações cutâneas (dermatite herpetiforme)
   - Sintomas neuropsiquiátricos: brain fog, ansiedade, depressão
   - História familiar de doença celíaca ou autoimunidade
   - Resposta prévia a retirada de glúten (se já testou)

2. **Investigação Laboratorial:**
   - **Primeira linha (com consumo de glúten):** Sorologia celíaca (IgA anti-transglutaminase tissular, IgA anti-endomísio, IgA total)
   - **Se sorologia negativa mas suspeita alta:** IgG anti-gliadina deamidada, teste genético HLA-DQ2/DQ8
   - **Marcadores de permeabilidade intestinal:** Zonulina sérica ou fecal (quando disponível)
   - **Marcadores inflamatórios:** PCR ultrassensível, VHS
   - **Avaliação nutricional:** Hemograma (anemia), ferritina, vitamina B12, folato, vitamina D, zinco (deficiências comuns em celíacos)

3. **Teste de Eliminação e Reintrodução (se sorologia negativa):**
   - Retirada completa de glúten por 4-6 semanas (período mínimo necessário)
   - Manter diário de sintomas detalhado (escala 0-10 para principais queixas)
   - Após período de eliminação, fazer teste de provocação oral controlado (reintrodução de glúten por 3-7 dias)
   - Avaliar recorrência de sintomas (geralmente ocorre em 6-48h)

4. **Intervenções Dietéticas:**
   - Se confirmada sensibilidade: dieta estritamente sem glúten (evitar contaminação cruzada)
   - Substituições: usar farinhas naturalmente sem glúten (arroz, milho, mandioca, amêndoas, coco)
   - Atenção a alimentos industrializados (verificar rótulos - glúten pode estar oculto em molhos, temperos, embutidos)
   - Evitar produtos "gluten-free" ultraprocessados (geralmente ricos em açúcar e gorduras trans)

5. **Suporte Nutricional e Reparação Intestinal:**
   - Probióticos multi-cepas: Lactobacillus e Bifidobacterium (10-50 bilhões UFC/dia)
   - L-glutamina: 5-10g/dia (reparação mucosa intestinal)
   - Enzimas digestivas: protease DPP-IV específica para glúten (uso transitório)
   - Ômega-3 EPA/DHA: 2-3g/dia (efeito anti-inflamatório)
   - Vitamina D: corrigir deficiência (2000-4000 UI/dia conforme níveis séricos)
   - Zinco: 30mg/dia se deficiente (importante para tight junctions)

6. **Monitoramento e Reavaliação:**
   - Primeira reavaliação: 4-6 semanas após início da dieta sem glúten
   - Avaliar resolução/melhora de sintomas (esperado >50% de melhora em SGNC)
   - Reavaliar marcadores laboratoriais: 3-6 meses (anticorpos celíacos, marcadores nutricionais)
   - Acompanhamento nutricional: garantir adequação da dieta, prevenir deficiências
   - Se doença celíaca confirmada: acompanhamento multidisciplinar (gastro + nutri), monitoramento anual vitalício

7. **Critérios de Sucesso:**
   - Melhora significativa de sintomas GI e sistêmicos
   - Redução de marcadores inflamatórios
   - Normalização de parâmetros nutricionais
   - Melhora de qualidade de vida (usar questionários validados)

8. **Sinais de Alerta:**
   - Perda de peso não intencional persistente
   - Anemia refratária a suplementação
   - Sintomas neurológicos progressivos (ataxia, neuropatia)
   - Ausência de resposta à dieta após 6 meses (considerar outras causas: SIBO, disbiose, parasitoses, doença inflamatória intestinal)"""
    },

    "Histamina": {
        "clinicalRelevance": """A intolerância à histamina (IH) é uma condição clínica subdiagnosticada caracterizada pelo acúmulo de histamina endógena e/ou exógena devido a déficit na capacidade de degradação ou excesso de produção. A histamina é uma amina biogênica com múltiplas funções fisiológicas: mediador de respostas alérgicas, neurotransmissor, regulador da secreção ácida gástrica e modulador imunológico.

O principal mecanismo fisiopatológico da IH envolve deficiência ou disfunção da enzima diamino oxidase (DAO), responsável pela degradação de histamina no trato gastrointestinal. Estudos demonstram que aproximadamente 1% da população apresenta IH, com maior prevalência em mulheres (80%) e frequentemente associada a condições que comprometem a mucosa intestinal (doença celíaca, SIBO, disbiose, doença de Crohn). Polimorfismos genéticos no gene AOC1 (que codifica a DAO) podem predispor a redução da atividade enzimática.

Além da via DAO intestinal, a histamina pode ser degradada pela histamina-N-metiltransferase (HNMT) hepática e neuronal. Deficiências em HNMT ou saturação desta via contribuem para acúmulo sistêmico. Fontes de histamina incluem: alimentos ricos (queijos envelhecidos, embutidos, fermentados, vinhos, peixes não-frescos), alimentos liberadores de histamina (chocolate, morangos, tomates, clara de ovo) e histamina endógena produzida por mastócitos e basófilos.

A síndrome de ativação mastocitária (MCAS) representa condição relacionada mas distinta, caracterizada por ativação inapropriada de mastócitos com liberação excessiva de mediadores, incluindo histamina, triptase, prostaglandinas e leucotrienos. Diferenciar IH de MCAS é crucial, pois o manejo difere significativamente.

As manifestações clínicas da IH são multissistêmicas e inespecíficas: gastrointestinais (diarreia, dor abdominal, distensão), cutâneas (urticária, prurido, flushing), respiratórias (rinorreia, congestão nasal, broncoespasmo), cardiovasculares (taquicardia, hipotensão, arritmias), neurológicas (cefaleia tipo enxaqueca, tontura, vertigem) e sistêmicas (fadiga, dismenorreia). A variabilidade sintomática dificulta o diagnóstico, frequentemente confundida com alergia alimentar IgE-mediada ou síndrome do intestino irritável.

Na medicina funcional integrativa, investigar IH é essencial em pacientes com sintomas multissistêmicos inexplicados, especialmente se exacerbados por alimentos específicos, álcool, ou em mulheres com piora peri-menstrual (flutuações hormonais afetam atividade DAO). O reconhecimento precoce permite intervenções dietéticas e terapêuticas que melhoram significativamente a qualidade de vida.""",

        "patientExplanation": """A histamina é uma substância natural presente no nosso corpo e em muitos alimentos. Ela tem funções importantes, como participar das defesas do organismo e regular processos no estômago e cérebro. Normalmente, nosso corpo possui enzimas que quebram e eliminam o excesso de histamina.

A intolerância à histamina acontece quando seu corpo não consegue quebrar a histamina adequadamente, seja porque produz pouca enzima DAO (responsável por degradar a histamina no intestino), ou porque você está consumindo muita histamina através dos alimentos. Isso faz com que a histamina se acumule no corpo, causando diversos sintomas.

Os sintomas podem incluir: dores de cabeça fortes (tipo enxaqueca), vermelhidão no rosto, coceira na pele ou urticária, congestão nasal, diarreia ou dor de barriga, coração acelerado, tontura e cansaço. Muitas vezes esses sintomas pioram depois de comer alimentos como queijos envelhecidos, vinho, embutidos, alimentos fermentados ou peixes não fresquíssimos.

Identificar a intolerância à histamina é importante porque fazer ajustes na alimentação e usar suplementos específicos pode eliminar ou reduzir muito esses sintomas incômodos.""",

        "conduct": """1. **Anamnese Detalhada - Reconhecimento de Padrões:**
   - Sintomas gastrointestinais: diarreia (especialmente pós-prandial), dor abdominal, distensão, náuseas
   - Sintomas cutâneos: urticária, prurido, flushing (rubor facial), angioedema
   - Sintomas respiratórios: rinorreia, congestão nasal, espirros, broncoespasmo
   - Sintomas cardiovasculares: taquicardia, palpitações, hipotensão ortostática, arritmias
   - Sintomas neurológicos: cefaleia (tipo enxaqueca), tontura, vertigem
   - Sintomas ginecológicos: dismenorreia, piora peri-menstrual
   - Correlação temporal: sintomas ocorrem 15min-2h após refeições ricas em histamina
   - Gatilhos: álcool (especialmente vinho tinto), queijos, embutidos, fermentados, chocolate
   - Uso de medicamentos inibidores de DAO: AINEs, alguns antibióticos (metronidazol), anti-hipertensivos (verapamil)

2. **Investigação Laboratorial:**
   - **Histamina plasmática:** Coleta após episódio sintomático (meia-vida curta, difícil interpretação)
   - **Metilhistamina urinária de 24h:** Mais estável, reflete carga histamínica total
   - **Atividade DAO sérica:** Disponível em laboratórios especializados (normal >80 U/mL; <40 U/mL sugere deficiência)
   - **Triptase sérica basal:** Para diferenciar MCAS (triptase >20 ng/mL sugere MCAS)
   - **IgE específicas:** Excluir alergia alimentar IgE-mediada verdadeira
   - **Investigar condições associadas:** Sorologia celíaca, teste respiratório para SIBO, calprotectina fecal (doença inflamatória intestinal), exame parasitológico

3. **Teste Diagnóstico Terapêutico - Dieta de Eliminação:**
   - **Fase de eliminação (4-6 semanas):**
     * Evitar alimentos ricos em histamina: queijos envelhecidos, embutidos/defumados, peixes não frescos, conservas, fermentados (chucrute, kimchi, kombucha), vinagre, molho de soja
     * Evitar alimentos liberadores de histamina: chocolate, morangos, tomates, clara de ovo, mariscos, álcool
     * Evitar alimentos bloqueadores de DAO: chá preto/verde, bebidas energéticas
     * Manter diário alimentar e de sintomas (escala 0-10)

   - **Fase de reintrodução (após melhora clínica):**
     * Reintroduzir um alimento por vez, a cada 3 dias
     * Avaliar recorrência de sintomas em 2-48h
     * Identificar limiar de tolerância individual

4. **Suplementação e Suporte Farmacológico:**
   - **DAO exógena:** Suplemento de diamino oxidase (1-2 cápsulas 15min antes de refeições)
   - **Vitamina C:** 500-1000mg/dia (cofator para degradação de histamina)
   - **Vitamina B6 (piridoxal-5-fosfato):** 50-100mg/dia (cofator DAO e HNMT)
   - **Cobre:** 1-2mg/dia se deficiente (cofator DAO)
   - **Zinco:** 30mg/dia (estabilizador de mastócitos)
   - **Quercetina:** 500-1000mg/dia (anti-histamínico natural, estabilizador mastocitário)
   - **Probióticos:** Cepas específicas baixas em histamina (Bifidobacterium infantis, B. longum, Lactobacillus rhamnosus) - EVITAR L. casei, L. bulgaricus (produtores de histamina)
   - **Anti-histamínicos:** Considerar H1 (loratadina, cetirizina) + H2 (famotidina) se sintomas severos

5. **Tratamento das Condições Subjacentes:**
   - **Se SIBO:** Tratamento com antibióticos/herbais, restauração da motilidade
   - **Se disbiose:** Probióticos, prebióticos seletivos, dieta anti-inflamatória
   - **Se doença celíaca:** Dieta estritamente sem glúten
   - **Se uso de medicamentos inibidores DAO:** Avaliar possibilidade de substituição com prescritor

6. **Orientações Dietéticas Práticas:**
   - Consumir alimentos frescos (histamina aumenta com fermentação/maturação)
   - Peixes e carnes: sempre fresquíssimos, congelar rapidamente após compra
   - Evitar sobras guardadas (histamina aumenta com tempo de armazenamento, mesmo refrigerado)
   - Preferir vegetais frescos ou congelados (evitar enlatados)
   - Cozinhar em pequenas porções, consumir imediatamente

7. **Monitoramento e Reavaliação:**
   - Primeira avaliação: 4 semanas após início da dieta de eliminação (esperado >50% melhora sintomas)
   - Reavaliação laboratorial: 3-6 meses (atividade DAO, marcadores inflamatórios)
   - Avaliar tolerância e flexibilização gradual da dieta conforme melhora
   - Acompanhamento nutricional: garantir adequação nutricional, prevenir restrição excessiva

8. **Critérios de Sucesso:**
   - Redução significativa (>50%) na frequência e intensidade de sintomas
   - Identificação clara de gatilhos alimentares
   - Melhora de qualidade de vida
   - Capacidade de manter dieta de forma sustentável

9. **Sinais de Alerta e Diagnósticos Diferenciais:**
   - Anafilaxia: encaminhar urgentemente para alergista (possível alergia IgE-mediada verdadeira)
   - Sintomas progressivos apesar de dieta: investigar MCAS (triptase basal e durante crise, cromogranina A, prostaglandina D2)
   - Sintomas não responsivos: considerar outras causas (síndrome carcinoide, feocromocitoma, mastocitose sistêmica)"""
    },

    "Programação Metabólica - Histórico Familiar": {
        "clinicalRelevance": """A programação metabólica, conceito central da hipótese DOHaD (Developmental Origins of Health and Disease), estabelece que exposições ambientais durante períodos críticos do desenvolvimento (gestação, primeira infância) podem induzir adaptações permanentes em sistemas fisiológicos, predispondo a doenças crônicas na vida adulta. O histórico familiar de alimentação representa um aspecto fundamental desta programação, operando através de mecanismos epigenéticos, culturais e comportamentais transgeracionais.

Estudos epidemiológicos clássicos, como o Dutch Hunger Winter e os estudos de fome chinesa, demonstraram que a desnutrição materna durante a gestação está associada a maior risco de obesidade, diabetes tipo 2, doenças cardiovasculares e distúrbios psiquiátricos na prole, mesmo décadas após o evento. Inversamente, a sobrenutrição materna (obesidade, diabetes gestacional) programa o feto para preferências alimentares por alimentos palatáveis (ricos em açúcar e gordura), alterações no sistema de recompensa dopaminérgico e maior suscetibilidade à obesidade e síndrome metabólica.

Os mecanismos epigenéticos incluem metilação do DNA, modificações de histonas e expressão de microRNAs, que alteram a expressão gênica sem modificar a sequência de DNA. Genes relacionados ao metabolismo de glicose (GLUT4, IRS1), regulação do apetite (leptina, NPY), adipogênese (PPARγ) e resposta ao estresse (eixo HPA) são particularmente suscetíveis a programação. Importante: estas marcas epigenéticas podem ser transmitidas transgeracionalmente (avós → pais → filhos), explicando padrões familiares de doença independentes da genética mendeliana clássica.

Além dos aspectos biológicos, o histórico familiar de alimentação estabelece padrões culturais e comportamentais que perpetuam hábitos alimentares inadequados. Crianças expostas precocemente a alimentos ultraprocessados desenvolvem preferências gustativas duradouras e menor aceitação de vegetais e alimentos integrais. O ambiente alimentar familiar (disponibilidade, frequência de refeições em família, modelagem parental) é preditor robusto de qualidade da dieta na vida adulta.

Na perspectiva da medicina funcional integrativa, investigar o histórico alimentar familiar permite: identificar padrões de risco transmitidos intergeracionalmente, compreender resistências e facilitadores para mudanças dietéticas, e implementar intervenções preventivas precoces. O conhecimento de que a programação metabólica é parcialmente reversível através de intervenções nutricionais, exercício físico e manejo do estresse oferece perspectiva terapêutica otimista, especialmente se iniciadas em períodos sensíveis (infância, adolescência, pré-concepção).""",

        "patientExplanation": """O histórico de alimentação da sua família - pais, irmãos, avós - tem uma influência muito maior na sua saúde do que muitas pessoas imaginam. Isso acontece por dois motivos principais: primeiro, pelos hábitos alimentares que você aprendeu e repetiu desde criança; segundo, por mudanças biológicas profundas que acontecem no corpo antes mesmo de você nascer.

Estudos científicos mostram que a alimentação da sua mãe durante a gravidez, e até mesmo a alimentação dos seus avós, pode "programar" seu corpo para reagir de certas formas aos alimentos. Por exemplo, se sua mãe passou fome ou teve diabetes na gravidez, isso pode ter alterado como seu corpo processa açúcar e gordura hoje, mesmo que você não saiba disso.

Além disso, os hábitos que você viu em casa - o que sua família comia, como cozinhavam, se comiam juntos à mesa - moldaram suas preferências e escolhas alimentares atuais. Entender esse histórico é muito importante porque ajuda a explicar dificuldades que você pode ter hoje (como tendência a ganhar peso, vontade exagerada de doces, dificuldade de comer vegetais) e permite que você tome decisões conscientes para mudar esses padrões, protegendo sua saúde e a das próximas gerações.""",

        "conduct": """1. **Anamnese Familiar Detalhada - Investigação Transgeracional:**
   - **Pais (especialmente mãe):**
     * Condições durante gestação: diabetes gestacional, hipertensão, pré-eclâmpsia, ganho de peso excessivo/insuficiente
     * Tipo de parto (cesárea vs normal) e amamentação (duração, exclusiva ou não)
     * Padrão alimentar familiar durante infância: frequência de ultraprocessados, fast-food, refrigerantes
     * Disponibilidade de frutas e vegetais em casa
     * Hábito de refeições em família vs alimentação individual/na frente da TV
     * Uso de comida como recompensa/punição ("se comportar ganha doce")

   - **Avós maternos e paternos:**
     * História de exposição a fome/escassez alimentar (guerras, crises econômicas)
     * Padrões culturais alimentares (tradições regionais, religiosas)
     * Doenças crônicas: diabetes, obesidade, doenças cardiovasculares, cânceres
     * Longevidade e causas de morte

   - **Irmãos:**
     * Padrões alimentares atuais (semelhantes ou diferentes do paciente)
     * Condições de saúde (obesidade, diabetes, doenças autoimunes)

2. **Correlacionar Histórico com Fenótipo Atual:**
   - Identificar possíveis programações metabólicas:
     * Diabetes gestacional materna → investigar resistência insulínica mesmo com peso normal
     * Obesidade materna/paterna → avaliar adipogênese, setpoint de leptina
     * Desnutrição fetal → avaliar fenótipo "thrifty" (propensão a ganho de peso com pequeno excesso calórico)

3. **Avaliação Laboratorial Direcionada:**
   - **Perfil metabólico:** Glicemia, insulina basal, HbA1c, curva glicêmica se indicado
   - **Perfil lipídico completo:** LDL, HDL, triglicerídeos, ApoB/ApoA
   - **Marcadores inflamatórios:** PCR ultrassensível, homocisteína
   - **Perfil hormonal:** Leptina, adiponectina (biomarcadores de adipogênese e resistência)
   - **Opcional:** Teste epigenético (metilação de genes específicos como LEP, POMC) - disponibilidade limitada

4. **Intervenções Nutricionais Personalizadas:**
   - Reconhecer que padrões familiares podem ter criado "programação" difícil de reverter rapidamente
   - Estabelecer objetivos realistas de mudança gradual (evitar dietas restritivas radicais que ativam resposta de "fome")
   - Focar em qualidade alimentar (comida de verdade, minimamente processada)
   - Envolver família quando possível (mudança de ambiente alimentar familiar)
   - Trabalhar preferências gustativas através de exposição repetida e preparos atrativos de vegetais

5. **Estratégias de Reprogramação Epigenética:**
   - **Nutrientes com impacto epigenético:**
     * Doadores de metil: folato (400-800mcg/dia), B12 (500-1000mcg/dia), colina (500mg/dia), betaína
     * Polifenóis: resveratrol, EGCG (chá verde), curcumina (modulam histona acetiltransferases)
     * Ômega-3 EPA/DHA: 2-3g/dia (modula expressão de genes inflamatórios e metabólicos)
     * Vitamina D: manter níveis 40-60ng/mL (regula >1000 genes)

   - **Jejum intermitente:** Considerar protocolos 16:8 ou 5:2 (induz autofagia, melhora flexibilidade metabólica, altera metilação de genes metabólicos)

   - **Exercício físico:** Atividade aeróbica + resistência (induz modificações epigenéticas benéficas em músculo esquelético e tecido adiposo)

   - **Manejo de estresse:** Meditação, yoga, biofeedback (reduz hiperativação do eixo HPA programada precocemente)

6. **Abordagem Psicológica e Comportamental:**
   - Psicoterapia para desconstruir relação emocional com comida estabelecida na infância
   - Técnicas de mindful eating (comer consciente)
   - Abordar crenças limitantes herdadas ("na minha família todos são gordos", "sempre tive compulsão por doces")
   - Empoderar paciente: programação é modificável, não é destino imutável

7. **Prevenção para Próxima Geração:**
   - **Se paciente está planejando gravidez:**
     * Otimização pré-concepção: normalização de peso, controle glicêmico, suplementação (folato, ômega-3)
     * Educação sobre importância da nutrição gestacional
   - **Se paciente tem filhos:**
     * Orientação sobre alimentação infantil (evitar ultraprocessados, amamentação prolongada, introdução alimentar adequada)
     * Modelagem parental (pais comem bem → filhos comem bem)

8. **Monitoramento Longitudinal:**
   - Reavaliações trimestrais no primeiro ano (mudança comportamental lenta)
   - Marcadores metabólicos a cada 6 meses
   - Avaliar não apenas perda de peso, mas melhora de parâmetros metabólicos (sensibilidade insulínica, perfil lipídico, marcadores inflamatórios)

9. **Critérios de Sucesso:**
   - Melhora de biomarcadores independente de perda de peso significativa
   - Mudança sustentável de padrão alimentar (não dieta temporária)
   - Melhora de relação emocional com comida
   - Transmissão de hábitos saudáveis para filhos (quebra do ciclo transgeracional)"""
    },

    "Programação Metabólica - Hábitos Parentais (Tabagismo, Etilismo, Drogas)": {
        "clinicalRelevance": """Os hábitos e vícios dos pais durante os períodos pré-concepcional, gestacional e perinatal exercem profundo impacto sobre a saúde da prole através de mecanismos epigenéticos, teratogênicos, neurodesenvolvimentais e transgeracionais. A evidência científica acumulada nas últimas décadas estabelece inequivocamente que exposições parentais a tabaco, álcool e outras substâncias psicoativas podem programar suscetibilidade a doenças metabólicas, neuropsiquiátricas, respiratórias e neoplásicas na descendência.

**Tabagismo:**
O tabagismo materno durante a gestação está associado a baixo peso ao nascer, prematuridade, síndrome da morte súbita infantil, e maior risco de obesidade infantil (paradoxalmente). Mecanismos incluem hipóxia fetal, disfunção placentária, exposição a mais de 4000 compostos químicos (nicotina, monóxido de carbono, hidrocarbonetos aromáticos policíclicos). A nicotina atravessa facilmente a placenta e o leite materno, afetando o desenvolvimento neurológico: filhos de mães fumantes apresentam maior risco de TDAH, distúrbios de aprendizagem, comportamentos disruptivos e dependência química na adolescência.

Evidências emergentes demonstram que o tabagismo **paterno** também é relevante. Alterações epigenéticas nos espermatozoides (metilação de DNA, modificações de histonas) podem ser transmitidas ao embrião. Estudos mostram que filhos de pais fumantes têm maior risco de asma, obesidade e até câncer, independentemente da exposição materna. O período pré-puberal masculino é particularmente sensível: tabagismo durante a espermatogênese pode alterar o epigenoma espermático.

**Etilismo:**
A síndrome alcoólica fetal (SAF) representa o espectro mais grave de exposição ao álcool gestacional, caracterizada por dismorfismo facial, retardo de crescimento e disfunção neurológica. Não existe dose segura de álcool na gestação. Mesmo consumo moderado está associado a déficits cognitivos sutis, problemas comportamentais e maior risco de dependência química futura. O álcool é teratogênico direto (interfere com diferenciação celular e migração neuronal) e induz estresse oxidativo, apoptose neuronal e alterações epigenéticas. O consumo paterno periconcepcional também foi associado a alterações espermáticas e maior risco de malformações congênitas.

**Outras Drogas:**
Cannabis, cocaína, opioides e outras substâncias cruzam a barreira placentária. Cannabis está associada a baixo peso ao nascer, parto prematuro e alterações neurocomportamentais (impulsividade, déficits atencionais). Cocaína causa vasoconstrição placentária, descolamento de placenta, malformações cardiovasculares e neurológicas. Opioides levam à síndrome de abstinência neonatal e alterações duradouras no sistema opioide endógeno da criança.

**Perspectiva Transgeracional:**
Estudos em modelos animais e coortes humanas sugerem que algumas alterações epigenéticas induzidas por exposições tóxicas podem ser transmitidas para a segunda (F2) e até terceira geração (F3), caracterizando herança epigenética transgeracional verdadeira. Isso significa que vícios dos avós podem influenciar a saúde dos netos, independentemente de exposição direta.

Na medicina funcional integrativa, investigar hábitos tóxicos parentais é crucial para compreender predisposições individuais, orientar intervenções preventivas em populações de risco e fornecer aconselhamento pré-concepcional para quebrar ciclos intergeracionais de doença.""",

        "patientExplanation": """Os hábitos dos seus pais - especialmente da sua mãe durante a gravidez - têm um impacto muito grande no desenvolvimento do bebê e na saúde ao longo da vida. Quando a mãe fuma, bebe álcool ou usa drogas durante a gravidez, essas substâncias passam para o bebê através da placenta e podem causar problemas sérios.

O cigarro diminui o oxigênio que chega ao bebê, pode fazer ele nascer menor e prematuro, e aumenta o risco de problemas respiratórios, de atenção (como TDAH) e até dependência química no futuro. O álcool é ainda mais perigoso: não existe quantidade segura durante a gravidez. Pode causar má-formação no rosto e no cérebro do bebê, atraso no desenvolvimento e dificuldades de aprendizado.

O interessante é que descobertas recentes mostram que os hábitos do pai também importam. Se o pai fuma ou bebe muito na época da concepção, pode afetar o espermatozoide de forma que influencia a saúde do filho. E mais surpreendente ainda: essas "marcas" podem passar até para os netos.

Entender esse histórico ajuda a explicar alguns problemas de saúde que você pode ter hoje e é fundamental se você está planejando ter filhos, pois mudanças nos hábitos antes da gravidez podem proteger a próxima geração.""",

        "conduct": """1. **Anamnese Direcionada - Investigação de Exposições Parentais:**
   - **Mãe durante gestação e lactação:**
     * Tabagismo: quantidade (cigarros/dia), período de exposição (qual trimestre), cessação (se ocorreu)
     * Etilismo: frequência, quantidade, padrão (binge drinking vs consumo regular)
     * Uso de drogas ilícitas: cannabis, cocaína, opioides, anfetaminas
     * Medicamentos psicotrópicos: antidepressivos, benzodiazepínicos, antipsicóticos
     * Exposição ocupacional a tóxicos: solventes, metais pesados, pesticidas

   - **Pai no período periconcepcional e durante gestação:**
     * Tabagismo (exposição passiva da gestante + alterações espermáticas)
     * Etilismo crônico
     * Uso de drogas
     * Exposição ocupacional (especialmente se trabalha com químicos, radiação)

   - **Ambiente doméstico:**
     * Exposição a fumo passivo (outros moradores fumantes)
     * Uso de álcool/drogas no ambiente familiar (mesmo se não pela gestante)

2. **Correlacionar Exposições com Fenótipo Clínico Atual:**
   - **Se exposição a tabaco materno:**
     * Investigar: asma, bronquite crônica, DPOC precoce
     * Avaliar: TDAH, distúrbios de aprendizagem, impulsividade
     * Rastrear: resistência insulínica, obesidade (paradoxal em crianças com baixo peso ao nascer)
     * Considerar: maior suscetibilidade a dependência de nicotina

   - **Se exposição a álcool gestacional:**
     * Avaliar: déficits cognitivos (memória, função executiva), QI
     * Investigar: dismorfismos faciais (se SAF completa: fissura palpebral curta, filtro nasal longo, lábio superior fino)
     * Rastrear: problemas comportamentais (impulsividade, dificuldades sociais)
     * Avaliar: maior risco de dependência de álcool/drogas

   - **Se exposição a outras drogas:**
     * Investigar malformações específicas conforme substância
     * Avaliar desenvolvimento neuropsicomotor

3. **Avaliação Laboratorial e Complementar:**
   - **Marcadores de estresse oxidativo e inflamação:**
     * PCR ultrassensível, homocisteína, 8-OHdG (dano oxidativo ao DNA)
   - **Avaliação de metais pesados:** (se exposição ocupacional parental relevante)
     * Chumbo, mercúrio, arsênio (cabelo, urina)
   - **Avaliação epigenética:** (se disponível e indicado)
     * Padrões de metilação de genes específicos (BDNF, NR3C1 - receptor de glicocorticoide)
   - **Avaliação neuropsicológica:** (se suspeita de déficits cognitivos/comportamentais)
     * Testes de QI, atenção, função executiva
   - **Avaliação respiratória:** (se exposição significativa a tabaco)
     * Espirometria, teste de broncoprovocação se suspeita de asma

4. **Intervenções Terapêuticas - Mitigação de Danos:**
   - **Suporte antioxidante intensivo:**
     * Vitamina C: 1000-2000mg/dia (neutraliza radicais livres)
     * Vitamina E: 400 UI/dia (proteção de membranas)
     * N-acetilcisteína (NAC): 600-1200mg/dia (precursor de glutationa, desintoxicação)
     * Ácido alfa-lipóico: 300-600mg/dia (antioxidante universal)
     * Glutationa lipossomal: 500mg/dia (se disponível)

   - **Suporte neurodesenvolvimento e neuroproteção:**
     * Ômega-3 DHA: 1000-2000mg/dia (essencial para cérebro, deficitário em filhos de fumantes)
     * Colina: 500mg/dia ou fosfatidilcolina (reparo de membranas, neurotransmissão)
     * Magnésio L-treonato: 144mg/dia (atravessa barreira hematoencefálica, melhora cognição)
     * Vitaminas B (B6, B9, B12): doadores de metil, reparo epigenético

   - **Suporte desintoxicação hepática:**
     * Sulforafano (brócolis): suplemento ou alimentar (induz enzimas de fase 2)
     * Cardo mariano (silimarina): 200-400mg/dia (hepatoproteção)
     * Cúrcuma (curcumina): 500-1000mg/dia (anti-inflamatório, induz detox)

   - **Modulação epigenética:**
     * Folato (metilfolato): 800-1000mcg/dia (doador de metil, reparo epigenético)
     * SAMe (S-adenosilmetionina): 400-800mg/dia (doador de metil universal)
     * Probióticos: cepas produtoras de butirato (modula histona acetilação)

5. **Intervenções Neurocomportamentais (se déficits presentes):**
   - Terapia cognitivo-comportamental (se TDAH, impulsividade)
   - Reabilitação neuropsicológica (se déficits executivos)
   - Terapia ocupacional (se atrasos motores)
   - Intervenções educacionais especializadas

6. **Manejo Respiratório (se asma/doença reativa das vias aéreas):**
   - Evitar exposição contínua a tabaco (fumo passivo)
   - Controle ambiental: ácaros, mofo, poluentes
   - Tratamento farmacológico conforme diretrizes (corticoides inalatórios, broncodilatadores)
   - Suplementação: Vitamina D (2000 UI/dia), Ômega-3, NAC

7. **Aconselhamento Pré-Concepcional (se paciente planeja ter filhos):**
   - **Quebrar ciclo transgeracional:**
     * Cessação absoluta de tabagismo (ambos os pais) pelo menos 3-6 meses antes da concepção
     * Abstinência completa de álcool e drogas
     * Desintoxicação: protocolo com antioxidantes, quelantes suaves se exposição a metais
     * Otimização nutricional: dieta anti-inflamatória, suplementação (folato, ômega-3, multivitamínico pré-natal)

   - **Janela de oportunidade:**
     * Educação sobre epigenética: hábitos saudáveis podem "resetar" marcas epigenéticas negativas
     * Espermatogênese dura ~74 dias: mudanças paternas 3 meses pré-concepção são cruciais
     * Oogenese: óvulos são formados na vida fetal, mas saúde materna periconcepcional modula qualidade ovular

8. **Monitoramento e Suporte Contínuo:**
   - Acompanhamento multidisciplinar: médico funcional, nutricionista, psicólogo
   - Reavaliações semestrais de parâmetros clínicos e laboratoriais
   - Suporte para cessação de vícios se paciente também desenvolveu dependências
   - Grupos de apoio/psicoterapia para lidar com histórico familiar complexo

9. **Educação e Empoderamento:**
   - Explicar: "Você não é culpado pelo que seus pais fizeram, mas pode tomar controle da sua saúde agora"
   - Epigenética é bidirecional: marcas negativas podem ser parcialmente revertidas
   - Intervenções precoces são mais efetivas, mas nunca é tarde para melhorar
   - Conhecimento permite prevenção para próxima geração (quebra de ciclo)

10. **Sinais de Alerta e Encaminhamentos:**
    - Déficits cognitivos significativos: encaminhar para neuropsicólogo/neurologista
    - Dependência química desenvolvida: encaminhar para especialista em adições
    - Malformações/síndromes genéticas: aconselhamento genético
    - Asma/DPOC severa: pneumologista"""
    }
}


def login():
    """Faz login e retorna o access token"""
    try:
        response = requests.post(
            f"{API_URL}/auth/login",
            json={"email": EMAIL, "password": PASSWORD},
            timeout=10
        )
        response.raise_for_status()
        return response.json()["accessToken"]
    except Exception as e:
        print(f"Erro no login: {str(e)}")
        return None


def get_score_item(token, item_id):
    """Busca um score item pelo ID"""
    headers = {"Authorization": f"Bearer {token}"}
    response = requests.get(
        f"{API_URL}/score-items/{item_id}",
        headers=headers,
        timeout=10
    )
    response.raise_for_status()
    return response.json()


def update_score_item(token, item_id, data):
    """Atualiza um score item"""
    headers = {
        "Authorization": f"Bearer {token}",
        "Content-Type": "application/json"
    }

    response = requests.put(
        f"{API_URL}/score-items/{item_id}",
        headers=headers,
        json=data,
        timeout=15
    )
    response.raise_for_status()
    return response.json()


def get_content_for_item(item_name):
    """Retorna o conteúdo clínico apropriado baseado no nome do item"""
    name_lower = item_name.lower()

    if "gluten" in name_lower:
        return CLINICAL_CONTENT["Gluten"]
    elif "histamina" in name_lower:
        return CLINICAL_CONTENT["Histamina"]
    elif any(term in name_lower for term in ["familiar", "parentes", "alimentação dos pais"]):
        return CLINICAL_CONTENT["Programação Metabólica - Histórico Familiar"]
    elif any(term in name_lower for term in ["mãe", "pai", "hábitos", "vícios", "tabagismo", "etilismo"]):
        return CLINICAL_CONTENT["Programação Metabólica - Hábitos Parentais (Tabagismo, Etilismo, Drogas)"]
    else:
        # Default para programação metabólica geral
        return CLINICAL_CONTENT["Programação Metabólica - Histórico Familiar"]


def main():
    print("=" * 80)
    print("ENRIQUECIMENTO DE 20 SCORE ITEMS DE ALIMENTAÇÃO")
    print("Conteúdo baseado em evidências de Medicina Funcional Integrativa")
    print("=" * 80)
    print()

    # Login
    print("🔐 Fazendo login...")
    token = login()
    if not token:
        print("❌ Falha no login")
        return 1
    print("✅ Login realizado\n")

    # Processar cada item
    results = []
    success_count = 0
    error_count = 0

    for idx, item_id in enumerate(TARGET_ITEMS, 1):
        print(f"\n{'=' * 80}")
        print(f"ITEM {idx}/{len(TARGET_ITEMS)}")
        print(f"{'=' * 80}")

        try:
            # 1. Buscar item
            print(f"📖 Buscando item...")
            item = get_score_item(token, item_id)
            item_name = item.get("name", "Sem nome")
            print(f"   Nome: {item_name}")
            print(f"   ID: {item_id}")

            # 2. Obter conteúdo apropriado
            print(f"\n📚 Selecionando conteúdo clínico...")
            content = get_content_for_item(item_name)

            cr_len = len(content.get("clinicalRelevance", ""))
            pe_len = len(content.get("patientExplanation", ""))
            cd_len = len(content.get("conduct", ""))

            print(f"   ✅ Conteúdo selecionado:")
            print(f"      - Clinical Relevance: {cr_len} caracteres")
            print(f"      - Patient Explanation: {pe_len} caracteres")
            print(f"      - Conduct: {cd_len} caracteres")

            # 3. Atualizar no banco
            print(f"\n💾 Salvando no banco de dados...")
            updated = update_score_item(token, item_id, content)
            print(f"   ✅ Item atualizado com sucesso!")

            success_count += 1
            results.append({
                "item_id": item_id,
                "name": item_name,
                "status": "success",
                "content_sizes": {
                    "clinical_relevance": cr_len,
                    "patient_explanation": pe_len,
                    "conduct": cd_len
                }
            })

        except Exception as e:
            print(f"   ❌ ERRO: {str(e)}")
            error_count += 1
            results.append({
                "item_id": item_id,
                "name": item.get("name", "Desconhecido") if 'item' in locals() else "Desconhecido",
                "status": "error",
                "error": str(e)
            })

    # Relatório final
    print(f"\n\n{'=' * 80}")
    print("RELATÓRIO FINAL")
    print(f"{'=' * 80}\n")

    print(f"✅ Sucesso: {success_count}/{len(TARGET_ITEMS)}")
    print(f"❌ Erros: {error_count}/{len(TARGET_ITEMS)}")
    print(f"📊 Taxa de sucesso: {(success_count/len(TARGET_ITEMS)*100):.1f}%")
    print()

    if error_count > 0:
        print("ERROS DETALHADOS:")
        for r in results:
            if r["status"] == "error":
                print(f"  - {r['name']} ({r['item_id']})")
                print(f"    Erro: {r.get('error', 'Desconhecido')}")
        print()

    # Salvar relatório JSON
    report_path = "/home/user/plenya/DIETARY-BATCH-20-REPORT.json"
    with open(report_path, "w", encoding="utf-8") as f:
        json.dump(results, f, indent=2, ensure_ascii=False)

    print(f"📄 Relatório JSON salvo em: {report_path}")
    print()

    # Sumário por categoria
    print("SUMÁRIO POR CATEGORIA:")
    categories = {}
    for r in results:
        if r["status"] == "success":
            name = r["name"]
            if "Gluten" in name:
                cat = "Glúten"
            elif "Histamina" in name:
                cat = "Histamina"
            elif "familiar" in name or "parentes" in name:
                cat = "Histórico Familiar Alimentação"
            elif "mãe" in name or "pai" in name:
                cat = "Hábitos Parentais (Vícios)"
            else:
                cat = "Outros"

            categories[cat] = categories.get(cat, 0) + 1

    for cat, count in sorted(categories.items()):
        print(f"  - {cat}: {count} itens")

    return 0 if error_count == 0 else 1


if __name__ == "__main__":
    sys.exit(main())
