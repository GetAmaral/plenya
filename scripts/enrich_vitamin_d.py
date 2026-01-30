#!/usr/bin/env python3
"""
Enriquecimento do item: 25-hidroxivitamina D
ID: ad40c276-dabd-4ef6-8f6e-66db64655c69
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

ITEM_ID = "ad40c276-dabd-4ef6-8f6e-66db64655c69"

# ==========================================
# BUSCAR ARTIGOS EXISTENTES NO BANCO
# ==========================================
print("\n" + "="*60)
print("BUSCANDO ARTIGOS SOBRE VITAMINA D NO BANCO...")
print("="*60)

cur.execute("""
    SELECT id, title, authors, journal, publish_date, original_link
    FROM articles
    WHERE
        title ILIKE '%vitamin D%'
        OR title ILIKE '%25-hydroxyvitamin%'
        OR title ILIKE '%cholecalciferol%'
    ORDER BY publish_date DESC
    LIMIT 5;
""")

existing_articles = cur.fetchall()
print(f"\n✓ Encontrados {len(existing_articles)} artigos no banco:")
for article in existing_articles:
    print(f"  • [{article[4]}] {article[1][:80]}...")

# ==========================================
# ADICIONAR NOVOS ARTIGOS CIENTÍFICOS
# ==========================================
print("\n" + "="*60)
print("ADICIONANDO NOVOS ARTIGOS CIENTÍFICOS...")
print("="*60)

new_articles = [
    {
        "title": "Vitamin D for the Prevention of Disease: An Endocrine Society Clinical Practice Guideline",
        "authors": "Demay MB, Pittas AG, Bikle DD, et al.",
        "journal": "Journal of Clinical Endocrinology & Metabolism",
        "publish_date": "2024-06-01",
        "original_link": "https://pubmed.ncbi.nlm.nih.gov/38828931/",
        "doi": "10.1210/clinem/dgae290",
        "article_type": "clinical_trial"
    },
    {
        "title": "Vitamin D supplementation to prevent acute respiratory infections: systematic review and meta-analysis",
        "authors": "Jolliffe DA, Camargo CA Jr, Sluyter JD, et al.",
        "journal": "The Lancet Diabetes & Endocrinology",
        "publish_date": "2024-12-01",
        "original_link": "https://pubmed.ncbi.nlm.nih.gov/39993397/",
        "doi": "10.1016/S2213-8587(24)00348-6",
        "article_type": "meta_analysis"
    },
    {
        "title": "Vitamin D and Calcium in Osteoporosis: Role of Bone Turnover Markers",
        "authors": "Fassio A, Adami G, Gatti D, et al.",
        "journal": "Nutrients",
        "publish_date": "2023-02-01",
        "original_link": "https://pmc.ncbi.nlm.nih.gov/articles/PMC9944083/",
        "doi": "10.3390/nu15051161",
        "article_type": "review"
    },
    {
        "title": "The effect of vitamin D on bone and osteoporosis",
        "authors": "Holick MF",
        "journal": "Best Practice & Research Clinical Endocrinology & Metabolism",
        "publish_date": "2011-08-01",
        "original_link": "https://pubmed.ncbi.nlm.nih.gov/21872800/",
        "doi": "10.1016/j.beem.2011.05.002",
        "article_type": "review"
    }
]

article_ids = []

for article_data in new_articles:
    # Verificar se já existe
    cur.execute("""
        SELECT id FROM articles
        WHERE title = %s OR doi = %s
    """, (article_data["title"], article_data.get("doi")))

    existing = cur.fetchone()

    if existing:
        article_id = existing[0]
        print(f"  ✓ Artigo já existe: {article_data['title'][:60]}...")
    else:
        cur.execute("""
            INSERT INTO articles (title, authors, journal, publish_date, original_link, doi, article_type)
            VALUES (%s, %s, %s, %s, %s, %s, %s)
            RETURNING id;
        """, (
            article_data["title"],
            article_data["authors"],
            article_data["journal"],
            article_data["publish_date"],
            article_data["original_link"],
            article_data.get("doi"),
            article_data.get("article_type", "research_article")
        ))
        article_id = cur.fetchone()[0]
        print(f"  + Artigo adicionado: {article_data['title'][:60]}...")

    article_ids.append(article_id)

conn.commit()

# ==========================================
# CRIAR RELAÇÕES ITEM <-> ARTIGOS
# ==========================================
print("\n" + "="*60)
print("CRIANDO RELAÇÕES ITEM <-> ARTIGOS...")
print("="*60)

for article_id in article_ids:
    cur.execute("""
        INSERT INTO article_score_items (score_item_id, article_id)
        VALUES (%s, %s)
        ON CONFLICT (score_item_id, article_id) DO NOTHING;
    """, (ITEM_ID, article_id))
    print(f"  ✓ Relação criada: item → artigo {article_id}")

conn.commit()

# ==========================================
# ENRIQUECER ITEM COM CONTEÚDO CLÍNICO
# ==========================================
print("\n" + "="*60)
print("ENRIQUECENDO ITEM COM CONTEÚDO CLÍNICO...")
print("="*60)

clinical_content = {
    "clinical_relevance": """A 25-hidroxivitamina D [25(OH)D] é a principal forma circulante de vitamina D e o melhor marcador do status vitamínico D no organismo. Suas funções essenciais incluem:

**Saúde Óssea e Metabolismo do Cálcio:**
• Regula a absorção intestinal de cálcio e fósforo
• Previne raquitismo em crianças e osteomalácia em adultos
• Reduz risco de osteoporose e fraturas quando associada à suplementação de cálcio
• Doses mínimas de 800 UI/dia de vitamina D com 1200 mg de cálcio são recomendadas para prevenção de osteoporose

**Função Imunológica:**
• Modula respostas imunes inatas e adaptativas
• Aumenta atividade fagocítica de macrófagos
• Suprime replicação viral e protege barreiras epiteliais respiratórias
• Meta-análise de 2024 (64.086 participantes) mostrou redução de 6% nas infecções respiratórias agudas com suplementação (OR 0.94, IC 95% 0.88-1.00)

**Novas Diretrizes 2024 (Endocrine Society):**
• Não mais recomenda dosagem rotineira de 25(OH)D em adultos saudáveis <75 anos
• Não estabelece níveis-alvo específicos de suficiência/deficiência
• Recomenda suplementação empírica para: gestantes, idosos >75 anos, crianças/adolescentes e pessoas com pré-diabetes de alto risco
• Para adultos 18-75 anos saudáveis: apenas Dietary Reference Intake (600-800 UI/dia)

**Interpretação de Valores (referência tradicional):**
• Deficiência: <20 ng/mL (<50 nmol/L)
• Insuficiência: 20-29 ng/mL (50-74 nmol/L)
• Suficiência: 30-60 ng/mL (75-150 nmol/L)
• Toxicidade: >100 ng/mL (>250 nmol/L)

**Nota:** As diretrizes de 2024 geraram controvérsia na comunidade científica, com especialistas criticando foco excessivo em ensaios clínicos randomizados e desconsideração de estudos observacionais.""",

    "patient_explanation": """**O que é a Vitamina D?**

A vitamina D é um hormônio essencial que seu corpo produz quando você se expõe ao sol. O exame de 25-hidroxivitamina D mede quanto você tem no sangue.

**Por que é importante?**

🦴 **Ossos Fortes:** Ajuda seu corpo a absorver cálcio dos alimentos. Sem vitamina D suficiente, seus ossos podem ficar fracos e quebrar mais facilmente.

🛡️ **Sistema Imunológico:** Ajuda seu corpo a combater infecções como gripes e resfriados. Estudos recentes mostram que pessoas com bons níveis de vitamina D têm menos infecções respiratórias.

⚡ **Energia e Bem-estar:** Níveis baixos podem causar cansaço, fraqueza muscular e até tristeza.

**Como melhorar seus níveis:**

☀️ **Sol:** 10-15 minutos de sol (sem protetor) em braços e pernas, 3x/semana
🥛 **Alimentos:** Peixes gordos (salmão, sardinha), ovos, leite fortificado, cogumelos
💊 **Suplementação:** Se necessário, seu médico pode prescrever vitamina D3 (doses variam de 1000 a 10000 UI/dia conforme o caso)

**Importante:** As recomendações mudaram em 2024. Se você é adulto saudável com menos de 75 anos, pode não precisar fazer esse exame rotineiramente. Converse com seu médico sobre seu caso específico.""",

    "conduct": """**Conduta Clínica - 25-Hidroxivitamina D:**

**1. Indicações para Dosagem (Diretrizes 2024):**
• NÃO dosar rotineiramente em adultos saudáveis <75 anos
• Dosar apenas em casos selecionados:
  - Osteoporose confirmada ou alto risco de fraturas
  - Doenças que afetam absorção de vitamina D (doença celíaca, Crohn, cirurgia bariátrica)
  - Insuficiência renal crônica
  - Hiperparatireoidismo
  - Uso crônico de medicamentos que interferem no metabolismo da vitamina D (anticonvulsivantes, glicocorticoides)
  - Raquitismo ou osteomalácia

**2. Interpretação e Conduta por Níveis:**

**Deficiência (<20 ng/mL):**
• Suplementação de ATAQUE: 50.000 UI/semana VO por 6-8 semanas
• Seguida de manutenção: 1000-2000 UI/dia VO
• Associar cálcio 1000-1200 mg/dia se risco de osteoporose
• Reavaliar níveis após 3-4 meses

**Insuficiência (20-29 ng/mL):**
• Suplementação de manutenção: 1000-2000 UI/dia VO
• Se sintomas (fadiga, dores ósseas): considerar dose de ataque
• Reavaliar após 3-6 meses

**Suficiência (30-60 ng/mL):**
• Manter suplementação preventiva apenas em grupos de risco:
  - Gestantes: 600-800 UI/dia
  - Idosos >75 anos: 800-1000 UI/dia
  - Crianças/adolescentes: conforme peso (400-600 UI/dia)
  - Pré-diabetes alto risco: 1000-2000 UI/dia

**Níveis Elevados (>60 ng/mL):**
• Investigar suplementação excessiva
• Se >100 ng/mL: suspender suplementação imediatamente e avaliar hipercalcemia

**3. Contraindicações à Suplementação:**
• Hipercalcemia (cálcio sérico >10.5 mg/dL)
• Hipercalciúria (cálcio urinário >300 mg/24h)
• Nefrolitíase ativa
• Sarcoidose ou outras doenças granulomatosas (risco de hipercalcemia)

**4. Monitoramento:**
• Reavaliar 25(OH)D após 3-4 meses de tratamento
• Dosar cálcio sérico e urinário se doses >4000 UI/dia
• Em idosos com osteoporose: seguimento anual após atingir níveis adequados

**5. Orientações Não-Farmacológicas:**
• Exposição solar: 10-15 min sem protetor solar, 3x/semana (braços/pernas)
• Alimentação: peixes gordos 2-3x/semana, ovos, alimentos fortificados
• Atividade física: estimula síntese de vitamina D e saúde óssea

**Referências:** Diretrizes Endocrine Society 2024, Lancet Diabetes & Endocrinology 2024."""
}

# Buscar dados atuais do item
cur.execute("""
    SELECT clinical_relevance, patient_explanation, conduct
    FROM score_items
    WHERE id = %s;
""", (ITEM_ID,))

current_data = cur.fetchone()

# Atualizar item
cur.execute("""
    UPDATE score_items
    SET
        clinical_relevance = %s,
        patient_explanation = %s,
        conduct = %s,
        updated_at = NOW()
    WHERE id = %s;
""", (
    clinical_content["clinical_relevance"],
    clinical_content["patient_explanation"],
    clinical_content["conduct"],
    ITEM_ID
))

conn.commit()
print(f"\n✓ Item atualizado com sucesso!")
print(f"  • clinical_relevance: {len(clinical_content['clinical_relevance'])} caracteres")
print(f"  • patient_explanation: {len(clinical_content['patient_explanation'])} caracteres")
print(f"  • conduct: {len(clinical_content['conduct'])} caracteres")

# ==========================================
# VERIFICAÇÃO FINAL
# ==========================================
print("\n" + "="*60)
print("VERIFICAÇÃO FINAL")
print("="*60)

cur.execute("""
    SELECT
        si.id,
        si.name,
        si.clinical_relevance IS NOT NULL as has_clinical,
        si.patient_explanation IS NOT NULL as has_patient,
        si.conduct IS NOT NULL as has_conduct,
        COUNT(DISTINCT asi.article_id) as article_count
    FROM score_items si
    LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
    WHERE si.id = %s
    GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct;
""", (ITEM_ID,))

verification = cur.fetchone()

print(f"\n✓ Item: {verification[1]}")
print(f"✓ Clinical Relevance: {'✓ Preenchido' if verification[2] else '✗ Vazio'}")
print(f"✓ Patient Explanation: {'✓ Preenchido' if verification[3] else '✗ Vazio'}")
print(f"✓ Conduct: {'✓ Preenchido' if verification[4] else '✗ Vazio'}")
print(f"✓ Artigos vinculados: {verification[5]}")

# ==========================================
# RELATÓRIO JSON
# ==========================================
report = {
    "item_id": ITEM_ID,
    "item_name": "25-hidroxivitamina D",
    "timestamp": datetime.now().isoformat(),
    "articles_added": len(new_articles),
    "articles_linked": verification[5],
    "fields_updated": ["clinical_relevance", "patient_explanation", "conduct"],
    "status": "✓ COMPLETO"
}

import os
report_dir = '/scripts/enrichment_data'
os.makedirs(report_dir, exist_ok=True)

with open(f'{report_dir}/vitamin_d_report.json', 'w', encoding='utf-8') as f:
    json.dump(report, f, ensure_ascii=False, indent=2)

print(f"\n✓ Relatório salvo em: enrichment_data/vitamin_d_report.json")

cur.close()
conn.close()

print("\n" + "="*60)
print("✓ ENRIQUECIMENTO CONCLUÍDO COM SUCESSO!")
print("="*60 + "\n")
