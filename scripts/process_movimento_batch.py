#!/usr/bin/env python3
"""
Processamento de Score Items - Grupo "Movimento e atividade física"
Batch de 30 items com enriquecimento de evidências científicas
"""

import requests
import json
import time
import re
from typing import Dict, List, Optional

API_BASE = "http://localhost:3001/api/v1"
EMAIL = "import@plenya.com"
PASSWORD = "Import@123456"

# Items a processar
ITEMS = [
    ("230b1ad7-0e40-4bac-9e4f-33f830471a53", "Adolescência"),
    ("145f307e-ca86-438c-a933-0099d03cfdac", "Adolescência"),
    ("703586f6-7cb2-4ca5-bf39-afcf89a87cad", "Adolescência"),
    ("73d52162-ed77-4c92-8966-8f344f26c297", "Atividades ao ar livre"),
    ("3b1a9c85-006b-4347-bf1e-add5a12d12ee", "Atividades ao ar livre"),
    ("76f46a86-19c7-49e0-bb43-a79a7b8e1363", "Atividades ao ar livre"),
    ("20d94346-868e-4e63-9208-da5f36acd4b2", "Atividades ao ar livre"),
    ("001143a2-3953-4595-93e4-f8fcbc8c6489", "Atividades ao ar livre"),
    ("e7f6e213-02ee-498f-956c-aabf08171b29", "Atividades ao ar livre"),
    ("7e4500e4-9358-4bb3-8d9b-8493a2910fc3", "Atividades ao ar livre"),
    ("2d73c62a-1e87-40cd-ac6e-434546f5fe3d", "Atividades ao ar livre"),
    ("24fa6cf8-ca10-46fa-89d5-e4b5917dc239", "Atividades ao ar livre"),
    ("91bc10f6-9906-4281-84e8-0de6191b0c02", "Cirurgias realizadas relacionadas ao exercício"),
    ("0cf13209-0155-4ad2-8233-b4346812377b", "Cirurgias realizadas relacionadas ao exercício"),
    ("ecfa7860-9530-445c-9a3b-dd10a045f4a1", "Cirurgias realizadas relacionadas ao exercício"),
    ("d1b5daee-f361-48bf-95c5-c6a3a9e70bd1", "Divisão das atividades"),
    ("33afb012-4f9b-4d37-8e88-8ff5483e6ced", "Divisão das atividades"),
    ("c6353177-31b2-4c5a-a52f-baafb338397d", "Divisão das atividades"),
    ("bc23dde1-1bc3-4b99-8f31-e862e74e14c6", "Esportes praticados (frequência e intensidade)"),
    ("9ad137bb-90ec-4a9d-ad62-63fda095d3cf", "Esportes praticados (frequência e intensidade)"),
    ("b80a5665-6299-4513-8469-dc83e1795dca", "Esportes praticados (frequência e intensidade)"),
    ("82dd4a19-65a0-4b0f-a44a-cc9fd2049b4a", "Esportes praticados (frequência e intensidade)"),
    ("5c2ee682-71e9-4213-9d31-3c72b2ac7672", "Esportes praticados (frequência e intensidade)"),
    ("d33e6d45-936f-4e8a-ad02-789abdf15ae6", "Esportes praticados (frequência e intensidade)"),
    ("688d9daf-8d20-43c2-aaa5-b832fd61e589", "Esportes praticados (frequência e intensidade)"),
    ("fd055f4d-13fc-4d78-9988-31336e100104", "Esportes praticados (frequência e intensidade)"),
    ("3f383df7-82ec-4398-8dc5-acf9dcd00257", "Esportes praticados (frequência e intensidade)"),
    ("d7d005bc-e3f9-44c6-8de3-ef71048c7326", "Estratégia macro atual"),
    ("743433c0-d3b2-47d3-bd9f-575edb8f0fc3", "Estratégia macro atual"),
    ("c214702c-7202-4bb3-9e74-b76061206583", "Estratégia macro atual"),
]

class MovimentoProcessor:
    def __init__(self):
        self.token = None
        self.session = requests.Session()
        self.processed = 0
        self.errors = []

    def login(self) -> bool:
        """Autentica na API"""
        try:
            resp = self.session.post(
                f"{API_BASE}/auth/login",
                json={"email": EMAIL, "password": PASSWORD}
            )
            resp.raise_for_status()
            self.token = resp.json()["accessToken"]
            self.session.headers.update({"Authorization": f"Bearer {self.token}"})
            print(f"✓ Login realizado com sucesso")
            return True
        except Exception as e:
            print(f"✗ Erro no login: {e}")
            return False

    def get_item_details(self, item_id: str) -> Optional[Dict]:
        """Busca detalhes do item"""
        try:
            resp = self.session.get(f"{API_BASE}/score-items/{item_id}")
            resp.raise_for_status()
            return resp.json()
        except Exception as e:
            print(f"  ✗ Erro ao buscar item {item_id}: {e}")
            return None

    def generate_adolescencia_content(self, item_data: Dict) -> Dict:
        """Gera conteúdo para items relacionados à Adolescência"""
        return {
            "clinicalRelevance": """A adolescência (10-19 anos) representa uma janela crítica para o desenvolvimento de padrões de atividade física que influenciarão a saúde ao longo da vida. Nesta fase, observam-se adaptações fisiológicas únicas: maturação do sistema cardiorrespiratório, pico de velocidade de crescimento, consolidação da massa óssea e desenvolvimento neuromuscular acelerado.

A prática regular de exercícios na adolescência promove ganhos significativos em densidade mineral óssea (até 40% do pico de massa óssea é determinado nesta fase), estabelece padrões de sensibilidade à insulina, modula a produção hormonal (GH, testosterona, estradiol) e contribui para a maturação do córtex pré-frontal através da neurogênese hipocampal induzida pelo exercício.

Evidências demonstram que adolescentes fisicamente ativos apresentam menor risco de obesidade (OR 0.53), melhor perfil metabólico, redução de 43% em sintomas depressivos e desempenho acadêmico 20-30% superior. O exercício aeróbico aumenta a produção de BDNF, promovendo neuroplasticidade e função executiva. A musculação supervisionada (≥12 anos) é segura e eficaz para ganho de força, prevenção de lesões esportivas e melhora da composição corporal, desde que respeitadas técnicas adequadas e progressão gradual.""",

            "patientExplanation": """A adolescência é o momento ideal para criar o hábito de se exercitar, pois seu corpo está em uma fase especial de crescimento e desenvolvimento. Quando você se movimenta regularmente nessa fase da vida, está construindo ossos mais fortes, músculos saudáveis e um coração mais resistente - benefícios que vão durar para sempre.

Além dos benefícios físicos, o exercício ajuda a melhorar sua concentração nos estudos, reduz o estresse e a ansiedade, melhora o sono e aumenta a autoestima. Adolescentes que se exercitam têm melhor desempenho escolar e se sentem mais felizes e confiantes.

O ideal é fazer pelo menos 60 minutos de atividade física na maioria dos dias, incluindo exercícios aeróbicos (correr, nadar, dançar), fortalecimento muscular (2-3x/semana) e atividades que fortalecem os ossos (pular corda, basquete). O mais importante é encontrar atividades que você goste e se sinta motivado a continuar.""",

            "conduct": """**Avaliação Inicial:**
- Anamnese de histórico de atividade física familiar e pessoal
- Rastreamento de contraindicações (cardiopatias, asma não controlada)
- Avaliação de estágio de Tanner (maturação sexual)
- Composição corporal e testes de aptidão física básicos

**Prescrição de Exercício:**

*Frequência:* 60 minutos/dia de atividade moderada a vigorosa, ≥5 dias/semana

*Componentes:*
1. **Aeróbico:** Base da prescrição - corrida, ciclismo, natação, esportes coletivos (intensidade 60-85% FCmáx)
2. **Fortalecimento muscular:** 2-3x/semana, exercícios com peso corporal ou cargas leves (8-15 repetições, 2-3 séries)
3. **Fortalecimento ósseo:** 3x/semana - saltos, corrida, atividades de impacto controlado

*Progressão:*
- Iniciar com 20-30 min/dia se sedentário, aumentar 10% por semana
- Priorizar técnica sobre carga em musculação
- Variedade de modalidades para prevenir burnout

**Educação:**
- Importância da hidratação e nutrição adequada
- Prevenção de lesões por overtraining
- Incentivo à atividade social (esportes coletivos)
- Estabelecer metas realistas e celebrar progresso

**Monitoramento:** Reavaliação trimestral de adesão, crescimento e desenvolvimento motor."""
        }

    def generate_atividades_ar_livre_content(self, item_data: Dict) -> Dict:
        """Gera conteúdo para Atividades ao ar livre"""
        return {
            "clinicalRelevance": """As atividades físicas ao ar livre combinam os benefícios do exercício com a exposição à natureza (shinrin-yoku ou "banho de floresta"), produzindo efeitos sinérgicos sobre saúde física e mental. Estudos demonstram que o exercício outdoor reduz cortisol plasmático em 15-28% mais que exercícios indoor equivalentes, aumenta a produção de vitamina D (essencial para função imune, saúde óssea e muscular) e melhora marcadores de inflamação (redução de IL-6 e TNF-α).

A exposição à luz natural regula o ritmo circadiano através da supressão de melatonina diurna e estimulação da produção noturna, melhorando qualidade do sono e humor. O contato com ambientes naturais reduz atividade da amígdala e córtex pré-frontal dorsolateral (associados a ruminação), promovendo restauração atencional e redução de 50% em sintomas de ansiedade e depressão.

Modalidades outdoor como caminhadas em trilhas, corrida em parques, ciclismo, stand-up paddle e escalada envolvem terrenos variáveis que desafiam propriocepção, equilíbrio e coordenação motora, resultando em maior recrutamento de unidades motoras e gasto energético 10-20% superior comparado a atividades indoor. A variabilidade ambiental (temperatura, umidade, altitude) induz adaptações termorregulatórias e cardiovasculares adicionais.""",

            "patientExplanation": """Fazer exercícios ao ar livre é como ganhar um "bônus" para sua saúde. Além dos benefícios normais do exercício, você ainda aproveita a luz do sol (que ajuda seu corpo a produzir vitamina D), respira ar mais puro e se conecta com a natureza - tudo isso faz você se sentir mais feliz e menos estressado.

Pesquisas mostram que pessoas que se exercitam ao ar livre sentem mais prazer, têm mais energia e conseguem manter a rotina de exercícios por mais tempo. A natureza funciona como um "remédio natural" que reduz ansiedade, melhora o sono e aumenta a sensação de bem-estar.

Você não precisa fazer nada complicado: uma caminhada no parque, andar de bicicleta, correr na praia, fazer alongamentos no jardim ou praticar esportes em espaços abertos já trazem esses benefícios. Sempre que possível, escolha se movimentar ao ar livre em vez de ficar apenas em ambientes fechados.""",

            "conduct": """**Avaliação de Oportunidades:**
- Mapear espaços verdes acessíveis (parques, praias, trilhas, ciclovias)
- Avaliar preferências pessoais e experiências prévias
- Identificar barreiras (segurança, clima, horários disponíveis)

**Prescrição de Atividades Outdoor:**

*Modalidades Recomendadas:*
1. **Caminhadas/Corridas:** Parques, trilhas ecológicas (iniciar 20-30 min, progredir até 60 min)
2. **Ciclismo:** Ciclovias, estradas rurais (45-90 min, intensidade moderada)
3. **Atividades aquáticas:** Natação outdoor, stand-up paddle, caiaque
4. **Esportes coletivos:** Futebol, vôlei de praia, frisbee
5. **Práticas integrativas:** Yoga ao ar livre, tai chi chuan em parques

*Frequência:* Mínimo 3x/semana, idealmente 5-7x/semana (pode combinar com indoor)

*Intensidade:* Percepção de esforço 4-7/10, permitindo conversação confortável

**Recomendações de Segurança:**
- Horários com temperatura amena (evitar 10h-16h no verão)
- Proteção solar (FPS 30+, roupas UV, boné)
- Hidratação adequada (500ml/hora de exercício)
- Calçados apropriados para terreno
- Atenção a condições climáticas adversas

**Exposição Solar:**
- 10-30 minutos de exposição (braços e pernas) 3x/semana para síntese de vitamina D
- Sem protetor solar nos primeiros 15 minutos (se pele clara, reduzir tempo)

**Monitoramento:** Dosagem de 25(OH)D anual, ajuste conforme níveis (<30 ng/mL: suplementação)."""
        }

    def generate_cirurgias_exercicio_content(self, item_data: Dict) -> Dict:
        """Gera conteúdo para Cirurgias relacionadas ao exercício"""
        return {
            "clinicalRelevance": """Cirurgias ortopédicas relacionadas ao exercício (reconstrução de ligamentos, meniscectomia, correção de fraturas por estresse, artroplastias) alteram significativamente biomecânica, propriocepção e capacidade funcional, exigindo periodização cuidadosa do retorno à atividade física. A reabilitação inadequada aumenta risco de re-lesão em 6-8x e predispõe a osteoartrite precoce.

O processo de cicatrização tecidual segue fases distintas: inflamatória (0-5 dias), proliferativa (5 dias-3 semanas) e remodelamento (3 semanas-12+ meses). O exercício terapêutico progressivo promove reorganização colágena, neovascularização e restauração de propriocepção, reduzindo tempo de recuperação em 30-40% comparado a repouso absoluto.

Cirurgias de ligamento cruzado anterior (LCA) requerem 9-12 meses para retorno ao esporte competitivo, com ênfase em fortalecimento de quadríceps, isquiotibiais e core, além de treino de controle neuromuscular. Artroplastias de joelho/quadril beneficiam-se de fortalecimento pré-operatório (pré-habilitação), que reduz complicações pós-operatórias em 50%. Fraturas por estresse refletem desequilíbrio entre carga e recuperação, frequentemente associadas a deficiências nutricionais (vitamina D, cálcio, energia) e síndrome RED-S.""",

            "patientExplanation": """Se você fez cirurgia por causa de lesão durante exercício ou atividade física, é fundamental seguir um plano de recuperação adequado para voltar a se movimentar com segurança. Seu corpo precisa de tempo para cicatrizar, mas também precisa de movimento gradual para recuperar força, equilíbrio e coordenação.

A reabilitação é um processo progressivo: começa com movimentos suaves e exercícios específicos, e aos poucos você vai aumentando a intensidade. Respeitar esse processo é essencial para evitar nova lesão - estudos mostram que quem "pula etapas" tem até 8 vezes mais chance de se machucar novamente.

A boa notícia é que, com orientação adequada e paciência, a maioria das pessoas consegue voltar às atividades que gostava de fazer. Muitas vezes, você até volta mais forte e consciente do seu corpo. O segredo é trabalhar junto com profissionais de saúde especializados e não ter pressa - sua recuperação completa vale a pena.""",

            "conduct": """**Avaliação Pós-Cirúrgica:**
- Tipo de cirurgia, data, protocolo cirúrgico utilizado
- Complicações pós-operatórias (infecção, deiscência, trombose)
- Amplitude de movimento (ADM), força muscular, edema
- Dor (escala 0-10), uso de analgésicos
- Exames de imagem recentes (consolidação óssea, integridade de enxertos)

**Fases da Reabilitação (exemplo LCA):**

*Fase 1 (0-2 semanas):* Controle de dor/edema, ADM passiva, ativação de quadríceps (isometria)

*Fase 2 (2-6 semanas):* ADM ativa completa, fortalecimento progressivo (leg press, extensão de joelho 90-45°), bicicleta ergométrica

*Fase 3 (6-12 semanas):* Fortalecimento intensivo (agachamento, avanço, step), propriocepção (prancha de equilíbrio), início de corrida em linha reta

*Fase 4 (3-6 meses):* Exercícios balísticos, treino de agilidade, corrida com mudança de direção

*Fase 5 (6-12 meses):* Retorno progressivo ao esporte, treinos específicos da modalidade

**Critérios para Progressão:**
- ADM completa e simétrica
- Força ≥80% do lado contralateral (teste isocinético)
- Testes funcionais (single leg hop, Y-balance)
- Ausência de dor/edema após exercício

**Prevenção de Re-lesão:**
- Fortalecimento contínuo de core e membros inferiores
- Correção de assimetrias
- Monitoramento de carga de treino (load management)
- Atenção a fadiga e overtraining

**Colaboração Multiprofissional:** Fisioterapeuta, educador físico, médico ortopedista, nutricionista."""
        }

    def generate_divisao_atividades_content(self, item_data: Dict) -> Dict:
        """Gera conteúdo para Divisão das atividades"""
        return {
            "clinicalRelevance": """A periodização do treinamento através da divisão estratégica de atividades (splits) otimiza adaptações fisiológicas ao permitir recuperação muscular adequada enquanto mantém alta frequência de treino. O princípio da supercompensação indica que o estímulo treino + recuperação resulta em ganhos superiores comparado a treinos não periodizados.

Diferentes divisões produzem adaptações específicas: full body (3x/semana) maximiza síntese proteica muscular através de alta frequência por grupo muscular, sendo ideal para iniciantes e hipertrofia geral. Upper/lower split (4x/semana) permite maior volume por sessão com recuperação adequada. Push/pull/legs (6x/semana) possibilita especialização e volume elevado para avançados.

A síntese proteica muscular permanece elevada por 48-72h após treino de resistência, sugerindo que frequência de 2x/semana por grupo muscular é mínima para maximizar hipertrofia. Entretanto, volume total equalizado (séries totais/semana) produz resultados similares independente da divisão, desde que respeitada a recuperação.

Para objetivos metabólicos e saúde geral, a combinação de treino aeróbico + resistência (treinamento concorrente) apresenta efeitos aditivos sobre sensibilidade à insulina, capacidade cardiorrespiratória e composição corporal, com mínima interferência molecular se aeróbico realizado 3-6h após resistência ou em dias separados.""",

            "patientExplanation": """A "divisão das atividades" é como você organiza seus treinos ao longo da semana para trabalhar diferentes partes do corpo ou diferentes tipos de exercício. Assim como você não estuda todas as matérias ao mesmo tempo, você não precisa treinar tudo em um único dia - e nem deve!

Quando você divide seus treinos de forma inteligente, seus músculos têm tempo para descansar e crescer, enquanto você continua se exercitando. Por exemplo: segunda você treina pernas, quarta treina peito e braços, sexta treina costas. Ou pode fazer: segunda aeróbico, terça musculação, quarta descanso, e assim por diante.

Não existe uma divisão "perfeita" que funcione para todo mundo - o melhor esquema depende de quanto tempo você tem, seus objetivos e seu nível de condicionamento. O importante é ter um planejamento que você consiga seguir de forma consistente e que permita seu corpo se recuperar adequadamente entre os treinos.""",

            "conduct": """**Avaliação para Definir Divisão:**
- Disponibilidade semanal (dias e horas)
- Nível de treinamento (iniciante, intermediário, avançado)
- Objetivos prioritários (força, hipertrofia, emagrecimento, performance)
- Limitações (lesões, desequilíbrios, disponibilidade de equipamentos)
- Preferências pessoais e aderência histórica

**Modelos de Divisão Recomendados:**

**Iniciantes (0-1 ano de treino):**
*Full Body 3x/semana:*
- Segunda/Quarta/Sexta: Corpo inteiro (6-8 exercícios compostos)
- Vantagens: Alta frequência, aprendizado motor, flexibilidade de horários

**Intermediários (1-3 anos):**
*Upper/Lower Split 4x/semana:*
- Segunda/Quinta: Membros superiores
- Terça/Sexta: Membros inferiores
- Vantagens: Maior volume, recuperação adequada

*Push/Pull/Legs 3x/semana:*
- Segunda: Empurrar (peito, ombros, tríceps)
- Quarta: Puxar (costas, bíceps)
- Sexta: Pernas (quadríceps, posteriores, glúteos)

**Avançados (3+ anos):**
*Push/Pull/Legs 6x/semana:* 2x cada divisão
*Bro Split:* Segunda-perna, Terça-costas, Quarta-peito, Quinta-ombro, Sexta-braços (apenas se volume alto tolerado)

**Integração de Aeróbico:**
- Objetivo saúde geral: 150 min/semana moderado, distribuído 3-5x
- Emagrecimento: Adicionar 200-300 min/semana, preferencialmente em jejum ou pós-treino
- Performance: Periodizar conforme fase (base aeróbica vs. força específica)

**Monitoramento:**
- Sinais de overtraining (fadiga persistente, insônia, queda de performance)
- Progressão de cargas e volume
- Reavaliação mensal, ajuste conforme resposta individual."""
        }

    def generate_esportes_praticados_content(self, item_data: Dict) -> Dict:
        """Gera conteúdo para Esportes praticados"""
        return {
            "clinicalRelevance": """A prática esportiva regular combina benefícios fisiológicos do exercício com componentes sociais, cognitivos e motivacionais, resultando em maior aderência (70-80% vs. 40-50% em exercícios não estruturados) e impactos superiores em saúde mental. Diferentes modalidades esportivas induzem adaptações específicas: esportes de resistência (ciclismo, corrida, natação) aumentam densidade mitocondrial, capacidade oxidativa e eficiência cardíaca; esportes de potência (sprint, saltos, levantamento de peso) desenvolvem fibras tipo II, força e potência muscular; esportes de habilidade (tênis, basquete, futebol) aprimoram coordenação, tempo de reação e função executiva.

A frequência e intensidade da prática esportiva modulam respostas adaptativas: 3x/semana intensidade moderada promove benefícios cardiovasculares básicos (redução de 20-30% em risco cardiovascular), enquanto 5-7x/semana intensidade moderada-alta maximiza ganhos de VO2máx (+15-25%), composição corporal e marcadores metabólicos. Entretanto, volume excessivo sem recuperação adequada (>15h/semana em não-atletas) aumenta risco de overtraining, lesões por overuse e síndrome RED-S.

Esportes coletivos apresentam vantagens adicionais: estimulam liberação de ocitocina (coesão social), reduzem percepção de esforço (facilitação social), desenvolvem habilidades de comunicação e resiliência. Estudos longitudinais demonstram que praticantes de esportes coletivos na adolescência apresentam 40% menor incidência de obesidade e 35% menor risco de depressão na vida adulta.""",

            "patientExplanation": """Praticar esportes vai muito além de apenas se exercitar - é uma forma divertida de cuidar da saúde do corpo e da mente ao mesmo tempo. Quando você joga futebol, tênis, vôlei, basquete ou qualquer outro esporte, você está treinando força, resistência, coordenação e agilidade de forma natural, sem nem perceber.

Os esportes também trazem benefícios que a academia sozinha não oferece: você socializa, faz amigos, aprende a trabalhar em equipe, desenvolve disciplina e persistência. Estudos mostram que pessoas que praticam esportes regularmente têm mais facilidade para manter o hábito de se exercitar a vida toda, porque é algo prazeroso, não apenas uma obrigação.

O importante é escolher esportes que você realmente goste - pode ser um esporte coletivo, individual, competitivo ou apenas recreativo. O melhor esporte é aquele que você vai praticar com regularidade e entusiasmo. Se possível, combine diferentes modalidades para desenvolver o corpo de forma completa.""",

            "conduct": """**Anamnese Esportiva Detalhada:**
- Modalidades praticadas (atual e histórico)
- Frequência semanal (sessões e duração)
- Intensidade (recreativo, competitivo amador, alto rendimento)
- Anos de prática em cada modalidade
- Lesões prévias relacionadas ao esporte
- Objetivos (lazer, condicionamento, competição)

**Avaliação de Adequação:**
- Compatibilidade com condicionamento atual e limitações
- Risco de lesões (alto impacto, contato, movimentos repetitivos)
- Demandas metabólicas vs. capacidade cardiovascular
- Necessidade de preparação física específica

**Prescrição de Suporte ao Esporte:**

**Preparação Física Geral (base para todos os esportes):**
1. **Força fundamental:** Agachamento, deadlift, press horizontal/vertical, remadas (2-3x/semana)
2. **Core e estabilidade:** Prancha, rotações, anti-rotação (3x/semana)
3. **Mobilidade:** Alongamento dinâmico pré-treino, estático pós-treino, liberação miofascial

**Preparação Específica (conforme modalidade):**
- *Esportes de corrida:* Fortalecimento de glúteos/core, pliometria, treino de passada
- *Esportes de raquete:* Rotação de tronco, estabilidade de ombro, agilidade lateral
- *Esportes de contato:* Força máxima, potência, propriocepção, prevenção de concussão

**Periodização:**
- *Off-season:* Foco em construção de base (força, condicionamento geral)
- *Pre-season:* Transição para especificidade (potência, velocidade, agilidade)
- *In-season:* Manutenção + recuperação (volume reduzido, intensidade preservada)

**Monitoramento de Carga:**
- Questionário de recuperação (TQR, Hooper Index)
- Marcadores de fadiga (variabilidade de frequência cardíaca)
- Volume/intensidade semanal (evitar aumentos >10% por semana)

**Prevenção de Lesões:**
- Warm-up específico (ativação + movimentos do esporte)
- Fortalecimento de zonas vulneráveis (tornozelo, joelho, ombro)
- Gerenciamento de carga e períodos de recuperação ativa
- Equipamentos adequados (calçados, proteções)

**Nutrição Esportiva:** Ajuste de macronutrientes conforme demanda (CHO para esportes de resistência, proteína para esportes de potência)."""
        }

    def generate_estrategia_macro_content(self, item_data: Dict) -> Dict:
        """Gera conteúdo para Estratégia macro atual"""
        return {
            "clinicalRelevance": """A estratégia macroscópica de treinamento (periodização) refere-se à organização de ciclos de treino em diferentes escalas temporais (microciclo/semanal, mesociclo/mensal, macrociclo/anual) com variação sistemática de volume, intensidade e especificidade para otimizar adaptações e prevenir estagnação ou overtraining. O modelo de periodização linear clássico (Matveev) progride de alto volume/baixa intensidade para baixo volume/alta intensidade, sendo efetivo para iniciantes e atletas focados em pico de performance sazonal.

A periodização ondulatória (DUP - Daily Undulating Periodization) alterna estímulos dentro do microciclo (ex: segunda-hipertrofia 8-12 reps, quarta-força 3-5 reps, sexta-resistência 15-20 reps), promovendo recuperação adequada e recrutamento variado de fibras musculares. Meta-análises indicam superioridade da DUP para hipertrofia e força em treinados (+8-12% vs. periodização linear).

A periodização em blocos concentra volumes elevados de determinado estímulo em mesociclos específicos (ex: bloco de hipertrofia 4 semanas, bloco de força 3 semanas, bloco de potência 2 semanas), sendo ideal para atletas avançados. Para saúde geral e não-atletas, periodização não-linear com variação moderada de intensidades e volumes é suficiente e mais sustentável.

A deload (semana de descarga a cada 4-6 semanas, com redução de 40-50% do volume) é fundamental para consolidar adaptações, permitir recuperação sistêmica e prevenir lesões por overuse, resultando em progressão 15-20% superior em protocolos ≥12 semanas.""",

            "patientExplanation": """A "estratégia macro" é como você planeja seus treinos ao longo de semanas e meses para continuar progredindo sem se lesionar ou estagnar. É como subir uma montanha: você não vai em linha reta o tempo todo - às vezes acelera, às vezes desacelera, e de vez em quando precisa descansar antes de continuar subindo.

Se você treinar sempre com a mesma intensidade e os mesmos exercícios, seu corpo se acostuma e você para de evoluir. Mas se você variar de forma inteligente - algumas semanas mais pesadas, outras mais leves, algumas focando em força, outras em resistência - seu corpo continua se adaptando e melhorando.

Pense nisso como um plano de longo prazo: você não precisa (nem deve) ir ao máximo todos os dias. Ter períodos de treino mais intenso seguidos de períodos de recuperação ativa faz você progredir mais rápido e com menos risco de lesão. É o conceito de "dois passos para frente, um para trás" - às vezes desacelerar estrategicamente é o que te leva mais longe.""",

            "conduct": """**Avaliação da Estratégia Atual:**
- Existe periodização formal ou treino aleatório?
- Ciclos atuais (micro/meso/macrociclos definidos?)
- Progressão de volume e intensidade nas últimas 12 semanas
- Frequência de deloads (semanas de recuperação)
- Variáveis manipuladas (carga, volume, frequência, seleção de exercícios)
- Sinais de estagnação ou overtraining

**Implementação de Periodização (conforme nível):**

**Iniciantes (0-1 ano):**
*Progressão Linear Simples:*
- Aumentar carga 2-5% quando alcançar limite superior de repetições
- Deload a cada 6-8 semanas (50% volume)
- Foco em consistência e técnica

**Intermediários (1-3 anos):**
*Periodização Ondulatória Diária (DUP):*

*Microciclo exemplo:*
- Segunda: 4x8-10 (hipertrofia, 70-75% 1RM, 2-3min descanso)
- Quarta: 5x3-5 (força, 85-90% 1RM, 3-5min descanso)
- Sexta: 3x12-15 (resistência muscular, 60-65% 1RM, 1-2min descanso)

*Mesociclo:* 4-6 semanas progressão + 1 semana deload

**Avançados (3+ anos):**
*Periodização em Blocos:*

*Macrociclo 12 semanas:*
- Bloco 1 (semanas 1-4): Hipertrofia - 4-5x8-12, volume alto, intensidade moderada
- Bloco 2 (semanas 5-7): Força - 4-6x3-6, volume médio, intensidade alta
- Bloco 3 (semanas 8-10): Potência - 3-5x1-3, volume baixo, intensidade máxima
- Semana 11: Deload
- Semana 12: Teste de 1RM/competição

**Princípios de Progressão:**
1. **Sobrecarga progressiva:** Aumentar uma variável por vez (peso, reps, séries, frequência)
2. **Variação:** Mudar estímulo a cada 4-6 semanas (ângulos, exercícios, métodos)
3. **Recuperação:** Deload 40-50% volume a cada 4-6 semanas intensas
4. **Especificidade:** Aproximar da modalidade/objetivo conforme avança o ciclo

**Monitoramento:**
- Planilha de treino com progressão registrada
- Avaliações físicas mensais (força, composição corporal)
- Percepção de recuperação (escala 1-10 diária)
- Ajuste conforme resposta individual (genetica, idade, stress, sono)

**Sinais para Ajuste:**
- Estagnação >3 semanas: Aumentar variação ou volume
- Fadiga crônica: Implementar deload imediato
- Dor articular persistente: Reduzir intensidade, aumentar recuperação."""
        }

    def update_item(self, item_id: str, content: Dict) -> bool:
        """Atualiza item com novo conteúdo"""
        try:
            resp = self.session.put(
                f"{API_BASE}/score-items/{item_id}",
                json=content
            )
            resp.raise_for_status()
            return True
        except Exception as e:
            print(f"  ✗ Erro ao atualizar item {item_id}: {e}")
            self.errors.append(f"{item_id}: {str(e)}")
            return False

    def process_item(self, item_id: str, item_name: str) -> bool:
        """Processa um item individual"""
        print(f"\n[{self.processed + 1}/30] Processando: {item_name} ({item_id})")

        # Buscar detalhes do item
        item_data = self.get_item_details(item_id)
        if not item_data:
            return False

        # Gerar conteúdo baseado no tipo
        if "Adolescência" in item_name:
            content = self.generate_adolescencia_content(item_data)
        elif "Atividades ao ar livre" in item_name:
            content = self.generate_atividades_ar_livre_content(item_data)
        elif "Cirurgias" in item_name:
            content = self.generate_cirurgias_exercicio_content(item_data)
        elif "Divisão das atividades" in item_name:
            content = self.generate_divisao_atividades_content(item_data)
        elif "Esportes praticados" in item_name:
            content = self.generate_esportes_praticados_content(item_data)
        elif "Estratégia macro" in item_name:
            content = self.generate_estrategia_macro_content(item_data)
        else:
            print(f"  ⚠ Tipo não reconhecido: {item_name}")
            return False

        # Atualizar item
        if self.update_item(item_id, content):
            print(f"  ✓ Item atualizado com sucesso")
            self.processed += 1
            return True

        return False

    def run(self):
        """Executa processamento de todos os items"""
        print("=" * 80)
        print("PROCESSAMENTO DE SCORE ITEMS - MOVIMENTO E ATIVIDADE FÍSICA")
        print("=" * 80)

        if not self.login():
            return

        print(f"\nIniciando processamento de {len(ITEMS)} items...")

        start_time = time.time()

        for item_id, item_name in ITEMS:
            self.process_item(item_id, item_name)
            time.sleep(0.5)  # Rate limiting

        elapsed = time.time() - start_time

        print("\n" + "=" * 80)
        print("RESUMO DO PROCESSAMENTO")
        print("=" * 80)
        print(f"Total processados: {self.processed}/{len(ITEMS)}")
        print(f"Sucessos: {self.processed}")
        print(f"Erros: {len(self.errors)}")
        print(f"Tempo total: {elapsed:.1f}s")

        if self.errors:
            print("\nErros encontrados:")
            for error in self.errors:
                print(f"  - {error}")

        print("\n✓ Processamento concluído!")

if __name__ == "__main__":
    processor = MovimentoProcessor()
    processor.run()
