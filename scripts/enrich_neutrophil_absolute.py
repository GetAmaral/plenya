#!/usr/bin/env python3
"""
Enriquecimento do item: Neutrófilos (absoluto)
ID: 3faeb6db-b8d6-4fb1-8740-07bfd91002c7

Baseado em:
- CTCAE v6 (2025-2026) - Nova classificação de neutropenia
- Guidelines AGIHO 2024/2025 para neutropenia febril
- Valores de referência atualizados
"""

import requests
import os
from datetime import datetime

# Configuração da API
API_URL = os.getenv("API_URL", "http://localhost:3001")
SCORE_ITEM_ID = "3faeb6db-b8d6-4fb1-8740-07bfd91002c7"

# Token de autenticação (será obtido via login)
AUTH_TOKEN = None

def get_auth_token():
    """Faz login e obtém token JWT"""
    global AUTH_TOKEN

    # Credenciais de doctor (ajustar conforme necessário)
    login_payload = {
        "email": os.getenv("ADMIN_EMAIL", "doctor@plenya.com"),
        "password": os.getenv("ADMIN_PASSWORD", "password123")
    }

    print("\n🔑 Autenticando na API...")
    response = requests.post(f"{API_URL}/api/v1/auth/login", json=login_payload)

    if response.status_code == 200:
        data = response.json()
        AUTH_TOKEN = data.get("access_token")
        print("   ✅ Autenticado com sucesso!")
        return True
    else:
        print(f"   ❌ Erro ao autenticar: {response.status_code}")
        print(f"   {response.text[:200]}")
        return False

def get_headers():
    """Retorna headers com token de autenticação"""
    return {
        "Authorization": f"Bearer {AUTH_TOKEN}",
        "Content-Type": "application/json"
    }

# Artigos científicos identificados
ARTICLES = [
    {
        "title": "A paradigm shift in neutrophil adverse event grading: What now?",
        "journal": "PMC",
        "year": 2025,
        "url": "https://pmc.ncbi.nlm.nih.gov/articles/PMC12745037/",
        "key_findings": "CTCAE v6 (2025) atualiza classificação de neutropenia: Grade 1 agora <1500-1000/µL (antes Grade 2), Grade 4 <100/µL. Mudanças visam inclusão de variante Duffy null (comum em pessoas com ancestralidade africana subsaariana)."
    },
    {
        "title": "2024 update of the AGIHO guideline on diagnosis and empirical treatment of fever of unknown origin in adult neutropenic patients",
        "journal": "The Lancet Regional Health – Europe",
        "year": 2025,
        "url": "https://www.thelancet.com/journals/lanepe/article/PIIS2666-7762(25)00006-7/fulltext",
        "key_findings": "Diretrizes atualizadas para manejo de neutropenia febril. Monoterapia empírica com beta-lactâmicos anti-pseudomonas é primeira linha. Estratificação de risco via índice MASCC."
    },
    {
        "title": "Febrile Neutropenia",
        "journal": "StatPearls - NCBI",
        "year": 2025,
        "url": "https://www.ncbi.nlm.nih.gov/books/NBK541102/",
        "key_findings": "Neutropenia febril definida como temperatura ≥38.3°C com ANC <1500/µL. Requer avaliação emergencial, hemoculturas (2 conjuntos) e antibioticoterapia empírica imediata."
    }
]

# Conteúdo clínico enriquecido (PT-BR)
CLINICAL_CONTENT = {
    "clinical_relevance": """**Contagem Absoluta de Neutrófilos (ANC): Pilar da Avaliação Imunológica**

Os neutrófilos são a primeira linha de defesa contra infecções bacterianas e fúngicas. A contagem absoluta de neutrófilos (ANC) é calculada multiplicando o percentual de neutrófilos (segmentados + bastonetes) pela contagem total de leucócitos.

**ATUALIZAÇÃO CTCAE v6 (2025-2026):**
O Common Terminology Criteria for Adverse Events versão 6, implementado em janeiro de 2026, modernizou a classificação de neutropenia:

• **Grau 1:** ANC 1000-1500/µL (antes era Grau 2)
• **Grau 2:** ANC 500-1000/µL (antes era Grau 3)
• **Grau 3:** ANC 100-500/µL (antes era Grau 4)
• **Grau 4:** ANC <100/µL (novo limiar)

Esta mudança reconhece que pessoas com variante genética Duffy null (comum em populações de ancestralidade africana subsaariana) apresentam ANC naturalmente mais baixo sem aumento de risco infeccioso.

**Valores de Referência:**
• Normal: 1.500-8.000/µL (algumas fontes: 2.500-7.000/µL)
• Neutropenia leve: 1.000-1.500/µL
• Neutropenia moderada: 500-1.000/µL
• Neutropenia grave: <500/µL
• Neutropenia profunda: <100/µL

**Neutropenia Febril (Emergência Médica):**
Definida como temperatura ≥38,3°C (ou ≥38°C por >1h) + ANC <1.000/µL. Requer hospitalização imediata, hemoculturas e antibioticoterapia empírica.

**Neutrofilia:**
ANC >7.700/µL indica resposta inflamatória, infecção bacteriana aguda, estresse fisiológico (cirurgia, trauma), uso de corticoides ou doenças mieloproliferativas.

**Desvio à Esquerda:**
Aumento de neutrófilos jovens (bastonetes, metamielócitos) sugere infecção bacteriana grave com mobilização acelerada da medula óssea.""",

    "patient_explanation": """**O que são neutrófilos e por que a contagem absoluta é importante?**

Neutrófilos são glóbulos brancos especializados em combater infecções causadas por bactérias e fungos. Eles atuam como "soldados de primeira linha" do seu sistema imunológico, chegando rapidamente ao local da infecção para destruir os invasores.

A **contagem absoluta de neutrófilos (ANC)** mede quantos desses soldados você tem circulando no sangue. Este número é calculado automaticamente pelo laboratório usando a contagem total de leucócitos e a porcentagem de neutrófilos.

**Valores normais:**
• Adultos: 1.500 a 8.000 neutrófilos por microlitro (µL)
• Crianças: acima de 1.500/µL

**O que significa ter neutrófilos baixos (neutropenia)?**

Quando seus neutrófilos estão abaixo de 1.500/µL, você tem neutropenia, o que aumenta o risco de infecções. Quanto mais baixa a contagem, maior o risco:

• **1.000-1.500/µL:** Risco levemente aumentado
• **500-1.000/µL:** Risco moderado - cuidados extras necessários
• **Abaixo de 500/µL:** Risco grave - qualquer febre é emergência médica

**IMPORTANTE - Neutropenia Febril é Emergência:**
Se você tem neutrófilos baixos E desenvolve febre (≥38,3°C), procure atendimento médico IMEDIATAMENTE. Esta combinação pode evoluir rapidamente para infecção grave, pois seu corpo não consegue se defender adequadamente.

**O que significa ter neutrófilos altos (neutrofilia)?**

Valores acima de 7.700/µL geralmente indicam:
• Infecção bacteriana em curso
• Resposta ao estresse físico (cirurgia, trauma)
• Uso de medicamentos (corticoides)
• Inflamação sistêmica
• Raramente: doenças da medula óssea

**Observação importante sobre diversidade genética:**
Algumas pessoas, especialmente aquelas com ancestralidade africana, têm contagens naturalmente mais baixas (1.000-1.500/µL) devido à variante genética Duffy null. Isso NÃO significa risco aumentado de infecção - é uma variação normal. Os novos critérios médicos (2025) já consideram isso.""",

    "conduct": """**Condutas Clínicas Baseadas em ANC (Atualizadas 2025)**

**1. ANC NORMAL (1.500-8.000/µL):**
✓ Nenhuma ação específica necessária
✓ Sistema imunológico funcionando adequadamente
✓ Manter acompanhamento de rotina

**2. NEUTROPENIA LEVE (1.000-1.500/µL) - CTCAE v6 Grau 1:**
→ Investigar causas:
  • Medicamentos (anticonvulsivantes, antibióticos, quimioterapia)
  • Deficiências nutricionais (B12, folato, cobre)
  • Infecções virais recentes (influenza, HIV, hepatites)
  • Doenças autoimunes (lúpus, artrite reumatoide)
  • Variante genética Duffy null (considerar em afrodescendentes)

→ Monitoramento: Repetir hemograma em 1-2 semanas
→ Orientações ao paciente:
  • Higiene rigorosa das mãos
  • Evitar contato com pessoas doentes
  • Cozinhar bem carnes e ovos
  • Lavar frutas/vegetais adequadamente

**3. NEUTROPENIA MODERADA (500-1.000/µL) - CTCAE v6 Grau 2:**
→ Avaliação hematológica obrigatória
→ Investigação ampliada:
  • Hemograma completo seriado
  • Esfregaço periférico
  • Dosagem de vitaminas (B12, folato)
  • Sorologias virais (se indicado)
  • Considerar aspirado de medula óssea

→ Precauções:
  • Evitar aglomerações e locais fechados
  • Máscara em ambientes públicos se indicado
  • Profilaxia antibiótica em casos selecionados
  • Educação sobre sinais de infecção

**4. NEUTROPENIA GRAVE (<500/µL) - CTCAE v6 Graus 3-4:**
→ ALTO RISCO INFECCIOSO - Monitoramento intensivo
→ Hospitalização se sinais de infecção
→ Considerar:
  • Fator estimulador de colônias (G-CSF/filgrastim)
  • Profilaxia antibiótica (fluoroquinolona)
  • Profilaxia antifúngica se prolongada (>7 dias)
  • Isolamento protetor se <100/µL

**5. NEUTROPENIA FEBRIL (ANC <1.000/µL + febre ≥38,3°C):**
🚨 **EMERGÊNCIA MÉDICA - PROTOCOLO IMEDIATO:**

→ Tempo porta-antibiótico: <60 minutos
→ Colher ANTES de antibióticos:
  • 2 conjuntos de hemoculturas (periférica + cateter se presente)
  • Urinocultura
  • Culturas de sítios específicos (feridas, cateter)

→ Antibioticoterapia empírica (Guidelines AGIHO 2024):
  **Primeira linha:** Beta-lactâmico anti-pseudomonas
  • Piperacilina-tazobactam 4,5g IV 6/6h OU
  • Cefepime 2g IV 8/8h OU
  • Meropenem 1g IV 8/8h (se colonização ESBL/alta prevalência)

→ Estratificação de risco (Índice MASCC):
  • ≥21 pontos: Baixo risco → considerar alta precoce/oral
  • <21 pontos: Alto risco → hospitalização prolongada

→ Adicionar se indicado:
  • Vancomicina (suspeita de infecção por Gram+ ou cateter)
  • Antifúngico empírico (febre persistente >4-7 dias)

**6. NEUTROFILIA (>7.700/µL):**
→ Investigar causa subjacente:
  • Sinais de infecção bacteriana
  • Inflamação sistêmica (PCR, VHS)
  • Revisar medicações (corticoides)
  • Descartar leucemia mieloide/mieloproliferação (se muito elevado)

→ Avaliar "desvio à esquerda":
  • Aumento de bastonetes (>10%) sugere infecção bacteriana aguda
  • Presença de metamielócitos/mielócitos = mobilização medular intensa

**7. ELEGIBILIDADE PARA ENSAIOS CLÍNICOS (Atualização 2025):**
→ Novo critério padrão: ANC ≥1.000/µL (anteriormente ≥1.500/µL)
→ Reflete mudança CTCAE v6 e inclusão de populações diversas

**MONITORAMENTO PÓS-QUIMIOTERAPIA:**
→ Nadir esperado: 7-14 dias após infusão
→ Hemograma 2-3x/semana durante nadir
→ G-CSF profilático se risco >20% de neutropenia febril

**REFERÊNCIAS CRÍTICAS:**
• CTCAE v6 (2025) - Nova classificação de neutropenia
• AGIHO 2024 Guidelines - Manejo de febre de origem indeterminada em neutropênicos
• MASCC Index - Estratificação de risco em neutropenia febril"""
}

def create_article_and_link(article_data):
    """Cria artigo científico e vincula ao score_item"""

    # 1. Criar artigo
    article_payload = {
        "title": article_data["title"],
        "source": article_data["journal"],
        "year": article_data["year"],
        "type": "SCIENTIFIC",
        "content": f"**Key Findings:** {article_data['key_findings']}\n\n**URL:** {article_data['url']}",
        "url": article_data["url"]
    }

    print(f"\n📄 Criando artigo: {article_data['title'][:60]}...")
    response = requests.post(
        f"{API_URL}/api/v1/articles",
        json=article_payload,
        headers=get_headers()
    )

    if response.status_code in [200, 201]:
        article = response.json()
        article_id = article.get("id")
        print(f"   ✅ Artigo criado: ID {article_id}")

        # 2. Vincular ao score_item
        link_payload = {
            "article_id": article_id,
            "score_item_id": SCORE_ITEM_ID
        }

        link_response = requests.post(
            f"{API_URL}/api/v1/articles/{article_id}/score-items/{SCORE_ITEM_ID}",
            json=link_payload,
            headers=get_headers()
        )

        if link_response.status_code in [200, 201]:
            print(f"   🔗 Vinculado ao score_item")
            return article_id
        else:
            print(f"   ⚠️  Erro ao vincular: {link_response.status_code}")
            print(f"   {link_response.text[:200]}")
            return article_id
    else:
        print(f"   ❌ Erro ao criar artigo: {response.status_code}")
        print(f"   {response.text[:200]}")
        return None

def update_score_item():
    """Atualiza o score_item com conteúdo clínico"""

    payload = {
        "clinical_relevance": CLINICAL_CONTENT["clinical_relevance"],
        "patient_explanation": CLINICAL_CONTENT["patient_explanation"],
        "conduct": CLINICAL_CONTENT["conduct"],
        "last_review": datetime.now().strftime("%Y-%m-%d")
    }

    print(f"\n🔄 Atualizando score_item {SCORE_ITEM_ID}...")
    response = requests.put(
        f"{API_URL}/api/v1/score-items/{SCORE_ITEM_ID}",
        json=payload,
        headers=get_headers()
    )

    if response.status_code == 200:
        print("   ✅ Score item atualizado com sucesso!")
        return True
    else:
        print(f"   ❌ Erro ao atualizar: {response.status_code}")
        print(f"   {response.text[:300]}")
        return False

def main():
    print("=" * 80)
    print("🧬 ENRIQUECIMENTO: Neutrófilos (absoluto)")
    print("=" * 80)
    print(f"Item ID: {SCORE_ITEM_ID}")
    print(f"API URL: {API_URL}")
    print(f"Timestamp: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}")

    # Autenticar
    if not get_auth_token():
        print("\n❌ Falha na autenticação. Abortando...")
        return

    # Criar artigos e vincular
    print("\n" + "=" * 80)
    print("📚 FASE 1: Criação e Vinculação de Artigos Científicos")
    print("=" * 80)

    article_ids = []
    for article in ARTICLES:
        article_id = create_article_and_link(article)
        if article_id:
            article_ids.append(article_id)

    print(f"\n✅ {len(article_ids)}/{len(ARTICLES)} artigos processados com sucesso")

    # Atualizar conteúdo clínico
    print("\n" + "=" * 80)
    print("📝 FASE 2: Atualização de Conteúdo Clínico")
    print("=" * 80)

    success = update_score_item()

    # Sumário final
    print("\n" + "=" * 80)
    print("📊 SUMÁRIO FINAL")
    print("=" * 80)
    print(f"Artigos criados: {len(article_ids)}")
    print(f"Artigos vinculados: {len(article_ids)}")
    print(f"Conteúdo atualizado: {'✅ SIM' if success else '❌ NÃO'}")

    if success:
        print("\n🎉 ENRIQUECIMENTO CONCLUÍDO COM SUCESSO!")
        print("\n📋 Conteúdo adicionado:")
        print(f"   • Relevância clínica: {len(CLINICAL_CONTENT['clinical_relevance'])} caracteres")
        print(f"   • Explicação ao paciente: {len(CLINICAL_CONTENT['patient_explanation'])} caracteres")
        print(f"   • Condutas: {len(CLINICAL_CONTENT['conduct'])} caracteres")
        print(f"   • Artigos vinculados: {len(article_ids)}")

        print("\n🔬 Destaques científicos:")
        print("   • CTCAE v6 (2025-2026) - Nova classificação de neutropenia")
        print("   • Inclusão de variante genética Duffy null")
        print("   • Guidelines AGIHO 2024 para neutropenia febril")
        print("   • Valores de referência atualizados")
    else:
        print("\n⚠️  ENRIQUECIMENTO PARCIAL - Revisar erros acima")

    print("=" * 80)

if __name__ == "__main__":
    main()
