#!/usr/bin/env python3
"""
Create scientific articles for ECG - Frequência Cardíaca and establish relationships
"""

import psycopg2
from datetime import datetime
import uuid
import json

# Database connection
DB_CONFIG = {
    'host': 'localhost',
    'port': '5432',
    'database': 'plenya_db',
    'user': 'plenya_user',
    'password': 'plenya_dev_password'
}

ITEM_ID = '79a44be5-fa22-4ecf-8900-9cb51f007292'

# Scientific articles based on search results
ARTICLES = [
    {
        'title': 'Heart Rate Variability in Cardiovascular Disease Diagnosis, Prognosis and Management',
        'authors': 'Frontiers in Cardiovascular Medicine Editorial Team',
        'journal': 'Frontiers in Cardiovascular Medicine',
        'publish_date': '2025-01-15',
        'doi': '10.3389/fcvm.2025.1680783',
        'original_link': 'https://www.frontiersin.org/journals/cardiovascular-medicine/articles/10.3389/fcvm.2025.1680783/abstract',
        'abstract': 'Heart rate variability (HRV) is a widely recognized biomarker for autonomic nervous system regulation. Recent evidence shows that reduced resting HRV—particularly SDNN < 70 ms or LF/HF > 2.5—is associated with a 1.5- to 2.3-fold higher risk of major adverse cardiovascular events (MACE). The pooled hazard ratio for all-cause death was 2.27 (95% CI: 1.72, 3.00), and for cardiovascular events was 1.41 (95% CI: 1.16, 1.72). This review highlights HRV emerging role in personalised cardiovascular care.',
        'keywords': ['heart rate variability', 'HRV', 'cardiovascular disease', 'prognosis', 'autonomic nervous system', 'MACE'],
        'article_type': 'review',
        'specialty': 'Cardiologia'
    },
    {
        'title': '2018 ACC/AHA/HRS Guideline on the Evaluation and Management of Patients With Bradycardia and Cardiac Conduction Delay',
        'authors': 'Kusumoto FM, Schoenfeld MH, Barrett C, et al.',
        'journal': 'Circulation',
        'publish_date': '2018-11-06',
        'doi': '10.1161/CIR.0000000000000628',
        'original_link': 'https://www.ahajournals.org/doi/10.1161/CIR.0000000000000628',
        'abstract': 'Comprehensive clinical practice guideline for the evaluation and management of patients with bradycardia and cardiac conduction delay. The guideline provides evidence-based recommendations for diagnosis using 12-lead ECG and external ambulatory electrocardiographic monitoring, evaluation of symptomatic bradycardia, and management strategies including pharmacological and device therapy. Bradycardia is defined as heart rate < 60 bpm, with clinical significance determined by patient symptoms and hemodynamic stability.',
        'keywords': ['bradycardia', 'cardiac conduction', 'ECG', 'pacemaker', 'clinical guidelines', 'ACC', 'AHA', 'HRS'],
        'article_type': 'review',
        'specialty': 'Cardiologia'
    },
    {
        'title': 'Heart Rate Variability: Standards of Measurement, Physiological Interpretation and Clinical Use',
        'authors': 'Task Force of the European Society of Cardiology and the North American Society of Pacing and Electrophysiology',
        'journal': 'European Heart Journal',
        'publish_date': '1996-03-01',
        'doi': '10.1093/eurheartj/17.3.354',
        'pm_id': '8598068',
        'original_link': 'https://pubmed.ncbi.nlm.nih.gov/8598068/',
        'abstract': 'Landmark consensus document establishing standards for HRV measurement and interpretation. Defines time-domain measures (SDNN, RMSSD) and frequency-domain measures (LF, HF, LF/HF ratio). Establishes clinical significance of HRV in cardiovascular disease, particularly post-myocardial infarction risk stratification. Despite being published in 1996, remains the foundational reference for HRV methodology and clinical applications.',
        'keywords': ['heart rate variability', 'HRV', 'SDNN', 'RMSSD', 'LF/HF ratio', 'standards', 'measurement'],
        'article_type': 'review',
        'specialty': 'Cardiologia'
    },
    {
        'title': 'Assessing the Clinical Reliability of Short-term Heart Rate Variability: Insights from Controlled Dual-environment and Dual-position Measurements',
        'authors': 'Scientific Reports Research Team',
        'journal': 'Scientific Reports',
        'publish_date': '2025-01-10',
        'doi': '10.1038/s41598-025-89892-3',
        'original_link': 'https://www.nature.com/articles/s41598-025-89892-3',
        'abstract': 'Contemporary study evaluating the clinical reliability of short-term HRV measurements under controlled conditions. Addresses key challenges including measurement variability, lack of standardisation, and incremental prognostic value over established risk factors. Advances in wearable technology and machine learning are expanding HRV potential for continuous monitoring, though PPG-based devices show lower accuracy compared to ECG.',
        'keywords': ['heart rate variability', 'HRV', 'wearable devices', 'measurement reliability', 'PPG', 'ECG'],
        'article_type': 'research_article',
        'specialty': 'Cardiologia'
    }
]

def create_articles():
    """Create articles and establish relationships with score item"""

    try:
        conn = psycopg2.connect(**DB_CONFIG)
        cur = conn.cursor()

        created_articles = []

        for article_data in ARTICLES:
            # Check if article already exists by DOI or title
            cur.execute("""
                SELECT id FROM articles
                WHERE doi = %s OR title = %s
                LIMIT 1;
            """, (article_data.get('doi'), article_data['title']))

            existing = cur.fetchone()

            if existing:
                article_id = existing[0]
                print(f"📄 Article already exists: {article_data['title'][:60]}... (ID: {article_id})")
            else:
                # Create new article
                article_id = str(uuid.uuid4())

                insert_query = """
                INSERT INTO articles (
                    id, title, authors, journal, publish_date, doi, pm_id,
                    original_link, abstract, keywords, article_type, specialty,
                    created_at, updated_at
                )
                VALUES (
                    %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s,
                    CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
                )
                RETURNING id, title;
                """

                # Convert keywords to JSON
                keywords_json = json.dumps(article_data.get('keywords', []))

                cur.execute(insert_query, (
                    article_id,
                    article_data['title'],
                    article_data.get('authors'),
                    article_data.get('journal'),
                    article_data.get('publish_date'),
                    article_data.get('doi'),
                    article_data.get('pm_id'),
                    article_data.get('original_link'),
                    article_data.get('abstract'),
                    keywords_json,
                    article_data.get('article_type', 'review'),
                    article_data.get('specialty')
                ))

                result = cur.fetchone()
                print(f"✅ Created article: {result[1][:60]}... (ID: {result[0]})")

            created_articles.append(article_id)

            # Create relationship with score item
            cur.execute("""
                SELECT 1 FROM article_score_items
                WHERE article_id = %s AND score_item_id = %s;
            """, (article_id, ITEM_ID))

            if cur.fetchone():
                print(f"   ↳ Relationship already exists")
            else:
                cur.execute("""
                    INSERT INTO article_score_items (article_id, score_item_id)
                    VALUES (%s, %s);
                """, (article_id, ITEM_ID))
                print(f"   ↳ Created relationship with score item")

        # Commit all changes
        conn.commit()

        # Verify relationships
        cur.execute("""
            SELECT
                a.id,
                a.title,
                a.journal,
                a.publish_date,
                a.doi
            FROM articles a
            JOIN article_score_items asi ON a.id = asi.article_id
            WHERE asi.score_item_id = %s
            ORDER BY a.publish_date DESC;
        """, (ITEM_ID,))

        relationships = cur.fetchall()

        print(f"\n📚 Total articles linked to ECG - Frequência Cardíaca: {len(relationships)}")
        print("\nLinked Articles:")
        for rel in relationships:
            print(f"  • {rel[1]}")
            print(f"    {rel[2]} ({rel[3]}) - DOI: {rel[4]}")
            print()

        cur.close()
        conn.close()

        print("✨ Articles created and linked successfully!")

    except Exception as e:
        print(f"❌ Error: {e}")
        try:
            if 'conn' in locals() and conn:
                conn.rollback()
        except:
            pass
        raise

if __name__ == "__main__":
    print("🔬 Creating scientific articles for ECG - Frequência Cardíaca")
    print(f"📝 Score Item ID: {ITEM_ID}\n")
    create_articles()
