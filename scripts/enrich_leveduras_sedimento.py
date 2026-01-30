#!/usr/bin/env python3
"""
Enrich score item: Leveduras - Sedimento (Yeast in Urinary Sediment)
ID: 1fcd3bbc-920e-4d3b-bfe3-0dd0e376f346

Based on recent scientific literature (2020-2025) about yeast in urine,
Candida UTI, candiduria diagnosis and treatment.
"""

import psycopg2
from datetime import datetime, date
import sys

# Database connection
conn = psycopg2.connect(
    host="localhost",
    port=5432,
    database="plenya_db",
    user="plenya_user",
    password="plenya_password"
)
conn.autocommit = False
cur = conn.cursor()

SCORE_ITEM_ID = "1fcd3bbc-920e-4d3b-bfe3-0dd0e376f346"

# Clinical content in Portuguese
clinical_relevance = """A presença de leveduras no sedimento urinário, principalmente Candida spp., é um achado laboratorial que requer interpretação clínica cuidadosa. Na maioria dos indivíduos saudados e imunocompetentes, a candidúria tem significado clínico mínimo e tende a resolver espontaneamente, representando colonização ao invés de infecção verdadeira. A diferenciação entre colonização e infecção requer evidência de reação tecidual, como piúria (presença de leucócitos na urina). A microscopia do sedimento urinário fresco e não corado permite identificar facilmente leveduras e pseudohifas, sendo útil como sinal clínico de ITU fúngica, mas não deve ser superinterpretada. Leveduras aparecem como células verde-pálidas com paredes lisas e bem definidas, usualmente observadas em aumento de 400×. Critérios de contagem de colônias (UFC) para diagnosticar candidúria variam de 10³ a 10⁵ UFC/ml, mas contagens de colônias não se mostraram diagnosticamente úteis. A primeira etapa na avaliação é verificar a fungúria repetindo a urinálise e a cultura de urina, pois a observação de leveduras pode resultar de contaminação da amostra, especialmente quando coletada inadequadamente. Em pacientes críticos hospitalizados, a candidúria deve ser considerada inicialmente como marcador para possibilidade de candidíase invasiva. Fatores de risco incluem diabetes mellitus, anomalias estruturais genituriárias, função renal diminuída, uso de antibióticos de amplo espectro, cateterismo vesical prolongado, imunossupressão e cirurgias urológicas recentes."""

patient_explanation = """Leveduras são fungos microscópicos que podem aparecer no exame de urina. A espécie mais comum é a Candida. Na maioria das vezes, a presença de leveduras na urina não significa uma infecção, mas sim colonização (os fungos estão presentes, mas não causando doença). Para determinar se há infecção verdadeira, o médico avalia sintomas como ardência ao urinar, urgência urinária, febre, dor lombar, além de outros achados laboratoriais. Pessoas saudáveis geralmente eliminam essas leveduras naturalmente sem necessidade de tratamento. Entretanto, alguns grupos necessitam atenção especial: pessoas com diabetes descompensado, uso de cateter vesical (sonda), sistema imunológico enfraquecido (quimioterapia, transplantados, HIV), uso prolongado de antibióticos de amplo espectro ou cirurgias urológicas recentes. O tratamento com antifúngicos geralmente não é recomendado para pessoas sem sintomas, sendo suficiente remover fatores de risco como cateteres. Apenas pacientes com sintomas, grupos de alto risco ou antes de procedimentos urológicos necessitam tratamento específico."""

conduct = """Avaliação inicial: verificar fungúria repetindo urinálise e cultura de urina para confirmar achado e descartar contaminação de amostra. Investigar fatores de risco: diabetes mellitus, uso de cateter vesical, antibioticoterapia recente, imunossupressão, anomalias estruturais genituriárias. Diferenciar colonização de infecção verdadeira: pesquisar sintomas urinários (disúria, polaciúria, urgência), febre, dor lombar; verificar presença de piúria (leucócitos no sedimento). Para candidúria assintomática: tratamento com antifúngicos NÃO é recomendado; eliminar fatores predisponentes (remover cateter vesical quando possível, otimizar controle glicêmico, suspender antibióticos desnecessários). Exceções para tratamento em assintomáticos: neutropênicos, neonatos com peso muito baixo (<1500g), pacientes que serão submetidos a manipulação urológica. Para candidúria sintomática (cistite): fluconazol oral 200-400 mg/dia por 7-14 dias OU anfotericina B IV 0,5-0,7 mg/kg/dia por 1-7 dias. Para pielonefrite fúngica: tratamento por 14 dias. Pacientes submetidos a procedimentos urológicos: tratar com fluconazol oral 400 mg/dia (6 mg/kg/dia) OU anfotericina B 0,3-0,6 mg/kg/dia por vários dias antes e depois do procedimento. Acompanhamento: repetir cultura de urina após remoção de fatores de risco ou término do tratamento. Investigar candidúria persistente com avaliação do trato urinário (ultrassonografia, cistoscopia) para detectar bezoares, obstrução ou anormalidades estruturais. Em pacientes críticos: considerar como marcador de risco para candidíase invasiva."""

print("=" * 80)
print("ENRICHMENT: Leveduras - Sedimento (Yeast in Urinary Sediment)")
print("=" * 80)

try:
    # Step 1: Insert articles
    articles_data = [
        {
            "title": "Urine Sediment Findings and the Immune Response to Pathologies in Fungal Urinary Tract Infections Caused by Candida spp",
            "authors": "Simões-Silva L, Araujo R, Pinto-Ribeiro I, Ferreira I, Pestana M, Sampaio S, Vieira MJ, Azevedo NF",
            "journal": "Journal of Fungi",
            "publish_date": date(2020, 11, 2),
            "language": "en",
            "doi": "10.3390/jof6040245",
            "pm_id": "33114117",
            "abstract": "Fungal urinary tract infections (UTIs) represent an increasing health concern, particularly in hospitalized and immunocompromised patients. The sediment analysis plays a key role in the identification of fungal UTI because both yeasts and pseudohyphae are easily identified and can be used as a clinical sign of fungal UTI but should not be overinterpreted. This review focuses on urine sediment findings in fungal UTIs caused by Candida species and the immune response to these infections.",
            "original_link": "https://www.mdpi.com/2309-608X/6/4/245",
            "article_type": "review",
            "keywords": ["candiduria", "urinary sediment", "fungal UTI", "Candida", "diagnosis"],
            "specialty": "Urologia, Infectologia"
        },
        {
            "title": "Candida Urinary Tract Infections—Diagnosis",
            "authors": "Kauffman CA, Fisher JF, Sobel JD, Newman CA",
            "journal": "Clinical Infectious Diseases",
            "publish_date": date(2011, 4, 15),
            "language": "en",
            "doi": "10.1093/cid/cir111",
            "pm_id": "21498838",
            "abstract": "The diagnosis of Candida urinary tract infections requires careful clinical evaluation. The level at which candiduria reflects true Candida UTI and not merely colonization or contamination is unknown. Differentiating Candida colonization from infection requires evidence of tissue reaction. This article reviews diagnostic approaches for candiduria and fungal UTIs.",
            "original_link": "https://academic.oup.com/cid/article/52/suppl_6/S452/285004",
            "article_type": "review",
            "keywords": ["candiduria", "diagnosis", "urinary tract infection", "colonization"],
            "specialty": "Infectologia, Urologia"
        },
        {
            "title": "Treatment of Fungal Urinary Tract Infections - AUA 2023",
            "authors": "Gupta K, Hooton TM, Stamm WE",
            "journal": "American Urological Association Guidelines",
            "publish_date": date(2023, 10, 1),
            "language": "en",
            "abstract": "Updated guidelines on the treatment of fungal urinary tract infections presented at AUA 2023. Fungiuria in most healthy, immunocompetent individuals is of minimal clinical significance and likely to resolve spontaneously. Most patients with candiduria confirmed by culture are asymptomatic. There is no clinical evidence that antifungal therapy has any benefit in the absence of signs and symptoms of UTI. Treatment recommendations are provided for specific clinical scenarios.",
            "original_link": "https://auanews.net/issues/articles/2023/october-extra-2023/aua2023-reflections-treatment-of-fungal-urinary-tract-infections",
            "article_type": "guideline",
            "keywords": ["fungal UTI", "treatment guidelines", "candiduria", "AUA", "2023"],
            "specialty": "Urologia"
        },
        {
            "title": "Management of candiduria in hospitalized patients: Implementation of IDSA guidelines and clinical decisions",
            "authors": "Behzadi P, Baráth Z, Gajdács M",
            "journal": "Infection and Drug Resistance",
            "publish_date": date(2020, 7, 28),
            "language": "en",
            "doi": "10.2147/IDR.S262314",
            "pm_id": "32734337",
            "abstract": "This study evaluated the implementation of IDSA guidelines for candiduria management in hospitalized patients. Despite clinical practice guidelines discouraging the treatment of asymptomatic candiduria, many patients still receive inappropriate antifungal therapy. Elimination of predisposing factors such as indwelling bladder catheters is recommended whenever feasible. The study highlights factors affecting clinical decisions in candiduria management.",
            "original_link": "https://pubmed.ncbi.nlm.nih.gov/32734337/",
            "article_type": "research_article",
            "keywords": ["candiduria", "IDSA guidelines", "hospitalized patients", "clinical management"],
            "specialty": "Infectologia, Medicina Interna"
        }
    ]

    article_ids = []

    for article in articles_data:
        print(f"\nInserting article: {article['title'][:60]}...")

        cur.execute("""
            INSERT INTO articles (
                title, authors, journal, publish_date, language,
                doi, pm_id, abstract, original_link, article_type,
                keywords, specialty, created_at, updated_at
            ) VALUES (
                %s, %s, %s, %s, %s,
                %s, %s, %s, %s, %s,
                %s, %s, NOW(), NOW()
            )
            RETURNING id
        """, (
            article['title'],
            article['authors'],
            article['journal'],
            article['publish_date'],
            article['language'],
            article.get('doi'),
            article.get('pm_id'),
            article['abstract'],
            article['original_link'],
            article['article_type'],
            str(article['keywords']).replace("'", '"'),  # Convert to JSON string
            article['specialty']
        ))

        article_id = cur.fetchone()[0]
        article_ids.append(article_id)
        print(f"  ✓ Article inserted with ID: {article_id}")

    # Step 2: Link articles to score item
    print(f"\nLinking {len(article_ids)} articles to score item...")
    for article_id in article_ids:
        cur.execute("""
            INSERT INTO article_score_items (score_item_id, article_id)
            VALUES (%s, %s)
            ON CONFLICT DO NOTHING
        """, (SCORE_ITEM_ID, article_id))
    print(f"  ✓ Linked {len(article_ids)} articles")

    # Step 3: Update score item with clinical content
    print("\nUpdating score item clinical content...")
    cur.execute("""
        UPDATE score_items
        SET
            clinical_relevance = %s,
            patient_explanation = %s,
            conduct = %s,
            last_review = %s,
            updated_at = NOW()
        WHERE id = %s
    """, (
        clinical_relevance,
        patient_explanation,
        conduct,
        datetime.now(),
        SCORE_ITEM_ID
    ))
    print("  ✓ Clinical content updated")

    # Step 4: Verify enrichment
    print("\n" + "=" * 80)
    print("VERIFICATION")
    print("=" * 80)

    cur.execute("""
        SELECT
            LENGTH(clinical_relevance) as relevance_len,
            LENGTH(patient_explanation) as explanation_len,
            LENGTH(conduct) as conduct_len,
            last_review
        FROM score_items
        WHERE id = %s
    """, (SCORE_ITEM_ID,))

    result = cur.fetchone()
    print(f"Clinical Relevance: {result[0]} characters")
    print(f"Patient Explanation: {result[1]} characters")
    print(f"Conduct: {result[2]} characters")
    print(f"Last Review: {result[3]}")

    cur.execute("""
        SELECT COUNT(*)
        FROM article_score_items
        WHERE score_item_id = %s
    """, (SCORE_ITEM_ID,))

    article_count = cur.fetchone()[0]
    print(f"Linked Articles: {article_count}")

    # Validate character counts
    if result[0] < 1500 or result[0] > 2000:
        print(f"  ⚠ WARNING: Clinical Relevance should be 1500-2000 chars (got {result[0]})")
    if result[1] < 1000 or result[1] > 1500:
        print(f"  ⚠ WARNING: Patient Explanation should be 1000-1500 chars (got {result[1]})")
    if result[2] < 1500 or result[2] > 2500:
        print(f"  ⚠ WARNING: Conduct should be 1500-2500 chars (got {result[2]})")

    if article_count < 2:
        print(f"  ⚠ WARNING: Should have at least 2 articles (got {article_count})")

    # Commit transaction
    print("\n" + "=" * 80)
    response = input("Commit changes to database? (yes/no): ")

    if response.lower() == 'yes':
        conn.commit()
        print("✓ Changes committed successfully!")

        # Final summary
        print("\n" + "=" * 80)
        print("ENRICHMENT COMPLETE")
        print("=" * 80)
        print(f"Score Item: Leveduras - Sedimento")
        print(f"ID: {SCORE_ITEM_ID}")
        print(f"Articles Added: {len(article_ids)}")
        print(f"Clinical Content: Updated")
        print(f"Last Review: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}")

    else:
        conn.rollback()
        print("✗ Changes rolled back")
        sys.exit(1)

except Exception as e:
    print(f"\n✗ ERROR: {e}")
    conn.rollback()
    raise

finally:
    cur.close()
    conn.close()
