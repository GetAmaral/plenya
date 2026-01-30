#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Enriquecimento: TC Tórax - Maior Nódulo Sólido
ID: dd6e920c-b203-4d40-b230-55f2074ac613
"""

import anthropic
import psycopg2
import os
import json
import sys
from datetime import datetime

# Configurações
ANTHROPIC_API_KEY = os.environ.get("ANTHROPIC_API_KEY")
DB_CONFIG = {
    "host": "localhost",
    "port": 5432,
    "database": "plenya_db",
    "user": "plenya_user",
    "password": "plenya_dev_password"
}

ITEM_ID = "dd6e920c-b203-4d40-b230-55f2074ac613"

def get_db_connection():
    """Conecta ao banco PostgreSQL"""
    return psycopg2.connect(**DB_CONFIG)

def fetch_item_details(conn):
    """Busca detalhes do item"""
    with conn.cursor() as cur:
        cur.execute("""
            SELECT id, name, subgroup_id
            FROM score_items
            WHERE id = %s
        """, (ITEM_ID,))
        return cur.fetchone()

def search_articles_web(query):
    """Busca artigos via Claude Web Search"""
    if not ANTHROPIC_API_KEY:
        print("ERROR: ANTHROPIC_API_KEY não encontrada")
        return []

    client = anthropic.Anthropic(api_key=ANTHROPIC_API_KEY)

    print(f"\n🔍 Buscando artigos: {query}")

    prompt = f"""
Busque artigos científicos sobre: {query}

Foco em:
1. Guidelines Fleischner Society
2. Manejo de nódulos pulmonares solitários
3. Screening de câncer de pulmão
4. Critérios de seguimento e malignidade

Para cada artigo relevante, forneça:
- Título completo
- Autores principais
- Journal/revista
- Ano de publicação
- DOI ou URL
- Resumo de 2-3 frases sobre o conteúdo

Busque 3-5 artigos de alta qualidade (últimos 10 anos preferencialmente).
"""

    try:
        message = client.messages.create(
            model="claude-sonnet-4-5-20250929",
            max_tokens=4000,
            temperature=0.2,
            messages=[{
                "role": "user",
                "content": prompt
            }],
            tools=[{
                "type": "web_search",
                "name": "web_search",
                "description": "Search the web for articles"
            }]
        )

        # Extrai resposta
        result_text = ""
        for block in message.content:
            if hasattr(block, 'text'):
                result_text += block.text

        print(f"✅ Busca concluída: {len(result_text)} caracteres")
        return result_text

    except Exception as e:
        print(f"❌ Erro na busca: {str(e)}")
        return ""

def generate_clinical_content(item_question, articles_text):
    """Gera conteúdo clínico enriquecido usando Claude"""
    if not ANTHROPIC_API_KEY:
        print("ERROR: ANTHROPIC_API_KEY não encontrada")
        return None

    client = anthropic.Anthropic(api_key=ANTHROPIC_API_KEY)

    print(f"\n🤖 Gerando conteúdo clínico enriquecido...")

    prompt = f"""
Você é um especialista em radiologia torácica e pneumologia.

**ITEM DO QUESTIONÁRIO:**
{item_question}

**ARTIGOS CIENTÍFICOS ENCONTRADOS:**
{articles_text}

**TAREFA:**
Gere conteúdo clínico em PT-BR para este item, baseado nos artigos científicos.

**FORMATO JSON (OBRIGATÓRIO):**
{{
  "clinical_relevance": "Explicação técnica para médicos (200-300 palavras). Inclua: significado clínico, guidelines Fleischner, critérios de malignidade (tamanho, contornos, densidade), indicações de seguimento vs biópsia, diagnósticos diferenciais.",

  "patient_explanation": "Explicação simples para pacientes (150-200 palavras). Use linguagem acessível, evite jargão médico, explique o que é um nódulo pulmonar, por que precisa ser acompanhado, quando há preocupação.",

  "conduct": "Condutas práticas (150-200 palavras). Inclua: protocolo de seguimento por tamanho (Fleischner), quando indicar TC follow-up, PET-CT, biópsia, cirurgia. Fatores de risco a considerar.",

  "articles": [
    {{
      "title": "Título completo do artigo",
      "authors": "Autor et al.",
      "journal": "Nome da revista",
      "year": 2023,
      "doi": "10.xxxx/xxxxx",
      "url": "https://...",
      "summary": "Resumo relevante do artigo"
    }}
  ]
}}

**CRITÉRIOS DE QUALIDADE:**
- Baseado em evidências científicas dos artigos fornecidos
- Português brasileiro correto
- Tom profissional mas acessível
- Foco prático e aplicável
- Inclua valores numéricos quando relevante (ex: "nódulos > 8mm")

Retorne APENAS o JSON, sem markdown ou texto adicional.
"""

    try:
        message = client.messages.create(
            model="claude-sonnet-4-5-20250929",
            max_tokens=6000,
            temperature=0.3,
            messages=[{
                "role": "user",
                "content": prompt
            }]
        )

        # Extrai resposta
        result_text = ""
        for block in message.content:
            if hasattr(block, 'text'):
                result_text += block.text

        # Parse JSON
        result_text = result_text.strip()
        if result_text.startswith("```json"):
            result_text = result_text.replace("```json", "").replace("```", "").strip()

        content = json.loads(result_text)
        print(f"✅ Conteúdo gerado com sucesso")
        return content

    except json.JSONDecodeError as e:
        print(f"❌ Erro ao parsear JSON: {str(e)}")
        print(f"Resposta recebida: {result_text[:500]}")
        return None
    except Exception as e:
        print(f"❌ Erro na geração: {str(e)}")
        return None

def save_articles(conn, articles):
    """Salva artigos científicos no banco"""
    article_ids = []

    with conn.cursor() as cur:
        for article in articles:
            try:
                # Verifica se artigo já existe (por DOI ou título)
                doi = article.get('doi')
                title = article['title']

                if doi:
                    cur.execute("""
                        SELECT id FROM articles
                        WHERE doi = %s
                        LIMIT 1
                    """, (doi,))
                else:
                    cur.execute("""
                        SELECT id FROM articles
                        WHERE title = %s
                        LIMIT 1
                    """, (title,))

                existing = cur.fetchone()
                if existing:
                    article_ids.append(existing[0])
                    print(f"   ↪ Artigo já existe: {title[:60]}...")
                    continue

                # Insere novo artigo
                # Campos obrigatórios: title, authors, journal, publish_date
                publish_year = article.get('year', 2024)
                publish_date = f"{publish_year}-01-01"

                cur.execute("""
                    INSERT INTO articles (
                        title, authors, journal, publish_date,
                        doi, original_link, abstract,
                        article_type, specialty,
                        created_at, updated_at
                    )
                    VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, NOW(), NOW())
                    RETURNING id
                """, (
                    title,
                    article.get('authors', 'Unknown'),
                    article.get('journal', 'Unknown Journal'),
                    publish_date,
                    doi,
                    article.get('url'),
                    article.get('summary'),
                    'research_article',
                    'Radiologia/Pneumologia'
                ))

                article_id = cur.fetchone()[0]
                article_ids.append(article_id)
                print(f"   ✅ Artigo salvo: {title[:60]}...")

            except Exception as e:
                print(f"   ❌ Erro ao salvar artigo '{article.get('title', 'N/A')[:40]}': {str(e)}")
                continue

    conn.commit()
    return article_ids

def link_articles_to_item(conn, item_id, article_ids):
    """Cria relações entre item e artigos"""
    if not article_ids:
        return

    with conn.cursor() as cur:
        for article_id in article_ids:
            try:
                # Verifica se relação já existe
                cur.execute("""
                    SELECT 1 FROM article_score_items
                    WHERE score_item_id = %s AND article_id = %s
                """, (item_id, article_id))

                if cur.fetchone():
                    continue

                # Insere relação
                cur.execute("""
                    INSERT INTO article_score_items (score_item_id, article_id)
                    VALUES (%s, %s)
                """, (item_id, article_id))

            except Exception as e:
                print(f"   ❌ Erro ao vincular artigo: {str(e)}")
                continue

    conn.commit()
    print(f"   ✅ {len(article_ids)} artigos vinculados ao item")

def update_item_content(conn, item_id, content):
    """Atualiza campos clínicos do item"""
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
            content['clinical_relevance'],
            content['patient_explanation'],
            content['conduct'],
            item_id
        ))

    conn.commit()
    print(f"✅ Item atualizado no banco")

def main():
    """Execução principal"""
    print("="*70)
    print("🔬 ENRIQUECIMENTO: TC Tórax - Maior Nódulo Sólido")
    print(f"📋 ID: {ITEM_ID}")
    print("="*70)

    # Conecta ao banco
    try:
        conn = get_db_connection()
        print("✅ Conectado ao banco de dados")
    except Exception as e:
        print(f"❌ Erro ao conectar ao banco: {str(e)}")
        return 1

    try:
        # 1. Busca detalhes do item
        print("\n📊 ETAPA 1: Buscando item no banco...")
        item = fetch_item_details(conn)
        if not item:
            print(f"❌ Item não encontrado: {ITEM_ID}")
            return 1

        item_id, item_name, subgroup_id = item
        print(f"   Item: {item_name}")
        print(f"   Subgrupo: {subgroup_id}")

        # 2. Busca artigos científicos
        print("\n📚 ETAPA 2: Buscando artigos científicos...")
        search_queries = [
            "pulmonary nodule management Fleischner guidelines 2023",
            "solitary pulmonary nodule malignancy risk",
            "lung cancer screening CT nodule follow-up"
        ]

        all_articles_text = ""
        for query in search_queries:
            articles_text = search_articles_web(query)
            all_articles_text += f"\n\n{articles_text}"

        if not all_articles_text.strip():
            print("⚠️  Nenhum artigo encontrado, usando conhecimento base")

        # 3. Gera conteúdo clínico
        print("\n✍️  ETAPA 3: Gerando conteúdo clínico...")
        content = generate_clinical_content(item_name, all_articles_text)

        if not content:
            print("❌ Falha ao gerar conteúdo")
            return 1

        # 4. Salva artigos no banco
        print("\n💾 ETAPA 4: Salvando artigos no banco...")
        article_ids = []
        if 'articles' in content and content['articles']:
            article_ids = save_articles(conn, content['articles'])
            print(f"   Total: {len(article_ids)} artigos processados")

        # 5. Vincula artigos ao item
        if article_ids:
            print("\n🔗 ETAPA 5: Vinculando artigos ao item...")
            link_articles_to_item(conn, item_id, article_ids)

        # 6. Atualiza conteúdo do item
        print("\n💾 ETAPA 6: Atualizando conteúdo do item...")
        update_item_content(conn, item_id, content)

        # 7. Resumo final
        print("\n" + "="*70)
        print("✅ ENRIQUECIMENTO CONCLUÍDO COM SUCESSO!")
        print("="*70)
        print(f"📋 Item ID: {item_id}")
        print(f"📚 Artigos vinculados: {len(article_ids)}")
        print(f"📝 Clinical relevance: {len(content['clinical_relevance'])} chars")
        print(f"👤 Patient explanation: {len(content['patient_explanation'])} chars")
        print(f"🏥 Conduct: {len(content['conduct'])} chars")
        print("="*70)

        return 0

    except Exception as e:
        print(f"\n❌ ERRO: {str(e)}")
        import traceback
        traceback.print_exc()
        return 1

    finally:
        conn.close()
        print("\n🔌 Conexão com banco encerrada")

if __name__ == "__main__":
    sys.exit(main())
