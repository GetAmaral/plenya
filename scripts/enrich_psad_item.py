#!/usr/bin/env python3
"""
Enrichment Script: USG Próstata - Densidade PSA (PSAD)
ID: 317acc85-3ce9-4f97-8e14-799354166f5e
Grupo: Exames > Imagem

Evidência científica baseada em:
- Peng et al. (2025) - BMC Urology - DOI: 10.1186/s12894-025-01719-5
- Chou et al. (2025) - Diagnostics - DOI: 10.3390/diagnostics15162027
- Yusim et al. (2020) - Scientific Reports - DOI: 10.1038/s41598-020-76786-9
"""

import psycopg2
import os
import sys
from datetime import datetime

# Database configuration
DB_CONFIG = {
    'host': os.getenv('DB_HOST', 'db'),
    'port': os.getenv('DB_PORT', '5432'),
    'database': os.getenv('DB_NAME', 'plenya_db'),
    'user': os.getenv('DB_USER', 'plenya_user'),
    'password': os.getenv('DB_PASSWORD', 'plenya_password')
}

# Target score item
ITEM_ID = '317acc85-3ce9-4f97-8e14-799354166f5e'

# Clinical content (PT-BR)
CLINICAL_RELEVANCE = """A Densidade do PSA (PSAD) é um parâmetro calculado pela divisão do valor do PSA sérico (ng/mL) pelo volume prostático obtido por ultrassonografia (cm³). Este índice melhora significativamente a capacidade de predição de câncer de próstata clinicamente significativo em comparação ao PSA isolado.

**Fundamento Clínico:**
A PSAD corrige o efeito do volume prostático no valor do PSA, permitindo diferenciar melhor entre elevações do PSA causadas por hiperplasia prostática benigna (HPB) e aquelas relacionadas a neoplasia maligna. Estudos demonstram que a PSAD apresenta área sob a curva ROC de 0.78-0.85, superior ao PSA isolado (AUC 0.64-0.72).

**Pontos de Corte Baseados em Evidências:**

1. **PSAD < 0.09-0.10 ng/mL/cm³:**
   - Baixíssimo risco de câncer clinicamente significativo (4-6%)
   - Pode-se considerar evitar biópsia, especialmente se RM negativa

2. **PSAD 0.10-0.15 ng/mL/cm³:**
   - Risco intermediário
   - Decisão compartilhada considerando RM e fatores individuais

3. **PSAD 0.15-0.20 ng/mL/cm³:**
   - Alto risco
   - Forte indicação de biópsia se lesão suspeita na RM

4. **PSAD ≥ 0.20-0.30 ng/mL/cm³:**
   - Risco muito alto de câncer significativo
   - Indicação clara de biópsia mesmo com RM negativa ou equívoca

**Integração com Ressonância Magnética:**
Para pacientes com PI-RADS 3 (achados equívocos), um limiar de PSAD ≥ 0.20 ng/mL/cm³ demonstra melhor balanço entre sensibilidade e especificidade. Em pacientes com RM negativa (PI-RADS ≤2) e PSAD < 0.15 ng/mL/cm³, pode-se considerar evitar biópsia através de decisão compartilhada.

**Considerações Especiais:**
- A acurácia da PSAD diminui em próstatas muito volumosas (>80 mL)
- A combinação de PSAD com cinética do PSA (variação percentual ao longo do tempo) pode reduzir ainda mais biópsias desnecessárias
- Em pacientes com HPB obstrutiva e RM negativa, o limiar de 0.30 ng/mL/cm³ demonstrou especificidade de 93% com sensibilidade de 65%

**Referências Científicas:**
- Peng Y, et al. BMC Urol. 2025;25:38. DOI: 10.1186/s12894-025-01719-5
- Chou YJ, et al. Diagnostics. 2025;15(16):2027. DOI: 10.3390/diagnostics15162027
- Yusim I, et al. Sci Rep. 2020;10:20015. DOI: 10.1038/s41598-020-76786-9"""

PATIENT_EXPLANATION = """A **Densidade do PSA (PSAD)** é um cálculo que ajuda os médicos a entenderem melhor se uma alteração no seu exame de PSA pode ser preocupante ou não.

**Como funciona:**
O médico pega o valor do seu PSA (aquele exame de sangue) e divide pelo tamanho da sua próstata (medido no ultrassom). Por exemplo, se seu PSA é 6 ng/mL e sua próstata tem 40 cm³ de volume, sua PSAD seria 6 ÷ 40 = 0.15 ng/mL/cm³.

**Por que isso é importante:**
Próstatas maiores naturalmente produzem mais PSA. Então, um PSA de 6 ng/mL pode ser normal se você tem uma próstata grande por hiperplasia benigna (crescimento não-canceroso), mas pode ser preocupante se sua próstata for pequena.

**O que os números significam:**

• **PSAD abaixo de 0.10:** Muito tranquilo. O risco de ter um câncer significativo é baixíssimo (cerca de 4%). Geralmente não precisa fazer biópsia.

• **PSAD entre 0.10 e 0.15:** Risco intermediário. Seu médico vai avaliar junto com outros exames (como a ressonância) se precisa investigar mais.

• **PSAD entre 0.15 e 0.20:** Risco aumentado. Provavelmente será necessária uma biópsia, especialmente se houver algo suspeito na ressonância.

• **PSAD acima de 0.20:** Risco alto. Forte indicação de fazer biópsia para investigar melhor, mesmo que a ressonância não mostre algo muito evidente.

**Vantagens deste cálculo:**
- Ajuda a evitar biópsias desnecessárias em homens com próstatas grandes e PSA elevado apenas por isso
- Identifica melhor os casos que realmente precisam de investigação aprofundada
- Quando combinado com a ressonância magnética, melhora ainda mais a precisão do diagnóstico

**Importante:**
A PSAD é uma ferramenta auxiliar. Seu médico vai considerar este valor junto com sua idade, histórico familiar, toque retal, resultado da ressonância e outros fatores para decidir os próximos passos do seu acompanhamento."""

CONDUCT = """**Interpretação e Conduta Baseada na PSAD:**

**1. PSAD < 0.10 ng/mL/cm³:**
- ✓ Considerar vigilância ativa sem biópsia imediata
- ✓ Repetir PSA e USG em 12 meses
- ✓ Se RM foi realizada e negativa (PI-RADS ≤2), risco de câncer significativo é ~4%
- ⚠️ Discussão compartilhada com paciente sobre risco residual

**2. PSAD 0.10-0.15 ng/mL/cm³:**
- ⚡ Zona cinzenta - individualizar decisão
- ✓ RM de próstata mandatória se ainda não realizada
- ✓ Se PI-RADS ≤2: considerar seguimento sem biópsia (risco ~6%)
- ✓ Se PI-RADS 3-5: indicar biópsia dirigida por fusão
- ✓ Avaliar cinética do PSA (variação temporal)
- ✓ Considerar testes complementares (PHI, 4K Score, PCA3)

**3. PSAD 0.15-0.20 ng/mL/cm³:**
- 🔴 Alto risco - biópsia geralmente indicada
- ✓ RM obrigatória antes da biópsia
- ✓ Biópsia dirigida por fusão RM-US se disponível
- ✓ Se PI-RADS ≥3, sensibilidade para câncer significativo é ~70%

**4. PSAD ≥ 0.20 ng/mL/cm³:**
- 🔴🔴 Risco muito alto - biópsia fortemente indicada
- ✓ Indicação de biópsia independente do resultado da RM
- ✓ Mesmo com PI-RADS ≤2, considerar biópsia sistemática
- ✓ Atenção especial ao Gleason score e volume tumoral

**5. PSAD ≥ 0.30 ng/mL/cm³ (HPB com RM negativa):**
- 🔴🔴🔴 Em contexto de HPB obstrutiva com RM negativa
- ✓ Especificidade de 93%, sensibilidade 65%
- ✓ Forte indicação de biópsia antes de considerar cirurgia para HPB

**Protocolo de Biópsia Recomendado:**
- **1ª linha:** Biópsia por fusão RM-ultrassom (se RM prévia)
- **2ª linha:** Biópsia sistemática 12-14 fragmentos se fusão indisponível
- **Seguimento pós-biópsia negativa:** Repetir PSA + PSAD em 6-12 meses

**Situações Especiais:**

**Próstatas Volumosas (>80 mL):**
- A PSAD perde acurácia
- Considerar peso maior para cinética do PSA e RM
- Limiares mais conservadores (favorecer biópsia se PSAD ≥ 0.15)

**Idade Avançada (>75 anos):**
- Avaliar expectativa de vida e comorbidades
- Considerar limiares mais altos (PSAD ≥ 0.20) para indicar biópsia
- Privilegiar qualidade de vida vs. diagnóstico precoce

**Vigilância Ativa de Câncer Confirmado:**
- PSAD > 0.15-0.20 pode sugerir reclassificação
- Aumento progressivo da PSAD indica progressão
- Trigger para biópsia de confirmação

**Integração com Outros Biomarcadores:**
- **PHI (Prostate Health Index):** Complementa PSAD na zona cinzenta
- **Cinética PSA:** Variação >20% ao ano aumenta risco
- **PSA livre/total < 15%:** Somado a PSAD >0.15 reforça indicação de biópsia

**Documentação Mandatória:**
- Volume prostático (método: elipsoide, planimetria)
- PSA sérico contemporâneo ao USG (idealmente <3 meses)
- Cálculo explícito: PSAD = PSA (ng/mL) / Volume (cm³)
- Contexto clínico: toque retal, sintomas urinários, histórico familiar"""

# Articles to insert (if not exist)
ARTICLES = [
    {
        'title': 'Optimal PSA density threshold for prostate biopsy in benign prostatic obstruction patients with elevated PSA levels but negative MRI findings',
        'authors': 'Peng Y, Wei C, Li Y, Zhao F, Liu Y, Jiang T, Chen Z, Zheng J, Fu J, Wang P, Shen W',
        'journal': 'BMC Urology',
        'year': 2025,
        'doi': '10.1186/s12894-025-01719-5',
        'url': 'https://pmc.ncbi.nlm.nih.gov/articles/PMC11874838/',
        'pmid': None,
        'summary': 'Study identifying optimal PSAD cutoff of 0.30 ng/ml/cm³ for biopsy decision in BPH patients with elevated PSA but negative MRI, demonstrating 93% specificity and 65% sensitivity. ROC analysis showed PSAD achieved AUC 0.848, outperforming PSA alone (0.722) and free/total PSA ratio (0.635).'
    },
    {
        'title': 'Integrating PSA Change with PSA Density Enhances Diagnostic Accuracy and Helps Avoid Unnecessary Prostate Biopsies',
        'authors': 'Chou YJ, Jong BE, Tsai YC',
        'journal': 'Diagnostics (Basel)',
        'year': 2025,
        'doi': '10.3390/diagnostics15162027',
        'url': 'https://pmc.ncbi.nlm.nih.gov/articles/PMC12385582/',
        'pmid': '40870878',
        'summary': 'Demonstrates that PSA density shows superior diagnostic performance (AUC 0.77-0.81) compared to PSA change alone. Combining both metrics provides optimal results, with >20% PSA decline criterion improving PSAD performance, especially valuable in prostates >80 mL where PSAD accuracy decreases.'
    },
    {
        'title': 'The use of prostate specific antigen density to predict clinically significant prostate cancer',
        'authors': 'Yusim I, Krenawi M, Mazor E, Novack V, Mabjeesh NJ',
        'journal': 'Scientific Reports',
        'year': 2020,
        'doi': '10.1038/s41598-020-76786-9',
        'url': 'https://pmc.ncbi.nlm.nih.gov/articles/PMC7672084/',
        'pmid': '33203873',
        'summary': 'Evaluated 992 men undergoing biopsy, finding PSAD AUC of 0.78 vs PSA AUC of 0.64 for predicting clinically significant cancer. Key thresholds: PSAD <0.09 ng/ml² only 4% risk; PSAD 0.09-0.19 ng/ml² risk increases with smaller prostates; PSAD ≥0.20 ng/ml² optimal cutoff with 70% sensitivity and 79% specificity.'
    }
]


def connect_db():
    """Connect to PostgreSQL database."""
    try:
        conn = psycopg2.connect(**DB_CONFIG)
        print(f"✓ Connected to database: {DB_CONFIG['database']}")
        return conn
    except Exception as e:
        print(f"✗ Database connection failed: {e}")
        sys.exit(1)


def verify_item_exists(cursor, item_id):
    """Verify that the score item exists."""
    cursor.execute("""
        SELECT id, question_text
        FROM score_items
        WHERE id = %s
    """, (item_id,))

    result = cursor.fetchone()
    if result:
        print(f"✓ Found score_item: {result[1]}")
        return True
    else:
        print(f"✗ Score item not found: {item_id}")
        return False


def insert_or_get_article(cursor, article):
    """Insert article if not exists, return article ID."""
    # Check if article exists by DOI
    cursor.execute("""
        SELECT id FROM articles WHERE doi = %s
    """, (article['doi'],))

    result = cursor.fetchone()
    if result:
        print(f"  ✓ Article exists: {article['title'][:60]}... (ID: {result[0]})")
        return result[0]

    # Insert new article
    cursor.execute("""
        INSERT INTO articles (
            title, authors, journal, year, doi, url, pmid, summary, created_at, updated_at
        ) VALUES (%s, %s, %s, %s, %s, %s, %s, %s, NOW(), NOW())
        RETURNING id
    """, (
        article['title'],
        article['authors'],
        article['journal'],
        article['year'],
        article['doi'],
        article['url'],
        article['pmid'],
        article['summary']
    ))

    article_id = cursor.fetchone()[0]
    print(f"  ✓ Inserted article: {article['title'][:60]}... (ID: {article_id})")
    return article_id


def link_article_to_item(cursor, item_id, article_id):
    """Create many-to-many relationship between score_item and article."""
    # Check if relationship already exists
    cursor.execute("""
        SELECT 1 FROM score_item_articles
        WHERE score_item_id = %s AND article_id = %s
    """, (item_id, article_id))

    if cursor.fetchone():
        print(f"  ✓ Link already exists: item {item_id} <-> article {article_id}")
        return

    # Create relationship
    cursor.execute("""
        INSERT INTO score_item_articles (score_item_id, article_id)
        VALUES (%s, %s)
    """, (item_id, article_id))

    print(f"  ✓ Created link: item {item_id} <-> article {article_id}")


def update_score_item(cursor, item_id, clinical_relevance, patient_explanation, conduct):
    """Update score_item with clinical content."""
    cursor.execute("""
        UPDATE score_items
        SET
            clinical_relevance = %s,
            patient_explanation = %s,
            conduct = %s,
            last_review_date = %s,
            updated_at = NOW()
        WHERE id = %s
    """, (
        clinical_relevance,
        patient_explanation,
        conduct,
        datetime.now().date(),
        item_id
    ))

    print(f"✓ Updated score_item with clinical content")


def main():
    """Main execution."""
    print("=" * 80)
    print("PSAD (Densidade PSA) - Clinical Enrichment Script")
    print("=" * 80)
    print()

    conn = connect_db()
    cursor = conn.cursor()

    try:
        # 1. Verify item exists
        print("\n[1/4] Verifying score_item exists...")
        if not verify_item_exists(cursor, ITEM_ID):
            print("\n✗ Item not found. Aborting.")
            return

        # 2. Insert articles
        print("\n[2/4] Processing scientific articles...")
        article_ids = []
        for article in ARTICLES:
            article_id = insert_or_get_article(cursor, article)
            article_ids.append(article_id)

        # 3. Link articles to item
        print("\n[3/4] Creating article-item relationships...")
        for article_id in article_ids:
            link_article_to_item(cursor, ITEM_ID, article_id)

        # 4. Update score_item with clinical content
        print("\n[4/4] Updating score_item with clinical content...")
        update_score_item(
            cursor,
            ITEM_ID,
            CLINICAL_RELEVANCE,
            PATIENT_EXPLANATION,
            CONDUCT
        )

        # Commit transaction
        conn.commit()
        print("\n" + "=" * 80)
        print("✓ SUCCESS! PSAD enrichment completed successfully")
        print("=" * 80)
        print(f"\nSummary:")
        print(f"  - Articles processed: {len(article_ids)}")
        print(f"  - Item updated: {ITEM_ID}")
        print(f"  - Clinical fields: 3 (clinical_relevance, patient_explanation, conduct)")
        print(f"  - Last review date: {datetime.now().date()}")

    except Exception as e:
        conn.rollback()
        print(f"\n✗ ERROR: {e}")
        import traceback
        traceback.print_exc()
        sys.exit(1)

    finally:
        cursor.close()
        conn.close()
        print("\n✓ Database connection closed")


if __name__ == '__main__':
    main()
