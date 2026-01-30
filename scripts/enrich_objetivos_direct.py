#!/usr/bin/env python3
"""
Script para enriquecer Score Items do grupo "Objetivos"
Busca items do banco e atualiza via API
"""

import requests
import subprocess
import json

API_URL = "http://localhost:3001/api/v1"
EMAIL = "import@plenya.com"
PASSWORD = "Import@123456"

def login():
    """Fazer login e obter token JWT"""
    try:
        response = requests.post(
            f"{API_URL}/auth/login",
            json={"email": EMAIL, "password": PASSWORD}
        )
        response.raise_for_status()
        data = response.json()
        token = data.get("accessToken")
        if not token:
            print("✗ Token não encontrado na resposta")
            return None
        print("✓ Login realizado com sucesso")
        return token
    except Exception as e:
        print(f"✗ Erro no login: {e}")
        return None

def get_items_from_db():
    """Buscar items do banco de dados"""
    cmd = [
        "docker", "compose", "exec", "-T", "db",
        "psql", "-U", "plenya_user", "-d", "plenya_db",
        "-t", "-A", "-F", "|",
        "-c", """
        SELECT si.id, si.name, sg.name AS subgroup, g.name AS grupo
        FROM score_items si
        LEFT JOIN score_subgroups sg ON si.subgroup_id = sg.id
        LEFT JOIN score_groups g ON sg.group_id = g.id
        WHERE g.name = 'Objetivos'
        ORDER BY sg.name, si.name;
        """
    ]

    try:
        result = subprocess.run(cmd, capture_output=True, text=True, check=True)
        lines = [line.strip() for line in result.stdout.strip().split('\n') if line.strip()]

        items = []
        for line in lines:
            parts = line.split('|')
            if len(parts) >= 3:
                items.append({
                    "id": parts[0],
                    "name": parts[1],
                    "subgroup": parts[2],
                    "group": parts[3] if len(parts) > 3 else "Objetivos"
                })

        print(f"✓ Encontrados {len(items)} items no banco de dados")
        return items
    except Exception as e:
        print(f"✗ Erro ao buscar items do banco: {e}")
        return []

def get_content_for_future_perception(item_name):
    """Conteúdo para items de percepção de futuro"""
    if "5 anos" in item_name:
        timeframe = "cinco anos"
        clinical_focus = "objetivos de curto-médio prazo"
        intervention = "mudanças comportamentais imediatas"
    elif "10 anos" in item_name:
        timeframe = "dez anos"
        clinical_focus = "planejamento de médio prazo"
        intervention = "estratégias preventivas consolidadas"
    elif "20 anos" in item_name:
        timeframe = "vinte anos"
        clinical_focus = "visão de longo prazo"
        intervention = "prevenção de doenças crônicas"
    elif "30 anos" in item_name:
        timeframe = "trinta anos"
        clinical_focus = "longevidade e qualidade de vida"
        intervention = "otimização do envelhecimento saudável"
    else:
        timeframe = "futuro"
        clinical_focus = "planejamento de saúde"
        intervention = "estratégias personalizadas"

    clinical_relevance = f"""A percepção que o paciente tem sobre sua saúde e vida em um horizonte de {timeframe} é um indicador crucial para o planejamento terapêutico em medicina funcional integrativa. Estudos em medicina centrada no paciente demonstram que a definição clara de objetivos de longo prazo melhora significativamente a adesão terapêutica e os resultados clínicos (Graffigna et al., 2014; Robinson et al., 2008).

Na medicina funcional, o estabelecimento de metas temporais permite ao profissional identificar prioridades do paciente, alinhar expectativas e desenvolver planos de intervenção personalizados. O foco em {clinical_focus} facilita a implementação de {intervention} que respeitem o contexto de vida único de cada indivíduo.

A literatura científica mostra que pacientes com visão clara de futuro apresentam maior motivação para mudanças de estilo de vida, melhor engajamento com planos terapêuticos e resultados superiores em marcadores de saúde metabólica, cardiovascular e neurológica. A medicina de longevidade moderna enfatiza que o planejamento antecipado, especialmente quando iniciado precocemente, é fundamental para prevenir doenças crônicas e otimizar o envelhecimento saudável (Seals et al., 2016; Kennedy et al., 2014).

Em termos práticos, compreender a percepção de futuro do paciente permite ao clínico adaptar recomendações nutricionais, protocolos de suplementação, estratégias de exercício e intervenções farmacológicas ao ritmo e às capacidades individuais, aumentando a probabilidade de sustentabilidade das mudanças propostas."""

    patient_explanation = f"""Quando perguntamos sobre como você se imagina daqui a {timeframe}, estamos tentando entender seus sonhos, preocupações e expectativas em relação à sua saúde e qualidade de vida. Essa visão de futuro é muito importante para construirmos juntos um plano de cuidado que faça sentido para você.

Saber o que você valoriza e deseja alcançar nos ajuda a criar metas realistas e personalizadas. Por exemplo, se você deseja estar ativo e independente no futuro, podemos focar em estratégias que fortaleçam sua saúde óssea, muscular e cardiovascular desde agora. Se sua preocupação é com a prevenção de doenças que existem na família, podemos implementar medidas preventivas específicas.

Ter clareza sobre seus objetivos também torna mais fácil manter a motivação durante o tratamento, porque você saberá exatamente por que está fazendo cada mudança no seu estilo de vida."""

    conduct = f"""**Avaliação Inicial:**
1. Explorar com o paciente sua visão de saúde e qualidade de vida em {timeframe}
2. Identificar preocupações específicas (doenças familiares, limitações físicas, cognição)
3. Avaliar realismo das expectativas considerando fatores de risco atuais
4. Documentar objetivos pessoais e profissionais relacionados à saúde

**Estabelecimento de Metas:**
1. Traduzir a visão de longo prazo em objetivos mensuráveis
2. Definir marcos intermediários (check-ins semestrais/anuais)
3. Priorizar intervenções com maior impacto preventivo para o horizonte temporal escolhido
4. Alinhar metas com valores pessoais e contexto de vida do paciente

**Plano de Intervenção para {timeframe.capitalize()}:**
1. **Prevenção primária:** Identificar e modificar fatores de risco modificáveis
2. **Otimização metabólica:** Nutrição, suplementação e exercício personalizados
3. **Monitoramento:** Biomarcadores relevantes para longevidade (inflamação, glicação, oxidação)
4. **Ajustes regulares:** Revisar e adaptar o plano conforme evolução e novas evidências

**Acompanhamento:**
- Revisões trimestrais de progresso e adesão
- Reavaliação de metas anualmente
- Celebrar conquistas intermediárias
- Ajustar estratégias conforme mudanças na vida do paciente"""

    return {
        "clinical_relevance": clinical_relevance,
        "patient_explanation": patient_explanation,
        "conduct": conduct
    }

def get_content_for_adherence(item_name):
    """Conteúdo para items de adesão e disciplina"""
    if "Muito disciplinado" in item_name:
        profile = "alta autodisciplina"
        approach = "empoderamento e autonomia"
        challenge = "perfeccionismo ou rigidez excessiva"
    elif "Disciplina moderada" in item_name:
        profile = "disciplina moderada"
        approach = "suporte estruturado e motivação contínua"
        challenge = "manutenção do engajamento a longo prazo"
    elif "Pouco disciplinado" in item_name:
        profile = "baixa autodisciplina"
        approach = "acompanhamento intensivo e health coaching"
        challenge = "estabelecimento de rotinas sustentáveis"
    else:
        profile = "perfil comportamental"
        approach = "estratégias personalizadas"
        challenge = "identificação de barreiras individuais"

    clinical_relevance = f"""A avaliação do perfil comportamental e da capacidade de adesão a planos terapêuticos é um dos pilares fundamentais da medicina funcional integrativa centrada no paciente. Pacientes com {profile} requerem abordagens específicas que respeitem suas características individuais e maximizem a probabilidade de sucesso terapêutico.

A literatura científica demonstra consistentemente que a adesão terapêutica é o principal determinante de resultados clínicos em intervenções de mudança de estilo de vida. Estudos mostram que a abordagem deve ser adaptada ao perfil motivacional e às capacidades de autorregulação de cada indivíduo, sendo que intervenções "tamanho único" apresentam baixa efetividade (Nieuwlaat et al., 2014; Martin et al., 2005).

No contexto da medicina funcional, compreender a auto-percepção de responsabilidade e disciplina permite ao profissional: (1) dimensionar adequadamente a complexidade inicial do plano terapêutico; (2) determinar a frequência ideal de acompanhamento; (3) identificar necessidade de health coaching ou suporte multidisciplinar; e (4) antecipar obstáculos potenciais à implementação.

A estratégia de {approach} tem demonstrado melhores resultados para este perfil. Health coaching, em particular, tem mostrado eficácia significativa no suporte a mudanças comportamentais sustentáveis, especialmente quando combinado com técnicas de entrevista motivacional e estabelecimento colaborativo de metas (Wolever et al., 2013; Kivelä et al., 2014). É essencial evitar sobrecarga cognitiva ou subestimação das capacidades do paciente, personalizando tanto a complexidade quanto o ritmo das mudanças propostas."""

    patient_explanation = f"""Entender como você se percebe em relação à disciplina e adesão a planos é fundamental para criarmos um tratamento que funcione para VOCÊ. Não existe certo ou errado aqui - cada pessoa tem um ritmo e uma forma de lidar com mudanças.

Se você se identifica como alguém que precisa de mais suporte, isso não é um problema. Na verdade, é muito positivo você ter essa consciência, pois podemos planejar acompanhamentos mais frequentes, estratégias de lembretes e apoio extra quando necessário. Se você é mais independente, podemos dar mais autonomia e focar em ajustes finos.

O importante é ser honesto sobre seu perfil, porque isso nos permite ajustar o plano às suas necessidades reais, aumentando muito suas chances de sucesso e evitando frustrações desnecessárias."""

    conduct = f"""**Avaliação do Perfil Comportamental:**
1. Aplicar questionários validados de autorregulação e motivação
2. Explorar experiências prévias com mudanças de hábitos (sucessos e fracassos)
3. Identificar gatilhos, barreiras e facilitadores específicos
4. Avaliar sistemas de suporte disponíveis (família, amigos, profissionais)

**Estratégias Específicas para Este Perfil:**
1. **Complexidade do plano:** Ajustar número e dificuldade das mudanças simultâneas
2. **Frequência de contato:** Determinar periodicidade ideal de consultas e check-ins
3. **Ferramentas de suporte:** Apps, lembretes, diários, grupos de apoio
4. **Foco comportamental:** Priorizar mudanças com maior probabilidade de adesão inicial

**Abordagem Terapêutica - {approach.capitalize()}:**
1. Estabelecer metas incrementais e celebrar pequenas vitórias
2. Utilizar técnicas de entrevista motivacional
3. Implementar estratégias de health coaching quando indicado
4. Criar sistemas de accountability personalizados

**Desafio Principal - {challenge.capitalize()}:**
1. Identificar precocemente sinais de desengajamento
2. Ajustar plano proativamente antes de falhas completas
3. Trabalhar crenças limitantes e autoeficácia
4. Desenvolver resiliência frente a recaídas temporárias

**Monitoramento:**
- Acompanhar adesão objetiva (registros, biomarcadores)
- Avaliar percepção subjetiva de dificuldade e sobrecarga
- Revisar e simplificar protocolo se necessário
- Evoluir gradualmente complexidade conforme ganho de confiança"""

    return {
        "clinical_relevance": clinical_relevance,
        "patient_explanation": patient_explanation,
        "conduct": conduct
    }

def get_content_for_initial_goals(item_name):
    """Conteúdo para items de objetivos iniciais"""
    clinical_relevance = """A identificação explícita dos objetivos iniciais que motivam o paciente a buscar um programa de medicina funcional integrativa é um passo fundamental para o sucesso terapêutico. A literatura em medicina centrada no paciente demonstra que o alinhamento entre as expectativas do paciente e os objetivos do profissional é preditor robusto de satisfação, adesão e resultados clínicos (Robinson et al., 2008; Stewart et al., 2000).

Em medicina funcional, os pacientes frequentemente apresentam múltiplas queixas inter-relacionadas, sendo essencial priorizar quais problemas serão abordados primeiro. A definição clara de objetivos permite: (1) estabelecer expectativas realistas sobre tempo e complexidade do tratamento; (2) criar marcos mensuráveis de progresso; (3) manter motivação durante fases de platô; e (4) ajustar estratégias quando os objetivos iniciais evoluem.

Estudos mostram que pacientes que participam ativamente da definição de metas terapêuticas apresentam maior engajamento, melhor compreensão do plano de tratamento e maior probabilidade de manutenção de mudanças de estilo de vida a longo prazo (Bodenheimer & Handley, 2009; Holman & Lorig, 2000). A abordagem colaborativa, onde profissional e paciente constroem juntos os objetivos, é superior a modelos paternalistas onde metas são impostas unilateralmente.

Do ponto de vista clínico, os objetivos iniciais do paciente revelam suas prioridades de saúde, valores pessoais e qualidade de vida desejada. Alguns pacientes priorizam energia e disposição, outros focam em emagrecimento, prevenção de doenças familiares, otimização cognitiva ou longevidade. Cada objetivo requer estratégias específicas, e o sucesso depende de honrar as motivações intrínsecas do indivíduo."""

    patient_explanation = """Os objetivos que você tem ao iniciar nosso programa são a bússola que guiará todo o seu tratamento. Queremos entender exatamente o que é mais importante para você neste momento: o que te motivou a buscar ajuda? O que você espera melhorar ou alcançar?

Algumas pessoas querem mais energia no dia a dia, outras buscam perder peso de forma saudável, prevenir doenças que existem na família, melhorar o sono, reduzir dores, otimizar o desempenho mental ou simplesmente envelhecer com mais qualidade de vida. Não existe objetivo "pequeno" ou "grande" - existe o que é importante para VOCÊ.

Quando deixamos claro o que queremos alcançar, fica mais fácil medir o progresso, fazer ajustes no caminho e manter a motivação. Além disso, seus objetivos podem evoluir ao longo do tratamento, e isso é completamente normal e bem-vindo."""

    conduct = """**Exploração de Objetivos Iniciais:**
1. **Conversa aberta e empática:**
   - Perguntar: "O que te motivou a buscar este programa agora?"
   - "Qual é sua maior preocupação de saúde atualmente?"
   - "Como você gostaria de se sentir daqui a 3-6 meses?"

2. **Categorização de objetivos:**
   - Sintomáticos (energia, dor, sono, digestão)
   - Estéticos/composição corporal
   - Preventivos (evitar doenças familiares)
   - Performance (física, cognitiva, sexual)
   - Longevidade e qualidade de vida

3. **Priorização conjunta:**
   - Identificar objetivo primário e secundários
   - Explicar inter-relações entre diferentes queixas
   - Estabelecer ordem de abordagem baseada em impacto e viabilidade

**Alinhamento de Expectativas:**
1. **Realismo temporal:** Explicar prazos típicos para diferentes objetivos
2. **Complexidade:** Esclarecer que alguns objetivos requerem múltiplas frentes de ação
3. **Comprometimento necessário:** Dimensionar mudanças de estilo de vida requeridas
4. **Monitoramento:** Definir como será medido o progresso (sintomas, exames, métricas)

**Estabelecimento de Metas SMART:**
1. **Específicas:** Traduzir objetivos vagos em metas concretas
2. **Mensuráveis:** Definir indicadores objetivos de sucesso
3. **Alcançáveis:** Ajustar expectativas irrealistas com empatia
4. **Relevantes:** Garantir que metas reflitam valores do paciente
5. **Temporais:** Estabelecer marcos de curto, médio e longo prazo

**Documentação e Contrato Terapêutico:**
1. Registrar objetivos declarados pelo paciente (textualmente)
2. Criar plano terapêutico alinhado com prioridades identificadas
3. Estabelecer pontos de revisão (90 dias, 6 meses, 1 ano)
4. Documentar consentimento informado sobre abordagem escolhida

**Revisão Contínua:**
- Revisar objetivos em cada consulta de acompanhamento
- Celebrar progressos, mesmo que parciais
- Ajustar metas conforme evolução clínica
- Adicionar novos objetivos conforme primários são alcançados
- Manter flexibilidade para mudanças nas prioridades do paciente"""

    return {
        "clinical_relevance": clinical_relevance,
        "patient_explanation": patient_explanation,
        "conduct": conduct
    }

def get_content_for_item(item):
    """Gerar conteúdo para um item específico"""
    subgroup = item["subgroup"]
    item_name = item["name"]

    if "Percepção de futuro" in subgroup:
        return get_content_for_future_perception(item_name)
    elif "Adesão a planos" in subgroup:
        return get_content_for_adherence(item_name)
    elif "Verificar objetivos" in subgroup:
        return get_content_for_initial_goals(item_name)
    else:
        return None

def update_item(token, item_id, content):
    """Atualizar item via API"""
    try:
        # Converter de snake_case para camelCase
        payload = {
            "clinicalRelevance": content["clinical_relevance"],
            "patientExplanation": content["patient_explanation"],
            "conduct": content["conduct"]
        }

        headers = {"Authorization": f"Bearer {token}"}
        response = requests.put(
            f"{API_URL}/score-items/{item_id}",
            json=payload,
            headers=headers
        )
        response.raise_for_status()
        return True
    except Exception as e:
        print(f"  ✗ Erro: {e}")
        return False

def main():
    print("="*80)
    print("Enriquecimento de Score Items - Grupo Objetivos")
    print("="*80)
    print()

    # Login
    token = login()
    if not token:
        return

    # Buscar items
    items = get_items_from_db()
    if not items:
        return

    print()
    print("="*80)
    print(f"Processando {len(items)} items")
    print("="*80)
    print()

    success = 0
    failed = 0
    skipped = 0

    for i, item in enumerate(items, 1):
        print(f"[{i}/{len(items)}] {item['name']}")
        print(f"           Subgrupo: {item['subgroup']}")

        # Gerar conteúdo
        content = get_content_for_item(item)

        if not content:
            print(f"           ⊘ Sem template para este item")
            skipped += 1
            print()
            continue

        # Atualizar
        if update_item(token, item["id"], content):
            print(f"           ✓ Atualizado com sucesso")
            success += 1
        else:
            print(f"           ✗ Falha na atualização")
            failed += 1

        print()

    print("="*80)
    print("Resumo:")
    print(f"  ✓ Sucesso:  {success}/{len(items)}")
    print(f"  ✗ Falhas:   {failed}/{len(items)}")
    print(f"  ⊘ Pulados:  {skipped}/{len(items)}")
    print("="*80)

if __name__ == "__main__":
    main()
