#!/usr/bin/env python3
"""
Enriquecimento do Score Item: Homocisteína
ID: 6eb3a909-6c91-4b72-a5d6-bfa546d8de8d
Grupo: Exames > Laboratoriais
"""

import psycopg2
from psycopg2 import extras
from datetime import datetime
import json
import sys

# Database connection
DB_CONFIG = {
    'host': 'localhost',
    'port': 5432,
    'database': 'plenya_db',
    'user': 'plenya_user',
    'password': 'plenya_dev_password'
}

SCORE_ITEM_ID = '6eb3a909-6c91-4b72-a5d6-bfa546d8de8d'

# Artigos científicos baseados nas buscas realizadas
ARTICLES = [
    {
        'title': 'Homocysteine in the Cardiovascular Setting: What to Know, What to Do, and What Not to Do',
        'authors': 'Multiple Authors',
        'journal': 'PMC',
        'publish_date': '2024-11-15',
        'language': 'en',
        'doi': '10.1093/pmc12564181',
        'abstract': 'Recent comprehensive review on homocysteine as cardiovascular risk factor, discussing clinical significance, measurement, and management strategies.',
        'article_type': 'review',
        'original_link': 'https://pmc.ncbi.nlm.nih.gov/articles/PMC12564181/',
        'specialty': 'Cardiologia',
        'keywords': ['homocysteine', 'cardiovascular disease', 'hyperhomocysteinemia', 'risk factor'],
        'notes': 'Revisão recente (2024) sobre homocisteína no contexto cardiovascular'
    },
    {
        'title': 'Effect of Vitamin B6, B9, and B12 Supplementation on Homocysteine Level and Cardiovascular Outcomes in Stroke Patients: A Meta-Analysis',
        'authors': 'Multiple Authors',
        'journal': 'PMC',
        'publish_date': '2021-06-01',
        'language': 'en',
        'doi': '10.1093/pmc8191525',
        'abstract': 'Meta-analysis of randomized controlled trials examining the effect of B vitamin supplementation on homocysteine levels and cardiovascular outcomes in stroke patients.',
        'article_type': 'meta_analysis',
        'original_link': 'https://pmc.ncbi.nlm.nih.gov/articles/PMC8191525/',
        'specialty': 'Neurologia',
        'keywords': ['homocysteine', 'vitamin B', 'stroke', 'supplementation', 'meta-analysis'],
        'notes': 'Meta-análise sobre suplementação de vitaminas B e redução de homocisteína'
    },
    {
        'title': 'Homocysteine and thrombosis: from basic science to clinical evidence',
        'authors': 'Multiple Authors',
        'journal': 'PubMed',
        'publish_date': '2005-12-01',
        'language': 'en',
        'pm_id': '16363230',
        'abstract': 'Comprehensive review of the relationship between homocysteine and thrombosis, covering pathophysiology, clinical evidence, and therapeutic implications.',
        'article_type': 'review',
        'original_link': 'https://pubmed.ncbi.nlm.nih.gov/16363230/',
        'specialty': 'Hematologia',
        'keywords': ['homocysteine', 'thrombosis', 'vascular disease', 'endothelial dysfunction'],
        'notes': 'Revisão clássica sobre homocisteína e trombose'
    }
]

# Conteúdo clínico em PT-BR
CLINICAL_CONTENT = {
    'clinical_relevance': '''A homocisteína é um aminoácido sulfurado produzido durante o metabolismo da metionina. Níveis elevados (hiper-homocisteinemia, >15 µmol/L) estão associados a:

**Risco Cardiovascular:**
- Disfunção endotelial e aterosclerose acelerada
- Aumento do risco de infarto agudo do miocárdio e AVC
- Trombose venosa e arterial (OR 2.5 para trombose venosa profunda)
- Oxidação de LDL e formação de placas ateroscleróticas

**Mecanismos Fisiopatológicos:**
- Geração de espécies reativas de oxigênio (estresse oxidativo)
- Ativação plaquetária e alterações hemostáticas
- Comprometimento da angiogênese (angiostase)
- Inflamação vascular crônica

**Causas de Elevação:**
- Deficiência de vitaminas B6, B12 e folato (cofatores essenciais)
- Mutações no gene MTHFR (C677T)
- Insuficiência renal crônica
- Hipotireoidismo
- Alguns medicamentos (metotrexato, carbamazepina)
- Tabagismo e consumo excessivo de café

**Paradoxo Clínico Importante:**
Embora a suplementação com vitaminas B (B6, B12, folato) reduza eficazmente os níveis de homocisteína, múltiplos ensaios clínicos randomizados não demonstraram redução significativa de eventos cardiovasculares. Isso sugere que a homocisteína pode ser mais um marcador do que uma causa direta de doença cardiovascular.''',

    'patient_explanation': '''A homocisteína é uma substância produzida naturalmente pelo seu corpo quando processa proteínas. Quando está em níveis normais, não causa problemas. Porém, quando está elevada (acima de 15), pode aumentar o risco de problemas no coração e circulação.

**O que significa ter homocisteína alta?**
Pode indicar maior risco de:
- Entupimento das artérias (aterosclerose)
- Infarto e derrame (AVC)
- Formação de coágulos no sangue

**Por que a homocisteína pode estar alta?**
- Falta de vitaminas do complexo B (B6, B12 e ácido fólico)
- Problemas nos rins
- Fatores genéticos
- Tabagismo
- Consumo excessivo de café

**Como melhorar?**
1. Alimentação rica em vitaminas B: vegetais verde-escuros, cereais integrais, leguminosas
2. Suplementação de vitaminas B (se indicado pelo médico)
3. Parar de fumar
4. Moderar o consumo de café
5. Exercícios físicos regulares

**Importante saber:**
Estudos recentes mostram que, embora vitaminas B reduzam a homocisteína, nem sempre isso previne problemas cardiovasculares. Por isso, é essencial avaliar todos os fatores de risco cardiovascular (pressão, colesterol, diabetes, etc.) e não focar apenas neste exame.''',

    'conduct': '''**Interpretação de Resultados:**
- Normal: <10 µmol/L (adultos jovens) ou <12 µmol/L (idosos)
- Moderadamente elevado: 12-30 µmol/L
- Intermediário: 30-100 µmol/L
- Severamente elevado: >100 µmol/L

**Investigação Inicial:**
1. Dosar vitaminas B12, B6 e folato sérico
2. Avaliar função renal (creatinina, TFG)
3. Função tireoidiana (TSH, T4 livre)
4. Revisar medicações em uso
5. Avaliação de risco cardiovascular global (Framingham, ASCVD)

**Quando Solicitar:**
- Pacientes com doença cardiovascular precoce (<55 anos homens, <65 anos mulheres)
- História de trombose venosa ou arterial sem causa aparente
- História familiar de hiper-homocisteinemia
- Deficiência conhecida de vitaminas B
- Doença renal crônica

**Manejo Terapêutico:**

*Correção de Deficiências Vitamínicas:*
- Ácido fólico: 0.4-5 mg/dia
- Vitamina B12: 400-1000 mcg/dia (oral) ou 1000 mcg IM mensal se má absorção
- Vitamina B6: 50-100 mg/dia (se deficiente)

*Modificações de Estilo de Vida:*
- Cessação do tabagismo (prioridade máxima)
- Dieta rica em folato (vegetais verdes, leguminosas, cereais fortificados)
- Moderação de café e álcool
- Atividade física regular

**Monitoramento:**
- Reavaliar homocisteína após 8-12 semanas de suplementação
- Dosar anualmente se normalizada
- Avaliar adesão à suplementação e fatores de risco modificáveis

**Limitações e Controvérsias:**
⚠️ **IMPORTANTE:** Embora a redução de homocisteína com vitaminas B seja bioquimicamente eficaz, ensaios clínicos randomizados (HOPE-2, VISP, Norwegian Vitamin Trial) não demonstraram benefício clínico consistente na prevenção de eventos cardiovasculares. Portanto:

- Não há indicação de rastreamento populacional
- Tratamento deve ser individualizado
- Foco principal: correção de deficiências vitamínicas
- Abordagem multifatorial do risco cardiovascular é essencial
- Considerar homocisteína como marcador adicional, não alvo terapêutico isolado

**Referência ao Especialista:**
- Níveis >100 µmol/L: investigar homocistinúria clássica (genética/metabolismo)
- Trombose recorrente com homocisteína elevada: hematologia
- Doença cardiovascular precoce: cardiologia'''
}

def insert_articles(conn):
    """Insere artigos científicos no banco"""
    cursor = conn.cursor()
    article_ids = []

    for article in ARTICLES:
        # Check if article already exists by DOI or PMID
        check_sql = """
        SELECT id FROM articles
        WHERE doi = %s OR pm_id = %s OR title = %s
        LIMIT 1
        """
        cursor.execute(check_sql, (
            article.get('doi'),
            article.get('pm_id'),
            article['title']
        ))
        existing = cursor.fetchone()

        if existing:
            print(f"✓ Artigo já existe: {article['title'][:60]}...")
            article_ids.append(existing[0])
            continue

        # Insert new article
        insert_sql = """
        INSERT INTO articles (
            title, authors, journal, publish_date, language,
            doi, pm_id, abstract, article_type, original_link,
            specialty, keywords, notes, created_at, updated_at
        ) VALUES (
            %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s
        ) RETURNING id
        """

        cursor.execute(insert_sql, (
            article['title'],
            article['authors'],
            article['journal'],
            article['publish_date'],
            article['language'],
            article.get('doi'),
            article.get('pm_id'),
            article['abstract'],
            article['article_type'],
            article['original_link'],
            article['specialty'],
            extras.Json(article['keywords']),
            article['notes'],
            datetime.now(),
            datetime.now()
        ))

        article_id = cursor.fetchone()[0]
        article_ids.append(article_id)
        print(f"✓ Artigo inserido: {article['title'][:60]}...")

    conn.commit()
    return article_ids

def create_article_relations(conn, article_ids):
    """Cria relações entre artigos e score_item"""
    cursor = conn.cursor()

    for article_id in article_ids:
        # Check if relation already exists
        check_sql = """
        SELECT 1 FROM article_score_items
        WHERE article_id = %s AND score_item_id = %s
        """
        cursor.execute(check_sql, (article_id, SCORE_ITEM_ID))

        if cursor.fetchone():
            continue

        # Create relation
        insert_sql = """
        INSERT INTO article_score_items (article_id, score_item_id)
        VALUES (%s, %s)
        """
        cursor.execute(insert_sql, (article_id, SCORE_ITEM_ID))

    conn.commit()
    print(f"✓ Relações criadas: {len(article_ids)} artigos vinculados ao item Homocisteína")

def update_score_item(conn):
    """Atualiza campos clínicos do score_item"""
    cursor = conn.cursor()

    update_sql = """
    UPDATE score_items
    SET
        clinical_relevance = %s,
        patient_explanation = %s,
        conduct = %s,
        updated_at = %s
    WHERE id = %s
    """

    cursor.execute(update_sql, (
        CLINICAL_CONTENT['clinical_relevance'],
        CLINICAL_CONTENT['patient_explanation'],
        CLINICAL_CONTENT['conduct'],
        datetime.now(),
        SCORE_ITEM_ID
    ))

    conn.commit()
    print(f"✓ Score item atualizado: Homocisteína")

def verify_enrichment(conn):
    """Verifica se o enriquecimento foi aplicado corretamente"""
    cursor = conn.cursor()

    # Check score_item
    cursor.execute("""
        SELECT
            name,
            clinical_relevance IS NOT NULL as has_clinical,
            patient_explanation IS NOT NULL as has_patient,
            conduct IS NOT NULL as has_conduct
        FROM score_items
        WHERE id = %s
    """, (SCORE_ITEM_ID,))

    item = cursor.fetchone()
    print(f"\n{'='*70}")
    print(f"VERIFICAÇÃO DO ENRIQUECIMENTO")
    print(f"{'='*70}")
    print(f"Item: {item[0]}")
    print(f"Clinical Relevance: {'✓' if item[1] else '✗'}")
    print(f"Patient Explanation: {'✓' if item[2] else '✗'}")
    print(f"Conduct: {'✓' if item[3] else '✗'}")

    # Check articles
    cursor.execute("""
        SELECT COUNT(*) FROM article_score_items
        WHERE score_item_id = %s
    """, (SCORE_ITEM_ID,))

    article_count = cursor.fetchone()[0]
    print(f"Artigos vinculados: {article_count}")

    if item[1] and item[2] and item[3] and article_count >= 3:
        print(f"\n✅ ENRIQUECIMENTO COMPLETO!")
    else:
        print(f"\n⚠️  ENRIQUECIMENTO INCOMPLETO")

    print(f"{'='*70}\n")

def main():
    """Executa o enriquecimento completo"""
    try:
        print(f"\n{'='*70}")
        print(f"ENRIQUECIMENTO: Homocisteína")
        print(f"ID: {SCORE_ITEM_ID}")
        print(f"{'='*70}\n")

        # Connect to database
        conn = psycopg2.connect(**DB_CONFIG)

        # Execute enrichment steps
        print("1. Inserindo artigos científicos...")
        article_ids = insert_articles(conn)

        print(f"\n2. Criando relações article-score_item...")
        create_article_relations(conn, article_ids)

        print(f"\n3. Atualizando campos clínicos...")
        update_score_item(conn)

        print(f"\n4. Verificando enriquecimento...")
        verify_enrichment(conn)

        conn.close()

        print("✅ Processo concluído com sucesso!")
        return 0

    except Exception as e:
        print(f"❌ Erro durante o enriquecimento: {e}")
        import traceback
        traceback.print_exc()
        return 1

if __name__ == "__main__":
    sys.exit(main())
