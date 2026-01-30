#!/usr/bin/env python3
"""
Enriquecimento: Hemácias - Mulheres
Item ID: 501fd84a-a440-4c13-9b11-35e2f69017d1
"""

import os
import re
import anthropic
import psycopg2
from datetime import datetime

# Configuração
ITEM_ID = "501fd84a-a440-4c13-9b11-35e2f69017d1"
ITEM_NAME = "Hemácias - Mulheres"

def get_db_connection():
    """Conecta ao PostgreSQL"""
    return psycopg2.connect(
        host=os.getenv("DB_HOST", "db"),
        port=os.getenv("DB_PORT", "5432"),
        database=os.getenv("DB_NAME", "plenya_db"),
        user=os.getenv("DB_USER", "plenya_user"),
        password=os.getenv("DB_PASSWORD", "plenya_password")
    )

def search_articles(client):
    """Busca artigos científicos via Claude com web search"""
    print(f"\n{'='*80}")
    print(f"BUSCANDO ARTIGOS CIENTÍFICOS: {ITEM_NAME}")
    print(f"{'='*80}\n")

    search_prompt = """Busque 3-4 artigos científicos sobre contagem de hemácias (eritrócitos) em mulheres:

TÓPICOS:
- Red blood cell count women (reference ranges 4.0-5.2 million/µL)
- Hemoglobin levels women vs men (sex differences)
- Iron deficiency anemia menstruation
- Menstrual blood loss and erythrocyte count
- Hormonal effects on RBC production

FONTES PRIORITÁRIAS: PubMed, Google Scholar, UpToDate

Para cada artigo, extrair:
1. Título completo
2. Autores principais
3. DOI/PMID
4. Ano
5. 2-3 insights clínicos principais

Formato JSON:
{
  "articles": [
    {
      "title": "...",
      "authors": "...",
      "doi": "...",
      "year": 2023,
      "key_findings": ["...", "..."]
    }
  ]
}"""

    message = client.messages.create(
        model="claude-sonnet-4-5-20250929",
        max_tokens=4000,
        tools=[{
            "type": "web_search_tool",
            "name": "web_search",
            "display_name": "Web Search",
            "display_number": 1,
            "is_enabled": True
        }],
        tool_choice={"type": "any"},
        messages=[{"role": "user", "content": search_prompt}]
    )

    # Extrair conteúdo da resposta
    articles_text = ""
    for block in message.content:
        if hasattr(block, 'text'):
            articles_text += block.text

    print(f"\nARTIGOS ENCONTRADOS:\n{articles_text[:500]}...\n")
    return articles_text

def enrich_content(client, articles_text):
    """Gera conteúdo clínico enriquecido em PT-BR"""
    print(f"\n{'='*80}")
    print(f"GERANDO CONTEÚDO CLÍNICO EM PT-BR")
    print(f"{'='*80}\n")

    enrichment_prompt = f"""Com base nos artigos científicos abaixo, gere conteúdo clínico para o item "{ITEM_NAME}":

ARTIGOS:
{articles_text}

CONTEXTO CLÍNICO:
- Hemácias (eritrócitos): células vermelhas que transportam oxigênio
- Valores de referência mulheres: 4.0-5.2 milhões/µL (vs homens 4.5-5.5)
- Diferenças por menstruação (perda ferro), hormônios, gravidez
- Baixo = anemia; alto = policitemia/desidratação

GERAR EM PT-BR (formato JSON):

{{
  "clinical_relevance": "Texto médico 150-200 palavras. Por que hemácias são menores em mulheres? Causas de alterações (menstruação, ferro, B12). Quando investigar?",

  "patient_explanation": "Texto leigo 100-150 palavras. 'Suas hemácias são as células que levam oxigênio...'. Explicar valores, sintomas anemia (cansaço, palidez).",

  "conduct": "Texto objetivo 80-120 palavras. Se baixo (<4.0): ferritina, ferro, B12, reticulócitos. Se alto (>5.2): hidratação, EPO, policitemia vera. Quando referenciar hematologista."
}}

REGRAS:
- PT-BR técnico mas compreensível
- Focar diferenças de gênero (menstruação, ferro)
- Incluir valores numéricos
- Ser objetivo e prático
- JSON válido (escapar aspas)"""

    message = client.messages.create(
        model="claude-sonnet-4-5-20250929",
        max_tokens=3000,
        messages=[{"role": "user", "content": enrichment_prompt}]
    )

    content = message.content[0].text
    print(f"CONTEÚDO GERADO:\n{content[:400]}...\n")

    # Extrair JSON
    json_match = re.search(r'\{.*\}', content, re.DOTALL)
    if json_match:
        import json
        return json.loads(json_match.group())
    return None

def save_articles(conn, articles_text):
    """Salva artigos no banco e retorna IDs"""
    print(f"\n{'='*80}")
    print(f"SALVANDO ARTIGOS NO BANCO")
    print(f"{'='*80}\n")

    article_ids = []

    # Parse simples de artigos (busca padrões comuns)
    article_pattern = r'(?:Title|Título):\s*(.+?)(?:\n|$).*?(?:DOI|PMID):\s*(.+?)(?:\n|$)'
    matches = re.finditer(article_pattern, articles_text, re.IGNORECASE | re.DOTALL)

    cursor = conn.cursor()

    for i, match in enumerate(matches, 1):
        title = match.group(1).strip()[:300]
        doi = match.group(2).strip()[:200]

        # Inserir artigo
        cursor.execute("""
            INSERT INTO articles (
                title,
                url,
                content,
                article_type,
                created_at,
                updated_at
            ) VALUES (%s, %s, %s, %s, NOW(), NOW())
            RETURNING id
        """, (
            f"{title} - RBC Women",
            f"https://doi.org/{doi}" if doi else f"https://pubmed.ncbi.nlm.nih.gov/search?term={title[:50]}",
            articles_text[match.start():match.start()+500],
            "scientific"
        ))

        article_id = cursor.fetchone()[0]
        article_ids.append(article_id)

        print(f"✓ Artigo {i} salvo: {article_id} - {title[:60]}...")

        # Criar relação com score_item
        cursor.execute("""
            INSERT INTO score_item_articles (
                score_item_id,
                article_id,
                created_at
            ) VALUES (%s, %s, NOW())
            ON CONFLICT DO NOTHING
        """, (ITEM_ID, article_id))

    conn.commit()
    cursor.close()

    print(f"\n✓ {len(article_ids)} artigos salvos e vinculados")
    return article_ids

def update_score_item(conn, enriched_data):
    """Atualiza score_item com conteúdo enriquecido"""
    print(f"\n{'='*80}")
    print(f"ATUALIZANDO SCORE ITEM: {ITEM_ID}")
    print(f"{'='*80}\n")

    cursor = conn.cursor()

    # Buscar dados atuais
    cursor.execute("""
        SELECT name, clinical_relevance, patient_explanation, conduct
        FROM score_items
        WHERE id = %s
    """, (ITEM_ID,))

    current = cursor.fetchone()
    if current:
        print(f"Dados atuais:")
        print(f"  Name: {current[0]}")
        print(f"  Clinical: {len(current[1] or '')} chars")
        print(f"  Patient: {len(current[2] or '')} chars")
        print(f"  Conduct: {len(current[3] or '')} chars")

    # Atualizar
    cursor.execute("""
        UPDATE score_items
        SET
            clinical_relevance = %s,
            patient_explanation = %s,
            conduct = %s,
            updated_at = NOW()
        WHERE id = %s
    """, (
        enriched_data['clinical_relevance'],
        enriched_data['patient_explanation'],
        enriched_data['conduct'],
        ITEM_ID
    ))

    conn.commit()
    cursor.close()

    print(f"\n✓ Score item atualizado com sucesso!")
    print(f"  Clinical: {len(enriched_data['clinical_relevance'])} chars")
    print(f"  Patient: {len(enriched_data['patient_explanation'])} chars")
    print(f"  Conduct: {len(enriched_data['conduct'])} chars")

def main():
    """Execução principal"""
    print(f"\n{'#'*80}")
    print(f"# ENRIQUECIMENTO: {ITEM_NAME}")
    print(f"# ID: {ITEM_ID}")
    print(f"# Timestamp: {datetime.now().isoformat()}")
    print(f"{'#'*80}")

    # Verificar API key
    api_key = os.getenv("ANTHROPIC_API_KEY")
    if not api_key:
        print("\n❌ ERRO: ANTHROPIC_API_KEY não configurada")
        return 1

    try:
        # Inicializar Claude
        client = anthropic.Anthropic(api_key=api_key)

        # 1. Buscar artigos
        articles_text = search_articles(client)

        # 2. Gerar conteúdo enriquecido
        enriched_data = enrich_content(client, articles_text)

        if not enriched_data:
            print("\n❌ ERRO: Falha ao gerar conteúdo enriquecido")
            return 1

        # 3. Conectar ao banco
        conn = get_db_connection()

        # 4. Salvar artigos
        article_ids = save_articles(conn, articles_text)

        # 5. Atualizar score_item
        update_score_item(conn, enriched_data)

        conn.close()

        print(f"\n{'#'*80}")
        print(f"# ✓ ENRIQUECIMENTO CONCLUÍDO COM SUCESSO")
        print(f"# Artigos salvos: {len(article_ids)}")
        print(f"# Item atualizado: {ITEM_ID}")
        print(f"{'#'*80}\n")

        return 0

    except Exception as e:
        print(f"\n❌ ERRO: {str(e)}")
        import traceback
        traceback.print_exc()
        return 1

if __name__ == "__main__":
    exit(main())
