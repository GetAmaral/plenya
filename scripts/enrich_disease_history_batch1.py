#!/usr/bin/env python3
"""
Script para enriquecer 20 Score Items do grupo "Histórico de doenças"
Batch 1: Primeiros 20 items
"""

import requests
import json
import time
from typing import Dict, List, Optional

# Configuração
API_BASE_URL = "http://localhost:3001/api/v1"
EMAIL = "import@plenya.com"
PASSWORD = "Import@123456"

# IDs dos 20 items (obtidos da consulta SQL)
ITEMS = [
    {"id": "ea288642-425c-496c-a577-6f6ee3acb468", "name": "Hipertensão arterial"},
    {"id": "8236ab3f-1ee1-4f17-a053-7f45b7162dd2", "name": "Diabetes mellitus"},
    {"id": "3e876264-ae80-4c5c-bc5b-00fbf09daf98", "name": "Pré-diabetes / resistência a insulina"},
    {"id": "c79cf41f-f3fe-41c2-b7af-1e9880f62583", "name": "DM estabelecido"},
    {"id": "126d82d2-e32d-4f47-bca1-5035e2abc8d2", "name": "Obesidade"},
    {"id": "58235b33-5431-48d9-83f4-77d158caf4da", "name": "Dislipidemia"},
    {"id": "9f646670-c0eb-40d5-bcb0-b2f069e41a29", "name": "Câncer"},
    {"id": "fb8c9e63-8d36-4e03-bb24-dbdc747e5fd4", "name": "Insuficiência cardíaca"},
    {"id": "fc4ea1da-ad8c-419c-be76-c58405f69b11", "name": "Arritmia"},
    {"id": "d4d8283f-42c4-4de7-a98d-603c41d12ae6", "name": "Doença cardiovascular (IAM, revascularização, AVC, etc)"},
    {"id": "82ba4b8f-708a-4adb-8bd5-a242485e0446", "name": "Doença renal crônica"},
    {"id": "3a1e71df-975f-4a7d-815c-a56b3e3735c9", "name": "Outras doenças renais"},
    {"id": "34fbbed1-e8a2-4103-bda8-a522bc357978", "name": "Nefrite"},
    {"id": "addbf7df-76b6-4bdb-b7c8-1c239c373389", "name": "Nefrótica"},
    {"id": "6ba04946-25de-49c2-8fc0-7dcd8dcb75bc", "name": "Litíase"},
    {"id": "a1a681ce-27a4-434a-aeae-0682b3375758", "name": "ITU"},
    {"id": "13bff813-f2dc-4d1f-b3bd-8f8e8f731634", "name": "Doenças virais crônicas"},
    {"id": "3b18f7b7-f600-44d0-837c-1ad343661dc2", "name": "HIV"},
    {"id": "8014a88c-4519-4e76-b57a-c79c8e45a162", "name": "Hepatite B"},
    {"id": "9ff6f838-bc35-4039-8b37-f70fdec68733", "name": "Hepatite C"}
]

class ScoreItemEnricher:
    def __init__(self):
        self.token = None
        self.session = requests.Session()

    def login(self) -> bool:
        """Faz login e obtém token JWT"""
        try:
            response = self.session.post(
                f"{API_BASE_URL}/auth/login",
                json={"email": EMAIL, "password": PASSWORD}
            )
            response.raise_for_status()
            data = response.json()
            self.token = data["accessToken"]
            self.session.headers.update({"Authorization": f"Bearer {self.token}"})
            print("Login realizado com sucesso!")
            return True
        except Exception as e:
            print(f"Erro no login: {e}")
            return False

    def update_score_item(self, item_id: str, clinical_relevance: str,
                         patient_explanation: str, conduct: str) -> bool:
        """Atualiza um score item com os textos clínicos"""
        try:
            payload = {
                "clinicalRelevance": clinical_relevance,
                "patientExplanation": patient_explanation,
                "conduct": conduct
            }

            response = self.session.put(
                f"{API_BASE_URL}/score-items/{item_id}",
                json=payload
            )
            response.raise_for_status()
            return True
        except Exception as e:
            print(f"Erro ao atualizar item {item_id}: {e}")
            return False

    def enrich_item(self, item: Dict) -> bool:
        """Enriquece um item específico"""
        print(f"\n{'='*80}")
        print(f"Processando: {item['name']}")
        print(f"ID: {item['id']}")
        print(f"{'='*80}\n")

        # Definir conteúdo baseado no item
        content = self.get_content_for_item(item)

        # Atualizar via API
        success = self.update_score_item(
            item['id'],
            content['clinical_relevance'],
            content['patient_explanation'],
            content['conduct']
        )

        if success:
            print(f"✓ Item '{item['name']}' enriquecido com sucesso!")
        else:
            print(f"✗ Falha ao enriquecer item '{item['name']}'")

        return success

    def get_content_for_item(self, item: Dict) -> Dict:
        """Retorna o conteúdo clínico para cada item"""

        # Hipertensão arterial
        if item['id'] == "ea288642-425c-496c-a577-6f6ee3acb468":
            return {
                "clinical_relevance": """A hipertensão arterial sistêmica (HAS) é uma condição crônica caracterizada por níveis pressóricos persistentemente elevados (≥140/90 mmHg). Na medicina funcional integrativa, a HAS é compreendida como manifestação de disfunções metabólicas subjacentes, incluindo resistência à insulina, disfunção endotelial, estresse oxidativo e desequilíbrio do eixo hipotálamo-hipófise-adrenal.

O histórico de hipertensão representa fator de risco cardiovascular independente, associado a aumento de morbimortalidade por doença arterial coronariana, acidente vascular cerebral, insuficiência cardíaca e doença renal crônica. A abordagem funcional reconhece que cerca de 95% dos casos são hipertensão essencial, frequentemente relacionada a fatores modificáveis como obesidade visceral, resistência à insulina, inflamação crônica de baixo grau e disfunção autonômica.

A avaliação do histórico pessoal e familiar de HAS permite estratificação de risco e identificação de janelas terapêuticas. Pacientes com história de hipertensão apresentam maior probabilidade de alterações metabólicas precoces, mesmo quando normotensos no momento da avaliação, justificando investigação mais detalhada de marcadores inflamatórios, função endotelial e parâmetros metabólicos.""",

                "patient_explanation": """A pressão alta (hipertensão) é uma condição em que o sangue circula com força excessiva nas artérias. Quando você tem histórico de hipertensão, significa que seu corpo já manifestou essa dificuldade de manter a pressão em níveis saudáveis.

Na medicina funcional, entendemos que a pressão alta geralmente não surge sozinha - ela costuma estar conectada com outros desequilíbrios no corpo, como resistência à insulina, inflamação silenciosa, estresse crônico ou excesso de peso. Por isso, mesmo que sua pressão esteja controlada agora, é importante investigar e tratar as causas raiz do problema.

O histórico de hipertensão nos alerta para monitorar sua saúde cardiovascular de perto e adotar medidas preventivas personalizadas.""",

                "conduct": """1. Monitoramento regular da pressão arterial (MAPA 24h quando indicado)

2. Avaliação metabólica abrangente:
   - Perfil glicêmico e insulinêmico
   - Marcadores inflamatórios (PCR-us, homocisteína)
   - Perfil lipídico avançado
   - Função renal e eletrólitos

3. Intervenções nutricionais:
   - Padrão alimentar anti-hipertensivo (DASH adaptado)
   - Restrição de sódio individualizada (<2g/dia)
   - Otimização de magnésio, potássio e cálcio
   - Redução de alimentos ultraprocessados

4. Modulação do estilo de vida:
   - Atividade física regular (150min/semana exercício aeróbico + resistido)
   - Técnicas de manejo de estresse (respiração diafragmática, meditação)
   - Higiene do sono (7-9h/noite)
   - Redução de consumo alcoólico

5. Suplementação direcionada quando indicado:
   - Magnésio (300-400mg/dia)
   - Ômega-3 (2-3g/dia EPA+DHA)
   - CoQ10 (100-200mg/dia)
   - Probióticos específicos

6. Reavaliação periódica e ajustes terapêuticos conforme resposta individual"""
            }

        # Diabetes mellitus
        elif item['id'] == "8236ab3f-1ee1-4f17-a053-7f45b7162dd2":
            return {
                "clinical_relevance": """O diabetes mellitus (DM) representa uma síndrome metabólica complexa caracterizada por hiperglicemia crônica resultante de defeitos na secreção ou ação da insulina. Na perspectiva funcional integrativa, o DM é compreendido como manifestação final de disfunções metabólicas progressivas, envolvendo resistência à insulina, estresse oxidativo, glicação avançada, disfunção mitocondrial e inflamação crônica.

O histórico de diabetes confere risco substancialmente aumentado para complicações macro e microvasculares, incluindo doença cardiovascular, nefropatia, retinopatia e neuropatia. A medicina funcional reconhece que mesmo em estados pré-diabéticos, já ocorrem alterações metabólicas significativas com impacto sistêmico em múltiplos órgãos e sistemas.

A avaliação do histórico diabético permite identificar o tempo de doença, padrão de controle glicêmico, presença de complicações e resposta a intervenções prévias. Esta informação é fundamental para estratificação de risco, individualização terapêutica e prevenção de progressão. Pacientes com história familiar de DM apresentam risco aumentado e se beneficiam de rastreamento metabólico precoce e intervenções preventivas intensivas.""",

                "patient_explanation": """O diabetes é uma condição em que o corpo perde a capacidade de controlar adequadamente os níveis de açúcar (glicose) no sangue. Isso acontece quando o pâncreas não produz insulina suficiente ou quando as células do corpo não respondem bem à insulina.

Ter histórico de diabetes significa que seu corpo já manifestou essa dificuldade metabólica. Na medicina funcional, entendemos que o diabetes não surge de repente - ele é resultado de anos de desequilíbrios que podemos identificar e tratar antes que causem complicações.

Mesmo que seu diabetes esteja controlado, é importante manter acompanhamento regular, pois a doença pode afetar diversos órgãos ao longo do tempo, como coração, rins, olhos e nervos. A boa notícia é que com tratamento adequado e mudanças no estilo de vida, muitas pessoas conseguem reverter ou controlar muito bem a condição.""",

                "conduct": """1. Monitoramento glicêmico rigoroso:
   - Hemoglobina glicada (HbA1c) trimestral
   - Glicemia de jejum e pós-prandial
   - Frutosamina quando indicado
   - Consideração de monitoramento contínuo de glicose (CGM)

2. Avaliação de complicações e comorbidades:
   - Perfil lipídico completo
   - Função renal (creatinina, TFG, microalbuminúria)
   - Função hepática
   - Marcadores inflamatórios
   - Vitamina B12 (especialmente em uso de metformina)

3. Rastreamento de complicações:
   - Exame oftalmológico anual
   - Avaliação neurológica (monofilamento, biotensiometria)
   - Ecocardiograma se indicado

4. Intervenção nutricional individualizada:
   - Controle de carga glicêmica
   - Distribuição adequada de macronutrientes
   - Inclusão de fibras solúveis (>30g/dia)
   - Restrição de carboidratos refinados e açúcares

5. Otimização do estilo de vida:
   - Atividade física regular (combinação aeróbica + resistida)
   - Controle de peso (perda de 5-10% se sobrepeso/obesidade)
   - Manejo de estresse
   - Sono reparador

6. Suplementação funcional quando indicado:
   - Cromo (200-400mcg/dia)
   - Ácido alfa-lipóico (300-600mg/dia)
   - Canela (1-6g/dia)
   - Berberina (900-1500mg/dia)
   - Magnésio
   - Ômega-3

7. Reavaliação regular e ajuste terapêutico conforme resposta individual"""
            }

        # Pré-diabetes / resistência a insulina
        elif item['id'] == "3e876264-ae80-4c5c-bc5b-00fbf09daf98":
            return {
                "clinical_relevance": """A pré-diabetes e a resistência à insulina representam estados metabólicos intermediários entre a homeostase glicêmica normal e o diabetes mellitus estabelecido. Na medicina funcional integrativa, essas condições são reconhecidas como janelas terapêuticas cruciais, onde intervenções podem prevenir ou reverter a progressão para diabetes tipo 2 e suas complicações.

A resistência à insulina é caracterizada por resposta celular inadequada à ação da insulina, levando a hiperinsulinemia compensatória e, eventualmente, hiperglicemia. Este estado está associado a inflamação crônica de baixo grau, estresse oxidativo, disfunção mitocondrial, desequilíbrio do eixo intestino-fígado e acúmulo de gordura ectópica, particularmente hepática e visceral.

O histórico de pré-diabetes indica risco aumentado (5-10% ao ano) de progressão para DM2, além de associação com síndrome metabólica, doença cardiovascular, esteatose hepática não alcoólica e síndrome dos ovários policísticos. A identificação precoce permite implementação de intervenções intensivas no estilo de vida, que demonstram eficácia superior a 50% na prevenção da progressão para diabetes.""",

                "patient_explanation": """A pré-diabetes e a resistência à insulina são sinais de alerta do seu corpo indicando que o metabolismo do açúcar não está funcionando perfeitamente, mas ainda não chegou ao diabetes completo.

Pense na resistência à insulina como uma "surdez" das suas células ao sinal da insulina (hormônio que ajuda o açúcar a entrar nas células). Seu pâncreas precisa produzir cada vez mais insulina para fazer o mesmo trabalho, e eventualmente isso pode levar ao diabetes.

A boa notícia é que esta é uma fase reversível! Com mudanças no estilo de vida, alimentação adequada e acompanhamento médico, você pode normalizar seu metabolismo e prevenir a progressão para diabetes. É como parar um trem antes que ele ganhe muita velocidade.""",

                "conduct": """1. Avaliação metabólica detalhada:
   - Glicemia de jejum e pós-prandial (1h e 2h)
   - Hemoglobina glicada (HbA1c)
   - Insulina de jejum
   - Índices de resistência insulínica (HOMA-IR)
   - Peptídeo C

2. Investigação de comorbidades associadas:
   - Perfil lipídico avançado
   - Função hepática e marcadores de esteatose
   - Marcadores inflamatórios
   - Hormônios tireoidianos
   - Avaliação de composição corporal

3. Intervenção nutricional intensiva:
   - Restrição de carboidratos refinados
   - Controle de carga glicêmica das refeições
   - Inclusão de fibras solúveis e insolúveis
   - Jejum intermitente quando apropriado
   - Educação sobre índice glicêmico

4. Programa de exercícios estruturado:
   - Atividade aeróbica (150min/semana)
   - Treinamento resistido (2-3x/semana)
   - HIIT quando apropriado

5. Modulação do estilo de vida:
   - Perda de peso (meta: 5-10% se sobrepeso)
   - Otimização do sono
   - Manejo de estresse crônico
   - Restrição de consumo alcoólico

6. Suplementação funcional direcionada:
   - Berberina (500mg 3x/dia)
   - Cromo picolinato (200-400mcg/dia)
   - Ácido alfa-lipóico (300-600mg/dia)
   - Magnésio (400-600mg/dia)
   - Inositol (2-4g/dia)
   - Canela (1-3g/dia)

7. Monitoramento regular (3-6 meses) e ajuste terapêutico conforme resposta"""
            }

        # DM estabelecido
        elif item['id'] == "c79cf41f-f3fe-41c2-b7af-1e9880f62583":
            return {
                "clinical_relevance": """O diabetes mellitus estabelecido representa estágio avançado de disfunção metabólica com hiperglicemia crônica sustentada, frequentemente acompanhada de complicações micro e macrovasculares. Esta condição indica falência progressiva das células beta pancreáticas e/ou resistência insulínica severa, com impacto sistêmico significativo.

Na medicina funcional integrativa, o DM estabelecido é abordado como síndrome metabólica complexa envolvendo múltiplos sistemas: disfunção mitocondrial, estresse oxidativo intenso, glicação avançada de proteínas (AGEs), inflamação crônica, disbiose intestinal e alterações da matriz extracelular. O histórico de diabetes estabelecido confere risco cardiovascular equivalente a evento coronariano prévio.

A presença de diabetes estabelecido exige avaliação abrangente de complicações (retinopatia, nefropatia, neuropatia, doença cardiovascular), otimização metabólica intensiva e estratégias para prevenção de progressão. O tempo de doença e padrão de controle glicêmico (memória metabólica) influenciam significativamente o prognóstico e risco de complicações futuras.""",

                "patient_explanation": """Ter diabetes estabelecido significa que a doença está presente há algum tempo e requer acompanhamento médico contínuo. Diferente da pré-diabetes, aqui o corpo já perdeu significativamente a capacidade de controlar o açúcar no sangue de forma natural.

Isso não significa que a situação não possa melhorar! Muitas pessoas com diabetes estabelecido conseguem controlar muito bem a doença e prevenir complicações através de tratamento adequado, mudanças na alimentação e no estilo de vida.

O mais importante agora é manter o açúcar no sangue dentro da meta estabelecida pelo seu médico e fazer acompanhamento regular para detectar e prevenir complicações nos olhos, rins, nervos e coração. Quanto melhor o controle, menor o risco de problemas futuros.""",

                "conduct": """1. Monitoramento glicêmico intensivo:
   - HbA1c trimestral (meta individualizada)
   - Perfil glicêmico completo
   - Monitoramento contínuo de glicose quando indicado
   - Registro de hipoglicemias

2. Rastreamento sistemático de complicações:
   - Fundo de olho anual (retinopatia)
   - Microalbuminúria e função renal
   - Monofilamento e biotensiometria (neuropatia)
   - ECG/ecocardiograma conforme risco
   - Índice tornozelo-braquial
   - Avaliação vascular periférica

3. Avaliação de comorbidades:
   - Perfil lipídico completo
   - Pressão arterial (meta <130/80 mmHg)
   - Função hepática
   - Marcadores inflamatórios
   - Vitamina B12 (anual se uso de metformina)
   - Vitamina D

4. Otimização terapêutica medicamentosa:
   - Revisão e ajuste de antidiabéticos
   - Tratamento de hipertensão e dislipidemia
   - Antiagregação plaquetária quando indicada

5. Intervenção nutricional individualizada:
   - Contagem de carboidratos
   - Controle de porções
   - Distribuição adequada de refeições
   - Educação nutricional contínua

6. Programa de exercícios supervisionado:
   - Atividade aeróbica regular
   - Treinamento resistido
   - Avaliação pré-exercício conforme complicações

7. Suplementação funcional:
   - Ácido alfa-lipóico (neuropatia)
   - Benfotiamina
   - Magnésio
   - Ômega-3
   - Antioxidantes (vitamina C, E, selênio)

8. Suporte multidisciplinar:
   - Educação em diabetes
   - Suporte psicológico
   - Acompanhamento de especialistas conforme complicações

9. Reavaliação regular e ajustes conforme resposta e evolução"""
            }

        # Obesidade
        elif item['id'] == "126d82d2-e32d-4f47-bca1-5035e2abc8d2":
            return {
                "clinical_relevance": """A obesidade é definida por acúmulo excessivo de gordura corporal (IMC ≥30 kg/m²) e representa condição metabólica complexa de etiologia multifatorial. Na medicina funcional integrativa, a obesidade é compreendida não apenas como excesso de peso, mas como síndrome de disfunção metabólica sistêmica, envolvendo resistência à insulina, inflamação crônica de baixo grau (metainflamação), disfunção mitocondrial, desequilíbrio do eixo intestino-fígado e alterações neuroendócrinas.

O histórico de obesidade confere risco substancialmente aumentado para desenvolvimento de comorbidades metabólicas (diabetes tipo 2, dislipidemia, hipertensão), doença cardiovascular, esteatose hepática não alcoólica, síndrome de apneia obstrutiva do sono, osteoartrite e determinados tipos de câncer. A distribuição da gordura corporal, particularmente a obesidade visceral/central, é determinante mais relevante de risco metabólico do que o IMC isolado.

A avaliação do histórico de obesidade deve incluir análise de padrões de ganho de peso, tentativas prévias de emagrecimento, presença de comorbidades, fatores ambientais, genéticos e comportamentais. A abordagem funcional reconhece que a obesidade frequentemente possui raízes em disfunções subjacentes (hipotireoidismo, resistência à insulina, disbiose, deficiências nutricionais, estresse crônico) que devem ser identificadas e tratadas.""",

                "patient_explanation": """A obesidade não é simplesmente uma questão de "comer demais" - é uma condição médica complexa que envolve desequilíbrios metabólicos, hormonais e inflamatórios no corpo. Quando há histórico de obesidade, isso indica que seu corpo tem dificuldade em regular o peso de forma natural.

Na medicina funcional, investigamos as causas raiz da obesidade: pode ser resistência à insulina, desequilíbrios hormonais (como tireoide), inflamação crônica, alterações no intestino, estresse crônico ou até deficiências nutricionais. Cada pessoa tem uma combinação única de fatores.

O objetivo não é apenas perder peso, mas restaurar o metabolismo saudável. Quando tratamos as causas subjacentes, o corpo naturalmente encontra seu peso ideal de forma mais sustentável. Isso envolve mudanças na alimentação, atividade física, sono, manejo de estresse e, às vezes, suplementação específica.""",

                "conduct": """1. Avaliação abrangente da composição corporal:
   - Bioimpedância ou DEXA
   - Circunferência abdominal
   - Relação cintura-quadril
   - Percentual de gordura corporal

2. Investigação metabólica completa:
   - Perfil glicêmico e insulinêmico (HOMA-IR)
   - Perfil lipídico avançado
   - Hormônios tireoidianos (TSH, T4L, T3L, anti-TPO)
   - Cortisol (ritmo circadiano)
   - Leptina e grelina quando indicado
   - Vitamina D

3. Rastreamento de comorbidades:
   - Pressão arterial
   - Função hepática e marcadores de esteatose
   - Marcadores inflamatórios (PCR-us)
   - Avaliação de apneia do sono (Epworth, polissonografia)

4. Investigação de causas subjacentes:
   - Avaliação da microbiota intestinal
   - Pesquisa de intolerâncias alimentares
   - Marcadores de estresse oxidativo
   - Avaliação psicoemocional

5. Programa nutricional individualizado:
   - Restrição calórica moderada (déficit 500-750kcal/dia)
   - Controle de carga glicêmica
   - Aumento de proteínas (1,6-2g/kg peso ideal)
   - Inclusão de fibras (>30g/dia)
   - Jejum intermitente quando apropriado
   - Dieta anti-inflamatória

6. Programa de exercícios progressivo:
   - Atividade aeróbica (200-300min/semana)
   - Treinamento resistido (3x/semana)
   - NEAT (Non-Exercise Activity Thermogenesis)

7. Modulação do estilo de vida:
   - Higiene do sono (7-9h/noite)
   - Técnicas de manejo de estresse
   - Mindful eating
   - Suporte comportamental

8. Suplementação funcional quando indicada:
   - Ômega-3 (2-3g/dia)
   - Probióticos específicos
   - Vitamina D (adequação níveis)
   - Magnésio
   - Cromo
   - 5-HTP ou triptofano (controle de apetite)

9. Monitoramento regular (mensal inicialmente) com ajustes conforme resposta

10. Consideração de terapias adjuvantes:
    - Acupuntura
    - Terapia cognitivo-comportamental
    - Farmacoterapia quando indicada"""
            }

        # Dislipidemia
        elif item['id'] == "58235b33-5431-48d9-83f4-77d158caf4da":
            return {
                "clinical_relevance": """A dislipidemia representa alterações quantitativas ou qualitativas dos lipídios plasmáticos, caracterizada por elevação de colesterol total, LDL-colesterol, triglicerídeos e/ou redução de HDL-colesterol. Na medicina funcional integrativa, a dislipidemia é compreendida como manifestação de desequilíbrios metabólicos sistêmicos, incluindo resistência à insulina, inflamação crônica, estresse oxidativo, disfunção hepática e alterações da microbiota intestinal.

O histórico de dislipidemia confere risco cardiovascular significativo, particularmente quando associado a outros fatores (hipertensão, diabetes, tabagismo, história familiar). A abordagem funcional reconhece que o padrão das partículas lipídicas (tamanho e densidade) é frequentemente mais relevante que os valores absolutos - partículas LDL pequenas e densas (fenótipo B) conferem maior aterogenicidade do que LDL grandes e flutuantes.

A avaliação do histórico de dislipidemia deve incluir análise do perfil lipídico avançado, marcadores inflamatórios, apolipoproteínas e fatores de risco residual. A origem da dislipidemia pode ser primária (genética) ou secundária (dieta, sedentarismo, resistência insulínica, hipotireoidismo), exigindo abordagem diagnóstica diferenciada.""",

                "patient_explanation": """A dislipidemia significa que os níveis de gorduras (lipídios) no seu sangue estão fora do equilíbrio ideal. Isso inclui o colesterol (LDL "ruim", HDL "bom") e os triglicerídeos. Ter histórico de dislipidemia indica que seu corpo tem dificuldade em regular essas gorduras adequadamente.

Na medicina funcional, entendemos que o colesterol alto não é apenas uma questão de "comer muita gordura". Muitas vezes está relacionado a inflamação no corpo, resistência à insulina, problemas na tireoide ou desequilíbrios no intestino. O tipo e o tamanho das partículas de colesterol também importam muito - nem todo colesterol LDL é igualmente prejudicial.

O objetivo é não apenas baixar os números, mas reduzir seu risco real de problemas cardiovasculares. Isso pode envolver mudanças na alimentação, exercícios, manejo de estresse e tratamento de causas subjacentes como inflamação e resistência à insulina.""",

                "conduct": """1. Avaliação lipídica avançada:
   - Perfil lipídico completo (jejum e pós-prandial)
   - LDL-colesterol calculado e direto
   - Apolipoproteína A1 e B
   - Lipoproteína (a)
   - Partículas LDL (tamanho e número)
   - Colesterol não-HDL
   - Índice TG/HDL

2. Investigação de causas secundárias:
   - Perfil glicêmico e insulinêmico
   - Hormônios tireoidianos
   - Função hepática
   - Função renal
   - Marcadores inflamatórios (PCR-us, homocisteína)

3. Estratificação de risco cardiovascular:
   - Escore de Framingham ou equivalente
   - Escore de cálcio coronário quando indicado
   - Espessura médio-intimal carotídea
   - Índice tornozelo-braquial

4. Intervenção nutricional direcionada:
   - Padrão alimentar mediterrâneo ou similar
   - Redução de gorduras trans e saturadas
   - Aumento de ômega-3 (peixes, linhaça, chia)
   - Inclusão de fitosteróis
   - Fibras solúveis (aveia, psyllium, leguminosas)
   - Restrição de carboidratos refinados (se triglicerídeos elevados)
   - Alimentos funcionais (alho, cúrcuma, chá verde)

5. Programa de exercícios regular:
   - Atividade aeróbica (150-300min/semana)
   - Exercícios resistidos
   - Foco em redução de gordura visceral

6. Modulação do estilo de vida:
   - Controle de peso
   - Cessação de tabagismo
   - Moderação de consumo alcoólico
   - Manejo de estresse

7. Suplementação funcional baseada em evidências:
   - Ômega-3 (2-4g/dia EPA+DHA)
   - Fitosteróis (2g/dia)
   - Niacina (quando indicado e monitorado)
   - Berberina (500mg 3x/dia)
   - Coenzima Q10 (100-200mg/dia)
   - Arroz vermelho fermentado (quando apropriado)
   - Alho envelhecido
   - Psyllium

8. Tratamento de comorbidades:
   - Otimização de diabetes/pré-diabetes
   - Controle de hipertensão
   - Tratamento de hipotireoidismo

9. Monitoramento regular (3-6 meses) e ajuste terapêutico conforme resposta

10. Consideração de farmacoterapia (estatinas, ezetimibe) quando indicado por alto risco ou resposta insuficiente a intervenções no estilo de vida"""
            }

        # Câncer
        elif item['id'] == "9f646670-c0eb-40d5-bcb0-b2f069e41a29":
            return {
                "clinical_relevance": """O histórico de câncer representa condição de importância clínica fundamental, indicando exposição prévia a processo neoplásico que pode ter implicações prognósticas, terapêuticas e preventivas significativas. Na medicina funcional integrativa, o câncer é compreendido como doença sistêmica resultante de disfunções em múltiplos níveis: genético-epigenético, imunológico, metabólico, inflamatório e ambiental.

A presença de histórico oncológico exige vigilância contínua para recidiva, desenvolvimento de segundo tumor primário e complicações tardias do tratamento (cirurgia, radioterapia, quimioterapia). Pacientes com câncer prévio apresentam risco aumentado de novos eventos neoplásicos, especialmente quando há predisposição genética ou exposição a fatores de risco persistentes.

A abordagem funcional reconhece a importância do terreno biológico na oncogênese e progressão tumoral. Fatores como inflamação crônica, estresse oxidativo, disfunção mitocondrial, resistência à insulina, disbiose intestinal e imunodeficiência podem contribuir para carcinogênese e representam alvos terapêuticos para prevenção de recorrência. O suporte metabólico e nutricional adequado durante e após o tratamento oncológico pode melhorar qualidade de vida, tolerância terapêutica e desfechos clínicos.""",

                "patient_explanation": """Ter histórico de câncer significa que você já passou pela experiência de ter células que cresceram de forma descontrolada no organismo. Mesmo após o tratamento bem-sucedido, é importante manter acompanhamento médico regular para monitorar possíveis recidivas (retorno do câncer) e detectar precocemente qualquer novo problema.

Na medicina funcional, trabalhamos para otimizar o "terreno" do seu corpo - ou seja, criar um ambiente interno menos favorável ao desenvolvimento de células cancerígenas. Isso envolve fortalecer o sistema imunológico, reduzir inflamação crônica, otimizar o metabolismo e corrigir desequilíbrios nutricionais.

Além do acompanhamento oncológico tradicional, focamos em: alimentação anti-inflamatória, suporte nutricional adequado, manejo de estresse, sono de qualidade, exercícios apropriados e suplementação direcionada. O objetivo é não apenas prevenir recidivas, mas também promover sua melhor qualidade de vida e vitalidade.""",

                "conduct": """1. Acompanhamento oncológico especializado:
   - Seguimento conforme protocolo específico do tipo de câncer
   - Exames de imagem periódicos
   - Marcadores tumorais quando aplicável
   - Vigilância para segundo tumor primário

2. Avaliação metabólica e imunológica:
   - Perfil glicêmico e insulinêmico
   - Marcadores inflamatórios (PCR-us, IL-6, TNF-alfa)
   - Perfil lipídico
   - Vitamina D (manter >40ng/mL)
   - Perfil de micronutrientes
   - Função imunológica (quando disponível)

3. Avaliação de sequelas do tratamento:
   - Função cardiovascular (quimio/radioterapia)
   - Função pulmonar
   - Função renal e hepática
   - Densidade óssea (quando indicado)
   - Neuropatia periférica
   - Função cognitiva
   - Saúde mental (ansiedade, depressão)

4. Programa nutricional anticâncer:
   - Dieta anti-inflamatória e antioxidante
   - Restrição de açúcares e carboidratos refinados
   - Aumento de vegetais crucíferos
   - Inclusão de alimentos funcionais (cúrcuma, chá verde, cogumelos medicinais)
   - Consideração de jejum intermitente (sob supervisão)
   - Hidratação adequada

5. Programa de exercícios adaptado:
   - Atividade física regular (reduz risco de recorrência)
   - Exercícios aeróbicos e resistidos
   - Yoga ou tai chi (benefícios psicofísicos)

6. Suporte psicoemocional:
   - Acompanhamento psicológico
   - Técnicas de manejo de estresse
   - Meditação e mindfulness
   - Grupos de suporte

7. Otimização do estilo de vida:
   - Sono reparador (7-9h/noite)
   - Redução de exposição a toxinas ambientais
   - Cessação de tabagismo
   - Moderação ou eliminação de álcool

8. Suplementação funcional baseada em evidências (coordenada com oncologista):
   - Vitamina D (2.000-5.000UI/dia)
   - Ômega-3 (2-3g/dia)
   - Probióticos de cepas específicas
   - Curcumina (biodisponível)
   - Quercetina
   - Vitamina C (doses conforme orientação)
   - Selênio
   - Coenzima Q10 (especialmente se uso prévio de antraciclinas)

9. Modulação da microbiota:
   - Probióticos e prebióticos
   - Alimentos fermentados
   - Fibras diversificadas

10. Monitoramento regular multidisciplinar e ajustes conforme evolução e necessidades individuais

**IMPORTANTE: Todas as intervenções complementares devem ser coordenadas com a equipe oncológica, especialmente durante tratamento ativo."""
            }

        # Insuficiência cardíaca
        elif item['id'] == "fb8c9e63-8d36-4e03-bb24-dbdc747e5fd4":
            return {
                "clinical_relevance": """A insuficiência cardíaca (IC) é síndrome clínica complexa resultante de disfunção estrutural ou funcional do coração, com consequente incapacidade de ejeção ou enchimento adequados para atender as demandas metabólicas teciduais. Na medicina funcional integrativa, a IC é compreendida como manifestação final de múltiplas disfunções cardiovasculares, metabólicas e inflamatórias, envolvendo disfunção mitocondrial cardiomiocitária, estresse oxidativo, neurohormonal e remodelamento ventricular.

O histórico de IC confere prognóstico reservado com morbimortalidade significativa, exigindo manejo multidisciplinar intensivo. A classificação funcional (NYHA) e fração de ejeção (IC com fração preservada vs reduzida) orientam estratificação de risco e abordagem terapêutica. A IC está frequentemente associada a comorbidades como hipertensão, diabetes, doença coronariana e valvopatias.

A abordagem funcional reconhece que déficits nutricionais (especialmente tiamina, coenzima Q10, magnésio, carnitina), inflamação sistêmica, caquexia cardíaca e disfunção mitocondrial contribuem para progressão da doença. O suporte metabólico adequado, otimização hemodinâmica e correção de desequilíbrios subjacentes podem melhorar capacidade funcional, qualidade de vida e desfechos clínicos.""",

                "patient_explanation": """A insuficiência cardíaca significa que o coração não está bombeando sangue de forma tão eficiente quanto deveria. Isso não significa que o coração vai "parar" - ele continua batendo, mas trabalha com mais dificuldade para suprir as necessidades do corpo, podendo causar cansaço, falta de ar e inchaço.

Ter histórico de insuficiência cardíaca exige cuidado contínuo e acompanhamento médico regular. Na medicina funcional, além do tratamento convencional essencial (medicamentos prescritos pelo cardiologista), trabalhamos para otimizar o suporte nutricional e metabólico do coração.

O músculo cardíaco precisa de nutrientes específicos para funcionar bem, especialmente coenzima Q10, magnésio e L-carnitina. Também é fundamental controlar o sal na alimentação, manter peso adequado, fazer exercícios leves supervisionados e monitorar sinais de descompensação (como ganho rápido de peso ou piora da falta de ar).""",

                "conduct": """1. Acompanhamento cardiológico especializado:
   - Avaliação clínica regular (NYHA, sinais congestivos)
   - Ecocardiograma periódico (função ventricular, valvas)
   - Eletrocardiograma
   - Peptídeos natriuréticos (BNP/NT-proBNP)
   - Teste de caminhada de 6 minutos

2. Monitoramento de parâmetros vitais:
   - Peso diário (alerta para ganho >2kg em 3 dias)
   - Pressão arterial
   - Frequência cardíaca
   - Saturação de oxigênio

3. Avaliação laboratorial abrangente:
   - Função renal (creatinina, TFG, eletrólitos)
   - Função hepática
   - Perfil lipídico
   - Perfil glicêmico
   - Hemograma (anemia agrava IC)
   - Perfil tireoidiano
   - Ferro e ferritina

4. Avaliação nutricional e metabólica:
   - Estado nutricional (risco de caquexia)
   - Vitamina B1 (tiamina) - déficit comum em IC
   - Vitamina D
   - Magnésio
   - Coenzima Q10 sérica

5. Otimização farmacológica (em coordenação com cardiologista):
   - IECA/BRA, betabloqueadores, antagonistas aldosterona
   - Diuréticos conforme status congestivo
   - SGLT2-i quando indicado

6. Intervenção nutricional específica:
   - Restrição de sódio (1,5-2g/dia)
   - Controle hídrico individualizado
   - Pequenas refeições frequentes (6-8x/dia)
   - Proteína adequada (1,1-1,5g/kg)
   - Evitar excesso de líquidos

7. Programa de exercícios supervisionado:
   - Reabilitação cardíaca quando disponível
   - Exercícios aeróbicos leves a moderados
   - Treino resistido leve
   - Progressão gradual conforme tolerância

8. Suplementação funcional baseada em evidências:
   - Coenzima Q10 (100-300mg/dia) - ESSENCIAL
   - Tiamina (B1) 100-300mg/dia
   - Magnésio (400-600mg/dia)
   - L-carnitina (1-3g/dia)
   - Taurina (2-3g/dia)
   - D-ribose (5g 3x/dia)
   - Ômega-3 (1g/dia)
   - Vitamina D (adequação níveis)

9. Modulação do estilo de vida:
   - Cessação de tabagismo
   - Restrição de álcool
   - Vacinação (influenza, pneumococo)
   - Controle de peso
   - Manejo de estresse

10. Monitoramento de sinais de descompensação:
    - Educação sobre sintomas de alerta
    - Plano de ação para piora aguda
    - Acesso rápido a atendimento especializado

11. Tratamento de comorbidades:
    - Controle rigoroso de hipertensão
    - Otimização de diabetes
    - Tratamento de anemia
    - Controle de arritmias

**IMPORTANTE: Todas as intervenções devem ser coordenadas com o cardiologista assistente. Suplementação com CoQ10 é particularmente importante em uso de estatinas."""
            }

        # Item 9: Arritmia
        elif item['id'] == "fc4ea1da-ad8c-419c-be76-c58405f69b11":
            return {
                "clinical_relevance": """Arritmia cardíaca refere-se a qualquer distúrbio no ritmo normal do coração, podendo manifestar-se como taquiarritmias (frequência cardíaca acelerada), bradiarritmias (frequência reduzida) ou batimentos ectópicos. Na medicina funcional integrativa, as arritmias são compreendidas como manifestações de desequilíbrios elétricos, metabólicos, autonômicos e estruturais do sistema cardiovascular.

O histórico de arritmia confere implicações clínicas significativas, variando desde condições benignas (extrassístoles atriais ou ventriculares isoladas) até arritmias potencialmente fatais (fibrilação ventricular, taquicardia ventricular sustentada). A fibrilação atrial, arritmia mais comum em adultos, está associada a risco quintuplicado de AVC embólico e piora progressiva da função cardíaca.

A abordagem funcional reconhece que arritmias frequentemente possuem gatilhos modificáveis: distúrbios eletrolíticos (especialmente magnésio e potássio), disfunção tireoidiana, desequilíbrio autonômico (excesso simpático), apneia obstrutiva do sono, consumo excessivo de estimulantes (cafeína, álcool), estresse oxidativo e inflamação sistêmica. O suporte nutricional adequado e correção de déficits podem reduzir significativamente a carga arrítmica e melhorar o prognóstico.""",

                "patient_explanation": """Arritmia significa que o coração não está batendo no ritmo regular esperado - pode bater muito rápido, muito devagar ou com batimentos "extras" fora do tempo. Nem toda arritmia é perigosa, mas ter histórico de arritmia exige acompanhamento médico.

Algumas arritmias são apenas incômodas (aquela sensação de "coração acelerado" ou "falhas"), enquanto outras podem ser sérias e aumentar o risco de derrame (AVC) ou problemas cardíacos. A fibrilação atrial, por exemplo, é uma arritmia comum que faz o coração bater de forma desorganizada e pode formar coágulos.

Na medicina funcional, investigamos fatores que podem estar provocando ou piorando a arritmia: falta de minerais como magnésio e potássio, problemas na tireoide, apneia do sono, excesso de cafeína ou álcool, estresse crônico. Corrigir esses desequilíbrios pode melhorar muito o controle da arritmia.""",

                "conduct": """1. Avaliação cardiológica especializada com ECG, Holter 24-48h, ecocardiograma

2. Investigação de causas: eletrólitos (magnésio, potássio), função tireoidiana, rastreamento apneia do sono

3. Modulação do estilo de vida: redução de cafeína e álcool, manejo de estresse

4. Suplementação direcionada: magnésio (400-600mg/dia), CoQ10, ômega-3, taurina

5. Tratamento farmacológico quando indicado (antiarrítmicos, anticoagulação)

6. Monitoramento regular"""
            }

        # Items 10-20 com conteúdo resumido para manter o script funcional
        elif item['id'] == "d4d8283f-42c4-4de7-a98d-603c41d12ae6":  # Doença cardiovascular
            return {
                "clinical_relevance": """O histórico de doença cardiovascular estabelecida (IAM, AVC, revascularização) representa condição de altíssimo risco para eventos recorrentes. Estes pacientes estão em prevenção secundária, exigindo manejo intensivo multifatorial. Na medicina funcional, eventos cardiovasculares prévios refletem disfunção endotelial crônica, aterosclerose avançada e inflamação vascular persistente. O risco de eventos recorrentes é substancial (15-20% em 5 anos), mas pode ser reduzido através de intervenções intensivas no estilo de vida, otimização farmacológica e suporte metabólico direcionado.""",
                "patient_explanation": """Ter histórico de doença cardiovascular significa que você já sofreu um evento grave no coração ou vasos sanguíneos. Isso indica problema de "entupimento" das artérias que precisa de atenção contínua. É fundamental acompanhamento médico rigoroso, medicamentos essenciais (antiagregantes, estatinas, anti-hipertensivos), alimentação anti-inflamatória, exercícios supervisionados e controle de todos os fatores de risco para prevenir novos eventos.""",
                "conduct": """1. Acompanhamento cardiológico rigoroso
2. Controle intensivo de fatores de risco (PA <130/80, LDL <70mg/dL, HbA1c <7%)
3. Terapia farmacológica otimizada (AAS, estatina, betabloqueador, IECA/BRA)
4. Reabilitação cardiovascular
5. Dieta mediterrânea
6. Suplementação: ômega-3, CoQ10, vitamina D, magnésio
7. Cessação de tabagismo absoluta"""
            }

        else:
            # Para os demais items (11-20), criar conteúdo genérico funcional
            disease_specific = {
                "82ba4b8f-708a-4adb-8bd5-a242485e0446": {  # Doença renal crônica
                    "cr": "DRC é disfunção renal progressiva com TFG <60mL/min/1,73m² ou lesão renal ≥3 meses. Confere risco cardiovascular elevado e requer controle rigoroso de PA, diabetes e proteinúria para retardar progressão.",
                    "pe": "Doença renal crônica significa que os rins perderam parte da capacidade de filtrar o sangue. É fundamental retardar a progressão através de controle da pressão arterial, diabetes, alimentação adaptada e evitar medicamentos que prejudicam os rins.",
                    "cd": "1. Acompanhamento nefrológico 2. Controle PA (<130/80) e diabetes 3. IECA/BRA 4. Dieta com restrição proteica moderada, sódio, fósforo, potássio 5. Manejo de complicações (anemia, distúrbio mineral-ósseo) 6. Evitar nefrotóxicos"
                },
                "3a1e71df-975f-4a7d-815c-a56b3e3735c9": {  # Outras doenças renais
                    "cr": "Doenças renais diversas incluem glomerulopatias, doenças túbulo-intersticiais, doenças císticas e vasculares renais. Requerem abordagem individualizada conforme etiologia específica e potencial de progressão.",
                    "pe": "Outras doenças renais incluem vários problemas nos rins com características próprias. O importante é fazer acompanhamento regular, identificar a causa específica, monitorar a função renal e proteger os rins de danos adicionais.",
                    "cd": "1. Diagnóstico etiológico específico 2. Tratamento conforme causa 3. Nefroproteção geral 4. Monitoramento de função renal 5. Prevenção de progressão"
                },
                "34fbbed1-e8a2-4103-bda8-a522bc357978": {  # Nefrite
                    "cr": "Nefrite refere-se a processos inflamatórios renais, geralmente glomerulonefrites com mecanismos imunológicos. Requer tratamento precoce para minimizar danos permanentes ao parênquima renal.",
                    "pe": "Nefrite é inflamação dos rins que pode ter várias causas (infecções, doenças autoimunes, medicamentos). O tratamento rápido é essencial para evitar danos permanentes, podendo incluir anti-inflamatórios, imunossupressores e controle da pressão.",
                    "cd": "1. Avaliação nefrológica urgente 2. Diagnóstico etiológico 3. Tratamento específico (imunossupressão quando indicada) 4. Controle de PA 5. Monitoramento de proteinúria e função renal"
                },
                "addbf7df-76b6-4bdb-b7c8-1c239c373389": {  # Nefrótica
                    "cr": "Síndrome nefrótica: proteinúria maciça (>3,5g/24h), hipoalbuminemia, edema e hipercolesterolemia. Confere risco de trombose, infecções e desnutrição. Tratamento visa reduzir proteinúria e prevenir complicações.",
                    "pe": "Síndrome nefrótica é quando os rins perdem muita proteína pela urina, causando inchaço, urina espumosa e alterações no colesterol. O tratamento pode incluir medicamentos para reduzir a perda de proteína, controlar o inchaço e proteger contra trombose.",
                    "cd": "1. Biópsia renal 2. Tratamento específico conforme etiologia 3. IECA/BRA 4. Diuréticos 5. Anticoagulação se albumina <2g/dL 6. Restrição de sódio 7. Estatinas"
                },
                "6ba04946-25de-49c2-8fc0-7dcd8dcb75bc": {  # Litíase
                    "cr": "Litíase renal tem prevalência de 10-15% e alta recorrência (50% em 5-10 anos). Formação de cálculos reflete desequilíbrios metabólicos específicos. Prevenção requer identificação e correção de fatores predisponentes.",
                    "pe": "Pedra nos rins se forma quando substâncias na urina se concentram demais e cristalizam. Quem já teve tem 50% de chance de ter novamente. Prevenção: muita água, ajustes na alimentação conforme tipo de pedra.",
                    "cd": "1. Análise do cálculo 2. Investigação metabólica (urina 24h) 3. Hidratação (2-2,5L/dia) 4. Dieta conforme tipo de cálculo 5. Citrato de potássio ou tiazídicos quando indicado"
                },
                "a1a681ce-27a4-434a-aeae-0682b3375758": {  # ITU
                    "cr": "ITUs são extremamente comuns, afetando 50% das mulheres ao longo da vida. ITUs recorrentes (≥2/6 meses) refletem desequilíbrios da microbiota urogenital, disfunção imunológica ou fatores anatômicos predisponentes.",
                    "pe": "Infecção urinária é infecção bacteriana da bexiga, uretra ou rins. Sintomas: ardor ao urinar, vontade frequente, urina turva. Prevenção: hidratação, urinar após relações sexuais, higiene adequada, cranberry.",
                    "cd": "1. Urocultura com antibiograma 2. Antibioticoterapia guiada 3. Prevenção: hidratação, urinar após sexo, cranberry (36mg PACs/dia), probióticos lactobacilos vaginais"
                },
                "13bff813-f2dc-4d1f-b3bd-8f8e8f731634": {  # Doenças virais crônicas
                    "cr": "Doenças virais crônicas (hepatites B/C, HIV, herpes vírus) caracterizam-se por persistência viral com ativação imunológica crônica e risco de complicações. Requerem monitoramento prolongado e suporte imunológico.",
                    "pe": "Doenças virais crônicas são infecções por vírus que permanecem no corpo por tempo prolongado. O acompanhamento médico é fundamental para monitorar atividade viral, prevenir complicações e em muitos casos usar medicamentos antivirais.",
                    "cd": "1. Identificação do agente viral 2. Tratamento antiviral específico quando indicado 3. Suporte imunológico (vitamina D, C, zinco, ômega-3) 4. Estilo de vida saudável 5. Prevenção de transmissão"
                },
                "3b18f7b7-f600-44d0-837c-1ad343661dc2": {  # HIV
                    "cr": "HIV é infecção crônica com depleção de TCD4+. Com TARV, tornou-se condição manejável com expectativa de vida próxima ao normal. Pacientes têm risco aumentado de comorbidades cardiovasculares, metabólicas e neoplásicas.",
                    "pe": "HIV ataca o sistema imunológico. Com tratamento (antirretrovirais), o vírus fica controlado, a imunidade se recupera e a expectativa de vida é normal. É fundamental tomar medicamentos corretamente todos os dias e fazer acompanhamento regular.",
                    "cd": "1. TARV imediato (1 comprimido/dia) 2. Objetivo: carga viral indetectável 3. Monitoramento CD4 e carga viral 4. Profilaxias conforme CD4 5. Manejo de comorbidades 6. Suporte nutricional (vitamina D, ômega-3)"
                },
                "8014a88c-4519-4e76-b57a-c79c8e45a162": {  # Hepatite B
                    "cr": "Hepatite B crônica afeta 240 milhões globalmente com risco de cirrose e hepatocarcinoma. Antivirais (tenofovir, entecavir) suprimem replicação viral e reduzem progressão. Monitoramento de hepatocarcinoma é essencial mesmo com supressão viral.",
                    "pe": "Hepatite B crônica pode danificar o fígado ao longo do tempo. O tratamento com antivirais controla o vírus e protege o fígado. É importante: medicação conforme prescrito, exames regulares, evitar álcool, rastrear câncer de fígado.",
                    "cd": "1. HBV-DNA, HBeAg, anti-HBe 2. Elastografia hepática 3. Tratamento: tenofovir ou entecavir 4. Rastreamento hepatocarcinoma (ultrassom + AFP semestral) 5. Cessação de álcool 6. Vacinação de contatos"
                },
                "9ff6f838-bc35-4039-8b37-f70fdec68733": {  # Hepatite C
                    "cr": "Hepatite C crônica (71 milhões globalmente) foi revolucionada pelos DAAs com taxas de cura >95% em 8-12 semanas. Mesmo após cura, pacientes com cirrose mantêm risco de hepatocarcinoma, exigindo vigilância contínua.",
                    "pe": "Hepatite C tem CURA com medicamentos modernos (>95% de sucesso em 8-12 semanas)! Mesmo após cura, quem teve cirrose precisa continuar rastreando câncer de fígado. Evitar álcool totalmente e manter peso saudável.",
                    "cd": "1. HCV-RNA, genotipagem 2. Elastografia hepática 3. Tratamento com DAAs (8-12 semanas) 4. RVS = cura 5. Se cirrose: rastreamento hepatocarcinoma contínuo 6. Cessação de álcool 7. Vacinação hepatite A/B"
                }
            }

            if item['id'] in disease_specific:
                spec = disease_specific[item['id']]
                return {
                    "clinical_relevance": spec["cr"],
                    "patient_explanation": spec["pe"],
                    "conduct": spec["cd"]
                }

            # Fallback genérico
            return {
                "clinical_relevance": f"Relevância clínica para {item['name']} em desenvolvimento. Condição requer avaliação especializada e acompanhamento regular.",
                "patient_explanation": f"{item['name']} é uma condição que requer acompanhamento médico. Consulte seu médico para orientações específicas.",
                "conduct": f"1. Avaliação médica especializada 2. Investigação diagnóstica apropriada 3. Tratamento conforme protocolo 4. Monitoramento regular"
            }

    def process_all_items(self):
        """Processa todos os 20 items"""
        results = {
            "success": [],
            "failed": [],
            "total": len(ITEMS)
        }

        print(f"\n{'='*80}")
        print(f"Iniciando processamento de {len(ITEMS)} items")
        print(f"{'='*80}\n")

        for item in ITEMS:
            success = self.enrich_item(item)
            if success:
                results["success"].append(item['name'])
            else:
                results["failed"].append(item['name'])

            # Pequeno delay entre requisições
            time.sleep(0.5)

        # Sumário final
        print(f"\n{'='*80}")
        print("SUMÁRIO DO PROCESSAMENTO")
        print(f"{'='*80}")
        print(f"Total de items: {results['total']}")
        print(f"Sucesso: {len(results['success'])}")
        print(f"Falhas: {len(results['failed'])}")

        if results['failed']:
            print(f"\nItems com falha:")
            for name in results['failed']:
                print(f"  - {name}")

        print(f"{'='*80}\n")

        return results

def main():
    enricher = ScoreItemEnricher()

    # Login
    if not enricher.login():
        print("Erro: Não foi possível fazer login. Encerrando.")
        return

    # Processar todos items
    results = enricher.process_all_items()

    # Salvar resultados em arquivo
    with open('/home/user/plenya/disease_history_batch1_results.json', 'w') as f:
        json.dump(results, f, indent=2, ensure_ascii=False)

    print(f"Resultados salvos em: /home/user/plenya/disease_history_batch1_results.json")

if __name__ == "__main__":
    main()
