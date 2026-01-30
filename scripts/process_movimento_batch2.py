#!/usr/bin/env python3
"""
Processamento de Score Items - Grupo "Movimento e atividade física" - Batch 2
Próximos 30 items com enriquecimento de evidências científicas
"""

import requests
import json
import time
from typing import Dict, List, Optional

API_BASE = "http://localhost:3001/api/v1"
EMAIL = "import@plenya.com"
PASSWORD = "Import@123456"

# Items a processar (Batch 2)
ITEMS = [
    ("7404cc64-61e3-4d65-aacf-88dd0d1eac54", "Histórico familiar de exercícios"),
    ("c5030479-8798-4da1-a385-f4fb16706b10", "Histórico familiar de exercícios"),
    ("1dff2de5-4dcc-416e-924a-f761b5bd82df", "Histórico familiar de exercícios"),
    ("9be31ce4-d174-47b2-ad4e-16134063ccf1", "Horários"),
    ("bbc0c563-c4ef-42b6-9e95-7f5a986a401d", "Horários"),
    ("7e019b01-4dc4-4f06-86a6-9119be8518af", "Horários"),
    ("0470945d-d26c-4fe8-b2f8-983e189ee327", "Infância"),
    ("1c1659ac-6511-46b4-a6e4-085fb2a7e934", "Infância"),
    ("49c880ac-1c90-47ce-8296-6d488eefc152", "Infância"),
    ("763fd69e-ea11-4970-ac11-f37196e59531", "Lesões relacionadas ao exercício"),
    ("37e3a409-7c1b-484e-a7b0-2960e52e54fb", "Lesões relacionadas ao exercício"),
    ("36b2c826-aca4-4de9-ad64-b75ed1037b5e", "Lesões relacionadas ao exercício"),
    ("d963c0f3-d0c9-47f7-9189-b90393cac2b0", "Melhores fases atléticas (idade, duração, esportes praticados)"),
    ("8c6761e6-40c6-4557-be33-e8102122bc16", "Melhores fases atléticas (idade, duração, esportes praticados)"),
    ("3aa65199-9371-41d9-aa16-6047f9ffb0c8", "Melhores fases atléticas (idade, duração, esportes praticados)"),
    ("9110f14b-f28f-4a2d-803b-7e39bddcb1c0", "Modalidades de esporte \"odiadas\""),
    ("609f1710-d357-4050-9f24-91e32ae49491", "Modalidades de esporte \"odiadas\""),
    ("0dd18174-e2ba-4be3-a79c-10b56226e425", "Modalidades de esporte \"odiadas\""),
    ("edd95777-9a26-4478-ba5d-6f8ee5b8b381", "Modalidades de esporte preferidas"),
    ("39f75c4f-47ef-44d3-af10-b4210f1ef00a", "Modalidades de esporte preferidas"),
    ("6d6b4c3b-46b6-47fc-8e1a-05dd744eba0e", "Modalidades de esporte preferidas"),
    ("e0ce855a-8d5a-4cd0-9069-797e9ccdedec", "Onde e como faz"),
    ("fbb1d7f4-0d43-42ed-832d-109f8863575e", "Onde e como faz"),
    ("53944c26-fc80-4b8a-a445-311a0f633963", "Onde e como faz"),
    ("ab970f66-09fc-4856-af2d-b66b4c06eaf9", "Piores fases de sedentarismo (idade, duração, fatores associados)"),
    ("29da00d5-fe6a-4469-b178-16c97cb0c25b", "Piores fases de sedentarismo (idade, duração, fatores associados)"),
    ("60ec1673-1072-40f0-8f4b-50cbdeb0cc6c", "Piores fases de sedentarismo (idade, duração, fatores associados)"),
    ("3de73c6e-3a9d-46f3-b535-83f3be573eb1", "Provas/desafios planejados para os próximos 6 meses"),
    ("1a0ce0bd-5886-45bb-ad6a-2f22013e4aca", "Provas/desafios planejados para os próximos 6 meses"),
    ("268ee3a0-1cee-42e0-b0b4-160b1385b749", "Provas/desafios planejados para os próximos 6 meses"),
]

class MovimentoBatch2Processor:
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

    def generate_historico_familiar_content(self, item_data: Dict) -> Dict:
        """Gera conteúdo para Histórico familiar de exercícios"""
        return {
            "clinicalRelevance": """O histórico familiar de atividade física apresenta influência significativa sobre padrões comportamentais individuais através de mecanismos genéticos, epigenéticos e socioculturais. Estudos de hereditabilidade estimam que 30-60% da variação no nível de atividade física é explicada por fatores genéticos, incluindo polimorfismos em genes relacionados a receptores dopaminérgicos (DRD2, DRD4), transportador de dopamina (DAT1), BDNF e ACTN3 (alfa-actinina-3, associada a performance em fibras tipo II).

A modelagem parental representa o fator ambiental mais forte: crianças com ambos os pais fisicamente ativos têm 5.8x mais probabilidade de serem ativas comparadas a crianças com pais sedentários. Este efeito persiste na vida adulta - adultos cujos pais foram ativos na infância apresentam 40% maior adesão a programas de exercício. Mecanismos incluem aprendizado observacional, reforço positivo, disponibilização de recursos (equipamentos, transporte para atividades) e normalização cultural da atividade física.

Famílias com histórico de prática esportiva transmitem não apenas comportamentos, mas também valores relacionados a disciplina, resiliência e orientação a metas. Estudos longitudinais demonstram que o ambiente familiar ativo na infância prediz maior capacidade cardiorrespiratória (+12% VO2máx), melhor composição corporal (-2.5% gordura corporal) e menor risco de doenças crônicas (redução de 28% em diabetes tipo 2, 18% em doenças cardiovasculares) na vida adulta, mesmo após ajuste para fatores socioeconômicos.""",

            "patientExplanation": """Sua história familiar com exercícios e esportes influencia mais do que você imagina nos seus hábitos atuais. Se você cresceu vendo seus pais, avós ou irmãos se exercitando, é mais provável que você também tenha desenvolvido esse hábito - não apenas porque pode ter "herdado" genes favoráveis, mas principalmente porque cresceu em um ambiente onde atividade física era normal e valorizada.

Isso não significa que quem não teve esse exemplo está condenado a ser sedentário! Significa apenas que você pode precisar fazer um esforço consciente maior para criar novos hábitos, já que não teve esses modelos na infância. A boa notícia é que você pode quebrar esse ciclo e se tornar o exemplo positivo para a próxima geração.

Conhecer o histórico familiar ajuda a entender suas próprias tendências e dificuldades. Se seus pais eram ativos e você também é, aproveite essa facilidade. Se não eram, reconheça que você está construindo algo novo, o que merece ser celebrado a cada pequena conquista.""",

            "conduct": """**Anamnese Detalhada do Histórico Familiar:**
- Pais: Nível de atividade física (sedentário, moderadamente ativo, muito ativo), modalidades praticadas, frequência
- Irmãos: Padrões de atividade, esportes praticados, nível competitivo
- Avós: Longevidade, capacidade funcional, histórico de atividade
- Cultura familiar: Valorização do esporte, participação em atividades recreativas conjuntas
- Barreiras familiares: Doenças, lesões, fatores socioeconômicos que limitaram atividade

**Interpretação e Uso Clínico:**

**Paciente com histórico familiar ativo:**
- Aproveitar motivação intrínseca pré-existente
- Explorar modalidades já conhecidas através da família
- Incentivar retorno/continuidade de práticas anteriores
- Usar comparação positiva (não competitiva) com familiares como motivador

**Paciente sem histórico familiar ativo:**
- Reconhecer desafio adicional e normalizar dificuldades iniciais
- Estabelecer objetivos mais graduais e celebrar pequenas vitórias
- Construir suporte social alternativo (grupos, comunidades esportivas)
- Enfatizar quebra de ciclo e criação de legado positivo para futuros descendentes
- Educação sobre benefícios além da performance (saúde, bem-estar)

**Histórico familiar de lesões/limitações:**
- Investigar fatores de risco hereditários (hipermobilidade, frouxidão ligamentar)
- Rastreamento preventivo (avaliação biomecânica, testes funcionais)
- Fortalecimento de zonas vulneráveis desde início
- Educação sobre progressão segura

**Intervenções Específicas:**
- Atividades familiares conjuntas (caminhadas, jogos, esportes recreativos) para criar novo contexto positivo
- Envolvimento de cônjuge/filhos em programa de exercício para mudar cultura familiar atual
- Para pais: Modelagem ativa para filhos (ser o exemplo que não teve)

**Aconselhamento Genético (quando disponível):**
- Testes de polimorfismos relacionados a resposta ao treino (não essencial, uso limitado)
- Foco principal permanece em modificação comportamental e ambiental

**Monitoramento:** Reavaliar influência do contexto familiar na adesão, ajustar estratégias conforme necessário."""
        }

    def generate_horarios_content(self, item_data: Dict) -> Dict:
        """Gera conteúdo para Horários de treino"""
        return {
            "clinicalRelevance": """O momento do dia em que o exercício é realizado (chronotype e timing) influencia adaptações fisiológicas, performance e aderência através de mecanismos circadianos. O ritmo circadiano regula temperatura corporal, secreção hormonal (cortisol, testosterona, GH, melatonina), tônus muscular e função cardiorrespiratória, com variações de 5-20% entre picos e nadires.

Performance física atinge pico no final da tarde/início da noite (16h-20h), quando temperatura corporal está mais elevada (+0.5-1°C), resultando em: maior potência muscular (+5-8%), tempo de reação reduzido (7-15%), melhor flexibilidade (+20%), e VO2máx ligeiramente superior (+3-5%). Exercícios matinais (6h-10h) ocorrem sob níveis elevados de cortisol (pico às 8h), promovendo maior lipólise e mobilização de ácidos graxos, sendo vantajosos para emagrecimento quando realizados em jejum ou baixo carboidrato.

Exercícios noturnos (pós 20h) podem interferir no sono em indivíduos sensíveis devido à elevação prolongada de catecolaminas e temperatura corporal, embora meta-análises recentes indiquem que exercício <1-2h antes de dormir não prejudica sono na maioria das pessoas, e pode até melhorar qualidade do sono profundo. Treino de força noturno aumenta síntese proteica durante o sono se acompanhado de ingestão proteica pré-sono (40g caseína).

A consistência do horário é mais importante que o horário "ideal" per se. Treinar no mesmo período diariamente por 4+ semanas induz adaptações circadianas específicas (phase entrainment), otimizando performance naquele horário específico. Indivíduos vespertinos (corujas) apresentam melhor performance e menor risco de lesão em treinos noturnos, enquanto matutinos (cotovias) performam melhor pela manhã.""",

            "patientExplanation": """Não existe um horário "mágico" para se exercitar que funcione para todo mundo - o melhor horário é aquele em que você consegue ser consistente e se sente bem. Dito isso, seu corpo funciona com um relógio biológico interno que faz você ter mais energia, força e coordenação em certos períodos do dia.

A maioria das pessoas tem melhor performance física no final da tarde e início da noite, quando a temperatura do corpo está mais alta e os músculos mais "acordados". Treinar de manhã tem vantagens para quem quer emagrecer (queima mais gordura) e para quem tem agenda imprevisível (faz logo antes de outras coisas aparecerem). Treinar à noite pode ser ótimo para construir músculos, mas para algumas pessoas pode atrapalhar o sono.

O mais importante é escolher um horário que caiba na sua rotina e que você consiga manter. Depois de 3-4 semanas treinando no mesmo horário, seu corpo se adapta e fica ainda mais fácil. Não se force a acordar às 5h da manhã se você não é uma pessoa matutina - isso pode prejudicar mais do que ajudar. Encontre SEU melhor horário.""",

            "conduct": """**Avaliação do Cronótipo e Preferências:**
- Questionário de cronotipo (MEQ - Morningness-Eveningness Questionnaire) para identificar se é matutino, vespertino ou intermediário
- Horários atuais de treino e percepção de energia/performance
- Rotina de trabalho, compromissos familiares, disponibilidade
- Qualidade do sono atual e horários de dormir/acordar
- Experiências prévias com diferentes horários

**Recomendações Baseadas em Objetivos:**

**Emagrecimento:**
- *Preferencial:* Manhã em jejum (6h-8h) - maior oxidação de gordura (+20-30%)
- *Alternativo:* Pós-trabalho (17h-19h) com intensidade moderada-alta
- *Importante:* Déficit calórico diário é mais importante que timing

**Hipertrofia/Força:**
- *Ótimo:* Tarde/Início da noite (16h-20h) - pico de força e potência
- *Bom:* Fim da manhã (10h-12h) após warm-up adequado
- *Evitar:* Muito cedo (6h-7h) - risco aumentado de lesões (rigidez matinal)

**Performance Esportiva:**
- Treinar no mesmo horário das competições (entrainment específico)
- Periodização: Treinos técnicos pela manhã, força/potência à tarde

**Saúde Geral/Bem-estar:**
- *Flexível:* Qualquer horário consistente
- *Considerações:* Evitar >2h antes de dormir se tiver dificuldade com sono

**Otimização por Cronotipo:**

**Matutinos (cotovias):**
- 6h-10h: Ideal para treinos aeróbicos, mobilidade, técnica
- Vantagens: Aderem melhor, agendas mais previsíveis
- Cuidados: Warm-up prolongado, evitar cargas máximas antes das 8h

**Vespertinos (corujas):**
- 16h-21h: Melhor janela de performance
- Vantagens: Maior força, potência, recuperação
- Cuidados: Suplementar cafeína se treino pré-19h for inviável

**Estratégias de Adaptação:**
- Iniciar treino 30min mais cedo/tarde a cada 2 semanas (transição gradual)
- Exposição à luz adequada: manhã para acordar, escurecer à noite
- Refeição pré-treino 60-90min antes para garantir energia
- Se treino noturno prejudica sono: terminar ≥2h antes de dormir, banho frio pós-treino, evitar cafeína pós-15h

**Consistência > Horário "Perfeito":**
- Priorizar aderência ao escolher horário
- Permitir flexibilidade (2 horários alternativos na semana)
- Monitorar: Performance percebida, qualidade do sono, adesão ao longo de 4 semanas

**Ajustes Sazonais:** Considerar mudanças de luz solar (inverno vs. verão), temperatura ambiente."""
        }

    def generate_infancia_content(self, item_data: Dict) -> Dict:
        """Gera conteúdo para Infância e atividade física"""
        return {
            "clinicalRelevance": """A infância (0-9 anos) constitui período crítico para desenvolvimento de competências motoras fundamentais, estabelecimento de padrões neuromotores e construção de capital ósseo e muscular que influenciarão capacidades físicas ao longo da vida. O desenvolvimento motor segue progressão previsível: controle postural (0-1 ano), locomoção básica (1-2 anos), manipulação de objetos (2-3 anos), e refinamento de habilidades complexas (4-9 anos).

Crianças fisicamente ativas apresentam desenvolvimento cognitivo superior: estudos longitudinais demonstram que atividade física na infância prediz melhor função executiva (+15-20%), memória de trabalho, atenção sustentada e desempenho acadêmico na adolescência. Mecanismos incluem aumento de BDNF, neuroplasticidade hipocampal, mielinização acelerada e desenvolvimento de redes pré-frontais.

O pico de ganho de massa óssea ocorre na infância/adolescência - atividades de impacto (saltos, corrida, esportes) durante esta fase aumentam densidade mineral óssea em até 8-12%, reduzindo risco de osteoporose na senescência em até 50%. Entretanto, especialização esportiva precoce (<12 anos) aumenta risco de lesões por overuse em 70-90%, burnout psicológico e abandono da prática (dropout rate 70% aos 13 anos).

Diretrizes pediátricas recomendam: lactentes (0-1 ano) atividade motora livre supervisionada múltiplas vezes ao dia; pré-escolares (3-5 anos) mínimo 180 min/dia de atividade; crianças (6-9 anos) 60+ min/dia de atividade moderada-vigorosa, incluindo variabilidade de movimentos e brincadeiras não estruturadas. Exposição a múltiplas modalidades (sampling approach) resulta em maior desenvolvimento motor global comparado a especialização precoce.""",

            "patientExplanation": """A infância é quando seu corpo "aprende" a se movimentar - é como construir a fundação de uma casa. Crianças que brincam, correm, pulam, escalam e praticam diferentes atividades desenvolvem músculos, ossos, coordenação e equilíbrio muito mais fortes. Esses benefícios duram para a vida toda.

Mas não é só o corpo: crianças ativas também têm melhor memória, concentração e desempenho na escola. O cérebro se desenvolve melhor quando o corpo está em movimento. Por isso brincar ao ar livre, praticar esportes e se movimentar é tão importante quanto estudar.

Se você teve uma infância ativa, isso pode explicar por que é relativamente fácil para você se exercitar hoje. Se não teve, não se preocupe - nunca é tarde para começar. O importante é entender que crianças precisam se movimentar livremente, experimentar diferentes atividades e brincar muito, sem pressão por resultados. O objetivo é desenvolver amor pelo movimento, não criar atletas profissionais.""",

            "conduct": """**Anamnese de Atividade Física na Infância:**
- Nível geral de atividade (muito ativo, moderadamente ativo, sedentário)
- Modalidades praticadas: esportes organizados, brincadeiras livres, atividades ao ar livre
- Frequência e duração das atividades
- Apoio familiar (pais incentivavam? participavam?)
- Ambiente (urbano vs. rural, acesso a espaços verdes, segurança)
- Marcos motores (idade de andar, correr, andar de bicicleta)
- Experiências positivas vs. negativas (traumas, bullying, sucessos)
- Transporte ativo (ia caminhando/bicicleta para escola?)

**Interpretação Clínica:**

**Infância ativa (60+ min/dia atividade vigorosa):**
- *Vantagens atuais:* Maior capital motor, facilidade de aprendizagem de novos movimentos, melhor propriocepção
- *Conduta:* Aproveitar base sólida, progredir rapidamente, explorar modalidades complexas

**Infância moderadamente ativa (30-60 min/dia):**
- *Avaliação:* Identificar déficits específicos (equilíbrio, coordenação, força)
- *Conduta:* Programa balanceado com ênfase em competências subdesenvolvidas

**Infância sedentária (<30 min/dia):**
- *Desafios esperados:* Menor consciência corporal, medo de movimento, coordenação limitada
- *Conduta:*
  - Iniciar com movimentos fundamentais (agachamento, empurrar, puxar, locomoção)
  - Progressão muito gradual (6-12 meses fase de adaptação)
  - Ênfase em prazer e autoeficácia, não performance
  - Considerar aulas de movimentos básicos (pilates, ginástica funcional adaptada)
  - Explorar atividades lúdicas (dança, artes marciais, escalada recreativa)

**Experiências Negativas na Infância:**
- *Trauma de educação física:* Bullying, humilhação pública, comparações negativas
- *Abordagem:*
  - Ambientes não-competitivos inicialmente
  - Reforço positivo constante
  - Foco em progresso individual, não comparação
  - Considerar treinos individuais ou grupos pequenos

**Especialização Precoce (histórico):**
- *Se continuou:* Avaliar overuse, desequilíbrios musculares, burnout
- *Se abandonou:* Explorar causas (pressão, lesão, perda de prazer), reconstruir relação positiva com movimento

**Recomendações para Pacientes com Filhos:**
- Educação sobre importância de infância ativa
- Encorajar brincadeiras livres, exploração motora variada
- Limitar tempo de tela (<2h/dia)
- Modelagem parental (ser ativo junto com criança)
- Evitar especialização antes dos 12 anos (sampling approach)
- Foco em diversão, não vitórias

**Recuperando Déficits Motores em Adultos:**
- Programa de alfabetização motora: locomoção, estabilidade, manipulação de objetos
- Atividades que desafiam coordenação: dança, artes marciais, esportes de raquete
- Trabalho de propriocepção e equilíbrio (1-2x/semana)
- Paciência: competências motoras podem levar meses para desenvolver

**Monitoramento:** Progressão de competências motoras, redução de medo/ansiedade de movimento, aumento de autoeficácia física."""
        }

    def generate_lesoes_exercicio_content(self, item_data: Dict) -> Dict:
        """Gera conteúdo para Lesões relacionadas ao exercício"""
        return {
            "clinicalRelevance": """Lesões relacionadas ao exercício classificam-se em agudas (trauma único: entorses, rupturas, fraturas) e crônicas por overuse (microtraumas repetitivos: tendinites, fraturas por estresse, síndromes compressivas). A incidência varia conforme modalidade: esportes de contato (8-15 lesões/1000h), corrida (20-70% lesionados anualmente), musculação (2.4-3.3 lesões/1000h). Fatores de risco incluem treinamento inadequado (aumento >10% carga/semana), assimetrias musculares (>15% diferença bilateral), déficits de mobilidade, fadiga acumulada e histórico prévio de lesão (recorrência em 30-50%).

Lesões por overuse resultam de desequilíbrio entre carga mecânica e capacidade de reparação tecidual. Tendões toleram cargas até 4-8% de strain antes de microlesões; repetição excede capacidade regenerativa (24-72h), progredindo de tendinopatia reativa → disrepair → degenerativa. Fraturas por estresse ocorrem quando remodelamento ósseo (reabsorção osteoclástica seguida de formação osteoblástica) é interrompido por carga contínua, prevalentes em tíbia, metatarsos e fêmur, associadas a tríade da atleta feminina (amenorreia, osteoporose, baixa disponibilidade energética).

Prevenção baseia-se em gerenciamento de carga (monitorar volume/intensidade via acute:chronic workload ratio <1.5), fortalecimento excêntrico de zonas vulneráveis (isquiotibiais, tendão de Aquiles), correção de padrões biomecânicos disfuncionais e periodização adequada com deloads. Retorno ao esporte segue critérios funcionais: amplitude de movimento completa, força ≥90% contralateral, testes de salto/agilidade >85% baseline, ausência de dor em progressão gradual (regra 10% - aumentar 10% carga/semana).""",

            "patientExplanation": """Lesões durante exercícios são frustrantes, mas na maioria das vezes podem ser prevenidas ou recuperadas completamente se tratadas adequadamente. Existem dois tipos principais: lesões agudas (acontecem de repente, como torcer o tornozelo) e lesões por uso excessivo (aparecem aos poucos, por fazer o mesmo movimento muitas vezes sem descanso suficiente).

A maioria das lesões não acontece por "azar" - acontece porque aumentamos a intensidade muito rápido, não descansamos o suficiente, temos desequilíbrios musculares, ou usamos técnica errada. O corpo precisa de tempo para se adaptar e ficar mais forte. Se você pular etapas ou ignorar sinais de alerta (dores que não passam, fadiga extrema), aumenta muito o risco de se machucar.

A boa notícia é que, respeitando limites, progredindo gradualmente, fortalecendo o corpo de forma equilibrada e dando tempo para recuperação, você pode se exercitar de forma segura por muitos anos. E se você já teve lesão, seguir um programa de reabilitação completo (não apenas até a dor passar) reduz muito a chance de acontecer de novo.""",

            "conduct": """**Anamnese Detalhada de Lesões:**
- Tipo de lesão: Aguda (trauma) vs. Crônica (overuse)
- Localização anatômica específica (articulação, músculo, tendão, osso)
- Circunstâncias: Atividade realizada, momento (início, meio, fim treino), fadiga presente
- Tratamento recebido: Fisioterapia, cirurgia, imobilização, medicamentos
- Tempo de recuperação até retorno ao exercício
- Recuperação completa? Limitações residuais, dor ocasional, medo de re-lesão
- Recorrências: Mesma lesão repetida?

**Classificação de Severidade:**
- *Grau I (leve):* Desconforto, sem limitação funcional, recuperação <2 semanas
- *Grau II (moderada):* Dor durante/após atividade, limitação parcial, 2-6 semanas
- *Grau III (severa):* Incapacidade funcional, necessidade de intervenção médica, >6 semanas

**Identificação de Fatores de Risco Persistentes:**

**Biomecânicos:**
- Avaliação de movimento (screening - FMS, Y-balance test)
- Assimetrias de força (teste isocinético ou dinamômetro)
- Déficits de mobilidade (flexibilidade de quadril, tornozelo, ombro)
- Análise de padrões de movimento (agachamento, single leg hop)

**Carga de Treino:**
- Histórico de progressão (aumentos >10%/semana?)
- Acute:chronic workload ratio (carga última semana / média 4 semanas - ideal 0.8-1.3)
- Frequência de deloads (redução de volume)
- Sono, nutrição, estresse (impactam recuperação)

**Prevenção de Re-lesão:**

**Fortalecimento Específico:**
- *Lesões de joelho:* Fortalecimento de quadríceps (especialmente VMO), glúteo médio, core
- *Lesões de tornozelo:* Propriocepção (prancha de equilíbrio), fortalecimento de fibulares
- *Lesões de ombro:* Manguito rotador, estabilizadores escapulares
- *Lesões de coluna lombar:* Core stability (prancha, bird dog, dead bug)
- *Isquiotibiais:* Exercícios excêntricos (Nordic hamstring, RDL)

**Correção Biomecânica:**
- Instrução técnica específica (agachamento, corrida, levantamentos)
- Feedback visual/tátil durante execução
- Progressão de complexidade (bilateral → unilateral → dinâmico)

**Gerenciamento de Carga:**
- Planilha de treino com volume/intensidade registrados
- Respeitar regra 10% (não aumentar >10% carga semanal)
- Deload programado a cada 4-6 semanas (40-50% volume)
- Monitorar sinais de overtraining (fadiga, dor persistente, queda de performance)

**Retorno ao Esporte Pós-Lesão:**

*Critérios Funcionais (todos devem ser atingidos):*
1. Amplitude de movimento completa e indolor
2. Força ≥90% do lado contralateral (ou baseline pré-lesão)
3. Testes funcionais >85% baseline:
   - Single leg hop for distance
   - Triple hop
   - Crossover hop
   - Y-balance test
4. Ausência de dor/edema após atividades progressivas

*Progressão de Retorno (exemplo para lesão de joelho):*
- Fase 1: Exercícios em cadeia fechada sem impacto (leg press, agachamento)
- Fase 2: Corrida em linha reta, baixa velocidade
- Fase 3: Corrida com mudanças de direção, acelerações graduais
- Fase 4: Treinamento específico do esporte, intensidade progressiva
- Fase 5: Retorno completo com monitoramento

**Abordagem Psicológica:**
- Kinesiofobia (medo de movimento) é comum pós-lesão
- Exposição gradual a movimentos "temidos"
- Reforço positivo, educação sobre segurança
- Considerar suporte psicológico se medo paralisante

**Monitoramento Contínuo:**
- Reavaliação funcional a cada 4-6 semanas
- Questionários de dor/função (VISA, LEFS, DASH conforme região)
- Ajuste de programa conforme resposta
- Atenção a sinais de recorrência (dor similar, edema, limitação)

**Colaboração Multidisciplinar:** Fisioterapeuta, educador físico, ortopedista, nutricionista (deficiências nutricionais podem predispor a lesões)."""
        }

    def generate_melhores_fases_atleticas_content(self, item_data: Dict) -> Dict:
        """Gera conteúdo para Melhores fases atléticas"""
        return {
            "clinicalRelevance": """Fases de alta performance atlética representam períodos de ótima convergência entre capacidade fisiológica, motivação psicológica, estrutura de treinamento adequada e ausência de lesões/doenças. A análise retrospectiva destas fases fornece insights sobre condições ideais que maximizam adaptações e podem ser replicadas ou adaptadas ao contexto atual.

Picos de performance esportiva variam conforme modalidade: esportes de resistência (maratona, ciclismo) atingem pico aos 27-35 anos quando VO2máx, economia de movimento e experiência tática convergem; esportes de potência (sprint, saltos) peak aos 23-28 anos devido a maior proporção de fibras tipo II e capacidade anaeróbica; esportes técnicos (tiro, golfe) podem manter alto nível até 35-45 anos pela compensação de declínio físico com experiência. Mulheres geralmente atingem pico 2-3 anos antes que homens.

Fatores fisiológicos associados a picos de performance incluem: maximização de adaptações específicas (hipertrofia muscular, densidade mitocondrial, economia de movimento), otimização de composição corporal (massa magra/gordura ideal para modalidade), maturação neuromuscular (coordenação intramuscular e intermuscular) e desenvolvimento psicológico (controle emocional, resiliência, autoeficácia).

A identificação de contextos que propiciaram melhores fases permite reconstrução estratégica de elementos-chave: estrutura de treino (volume, intensidade, periodização), suporte nutricional, recuperação adequada, ambiente social motivador, e propósito/objetivo claro. Estudos de carreira de atletas indicam que fases de pico não são randômicas, mas resultantes de 4-8 anos de progressão estruturada específica.""",

            "patientExplanation": """Todos nós temos fases na vida em que estamos "no nosso melhor" fisicamente - seja competindo em esportes, treinando regularmente, ou simplesmente nos sentindo muito bem e fortes. Olhar para trás e identificar esses momentos pode ensinar muito sobre o que funcionou bem para você.

Pergunte-se: naquela época, como eu treinava? Quanto tempo dedicava? Estava motivado? Tinha uma meta clara? Dormia bem? Me alimentava adequadamente? Tinha uma rotina consistente? Geralmente, essas fases boas não acontecem por acaso - são resultado de vários fatores alinhados ao mesmo tempo.

A boa notícia é que você pode recriar muitos desses elementos agora, adaptados à sua realidade atual. Não é sobre voltar a ser exatamente como era antes, mas sim entender os princípios que funcionaram e aplicá-los de forma realista hoje. Às vezes a melhor fase ainda está por vir, quando você souber combinar experiência com estratégia adequada.""",

            "conduct": """**Anamnese de Fases de Alta Performance:**
- Idade(s) em que se sentiu no melhor condicionamento físico
- Duração de cada fase (meses/anos)
- Esportes/atividades praticados naquele período
- Frequência e volume semanal de treino
- Tipo de treino (estruturado vs. livre, sozinho vs. grupo, com treinador?)
- Competições, provas, desafios realizados (resultados alcançados)
- Motivação predominante (intrínseca vs. extrínseca, objetivos claros?)
- Contexto de vida (estudante, início de carreira, período de menor stress)
- Sono, nutrição, recuperação (padrões na época)
- Suporte social (amigos treino, família, comunidade)
- Fim da fase: O que mudou? (lesão, mudança de cidade, trabalho, desmotivação)

**Análise de Fatores de Sucesso:**

**Elementos Estruturais:**
- Periodização (havia planejamento de longo prazo?)
- Consistência (treino regular vs. intermitente)
- Progressão (aumento gradual de desafios)
- Variedade (múltiplas modalidades vs. especialização)

**Elementos Motivacionais:**
- Objetivos específicos (prova, competição, estético)
- Prazer intrínseco na atividade
- Suporte social e senso de pertencimento
- Feedback positivo (resultados, reconhecimento)

**Elementos Contextuais:**
- Disponibilidade de tempo (horários flexíveis)
- Recursos (acesso a instalações, equipamentos, orientação)
- Baixo stress externo (trabalho, relações)
- Saúde física e mental estável

**Estratégias de Replicação Adaptada:**

**Recriar Estrutura de Treino:**
- Adaptar volume/intensidade da melhor fase à capacidade atual (ex: se treinava 10h/semana, iniciar com 5h)
- Implementar periodização similar (se funcionou antes, provavelmente funcionará novamente)
- Retomar modalidades que mais gostava (motivação intrínseca)

**Resgatar Motivação:**
- Estabelecer objetivos concretos similares (se era prova, buscar novos desafios)
- Reconstruir ambiente social (grupos de treino, comunidades esportivas)
- Criar rituais e rotinas que remetam a fase positiva

**Adaptar ao Contexto Atual:**
- Se disponibilidade de tempo menor: Otimizar eficiência (HIIT, treinos compostos, 30-45min/sessão)
- Se idade mais avançada: Maior ênfase em recuperação, mobilidade, prevenção de lesões
- Se objetivos mudaram: Adaptar treinamento (saúde vs. performance)

**Evitar Armadilhas:**
- *Nostalgia excessiva:* Não tentar replicar exatamente (contexto mudou)
- *Comparação negativa:* Celebrar progresso atual, não lamentar "perda"
- *Ignorar idade/lesões:* Respeitar limitações atuais (evitar re-lesões por impaciência)
- *Pressão excessiva:* Permitir que processo seja prazeroso, não apenas focado em resultado

**Construindo Nova Fase de Pico:**

*Princípios Universais de Fases de Alta Performance:*
1. **Consistência:** Treino regular (4-6x/semana), 12+ meses
2. **Progressão inteligente:** Aumento gradual, periodização
3. **Propósito claro:** Meta específica, prazo definido
4. **Recuperação adequada:** Sono 7-9h, nutrição, gerenciamento de stress
5. **Suporte social:** Comunidade, accountability
6. **Monitoramento:** Tracking de progresso, ajustes conforme feedback

**Planejamento de Nova Fase (12 meses):**
- Meses 1-3: Construção de base, estabelecimento de rotina
- Meses 4-6: Aumento de volume, introdução de especificidade
- Meses 7-9: Intensificação, preparação para pico
- Meses 10-12: Fase de pico (competição, desafio, avaliação)

**Monitoramento:**
- Comparar métricas com fase prévia (força, resistência, composição corporal)
- Tracking de aderência (% treinos realizados vs. planejados)
- Satisfação e prazer subjetivo (escala 1-10)
- Reavaliação trimestral, ajuste de estratégias conforme resposta

**Manutenção Pós-Pico:** Após atingir novo pico, estabelecer rotina de manutenção (80% do volume) para sustentar ganhos a longo prazo."""
        }

    def generate_modalidades_odiadas_content(self, item_data: Dict) -> Dict:
        """Gera conteúdo para Modalidades de esporte odiadas"""
        return {
            "clinicalRelevance": """A aversão a determinadas modalidades esportivas resulta de interações complexas entre experiências negativas prévias (trauma, humilhação, lesão), incompatibilidade biomecânica/fisiológica, fatores psicológicos (ansiedade, perfeccionismo) e preferências individuais de estímulo sensorial e social. Compreender estas aversões é fundamental para prescrição personalizada que maximize aderência, evitando imposição de modalidades que geram resistência psicológica e dropout (taxa de abandono 40-70% em 6 meses quando atividade não é prazerosa).

Aspectos neurofisiológicos da aversão incluem condicionamento aversivo clássico (associação entre modalidade e experiência negativa ativa amígdala e resposta de stress) e baixa recompensa dopaminérgica (atividades que não geram prazer têm menor liberação de dopamina no nucleus accumbens, reduzindo motivação intrínseca). Modalidades que exigem habilidades não desenvolvidas na infância (coordenação complexa, ritmo, equilíbrio dinâmico) geram frustração e evitação.

Aversões comuns incluem: corrida (monotonia, impacto articular, dispneia desconfortável em descondicionados), natação (desconforto aquático, necessidade técnica alta, exposição corporal), musculação (ambiente intimidador, comparação social, percepção de "não pertencimento"), esportes coletivos (ansiedade social, medo de julgamento, competitividade excessiva). Identificação precisa permite evitar gatilhos enquanto explora alternativas prazerosas.

Estratégia terapêutica foca em dessensibilização gradual (se aversão é objetivamente limitante) ou substituição por modalidades com benefícios fisiológicos equivalentes mas experiência subjetiva positiva. Por exemplo: aversão a corrida pode ser substituída por ciclismo, natação ou caminhadas rápidas (similares em gasto energético e adaptações cardiovasculares), mantendo aderência.""",

            "patientExplanation": """É completamente normal ter tipos de exercício que você simplesmente não suporta - e tudo bem! Não existe uma regra que diz que você tem que gostar de tudo. O importante é entender POR QUE você não gosta, para que possamos encontrar alternativas que você realmente goste e consiga manter.

Muitas vezes, a aversão vem de uma experiência ruim no passado: talvez tenha sido forçado a fazer aquilo na escola, foi humilhado, se lesionou, ou simplesmente descobriu que aquele tipo de movimento não combina com você. Outras vezes, é só uma questão de preferência pessoal - algumas pessoas adoram a solidão da corrida, outras acham entediante; alguns amam a competição dos esportes coletivos, outros sentem ansiedade.

O segredo é: NÃO se force a fazer algo que você odeia, porque você vai acabar desistindo. Em vez disso, vamos encontrar atividades que dão os mesmos benefícios, mas que você realmente gosta. Existem tantas formas de se movimentar que sempre há alternativas. Exercício precisa ser algo sustentável, não uma tortura.""",

            "conduct": """**Investigação de Aversões:**
- Modalidades/exercícios que o paciente evita ativamente ou tem experiência muito negativa
- Razões específicas para aversão:
  - *Experiencial:* Trauma, humilhação, lesão prévia
  - *Física:* Dor, desconforto excessivo, incompatibilidade biomecânica
  - *Psicológica:* Ansiedade, medo, tédio, falta de competência percebida
  - *Social:* Ambiente intimidador, comparação social, exposição
  - *Sensorial:* Monotonia, muito barulho, muito silêncio, temperatura

**Análise do Impacto:**
- A aversão é a modalidade única ou categoria ampla? (ex: odeia apenas corrida ou todo exercício aeróbico?)
- Limitação prática: A aversão impede objetivos importantes? (ex: precisa melhorar cardio mas odeia tudo cardiovascular?)
- Intensidade da aversão: Leve desinteresse vs. ansiedade intensa

**Estratégias de Manejo:**

**1. Substituição por Equivalente Prazeroso:**

*Aversão: Corrida*
- Alternativas cardiovasculares: Ciclismo, natação, remo, elíptico, dança, esportes coletivos, caminhada rápida em trilhas
- Equivalência: Todas melhoram VO2máx, queimam calorias, fortalecem sistema cardiovascular

*Aversão: Musculação em academia*
- Alternativas de força: Treino em casa (peso corporal, faixas elásticas), pilates, yoga power, escalada, CrossFit (se preferir ambiente comunitário), treino ao ar livre (barras, parques)

*Aversão: Natação*
- Alternativas baixo impacto: Ciclismo, elíptico, remo, hidroginástica (se desconforto é com técnica, não com água), yoga

*Aversão: Esportes coletivos*
- Alternativas: Atividades individuais (corrida, ciclismo, natação), esportes de raquete (menos pessoas), artes marciais (estrutura individual mesmo em grupo)

*Aversão: Yoga/Alongamento*
- Alternativas de mobilidade: Dança, tai chi, pilates dinâmico, treino funcional com ênfase em amplitude

**2. Dessensibilização Gradual (se aversão é limitante):**
- Útil quando modalidade é objetivamente benéfica mas há barreira psicológica superável
- Exemplo: Aversão a musculação por ansiedade social
  - Fase 1: Treino em casa 4 semanas (construir confiança)
  - Fase 2: Academia em horários vazios (6h ou 22h)
  - Fase 3: Sessões com personal (suporte social)
  - Fase 4: Gradualmente treinar em horários normais

**3. Reformulação da Experiência:**
- Mudança de contexto pode transformar aversão
- Exemplo: Odiava corrida em esteira → adora trail running (natureza, variabilidade)
- Exemplo: Odiava musculação solo → adora treino em grupo (CrossFit, bootcamp)

**4. Aceitação e Foco em Preferências:**
- Se aversão é genuína e há alternativas viáveis: simplesmente evitar e focar no que funciona
- Princípio: Aderência > Modalidade "ideal"
- 70% de aderência em atividade sub-ótima > 20% de aderência em atividade "perfeita"

**Prevenção de Novas Aversões:**
- Progressão gradual (evitar experiências muito difíceis/dolorosas que criam aversão)
- Reforço positivo constante
- Permitir experimentação sem pressão
- Respeitar sinais de desconforto psicológico (não apenas físico)

**Educação:**
- Não existe exercício universal obrigatório
- Diversidade de modalidades permite personalização total
- Mudança de preferências ao longo do tempo é normal (reavaliar periodicamente)

**Atenção a Evitação Patológica:**
- Se paciente "odeia tudo": Investigar barreiras mais profundas (depressão, trauma, ansiedade generalizada)
- Considerar suporte psicológico concomitante
- Iniciar com atividades de baixíssima ameaça (caminhadas curtas, movimentos suaves)

**Documentação:**
- Registrar aversões em prontuário para evitar prescrições inadequadas futuras
- Revisar anualmente (preferências podem mudar)

**Princípio Central:** Exercício sustentável é exercício prazeroso (ou ao menos tolerável). Respeitar aversões aumenta aderência e resultados a longo prazo."""
        }

    def generate_modalidades_preferidas_content(self, item_data: Dict) -> Dict:
        """Gera conteúdo para Modalidades de esporte preferidas"""
        return {
            "clinicalRelevance": """Modalidades esportivas preferidas refletem convergência entre aptidões físicas individuais, traços de personalidade, necessidades psicológicas e recompensas neurobiológicas específicas. A identificação e priorização de atividades prazerosas aumenta aderência em 200-300% comparado a prescrições genéricas, sendo o fator preditivo mais forte de manutenção de exercício em longo prazo (>12 meses).

Preferências esportivas correlacionam-se com perfis psicológicos: indivíduos com alta necessidade de autonomia preferem atividades individuais (corrida, ciclismo, natação); alta necessidade de afiliação social prefere esportes coletivos; alta sensação-seeking prefere atividades de risco/aventura (escalada, mountain bike, esportes radicais); alta orientação a competitividade prefere modalidades com rankings/performance mensurável.

Neurobiologicamente, atividades preferidas induzem maior liberação de dopamina no sistema de recompensa (nucleus accumbens, área tegmental ventral), endorfinas (analgesia natural e euforia) e endocanabinoides (runner's high), criando associações positivas que reforçam comportamento. Estudos de neuroimagem demonstram que exercícios prazerosos ativam córtex pré-frontal ventromedial (processamento de valor hedônico) e reduzem ativação de ínsula anterior (processamento de esforço negativo).

Fisiologicamente, indivíduos apresentam aptidões variadas: predominância de fibras tipo I (resistência) vs. tipo II (potência), VO2máx geneticamente influenciado, flexibilidade articular, relação potência/peso. Alinhamento entre aptidões e modalidade escolhida aumenta percepção de competência (autoeficácia), fundamental para motivação intrínseca sustentada. Contudo, desenvolvimento de competências através de prática pode transformar atividades inicialmente neutras em preferidas.""",

            "patientExplanation": """As atividades físicas que você mais gosta são seu maior trunfo para manter uma vida ativa e saudável. Quando você pratica algo que realmente gosta, seu cérebro libera substâncias que fazem você se sentir bem, criar memórias positivas e querer repetir a experiência. É por isso que você consegue manter por anos algo que ama, mas desiste em semanas de algo que odeia.

Suas preferências não são aleatórias - elas dizem muito sobre você. Se você gosta de correr sozinho, talvez valorize momentos de introspecção e autonomia. Se prefere esportes coletivos, provavelmente a conexão social é importante para sua motivação. Se gosta de desafios físicos intensos, pode ter uma personalidade mais aventureira. Não há certo ou errado, apenas o que funciona para VOCÊ.

O ideal é construir sua rotina de exercícios baseada principalmente no que você gosta, complementando com algumas atividades que são importantes para saúde mesmo que não sejam suas favoritas. Por exemplo: se você ama correr mas precisa ganhar força muscular, pode fazer musculação 2x/semana e correr 3x. Assim você se mantém motivado e ainda trabalha todas as áreas importantes.""",

            "conduct": """**Investigação Detalhada de Preferências:**

**Modalidades Atuais/Históricas Preferidas:**
- Quais atividades você mais gosta ou gostava de praticar?
- O que especificamente você gosta nelas? (movimento, ambiente, social, sensação física, desafio mental)
- Frequência ideal que praticaria se não houvesse barreiras
- Nível de habilidade/competência percebida (iniciante, intermediário, avançado)

**Categorização de Preferências:**

*Por Características Físicas:*
- Aeróbico vs. Força vs. Flexibilidade vs. Misto
- Alto impacto vs. Baixo impacto
- Contínuo (steady state) vs. Intervalado
- Controlado vs. Imprevisível/Reativo

*Por Contexto Social:*
- Individual vs. Coletivo vs. Duplas
- Competitivo vs. Cooperativo vs. Não-competitivo
- Instruído (aulas) vs. Livre

*Por Ambiente:*
- Indoor vs. Outdoor
- Natureza vs. Urbano vs. Academia
- Silencioso vs. Musical vs. Social

*Por Experiência Sensorial:*
- Ritmo/Música (dança, spinning) vs. Sem música
- Meditativo/Mindful (yoga, tai chi) vs. Alta intensidade
- Técnico/Precisão vs. Intuitivo/Livre

**Análise Motivacional:**
- Motivação predominante: Saúde, estética, performance, social, prazer intrínseco, redução de stress?
- Qual aspecto é mais recompensador: Progresso mensurável, sensação durante, sensação pós, socialização, competição?

**Prescrição Centrada em Preferências:**

**Estratégia 80/20:**
- 80% do volume de treino em atividades preferidas (garante aderência)
- 20% em atividades complementares importantes (endereça gaps fisiológicos)

**Exemplo 1: Preferência por corrida**
- *Base (80%):* Corrida outdoor 4-5x/semana (variando intensidades)
- *Complemento (20%):* Fortalecimento (2x/semana, 30min) para prevenir lesões, mobilidade (10min pós-corrida)
- *Resultado:* Paciente adora a maior parte do treino, tolera necessário

**Exemplo 2: Preferência por esportes coletivos (futebol)**
- *Base (80%):* Futebol recreativo 3x/semana
- *Complemento (20%):* Treino de força/potência 2x/semana (melhora performance no futebol), flexibilidade
- *Resultado:* Complemento serve ao prazer principal

**Exemplo 3: Preferência por yoga/atividades mente-corpo**
- *Base (80%):* Yoga 4-5x/semana
- *Complemento (20%):* Caminhadas vigorosas 2x/semana (cardiovascular), treino funcional 1x (força não desenvolvida em yoga)

**Maximização de Prazer nas Preferências:**
- Otimizar contexto (se gosta de correr outdoor, não forçar esteira)
- Variedade dentro da preferência (se gosta de musculação: variar exercícios, métodos, rep ranges)
- Progressão adequada (desafio suficiente para engajamento, mas não frustração)
- Social: Conectar com comunidades da modalidade preferida (grupos de corrida, clubes, aulas)

**Expansão Gradual de Preferências:**
- Exposição a novas modalidades relacionadas (se gosta de ciclismo, tentar mountain bike, spinning, gravel)
- Desenvolvimento de competências pode criar novas preferências (aulas de natação podem transformar natação de "difícil" para "prazerosa")
- Reavaliar preferências anualmente (mudanças são normais)

**Adaptação a Fases de Vida:**
- Preferências podem precisar adaptação temporária (lesão, gravidez, mudança de cidade)
- Buscar versões adaptadas da preferência (se corria mas lesionou: ciclismo/natação temporariamente)
- Manter conexão emocional com modalidade preferida mesmo se modificada

**Barreiras a Preferências:**
- Identificar obstáculos práticos (custo, acesso, tempo, clima)
- Resolução criativa de problemas:
  - Custo alto (natação): Buscar piscinas públicas, programas comunitários
  - Clima (ciclismo outdoor): Rolo indoor, spinning, roupas adequadas
  - Tempo limitado: Sessões mais curtas mas frequentes, proximidade de casa

**Educação:**
- Reforçar importância de prazer para sustentabilidade
- Validar preferências individuais (não há hierarquia de modalidades)
- Encorajar exploração de novas atividades sem pressão

**Monitoramento:**
- Satisfação subjetiva (escala 1-10 pós-treino)
- Aderência (% sessões planejadas vs. realizadas)
- Sinais de tédio ou desmotivação (indicam necessidade de variação)
- Reavaliação trimestral: Preferências mudaram? Novas modalidades a explorar?

**Princípio Fundamental:** Prescrição eficaz alinha-se com preferências individuais, não com dogmas de "melhor exercício". Exercício preferido realizado > Exercício "ótimo" abandonado."""
        }

    def generate_onde_como_faz_content(self, item_data: Dict) -> Dict:
        """Gera conteúdo para Onde e como faz exercício"""
        return {
            "clinicalRelevance": """O contexto de realização do exercício (local, modalidade de prática, suporte profissional) influencia significativamente aderência, segurança, eficácia e satisfação. Estudos de aderência demonstram que taxas de manutenção variam drasticamente conforme setting: treino supervisionado por profissional (68-75% aderência 6 meses), treino em grupo/aulas (55-65%), treino individual em academia (35-45%), treino em casa sem orientação (20-30%).

Academias oferecem vantagens de equipamentos diversificados, ambiente social motivador e facilidade de progressão de carga, mas apresentam barreiras de custo (R$80-300/mês), deslocamento (média 20-30min), intimidação social (gym anxiety em 50% de iniciantes) e horários fixos. Treino domiciliar elimina barreiras logísticas e financeiras, mas requer maior autodisciplina, tem limitações de equipamento e ausência de feedback técnico imediato, aumentando risco de técnica inadequada e lesões (15-20% maior em não-supervisionados).

Treino ao ar livre (parques, praias, trilhas) combina benefícios do exercício com exposição à natureza, reduzindo cortisol adicional (-20-25% vs. indoor), melhorando humor e aumentando aderência (65-70% manutenção). Modalidades outdoor incluem corrida, calistenia (barras, exercícios corporais), ciclismo e esportes recreativos. Limitações incluem dependência climática e disponibilidade de espaços seguros.

Supervisão profissional (personal trainer, educador físico) aumenta eficácia através de periodização adequada, correção técnica em tempo real, ajuste individualizado e accountability. Meta-análises indicam ganhos 30-40% superiores em força e hipertrofia com supervisão vs. autogerenciado, além de 60% menor incidência de lesões. Modalidades grupais (CrossFit, spinning, bootcamps) adicionam componente de coesão social e competição amigável, potentes motivadores extrínsecos.""",

            "patientExplanation": """Onde e como você faz exercício é quase tão importante quanto o exercício em si, porque isso afeta se você vai conseguir manter a rotina ou não. Cada opção tem vantagens e desvantagens, e o ideal depende da sua personalidade, recursos e objetivos.

Academia é ótima se você gosta de variedade de equipamentos, precisa do ambiente social para se motivar, ou se beneficia de ter um local "dedicado" ao treino. Treino em casa é perfeito se você valoriza praticidade, tem pouco tempo ou budget limitado, mas exige mais disciplina. Treino ao ar livre é maravilhoso para quem gosta de natureza e quer combinar exercício com bem-estar mental, mas depende do clima.

Ter orientação profissional (personal, educador físico) faz enorme diferença nos resultados e na segurança, especialmente no começo. Mas se não for viável financeiramente, existem alternativas como aulas em grupo (mais baratas), aplicativos bem estruturados, ou aprender com profissionais em períodos curtos para depois continuar sozinho com base sólida. O importante é escolher um formato que você consiga sustentar.""",

            "conduct": """**Avaliação Detalhada do Contexto Atual:**

**Onde treina atualmente?**
- Academia (qual tipo: tradicional, CrossFit, boutique)
- Casa (estrutura: sem equipamento, equipamento básico, home gym completo)
- Ao ar livre (parques, ruas, trilhas, praias)
- Estúdios especializados (pilates, yoga, artes marciais)
- Locais de trabalho (ginástica laboral, academia corporativa)
- Combinação de locais

**Como treina?**
- Com supervisão profissional (personal 1:1, semi-personal, coach)
- Aulas em grupo (spinning, crossfit, bootcamp, dança)
- Sozinho seguindo programa prescrito
- Sozinho sem programa definido (improvisa)
- Com amigos/parceiros de treino
- Aplicativos/programas online

**Análise de Adequação:**

**Compatibilidade com Objetivos:**
- Iniciantes: Necessitam supervisão ou estrutura clara (academia com orientação, aulas, personal)
- Objetivos específicos (hipertrofia, força máxima): Requerem equipamentos e periodização (academia ou home gym bem equipado)
- Saúde geral: Flexível (qualquer setting funciona se consistente)
- Reabilitação/limitações: Supervisão essencial inicialmente

**Compatibilidade com Perfil:**
- Personalidade autodisciplinada: Casa/outdoor funciona bem
- Necessita accountability externo: Academia, aulas, personal
- Ansiedade social: Casa ou outdoor inicialmente, transição gradual
- Alta necessidade de socialização: Aulas em grupo, esportes coletivos, academias

**Barreiras e Facilitadores:**
- Tempo: Casa/outdoor (sem deslocamento) vs. Academia (deslocamento 20-40min)
- Custo: Casa (R$0-500 investimento inicial) vs. Academia (R$80-300/mês) vs. Personal (R$100-300/sessão)
- Acesso: Proximidade de academia/parques, segurança do bairro
- Clima: Dependência de outdoor em regiões com clima instável

**Otimização do Setting Atual:**

**Se treina em academia:**
- Maximizar uso de recursos (aulas inclusas, avaliações, apps da academia)
- Considerar horários menos lotados se gym anxiety
- Explorar zonas diferentes (funcional, livre, máquinas)
- Buscar orientação gratuita se disponível (plantonista, avaliação)

**Se treina em casa:**
- *Equipamento mínimo efetivo:*
  - Faixas elásticas (R$50-150): Resistência variada, portátil
  - Halteres ajustáveis (R$300-800): Progressão de carga
  - Barra/Anilhas (R$400-1500): Permite treino completo
  - TRX/Suspension (R$200-400): Versátil, baixo espaço
  - Tapete yoga (R$80-200): Mobilidade, core
- *Estruturação:* Seguir programa online bem avaliado ou contratar prescrição profissional (consulta mensal)
- *Accountability:* Parceiro treino virtual, grupos online, tracking apps
- *Espaço:* Dedicar área específica (associação contextual aumenta consistência)

**Se treina ao ar livre:**
- *Locais:* Mapear parques com barras, academias ao ar livre, circuitos de corrida/bike
- *Equipamento:* Faixas elásticas portáteis, app de treino, relógio GPS
- *Plano B:* Alternativa para dias de chuva (casa, academia pay-per-use)
- *Segurança:* Horários adequados, não treinar isolado, atenção a ambiente

**Supervisão Profissional:**

**Quando é essencial:**
- Iniciantes absolutos (evitar lesões, aprender técnicas fundamentais)
- Reabilitação pós-lesão ou condições médicas
- Objetivos competitivos ou avançados
- Histórico de múltiplas tentativas falhas (necessita accountability externo)

**Modelos de Supervisão:**
- *Personal 1:1:* Máxima personalização, mais caro (R$100-300/sessão)
- *Semi-personal:* Pequenos grupos 2-5 pessoas, custo intermediário (R$50-100/sessão)
- *Consultoria mensal:* 1 sessão/mês para prescrição + execução independente (R$200-400/mês)
- *Aulas em grupo:* Estrutura e supervisão, social, custo menor (R$20-80/aula)
- *Online coaching:* Prescrição remota + check-ins virtuais (R$100-500/mês)

**Transições Estratégicas:**

**Casa → Academia:**
- Motivação: Necessidade de progressão, tédio de equipamento limitado, desejo de socialização
- Estratégia: Período teste (day pass, plano mensal), horários menos lotados inicialmente

**Academia → Casa:**
- Motivação: Custo, tempo, pandemia, preferência por autonomia
- Estratégia: Investimento gradual em equipamentos, manter supervisão remota inicialmente

**Não-supervisionado → Supervisionado:**
- Motivação: Estagnação, lesão, desejo de otimizar resultados
- Estratégia: Contratar 4-8 sessões para "reset" (aprender técnicas, receber programa), depois avaliar necessidade

**Modelo Híbrido (ideal para muitos):**
- Exemplo 1: Academia 3x/semana (força supervisionada) + Outdoor 2x/semana (cardio autogerenciado)
- Exemplo 2: Aulas grupo 2x/semana (social, motivacional) + Casa 2x/semana (complementar, conveniência)
- Exemplo 3: Personal 1x/semana (técnica, prescrição) + Independente 3-4x/semana (execução programa)

**Fatores de Decisão Principais:**
1. Aderência esperada (qual setting você realisticamente manterá?)
2. Custo-benefício (investimento justifica retorno em resultados/saúde?)
3. Segurança (competência técnica suficiente se não-supervisionado?)
4. Compatibilidade com rotina (logística viável?)

**Monitoramento:**
- Aderência (% sessões planejadas realizadas) - se <70%, reavaliar setting
- Progressão (ganhos estagnados podem indicar necessidade de supervisão/equipamento)
- Satisfação (prazer subjetivo no setting atual)
- Reavaliação semestral: Setting atual ainda é ótimo ou mudanças de vida exigem adaptação?

**Princípio Central:** Melhor setting é aquele que você consegue manter consistentemente com segurança, não necessariamente o "ideal teórico"."""
        }

    def generate_sedentarismo_content(self, item_data: Dict) -> Dict:
        """Gera conteúdo para Piores fases de sedentarismo"""
        return {
            "clinicalRelevance": """Períodos prolongados de sedentarismo induzem descondicionamento físico profundo e reversão de adaptações prévias (detraining), com consequências metabólicas, cardiovasculares, musculoesqueléticas e psicológicas. Após 2-4 semanas de inatividade, observa-se redução de 5-10% em VO2máx, diminuição de 10-15% em força muscular e aumento de resistência à insulina (+25-30%). Sedentarismo >3 meses resulta em perda de 20-30% da massa muscular em indivíduos previamente treinados, redução de densidade mitocondrial (-40-50%) e aumento de marcadores inflamatórios (IL-6, TNF-α, PCR).

Fatores precipitantes de fases sedentárias incluem: transições de vida (entrada na vida profissional, casamento, nascimento de filhos), mudanças geográficas (perda de redes sociais/infraestrutura), lesões/cirurgias prolongadas, doenças crônicas, períodos de alto stress ocupacional, e quadros de saúde mental (depressão, ansiedade). Identificação de padrões históricos permite antecipação e desenvolvimento de estratégias preventivas para futuros períodos de risco.

Consequências psicológicas do sedentarismo incluem redução de autoeficácia física, aumento de sintomas depressivos (50-70% maior em sedentários vs. ativos), piora de qualidade do sono e redução de resiliência ao stress. A inatividade reduz produção de BDNF, neurogênese hipocampal e conectividade funcional em redes de modo padrão, contribuindo para declínio cognitivo acelerado.

Retorno à atividade após sedentarismo prolongado requer cuidado especial: risco aumentado de lesões musculoesqueléticas (especialmente primeiras 4-6 semanas), eventos cardiovasculares agudos em descondicionados (avaliação médica pré-participação essencial), e frustração psicológica pela perda de capacidades prévias. Progressão deve ser 50-70% mais conservadora que em iniciantes sem histórico ativo.""",

            "patientExplanation": """Todos nós temos fases da vida em que paramos de nos exercitar - e isso é normal. Pode ser por causa do trabalho muito intenso, nascimento de filho, mudança de cidade, lesão, doença, ou simplesmente porque a vida ficou muito corrida e o exercício foi a primeira coisa a sair da lista. Não há motivo para culpa, mas é importante entender o que aconteceu para evitar que se repita.

Quando você fica muito tempo parado, seu corpo perde força, resistência e flexibilidade - e quanto mais tempo parado, mais difícil é voltar. Mas o corpo tem uma "memória": se você já foi ativo antes, voltar é mais rápido do que seria para alguém que nunca treinou. A chave é recomeçar com paciência, sem tentar voltar imediatamente ao nível anterior (isso leva a lesões).

Olhar para trás e identificar o que causou essas fases de sedentarismo ajuda a criar estratégias para não cair na mesma armadilha novamente. Por exemplo: se você sempre para quando o trabalho fica intenso, talvez precise de treinos curtos de 20-30 minutos em vez de sessões longas. Se para quando se lesiona, talvez precise trabalhar prevenção e ter planos alternativos.""",

            "conduct": """**Anamnese de Fases Sedentárias:**

**Identificação dos Períodos:**
- Idades/anos em que esteve predominantemente sedentário (<150 min atividade/semana)
- Duração de cada fase (meses/anos)
- Nível de sedentarismo (sedentarismo completo vs. atividade mínima esporádica)
- Consequências percebidas (ganho de peso, perda de força, piora de humor, doenças)

**Fatores Precipitantes:**
- *Profissionais:* Entrada no mercado de trabalho, promoção com maior carga, múltiplos empregos
- *Familiares:* Casamento, nascimento de filhos, cuidado de familiares doentes
- *Geográficos:* Mudança de cidade/país, perda de infraestrutura (academia fechou)
- *Saúde:* Lesões, cirurgias, doenças agudas ou crônicas, gravidez
- *Psicológicos:* Depressão, ansiedade, burnout, perda de motivação
- *Sociais:* Perda de parceiros de treino, isolamento social, mudança de círculo social
- *Financeiros:* Perda de emprego, redução de renda

**Fatores Mantenedores:**
- O que impediu retorno: Falta de tempo, baixa energia, dor crônica, vergonha de condição atual, falta de conhecimento
- Tentativas de retorno frustradas? (recaídas, lesões ao retomar)

**Consequências das Fases:**
- Físicas: Ganho de peso (+__kg), perda de massa muscular, dores articulares, diagnóstico de doenças (DM2, HAS, dislipidemia)
- Mentais: Piora de humor, ansiedade, qualidade de sono
- Funcionais: Limitações em AVDs (subir escadas, carregar peso)

**Análise de Padrões:**
- Existe padrão repetitivo? (sempre para quando stress aumenta, sempre para no inverno)
- Fatores comuns às fases sedentárias
- Sinais precoces de início de fase sedentária (permite intervenção preventiva)

**Estratégias de Prevenção de Recaídas:**

**Identificação de Gatilhos de Risco:**
- Listar situações de vida que historicamente levaram a sedentarismo
- Desenvolver planos de contingência específicos para cada gatilho

**Planos de Contingência por Gatilho:**

**Gatilho: Aumento de carga de trabalho**
- *Estratégia:* Treinos curtos de alta intensidade (20-30min, 3x/semana)
- *Implementação:* HIIT em casa, caminhar no horário de almoço, treino às 6h antes do trabalho
- *Mindset:* "Algo é sempre melhor que nada" - 20min/dia mantém base

**Gatilho: Nascimento de filho/cuidado de dependente**
- *Estratégia:* Flexibilidade total, treinos interruptíveis em casa
- *Implementação:* Equipamento em casa, sessões de 10-15min múltiplas vezes, incluir bebê/criança em atividades (caminhadas com carrinho)
- *Suporte:* Negociar com cônjuge janelas de treino, usar período de sono do dependente

**Gatilho: Lesão/Doença**
- *Estratégia:* Manter atividade adaptada em partes não afetadas
- *Implementação:* Lesão de perna → treino de MS e core, lesão de braço → caminhadas/bicicleta
- *Colaboração:* Fisioterapeuta + educador físico para prescrição segura durante recuperação

**Gatilho: Mudança de cidade**
- *Estratégia pré-mudança:* Pesquisar infraestrutura de exercício no novo local (academias, parques, grupos)
- *Implementação:* Priorizar escolher moradia próxima a locais de atividade, conectar-se com comunidades esportivas locais nas primeiras semanas

**Gatilho: Falta de motivação/Depressão**
- *Estratégia:* Exercício como tratamento não-negociável, suporte profissional
- *Implementação:* Compromissos sociais (treino com amigo - não cancelar), acompanhamento psicológico concomitante, iniciar com atividades de prazer (não obrigação)
- *Dose mínima:* 10min caminhada ao ar livre diariamente como baseline absoluto

**Gatilho: Inverno/Clima adverso**
- *Estratégia:* Alternativas indoor preparadas antes da estação
- *Implementação:* Equipamento em casa, assinatura de academia indoor, aulas online, roupas adequadas para outdoor

**Sistema de Alerta Precoce:**
- Identificar sinais de início de declínio (faltou 2 treinos seguidos, não treinou há 1 semana)
- Protocolo automático quando sinal aparece:
  - Semana 0 atividade: Acionar plano mínimo absoluto (3x 15min o que for possível)
  - 2 semanas <50% frequência: Reavaliar barreiras, ajustar plano, considerar suporte profissional/social

**Dose Mínima Efetiva (DME):**
- Definir baseline absoluto que mantém condicionamento mínimo mesmo em fases difíceis
- Exemplo: 75 min/semana atividade moderada (3x 25min) como "não-negociável"
- Conceito: Melhor manter 30% do treino ideal durante fase difícil do que abandonar 100% e ter que recomeçar do zero

**Retorno Pós-Fase Sedentária:**

**Avaliação Inicial:**
- Duração do sedentarismo (< 3 meses vs. 3-12 meses vs. >1 ano)
- Condição atual (composição corporal, capacidade cardiorrespiratória básica, força)
- Barreiras residuais (fatores que causaram sedentarismo ainda presentes?)
- Avaliação médica se: >6 meses sedentário + >40 anos + fator de risco cardiovascular

**Protocolo de Retorno Gradual:**

*Sedentarismo 3-6 meses:*
- Semanas 1-2: 40-50% do volume prévio, intensidade baixa-moderada (RPE 4-6/10)
- Semanas 3-4: 60-70% volume
- Semanas 5-8: 80-90% volume
- Semana 9+: Volume normal progressivamente

*Sedentarismo 6-12 meses:*
- Semanas 1-4: 30-40% volume prévio, iniciar com movimentos fundamentais
- Semanas 5-8: 50-60% volume
- Semanas 9-16: 70-90% volume
- Meses 4-6: Retorno completo

*Sedentarismo >12 meses:*
- Tratar como iniciante (evitar comparação com capacidades passadas)
- Programa de 6-12 meses progressão cuidadosa
- Ênfase em reconstrução de base, não retorno rápido

**Prevenção de Lesões no Retorno:**
- Warm-up prolongado (10-15min vs. 5min usual)
- Evitar cargas máximas/intensidades extremas primeiros 2 meses
- Atenção especial a tendões (adaptam lentamente, risco de tendinites)
- Fortalecimento de core e estabilizadores antes de progressões complexas

**Gestão Psicológica:**
- Normalizar frustração com perda de capacidades
- Celebrar progresso atual (não lamentar o passado)
- Memória muscular: Recuperação é 2-3x mais rápida que construção inicial
- Foco em processo (consistência) vs. resultado (performance imediata)

**Monitoramento de Risco:**
- Check-in semanal: Frequência de treino se mantendo?
- Identificação precoce de declínio: Intervenção imediata
- Reavaliação trimestral de barreiras e ajuste de estratégias

**Princípio Central:** Prevenir recaídas é mais fácil que reiniciar. Manter atividade mínima em fases difíceis > Parar completamente e ter que reconstruir."""
        }

    def generate_provas_desafios_content(self, item_data: Dict) -> Dict:
        """Gera conteúdo para Provas/desafios planejados"""
        return {
            "clinicalRelevance": """Objetivos concretos e temporalmente definidos (provas, competições, desafios) constituem poderosos motivadores extrínsecos que aumentam aderência, intensidade de treino e consistência. A Teoria da Autodeterminação (Deci & Ryan) indica que metas específicas aumentam percepção de competência e autonomia, componentes essenciais de motivação intrínseca sustentada. Estudos demonstram que indivíduos treinando para evento específico apresentam 70-85% aderência vs. 35-50% em treino sem objetivo definido.

A preparação para desafios estrutura periodização natural: fase de base (construção de capacidade aeróbica/força), fase específica (treino de modalidade/distância do evento), fase de afilamento (taper - redução de volume 40-60% mantendo intensidade 7-14 dias pré-evento para supercompensação). Este ciclo otimiza adaptações fisiológicas e previne monotonia de treino não-direcionado.

Fisiologicamente, a antecipação de competição ativa sistema de stress controlado: aumento de catecolaminas (adrenalina, noradrenalina), cortisol transitório (mobilização energética), e elevação de testosterona (comportamento competitivo), criando estado de alerta ótimo. Pós-evento, ativação do sistema de recompensa (dopamina, endorfinas) reforça comportamento, especialmente se objetivo alcançado, estabelecendo ciclo motivacional positivo.

Desafios graduados permitem progressão mensurável de capacidades: 5k → 10k → meia-maratona → maratona em corrida; competições de powerlifting para avaliar força; travessias de natação; competições de CrossFit. Cada nível atingido aumenta autoeficácia e estabelece novo baseline, permitindo objetivos progressivamente ambiciosos. Falha em atingir objetivo, se bem gerida psicologicamente, oferece feedback valioso para ajuste de treino futuro.""",

            "patientExplanation": """Ter uma prova, competição ou desafio marcado no calendário é como ter um destino claro em uma viagem - você sabe para onde está indo e pode planejar o caminho. Pessoas que treinam para algo específico costumam ser muito mais consistentes e motivadas do que quem treina "só por treinar".

O desafio não precisa ser algo extremo como uma maratona ou competição profissional. Pode ser uma corrida de 5km, uma trilha com amigos, um torneio recreativo, ou até um objetivo pessoal como fazer 10 flexões ou correr 30 minutos sem parar. O importante é que seja algo que te desafie (mas seja realista), tenha uma data definida e que você possa se preparar para alcançar.

Quando você tem um objetivo assim, cada treino ganha propósito - você não está apenas "fazendo exercício", está se preparando para conquistar algo. E quando alcança, a sensação de realização é incrível e te motiva a buscar o próximo desafio. É uma forma de transformar exercício de uma obrigação em uma jornada empolgante.""",

            "conduct": """**Investigação de Objetivos Futuros:**

**Desafios Planejados nos Próximos 6 Meses:**
- Eventos específicos: Nome, data, local, formato (prova, competição, desafio pessoal)
- Distância/modalidade/objetivo específico (5k, 10k, maratona, triathlon, competição de crossfit, etc.)
- Nível de experiência: Primeira vez vs. Repetição de evento prévio
- Motivação: Por que escolheu este desafio? (intrínseca vs. social vs. saúde)
- Preparação atual: Já iniciou treino específico? Quanto tempo até o evento?

**Se NÃO tem desafios planejados:**
- Interesse em estabelecer algum objetivo futuro?
- Barreiras: Falta de conhecimento, medo de não conseguir, custo, disponibilidade
- Histórico: Já participou de eventos/desafios no passado? Experiência positiva ou negativa?

**Benefícios de Estabelecer Objetivos:**

**Motivacionais:**
- Propósito claro para cada treino
- Accountability (compromisso público)
- Estrutura temporal (treino progressivo com deadline)
- Senso de progresso mensurável

**Fisiológicos:**
- Periodização estruturada otimiza adaptações
- Intensidade de treino geralmente maior (esforço direcionado)
- Progressão de capacidades de forma sistematizada

**Psicológicos:**
- Aumento de autoeficácia ao atingir objetivos
- Superação de limites auto-impostos
- Experiência de flow durante evento
- Construção de identidade ("sou corredor", "sou atleta")

**Seleção de Desafio Apropriado:**

**Critérios para Objetivo SMART:**
- *Specific (Específico):* "Completar corrida de 10k" vs. "Melhorar condicionamento"
- *Measurable (Mensurável):* Distância, tempo, colocação, carga levantada
- *Achievable (Alcançável):* Desafiador mas realista dado tempo de preparação e nível atual
- *Relevant (Relevante):* Alinhado com valores e preferências (não fazer maratona se odeia correr)
- *Time-bound (Temporal):* Data específica (cria urgência produtiva)

**Progressão de Desafios por Nível:**

**Iniciantes (0-6 meses treino):**
- Corrida: 5km (após 8-12 semanas treino)
- Força: Competição interna de academia, estabelecer 1RM
- Funcional: Completar WOD (workout of the day) RX (sem modificações)
- Geral: 30 dias consecutivos de atividade física, atingir x% gordura corporal

**Intermediários (6-24 meses):**
- Corrida: 10km, meia-maratona (após 12-16 semanas treino específico)
- Triathlon: Sprint distance (750m natação, 20km bike, 5km corrida)
- Força: Competição local de powerlifting, atingir total específico (soma agachamento + supino + deadlift)
- Funcionais: Competições de CrossFit escalões iniciante/intermediário
- Aventura: Trilhas longas (15-25km), cicloturismo (50-100km)

**Avançados (>24 meses):**
- Corrida: Maratona, ultramaratonas
- Triathlon: Olímpico, Ironman 70.3, Ironman
- Força: Competições nacionais, atingir standards elite
- Aventura: Expedições multi-dia, corridas de montanha

**Desenvolvimento de Plano de Preparação:**

**Estrutura Temporal:**
- Determinar tempo até evento (ideal: 12-20 semanas para primeira meia-maratona, 20-24 para maratona)
- Dividir em fases: Base → Específico → Afilamento

**Exemplo: Preparação para 10km (12 semanas, intermediário)**

*Fase Base (semanas 1-4):*
- Volume: 20-30km/semana
- Treinos: 3-4x corrida fácil 30-45min, 1x interval curto (6-8x 400m)
- Complementar: 2x força de suporte (perna, core)

*Fase Específica (semanas 5-10):*
- Volume: 30-40km/semana
- Treinos: 1x long run progressivo (até 12-14km), 1x tempo run (3-5km ritmo de prova), 1x intervals (5-6x 1km), 2x corrida fácil
- Complementar: 1-2x força manutenção

*Fase Afilamento (semanas 11-12):*
- Volume: Reduzir 40-50% (20-25km/semana)
- Manter intensidade mas reduzir volume de intervalos
- Foco em recuperação, sono, nutrição

**Ajustes por Modalidade:**
- Desafios de força: Ênfase em progressão de carga, menor ênfase cardiovascular
- Desafios técnicos (natação, escalada): Maior frequência de prática técnica (5-6x/semana)
- Desafios de endurance extrema (ultra, Ironman): Volume muito alto, múltiplas sessões/dia

**Gestão de Expectativas:**
- Tempo de conclusão realista baseado em treinos (pace médio em long runs + 5-10%)
- Primeiros eventos: Foco em completar, não em tempo
- Eventos subsequentes: Objetivos progressivos (melhorar tempo, colocação)

**Estratégia de Prova:**
- Pacing: Começar conservador (5-10% mais lento que ritmo alvo primeiros 30-40% da prova)
- Nutrição/Hidratação: Estratégia testada em treinos longos (não inovar no dia)
- Mental: Dividir prova em segmentos gerenciáveis, mantras para momentos difíceis

**Pós-Evento:**
- Recuperação ativa 1-2 semanas (volume muito reduzido, intensidade leve)
- Reflexão: O que funcionou? O que melhorar?
- Próximo objetivo: Estabelecer novo desafio (mantém motivação) ou período de treino livre (evitar burnout)

**Gestão de Falha em Objetivo:**
- Normalizar: Fatores incontroláveis (clima, dia ruim, doença)
- Aprendizado: Ajuste de preparação futura
- Redefinir: Tentar novamente ou ajustar expectativa (objetivo era agressivo demais?)
- Evitar: Desistir de futuros objetivos por uma falha

**Considerações Especiais:**

**Financeiras:**
- Inscrição de provas (R$50-500 dependendo do evento)
- Transporte/hospedagem se evento fora da cidade
- Equipamentos específicos (tênis, roupa, equipamentos técnicos)

**Logísticas:**
- Disponibilidade no dia do evento (família, trabalho)
- Viagem e planejamento se evento distante

**Sociais:**
- Envolver amigos/família (espectadores, participar juntos)
- Compartilhar preparação (accountability, suporte)

**Alternativas a Eventos Formais (se barreiras a competições):**
- Desafios pessoais (correr x km no mês, completar programa específico)
- Desafios virtuais (apps com badges, desafios online)
- Grupos de treino com objetivos compartilhados
- Auto-competição (bater PRs pessoais em treinos)

**Monitoramento:**
- Progressão no treino específico (tempos melhorando? volume tolerado?)
- Fadiga e sinais de overtraining (ajustar se necessário)
- Entusiasmo e motivação (se declinar, reavaliar objetivo ou estratégia)

**Princípio Central:** Objetivos concretos transformam treino de obrigação em jornada com propósito. Selecionar desafios realistas mas estimulantes maximiza motivação e resultados."""
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
        if "Histórico familiar" in item_name:
            content = self.generate_historico_familiar_content(item_data)
        elif "Horários" in item_name:
            content = self.generate_horarios_content(item_data)
        elif "Infância" in item_name:
            content = self.generate_infancia_content(item_data)
        elif "Lesões" in item_name:
            content = self.generate_lesoes_exercicio_content(item_data)
        elif "Melhores fases" in item_name:
            content = self.generate_melhores_fases_atleticas_content(item_data)
        elif "odiadas" in item_name:
            content = self.generate_modalidades_odiadas_content(item_data)
        elif "preferidas" in item_name:
            content = self.generate_modalidades_preferidas_content(item_data)
        elif "Onde e como" in item_name:
            content = self.generate_onde_como_faz_content(item_data)
        elif "sedentarismo" in item_name:
            content = self.generate_sedentarismo_content(item_data)
        elif "Provas/desafios" in item_name:
            content = self.generate_provas_desafios_content(item_data)
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
        print("PROCESSAMENTO DE SCORE ITEMS - MOVIMENTO E ATIVIDADE FÍSICA (BATCH 2)")
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
        print("RESUMO DO PROCESSAMENTO - BATCH 2")
        print("=" * 80)
        print(f"Total processados: {self.processed}/{len(ITEMS)}")
        print(f"Sucessos: {self.processed}")
        print(f"Erros: {len(self.errors)}")
        print(f"Tempo total: {elapsed:.1f}s")

        if self.errors:
            print("\nErros encontrados:")
            for error in self.errors:
                print(f"  - {error}")

        print("\n✓ Processamento Batch 2 concluído!")

if __name__ == "__main__":
    processor = MovimentoBatch2Processor()
    processor.run()
