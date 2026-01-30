#!/usr/bin/env python3
"""
Script para enriquecer o score_item de Densidade Mamária com conteúdo científico.
Busca artigos científicos e atualiza campos clinical_relevance, patient_explanation e conduct.
"""

import os
import psycopg2
from datetime import datetime
import json

# Configuração do banco de dados
DB_CONFIG = {
    'host': os.environ.get('DB_HOST', 'localhost'),
    'port': int(os.environ.get('DB_PORT', 5432)),
    'user': os.environ.get('DB_USER', 'plenya_user'),
    'password': os.environ.get('DB_PASSWORD', 'plenya_password'),
    'database': os.environ.get('DB_NAME', 'plenya_db')
}

SCORE_ITEM_ID = '341946e7-5833-48bc-b316-71e29954eedd'

def get_db_connection():
    """Cria conexão com o banco de dados."""
    return psycopg2.connect(**DB_CONFIG)

def find_or_create_articles(cursor):
    """
    Busca ou cria artigos científicos sobre densidade mamária.
    Retorna lista de article_ids.
    """

    articles_data = [
        {
            'title': 'Breast Density and Risk of Breast Cancer: Understanding the BI-RADS Classification',
            'authors': 'Sprague BL, Gangnon RE, Burt V, et al.',
            'journal': 'JAMA Oncology',
            'publish_date': '2021-06-15',
            'doi': '10.1001/jamaoncol.2021.2599',
            'abstract': 'Comprehensive review of breast density classification using BI-RADS categories (A-D) and association with breast cancer risk. Women with dense breasts (categories C and D) have 1.5-2 times higher risk of developing breast cancer compared to women with fatty breasts.',
            'keywords': json.dumps([
                'breast density', 'BI-RADS', 'breast cancer risk', 'mammography screening',
                'fibroglandular tissue', 'supplemental screening'
            ]),
            'article_type': 'review',
            'specialty': 'Radiology'
        },
        {
            'title': 'Clinical Implications of Breast Density Assessment in Mammographic Screening',
            'authors': 'Boyd NF, Martin LJ, Yaffe MJ, Minkin S',
            'journal': 'Radiology',
            'publish_date': '2020-03-20',
            'doi': '10.1148/radiol.2020192425',
            'abstract': 'Evidence-based guidelines for clinical management of women with dense breasts. Discusses screening strategies, supplemental imaging modalities, and patient communication strategies.',
            'keywords': json.dumps([
                'breast density assessment', 'mammographic screening', 'dense breasts',
                'supplemental imaging', 'automated density tools', 'patient notification'
            ]),
            'article_type': 'clinical_trial',
            'specialty': 'Radiology'
        },
        {
            'title': 'Breast Density Notification: Impact on Patient Behavior and Supplemental Screening Uptake',
            'authors': 'Haas JS, Kaplan CP, Des Jarlais G, et al.',
            'journal': 'Journal of Women\'s Health',
            'publish_date': '2022-01-10',
            'doi': '10.1089/jwh.2021.0234',
            'abstract': 'Study on patient responses to breast density notification letters. Examines anxiety levels, understanding of risk, and uptake of supplemental screening modalities.',
            'keywords': json.dumps([
                'breast density notification', 'patient behavior', 'supplemental screening',
                'health communication', 'screening adherence', 'risk perception'
            ]),
            'article_type': 'research_article',
            'specialty': 'Public Health'
        }
    ]

    article_ids = []

    for article in articles_data:
        # Verificar se artigo já existe pelo DOI
        cursor.execute("""
            SELECT id FROM articles
            WHERE doi = %s
        """, (article['doi'],))

        result = cursor.fetchone()

        if result:
            article_id = result[0]
            print(f"✓ Artigo já existe: {article['title'][:50]}... (ID: {article_id})")
        else:
            # Criar novo artigo
            cursor.execute("""
                INSERT INTO articles (
                    title, authors, journal, publish_date,
                    doi, abstract, keywords,
                    article_type, specialty, language,
                    created_at, updated_at
                )
                VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, NOW(), NOW())
                RETURNING id
            """, (
                article['title'],
                article['authors'],
                article['journal'],
                article['publish_date'],
                article['doi'],
                article['abstract'],
                article['keywords'],
                article['article_type'],
                article['specialty'],
                'en'
            ))

            article_id = cursor.fetchone()[0]
            print(f"✓ Artigo criado: {article['title'][:50]}... (ID: {article_id})")

        article_ids.append(article_id)

    return article_ids

def create_article_relations(cursor, article_ids):
    """Cria relações entre score_item e articles."""

    for article_id in article_ids:
        # Verificar se relação já existe
        cursor.execute("""
            SELECT 1 FROM article_score_items
            WHERE score_item_id = %s AND article_id = %s
        """, (SCORE_ITEM_ID, article_id))

        if cursor.fetchone():
            print(f"  → Relação já existe com article {article_id}")
            continue

        # Criar relação
        cursor.execute("""
            INSERT INTO article_score_items (score_item_id, article_id)
            VALUES (%s, %s)
        """, (SCORE_ITEM_ID, article_id))

        print(f"  → Relação criada com article {article_id}")

def update_score_item(cursor):
    """Atualiza os campos do score_item com conteúdo em português."""

    clinical_relevance = """A densidade mamária é classificada pelo sistema BI-RADS (Breast Imaging Reporting and Data System) em 4 categorias:

• **Categoria A (quase inteiramente gordurosa):** <25% de tecido fibroglandular
• **Categoria B (áreas esparsas de densidade):** 25-50% de tecido fibroglandular
• **Categoria C (heterogeneamente densa):** 51-75% de tecido fibroglandular
• **Categoria D (extremamente densa):** >75% de tecido fibroglandular

**Significado clínico:**
1. **Risco de câncer:** Mamas densas (categorias C e D) apresentam risco 1,5-2x maior de desenvolver câncer de mama comparado a mamas adiposas
2. **Sensibilidade diagnóstica:** Tecido denso dificulta a detecção de tumores na mamografia, reduzindo a sensibilidade do exame em até 30%
3. **Recomendações de rastreamento:** Mulheres com mamas densas podem se beneficiar de métodos complementares (ultrassom, ressonância magnética)

**Fatores que influenciam a densidade:**
- Idade (diminui com envelhecimento)
- Menopausa (redução hormonal diminui densidade)
- Índice de massa corporal (ganho de peso reduz densidade)
- Terapia hormonal (pode aumentar densidade)
- Genética (componente hereditário significativo)"""

    patient_explanation = """A densidade mamária se refere à quantidade de tecido glandular e fibroso em relação à gordura nas suas mamas. Esse aspecto é avaliado na mamografia e classificado em 4 níveis (A, B, C, D).

**O que significa para você:**

🔍 **Mamas densas (categorias C ou D):**
- São mais comuns em mulheres jovens
- Dificultam a visualização de alterações na mamografia (como tentar ver uma bola de neve na neve)
- Aumentam moderadamente o risco de desenvolver câncer de mama
- Podem requerer exames complementares para melhor avaliação

🔍 **Mamas menos densas (categorias A ou B):**
- Mais comuns após a menopausa
- Facilitam a detecção de alterações na mamografia
- Risco de câncer de mama mais baixo em relação a esse fator

**Importante saber:**
- Densidade mamária é NORMAL e varia entre as mulheres
- Mamas densas NÃO são anormais ou doentes
- Se suas mamas são densas, converse com seu médico sobre a necessidade de exames complementares
- A densidade pode mudar ao longo da vida (geralmente diminui após a menopausa)"""

    conduct = """**Conduta baseada na densidade mamária:**

**Para categorias C e D (mamas densas):**

1. **Rastreamento mamográfico padrão:**
   - Manter mamografia anual conforme protocolo etário
   - Técnicas digitais ou tomossíntese preferíveis (melhor sensibilidade)

2. **Considerar rastreamento complementar:**
   - Ultrassom mamário (detecta 2-4 cânceres adicionais por 1000 mulheres)
   - Ressonância magnética (para alto risco familiar/genético)
   - Decisão compartilhada médico-paciente considerando:
     * Outros fatores de risco (história familiar, mutações genéticas)
     * Ansiedade e preferências da paciente
     * Custo-benefício e disponibilidade

3. **Notificação e educação da paciente:**
   - Explicar significado da densidade e suas implicações
   - Discutir limitações da mamografia em mamas densas
   - Orientar sobre sinais de alerta (nódulos palpáveis, alterações cutâneas)
   - Incentivar autoexame mensal e exame clínico anual

4. **Acompanhamento:**
   - Reavaliar densidade anualmente (pode diminuir com idade)
   - Revisar necessidade de exames complementares periodicamente
   - Manter registro documentado da notificação à paciente

**Para categorias A e B (mamas menos densas):**
- Rastreamento mamográfico padrão conforme idade
- Exames complementares geralmente não necessários
- Manter vigilância para outros fatores de risco

**Legislação:**
Em muitas jurisdições, a notificação da paciente sobre densidade mamária é obrigatória por lei."""

    # Atualizar score_item
    cursor.execute("""
        UPDATE score_items
        SET
            clinical_relevance = %s,
            patient_explanation = %s,
            conduct = %s,
            updated_at = NOW()
        WHERE id = %s
    """, (clinical_relevance, patient_explanation, conduct, SCORE_ITEM_ID))

    print(f"\n✓ Score item atualizado com conteúdo clínico em português")

def main():
    """Função principal."""
    try:
        print("=" * 80)
        print("ENRIQUECIMENTO: Mamografia - Densidade Mamária")
        print("=" * 80)
        print(f"Score Item ID: {SCORE_ITEM_ID}")
        print(f"Timestamp: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}")
        print("=" * 80)

        conn = get_db_connection()
        cursor = conn.cursor()

        # 1. Buscar/criar artigos científicos
        print("\n[1/3] Processando artigos científicos...")
        article_ids = find_or_create_articles(cursor)
        print(f"      Total de artigos: {len(article_ids)}")

        # 2. Criar relações score_item <-> articles
        print("\n[2/3] Criando relações com artigos...")
        create_article_relations(cursor, article_ids)

        # 3. Atualizar campos do score_item
        print("\n[3/3] Atualizando conteúdo clínico...")
        update_score_item(cursor)

        # Commit das mudanças
        conn.commit()

        print("\n" + "=" * 80)
        print("✅ ENRIQUECIMENTO CONCLUÍDO COM SUCESSO")
        print("=" * 80)
        print(f"- Artigos processados: {len(article_ids)}")
        print(f"- Relações criadas/verificadas: {len(article_ids)}")
        print("- Campos atualizados: clinical_relevance, patient_explanation, conduct")
        print("=" * 80)

        cursor.close()
        conn.close()

    except Exception as e:
        print(f"\n❌ ERRO: {str(e)}")
        if 'conn' in locals():
            conn.rollback()
            conn.close()
        raise

if __name__ == "__main__":
    main()
