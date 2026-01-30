#!/usr/bin/env python3
"""
Script autônomo para enriquecimento de items com AI
Não requer confirmações - executa tudo automaticamente
"""

import subprocess
import json
import os
from anthropic import Anthropic

# Configurar cliente Anthropic
client = Anthropic(api_key=os.environ.get("ANTHROPIC_API_KEY"))

def run_sql(sql):
    """Executa SQL sem pedir confirmação"""
    cmd = ['docker', 'compose', 'exec', '-T', 'db', 'psql',
           '-U', 'plenya_user', '-d', 'plenya_db', '-c', sql]
    result = subprocess.run(cmd, capture_output=True, text=True)
    return result.returncode == 0, result.stdout, result.stderr

def get_articles():
    """Busca artigos disponíveis"""
    sql = "SELECT id, title FROM articles ORDER BY title LIMIT 100;"
    success, stdout, _ = run_sql(sql)

    articles = []
    if success:
        lines = stdout.strip().split('\n')[2:-1]  # Skip header and footer
        for line in lines:
            parts = line.split('|')
            if len(parts) >= 2:
                articles.append({
                    'id': parts[0].strip(),
                    'title': parts[1].strip()
                })
    return articles

def enrich_item_with_ai(item, group_context, articles):
    """
    Usa Claude para gerar conteúdo científico para um item
    """

    articles_text = "\n".join([f"- {a['id']}: {a['title']}" for a in articles[:30]])

    prompt = f"""Você é um especialista em Medicina Funcional Integrativa.

TAREFA: Criar conteúdo científico de alta qualidade para um item de avaliação clínica.

ITEM: {item['name']}
GRUPO: {group_context['group']}
SUBGRUPO: {item['subgroup']}

CONTEXTO DO GRUPO:
{group_context.get('description', '')}

Você deve gerar APENAS um JSON válido com esta estrutura exata:

{{
  "clinical_relevance": "Texto técnico para médicos (150-300 palavras). Mecanismos fisiopatológicos, evidências científicas, correlações clínicas. Base em medicina funcional integrativa.",
  "patient_explanation": "Texto acessível para pacientes (100-200 palavras). Linguagem simples, analogias práticas, empoderamento.",
  "conduct": "Orientações práticas para conduta clínica (80-150 palavras). Questões de anamnese, exames sugeridos, estratégias de intervenção.",
  "article_ids": ["id1", "id2", "id3"]
}}

ARTIGOS DISPONÍVEIS (selecione 2-5 relevantes):
{articles_text}

IMPORTANTE:
- Textos em português brasileiro
- Base científica sólida
- Sem emojis
- Retorne APENAS o JSON, sem texto adicional
- article_ids: array com IDs dos artigos mais relevantes"""

    try:
        message = client.messages.create(
            model="claude-sonnet-4-20250514",
            max_tokens=4000,
            messages=[{
                "role": "user",
                "content": prompt
            }]
        )

        # Extrair JSON da resposta
        content = message.content[0].text.strip()

        # Tentar parsear JSON
        if content.startswith('```'):
            # Remover markdown code blocks
            content = content.split('```')[1]
            if content.startswith('json'):
                content = content[4:]

        data = json.loads(content.strip())
        return data

    except Exception as e:
        print(f"  ❌ Erro AI para {item['name']}: {e}")
        return None

def update_item_in_db(item_id, content):
    """Atualiza item no banco de dados"""

    # Escapar aspas simples
    clinical = content['clinical_relevance'].replace("'", "''")
    patient = content['patient_explanation'].replace("'", "''")
    conduct = content['conduct'].replace("'", "''")

    # UPDATE do item
    sql_update = f"""
    UPDATE score_items
    SET
      clinical_relevance = '{clinical}',
      patient_explanation = '{patient}',
      conduct = '{conduct}',
      updated_at = NOW()
    WHERE id = '{item_id}';
    """

    success, _, err = run_sql(sql_update)
    if not success:
        print(f"  ❌ Erro UPDATE: {err}")
        return False

    # Linkar artigos
    if 'article_ids' in content and content['article_ids']:
        for article_id in content['article_ids']:
            sql_link = f"""
            INSERT INTO article_score_items (article_id, score_item_id)
            VALUES ('{article_id}', '{item_id}')
            ON CONFLICT DO NOTHING;
            """
            run_sql(sql_link)

    return True

def enrich_batch(batch_file, group_context):
    """
    Enriquece um batch completo de items
    """

    print("="*80)
    print(f"ENRIQUECIMENTO AUTÔNOMO: {group_context['group']}")
    print("="*80)
    print()

    # Carregar items
    with open(batch_file, 'r', encoding='utf-8') as f:
        data = json.load(f)

    items = data['items']
    total = len(items)

    print(f"📦 Total de items: {total}")
    print()

    # Buscar artigos
    print("📚 Buscando artigos disponíveis...")
    articles = get_articles()
    print(f"  ✓ {len(articles)} artigos encontrados")
    print()

    # Processar cada item
    success_count = 0
    failed = []

    for idx, item in enumerate(items, 1):
        print(f"[{idx}/{total}] {item['name']}")

        # Gerar conteúdo com AI
        print("  → Gerando conteúdo com Claude...")
        content = enrich_item_with_ai(item, group_context, articles)

        if not content:
            failed.append(item['name'])
            continue

        # Atualizar no banco
        print("  → Salvando no banco de dados...")
        if update_item_in_db(item['id'], content):
            success_count += 1
            print(f"  ✅ Sucesso! ({success_count}/{total})")
        else:
            failed.append(item['name'])

        print()

    # Resumo
    print("="*80)
    print("RESUMO")
    print("="*80)
    print(f"Total: {total}")
    print(f"Sucesso: {success_count}")
    print(f"Falhas: {len(failed)}")

    if failed:
        print("\nItems que falharam:")
        for name in failed:
            print(f"  - {name}")

    print()
    print("✓ Batch concluído!")

if __name__ == "__main__":
    import sys

    if len(sys.argv) < 2:
        print("Uso: python3 enrich_batch_autonomous.py <batch_file> [group_description]")
        print("\nExemplo:")
        print("  python3 enrich_batch_autonomous.py scripts/enrichment_data/batch3_composicao.json")
        sys.exit(1)

    batch_file = sys.argv[1]

    # Contexto do grupo
    group_context = {
        'group': 'Composição Corporal',
        'description': 'Avaliação histórica e atual de composição corporal, peso, gordura, massa muscular, períodos de mudança.'
    }

    if len(sys.argv) > 2:
        group_context['description'] = sys.argv[2]

    enrich_batch(batch_file, group_context)
