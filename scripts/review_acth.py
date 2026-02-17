#!/usr/bin/env python3
"""
Script de revisão completa do ScoreItem ACTH
Segue workflow: verificar artigos → buscar RAG → buscar PubMed → enriquecer LLM → salvar
"""

import os
import sys
import json
import time
import requests
import psycopg2
from datetime import datetime
from typing import List, Dict, Optional
import anthropic

# Configuração
ACTH_ID = "c77cedd3-2800-77e4-b510-d4cccf642c0d"
MIN_ARTICLES = 7
RAG_THRESHOLD = 0.7

# Database connection
DB_CONFIG = {
    'host': 'localhost',
    'port': 5432,
    'user': 'plenya_user',
    'password': 'plenya_dev_password',
    'database': 'plenya_db'
}

# API Keys
ANTHROPIC_API_KEY = os.getenv('ANTHROPIC_API_KEY')
if not ANTHROPIC_API_KEY:
    raise ValueError("ANTHROPIC_API_KEY environment variable is required")
PUBMED_EMAIL = os.getenv('PUBMED_EMAIL', 'contato@plenya.com.br')


def get_db_connection():
    """Conectar ao banco PostgreSQL"""
    return psycopg2.connect(**DB_CONFIG)


def fetch_score_item(conn, item_id: str) -> Dict:
    """Buscar ScoreItem do banco"""
    with conn.cursor() as cur:
        cur.execute("""
            SELECT id, name, unit, points, clinical_relevance,
                   patient_explanation, conduct, last_review
            FROM score_items
            WHERE id = %s
        """, (item_id,))

        row = cur.fetchone()
        if not row:
            raise ValueError(f"ScoreItem {item_id} não encontrado")

        return {
            'id': str(row[0]),
            'name': row[1],
            'unit': row[2],
            'points': float(row[3]) if row[3] else 0,
            'clinical_relevance': row[4],
            'patient_explanation': row[5],
            'conduct': row[6],
            'last_review': row[7]
        }


def fetch_linked_articles(conn, item_id: str) -> List[Dict]:
    """Buscar artigos já linkados ao ScoreItem"""
    with conn.cursor() as cur:
        cur.execute("""
            SELECT a.id, a.pm_id, a.title, EXTRACT(YEAR FROM a.publish_date)::integer as year,
                   a.internal_link, a.abstract
            FROM articles a
            INNER JOIN article_score_items asi ON a.id = asi.article_id
            WHERE asi.score_item_id = %s
            ORDER BY a.publish_date DESC
        """, (item_id,))

        articles = []
        for row in cur.fetchall():
            articles.append({
                'id': str(row[0]),
                'pmid': row[1],
                'title': row[2],
                'year': row[3],
                'has_pdf': bool(row[4]),
                'abstract': row[5]
            })

        return articles


def search_pubmed(query: str, max_results: int = 10) -> List[str]:
    """Buscar artigos no PubMed"""
    base_url = "https://eutils.ncbi.nlm.nih.gov/entrez/eutils/esearch.fcgi"

    params = {
        'db': 'pubmed',
        'term': query,
        'retmax': max_results,
        'retmode': 'json',
        'tool': 'plenya',
        'email': PUBMED_EMAIL
    }

    response = requests.get(base_url, params=params)
    response.raise_for_status()

    data = response.json()
    return data.get('esearchresult', {}).get('idlist', [])


def fetch_pubmed_details(pmids: List[str]) -> List[Dict]:
    """Buscar detalhes de artigos do PubMed"""
    if not pmids:
        return []

    base_url = "https://eutils.ncbi.nlm.nih.gov/entrez/eutils/esummary.fcgi"

    params = {
        'db': 'pubmed',
        'id': ','.join(pmids),
        'retmode': 'json',
        'tool': 'plenya',
        'email': PUBMED_EMAIL
    }

    response = requests.get(base_url, params=params)
    response.raise_for_status()

    data = response.json()
    result = data.get('result', {})

    articles = []
    for pmid in pmids:
        if pmid in result:
            article_data = result[pmid]
            articles.append({
                'pmid': pmid,
                'title': article_data.get('title', ''),
                'authors': ', '.join([a.get('name', '') for a in article_data.get('authors', [])[:5]]),
                'journal': article_data.get('source', ''),
                'year': int(article_data.get('pubdate', '2024').split()[0])
            })

    return articles


def create_article_in_db(conn, article: Dict) -> str:
    """Criar artigo no banco e retornar ID"""
    with conn.cursor() as cur:
        # Verificar se já existe
        cur.execute("SELECT id FROM articles WHERE pm_id = %s", (article['pmid'],))
        existing = cur.fetchone()
        if existing:
            return str(existing[0])

        # Criar novo
        cur.execute("""
            INSERT INTO articles
            (id, title, authors, journal, publish_date, language, pm_id, article_type, embedding_status, chunk_count)
            VALUES (uuid_generate_v7(), %s, %s, %s, %s, 'en', %s, 'research_article', 'pending', 0)
            RETURNING id
        """, (
            article['title'],
            article['authors'],
            article['journal'],
            f"{article['year']}-01-01",
            article['pmid']
        ))

        article_id = cur.fetchone()[0]
        conn.commit()
        return str(article_id)


def link_article_to_score_item(conn, article_id: str, score_item_id: str):
    """Linkar artigo ao ScoreItem"""
    with conn.cursor() as cur:
        # Verificar se já existe
        cur.execute("""
            SELECT 1 FROM article_score_items
            WHERE article_id = %s AND score_item_id = %s
        """, (article_id, score_item_id))

        if cur.fetchone():
            return  # Já existe

        cur.execute("""
            INSERT INTO article_score_items (article_id, score_item_id, auto_linked, confidence_score)
            VALUES (%s, %s, true, 0.8)
            ON CONFLICT (score_item_id, article_id) DO NOTHING
        """, (article_id, score_item_id))

        conn.commit()


def enrich_with_llm(item: Dict, articles: List[Dict]) -> Dict:
    """Gerar enriquecimento usando Claude API"""

    # Construir contexto de artigos
    articles_context = "Artigos científicos disponíveis:\n\n"
    for i, article in enumerate(articles[:10], 1):
        articles_context += f"{i}. [{article.get('pmid', 'N/A')}] {article['title']} ({article.get('year', 'N/A')})\n"
        if article.get('abstract'):
            abstract = article['abstract'][:500] + "..." if len(article.get('abstract', '')) > 500 else article.get('abstract', '')
            articles_context += f"   Resumo: {abstract}\n"
        articles_context += "\n"

    # Construir prompt
    prompt = f"""Você é especialista clínico escrevendo conteúdo educacional para sistema EMR.

ScoreItem: {item['name']}
Unidade: {item.get('unit', 'N/A')}
Pontuação atual: {item['points']}

{articles_context}

TAREFA: GERAR campos clínicos completos sobre ACTH (hormônio adrenocorticotrópico), baseado nos artigos fornecidos.

**IMPORTANTE**: Os dados atuais estão ERRADOS (são sobre Anti-TPO, não ACTH). Você deve REESCREVER completamente com informações corretas sobre ACTH.

**Diretrizes de formato:**

1. clinical_relevance (1000-1500 caracteres):
   - Definição: O que é ACTH (hormônio adrenocorticotrópico)
   - Fisiologia: Papel no eixo HPA (hipotálamo-pituitária-adrenal)
   - Função: Estímulo da produção de cortisol pelas adrenais
   - Indicações clínicas: Quando dosar (suspeita de Cushing, Addison, tumores hipofisários)
   - Dados epidemiológicos e significância clínica quantificada
   - Interpretação: valores baixos vs altos

2. patient_explanation (500-800 caracteres):
   - Linguagem acessível (8º ano escolar)
   - O QUE é ACTH (hormônio que controla cortisol)
   - POR QUE importa (regula resposta ao estresse)
   - O QUE significam valores anormais
   - Tom: empático e empoderador (evitar alarmismo)

3. conduct (800-1200 caracteres):
   - Estruturado com headers e bullets markdown
   - Valores de referência normais (pg/mL)
   - Workup diagnóstico quando alterado
   - Exames complementares (cortisol, teste de estímulo, imagem)
   - Protocolos de tratamento baseados em evidências
   - Critérios de referral ao endocrinologista
   - Periodicidade de monitoramento

4. max_points (0-50):
   - Baseado em importância clínica do ACTH
   - ACTH é marcador crucial de disfunção adrenal e hipofisária
   - Sugestão: 25-35 pontos (alta relevância clínica)
   - Justificativa breve

Retorne APENAS JSON válido (sem markdown, sem explicações):
{{
  "clinical_relevance": "...",
  "patient_explanation": "...",
  "conduct": "...",
  "max_points": 30,
  "justification": "...",
  "confidence": 0.92
}}"""

    # Chamar Claude API
    client = anthropic.Anthropic(api_key=ANTHROPIC_API_KEY)

    try:
        message = client.messages.create(
            model="claude-sonnet-4-5-20250929",
            max_tokens=4000,
            messages=[{
                "role": "user",
                "content": prompt
            }]
        )

        # Extrair resposta
        response_text = message.content[0].text

        # Parse JSON
        response_text = response_text.strip()
        if response_text.startswith('```json'):
            response_text = response_text[7:]
        if response_text.startswith('```'):
            response_text = response_text[3:]
        if response_text.endswith('```'):
            response_text = response_text[:-3]
        response_text = response_text.strip()

        result = json.loads(response_text)
        return result

    except Exception as e:
        print(f"❌ Erro ao chamar Claude API: {e}")
        raise


def save_enrichment(conn, item_id: str, enrichment: Dict):
    """Salvar enriquecimento no banco"""
    with conn.cursor() as cur:
        # Atualizar ScoreItem
        cur.execute("""
            UPDATE score_items
            SET clinical_relevance = %s,
                patient_explanation = %s,
                conduct = %s,
                points = %s,
                last_review = NOW(),
                updated_at = NOW()
            WHERE id = %s
        """, (
            enrichment['clinical_relevance'],
            enrichment['patient_explanation'],
            enrichment['conduct'],
            enrichment['max_points'],
            item_id
        ))

        conn.commit()
        print(f"✅ ScoreItem atualizado com sucesso!")


def main():
    print("=== REVISÃO COMPLETA: ScoreItem ACTH ===\n")

    # Conectar ao banco
    conn = get_db_connection()

    try:
        # 1. Buscar ScoreItem
        print("1. Buscando ScoreItem...")
        item = fetch_score_item(conn, ACTH_ID)
        print(f"   ✓ {item['name']} (ID: {item['id']})")
        print(f"   Unidade: {item.get('unit', 'N/A')}")
        print(f"   Pontos atuais: {item['points']}")

        # 2. Verificar artigos existentes
        print("\n2. Verificando artigos linkados...")
        articles = fetch_linked_articles(conn, ACTH_ID)
        print(f"   Total: {len(articles)} artigos")

        scientific_count = sum(1 for a in articles if a['pmid'])
        print(f"   Científicos (com PMID): {scientific_count}")

        # Mostrar alguns
        for i, article in enumerate(articles[:3], 1):
            pmid = article['pmid'] or 'N/A'
            print(f"   {i}. [PMID: {pmid}] {article['title'][:80]}...")

        # 3. Buscar mais artigos no PubMed se necessário
        if scientific_count < MIN_ARTICLES:
            needed = MIN_ARTICLES - scientific_count
            print(f"\n3. Buscando {needed} artigos adicionais no PubMed...")

            query = "ACTH[tiab] AND (reference values[tiab] OR normal range[tiab] OR diagnosis[tiab] OR clinical significance[tiab] OR cushing[tiab] OR addison[tiab]) AND (Review[ptyp] OR Meta-Analysis[ptyp]) AND 2019:2026[dp]"
            print(f"   Query: {query[:100]}...")

            pmids = search_pubmed(query, max_results=needed * 2)
            print(f"   PMIDs encontrados: {len(pmids)}")

            if pmids:
                print("   Buscando detalhes...")
                pubmed_articles = fetch_pubmed_details(pmids)
                print(f"   Artigos obtidos: {len(pubmed_articles)}")

                # Salvar e linkar artigos
                for article in pubmed_articles[:needed]:
                    article_id = create_article_in_db(conn, article)
                    link_article_to_score_item(conn, article_id, ACTH_ID)
                    articles.append(article)  # Adicionar à lista
                    print(f"   ✓ Linkado: {article['title'][:60]}...")

                time.sleep(0.5)  # Rate limiting PubMed
        else:
            print(f"\n3. ✓ Artigos suficientes ({scientific_count} ≥ {MIN_ARTICLES})")

        # 4. Gerar enriquecimento LLM
        print("\n4. Gerando enriquecimento com Claude Sonnet 4.5...")
        print("   ⚠️  Conteúdo atual INCORRETO: dados de Anti-TPO, não ACTH")
        print("   Modo: REWRITE (reescrever do zero)\n")

        enrichment = enrich_with_llm(item, articles)

        # 5. Mostrar preview
        print("=== PREVIEW DO RESULTADO ===\n")
        print(f"Confidence: {enrichment['confidence']:.2f}")
        print(f"Max Points: {enrichment['max_points']}\n")

        print("Clinical Relevance:")
        print(enrichment['clinical_relevance'][:300] + "...\n")

        print("Patient Explanation:")
        print(enrichment['patient_explanation'][:200] + "...\n")

        print("Conduct:")
        print(enrichment['conduct'][:300] + "...\n")

        print(f"Justification: {enrichment['justification']}\n")

        # 6. Confirmar e salvar
        response = input("Salvar mudanças no banco de dados? (y/N): ")
        if response.lower() == 'y':
            save_enrichment(conn, ACTH_ID, enrichment)
            print(f"\n✅ CONCLUÍDO! Last Review: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}")
        else:
            print("\n❌ Abortado pelo usuário")

    finally:
        conn.close()


if __name__ == '__main__':
    main()
