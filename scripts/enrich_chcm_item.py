#!/usr/bin/env python3
"""
Script para enriquecer o item CHCM (MCHC) com conteúdo clínico baseado em evidências
ID: f1a1d106-b5ac-4483-9de9-ee90ae103d26
"""

import psycopg2
import os
from datetime import datetime

# Conteúdo científico baseado nas pesquisas (2023-2024)
CHCM_DATA = {
    "clinical_relevance": """A Concentração de Hemoglobina Corpuscular Média (CHCM/MCHC) é um índice hematimétrico que mede a concentração média de hemoglobina dentro de um determinado volume de hemácias. Diferentemente do HCM, que correlaciona a massa absoluta de hemoglobina, o CHCM avalia a concentração relativa, sendo calculado pela divisão da hemoglobina pelo hematócrito (Hb/Ht × 100).

Valores de referência: 31-36 g/dL (adultos)

**CHCM Baixo (Hipocromia - <31 g/dL):**
Indica hemácias pálidas com menor concentração de hemoglobina. Ocorre principalmente em:
- Anemia ferropriva (deficiência de ferro) - causa mais comum
- Talassemias (trait talassêmico)
- Anemia de doença crônica
- Anemia sideroblástica

O CHCM baixo reflete uma síntese inadequada de hemoglobina que excede a redução no tamanho das hemácias, resultando em células hipocrômicas ao microscópio.

**CHCM Alto (Hipercromia - >36 g/dL):**
Valores elevados (36-37 g/dL) podem indicar:
- Esferocitose hereditária (causa mais comum de CHCM elevado)
- Xerocitose
- Hemoglobinopatias (doença falciforme, hemoglobina C)

**IMPORTANTE:** CHCM >36 g/dL pode representar artefato laboratorial em situações como lipemia (eleva Hb artificialmente), hemólise da amostra, ou aglutinação de hemácias (reduz artificialmente a contagem de células e eleva o VCM).

**Diagnóstico de Esferocitose Hereditária:**
Estudos de 2024 demonstram que MCHC >35,5 g/dL (355 g/L) é um dos primeiros parâmetros utilizados no rastreamento de esferocitose hereditária (HS). Analisadores automatizados modernos (Advia) fornecem o percentual de células hiperdensas (% Hyper), que reflete hipercronia das hemácias e auxilia na detecção. O diagnóstico é confirmado com Coombs negativo, história familiar positiva, esferócitos ao microscópio e testes especializados (fragilidade osmótica, EMA, AGLT).

**Importante:** Percentuais normais de hemácias hipercrômicas podem ocorrer durante fases de anemia e/ou deficiência de ferro mesmo em pacientes com HS, não excluindo o diagnóstico.""",

    "patient_explanation": """**O que é CHCM?**

CHCM significa Concentração de Hemoglobina Corpuscular Média. É um exame que mostra quanto de hemoglobina (proteína que transporta oxigênio) está "empacotado" dentro de cada glóbulo vermelho do seu sangue.

**Valor normal:** 31 a 36 g/dL

**Por que é importante?**
Ajuda a identificar diferentes tipos de anemia e problemas no sangue.

**CHCM baixo (abaixo de 31):**
Significa que suas hemácias estão "pálidas" e têm pouca hemoglobina. As causas mais comuns são:
- Falta de ferro (anemia por deficiência de ferro)
- Problemas genéticos na produção de hemoglobina (talassemia)

**Sintomas possíveis:** cansaço, fraqueza, palidez, falta de ar, dificuldade de concentração.

**CHCM alto (acima de 36):**
Significa que suas hemácias têm muita hemoglobina concentrada. Pode indicar:
- Esferocitose hereditária (doença genética que muda o formato das hemácias)
- Às vezes pode ser um erro do laboratório

**Sintomas possíveis (esferocitose):** icterícia (pele amarelada), anemia leve, aumento do baço.

**O que fazer?**
Se o seu CHCM estiver alterado, seu médico vai solicitar outros exames para entender a causa. Cada situação tem um tratamento específico - desde suplementação de ferro até acompanhamento especializado.""",

    "conduct": """**Interpretação Inicial:**
1. Avaliar CHCM em conjunto com VCM, HCM, RDW e morfologia eritrocitária
2. Correlacionar com história clínica, sintomas e exame físico
3. Verificar qualidade da amostra (lipemia, hemólise podem gerar artefatos)

**CHCM Baixo (<31 g/dL) - Hipocromia:**

**Investigação:**
- Ferritina sérica, ferro sérico, saturação de transferrina (avaliar deficiência de ferro)
- Hemograma completo com reticulócitos
- Esfregaço de sangue periférico (confirmar hipocromia, microcitose)
- Eletroforese de hemoglobina (se suspeita de talassemia ou hemoglobinopatia)
- PCR, VHS (avaliar inflamação crônica)

**Conduta:**
- **Anemia ferropriva:** Sulfato ferroso VO 200-300mg/dia de ferro elementar por 3-6 meses. Investigar causa do sangramento (endoscopia digestiva se indicado).
- **Talassemia minor:** Aconselhamento genético, evitar suplementação desnecessária de ferro.
- **Anemia de doença crônica:** Tratar condição de base.
- Reavaliação em 4-6 semanas após início do tratamento.

**CHCM Alto (>36 g/dL) - Hipercromia:**

**Investigação:**
1. **Afastar artefatos:** Repetir amostra, verificar lipemia, hemólise, aglutinação
2. **Se CHCM persistentemente elevado:**
   - Esfregaço de sangue periférico (pesquisar esferócitos)
   - Teste de Coombs direto (excluir anemia hemolítica autoimune)
   - História familiar de anemia, icterícia, esplenomegalia
   - Bilirrubinas (avaliar hemólise)
   - Reticulócitos (elevados em hemólise)

3. **Se suspeita de Esferocitose Hereditária:**
   - Teste de fragilidade osmótica (incubada)
   - Citometria de fluxo com eosin-5-maleimide (EMA) - teste de escolha atual
   - Teste de lise em glicerol acidificado (AGLT)
   - Ultrassom abdominal (avaliar esplenomegalia, cálculos biliares)

**Conduta Esferocitose:**
- **Leve (Hb >11 g/dL):** Acompanhamento clínico, suplementação de ácido fólico 1mg/dia
- **Moderada a grave:** Encaminhar para hematologia, considerar esplenectomia após 6 anos de idade
- Monitorar crises aplásticas (parvovírus B19)
- Rastreamento de cálculos biliares

**Encaminhamento para Hematologia:**
- CHCM persistentemente alto sem causa identificada
- Suspeita confirmada de esferocitose hereditária
- Hemólise inexplicada
- Hemoglobinopatias

**Seguimento:**
- Repetir hemograma em 1-3 meses conforme gravidade
- Documentar resposta terapêutica (ferritina, hemoglobina)
- Avaliar necessidade de investigação familiar em casos hereditários"""
}

# Artigos científicos identificados (2023-2024)
ARTICLES = [
    {
        "title": "Recommendations for diagnosis, treatment, and prevention of iron deficiency and iron deficiency anemia",
        "url": "https://pmc.ncbi.nlm.nih.gov/articles/PMC11247274/",
        "year": 2024,
        "journal": "HemaSphere",
        "type": "clinical_guideline"
    },
    {
        "title": "Overview on Hereditary Spherocytosis Diagnosis",
        "url": "https://pubmed.ncbi.nlm.nih.gov/39467036/",
        "year": 2024,
        "journal": "International Journal of Laboratory Hematology",
        "type": "review"
    },
    {
        "title": "Red Cell Indices - Clinical Methods",
        "url": "https://www.ncbi.nlm.nih.gov/books/NBK260/",
        "year": 2023,
        "journal": "NCBI Bookshelf",
        "type": "reference"
    },
    {
        "title": "Normal and Abnormal Complete Blood Count With Differential",
        "url": "https://www.ncbi.nlm.nih.gov/books/NBK604207/",
        "year": 2024,
        "journal": "StatPearls",
        "type": "reference"
    },
    {
        "title": "Hereditary Spherocytosis - Clinical Features and Diagnosis",
        "url": "https://www.ncbi.nlm.nih.gov/books/NBK539797/",
        "year": 2024,
        "journal": "StatPearls",
        "type": "reference"
    }
]

def get_db_connection():
    """Cria conexão com o banco de dados"""
    return psycopg2.connect(
        host=os.getenv('DB_HOST', 'localhost'),
        port=os.getenv('DB_PORT', '5432'),
        database=os.getenv('DB_NAME', 'plenya_db'),
        user=os.getenv('DB_USER', 'plenya_user'),
        password=os.getenv('DB_PASSWORD', 'senha_segura')
    )

def insert_or_get_article(cursor, article):
    """Insere ou recupera artigo científico"""
    # Verificar se artigo já existe
    cursor.execute("""
        SELECT id FROM articles
        WHERE url = %s OR title = %s
        LIMIT 1
    """, (article['url'], article['title']))

    result = cursor.fetchone()
    if result:
        print(f"  ✓ Artigo já existe: {article['title'][:60]}...")
        return result[0]

    # Inserir novo artigo
    cursor.execute("""
        INSERT INTO articles (
            title, url, publication_year, journal, article_type,
            created_at, updated_at
        )
        VALUES (%s, %s, %s, %s, %s, NOW(), NOW())
        RETURNING id
    """, (
        article['title'],
        article['url'],
        article['year'],
        article['journal'],
        article['type']
    ))

    article_id = cursor.fetchone()[0]
    print(f"  ✓ Artigo inserido: {article['title'][:60]}...")
    return article_id

def link_article_to_score_item(cursor, score_item_id, article_id):
    """Cria relação entre artigo e score_item"""
    # Verificar se já existe
    cursor.execute("""
        SELECT 1 FROM score_item_articles
        WHERE score_item_id = %s AND article_id = %s
    """, (score_item_id, article_id))

    if cursor.fetchone():
        return False

    # Criar relação
    cursor.execute("""
        INSERT INTO score_item_articles (score_item_id, article_id, created_at)
        VALUES (%s, %s, NOW())
    """, (score_item_id, article_id))

    return True

def enrich_chcm_item():
    """Enriquece o item CHCM com conteúdo clínico e artigos"""

    score_item_id = 'f1a1d106-b5ac-4483-9de9-ee90ae103d26'

    print("=" * 80)
    print("ENRIQUECIMENTO DO ITEM: CHCM (MCHC)")
    print("=" * 80)

    try:
        conn = get_db_connection()
        cursor = conn.cursor()

        # 1. Atualizar campos do score_item
        print("\n1. Atualizando conteúdo clínico do score_item...")
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
            CHCM_DATA['clinical_relevance'],
            CHCM_DATA['patient_explanation'],
            CHCM_DATA['conduct'],
            datetime.now().date(),
            score_item_id
        ))

        print(f"  ✓ Campos atualizados: clinical_relevance, patient_explanation, conduct, last_review")

        # 2. Inserir artigos científicos
        print(f"\n2. Processando {len(ARTICLES)} artigos científicos...")
        article_ids = []

        for article in ARTICLES:
            article_id = insert_or_get_article(cursor, article)
            article_ids.append(article_id)

        # 3. Criar relações com artigos
        print(f"\n3. Criando relações score_item <-> articles...")
        links_created = 0

        for article_id in article_ids:
            if link_article_to_score_item(cursor, score_item_id, article_id):
                links_created += 1

        print(f"  ✓ {links_created} novas relações criadas")

        # 4. Commit e verificação final
        conn.commit()

        print("\n4. Verificação final...")
        cursor.execute("""
            SELECT
                name,
                unit,
                LENGTH(clinical_relevance) as clinical_len,
                LENGTH(patient_explanation) as patient_len,
                LENGTH(conduct) as conduct_len,
                last_review,
                (SELECT COUNT(*) FROM score_item_articles WHERE score_item_id = %s) as article_count
            FROM score_items
            WHERE id = %s
        """, (score_item_id, score_item_id))

        result = cursor.fetchone()

        print("\n" + "=" * 80)
        print("RESULTADO FINAL")
        print("=" * 80)
        print(f"Nome: {result[0]}")
        print(f"Unidade: {result[1]}")
        print(f"Clinical relevance: {result[2]} caracteres")
        print(f"Patient explanation: {result[3]} caracteres")
        print(f"Conduct: {result[4]} caracteres")
        print(f"Last review: {result[5]}")
        print(f"Artigos vinculados: {result[6]}")
        print("=" * 80)
        print("✓ ENRIQUECIMENTO CONCLUÍDO COM SUCESSO!")
        print("=" * 80)

        cursor.close()
        conn.close()

    except Exception as e:
        print(f"\n✗ ERRO: {e}")
        if conn:
            conn.rollback()
        raise

if __name__ == "__main__":
    enrich_chcm_item()
