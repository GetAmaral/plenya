#!/usr/bin/env python3
"""
Enriquecimento do Score Item: PSA Total
ID: 2f01d97b-28ef-4016-96ad-2fa95ae09f99
Categoria: Exames > Laboratoriais
"""

import os
import psycopg2
from psycopg2.extras import RealDictCursor
import json
from datetime import datetime

# Database connection
DB_CONFIG = {
    'host': os.getenv('DB_HOST', 'localhost'),
    'port': os.getenv('DB_PORT', '5432'),
    'database': os.getenv('DB_NAME', 'plenya_db'),
    'user': os.getenv('DB_USER', 'plenya_user'),
    'password': os.getenv('DB_PASSWORD', 'senha_segura')
}

# Item ID
ITEM_ID = '2f01d97b-28ef-4016-96ad-2fa95ae09f99'

# Clinical content in PT-BR
CLINICAL_CONTENT = {
    'clinical_relevance': '''O PSA (Antígeno Prostático Específico) é uma proteína produzida pela próstata e seu nível sanguíneo é utilizado como marcador para rastreamento e monitoramento de condições prostáticas. O teste PSA Total mede a quantidade total dessa proteína no sangue.

INTERPRETAÇÃO DOS VALORES:
• < 1,0 ng/mL: Considerado benigno, baixo risco
• 1,0-3,0 ng/mL: Faixa geralmente segura na maioria dos homens
• 3,0-4,0 ng/mL: Zona de atenção, pode necessitar avaliação adicional
• 4,0-10,0 ng/mL: Valores elevados, indicam necessidade de investigação adicional
• > 10,0 ng/mL: Fortemente sugestivo de patologia prostática
• > 20,0 ng/mL: Mais indicativo de câncer de próstata
• > 50,0 ng/mL: Forte indicador de câncer de próstata

VALORES POR IDADE:
• Homens < 60 anos: PSA > 2,5 ng/mL requer avaliação
• Homens ≥ 60 anos: PSA > 4,0 ng/mL requer avaliação

CAUSAS DE ELEVAÇÃO DO PSA:
1. Câncer de próstata (principal preocupação)
2. Hiperplasia prostática benigna (HPB)
3. Prostatite (inflamação/infecção)
4. Manipulação prostática (toque retal, biópsia recente)
5. Exercício vigoroso (ciclismo)
6. Ejaculação recente (últimas 48h)

LIMITAÇÕES DO TESTE:
• PSA elevado não significa necessariamente câncer
• PSA normal não exclui totalmente a possibilidade de câncer
• Deve ser interpretado em conjunto com outros fatores: idade, raça, história familiar, toque retal, PSA livre/total''',

    'patient_explanation': '''O PSA Total é um exame de sangue que mede uma proteína produzida pela sua próstata. Ele é usado principalmente para detectar problemas na próstata, incluindo câncer.

O QUE SIGNIFICA O RESULTADO?
• Valores normais geralmente ficam abaixo de 4,0 ng/mL
• Valores entre 4,0 e 10,0 podem indicar problemas e geralmente requerem mais testes
• Valores acima de 10,0 são mais preocupantes e exigem investigação

IMPORTANTE SABER:
• PSA elevado NÃO significa automaticamente que você tem câncer
• Várias condições benignas podem elevar o PSA (aumento natural da próstata com idade, infecções)
• O resultado deve ser analisado junto com sua idade, sintomas e outros exames

FATORES QUE PODEM ALTERAR O EXAME:
• Evite ejaculação 48 horas antes do exame
• Evite exercícios intensos (especialmente ciclismo) antes do exame
• Informe ao médico sobre medicamentos para próstata (finasterida, dutasterida)
• Não faça após massagem prostática ou toque retal recente

QUANDO FAZER O EXAME?
As diretrizes recomendam conversar com seu médico sobre rastreamento a partir dos:
• 45 anos para a maioria dos homens
• 40 anos se você é negro ou tem histórico familiar de câncer de próstata

A decisão de fazer o PSA deve ser individualizada, considerando benefícios e riscos.''',

    'conduct': '''ABORDAGEM CONFORME DIRETRIZES USPSTF (2018):

HOMENS 55-69 ANOS:
• Decisão individualizada após discussão sobre benefícios/riscos (Recomendação C)
• Intervalo de rastreamento: geralmente a cada 2-4 anos se optado pelo rastreamento
• Considerar fatores de risco: raça negra, história familiar

HOMENS ≥ 70 ANOS:
• Rastreamento de rotina NÃO recomendado (Recomendação D)
• Exceções: expectativa de vida > 10 anos e risco elevado

HOMENS < 55 ANOS:
• Rastreamento não recomendado de rotina
• Avaliar caso a caso se alto risco (história familiar forte, sintomas)

CONDUTAS CONFORME RESULTADO DO PSA:

PSA < 1,0 ng/mL:
• Risco muito baixo
• Repetir em 2-4 anos se rastreamento optado
• Considerar intervalos maiores ou interromper rastreamento

PSA 1,0-3,0 ng/mL:
• Risco baixo-moderado
• Repetir anualmente ou a cada 2 anos
• Avaliar velocidade de aumento (PSA velocity)
• Considerar PSA livre/total se disponível

PSA 3,0-10,0 ng/mL:
• Risco moderado-alto, requer avaliação
• Solicitar PSA livre/total (% PSA livre < 10% = alto risco)
• Exame de toque retal
• Considerar ressonância multiparamétrica de próstata
• Avaliar PSA density (PSA/volume prostático)
• Discutir biópsia se indicadores de risco

PSA > 10,0 ng/mL:
• Alto risco de câncer de próstata
• Exame de toque retal obrigatório
• Ressonância multiparamétrica de próstata
• Encaminhamento para urologista
• Biópsia geralmente indicada

MONITORAMENTO:
• Velocidade do PSA (PSA velocity): aumento > 0,75 ng/mL/ano = preocupante
• Tempo de duplicação do PSA (PSA doubling time)
• PSA livre/total (% PSA livre)

REFERÊNCIAS CIENTÍFICAS:
1. USPSTF Recommendation 2018: https://www.uspreventiveservicestaskforce.org/uspstf/recommendation/prostate-cancer-screening
2. NCI PSA Fact Sheet: https://www.cancer.gov/types/prostate/psa-fact-sheet
3. StatPearls PSA Review: https://www.ncbi.nlm.nih.gov/books/NBK557495/'''
}

# Articles to link
ARTICLES = [
    {
        'title': 'Screening for Prostate Cancer: US Preventive Services Task Force Recommendation Statement',
        'authors': 'US Preventive Services Task Force',
        'journal': 'JAMA',
        'publish_date': '2018-05-08',
        'pm_id': '29801017',
        'doi': '10.1001/jama.2018.3710',
        'original_link': 'https://www.uspreventiveservicestaskforce.org/uspstf/recommendation/prostate-cancer-screening',
        'article_type': 'clinical_trial',
        'keywords': ['PSA', 'prostate cancer', 'screening', 'USPSTF', 'guidelines']
    },
    {
        'title': 'Prostate-Specific Antigen (PSA) - Clinical Interpretation and Use',
        'authors': 'National Cancer Institute',
        'journal': 'NCI Fact Sheet',
        'publish_date': '2024-01-15',
        'original_link': 'https://www.cancer.gov/types/prostate/psa-fact-sheet',
        'article_type': 'review',
        'keywords': ['PSA', 'interpretation', 'clinical use', 'biomarker']
    },
    {
        'title': 'Prostate-Specific Antigen: StatPearls Review',
        'authors': 'Leslie SW, Soon-Sutton TL, Sajjad H, Siref LE',
        'journal': 'StatPearls Publishing',
        'publish_date': '2024-06-01',
        'pm_id': '32491364',
        'original_link': 'https://www.ncbi.nlm.nih.gov/books/NBK557495/',
        'article_type': 'review',
        'keywords': ['PSA', 'physiology', 'clinical significance', 'testing']
    }
]

def connect_db():
    """Connect to PostgreSQL database"""
    try:
        conn = psycopg2.connect(**DB_CONFIG)
        return conn
    except Exception as e:
        print(f"❌ Erro ao conectar ao banco: {e}")
        raise

def insert_or_get_article(cursor, article_data):
    """Insert article if not exists and return its ID"""

    # Check if article exists by title or DOI/PM_ID
    query_check = """
        SELECT id FROM articles
        WHERE title = %s
        OR (doi = %s AND doi IS NOT NULL)
        OR (pm_id = %s AND pm_id IS NOT NULL)
        LIMIT 1
    """
    cursor.execute(query_check, (
        article_data['title'],
        article_data.get('doi'),
        article_data.get('pm_id')
    ))

    result = cursor.fetchone()
    if result:
        print(f"  ✓ Artigo já existe: {article_data['title'][:60]}...")
        return result['id']

    # Insert new article
    query_insert = """
        INSERT INTO articles (
            title, authors, journal, publish_date,
            doi, pm_id, original_link, article_type, keywords,
            created_at, updated_at
        ) VALUES (
            %s, %s, %s, %s, %s, %s, %s, %s, %s, NOW(), NOW()
        ) RETURNING id
    """

    cursor.execute(query_insert, (
        article_data['title'],
        article_data['authors'],
        article_data['journal'],
        article_data['publish_date'],
        article_data.get('doi'),
        article_data.get('pm_id'),
        article_data['original_link'],
        article_data['article_type'],
        json.dumps(article_data['keywords'])
    ))

    article_id = cursor.fetchone()['id']
    print(f"  ✓ Artigo inserido: {article_data['title'][:60]}...")
    return article_id

def link_article_to_item(cursor, article_id, item_id):
    """Link article to score item (avoid duplicates)"""

    # Check if relation already exists
    query_check = """
        SELECT 1 FROM article_score_items
        WHERE score_item_id = %s AND article_id = %s
    """
    cursor.execute(query_check, (item_id, article_id))

    if cursor.fetchone():
        return False  # Already linked

    # Create link
    query_insert = """
        INSERT INTO article_score_items (score_item_id, article_id)
        VALUES (%s, %s)
    """
    cursor.execute(query_insert, (item_id, article_id))
    return True  # New link created

def update_score_item(cursor, item_id, clinical_content):
    """Update score item with clinical content"""

    query = """
        UPDATE score_items
        SET
            clinical_relevance = %s,
            patient_explanation = %s,
            conduct = %s,
            last_review = NOW(),
            updated_at = NOW()
        WHERE id = %s
        RETURNING name, unit
    """

    cursor.execute(query, (
        clinical_content['clinical_relevance'],
        clinical_content['patient_explanation'],
        clinical_content['conduct'],
        item_id
    ))

    return cursor.fetchone()

def main():
    print("=" * 80)
    print("ENRIQUECIMENTO: PSA Total")
    print("=" * 80)
    print(f"Item ID: {ITEM_ID}")
    print(f"Timestamp: {datetime.now().isoformat()}")
    print()

    conn = connect_db()
    cursor = conn.cursor(cursor_factory=RealDictCursor)

    try:
        # Step 1: Insert articles
        print("📚 INSERINDO ARTIGOS CIENTÍFICOS...")
        print("-" * 80)
        article_ids = []
        for article in ARTICLES:
            article_id = insert_or_get_article(cursor, article)
            article_ids.append(article_id)
        print()

        # Step 2: Link articles to item
        print("🔗 VINCULANDO ARTIGOS AO ITEM...")
        print("-" * 80)
        new_links = 0
        for article_id in article_ids:
            if link_article_to_item(cursor, article_id, ITEM_ID):
                new_links += 1
                print(f"  ✓ Artigo {article_id} vinculado")
            else:
                print(f"  → Artigo {article_id} já vinculado")
        print(f"\n  Total: {new_links} novos vínculos criados")
        print()

        # Step 3: Update clinical content
        print("📝 ATUALIZANDO CONTEÚDO CLÍNICO...")
        print("-" * 80)
        result = update_score_item(cursor, ITEM_ID, CLINICAL_CONTENT)
        print(f"  ✓ Item atualizado: {result['name']}")
        if result['unit']:
            print(f"  ✓ Unidade: {result['unit']}")
        print()

        # Step 4: Verify update
        print("✅ VERIFICANDO ATUALIZAÇÃO...")
        print("-" * 80)
        cursor.execute("""
            SELECT
                si.name,
                si.unit,
                si.last_review,
                LENGTH(si.clinical_relevance) as relevance_len,
                LENGTH(si.patient_explanation) as explanation_len,
                LENGTH(si.conduct) as conduct_len,
                COUNT(asi.article_id) as article_count
            FROM score_items si
            LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
            WHERE si.id = %s
            GROUP BY si.id, si.name, si.unit, si.last_review,
                     si.clinical_relevance, si.patient_explanation, si.conduct
        """, (ITEM_ID,))

        verification = cursor.fetchone()
        print(f"  Item: {verification['name']}")
        print(f"  Unidade: {verification['unit']}")
        print(f"  Última revisão: {verification['last_review']}")
        print(f"  Clinical Relevance: {verification['relevance_len']} caracteres")
        print(f"  Patient Explanation: {verification['explanation_len']} caracteres")
        print(f"  Conduct: {verification['conduct_len']} caracteres")
        print(f"  Artigos vinculados: {verification['article_count']}")
        print()

        # Commit transaction
        conn.commit()

        print("=" * 80)
        print("✅ ENRIQUECIMENTO CONCLUÍDO COM SUCESSO!")
        print("=" * 80)
        print()
        print("FONTES CIENTÍFICAS:")
        print("-" * 80)
        for article in ARTICLES:
            print(f"• {article['title']}")
            print(f"  {article['journal']}, {article['publish_date']}")
            print(f"  {article['original_link']}")
            print()

    except Exception as e:
        conn.rollback()
        print(f"\n❌ ERRO: {e}")
        raise
    finally:
        cursor.close()
        conn.close()

if __name__ == '__main__':
    main()
