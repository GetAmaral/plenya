#!/usr/bin/env python3
"""
Enriquecimento do item: Leucócitos Totais (WBC)
ID: 56350a39-c589-4ed5-b757-0424399b7a61
Grupo: Exames > Laboratoriais

Baseado em fontes científicas atualizadas (2024-2025):
- StatPearls (NCBI): Leukocytosis & Normal/Abnormal CBC
- Medscape: Differential Blood Count
- PubMed: Clinical guidelines e evidências
"""

import psycopg2
import os
from datetime import datetime

# Configurações do banco
DB_CONFIG = {
    'host': os.getenv('DB_HOST', 'localhost'),
    'port': os.getenv('DB_PORT', '5432'),
    'database': os.getenv('DB_NAME', 'plenya_db'),
    'user': os.getenv('DB_USER', 'plenya_user'),
    'password': os.getenv('DB_PASSWORD', 'plenya_password')
}

SCORE_ITEM_ID = '56350a39-c589-4ed5-b757-0424399b7a61'

# Artigos científicos para inserir
ARTICLES = [
    {
        'title': 'Leukocytosis',
        'authors': 'Riley LK, Rupert J',
        'journal': 'StatPearls [Internet]',
        'year': 2024,
        'url': 'https://www.ncbi.nlm.nih.gov/books/NBK560882/',
        'doi': None,
        'pubmed_id': '32809717',
        'summary': 'Comprehensive review of leukocytosis including definitions, age-specific normal ranges, etiology by cell type (neutrophilia, lymphocytosis, eosinophilia, monocytosis, basophilia), leukemoid reactions, clinical evaluation guidelines, differential diagnosis, and management of hyperleukocytosis.',
        'key_findings': 'Normal adult WBC: 4,500-11,000 cells/µL. Hyperleukocytosis (>100,000 cells/µL) requires urgent evaluation. Neutrophilia (>7,700/µL) is most common cause. Leukostasis complications include CNS/pulmonary symptoms. Prognostic significance in cardiovascular events.',
        'study_type': 'clinical_guideline',
        'evidence_level': 'A'
    },
    {
        'title': 'Normal and Abnormal Complete Blood Count With Differential',
        'authors': 'Neapolitan R, Malki J, Cohen B',
        'journal': 'StatPearls [Internet]',
        'year': 2024,
        'url': 'https://www.ncbi.nlm.nih.gov/books/NBK604207/',
        'doi': None,
        'pubmed_id': None,
        'summary': 'Detailed reference guide for CBC with differential interpretation, including normal reference ranges for WBC and differential counts, clinical significance of leukocytosis and leukopenia, spurious causes, and interpretation guidelines.',
        'key_findings': 'Normal WBC: 4,500-11,000 cells/µL. Differential ranges: Neutrophils 40-60% (1,500-8,000/µL), Lymphocytes 20-40% (1,000-4,000/µL), Monocytes 2-8% (200-1,000/µL), Eosinophils 0-4% (0-500/µL), Basophils 0.5-1% (0-200/µL). Results must be interpreted in clinical context.',
        'study_type': 'clinical_guideline',
        'evidence_level': 'A'
    },
    {
        'title': 'Differential Blood Count: Reference Range, Interpretation, Collection and Panels',
        'authors': 'Medscape Medical Reference',
        'journal': 'Medscape',
        'year': 2024,
        'url': 'https://emedicine.medscape.com/article/2085133-overview',
        'doi': None,
        'pubmed_id': None,
        'summary': 'Clinical reference for differential blood count utility in generating absolute values for each WBC type, diagnostic applications in identifying neutropenia, neutrophilia, lymphopenia, and lymphocytosis, and clinical significance of neutrophil-lymphocyte ratio.',
        'key_findings': 'Absolute values more meaningful than percentages. Neutrophil-lymphocyte count ratio (NLCR) is simple promising method to evaluate systemic inflammation in critically ill. Severity of clinical course correlates with divergence of neutrophil/lymphocyte counts.',
        'study_type': 'clinical_guideline',
        'evidence_level': 'B'
    },
    {
        'title': 'Neutropenia',
        'authors': 'Newburger PE, Dale DC',
        'journal': 'StatPearls [Internet]',
        'year': 2024,
        'url': 'https://www.ncbi.nlm.nih.gov/books/NBK507702/',
        'doi': None,
        'pubmed_id': '29939530',
        'summary': 'Comprehensive review of neutropenia including benign ethnic neutropenia, causes (infection, drugs, malignancy, immunoneutropenia), evaluation approaches, and management including G-CSF therapy for chemotherapy-induced neutropenia.',
        'key_findings': 'Leukopenia defined as WBC <4,000/mm³. Life-threatening in agranulocytosis with fever (requires immediate broad-spectrum antibiotics). G-CSF stimulates bone marrow to produce more WBC. Check previous counts to assess dynamic development.',
        'study_type': 'systematic_review',
        'evidence_level': 'A'
    }
]

# Conteúdo clínico em português
CLINICAL_CONTENT = {
    'clinical_relevance': '''**Valores de Referência:**
- Adultos: 4.000-11.000 células/µL (4-11 mil/µL)
- Variação por idade: neonatos (13.000-38.000), crianças (5.000-17.500), gestantes 3º trimestre (5.800-13.200)

**Leucocitose (>11.000/µL):**
Elevação da contagem de leucócitos totais, podendo indicar:
- **Neutrofilia (>7.700/µL):** Infecções bacterianas, inflamação, estresse, uso de corticosteroides, neoplasias mieloproliferativas, tabagismo
- **Linfocitose:** Infecções virais (EBV, CMV, influenza), coqueluche, leucemia, linfoma, doenças autoimunes
- **Eosinofilia (>500/µL):** Alergias, parasitoses, asma, doenças mieloproliferativas
- **Monocitose (>1.000/µL):** Infecções crônicas (tuberculose, endocardite), doença inflamatória intestinal
- **Basofilia:** Rara, sugere neoplasia mieloproliferativa se persistente >8 semanas

**Hiperleucocitose (>100.000/µL):** Emergência médica potencial. Risco de leucostase (sintomas neurológicos/pulmonares), requer avaliação hematológica urgente.

**Reação Leucemoide:** WBC >50.000/µL com neutrofilia sem neoplasia mieloproliferativa. Diferencia-se da leucemia pela presença de neutrófilos maduros (não blastos) e resolução com tratamento da causa base.

**Leucopenia (<4.000/µL):**
Redução da contagem de leucócitos, podendo indicar:
- Falência ou supressão medular (quimioterapia, radiação)
- Infecções virais (HIV, CMV, sarampo) ou bacterianas graves (tifoide, sepse)
- Doenças autoimunes (lúpus)
- Medicamentos (antimicrobianos, anticonvulsivantes, anti-inflamatórios)
- Hiperesplenismo
- Neutropenia étnica benigna (africanos, oriente médio, caribenhos)

**Significado Prognóstico:**
- Em eventos cardiovasculares/cerebrovasculares: leucocitose correlaciona-se com gravidade da lesão isquêmica e desfechos adversos
- No infarto do miocárdio: WBC elevado associa-se a maior mortalidade
- Agranulocitose com febre: risco de vida, requer internação e antibióticos imediatos

**Interpretação do Diferencial:**
- Neutrófilos 40-60% (1.500-8.000/µL)
- Linfócitos 20-40% (1.000-4.000/µL)
- Monócitos 2-8% (200-1.000/µL)
- Eosinófilos 0-4% (0-500/µL)
- Basófilos 0,5-1% (0-200/µL)

**Relação Neutrófilo-Linfócito (RNL):**
Método simples e promissor para avaliar inflamação sistêmica em pacientes críticos. A gravidade do quadro clínico correlaciona-se com a divergência das contagens de neutrófilos e linfócitos.''',

    'patient_explanation': '''**O que são os Leucócitos (Glóbulos Brancos)?**

Os leucócitos, também conhecidos como glóbulos brancos, são células de defesa do seu organismo. Eles protegem você contra infecções, combatem microrganismos invasores e participam de reações alérgicas e inflamatórias.

**Valor Normal:**
- Entre 4.000 e 11.000 células por microlitro (4-11 mil/µL)

**Quando os Leucócitos Estão Aumentados (Leucocitose):**

*Valores elevados geralmente indicam:*
- **Infecção:** Seu corpo está combatendo bactérias, vírus ou outros microrganismos
- **Inflamação:** Resposta a lesões, cirurgias ou doenças inflamatórias
- **Estresse físico ou emocional:** Situações de estresse podem elevar temporariamente
- **Uso de medicamentos:** Especialmente corticosteroides
- **Tabagismo:** Fumantes frequentemente apresentam valores mais altos
- **Gravidez:** Valores mais elevados são normais no 3º trimestre
- **Alergias ou parasitoses:** Especialmente quando eosinófilos estão elevados
- **Doenças do sangue:** Em casos mais raros

**Atenção:** Valores muito altos (acima de 100.000/µL) requerem avaliação médica urgente.

**Quando os Leucócitos Estão Diminuídos (Leucopenia):**

*Valores baixos podem indicar:*
- **Infecções virais:** Como gripe, sarampo ou infecções mais graves (HIV, CMV)
- **Medicamentos:** Alguns antibióticos, anti-inflamatórios ou medicamentos para tireoide
- **Doenças autoimunes:** Como lúpus
- **Quimioterapia ou radioterapia:** Tratamentos oncológicos frequentemente reduzem leucócitos
- **Deficiência de vitaminas:** Especialmente vitamina B12 ou ácido fólico
- **Doenças da medula óssea:** Onde as células são produzidas
- **Variação étnica normal:** Pessoas de origem africana, do oriente médio ou caribenha podem ter valores naturalmente mais baixos

**Atenção:** Valores muito baixos aumentam o risco de infecções. Febre associada a leucopenia baixa é emergência médica.

**O Diferencial de Leucócitos:**

Existem 5 tipos principais de leucócitos, cada um com função específica:
- **Neutrófilos (40-60%):** Primeira linha de defesa contra bactérias
- **Linfócitos (20-40%):** Combatem vírus e produzem anticorpos
- **Monócitos (2-8%):** Eliminam células mortas e combatem infecções crônicas
- **Eosinófilos (0-4%):** Atuam em alergias e parasitoses
- **Basófilos (0,5-1%):** Participam de reações alérgicas

**O que fazer?**
- Valores alterados isoladamente nem sempre indicam doença
- Sempre correlacione com seus sintomas e histórico médico
- Exercícios físicos e estresse podem alterar temporariamente os valores
- Consulte seu médico para interpretação adequada dos resultados''',

    'conduct': '''**1. AVALIAÇÃO INICIAL:**

**Leucocitose (WBC >11.000/µL):**
- Revisar hemograma completo com diferencial (percentual e valores absolutos)
- Analisar esfregaço de sangue periférico para excluir artefatos e avaliar morfologia
- Comparar com hemogramas anteriores para identificar tendências
- História clínica detalhada:
  * Sintomas de infecção (febre, sinais flogísticos)
  * Medicamentos em uso (corticosteroides, lítio, catecolaminas)
  * Exposições ocupacionais, tabagismo
  * História familiar de doenças hematológicas
- Exame físico: linfadenopatia, esplenomegalia, sinais de infecção

**Leucopenia (WBC <4.000/µL):**
- Hemograma completo com diferencial e esfregaço periférico
- Comparar com valores prévios (leucopenia crônica vs. aguda)
- Verificar se há bi ou pancitopenia (sugere problema medular)
- História de:
  * Medicamentos (antimicrobianos, anticonvulsivantes, AINEs)
  * Infecções recentes ou crônicas
  * Doenças autoimunes
  * Exposição a radiação ou quimioterápicos
  * Origem étnica (neutropenia étnica benigna)
- Avaliar sinais/sintomas de infecção ativa

**2. INVESTIGAÇÃO POR GRAVIDADE:**

**URGENTE (Avaliação Hematológica Imediata):**
- WBC >100.000/µL (hiperleucocitose)
- Linfocitose >50.000/µL (investigar malignidade)
- Leucopenia com neutropenia grave (<500/µL) + febre
- Sintomas de leucostase: alterações visuais, cefaleia, confusão mental, dispneia, hipóxia
- Presença de blastos no esfregaço periférico

**INVESTIGAÇÃO COMPLEMENTAR (Leucocitose Moderada):**
- Função renal e hepática
- Marcadores inflamatórios: PCR, VHS
- Hemoculturas e culturas conforme suspeita clínica
- Se neutrofilia: investigar foco infeccioso ou inflamatório
- Se linfocitose: sorologias virais (EBV, CMV), se persistente considerar citometria de fluxo
- Se eosinofilia >1.500/µL: repetir hemograma em 1-2 semanas; se persistente, investigar síndrome hipereosinofílica
- Se monocitose: avaliar infecções crônicas (tuberculose, endocardite)
- Se basofilia persistente >8 semanas: investigar doença mieloproliferativa

**INVESTIGAÇÃO COMPLEMENTAR (Leucopenia):**
- Hemograma de controle em 1-2 semanas se assintomático
- Função tireoidiana
- Dosagem de vitamina B12 e ácido fólico
- Sorologias: HIV, CMV, hepatites
- Autoanticorpos se suspeita de doença autoimune
- Considerar mielograma se:
  * Leucopenia persistente sem causa identificada
  * Bi ou pancitopenia
  * Células anormais no esfregaço periférico

**3. MANEJO CLÍNICO:**

**Leucocitose:**
- Maioria resolve sem intervenção, tratar causa subjacente
- Hiperleucocitose com leucostase:
  * Hidratação vigorosa
  * Considerar leucaférese
  * Quimioterapia se leucemia/linfoma
  * Monitorar complicações: CID, síndrome de lise tumoral
- Reação leucemoide: resolver condição desencadeante
- Retirar medicamentos causadores quando possível

**Leucopenia:**
- **Agranulocitose com febre (EMERGÊNCIA):**
  * Internação imediata
  * Antibióticos de amplo espectro empiricamente
  * Isolamento protetor
  * Considerar G-CSF (fator estimulador de colônias de granulócitos)
- Leucopenia leve-moderada assintomática:
  * Observação e repetição de hemograma
  * Orientar sobre sinais de infecção
  * Evitar exposições de risco
- Se relacionada a medicamento:
  * Suspender ou substituir fármaco quando possível
  * Reavaliar hemograma após suspensão
- Se secundária a quimioterapia:
  * G-CSF profilático ou terapêutico conforme protocolo
  * Profilaxia antimicrobiana se neutropenia grave prolongada

**4. CAUSAS ESPÚRIAS (Excluir):**
- Leucocitose: agregados plaquetários, coágulos de fibrina, hemácias nucleadas contadas como leucócitos
- Leucopenia: agregação de leucócitos

**5. MONITORAMENTO:**
- WBC >50.000/µL: seguimento hematológico regular
- Eosinofilia >1.500/µL persistente: investigar dano orgânico (cardíaco, pulmonar, neurológico)
- Leucopenia em quimioterapia: hemogramas seriados conforme protocolo
- Leucocitose em eventos cardiovasculares: associa-se a pior prognóstico, monitorar evolução

**6. CONSIDERAÇÕES ESPECIAIS:**
- **Etnia:** Neutropenia étnica benigna é normal em certas populações (africanos, oriente médio, caribenhos)
- **Gestação:** Leucocitose fisiológica no 3º trimestre (até 13.200/µL)
- **Exercício e estresse:** Podem elevar temporariamente WBC e diferencial
- **Tabagismo:** Causa leucocitose crônica leve
- **Reconciliação medicamentosa:** AINEs, alopurinol, fenitoína, penicilinas, cefalosporinas podem causar leucocitose inespecífica'''
}

def create_article(cursor, article_data):
    """Cria um artigo científico e retorna seu ID"""
    insert_article = """
        INSERT INTO articles
        (title, authors, journal, publish_date, language, doi, pm_id, abstract,
         notes, original_link, article_type, created_at, updated_at)
        VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, NOW(), NOW())
        RETURNING id;
    """

    # Mapear study_type para article_type (validado pelo CHECK constraint)
    study_type_map = {
        'clinical_guideline': 'review',
        'systematic_review': 'review',
        'meta_analysis': 'meta_analysis',
        'randomized_controlled_trial': 'clinical_trial',
        'cohort_study': 'research_article',
        'case_control': 'case_study',
        'cross_sectional': 'research_article'
    }

    article_type = study_type_map.get(article_data['study_type'], 'review')

    # Criar data de publicação no formato YYYY-MM-DD
    publish_date = f"{article_data['year']}-01-01"

    # Combinar summary e key_findings no abstract
    abstract = f"{article_data['summary']}\n\nKey Findings: {article_data['key_findings']}"

    # Usar notes para evidence_level
    notes = f"Evidence Level: {article_data['evidence_level']}"

    cursor.execute(insert_article, (
        article_data['title'],
        article_data['authors'],
        article_data['journal'],
        publish_date,
        'en',  # language
        article_data['doi'],
        article_data['pubmed_id'],
        abstract,
        notes,
        article_data['url'],
        article_type
    ))

    article_id = cursor.fetchone()[0]
    print(f"✓ Artigo criado: {article_data['title']} (ID: {article_id})")
    return article_id

def link_article_to_score_item(cursor, article_id, score_item_id):
    """Cria a relação entre artigo e score_item"""
    insert_link = """
        INSERT INTO article_score_items (score_item_id, article_id)
        VALUES (%s, %s)
        ON CONFLICT (score_item_id, article_id) DO NOTHING;
    """

    cursor.execute(insert_link, (score_item_id, article_id))
    print(f"✓ Artigo {article_id} vinculado ao score_item {score_item_id}")

def update_score_item_content(cursor, score_item_id, content):
    """Atualiza o conteúdo clínico do score_item"""
    update_sql = """
        UPDATE score_items
        SET
            clinical_relevance = %s,
            patient_explanation = %s,
            conduct = %s,
            last_review = %s,
            updated_at = NOW()
        WHERE id = %s;
    """

    cursor.execute(update_sql, (
        content['clinical_relevance'],
        content['patient_explanation'],
        content['conduct'],
        datetime.now().date(),
        score_item_id
    ))

    print(f"✓ Conteúdo clínico atualizado para score_item {score_item_id}")

def verify_score_item_exists(cursor, score_item_id):
    """Verifica se o score_item existe"""
    cursor.execute("SELECT name FROM score_items WHERE id = %s", (score_item_id,))
    result = cursor.fetchone()
    if result:
        print(f"✓ Score item encontrado: {result[0]}")
        return True
    else:
        print(f"✗ Score item {score_item_id} não encontrado!")
        return False

def main():
    print("=" * 80)
    print("ENRIQUECIMENTO: Leucócitos Totais (WBC)")
    print("=" * 80)
    print(f"Score Item ID: {SCORE_ITEM_ID}")
    print(f"Artigos a inserir: {len(ARTICLES)}")
    print()

    try:
        # Conectar ao banco
        conn = psycopg2.connect(**DB_CONFIG)
        conn.autocommit = False
        cursor = conn.cursor()

        print("✓ Conectado ao banco de dados")
        print()

        # Verificar se o score_item existe
        if not verify_score_item_exists(cursor, SCORE_ITEM_ID):
            print("\n✗ Abortando: Score item não encontrado")
            return

        print()
        print("-" * 80)
        print("FASE 1: Inserindo artigos científicos")
        print("-" * 80)

        article_ids = []
        for i, article in enumerate(ARTICLES, 1):
            print(f"\n[{i}/{len(ARTICLES)}] {article['title']}")
            article_id = create_article(cursor, article)
            article_ids.append(article_id)

        print()
        print("-" * 80)
        print("FASE 2: Vinculando artigos ao score_item")
        print("-" * 80)
        print()

        for article_id in article_ids:
            link_article_to_score_item(cursor, article_id, SCORE_ITEM_ID)

        print()
        print("-" * 80)
        print("FASE 3: Atualizando conteúdo clínico em português")
        print("-" * 80)
        print()

        update_score_item_content(cursor, SCORE_ITEM_ID, CLINICAL_CONTENT)

        # Commit das alterações
        conn.commit()

        print()
        print("=" * 80)
        print("✓ ENRIQUECIMENTO CONCLUÍDO COM SUCESSO!")
        print("=" * 80)
        print(f"✓ {len(article_ids)} artigos científicos inseridos")
        print(f"✓ {len(article_ids)} vínculos criados")
        print("✓ Conteúdo clínico atualizado (clinical_relevance, patient_explanation, conduct)")
        print("✓ Campo last_review atualizado")
        print()

        # Verificação final
        cursor.execute("""
            SELECT
                si.name,
                COUNT(asi.article_id) as num_articles,
                LENGTH(si.clinical_relevance) as len_clinical,
                LENGTH(si.patient_explanation) as len_patient,
                LENGTH(si.conduct) as len_conduct,
                si.last_review
            FROM score_items si
            LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
            WHERE si.id = %s
            GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct, si.last_review
        """, (SCORE_ITEM_ID,))

        result = cursor.fetchone()
        if result:
            print("VERIFICAÇÃO FINAL:")
            print(f"  Nome: {result[0]}")
            print(f"  Artigos vinculados: {result[1]}")
            print(f"  Tamanho clinical_relevance: {result[2]} caracteres")
            print(f"  Tamanho patient_explanation: {result[3]} caracteres")
            print(f"  Tamanho conduct: {result[4]} caracteres")
            print(f"  Última revisão: {result[5]}")

        cursor.close()
        conn.close()

    except psycopg2.Error as e:
        print(f"\n✗ Erro no banco de dados: {e}")
        if conn:
            conn.rollback()
            conn.close()
        raise
    except Exception as e:
        print(f"\n✗ Erro inesperado: {e}")
        if conn:
            conn.rollback()
            conn.close()
        raise

if __name__ == "__main__":
    main()
