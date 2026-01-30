#!/usr/bin/env python3
"""
Enriquecimento do score_item: Plaquetas
ID: 91a9cf25-2622-4296-bf16-64b9f95b1e8d
Grupo: Exames > Laboratoriais
"""

import psycopg2
from datetime import datetime
import json

# Database connection
conn = psycopg2.connect(
    dbname="plenya_db",
    user="plenya_user",
    password="plenya_dev_password",
    host="db",
    port="5432"
)
cur = conn.cursor()

# Item ID
ITEM_ID = "91a9cf25-2622-4296-bf16-64b9f95b1e8d"

# ============================================================================
# STEP 1: Insert Scientific Articles
# ============================================================================

articles = [
    {
        "title": "Thrombocytopenia",
        "authors": "Sruthi Jinna, Samprit Karra, Scott W. Penney, Paras B. Khandhar",
        "publish_date": "2025-12-01",
        "journal": "StatPearls Publishing",
        "original_link": "https://www.ncbi.nlm.nih.gov/books/NBK542208/",
        "doi": None,
        "abstract": "Comprehensive review of thrombocytopenia covering definition, etiology, clinical manifestations, and management. Defines platelet count below 150,000/µL as thrombocytopenia, categorizes severity (mild >100k, moderate 50-100k, severe <50k), and discusses causes including autoimmune disorders, infections, drug reactions, and pregnancy complications. Emphasizes paradoxical thrombosis risk and modern treatment approaches including thrombopoietin receptor agonists.",
        "article_type": "review",
        "specialty": "Hematologia"
    },
    {
        "title": "Thrombocytosis: Diagnostic Evaluation, Thrombotic Risk Stratification, and Risk-Based Management Strategies",
        "authors": "Jonathan S Bleeker, William J Hogan",
        "publish_date": "2011-01-01",
        "journal": "Thrombosis",
        "original_link": "https://pmc.ncbi.nlm.nih.gov/articles/PMC3200282/",
        "doi": "10.1155/2011/536062",
        "abstract": "Detailed analysis of thrombocytosis classification and management. Identifies three categories: spurious (artifact), reactive (88-97% of cases, secondary to infection/inflammation), and clonal (myeloproliferative neoplasms). Provides risk stratification for clonal disease based on age >60, prior thrombosis, and leukocytosis ≥8.7×10⁹/L. Discusses personalized treatment strategies including hydroxyurea and aspirin for high-risk essential thrombocythemia.",
        "article_type": "research_article",
        "specialty": "Hematologia"
    },
    {
        "title": "Thrombocytopenia: Evaluation and Management",
        "authors": "Robert L. Gauer, Daniel J. Whitaker",
        "publish_date": "2022-09-01",
        "journal": "American Family Physician",
        "original_link": "https://www.aafp.org/pubs/afp/issues/2022/0900/thrombocytopenia.html",
        "doi": None,
        "abstract": "Practical primary care guidelines for thrombocytopenia. Emphasizes excluding pseudothrombocytopenia first, then distinguishing acute vs chronic. Defines bleeding risk by count: >50k asymptomatic, 20-50k petechiae/bruising, <10k serious bleeding risk. Provides procedural thresholds (40-50k for most procedures, 100k for neurosurgery) and treatment protocols for immune, drug-induced, and heparin-induced thrombocytopenia. Recommends activity restrictions for counts <50k.",
        "article_type": "review",
        "specialty": "Medicina de Família"
    }
]

article_ids = []

print("=" * 80)
print("STEP 1: Inserting scientific articles")
print("=" * 80)

for article in articles:
    # Check if article exists
    cur.execute("""
        SELECT id FROM articles
        WHERE original_link = %s OR (title = %s AND authors = %s)
    """, (article["original_link"], article["title"], article["authors"]))

    existing = cur.fetchone()

    if existing:
        article_id = existing[0]
        print(f"✓ Article already exists: {article['title'][:60]}... (ID: {article_id})")
    else:
        cur.execute("""
            INSERT INTO articles (
                title, authors, publish_date, journal, original_link, doi, abstract,
                article_type, specialty, created_at
            ) VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, NOW())
            RETURNING id
        """, (
            article["title"],
            article["authors"],
            article["publish_date"],
            article["journal"],
            article["original_link"],
            article["doi"],
            article["abstract"],
            article["article_type"],
            article["specialty"]
        ))
        article_id = cur.fetchone()[0]
        print(f"✓ Inserted: {article['title'][:60]}... (ID: {article_id})")

    article_ids.append(article_id)

conn.commit()
print(f"\n✓ {len(article_ids)} articles processed\n")

# ============================================================================
# STEP 2: Create article-score_item relationships
# ============================================================================

print("=" * 80)
print("STEP 2: Creating article-score_item relationships")
print("=" * 80)

for article_id in article_ids:
    # Check if relationship exists
    cur.execute("""
        SELECT 1 FROM article_score_items
        WHERE article_id = %s AND score_item_id = %s
    """, (article_id, ITEM_ID))

    if cur.fetchone():
        print(f"✓ Relationship already exists: article {article_id}")
    else:
        cur.execute("""
            INSERT INTO article_score_items (article_id, score_item_id)
            VALUES (%s, %s)
        """, (article_id, ITEM_ID))
        print(f"✓ Created relationship: article {article_id} <-> item {ITEM_ID}")

conn.commit()
print(f"\n✓ All relationships created\n")

# ============================================================================
# STEP 3: Update score_item with clinical content (PT-BR)
# ============================================================================

print("=" * 80)
print("STEP 3: Updating score_item with clinical content")
print("=" * 80)

clinical_relevance = """A contagem de plaquetas é um marcador essencial da capacidade de coagulação sanguínea e da saúde hematológica. Valores normais variam entre 150.000-450.000/µL, podendo apresentar pequenas variações entre laboratórios.

**TROMBOCITOPENIA (Plaquetas Baixas - <150.000/µL):**
A redução das plaquetas compromete a hemostasia e aumenta o risco de sangramentos. A gravidade clínica correlaciona-se diretamente com a magnitude da queda:

• **Leve (100.000-150.000/µL):** Geralmente assintomática, sem risco significativo
• **Moderada (50.000-100.000/µL):** Pode causar equimoses e sangramento excessivo em traumas
• **Grave (<50.000/µL):** Risco de sangramento espontâneo em procedimentos invasivos
• **Crítica (<10.000/µL):** Alto risco de hemorragias espontâneas graves

Paradoxalmente, algumas causas de trombocitopenia (como trombocitopenia induzida por heparina e síndrome antifosfolípide) aumentam o risco de trombose, não de sangramento.

**Causas principais:** Distúrbios autoimunes (PTI), infecções virais/bacterianas, reações medicamentosas, neoplasias, doenças da medula óssea, quimioterapia, e complicações gestacionais (pré-eclâmpsia, síndrome HELLP).

**TROMBOCITOSE (Plaquetas Elevadas - >450.000/µL):**
A elevação das plaquetas pode ser reativa (secundária) ou clonal (neoplásica):

• **Reativa (88-97% dos casos):** Secundária a infecções, inflamações crônicas, deficiência de ferro, cirurgias recentes, ou traumas. Geralmente não requer tratamento específico.
• **Clonal (Trombocitemia Essencial):** Doença mieloproliferativa com proliferação anormal da medula óssea. Associada a mutação JAK2V617F. Aumenta risco de trombose (AVE, TEP, TVP) e sangramento paradoxal.

**Fatores de alto risco para trombose:** Idade >60 anos, história prévia de trombose, leucocitose ≥8.700/µL, presença de mutação JAK2V617F.

**Relevância clínica:** A interpretação correta da contagem plaquetária orienta desde a segurança de procedimentos invasivos até o diagnóstico de doenças graves como leucemias e distúrbios autoimunes."""

patient_explanation = """As plaquetas são pequenas células do sangue responsáveis pela coagulação e cicatrização. Quando você se machuca, elas formam um "tampão" para estancar o sangramento.

**Valores normais:** Entre 150.000 e 450.000 plaquetas por microlitro de sangue (k/µL).

**O que significa ter POUCAS plaquetas (trombocitopenia)?**
Quando suas plaquetas estão baixas, seu sangue tem mais dificuldade para coagular, aumentando o risco de:
• Hematomas (manchas roxas na pele) que aparecem facilmente
• Sangramento nasal frequente
• Sangramento gengival ao escovar os dentes
• Menstruação mais intensa que o normal
• Petéquias (pontinhos vermelhos na pele)
• Em casos graves (<20.000/µL): sangramentos espontâneos que podem ser perigosos

**Causas comuns de plaquetas baixas:**
- Infecções virais (dengue, COVID-19)
- Reações a medicamentos
- Doenças autoimunes (seu corpo ataca suas próprias plaquetas)
- Tratamentos como quimioterapia
- Problemas na medula óssea
- Gravidez (especialmente pré-eclâmpsia)

**O que significa ter MUITAS plaquetas (trombocitose)?**
Na maioria das vezes, plaquetas altas são uma reação do corpo a:
• Infecções ou inflamações
• Anemia por falta de ferro
• Recuperação pós-cirúrgica
• Uso de alguns medicamentos

Raramente, pode indicar doenças da medula óssea que aumentam o risco de:
• Coágulos sanguíneos (trombose) em veias ou artérias
• AVC ou infarto
• Paradoxalmente, também pode causar sangramentos

**Quando se preocupar:**
• Plaquetas <50.000/µL: Evite atividades de impacto e procure orientação médica
• Plaquetas <20.000/µL: Busque atendimento urgente
• Plaquetas muito altas (>1.000.000/µL) associadas a sintomas: Investigue com hematologista"""

conduct = """**TROMBOCITOPENIA - Protocolo de Avaliação e Conduta:**

**1. Excluir pseudotrombocitopenia:**
- Primeiro passo: Recoletar sangue em tubo de heparina ou citrato de sódio
- Avaliar presença de agregados plaquetários no esfregaço sanguíneo
- Pseudotrombocitopenia ocorre por aglutinação de plaquetas in vitro (artefato de laboratório)

**2. Avaliar gravidade e risco de sangramento:**

**Contagem >50.000/µL:**
- Geralmente assintomática
- Investigação ambulatorial
- Solicitar: Hemograma completo, esfregaço de sangue periférico
- Investigar causas: Medicamentos, infecções, doenças autoimunes

**Contagem 20.000-50.000/µL:**
- Risco de sangramento em traumas/procedimentos
- Afastar medicamentos suspeitos
- Solicitar: Anti-HIV, anti-HCV, sorologias virais, TSH, FAN
- Considerar encaminhamento ao hematologista

**Contagem <20.000/µL:**
- **EMERGÊNCIA HEMATOLÓGICA**
- Internação hospitalar ou avaliação urgente por hematologista
- Orientações ao paciente:
  • Evitar atividades de contato e exercícios vigorosos
  • Não usar AINEs, AAS ou anticoagulantes
  • Retornar imediatamente se: sangramento que não para, petéquias difusas, hematúria, melena

**3. Investigação diagnóstica:**
- **Hemograma completo + esfregaço:** Avaliar outras linhagens celulares
- **Mielograma/biópsia de medula:** Se pancitopenia, blastos, ou suspeita de doença medular
- **Testes para PTI:** Anti-glicoproteína plaquetária, FAN, anticoagulante lúpico
- **Sorologias:** HIV, HCV, H. pylori (associado a PTI)
- **Pesquisa de HIT:** Se uso recente de heparina

**4. Tratamento conforme etiologia:**

**Trombocitopenia Imune Primária (PTI):**
- **1ª linha:** Corticosteroides (prednisona 1mg/kg/dia) ou Imunoglobulina IV
- **2ª linha:** Agonistas de receptor de trombopoetina (eltrombopague, romiplostim)
- **Refratários:** Rituximabe, esplenectomia

**Induzida por drogas:**
- Suspender medicamento suspeito imediatamente
- Recuperação esperada em 7-10 dias
- Evitar re-exposição ao medicamento

**Trombocitopenia Induzida por Heparina (HIT):**
- **Suspender heparina IMEDIATAMENTE**
- Iniciar anticoagulação alternativa (argatroban, fondaparinux)
- Alto risco de trombose paradoxal

**Thresholds para procedimentos:**
- Procedimentos menores: 40.000-50.000/µL
- Procedimentos de médio risco: 50.000/µL
- Neurocirurgia/oftalmologia: 100.000/µL
- Transfusão profilática: Considerar se <10.000/µL ou <50.000/µL com sangramento ativo

---

**TROMBOCITOSE - Protocolo de Avaliação e Conduta:**

**1. Classificar como reativa ou clonal:**

**Investigação inicial:**
- Hemograma completo seriado
- PCR, VHS, ferritina (marcadores inflamatórios)
- Ferro sérico, capacidade de ligação, saturação de transferrina
- Investigar infecções, neoplasias ocultas, doenças inflamatórias

**Trombocitose Reativa (>450.000/µL):**
- **Causas comuns:** Infecções, inflamações crônicas, deficiência de ferro, pós-operatório, trauma, malignidades
- Geralmente não causa sintomas
- **Conduta:** Tratar doença de base; plaquetas normalizam com resolução da causa
- Não requer tratamento antitrombótico específico

**Trombocitose Clonal - Suspeitar se:**
- Contagem persistente >600.000/µL sem causa secundária aparente
- Esplenomegalia
- Sintomas vasomotores (cefaleia, parestesias, eritromelalgia)
- História de trombose ou sangramento

**2. Investigação para Trombocitemia Essencial:**
- **Encaminhar ao hematologista** se suspeita de causa clonal
- Solicitar:
  • Pesquisa de mutação JAK2V617F (positiva em 50-60% dos casos)
  • Mutações CALR e MPL
  • Biópsia de medula óssea (confirma diagnóstico)

**3. Estratificação de risco (Trombocitemia Essencial):**

**Alto risco:**
- Idade >60 anos
- História prévia de trombose
- Leucocitose ≥8.700/µL

**Baixo risco:**
- Idade <60 anos
- Sem história de trombose
- Sem leucocitose

**4. Tratamento:**

**Trombocitemia Essencial - Baixo Risco:**
- AAS 100mg/dia (reduz risco de trombose microvascular)
- Observação clínica regular

**Trombocitemia Essencial - Alto Risco:**
- **Citorredução:** Hidroxiureia (1ª linha) para manter plaquetas <400.000/µL
- AAS 100mg/dia
- Alternativas: Anagrelida, interferon-alfa (em jovens/gestantes)

**5. Monitoramento:**
- Hemograma mensal durante ajuste de tratamento
- Após estabilização: Hemograma a cada 3 meses
- Avaliar sintomas de trombose ou sangramento em todas as consultas
- Ecocardiograma anual se alto risco cardiovascular

**6. Orientações ao paciente com trombocitose clonal:**
- Manter boa hidratação
- Evitar tabagismo
- Controlar fatores de risco cardiovascular (HAS, DM, dislipidemia)
- Sinais de alerta: Dor torácica, dispneia súbita, cefaleia intensa, déficits neurológicos"""

# Update the score_item
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
    ITEM_ID
))

conn.commit()
print(f"✓ Updated score_item {ITEM_ID} with clinical content")
print(f"  - clinical_relevance: {len(clinical_relevance)} chars")
print(f"  - patient_explanation: {len(patient_explanation)} chars")
print(f"  - conduct: {len(conduct)} chars")
print(f"  - last_review: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}")

# ============================================================================
# STEP 4: Verification
# ============================================================================

print("\n" + "=" * 80)
print("STEP 4: Verification")
print("=" * 80)

# Verify article count
cur.execute("""
    SELECT COUNT(*) FROM article_score_items
    WHERE score_item_id = %s
""", (ITEM_ID,))
article_count = cur.fetchone()[0]
print(f"✓ Articles linked: {article_count}")

# Verify clinical content
cur.execute("""
    SELECT
        name,
        unit,
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
        last_review
    FROM score_items
    WHERE id = %s
""", (ITEM_ID,))

result = cur.fetchone()
print(f"\nScore Item: {result[0]}")
print(f"Unit: {result[1]}")
print(f"Clinical Relevance: {result[2]}")
print(f"Patient Explanation: {result[3]}")
print(f"Conduct: {result[4]}")
print(f"Last Review: {result[5]}")

# Close connection
cur.close()
conn.close()

print("\n" + "=" * 80)
print("✓ ENRICHMENT COMPLETED SUCCESSFULLY")
print("=" * 80)
print(f"\nItem ID: {ITEM_ID}")
print(f"Item: Plaquetas")
print(f"Articles: {article_count}")
print(f"Clinical content: COMPLETE")
print("\n" + "=" * 80)
