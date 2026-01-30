#!/usr/bin/env python3
"""
Enriquecimento de 20 items SONO - Parte 2
Items: Campos eletromagnéticos até Motivos
"""

import requests
import json
import sys
from typing import Dict, Any

API_BASE = "http://localhost:3001/api/v1"
LOGIN_EMAIL = "import@plenya.com"
LOGIN_PASSWORD = "Import@123456"

# 20 IDs dos items de SONO - Parte 2
ITEM_IDS = [
    "073e9b8b-f3b8-411d-996c-e4b6ae6b71e1",  # Campos eletromagnéticos
    "339da4f9-89ef-4990-bbc5-e15de8eb96be",  # Colchão
    "c9261b30-8c20-4a31-af68-1a702327b2e2",  # Despertares noturnos
    "7b24a7f0-7c7a-415c-a182-eba24bbcfd43",  # Dieta noturna
    "5a4f8b02-04b7-4815-a7d3-62844e1d3dae",  # Equipamento do sono
    "428052bb-0cbd-4ab2-977c-a6a39cedc3fa",  # Estimulantes noturnos (cafeína)
    "88f10871-8821-42db-bee4-01c31f2187f8",  # Higiene do sono
    "98a7ddf2-4c79-44c9-8724-bd5baf7c709f",  # Histórico medicamentos/suplementos
    "f0fed096-a535-4388-bb84-343c3b42b84c",  # Histórico familiar
    "2ce2bbb8-31d7-44e1-88b3-d7282aeb79a7",  # Hábitos cônjuge
    "b991b0a9-7b65-4fd8-87b3-ed45772c794f",  # Idade desfralde
    "d60fb723-4152-46ec-8056-4135326c1f8b",  # Iluminação
    "9ca5a90d-ed46-49f7-9e53-013004f6e3db",  # Infância
    "67185b74-1b99-4397-90e1-9971e7d24353",  # Janelas/claridade
    "2454de7e-70e5-45e4-83d7-0091b5847b49",  # Lençóis/cobertas
    "25aa750f-d458-4e39-b37c-b5a66502340a",  # Melhores épocas
    "b256e241-ac54-4043-82a3-0ec5087702bc",  # Motivos
    "df8df017-2f55-443a-8f14-683006b14adb",  # (reserva)
    "d3647336-4f00-4ca9-82cf-f55897d2a862",  # (reserva)
    "20f0c9f4-7cf2-4e7d-94ea-33e0916adbc7",  # (reserva)
]

# Conteúdo clínico para cada item
CLINICAL_CONTENT = {
    "073e9b8b-f3b8-411d-996c-e4b6ae6b71e1": {  # Campos eletromagnéticos
        "name": "Campos eletromagnéticos",
        "clinical_relevance": "Exposição a campos eletromagnéticos (CEM) de dispositivos eletrônicos próximos à cama pode interferir na produção de melatonina e na qualidade do sono. Celulares, tablets, Wi-Fi, relógios digitais e outros dispositivos emitem radiação não-ionizante que pode afetar os ritmos circadianos. A luz azul emitida por telas suprime a secreção noturna de melatonina, atrasando o início do sono. Recomenda-se manter dispositivos eletrônicos a pelo menos 1-2 metros da cama, usar modo avião em celulares durante a noite, e evitar uso de telas 2 horas antes de dormir.",
        "interpretation_guide": "Questionar sobre dispositivos eletrônicos no quarto: celular próximo à cama, carregadores, computadores, televisão, roteador Wi-Fi. Avaliar se o paciente usa celular como despertador (mantendo-o muito próximo). Investigar sintomas de hipersensibilidade eletromagnética (cefaleia, insônia, irritabilidade). Quantificar número de dispositivos no ambiente de sono e distância da cama.",
        "scientific_references": "Pall ML. Wi-Fi is an important threat to human health. Environ Res. 2018;164:405-416. Halgamuge MN. Pineal melatonin level disruption in humans due to electromagnetic fields and ICNIRP limits. Radiat Prot Dosimetry. 2013;154(4):405-416. Wood AW et al. Changes in human plasma melatonin profiles in response to 50 Hz magnetic field exposure. J Pineal Res. 1998;25(2):116-127.",
        "low_description": "Exposição mínima a CEM: sem dispositivos eletrônicos no quarto, celular em modo avião longe da cama, ausência de Wi-Fi no ambiente de sono.",
        "medium_description": "Exposição moderada: 1-2 dispositivos no quarto (ex: celular como despertador), carregadores conectados, uso ocasional de telas antes de dormir.",
        "high_description": "Exposição alta a CEM: múltiplos dispositivos eletrônicos ligados próximos à cama, celular sob o travesseiro, uso frequente de telas na cama, presença de roteador Wi-Fi no quarto."
    },
    "339da4f9-89ef-4990-bbc5-e15de8eb96be": {  # Colchão
        "name": "Colchão",
        "clinical_relevance": "A qualidade e adequação do colchão são fundamentais para suporte da coluna vertebral, alívio de pontos de pressão e conforto térmico durante o sono. Colchões inadequados (muito moles, muito duros, desgastados) causam microdespertares, dor musculoesquelética e fragmentação do sono. Recomenda-se substituição a cada 7-10 anos. A firmeza ideal varia: colchões mais firmes para dormidores de barriga ou costas, intermediários para laterais. Material também importa: espuma viscoelástica (adaptação), molas ensacadas (suporte), látex (hipoalergênico). Temperatura: espuma retém calor, molas arejam melhor.",
        "interpretation_guide": "Avaliar idade do colchão, firmeza percebida (muito mole/duro), presença de deformações ou afundamentos, conforto térmico (acordar com calor/frio). Questionar sobre dores matinais (cervical, lombar, pontos de pressão), melhor qualidade de sono em outros locais (hotéis, sofá). Identificar posição preferencial de sono para recomendar firmeza adequada. Investigar presença de ácaros/alergias relacionadas ao colchão.",
        "scientific_references": "Jacobson BH et al. Effect of prescribed sleep surfaces on back pain and sleep quality in patients diagnosed with low back and shoulder pain. Appl Ergon. 2010;42(1):91-97. Verhaert V et al. Ergonomics in bed design: the effect of spinal alignment on sleep parameters. Ergonomics. 2011;54(2):169-178. Radwan A et al. Effect of different mattress designs on promoting sleep quality, pain reduction, and spinal alignment. Sleep Med. 2015;16(1):1-10.",
        "low_description": "Colchão adequado: menos de 5 anos de uso, firmeza apropriada para posição de sono, sem deformações, conforto térmico ideal, ausência de dores relacionadas.",
        "medium_description": "Colchão aceitável: 5-10 anos de uso, firmeza não ideal mas tolerável, pequenas deformações, desconforto ocasional, dores leves ao acordar.",
        "high_description": "Colchão inadequado: mais de 10 anos, deformações visíveis, firmeza inapropriada (muito mole/duro), desconforto térmico, dores matinais frequentes, sono melhor em outros locais."
    },
    "c9261b30-8c20-4a31-af68-1a702327b2e2": {  # Despertares noturnos
        "name": "Despertares noturnos",
        "clinical_relevance": "Despertares noturnos (awakening, WASO - Wake After Sleep Onset) fragmentam a arquitetura do sono, impedindo progressão adequada pelos estágios e reduzindo sono REM e profundo. Causas: noctúria, apneia do sono, refluxo gastroesofágico, dor crônica, ansiedade, medicamentos, distúrbios do movimento (síndrome das pernas inquietas). Mais de 2-3 despertares por noite com dificuldade de retornar ao sono indicam insônia de manutenção. Avaliar duração dos despertares, capacidade de retornar ao sono, horários (primeira/segunda metade da noite), sintomas associados.",
        "interpretation_guide": "Quantificar número de despertares por noite, duração média (segundos/minutos/horas), horários típicos. Investigar causas: acordar para urinar (quantas vezes), engasgos/roncos (apneia), azia (refluxo), dor, ansiedade/pesadelos, ruídos externos, calor/frio. Avaliar latência para retornar ao sono. Diferenciar despertares conscientes (insônia) de microdespertares inconscientes (fragmentação do sono). Questionar sobre uso de medicamentos ou álcool como 'facilitadores'.",
        "scientific_references": "Stepanski E et al. The effect of sleep fragmentation on daytime function. Sleep. 2002;25(3):268-276. Bonnet MH, Arand DL. Clinical effects of sleep fragmentation versus sleep deprivation. Sleep Med Rev. 2003;7(4):297-310. Ohayon MM et al. Meta-analysis of quantitative sleep parameters from childhood to old age in healthy individuals. Sleep. 2004;27(7):1255-1273.",
        "low_description": "Despertares mínimos: 0-1 despertar por noite, rápido retorno ao sono (< 5 minutos), sem impacto na qualidade do sono, sensação de sono contínuo.",
        "medium_description": "Despertares moderados: 2-3 despertares, retorno ao sono em 5-20 minutos, algumas noites de sono fragmentado, leve impacto diurno.",
        "high_description": "Despertares frequentes: 4+ despertares por noite, dificuldade de retornar ao sono (> 20 minutos), sono fragmentado crônico, cansaço diurno importante, causas identificáveis (noctúria, apneia, dor)."
    },
    "7b24a7f0-7c7a-415c-a182-eba24bbcfd43": {  # Dieta noturna
        "name": "Dieta noturna",
        "clinical_relevance": "Alimentação próxima ao horário de dormir afeta qualidade do sono. Refeições pesadas, gordurosas ou picantes (< 3h antes) causam desconforto digestivo, refluxo gastroesofágico e fragmentação do sono. Açúcares elevam glicemia/insulina, gerando picos energéticos seguidos de hipoglicemia noturna (despertar). Ideal: jantar leve 3-4h antes, evitar cafeína 6h antes, limitar líquidos 2h antes (reduzir noctúria). Alimentos promotores do sono: triptofano (peru, banana, nozes), magnésio (amêndoas, espinafre), carboidratos complexos (aveia). Evitar: frituras, álcool, chocolate, alimentos ácidos.",
        "interpretation_guide": "Avaliar horário do jantar em relação ao sono, composição da refeição (leve/pesada, gordurosa, picante), presença de refluxo/azia noturna. Questionar lanches noturnos (tipo, frequência, motivação - fome real vs ansiedade). Identificar consumo noturno de cafeína (café, chá preto/mate, refrigerantes, chocolate), álcool (quantidade, horário). Investigar sede noturna ou ingestão excessiva de líquidos antes de dormir (correlação com noctúria).",
        "scientific_references": "St-Onge MP et al. Effects of diet on sleep quality. Adv Nutr. 2016;7(5):938-949. Crispim CA et al. Relationship between food intake and sleep pattern in healthy individuals. J Clin Sleep Med. 2011;7(6):659-664. Afaghi A et al. High-glycemic-index carbohydrate meals shorten sleep onset. Am J Clin Nutr. 2007;85(2):426-430. Peuhkuri K et al. Diet promotes sleep duration and quality. Nutr Res. 2012;32(5):309-319.",
        "low_description": "Dieta noturna adequada: jantar leve 3-4h antes de dormir, sem alimentos pesados/picantes/gordurosos, ausência de cafeína 6h antes, líquidos limitados 2h antes, sem lanches noturnos.",
        "medium_description": "Dieta noturna moderada: jantar 2-3h antes, ocasionalmente pesado, lanche leve à noite, consumo moderado de líquidos, refluxo ocasional.",
        "high_description": "Dieta noturna inadequada: jantar pesado < 2h antes de dormir, alimentos gordurosos/picantes/açucarados, lanches noturnos frequentes, cafeína à noite, excesso de líquidos, refluxo/desconforto digestivo recorrente."
    },
    "5a4f8b02-04b7-4815-a7d3-62844e1d3dae": {  # Equipamento do sono
        "name": "Equipamento do sono",
        "clinical_relevance": "Equipamentos auxiliares podem melhorar significativamente a qualidade do sono: CPAP/BiPAP para apneia, umidificadores para ressecamento de vias aéreas, máscaras de olhos para bloqueio de luz, protetores auriculares para ruído, dispositivos de ruído branco, travesseiros especiais (cervicais, anti-ronco), rastreadores de sono (wearables). Essencial avaliar adesão ao CPAP em pacientes com apneia obstrutiva do sono (AOS) - uso mínimo de 4h/noite em 70% das noites. Investigar barreiras à adesão: desconforto da máscara, ressecamento nasal, claustrofobia, parceiro.",
        "interpretation_guide": "Identificar uso de CPAP/BiPAP: pressão prescrita, adesão (horas/noite, noites/semana), efeitos colaterais (vazamentos, claustrofobia, ressecamento), limpeza adequada. Avaliar uso de umidificadores (tipo, frequência de limpeza). Questionar sobre máscaras de olhos (blackout do ambiente), protetores auriculares/fones com cancelamento de ruído, aplicativos/dispositivos de rastreamento do sono (interpretação correta dos dados). Investigar uso de travesseiros especiais, camas ajustáveis, dispositivos de posicionamento (evitar decúbito dorsal).",
        "scientific_references": "Weaver TE, Grunstein RR. Adherence to continuous positive airway pressure therapy. Proc Am Thorac Soc. 2008;5(2):173-178. Aloia MS et al. Treatment adherence and outcomes in flexible vs standard continuous positive airway pressure therapy. Chest. 2005;127(6):2085-2093. Ebben MR, Yan P. The effect of eye masks and earplugs on nocturnal sleep in a simulated intensive care environment. Crit Care Med. 2021;49(11):1846-1854.",
        "low_description": "Equipamentos adequados: uso correto e aderente de CPAP (se indicado), umidificador funcional, máscara de olhos/protetores se necessário, travesseiro adequado, dispositivos bem mantidos.",
        "medium_description": "Equipamentos parcialmente utilizados: adesão irregular ao CPAP (3-5h/noite), uso ocasional de máscaras/protetores, equipamentos sem manutenção regular, benefícios parciais.",
        "high_description": "Equipamentos inadequados ou não utilizados: não adesão ao CPAP (< 2h/noite ou abandono), ausência de equipamentos necessários, má manutenção (CPAP sem limpeza), agravamento de sintomas."
    },
    "428052bb-0cbd-4ab2-977c-a6a39cedc3fa": {  # Estimulantes noturnos (cafeína)
        "name": "Estimulantes noturnos (cafeína)",
        "clinical_relevance": "Cafeína é o estimulante psicoativo mais consumido mundialmente, com meia-vida de 5-6 horas. Bloqueia receptores de adenosina, neurotransmissor promotor do sono, causando alerta e reduzindo sono profundo e REM. Consumo após 14h pode atrasar início do sono, reduzir duração total e piorar qualidade. Tolerância individual varia (metabolizadores rápidos vs lentos do CYP1A2). Fontes: café, chá preto/mate/verde, refrigerantes (cola), energéticos, chocolate, medicamentos (analgésicos, descongestionantes). Limite: < 400mg/dia (4 xícaras café), evitar 6h antes de dormir.",
        "interpretation_guide": "Quantificar consumo diário de cafeína: número de xícaras de café/chá, refrigerantes, energéticos, chocolate. Identificar horário da última dose em relação ao sono (crítico se < 6h). Avaliar sensibilidade individual (alguns toleram café à noite, outros são sensíveis mesmo pela manhã). Investigar consumo de outros estimulantes: nicotina (cigarros, vapes), medicamentos (pseudoefedrina, modafinil, metilfenidato), suplementos pré-treino. Questionar sobre sintomas de abstinência (cefaleia ao reduzir) e tentativas prévias de redução.",
        "scientific_references": "Drake C et al. Caffeine effects on sleep taken 0, 3, or 6 hours before going to bed. J Clin Sleep Med. 2013;9(11):1195-1200. Clark I, Landolt HP. Coffee, caffeine, and sleep: A systematic review of epidemiological studies and randomized controlled trials. Sleep Med Rev. 2017;31:70-78. Roehrs T, Roth T. Caffeine: sleep and daytime sleepiness. Sleep Med Rev. 2008;12(2):153-162.",
        "low_description": "Consumo adequado: < 200mg cafeína/dia, última dose > 6h antes de dormir, ausência de estimulantes noturnos, sem impacto no sono.",
        "medium_description": "Consumo moderado: 200-400mg cafeína/dia, última dose 4-6h antes de dormir, consumo ocasional de chá/refrigerante à noite, leve impacto na latência do sono.",
        "high_description": "Consumo excessivo: > 400mg cafeína/dia, consumo < 4h antes de dormir, café/energéticos noturnos, múltiplos estimulantes (cafeína + nicotina), insônia de início relacionada, dependência."
    },
    "88f10871-8821-42db-bee4-01c31f2187f8": {  # Higiene do sono
        "name": "Higiene do sono",
        "clinical_relevance": "Higiene do sono compreende práticas comportamentais e ambientais que promovem sono de qualidade. Componentes principais: horários regulares (deitar/acordar), ambiente escuro/silencioso/fresco (16-19°C), rotina relaxante pré-sono (30-60min), evitar telas 1-2h antes, cama exclusiva para sono/sexo (não trabalho/TV), exposição à luz solar matinal, exercícios regulares (não < 3h antes de dormir), evitar cochilos longos (> 30min) ou tardios (após 15h). Má higiene do sono é fator perpetuador de insônia crônica.",
        "interpretation_guide": "Avaliar cada componente: regularidade de horários (variação fim de semana vs dias úteis), rotina pré-sono (relaxante vs estimulante - TV, trabalho, discussões), ambiente do quarto (temperatura, luz, ruído), uso de cama para atividades não relacionadas ao sono, exposição à luz (manhã/noite), padrão de exercícios (horário, intensidade), cochilos diurnos (duração, horário, impacto noturno). Investigar comportamentos contraproducentes: ficar na cama acordado > 20min (reforço negativo), check de celular durante despertares noturnos, compensação de sono perdido (dormir até tarde, cochilos longos).",
        "scientific_references": "Irish LA et al. The role of sleep hygiene in promoting public health: A review of empirical evidence. Sleep Med Rev. 2015;22:23-36. Mastin DF et al. Assessment of sleep hygiene using the Sleep Hygiene Index. J Behav Med. 2006;29(3):223-227. Stepanski EJ, Wyatt JK. Use of sleep hygiene in the treatment of insomnia. Sleep Med Rev. 2003;7(3):215-225.",
        "low_description": "Higiene do sono excelente: horários regulares (< 30min variação), rotina relaxante pré-sono, ambiente ideal, cama exclusiva para sono, exposição solar matinal, exercícios regulares em horário adequado, sem cochilos tardios.",
        "medium_description": "Higiene do sono moderada: horários irregulares ocasionais, rotina pré-sono inconsistente, ambiente aceitável mas não otimizado, uso ocasional de telas antes de dormir, cochilos esporádicos.",
        "high_description": "Higiene do sono inadequada: horários caóticos (> 2h variação), ausência de rotina, ambiente inapropriado (claro/ruidoso/quente), uso de telas na cama, cama para múltiplas atividades, cochilos longos/tardios, exercícios noturnos."
    },
    "98a7ddf2-4c79-44c9-8724-bd5baf7c709f": {  # Histórico medicamentos/suplementos
        "name": "Histórico medicamentos/suplementos",
        "clinical_relevance": "Diversos medicamentos afetam o sono: Sedativos/hipnóticos (benzodiazepínicos, Z-drugs) - dependência, tolerância, rebote; Antidepressivos (SSRIs) - insônia inicial ou sonolência; Betabloqueadores - pesadelos, insônia; Corticoides - insônia, agitação; Broncodilatadores - estimulação; Diuréticos - noctúria. Suplementos: melatonina (1-3mg, eficaz para jet lag, turnos), magnésio (relaxamento muscular), L-triptofano/5-HTP (precursores serotonina), valeriana (sedação leve), CBD (evidências limitadas). Importante: histórico de uso de hipnóticos (risco dependência), tentativas de descontinuação, efeitos colaterais.",
        "interpretation_guide": "Listar TODOS os medicamentos atuais e recentes (nome, dose, horário, duração): hipnóticos (benzodiazepínicos, zolpidem, eszopiclona), antidepressivos, ansiolíticos, anti-hipertensivos, broncodilatadores, anti-histamínicos, corticoides, analgésicos opioides. Avaliar indicação, eficácia percebida, efeitos colaterais (sonolência diurna, tolerância, rebote). Investigar uso de suplementos: melatonina (dose, horário), magnésio, fitoterápicos (valeriana, camomila, passiflora), CBD, triptofano. Questionar automedicação, uso irregular ou abuso de sedativos.",
        "scientific_references": "Wichniak A et al. Effects of antidepressants on sleep. Curr Psychiatry Rep. 2017;19(9):63. Wilt TJ et al. Pharmacologic treatment of insomnia disorder: an evidence report for a clinical practice guideline by the American College of Physicians. Ann Intern Med. 2016;165(2):103-112. Auld F et al. Evidence for the efficacy of melatonin in the treatment of primary adult sleep disorders. Sleep Med Rev. 2017;34:10-22.",
        "low_description": "Medicamentos/suplementos sem impacto negativo: ausência de hipnóticos ou uso racional de melatonina/magnésio, medicações crônicas sem efeito no sono, sem automedicação.",
        "medium_description": "Medicamentos com impacto moderado: uso ocasional de hipnóticos (< 3x/semana), antidepressivos com leve sedação/insônia, tentativas de suplementação natural, efeitos colaterais toleráveis.",
        "high_description": "Medicamentos com impacto significativo: uso crônico de benzodiazepínicos/Z-drugs (> 4 semanas), dependência, tolerância, rebote na descontinuação, polifarmácia com efeitos no sono, automedicação, noctúria por diuréticos noturnos."
    },
    "f0fed096-a535-4388-bb84-343c3b42b84c": {  # Histórico familiar
        "name": "Histórico familiar de distúrbios do sono",
        "clinical_relevance": "Diversos distúrbios do sono têm componente hereditário significativo: apneia obstrutiva do sono (AOS) - hereditariedade 40%, associada a anatomia craniofacial e obesidade familiar; insônia crônica - risco 3-5x maior se pais afetados; síndrome das pernas inquietas (SPI) - 50-70% casos com história familiar; narcolepsia - risco 10-40x se familiar de 1º grau; distúrbios do ritmo circadiano (síndrome fase avançada/atrasada) - padrão familiar comum. Investigar também: roncos, sonambulismo, terror noturno (parassônias), bruxismo. História familiar auxilia no diagnóstico diferencial e alerta para rastreamento precoce.",
        "interpretation_guide": "Questionar sobre distúrbios do sono em familiares de 1º grau (pais, irmãos, filhos): apneia do sono/roncos, insônia crônica, síndrome das pernas inquietas, narcolepsia, sonambulismo, terror noturno, bruxismo. Identificar padrões familiares: todos 'coruja' ou 'cotovia' (cronotipos), tendência a sono curto/longo, dificuldade crônica para dormir. Avaliar se familiar foi diagnosticado formalmente (polissonografia) ou apenas relato. Investigar história familiar de obesidade, diabetes, hipertensão (fatores de risco para apneia).",
        "scientific_references": "Redline S et al. The genetics of sleep apnea. Sleep Med Rev. 2000;4(6):583-602. Dauvilliers Y et al. Family studies in insomnia. J Psychosom Res. 2005;58(3):271-278. Winkelmann J et al. Genome-wide association study of restless legs syndrome identifies common variants in three genomic regions. Nat Genet. 2007;39(8):1000-1006.",
        "low_description": "História familiar negativa: ausência de distúrbios do sono diagnosticados em familiares de 1º grau, sono normal em pais/irmãos.",
        "medium_description": "História familiar moderada: roncos/insônia ocasional em familiares, sem diagnóstico formal, sintomas leves, sem impacto funcional significativo.",
        "high_description": "História familiar positiva: apneia do sono diagnosticada (CPAP), insônia crônica, síndrome das pernas inquietas, narcolepsia, parassônias recorrentes em múltiplos familiares de 1º grau, padrão hereditário claro."
    },
    "2ce2bbb8-31d7-44e1-88b3-d7282aeb79a7": {  # Hábitos cônjuge
        "name": "Hábitos de sono do cônjuge/parceiro",
        "clinical_relevance": "O sono do parceiro pode impactar significativamente a qualidade do sono: roncos (principal queixa - até 95 dB, equivalente a cortador de grama), apneia não tratada, movimentos noturnos excessivos, síndrome das pernas inquietas, horários incompatíveis (cronotipo diferente, trabalho em turnos), uso de telas/luzes no quarto, diferenças na temperatura preferencial. Estudos mostram que parceiros de roncadores perdem até 1-2h de sono/noite. Considerar quartos separados quando impacto é severo (não implica problema conjugal, é estratégia de saúde). Importante: se parceiro tem apneia não tratada, encaminhar para avaliação.",
        "interpretation_guide": "Avaliar hábitos do parceiro que afetam o sono: roncos (intensidade, frequência, pausas respiratórias observadas), movimentos noturnos (chutes, inquietação, SPI), horários de sono (deitar/acordar muito diferentes), uso de dispositivos eletrônicos na cama, preferências incompatíveis (temperatura, luz, ruído branco). Questionar se paciente acorda devido ao parceiro, se já consideraram quartos separados, impacto na relação. Investigar se parceiro tem diagnóstico de distúrbio do sono ou resistência a buscar avaliação médica.",
        "scientific_references": "Parish JM, Lyng PJ. Quality of life in bed partners of patients with obstructive sleep apnea or hypopnea after treatment with continuous positive airway pressure. Chest. 2003;124(3):942-947. Strawbridge WJ et al. Marital quality and health: A longitudinal study. J Psychosom Res. 1997;42(4):315-323. Beninati W et al. The effect of snoring and obstructive sleep apnea on the sleep quality of bed partners. Mayo Clin Proc. 1999;74(10):955-958.",
        "low_description": "Hábitos do parceiro compatíveis: sem roncos, movimentos mínimos, horários similares, respeito ao ambiente de sono, ausência de impacto negativo.",
        "medium_description": "Hábitos do parceiro com impacto moderado: roncos ocasionais/leves, movimentos noturnos esporádicos, horários ligeiramente incompatíveis, adaptação parcial, impacto tolerável.",
        "high_description": "Hábitos do parceiro com impacto severo: roncos intensos/apneia não tratada, movimentos noturnos frequentes/SPI, horários muito incompatíveis, uso de telas/luzes perturbadoras, despertar frequente do paciente, consideração de quartos separados."
    },
    "b991b0a9-7b65-4fd8-87b3-ed45772c794f": {  # Idade desfralde
        "name": "Idade do desfralde noturno",
        "clinical_relevance": "Controle esfincteriano noturno (desfralde noturno) ocorre tipicamente entre 2-5 anos, após desfralde diurno. Enurese noturna (urinar na cama) após 5 anos é considerada patológica se ocorre ≥2x/semana por ≥3 meses. Causas: imaturidade vesical, produção excessiva de urina noturna (deficiência de vasopressina), sono profundo (dificuldade de acordar com bexiga cheia), fatores genéticos (70% se ambos pais com histórico), constipação, apneia do sono. Desfralde tardio ou enurese persistente pode indicar distúrbio do sono (dificuldade de acordar), apneia (pressão intra-abdominal), problemas urológicos ou constipação crônica. Relevante em adultos: história de enurese até adolescência sugere sono muito profundo ou apneia de longa data.",
        "interpretation_guide": "Questionar idade do desfralde diurno e noturno (se houve diferença significativa). Em crianças: frequência atual de enurese, volume de urina, horário (primeira metade da noite = sono profundo), padrão familiar (pais/irmãos), associação com roncos/apneia, constipação. Em adultos: histórico de enurese na infância/adolescência (até que idade), se tratou/resolveu espontaneamente, relação com sono profundo. Investigar noctúria atual (acordar para urinar) vs enurese (urinar dormindo). Avaliar se há sintomas de apneia do sono concomitantes.",
        "scientific_references": "Neveus T et al. Evaluation and treatment of monosymptomatic enuresis: a standardization document from the International Children's Continence Society. J Urol. 2010;183(2):441-447. Van Hoeck K et al. Economic burden of nocturnal enuresis. BJU Int. 2013;111(6):889-897. Soster LA et al. Association between sleep apnea and enuresis in children. J Pediatr. 2012;88(4):310-316.",
        "low_description": "Desfralde normal: controle noturno estabelecido entre 2-4 anos, sem enurese após 5 anos, padrão adequado de despertar para urinar.",
        "medium_description": "Desfralde com atraso leve: controle noturno entre 4-6 anos, enurese ocasional até 7-8 anos, resolução espontânea, sem impacto atual.",
        "high_description": "Desfralde tardio ou enurese persistente: controle noturno após 6 anos, enurese frequente até adolescência, noctúria atual significativa (≥2x/noite), possível associação com apneia/sono profundo, impacto psicossocial."
    },
    "d60fb723-4152-46ec-8056-4135326c1f8b": {  # Iluminação
        "name": "Iluminação do ambiente de sono",
        "clinical_relevance": "Luz é o sincronizador mais potente do ritmo circadiano. Luz azul (450-480nm) suprime melatonina via células ganglionares retinianas intrinsecamente fotossensíveis (ipRGCs). Exposição à luz à noite atrasa fase circadiana, reduz duração e qualidade do sono. Ideal: < 1 lux durante sono (escuridão total), < 50 lux 2h antes de dormir, bloqueio de luz azul (filtros, óculos âmbar). Fontes problemáticas: telas (smartphones, tablets, computadores, TV), LEDs brancos/azuis, luz da rua através de janelas, relógios digitais brilhantes. Luz matinal (> 1000 lux) é benéfica: consolida ritmo, promove alerta diurno. Trabalhadores noturnos: exposição à luz intensa durante trabalho + escuridão absoluta durante sono diurno.",
        "interpretation_guide": "Avaliar iluminação noturna: escuridão do quarto (luz externa, cortinas blackout), dispositivos luminosos (relógios digitais, LEDs de aparelhos, luz de emergência), uso de máscaras de olhos. Investigar exposição à luz 2h antes de dormir: telas sem filtro de luz azul, iluminação intensa no ambiente, luzes azuis/brancas vs âmbar/vermelhas. Questionar exposição matinal à luz natural (acordar com luz do sol, sair ao ar livre nas primeiras horas). Em trabalhadores noturnos: estratégias de iluminação (luz intensa no trabalho, óculos escuros no retorno, escuridão absoluta no quarto).",
        "scientific_references": "Chang AM et al. Evening use of light-emitting eReaders negatively affects sleep, circadian timing, and next-morning alertness. Proc Natl Acad Sci USA. 2015;112(4):1232-1237. Gooley JJ et al. Exposure to room light before bedtime suppresses melatonin onset and shortens melatonin duration in humans. J Clin Endocrinol Metab. 2011;96(3):E463-472. Lockley SW et al. Short-wavelength sensitivity for the direct effects of light on alertness. Sleep. 2006;29(2):161-168.",
        "low_description": "Iluminação adequada: escuridão completa durante sono (< 1 lux, cortinas blackout), iluminação âmbar/reduzida 2h antes, bloqueio de luz azul (filtros, óculos), exposição solar matinal.",
        "medium_description": "Iluminação moderada: quarto razoavelmente escuro mas com alguma luz externa, uso ocasional de telas sem filtro antes de dormir, relógios digitais brilhantes, exposição matinal irregular.",
        "high_description": "Iluminação inadequada: quarto com iluminação significativa (rua, janelas sem cortina, dispositivos luminosos), uso intenso de telas antes de dormir sem filtros, ausência de exposição solar matinal, atraso de fase circadiana."
    },
    "9ca5a90d-ed46-49f7-9e53-013004f6e3db": {  # Infância
        "name": "Padrão de sono na infância",
        "clinical_relevance": "Problemas de sono na infância frequentemente persistem na vida adulta. Crianças com insônia comportamental (dificuldade de iniciar sono independentemente, associação com objetos/rotinas específicas, despertares noturnos exigindo intervenção parental) têm risco aumentado de insônia crônica na adultez. Parassônias (sonambulismo, terror noturno, confusão ao acordar) são comuns em crianças (pico 3-8 anos), geralmente resolvem na adolescência; persistência sugere sono fragmentado ou apneia. Histórico de roncos/apneia na infância (hipertrofia adenoamigdaliana) pode evoluir para apneia adulta se fatores de risco presentes (obesidade). Investigar: necessidade de sono excessivo ou insuficiente, dificuldades escolares relacionadas a sonolência.",
        "interpretation_guide": "Questionar padrão de sono na infância (0-12 anos): dificuldade para dormir sozinho, necessidade de presença parental, despertares noturnos frequentes, parassônias (sonambulismo, terror noturno, pesadelos), roncos persistentes, respiração oral, sono agitado. Avaliar se houve tratamento (adenoamigdalectomia, terapia comportamental). Investigar se problemas persistiram na adolescência/adultez ou resolveram. Questionar impacto escolar: sonolência diurna, dificuldade de concentração, notas baixas relacionadas. Identificar eventos traumáticos na infância que afetaram o sono (perdas, mudanças, hospitalizações).",
        "scientific_references": "Simola P et al. Persistence of sleep problems from childhood to adolescence: a longitudinal study of 1,729 children. Pediatrics. 2014;133(4):e969-e977. Gregory AM, O'Connor TG. Sleep problems in childhood: a longitudinal study of developmental change and association with behavioral problems. J Am Acad Child Adolesc Psychiatry. 2002;41(8):964-971. Mindell JA, Owens JA. A Clinical Guide to Pediatric Sleep: Diagnosis and Management of Sleep Problems. 3rd ed. Wolters Kluwer; 2015.",
        "low_description": "Sono na infância normal: sono independente desde cedo, sem despertares noturnos significativos, ausência de parassônias persistentes, sono restaurador.",
        "medium_description": "Sono na infância com dificuldades transitórias: insônia comportamental leve, parassônias ocasionais (resolvidas na adolescência), roncos intermitentes, impacto funcional mínimo.",
        "high_description": "Sono na infância problemático: insônia crônica desde criança, despertares noturnos frequentes, parassônias persistentes até adolescência, roncos/apneia (adenoamigdalectomia), impacto escolar (sonolência, baixo rendimento), persistência de problemas na adultez."
    },
    "67185b74-1b99-4397-90e1-9971e7d24353": {  # Janelas/claridade
        "name": "Janelas e entrada de claridade",
        "clinical_relevance": "Controle da entrada de luz natural no quarto é fundamental para sincronização circadiana e qualidade do sono. Ideal: exposição à luz matinal (acordar com luz do sol ou abrir janelas logo ao acordar - sinal de início do dia, suprime melatonina, promove alerta), escuridão completa durante o sono noturno (cortinas blackout bloqueando 100% luz externa). Problemas comuns: luz da rua (postes, carros), iluminação de vizinhos, amanhecer precoce no verão (acordar prematuramente por luz), ausência de cortinas adequadas. Trabalhadores noturnos requerem blackout absoluto para sono diurno. Janelas também impactam temperatura e ruído (ventilação vs isolamento).",
        "interpretation_guide": "Avaliar tipo de janela: tamanho, orientação (leste = luz matinal, oeste = luz entardecer), presença de cortinas/persianas (blackout vs translúcidas). Questionar entrada de luz externa durante sono: luz da rua, postes, vizinhos, amanhecer (horário que começa a clarear no quarto). Investigar uso de máscaras de olhos (compensação para falta de blackout). Avaliar exposição à luz matinal: acordar com luz natural (benéfico) vs quarto sempre escuro pela manhã. Em trabalhadores noturnos: efetividade do blackout para sono diurno. Considerar impacto de janelas no ruído (rua movimentada) e temperatura (sol direto aquecendo quarto).",
        "scientific_references": "Burgess HJ, Molina TA. Home lighting before usual bedtime impacts circadian timing: A field study. Photochem Photobiol. 2014;90(3):723-726. Wright KP Jr et al. Entrainment of the human circadian clock to the natural light-dark cycle. Curr Biol. 2013;23(16):1554-1558. Figueiro MG et al. The impact of daytime light exposures on sleep and mood in office workers. Sleep Health. 2017;3(3):204-215.",
        "low_description": "Janelas/claridade adequadas: cortinas blackout efetivas (escuridão completa noturna), exposição matinal à luz natural, controle total da entrada de luz conforme necessidade.",
        "medium_description": "Janelas/claridade moderadas: cortinas parcialmente efetivas (alguma luz externa entra), exposição matinal irregular, uso ocasional de máscaras de olhos, leve impacto no sono.",
        "high_description": "Janelas/claridade inadequadas: ausência de blackout (luz intensa da rua/postes), acordar prematuramente com amanhecer, quarto sempre escuro (ausência de luz matinal), trabalhador noturno sem blackout efetivo, impacto significativo no ritmo circadiano."
    },
    "2454de7e-70e5-45e4-83d7-0091b5847b49": {  # Lençóis/cobertas
        "name": "Lençóis, cobertas e roupas de cama",
        "clinical_relevance": "Termorregulação adequada durante o sono é essencial. Temperatura corporal central cai ~1-1.5°C durante sono, facilitada por perda de calor periférico. Roupas de cama inadequadas interferem nesse processo: cobertores muito pesados retêm calor excessivo, sintéticos impedem dissipação de umidade, alergênicos (ácaros, mofo) causam sintomas respiratórios/alérgicos que fragmentam sono. Materiais recomendados: algodão (respirável), linho (fresco), bambu (antibacteriano), lã (termorregulação natural). Evitar: poliéster (retém calor e umidade), cobertores excessivamente pesados. Importante: higiene adequada (lavagem semanal, eliminação de ácaros), temperatura ambiente 16-19°C.",
        "interpretation_guide": "Avaliar tipo de lençóis/cobertas: material (algodão, sintético, lã), peso (leves vs pesados), número de camadas, adequação à estação do ano. Questionar desconforto térmico: acordar com calor/suor excessivo ou frio, necessidade de descobrir/cobrir durante noite, suor noturno (distinguir de fogachos menopausais ou hiperhidrose). Investigar sintomas alérgicos: coriza, espirros, coceira, congestão nasal matinal, tosse noturna (possível alergia a ácaros). Avaliar frequência de lavagem (semanal recomendado), presença de mofo/umidade, idade dos travesseiros (substituir a cada 1-2 anos). Questionar uso de cobertores pesados terapêuticos (weighted blankets - 5-10% do peso corporal, benefício em ansiedade).",
        "scientific_references": "Okamoto-Mizuno K, Mizuno K. Effects of thermal environment on sleep and circadian rhythm. J Physiol Anthropol. 2012;31(1):14. Ackerley R et al. Positive effects of a weighted blanket on insomnia. J Sleep Med Disord. 2015;2(3):1022. Porcheret K et al. Investigation of the impact of thermal hugs during sleep on insomnia. Sleep. 2019;42(Suppl 1):A95.",
        "low_description": "Roupas de cama adequadas: materiais respiráveis (algodão, linho), peso apropriado para estação, conforto térmico ideal, lavagem semanal, ausência de alergias.",
        "medium_description": "Roupas de cama aceitáveis: materiais mistos, peso não ideal (ocasionalmente quente/frio), lavagem quinzenal, sintomas alérgicos leves, ajustes noturnos (cobrir/descobrir).",
        "high_description": "Roupas de cama inadequadas: materiais sintéticos/não respiráveis, peso inapropriado (muito pesado ou insuficiente), desconforto térmico frequente (suor/frio), lavagem irregular, sintomas alérgicos (ácaros), qualidade/higiene precárias."
    },
    "25aa750f-d458-4e39-b37c-b5a66502340a": {  # Melhores épocas
        "name": "Melhores épocas de sono",
        "clinical_relevance": "Variações sazonais, etapas da vida e contextos específicos afetam a qualidade do sono. Fatores sazonais: inverno (noites mais longas, maior produção de melatonina, risco de SAD - seasonal affective disorder), verão (amanhecer precoce, calor, sono mais curto). Etapas da vida: adolescência (atraso de fase natural, necessidade de 8-10h), gestação (1º trimestre sonolência, 3º trimestre desconforto/noctúria), pós-parto (fragmentação extrema), menopausa (fogachos, insônia), envelhecimento (sono mais leve, despertares). Contextos: férias (sem alarmes, recuperação de débito de sono), períodos de estresse reduzido, mudanças de altitude/fuso (jet lag). Identificar 'melhores épocas' auxilia a reconhecer fatores modificáveis vs imutáveis.",
        "interpretation_guide": "Questionar variações sazonais: sono melhor no inverno (noites longas, frio) ou verão, impacto da luminosidade (SAD no inverno, acordar com sol no verão). Investigar etapas da vida: sono na adolescência vs adultez vs atual, impacto de gestações, menopausa (fogachos noturnos), uso de TRH. Avaliar contextos específicos: sono em férias vs trabalho, fins de semana vs dias úteis, períodos sem estresse vs estresse crônico. Identificar momentos de sono 'ideal' na vida do paciente (quando dormia melhor, em que circunstâncias), comparar com período atual. Questionar viagens: jet lag (facilidade de adaptação), sono em altitudes diferentes.",
        "scientific_references": "Rosenthal NE et al. Seasonal affective disorder: a description of the syndrome and preliminary findings with light therapy. Arch Gen Psychiatry. 1984;41(1):72-80. Terman M. Evolving applications of light therapy. Sleep Med Rev. 2007;11(6):497-507. Carrier J et al. Sleep and circadian rhythm changes during pregnancy. Sleep Med Clin. 2018;13(3):427-436. Carskadon MA et al. A self-administered rating scale for pubertal development. J Adolesc Health. 1993;14(3):190-195.",
        "low_description": "Sono consistentemente bom: qualidade estável ao longo do ano, sem variações sazonais significativas, adaptação adequada a mudanças, sono satisfatório na maioria dos contextos.",
        "medium_description": "Sono com variações moderadas: melhor em determinada estação ou contexto (férias, fins de semana), impacto leve de mudanças (viagens, estações), qualidade aceitável na maior parte do tempo.",
        "high_description": "Sono com variações extremas: grande discrepância entre melhores/piores épocas, SAD (depressão sazonal), sono apenas adequado em contextos ideais (férias, ausência de estresse), dificuldade significativa em períodos específicos da vida (menopausa, pós-parto), jet lag severo."
    },
    "b256e241-ac54-4043-82a3-0ec5087702bc": {  # Motivos
        "name": "Motivos para distúrbios do sono",
        "clinical_relevance": "Identificar causas subjacentes de distúrbios do sono é fundamental para tratamento eficaz. Categorias principais: 1) Médicas - apneia, refluxo, dor crônica, noctúria, distúrbios neurológicos (Parkinson, demência), doenças respiratórias (asma, DPOC), cardíacas (ICC), endócrinas (hiper/hipotireoidismo, diabetes); 2) Psiquiátricas - ansiedade, depressão, TEPT, transtorno bipolar; 3) Comportamentais - má higiene do sono, turnos, jet lag crônico, uso de substâncias (álcool, drogas); 4) Ambientais - ruído, luz, temperatura, parceiro; 5) Medicamentosas - efeito colateral, abstinência. Abordagem: tratar causa raiz (não apenas sintoma insônia). Insônia primária (sem causa identificável) requer TCC-I (terapia cognitivo-comportamental para insônia).",
        "interpretation_guide": "Investigar sistematicamente causas médicas: sintomas respiratórios (roncos, apneia), digestivos (refluxo), urológicos (noctúria), dor (localização, intensidade noturna), condições crônicas (listar diagnósticos). Avaliar causas psiquiátricas: transtornos diagnosticados (em tratamento vs não tratados), sintomas de ansiedade (ruminação, preocupações), depressão (despertar precoce, anedonia), TEPT (pesadelos, hipervigilância). Identificar causas comportamentais: trabalho em turnos, viagens frequentes, uso de álcool/drogas, padrão irregular de sono. Explorar causas ambientais: problemas no quarto (ruído, luz, temperatura, cama, parceiro). Revisar medicações (ver item específico). Atribuição do paciente: o que ele acredita ser a causa?",
        "scientific_references": "Roth T. Insomnia: definition, prevalence, etiology, and consequences. J Clin Sleep Med. 2007;3(5 Suppl):S7-10. Sateia MJ. International classification of sleep disorders-third edition: highlights and modifications. Chest. 2014;146(5):1387-1394. Riemann D et al. European guideline for the diagnosis and treatment of insomnia. J Sleep Res. 2017;26(6):675-700. Morin CM et al. Cognitive behavioral therapy for chronic insomnia: a systematic review. Ann Intern Med. 2015;163(3):191-204.",
        "low_description": "Motivos identificáveis e tratáveis: causa clara (ex: apneia diagnosticada com CPAP, refluxo controlado, depressão em remissão), resposta adequada ao tratamento, ausência de múltiplos fatores perpetuadores.",
        "medium_description": "Motivos parcialmente identificados: causa provável mas não confirmada (ex: possível apneia sem polissonografia), múltiplos fatores leves (má higiene + estresse + ambiente), tratamento em andamento com resposta parcial.",
        "high_description": "Motivos múltiplos ou não identificados: causas complexas entrelaçadas (médica + psiquiátrica + comportamental), insônia primária (sem causa clara), múltiplos fatores perpetuadores não tratados, cronicidade (> 3 meses), impacto funcional significativo, tratamentos prévios sem sucesso."
    },
}

def login() -> str:
    """Faz login e retorna o token de autenticação."""
    try:
        response = requests.post(
            f"{API_BASE}/auth/login",
            json={"email": LOGIN_EMAIL, "password": LOGIN_PASSWORD}
        )
        response.raise_for_status()
        token = response.json()["accessToken"]
        print(f"✓ Login realizado com sucesso")
        return token
    except Exception as e:
        print(f"✗ Erro no login: {e}")
        sys.exit(1)

def get_score_item(token: str, item_id: str) -> Dict[str, Any]:
    """Busca um score item pelo ID."""
    try:
        headers = {"Authorization": f"Bearer {token}"}
        response = requests.get(
            f"{API_BASE}/score-items/{item_id}",
            headers=headers
        )
        response.raise_for_status()
        data = response.json()
        # A resposta pode vir com ou sem wrapper "data"
        if "data" in data:
            return data["data"]
        return data
    except Exception as e:
        print(f"✗ Erro ao buscar item {item_id}: {e}")
        return None

def update_score_item(token: str, item_id: str, data: Dict[str, Any]) -> bool:
    """Atualiza um score item."""
    try:
        headers = {"Authorization": f"Bearer {token}"}
        response = requests.put(
            f"{API_BASE}/score-items/{item_id}",
            headers=headers,
            json=data
        )
        response.raise_for_status()
        return True
    except Exception as e:
        print(f"✗ Erro ao atualizar item {item_id}: {e}")
        return False

def enrich_item(token: str, item_id: str, content: Dict[str, Any]) -> bool:
    """Enriquece um score item com conteúdo clínico."""
    # Buscar item atual
    item = get_score_item(token, item_id)
    if not item:
        return False

    current_name = item.get("name", "")
    print(f"\n📝 Processando: {current_name}")

    # Preparar dados de atualização
    update_data = {
        "name": content.get("name", current_name),
        "clinical_relevance": content["clinical_relevance"],
        "interpretation_guide": content["interpretation_guide"],
        "scientific_references": content["scientific_references"],
        "low_description": content["low_description"],
        "medium_description": content["medium_description"],
        "high_description": content["high_description"]
    }

    # Atualizar
    if update_score_item(token, item_id, update_data):
        print(f"✓ Item atualizado: {content['name']}")
        return True
    else:
        print(f"✗ Falha ao atualizar: {content['name']}")
        return False

def main():
    print("=" * 80)
    print("ENRIQUECIMENTO DE ITEMS - SONO PARTE 2")
    print("20 items: Campos eletromagnéticos até Motivos")
    print("=" * 80)

    # Login
    token = login()

    # Processar cada item
    success_count = 0
    fail_count = 0

    for item_id in ITEM_IDS:
        if item_id in CLINICAL_CONTENT:
            if enrich_item(token, item_id, CLINICAL_CONTENT[item_id]):
                success_count += 1
            else:
                fail_count += 1
        else:
            print(f"⚠ Sem conteúdo definido para ID: {item_id}")
            fail_count += 1

    # Resumo
    print("\n" + "=" * 80)
    print("RESUMO DA EXECUÇÃO")
    print("=" * 80)
    print(f"✓ Sucesso: {success_count}")
    print(f"✗ Falhas: {fail_count}")
    print(f"📊 Total processado: {success_count + fail_count}")
    print("=" * 80)

if __name__ == "__main__":
    main()
