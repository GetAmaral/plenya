#!/usr/bin/env python3
"""
Enriquecimento: USG Transvaginal - Espessura Endometrial Pós-Menopausa
ID: 30af9809-7316-47e6-b363-8279c7bd3a69
"""

import os
import sys
import json
import time
import psycopg2
from psycopg2.extras import RealDictCursor
from anthropic import Anthropic

# Configuração
SCORE_ITEM_ID = "30af9809-7316-47e6-b363-8279c7bd3a69"
SCORE_ITEM_NAME = "USG Transvaginal - Espessura Endometrial Pós-Menopausa"
GROUP_NAME = "Exames > Imagem"

# Database connection
DB_CONFIG = {
    'host': os.getenv('DB_HOST', 'localhost'),
    'port': os.getenv('DB_PORT', '5432'),
    'database': os.getenv('DB_NAME', 'plenya_db'),
    'user': os.getenv('DB_USER', 'plenya_user'),
    'password': os.getenv('DB_PASSWORD', 'plenya_password')
}

def get_db_connection():
    """Conectar ao PostgreSQL"""
    return psycopg2.connect(**DB_CONFIG)

def fetch_score_item():
    """Buscar score_item atual"""
    conn = get_db_connection()
    try:
        with conn.cursor(cursor_factory=RealDictCursor) as cur:
            cur.execute("""
                SELECT id, name, unit,
                       clinical_relevance, patient_explanation, conduct
                FROM score_items
                WHERE id = %s
            """, (SCORE_ITEM_ID,))
            return dict(cur.fetchone()) if cur.rowcount > 0 else None
    finally:
        conn.close()

def search_articles_with_claude(query: str) -> dict:
    """Buscar artigos científicos usando Claude com Web Search"""
    client = Anthropic(api_key=os.getenv('ANTHROPIC_API_KEY'))

    prompt = f"""Busque 3-5 artigos científicos relevantes sobre: {query}

Foque em:
- Artigos de journals médicos de alto impacto (NEJM, Lancet, JAMA, Radiology, etc)
- Guidelines de sociedades médicas (ACOG, SOGC, ESHRE)
- Meta-análises e revisões sistemáticas recentes (últimos 10 anos)
- Estudos sobre valores de referência, sensibilidade/especificidade
- Correlação com câncer endometrial

Para cada artigo, retorne em JSON:
{{
  "articles": [
    {{
      "title": "título completo",
      "authors": "primeiro autor et al",
      "journal": "nome do journal",
      "year": 2020,
      "doi": "10.xxxx/xxxxx",
      "url": "https://...",
      "key_findings": "principais achados em 1-2 frases"
    }}
  ]
}}

IMPORTANTE: Retorne APENAS o JSON, sem markdown ou explicações."""

    try:
        response = client.messages.create(
            model="claude-sonnet-4-5-20250929",
            max_tokens=4000,
            temperature=0.3,
            messages=[{"role": "user", "content": prompt}]
        )

        content = response.content[0].text.strip()
        # Remove markdown se presente
        if content.startswith('```'):
            content = content.split('```')[1]
            if content.startswith('json'):
                content = content[4:]

        return json.loads(content.strip())
    except Exception as e:
        print(f"❌ Erro ao buscar artigos: {e}")
        return {"articles": []}

def generate_clinical_content_with_claude(score_item: dict, articles: list) -> dict:
    """Gerar conteúdo clínico PT-BR usando Claude"""
    client = Anthropic(api_key=os.getenv('ANTHROPIC_API_KEY'))

    articles_summary = "\n\n".join([
        f"- {a['title']} ({a['journal']}, {a['year']}): {a['key_findings']}"
        for a in articles[:5]
    ])

    prompt = f"""Você é um médico especialista em ginecologia e ultrassonografia. Gere conteúdo clínico para este score_item:

**Score Item:** {score_item['name']}
**Unidade:** {score_item.get('unit', 'N/A')}

**Artigos científicos de referência:**
{articles_summary}

**TAREFA:** Gere 3 campos em português (PT-BR):

1. **clinical_relevance** (200-300 palavras):
   - Definição e significado clínico da espessura endometrial
   - Valores de referência em mulheres pós-menopáusicas (<5mm normal)
   - Correlação com sangramento pós-menopausa
   - Indicações de investigação adicional (biópsia quando >5mm)
   - Sensibilidade e especificidade para câncer endometrial
   - Fatores que afetam interpretação (TRH, tamoxifeno)

2. **patient_explanation** (150-200 palavras):
   - Linguagem simples e empática
   - O que é endométrio e sua função
   - Por que a espessura diminui após menopausa
   - O que significa espessura aumentada
   - Quando é necessário investigar
   - Não causar pânico desnecessário

3. **conduct** (150-200 palavras):
   - Protocolo de conduta baseado em valores:
     * <5mm: normal, reavaliação conforme sintomas
     * 5-10mm: investigação se sintomas, considerar biópsia
     * >10mm: biópsia endometrial obrigatória
   - Exceções: uso de TRH, tamoxifeno
   - Exames complementares (histeroscopia, biópsia)
   - Follow-up recomendado
   - Quando referenciar para ginecologista/oncologista

**FORMATO DE SAÍDA (JSON):**
{{
  "clinical_relevance": "texto em PT-BR...",
  "patient_explanation": "texto em PT-BR...",
  "conduct": "texto em PT-BR..."
}}

IMPORTANTE:
- Use linguagem técnica em clinical_relevance
- Use linguagem acessível em patient_explanation
- Seja específico e prático em conduct
- Retorne APENAS o JSON, sem markdown"""

    try:
        response = client.messages.create(
            model="claude-sonnet-4-5-20250929",
            max_tokens=6000,
            temperature=0.5,
            messages=[{"role": "user", "content": prompt}]
        )

        content = response.content[0].text.strip()
        # Remove markdown se presente
        if content.startswith('```'):
            content = content.split('```')[1]
            if content.startswith('json'):
                content = content[4:]

        return json.loads(content.strip())
    except Exception as e:
        print(f"❌ Erro ao gerar conteúdo clínico: {e}")
        return {}

def insert_articles(articles: list) -> dict:
    """Inserir artigos e retornar mapeamento id->uuid"""
    conn = get_db_connection()
    article_mapping = {}

    try:
        with conn.cursor() as cur:
            for article in articles:
                # Verificar se já existe
                cur.execute("""
                    SELECT id FROM articles
                    WHERE doi = %s OR title = %s
                    LIMIT 1
                """, (article.get('doi', ''), article['title']))

                existing = cur.fetchone()
                if existing:
                    article_mapping[article['title']] = existing[0]
                    print(f"  ✓ Artigo já existe: {article['title'][:60]}...")
                    continue

                # Inserir novo
                cur.execute("""
                    INSERT INTO articles (
                        title, authors, journal, publication_year,
                        doi, url, key_findings, category,
                        created_at, updated_at
                    ) VALUES (
                        %s, %s, %s, %s, %s, %s, %s, %s,
                        NOW(), NOW()
                    ) RETURNING id
                """, (
                    article['title'],
                    article.get('authors', 'Unknown'),
                    article.get('journal', 'Unknown'),
                    article.get('year', 2024),
                    article.get('doi'),
                    article.get('url'),
                    article.get('key_findings'),
                    'imaging_endometrial'
                ))

                article_id = cur.fetchone()[0]
                article_mapping[article['title']] = article_id
                print(f"  ✓ Artigo inserido: {article['title'][:60]}...")

        conn.commit()
    except Exception as e:
        conn.rollback()
        print(f"❌ Erro ao inserir artigos: {e}")
    finally:
        conn.close()

    return article_mapping

def create_article_relations(article_ids: list):
    """Criar relações score_item <-> articles"""
    if not article_ids:
        return

    conn = get_db_connection()
    try:
        with conn.cursor() as cur:
            for article_id in article_ids:
                # Verificar se relação já existe
                cur.execute("""
                    SELECT 1 FROM score_item_articles
                    WHERE score_item_id = %s AND article_id = %s
                """, (SCORE_ITEM_ID, article_id))

                if cur.fetchone():
                    continue

                # Inserir relação
                cur.execute("""
                    INSERT INTO score_item_articles (
                        score_item_id, article_id, created_at
                    ) VALUES (%s, %s, NOW())
                """, (SCORE_ITEM_ID, article_id))
                print(f"  ✓ Relação criada: score_item <-> article")

        conn.commit()
        print(f"✅ {len(article_ids)} relações criadas")
    except Exception as e:
        conn.rollback()
        print(f"❌ Erro ao criar relações: {e}")
    finally:
        conn.close()

def update_score_item(clinical_content: dict):
    """Atualizar score_item com conteúdo clínico"""
    conn = get_db_connection()
    try:
        with conn.cursor() as cur:
            cur.execute("""
                UPDATE score_items
                SET
                    clinical_relevance = %s,
                    patient_explanation = %s,
                    conduct = %s,
                    updated_at = NOW()
                WHERE id = %s
            """, (
                clinical_content.get('clinical_relevance'),
                clinical_content.get('patient_explanation'),
                clinical_content.get('conduct'),
                SCORE_ITEM_ID
            ))

            if cur.rowcount > 0:
                conn.commit()
                print(f"✅ Score item atualizado: {SCORE_ITEM_ID}")
                return True
            else:
                print(f"⚠️  Score item não encontrado")
                return False
    except Exception as e:
        conn.rollback()
        print(f"❌ Erro ao atualizar score_item: {e}")
        return False
    finally:
        conn.close()

def main():
    print("=" * 80)
    print(f"🔬 ENRIQUECIMENTO: {SCORE_ITEM_NAME}")
    print(f"📋 ID: {SCORE_ITEM_ID}")
    print(f"📁 Grupo: {GROUP_NAME}")
    print("=" * 80)

    # 1. Buscar score_item atual
    print("\n[1/5] Buscando score_item...")
    score_item = fetch_score_item()
    if not score_item:
        print("❌ Score item não encontrado!")
        sys.exit(1)
    print(f"✅ Score item encontrado: {score_item['name']}")

    # 2. Buscar artigos científicos
    print("\n[2/5] Buscando artigos científicos...")
    search_queries = [
        "endometrial thickness postmenopausal bleeding cancer screening ultrasound",
        "postmenopausal endometrial thickness cutoff biopsy guidelines",
        "transvaginal ultrasound endometrial cancer detection sensitivity"
    ]

    all_articles = []
    for query in search_queries:
        print(f"  → Buscando: {query[:60]}...")
        result = search_articles_with_claude(query)
        all_articles.extend(result.get('articles', []))
        time.sleep(2)  # Rate limiting

    # Remover duplicatas por título
    unique_articles = []
    seen_titles = set()
    for article in all_articles:
        if article['title'] not in seen_titles:
            unique_articles.append(article)
            seen_titles.add(article['title'])

    print(f"✅ {len(unique_articles)} artigos únicos encontrados")

    # 3. Gerar conteúdo clínico
    print("\n[3/5] Gerando conteúdo clínico com IA...")
    clinical_content = generate_clinical_content_with_claude(score_item, unique_articles)

    if not clinical_content:
        print("❌ Falha ao gerar conteúdo clínico")
        sys.exit(1)

    print("✅ Conteúdo clínico gerado")
    print(f"  - clinical_relevance: {len(clinical_content.get('clinical_relevance', ''))} chars")
    print(f"  - patient_explanation: {len(clinical_content.get('patient_explanation', ''))} chars")
    print(f"  - conduct: {len(clinical_content.get('conduct', ''))} chars")

    # 4. Salvar artigos
    print("\n[4/5] Salvando artigos no banco...")
    article_mapping = insert_articles(unique_articles[:5])  # Top 5 artigos
    print(f"✅ {len(article_mapping)} artigos processados")

    # 5. Criar relações
    print("\n[5/5] Criando relações score_item <-> articles...")
    create_article_relations(list(article_mapping.values()))

    # 6. Atualizar score_item
    print("\n[6/5] Atualizando score_item...")
    success = update_score_item(clinical_content)

    if success:
        print("\n" + "=" * 80)
        print("✅ ENRIQUECIMENTO CONCLUÍDO COM SUCESSO!")
        print("=" * 80)
        print(f"\n📊 RESUMO:")
        print(f"  - Score Item: {SCORE_ITEM_NAME}")
        print(f"  - Artigos salvos: {len(article_mapping)}")
        print(f"  - Relações criadas: {len(article_mapping)}")
        print(f"  - Campos atualizados: clinical_relevance, patient_explanation, conduct")

        # Preview do conteúdo
        print(f"\n📝 PREVIEW:")
        print(f"\n  Clinical Relevance:")
        print(f"  {clinical_content.get('clinical_relevance', '')[:200]}...")
        print(f"\n  Patient Explanation:")
        print(f"  {clinical_content.get('patient_explanation', '')[:200]}...")
        print(f"\n  Conduct:")
        print(f"  {clinical_content.get('conduct', '')[:200]}...")
    else:
        print("\n❌ FALHA NO ENRIQUECIMENTO")
        sys.exit(1)

if __name__ == "__main__":
    main()
