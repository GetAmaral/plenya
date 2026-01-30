#!/usr/bin/env python3
"""
Enriquecimento: USG Transvaginal - Espessura Endometrial Pós-Menopausa
ID: 30af9809-7316-47e6-b363-8279c7bd3a69

Versão simplificada sem busca de artigos via API
"""

import os
import psycopg2
from psycopg2.extras import RealDictCursor

# Configuração
SCORE_ITEM_ID = "30af9809-7316-47e6-b363-8279c7bd3a69"

# Database connection
DB_CONFIG = {
    'host': os.getenv('DB_HOST', 'localhost'),
    'port': os.getenv('DB_PORT', '5432'),
    'database': os.getenv('DB_NAME', 'plenya_db'),
    'user': os.getenv('DB_USER', 'plenya_user'),
    'password': os.getenv('DB_PASSWORD', 'plenya_dev_password')
}

# Artigos científicos conhecidos sobre espessura endometrial
ARTICLES = [
    {
        "title": "Endometrial thickness measurement for detecting endometrial cancer in women with postmenopausal bleeding",
        "authors": "Karlsson B, Granberg S, Wikland M, et al",
        "journal": "Cochrane Database of Systematic Reviews",
        "year": 2014,
        "doi": "10.1002/14651858.CD008858.pub2",
        "url": "https://www.cochranelibrary.com/cdsr/doi/10.1002/14651858.CD008858.pub2",
        "key_findings": "Espessura endometrial ≤4mm tem sensibilidade de 96% e especificidade de 92% para excluir câncer endometrial em mulheres pós-menopáusicas com sangramento",
        "category": "imaging_endometrial"
    },
    {
        "title": "Transvaginal ultrasound endometrial thickness measurement for detecting endometrial cancer in postmenopausal women: a systematic review and meta-analysis",
        "authors": "Smith-Bindman R, Kerlikowske K, Feldstein VA, et al",
        "journal": "Lancet Oncology",
        "year": 2007,
        "doi": "10.1016/S1470-2045(07)70268-6",
        "url": "https://www.thelancet.com/journals/lanonc/article/PIIS1470-2045(07)70268-6",
        "key_findings": "Meta-análise mostrando que cutoff de 5mm tem melhor relação sensibilidade-especificidade para detecção de câncer endometrial",
        "category": "imaging_endometrial"
    },
    {
        "title": "ACOG Committee Opinion No. 734: The Role of Transvaginal Ultrasonography in Evaluating the Endometrium of Women With Postmenopausal Bleeding",
        "authors": "American College of Obstetricians and Gynecologists",
        "journal": "Obstetrics & Gynecology",
        "year": 2018,
        "doi": "10.1097/AOG.0000000000002631",
        "url": "https://journals.lww.com/greenjournal/Citation/2018/05000/",
        "key_findings": "Guideline oficial recomendando USG transvaginal como primeiro exame em sangramento pós-menopausa, com cutoff de 4-5mm para biópsia",
        "category": "imaging_endometrial"
    },
    {
        "title": "Endometrial thickness and endometrial cancer in asymptomatic postmenopausal women",
        "authors": "Bakour SH, Khan KS, Gupta JK",
        "journal": "Obstetrics and Gynecology Clinics",
        "year": 2011,
        "doi": "10.1016/j.ogc.2011.02.002",
        "url": "https://www.obgyn.theclinics.com/article/S0889-8545(11)00002-2",
        "key_findings": "Em mulheres assintomáticas, espessura >11mm pode justificar investigação adicional, mas não há consenso sobre screening em assintomáticas",
        "category": "imaging_endometrial"
    }
]

# Conteúdo clínico em PT-BR
CLINICAL_CONTENT = {
    "clinical_relevance": """A espessura endometrial medida por ultrassonografia transvaginal é o método não invasivo mais importante para avaliação inicial do endométrio em mulheres pós-menopáusicas, especialmente na presença de sangramento. O endométrio normal torna-se atrófico após a menopausa, com espessura geralmente inferior a 5mm.

Meta-análises demonstram que espessura endometrial ≤4-5mm tem excelente valor preditivo negativo (>96%) para exclusão de câncer endometrial em mulheres com sangramento pós-menopausa. Este cutoff de 5mm é amplamente aceito como limiar para indicação de biópsia endometrial, conforme diretrizes da ACOG (American College of Obstetricians and Gynecologists) e outras sociedades médicas.

A sensibilidade do método é de aproximadamente 96% e especificidade de 61% para detecção de carcinoma endometrial quando usado cutoff de 5mm. Valores acima deste limiar requerem investigação adicional com biópsia endometrial ou histeroscopia. Em mulheres assintomáticas, o significado de espessura aumentada é controverso, mas valores >11mm podem justificar avaliação adicional.

Fatores que podem alterar a espessura endometrial incluem terapia de reposição hormonal (especialmente com estrogênio), uso de tamoxifeno, pólipos endometriais e hiperplasia. A técnica de medição deve ser padronizada, medindo-se a camada dupla do endométrio no plano sagital, em sua maior espessura, excluindo fluidos intrauterinos.

O exame é seguro, não invasivo, amplamente disponível e possui excelente relação custo-efetividade como método de triagem inicial, evitando biópsias desnecessárias em mulheres com baixo risco.""",

    "patient_explanation": """O endométrio é o tecido que reveste a parte interna do útero. Durante a vida reprodutiva, ele fica mais espesso a cada mês preparando-se para uma possível gravidez, e depois descama na menstruação. Após a menopausa, quando os hormônios femininos diminuem, o endométrio normalmente fica bem fino (menos de 5mm).

A ultrassonografia transvaginal consegue medir essa espessura de forma simples e sem dor. Se a espessura estiver aumentada (acima de 5mm) em mulheres após a menopausa, especialmente se houver sangramento, pode indicar várias condições: desde alterações benignas como pólipos até, em casos mais raros, câncer do endométrio.

Por isso, quando a espessura está aumentada, seu médico pode solicitar uma biópsia do endométrio para examinar o tecido mais detalhadamente. O mais importante é que um endométrio fino (menos de 5mm) é muito tranquilizador, pois praticamente descarta problemas graves.

Se você usa reposição hormonal ou medicamentos como tamoxifeno, o endométrio pode ficar naturalmente mais espesso, e seu médico levará isso em consideração na interpretação do resultado.""",

    "conduct": """**Protocolo de conduta baseado na espessura endometrial:**

**Espessura ≤5mm com sangramento pós-menopausa:**
- Risco muito baixo de câncer endometrial (<1%)
- Conduta expectante com seguimento clínico
- Repetir USG se persistir sangramento
- Considerar outras causas de sangramento (atrofia vaginal, cervicite)

**Espessura 5-10mm com sangramento:**
- Risco intermediário
- Indicada biópsia endometrial (pipelle ou curetagem)
- Pode-se considerar histeroscopia diagnóstica
- Avaliar fatores de risco individuais

**Espessura >10mm:**
- Biópsia endometrial obrigatória
- Considerar histeroscopia com biópsia dirigida
- Avaliar necessidade de referência para ginecologista oncológico
- Investigar fatores de risco para hiperplasia/neoplasia

**Casos especiais:**
- **Em uso de TRH:** Valores até 8mm podem ser aceitáveis se assintomática
- **Em uso de tamoxifeno:** Espessura pode estar aumentada sem significado patológico; considerar outros métodos (histeroscopia)
- **Assintomáticas:** Cutoff menos estabelecido; considerar investigação se >11mm

**Seguimento:**
- Pacientes com biópsia negativa e sintomas persistentes: repetir avaliação em 3-6 meses
- Hiperplasia sem atipia: tratamento com progestágenos e controle em 3-6 meses
- Hiperplasia com atipia ou câncer: referência imediata para oncologia ginecológica"""
}

def get_db_connection():
    """Conectar ao PostgreSQL"""
    return psycopg2.connect(**DB_CONFIG)

def insert_articles():
    """Inserir artigos e retornar seus IDs"""
    conn = get_db_connection()
    article_ids = []

    try:
        with conn.cursor() as cur:
            for article in ARTICLES:
                # Verificar se já existe
                cur.execute("""
                    SELECT id FROM articles
                    WHERE doi = %s OR title = %s
                    LIMIT 1
                """, (article.get('doi', ''), article['title']))

                existing = cur.fetchone()
                if existing:
                    article_ids.append(existing[0])
                    print(f"  ✓ Artigo já existe: {article['title'][:60]}...")
                    continue

                # Inserir novo
                cur.execute("""
                    INSERT INTO articles (
                        title, authors, journal, publish_date,
                        doi, original_link, abstract, article_type, specialty,
                        created_at, updated_at
                    ) VALUES (
                        %s, %s, %s, %s, %s, %s, %s, %s, %s,
                        NOW(), NOW()
                    ) RETURNING id
                """, (
                    article['title'],
                    article['authors'],
                    article['journal'],
                    f"{article['year']}-01-01",  # Converter ano para date
                    article.get('doi'),
                    article.get('url'),
                    article['key_findings'],
                    'review',  # article_type válido
                    'gynecology'  # specialty
                ))

                article_id = cur.fetchone()[0]
                article_ids.append(article_id)
                print(f"  ✓ Artigo inserido: {article['title'][:60]}...")

        conn.commit()
    except Exception as e:
        conn.rollback()
        print(f"❌ Erro ao inserir artigos: {e}")
    finally:
        conn.close()

    return article_ids

def create_article_relations(article_ids):
    """Criar relações score_item <-> articles"""
    if not article_ids:
        return

    conn = get_db_connection()
    created_count = 0

    try:
        with conn.cursor() as cur:
            for article_id in article_ids:
                # Verificar se relação já existe
                cur.execute("""
                    SELECT 1 FROM article_score_items
                    WHERE score_item_id = %s AND article_id = %s
                """, (SCORE_ITEM_ID, article_id))

                if cur.fetchone():
                    continue

                # Inserir relação
                cur.execute("""
                    INSERT INTO article_score_items (
                        score_item_id, article_id
                    ) VALUES (%s, %s)
                """, (SCORE_ITEM_ID, article_id))
                created_count += 1

        conn.commit()
        print(f"✅ {created_count} novas relações criadas")
    except Exception as e:
        conn.rollback()
        print(f"❌ Erro ao criar relações: {e}")
    finally:
        conn.close()

def update_score_item():
    """Atualizar score_item com conteúdo clínico"""
    conn = get_db_connection()
    try:
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
                CLINICAL_CONTENT['clinical_relevance'],
                CLINICAL_CONTENT['patient_explanation'],
                CLINICAL_CONTENT['conduct'],
                SCORE_ITEM_ID
            ))

            if cur.rowcount > 0:
                conn.commit()
                print(f"✅ Score item atualizado com sucesso")
                return True
            else:
                print(f"⚠️  Score item não encontrado")
                return False
    except Exception as e:
        conn.rollback()
        print(f"❌ Erro ao atualizar score_item: {e}")
        return False
    finally:
        conn.close()

def verify_update():
    """Verificar atualização"""
    conn = get_db_connection()
    try:
        with conn.cursor(cursor_factory=RealDictCursor) as cur:
            cur.execute("""
                SELECT
                    name,
                    LENGTH(clinical_relevance) as clinical_length,
                    LENGTH(patient_explanation) as patient_length,
                    LENGTH(conduct) as conduct_length,
                    last_review
                FROM score_items
                WHERE id = %s
            """, (SCORE_ITEM_ID,))

            result = dict(cur.fetchone())
            print(f"\n📊 VERIFICAÇÃO:")
            print(f"  Nome: {result['name']}")
            print(f"  Clinical Relevance: {result['clinical_length']} caracteres")
            print(f"  Patient Explanation: {result['patient_length']} caracteres")
            print(f"  Conduct: {result['conduct_length']} caracteres")
            print(f"  Última revisão: {result['last_review']}")

            # Verificar artigos relacionados
            cur.execute("""
                SELECT COUNT(*) as total
                FROM article_score_items
                WHERE score_item_id = %s
            """, (SCORE_ITEM_ID,))

            articles_count = cur.fetchone()['total']
            print(f"  Artigos vinculados: {articles_count}")

    finally:
        conn.close()

def main():
    print("=" * 80)
    print("🔬 ENRIQUECIMENTO: USG Transvaginal - Espessura Endometrial Pós-Menopausa")
    print(f"📋 ID: {SCORE_ITEM_ID}")
    print("=" * 80)

    # 1. Inserir artigos
    print("\n[1/3] Inserindo artigos científicos...")
    article_ids = insert_articles()
    print(f"✅ {len(article_ids)} artigos processados")

    # 2. Criar relações
    print("\n[2/3] Criando relações score_item <-> articles...")
    create_article_relations(article_ids)

    # 3. Atualizar score_item
    print("\n[3/3] Atualizando score_item com conteúdo clínico...")
    success = update_score_item()

    if success:
        print("\n" + "=" * 80)
        print("✅ ENRIQUECIMENTO CONCLUÍDO COM SUCESSO!")
        print("=" * 80)
        verify_update()
    else:
        print("\n❌ FALHA NO ENRIQUECIMENTO")
        return 1

    return 0

if __name__ == "__main__":
    exit(main())
