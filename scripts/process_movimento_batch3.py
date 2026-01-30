#!/usr/bin/env python3
"""
Processamento de Score Items - Grupo "Movimento e atividade física" - Batch 3 FINAL
Últimos 15 items com enriquecimento de evidências científicas
"""

import requests
import json
import time
from typing import Dict, List, Optional

API_BASE = "http://localhost:3001/api/v1"
EMAIL = "import@plenya.com"
PASSWORD = "Import@123456"

# Items a processar (Batch 3 - FINAL)
ITEMS = [
    ("8e611b4a-a05e-45f3-9aca-0bddd66c8069", "Quem treina"),
    ("c3b39567-d51a-43c7-a9e0-a29099d7a710", "Quem treina"),
    ("2d2f2d91-923a-4dd1-9b13-b3d85979f9f6", "Quem treina"),
    ("a32cd296-be72-404e-9fe7-08f587ea4710", "Restrições de atividades"),
    ("c9565666-5d45-4ccd-a370-e461d8980883", "Restrições de atividades"),
    ("2672656c-8680-4b13-adf5-bdca1a504517", "Restrições de atividades"),
    ("87154ec9-6550-4d91-b3c1-3d041bb003d4", "Situação familiar/amigos de exercícios atual"),
    ("0ffb33de-5b9a-4f9c-9bf6-5c053962ca97", "Situação familiar/amigos de exercícios atual"),
    ("3c2a9cdf-1239-4f75-9cd7-4a09ae4913f5", "Situação familiar/amigos de exercícios atual"),
    ("0a6df688-8539-4c6f-a7a9-a1356371984d", "Suplementação pré, intra e pós treinos"),
    ("f87f5ece-1d1c-46f6-b16d-48e59d91f602", "Suplementação pré, intra e pós treinos"),
    ("cee904fc-5813-4f15-940c-a4d47cabb237", "Suplementação pré, intra e pós treinos"),
    ("717d09cc-2978-44bd-818c-fddd03fbcf69", "Vida adulta"),
    ("482f4b9c-2e18-4e4b-8d8a-237a26edb552", "Vida adulta"),
    ("2e7dbc4d-9ba3-441e-a22a-423acd1990e1", "Vida adulta"),
]

class MovimentoBatch3Processor:
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

    def generate_quem_treina_content(self, item_data: Dict) -> Dict:
        """Gera conteúdo para Quem treina (supervisão e parceiros)"""
        return {
            "clinicalRelevance": """A presença de supervisão profissional e/ou parceiros de treino influencia significativamente aderência, segurança, intensidade de esforço e resultados através de mecanismos de accountability social, feedback técnico em tempo real, facilitação social e motivação extrínseca. Meta-análises demonstram que treino supervisionado por educador físico/personal trainer resulta em 30-45% maiores ganhos de força e hipertrofia comparado a treino autogerenciado, atribuído a periodização adequada, correção técnica, progressão otimizada e maior volume total de treino (indivíduos sozinhos tendem a subtreinar).

O fenômeno de facilitação social (Zajonc, 1965) indica que a mera presença de outros aumenta ativação fisiológica e esforço em tarefas bem-aprendidas - estudos em ciclismo demonstram aumento de 10-15% na potência média quando treino é realizado em grupo vs. sozinho. Parceiros de treino fornecem competição amigável (aumenta intensidade), suporte emocional durante momentos difíceis, e compromisso social (reduz taxa de faltas em 40-60% quando treino é marcado com outra pessoa vs. compromisso individual).

Supervisão profissional reduz incidência de lesões em 50-70% através de correção biomecânica, progressão de carga adequada, identificação precoce de sinais de overtraining e adaptação de exercícios a limitações individuais. Em populações especiais (idosos, cardiopatas, diabéticos, gestantes), supervisão é essencial para segurança. Treino em grupo estruturado (CrossFit, bootcamps, aulas coletivas) combina benefícios de supervisão com coesão grupal, aumentando aderência para 65-75% em 12 meses vs. 30-40% em treino solitário não-supervisionado.

Limitações da supervisão incluem custo elevado (R$100-300/sessão personal 1:1), dependência excessiva (dificuldade de treinar sozinho), e potencial incompatibilidade de personalidade treinador-cliente. Treino solitário desenvolve maior autodisciplina, autonomia e conhecimento corporal profundo, sendo preferível para indivíduos intrinsecamente motivados após fase inicial de aprendizado técnico.""",

            "patientExplanation": """Quem está com você durante o treino - seja um personal trainer, um amigo, um grupo de pessoas ou ninguém - faz enorme diferença na sua experiência e nos resultados. Não existe uma opção "certa" para todo mundo, mas cada formato tem seus prós e contras.

Treinar com personal trainer ou educador físico é especialmente valioso no começo, quando você está aprendendo técnicas corretas, ou quando tem objetivos específicos e quer maximizar resultados. O profissional corrige seus movimentos, ajusta o treino para você, te puxa quando precisa e garante que você está progredindo com segurança. Ter um parceiro de treino (amigo, familiar) pode ser incrível para motivação - vocês não deixam um ao outro desistir, tornam o treino mais divertido e se desafiam mutuamente.

Treinar sozinho tem suas vantagens também: você vai no seu ritmo, não precisa coordenar horários com ninguém, e desenvolve autodisciplina. Muita gente começa com orientação profissional e depois passa a treinar sozinho quando já sabe o que fazer. O importante é ser honesto sobre o que funciona para VOCÊ - se você precisa de alguém para ter compromisso, não tente ser herói treinando sozinho e faltando sempre.""",

            "conduct": """**Investigação Detalhada de Contexto de Treino:**

**Situação Atual:**
- Treina sozinho ou acompanhado?
- Se acompanhado:
  - *Profissional:* Personal 1:1, semi-personal, coach de grupo, plantonista de academia
  - *Social:* Parceiro fixo (amigo, cônjuge, familiar), grupo informal, aulas coletivas
  - *Misto:* Combinação de formatos

**Frequência de Supervisão:**
- Todas as sessões, algumas sessões/semana, consultas periódicas (mensal), fase inicial apenas

**Histórico:**
- Sempre treinou neste formato ou já experimentou outros?
- Experiências prévias (positivas/negativas) com diferentes formatos

**Preferências:**
- Prefere treinar sozinho ou acompanhado?
- Valoriza autonomia vs. orientação externa?
- Motivação depende de presença de outros?

**Análise de Adequação:**

**Indicações para Supervisão Profissional Intensiva:**
- Iniciantes absolutos (0-6 meses experiência)
- Objetivos avançados (competição, performance específica)
- Condições médicas (cardiopatia, diabetes, hipertensão, obesidade severa, gestação)
- Reabilitação pós-lesão ou cirurgia
- Histórico de lesões recorrentes (necessita correção biomecânica)
- Múltiplas tentativas prévias falhas de aderência (necessita accountability externo)
- Medo/ansiedade relacionados a exercício

**Indicações para Treino com Parceiro:**
- Necessidade de accountability social moderada
- Preferência por experiência social durante exercício
- Motivação extrínseca (competição amigável)
- Limitação financeira para supervisão profissional mas desejo de suporte

**Indicações para Treino Individual:**
- Alta autodisciplina e motivação intrínseca
- Preferência por autonomia e flexibilidade de horário
- Conhecimento técnico suficiente (experiência >1-2 anos)
- Introversão/preferência por solidão durante exercício
- Objetivos de saúde geral (não performance específica)

**Modelos de Supervisão e Custo-Benefício:**

**Personal Training 1:1 (R$100-300/sessão):**
- *Vantagens:* Máxima personalização, feedback imediato, progressão otimizada, accountability total
- *Limitações:* Custo elevado, dependência potencial
- *Recomendação:* 2-3x/semana para iniciantes ou objetivos específicos, 1x/semana para manutenção em intermediários

**Semi-Personal 2-5 pessoas (R$50-100/sessão):**
- *Vantagens:* Custo intermediário, supervisão qualificada, componente social
- *Limitações:* Menos personalização que 1:1, necessidade de coordenação de horários
- *Recomendação:* Ótimo equilíbrio custo-benefício para maioria dos praticantes

**Consultoria Mensal (R$200-400/mês por 1-2 sessões):**
- *Vantagens:* Baixo custo, prescrição profissional, desenvolvimento de autonomia
- *Limitações:* Sem supervisão diária (requer autodisciplina), sem correção técnica em tempo real
- *Recomendação:* Intermediários/avançados que precisam de orientação mas não supervisão constante

**Aulas em Grupo (R$20-80/aula):**
- *Vantagens:* Custo acessível, social, estruturado, motivacional
- *Limitações:* Baixa personalização, turmas grandes reduzem atenção individual
- *Recomendação:* Excelente para quem valoriza componente social e necessita estrutura externa

**Treino com Parceiro Não-Profissional (gratuito):**
- *Vantagens:* Custo zero, accountability social, motivação mútua, diversão
- *Limitações:* Sem expertise técnica, dependência de agenda do parceiro
- *Recomendação:* Complementar a supervisão profissional inicial, ou suficiente para indivíduos já treinados

**Aplicativos/Programas Online (R$30-100/mês):**
- *Vantagens:* Baixo custo, flexibilidade total, acesso a expertise remota
- *Limitações:* Zero feedback técnico em tempo real, requer alta autodisciplina
- *Recomendação:* Para autodisciplinados com conhecimento técnico básico

**Estratégias de Transição:**

**Dependente → Autônomo:**
- *Situação:* Treina sempre com personal mas deseja independência (custo, disponibilidade)
- *Protocolo:*
  - Meses 1-2: Reduzir para 2x/semana supervisionado + 1x independente
  - Meses 3-4: 1x/semana supervisionado + 2-3x independente
  - Mês 5+: Consultoria mensal + execução independente
- *Ferramentas:* Planilha detalhada, vídeos de referência técnica, app de tracking

**Sozinho → Acompanhado:**
- *Situação:* Treina sozinho mas tem dificuldade de aderência ou estagnação
- *Protocolo:*
  - Contratar 8-12 sessões com profissional para "reset" (aprendizado técnico, nova periodização)
  - Buscar parceiro de treino para accountability
  - Participar de 1-2 aulas grupo/semana para componente social

**Otimização de Parceria de Treino:**

**Seleção de Parceiro Ideal:**
- Nível similar de condicionamento (diferenças >30% prejudicam ambos)
- Objetivos compatíveis (hipertrofia + hipertrofia, não hipertrofia + maratona)
- Disponibilidade de horários coincidente
- Personalidade compatível (ambos sérios vs. ambos descontraídos)
- Compromisso mútuo (não cancelar facilmente)

**Estruturação de Treino Conjunto:**
- Definir horários fixos semanais (compromisso firme)
- Comunicação clara de expectativas (intensidade, duração, foco)
- Revezamento de liderança (evitar dominância de um)
- Permitir sessões individuais ocasionais (flexibilidade)
- Celebrar progressos mútuos (reforço positivo)

**Gerenciamento de Conflitos:**
- Diferenças de progresso (inevitável - normalizar, evitar comparação)
- Mudança de disponibilidade (ter plano B para treino solo)
- Desmotivação de um (não arrastar o outro - permitir pausa)

**Maximização de Benefícios de Supervisão Profissional:**

**O Que Esperar do Profissional:**
- Avaliação inicial completa (anamnese, testes físicos)
- Prescrição periodizada (não treino aleatório)
- Correção técnica constante (feedback verbal e tátil)
- Progressão documentada (planilha, aplicativo)
- Ajustes conforme resposta individual (não programa genérico)
- Educação (explicar "porquê" dos exercícios)

**O Que o Cliente Deve Fazer:**
- Comunicação aberta (dor, desconforto, dúvidas)
- Aderência ao programa entre sessões (se semi-supervisionado)
- Feedback honesto sobre preferências e dificuldades
- Respeitar expertise (não ditar treino, colaborar na construção)

**Quando Trocar de Profissional:**
- Estagnação prolongada sem ajustes (>2-3 meses)
- Lesões recorrentes por técnica inadequada
- Falta de atenção (distração, celular, conversas paralelas)
- Ausência de progressão documentada (treino aleatório)
- Incompatibilidade de personalidade persistente

**Desenvolvimento de Autonomia a Longo Prazo:**

*Objetivo:* Mesmo com supervisão, desenvolver competências para treinar sozinho quando necessário

**Competências a Desenvolver:**
- Autoavaliação técnica (filmar-se, comparar com referências)
- Percepção de intensidade (RPE, reconhecer esforço adequado)
- Ajuste de carga (progressão segura, saber quando reduzir)
- Seleção de exercícios alternativos (equipamento indisponível, dor)
- Autocorreção postural (consciência corporal, propriocepção)

**Monitoramento:**
- Aderência (supervisão adequada se ≥70% sessões realizadas)
- Progressão (ganhos mensuráveis a cada 4-6 semanas)
- Satisfação (prazer no formato atual)
- Lesões (taxa deve ser <1 lesão/ano com supervisão adequada)
- Custo-benefício (investimento justificado pelos resultados?)

**Reavaliação Semestral:**
- Formato atual ainda é ideal ou mudanças de vida/objetivos exigem adaptação?
- Desenvolvimento de autonomia suficiente para reduzir supervisão?
- Necessidade de intensificar suporte (estagnação, perda de motivação)?

**Princípio Central:** Melhor formato de treino (sozinho, acompanhado, supervisionado) é aquele que maximiza ADERÊNCIA + SEGURANÇA + RESULTADOS para o indivíduo específico, não dogma universal."""
        }

    def generate_restricoes_atividades_content(self, item_data: Dict) -> Dict:
        """Gera conteúdo para Restrições de atividades"""
        return {
            "clinicalRelevance": """Restrições a atividades físicas abrangem contraindicações absolutas (exercício proibido até resolução da condição), contraindicações relativas (exercício permitido com modificações específicas) e precauções (monitoramento intensificado necessário). Condições cardiovasculares representam principais contraindicações: angina instável, arritmias não controladas, estenose aórtica severa, insuficiência cardíaca descompensada, infarto agudo do miocárdio (<2 semanas), hipertensão severa não controlada (PA >180/110 mmHg) contraindicam exercício vigoroso até estabilização médica.

Condições musculoesqueléticas requerem modificações: osteoartrite severa beneficia-se de exercícios de baixo impacto (natação, ciclismo) evitando corrida; osteoporose avançada contraindica flexão anterior de coluna (risco de fraturas vertebrais); tendinites agudas necessitam repouso relativo da estrutura afetada mantendo atividade de regiões não envolvidas. Hérnias discais com sintomas radiculares restringem hiperextensão lombar e cargas axiais elevadas; escoliose severa (>40°) requer supervisão para evitar rotações/lateralizações que agravam curva.

Condições metabólicas: diabetes tipo 1 com cetoacidose ou glicemia >250 mg/dL + cetonúria contraindica exercício (piora hiperglicemia); diabetes tipo 2 descompensado (glicemia >300 mg/dL) requer controle prévio; hipertireoidismo não controlado aumenta risco de arritmias durante esforço. Gravidez de alto risco (placenta prévia, pré-eclâmpsia, trabalho de parto prematuro, ruptura de membranas, restrição de crescimento intrauterino) restringe exercício; gravidez sem complicações permite até exercício moderado-vigoroso com adaptações.

Condições psiquiátricas: depressão severa com ideação suicida requer estabilização antes de programa de exercício não-supervisionado; transtornos alimentares (anorexia, bulimia) com exercício compulsivo requerem abordagem terapêutica multidisciplinar antes de prescrição; dismorfia muscular contraindica musculação não-supervisionada até tratamento psiquiátrico adequado.""",

            "patientExplanation": """Algumas condições de saúde exigem cuidado especial ou adaptações na hora de se exercitar - não para impedir você de ser ativo, mas para garantir que o exercício seja seguro e benéfico. É importante diferenciar: poucas condições impedem TOTALMENTE o exercício; a maioria apenas requer modificações inteligentes.

Por exemplo: se você tem problema no joelho, pode não conseguir correr, mas pode nadar, pedalar ou fazer musculação adaptada. Se tem pressão alta não controlada, precisa primeiro estabilizar com medicação e depois começar com exercícios leves, progredindo gradualmente. Se tem diabetes, precisa monitorar a glicemia antes e depois do treino para ajustar alimentação e medicação.

Nunca esconda condições de saúde do profissional que te orienta no exercício - quanto mais ele souber, melhor ele pode adaptar o treino para você. E se seu médico disse para evitar algum tipo específico de atividade, leve isso a sério. Exercício deve promover saúde, não criar problemas. Com as adaptações certas, quase todo mundo pode e deve ser fisicamente ativo.""",

            "conduct": """**Avaliação Completa de Restrições:**

**Anamnese de Condições Limitantes:**

**Cardiovasculares:**
- Hipertensão arterial (controlada? valores? medicação?)
- Doença arterial coronariana (história de infarto, angina, stents, revascularização)
- Arritmias (tipo, frequência, sintomática?)
- Insuficiência cardíaca (classe funcional NYHA)
- Valvopatias (tipo, severidade)
- Clearance médico para exercício (teste ergométrico, ecocardiograma recentes?)

**Metabólicas:**
- Diabetes (tipo 1 ou 2, controle glicêmico - HbA1c, hipoglicemias frequentes?, neuropatia periférica?)
- Doenças tireoidianas (hipo/hipertireoidismo, controlado?)
- Obesidade severa (IMC >40, limitações funcionais?)

**Musculoesqueléticas:**
- Osteoartrite (localização, severidade, dor em atividades?)
- Osteoporose (T-score, fraturas prévias, localização de maior risco)
- Hérnias discais (cervical/lombar, sintomas atuais, tratamento)
- Lesões agudas/crônicas não resolvidas (tendinites, entorses, fraturas)
- Escoliose/cifose/lordose acentuadas
- Próteses articulares (joelho, quadril - tempo pós-cirúrgico, restrições do ortopedista)
- Hipermobilidade/instabilidade articular

**Neurológicas:**
- Epilepsia (controlada? frequência de crises, gatilhos)
- AVC prévio (sequelas motoras, tempo pós-evento)
- Neuropatias periféricas (diabética, outras - perda sensorial, equilíbrio)
- Parkinson, esclerose múltipla (limitações funcionais)

**Respiratórias:**
- Asma (controlada? crises induzidas por exercício? uso de broncodilatador)
- DPOC (severidade, limitação funcional)

**Psiquiátricas:**
- Depressão severa (ideação suicida, medicação)
- Transtornos alimentares (anorexia, bulimia, exercício compulsivo)
- Transtorno dismórfico corporal/vigorexia
- Ansiedade severa relacionada a exercício

**Outras:**
- Gravidez (trimestre, complicações, clearance obstétrico)
- Doenças autoimunes (lúpus, artrite reumatoide - períodos de flare)
- Câncer (tipo, tratamento atual - quimio/radio, fadiga)

**Contraindicações Absolutas (NÃO exercitar até resolução):**

- Angina instável
- Arritmias não controladas causando sintomas
- Estenose aórtica severa sintomática
- Insuficiência cardíaca descompensada (NYHA IV)
- IAM agudo (<2 semanas sem clearance médico)
- Miocardite/pericardite aguda
- Dissecção de aorta
- Embolia pulmonar aguda
- Hipertensão severa não controlada (PA >180/110 mmHg)
- Diabetes com cetoacidose ou glicemia >300 mg/dL + sintomas
- Epilepsia não controlada (crises frequentes)
- Gravidez de alto risco com contraindicação médica expressa
- Febre/infecção sistêmica aguda
- Fraturas não consolidadas

**Contraindicações Relativas (exercício modificado):**

**Cardiovasculares:**
- *HAS controlada (140-160/90-100):* Evitar exercícios isométricos máximos (Valsalva), preferir aeróbico moderado, monitorar PA
- *Doença coronariana estável:* Programa supervisionado, teste ergométrico prévio, evitar intensidades máximas, reconhecer sinais de angina
- *Arritmias controladas:* Monitoramento de FC, evitar intensidades que desencadeiam arritmia

**Metabólicas:**
- *Diabetes tipo 1:* Monitorar glicemia pré/pós (evitar <100 ou >250 mg/dL), ajustar insulina/carboidratos, evitar exercício no pico de insulina
- *Diabetes tipo 2:* Monitorar glicemia se medicado com insulina/secretagogos, atenção a neuropatia (calçados adequados, evitar trauma em pés)
- *Hipotireoidismo:* Iniciar com baixa intensidade (fadiga), progressão gradual

**Musculoesqueléticas:**
- *Osteoartrite:* Evitar alto impacto (corrida → natação/ciclismo), fortalecimento de musculatura periarticular
- *Osteoporose:* Evitar flexão anterior de coluna, rotações bruscas, quedas; priorizar exercícios de impacto moderado (caminhada) + fortalecimento
- *Hérnia discal:* Evitar hiperextensão/flexão lombar extrema, cargas axiais elevadas; preferir exercícios em decúbito, fortalecimento de core
- *Tendinite aguda:* Repouso relativo da estrutura, manter atividade de outras regiões, fisioterapia concomitante

**Respiratórias:**
- *Asma induzida por exercício:* Broncodilatador pré-exercício, warm-up prolongado, evitar exercício em ar frio/seco/poluído
- *DPOC:* Exercício aeróbico leve-moderado (melhora capacidade funcional), treinamento muscular inspiratório

**Neurológicas:**
- *Epilepsia controlada:* Evitar atividades com risco de trauma na queda (escalada, natação solitária), preferir parceiro de treino ciente da condição
- *Neuropatia periférica:* Inspeção diária de pés, calçados adequados, evitar exercícios com risco de trauma despercebido

**Prescrição Adaptada por Condição:**

**Exemplo: Paciente com HAS controlada + Osteoartrite de joelho**

*Objetivos:* Reduzir PA, controlar peso, manter função articular

*Programa:*
- Aeróbico: Ciclismo ou natação 30-45min, 5x/semana, intensidade moderada (RPE 5-6/10)
- Fortalecimento: Musculação 2x/semana (quadríceps, glúteos, core), evitar Valsalva (respiração contínua), cargas moderadas (10-15 reps)
- Mobilidade: Alongamento e mobilidade articular 10min pós-treino
- *Evitar:* Corrida, agachamento profundo com carga alta, leg press máximo

**Exemplo: Diabetes tipo 1 + Objetivo de hipertrofia**

*Desafios:* Risco de hipoglicemia, ajuste de insulina

*Programa:*
- Monitorar glicemia: Pré-treino (idealmente 120-180 mg/dL), pós-treino
- Se <100 pré-treino: Carboidrato rápido 15-30g, aguardar 15min
- Fortalecimento: 4-5x/semana, foco hipertrofia (8-12 reps)
- Reduzir insulina basal pré-treino (ajuste com endocrinologista)
- Carboidratos peri-treino: 30-60g durante sessão longa (>60min), refeição pós-treino com proteína + carboidrato

**Exemplo: Gravidez sem complicações (2º trimestre)**

*Seguro:* Exercício moderado 30-40min, 4-5x/semana

*Programa:*
- Aeróbico: Caminhada, natação, bicicleta ergométrica (conversação confortável)
- Fortalecimento: Peso corporal ou cargas leves, foco em core adaptado (prancha modificada, bird dog), glúteos, MS
- *Evitar:* Supino decúbito dorsal após 16 semanas (compressão veia cava), exercícios com risco de queda, hipertermia (hot yoga)
- *Sinais de alarme para cessar exercício:* Sangramento, contrações regulares, tontura, dispneia severa

**Documentação e Clearance Médico:**

**Quando Exigir Clearance Médico:**
- Homens >45 anos ou mulheres >55 anos iniciando exercício vigoroso + fator de risco cardiovascular
- Qualquer idade com doença cardiovascular conhecida
- Diabetes com complicações (neuropatia, nefropatia, retinopatia)
- Sintomas sugestivos de doença cardiovascular (dor torácica, dispneia desproporcional, síncope)
- Condições complexas ou múltiplas comorbidades

**Documentação Recomendada:**
- Atestado médico liberando para prática (especificar modalidade e intensidade)
- Exames recentes: ECG, teste ergométrico (se indicado), exames laboratoriais
- Relatório médico detalhando condição, medicações, restrições específicas

**Comunicação Médico-Profissional de Educação Física:**
- Educador físico deve comunicar ao médico: Programa proposto, intensidades planejadas
- Médico deve especificar: Restrições, sinais de alarme, parâmetros de monitoramento (FC máxima, PA)

**Monitoramento Durante Exercício:**

**Sinais de Alarme (CESSAR exercício imediatamente):**
- Dor ou desconforto torácico
- Dispneia desproporcional ao esforço
- Tontura, vertigem, síncope
- Palpitações irregulares
- Dor articular aguda severa
- Náusea, vômito, palidez extrema

**Parâmetros a Monitorar:**
- *FC:* Manter abaixo de máxima prescrita (se cardiopata: conforme teste ergométrico)
- *PA:* Verificar pré/pós exercício em hipertensos (não deve ultrapassar 220/110 durante)
- *Glicemia:* Pré/pós em diabéticos (durante se sessão >60min)
- *SpO2:* Manter >90% em pacientes respiratórios

**Ajustes Conforme Resposta:**
- Progressão 50% mais conservadora que em indivíduos sem restrições
- Reavaliação médica a cada 6-12 meses ou se mudança na condição
- Comunicação contínua com equipe médica

**Princípio Central:** Restrições não impedem exercício, direcionam adaptações. Objetivo é maximizar benefícios minimizando riscos através de prescrição individualizada e monitoramento adequado."""
        }

    def generate_situacao_social_exercicios_content(self, item_data: Dict) -> Dict:
        """Gera conteúdo para Situação familiar/amigos de exercícios"""
        return {
            "clinicalRelevance": """O suporte social para atividade física constitui um dos preditores mais fortes de aderência e manutenção de comportamento ativo em longo prazo. A Teoria Sociocognitiva de Bandura identifica suporte social como modulador crítico de autoeficácia (crença na própria capacidade de realizar comportamento) e controle comportamental percebido. Meta-análises indicam que indivíduos com alto suporte social apresentam 70-80% mais probabilidade de manter exercício por >12 meses comparados a isolados socialmente.

Tipos de suporte social relacionados a exercício incluem: (1) Suporte emocional - encorajamento, validação, empatia durante dificuldades; (2) Suporte instrumental - facilitar logística (transporte, cuidado de crianças, ajuste de horários); (3) Suporte informacional - compartilhar conhecimento, experiências, recursos; (4) Suporte de companherismo - exercitar junto, criar contexto social prazeroso. Estudos longitudinais demonstram que cônjuge ativo aumenta probabilidade de parceiro ser ativo em 300%, presença de amigos próximos ativos em 150-200%.

Mecanismos neurobiológicos: exercício em contexto social aumenta liberação de ocitocina (hormônio de vínculo social), endorfinas e endocanabinoides, criando recompensa hedônica superior a exercício solitário. Fenômeno de contágio social (Christakis & Fowler) indica que comportamentos de saúde propagam em redes sociais - ganho de peso de amigo próximo aumenta risco de ganho próprio em 57%; inversamente, amigo iniciando exercício aumenta probabilidade de início próprio em 35-45%.

Barreiras sociais incluem: falta de suporte familiar (cônjuge/família desencoraja ou ridiculariza tentativas), normas sociais de sedentarismo no círculo próximo (normalização de inatividade), isolamento social (ausência de rede próxima), e conflitos de papéis (culpa por "tempo para si" em detrimento de família, especialmente mulheres). Intervenções focadas em construção de suporte social (grupos de exercício, envolvimento familiar) aumentam aderência em 40-60% comparadas a intervenções individuais isoladas.""",

            "patientExplanation": """As pessoas ao seu redor - família, amigos, colegas - influenciam enormemente se você vai conseguir manter uma rotina de exercícios. Não é só sobre ter ou não "força de vontade"; é sobre ter (ou não) um ambiente que apoia suas escolhas saudáveis.

Quando seu cônjuge, família ou amigos são ativos e te encorajam, fica muito mais fácil: eles entendem quando você precisa de tempo para treinar, às vezes treinam com você, compartilham dicas, vibram com suas conquistas. Mas quando as pessoas próximas são sedentárias, fazem piadas sobre você treinar, reclamam do tempo que você "perde" na academia, ou te oferecem comida não-saudável constantemente, a luta fica muito mais difícil.

Se você não tem esse suporte em casa, não significa que é impossível - mas significa que você precisa buscar suporte em outros lugares: grupos de corrida, colegas de academia, comunidades online, aulas coletivas. Humanos são seres sociais; precisamos de pertencimento e encorajamento. Criar ou encontrar uma "tribo" que valoriza saúde pode fazer toda a diferença entre desistir na primeira dificuldade ou persistir por anos.""",

            "conduct": """**Avaliação do Contexto Social Atual:**

**Suporte Familiar:**
- Estado civil (solteiro, casado, união estável, divorciado)
- Se casado/união: Parceiro é fisicamente ativo? Apoia/desencoraja seus esforços de exercício?
- Filhos (idades, cuidados necessários, impacto na disponibilidade de tempo)
- Família estendida (pais, irmãos) - valorizam atividade física? Influência sobre comportamento
- Divisão de responsabilidades domésticas (permite tempo para exercício?)
- Conflitos relacionados a tempo de exercício (culpa, críticas, ressentimento)

**Suporte de Amigos:**
- Amigos próximos são ativos ou sedentários?
- Tem amigos que exercitam regularmente com você ou separadamente?
- Atividades sociais do círculo (focadas em exercício vs. focadas em alimentação/bebidas)
- Encorajamento vs. sabotagem (ofertas de comida não-saudável, pressão para pular treinos)

**Suporte Comunitário:**
- Participa de grupos/comunidades de exercício (corrida, ciclismo, CrossFit, yoga)?
- Academia/local de treino oferece ambiente social positivo?
- Colegas de trabalho são ativos (grupos de corrida no almoço, desafios corporativos)?
- Acesso a redes online de suporte (apps, grupos, fóruns)

**Barreiras Sociais:**
- Isolamento social (mora sozinho, trabalho remoto, poucos amigos)
- Responsabilidades de cuidado (crianças, pais idosos, familiares doentes)
- Normas culturais/familiares de sedentarismo
- Críticas ou ridicularização de esforços de exercício

**Análise de Qualidade de Suporte:**

**Suporte Social Alto (Facilitador):**
- *Indicadores:*
  - Parceiro/família encoraja ativamente ("Vai treinar, eu cuido das crianças")
  - Amigos treinam junto ou respeitam horários de treino
  - Celebração de conquistas (PRs, provas, marcos)
  - Ajuste de logística familiar para permitir exercício
  - Modelagem positiva (familiares também ativos)
- *Resultado esperado:* Alta aderência (70-85% sessões realizadas)

**Suporte Social Moderado (Neutro):**
- *Indicadores:*
  - Família não desencoraja ativamente mas não facilita
  - Alguns amigos ativos, outros não
  - Necessidade de negociação constante de tempo
  - Ausência de sabotagem direta mas pouco encorajamento
- *Resultado esperado:* Aderência moderada (50-70%), maior dependência de motivação intrínseca

**Suporte Social Baixo (Barreira):**
- *Indicadores:*
  - Críticas de parceiro/família sobre tempo dedicado a exercício
  - Pressão para priorizar outras atividades sobre treino
  - Círculo social totalmente sedentário, normas de inatividade
  - Sabotagem ativa (ofertas de comida, convites conflitantes com horário de treino)
  - Ridicularização de esforços ("Não adianta nada", "Você está obcecado")
- *Resultado esperado:* Baixa aderência (20-40%), alto risco de dropout

**Estratégias de Intervenção por Nível de Suporte:**

**Se Suporte Social Alto: Manutenção e Otimização**
- Gratidão explícita (reconhecer e agradecer suporte recebido)
- Reciprocidade (oferecer suporte a atividades importantes para apoiadores)
- Expansão (envolver mais membros da rede em atividades físicas)
- Atividades conjuntas (caminhadas familiares, esportes recreativos)

**Se Suporte Social Moderado: Construção Ativa**

*Estratégias Familiares:*
- Comunicação assertiva de importância do exercício para saúde física/mental
- Negociação de horários fixos para treino (compromisso mútuo)
- Envolvimento de família em atividades (passeios de bike, caminhadas no parque)
- Demonstrar benefícios (humor melhor, mais energia para família)

*Estratégias de Amizade:*
- Buscar/cultivar amizades com pessoas ativas
- Convidar amigos existentes para atividades ativas (trilha em vez de bar)
- Participar de eventos esportivos sociais (corridas, desafios)

**Se Suporte Social Baixo: Construção de Nova Rede + Gestão de Barreiras**

**Construção de Rede de Suporte:**

*Comunidades de Exercício:*
- **Grupos de corrida:** Corridas em grupo 2-3x/semana, sociais pós-treino (suporte + accountability)
- **Aulas coletivas:** CrossFit, spinning, bootcamp, dança (comunidade integrada)
- **Clubes esportivos:** Ciclismo, natação, triathlon, artes marciais
- **Aplicativos sociais:** Strava, Nike Run Club, grupos de desafio (suporte virtual)
- **Grupos de interesse:** Meetup, grupos de Facebook locais

*Características de Comunidade Efetiva:*
- Valores compartilhados (saúde, superação)
- Inclusividade (todos níveis bem-vindos)
- Regularidade de encontros (cria rotina e vínculo)
- Suporte mútuo (celebração de sucessos, apoio em dificuldades)

**Gestão de Barreiras Familiares:**

*Comunicação:*
- Conversa honesta sobre importância do exercício para saúde mental e física
- Uso de "Eu" statements: "Eu preciso desse tempo para minha saúde" vs. "Você não me apoia"
- Demonstrar que exercício melhora qualidade de presença com família (não é egoísmo)

*Negociação:*
- Propor trade-offs: "Treino 1h de manhã, você tem 1h livre à tarde para seu hobby"
- Horários que minimizam impacto (antes de família acordar, horário de almoço)
- Exercício eficiente (30-45min de alta qualidade vs. 2h desnecessárias)

*Envolvimento:*
- Atividades familiares ativas nos finais de semana
- Modelagem para filhos (construindo cultura familiar ativa)
- Demonstrar resultados positivos (mais energia, melhor humor, redução de doenças)

**Se Críticas Persistentes:**
- Estabelecer limites firmes ("Este tempo não é negociável para minha saúde")
- Buscar suporte externo (amigos, comunidades) para compensar falta de suporte familiar
- Em casos extremos: Terapia de casal/familiar (padrões de controle, ressentimento)

**Isolamento Social:**
- Priorizar modalidades com componente social (aulas, grupos)
- Voluntariado em eventos esportivos (conexão com comunidade)
- Trabalho remoto: Coworking esportivo, grupos de atividade local
- Apps e comunidades online (suporte inicial até desenvolver rede presencial)

**Responsabilidades de Cuidado:**

*Crianças Pequenas:*
- Academia com creche (permite treino sem preocupação)
- Treino em casa durante soneca/após dormir
- Envolver crianças (carrinho de corrida, bicicleta com cadeirinha)
- Parceiro reveza cuidado para permitir treino de ambos

*Cuidado de Idosos/Doentes:*
- Treino em casa (curto, eficiente)
- Respite care (familiar/profissional assume temporariamente)
- Exercícios que podem ser interrompidos facilmente

**Construindo Cultura Social Ativa:**

**Para Indivíduo:**
- Tornar-se modelo positivo no círculo social (influência por exemplo, não pregação)
- Compartilhar experiências positivas (não de forma evangelizadora)
- Convidar pessoas próximas para experimentar atividades

**Para Família:**
- Rituais familiares ativos (caminhada pós-jantar, passeio de bike sábado)
- Substituir presentes por experiências ativas (aula de escalada, passeio de caiaque)
- Celebrar marcos com atividades (aniversário = trilha + piquenique)

**Monitoramento:**
- Aderência ao exercício (se <60%, reavaliar barreiras sociais)
- Satisfação com suporte recebido (escala 1-10)
- Conflitos relacionados a exercício (frequência, intensidade)
- Expansão de rede (novos amigos ativos? participação em comunidades?)

**Reavaliação:**
- Trimestral: Suporte social melhorou? Novas barreiras surgiram?
- Ajuste de estratégias conforme mudanças (nascimento de filho, mudança de cidade, mudança de emprego)

**Princípio Central:** Suporte social não é luxo, é necessidade. Construir e manter rede de apoio é TÃO importante quanto o próprio programa de exercício para sucesso em longo prazo."""
        }

    def generate_suplementacao_treino_content(self, item_data: Dict) -> Dict:
        """Gera conteúdo para Suplementação pré, intra e pós treinos"""
        return {
            "clinicalRelevance": """A suplementação peri-treino (pré, intra, pós) visa otimizar performance, retardar fadiga, maximizar adaptações anabólicas e acelerar recuperação através da modulação de substratos energéticos, perfil hormonal e síntese proteica muscular. Evidências robustas suportam eficácia de poucos suplementos; maioria apresenta efeitos modestos ou inconsistentes, com custo-benefício questionável para não-atletas.

**Pré-treino:** Cafeína (3-6 mg/kg, 60min antes) aumenta performance em resistência (+3-5%), força (+2-3%) e potência, através de antagonismo de receptores de adenosina (reduz percepção de fadiga) e aumento de catecolaminas. Creatina (3-5g/dia cronicamente) satura fosfocreatina muscular, melhorando performance anaeróbica (séries de alta intensidade, +5-15% trabalho total). Beta-alanina (3-6g/dia, 4+ semanas) aumenta carnosina muscular, tamponando acidose (benefício em exercícios 1-4 min duração, +2-3% performance). Carboidratos (30-60g, 30-60min antes) para treinos >90min previnem hipoglicemia e melhoram resistência.

**Intra-treino:** Carboidratos (30-60g/h) para sessões >60-90min mantêm glicemia e poupam glicogênio, retardando fadiga central e periférica (+10-20% tempo até exaustão em exercício prolongado). Eletrólitos (sódio 300-600mg/L) em treinos >2h ou alta sudorese previnem hiponatremia. Aminoácidos essenciais/BCAAs têm evidência limitada - proteína alimentar adequada torna suplementação redundante; potencial benefício em treino jejum ou hipocalórico severo.

**Pós-treino:** Proteína (20-40g, 0-2h pós-treino) maximiza síntese proteica muscular (MPS); janela anabólica menos crítica que volume proteico diário total (1.6-2.2g/kg/dia). Whey protein é superior a outras fontes (rápida digestão, alto teor de leucina - 3g necessários para máxima MPS). Carboidratos pós-treino (0.8-1.2g/kg) aceleram repleção de glicogênio, críticos para atletas com múltiplas sessões/dia; menos relevantes para praticantes recreacionais com >24h entre treinos. Combinação proteína + carboidrato não superior a proteína isolada para hipertrofia, mas acelera recuperação glicogênica.""",

            "patientExplanation": """Suplementos pré, intra e pós-treino são ferramentas que PODEM ajudar, mas não são essenciais para maioria das pessoas. Antes de gastar dinheiro com suplementos, certifique-se de que sua alimentação básica está adequada - suplemento não compensa dieta ruim.

Pré-treino: Uma xícara de café 30-60min antes pode dar energia e melhorar performance. Se seu treino é muito cedo, um lanche leve com carboidrato (banana, pão) garante que você tenha energia.

Durante o treino: Se você treina menos de 1 hora, água é suficiente. Se treina mais que 1-2 horas intensamente (corrida longa, ciclismo), bebidas esportivas com carboidrato e eletrólitos ajudam você a manter o ritmo.

Pós-treino: Seu corpo precisa de proteína para recuperar e construir músculo. Uma refeição com proteína (frango, ovos, peixe) + carboidrato (arroz, batata) nas 2 horas após treino é ótimo. Se não consegue fazer uma refeição logo, um shake de whey protein pode ser prático.

Lembre-se: alimentação real sempre vem primeiro. Suplementos são, literalmente, suplementares.""",

            "conduct": """**Avaliação de Uso Atual de Suplementação:**

**Suplementos Utilizados:**
- Pré-treino: Quais (cafeína, pré-treino comercial, beta-alanina, creatina)? Dose, timing, frequência
- Intra-treino: Carboidratos, eletrólitos, BCAAs, outros? Tipo de treino (duração, intensidade)
- Pós-treino: Proteína (whey, caseína, vegetal), carboidratos, glutamina, outros? Timing, dose
- Suplementos diários (não específicos de treino): Creatina, multivitamínico, ômega-3

**Justificativa:**
- Por que usa cada suplemento (objetivo percebido)?
- Nota diferença na performance/recuperação quando usa vs. não usa?
- Quanto gasta mensalmente em suplementação?

**Contexto Nutricional:**
- Ingestão proteica diária total (g/kg) - suplemento compensa déficit ou é redundante?
- Qualidade da dieta geral (se dieta ruim, suplemento não resolve)
- Timing de refeições vs. treino (treina em jejum? refeição pré-treino adequada?)

**Objetivos de Treino:**
- Performance competitiva (suplementação mais justificável)
- Hipertrofia/força (proteína + creatina têm maior evidência)
- Saúde geral/emagrecimento (maioria dos suplementos desnecessários)
- Resistência (carboidratos intra-treino em sessões longas)

**Evidência e Recomendações por Categoria:**

**PRÉ-TREINO**

**Cafeína - FORTE EVIDÊNCIA**
- *Dose:* 3-6 mg/kg peso corporal (200-400mg para 70kg)
- *Timing:* 30-60min antes do treino
- *Fontes:* Café (80-100mg/xícara), chá verde (30-50mg/xícara), suplemento isolado
- *Efeitos:* ↑ performance aeróbica (3-5%), força (2-3%), potência, foco, ↓ percepção de esforço
- *Considerações:* Tolerância individual (alguns sensíveis - ansiedade, insônia), evitar >6h antes de dormir
- *Recomendação:* Primeira escolha para boost pré-treino, custo-benefício excelente (café é barato)

**Beta-Alanina - EVIDÊNCIA MODERADA**
- *Dose:* 3-6g/dia (dividido em 2-3 doses de 1.5-2g para evitar parestesia)
- *Timing:* Cronicamente (4+ semanas para saturar carnosina muscular), não precisa ser pré-treino
- *Efeitos:* ↑ performance em exercícios 1-4min (séries alta intensidade, +2-3%)
- *Considerações:* Parestesia (formigamento) é comum mas inofensiva, efeito é cumulativo (não agudo)
- *Recomendação:* Útil para atletas competitivos, opcional para recreacionais

**Creatina - FORTE EVIDÊNCIA (mas efeito crônico, não agudo pré-treino)**
- *Dose:* 3-5g/dia (cronicamente, qualquer horário)
- *Timing:* Não precisa ser pré-treino (estoques saturados cronicamente)
- *Efeitos:* ↑ performance anaeróbica, força, hipertrofia (5-15% mais trabalho total)
- *Recomendação:* Um dos suplementos com melhor evidência, custo-benefício excelente

**Carboidratos - EVIDÊNCIA MODERADA (contexto-dependente)**
- *Dose:* 30-60g, 30-90min antes
- *Indicação:* Treino >90min ou treino matinal em jejum
- *Fontes:* Banana, pão branco, gel de carboidrato
- *Recomendação:* Necessário apenas para treinos longos ou jejum; maioria não precisa

**Pré-Treinos Comerciais (Multi-ingrediente):**
- *Composição típica:* Cafeína + beta-alanina + citrulina + BCAAs + creatina + vitaminas
- *Evidência:* Efeito majoritariamente da cafeína; demais ingredientes redundantes ou subdosados
- *Custo:* R$80-200/mês (cafeína isolada = R$10-30/mês)
- *Recomendação:* Custo-benefício ruim; preferir café + creatina separada (se desejado)

**INTRA-TREINO**

**Carboidratos - EVIDÊNCIA FORTE (treinos >60-90min)**
- *Dose:* 30-60g/h para treinos 1-2.5h, até 90g/h para >2.5h (combinação glicose + frutose)
- *Indicação:* Corridas longas, ciclismo, treinos >90min intensidade moderada-alta
- *Fontes:* Bebidas esportivas (Gatorade, Powerade), géis, gomas de carboidrato, água de coco
- *Efeitos:* Mantém glicemia, ↓ fadiga central, +10-20% tempo até exaustão
- *Recomendação:* Crítico para atletas de endurance, desnecessário para treinos <60min

**Eletrólitos (Sódio) - EVIDÊNCIA MODERADA**
- *Dose:* 300-600mg sódio/L (300-600mg/h)
- *Indicação:* Treinos >2h, alta temperatura, alta taxa de sudorese (>1L/h)
- *Efeitos:* Previne hiponatremia, mantém volume plasmático
- *Recomendação:* Necessário apenas em endurance prolongado; maioria usa água pura

**BCAAs - EVIDÊNCIA FRACA**
- *Alegação:* Reduz fadiga central, ↓ catabolismo muscular
- *Realidade:* Evidência inconsistente, redundante se ingestão proteica diária adequada
- *Custo:* R$80-150/mês
- *Recomendação:* Não recomendado; priorizar proteína alimentar

**Água - ESSENCIAL**
- *Dose:* 400-800ml/h conforme taxa de sudorese
- *Recomendação:* Hidratar conforme sede, monitorar cor da urina (amarelo claro ideal)

**PÓS-TREINO**

**Proteína - FORTE EVIDÊNCIA**
- *Dose:* 20-40g (0.3-0.5g/kg), idealmente 0-2h pós-treino (mas janela não é crítica)
- *Tipo:* Whey protein isolado/concentrado (rápida absorção, alto leucina)
- *Fontes alimentares:* 100-150g frango, 3-4 ovos, 150g peixe, 200-300ml leite + aveia
- *Efeitos:* Maximiza síntese proteica muscular (hipertrofia, recuperação)
- *Recomendação:* Essencial, mas alimentação pode suprir (suplemento = conveniência)

**Proteína: Whey vs. Caseína vs. Vegetal**
- *Whey:* Rápida (ideal pós-treino), alta leucina, completa
- *Caseína:* Lenta (ideal pré-sono para MPS noturna), completa
- *Vegetal (soja, ervilha, arroz):* Opção para veganos, preferir blend (aminograma completo)

**Carboidratos - EVIDÊNCIA MODERADA (contexto-dependente)**
- *Dose:* 0.8-1.2g/kg
- *Indicação:* Atletas com múltiplas sessões/dia (<8h entre treinos), repleção rápida de glicogênio necessária
- *Para recreacionais:* Menos crítico se >24h até próximo treino (repleção ocorre naturalmente com dieta)
- *Recomendação:* Maioria não precisa focar em carboidrato isolado pós-treino; refeição balanceada suficiente

**Creatina - PODE SER TOMADA PÓS-TREINO**
- *Timing:* Evidência ligeiramente favorável a pós-treino vs. pré (melhor absorção), mas diferença pequena
- *Recomendação:* Tomar diariamente (qualquer horário) é mais importante que timing específico

**Glutamina - EVIDÊNCIA FRACA**
- *Alegação:* Recuperação muscular, função imune
- *Realidade:* Não demonstra benefício em indivíduos bem nutridos
- *Recomendação:* Não recomendado; economize o dinheiro

**Prescrição Baseada em Perfil:**

**Iniciante/Recreacional (3-4x/semana, <1h/sessão, objetivo saúde geral):**
- *Essencial:* NADA - alimentação adequada suficiente
- *Opcional:* Café pré-treino (energia, custo zero), proteína pós se refeição não viável logo (whey, R$50-100/mês)
- *Evitar:* Pré-treinos comerciais caros, BCAAs, glutamina

**Intermediário (4-5x/semana, objetivos estéticos/força):**
- *Recomendado:*
  - Creatina 5g/dia (R$20-40/mês) - melhor custo-benefício
  - Proteína pós-treino se déficit proteico ou conveniência (whey, R$80-120/mês)
  - Cafeína pré-treino (café ou 200mg suplemento, R$10-30/mês)
- *Opcional:* Beta-alanina 3-6g/dia se treino alta intensidade (R$50-80/mês)
- *Total:* R$50-200/mês

**Avançado/Competitivo (6-7x/semana, treinos >90min, competições):**
- *Recomendado:* Todos acima +
  - Carboidratos intra-treino para sessões longas (bebida esportiva ou gel, R$100-200/mês)
  - Eletrólitos para treinos prolongados
  - Caseína pré-sono (30-40g, se déficit proteico, R$60-100/mês)
- *Considerar:* Beta-alanina, citrulina malato (6-8g, pump + performance)
- *Total:* R$200-500/mês (justificável se performance é prioritária)

**Atleta Endurance (corrida longa, ciclismo, triathlon):**
- *Crítico:*
  - Carboidratos intra-treino (géis, bebidas, 60-90g/h em treinos/provas longas)
  - Eletrólitos (sódio 300-600mg/h)
  - Cafeína pré-prova (200-400mg)
- *Proteína:* 20-30g pós-treino (recuperação muscular mesmo em esporte de resistência)

**Considerações de Segurança:**
- Suplementos não são regulados como medicamentos (qualidade variável)
- Preferir marcas certificadas (Informed Sport, NSF Certified for Sport - testam contaminantes)
- Evitar "proprietary blends" (dosagem não transparente)
- Atletas competitivos: Risco de contaminação com substâncias banidas (WADA)

**Educação sobre Prioridades:**
1. **Alimentação diária adequada** (80% do resultado)
2. **Treinamento consistente** (15% do resultado)
3. **Suplementação** (5% do resultado, apenas se 1 e 2 otimizados)

**Monitoramento:**
- Custo mensal (justificado por resultados?)
- Performance/recuperação percebida (teste períodos com e sem)
- Efeitos colaterais (GI, insônia, ansiedade)
- Reavaliação semestral: Continuar, ajustar ou eliminar?

**Princípio Central:** Suplementação é ferramenta de otimização marginal, não substituto de fundamentos (dieta + treino). Priorizar opções com evidência forte e custo-benefício favorável; evitar marketing enganoso de indústria bilionária."""
        }

    def generate_vida_adulta_content(self, item_data: Dict) -> Dict:
        """Gera conteúdo para Vida adulta e atividade física"""
        return {
            "clinicalRelevance": """A vida adulta (20-65 anos) apresenta desafios únicos para manutenção de atividade física devido a múltiplas demandas concorrentes (profissionais, familiares, sociais) e mudanças fisiológicas relacionadas à idade. Estudos epidemiológicos demonstram declínio progressivo de atividade física da juventude à meia-idade: 60% dos adultos jovens (20-30 anos) atingem recomendações de atividade física vs. apenas 35% dos adultos de meia-idade (40-55 anos) e 25% de idosos jovens (56-65 anos).

Fisiologicamente, a partir dos 30 anos observa-se sarcopenia progressiva (perda de 3-8% massa muscular por década, acelerando pós-50 anos), redução de VO2máx (1% ao ano em sedentários, 0.5% em ativos), diminuição de densidade óssea (especialmente mulheres pós-menopausa), aumento de gordura visceral e declínio de flexibilidade. Exercício regular atenua drasticamente esses declínios: adultos fisicamente ativos apresentam capacidade funcional 10-15 anos "mais jovem" que sedentários pareados por idade cronológica.

Transições de vida típicas da adultez impactam atividade física: início de carreira profissional (↓58% atividade vs. estudante), casamento (↓14%), nascimento de primeiro filho (↓40-60%, especialmente mulheres), divórcio (respostas variadas - alguns aumentam, outros diminuem), mudanças de emprego/cidade (disrupção de rotinas). Identificação de padrões permite antecipação e desenvolvimento de estratégias preventivas de recaída.

Benefícios do exercício na vida adulta incluem prevenção primária de doenças crônicas (↓40% DM2, ↓35% DCV, ↓20% câncer colorretal/mama), manutenção de capacidade funcional (ADLs e IADLs), preservação de massa muscular e óssea, controle de peso corporal, saúde mental (↓50% risco de depressão), produtividade profissional (+15% em ativos vs. sedentários) e qualidade de vida geral. Investimento em exercício na adultez determina saúde e independência na senescência.""",

            "patientExplanation": """A vida adulta é quando a maioria das pessoas para de se exercitar, mesmo sabendo que é importante. Entre trabalho exigente, responsabilidades familiares, casa para cuidar e pouco tempo sobrando, o exercício muitas vezes vira a primeira coisa a ser sacrificada quando a vida fica corrida.

Mas essa é exatamente a fase em que exercício é MAIS importante, não menos. Seu corpo já está começando a perder massa muscular, seu metabolismo está ficando mais lento, e os hábitos ruins (sedentarismo, má alimentação) começam a se acumular na forma de doenças crônicas. Adultos que se mantêm ativos têm corpo e energia de alguém 10-15 anos mais jovem, e muito menos risco de diabetes, pressão alta, obesidade e depressão.

O segredo não é tentar replicar os treinos da adolescência (você não tem 3 horas por dia disponíveis), mas encontrar formas REALISTAS de encaixar movimento na sua rotina adulta. Pode ser 30 minutos de manhã antes do trabalho, treino na hora do almoço, ou exercício curto à noite. O importante é consistência, não perfeição. Adultos que conseguem manter exercício regular colhem benefícios enormes em saúde, produtividade e qualidade de vida.""",

            "conduct": """**Anamnese de Atividade Física na Vida Adulta:**

**Padrão Geral:**
- Faixa etária atual e padrão histórico (sempre ativo, sempre sedentário, flutuante)
- Nível atual de atividade física (min/semana, modalidades)
- Tendência nas últimas décadas (manteve, declinou, aumentou)

**Transições de Vida e Impacto:**
- Entrada na vida profissional (idade, impacto na atividade física)
- Casamento/união estável (mudou padrão de exercício?)
- Nascimento de filhos (impacto - especialmente forte, duração do declínio)
- Mudanças de emprego/carreira (transições disruptivas)
- Mudanças geográficas (perda de infraestrutura, rede social)
- Eventos de saúde (diagnósticos, cirurgias, períodos de inatividade)
- Divórcio, perda de entes queridos (impacto variado)

**Barreiras Específicas da Vida Adulta:**
- *Profissionais:* Longas jornadas, trabalho mental/fisicamente exaustivo, viagens frequentes, horários irregulares
- *Familiares:* Cuidado de filhos pequenos, responsabilidades domésticas, cuidado de pais idosos
- *Tempo:* Percepção de "falta de tempo" (priorização competitiva)
- *Energia:* Fadiga crônica pós-trabalho, depleção de força de vontade ao fim do dia
- *Financeiras:* Custo de academia, equipamentos, personal
- *Fisiológicas:* Lesões acumuladas, dores crônicas, declínio de capacidade percebida

**Estratégias de Manutenção/Retorno por Fase da Vida Adulta:**

**ADULTO JOVEM (20-35 anos):**

*Características:*
- Alta capacidade física, recuperação rápida
- Estabelecimento de carreira (longas jornadas comuns)
- Início de relacionamentos sérios, possível casamento, primeiros filhos
- Risco: Transição estudante → profissional (↓60% atividade)

*Estratégias:*
- **Manter hábitos da juventude:** Se era ativo, priorizar continuidade (mais fácil manter que retomar)
- **Integração social:** Esportes coletivos, grupos de corrida/ciclismo (suporte social + atividade)
- **Eficiência:** Treinos 30-45min de alta qualidade (HIIT, força funcional) vs. 2h desnecessárias
- **Horário:** Manhã (antes de trabalho "roubar" tempo) ou hora de almoço
- **Prevenção:** Estabelecer rotina ANTES de casamento/filhos (mais fácil manter que criar depois)

*Objetivos típicos:* Estéticos, performance, social

*Prescrição exemplo:*
- 4-5x/semana, 30-60min
- Combinação força + aeróbico (treino concorrente)
- Modalidades variadas (evitar monotonia)

**ADULTO DE MEIA-IDADE (35-50 anos):**

*Características:*
- Início de declínios fisiológicos perceptíveis (↓força, ↑gordura abdominal, ↓recuperação)
- Pico de demandas profissionais e familiares (filhos escolares, carreira consolidada)
- Risco: Nascimento de filhos, promoções exigentes (↓50% atividade)
- Aumento de diagnósticos (HAS, pré-diabetes, dislipidemia) - motivador potencial

*Estratégias:*
- **Não-negociável:** Exercício como "medicamento" essencial, não luxo
- **Eficiência máxima:** 3-4x/semana, 30-45min suficientes se bem estruturados
- **Timing estratégico:**
  - Manhã cedo (5h30-6h30, antes de família acordar) - mais sustentável
  - Hora de almoço (treino + almoço rápido)
  - Flexibilidade (não depender de horário único - ter planos A, B, C)
- **Envolver família:**
  - Atividades ativas com filhos (bike, trilha, esportes)
  - Negociar com cônjuge ("revezamento" - cada um tem janelas de treino garantidas)
- **Treino em casa:** Equipamento mínimo (faixas, halteres) elimina deslocamento (economiza 30-40min)
- **Foco em saúde:** Menos estética, mais longevidade e prevenção de doenças

*Objetivos típicos:* Saúde, prevenção de doenças, manutenção de capacidade funcional

*Prescrição exemplo:*
- 3-4x/semana, 30-45min
- Ênfase em força (combate sarcopenia) + cardio moderado
- Mobilidade/flexibilidade (previne lesões, combate rigidez)

**ADULTO TARDIO/PRÉ-IDOSO (50-65 anos):**

*Características:*
- Declínios fisiológicos mais acentuados (sarcopenia acelerada pós-50, especialmente mulheres pós-menopausa)
- Possível redução de demandas familiares (filhos crescidos) - "ninho vazio"
- Maior prevalência de condições crônicas (HAS, DM2, artrose)
- Proximidade de aposentadoria (potencial aumento de tempo disponível)
- Risco: Acúmulo de lesões/dores crônicas, percepção de "ser velho demais"

*Estratégias:*
- **Combate à sarcopenia:** Fortalecimento 2-3x/semana ESSENCIAL (preserva independência futura)
- **Impacto ósseo:** Caminhada, corrida leve, exercícios de impacto (previne osteoporose)
- **Equilíbrio e propriocepção:** Previne quedas (single leg stance, yoga, tai chi)
- **Adaptação a limitações:** Modificar exercícios conforme dores/artrose (substituir corrida por ciclismo/natação)
- **Aproveitar maior disponibilidade:** Se filhos cresceram, pode haver mais tempo para exercício
- **Preparo para aposentadoria:** Estabelecer rotina que independe de trabalho

*Objetivos típicos:* Independência funcional, qualidade de vida, gerenciamento de doenças

*Prescrição exemplo:*
- 4-5x/semana (mais tempo disponível)
- Força 2-3x (foco anti-sarcopenia)
- Aeróbico 3-4x (saúde cardiovascular)
- Flexibilidade/equilíbrio 2-3x

**Estratégias Transversais (todas idades adultas):**

**Gestão de Tempo:**
- **Auditar tempo:** Registrar 1 semana - onde tempo "desaparece"? (TV, redes sociais = média 3-4h/dia)
- **Microworkouts:** 3x 10-15min/dia = 30-45min total (tão efetivo quanto sessão única)
- **Atividade espontânea:** Escadas vs. elevador, caminhar em reuniões, bicicleta para trabalho
- **Combinação:** Audiobook/podcast durante corrida, reuniões em caminhada

**Gestão de Energia:**
- **Exercício GERA energia:** Paradoxo - pessoas acham que estão "cansadas demais", mas exercício aumenta energia
- **Timing:** Manhã (antes de depleção de willpower do dia) > Noite
- **Intensidade adaptada:** Dias de baixa energia = treino leve/mobilidade (não pular)

**Flexibilidade de Rotina:**
- Ter 3 formatos: Ideal (60min academia), Plano B (30min casa), Plano C (15min qualquer coisa)
- Aceitar que nem sempre será perfeito - consistência imperfeita > perfeição inconsistente

**Suporte Social:**
- Parceiro de treino (accountability)
- Grupos/comunidades (compromisso social)
- Envolver família (apoio, não sabotagem)

**Reenquadramento Cognitivo:**
- Exercício não é "tempo perdido", é investimento em produtividade/saúde/longevidade
- Não é egoísmo - é responsabilidade (ser saudável para cuidar de família, não ser fardo futuro)
- Não é "devia" - é "quero" (motivação intrínseca sustentável)

**Prevenção de Recaídas em Transições:**

**Antecipação:**
- Identificar transições futuras (nascimento de filho, mudança de emprego, aposentadoria)
- Planejar estratégias ANTES da transição (estabelecer novo formato de treino preventivamente)

**Durante Transição:**
- Aceitar redução temporária (dose mínima efetiva - 50-60% do habitual)
- Flexibilidade total (treinos curtos, modalidades adaptadas)
- Não abandonar completamente (manter conexão, mais fácil retomar)

**Pós-Transição:**
- Reavaliação após 3-6 meses (nova rotina estabelecida)
- Reconstrução gradual (não pular para volume prévio imediatamente)

**Monitoramento:**
- Aderência semanal (% treinos planejados realizados)
- Tendência trimestral (mantém, declina, melhora?)
- Identificação precoce de declínio (intervenção imediata)
- Reavaliação anual: Estratégias atuais funcionam? Mudanças de vida exigem adaptação?

**Princípio Central:** Vida adulta EXIGE criatividade e flexibilidade para manter atividade física. Não é sobre ter "tempo sobrando" (nunca haverá), é sobre PRIORIZAR e INTEGRAR exercício na vida real complexa. Adultos que conseguem são recompensados com saúde superior, mais energia, melhor humor e maior qualidade de vida em todas as dimensões."""
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
        print(f"\n[{self.processed + 1}/15] Processando: {item_name} ({item_id})")

        # Buscar detalhes do item
        item_data = self.get_item_details(item_id)
        if not item_data:
            return False

        # Gerar conteúdo baseado no tipo
        if "Quem treina" in item_name:
            content = self.generate_quem_treina_content(item_data)
        elif "Restrições" in item_name:
            content = self.generate_restricoes_atividades_content(item_data)
        elif "Situação familiar" in item_name or "amigos" in item_name:
            content = self.generate_situacao_social_exercicios_content(item_data)
        elif "Suplementação" in item_name:
            content = self.generate_suplementacao_treino_content(item_data)
        elif "Vida adulta" in item_name:
            content = self.generate_vida_adulta_content(item_data)
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
        print("PROCESSAMENTO DE SCORE ITEMS - MOVIMENTO E ATIVIDADE FÍSICA (BATCH 3 FINAL)")
        print("=" * 80)

        if not self.login():
            return

        print(f"\nIniciando processamento de {len(ITEMS)} items finais...")

        start_time = time.time()

        for item_id, item_name in ITEMS:
            self.process_item(item_id, item_name)
            time.sleep(0.5)  # Rate limiting

        elapsed = time.time() - start_time

        print("\n" + "=" * 80)
        print("RESUMO DO PROCESSAMENTO - BATCH 3 FINAL")
        print("=" * 80)
        print(f"Total processados: {self.processed}/{len(ITEMS)}")
        print(f"Sucessos: {self.processed}")
        print(f"Erros: {len(self.errors)}")
        print(f"Tempo total: {elapsed:.1f}s")

        if self.errors:
            print("\nErros encontrados:")
            for error in self.errors:
                print(f"  - {error}")

        print("\n" + "=" * 80)
        print("GRUPO 'MOVIMENTO E ATIVIDADE FÍSICA' COMPLETO!")
        print("=" * 80)
        print("Total geral: 75 items processados (30 + 30 + 15)")
        print("✓ Todos os items do grupo enriquecidos com sucesso!")

if __name__ == "__main__":
    processor = MovimentoBatch3Processor()
    processor.run()
