#!/usr/bin/env python3
"""
Enriquecimento do item T3 Livre com busca de artigos científicos e conteúdo clínico.
"""

import anthropic
import os
import psycopg2
from datetime import datetime

# Configuração
ITEM_ID = "d164eacf-a0d7-48f2-899d-3f0d57ec7cc3"
ITEM_NAME = "T3 Livre"

# Cliente Anthropic
client = anthropic.Anthropic(api_key=os.environ.get("ANTHROPIC_API_KEY"))

def get_db_connection():
    """Conectar ao PostgreSQL"""
    return psycopg2.connect(
        host="db",
        port=5432,
        dbname="plenya_db",
        user="plenya_user",
        password="plenya_password"
    )

def search_articles():
    """Buscar artigos científicos sobre T3 Livre"""

    print("🔍 Buscando artigos científicos sobre T3 Livre...")

    prompt = """Você é um especialista em endocrinologia. Preciso de 3 artigos científicos de referência sobre T3 Livre (triiodotironina livre).

CRITÉRIOS:
- Artigos de periódicos indexados (PubMed, Endocrine Reviews, JCEM, etc.)
- Foco em interpretação clínica, valores de referência, T3 toxicosis
- Publicações dos últimos 10 anos preferencialmente
- Artigos em inglês de alta qualidade

Para cada artigo, forneça:
1. Título completo
2. Autores (primeiro autor et al. se mais de 3)
3. Journal e ano
4. DOI ou PMID
5. Resumo de 2-3 linhas sobre o conteúdo relevante

Formato JSON:
{
  "articles": [
    {
      "title": "...",
      "authors": "...",
      "journal": "...",
      "year": 2023,
      "doi": "...",
      "pmid": "...",
      "summary": "..."
    }
  ]
}"""

    message = client.messages.create(
        model="claude-sonnet-4-5-20250929",
        max_tokens=4000,
        temperature=0.3,
        messages=[{"role": "user", "content": prompt}]
    )

    return message.content[0].text

def generate_clinical_content(articles_info):
    """Gerar conteúdo clínico em PT-BR baseado nos artigos"""

    print("📝 Gerando conteúdo clínico em português...")

    prompt = f"""Você é um endocrinologista especializado em doenças tireoidianas. Com base nos artigos científicos abaixo, crie conteúdo clínico sobre T3 Livre.

ARTIGOS DE REFERÊNCIA:
{articles_info}

INFORMAÇÕES TÉCNICAS:
- Nome: T3 Livre (Triiodotironina Livre)
- Tipo: Hormônio tireoidiano ativo
- Valores de referência típicos: 2.3-4.2 pg/mL (pode variar por laboratório)
- Patologias: Hipertireoidismo, T3 toxicosis, hipotireoidismo

Gere 3 campos em PORTUGUÊS BRASILEIRO:

1. **clinical_relevance** (relevância clínica para médicos):
   - O que é T3 livre e sua função fisiológica
   - Quando solicitar o exame
   - Interpretação de valores elevados e reduzidos
   - Diferença entre T3 total e T3 livre
   - T3 toxicosis e sua importância diagnóstica
   - Relação com TSH e T4 livre
   - 200-300 palavras, linguagem técnica médica

2. **patient_explanation** (explicação para pacientes):
   - O que é o T3 livre em linguagem simples
   - Para que serve o exame
   - O que significam valores alterados
   - Sintomas associados
   - 150-200 palavras, linguagem acessível, empática

3. **conduct** (conduta clínica):
   - Protocolo de investigação quando T3 livre está alterado
   - Exames complementares necessários
   - Quando encaminhar ao endocrinologista
   - Orientações sobre reavaliação
   - 150-200 palavras, formato de lista ou parágrafos curtos

FORMATO DE SAÍDA (JSON):
{{
  "clinical_relevance": "...",
  "patient_explanation": "...",
  "conduct": "..."
}}

IMPORTANTE:
- Use terminologia médica correta em PT-BR
- Seja preciso e baseado em evidências
- Mantenha tom profissional mas acessível no patient_explanation
- Cite valores de referência quando relevante"""

    message = client.messages.create(
        model="claude-sonnet-4-5-20250929",
        max_tokens=6000,
        temperature=0.3,
        messages=[{"role": "user", "content": prompt}]
    )

    return message.content[0].text

def save_articles(conn, articles_data):
    """Salvar artigos no banco e retornar IDs"""

    print("💾 Salvando artigos no banco de dados...")

    import json

    # Parse do JSON
    articles = json.loads(articles_data)["articles"]
    article_ids = []

    cursor = conn.cursor()

    for article in articles:
        # Verificar se artigo já existe por DOI ou PMID
        doi = article.get("doi", "")
        pmid = article.get("pmid", "")

        if doi:
            cursor.execute(
                "SELECT id FROM articles WHERE doi = %s",
                (doi,)
            )
        elif pmid:
            cursor.execute(
                "SELECT id FROM articles WHERE pmid = %s",
                (pmid,)
            )
        else:
            cursor.execute(
                "SELECT id FROM articles WHERE title = %s",
                (article["title"],)
            )

        existing = cursor.fetchone()

        if existing:
            article_id = existing[0]
            print(f"   ✓ Artigo já existe: {article['title'][:50]}...")
        else:
            # Inserir novo artigo
            cursor.execute("""
                INSERT INTO articles (
                    title, authors, journal, publication_year,
                    doi, pmid, summary, category, created_at, updated_at
                )
                VALUES (%s, %s, %s, %s, %s, %s, %s, %s, NOW(), NOW())
                RETURNING id
            """, (
                article["title"],
                article["authors"],
                article["journal"],
                article["year"],
                article.get("doi"),
                article.get("pmid"),
                article.get("summary"),
                "Endocrinologia"
            ))

            article_id = cursor.fetchone()[0]
            print(f"   ✓ Artigo salvo: {article['title'][:50]}...")

        article_ids.append(article_id)

    conn.commit()
    cursor.close()

    return article_ids

def link_articles_to_item(conn, article_ids):
    """Criar relações entre artigos e score_item"""

    print("🔗 Vinculando artigos ao item T3 Livre...")

    cursor = conn.cursor()

    for article_id in article_ids:
        # Verificar se relação já existe
        cursor.execute("""
            SELECT 1 FROM article_score_items
            WHERE article_id = %s AND score_item_id = %s
        """, (article_id, ITEM_ID))

        if cursor.fetchone():
            print(f"   ⚠ Relação já existe: {article_id}")
            continue

        # Criar relação
        cursor.execute("""
            INSERT INTO article_score_items (article_id, score_item_id, created_at)
            VALUES (%s, %s, NOW())
        """, (article_id, ITEM_ID))

        print(f"   ✓ Artigo vinculado: {article_id}")

    conn.commit()
    cursor.close()

def update_score_item(conn, clinical_content):
    """Atualizar score_item com conteúdo clínico"""

    print("📊 Atualizando score_item com conteúdo clínico...")

    import json

    # Parse do JSON
    content = json.loads(clinical_content)

    cursor = conn.cursor()

    cursor.execute("""
        UPDATE score_items
        SET
            clinical_relevance = %s,
            patient_explanation = %s,
            conduct = %s,
            last_review = %s,
            updated_at = NOW()
        WHERE id = %s
    """, (
        content["clinical_relevance"],
        content["patient_explanation"],
        content["conduct"],
        datetime.now(),
        ITEM_ID
    ))

    conn.commit()
    cursor.close()

    print(f"   ✓ Score item atualizado: {ITEM_ID}")

def main():
    """Fluxo principal"""

    print("=" * 80)
    print(f"🚀 ENRIQUECIMENTO: {ITEM_NAME}")
    print(f"   ID: {ITEM_ID}")
    print("=" * 80)
    print()

    try:
        # 1. Buscar artigos
        articles_data = search_articles()
        print()
        print("📄 Artigos encontrados:")
        print(articles_data)
        print()

        # 2. Gerar conteúdo clínico
        clinical_content = generate_clinical_content(articles_data)
        print()
        print("📋 Conteúdo clínico gerado:")
        print(clinical_content)
        print()

        # 3. Conectar ao banco
        conn = get_db_connection()

        # 4. Salvar artigos
        article_ids = save_articles(conn, articles_data)
        print(f"   ✓ {len(article_ids)} artigos processados")
        print()

        # 5. Vincular artigos ao item
        link_articles_to_item(conn, article_ids)
        print()

        # 6. Atualizar score_item
        update_score_item(conn, clinical_content)
        print()

        conn.close()

        print("=" * 80)
        print("✅ ENRIQUECIMENTO CONCLUÍDO COM SUCESSO!")
        print("=" * 80)

    except Exception as e:
        print()
        print("=" * 80)
        print(f"❌ ERRO: {str(e)}")
        print("=" * 80)
        raise

if __name__ == "__main__":
    main()
