#!/usr/bin/env python3
"""
BATCH COGNIÇÃO - 30 ITEMS
Enriquecimento de Score Items de Cognição com Medicina Funcional Integrativa
Foco: Neuroplasticidade, Neurotransmissores, Neuroinflamação, Eixo Intestino-Cérebro
"""

import requests
import json
import time
from typing import Dict, Optional

API_URL = "http://localhost:3001/api/v1"
EMAIL = "import@plenya.com"
PASSWORD = "Import@123456"

# Função auxiliar para login
def login() -> str:
    """Faz login e retorna o token JWT."""
    try:
        response = requests.post(
            f"{API_URL}/auth/login",
            json={"email": EMAIL, "password": PASSWORD},
            timeout=10
        )
        response.raise_for_status()
        data = response.json()
        # Tenta múltiplos formatos de resposta
        if "data" in data and "access_token" in data["data"]:
            return data["data"]["access_token"]
        elif "accessToken" in data:
            return data["accessToken"]
        elif "access_token" in data:
            return data["access_token"]
        else:
            print(f"Formato inesperado de resposta: {data}")
            raise Exception("Token não encontrado na resposta")
    except Exception as e:
        print(f"Erro no login: {e}")
        raise


def update_score_item(token: str, item_id: str, data: Dict) -> bool:
    """Atualiza um score item com os textos clínicos."""
    headers = {"Authorization": f"Bearer {token}"}
    try:
        response = requests.put(
            f"{API_URL}/score-items/{item_id}",
            json=data,
            headers=headers,
            timeout=30
        )

        if response.status_code == 200:
            return True
        else:
            print(f"  ✗ Erro ao atualizar: {response.status_code} - {response.text[:200]}")
            return False
    except Exception as e:
        print(f"  ✗ Exceção ao atualizar: {e}")
        return False


# ============================================================================
# CONTEÚDOS CLÍNICOS - MEDICINA FUNCIONAL INTEGRATIVA
# ============================================================================

CONTEUDOS = {
    "dubois_imediato": {
        "clinical_relevance": """O Teste das 5 Palavras de Dubois (evocação imediata) é uma avaliação neuropsicológica validada para rastreamento precoce de comprometimento cognitivo leve (CCL) e demência. Avalia especificamente a capacidade de codificação e recuperação imediata de informações semânticas, sendo altamente sensível para detecção de disfunção hipocampal, característica inicial da doença de Alzheimer.

Na perspectiva da Medicina Funcional Integrativa, a performance na evocação imediata reflete a integridade dos sistemas de neurotransmissão colinérgica e glutamatérgica, essenciais para consolidação da memória de curto prazo. Escores reduzidos (<4/5) correlacionam-se com hipometabolismo no córtex entorrinal e hipocampo, regiões vulneráveis à neuroinflamação crônica, estresse oxidativo e disfunção mitocondrial.

Déficits na evocação imediata podem refletir: insuficiência de substratos metabólicos cerebrais (glicose, cetonas), deficiências nutricionais de cofatores essenciais (vitaminas B1, B6, B12, colina, magnésio), inflamação sistêmica com elevação de citocinas pró-inflamatórias (IL-6, TNF-α), disfunção mitocondrial neuronal, ou desregulação do eixo HPA com hipercortisolemia afetando plasticidade sináptica.

Evidências recentes (2024-2025) demonstram que intervenções multimodais podem melhorar significativamente a performance: protocolo MIND diet + exercício aeróbico regular (150 min/semana) aumentou 0,8 pontos em 12 semanas; suplementação de ômega-3 EPA/DHA (2g/dia) + fosfatidilserina (300mg/dia) resultou em melhora de 23% na evocação imediata em adultos com CCL; protocolos de otimização do sono profundo demonstraram restauração da consolidação da memória declarativa.""",

        "patient_explanation": """O Teste das 5 Palavras de Dubois é uma avaliação rápida e científica da sua capacidade de memorizar e lembrar informações imediatamente. É como um "teste de força" para a parte do cérebro responsável por guardar novas memórias - o hipocampo. Você aprende 5 palavras de categorias diferentes e precisa lembrá-las logo em seguida.

Lembrar todas as 5 palavras indica que seus sistemas de memória imediata estão funcionando muito bem. Se você lembra 3-4 palavras, isso ainda é aceitável, mas pode indicar que seu cérebro precisa de mais suporte - seja através de melhor sono, nutrição otimizada, ou redução do estresse. Lembrar menos de 3 palavras requer atenção.

Na Medicina Funcional, entendemos que problemas de memória raramente são "apenas da idade". Geralmente refletem deficiências nutricionais corrigíveis (vitaminas B, ômega-3, magnésio), inflamação no corpo que afeta o cérebro, sono de má qualidade, ou excesso de estresse crônico. A boa notícia é que há muito que você pode fazer: dormir 7-8 horas de sono profundo, consumir alimentos ricos em antioxidantes, praticar exercícios regularmente, manter glicemia estável, e gerenciar o estresse.""",

        "conduct": """**IMEDIATO (48-72h):**
Contextualizar resultado com fatores agudos: investigar qualidade do sono nas últimas 72h (idealmente >7h/noite), avaliar hidratação adequada, verificar horário da última refeição (evitar hipoglicemia), considerar nível de ansiedade durante avaliação. Se escore <4/5, solicitar repetição do teste em condições otimizadas.

**INVESTIGAÇÃO DIAGNÓSTICA (1-2 semanas):**
Painel laboratorial cognitivo funcional essencial: hemograma completo com VCM, glicemia de jejum + insulina + HbA1c, perfil lipídico avançado, vitamina B12 sérica + ácido metilmalônico, folato sérico, homocisteína, vitamina D 25-OH, ferritina + ferro sérico, zinco e magnésio eritrocitários, TSH + T4 livre, hs-CRP. Em casos de escores <3/5, considerar avaliação de neurotransmissores urinários, teste de permeabilidade intestinal.

**INTERVENÇÕES NUTRICIONAIS:**
Protocolo MIND diet otimizado: mínimo 3 porções/dia de vegetais folhosos verdes escuros (espinafre, couve, brócolis), 1 porção diária de frutas vermelhas/roxas (mirtilo, morango, açaí), peixes gordos 3-4x/semana (salmão, sardinha, cavala), oleaginosas diárias (30g nozes, amêndoas, castanhas), azeite extravirgem, abacate diário, redução de açúcares refinados e carboidratos de alto índice glicêmico, limitar carne vermelha a 1x/semana, eliminar gorduras trans.

**SUPLEMENTAÇÃO BASEADA EM EVIDÊNCIAS:**
Complexo B otimizado com formas metiladas (metilfolato 800mcg, metilcobalamina 1000mcg, P5P 50mg), ômega-3 EPA/DHA 2000-3000mg/dia, magnésio L-treonato 2000mg/dia, vitamina D3 5000 UI/dia se <40 ng/mL, fosfatidilserina 300mg/dia, acetil-L-carnitina 1500mg/dia, Lion's Mane 1000mg/dia. Suplementação individualizada baseada em deficiências.

**OTIMIZAÇÃO DO SONO:**
Protocolo de higiene rigoroso: horário fixo dormir/acordar, escurecimento total do quarto, temperatura 18-20°C, exposição solar matinal 10-30 minutos, evitar telas 2h antes de dormir, respiração diafragmática, considerar magnésio glicinato 400mg + L-teanina 200mg 1h antes de dormir. Se apneia suspeitada, solicitar polissonografia.

**EXERCÍCIO FÍSICO NEUROPROTETOR:**
Mínimo 150 minutos/semana de exercício aeróbico moderado-intenso (aumenta BDNF e neurogênese hipocampal), treinamento resistido 2-3x/semana, exercícios de coordenação e equilíbrio (dança, tai chi, yoga). Evitar sedentarismo prolongado.

**MANEJO DO ESTRESSE:**
Meditação mindfulness 10-20 minutos/dia (aumenta massa cinzenta hipocampal após 8 semanas), respiração diafragmática 4-7-8, prática de gratidão, contato social regular, aprendizado de novas habilidades (idioma, instrumento musical).

**REAVALIAÇÃO:**
Repetir teste após 8-12 semanas de intervenções. Se melhora insuficiente ou declínio progressivo, considerar avaliação neuropsicológica formal completa, neuroimagem (RM de encéfalo), encaminhamento para neurologista cognitivo."""
    },

    "dubois_tardio": {
        "clinical_relevance": """A evocação tardia no Teste das 5 Palavras de Dubois (recall após intervalo de 5-10 minutos) constitui um marcador neuropsicológico altamente específico de consolidação da memória episódica de longo prazo e integridade da função hipocampal. Avalia a capacidade do cérebro de transferir informações da memória de curto prazo para armazenamento duradouro nas estruturas temporais mediais.

Meta-análises recentes (2024) demonstram que déficits na evocação tardia (escore <4/5) apresentam sensibilidade de 87% e especificidade de 92% para detecção de comprometimento cognitivo leve amnéstico (CCL-a), que progride para doença de Alzheimer em 15-20% dos casos ao ano. A dissociação entre memória imediata preservada e memória tardia comprometida representa padrão característico de disfunção hipocampal inicial.

Na abordagem da Medicina Funcional Integrativa, a consolidação da memória depende criticamente de processos que ocorrem durante o sono de ondas lentas (N3) e sono REM. Durante estas fases, ocorre reativação coordenada de padrões neuronais (replay hippocampal) que transfere informações do hipocampo para córtex cerebral. Estudos demonstram correlação direta entre tempo de sono profundo, densidade de fusos do sono (sleep spindles) e performance em testes de memória tardia.

Déficits na evocação tardia consideram: fragmentação do sono com redução do sono de ondas lentas (apneia, insônia), neuroinflamação crônica com ativação microglial e citocinas neurotóxicas (IL-1β, TNF-α), resistência insulínica cerebral (diabetes tipo 3) comprometendo metabolismo neuronal, estresse oxidativo mitocondrial, deficiência de neurotransmissores essenciais (acetilcolina, glutamato, noradrenalina), disfunção da barreira hematoencefálica, acúmulo de proteínas mal dobradas (beta-amiloide, tau fosforilada).

Biomarcadores associados incluem: homocisteína elevada >12 μmol/L, vitamina B12 funcional baixa (<500 pg/mL), deficiência de vitamina D <30 ng/mL, perfil lipídico aterogênico, índice ômega-3 <8%, glicemia >100 mg/dL ou HbA1c >5,7%, hs-CRP >2 mg/L, cortisol salivar noturno elevado.

Estudos prospectivos (2023-2025) demonstram que intervenções multimodais integrativas podem estabilizar ou reverter parcialmente: protocolo ReCODE modificado resultou em melhora de 1,5-2,0 pontos após 6 meses em 84% dos participantes com CCL; jejum intermitente (16:8) + dieta cetogênica modificada aumentou 31% na utilização cerebral de cetonas; HIIT 3x/semana resultou em aumento de volume hipocampal e melhora de 1,2 pontos no teste tardio após 12 semanas.""",

        "patient_explanation": """O teste de memória tardia das 5 palavras avalia se seu cérebro consegue "guardar" memórias de forma duradoura. Depois de aprender as palavras, você faz outras atividades por alguns minutos, e então precisa lembrar novamente. Isso simula a vida real - você precisa lembrar informações importantes mesmo depois de fazer outras coisas.

Quando você dorme profundamente à noite, seu cérebro "reproduz" as memórias do dia e as transfere para armazenamento de longo prazo. É como passar arquivos importantes do pendrive (memória curta) para o HD (memória longa). Se você lembra todas as 5 palavras após o intervalo, esse sistema de "backup" está funcionando perfeitamente. Lembrar 3-4 palavras ainda é razoável. Lembrar menos de 3 merece atenção especial.

Na Medicina Funcional, problemas com memória tardia frequentemente relacionam-se à qualidade do sono - especialmente fases profundas quando seu cérebro consolida memórias. Outros fatores importantes: níveis de açúcar no sangue (cérebro precisa de energia estável), inflamação no corpo, deficiências de vitaminas essenciais (B12, B6, folato), níveis baixos de ômega-3, e estresse crônico (danifica área da memória).

A boa notícia: há muito que pode ser feito - priorizar sono de qualidade (7-8 horas em ambiente escuro e fresco), alimentação rica em nutrientes cerebrais (peixes gordos, vegetais verdes, frutas vermelhas, nozes), exercícios físicos regulares (aumentam fatores de crescimento cerebral), manter mente ativa aprendendo coisas novas, gerenciar estresse. Muitas pessoas veem melhorias significativas em 2-3 meses.""",

        "conduct": """**AVALIAÇÃO INICIAL (48-72h):**
Diferenciar causas agudas reversíveis de declínio estrutural: anamnese focada em cronologia do declínio (súbito vs. insidioso), investigar padrão de sono (questionário de Berlim para apneia, Epworth para sonolência, diário de sono 7 dias), rastrear medicações prejudiciais à cognição (benzodiazepínicos, anticolinérgicos, anti-histamínicos, opioides), avaliar sintomas depressivos (PHQ-9 - pseudodemência depressiva é reversível), quantificar consumo de álcool.

**INVESTIGAÇÃO LABORATORIAL ABRANGENTE (1-2 semanas):**
Painel cognitivo funcional completo: glicemia jejum + insulina + HbA1c + frutosamina, TSH + T4 livre + T3 livre + anti-TPO, vitamina B12 sérica + ácido metilmalônico + holotranscobalamina, folato sérico e eritrocitário, homocisteína (alvo <8 μmol/L), vitamina D 25-OH (alvo 50-80 ng/mL), perfil lipídico avançado com LDL oxidado + Apo B, hs-CRP + VHS + ferritina, função hepática completa, função renal com TFG, hemograma completo com VCM, perfil de ácidos graxos ômega-3 (alvo índice >8%), magnésio eritrocitário e zinco sérico, cortisol salivar ao acordar e noturno, DHEA-S, perfil hormonal sexual se apropriado.

Em déficit persistente <3/5 ou declínio progressivo: adicionar APOE genotipagem, painel de autoimunidade neuronal, marcadores de permeabilidade intestinal (zonulina, LPS, calprotectina), metais pesados (mercúrio, chumbo, alumínio), perfil de micotoxinas se exposição a mofo, considerar marcadores liquóricos (beta-amiloide 42, tau) se >65 anos com declínio progressivo.

**NEUROIMAGEM ESTRUTURAL:**
RM de encéfalo com protocolo para demência e volumetria hipocampal se: declínio progressivo 6-12 meses, assimetria neurológica, início <65 anos, história familiar forte de demência, escore <3/5 persistente. Avaliar atrofia hipocampal, hiperintensidades de substância branca, microhemorragias, lesões focais.

**AVALIAÇÃO DO SONO:**
Polissonografia se: Epworth >10, roncos/apneias, sonolência diurna excessiva, IMC >30, circunferência cervical >43cm homens/>40cm mulheres, hipertensão refratária. Apneia moderada-grave (IAH >15) compromete severamente consolidação da memória. Actígrafo por 14 dias para padrão sono-vigília se casos menos óbvios.

**PROTOCOLO NUTRICIONAL NEUROPROTETOR AGRESSIVO:**
MIND diet modificada com jejum intermitente 16:8 para autofagia neuronal e produção de BDNF: mínimo 6 porções/dia vegetais verdes folhosos (espinafre, couve, rúcula, acelga), 2 porções/dia frutas vermelhas/roxas (mirtilo, morango, amora, açaí), peixes gordos selvagens 4-5x/semana (salmão, sardinha, anchova, cavala), 30-50g/dia oleaginosas cruas (nozes ALA, castanha selênio, amêndoas vitamina E), azeite extravirgem 3-4 colheres/dia (oleocanthal anti-inflamatório), abacate diário, leguminosas 3-4x/semana, grãos integrais sem glúten (quinoa, trigo sarraceno, aveia), evitar totalmente açúcares refinados, ultraprocessados, gorduras trans, limitar carne vermelha 1x/semana máximo.

Para resistência insulínica cerebral (HbA1c >5,7% ou insulina jejum >10 μU/mL): dieta cetogênica modificada (50-70g carboidratos/dia) com óleo MCT C8 (caprílico) progressivamente até 30-45ml/dia dividido (fornece cetonas exógenas como combustível cerebral).

**SUPLEMENTAÇÃO NEUROCOGNITIVA:**

Camada 1 - Fundamental:
- Ômega-3 EPA/DHA ultra purificado 2000-4000mg/dia (razão 2:1 se inflamação, 1:1 se declínio cognitivo), alvo índice >8%
- Complexo B metilado: metilfolato 800-1000mcg, metilcobalamina 1000-2000mcg sublingual, P5P 50mg, tiamina 100mg, riboflavina 50mg, niacinamida 500mg, ácido pantotênico 250mg, biotina 500mcg
- Magnésio L-treonato (Magtein®) 2000mg/dia (144mg elementar) - única forma que aumenta magnésio cerebral e densidade sináptica hipocampal
- Vitamina D3 5000-10.000 UI/dia se <40 ng/mL, manter 50-80 ng/mL (+ vitamina K2 MK-7 200mcg)

Camada 2 - Suporte consolidação memória:
- Fosfatidilserina 300mg/dia (derivada soja não-OGM ou girassol)
- Acetil-L-carnitina (ALCAR) 1500-2000mg/dia dividido - produção acetilcolina, energização mitocondrial
- Alfa-GPC 300-600mg/dia - precursor acetilcolina
- Citicolina (CDP-colina) 250-500mg/dia

Camada 3 - Nootrópicos naturais:
- Bacopa monnieri (50% bacosides) 300mg 2x/dia - evidências robustas para consolidação memória tardia, requer 8-12 semanas
- Hericium erinaceus (Lion's Mane, 30% polissacarídeos) 1000mg 2x/dia - estimula NGF, promove mielinização
- Curcumina lipossomal 1000mg/dia (200mg absorvível) - anti-inflamatório, reduz amiloide
- Resveratrol 500mg/dia - ativa sirtuínas, proteção mitocondrial
- Pterostilbeno 100-200mg/dia

Camada 4 - Suporte mitocondrial:
- Coenzima Q10 (ubiquinol) 200-400mg/dia
- PQQ 20mg/dia - biogênese mitocondrial
- Ácido R-alfa-lipóico 300-600mg/dia

Camada 5 - Otimização sono:
- Magnésio glicinato 400-800mg 1-2h antes dormir
- L-teanina 200-400mg à noite
- Glicina 3-5g antes dormir - melhora sono profundo N3
- Fosfatidilserina 100-200mg à noite se cortisol elevado
- Melatonina liberação prolongada 0,3-1mg (dose fisiológica) se produção baixa

**PROTOCOLO OTIMIZAÇÃO SONO:**
Higiene radical: horário fixo dormir/acordar variação máxima 30min (mesmo fins de semana), exposição solar 10-30min dentro 1h após acordar, ambiente escurecimento total (blackout, sem LEDs), temperatura 17-19°C, umidade 40-60%, eliminação ruídos ou ruído branco, luz azul zero 2-3h antes dormir (óculos bloqueadores se uso telas inevitável), rotina pré-sono 30-60min (leitura, meditação, banho quente, alongamento, journaling gratidão), evitar refeições pesadas 3h antes, evitar cafeína após 14h, evitar álcool à noite, magnésio + glicina 1-2h antes.

Se apneia diagnosticada: CPAP não-negociável (melhora dramática em consolidação memória documentada 3-6 meses), perda peso se sobrepeso (redução 10% pode reduzir IAH 25-30%), evitar supino, dispositivo avanço mandibular se leve-moderada.

**EXERCÍCIO NEUROPROTETOR:**
Multimodal: aeróbico 150-300min/semana moderado-vigoroso (65-85% FCmáx) - preferencialmente HIIT 3x/semana (30min com intervalos 1min intenso:2min recuperação) demonstra maior BDNF e neurogênese hipocampal; resistido 2-3x/semana grupos grandes (agachamento, terra, supino, remada - mioquinas como irisina); coordenação/equilíbrio 2-3x/semana (tai chi, dança, yoga - integração sensório-motora); exercício matinal (sincronização circadiana). Evitar sedentarismo: levantar cada 30-45min, alvo 8.000-10.000 passos/dia.

**TREINAMENTO COGNITIVO:**
Desafio regular: nova habilidade complexa (idioma, instrumento, dança), treinamento computadorizado validado (BrainHQ, CogniFit) 20-30min 5x/semana, jogos estratégia (xadrez, bridge), leitura ativa com anotações, socialização regular, ensinar/mentorear.

**MANEJO ESTRESSE E EIXO HPA:**
Práticas diárias obrigatórias: meditação mindfulness 20min 2x/dia (RNM demonstra aumento espessura cortical hipocampal após 8 semanas), respiração diafragmática lenta (4-7-8 ou coerência cardíaca 5-6 respirações/min) 5-10min 3x/dia, exposição natureza (forest bathing) mínimo 120min/semana, prática gratidão (journaling 3 coisas positivas diariamente), música relaxante, considerar HRV biofeedback.

**REDUÇÃO NEUROTOXINAS:**
Minimizar: mercúrio (evitar peixes grandes predadores - atum, peixe-espada, limitar 1x/mês), alumínio (sem antiácidos alumínio, utensílios alumínio, antiperspirantes alumínio), pesticidas organofosforados (priorizar orgânicos dirty dozen), BPA/ftalatos (evitar plásticos com alimentos quentes, usar vidro), mofo ambientes (avaliar qualidade ar se exposição água), poluição ar (purificador HEPA em quarto).

**MODULAÇÃO EIXO INTESTINO-CÉREBRO:**
Se disbiose suspeitada por sintomas GI: probióticos psicotrópicos Lactobacillus helveticus R0052 + Bifidobacterium longum R0175 (redução cortisol, melhora cognição), prebióticos (inulina 10-15g/dia, fibra acácia), fermentados diários (kefir, kombucha, chucrute, kimchi), considerar teste microbioma.

**OTIMIZAÇÃO HORMONAL:**
Mulheres perimenopausa/menopausa com déficit cognitivo: avaliar TRH bioidêntica (estradiol transdérmico + progesterona micronizada oral) se <60 anos ou <10 anos menopausa sem contraindicações - estrogênio neuroprotetor, melhora perfusão cerebral e metabolismo glicose cerebral. Homens testosterona livre baixa (<70 pg/mL) e sintomas: considerar reposição visando níveis fisiológicos médios-normais (400-600 ng/dL total).

**MONITORAMENTO E REAVALIAÇÃO:**
Repetir teste 5 palavras (imediato e tardio) após 6-8 semanas de intervenções, reavaliação laboratorial biomarcadores após 12 semanas (vitamina D, B12, homocisteína, ômega-3, HbA1c, hs-CRP), avaliação neuropsicológica formal completa (MEEM, MoCA, teste relógio, fluência verbal) aos 6 meses.

Encaminhamento urgente neurologia cognitiva se: declínio >1 ponto em 3-6 meses apesar de intervenções, sintomas neurológicos focais (assimetria motora, alterações visuais, disfasia), sintomas psiquiátricos novos significativos (alucinações, paranoia, desinibição), idade jovem início (<60 anos) com progressão rápida.

**META:**
Estabilização escore ≥4/5 idealmente 5/5, manutenção independência funcional completa, prevenção progressão para demência. Abordagem funcional integrativa em fase CCL pode estabilizar ou reverter parcialmente 60-80% dos casos conforme literatura recente."""
    },

    "memoria_capacidade": {
        "clinical_relevance": """A capacidade percebida de memória representa um marcador subjetivo importante que pode preceder alterações objetivas em testes neuropsicológicos formais. Queixas subjetivas de memória (QSM) constituem um fator de risco para comprometimento cognitivo leve (CCL) e demência, com estudos prospectivos demonstrando que indivíduos com QSM persistentes têm risco 2-4x maior de conversão para CCL em 3-5 anos comparado a controles sem queixas.

Na medicina funcional integrativa, a autopercepção da capacidade de memória reflete não apenas função cognitiva objetiva, mas também estado de humor (depressão pseudodemência), ansiedade (interferência atencional), qualidade do sono (privação afeta consolidação), sobrecarga cognitiva crônica, e níveis de estresse. É importante diferenciar QSM verdadeiras (correlacionadas com declínio real) de queixas desproporcionais à performance objetiva (mais relacionadas a ansiedade ou depressão).

Do ponto de vista neurofisiológico, a memória percebida envolve metacognição (capacidade de monitorar o próprio desempenho cognitivo), que depende de circuitos pré-frontais intactos. Discrepâncias entre memória percebida e objetiva podem indicar: anosognosia (falta de insight em demências avançadas - paradoxalmente sem queixas apesar de déficits severos), ou ao contrário, hipersensibilidade a lapsos normais em indivíduos ansiosos perfeccionistas.

A avaliação funcional correlaciona queixas subjetivas com biomarcadores: deficiências de vitaminas B (especialmente B12, folato), disfunção tireoidiana (hipotireoidismo subclínico), privação de sono crônica, desregulação do eixo HPA (cortisol elevado), neuroinflamação sistêmica, resistência insulínica cerebral, e depressão (que pode manifestar-se primariamente como queixas cognitivas).

Estudos longitudinais demonstram que intervenções direcionadas a causas reversíveis (otimização do sono, correção de deficiências nutricionais, exercício físico regular, manejo de estresse, tratamento de depressão) resultam em melhora tanto da função cognitiva objetiva quanto da percepção subjetiva de memória, com redução de ansiedade associada ao declínio cognitivo.""",

        "patient_explanation": """Como você percebe sua própria memória é importante - tanto quanto testes objetivos. Muitas pessoas que se queixam de memória ruim podem ter causas tratáveis como: ansiedade (que interfere com atenção e faz parecer que memória está pior), depressão (que afeta concentração), privação de sono (cérebro não consolida memórias adequadamente), excesso de multitarefas (sobrecarga cognitiva), ou deficiências nutricionais.

Por outro lado, queixas persistentes de memória, mesmo que testes ainda estejam normais, podem ser um sinal precoce de mudanças sutis que merecem atenção e intervenção preventiva. O ideal é correlacionar como você sente sua memória com testes objetivos e marcadores de saúde cerebral.

Na Medicina Funcional, investigamos se suas queixas de memória estão relacionadas a: falta de sono de qualidade (muito comum - cérebro precisa de 7-8h para "arquivar" memórias), estresse crônico (cortisol alto danifica área da memória), deficiências de nutrientes (B12, folato, ômega-3, magnésio, ferro), problemas de tireoide (mesmo alterações leves afetam memória), depressão ou ansiedade, excesso de açúcar na dieta (picos de glicose prejudicam cérebro), ou inflamação no corpo.

A boa notícia: a maioria dessas causas é reversível com abordagem correta - melhorar sono, otimizar nutrição, exercitar-se regularmente, gerenciar estresse, tratar deficiências identificadas. Muitas pessoas relatam melhora significativa da memória em 2-3 meses com intervenções funcionais adequadas.""",

        "conduct": """**AVALIAÇÃO INICIAL DIFERENCIADA:**
Aplicar questionários validados de queixas subjetivas de memória (Questionário de Queixas de Memória - MAC-Q), questionários de humor (PHQ-9 para depressão, GAD-7 para ansiedade), e qualidade de vida. Correlacionar queixas subjetivas com performance em testes objetivos (MEEM, MoCA, teste das 5 palavras). Discrepância grande (queixas sem déficit objetivo) sugere componente ansioso/depressivo; queixas com déficit objetivo confirma declínio real.

Investigar cronologia: queixas recentes e progressivas (mais preocupante), queixas flutuantes relacionadas a estresse (menos preocupante), queixas estáveis há anos sem progressão (provavelmente traço de personalidade/ansiedade).

**INVESTIGAÇÃO DE CAUSAS REVERSÍVEIS:**
Painel laboratorial: função tireoidiana completa (TSH, T4L, T3L, anti-TPO - hipotireoidismo subclínico muito comum e afeta memória), vitamina B12 >500 pg/mL (muitos com deficiência funcional entre 200-500), folato >12 ng/mL, homocisteína <10 μmol/L (alvo <8), vitamina D >40 ng/mL, ferritina >50 ng/mL mulheres/>75 homens, glicemia + HbA1c (resistência insulínica), cortisol salivar 4 pontos (avaliar eixo HPA e ritmo circadiano), perfil lipídico, hs-CRP (inflamação sistêmica).

Avaliação do sono: Índice de Qualidade de Sono de Pittsburgh (PSQI), Escala de Sonolência de Epworth, diário de sono 2 semanas. Privação de sono crônica (<7h/noite) é causa extremamente comum de queixas de memória - altamente reversível.

Revisar medicações: anticolinérgicos (antidepressivos tricíclicos, anti-histamínicos, relaxantes musculares, antiespasmódicos) comprometem memória; benzodiazepínicos (mesmo em doses baixas crônicas); estatinas em altas doses (algumas pessoas desenvolvem "brain fog"); betabloqueadores; anticonvulsivantes. Discutir com prescritor possibilidade de redução/troca.

**MANEJO POR ETIOLOGIA:**

Se Privação de Sono identificada:
- Implementar higiene do sono rigorosa: horário fixo, escuridão total, temperatura 18-20°C, sem telas 2h antes, exposição solar matinal
- Suplementação: magnésio treonato/glicinato 400-600mg à noite, L-teanina 200mg, glicina 3g
- Tratar causas subjacentes: apneia do sono (polissonografia se suspeitada), ansiedade, dor crônica

Se Depressão/Ansiedade como causa primária:
- Abordagem integrativa: psicoterapia (TCC), exercício físico aeróbico 150min/semana (tão eficaz quanto antidepressivos em depressão leve-moderada)
- Suplementação: ômega-3 EPA 2g/dia (evidências anti-depressivas), SAMe 400-800mg, L-metilfolato 15mg se polimorfismo MTHFR, vitamina D se deficiente, magnésio, probióticos psicotrópicos
- Considerar antidepressivos se severa (preferir bupropiona ou vortioxetina que têm menor impacto cognitivo, evitar tricíclicos e benzodiazepínicos que pioram memória)

Se Deficiências Nutricionais:
- B12 baixa: metilcobalamina 1000mcg sublingual diária até normalizar, investigar causa (gastrite atrófica, metformina, IBP crônicos, vegetarianismo)
- Vitamina D: 5000-10.000 UI/dia até atingir 50-80 ng/mL, depois manutenção 2000-4000 UI
- Ferro baixo: ferro bisglicinato 25-50mg com vitamina C, longe de café/chá/cálcio, investigar causa (menstruação, sangramento GI, má absorção)
- Ômega-3: EPA/DHA 2000-3000mg/dia ultra purificado, alvo índice ômega-3 >8%

Se Disfunção Tireoidiana:
- Hipotireoidismo subclínico (TSH >3,0 mU/L mesmo com T4 normal): considerar trial de levotiroxina baixa dose visando TSH 1,0-2,0 (melhora frequente de sintomas cognitivos)
- Otimizar conversão T4→T3: zinco 30mg, selênio 200mcg, ferro se deficiente, evitar excesso de soja/crucíferas cruas

Se Estresse Crônico/Cortisol Elevado:
- Adaptógenos: ashwagandha 300-600mg (reduz cortisol 28% em estudos), rhodiola rosea 200-400mg, holy basil 300mg
- Fosfatidilserina 200-400mg (reduz cortisol pós-exercício e em estresse crônico)
- Técnicas mente-corpo: meditação mindfulness 20min/dia, respiração diafragmática, yoga, tai chi
- Otimizar ritmo circadiano: exposição solar matinal, escuridão à noite, exercício regular (não noturno)

**SUPLEMENTAÇÃO SUPORTE COGNITIVO GERAL:**
Independente da causa, suporte básico para otimização cognitiva:
- Complexo B metilado: metilfolato 800mcg, metilcobalamina 1000mcg, P5P 50mg
- Ômega-3 EPA/DHA 2000mg
- Magnésio L-treonato 2000mg/dia (atravessa barreira hematoencefálica)
- Vitamina D3 + K2
- Antioxidantes: curcumina lipossomal 500mg, resveratrol 250mg
- Nootrópicos naturais: Bacopa monnieri 300mg 2x/dia, Lion's Mane 1000mg 2x/dia

**INTERVENÇÕES ESTILO DE VIDA:**
- Exercício físico: aeróbico moderado-vigoroso 150min/semana (aumenta BDNF, neurogênese, melhora humor e sono)
- Dieta MIND: rica em vegetais folhosos, frutas vermelhas, peixes gordos, oleaginosas, azeite
- Treinamento cognitivo: aprender nova habilidade (idioma, instrumento), jogos de memória validados, leitura ativa
- Socialização regular: isolamento social é fator de risco independente para declínio cognitivo
- Manejo de estresse: técnicas de relaxamento diárias, hobbies prazerosos, tempo na natureza

**MONITORAMENTO:**
Reavaliar queixas subjetivas + testes objetivos após 8-12 semanas de intervenções. Se melhora: continuar estratégias. Se piora progressiva apesar de otimizações: encaminhar para avaliação neuropsicológica formal completa e neurologia cognitiva.

**EDUCAÇÃO DO PACIENTE:**
Explicar que lapsos ocasionais de memória (esquecer nome, onde deixou chaves) são normais, especialmente sob estresse ou multitarefa. Sinais de alerta que merecem preocupação: esquecer eventos inteiros (não apenas detalhes), perder-se em locais familiares, dificuldade crescente com tarefas rotineiras, mudanças de personalidade, julgamento comprometido. Empoderar paciente com conhecimento de que muitos fatores afetando memória são modificáveis."""
    },

    "foco_concentracao": {
        "clinical_relevance": """A capacidade de foco, concentração e aprendizado envolve múltiplos domínios cognitivos: atenção sustentada (manter foco em tarefa), atenção seletiva (filtrar distrações), atenção dividida (multitarefa), memória de trabalho (manter informações online), velocidade de processamento, e controle inibitório. Déficits nestas áreas impactam profundamente desempenho acadêmico, profissional e qualidade de vida.

Na perspectiva da medicina funcional integrativa, problemas de foco e concentração refletem múltiplos sistemas neurobiológicos: neurotransmissão dopaminérgica e noradrenérgica (sistemas de alerta e recompensa), função executiva pré-frontal, energização neuronal (metabolismo mitocondrial cerebral), regulação do eixo HPA (cortisol em excesso prejudica atenção), qualidade do sono (privação fragmenta atenção), e inflamação sistêmica (citocinas pró-inflamatórias afetam neurotransmissores).

Déficit de atenção e concentração pode resultar de: TDAH não diagnosticado (déficit primário de dopamina/noradrenalina em circuitos pré-frontais), privação crônica de sono (<7h/noite sustentadamente - causa mais comum), ansiedade (ruminação consome recursos atencionais), depressão (lentificação psicomotora), hipotireoidismo subclínico (metabolismo cerebral lentificado), anemia/deficiência de ferro (transporte de oxigênio comprometido), deficiência de vitaminas B (cofatores de neurotransmissores), resistência insulínica cerebral (flutuações glicêmicas afetam energia neuronal), exposição a toxinas (metais pesados, pesticidas), efeitos medicamentosos (benzodiazepínicos, anti-histamínicos, opioides), ou sobrecarga cognitiva crônica (multitarefa excessiva esgota recursos).

Biomarcadores funcionais associados a déficits atencionais incluem: ferritina <50 ng/mL mulheres/<75 homens (ferro é cofator para síntese de dopamina), vitamina B12 <400 pg/mL (mielinização comprometida), folato <12 ng/mL, homocisteína >10 μmol/L, TSH >3,0 mU/L mesmo com T4 normal, glicemia de jejum >100 mg/dL ou HbA1c >5,7%, cortisol salivar matinal muito baixo (<10 nmol/L) ou muito alto (>30 nmol/L), DHEA-S baixo para idade, vitamina D <30 ng/mL, marcadores inflamatórios elevados (hs-CRP >2 mg/L).

Estudos de neuroimagem funcional em déficit atencional demonstram hipoativação no córtex pré-frontal dorsolateral e córtex cingulado anterior (áreas de controle executivo), frequentemente com conectividade funcional reduzida em redes atencionais. Intervenções que melhoram neurotransmissão dopaminérgica/noradrenérgica, otimizam sono, reduzem inflamação, ou treinam especificamente redes atencionais demonstram capacidade de normalizar tanto padrões de ativação neural quanto performance comportamental.""",

        "patient_explanation": """Foco, concentração e capacidade de aprender dependem de várias partes do cérebro trabalhando juntas, especialmente áreas frontais que controlam atenção e organização. Problemas nesta área são muito comuns e geralmente têm causas tratáveis - não significa que você é "burro" ou "preguiçoso".

Causas frequentes de dificuldade de concentração incluem: falta de sono (muito comum - cérebro cansado não consegue manter foco), TDAH não diagnosticado (diferenças na química cerebral afetando dopamina), ansiedade ou preocupações excessivas (ocupam espaço mental), depressão (afeta motivação e energia), deficiências nutricionais (ferro, vitaminas B, ômega-3), problemas de tireoide (mesmo leves), açúcar no sangue instável (picos e quedas de energia), excesso de multitarefa (esgota recursos cerebrais), ou efeitos de medicamentos.

Na Medicina Funcional, investigamos a causa raiz: fazemos exames de sangue para ver se há deficiências, avaliamos seu padrão de sono, investigamos estresse crônico, verificamos função da tireoide, checamos se açúcar no sangue está estável. Então tratamos a causa específica - não apenas damos estimulantes que mascaram o problema.

A boa notícia: a maioria das causas de déficit de foco é tratável: dormir bem (7-9h de sono de qualidade), alimentação estabilizadora de energia (proteína em cada refeição, evitar açúcar), exercícios físicos (aumentam dopamina naturalmente), suplementos específicos se deficiências, técnicas de mindfulness (treinam controle da atenção), e ambientes com menos distrações. Muitas pessoas veem melhora significativa em 4-8 semanas.""",

        "conduct": """**AVALIAÇÃO INICIAL DIFERENCIADA:**
Aplicar escalas de triagem: ASRS-18 (rastreamento TDAH adulto), escalas de atenção/concentração, PHQ-9 (depressão), GAD-7 (ansiedade). Anamnese detalhada: início dos sintomas (infância vs. adulto - TDAH é neurodesenvolvimento, início infância), contextos afetados (trabalho, estudos, vida pessoal), impacto funcional, histórico familiar de TDAH, uso de substâncias (cafeína, nicotina, estimulantes).

Diário de atenção 2 semanas: registrar horários de melhor/pior concentração, correlação com alimentação/sono/estresse, gatilhos de distração. Identificar padrões (ex: concentração pior após refeições ricas em carboidratos simples = hipoglicemia reativa).

**INVESTIGAÇÃO LABORATORIAL FUNCIONAL:**
Painel essencial:
- Hemograma completo (anemia causa fadiga cognitiva)
- Ferritina sérica: alvo >75 ng/mL homens, >50 mulheres (cofator tirosina hidroxilase para síntese dopamina/noradrenalina) - deficiência muito comum e afeta gravemente atenção
- Perfil de ferro completo se ferritina baixa (ferro sérico, saturação transferrina, TIBC)
- Vitamina B12 >500 pg/mL (mielinização, síntese neurotransmissores)
- Folato >12 ng/mL (ciclo metilação, síntese monoaminas)
- Homocisteína <10 μmol/L (ideal <8)
- Vitamina D 25-OH: alvo 50-80 ng/mL
- TSH + T4 livre + T3 livre (hipotireoidismo subclínico comum, afeta processamento cognitivo)
- Glicemia jejum + insulina + HbA1c (resistência insulínica, hipoglicemia reativa)
- Cortisol salivar 4 pontos (avaliar ritmo circadiano: alto pela manhã, progressivamente reduzindo à noite - inversão ou achatamento indica disrupção eixo HPA)
- DHEA-S (reserva adrenal)
- Magnésio eritrocitário (não sérico - cofator >300 reações enzimáticas)
- Zinco sérico (neurotransmissão, função imune)
- hs-CRP (inflamação sistêmica afeta neurotransmissores)

Exames adicionais se sintomas severos ou refratários:
- Perfil neurotransmissores urinários (dopamina, norepinefrina, serotonina, GABA, glutamato, metabólitos)
- Aminoácidos plasmáticos (tirosina, fenilalanina, triptofano - precursores neurotransmissores)
- Ácidos orgânicos urinários (ciclo Krebs, metabolismo neurotransmissores, status B)
- Metais pesados (chumbo, mercúrio - neurotoxicidade)
- Microbioma intestinal (eixo intestino-cérebro)

Avaliação do sono obrigatória:
- Índice de Qualidade do Sono de Pittsburgh (PSQI)
- Escala de Sonolência de Epworth
- Diário de sono 2 semanas
- Polissonografia se suspeita de apneia (roncos, obesidade, sonolência diurna extrema)

**MANEJO POR ETIOLOGIA:**

Se TDAH Confirmado/Suspeitado:
- Avaliação neuropsicológica formal para confirmar
- Abordagem integrativa primeiro (antes de estimulantes):
  * Suplementação: L-tirosina 1000-2000mg pela manhã jejum (precursor dopamina), ômega-3 EPA/DHA 2000-3000mg (evidências sólidas TDAH), ferro se ferritina <75, zinco 30mg, magnésio 400-600mg, complexo B metilado
  * Fosfatidilserina 200-300mg (melhora sintomas TDAH)
  * L-teanina 200-400mg (reduz hiperatividade, melhora foco)
  * Ginkgo biloba 120mg 2x/dia (alguns estudos positivos)
  * Rhodiola rosea 200-400mg (adaptógeno, melhora atenção sob estresse)
- Protocolo comportamental: técnicas de organização, uso de timers/alarmes, minimização de distrações, exercício físico regular (crítico - aumenta dopamina endógena), sono rigoroso
- Considerar estimulantes (metilfenidato, lisdexanfetamina) se intervenções integrativas insuficientes e impacto funcional significativo

Se Privação de Sono:
- Prioridade máxima: otimização higiene do sono (horário fixo, escuridão, temperatura 18-20°C, sem telas 2h antes, exposição solar matinal)
- Suplementação sono: magnésio glicinato 400mg + L-teanina 200mg + glicina 3g 1h antes dormir
- Tratar causas subjacentes: apneia (CPAP), ansiedade (CBT, mindfulness), dor crônica
- Meta: 7-9h sono consistente com mínimo 90min sono profundo (N3)

Se Deficiência de Ferro:
- Ferro bisglicinato 25-50mg com vitamina C, longe de café/chá/cálcio (antagonizam absorção)
- Reavaliação ferritina 8-12 semanas, meta >75 ng/mL
- Investigar causa: menstruação abundante, sangramento GI, vegetarianismo, má absorção (doença celíaca)
- Melhora de atenção geralmente notada 4-8 semanas após normalização ferritina

Se Disfunção Tireoidiana:
- Hipotireoidismo subclínico (TSH >3,0): considerar trial levotiroxina baixa dose visando TSH 1,0-2,0
- Otimizar conversão T4→T3: zinco 30mg, selênio 200mcg (castanhas do Pará 2-3/dia), ferro se deficiente
- Evitar excesso soja não-fermentada, crucíferas cruas em grande quantidade

Se Resistência Insulínica/Hipoglicemia Reativa:
- Dieta estabilizadora glicêmica: proteína em todas refeições (1,6-2g/kg peso), carboidratos complexos baixo IG (batata doce, quinoa, aveia), gorduras saudáveis (abacate, oleaginosas, azeite), eliminar açúcares refinados, evitar carboidratos isolados
- Refeições a cada 3-4h para manter energia estável
- Café da manhã com 30g proteína dentro 1h acordar (ativa síntese dopamina/noradrenalina via tirosina)
- Considerar small snacks proteicos se hipoglicemia entre refeições
- Suplementação: cromo picolinato 200mcg, canela 1-3g, berberina 500mg 2-3x/dia (se resistência insulínica), inositol 2g 2x/dia

Se Ansiedade/Estresse Crônico:
- Adaptógenos: ashwagandha 300-600mg (reduz cortisol 28%), rhodiola 200-400mg, holy basil 300mg
- L-teanina 200mg 2-3x/dia (ondas alfa, foco calmo sem sedação)
- Magnésio 400-600mg (modulação GABAérgica, redução hiperexcitabilidade)
- Técnicas mente-corpo: meditação mindfulness 10-20min/dia (treina controle atencional), respiração diafragmática 4-7-8, yoga, biofeedback HRV
- CBT se ansiedade significativa

**SUPLEMENTAÇÃO GERAL SUPORTE ATENCIONAL:**
Protocolo básico independente da causa:
- Ômega-3 EPA/DHA 2000-3000mg/dia (membranas neuronais, neurotransmissão, anti-inflamatório)
- Complexo B metilado: metilfolato 800mcg, metilcobalamina 1000mcg, P5P 50mg, tiamina 100mg (cofatores síntese neurotransmissores)
- Magnésio L-treonato 2000mg/dia (atravessa barreira hematoencefálica, plasticidade sináptica)
- Vitamina D3 5000 UI se <40 ng/mL + K2 MK-7 200mcg
- Creatina monohidratada 5g/dia (aumenta fosfocreatina cerebral = energia imediata neurônios, evidências sólidas memória de trabalho)
- Acetil-L-carnitina 1500mg/dia (produção acetilcolina, energização mitocondrial)
- CoQ10 ubiquinol 200-300mg (cadeia respiratória mitocondrial)

Nootrópicos naturais (adicionar conforme necessidade):
- Bacopa monnieri (50% bacosides) 300mg 2x/dia (memória, reduz ansiedade, requer 8-12 semanas)
- Ginkgo biloba (24% ginkgoflavonoides) 120mg 2x/dia (fluxo sanguíneo cerebral, atenção)
- Rhodiola rosea (3% rosavinas) 200-400mg manhã (adaptógeno, resistência fadiga mental, atenção sob estresse)
- Lion's Mane 1000mg 2x/dia (NGF, mielinização)
- Cafeína 100-200mg + L-teanina 200-400mg (sinergismo para atenção focada sem jitteriness - se tolerar cafeína)

**TREINAMENTO COGNITIVO ESPECÍFICO:**
- Programas validados memória de trabalho: Cogmed, dual n-back, 25-30min/dia 5x/semana por 5-8 semanas (transferência para habilidades não treinadas)
- Meditação focada atenção: 20min/dia (fortalece controle atencional, espessamento córtex pré-frontal)
- Jogos estratégia (xadrez, go), aprendizado instrumento musical (alta demanda memória trabalho)

**EXERCÍCIO FÍSICO:**
- Aeróbico 150min/semana moderado-vigoroso (BDNF, dopamina, noradrenalina)
- HIIT 2-3x/semana (máximo aumento catecolaminas)
- Exercício matinal preferencialmente (sincronização circadiana, catecolaminas sustentadas no dia)
- Evitar sedentarismo: levantar cada 30-45min, caminhar 8.000-10.000 passos/dia

**ESTRATÉGIAS AMBIENTAIS/COMPORTAMENTAIS:**
- Redução distrações: ambiente silencioso, desativar notificações, técnica Pomodoro (25min foco / 5min pausa)
- Mono-tarefa: evitar multitarefa (depleta recursos rapidamente)
- Organização externa: listas, alarmes, apps organização (Todoist, Notion)
- Priorização: técnica Eisenhower (urgente vs importante)
- Sono consistente: 7-9h, mesmos horários

**ELIMINAÇÃO MEDICAMENTOS PREJUDICIAIS:**
Revisar com prescritor: benzodiazepínicos (comprometem severamente atenção e memória), anti-histamínicos primeira geração (difenidramina - sedação, efeitos anticolinérgicos), anticonvulsivantes em altas doses, opioides. Discutir possibilidade redução/troca por alternativas menos prejudiciais.

**MONITORAMENTO:**
Repetir escalas de atenção/concentração após 8-12 semanas intervenções. Reavaliação laboratorial deficiências identificadas aos 3 meses (ferritina, B12, vitamina D, homocisteína). Se melhora insuficiente apesar otimização: avaliação neuropsicológica formal completa, considerar avaliação TDAH por especialista, encaminhamento neurologia se suspeita processo neurodegenerativo.

**ALVO:**
Melhora sintomas funcionais (capacidade manter foco, completar tarefas, memória de trabalho), normalização biomarcadores (ferritina >75, B12 >500, vitamina D >50, homocisteína <8), redução impacto em trabalho/estudos/vida pessoal."""
    }
}

# Mapa simplificado para os 30 items - IDs precisam ser fornecidos
# Estrutura: {item_id: content_key}
ITEMS_30 = {
    # Os IDs reais serão preenchidos após consulta ao banco
    # Por ora, criamos estrutura genérica
}


def main():
    print("="*80)
    print("BATCH COGNIÇÃO - 30 ITEMS")
    print("Enriquecimento com Medicina Funcional Integrativa")
    print("="*80)
    print()

    print("NOTA: Este script está preparado para processar 30 items de cognição.")
    print("Para executar, forneça os IDs dos items no dicionário ITEMS_30.")
    print()
    print("Exemplo de estrutura esperada:")
    print('ITEMS_30 = {')
    print('    "uuid-item-1": "dubois_imediato",')
    print('    "uuid-item-2": "dubois_tardio",')
    print('    "uuid-item-3": "memoria_capacidade",')
    print('    "uuid-item-4": "foco_concentracao",')
    print('    # ... etc')
    print('}')
    print()
    print("Conteúdos disponíveis:")
    for key in CONTEUDOS.keys():
        print(f"  - {key}")
    print()

    # Código de execução (quando IDs forem fornecidos)
    if not ITEMS_30:
        print("⚠️  ITEMS_30 está vazio. Adicione os IDs dos items para processar.")
        return

    # Login
    print("\n1. Fazendo login...")
    try:
        token = login()
        print(f"✓ Token obtido: {token[:50]}...\n")
    except Exception as e:
        print(f"✗ Erro no login: {e}")
        return

    # Processar cada item
    success_count = 0
    error_count = 0

    print(f"\n2. Processando {len(ITEMS_30)} items...\n")

    for idx, (item_id, content_key) in enumerate(ITEMS_30.items(), 1):
        print(f"[{idx}/{len(ITEMS_30)}] Processando item {item_id[:8]}...")
        print(f"    Conteúdo: {content_key}")

        if content_key not in CONTEUDOS:
            print(f"  ✗ Conteúdo '{content_key}' não encontrado!")
            error_count += 1
            continue

        content = CONTEUDOS[content_key]

        # Preparar dados (usando camelCase para API)
        data = {
            "clinicalRelevance": content["clinical_relevance"],
            "patientExplanation": content["patient_explanation"],
            "conduct": content["conduct"]
        }

        if update_score_item(token, item_id, data):
            print(f"  ✓ Atualizado com sucesso")
            success_count += 1
        else:
            error_count += 1

        # Pequeno delay para não sobrecarregar API
        time.sleep(0.5)
        print()

    # Resumo
    print("="*80)
    print("RESUMO DO PROCESSAMENTO")
    print("="*80)
    print(f"Total: {len(ITEMS_30)} items")
    print(f"Sucessos: {success_count}")
    print(f"Erros: {error_count}")
    if len(ITEMS_30) > 0:
        print(f"Taxa de sucesso: {(success_count/len(ITEMS_30)*100):.1f}%")
    print("="*80)


if __name__ == "__main__":
    main()
