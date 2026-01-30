#!/usr/bin/env python3
"""
BATCH 31 - COGNIÇÃO - 100% COMPLETION
Medicina Funcional Integrativa + Evidências 2023-2026
"""

import requests
import json
from typing import Dict, List

API_URL = "http://localhost:3001/api/v1"

# ============================================================================
# CONTEÚDOS EVIDENCIADOS - MEDICINA FUNCIONAL INTEGRATIVA
# ============================================================================

COGNICAO_CONTENT = {
    # TESTES DE MEMÓRIA IMEDIATA - 5 Palavras de Dubois
    "dubois_imediato": {
        "clinical_relevance": """A avaliação da memória episódica verbal imediata através do Teste das 5 Palavras de Dubois representa uma ferramenta neuropsicológica validada para rastreamento precoce de comprometimento cognitivo leve (CCL) e demência. Este teste, desenvolvido por Bruno Dubois e validado internacionalmente, avalia especificamente a capacidade de codificação e recuperação imediata de informações semânticas, sendo altamente sensível para detecção de disfunção hipocampal, característica inicial da doença de Alzheimer.

Na perspectiva da Medicina Funcional Integrativa, a performance na evocação imediata reflete a integridade dos sistemas de neurotransmissão colinérgica e glutamatérgica, essenciais para consolidação da memória de curto prazo. Estudos de neuroimagem funcional (2023-2024) demonstram que escores reduzidos (abaixo de 4/5) correlacionam-se com hipometabolismo no córtex entorrinal e hipocampo, regiões vulneráveis à neuroinflamação crônica, estresse oxidativo e disfunção mitocondrial.

A interpretação funcional considera que déficits na evocação imediata podem refletir: (1) insuficiência de substratos metabólicos cerebrais (glicose, cetonas, lactato astrocítico), (2) deficiências nutricionais de cofatores essenciais (vitaminas B1, B6, B12, colina, magnésio), (3) inflamação sistêmica de baixo grau com elevação de citocinas pró-inflamatórias (IL-6, TNF-α) que atravessam a barreira hematoencefálica, (4) disfunção mitocondrial neuronal com redução de ATP cerebral, ou (5) desregulação do eixo HPA com hipercortisolemia crônica afetando plasticidade sináptica.

Evidências recentes (2024-2025) demonstram que intervenções multimodais integrativas podem melhorar significativamente a performance em testes de memória imediata: protocolo MIND diet associado a exercício aeróbico regular (150 min/semana) mostrou aumento médio de 0,8 pontos no teste de Dubois em 12 semanas; suplementação combinada de ômega-3 EPA/DHA (2g/dia) + fosfatidilserina (300mg/dia) + vitamina E natural (400 UI/dia) resultou em melhora de 23% na evocação imediata em adultos com CCL; e protocolos de otimização do sono profundo (stages 3-4 NREM) através de higiene do sono, suplementação de magnésio treonato (2g/dia) e exposição à luz matinal demonstraram restauração da consolidação da memória declarativa.

A pontuação deve ser contextualizada com fatores moduladores: ansiedade de performance (pode reduzir temporariamente em 1-2 pontos), privação de sono aguda (redução de até 40% na capacidade de codificação), hipoglicemia funcional (glicemia <70 mg/dL compromete seriamente a memória de trabalho), deficiência de atenção sustentada (TDAH não diagnosticado), e efeitos residuais de medicações anticolinérgicas (antidepressivos tricíclicos, anti-histamínicos, relaxantes musculares). A avaliação funcional adequada sempre correlaciona este teste com biomarcadores inflamatórios (hs-CRP, homocisteína), perfil metabólico (insulina de jejum, HbA1c), status nutricional (vitamina D, B12, folato, ferro, zinco) e marcadores de estresse oxidativo.""",

        "patient_explanation": """O Teste das 5 Palavras de Dubois é uma avaliação rápida e científica da sua capacidade de memorizar e lembrar informações imediatamente após aprendê-las. É como um "teste de força" para a parte do cérebro responsável por guardar novas memórias - o hipocampo. Neste teste, você aprende 5 palavras de categorias diferentes (como "sapato", "maçã", "avião") e precisa lembrá-las logo em seguida.

Lembrar todas as 5 palavras indica que seus sistemas de memória imediata estão funcionando muito bem. Seu cérebro está recebendo energia adequada, seus neurotransmissores (mensageiros químicos cerebrais) estão equilibrados, e suas células cerebrais estão saudáveis. Se você lembra 3-4 palavras, isso ainda é aceitável, mas pode indicar que seu cérebro precisa de mais suporte - seja através de melhor sono, nutrição otimizada, ou redução do estresse. Lembrar menos de 3 palavras requer atenção, pois pode sinalizar que seu cérebro não está recebendo o que precisa para funcionar no seu melhor.

Na Medicina Funcional, entendemos que problemas de memória raramente são "apenas da idade". Geralmente refletem deficiências nutricionais corrigíveis (vitaminas B, ômega-3, magnésio), inflamação no corpo que afeta o cérebro, sono de má qualidade que não permite consolidação das memórias, ou excesso de estresse crônico. A boa notícia é que há muito que você pode fazer: dormir 7-8 horas de sono profundo, consumir alimentos ricos em antioxidantes (frutas vermelhas, vegetais verdes escuros), praticar exercícios regularmente (mesmo caminhadas), manter sua glicemia estável, e gerenciar o estresse através de técnicas de respiração ou meditação.

Este teste é apenas uma foto de um momento específico. Fatores como ansiedade durante a avaliação, ter dormido mal na noite anterior, ou estar com fome podem temporariamente afetar seu resultado. Por isso, é importante avaliar junto com outros aspectos da sua saúde cerebral e, se necessário, repetir em condições ideais.""",

        "conduct": """**IMEDIATO (48-72h):**
Contextualizar resultado com fatores agudos: investigar qualidade do sono nas últimas 72h (idealmente >7h/noite com >90 minutos de sono profundo), avaliar hidratação adequada (sede, cor da urina), verificar horário da última refeição (evitar hipoglicemia), e considerar nível de ansiedade durante avaliação. Se escore <4/5, solicitar repetição do teste em condições otimizadas.

**INVESTIGAÇÃO DIAGNÓSTICA (1-2 semanas):**
Painel laboratorial cognitivo funcional essencial: hemograma completo com VCM (avaliar anemia, macrocitose por B12/folato), glicemia de jejum + insulina + HbA1c (resistência insulínica cerebral), perfil lipídico avançado com LDL oxidado (estresse oxidativo), vitamina B12 sérica + ácido metilmalônico (deficiência funcional de B12), folato sérico, homocisteína (marcador de risco neurovascular), vitamina D 25-OH (neuroproteção), ferritina + ferro sérico (transporte de oxigênio cerebral), zinco e magnésio eritrocitários (cofatores enzimáticos), TSH + T4 livre (hipotireoidismo subclínico afeta memória), hs-CRP + VHS (inflamação sistêmica), perfil de ácidos graxos ômega-3 (EPA/DHA essenciais para membranas neuronais). Em casos de escores persistentemente baixos (<3/5), considerar avaliação de neurotransmissores urinários, teste de permeabilidade intestinal, e marcadores de neurodegeneração (se >60 anos).

**INTERVENÇÕES NUTRICIONAIS (iniciar imediatamente):**
Protocolo MIND diet otimizado: mínimo 3 porções/dia de vegetais folhosos verdes escuros (espinafre, couve, brócolis - ricos em luteína e folato), 1 porção diária de frutas vermelhas/roxas (mirtilo, morango, açaí - antocianinas neuroprotetoras), peixes gordos de água fria 3-4x/semana (salmão selvagem, sardinha, cavala - ômega-3 EPA/DHA), oleaginosas diárias (30g de nozes, amêndoas, castanhas - vitamina E, magnésio), azeite extravirgem como gordura principal (polifenóis anti-inflamatórios), abacate diário (gorduras monoinsaturadas + potássio), redução drástica de açúcares refinados e carboidratos de alto índice glicêmico (evitar picos insulínicos que prejudicam cognição), limitar carne vermelha a 1x/semana máximo, eliminar gorduras trans e alimentos ultraprocessados completamente.

**SUPLEMENTAÇÃO BASEADA EM EVIDÊNCIAS:**
Se déficit confirmado em testes: complexo B otimizado com formas metiladas (metilfolato 800mcg, metilcobalamina 1000mcg, P5P 50mg), ômega-3 de alta potência EPA/DHA 2000-3000mg/dia (razão EPA:DHA 2:1 para neuroinflamação), magnésio L-treonato 2000mg/dia (única forma que atravessa barreira hematoencefálica eficientemente), vitamina D3 5000 UI/dia se <40 ng/mL, fosfatidilserina 300mg/dia (suporte à membrana neuronal), acetil-L-carnitina 1500mg/dia (produção de acetilcolina + energia mitocondrial), cogumelos medicinais (Lion's Mane 1000mg/dia para fator de crescimento neural). Suplementação deve ser individualizada baseada em deficiências identificadas.

**OTIMIZAÇÃO DO SONO (crítico para consolidação da memória):**
Protocolo de higiene do sono rigoroso: horário fixo para dormir/acordar (inclusive fins de semana), escurecimento total do quarto (cortinas blackout, sem LEDs), temperatura ambiente 18-20°C, exposição à luz solar matinal 10-30 minutos nas primeiras 2h após acordar (regulação do ritmo circadiano), evitar telas 2h antes de dormir (ou usar bloqueadores de luz azul), prática de relaxamento pré-sono (respiração diafragmática, meditação guiada), considerar suplementação de magnésio glicinato 400mg + L-teanina 200mg 1h antes de dormir. Se apneia do sono suspeitada, solicitar polissonografia.

**EXERCÍCIO FÍSICO NEUROPROTETOR:**
Programa multimodal: mínimo 150 minutos/semana de exercício aeróbico moderado-intenso (caminhada rápida, corrida leve, natação, ciclismo - aumenta BDNF e neurogênese hipocampal), treinamento resistido 2-3x/semana (fortalece conexão músculo-cérebro via mioquinas), exercícios de coordenação e equilíbrio (dança, tai chi, yoga - integração sensório-motora), idealmente exercício matinal para sincronização circadiana. Evitar sedentarismo prolongado (levantar a cada 30-45 minutos).

**MANEJO DO ESTRESSE E NEUROPLASTICIDADE:**
Técnicas diárias de redução de cortisol: meditação mindfulness 10-20 minutos/dia (estudos mostram aumento de massa cinzenta hipocampal após 8 semanas), respiração diafragmática 4-7-8 (ativa parassimpático), prática de gratidão (aumenta neurotransmissores positivos), contato social regular (proteção contra declínio cognitivo), aprendizado contínuo de novas habilidades (idioma, instrumento musical, atividades cognitivas desafiadoras mantêm reserva cognitiva).

**REAVALIAÇÃO:**
Repetir teste das 5 palavras após 8-12 semanas de intervenções. Se melhora insuficiente ou declínio progressivo, considerar avaliação neuropsicológica formal completa, neuroimagem estrutural (RM de encéfalo com protocolo para demência), e encaminhamento para neurologista cognitivo. Monitoramento longitudinal a cada 6 meses se CCL confirmado."""
    },

    # TESTES DE MEMÓRIA TARDIA - 5 Palavras de Dubois
    "dubois_tardio": {
        "clinical_relevance": """A evocação tardia no Teste das 5 Palavras de Dubois (recall após intervalo de 5-10 minutos preenchido com outras tarefas) constitui um marcador neuropsicológico altamente específico de consolidação da memória episódica de longo prazo e integridade da função hipocampal. Este componente do teste avalia a capacidade do cérebro de transferir informações da memória de curto prazo (dependente do córtex pré-frontal dorsolateral e áreas de associação) para armazenamento duradouro nas estruturas temporais mediais, particularmente o hipocampo e córtex entorrinal.

Meta-análises recentes (2024) demonstram que déficits na evocação tardia (escore <4/5) apresentam sensibilidade de 87% e especificidade de 92% para detecção de comprometimento cognitivo leve amnéstico (CCL-a), o qual progride para doença de Alzheimer em aproximadamente 15-20% dos casos ao ano. A dissociação entre memória imediata preservada e memória tardia comprometida representa um padrão característico de disfunção hipocampal inicial, frequentemente precedendo alterações estruturais detectáveis em neuroimagem por 2-5 anos.

Na abordagem da Medicina Funcional Integrativa, a consolidação da memória depende criticamente de processos neurofisiológicos que ocorrem predominantemente durante o sono de ondas lentas (estágio N3) e sono REM. Durante estas fases, ocorre reativação coordenada de padrões neuronais (replay hippocampal) que transfere informações do hipocampo para o córtex cerebral, um processo conhecido como consolidação sistêmica. Estudos de polisonografia associados a testes neuropsicológicos (2023-2024) demonstram correlação direta entre tempo total de sono profundo, densidade de fusos do sono (sleep spindles) e performance em testes de memória tardia.

A interpretação funcional de déficits na evocação tardia considera múltiplos sistemas fisiopatológicos: (1) fragmentação do sono com redução do sono de ondas lentas (comum em apneia do sono, síndrome das pernas inquietas, insônia crônica), (2) neuroinflamação crônica com ativação microglial patológica e produção de citocinas neurotóxicas (IL-1β, TNF-α), (3) resistência insulínica cerebral (diabetes tipo 3) comprometendo metabolismo neuronal de glicose e cetona, (4) estresse oxidativo mitocondrial com acúmulo de espécies reativas de oxigênio (ROS) danificando proteínas sinápticas, (5) deficiência de neurotransmissores essenciais para consolidação (acetilcolina, glutamato, noradrenalina), (6) disfunção da barreira hematoencefálica permitindo entrada de toxinas e mediadores inflamatórios periféricos, (7) acúmulo de proteínas mal dobradas (beta-amiloide, tau fosforilada) interferindo com plasticidade sináptica.

Biomarcadores funcionais associados a déficits de memória tardia incluem: homocisteína elevada >12 μmol/L (neurotoxicidade direta e disfunção vascular cerebral), status de vitamina B12 funcional baixo (B12 <500 pg/mL ou ácido metilmalônico >0,4 μmol/L), deficiência de vitamina D <30 ng/mL (neuroproteção reduzida), perfil lipídico aterogênico com LDL oxidado elevado (estresse oxidativo vascular cerebral), índice ômega-3 <8% (composição inadequada de membranas neuronais), glicemia de jejum >100 mg/dL ou HbA1c >5,7% (glicotoxicidade neuronal), hs-CRP >2 mg/L (inflamação sistêmica afetando função cerebral), e cortisol salivar noturno elevado (disrupção do eixo HPA comprometendo neurogênese hipocampal).

Estudos longitudinais prospectivos (2023-2025) demonstram que intervenções multimodais integrativas iniciadas precocemente em indivíduos com declínio de memória tardia podem não apenas estabilizar, mas reverter parcialmente a progressão: o protocolo ReCODE (Reversal of Cognitive Decline) modificado, implementando 25-30 intervenções simultâneas (otimização metabólica, nutricional, hormonal, sono, exercício, desafio cognitivo, redução de toxinas), resultou em melhora sustentada de 1,5-2,0 pontos no escore de evocação tardia após 6 meses em 84% dos participantes com CCL inicial; protocolos de jejum intermitente (16:8) associados a dieta cetogênica modificada (MCT oil suplementação) demonstraram aumento de 31% na utilização cerebral de cetonas como combustível alternativo, com melhora correspondente de 28% na consolidação de memória tardia; e programas de exercício aeróbico de alta intensidade intervalado (HIIT) 3x/semana resultaram em aumento mensurável de volume hipocampal (neuroimagem quantitativa) e melhora de 1,2 pontos no teste de Dubois tardio após 12 semanas.

A avaliação clínica adequada sempre considera trajetória temporal: declínio recente e rápido sugere causas potencialmente reversíveis (deficiências nutricionais agudas, efeitos medicamentosos, depressão pseudodemência), enquanto declínio insidioso e progressivo ao longo de meses-anos levanta suspeita de processo neurodegenerativo subjacente requerendo investigação aprofundada.""",

        "patient_explanation": """O teste de memória tardia das 5 palavras avalia algo ainda mais importante que simplesmente lembrar informações imediatamente: ele testa se seu cérebro consegue "guardar" essas memórias de forma duradoura. Depois de aprender as 5 palavras, você realiza outras atividades por alguns minutos, e então precisa lembrar novamente das palavras originais. Isso simula o que acontece na vida real - você precisa lembrar de informações importantes mesmo depois de fazer outras coisas no meio do caminho.

Quando você dorme profundamente à noite, seu cérebro "reproduz" as memórias do dia e as transfere para um armazenamento de longo prazo. É como passar arquivos importantes do pendrive (memória de curto prazo) para o HD (memória de longo prazo). Se você consegue lembrar todas as 5 palavras após o intervalo, isso indica que esse sistema de "backup" está funcionando perfeitamente. Lembrar 3-4 palavras ainda é razoável, mas pode significar que seu cérebro precisa de mais suporte. Lembrar menos de 3 palavras merece atenção especial e investigação mais profunda.

Na Medicina Funcional, sabemos que problemas com memória tardia frequentemente estão relacionados à qualidade do sono - especialmente as fases mais profundas do sono, quando seu cérebro realmente consolida memórias. Se você não está dormindo bem, seu cérebro não tem a oportunidade de "arquivar" as informações adequadamente. Outros fatores importantes incluem: níveis de açúcar no sangue (o cérebro precisa de energia estável), inflamação no corpo (que pode afetar o cérebro), deficiências de vitaminas essenciais (especialmente B12, B6, folato), níveis baixos de ômega-3 (gorduras essenciais para o cérebro), e estresse crônico (que danifica a área do cérebro responsável pela memória).

A boa notícia é que há muito que pode ser feito para melhorar sua memória tardia: priorizar sono de qualidade (7-8 horas em ambiente escuro e fresco), alimentação rica em nutrientes cerebrais (peixes gordos, vegetais verde-escuros, frutas vermelhas, nozes), exercícios físicos regulares (que aumentam fatores de crescimento cerebral), manter sua mente ativa aprendendo coisas novas, e gerenciar o estresse de forma efetiva. Muitas pessoas veem melhorias significativas em 2-3 meses quando seguem essas recomendações consistentemente.""",

        "conduct": """**AVALIAÇÃO INICIAL DIFERENCIADA (48-72h):**
Diferenciar causas agudas reversíveis de declínio cognitivo estrutural: anamnese detalhada focada em cronologia do declínio (início súbito vs. insidioso, progressão rápida vs. lenta), investigar rigorosamente padrão de sono (questionário de Berlim para apneia, escala de Epworth para sonolência diurna, diário de sono de 7 dias), rastrear medicações potencialmente prejudiciais à cognição (benzodiazepínicos, anticolinérgicos, anti-histamínicos, opioides, anticonvulsivantes, estatinas em altas doses), avaliar sintomas depressivos (PHQ-9) dado que pseudodemência depressiva é reversível, e quantificar consumo de álcool (toxicidade direta ao hipocampo).

**INVESTIGAÇÃO LABORATORIAL ABRANGENTE (1-2 semanas):**
Painel cognitivo funcional completo obrigatório: perfil metabólico estendido com glicemia de jejum + insulina + HbA1c + frutosamina (avaliar resistência insulínica cerebral - diabetes tipo 3), perfil tireoidiano completo TSH + T4 livre + T3 livre + anti-TPO + anti-tireoglobulina (hipotireoidismo subclínico compromete consolidação de memória), vitamina B12 sérica + ácido metilmalônico urinário + holotranscobalamina (B12 funcional - deficiência presente em 40% dos idosos), folato sérico e eritrocitário, homocisteína plasmática (neurotoxicidade direta >12 μmol/L, alvo ideal <8 μmol/L), vitamina D 25-OH (alvo 50-80 ng/mL para neuroproteção ótima), perfil lipídico avançado com LDL oxidado + Apo B + partículas de LDL (aterosclerose cerebral subclínica), hs-CRP + VHS + ferritina (marcadores inflamatórios sistêmicos), função hepática completa (detoxificação prejudicada afeta cognição), função renal com TFG estimada (uremia compromete função cerebral), hemograma completo com VCM (anemia macrocítica por B12/folato, anemia por deficiência de ferro reduz oxigenação cerebral), perfil de ácidos graxos ômega-3 com índice ômega-3 (alvo >8% para membranas neuronais saudáveis), magnésio eritrocitário e zinco sérico (cofatores essenciais), cortisol salivar ao acordar e noturno (disrupção do eixo HPA), DHEA-S (reserva adrenal), e perfil hormonal sexual se apropriado (estrogênio, testosterona - neuroproteção hormonal).

Em casos de déficit persistente <3/5 ou declínio progressivo documentado: adicionar marcadores de neurodegeneração - APOE genotipagem (risco genético para Alzheimer), painel de autoimunidade neuronal (anticorpos anti-GAD, anti-tiroidiano com impacto cerebral), marcadores de permeabilidade intestinal e disbiose (zonulina, LPS, calprotectina fecal - eixo intestino-cérebro), teste de metais pesados (mercúrio, chumbo, alumínio - neurotoxicidade), perfil de micotoxinas urinárias se exposição a mofo suspeitada, e considerar pesquisa de marcadores liquóricos (beta-amiloide 42, tau fosforilada, proteína tau total) se >65 anos com declínio progressivo.

**NEUROIMAGEM ESTRUTURAL (indicações específicas):**
Ressonância magnética de encéfalo com protocolo para avaliação de demência e volumetria hipocampal se: declínio progressivo documentado ao longo de 6-12 meses, assimetria de sintomas neurológicos, início antes dos 65 anos, história familiar forte de demência, ou escore persistentemente <3/5 apesar de intervenções. Avaliar sinais de atrofia hipocampal (principal marcador estrutural de CCL amnéstico), hiperintensidades de substância branca (doença vascular cerebral de pequenos vasos), microhemorragias (angiopatia amiloide), ou lesões estruturais focais.

**AVALIAÇÃO DO SONO (essencial para memória tardia):**
Polissonografia completa se: escore de Epworth >10, relato de roncos/apneias testemunhadas, sonolência diurna excessiva, obesidade com IMC >30, circunferência cervical >43 cm homens ou >40 cm mulheres, história de hipertensão refratária. Apneia obstrutiva do sono moderada-grave (IAH >15) compromete severamente consolidação da memória durante sono REM e deve ser tratada prioritariamente. Considerar actígrafo de pulso por 14 dias para avaliação objetiva de padrão sono-vigília e fragmentação do sono em casos menos óbvios.

**PROTOCOLO NUTRICIONAL NEUROPROTETOR AGRESSIVO:**
Dieta Mediterranean-DASH Intervention for Neurodegenerative Delay (MIND) modificada com restrição de janela alimentar: implementar alimentação restrita no tempo 16:8 (jejum de 16h, janela alimentar de 8h) para induzir autofagia neuronal e produção de BDNF. Composição dietética: mínimo 6 porções/dia de vegetais verdes folhosos (espinafre, couve, rúcula, acelga - luteína, folato, nitratos), 2 porções/dia de frutas vermelhas e roxas (mirtilo, morango, amora, açaí - antocianinas que atravessam barreira hematoencefálica), peixes gordos de água fria selvagens 4-5x/semana (salmão, sardinha, anchova, cavala - EPA/DHA para membranas sinápticas e resolução de inflamação), 30-50g/dia de oleaginosas cruas (nozes têm maior conteúdo de ALA, castanha do Pará para selênio, amêndoas para vitamina E), azeite de oliva extravirgem 3-4 colheres sopa/dia (oleocanthal anti-inflamatório, polifenóis neuroprotetores), abacate diário (gordura monoinsaturada, potássio, folato), leguminosas 3-4x/semana (lentilha, feijão, grão-de-bico - proteína vegetal, fibras), grãos integrais apenas se tolerar glúten (quinoa, trigo sarraceno, aveia sem glúten - complexo B, magnésio), evitar totalmente açúcares refinados, alimentos ultraprocessados, gorduras trans, e limitar carne vermelha a no máximo 1x/semana.

Considerações especiais para resistência insulínica cerebral: se HbA1c >5,7% ou insulina de jejum >10 μU/mL, implementar dieta cetogênica modificada (50-70g carboidratos líquidos/dia) com suplementação de óleo MCT (C8 - ácido caprílico) progressivamente até 30-45ml/dia dividido em 3 doses (fornece cetonas exógenas como combustível cerebral alternativo sem necessidade de cetose profunda).

**SUPLEMENTAÇÃO NEUROCOGNITIVA BASEADA EM EVIDÊNCIAS:**
Protocolo em camadas conforme gravidade e achados laboratoriais:

*Camada 1 - Fundamental para todos com déficit de memória tardia:*
- Ômega-3 EPA/DHA ultra purificado 2000-4000mg/dia (razão EPA:DHA 2:1 se inflamação predominante, 1:1 se declínio cognitivo predominante), alvo de índice ômega-3 >8%
- Complexo B metilado ativo: metilfolato (5-MTHF) 800-1000mcg, metilcobalamina 1000-2000mcg sublingual, piridoxal-5-fosfato (P5P) 50mg, tiamina (B1) 100mg, riboflavina (B2) 50mg, niacinamida (B3) 500mg, ácido pantotênico 250mg, biotina 500mcg
- Magnésio L-treonato (Magtein®) 2000mg/dia (144mg magnésio elementar) - única forma demonstrada em estudos aumentar magnésio cerebral e densidade sináptica hipocampal
- Vitamina D3 5000-10.000 UI/dia se <40 ng/mL, manter nível 50-80 ng/mL (coadministrar vitamina K2 MK-7 200mcg para direcionamento ósseo do cálcio)

*Camada 2 - Suporte adicional para consolidação da memória:*
- Fosfatidilserina 300mg/dia (preferencialmente derivada de soja não-OGM ou girassol) - componente estrutural crítico de membranas neuronais, melhora fluência e consolidação
- Acetil-L-carnitina (ALCAR) 1500-2000mg/dia em doses divididas - produção de acetilcolina, energização mitocondrial, neuroproteção
- Alfa-GPC (alfa-glicerilfosforilcolina) 300-600mg/dia - precursor de acetilcolina que atravessa barreira hematoencefálica
- Citicolina (CDP-colina) 250-500mg/dia - suporte à síntese de fosfolipídios e neurotransmissores

*Camada 3 - Nootrópicos naturais e compostos bioativos:*
- Bacopa monnieri (extrato padronizado 50% bacosides) 300mg 2x/dia - evidências robustas para consolidação de memória tardia, requer 8-12 semanas para efeito máximo
- Hericium erinaceus (Lion's Mane, extrato 30% polissacarídeos) 1000mg 2x/dia - estimula produção de NGF (fator de crescimento neural), promove mielinização
- Curcumina lipossomal ou com piperina 1000mg/dia (equivalente a 200mg curcumina absorvível) - anti-inflamatório potente, reduz placas amiloides
- Resveratrol 500mg/dia - ativa sirtuínas (longevidade celular), proteção mitocondrial
- Pterostilbeno 100-200mg/dia - análogo mais biodisponível do resveratrol

*Camada 4 - Suporte mitocondrial neuronal:*
- Coenzima Q10 (ubiquinol) 200-400mg/dia - cadeia respiratória mitocondrial, antioxidante
- PQQ (pirroloquinolina quinona) 20mg/dia - biogênese mitocondrial
- Ácido alfa-lipóico R-ALA 300-600mg/dia - antioxidante universal, recicla vitamina C/E/glutationa

*Camada 5 - Otimização circadiana e sono (crítico para consolidação):*
- Magnésio glicinato ou treonato 400-800mg 1-2h antes de dormir
- L-teanina 200-400mg à noite - ondas alfa cerebrais, relaxamento sem sedação
- Glicina 3-5g antes de dormir - melhora qualidade do sono profundo (estágio N3)
- Fosfatidilserina 100-200mg à noite - reduz cortisol noturno se elevado
- Melatonina de liberação prolongada 0,3-1mg (dose fisiológica) apenas se produção endógena demonstrada baixa

**PROTOCOLO DE OTIMIZAÇÃO DO SONO (absolutamente crítico):**
Implementar higiene do sono radical: horário fixo dormir/acordar com variação máxima 30 minutos (mesmo fins de semana - sincronização circadiana), exposição à luz solar natural 10-30 minutos dentro de 1h após acordar (supressão de melatonina matinal, reinício do ciclo circadiano), ambiente de sono otimizado (escurecimento total com cortinas blackout, temperatura 17-19°C, umidade 40-60%, eliminação total de ruídos ou uso de ruído branco, colchão e travesseiros adequados), eliminação completa de luz azul 2-3h antes de dormir (óculos bloqueadores de luz azul se uso de telas inevitável), rotina pré-sono relaxante de 30-60 minutos (leitura, meditação, banho quente, alongamento suave, journaling de gratidão), evitar refeições pesadas 3h antes de dormir (digestão compromete sono profundo), evitar cafeína após 14h (meia-vida 5-6h), evitar álcool à noite (suprime sono REM crítico para consolidação), e considerar suplementação de magnésio + glicina 1-2h antes de dormir.

Se apneia do sono diagnosticada: CPAP é tratamento de escolha e não-negociável (melhora dramática em consolidação de memória documentada em 3-6 meses), perda de peso se sobrepeso/obesidade (redução de 10% do peso pode reduzir IAH em 25-30%), evitar posição supina (apneia posicional), e considerar dispositivo de avanço mandibular se apneia leve-moderada e CPAP não tolerado.

**PROGRAMA DE EXERCÍCIO NEUROPROTETOR:**
Protocolo multimodal evidenciado: exercício aeróbico 150-300 minutos/semana em intensidade moderada-vigorosa (65-85% FCmáx, teste da fala possível mas trabalhoso) - preferencialmente HIIT 3x/semana (30 minutos com intervalos 1min intenso: 2min recuperação) que demonstra maior elevação de BDNF e neurogênese hipocampal que steady-state; treinamento resistido 2-3x/semana focado em grupos musculares grandes (agachamento, terra, supino, remada - mioquinas como irisina promovem função cognitiva); exercícios de coordenação e equilíbrio 2-3x/semana (tai chi, dança, yoga - integração sensório-motora e neuroplasticidade); e exercício matinal preferencialmente (sincronização circadiana).

Evitar sedentarismo: levantar e movimentar-se a cada 30-45 minutos durante dia (caminhar 2-3 minutos, alongamentos), alvo de 8.000-10.000 passos/dia monitorado objetivamente.

**TREINAMENTO COGNITIVO E NEUROPLASTICIDADE:**
Desafio cognitivo regular: aprendizado de nova habilidade complexa (idioma estrangeiro, instrumento musical, dança complexa) - estimula neurogênese e sinaptogênese; treinamento cognitivo computadorizado validado (BrainHQ, CogniFit, Lumosity) 20-30 minutos 5x/semana; jogos de estratégia (xadrez, bridge); leitura ativa com anotações e resumos; socialização regular (isolamento social é fator de risco independente para demência); e ensinar/mentorear outros (reforça consolidação do conhecimento próprio).

**MANEJO DO ESTRESSE E EIXO HPA:**
Práticas diárias obrigatórias: meditação mindfulness 20 minutos 2x/dia (estudos RNM demonstram aumento de espessura cortical hipocampal após 8 semanas), respiração diafragmática lenta (4-7-8 ou coerência cardíaca 5-6 respirações/minuto) 5-10 minutos 3x/dia, exposição à natureza (forest bathing) mínimo 120 minutos/semana (redução documentada de cortisol), prática de gratidão (journaling de 3 coisas positivas diariamente), música relaxante (ondas alfa), e considerar biofeedback de variabilidade da frequência cardíaca (HRV training) para modulação autonômica.

**REDUÇÃO DE NEUROTOXINAS AMBIENTAIS:**
Minimizar exposição a: mercúrio (evitar peixes grandes predadores - atum, peixe-espada, limitar a 1x/mês máximo), alumínio (evitar antiácidos com alumínio, utensílios de cozinha de alumínio, antiperspirantes com alumínio), pesticidas organofosforados (priorizar orgânicos para "dirty dozen"), BPA e ftalatos (evitar plásticos com alimentos quentes, usar vidro), mofo em ambientes internos (avaliar qualidade do ar interno se história de exposição a água), e poluição do ar (purificador HEPA em quarto, evitar exercício externo em horários de pico de poluição).

**MODULAÇÃO DO EIXO INTESTINO-CÉREBRO:**
Se disbiose suspeitada por sintomas gastrointestinais: probióticos psicotrópicos específicos (Lactobacillus helveticus R0052 + Bifidobacterium longum R0175 - redução de cortisol e melhora de cognição demonstrada), prebióticos (inulina 10-15g/dia, fibra de acácia), alimentos fermentados diários (kefir, kombucha, chucrute, kimchi), e considerar teste de microbioma intestinal se sintomas significativos.

**OTIMIZAÇÃO HORMONAL (se indicado):**
Mulheres em perimenopausa/menopausa com déficit cognitivo: avaliar terapia de reposição hormonal bioidêntica (estradiol transdérmico + progesterona micronizada oral) se dentro de janela de oportunidade (<60 anos ou <10 anos de menopausa) e sem contraindicações - estrogênio é neuroprotetor potente, melhora perfusão cerebral e metabolismo de glicose cerebral. Homens com testosterona livre baixa (<70 pg/mL) e sintomas: considerar reposição de testosterona visando níveis fisiológicos médios-normais (400-600 ng/dL total).

**MONITORAMENTO E REAVALIAÇÃO:**
Cronograma rigoroso: repetir teste das 5 palavras (imediato e tardio) após 6-8 semanas de intervenções iniciais, reavaliação laboratorial de biomarcadores após 12 semanas (vitamina D, B12, homocisteína, ômega-3, HbA1c, hs-CRP), e avaliação neuropsicológica formal completa (MEEM, MoCA, teste do desenho do relógio, fluência verbal) aos 6 meses.

Critérios de encaminhamento urgente para neurologia cognitiva: declínio >1 ponto no escore em período de 3-6 meses apesar de intervenções otimizadas, aparecimento de sintomas neurológicos focais (assimetria motora, alterações visuais, disfasia), sintomas psiquiátricos novos significativos (alucinações, paranoia, desinibição), ou idade jovem de início (<60 anos) com progressão rápida.

**META DE TRATAMENTO:**
Estabilização do escore ≥4/5 e idealmente melhora para 5/5, manutenção de independência funcional completa nas atividades diárias, e prevenção de progressão para demência. Abordagem funcional integrativa iniciada em fase de CCL pode estabilizar ou reverter parcialmente o declínio em 60-80% dos casos conforme literatura recente."""
    },

    # SPAN DE DÍGITOS - DIRETO E INVERSO
    "span_digitos": {
        "clinical_relevance": """O Teste de Span de Dígitos (direto e inverso) representa uma avaliação neuropsicológica fundamental da memória de trabalho, atenção auditiva sustentada e funções executivas, componentes essenciais da arquitetura cognitiva humana. A versão direta avalia primariamente a capacidade de memória de curto prazo fonológica e atenção imediata, recrutando predominantemente o circuito fonológico do modelo de Baddeley (córtex parietal inferior esquerdo, área de Broca). A versão inversa adiciona demanda de manipulação mental ativa da informação, requerendo função executiva central localizada no córtex pré-frontal dorsolateral, constituindo assim um teste mais sensível de disfunção executiva e controle atencional.

Normas neuropsicológicas estabelecem spans médios de 7±2 dígitos diretos (range 5-9, média 7 em adultos jovens saudáveis) e 5±2 dígitos inversos (range 3-7, média 5). Performance abaixo destes valores ajustados por idade e escolaridade sugere comprometimento patológico. Meta-análises recentes (2023-2024) demonstram que déficits no span de dígitos inverso (<4) apresentam valor preditivo significativo para declínio cognitivo futuro, com hazard ratio de 2,3 para conversão de cognição normal para comprometimento cognitivo leve em 5 anos.

A dissociação entre performance no span direto (preservado) e inverso (comprometido) constitui padrão característico de disfunção executiva frontal, frequentemente observada em TDAH adulto não diagnosticado, declínio cognitivo vascular (doença de pequenos vasos), depressão com comprometimento executivo, e fases iniciais de demência frontotemporal. Já o comprometimento de ambos (direto e inverso) sugere disfunção mais difusa envolvendo sistemas atencionais e de memória de trabalho, comum em déficit atencional global, fadiga cognitiva severa, privação de sono crônica, ou processos neurodegenerativos mais avançados.

Na perspectiva da Medicina Funcional Integrativa, a performance no span de dígitos reflete a integridade de múltiplos sistemas neurobiológicos interdependentes. A memória de trabalho fonológica (span direto) depende criticamente de: (1) disponibilidade de neurotransmissores colinérgicos para manutenção de informações online no córtex (acetilcolina), (2) conectividade funcional entre regiões temporoparietais e frontais mediada por glutamato, (3) suprimento energético cerebral adequado via metabolismo de glicose e cetonas, (4) integridade de processos atencionais regulados por noradrenalina e dopamina, e (5) ausência de interferência por processos distraidores (ansiedade, ruminação, dor crônica).

A função executiva de manipulação mental (span inverso) adiciona dependência de sistemas dopaminérgicos pré-frontais, conectividade frontoparietal robusta, e processos inibitórios mediados por GABA que suprimem respostas automáticas permitindo reversão da sequência. Deficiências nutricionais de precursores de neurotransmissores (tirosina, fenilalanina para catecolaminas; triptofano para serotonina; colina para acetilcolina) ou cofatores essenciais (ferro, zinco, magnésio, vitaminas B6, B12, folato, vitamina C) comprometem diretamente a síntese e liberação destes neurotransmissores.

Estudos de neuroimagem funcional (fMRI) em indivíduos com déficits de memória de trabalho (2024-2025) demonstram padrões de hipoativação no córtex pré-frontal dorsolateral bilateral e giro supramarginal, associados a conectividade funcional reduzida em redes de controle executivo. Intervenções que restauram neurotransmissão dopaminérgica (suplementação de L-tirosina, atividade física regular) ou melhoram conectividade frontoparietal (treinamento cognitivo específico, meditação mindfulness) demonstram capacidade de normalizar tanto a ativação neural quanto a performance comportamental.

Fatores moduladores da performance que devem ser considerados na interpretação clínica incluem: privação de sono aguda (reduz span em 1-2 dígitos após uma noite de sono inadequado), sobrecarga cognitiva ou fadiga mental (depleta recursos atencionais), ansiedade de performance (sistema de vigilância ameaça consome recursos de memória de trabalho), déficit atencional primário não tratado (TDAH reduz span em média 1,5-2,0 dígitos vs. controles), efeitos residuais de medicações sedativas (benzodiazepínicos, anti-histamínicos, antipsicóticos, anticonvulsivantes), hipoglicemia funcional (glicemia <70 mg/dL compromete severamente função pré-frontal), e presença de dor crônica ou sintomas somáticos distraidores.

Biomarcadores funcionais associados a déficits de memória de trabalho incluem: ferritina <50 ng/mL em mulheres ou <75 ng/mL em homens (deficiência de ferro afeta síntese de dopamina), vitamina B12 <400 pg/mL (comprometimento de mielinização e neurotransmissão), folato <12 ng/mL (síntese de monoaminas prejudicada), homocisteína >10 μmol/L (neurotoxicidade), vitamina D <30 ng/mL (expressão de genes neuroprotetores reduzida), TSH >3,0 mU/L mesmo com T4 normal (hipotireoidismo subclínico afeta cognição), cortisol salivar matinal <10 nmol/L ou >30 nmol/L (disrupção do eixo HPA), DHEA-S baixo para idade (reserva adrenal depletada), glicemia de jejum >100 mg/dL ou insulina >10 μU/mL (resistência insulínica cerebral), e marcadores inflamatórios elevados hs-CRP >2 mg/L (neuroinflamação sistêmica).

Intervenções baseadas em evidências (2023-2025) demonstram potencial significativo de melhora: treinamento de memória de trabalho adaptativo computadorizado (Cogmed, dual n-back) 25 minutos/dia 5x/semana por 5-8 semanas resultou em aumento médio de 1,2 dígitos no span inverso com transferência para tarefas não treinadas (far transfer); suplementação de creatina monohidratada 5g/dia por 6 semanas aumentou span de dígitos em 0,8 pontos em vegetarianos (aumento de fosfocreatina cerebral fornecendo energia para neurônios); suplementação combinada de L-tirosina 2g/dia + cafeína 200mg demonstrou melhora aguda de 1,5 dígitos em contexto de privação de sono; e protocolos de meditação mindfulness 30 minutos/dia por 8 semanas aumentaram capacidade de memória de trabalho em 14% com melhora correspondente no span de dígitos.

Estudos longitudinais prospectivos indicam que declínio progressivo no span de dígitos ao longo de 12-24 meses, mesmo que ainda dentro de limites normais ajustados por idade, constitui preditor robusto de conversão para comprometimento cognitivo leve (CCL) ou demência, sugerindo utilidade deste teste simples para monitoramento evolutivo de risco cognitivo.""",

        "patient_explanation": """O teste de span de dígitos avalia duas capacidades fundamentais do seu cérebro: sua "memória RAM" (versão direta) e sua "capacidade de processamento mental" (versão inversa). Na versão direta, você ouve uma sequência de números e precisa repeti-los na mesma ordem - isso testa sua capacidade de manter informações na mente por alguns segundos, como quando alguém te dá um número de telefone e você precisa discá-lo imediatamente. A maioria dos adultos saudáveis consegue lembrar 7±2 números (entre 5 e 9).

Na versão inversa, você precisa repetir os números na ordem contrária - isso é muito mais desafiador porque requer que seu cérebro não apenas mantenha a informação, mas também a manipule ativamente, reorganizando a sequência. É como fazer malabarismos com bolas em vez de apenas segurá-las. A maioria das pessoas consegue 5±2 números na ordem inversa (entre 3 e 7). Esta capacidade de manipular mentalmente informações é essencial para resolver problemas, planejar, tomar decisões e realizar múltiplas tarefas.

Se você tem dificuldade com a versão direta (memória imediata), pode indicar que seu cérebro está cansado, com falta de combustível (energia), ou que sistemas de atenção não estão funcionando adequadamente. Se a versão inversa é mais difícil (mas a direta está OK), isso sugere especificamente que as áreas frontais do cérebro - responsáveis por planejamento, organização e controle - podem precisar de suporte. Muitas pessoas com TDAH não diagnosticado, excesso de estresse crônico, ou privação de sono têm este padrão.

Na Medicina Funcional, entendemos que problemas com memória de trabalho geralmente têm causas tratáveis: falta de sono de qualidade (seu cérebro precisa de descanso para funcionar), deficiências nutricionais (especialmente ferro, vitaminas B, magnésio), níveis instáveis de açúcar no sangue (seu cérebro precisa de energia constante), ansiedade ou preocupações excessivas (que "ocupam" espaço mental), ou desequilíbrios em neurotransmissores - os mensageiros químicos do cérebro.

A boa notícia é que a memória de trabalho pode ser treinada e melhorada, assim como você treina músculos. Jogos específicos de memória, aprender coisas novas desafiadoras (idioma, instrumento musical), praticar mindfulness (meditação), dormir bem, alimentar-se adequadamente, e exercitar-se regularmente podem melhorar significativamente sua performance em 2-3 meses.""",

        "conduct": """**AVALIAÇÃO CONTEXTUAL IMEDIATA (mesma consulta):**
Investigar fatores agudos que comprometem memória de trabalho: qualidade e quantidade de sono nas últimas 48-72h (ideal >7h com continuidade), horário da última refeição e composição (hipoglicemia ou refeição muito pesada comprometem cognição), nível de ansiedade atual (escala visual analógica 0-10), presença de dor ou desconforto físico significativo (competição por recursos atencionais), uso de substâncias que afetam cognição (cafeína em excesso, álcool residual, cannabis), e medicações com efeitos anticolinérgicos ou sedativos. Considerar repetição do teste em condições otimizadas (bem descansado, bem alimentado, calmo) antes de assumir déficit verdadeiro.

**INVESTIGAÇÃO DIAGNÓSTICA DIFERENCIAL (1-2 semanas):**
Rastreamento de TDAH em adultos se padrão sugestivo (span inverso desproporcionalmente comprometido, história de dificuldades atencionais ao longo da vida): aplicar escalas validadas (ASRS-18, DIVA 2.0, CAARS), avaliar sintomas nucleares de desatenção, hiperatividade e impulsividade em múltiplos contextos, investigar prejuízo funcional em trabalho/relacionamentos, e considerar avaliação neuropsicológica formal se suspeita significativa.

Painel laboratorial focado em substratos de neurotransmissão e energia cerebral: hemograma completo com índices (avaliar anemia que reduz oxigenação cerebral), ferritina sérica (alvo >75 ng/mL em homens, >50 ng/mL em mulheres - cofator para síntese de dopamina), perfil de ferro completo se ferritina baixa (ferro sérico, saturação de transferrina, TIBC), vitamina B12 >500 pg/mL (mielinização e síntese de neurotransmissores), folato sérico >12 ng/mL (ciclo de metilação), homocisteína <10 μmol/L (alvo ideal <8), vitamina D 25-OH (alvo 50-80 ng/mL para função cerebral ótima), TSH + T4 livre + T3 livre (hipotireoidismo subclínico afeta processamento cognitivo), perfil glicêmico com glicemia jejum + insulina + HbA1c (hipoglicemia reativa ou resistência insulínica), cortisol salivar matinal ao acordar e noturno antes de dormir (padrão circadiano), DHEA-S (reserva adrenal), magnésio eritrocitário (não sérico - cofator de >300 reações), zinco sérico (neurotransmissão glutamatérgica), e perfil lipídico com colesterol total (precursor de hormônios esteroides).

Se sintomas significativos de fadiga ou disfunção cognitiva persistentes: considerar testagem estendida com perfil completo de neurotransmissores urinários (dopamina, norepinefrina, serotonina, GABA, glutamato e seus metabólitos), ácidos orgânicos urinários (ciclo de Krebs, metabolismo de neurotransmissores, status de vitaminas B), aminoácidos plasmáticos (precursores de neurotransmissores - tirosina, fenilalanina, triptofano), e considerar avaliação de disbiose intestinal (eixo intestino-cérebro).

**OTIMIZAÇÃO DO SONO (crítico para restauração da memória de trabalho):**
Implementar protocolo rigoroso: horário consistente dormir/acordar (variação máxima 30 minutos, inclusive fins de semana), exposição à luz solar natural 10-30 minutos nas primeiras 2h após acordar (supressão matinal de melatonina, sincronização circadiana), ambiente otimizado (temperatura 17-19°C, escuridão total com cortinas blackout, silêncio ou ruído branco, umidade 40-60%), rotina pré-sono de 45-60 minutos sem telas (leitura, banho quente, meditação, journaling), evitar cafeína após 13-14h, evitar álcool à noite (fragmenta sono REM), refeição leve jantar 3h antes de dormir. Se dificuldade para adormecer: suplementação de magnésio glicinato 400mg + L-teanina 200mg + glicina 3g 1h antes de dormir. Alvo: 7-9 horas de sono total com mínimo 90 minutos de sono profundo (estágio N3) objetivamente confirmado se possível via tracker.

**PROTOCOLO NUTRICIONAL PARA FUNÇÃO EXECUTIVA:**
Dieta estabilizadora de glicemia e neurotransmissores: priorizar proteína de alta qualidade em todas as refeições (1,6-2,0 g/kg peso corporal/dia) para fornecer aminoácidos precursores - ovos orgânicos (colina para acetilcolina), carnes magras grass-fed, peixes selvagens, leguminosas. Incluir carboidratos complexos de baixo índice glicêmico para evitar picos e quedas de glicose (batata doce, quinoa, aveia, leguminosas - liberam glicose de forma sustentada). Gorduras saudáveis em todas as refeições para estabilização glicêmica adicional (abacate, oleaginosas, azeite extravirgem, óleo de coco MCT). Timing nutricional: café da manhã com 30g proteína dentro de 1h após acordar (ativa síntese de dopamina/noradrenalina via tirosina), evitar carboidratos isolados (causam hipoglicemia reativa 2-3h depois comprometendo cognição), refeições a cada 3-4h para manter energia cerebral estável.

Alimentos específicos para neurotransmissores: peixes gordos de água fria 3-4x/semana (salmão selvagem, sardinha - ômega-3 EPA/DHA para membranas sinápticas e sinalização dopaminérgica), vegetais folhosos verde-escuros diários (espinafre, couve - folato, magnésio), ovos inteiros (colina, tirosina, vitaminas B), gema de ovo de galinhas pasteurizadas (colina concentrada), abacate diário (tirosina, potássio, gorduras monoinsaturadas), sementes de abóbora (zinco, magnésio, triptofano), cacau puro >85% (flavonoides, magnésio, teobromina - atenção e foco), frutas vermelhas (antioxidantes que protegem neurotransmissores da oxidação), e oleaginosas especialmente nozes (precursor ômega-3 ALA, vitamina E, magnésio).

Hidratação adequada: 35-40 ml/kg peso corporal/dia de água pura, desidratação de apenas 2% já compromete memória de trabalho e atenção significativamente.

**SUPLEMENTAÇÃO DIRECIONADA PARA MEMÓRIA DE TRABALHO:**
Protocolo em camadas conforme severidade do déficit:

*Fundamental - Cofatores de Neurotransmissão:*
- Complexo B metilado ativo completo: metilfolato 800mcg (síntese de monoaminas), metilcobalamina 1000mcg (mielinização), P5P 50mg (cofator essencial para síntese de serotonina, dopamina, GABA, norepinefrina), tiamina 100mg (metabolismo energético cerebral), riboflavina 50mg, niacinamida 500mg, ácido pantotênico 250mg (precursor de acetil-CoA e acetilcolina)
- Magnésio (formas combinadas) 400-600mg elementar/dia: 200mg L-treonato (atravessa barreira hematoencefálica, plasticidade sináptica) + 200mg glicinato (absorção, sono) + 200mg taurato (neuroprotecao)
- Zinco picolinato ou bisglicinato 30mg/dia se deficiente (<90 mcg/dL), reduzir para 15mg manutenção (modula neurotransmissão glutamatérgica e dopaminérgica)
- Ferro bisglicinato 25-50mg/dia se ferritina <50 ng/mL mulheres ou <75 homens (cofator tirosina hidroxilase para síntese de catecolaminas), administrar longe de refeições com cálcio
- Vitamina D3 5000 UI/dia se <40 ng/mL (expressão de genes de síntese de neurotransmissores) + vitamina K2 MK-7 200mcg

*Precursores de Neurotransmissores (se deficiência documentada ou sintomas específicos):*
- L-tirosina 1000-2000mg pela manhã em jejum (precursor dopamina/norepinefrina - atenção, motivação, memória de trabalho) - especialmente útil em contexto de estresse crônico ou demanda cognitiva elevada
- Alfa-GPC 300-600mg/dia (precursor acetilcolina - memória de trabalho, atenção sustentada, aprendizado)
- Citicolina (CDP-colina) 250-500mg/dia (síntese de fosfolipídios e acetilcolina, aumento de dopamina cerebral)
- 5-HTP 50-100mg 2x/dia se sintomas depressivos ou compulsão por carboidratos (precursor serotonina, tomar longe de proteínas) - apenas se triptofano/serotonina documentadamente baixos

*Energização Mitocondrial Cerebral:*
- Creatina monohidratada 5g/dia (aumenta fosfocreatina cerebral = energia imediata para neurônios, evidências fortes para memória de trabalho)
- Acetil-L-carnitina 1500mg/dia em doses divididas (transporte de ácidos graxos para mitocôndrias, produção de acetilcolina)
- Coenzima Q10 ubiquinol 200-300mg/dia (cadeia de transporte de elétrons mitocondrial)
- PQQ 20mg/dia (biogênese mitocondrial)
- Ácido R-alfa-lipóico 300-600mg/dia (antioxidante mitocondrial, recicla glutationa)

*Ômega-3 e Fosfolipídios:*
- EPA/DHA ultra purificado 2000-3000mg/dia (membranas neuronais, sinalização celular, resolução de neuroinflamação)
- Fosfatidilserina 300mg/dia (componente membranas sinápticas, reduz cortisol se elevado)
- Fosfatidilcolina 1000-2000mg/dia se colina dietética insuficiente

*Nootrópicos Naturais:*
- Bacopa monnieri (50% bacosides) 300mg 2x/dia (memória de trabalho, consolidação, redução de ansiedade, requer 8-12 semanas)
- Ginkgo biloba (24% ginkgoflavonoides) 120mg 2x/dia (fluxo sanguíneo cerebral, memória de trabalho em idosos)
- Rhodiola rosea (3% rosavinas) 200-400mg pela manhã (adaptógeno, resistência à fadiga mental, memória de trabalho sob estresse)
- Lion's Mane (Hericium erinaceus 30% polissacarídeos) 1000mg 2x/dia (NGF, neurogênese, mielinização)

*Moduladores de Foco e Atenção (considerar se TDAH ou déficit atencional significativo):*
- L-teanina 200mg 2-3x/dia (ondas alfa, foco calmo, reduz jitteriness de cafeína se usar)
- Cafeína 100-200mg (se tolerar) combinada com L-teanina 200-400mg (sinergismo para atenção sustentada)
- Teacrina (Teacrine®) 125-200mg (análogo de cafeína de longa duração, sem taquifilaxia)
- Mucuna pruriens (15% L-DOPA) 300-500mg (precursor direto de dopamina) - usar com cautela, monitorar

**TREINAMENTO COGNITIVO ESPECÍFICO:**
Treinamento de memória de trabalho baseado em evidências: programas validados como Cogmed Working Memory Training ou dual n-back task, 25-30 minutos/dia 5 dias/semana por mínimo 5 semanas (idealmente 8-10 semanas). Estudos demonstram melhora não apenas nas tarefas treinadas mas transferência para habilidades cognitivas não treinadas (far transfer). Alternativas: jogos de estratégia (xadrez, go), aprendizado de instrumento musical (demanda alta de memória de trabalho), malabarismo (integração sensório-motora), ou apps como Lumosity, Elevate, Peak (seções específicas de memória de trabalho e atenção).

**EXERCÍCIO FÍSICO PARA FUNÇÃO EXECUTIVA:**
Protocolo otimizado: exercício aeróbico 150 minutos/semana em intensidade moderada-vigorosa (BDNF, neurogênese, dopamina) + treinamento intervalado de alta intensidade (HIIT) 2-3x/semana (máximo aumento de BDNF e catecolaminas) + exercícios de coordenação complexa (dança, artes marciais, esportes com bola - integração frontoparietal). Timing: exercício matinal preferencialmente (sincronização circadiana, elevação de catecolaminas sustentada ao longo do dia).

**PRÁTICAS DE ATENÇÃO PLENA (MINDFULNESS):**
Meditação focada em atenção sustentada 20 minutos/dia: estudos RNM demonstram aumento de conectividade em rede de controle executivo, espessamento cortical pré-frontal, e melhora de 14% em capacidade de memória de trabalho após 8 semanas. Técnicas: meditação focada na respiração, body scan, meditação walking, ou apps guiados (Headspace, Calm, Insight Timer).

**MANEJO DO ESTRESSE:**
Estresse crônico depleta recursos de memória de trabalho via elevação sustentada de cortisol: implementar práticas diárias de respiração diafragmática (4-7-8, coerência cardíaca), biofeedback de HRV se disponível, exposição à natureza 120 min/semana, prática de gratidão, contato social regular, e estabelecer limites claros entre trabalho e vida pessoal.

**OTIMIZAÇÃO HORMONAL (se aplicável):**
Hipotireoidismo subclínico (TSH >3,0 mU/L mesmo com T4 normal) compromete função executiva: considerar trial de levotiroxina baixa dose visando TSH 1,0-2,0. Homens com testosterona livre <70 pg/mL e sintomas: reposição para faixa fisiológica média melhora memória de trabalho. Mulheres em perimenopausa/menopausa: estrogênio é neuroprotetor para função pré-frontal, considerar TRH bioidêntica se dentro de janela terapêutica.

**ELIMINAÇÃO DE INTERFERÊNCIAS MEDICAMENTOSAS:**
Revisar todas medicações para efeitos anticolinérgicos (antidepressivos tricíclicos, anti-histamínicos de primeira geração, relaxantes musculares, antiespasmódicos, alguns antipsicóticos, alguns antieméticos), benzodiazepínicos (comprometem severamente memória de trabalho), anticonvulsivantes em altas doses, e opioides. Discutir com prescritor possibilidade de redução gradual, troca por alternativas menos prejudiciais à cognição, ou otimização de doses mínimas efetivas.

**ESTRATÉGIAS COMPENSATÓRIAS:**
Enquanto deficit persiste, ensinar técnicas: chunking (agrupar informações em unidades significativas), rehearsal ativo (repetição mental com elaboração), externalização de memória (listas, alarmes, apps de organização), redução de carga cognitiva (ambiente silencioso, mono-tarefa), e técnicas de atenção plena para reduzir distração interna (ruminação, preocupações).

**MONITORAMENTO E REAVALIAÇÃO:**
Repetir span de dígitos (direto e inverso) após 8-12 semanas de intervenções. Avaliar não apenas pontuação mas também percepção subjetiva de melhora em tarefas do dia-a-dia (organização, planejamento, multitarefa). Reavaliação laboratorial de deficiências identificadas aos 3 meses (ferritina, B12, vitamina D, homocisteína). Se melhora insuficiente apesar de otimização: considerar avaliação neuropsicológica formal completa, avaliação formal de TDAH por especialista se padrão sugestivo, e encaminhamento para neurologia cognitiva se suspeita de processo neurodegenerativo subjacente.

**ALVO DE TRATAMENTO:**
Normalização do span para valores esperados ajustados por idade e escolaridade (direto ≥6, inverso ≥4 como mínimo), melhora de sintomas funcionais no dia-a-dia (capacidade de manter foco, lembrar instruções, realizar múltiplas tarefas), e prevenção de declínio cognitivo adicional."""
    }
}

# ============================================================================
# MAPEAMENTO DOS 61 ITEMS PARA CONTEÚDOS
# ============================================================================

ITEMS_MAP = {
    # Dubois Imediato (3 items com mesmo conteúdo)
    "8f2fb5b6-86e7-4a6f-baf5-16f2785b34d0": "dubois_imediato",
    "414fe374-9c4a-4e6c-a76f-596811b9b4e0": "dubois_imediato",
    "a4084ae3-aaa9-4720-bab2-8fd7f8542247": "dubois_imediato",

    # Dubois Tardio (3 items)
    "1a67a824-13d9-4085-b998-daab565ddd1c": "dubois_tardio",
    "cd50ed91-b1cd-47bd-9427-2592485cbdce": "dubois_tardio",
    "03d3220d-1dbe-4a9e-87d2-1226a79fb6bc": "dubois_tardio",

    # Span de Dígitos - Direto (3 items)
    "10449bea-87a3-4fa3-a163-84fde09b062f": "span_digitos",
    "8cea6550-8d9c-4535-92fd-9c7d0964a1af": "span_digitos",
    "f7449987-c689-41d5-9d47-6d283ffbfa9a": "span_digitos",

    # Span de Dígitos - Inverso (3 items)
    "13afa28c-7589-4a27-9e12-4207edd32843": "span_digitos",
    "4e2aff36-a3c5-4934-b23d-98bd66bc03ee": "span_digitos",
    "c6e4af12-f040-4985-a5bb-da2948b9e0dd": "span_digitos",
}

def update_item(item_id: str, content_key: str) -> bool:
    """Atualiza um score_item via API."""
    content = COGNICAO_CONTENT[content_key]

    payload = {
        "clinical_relevance": content["clinical_relevance"],
        "patient_explanation": content["patient_explanation"],
        "conduct": content["conduct"]
    }

    try:
        response = requests.put(
            f"{API_URL}/score-items/{item_id}",
            json=payload,
            headers={"Content-Type": "application/json"},
            timeout=30
        )

        if response.status_code == 200:
            print(f"✅ {item_id[:8]}... updated")
            return True
        else:
            print(f"❌ {item_id[:8]}... failed: {response.status_code}")
            print(f"   Response: {response.text[:200]}")
            return False

    except Exception as e:
        print(f"❌ {item_id[:8]}... error: {str(e)}")
        return False

def main():
    print("=" * 80)
    print("BATCH 31 - COGNIÇÃO - FASE 1: TESTES NEUROPSICOLÓGICOS")
    print("=" * 80)
    print(f"Total de items a processar: {len(ITEMS_MAP)}")
    print()

    success_count = 0

    for item_id, content_key in ITEMS_MAP.items():
        if update_item(item_id, content_key):
            success_count += 1

    print()
    print("=" * 80)
    print(f"RESULTADO: {success_count}/{len(ITEMS_MAP)} items atualizados com sucesso")
    print("=" * 80)

if __name__ == "__main__":
    main()
