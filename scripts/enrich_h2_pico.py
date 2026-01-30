#!/usr/bin/env python3
"""
Enriquecimento: Hidrogênio Pico (H₂ Máximo)
Item ID: 210eedab-08e7-4f47-ae6a-37aecea9bc16
"""

import psycopg2
from datetime import datetime
import json

# Conexão com banco
conn = psycopg2.connect(
    host="localhost",
    port=5432,
    database="plenya_db",
    user="plenya_user",
    password="plenya_password"
)
cur = conn.cursor()

ITEM_ID = "210eedab-08e7-4f47-ae6a-37aecea9bc16"

# Artigos científicos
articles = [
    {
        "title": "Understanding Our Tests: Hydrogen-Methane Breath Testing to Diagnose Small Intestinal Bacterial Overgrowth",
        "authors": "Rezaie A, et al.",
        "journal": "Clinical and Translational Gastroenterology",
        "year": 2023,
        "doi": "10.14309/ctg.0000000000000567",
        "url": "https://pmc.ncbi.nlm.nih.gov/articles/PMC10132719/",
        "article_type": "review"
    },
    {
        "title": "Hydrogen and Methane-Based Breath Testing in Gastrointestinal Disorders: The North American Consensus",
        "authors": "Rezaie A, Buresi M, Lembo A, et al.",
        "journal": "American Journal of Gastroenterology",
        "year": 2017,
        "doi": "10.1038/ajg.2017.46",
        "url": "https://pubmed.ncbi.nlm.nih.gov/28323273/",
        "article_type": "consensus"
    },
    {
        "title": "Refined Lactulose Hydrogen Breath Test for Small Intestinal Bacterial Overgrowth Subgrouping Irritable Bowel Syndrome",
        "authors": "Dahlgren A, et al.",
        "journal": "Gastroenterology Research and Practice",
        "year": 2025,
        "doi": "10.1155/grp/5597071",
        "url": "https://onlinelibrary.wiley.com/doi/10.1155/grp/5597071",
        "article_type": "clinical_trial"
    },
    {
        "title": "How to Interpret Hydrogen Breath Tests",
        "authors": "Rana SV, Malik A",
        "journal": "Journal of Neurogastroenterology and Motility",
        "year": 2014,
        "doi": "10.5056/jnm.2014.20.3.312",
        "url": "https://pmc.ncbi.nlm.nih.gov/articles/PMC4155069/",
        "article_type": "review"
    },
    {
        "title": "Carbohydrate malabsorption in patients with non-specific abdominal complaints",
        "authors": "Yao CK, Gibson PR, Shepherd SJ",
        "journal": "Current Opinion in Clinical Nutrition & Metabolic Care",
        "year": 2014,
        "doi": "10.1097/MCO.0000000000000094",
        "url": "https://pmc.ncbi.nlm.nih.gov/articles/PMC4171253/",
        "article_type": "review"
    }
]

# Inserir artigos e criar relações
article_ids = []
for article in articles:
    # Verificar se já existe
    cur.execute("""
        SELECT id FROM scientific_articles
        WHERE doi = %s OR title = %s
    """, (article.get('doi'), article['title']))

    existing = cur.fetchone()

    if existing:
        article_id = existing[0]
        print(f"✓ Artigo já existe: {article['title'][:60]}...")
    else:
        cur.execute("""
            INSERT INTO scientific_articles
            (title, authors, journal, publication_year, doi, url, article_type, relevance_score)
            VALUES (%s, %s, %s, %s, %s, %s, %s, %s)
            RETURNING id
        """, (
            article['title'],
            article['authors'],
            article['journal'],
            article['year'],
            article.get('doi'),
            article['url'],
            article['article_type'],
            95  # Alta relevância para teste diagnóstico
        ))
        article_id = cur.fetchone()[0]
        print(f"✓ Artigo inserido: {article['title'][:60]}...")

    article_ids.append(article_id)

    # Criar relação com score_item
    cur.execute("""
        INSERT INTO score_item_articles (score_item_id, article_id, relevance_context)
        VALUES (%s, %s, %s)
        ON CONFLICT (score_item_id, article_id) DO NOTHING
    """, (
        ITEM_ID,
        article_id,
        "Diagnóstico de SIBO e má absorção de carboidratos via teste respiratório"
    ))

print(f"\n✓ {len(article_ids)} artigos vinculados ao item")

# Conteúdo clínico em PT-BR
clinical_relevance = """
O Hidrogênio Pico (H₂ Máximo) é o valor máximo de hidrogênio exalado durante o teste respiratório de hidrogênio, sendo fundamental para diagnóstico de:

**SIBO (Supercrescimento Bacteriano no Intestino Delgado):**
- Critério diagnóstico: elevação ≥20 ppm nos primeiros 90 minutos (consenso norte-americano)
- Sensibilidade: 77% / Especificidade: 88% (usando cutoff de 20 ppm)
- Pacientes com SIBO apresentam pico médio de 38 ppm vs. 8 ppm em indivíduos saudáveis

**Má Absorção de Carboidratos:**
- Intolerância à lactose: pico ≥20 ppm acima do basal após carga de lactose
- Má absorção de frutose: pico ≥20 ppm acima do basal após carga de frutose
- Cutoff diferenciado: 12 ppm para SIBO vs. 20 ppm para má absorção

**Interpretação Conjunta:**
- Metano ≥10 ppm indica IMO (Intestinal Methanogen Overgrowth)
- Produtores mistos: considerar média de H₂ + CH₄ (cutoff 15 ppm)
- O valor do pico NÃO se correlaciona com intensidade dos sintomas

**Limitações:**
- 15-27% da população são "não produtores" de H₂ (resultado falso-negativo)
- Preparação inadequada pode gerar resultados inválidos
- Deve ser interpretado junto com tempo de trânsito orocecal
"""

patient_explanation = """
Este exame mede o hidrogênio no ar que você expira após beber uma solução açucarada, ajudando a identificar problemas digestivos.

**Como Funciona:**
Normalmente, o açúcar é absorvido no intestino delgado. Se houver:
- Bactérias em excesso no intestino delgado (SIBO)
- Dificuldade em absorver lactose ou frutose

As bactérias fermentam o açúcar e produzem hidrogênio, que vai para o sangue e é eliminado pelos pulmões.

**Valores de Referência:**
- Normal: pico até 20 ppm (partes por milhão)
- SIBO: pico ≥20 ppm nos primeiros 90 minutos
- Má absorção: pico ≥20 ppm acima do valor inicial

**O Que Significa Resultado Alterado:**
- Pode indicar SIBO (supercrescimento bacteriano)
- Pode indicar intolerância à lactose ou frutose
- Pode explicar sintomas como inchaço, diarreia, gases e dor abdominal

**Importante:**
O resultado precisa ser interpretado pelo médico junto com seus sintomas e o tempo que levou para o hidrogênio subir. Algumas pessoas não produzem hidrogênio (são "não produtoras") e podem ter resultado falso normal.
"""

conduct = """
**Interpretação do Resultado:**

1. **H₂ Pico <20 ppm:**
   - Resultado negativo para SIBO e má absorção
   - ATENÇÃO: descartar paciente "não produtor" de H₂
   - Considerar teste com metano se sintomas persistentes

2. **H₂ Pico ≥20 ppm nos primeiros 90 min:**
   - SIBO provável (sensibilidade 77%, especificidade 88%)
   - Correlacionar com tempo de pico (antes ou depois de 90 min)
   - Verificar metano concomitante (IMO se ≥10 ppm)
   - **Conduta:**
     - Antibioticoterapia (rifaximina 550mg 3x/dia por 14 dias)
     - Dieta pobre em FODMAPs durante tratamento
     - Investigar causa base (dismotilidade, hipocloridria, cirurgia prévia)
     - Repetir teste 2-4 semanas após tratamento

3. **H₂ Pico ≥20 ppm após 90 min (>120 min):**
   - Má absorção de carboidrato provável
   - Identificar qual carboidrato (lactose, frutose, lactulose)
   - **Conduta:**
     - Dieta de exclusão do carboidrato identificado
     - Suplementação com enzimas digestivas (ex: lactase para lactose)
     - Orientação nutricional
     - Reavaliar sintomas em 4-6 semanas

4. **H₂ Pico 10-19 ppm (zona cinzenta):**
   - Interpretar com cautela
   - Considerar contexto clínico e sintomas
   - Pode indicar disbiose leve
   - Repetir teste ou complementar com metano

**Preparação Adequada do Paciente (Essencial):**
- Suspender antibióticos: 4 semanas antes
- Suspender probióticos: 2 semanas antes
- Suspender laxantes/procinéticos: 1 semana antes
- Jejum de 12 horas (exceto água)
- Não fumar/exercitar no dia do teste
- Dieta restrita no dia anterior (arroz branco, frango grelhado)

**Follow-up:**
- Sucesso terapêutico: redução de ≥50% no pico ou normalização
- Se falha: investigar resistência, IMO concomitante, causa não corrigida
- Recorrência comum: abordar fatores predisponentes
"""

# Atualizar score_item
cur.execute("""
    UPDATE score_items
    SET
        clinical_relevance = %s,
        patient_explanation = %s,
        conduct = %s,
        reference_range = %s,
        updated_at = NOW()
    WHERE id = %s
""", (
    clinical_relevance,
    patient_explanation,
    conduct,
    "<20 ppm (normal), ≥20 ppm (SIBO ou má absorção)",
    ITEM_ID
))

print(f"\n✓ Conteúdo clínico atualizado")

# Commit
conn.commit()

# Verificar resultado final
cur.execute("""
    SELECT
        si.name,
        si.unit,
        si.reference_range,
        COUNT(sia.article_id) as num_articles,
        CASE
            WHEN si.clinical_relevance IS NOT NULL THEN 'OK'
            ELSE 'PENDENTE'
        END as clinical_content_status
    FROM score_items si
    LEFT JOIN score_item_articles sia ON si.id = sia.score_item_id
    WHERE si.id = %s
    GROUP BY si.id, si.name, si.unit, si.reference_range, si.clinical_relevance
""", (ITEM_ID,))

result = cur.fetchone()
print("\n" + "="*70)
print("RESULTADO FINAL")
print("="*70)
print(f"Item: {result[0]}")
print(f"Unidade: {result[1]}")
print(f"Referência: {result[2]}")
print(f"Artigos vinculados: {result[3]}")
print(f"Conteúdo clínico: {result[4]}")
print("="*70)

# Exportar JSON para validação
export_data = {
    "item_id": ITEM_ID,
    "name": "Hidrogênio Pico (H₂ Máximo)",
    "unit": "ppm",
    "reference_range": "<20 ppm (normal), ≥20 ppm (SIBO ou má absorção)",
    "clinical_relevance_preview": clinical_relevance[:200] + "...",
    "patient_explanation_preview": patient_explanation[:200] + "...",
    "conduct_preview": conduct[:200] + "...",
    "articles_count": len(article_ids),
    "articles": [
        {
            "title": a['title'],
            "year": a['year'],
            "journal": a['journal']
        }
        for a in articles
    ],
    "timestamp": datetime.now().isoformat()
}

with open('/home/user/plenya/h2_pico_enrichment_result.json', 'w', encoding='utf-8') as f:
    json.dump(export_data, f, ensure_ascii=False, indent=2)

print("\n✓ Resultado exportado: h2_pico_enrichment_result.json")

cur.close()
conn.close()

print("\n✅ ENRIQUECIMENTO CONCLUÍDO COM SUCESSO!")
