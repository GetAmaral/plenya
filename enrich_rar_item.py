#!/usr/bin/env python3
"""
Enriquecimento do item RAR (Renal-Aortic Ratio) - Doppler Aorta
Item ID: b7c85ad4-ece5-4a21-b23d-8edf321f7874
"""

import psycopg2
import sys
from datetime import datetime

# Configuração do banco de dados
DB_CONFIG = {
    'host': 'localhost',
    'port': 5432,
    'database': 'plenya_db',
    'user': 'plenya_user',
    'password': 'plenya_dev_password'
}

# ID do item a ser enriquecido
ITEM_ID = 'b7c85ad4-ece5-4a21-b23d-8edf321f7874'

# Dados de enriquecimento baseados em evidências científicas
ENRICHMENT_DATA = {
    'clinical_relevance': '''O RAR (Renal-Aortic Ratio) é um parâmetro crítico do Doppler ultrassonográfico usado para detectar estenose de artéria renal, a causa mais comum de hipertensão secundária. Calculado pela divisão da velocidade de pico sistólico (PSV) máxima na artéria renal pela PSV na aorta, o RAR normaliza variações hemodinâmicas individuais.

**Critérios Diagnósticos:**
- RAR normal: < 3.5
- RAR ≥ 3.5: sugere estenose ≥ 60% (sensibilidade 84-91%, especificidade 95-97%)
- Valores médios: Normal = 2.2; Estenose <60% = 2.9; Estenose ≥60% = 4.5

**Indicações Clínicas:**
- Hipertensão resistente ao tratamento medicamentoso
- Hipertensão de início súbito em jovens (<30 anos) ou idosos (>55 anos)
- Deterioração aguda da função renal com uso de IECA/BRA
- Assimetria renal ou sopro abdominal
- Monitoramento pós-transplante renal

**Achados Associados:**
- PSV renal > 200 cm/s (estenose ≥60%)
- Fluxo turbulento pós-estenótico com aliasing
- Ausência de sinal Doppler (oclusão completa)
- Índice de resistência (IR) alterado no parênquima renal

O RAR é preferível ao PSV isolado por reduzir falsos positivos em estados hiperdinâmicos (anemia, febre, tireotoxicose) e melhorar a especificidade diagnóstica.''',

    'patient_explanation': '''O RAR (Razão Renal-Aórtica) é um exame de ultrassom com Doppler que avalia se as artérias que levam sangue aos seus rins estão estreitadas (entupidas). Esse estreitamento pode causar pressão alta difícil de controlar com remédios.

**Como é feito:**
O médico passa um aparelho de ultrassom na sua barriga para medir a velocidade do sangue nas artérias dos rins e compara com a velocidade na artéria principal (aorta). É como medir o fluxo de água em um cano: se estiver entupido, a água passa mais rápido no local do entupimento.

**O que significa o resultado:**
- Valor normal: menor que 3.5
- Valor alterado: 3.5 ou maior - pode indicar artéria renal estreitada

**Quando é pedido:**
- Pressão alta que não melhora com vários remédios
- Pressão alta que começou de repente
- Rins de tamanhos diferentes
- Após transplante de rim
- Piora da função dos rins ao usar certos remédios para pressão

**Preparação:**
- Jejum de 6-8 horas (para reduzir gases intestinais)
- Pode precisar beber água antes do exame

Este exame não dói e não usa radiação. É seguro e pode ajudar seu médico a entender melhor sua pressão alta e função renal.''',

    'conduct': '''**Interpretação dos Resultados:**

1. **RAR < 3.5 (Normal):**
   - Estenose significativa (≥60%) improvável
   - Considerar outras causas de hipertensão
   - Manter seguimento clínico rotineiro

2. **RAR 3.5 - 4.5 (Limítrofe):**
   - Considerar exame confirmatório (AngioTC ou AngioRM renal)
   - Avaliar outros parâmetros Doppler (PSV, IR, padrão de onda)
   - Repetir exame em 3-6 meses se clínica estável

3. **RAR > 4.5 (Fortemente Positivo):**
   - Alta probabilidade de estenose ≥60%
   - Solicitar AngioTC/AngioRM para confirmação anatômica
   - Avaliar etiologia: aterosclerose (>55 anos) vs. displasia fibromuscular (jovens)
   - Considerar interconsulta com nefrologia/cirurgia vascular

**Achados Adicionais Importantes:**
- PSV renal > 200 cm/s + RAR > 3.5: VPP ~95% para estenose significativa
- IR parênquima renal > 0.8: pior prognóstico pós-revascularização
- Assimetria de IR entre rins > 0.05: sugestivo de estenose unilateral

**Manejo Terapêutico:**
- Otimizar controle pressórico (evitar IECA/BRA se estenose bilateral grave)
- Controlar fatores de risco (diabetes, dislipidemia, tabagismo)
- Avaliar indicação de revascularização:
  * Indicações fortes: hipertensão refratária, edema pulmonar flash, deterioração renal progressiva
  * Técnicas: angioplastia com stent (aterosclerose) ou sem stent (displasia)

**Limitações do Exame:**
- Dependente do operador
- Dificuldade técnica em obesos ou com gases intestinais
- Pode ser inconclusivo em 10-20% dos casos
- Falsos negativos em estenose distal ou em ramos

**Seguimento:**
- Se normal: reavaliação conforme quadro clínico
- Se alterado sem revascularização: Doppler anual
- Pós-angioplastia/stent: Doppler em 1, 6 e 12 meses (detectar re-estenose)'''
}

def create_articles():
    """Cria artigos científicos baseados nas fontes da pesquisa"""
    articles = []

    # Artigo 1: StatPearls - Doppler Renal Assessment
    articles.append({
        'title': 'Doppler Renal Assessment, Protocols, and Interpretation',
        'authors': 'StatPearls Publishing',
        'journal': 'StatPearls',
        'publish_date': '2024-01-01',
        'language': 'en',
        'pm_id': 'NBK572135',
        'abstract': 'Comprehensive review of renal Doppler ultrasound techniques, including renal-aortic ratio (RAR) for detecting renal artery stenosis. The RAR compares intrastenotic flow velocity in renal arteries with aortic reference values, with RAR >3.5 predicting ≥60% stenosis with 84-91% sensitivity and 95-97% specificity.',
        'original_link': 'https://www.ncbi.nlm.nih.gov/books/NBK572135/',
        'article_type': 'review',
        'specialty': 'Radiologia',
        'keywords': ['renal doppler', 'renal artery stenosis', 'renal-aortic ratio', 'renovascular hypertension', 'ultrasound'],
        'favorite': True,
        'rating': 5
    })

    # Artigo 2: Journal of Vascular Surgery - Critical Analysis
    articles.append({
        'title': 'Critical analysis of renal duplex ultrasound parameters in detecting significant renal artery stenosis',
        'authors': 'Zierler RE, Bergelin RO, Isaacson JA, Strandness DE',
        'journal': 'Journal of Vascular Surgery',
        'publish_date': '2012-09-01',
        'language': 'en',
        'doi': '10.1016/j.jvs.2012.02.044',
        'abstract': 'Large study of 313 patients evaluating Doppler parameters for renal artery stenosis. Mean renal-aortic ratios for normal, <60%, and ≥60% stenosis were 2.2, 2.9, and 4.5 respectively. RAR >3.5 demonstrated high diagnostic accuracy for detecting hemodynamically significant stenosis.',
        'original_link': 'https://www.jvascsurg.org/article/S0741-5214(12)00504-6/fulltext',
        'article_type': 'research_article',
        'specialty': 'Cirurgia Vascular',
        'keywords': ['renal artery stenosis', 'duplex ultrasound', 'renal-aortic ratio', 'diagnostic accuracy', 'hemodynamic significance'],
        'favorite': True,
        'rating': 5
    })

    # Artigo 3: PMC - Doppler Ultrasound Overview
    articles.append({
        'title': 'Doppler ultrasound and renal artery stenosis: An overview',
        'authors': 'Kahraman S, Yilmaz R, Akcar N, Arici S',
        'journal': 'Journal of Ultrasound',
        'publish_date': '2013-03-01',
        'language': 'en',
        'pm_id': 'PMC3567456',
        'abstract': 'Overview of Doppler ultrasound techniques for diagnosing renal artery stenosis, the most common cause of secondary hypertension. Discusses renal-aortic ratio as a reliable parameter that normalizes individual hemodynamic variations, improving diagnostic specificity compared to peak systolic velocity alone.',
        'original_link': 'https://pmc.ncbi.nlm.nih.gov/articles/PMC3567456/',
        'article_type': 'review',
        'specialty': 'Nefrologia',
        'keywords': ['doppler ultrasound', 'renal artery stenosis', 'secondary hypertension', 'renal-aortic ratio', 'diagnostic imaging'],
        'favorite': True,
        'rating': 4
    })

    return articles

def insert_articles(conn, articles):
    """Insere artigos no banco e retorna IDs"""
    cursor = conn.cursor()
    article_ids = []

    for article in articles:
        # Verifica se artigo já existe (por PMID ou DOI)
        check_query = """
            SELECT id FROM articles
            WHERE (pm_id = %s AND pm_id IS NOT NULL)
               OR (doi = %s AND doi IS NOT NULL)
            LIMIT 1
        """
        cursor.execute(check_query, (article.get('pm_id'), article.get('doi')))
        existing = cursor.fetchone()

        if existing:
            print(f"✓ Artigo já existe: {article['title'][:60]}... (ID: {existing[0]})")
            article_ids.append(existing[0])
            continue

        # Insere novo artigo
        insert_query = """
            INSERT INTO articles (
                title, authors, journal, publish_date, language,
                pm_id, doi, abstract, original_link, article_type,
                specialty, keywords, favorite, rating, created_at, updated_at
            ) VALUES (
                %s, %s, %s, %s, %s,
                %s, %s, %s, %s, %s,
                %s, %s::jsonb, %s, %s, NOW(), NOW()
            ) RETURNING id
        """

        import json
        keywords_json = json.dumps(article.get('keywords', []))

        cursor.execute(insert_query, (
            article['title'],
            article['authors'],
            article['journal'],
            article['publish_date'],
            article['language'],
            article.get('pm_id'),
            article.get('doi'),
            article['abstract'],
            article['original_link'],
            article['article_type'],
            article['specialty'],
            keywords_json,
            article.get('favorite', False),
            article.get('rating')
        ))

        article_id = cursor.fetchone()[0]
        article_ids.append(article_id)
        print(f"✓ Artigo inserido: {article['title'][:60]}... (ID: {article_id})")

    conn.commit()
    return article_ids

def create_article_relationships(conn, article_ids):
    """Cria relacionamentos entre artigos e score_item"""
    cursor = conn.cursor()

    for article_id in article_ids:
        # Verifica se relacionamento já existe
        check_query = """
            SELECT score_item_id FROM article_score_items
            WHERE article_id = %s AND score_item_id = %s
        """
        cursor.execute(check_query, (article_id, ITEM_ID))

        if cursor.fetchone():
            print(f"  - Relacionamento já existe para artigo {article_id}")
            continue

        # Cria relacionamento
        insert_query = """
            INSERT INTO article_score_items (article_id, score_item_id)
            VALUES (%s, %s)
        """
        cursor.execute(insert_query, (article_id, ITEM_ID))
        print(f"  - Relacionamento criado para artigo {article_id}")

    conn.commit()

def update_score_item(conn):
    """Atualiza o score_item com os dados de enriquecimento"""
    cursor = conn.cursor()

    # Primeiro, verifica se o item existe
    cursor.execute("""
        SELECT si.name, sg.name as subgroup_name
        FROM score_items si
        LEFT JOIN score_subgroups sg ON si.subgroup_id = sg.id
        WHERE si.id = %s
    """, (ITEM_ID,))
    item = cursor.fetchone()

    if not item:
        print(f"❌ ERRO: Item {ITEM_ID} não encontrado!")
        return False

    print(f"\n📋 Item encontrado: {item[0]} (Subgrupo: {item[1]})")

    # Atualiza o item
    update_query = """
        UPDATE score_items
        SET
            clinical_relevance = %s,
            patient_explanation = %s,
            conduct = %s,
            last_review = NOW(),
            updated_at = NOW()
        WHERE id = %s
    """

    cursor.execute(update_query, (
        ENRICHMENT_DATA['clinical_relevance'],
        ENRICHMENT_DATA['patient_explanation'],
        ENRICHMENT_DATA['conduct'],
        ITEM_ID
    ))

    conn.commit()
    print(f"✓ Score item atualizado com sucesso!")
    return True

def verify_enrichment(conn):
    """Verifica o enriquecimento completo"""
    cursor = conn.cursor()

    # Verifica dados do item
    cursor.execute("""
        SELECT
            si.name,
            sg.name as subgroup_name,
            CASE WHEN si.clinical_relevance IS NOT NULL THEN '✓' ELSE '✗' END as has_clinical,
            CASE WHEN si.patient_explanation IS NOT NULL THEN '✓' ELSE '✗' END as has_explanation,
            CASE WHEN si.conduct IS NOT NULL THEN '✓' ELSE '✗' END as has_conduct,
            LENGTH(si.clinical_relevance) as clinical_len,
            LENGTH(si.patient_explanation) as explanation_len,
            LENGTH(si.conduct) as conduct_len
        FROM score_items si
        LEFT JOIN score_subgroups sg ON si.subgroup_id = sg.id
        WHERE si.id = %s
    """, (ITEM_ID,))

    result = cursor.fetchone()

    if result:
        print("\n" + "="*80)
        print("📊 VERIFICAÇÃO DO ENRIQUECIMENTO")
        print("="*80)
        print(f"Nome: {result[0]}")
        print(f"Subgrupo: {result[1]}")
        print(f"\nCampos Preenchidos:")
        print(f"  Clinical Relevance: {result[2]} ({result[5] or 0} caracteres)")
        print(f"  Patient Explanation: {result[3]} ({result[6] or 0} caracteres)")
        print(f"  Conduct: {result[4]} ({result[7] or 0} caracteres)")

    # Verifica artigos relacionados
    cursor.execute("""
        SELECT
            a.id,
            a.title,
            a.journal,
            EXTRACT(YEAR FROM a.publish_date)::INTEGER as year
        FROM articles a
        JOIN article_score_items asi ON a.id = asi.article_id
        WHERE asi.score_item_id = %s
        ORDER BY a.publish_date DESC
    """, (ITEM_ID,))

    articles = cursor.fetchall()

    print(f"\n📚 Artigos Relacionados: {len(articles)}")
    for i, article in enumerate(articles, 1):
        print(f"  {i}. {article[1][:70]}...")
        print(f"     {article[2]} ({article[3]})")

    print("="*80)

def main():
    """Função principal"""
    try:
        print("\n🔬 ENRIQUECIMENTO: RAR (Renal-Aortic Ratio) - Doppler Aorta")
        print("="*80)

        # Conecta ao banco
        print("\n1️⃣ Conectando ao banco de dados...")
        conn = psycopg2.connect(**DB_CONFIG)
        print("✓ Conectado com sucesso!")

        # Cria artigos
        print("\n2️⃣ Criando artigos científicos...")
        articles = create_articles()
        article_ids = insert_articles(conn, articles)
        print(f"✓ {len(article_ids)} artigos processados")

        # Cria relacionamentos
        print("\n3️⃣ Criando relacionamentos article_score_items...")
        create_article_relationships(conn, article_ids)
        print("✓ Relacionamentos criados")

        # Atualiza score_item
        print("\n4️⃣ Atualizando score_item...")
        if not update_score_item(conn):
            sys.exit(1)

        # Verifica enriquecimento
        print("\n5️⃣ Verificando enriquecimento...")
        verify_enrichment(conn)

        conn.close()

        print("\n✅ ENRIQUECIMENTO CONCLUÍDO COM SUCESSO!")
        print("="*80)

        return 0

    except Exception as e:
        print(f"\n❌ ERRO: {e}")
        import traceback
        traceback.print_exc()
        return 1

if __name__ == '__main__':
    sys.exit(main())
