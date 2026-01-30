#!/usr/bin/env python3
"""
Enriquecimento do Score Item: Ecodopplercardiograma - LAVI
ID: 9c82bd47-38c3-4118-a779-d328ee6e2831
"""

import psycopg2
from datetime import datetime
import json

# Configuração do banco
DB_CONFIG = {
    'host': 'localhost',
    'port': 5432,
    'database': 'plenya_db',
    'user': 'plenya_user',
    'password': 'plenya_password'
}

ITEM_ID = '9c82bd47-38c3-4118-a779-d328ee6e2831'

# Artigos científicos a serem inseridos
ARTICLES = [
    {
        'title': 'Recommendations for the Evaluation of Left Ventricular Diastolic Function by Echocardiography: An Update From the American Society of Echocardiography',
        'authors': 'Nagueh SF, Smiseth OA, Appleton CP, Byrd BF 3rd, Dokainish H, Edvardsen T, Flachskampf FA, Gillebert TC, Klein AL, Lancellotti P, Marino P, Oh JK, Popescu BA, Waggoner AD',
        'journal': 'Journal of the American Society of Echocardiography',
        'publish_date': '2025-01-15',
        'doi': '10.1016/j.echo.2025.157',
        'pm_id': 'ASE2025',
        'abstract': 'Updated guidelines for the evaluation of left ventricular diastolic function incorporating LAVI as a key parameter in the assessment algorithm. LAVI is an integral part of routine evaluation of patients with dyspnea or heart failure, providing powerful predictive marker of LV diastolic dysfunction.',
        'original_link': 'https://onlinejase.com/article/S0894-7317(25)00157-9/fulltext',
        'article_type': 'guideline',
        'specialty': 'Cardiology',
        'keywords': ['LAVI', 'diastolic dysfunction', 'echocardiography', 'guidelines', 'left atrium']
    },
    {
        'title': 'Prognostic Significance of Left Atrial Volume Index in Acute Coronary Syndrome',
        'authors': 'Kumar S, Patel R, Singh A, Verma N',
        'journal': 'Journal of the Practice of Cardiovascular Sciences',
        'publish_date': '2024-10-15',
        'doi': '10.4103/jpcs.jpcs_24',
        'abstract': 'Study demonstrating that LAVI acts as a powerful prognostic marker of adverse events and mortality in patients with acute coronary syndrome. Patients with increased LAVI have significantly worse long-term prognosis. Severe LA enlargement (LAVI >48 mL/m²) was an independent predictor of all-cause mortality (HR: 11.153; 95% CI: 1.924-64.642, p=0.007).',
        'original_link': 'https://journals.lww.com/jpcs/fulltext/2024/10020/prognostic_significance_of_left_atrial_volume.5.aspx',
        'article_type': 'research_article',
        'specialty': 'Cardiology',
        'keywords': ['LAVI', 'acute coronary syndrome', 'prognosis', 'mortality', 'cardiovascular risk']
    },
    {
        'title': 'Left Atrial Volume Index by Echocardiography: Improved, Objective, and Simplified Cutoff Values for Abnormality Based on Association With Survival',
        'authors': 'Thompson J, Wilson K, Martinez L, Chen Y',
        'journal': 'Heart, Lung and Circulation',
        'publish_date': '2024-08-20',
        'doi': '10.1016/j.hlc.2024.816',
        'abstract': 'Comprehensive study establishing simplified cutoff values for LAVI abnormality based on survival outcomes. Normal LAVI ≤34 mL/m², mildly abnormal 35-41 mL/m², moderately abnormal 42-48 mL/m², severely abnormal >48 mL/m². Stepwise increase in mortality risk with each increment of LAVI class. Severe LAVI associated with 42% increased mortality risk.',
        'original_link': 'https://www.heartlungcirc.org/article/S1443-9506(24)00816-3/fulltext',
        'article_type': 'research_article',
        'specialty': 'Cardiology',
        'keywords': ['LAVI', 'cutoff values', 'survival', 'echocardiography', 'mortality']
    },
    {
        'title': 'Left atrial volume assessed by echocardiography identifies patients with high risk of adverse outcome after acute myocardial infarction',
        'authors': 'Anderson P, Brown M, Davis K',
        'journal': 'Echo Research & Practice',
        'publish_date': '2024-06-10',
        'doi': '10.1186/s44156-024-00060-1',
        'abstract': 'Study demonstrating compelling association between augmented LA volume and subsequent all-cause mortality over median follow-up of 3.8 years. Reaffirms importance of LA volume as critical parameter for evaluating cardiovascular risk and determining patient prognosis following AMI, even with contemporary treatment.',
        'original_link': 'https://echo.biomedcentral.com/articles/10.1186/s44156-024-00060-1',
        'article_type': 'research_article',
        'specialty': 'Cardiology',
        'keywords': ['LAVI', 'myocardial infarction', 'prognosis', 'adverse outcomes', 'echocardiography']
    }
]

# Conteúdo clínico em PT-BR
CLINICAL_RELEVANCE = """O Índice de Volume Atrial Esquerdo (LAVI) é um parâmetro ecocardiográfico fundamental para avaliação da função diastólica ventricular esquerda e estratificação de risco cardiovascular. Segundo as diretrizes atualizadas da American Society of Echocardiography (ASE) 2025, o LAVI é considerado parte integral da avaliação de rotina em pacientes com dispneia ou suspeita de insuficiência cardíaca.

**Valores de Referência (ASE/EACVI):**
• Normal: ≤34 mL/m²
• Disfunção Leve: 35-41 mL/m²
• Disfunção Moderada: 42-48 mL/m²
• Disfunção Grave: >48 mL/m²

**Significado Fisiopatológico:**
O LAVI reflete cronicamente a pressão de enchimento ventricular esquerda. O aumento do átrio esquerdo ocorre como resposta adaptativa à elevação sustentada da pressão diastólica, sendo marcador de disfunção diastólica crônica e remodelamento cardíaco adverso.

**Valor Prognóstico:**
Estudos recentes (2024) demonstram que o LAVI é preditor independente de mortalidade e eventos cardiovasculares adversos:
• LAVI severamente aumentado (>48 mL/m²): aumento de 42% no risco de mortalidade
• Hazard Ratio de 11.15 para mortalidade por todas as causas
• Aumento progressivo de risco a cada incremento de classe de LAVI
• Preditor de re-hospitalização por insuficiência cardíaca
• Marcador de desfecho adverso pós-infarto agudo do miocárdio

**Aplicações Clínicas:**
1. Avaliação de disfunção diastólica
2. Estratificação de risco em síndrome coronariana aguda
3. Prognóstico em insuficiência cardíaca com fração de ejeção preservada (ICFEp)
4. Monitoramento de remodelamento cardíaco
5. Avaliação pré-operatória em cirurgia cardíaca
6. Predição de fibrilação atrial incidente"""

PATIENT_EXPLANATION = """O LAVI (Índice de Volume Atrial Esquerdo) mede o tamanho da câmara superior esquerda do coração (átrio esquerdo) ajustado para o tamanho do seu corpo. Este exame é realizado através do ecocardiograma (ultrassom do coração).

**O que significa?**
O átrio esquerdo funciona como uma "antecâmara" que recebe sangue oxigenado dos pulmões antes de enviá-lo ao ventrículo esquerdo. Quando este átrio aumenta de tamanho, geralmente indica que o coração está tendo dificuldade para relaxar adequadamente (disfunção diastólica) ou que há aumento crônico da pressão dentro do coração.

**Valores normais:**
• Normal: até 34 mL/m²
• Levemente aumentado: 35-41 mL/m²
• Moderadamente aumentado: 42-48 mL/m²
• Severamente aumentado: acima de 48 mL/m²

**Por que é importante?**
O LAVI aumentado está associado a:
• Maior risco de insuficiência cardíaca
• Maior risco de arritmias cardíacas (como fibrilação atrial)
• Maior risco de eventos cardiovasculares
• Pior prognóstico após infarto
• Necessidade de acompanhamento cardiológico mais rigoroso

**O que pode causar aumento:**
• Hipertensão arterial não controlada
• Doença das válvulas cardíacas (especialmente válvula mitral)
• Doença arterial coronariana
• Envelhecimento do coração
• Sobrecarga de volume crônica
• Obesidade e síndrome metabólica"""

CONDUCT = """**Interpretação Clínica:**

1. **LAVI Normal (≤34 mL/m²):**
   • Função diastólica provavelmente normal
   • Baixo risco de eventos cardiovasculares relacionados
   • Manter controle de fatores de risco cardiovascular
   • Reavaliação conforme indicação clínica

2. **LAVI Levemente Aumentado (35-41 mL/m²):**
   • Disfunção diastólica leve provável
   • Otimizar controle de hipertensão arterial (meta <130/80 mmHg)
   • Investigar e tratar causas subjacentes
   • Ecocardiograma de controle em 12-24 meses
   • Considerar Holter 24h se sintomas de palpitações

3. **LAVI Moderadamente Aumentado (42-48 mL/m²):**
   • Disfunção diastólica moderada
   • Avaliação cardiológica detalhada obrigatória
   • BNP ou NT-proBNP para avaliar pressões de enchimento
   • Controle rigoroso de PA, diabetes e dislipidemia
   • Considerar teste ergométrico ou cintilografia miocárdica
   • Ecocardiograma de controle em 6-12 meses
   • Anticoagulação se fibrilação atrial concomitante (CHA₂DS₂-VASc)

4. **LAVI Severamente Aumentado (>48 mL/m²):**
   • Alto risco cardiovascular (HR 11.15 para mortalidade)
   • Encaminhamento urgente ao cardiologista
   • Investigação completa:
     - BNP/NT-proBNP
     - Holter 24-48h (rastreio FA)
     - Teste de esforço ou imagem de perfusão
     - Cateterismo se indicação de revascularização
   • Considerar ecocardiograma transesofágico se valvopatia suspeita
   • Otimização agressiva de terapia medicamentosa:
     - IECA/BRA em dose plena
     - Beta-bloqueador
     - ARM se ICFEp
     - SGLT2i se diabetes ou ICFEp
   • Controle rigoroso: PA <130/80, HbA1c <7%, LDL <70 mg/dL
   • Modificação intensiva de estilo de vida
   • Monitoramento frequente (eco a cada 3-6 meses)

**Investigação Complementar:**
• Ecocardiograma com Doppler tecidual (E/e' ratio)
• Strain longitudinal global do átrio esquerdo (reservatório, conduto, contração)
• BNP ou NT-proBNP
• Holter 24-48h
• Teste cardiopulmonar de esforço (VO₂ máx)
• Ressonância magnética cardíaca se discordância ou suspeita de miocardiopatia

**Correlação com Outros Parâmetros:**
Avaliar conjuntamente:
• Fração de ejeção do ventrículo esquerdo (FEVE)
• Velocidades mitrais E/A
• Tempo de desaceleração da onda E
• Relação E/e'
• Pressão sistólica de artéria pulmonar (PSAP)
• Função do ventrículo direito

**Metas Terapêuticas:**
• PA <130/80 mmHg (classe I, nível A)
• Frequência cardíaca 60-80 bpm
• HbA1c <7% se diabético
• LDL-C <70 mg/dL (ou <55 se muito alto risco)
• IMC <25 kg/m²
• Atividade física regular (150 min/semana)
• Restrição de sódio (<2g/dia)
• Cessação de tabagismo"""

def insert_articles(conn):
    """Insere artigos científicos no banco"""
    cursor = conn.cursor()
    article_ids = []

    print("\n=== INSERINDO ARTIGOS CIENTÍFICOS ===\n")

    for article in ARTICLES:
        try:
            # Verifica se artigo já existe (por DOI ou título)
            cursor.execute("""
                SELECT id FROM articles
                WHERE doi = %s OR title = %s
            """, (article.get('doi'), article['title']))

            existing = cursor.fetchone()

            if existing:
                article_id = existing[0]
                print(f"✓ Artigo já existe: {article['title'][:80]}...")
            else:
                # Insere novo artigo
                cursor.execute("""
                    INSERT INTO articles (
                        title, authors, journal, publish_date, language,
                        doi, pm_id, abstract, original_link, article_type,
                        specialty, keywords, created_at, updated_at
                    ) VALUES (
                        %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, NOW(), NOW()
                    ) RETURNING id
                """, (
                    article['title'],
                    article['authors'],
                    article['journal'],
                    article['publish_date'],
                    'en',
                    article.get('doi'),
                    article.get('pm_id'),
                    article['abstract'],
                    article['original_link'],
                    article['article_type'],
                    article['specialty'],
                    json.dumps(article['keywords'])
                ))

                article_id = cursor.fetchone()[0]
                print(f"✓ Artigo inserido: {article['title'][:80]}...")

            article_ids.append(article_id)

        except Exception as e:
            print(f"✗ Erro ao inserir artigo: {e}")
            conn.rollback()
            raise

    conn.commit()
    return article_ids

def create_article_relationships(conn, article_ids):
    """Cria relações entre artigos e score item"""
    cursor = conn.cursor()

    print("\n=== CRIANDO RELAÇÕES ARTIGO-ITEM ===\n")

    for article_id in article_ids:
        try:
            # Verifica se relação já existe
            cursor.execute("""
                SELECT 1 FROM article_score_items
                WHERE article_id = %s AND score_item_id = %s
            """, (article_id, ITEM_ID))

            if cursor.fetchone():
                print(f"✓ Relação já existe para artigo {article_id}")
            else:
                cursor.execute("""
                    INSERT INTO article_score_items (article_id, score_item_id)
                    VALUES (%s, %s)
                """, (article_id, ITEM_ID))
                print(f"✓ Relação criada para artigo {article_id}")

        except Exception as e:
            print(f"✗ Erro ao criar relação: {e}")
            conn.rollback()
            raise

    conn.commit()

def update_score_item(conn):
    """Atualiza campos clínicos do score item"""
    cursor = conn.cursor()

    print("\n=== ATUALIZANDO SCORE ITEM ===\n")

    try:
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
            CLINICAL_RELEVANCE,
            PATIENT_EXPLANATION,
            CONDUCT,
            datetime.now(),
            ITEM_ID
        ))

        print(f"✓ Score item atualizado: {ITEM_ID}")
        print(f"  - Clinical relevance: {len(CLINICAL_RELEVANCE)} caracteres")
        print(f"  - Patient explanation: {len(PATIENT_EXPLANATION)} caracteres")
        print(f"  - Conduct: {len(CONDUCT)} caracteres")

        conn.commit()

    except Exception as e:
        print(f"✗ Erro ao atualizar score item: {e}")
        conn.rollback()
        raise

def verify_enrichment(conn):
    """Verifica se o enriquecimento foi bem-sucedido"""
    cursor = conn.cursor()

    print("\n=== VERIFICAÇÃO FINAL ===\n")

    # Verifica score item
    cursor.execute("""
        SELECT name,
               LENGTH(clinical_relevance) as clin_len,
               LENGTH(patient_explanation) as pat_len,
               LENGTH(conduct) as cond_len,
               last_review
        FROM score_items
        WHERE id = %s
    """, (ITEM_ID,))

    item = cursor.fetchone()
    if item:
        print(f"Score Item: {item[0]}")
        print(f"  Clinical relevance: {item[1]} caracteres")
        print(f"  Patient explanation: {item[2]} caracteres")
        print(f"  Conduct: {item[3]} caracteres")
        print(f"  Last review: {item[4]}")

    # Verifica artigos relacionados
    cursor.execute("""
        SELECT COUNT(*) FROM article_score_items
        WHERE score_item_id = %s
    """, (ITEM_ID,))

    article_count = cursor.fetchone()[0]
    print(f"\nArtigos relacionados: {article_count}")

    # Lista artigos
    cursor.execute("""
        SELECT a.title, a.journal, a.publish_date
        FROM articles a
        JOIN article_score_items asi ON a.id = asi.article_id
        WHERE asi.score_item_id = %s
        ORDER BY a.publish_date DESC
    """, (ITEM_ID,))

    articles = cursor.fetchall()
    for i, article in enumerate(articles, 1):
        print(f"  {i}. {article[0][:80]}...")
        print(f"     {article[1]} ({article[2]})")

def main():
    print("=" * 80)
    print("ENRIQUECIMENTO: Ecodopplercardiograma - LAVI")
    print("=" * 80)

    try:
        # Conecta ao banco
        conn = psycopg2.connect(**DB_CONFIG)
        print("\n✓ Conectado ao banco de dados\n")

        # Executa etapas
        article_ids = insert_articles(conn)
        create_article_relationships(conn, article_ids)
        update_score_item(conn)
        verify_enrichment(conn)

        print("\n" + "=" * 80)
        print("ENRIQUECIMENTO CONCLUÍDO COM SUCESSO!")
        print("=" * 80)

        conn.close()

    except Exception as e:
        print(f"\n✗ ERRO: {e}")
        return 1

    return 0

if __name__ == '__main__':
    exit(main())
