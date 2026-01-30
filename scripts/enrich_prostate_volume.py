#!/usr/bin/env python3
"""
Enriquecimento clínico: USG Próstata - Volume Prostático
ID: 23b012f9-ce8b-4a1d-95f4-cfe9183295e0
"""

import os
import sys
import psycopg2
import anthropic
from datetime import datetime

# Database configuration
DB_CONFIG = {
    'host': 'localhost',
    'port': 5432,
    'user': 'plenya_user',
    'password': 'plenya_dev_password',
    'database': 'plenya_db'
}

ITEM_ID = '23b012f9-ce8b-4a1d-95f4-cfe9183295e0'

def get_db_connection():
    """Estabelece conexão com PostgreSQL"""
    return psycopg2.connect(**DB_CONFIG)

def verify_item_exists(conn):
    """Verifica se o item existe"""
    with conn.cursor() as cur:
        cur.execute("""
            SELECT id, name, subgroup_id
            FROM score_items
            WHERE id = %s
        """, (ITEM_ID,))
        result = cur.fetchone()
        if result:
            print(f"✓ Item encontrado: {result[1]} (ID: {result[0]})")
            return True
        else:
            print(f"✗ ERRO: Item {ITEM_ID} não encontrado no banco")
            return False

def search_articles():
    """Busca artigos científicos via Claude Web Search"""
    client = anthropic.Anthropic(api_key=os.environ.get("ANTHROPIC_API_KEY"))

    prompt = """Find 3 high-quality scientific articles about prostate volume measurement and benign prostatic hyperplasia (BPH).

Focus on:
1. Normal prostate volume by age
2. BPH diagnosis and volume thresholds
3. Surgical indications based on prostate volume

Return ONLY a JSON array with this exact structure:
[
  {
    "title": "Article title",
    "authors": "Author names",
    "year": 2023,
    "journal": "Journal name",
    "doi": "10.xxxx/xxxxx",
    "url": "https://...",
    "key_findings": "Brief summary of key findings relevant to prostate volume"
  }
]

Search terms to use: "prostate volume measurement BPH" OR "benign prostatic hyperplasia volume surgical indication"
"""

    print("\n🔍 Buscando artigos científicos sobre volume prostático...")

    message = client.messages.create(
        model="claude-sonnet-4-5-20250929",
        max_tokens=4000,
        tools=[{
            "type": "web_search_20250124",
            "name": "web_search",
            "display_name": "Web Search",
            "display_number": 1
        }],
        messages=[{"role": "user", "content": prompt}]
    )

    # Extract JSON from response
    response_text = ""
    for block in message.content:
        if hasattr(block, 'text'):
            response_text += block.text

    print(f"✓ Artigos encontrados")
    return response_text

def generate_enrichment_content(articles_json):
    """Gera conteúdo clínico enriquecido em PT-BR"""
    client = anthropic.Anthropic(api_key=os.environ.get("ANTHROPIC_API_KEY"))

    prompt = f"""Você é um médico urologista especializado em hiperplasia prostática benigna (HBP).

Com base nos artigos científicos abaixo, crie conteúdo clínico em PORTUGUÊS BRASILEIRO para o item:
**USG Próstata - Volume Prostático**

Artigos científicos:
{articles_json}

Retorne APENAS um JSON válido com esta estrutura EXATA:

{{
  "clinical_relevance": "Explicação técnica para médicos sobre volume prostático, valores normais por idade, relação com HBP, indicações de tratamento. 2-3 parágrafos, linguagem médica. Inclua valores de referência (ex: <30mL normal, 30-50mL HBP leve, >80mL HBP severa).",

  "patient_explanation": "Explicação simples para o paciente sobre o que é volume prostático, por que aumenta com idade, quando é preocupante. Use linguagem acessível, evite jargão. 2-3 parágrafos.",

  "conduct": "Protocolo clínico objetivo: quando repetir exame, quando encaminhar para urologista, quando considerar tratamento medicamentoso ou cirúrgico baseado em volume. Use bullet points ou parágrafos curtos."
}}

IMPORTANTE:
- Use linguagem técnica em clinical_relevance
- Use linguagem simples em patient_explanation
- Seja objetivo em conduct
- Base tudo nos artigos fornecidos
- Valores numéricos devem ser precisos
"""

    print("\n🤖 Gerando conteúdo clínico com Claude...")

    message = client.messages.create(
        model="claude-sonnet-4-5-20250929",
        max_tokens=3000,
        messages=[{"role": "user", "content": prompt}]
    )

    content = message.content[0].text
    print(f"✓ Conteúdo gerado ({len(content)} caracteres)")

    return content

def save_articles_to_db(conn, articles_json):
    """Salva artigos no banco e retorna IDs"""
    import json

    try:
        articles = json.loads(articles_json)
    except json.JSONDecodeError as e:
        print(f"✗ ERRO ao parsear JSON de artigos: {e}")
        return []

    article_ids = []

    with conn.cursor() as cur:
        for article in articles:
            # Verifica se artigo já existe (por DOI ou URL)
            cur.execute("""
                SELECT id FROM articles
                WHERE doi = %s OR url = %s
                LIMIT 1
            """, (article.get('doi'), article.get('url')))

            existing = cur.fetchone()

            if existing:
                article_id = existing[0]
                print(f"  ↪ Artigo já existe: {article['title'][:60]}... (ID: {article_id})")
            else:
                # Insere novo artigo
                cur.execute("""
                    INSERT INTO articles (
                        title, authors, publication_year, journal,
                        doi, url, summary, created_at, updated_at
                    )
                    VALUES (%s, %s, %s, %s, %s, %s, %s, NOW(), NOW())
                    RETURNING id
                """, (
                    article['title'],
                    article.get('authors'),
                    article.get('year'),
                    article.get('journal'),
                    article.get('doi'),
                    article.get('url'),
                    article.get('key_findings')
                ))

                article_id = cur.fetchone()[0]
                print(f"  ✓ Artigo salvo: {article['title'][:60]}... (ID: {article_id})")

            article_ids.append(article_id)

        conn.commit()

    return article_ids

def link_articles_to_item(conn, article_ids):
    """Cria relações entre artigos e score_item"""
    with conn.cursor() as cur:
        for article_id in article_ids:
            # Verifica se relação já existe
            cur.execute("""
                SELECT 1 FROM article_score_items
                WHERE score_item_id = %s AND article_id = %s
            """, (ITEM_ID, article_id))

            if cur.fetchone():
                print(f"  ↪ Relação já existe: item ↔ article {article_id}")
            else:
                cur.execute("""
                    INSERT INTO article_score_items (score_item_id, article_id)
                    VALUES (%s, %s)
                """, (ITEM_ID, article_id))
                print(f"  ✓ Relação criada: item ↔ article {article_id}")

        conn.commit()

def update_score_item(conn, enrichment_json):
    """Atualiza score_item com conteúdo enriquecido"""
    import json

    try:
        enrichment = json.loads(enrichment_json)
    except json.JSONDecodeError as e:
        print(f"✗ ERRO ao parsear JSON de enriquecimento: {e}")
        return False

    with conn.cursor() as cur:
        cur.execute("""
            UPDATE score_items
            SET
                clinical_relevance = %s,
                patient_explanation = %s,
                conduct = %s,
                last_review = NOW(),
                updated_at = NOW()
            WHERE id = %s
        """, (
            enrichment['clinical_relevance'],
            enrichment['patient_explanation'],
            enrichment['conduct'],
            ITEM_ID
        ))

        conn.commit()

        if cur.rowcount > 0:
            print(f"✓ Score item atualizado (ID: {ITEM_ID})")
            return True
        else:
            print(f"✗ ERRO: Nenhuma linha atualizada")
            return False

def verify_enrichment(conn):
    """Verifica o resultado final"""
    with conn.cursor() as cur:
        cur.execute("""
            SELECT
                name,
                LENGTH(clinical_relevance) as clinical_len,
                LENGTH(patient_explanation) as patient_len,
                LENGTH(conduct) as conduct_len,
                last_review,
                (SELECT COUNT(*) FROM article_score_items WHERE score_item_id = %s) as article_count
            FROM score_items
            WHERE id = %s
        """, (ITEM_ID, ITEM_ID))

        result = cur.fetchone()

        if result:
            print("\n" + "="*70)
            print("📊 VERIFICAÇÃO FINAL")
            print("="*70)
            print(f"Item: {result[0]}")
            print(f"Clinical Relevance: {result[1]} caracteres")
            print(f"Patient Explanation: {result[2]} caracteres")
            print(f"Conduct: {result[3]} caracteres")
            print(f"Last Review: {result[4]}")
            print(f"Artigos vinculados: {result[5]}")
            print("="*70)

            if result[1] > 0 and result[2] > 0 and result[3] > 0 and result[5] > 0:
                print("✅ ENRIQUECIMENTO COMPLETO!")
                return True
            else:
                print("⚠️  ENRIQUECIMENTO INCOMPLETO")
                return False

        return False

def main():
    """Executa workflow completo"""
    print("="*70)
    print("🧬 ENRIQUECIMENTO CLÍNICO: USG Próstata - Volume Prostático")
    print("="*70)
    print(f"Item ID: {ITEM_ID}")
    print(f"Data: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}")
    print("="*70)

    try:
        # 1. Conectar ao banco
        print("\n📦 Conectando ao banco de dados...")
        conn = get_db_connection()
        print("✓ Conexão estabelecida")

        # 2. Verificar se item existe
        if not verify_item_exists(conn):
            return 1

        # 3. Buscar artigos
        articles_json = search_articles()

        # 4. Salvar artigos
        print("\n💾 Salvando artigos no banco...")
        article_ids = save_articles_to_db(conn, articles_json)
        print(f"✓ {len(article_ids)} artigos processados")

        # 5. Vincular artigos ao item
        print("\n🔗 Vinculando artigos ao score_item...")
        link_articles_to_item(conn, article_ids)

        # 6. Gerar conteúdo enriquecido
        enrichment_json = generate_enrichment_content(articles_json)

        # 7. Atualizar score_item
        print("\n💾 Salvando enriquecimento no banco...")
        if not update_score_item(conn, enrichment_json):
            return 1

        # 8. Verificar resultado
        if not verify_enrichment(conn):
            return 1

        conn.close()

        print("\n" + "="*70)
        print("🎉 PROCESSO CONCLUÍDO COM SUCESSO!")
        print("="*70)
        return 0

    except Exception as e:
        print(f"\n❌ ERRO: {e}")
        import traceback
        traceback.print_exc()
        return 1

if __name__ == "__main__":
    sys.exit(main())
