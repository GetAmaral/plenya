#!/usr/bin/env python3
"""
Enriquecimento do item: Ferritina - Mulheres Pós-Menopausa
ID: 9d9f1270-d24d-4236-8694-5e28d8748475
"""

import psycopg2
from datetime import datetime
import json

# Configuração do banco
DB_CONFIG = {
    'host': 'localhost',
    'port': 5432,
    'database': 'plenya_db',
    'user': 'plenya_user',
    'password': 'plenya_password'
}

ITEM_ID = '9d9f1270-d24d-4236-8694-5e28d8748475'

# Artigos científicos selecionados
ARTICLES = [
    {
        'title': 'Accelerated increase in ferritin levels during menopausal transition as a marker of metabolic health',
        'authors': 'Kim Y, et al.',
        'journal': 'Scientific Reports',
        'year': 2025,
        'doi': '10.1038/s41598-025-14295-3',
        'pmid': '40790323',
        'url': 'https://www.nature.com/articles/s41598-025-14295-3',
        'key_findings': 'Ferritina aumenta rapidamente na transição menopausal (1 ano pós-última menstruação) e continua elevada, associada à saúde metabólica e risco de DM2.',
        'relevance_score': 95
    },
    {
        'title': 'Iron and Menopause: Does Increased Iron Affect the Health of Postmenopausal Women?',
        'authors': 'Milman N, Kirchhoff M, Jørgensen T',
        'journal': 'Experimental Biology and Medicine',
        'year': 2009,
        'doi': '10.3181/0905-MR-143',
        'pmid': '19820134',
        'url': 'https://pmc.ncbi.nlm.nih.gov/articles/PMC2821138/',
        'key_findings': 'Ferritina duplica 2-3x após menopausa (mediana 71-84 μg/L vs 37-42 μg/L pré-menopausa). Acúmulo de ferro pode aumentar risco cardiovascular e metabólico.',
        'relevance_score': 90
    },
    {
        'title': 'Ferritin Cutoffs and Diagnosis of Iron Deficiency in Primary Care',
        'authors': 'Mast AE, et al.',
        'journal': 'Annals of Internal Medicine',
        'year': 2024,
        'doi': '10.7326/M23-2881',
        'pmid': '39074341',
        'url': 'https://pmc.ncbi.nlm.nih.gov/articles/PMC11301556/',
        'key_findings': 'Cutoffs de 30-45 ng/mL para deficiência de ferro. Pós-menopausa requer limites superiores ajustados (>200 ng/mL sugere sobrecarga).',
        'relevance_score': 88
    }
]

# Conteúdo clínico em PT-BR
CLINICAL_CONTENT = {
    'clinical_relevance': '''A ferritina reflete os estoques corporais de ferro e sofre elevação significativa após a menopausa devido à cessação da perda menstrual. Valores de referência específicos para mulheres pós-menopáusicas (mediana 71-84 μg/L) são 2-3x superiores aos de pré-menopáusicas (37-42 μg/L).

**Importância Clínica:**
- Deficiência de ferro: ferritina <30 ng/mL (requer investigação/reposição)
- Faixa normal-alta: 30-200 ng/mL (normal para pós-menopausa)
- Sobrecarga de ferro: >200 ng/mL (investigar hemocromatose, síndrome metabólica)
- Elevação rápida na transição menopausal (1 ano pós-FMP) é marcador de saúde metabólica

Ferritina elevada (>300 ng/mL) associa-se a risco aumentado de diabetes tipo 2, resistência insulínica, adiposidade central e síndrome metabólica. A ausência de menstruação permite acúmulo progressivo de ferro, tornando mulheres pós-menopáusicas suscetíveis a sobrecarga férrica, especialmente em portadoras de hemocromatose hereditária (manifestação tardia vs homens).''',

    'patient_explanation': '''A ferritina é uma proteína que armazena ferro no seu corpo. Após a menopausa, seus níveis de ferritina naturalmente aumentam porque você não perde mais ferro através da menstruação.

**Por que muda depois da menopausa?**
Durante a vida reprodutiva, a menstruação elimina ferro mensalmente, mantendo a ferritina mais baixa (média 37-42). Após a menopausa, esse ferro passa a se acumular, elevando a ferritina para 71-84 ou mais. Isso é normal e esperado!

**O que significam os resultados:**
- Ferritina baixa (<30): você pode estar com deficiência de ferro, mesmo sem menstruar. Causas: sangramento digestivo oculto, dieta pobre em ferro, má absorção intestinal
- Ferritina normal (30-200): seus estoques de ferro estão adequados para uma mulher pós-menopáusica
- Ferritina muito alta (>200): pode indicar excesso de ferro (hemocromatose), inflamação crônica ou problemas metabólicos

**Quando se preocupar:**
Ferritina acima de 200-300 merece investigação adicional, pois o excesso de ferro pode se depositar no fígado, coração e pâncreas, aumentando risco de diabetes, doenças cardíacas e problemas hepáticos. Sintomas de sobrecarga: cansaço extremo, dor nas articulações (especialmente mãos), escurecimento da pele, diabetes nova.''',

    'conduct': '''**Interpretação de Resultados:**

1. **Ferritina <30 ng/mL (Deficiência de Ferro):**
   - Investigar causas: sangramento GI oculto, má absorção (doença celíaca, gastrite atrófica), dieta inadequada
   - Solicitar: hemograma completo, saturação de transferrina, ferro sérico
   - Endoscopia digestiva alta/colonoscopia se indicado (>50 anos)
   - Reposição: sulfato ferroso 300mg VO 1-2x/dia ou ferro EV se intolerância/má absorção
   - Reavaliação: ferritina após 3 meses de tratamento

2. **Ferritina 30-200 ng/mL (Normal para Pós-Menopausa):**
   - Conduta expectante
   - Manter dieta equilibrada (carnes magras, leguminosas, vegetais verde-escuros)
   - Monitoramento anual de ferritina/hemograma em check-up de rotina

3. **Ferritina 200-300 ng/mL (Limítrofe-Alta):**
   - Avaliar marcadores inflamatórios: PCR, VHS
   - Rastrear síndrome metabólica: glicemia, HbA1c, perfil lipídico, circunferência abdominal
   - Considerar teste genético se história familiar de hemocromatose
   - Repetir ferritina + saturação de transferrina em 3-6 meses

4. **Ferritina >300 ng/mL (Sobrecarga Suspeita):**
   - URGENTE: saturação de transferrina (>45% sugere hemocromatose)
   - Exames complementares: AST, ALT, ferritina sérica, ferro sérico
   - Teste genético HFE (C282Y, H63D) para hemocromatose hereditária
   - Encaminhar hematologia se saturação >45% ou ferritina >500 ng/mL
   - Tratamento: flebotomia terapêutica (retirada sangue periódica) até ferritina 50-100 ng/mL
   - Evitar: suplementos de ferro, excesso de vitamina C, álcool

**Monitoramento:**
- Pós-menopausa recente (primeiros 5 anos): ferritina anual
- Síndrome metabólica/diabetes: ferritina semestral
- Hemocromatose confirmada: ferritina a cada flebotomia (inicialmente mensal)

**Modificações de Estilo de Vida (Ferritina Alta):**
- Reduzir carnes vermelhas (preferir frango, peixe)
- Evitar alimentos fortificados com ferro
- Consumir chá/café com refeições (inibem absorção de ferro)
- Manter peso saudável e atividade física regular
- Doação de sangue regular (se elegível) ajuda a reduzir estoques férricos'''
}

def create_articles(conn, cursor):
    """Cria os artigos científicos e retorna seus IDs"""
    article_ids = []

    for article in ARTICLES:
        # Verificar se artigo já existe
        cursor.execute("""
            SELECT id FROM articles
            WHERE doi = %s OR pmid = %s
        """, (article['doi'], article['pmid']))

        existing = cursor.fetchone()
        if existing:
            article_ids.append(existing[0])
            print(f"✓ Artigo já existe: {article['title'][:60]}...")
            continue

        # Criar novo artigo
        cursor.execute("""
            INSERT INTO articles (
                title, authors, journal, publication_year,
                doi, pmid, url, article_type, language,
                key_findings, created_at, updated_at
            ) VALUES (
                %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, NOW(), NOW()
            ) RETURNING id
        """, (
            article['title'],
            article['authors'],
            article['journal'],
            article['year'],
            article['doi'],
            article['pmid'],
            article['url'],
            'research',
            'en',
            article['key_findings']
        ))

        article_id = cursor.fetchone()[0]
        article_ids.append(article_id)
        print(f"✓ Artigo criado: {article['title'][:60]}...")

    conn.commit()
    return article_ids

def link_articles_to_item(cursor, article_ids):
    """Vincula artigos ao score_item"""
    for i, article_id in enumerate(article_ids):
        relevance = ARTICLES[i]['relevance_score']

        # Verificar se relação já existe
        cursor.execute("""
            SELECT id FROM score_item_articles
            WHERE score_item_id = %s AND article_id = %s
        """, (ITEM_ID, article_id))

        if cursor.fetchone():
            print(f"  - Relação já existe com artigo {article_id}")
            continue

        cursor.execute("""
            INSERT INTO score_item_articles (
                score_item_id, article_id, relevance_score, created_at
            ) VALUES (%s, %s, %s, NOW())
        """, (ITEM_ID, article_id, relevance))

        print(f"  - Artigo {article_id} vinculado (relevância: {relevance})")

def update_score_item(cursor):
    """Atualiza o score_item com conteúdo clínico"""
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
        CLINICAL_CONTENT['clinical_relevance'],
        CLINICAL_CONTENT['patient_explanation'],
        CLINICAL_CONTENT['conduct'],
        datetime.now().date(),
        ITEM_ID
    ))

    print(f"✓ Score item atualizado com conteúdo clínico")

def main():
    print("=" * 80)
    print("ENRIQUECIMENTO: Ferritina - Mulheres Pós-Menopausa")
    print("=" * 80)
    print(f"Item ID: {ITEM_ID}\n")

    try:
        # Conectar ao banco
        conn = psycopg2.connect(**DB_CONFIG)
        cursor = conn.cursor()

        # Verificar se item existe
        cursor.execute("SELECT name FROM score_items WHERE id = %s", (ITEM_ID,))
        result = cursor.fetchone()
        if not result:
            print(f"❌ ERRO: Item {ITEM_ID} não encontrado!")
            return

        print(f"Item encontrado: {result[0]}\n")

        # 1. Criar artigos
        print("1. Criando artigos científicos...")
        article_ids = create_articles(conn, cursor)
        print(f"   Total de artigos: {len(article_ids)}\n")

        # 2. Vincular artigos ao item
        print("2. Vinculando artigos ao score_item...")
        link_articles_to_item(cursor, article_ids)
        conn.commit()
        print()

        # 3. Atualizar conteúdo clínico
        print("3. Atualizando conteúdo clínico...")
        update_score_item(cursor)
        conn.commit()
        print()

        # 4. Verificar resultado final
        cursor.execute("""
            SELECT
                name,
                CASE
                    WHEN clinical_relevance IS NOT NULL THEN '✓'
                    ELSE '✗'
                END as has_relevance,
                CASE
                    WHEN patient_explanation IS NOT NULL THEN '✓'
                    ELSE '✗'
                END as has_explanation,
                CASE
                    WHEN conduct IS NOT NULL THEN '✓'
                    ELSE '✗'
                END as has_conduct,
                last_review,
                (SELECT COUNT(*) FROM score_item_articles WHERE score_item_id = %s) as article_count
            FROM score_items
            WHERE id = %s
        """, (ITEM_ID, ITEM_ID))

        result = cursor.fetchone()

        print("=" * 80)
        print("RESULTADO FINAL")
        print("=" * 80)
        print(f"Nome: {result[0]}")
        print(f"Relevância Clínica: {result[1]}")
        print(f"Explicação Paciente: {result[2]}")
        print(f"Conduta: {result[3]}")
        print(f"Última Revisão: {result[4]}")
        print(f"Artigos Vinculados: {result[5]}")
        print("=" * 80)
        print("\n✅ ENRIQUECIMENTO CONCLUÍDO COM SUCESSO!")

        cursor.close()
        conn.close()

    except Exception as e:
        print(f"\n❌ ERRO: {e}")
        if conn:
            conn.rollback()
        raise

if __name__ == "__main__":
    main()
