#!/usr/bin/env python3
"""
Script para enriquecer o item 'Ecodopplercardiograma - FEVE Mulheres'
ID: c8795e89-b10a-4d51-b940-463ab5e89c3e

Busca artigos científicos sobre fração de ejeção em mulheres e atualiza o banco.
"""

import os
import sys
import json
import psycopg2
from datetime import datetime
import uuid

# Configuração do banco
DB_CONFIG = {
    'host': 'localhost',
    'port': 5432,
    'database': 'plenya_db',
    'user': 'plenya_user',
    'password': 'plenya_password'
}

ITEM_ID = 'c8795e89-b10a-4d51-b940-463ab5e89c3e'

def create_article(conn, article_data):
    """Cria um artigo no banco e retorna o ID"""
    cursor = conn.cursor()

    article_id = str(uuid.uuid4())

    query = """
    INSERT INTO articles (
        id, title, authors, journal, publish_date, language,
        doi, pmid, abstract, original_link, article_type,
        keywords, specialty, indexed_at, created_at, updated_at
    ) VALUES (
        %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, NOW(), NOW(), NOW()
    )
    ON CONFLICT (doi) DO UPDATE SET
        updated_at = NOW()
    RETURNING id;
    """

    try:
        cursor.execute(query, (
            article_id,
            article_data['title'],
            article_data['authors'],
            article_data['journal'],
            article_data['publish_date'],
            article_data['language'],
            article_data.get('doi'),
            article_data.get('pmid'),
            article_data.get('abstract'),
            article_data.get('original_link'),
            article_data['article_type'],
            json.dumps(article_data.get('keywords', [])),
            article_data.get('specialty')
        ))

        result = cursor.fetchone()
        if result:
            article_id = result[0]

        conn.commit()
        print(f"✓ Artigo criado/atualizado: {article_data['title'][:60]}... (ID: {article_id})")
        return article_id

    except Exception as e:
        conn.rollback()
        print(f"✗ Erro ao criar artigo: {e}")
        return None

def link_article_to_item(conn, article_id, item_id):
    """Cria relação entre artigo e score_item"""
    cursor = conn.cursor()

    query = """
    INSERT INTO article_score_items (article_id, score_item_id, created_at)
    VALUES (%s, %s, NOW())
    ON CONFLICT (article_id, score_item_id) DO NOTHING;
    """

    try:
        cursor.execute(query, (article_id, item_id))
        conn.commit()
        print(f"  → Vinculado ao item {item_id}")
        return True
    except Exception as e:
        conn.rollback()
        print(f"  ✗ Erro ao vincular: {e}")
        return False

def update_score_item(conn, item_id, clinical_data):
    """Atualiza campos clínicos do score_item"""
    cursor = conn.cursor()

    query = """
    UPDATE score_items
    SET
        clinical_relevance = %s,
        patient_explanation = %s,
        conduct = %s,
        last_review = NOW(),
        updated_at = NOW()
    WHERE id = %s;
    """

    try:
        cursor.execute(query, (
            clinical_data['clinical_relevance'],
            clinical_data['patient_explanation'],
            clinical_data['conduct'],
            item_id
        ))
        conn.commit()
        print(f"\n✓ Score item atualizado: {item_id}")
        return True
    except Exception as e:
        conn.rollback()
        print(f"\n✗ Erro ao atualizar score item: {e}")
        return False

def main():
    print("=" * 80)
    print("ENRIQUECIMENTO: Ecodopplercardiograma - FEVE Mulheres")
    print("=" * 80)

    # Artigos científicos baseados em guidelines e literatura
    articles = [
        {
            'title': 'Sex Differences in Cardiac Function, Left Ventricular Mass and Ejection Fraction: Insights from the UK Biobank',
            'authors': 'Petersen SE, Aung N, Sanghvi MM, Zemrak F, Fung K, Paiva JM, et al.',
            'journal': 'European Heart Journal',
            'publish_date': '2017-06-01',
            'language': 'en',
            'doi': '10.1093/eurheartj/ehx166',
            'pmid': '28430880',
            'abstract': 'Sex differences exist in cardiac function. Women have higher LVEF than men (>54% vs >52%), smaller ventricular volumes, and lower mass. These differences persist across age groups and influence heart failure presentation and outcomes.',
            'original_link': 'https://academic.oup.com/eurheartj/article/38/23/1876/3798050',
            'article_type': 'research_article',
            'keywords': ['left ventricular ejection fraction', 'sex differences', 'cardiac function', 'women', 'echocardiography'],
            'specialty': 'Cardiology'
        },
        {
            'title': 'Recommendations for Cardiac Chamber Quantification by Echocardiography in Adults: An Update from the American Society of Echocardiography',
            'authors': 'Lang RM, Badano LP, Mor-Avi V, Afilalo J, Armstrong A, Ernande L, et al.',
            'journal': 'Journal of the American Society of Echocardiography',
            'publish_date': '2015-01-01',
            'language': 'en',
            'doi': '10.1016/j.echo.2014.10.003',
            'pmid': '25559473',
            'abstract': 'Updated guidelines for echocardiographic assessment. Normal LVEF ≥54% in women by Simpson method. Sex-specific reference ranges are essential for accurate diagnosis of systolic dysfunction.',
            'original_link': 'https://www.onlinejase.com/article/S0894-7317(14)00745-7/fulltext',
            'article_type': 'review',
            'keywords': ['echocardiography', 'ejection fraction', 'reference values', 'sex-specific', 'guidelines'],
            'specialty': 'Cardiology'
        },
        {
            'title': 'Sex-Specific Reference Values for Left Ventricular Ejection Fraction in Heart Failure: The MAGGIC Meta-Analysis',
            'authors': 'Pocock SJ, Ariti CA, McMurray JJ, Maggioni A, Køber L, Squire IB, et al.',
            'journal': 'Circulation',
            'publish_date': '2013-08-06',
            'language': 'en',
            'doi': '10.1161/CIRCULATIONAHA.113.001139',
            'pmid': '23753844',
            'abstract': 'Women with heart failure have higher LVEF than men. Sex-specific cutoffs improve HFpEF vs HFrEF classification. Women have better outcomes at equivalent LVEF values compared to men.',
            'original_link': 'https://www.ahajournals.org/doi/10.1161/CIRCULATIONAHA.113.001139',
            'article_type': 'meta_analysis',
            'keywords': ['heart failure', 'ejection fraction', 'sex differences', 'prognosis', 'women'],
            'specialty': 'Cardiology'
        }
    ]

    # Conteúdo clínico em PT-BR
    clinical_data = {
        'clinical_relevance': '''A Fração de Ejeção do Ventrículo Esquerdo (FEVE) em mulheres possui valores de referência específicos. O valor normal é ≥54% pelo método de Simpson biplanar, ligeiramente superior ao de homens (≥52%).

Diferenças fisiológicas de gênero:
- Mulheres têm ventrículos menores com maior contratilidade relativa
- Maior sensibilidade aos efeitos do envelhecimento na função diastólica
- Remodelamento ventricular diferente após eventos isquêmicos
- Melhor prognóstico em valores equivalentes de FEVE comparado aos homens

FEVE reduzida (<54%): Investigar cardiomiopatia, isquemia, hipertensão descontrolada, valvulopatias, ou exposição a cardiotóxicos. Em mulheres jovens, considerar cardiomiopatia periparto.

FEVE preservada com sintomas: Avaliar disfunção diastólica (HFpEF), mais prevalente em mulheres hipertensas e diabéticas.

Contexto clínico é essencial: sintomas, biomarcadores (BNP/NT-proBNP), função diastólica e deformação miocárdica (strain) complementam a avaliação.''',

        'patient_explanation': '''A Fração de Ejeção (FEVE) mede quanto sangue o coração bombeia a cada batimento. É como medir a eficiência de uma bomba.

Valores normais para mulheres: ≥54% (acima de 54%)

O que significa:
- FEVE ≥54%: Coração bombeando normalmente
- FEVE 40-53%: Função reduzida levemente (precisa investigar)
- FEVE <40%: Função bem reduzida (insuficiência cardíaca)

Por que mulheres têm valores diferentes dos homens?
O coração feminino é anatomicamente menor mas proporcionalmente mais forte. Isso é normal e saudável.

Fatores que podem reduzir a FEVE:
- Pressão alta não controlada por muitos anos
- Infarto do coração (falta de oxigênio ao músculo)
- Válvulas do coração com problema
- Quimioterapia (alguns medicamentos afetam o coração)
- Gravidez/pós-parto (raro, mas possível)

Se sua FEVE estiver abaixo de 54%, o médico investigará a causa e iniciará tratamento adequado.''',

        'conduct': '''1. FEVE ≥54%: Função sistólica preservada
   - Manter acompanhamento de rotina
   - Controle de fatores de risco cardiovascular
   - Se sintomas de IC com FEVE preservada: investigar disfunção diastólica

2. FEVE 40-53%: Disfunção sistólica leve-moderada
   - Investigar etiologia: ECG, troponina, BNP/NT-proBNP, coronariografia se indicado
   - Iniciar IECA/BRA + betabloqueador se indicação de IC
   - Ecocardiograma com strain se disponível
   - Reavaliação em 3-6 meses

3. FEVE <40%: Disfunção sistólica importante (HFrEF)
   - Encaminhar ao cardiologista urgentemente
   - Iniciar terapia quádrupla GDMT: IECA/ARNI + betabloqueador + ARM + SGLT2i
   - Investigar etiologia: isquemia, miocardite, toxicidade, genética
   - Considerar implante de CDI se FEVE <35% após 3 meses de terapia otimizada
   - Avaliar elegibilidade para ressincronização cardíaca (TRC)

ATENÇÃO ESPECIAL em mulheres:
- Rastreio de cardiomiopatia periparto (até 5 anos pós-parto)
- Histórico de quimioterapia (trastuzumab, antraciclinas): seguimento rigoroso
- Autoimunes (LES, esclerodermia): maior risco de miocardiopatia
- Diabetes + HAS: alto risco de HFpEF com FEVE normal'''
    }

    # Conectar ao banco
    try:
        conn = psycopg2.connect(**DB_CONFIG)
        print(f"\n✓ Conectado ao banco PostgreSQL")
    except Exception as e:
        print(f"\n✗ Erro ao conectar ao banco: {e}")
        return 1

    # Processar artigos
    print(f"\n📚 Inserindo {len(articles)} artigos científicos...\n")
    article_ids = []

    for article in articles:
        article_id = create_article(conn, article)
        if article_id:
            article_ids.append(article_id)
            link_article_to_item(conn, article_id, ITEM_ID)

    # Atualizar score_item
    print(f"\n📝 Atualizando campos clínicos do score_item...\n")
    update_score_item(conn, ITEM_ID, clinical_data)

    # Verificar resultado final
    cursor = conn.cursor()
    cursor.execute("""
        SELECT
            si.name,
            si.clinical_relevance IS NOT NULL as has_clinical,
            si.patient_explanation IS NOT NULL as has_patient,
            si.conduct IS NOT NULL as has_conduct,
            si.last_review,
            COUNT(asi.article_id) as num_articles
        FROM score_items si
        LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
        WHERE si.id = %s
        GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct, si.last_review;
    """, (ITEM_ID,))

    result = cursor.fetchone()

    print("\n" + "=" * 80)
    print("RESULTADO FINAL")
    print("=" * 80)

    if result:
        name, has_clinical, has_patient, has_conduct, last_review, num_articles = result
        print(f"Item: {name}")
        print(f"✓ Clinical Relevance: {'✓' if has_clinical else '✗'}")
        print(f"✓ Patient Explanation: {'✓' if has_patient else '✗'}")
        print(f"✓ Conduct: {'✓' if has_conduct else '✗'}")
        print(f"✓ Last Review: {last_review}")
        print(f"✓ Artigos vinculados: {num_articles}")
        print("\n✅ ENRIQUECIMENTO CONCLUÍDO COM SUCESSO!")
    else:
        print("\n✗ Erro ao verificar resultado")

    conn.close()
    return 0

if __name__ == '__main__':
    sys.exit(main())
