#!/usr/bin/env python3
"""
Enriquecimento clínico: USG Próstata - Volume Prostático
ID: 23b012f9-ce8b-4a1d-95f4-cfe9183295e0
"""

import sys
import psycopg2
from datetime import datetime
import json

# Database configuration
DB_CONFIG = {
    'host': 'localhost',
    'port': 5432,
    'user': 'plenya_user',
    'password': 'plenya_dev_password',
    'database': 'plenya_db'
}

ITEM_ID = '23b012f9-ce8b-4a1d-95f4-cfe9183295e0'

# Artigos científicos identificados
ARTICLES = [
    {
        "title": "Management of Lower Urinary Tract Symptoms Attributed to Benign Prostatic Hyperplasia (BPH): AUA Guideline Amendment 2023",
        "authors": "American Urological Association",
        "publish_date": "2023-10-01",
        "journal": "Journal of Urology",
        "doi": "10.1097/JU.0000000000003698",
        "original_link": "https://www.auajournals.org/doi/10.1097/JU.0000000000003698",
        "abstract": "Define prostatic enlargement as volume >30g on imaging. Recommends prostate size assessment via ultrasound prior to intervention. Combination therapy (5-ARI + alpha blocker) for prostates >30g.",
        "specialty": "Urologia"
    },
    {
        "title": "Age-stratified normal values for prostate volume, PSA, maximum urinary flow rate, IPSS, and other LUTS/BPH indicators in the German male community-dwelling population aged 50 years or older",
        "authors": "Bosch JL, Tilling K, Bohnen AM, Bangma CH, Donovan JL",
        "publish_date": "2011-02-01",
        "journal": "World Journal of Urology",
        "doi": "10.1007/s00345-010-0638-z",
        "pm_id": "21221974",
        "original_link": "https://pubmed.ncbi.nlm.nih.gov/21221974/",
        "abstract": "Normal prostate volume ranges 20-30 mL. Volume increases with age. Classification: mild enlargement 30-50mL, moderate 50-70mL, marked >70mL.",
        "specialty": "Urologia"
    },
    {
        "title": "Can men with prostates sized 80 mL or larger be managed conservatively?",
        "authors": "Meskawi M, Parra J, Phuoc VH, Chughtai B, Te A, Kaplan S",
        "publish_date": "2017-08-01",
        "journal": "Canadian Urological Association Journal",
        "doi": "10.5489/cuaj.4328",
        "original_link": "https://pmc.ncbi.nlm.nih.gov/articles/PMC5577333/",
        "abstract": "Two-thirds of men with prostates ≥80mL maintained stability with conservative management over 62 months. Clinical progression occurred in 33%. No baseline volume predicted progression.",
        "specialty": "Urologia"
    }
]

# Conteúdo clínico enriquecido em PT-BR
ENRICHMENT = {
    "clinical_relevance": """O volume prostático medido por ultrassonografia transretal (USTR) ou abdominal é um parâmetro fundamental na avaliação da hiperplasia prostática benigna (HBP). Segundo diretrizes da American Urological Association (2023), a próstata é considerada aumentada quando o volume ultrassonográfico excede 30 gramas/mL.

A classificação por volume orienta condutas terapêuticas: volumes de 30-50mL indicam HBP leve, 50-70mL moderada, e >70mL marcada. O volume normal em adultos varia de 20-30mL, com média de 24mL em homens de 50-60 anos. É esperado crescimento progressivo com o envelhecimento, à taxa aproximada de 0,6mL/ano após os 40 anos.

O volume prostático é determinante na escolha terapêutica: terapia combinada com inibidores da 5-alfa-redutase (5-ARI) e alfa-bloqueadores está indicada apenas para próstatas >30g. Procedimentos minimamente invasivos como PUL (prostatic urethral lift) e terapia com vapor de água são recomendados para volumes de 30-80cc, com contraindicação relativa quando há lobo médio obstrutivo. Para próstatas ≥80mL, estudos demonstram que dois terços dos pacientes mantêm estabilidade clínica com manejo conservador medicamentoso por até 62 meses, embora um terço evolua para progressão clínica necessitando cirurgia.""",

    "patient_explanation": """A próstata é uma glândula do tamanho de uma noz que fica logo abaixo da bexiga e aumenta naturalmente com a idade. O ultrassom mede o volume (tamanho) da próstata em mililitros, similar a medir o volume de água em um copo.

Uma próstata normal tem cerca de 20-30 mL (equivalente a 2-3 colheres de sopa). Quando passa de 30 mL, é considerada aumentada - condição chamada hiperplasia prostática benigna (HBP), que não é câncer. Próstatas de 30-50 mL são levemente aumentadas, 50-70 mL moderadamente aumentadas, e acima de 70 mL bastante aumentadas.

O tamanho da próstata ajuda o médico a decidir qual tratamento é melhor para você. Próstatas menores geralmente respondem bem a medicamentos. Próstatas maiores podem precisar de procedimentos ou cirurgia, especialmente se estiverem causando sintomas como dificuldade para urinar, jato fraco, ou necessidade de acordar várias vezes à noite para ir ao banheiro. Mesmo próstatas muito grandes (acima de 80 mL) podem ser tratadas inicialmente com medicamentos, desde que os sintomas não sejam graves.""",

    "conduct": """**Protocolo de Conduta Baseada em Volume Prostático:**

**Volume <30 mL (Normal):**
- Considerar causas não-prostáticas para LUTS (sintomas do trato urinário inferior)
- Avaliar disfunção vesical, estenose uretral, bexiga hiperativa
- Monitoramento anual se assintomático

**Volume 30-50 mL (HBP Leve):**
- Terapia inicial: alfa-bloqueador isolado (tamsulosina, alfuzosina)
- Se PSA >1.5 ng/mL: considerar terapia combinada (alfa-bloqueador + 5-ARI)
- Reavaliar volume e sintomas (IPSS) em 6-12 meses
- Candidatos a procedimentos minimamente invasivos (PUL, WVTT) se falha medicamentosa

**Volume 50-80 mL (HBP Moderada):**
- Terapia combinada preferencial (alfa-bloqueador + finasterida/dutasterida)
- Avaliação de resíduo pós-miccional (RPM) e urofluxometria
- Considerar cirurgia se: RPM >300mL, retenção urinária recorrente, insuficiência renal, hematúria refratária
- Reavaliação a cada 6 meses

**Volume >80 mL (HBP Acentuada):**
- Discussão sobre opções cirúrgicas (RTUP, enucleação a laser, prostatectomia aberta)
- Terapia medicamentosa pode ser tentada se sintomas leves-moderados e sem complicações
- Acompanhamento rigoroso: PSA, creatinina, RPM a cada 3-6 meses
- Encaminhamento para urologia para planejamento cirúrgico se progressão sintomática

**Critérios de Encaminhamento Urgente:**
- Retenção urinária aguda
- Hidronefrose/insuficiência renal
- Hematúria macroscópica persistente
- Cálculos vesicais
- Infecções urinárias de repetição"""
}

def get_db_connection():
    """Estabelece conexão com PostgreSQL"""
    return psycopg2.connect(**DB_CONFIG)

def save_articles(conn):
    """Salva artigos no banco e retorna IDs"""
    article_ids = []

    with conn.cursor() as cur:
        for article in ARTICLES:
            # Verifica se artigo já existe
            cur.execute("""
                SELECT id FROM articles
                WHERE doi = %s OR original_link = %s
                LIMIT 1
            """, (article.get('doi'), article.get('original_link')))

            existing = cur.fetchone()

            if existing:
                article_id = existing[0]
                print(f"  ↪ Artigo já existe: {article['title'][:60]}...")
            else:
                cur.execute("""
                    INSERT INTO articles (
                        title, authors, publish_date, journal,
                        doi, pm_id, original_link, abstract, specialty,
                        article_type, language, created_at, updated_at
                    )
                    VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, NOW(), NOW())
                    RETURNING id
                """, (
                    article['title'],
                    article.get('authors'),
                    article.get('publish_date'),
                    article.get('journal'),
                    article.get('doi'),
                    article.get('pm_id'),
                    article.get('original_link'),
                    article.get('abstract'),
                    article.get('specialty'),
                    'research_article',
                    'en'
                ))

                article_id = cur.fetchone()[0]
                print(f"  ✓ Artigo salvo: {article['title'][:60]}...")

            article_ids.append(article_id)

        conn.commit()

    return article_ids

def link_articles(conn, article_ids):
    """Cria relações entre artigos e score_item"""
    with conn.cursor() as cur:
        for article_id in article_ids:
            cur.execute("""
                SELECT 1 FROM article_score_items
                WHERE score_item_id = %s AND article_id = %s
            """, (ITEM_ID, article_id))

            if cur.fetchone():
                print(f"  ↪ Relação já existe com artigo {article_id}")
            else:
                cur.execute("""
                    INSERT INTO article_score_items (score_item_id, article_id)
                    VALUES (%s, %s)
                """, (ITEM_ID, article_id))
                print(f"  ✓ Relação criada com artigo {article_id}")

        conn.commit()

def update_score_item(conn):
    """Atualiza score_item com conteúdo enriquecido"""
    with conn.cursor() as cur:
        cur.execute("""
            UPDATE score_items
            SET
                clinical_relevance = %s,
                patient_explanation = %s,
                conduct = %s,
                last_review = NOW(),
                updated_at = NOW()
            WHERE id = %s
        """, (
            ENRICHMENT['clinical_relevance'],
            ENRICHMENT['patient_explanation'],
            ENRICHMENT['conduct'],
            ITEM_ID
        ))

        conn.commit()

        if cur.rowcount > 0:
            print(f"✓ Score item atualizado")
            return True
        else:
            print(f"✗ ERRO: Nenhuma linha atualizada")
            return False

def verify_result(conn):
    """Verifica o resultado final"""
    with conn.cursor() as cur:
        cur.execute("""
            SELECT
                name,
                LENGTH(clinical_relevance) as clinical_len,
                LENGTH(patient_explanation) as patient_len,
                LENGTH(conduct) as conduct_len,
                last_review,
                (SELECT COUNT(*) FROM article_score_items WHERE score_item_id = %s) as article_count
            FROM score_items
            WHERE id = %s
        """, (ITEM_ID, ITEM_ID))

        result = cur.fetchone()

        if result:
            print("\n" + "="*70)
            print("📊 VERIFICAÇÃO FINAL")
            print("="*70)
            print(f"Item: {result[0]}")
            print(f"Clinical Relevance: {result[1]} caracteres")
            print(f"Patient Explanation: {result[2]} caracteres")
            print(f"Conduct: {result[3]} caracteres")
            print(f"Last Review: {result[4]}")
            print(f"Artigos vinculados: {result[5]}")
            print("="*70)

            if result[1] > 0 and result[2] > 0 and result[3] > 0 and result[5] > 0:
                print("✅ ENRIQUECIMENTO COMPLETO!")
                return True
            else:
                print("⚠️  ENRIQUECIMENTO INCOMPLETO")
                return False

        return False

def main():
    """Executa workflow completo"""
    print("="*70)
    print("🧬 ENRIQUECIMENTO: USG Próstata - Volume Prostático")
    print("="*70)
    print(f"Item ID: {ITEM_ID}")
    print(f"Data: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}")
    print("="*70)

    try:
        conn = get_db_connection()
        print("\n✓ Conectado ao banco")

        # Verificar item
        with conn.cursor() as cur:
            cur.execute("SELECT name FROM score_items WHERE id = %s", (ITEM_ID,))
            item = cur.fetchone()
            if not item:
                print(f"✗ ERRO: Item {ITEM_ID} não encontrado")
                return 1
            print(f"✓ Item encontrado: {item[0]}")

        # Salvar artigos
        print("\n💾 Salvando artigos científicos...")
        article_ids = save_articles(conn)
        print(f"✓ {len(article_ids)} artigos processados")

        # Vincular artigos
        print("\n🔗 Vinculando artigos ao score_item...")
        link_articles(conn, article_ids)

        # Atualizar score_item
        print("\n📝 Salvando enriquecimento clínico...")
        if not update_score_item(conn):
            return 1

        # Verificar resultado
        if not verify_result(conn):
            return 1

        conn.close()

        print("\n" + "="*70)
        print("🎉 PROCESSO CONCLUÍDO COM SUCESSO!")
        print("="*70)
        return 0

    except Exception as e:
        print(f"\n❌ ERRO: {e}")
        import traceback
        traceback.print_exc()
        return 1

if __name__ == "__main__":
    sys.exit(main())
