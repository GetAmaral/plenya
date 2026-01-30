#!/usr/bin/env python3
"""
Script para enriquecer 20 Score Items do grupo "Histórico de doenças" - PARTE 3
Foco: Medicações e condições específicas (Antidiabéticos, GLP-1, Arritmia, Artrite, etc)
"""

import requests
import json
import time
from typing import Dict, List

# Configuração
API_BASE_URL = "http://localhost:3001/api/v1"
EMAIL = "import@plenya.com"
PASSWORD = "Import@123456"

# IDs dos 20 items - PARTE 3
ITEMS = [
    {"id": "2aa595d8-a3e5-4f7f-964b-88f5c2451e79", "name": "Antidiabéticos"},
    {"id": "791917c3-1c37-46fa-89da-079a8adb2bd7", "name": "Antiosteoporóticos"},
    {"id": "6dcbb3d9-611d-4c46-aa62-a9dbb13341dd", "name": "Antiparkinsonianos"},
    {"id": "3511d78f-cfc2-4d9b-bd1f-32ed3a3120ae", "name": "Antipsicóticos"},
    {"id": "9323cdf9-9b56-4638-b170-da2008d298b5", "name": "Antivirais"},
    {"id": "9d6ff60e-000b-493d-bb4e-43b4db957188", "name": "GLP-1"},
    {"id": "e04bd4fa-2076-4068-9a9d-ab0e9c1d16a2", "name": "Arritmia (duplicata - item medicação)"},
    {"id": "78c36d3f-b74e-436d-a224-3d66f3a5d7c1", "name": "Artrite"},
    {"id": "d61bf844-b1a2-4096-b954-53031694c590", "name": "Artrite reumatoide"},
    {"id": "47ca7e29-247c-4dbb-a1a8-ae723ade3004", "name": "Osteoartrite"},
    {"id": "5cb9c7b1-fb72-487b-ab16-51343f66bc18", "name": "Outras artrites"},
    {"id": "62d6c98a-cede-45b4-9993-c27955927d58", "name": "Bursite"},
    {"id": "35225e38-b0d2-4636-a5e9-bd52087a3550", "name": "Tendinite"},
    {"id": "0846cfba-7290-4f85-b7b8-d265f252fb99", "name": "Dor lombar"},
    {"id": "e6f0b0eb-07de-469d-874f-ade1376eceea", "name": "Fibromialgia"},
    {"id": "9dec1d1b-e290-470d-b344-7b48d9562876", "name": "Espondilite anquilosante"},
    {"id": "fdeb45bc-f0c7-43f8-bc20-2c0c1b907f49", "name": "Gota"},
    {"id": "29b6d4a8-4fd3-4b94-ac6e-c404259efa2e", "name": "Lúpus"},
    {"id": "d77e86c7-61ff-4178-b19a-686763310801", "name": "Esclerose múltipla"},
    {"id": "4753fb06-160c-431b-a1dd-edf619700cca", "name": "Doença de Crohn"}
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
            print("✓ Login realizado com sucesso!")
            return True
        except Exception as e:
            print(f"✗ Erro no login: {e}")
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
            print(f"✗ Erro ao atualizar item {item_id}: {e}")
            return False

    def enrich_item(self, item: Dict) -> bool:
        """Enriquece um item específico"""
        print(f"\n{'='*80}")
        print(f"Processando: {item['name']}")
        print(f"ID: {item['id']}")
        print(f"{'='*80}\n")

        content = self.get_content_for_item(item)

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

        # 1. Antidiabéticos
        if item['id'] == "2aa595d8-a3e5-4f7f-964b-88f5c2451e79":
            return {
                "clinical_relevance": """O histórico de uso de antidiabéticos indica diagnóstico estabelecido de diabetes mellitus ou pré-diabetes com necessidade de controle farmacológico da glicemia. Na medicina funcional integrativa, o uso desses medicamentos é contextualizado dentro de estratégia terapêutica multimodal que inclui modificação intensiva do estilo de vida, otimização nutricional e correção de disfunções metabólicas subjacentes.

As classes principais incluem: metformina (primeira linha, melhora sensibilidade insulínica e pode ter efeitos anti-envelhecimento), sulfonilureias (estimulam secreção insulínica, risco de hipoglicemia), inibidores DPP-4, inibidores SGLT2 (benefícios cardiovasculares e renais), agonistas GLP-1 (perda de peso, proteção cardiovascular) e insulina (estágios avançados ou DM1).

O histórico de antidiabéticos fornece informações sobre: gravidade e duração do diabetes, grau de disfunção das células beta pancreáticas, presença de complicações, e resposta a terapias prévias. A abordagem funcional busca minimizar a necessidade medicamentosa através de intervenções no estilo de vida, correção de resistência insulínica, otimização da composição corporal e suporte nutricional direcionado, embora reconheça que farmacoterapia é frequentemente necessária e benéfica.""",

                "patient_explanation": """Usar medicamentos antidiabéticos significa que seu corpo precisa de ajuda para controlar os níveis de açúcar no sangue. Existem diferentes tipos de remédios para diabetes, cada um funcionando de maneira específica - alguns ajudam seu corpo a usar melhor a insulina, outros fazem o pâncreas produzir mais insulina, e alguns ajudam a eliminar o excesso de açúcar pela urina.

Na medicina funcional, entendemos que o medicamento é importante, mas não é suficiente sozinho. Trabalhamos para otimizar sua alimentação, atividade física, sono e estresse para que seu corpo funcione melhor e, em alguns casos, você possa reduzir a dose ou até parar o medicamento (sempre com orientação médica).

É fundamental nunca parar ou alterar seus medicamentos por conta própria. Mudanças devem ser feitas gradualmente, sob supervisão médica, conforme seu metabolismo melhora com as intervenções no estilo de vida.""",

                "conduct": """1. Revisão detalhada da terapia antidiabética atual:
   - Quais medicamentos, doses, horários
   - Adesão ao tratamento
   - Efeitos colaterais (hipoglicemias, desconforto gastrointestinal, ganho de peso)
   - Resposta terapêutica (HbA1c, glicemias)

2. Monitoramento glicêmico abrangente:
   - Hemoglobina glicada trimestral
   - Perfil glicêmico (jejum, pós-prandial)
   - Frutosamina quando indicado
   - Consideração de monitoramento contínuo de glicose (CGM)
   - Registro de hipoglicemias

3. Avaliação de deficiências induzidas por medicamentos:
   - Vitamina B12 (metformina causa má absorção)
   - Coenzima Q10 (se uso de estatinas concomitante)
   - Vitamina D e cálcio (se uso de tiazolidinedionas)

4. Otimização da terapia medicamentosa:
   - Priorização de agentes com benefícios cardiovasculares (iSGLT2, aGLP-1)
   - Minimização de medicamentos com risco de hipoglicemia
   - Simplificação de esquemas quando possível (adesão)
   - Individualização de metas conforme perfil do paciente

5. Intervenção nutricional intensiva para redução medicamentosa:
   - Controle de carga glicêmica
   - Dieta baixa em carboidratos refinados
   - Jejum intermitente supervisionado (ajustar medicação)
   - Monitoramento rigoroso durante mudanças dietéticas

6. Programa de exercícios estruturado:
   - Combinação aeróbico + resistido
   - Atenção especial a risco de hipoglicemia
   - Ajustes medicamentosos conforme atividade física

7. Suplementação funcional complementar:
   - Vitamina B12 (1000mcg/dia se uso de metformina)
   - Berberina (pode potencializar efeito - monitorar glicemia)
   - Cromo, ácido alfa-lipóico, magnésio
   - Canela, inositol

8. Educação do paciente:
   - Reconhecimento e manejo de hipoglicemias
   - Importância da adesão medicamentosa
   - Automonitoramento glicêmico
   - Comunicação com equipe sobre mudanças

9. Reavaliação periódica para ajuste de doses conforme melhora metabólica

10. Coordenação com endocrinologista para otimização terapêutica contínua"""
            }

        # 2. Antiosteoporóticos
        elif item['id'] == "791917c3-1c37-46fa-89da-079a8adb2bd7":
            return {
                "clinical_relevance": """O histórico de uso de antiosteoporóticos (bifosfonatos, denosumabe, teriparatida, raloxifeno) indica diagnóstico de osteoporose ou osteopenia com risco elevado de fraturas. Essas medicações visam reduzir reabsorção óssea (antirreabsortivos) ou estimular formação óssea (anabólicos), reduzindo risco de fraturas vertebrais, quadril e outros sítios.

Na medicina funcional integrativa, a saúde óssea é compreendida como resultado de múltiplos fatores: equilíbrio hormonal (estrogênio, testosterona, paratormônio, vitamina D), estado nutricional (cálcio, magnésio, vitamina K2, proteínas), carga mecânica (exercícios resistidos), inflamação sistêmica, saúde gastrointestinal (absorção de nutrientes) e eixo intestino-osso (microbiota).

O uso de antiosteoporóticos requer atenção a: duração do tratamento (bifosfonatos: geralmente 3-5 anos, "drug holiday" subsequente), efeitos adversos raros mas graves (osteonecrose de mandíbula, fraturas atípicas de fêmur), adequação de cálcio e vitamina D, saúde dentária. A abordagem funcional complementa a farmacoterapia com otimização nutricional, exercícios de impacto e resistência, equilíbrio hormonal e suporte à saúde óssea metabólica.""",

                "patient_explanation": """Usar medicamentos para osteoporose significa que seus ossos estão mais frágeis e com maior risco de fraturas. Esses remédios ajudam a fortalecer os ossos ao reduzir a perda óssea ou estimular a formação de osso novo.

Na medicina funcional, além do medicamento, trabalhamos para otimizar todos os fatores que afetam a saúde dos ossos: garantir níveis adequados de vitamina D e cálcio (mas também magnésio e vitamina K2, que são igualmente importantes), corrigir desequilíbrios hormonais, melhorar a absorção intestinal de nutrientes, e implementar exercícios de resistência e impacto que estimulam a formação óssea.

É importante fazer exames dentários regulares enquanto usa esses medicamentos e informar sempre seu dentista sobre o tratamento. Também é fundamental não confiar apenas no medicamento - a combinação com alimentação rica em nutrientes e exercícios adequados é essencial para ossos saudáveis.""",

                "conduct": """1. Revisão da terapia antiosteoporótica:
   - Medicamento em uso, dose, duração do tratamento
   - Adesão (bifosfonatos: importante tomar em jejum, posição vertical)
   - Efeitos adversos (esofagite, dor óssea/muscular)
   - Tempo de uso (considerar "drug holiday" após 3-5 anos de bifosfonatos)

2. Monitoramento de eficácia e segurança:
   - Densitometria óssea (DEXA) anual ou bianual
   - Marcadores de remodelação óssea (CTX, P1NP)
   - Cálcio, fósforo, fosfatase alcalina
   - Função renal (bifosfonatos contraindicados se TFG <35mL/min)

3. Avaliação odontológica:
   - Exame dentário antes de iniciar/durante tratamento
   - Procedimentos invasivos devem ser feitos antes de iniciar (quando possível)
   - Informar dentista sobre uso de antirreabsortivos

4. Otimização de vitamina D e cálcio:
   - Vitamina D: manter níveis >30ng/mL (idealmente 40-60ng/mL)
   - Cálcio: 1000-1200mg/dia (preferir fontes alimentares)
   - Suplementação quando ingestão dietética insuficiente

5. Avaliação e suplementação de cofatores essenciais:
   - Vitamina K2 (MK-7): 100-200mcg/dia (direciona cálcio para ossos)
   - Magnésio: 400-600mg/dia (essencial para saúde óssea)
   - Boro: 3mg/dia
   - Silício
   - Proteína adequada (1-1,2g/kg)

6. Programa de exercícios para saúde óssea:
   - Exercícios de impacto (caminhada, corrida leve, pular corda)
   - Treinamento de resistência/força (2-3x/semana)
   - Exercícios de equilíbrio (prevenção de quedas)

7. Avaliação e otimização hormonal:
   - Hormônios tireoidianos (hipertireoidismo acelera perda óssea)
   - Hormônios sexuais (estrogênio, testosterona)
   - Paratormônio (hiperparatireoidismo)
   - Cortisol (excesso causa perda óssea)

8. Prevenção de quedas:
   - Avaliação de risco de quedas
   - Revisão de medicamentos que aumentam risco (benzodiazepínicos, etc)
   - Adequação do ambiente domiciliar
   - Correção visual

9. Investigação e correção de causas secundárias:
   - Doença celíaca, má absorção
   - Hiperparatireoidismo primário
   - Deficiência de vitamina D
   - Uso crônico de corticoides
   - Hipogonadismo

10. Reavaliação periódica e decisão sobre continuidade ou pausa do tratamento farmacológico"""
            }

        # 3. Antiparkinsonianos
        elif item['id'] == "6dcbb3d9-611d-4c46-aa62-a9dbb13341dd":
            return {
                "clinical_relevance": """O histórico de uso de antiparkinsonianos indica diagnóstico de doença de Parkinson ou parkinsonismo secundário. A doença de Parkinson é distúrbio neurodegenerativo caracterizado por degeneração de neurônios dopaminérgicos da substância negra, manifestando-se com tremor de repouso, rigidez, bradicinesia e instabilidade postural.

As medicações incluem: levodopa/carbidopa (reposição dopaminérgica, gold standard), agonistas dopaminérgicos (pramipexol, ropinirol), inibidores MAO-B (selegilina, rasagilina - neuroproteção potencial), inibidores COMT (entacapona) e anticolinérgicos. O tratamento sintomático melhora qualidade de vida mas não altera progressão da doença.

Na medicina funcional integrativa, a doença de Parkinson é abordada considerando: estresse oxidativo intenso, disfunção mitocondrial neuronal, neuroinflamação, desequilíbrio do eixo intestino-cérebro (disbiose precede sintomas motores), acúmulo de alfa-sinucleína, deficiências nutricionais (CoQ10, vitaminas B). Estratégias complementares incluem otimização nutricional, suporte mitocondrial, modulação da neuroinflamação, exercícios físicos específicos e manejo de complicações não-motoras.""",

                "patient_explanation": """Usar medicamentos antiparkinsonianos significa que você tem doença de Parkinson ou sintomas semelhantes. Essa é uma condição neurológica em que uma parte do cérebro produz menos dopamina, um neurotransmissor importante para controlar os movimentos. Os remédios ajudam a repor ou potencializar a dopamina, melhorando tremores, rigidez e lentidão dos movimentos.

Na medicina funcional, além dos medicamentos essenciais prescritos pelo neurologista, trabalhamos para proteger os neurônios e melhorar a função cerebral através de: alimentação anti-inflamatória rica em antioxidantes, suporte à saúde intestinal (estudos mostram conexão forte entre intestino e Parkinson), suplementação com CoQ10 e outros nutrientes neuroprotetores, exercícios físicos regulares (muito importantes para desacelerar a progressão), e manejo de sintomas não-motores como constipação, depressão e distúrbios do sono.

É fundamental manter o tratamento medicamentoso rigorosamente conforme prescrito pelo neurologista e fazer acompanhamento regular.""",

                "conduct": """1. Acompanhamento neurológico especializado contínuo

2. Revisão da terapia antiparkinsoniana:
   - Medicamentos, doses, horários (levodopa requer horários precisos)
   - Eficácia no controle de sintomas motores
   - Complicações motoras (flutuações, discinesias)
   - Efeitos adversos (náuseas, hipotensão ortostática, sonolência, alucinações)

3. Avaliação de sintomas não-motores (frequentemente subdiagnosticados):
   - Constipação (presente anos antes dos sintomas motores)
   - Distúrbios do sono (insônia, REM behavior disorder)
   - Disfunção autonômica (hipotensão ortostática)
   - Sintomas neuropsiquiátricos (depressão, ansiedade, apatia)
   - Disfunção cognitiva

4. Otimização nutricional:
   - Proteína: distribuir ao longo do dia, evitar excesso em refeição com levodopa (compete absorção)
   - Hidratação adequada
   - Fibras para constipação (30-40g/dia)
   - Dieta mediterrânea (evidências de neuroproteção)

5. Suporte mitocondrial e antioxidante:
   - Coenzima Q10: 600-1200mg/dia (estudos mostram benefício em altas doses)
   - Vitamina E
   - NAC (N-acetilcisteína)
   - Glutationa
   - Ácido alfa-lipóico

6. Suporte à saúde intestinal:
   - Probióticos de cepas específicas
   - Prebióticos
   - Tratamento de disbiose e SIBO
   - Otimização da motilidade intestinal

7. Vitaminas B essenciais:
   - Complexo B (B1, B6, B12)
   - Folato (metilfolato)
   - Atenção: vitamina B6 em altas doses pode reduzir eficácia de levodopa isolada

8. Outros suplementos com evidências:
   - Creatina (3-5g/dia - suporte energético neuronal)
   - Curcumina
   - Ômega-3
   - Vitamina D (adequação de níveis)

9. Programa de exercícios terapêuticos:
   - Exercícios aeróbicos regulares (neuroproteção)
   - Treinamento de força
   - Exercícios de equilíbrio e marcha
   - Tai chi, dança, boxing for Parkinson's
   - Fisioterapia especializada

10. Terapias complementares:
    - Fonoaudiologia (voz, deglutição)
    - Terapia ocupacional
    - Suporte psicológico

11. Monitoramento e prevenção de complicações:
    - Quedas
    - Desnutrição
    - Depressão
    - Demência

12. Educação de paciente e cuidadores sobre progressão, manejo de sintomas e qualidade de vida"""
            }

        # 4. Antipsicóticos
        elif item['id'] == "3511d78f-cfc2-4d9b-bd1f-32ed3a3120ae":
            return {
                "clinical_relevance": """O histórico de uso de antipsicóticos indica diagnóstico de transtorno psiquiátrico grave (esquizofrenia, transtorno bipolar, depressão psicótica) ou uso off-label para insônia, agitação, náuseas. Os antipsicóticos dividem-se em típicos/primeira geração (haloperidol, clorpromazina - alto risco de sintomas extrapiramidais) e atípicos/segunda geração (risperidona, quetiapina, olanzapina - menor risco extrapiramidal, maior risco metabólico).

Estes medicamentos apresentam efeitos adversos significativos: síndrome metabólica (ganho de peso, resistência insulínica, dislipidemia, diabetes), sintomas extrapiramidais (parkinsonismo, acatisia, discinesia tardia), hiperprolactinemia, prolongamento QT, sedação. O manejo desses efeitos adversos é crucial para adesão e qualidade de vida.

Na medicina funcional integrativa, pacientes em uso de antipsicóticos requerem atenção especial a: monitoramento metabólico rigoroso, prevenção de ganho de peso e síndrome metabólica, correção de deficiências nutricionais (especialmente vitamina D, ômega-3), suporte à função mitocondrial, modulação da neuroinflamação. A abordagem complementar visa minimizar efeitos adversos metabólicos e otimizar saúde geral, sempre em coordenação com psiquiatra assistente.""",

                "patient_explanation": """Usar antipsicóticos significa que você está tratando uma condição psiquiátrica que requer medicação para estabilizar os sintomas. Esses remédios são muito importantes e eficazes, mas podem causar efeitos colaterais como ganho de peso, alterações no metabolismo do açúcar e colesterol, sonolência e às vezes movimentos involuntários.

Na medicina funcional, trabalhamos em conjunto com seu psiquiatra para minimizar esses efeitos colaterais através de: alimentação balanceada para evitar ganho de peso excessivo, exercícios regulares, suplementação com nutrientes que protegem o cérebro e o metabolismo, e monitoramento rigoroso de peso, glicemia e colesterol.

É absolutamente fundamental NUNCA parar ou reduzir seu antipsicótico por conta própria, pois isso pode causar recaída grave da doença. Qualquer ajuste deve ser feito apenas pelo psiquiatra.""",

                "conduct": """1. Coordenação estreita com psiquiatra assistente

2. Revisão da terapia antipsicótica:
   - Medicamento, dose, adesão
   - Controle de sintomas psiquiátricos
   - Efeitos adversos (ganho de peso, sedação, sintomas extrapiramidais)
   - Consideração de troca para agente com melhor perfil metabólico

3. Monitoramento metabólico rigoroso:
   - Peso, IMC, circunferência abdominal (mensal inicialmente)
   - Glicemia de jejum, HbA1c (trimestral)
   - Perfil lipídico (trimestral no primeiro ano, depois semestral)
   - Pressão arterial
   - Prolactina (se sintomas de hiperprolactinemia)
   - ECG (QTc se uso de ziprasidona, outros conforme indicação)

4. Prevenção e manejo de ganho de peso:
   - Intervenção nutricional precoce e intensiva
   - Controle calórico moderado
   - Dieta de baixa carga glicêmica
   - Evitar bebidas calóricas
   - Educação sobre controle de porções

5. Programa de exercícios estruturado:
   - Atividade física regular (mínimo 150min/semana)
   - Treinamento de resistência
   - Benefícios duais: controle de peso + melhora de sintomas psiquiátricos

6. Suporte nutricional e suplementação:
   - Ômega-3: 2-3g/dia (evidências em esquizofrenia)
   - Vitamina D: adequação de níveis (deficiência comum)
   - Complexo B (especialmente B12, folato)
   - NAC (N-acetilcisteína): 1000-2000mg/dia (evidências emergentes)
   - Antioxidantes (vitamina C, E)

7. Manejo de resistência insulínica:
   - Metformina quando indicado
   - Berberina, cromo, ácido alfa-lipóico
   - Controle de carboidratos refinados

8. Avaliação e manejo de sintomas extrapiramidais:
   - Escalas de avaliação (AIMS para discinesia tardia)
   - Redução de dose quando possível
   - Troca para atípico se uso de típico
   - Vitamina E (discinesia tardia - evidências limitadas)

9. Suporte à saúde mitocondrial e neuroproteção:
   - Coenzima Q10
   - L-carnitina
   - Creatina
   - Magnésio

10. Monitoramento de complicações:
    - Síndrome metabólica
    - Diabetes tipo 2
    - Doença cardiovascular
    - Discinesia tardia

11. Atenção a interações medicamentosas (especialmente se polifarmácia)

12. Suporte psicossocial contínuo e educação sobre importância da adesão"""
            }

        # 5. Antivirais
        elif item['id'] == "9323cdf9-9b56-4638-b170-da2008d298b5":
            return {
                "clinical_relevance": """O histórico de uso de antivirais indica infecção viral aguda ou crônica requerendo terapia específica. As principais indicações incluem: HIV (terapia antirretroviral - TARV), hepatites B e C (análogos nucleos(t)ídicos, DAAs), herpes vírus (aciclovir, valaciclovir), influenza (oseltamivir), CMV (ganciclovir), e mais recentemente COVID-19.

Na era moderna, antivirais transformaram doenças anteriormente fatais em condições manejáveis cronicamente (HIV) ou até curáveis (hepatite C com >95% de cura). A TARV para HIV alcançou formulações de comprimido único diário com excelente perfil de tolerabilidade, permitindo expectativa de vida próxima ao normal quando carga viral indetectável. Hepatite C é curável em 8-12 semanas com DAAs. Hepatite B é suprimível a longo prazo.

Na medicina funcional integrativa, pacientes em terapia antiviral crônica requerem: monitoramento de efeitos adversos medicamentosos (disfunção mitocondrial, dislipidemias, nefrotoxicidade, hepatotoxicidade), suporte nutricional adequado, otimização da função imunológica, prevenção de comorbidades metabólicas e cardiovasculares (particularmente em HIV), e atenção à saúde mental. A abordagem complementar visa maximizar qualidade de vida, minimizar efeitos adversos e reduzir comorbidades não-relacionadas à infecção.""",

                "patient_explanation": """Usar antivirais significa que você está tratando uma infecção por vírus que requer medicação específica. Dependendo do vírus, o tratamento pode ser de curta duração (como para gripe ou herpes) ou de longo prazo (como para HIV ou hepatite B).

Os antivirais modernos são muito eficazes. Por exemplo, no caso do HIV, tomar a medicação corretamente mantém o vírus indetectável no sangue, preserva a imunidade e permite vida normal e longa. Na hepatite C, o tratamento cura a infecção em mais de 95% dos casos.

Na medicina funcional, além de garantir que você tome os antivirais corretamente conforme prescrito, trabalhamos para: fortalecer seu sistema imunológico através de alimentação adequada e suplementação, prevenir efeitos colaterais dos medicamentos, monitorar sua saúde metabólica e cardiovascular, e otimizar sua qualidade de vida geral.""",

                "conduct": """1. Identificação do agente viral e regime antiviral específico

2. Para HIV - Terapia Antirretroviral (TARV):
   - Adesão >95% essencial (doses perdidas levam a resistência)
   - Monitoramento: CD4, carga viral (objetivo: indetectável)
   - Rastreamento de comorbidades não-AIDS (cardiovascular, renal, hepática, óssea)
   - Lipodistrofia, dislipidemia, resistência insulínica
   - Suplementação: vitamina D, ômega-3, probióticos
   - Vacinações

3. Para Hepatite B crônica:
   - Antivirais: tenofovir ou entecavir (supressão viral)
   - Monitoramento: HBV-DNA, transaminases, AFP
   - Rastreamento hepatocarcinoma (ultrassom + AFP semestral)
   - Elastografia hepática
   - Evitar álcool absolutamente

4. Para Hepatite C:
   - DAAs: 8-12 semanas, taxa de cura >95%
   - RVS (resposta virológica sustentada) = cura
   - Pós-cura: se cirrose, manter rastreamento hepatocarcinoma
   - Suporte hepático: silimarina, NAC

5. Monitoramento de toxicidade medicamentosa:
   - Função renal (tenofovir pode ser nefrotóxico)
   - Função hepática
   - Perfil lipídico (inibidores de protease HIV)
   - Densidade óssea (tenofovir)
   - Lactato (análogos nucleosídicos - disfunção mitocondrial rara)

6. Suporte nutricional e suplementação:
   - Dieta anti-inflamatória rica em antioxidantes
   - Vitamina D: adequação de níveis (imunomodulação)
   - Ômega-3: 2-3g/dia (anti-inflamatório)
   - Probióticos de cepas específicas (saúde intestinal/imune)
   - NAC, glutationa (suporte hepático e antioxidante)
   - Vitamina C, zinco, selênio (função imunológica)

7. Suporte mitocondrial (especialmente análogos nucleosídicos):
   - Coenzima Q10
   - L-carnitina
   - Ácido alfa-lipóico
   - Complexo B

8. Prevenção de comorbidades metabólicas:
   - Controle de peso
   - Exercícios regulares
   - Evitar tabagismo e excesso de álcool
   - Manejo de resistência insulínica e dislipidemia

9. Saúde mental e qualidade de vida:
   - Rastreamento de depressão e ansiedade
   - Suporte psicossocial
   - Grupos de apoio quando apropriado

10. Educação sobre:
    - Importância da adesão estrita
    - Prevenção de transmissão
    - Indetectável = Intransmissível (I=I) para HIV
    - Interações medicamentosas

11. Vacinações conforme status imunológico

12. Acompanhamento especializado regular (infectologia)"""
            }

        # 6. GLP-1 (Agonistas do receptor de GLP-1)
        elif item['id'] == "9d6ff60e-000b-493d-bb4e-43b4db957188":
            return {
                "clinical_relevance": """O histórico de uso de agonistas do receptor de GLP-1 (liraglutida, semaglutida, dulaglutida, tirzepatida) indica diagnóstico de diabetes tipo 2 e/ou obesidade. Estes medicamentos revolucionaram o tratamento ao oferecer controle glicêmico superior, perda de peso substancial (10-15% com semaglutida, até 22% com tirzepatida) e benefícios cardiovasculares comprovados em estudos de desfechos (redução de eventos cardiovasculares maiores).

Os aGLP-1 mimetizam o hormônio incretina GLP-1, estimulando secreção insulínica glicose-dependente, suprimindo glucagon, retardando esvaziamento gástrico e reduzindo apetite. Seu uso expandiu além do diabetes para obesidade primária (indicação aprovada), doença hepática gordurosa não alcoólica, e potencialmente outras condições metabólicas.

Na medicina funcional integrativa, aGLP-1 são reconhecidos como ferramentas valiosas quando intervenções no estilo de vida isoladas são insuficientes. Entretanto, a abordagem funcional enfatiza: os medicamentos não substituem mudanças fundamentais na alimentação e atividade física; a perda de peso deve incluir preservação de massa muscular; efeitos colaterais gastrointestinais podem ser minimizados com progressão lenta de dose; e a descontinuação frequentemente leva a reganho de peso, exigindo estratégias de manutenção sustentáveis.""",

                "patient_explanation": """Os medicamentos GLP-1 (como Ozempic, Wegovy, Mounjaro) são injeções semanais ou diárias que ajudam no controle do diabetes e na perda de peso. Eles funcionam imitando um hormônio natural do seu corpo que: aumenta a produção de insulina quando você come, reduz o apetite, faz você se sentir satisfeito mais rapidamente e retarda a digestão.

Esses remédios têm mostrado resultados impressionantes de perda de peso (10-20kg ou mais) e também protegem o coração em pessoas com diabetes. No entanto, podem causar náuseas, vômitos e outros desconfortos gastrointestinais, especialmente no início.

Na medicina funcional, entendemos que esses medicamentos são ferramentas poderosas, mas não são mágicos nem definitivos. Você ainda precisa fazer mudanças na alimentação e aumentar a atividade física - o remédio facilita essas mudanças ao reduzir a fome. Além disso, é importante preservar massa muscular durante a perda de peso (através de exercícios de resistência e proteína adequada), pois parte do peso perdido pode ser músculo, não apenas gordura.""",

                "conduct": """1. Revisão do tratamento com aGLP-1:
   - Medicamento específico (liraglutida, semaglutida, dulaglutida, tirzepatida)
   - Dose atual (titulação deve ser lenta para minimizar efeitos GI)
   - Via de administração, frequência
   - Adesão, técnica de injeção
   - Efeitos colaterais (náuseas, vômitos, diarreia, constipação, gastroparesia)
   - Resposta terapêutica (HbA1c, peso)

2. Monitoramento de eficácia:
   - Peso corporal (semanal/quinzenal)
   - Hemoglobina glicada (trimestral)
   - Glicemias (jejum e pós-prandial)
   - Redução de outros antidiabéticos conforme melhora glicêmica

3. Avaliação de composição corporal:
   - Bioimpedância ou DEXA
   - Objetivo: maximizar perda de gordura, preservar massa muscular
   - Perda de peso rápida pode incluir 20-40% de massa magra

4. Otimização nutricional durante uso de aGLP-1:
   - Proteína adequada: 1,6-2g/kg peso ideal (preservar massa muscular)
   - Refeições menores e mais frequentes (melhor tolerabilidade)
   - Evitar alimentos gordurosos, muito condimentados
   - Hidratação adequada
   - Suplementação de multivitamínico/mineral

5. Programa de exercícios essencial:
   - Treinamento de resistência/força (3-4x/semana) - CRÍTICO para preservar músculo
   - Atividade aeróbica moderada
   - Progressão gradual conforme perda de peso

6. Manejo de efeitos colaterais gastrointestinais:
   - Titulação lenta de dose
   - Gengibre (náuseas)
   - Refeições pequenas, evitar deitar após comer
   - Antiemético quando necessário
   - Fibras e hidratação (constipação)

7. Monitoramento de segurança:
   - Amilase/lipase (pancreatite rara mas grave)
   - Frequência cardíaca (pode aumentar)
   - Sintomas de gastroparesia
   - Sintomas biliares (risco de colelitíase com perda de peso rápida)
   - Monitoramento de hipoglicemia (se uso concomitante de sulfonilureia/insulina)

8. Avaliação de risco cardiovascular e renal:
   - aGLP-1 têm benefícios cardiovasculares comprovados (semaglutida, liraglutida)
   - Benefícios renais (redução de albuminúria, progressão de DRC)

9. Suplementação durante perda de peso:
   - Multivitamínico/mineral completo
   - Cálcio, vitamina D
   - Vitamina B12
   - Ômega-3
   - Probióticos

10. Planejamento de manutenção a longo prazo:
    - Medicamento geralmente requer uso contínuo
    - Descontinuação leva a reganho de peso na maioria
    - Desenvolver hábitos sustentáveis de alimentação e exercício
    - Considerar dose de manutenção mínima efetiva

11. Educação do paciente:
    - Expectativas realistas (perda gradual de peso)
    - Importância de mudanças no estilo de vida concomitantes
    - Reconhecimento de efeitos adversos graves (pancreatite)
    - Custo (medicamentos caros, planejar sustentabilidade)

12. Coordenação com endocrinologista para otimização terapêutica"""
            }

        # Items 7-20: Conteúdo conciso mas clinicamente relevante

        # 7. Arritmia (duplicata - contexto medicação)
        elif item['id'] == "e04bd4fa-2076-4068-9a9d-ab0e9c1d16a2":
            return {
                "clinical_relevance": """Uso de medicações antiarrítmicas ou anticoagulantes indica histórico de arritmia significativa. Fibrilação atrial requer anticoagulação (warfarina, NOACs) para prevenção de AVC. Antiarrítmicos (betabloqueadores, amiodarona, flecainida) controlam ritmo/frequência. Requer monitoramento de eficácia, toxicidade e adesão.""",
                "patient_explanation": """Medicamentos para arritmia controlam batimentos irregulares do coração. Alguns controlam a velocidade ou ritmo, outros previnem formação de coágulos que poderiam causar derrame. É fundamental tomar rigorosamente conforme prescrito e fazer acompanhamento regular com cardiologista.""",
                "conduct": """1. Acompanhamento cardiológico 2. Monitoramento com ECG/Holter 3. INR se warfarina (meta 2-3 para FA) 4. Função renal e hepática (NOACs) 5. Avaliação de eletrólitos (magnésio, potássio) 6. Suplementação: magnésio, CoQ10, ômega-3, taurina 7. Evitar álcool e cafeína em excesso"""
            }

        # 8. Artrite (genérica)
        elif item['id'] == "78c36d3f-b74e-436d-a224-3d66f3a5d7c1":
            return {
                "clinical_relevance": """Artrite refere-se a inflamação articular de múltiplas etiologias (autoimune, degenerativa, infecciosa, metabólica). Manifestações incluem dor, edema, rigidez matinal e limitação funcional. Na medicina funcional, artrites são compreendidas como processos inflamatórios sistêmicos influenciados por fatores dietéticos, microbiota intestinal, permeabilidade intestinal, estresse oxidativo e desequilíbrios imunológicos. Abordagem anti-inflamatória multimodal pode reduzir sintomas e progressão.""",
                "patient_explanation": """Artrite significa inflamação nas articulações (juntas), causando dor, inchaço e dificuldade de movimento. Existem vários tipos com causas diferentes. Na medicina funcional, além dos tratamentos convencionais, usamos alimentação anti-inflamatória, suplementação (ômega-3, cúrcuma), exercícios apropriados e cuidados com a saúde intestinal para reduzir a inflamação.""",
                "conduct": """1. Diagnóstico etiológico específico (AR vs osteoartrite vs outras) 2. Marcadores inflamatórios (PCR, VHS) 3. Dieta anti-inflamatória (eliminar processados, açúcar, glúten trial) 4. Ômega-3 (2-4g/dia) 5. Cúrcuma/curcumina 6. Vitamina D adequada 7. Suporte à saúde intestinal 8. Exercícios de baixo impacto 9. Fisioterapia"""
            }

        # 9. Artrite reumatoide
        elif item['id'] == "d61bf844-b1a2-4096-b954-53031694c590":
            return {
                "clinical_relevance": """Artrite reumatoide (AR) é doença autoimune sistêmica caracterizada por poliartrite simétrica erosiva, podendo causar deformidades e incapacidade funcional. Manifestações extra-articulares incluem nódulos reumatoides, vasculite, doença pulmonar intersticial e risco cardiovascular aumentado. O tratamento precoce e agressivo com DMARDs (metotrexato, biológicos) é fundamental para prevenir danos articulares irreversíveis. Na medicina funcional, atenção especial à saúde intestinal (permeabilidade, disbiose), modulação imunológica nutricional e redução de gatilhos inflamatórios.""",
                "patient_explanation": """Artrite reumatoide é uma doença autoimune em que o sistema imunológico ataca as próprias articulações, causando inflamação, dor e, se não tratada, deformidades. O tratamento médico com medicamentos imunossupressores é essencial. Na medicina funcional, complementamos com dieta anti-inflamatória, cuidados intestinais, suplementação e técnicas de redução de estresse para controlar melhor a doença.""",
                "conduct": """1. Acompanhamento reumatológico 2. DMARDs (metotrexato +/- biológicos) 3. Monitoramento: PCR, VHS, FR, anti-CCP 4. Ácido fólico (se metotrexato) 5. Dieta anti-inflamatória estrita 6. Ômega-3 (3-4g/dia) 7. Vitamina D (>40ng/mL) 8. Probióticos 9. Cúrcuma 10. Evitar glúten e laticínios (trial) 11. Exercícios preservando articulações"""
            }

        # 10. Osteoartrite
        elif item['id'] == "47ca7e29-247c-4dbb-a1a8-ae723ade3004":
            return {
                "clinical_relevance": """Osteoartrite (OA) é doença articular degenerativa mais comum, caracterizada por degradação progressiva da cartilagem articular, esclerose subcondral e formação de osteófitos. Afeta principalmente articulações de carga (joelhos, quadris, coluna) e mãos. Fatores de risco incluem idade, obesidade, trauma prévio, desalinhamento articular. Na medicina funcional, OA é vista como processo inflamatório de baixo grau modulável através de perda de peso, suplementação condroprotora, anti-inflamatórios naturais e otimização biomecânica.""",
                "patient_explanation": """Osteoartrite é o "desgaste" da cartilagem das articulações que acontece com a idade, uso repetitivo ou sobrepeso. Causa dor, rigidez e limitação dos movimentos. Na medicina funcional, trabalhamos para: reduzir peso (alivia carga nas articulações), fortalecer músculos ao redor, suplementar com glucosamina e condroitina, usar anti-inflamatórios naturais (cúrcuma, ômega-3) e melhorar a biomecânica do movimento.""",
                "conduct": """1. Avaliação clínica e radiológica 2. Perda de peso (cada 1kg reduz 4kg de carga no joelho) 3. Exercícios de fortalecimento muscular e baixo impacto 4. Fisioterapia 5. Glucosamina (1500mg/dia) + Condroitina (1200mg/dia) 6. Colágeno tipo II não desnaturado (UC-II 40mg/dia) 7. Ômega-3, cúrcuma, boswellia 8. Vitamina D adequada 9. Analgesia quando necessário 10. Considerar infiltrações (ácido hialurônico, PRP)"""
            }

        # 11. Outras artrites
        elif item['id'] == "5cb9c7b1-fb72-487b-ab16-51343f66bc18":
            return {
                "clinical_relevance": """Outras artrites incluem artrite psoriásica, espondiloartropatias, artrite por cristais (gota, pseudogota), artrite reativa, artrite infecciosa. Cada uma requer diagnóstico específico e tratamento direcionado. Abordagem funcional adapta intervenções nutricionais e suplementares conforme etiologia.""",
                "patient_explanation": """Existem vários tipos de artrite além das mais comuns. Cada tipo tem características próprias e tratamento específico. O importante é ter diagnóstico correto através do reumatologista e complementar com medidas anti-inflamatórias naturais apropriadas.""",
                "conduct": """1. Diagnóstico etiológico preciso 2. Tratamento específico conforme causa 3. Dieta anti-inflamatória 4. Suplementação anti-inflamatória 5. Exercícios apropriados 6. Acompanhamento especializado"""
            }

        # 12. Bursite
        elif item['id'] == "62d6c98a-cede-45b4-9993-c27955927d58":
            return {
                "clinical_relevance": """Bursite é inflamação da bursa sinovial (saco com líquido que reduz atrito entre tendões, músculos e ossos). Causas: trauma repetitivo, sobrecarga, trauma direto, infecção, artrite. Locais comuns: ombro, quadril, cotovelo, joelho. Tratamento: repouso relativo, gelo, anti-inflamatórios, fisioterapia. Na medicina funcional, foco em correção biomecânica, modulação anti-inflamatória e suporte à recuperação tecidual.""",
                "patient_explanation": """Bursite é inflamação de pequenas bolsas de líquido que protegem suas articulações. Geralmente causada por movimentos repetitivos ou sobrecarga. Tratamento: repouso da articulação afetada, gelo, anti-inflamatórios e fisioterapia. Suplementação com ômega-3, cúrcuma e proteolíticos pode ajudar.""",
                "conduct": """1. Repouso relativo da articulação 2. Crioterapia 3. Fisioterapia 4. Correção biomecânica 5. Anti-inflamatórios naturais (ômega-3, cúrcuma, bromelina) 6. Infiltração se necessário 7. Prevenir recorrência"""
            }

        # 13. Tendinite
        elif item['id'] == "35225e38-b0d2-4636-a5e9-bd52087a3550":
            return {
                "clinical_relevance": """Tendinite é inflamação/degeneração de tendões por sobrecarga, trauma repetitivo ou biomecânica inadequada. Locais comuns: ombro (manguito rotador), cotovelo (epicondilite), punho, aquileu. Tratamento: repouso, modificação de atividade, fisioterapia, fortalecimento progressivo. Na medicina funcional, suporte com nutrientes para síntese de colágeno (vitamina C, cobre, manganês), anti-inflamatórios naturais e correção de deficiências.""",
                "patient_explanation": """Tendinite é inflamação dos tendões (estruturas que conectam músculos aos ossos) geralmente por uso excessivo ou movimentos repetitivos. Tratamento: repouso, gelo, fisioterapia e fortalecimento gradual. Suplementação com colágeno, vitamina C e anti-inflamatórios naturais pode acelerar recuperação.""",
                "conduct": """1. Repouso relativo, modificação de atividade 2. Fisioterapia (fortalecimento excêntrico) 3. Colágeno hidrolisado (10-15g/dia) 4. Vitamina C (1-2g/dia) 5. Ômega-3, cúrcuma 6. Enzimas proteolíticas (bromelina, papaína) 7. Correção biomecânica"""
            }

        # 14. Dor lombar
        elif item['id'] == "0846cfba-7290-4f85-b7b8-d265f252fb99":
            return {
                "clinical_relevance": """Dor lombar é extremamente prevalente (80% da população em algum momento da vida). Causas: mecânicas (tensão muscular, hérnia discal, estenose, espondilolistese), inflamatórias, infecciosas, neoplásicas, referidas. Maioria é inespecífica/mecânica, autolimitada em 4-6 semanas. Bandeiras vermelhas (déficit neurológico, síndrome de cauda equina, febre, perda de peso) exigem investigação urgente. Tratamento: manutenção de atividade, fisioterapia, exercícios de fortalecimento core. Na medicina funcional, avaliação postural, biomecânica, inflamação sistêmica.""",
                "patient_explanation": """Dor lombar (dor nas costas) é muito comum e geralmente melhora em semanas. Causas: má postura, fraqueza muscular, sobrecarga, hérnia de disco. Importante: manter-se ativo (repouso prolongado piora), fazer fisioterapia, fortalecer abdômen e costas, corrigir postura. Sinais de alerta: fraqueza nas pernas, perda de controle da bexiga, febre.""",
                "conduct": """1. Avaliação clínica (bandeiras vermelhas) 2. Imagem se indicação específica 3. Manter atividade, evitar repouso 4. Fisioterapia, exercícios de core 5. Quiropraxia/osteopatia 6. Magnésio, cúrcuma 7. Peso saudável 8. Ergonomia 9. Técnicas de relaxamento"""
            }

        # 15. Fibromialgia
        elif item['id'] == "e6f0b0eb-07de-469d-874f-ade1376eceea":
            return {
                "clinical_relevance": """Fibromialgia é síndrome de dor crônica generalizada com sensibilização central do sistema nervoso, fadiga, distúrbios do sono, disfunção cognitiva ("fibrofog") e sintomas somáticos variados. Fisiopatologia envolve processamento anormal da dor, neuroinflamação, disfunção mitocondrial, alterações neurotransmissoras (serotonina, noradrenalina), eixo HPA. Diagnóstico clínico (critérios ACR). Tratamento multimodal: exercícios aeróbicos graduais, terapia cognitivo-comportamental, pregabalina/duloxetina. Na medicina funcional, foco em sono, estresse, inflamação sistêmica, deficiências nutricionais, saúde intestinal.""",
                "patient_explanation": """Fibromialgia é dor generalizada crônica no corpo todo, acompanhada de cansaço, sono ruim e dificuldade de concentração. O sistema nervoso fica "supersensível" à dor. Tratamento combina: exercícios leves regulares, melhora do sono, manejo de estresse, medicamentos quando necessário. Na medicina funcional, corrigimos deficiências de vitamina D, magnésio, coenzima Q10, melhoramos sono e reduzimos inflamação.""",
                "conduct": """1. Diagnóstico clínico (critérios ACR) 2. Excluir outras causas 3. Exercícios aeróbicos graduais (essencial) 4. Higiene do sono rigorosa 5. TCC, mindfulness 6. Vitamina D (>40ng/mL) 7. Magnésio 8. CoQ10 9. 5-HTP ou SAMe 10. Ômega-3 11. Dieta anti-inflamatória 12. Pregabalina ou duloxetina se necessário"""
            }

        # 16. Espondilite anquilosante
        elif item['id'] == "9dec1d1b-e290-470d-b344-7b48d9562876":
            return {
                "clinical_relevance": """Espondilite anquilosante (EA) é espondiloartropatia inflamatória crônica afetando coluna axial e articulações sacroilíacas, podendo levar a fusão vertebral ("coluna em bambu"). Associada a HLA-B27 (90%). Manifestações: dor lombar inflamatória (melhora com atividade, piora com repouso), rigidez matinal, redução de expansibilidade torácica. Manifestações extra-articulares: uveíte, doença inflamatória intestinal, psoriase. Tratamento: AINEs contínuos, anti-TNF (biológicos), exercícios (preservar mobilidade). Na medicina funcional, dieta baixa em amido (protocolo London AS), saúde intestinal, anti-inflamatórios naturais.""",
                "patient_explanation": """Espondilite anquilosante é doença inflamatória crônica que afeta principalmente a coluna e a bacia, podendo causar rigidez progressiva. A dor piora com repouso e melhora com movimento. Tratamento: medicamentos anti-inflamatórios ou biológicos, exercícios diários (essenciais para manter mobilidade), fisioterapia. Dieta baixa em amido pode ajudar.""",
                "conduct": """1. Acompanhamento reumatológico 2. Anti-TNF se AINE insuficiente 3. Exercícios diários (mobilidade, alongamento) 4. Fisioterapia 5. Dieta baixa em amido (London AS Diet) 6. Ômega-3 7. Vitamina D 8. Probióticos 9. Cúrcuma 10. Postura adequada"""
            }

        # 17. Gota
        elif item['id'] == "fdeb45bc-f0c7-43f8-bc20-2c0c1b907f49":
            return {
                "clinical_relevance": """Gota é artrite inflamatória por deposição de cristais de urato monossódico resultante de hiperuricemia crônica (>6,8mg/dL). Manifestações: artrite aguda (podagra - 1ª metatarsofalangiana), tofos, nefrolitíase. Associada a síndrome metabólica, obesidade, hipertensão, doença renal. Tratamento agudo: colchicina, AINEs, corticoides. Profilaxia: alopurinol, febuxostat (meta: ácido úrico <6mg/dL). Na medicina funcional, foco em: dieta pobre em purinas, redução de frutose, perda de peso, hidratação, cereja tart (reduz ataques), controle de síndrome metabólica.""",
                "patient_explanation": """Gota é causada por excesso de ácido úrico no sangue que forma cristais nas articulações, causando dor intensa, vermelhidão e inchaço (geralmente no dedão do pé). Relacionada a alimentação rica em carnes, frutos do mar, bebidas alcoólicas e frutose. Tratamento: remédios para baixar ácido úrico, mudanças na dieta, muito líquido, perder peso. Suco de cereja ajuda a prevenir crises.""",
                "conduct": """1. Diagnóstico: ácido úrico, artrocentese (cristais) 2. Fase aguda: colchicina, AINE, corticoide 3. Profilaxia: alopurinol ou febuxostat (meta AU <6mg/dL) 4. Dieta: reduzir carnes vermelhas, vísceras, frutos do mar, álcool, frutose 5. Hidratação (2-3L/dia) 6. Perda de peso 7. Cereja tart ou extrato 8. Vitamina C 9. Controle de síndrome metabólica"""
            }

        # 18. Lúpus
        elif item['id'] == "29b6d4a8-4fd3-4b94-ac6e-c404259efa2e":
            return {
                "clinical_relevance": """Lúpus eritematoso sistêmico (LES) é doença autoimune multissistêmica caracterizada por produção de autoanticorpos (anti-DNA, anti-Sm, anti-fosfolípides) com manifestações variadas: rash malar, fotossensibilidade, artrite, nefrite, serosite, citopenias, manifestações neuropsiquiátricas. Curso recorrente-remitente. Tratamento: hidroxicloroquina (base), corticoides, imunossupressores (MMF, azatioprina), biológicos (belimumabe). Na medicina funcional, modulação imunológica através de vitamina D, ômega-3, antioxidantes, saúde intestinal, redução de estresse, proteção solar rigorosa.""",
                "patient_explanation": """Lúpus é doença autoimune em que o sistema imunológico ataca vários órgãos do corpo. Sintomas variam: rash facial em borboleta, dor nas juntas, fadiga, febre, problemas nos rins. Requer acompanhamento médico rigoroso com medicamentos imunossupressores. Proteção solar é essencial (sol piora a doença). Na medicina funcional, complementamos com suplementação anti-inflamatória, vitamina D e cuidados com estresse.""",
                "conduct": """1. Acompanhamento reumatológico 2. Hidroxicloroquina (base) 3. Imunossupressores conforme manifestações 4. Proteção solar rigorosa (FPS 50+) 5. Vitamina D (>40ng/mL - imunomodulação) 6. Ômega-3 (3-4g/dia) 7. DHEA (25-50mg/dia) 8. Antioxidantes 9. Probióticos 10. Dieta anti-inflamatória 11. Manejo de estresse 12. Evitar sulfa, alfalfa"""
            }

        # 19. Esclerose múltipla
        elif item['id'] == "d77e86c7-61ff-4178-b19a-686763310801":
            return {
                "clinical_relevance": """Esclerose múltipla (EM) é doença autoimune desmielinizante do SNC com curso recorrente-remitente (mais comum), progressiva secundária ou primária progressiva. Manifestações: neurite óptica, sintomas motores/sensitivos, fadiga, disfunção vesical, espasticidade. Diagnóstico: critérios McDonald (clínica + RNM + LCR). Tratamento: drogas modificadoras de doença (interferons, fingolimod, natalizumabe, ocrelizumabe), corticoides para surtos. Na medicina funcional, evidências para: vitamina D (>40ng/mL), protocolo Wahls (dieta paleolítica modificada), ômega-3, controle de infecções (EBV?), saúde intestinal, redução de estresse.""",
                "patient_explanation": """Esclerose múltipla é doença autoimune em que o sistema imunológico ataca a bainha de mielina (proteção dos nervos) no cérebro e medula. Sintomas variam: dormência, fraqueza, problemas de visão, fadiga, dificuldade de equilíbrio. Tratamento com medicamentos imunossupressores é essencial. Na medicina funcional, complementamos com: vitamina D alta, dieta específica (Wahls Protocol), ômega-3, evitar glúten e laticínios, manejo de estresse.""",
                "conduct": """1. Acompanhamento neurológico 2. Drogas modificadoras de doença 3. Vitamina D (>40ng/mL, alguns protocolos usam 10.000UI/dia) 4. Ômega-3 (3-4g/dia) 5. Protocolo Wahls ou dieta paleolítica 6. Evitar glúten, laticínios (trial) 7. Probióticos 8. Ácido alfa-lipóico 9. Biotina (alta dose - evidências emergentes) 10. Exercícios regulares 11. Manejo de fadiga 12. Evitar superaquecimento"""
            }

        # 20. Doença de Crohn
        elif item['id'] == "4753fb06-160c-431b-a1dd-edf619700cca":
            return {
                "clinical_relevance": """Doença de Crohn (DC) é doença inflamatória intestinal transmural, podendo afetar qualquer segmento do TGI (mais comum: íleo terminal, cólon). Manifestações: dor abdominal, diarreia crônica (às vezes sanguinolenta), perda de peso, fístulas, abscessos, estenoses. Complicações: má absorção, deficiências nutricionais (B12, ferro, vitamina D, zinco), risco aumentado de câncer colorretal. Tratamento: 5-ASA, corticoides, imunossupressores (azatioprina, metotrexato), biológicos (anti-TNF). Na medicina funcional, foco em: modulação da microbiota, integridade da barreira intestinal, nutrição adequada (risco de desnutrição), identificação de gatilhos alimentares, omega-3, curcumina, probióticos específicos.""",
                "patient_explanation": """Doença de Crohn é inflamação crônica do intestino que causa dor abdominal, diarreia, perda de peso e cansaço. Pode ter períodos de crise e de melhora. O tratamento médico com imunossupressores é fundamental para controlar a inflamação e evitar complicações. Na medicina funcional, trabalhamos para: identificar alimentos que pioram (comuns: glúten, laticínios), melhorar a flora intestinal, corrigir deficiências nutricionais (você tem risco alto), reduzir inflamação com ômega-3 e cúrcuma.""",
                "conduct": """1. Acompanhamento gastroenterológico 2. Imunossupressores/biológicos conforme gravidade 3. Monitoramento: PCR, calprotectina fecal 4. Rastreamento de deficiências: ferro, B12, vitamina D, zinco, folato 5. Dieta: identificar gatilhos (glúten, laticínios, FODMAPs trial) 6. Dieta específica de carboidratos (SCD) ou LOFFLEX 7. Ômega-3 (2-4g/dia) 8. Curcumina (biodisponível) 9. Probióticos (VSL#3, outras cepas) 10. Glutamina, vitamina D 11. Suporte nutricional (risco desnutrição) 12. Manejo de estresse"""
            }

        # Fallback
        return {
            "clinical_relevance": f"Conteúdo clínico para {item['name']} em desenvolvimento.",
            "patient_explanation": f"{item['name']} requer acompanhamento médico especializado.",
            "conduct": "1. Avaliação especializada 2. Tratamento conforme protocolo 3. Monitoramento regular"
        }

    def process_all_items(self):
        """Processa todos os 20 items"""
        results = {
            "success": [],
            "failed": [],
            "total": len(ITEMS)
        }

        print(f"\n{'='*80}")
        print(f"ENRIQUECIMENTO - HISTÓRICO DE DOENÇAS PARTE 3")
        print(f"Iniciando processamento de {len(ITEMS)} items")
        print(f"Foco: Medicações e condições específicas")
        print(f"{'='*80}\n")

        for idx, item in enumerate(ITEMS, 1):
            print(f"\nItem {idx}/{len(ITEMS)}")
            success = self.enrich_item(item)

            if success:
                results["success"].append(item['name'])
            else:
                results["failed"].append(item['name'])

            time.sleep(0.5)

        # Sumário final
        print(f"\n{'='*80}")
        print("SUMÁRIO DO PROCESSAMENTO")
        print(f"{'='*80}")
        print(f"Total de items: {results['total']}")
        print(f"✓ Sucesso: {len(results['success'])}")
        print(f"✗ Falhas: {len(results['failed'])}")
        print(f"Taxa de sucesso: {len(results['success'])/results['total']*100:.1f}%")

        if results['failed']:
            print(f"\nItems com falha:")
            for name in results['failed']:
                print(f"  - {name}")

        print(f"{'='*80}\n")

        return results

def main():
    enricher = ScoreItemEnricher()

    if not enricher.login():
        print("✗ Erro: Não foi possível fazer login. Encerrando.")
        return

    results = enricher.process_all_items()

    # Salvar resultados
    output_file = '/home/user/plenya/disease_history_batch3_results.json'
    with open(output_file, 'w', encoding='utf-8') as f:
        json.dump(results, f, indent=2, ensure_ascii=False)

    print(f"\n✓ Resultados salvos em: {output_file}")

if __name__ == "__main__":
    main()
