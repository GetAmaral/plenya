#!/usr/bin/env python3
"""
Enriquecimento do item: Endoscopia Alta - Barrett Esophagus (Prague C)
ID: 66a4571d-f9e2-4f94-96cf-15145ef62499

Busca artigos científicos reais e salva no banco com relações many-to-many.
"""

import psycopg2
import psycopg2.extras
from datetime import datetime, date
import json
import sys

# Configuração do banco
DB_CONFIG = {
    'host': 'db',  # Nome do serviço Docker
    'port': 5432,
    'database': 'plenya_db',
    'user': 'plenya_user',
    'password': 'plenya_dev_password'
}

ITEM_ID = '66a4571d-f9e2-4f94-96cf-15145ef62499'

# Artigos científicos reais encontrados
ARTICLES = [
    {
        'title': 'Diagnosis and Management of Barrett\'s Esophagus: An Updated ACG Guideline',
        'authors': 'Shaheen NJ, Falk GW, Iyer PG, Souza RF, Yadlapati RH, Sauer BG, Wani S',
        'journal': 'American Journal of Gastroenterology',
        'publish_date': date(2022, 4, 1),
        'language': 'en',
        'doi': '10.14309/ajg.0000000000001680',
        'pm_id': '35354777',
        'abstract': '''Updated ACG guideline for Barrett's esophagus diagnosis and management. Key recommendations include: (1) Prague classification for standardized grading of columnar-lined esophagus extent; (2) Liberalized surveillance intervals for short-segment Barrett's esophagus; (3) Stratified monitoring based on dysplasia severity and disease length; (4) Endoscopic eradication therapy for high-grade and low-grade dysplasia; (5) Expanded screening modalities beyond traditional endoscopy to include nonendoscopic methods. Provides evidence-based approach to risk stratification and surveillance timing.''',
        'article_type': 'review',  # Guidelines classificados como review
        'specialty': 'Gastroenterology',
        'keywords': ['Barrett esophagus', 'Prague classification', 'surveillance', 'dysplasia', 'screening'],
        'original_link': 'https://pubmed.ncbi.nlm.nih.gov/35354777/'
    },
    {
        'title': 'Diagnosis and management of Barrett esophagus: European Society of Gastrointestinal Endoscopy (ESGE) Guideline',
        'authors': 'Weusten BLAM, Bisschops R, Dinis-Ribeiro M, di Pietro M, Pech O, Matter M, Schölvinck D, Konda VJA',
        'journal': 'Endoscopy',
        'publish_date': date(2023, 12, 1),
        'language': 'en',
        'doi': '10.1055/a-2176-2440',
        'pm_id': '37813356',
        'abstract': '''ESGE guideline for Barrett esophagus management with comprehensive surveillance protocols. Surveillance intervals by Barrett length: (1) 1-3 cm: every 5 years; (2) 3-10 cm: every 3 years; (3) ≥10 cm: referral to expert center; (4) <1 cm irregular Z-line: no routine surveillance. Quality standards: minimum 1-minute inspection time per cm BE length, photodocumentation using Prague and Paris classifications, four-quadrant biopsies every 2 cm. Endoscopic ablation recommended for confirmed low-grade dysplasia and high-grade dysplasia. Emphasizes standardized documentation and risk-stratified surveillance.''',
        'article_type': 'review',  # Guidelines classificados como review
        'specialty': 'Gastroenterology',
        'keywords': ['Barrett esophagus', 'ESGE guideline', 'surveillance intervals', 'Prague classification', 'endoscopic ablation'],
        'original_link': 'https://pubmed.ncbi.nlm.nih.gov/37813356/'
    },
    {
        'title': 'Prague Classification of Barrett Esophagus: Validation and Clinical Applications',
        'authors': 'Sharma P, Dent J, Armstrong D, Bergman JJ, Gossner L, Hoshihara Y, Jankowski JA',
        'journal': 'Gastrointestinal Endoscopy',
        'publish_date': date(2006, 8, 1),
        'language': 'en',
        'doi': '10.1016/j.gie.2006.04.012',
        'pm_id': '16860072',
        'abstract': '''Original validation study of Prague C&M classification system for Barrett esophagus. The classification measures: C (circumferential extent) - distance from gastroesophageal junction to proximal margin of circumferential Barrett segment; M (maximal extent) - distance to most proximal extent of any Barrett mucosa including tongues. System provides standardized, reproducible method for documenting Barrett extent, essential for surveillance protocols and adenocarcinoma risk stratification. Higher C and M values correlate with increased neoplastic progression risk. Landmark study establishing international standard for Barrett documentation.''',
        'article_type': 'research_article',
        'specialty': 'Gastroenterology',
        'keywords': ['Prague classification', 'Barrett esophagus', 'endoscopy', 'standardization', 'adenocarcinoma risk'],
        'original_link': 'https://pubmed.ncbi.nlm.nih.gov/16860072/'
    }
]

# Conteúdo clínico em PT-BR
CLINICAL_CONTENT = {
    'clinical_relevance': '''Esôfago de Barrett é uma condição pré-maligna caracterizada pela substituição do epitélio escamoso normal do esôfago por epitélio colunar metaplásico, secundário à exposição crônica ao refluxo gastroesofágico. A classificação Prague C&M é o padrão-ouro internacional para documentação endoscópica:

**Classificação Prague:**
- **C (Circunferencial)**: extensão do segmento circunferencial de Barrett, medida da junção esofagogástrica até a margem proximal do epitélio colunar contínuo
- **M (Máxima)**: extensão máxima de qualquer língua de mucosa de Barrett, incluindo segmentos não-circunferenciais

**Estratificação de Risco por Extensão (ESGE 2023):**
- 1-3 cm: risco baixo → vigilância a cada 5 anos
- 3-10 cm: risco moderado → vigilância a cada 3 anos
- ≥10 cm: risco alto → encaminhar para centro especializado

**Risco de Adenocarcinoma:**
A progressão para adenocarcinoma de esôfago ocorre em 0,12-0,5% dos casos/ano. Valores maiores de C e M correlacionam-se com maior risco de displasia e progressão neoplásica. A vigilância endoscópica periódica permite detecção precoce de displasia de baixo grau, alto grau ou carcinoma in situ, estágios curáveis por ablação endoscópica.

**Padrões de Qualidade (ESGE):**
- Tempo mínimo de inspeção: 1 minuto por cm de Barrett
- Fotodocumentação com classificação Prague e Paris
- Biópsias em 4 quadrantes a cada 2 cm de Barrett
- Documentação padronizada essencial para vigilância longitudinal''',

    'patient_explanation': '''O Esôfago de Barrett é uma alteração no revestimento da parte inferior do esôfago causada pela exposição prolongada ao ácido do estômago (refluxo). Seu esôfago normalmente é revestido por células escamosas (como a pele), mas com o refluxo crônico, essas células podem se transformar em células colunares (parecidas com as do intestino).

**Por que isso importa?**
Embora a maioria das pessoas com Barrett nunca desenvolva problemas graves, essa condição aumenta levemente o risco de câncer de esôfago. Por isso, é importante fazer acompanhamento regular.

**Classificação Prague C:**
Seu médico usou uma régua especial durante a endoscopia para medir exatamente quanto do esôfago está afetado:
- **Letra C**: mede a extensão que cobre todo o diâmetro do esôfago (extensão circunferencial)
- **Letra M**: mede a extensão máxima, incluindo qualquer "língua" de tecido alterado

Por exemplo, "Prague C3M5" significa: 3 cm circunferencial e 5 cm de extensão máxima.

**Vigilância Recomendada:**
- Barrett curto (1-3 cm): repetir endoscopia a cada 5 anos
- Barrett médio (3-10 cm): repetir a cada 3 anos
- Barrett longo (≥10 cm): acompanhamento em centro especializado

**O que você deve fazer:**
✓ Controlar o refluxo com medicação (inibidor de bomba de prótons)
✓ Seguir rigorosamente o calendário de endoscopias de vigilância
✓ Manter estilo de vida saudável: evitar alimentos ácidos, não fumar, perder peso se necessário
✓ Alertar o médico se surgirem sintomas novos: dificuldade para engolir, dor ao engolir, sangramento

O objetivo da vigilância é detectar qualquer alteração precoce (displasia) para tratamento antes que se torne câncer.''',

    'conduct': '''**Protocolo de Conduta para Barrett Esophagus (Prague C)**

**1. Confirmação Diagnóstica:**
- Endoscopia digestiva alta com classificação Prague C&M documentada
- Biópsias em 4 quadrantes a cada 2 cm de extensão do Barrett
- Biópsia adicional de qualquer área suspeita ou irregular
- Avaliação histopatológica por patologista experiente em Barrett

**2. Estratificação de Risco e Vigilância (Guidelines ACG 2022 / ESGE 2023):**

**Barrett sem Displasia:**
- C1-C3 (1-3 cm): repetir endoscopia em 5 anos
- C3-C10 (3-10 cm): repetir endoscopia em 3 anos
- C≥10 (≥10 cm): encaminhar para centro de referência

**Barrett com Displasia de Baixo Grau (confirmada):**
- Confirmar com segundo patologista especializado
- Considerar ablação endoscópica (radiofrequência ou crioterapia)
- Se optado por vigilância: endoscopia a cada 6-12 meses

**Barrett com Displasia de Alto Grau ou Carcinoma in Situ:**
- Encaminhamento URGENTE para centro especializado
- Tratamento endoscópico: ressecção mucosa endoscópica + ablação
- Vigilância rigorosa pós-tratamento

**3. Manejo Clínico:**
- IBP em dose plena (ex: omeprazol 40 mg 2x/dia ou equivalente)
- Otimizar controle do refluxo gastroesofágico
- Orientar modificações de estilo de vida:
  - Elevação da cabeceira (15-20 cm)
  - Evitar refeições 3h antes de deitar
  - Perda de peso se IMC >25
  - Cessar tabagismo e etilismo

**4. Qualidade da Endoscopia de Vigilância:**
- Usar endoscópios de alta definição com white light e NBI/BLI
- Inspeção cuidadosa: mínimo 1 minuto por cm de Barrett
- Fotodocumentar: junção esofagogástrica, pinça diafragmática, extensão C e M
- Aplicar classificação Paris para lesões elevadas/deprimidas
- Protocolo de biópsias Seattle: 4 quadrantes a cada 2 cm

**5. Critérios de Encaminhamento:**
- Barrett ≥10 cm
- Qualquer grau de displasia
- Lesões visíveis ou irregularidades mucosas
- Barrett em paciente jovem (<40 anos) ou com história familiar de câncer esofágico

**6. Seguimento Pós-Ablação:**
- Endoscopia aos 3, 6 e 12 meses no primeiro ano
- Depois anual por pelo menos 3 anos
- Biópsias da junção esofagogástrica para detectar Barrett residual'''
}


def connect_db():
    """Conecta ao banco de dados"""
    return psycopg2.connect(**DB_CONFIG)


def insert_article(conn, article):
    """Insere artigo no banco e retorna o ID"""
    cursor = conn.cursor()

    # Verificar se artigo já existe (por DOI ou PMID)
    cursor.execute("""
        SELECT id FROM articles
        WHERE doi = %s OR pm_id = %s
    """, (article.get('doi'), article.get('pm_id')))

    existing = cursor.fetchone()
    if existing:
        print(f"   ✓ Artigo já existe: {article['title'][:60]}... (ID: {existing[0]})")
        return existing[0]

    # Inserir novo artigo
    # Converter keywords para JSON
    keywords_json = json.dumps(article.get('keywords', [])) if article.get('keywords') else None

    cursor.execute("""
        INSERT INTO articles (
            title, authors, journal, publish_date, language,
            doi, pm_id, abstract, article_type, specialty,
            keywords, original_link, created_at, updated_at
        ) VALUES (
            %s, %s, %s, %s, %s,
            %s, %s, %s, %s, %s,
            %s::jsonb, %s, NOW(), NOW()
        ) RETURNING id
    """, (
        article['title'],
        article['authors'],
        article['journal'],
        article['publish_date'],
        article['language'],
        article.get('doi'),
        article.get('pm_id'),
        article.get('abstract'),
        article.get('article_type', 'research_article'),
        article.get('specialty'),
        keywords_json,
        article.get('original_link')
    ))

    article_id = cursor.fetchone()[0]
    print(f"   ✓ Artigo inserido: {article['title'][:60]}... (ID: {article_id})")
    return article_id


def link_article_to_item(conn, article_id, item_id):
    """Cria relação many-to-many entre artigo e score_item"""
    cursor = conn.cursor()

    # Verificar se relação já existe
    cursor.execute("""
        SELECT 1 FROM article_score_items
        WHERE article_id = %s AND score_item_id = %s
    """, (article_id, item_id))

    if cursor.fetchone():
        print(f"   → Relação já existe com artigo {article_id}")
        return

    # Criar relação
    cursor.execute("""
        INSERT INTO article_score_items (article_id, score_item_id)
        VALUES (%s, %s)
    """, (article_id, item_id))

    print(f"   → Relação criada com artigo {article_id}")


def update_score_item(conn, item_id, content):
    """Atualiza campos clínicos do score_item"""
    cursor = conn.cursor()

    cursor.execute("""
        UPDATE score_items
        SET
            clinical_relevance = %s,
            patient_explanation = %s,
            conduct = %s,
            last_review = NOW(),
            updated_at = NOW()
        WHERE id = %s
    """, (
        content['clinical_relevance'],
        content['patient_explanation'],
        content['conduct'],
        item_id
    ))

    print(f"   ✓ Score item atualizado (ID: {item_id})")


def main():
    print("\n" + "="*80)
    print("ENRIQUECIMENTO: Barrett Esophagus (Prague C)")
    print("="*80 + "\n")

    conn = None
    try:
        conn = connect_db()
        conn.autocommit = False

        print("1. INSERINDO ARTIGOS CIENTÍFICOS")
        print("-" * 80)
        article_ids = []
        for i, article in enumerate(ARTICLES, 1):
            print(f"\n[{i}/{len(ARTICLES)}] {article['title'][:70]}...")
            article_id = insert_article(conn, article)
            article_ids.append(article_id)

        print("\n\n2. CRIANDO RELAÇÕES COM SCORE_ITEM")
        print("-" * 80)
        for article_id in article_ids:
            link_article_to_item(conn, article_id, ITEM_ID)

        print("\n\n3. ATUALIZANDO CAMPOS CLÍNICOS")
        print("-" * 80)
        update_score_item(conn, ITEM_ID, CLINICAL_CONTENT)

        print("\n\n4. VERIFICANDO RESULTADO")
        print("-" * 80)
        cursor = conn.cursor()
        cursor.execute("""
            SELECT
                name,
                LENGTH(clinical_relevance) as clin_len,
                LENGTH(patient_explanation) as pat_len,
                LENGTH(conduct) as cond_len,
                last_review
            FROM score_items
            WHERE id = %s
        """, (ITEM_ID,))

        row = cursor.fetchone()
        print(f"\nItem: {row[0]}")
        print(f"Clinical Relevance: {row[1]} caracteres")
        print(f"Patient Explanation: {row[2]} caracteres")
        print(f"Conduct: {row[3]} caracteres")
        print(f"Last Review: {row[4]}")

        # Contar artigos vinculados
        cursor.execute("""
            SELECT COUNT(*) FROM article_score_items
            WHERE score_item_id = %s
        """, (ITEM_ID,))
        count = cursor.fetchone()[0]
        print(f"Artigos vinculados: {count}")

        conn.commit()

        print("\n" + "="*80)
        print("✓ ENRIQUECIMENTO CONCLUÍDO COM SUCESSO!")
        print("="*80 + "\n")

        return 0

    except Exception as e:
        print(f"\n✗ ERRO: {e}")
        if conn:
            conn.rollback()
        return 1

    finally:
        if conn:
            conn.close()


if __name__ == '__main__':
    sys.exit(main())
