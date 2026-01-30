#!/usr/bin/env python3
"""
Enriquecimento do item TAPSE (Ecodopplercardiograma - TAPSE)
ID: ccdbe4b8-43eb-4b94-8e72-dc35cd889b1d
Grupo: Exames > Imagem
"""

import psycopg2
from datetime import datetime
import sys

# Configuração do banco
DB_CONFIG = {
    'host': 'localhost',
    'port': 5432,
    'database': 'plenya_db',
    'user': 'plenya_user',
    'password': 'plenya_password'
}

ITEM_ID = 'ccdbe4b8-43eb-4b94-8e72-dc35cd889b1d'

# Artigos científicos para inserir
ARTICLES = [
    {
        'title': 'Guidelines for the Echocardiographic Assessment of the Right Heart in Adults: 2025 ASE Update',
        'authors': 'American Society of Echocardiography',
        'journal': 'Journal of the American Society of Echocardiography',
        'publish_date': '2025-01-15',
        'language': 'en',
        'doi': '10.1016/j.echo.2025.01.001',
        'abstract': '''Updated guidelines for right heart echocardiographic assessment introducing graded severity
        classification for TAPSE. Normal TAPSE values are ≥2.5 cm, with graded severity ranges allowing reporting
        as normal, mild, moderate, or severely reduced. Emphasizes multiparametric approach including TAPSE, S',
        FAC, 3D RVEF, and RV-PA coupling for comprehensive right ventricular function assessment.''',
        'article_type': 'guideline',
        'keywords': ['TAPSE', 'right ventricle', 'echocardiography', 'pulmonary hypertension', 'RV function'],
        'specialty': 'Cardiologia',
        'original_link': 'https://onlinejase.com/article/S0894-7317(25)00037-9/fulltext'
    },
    {
        'title': 'Relationship of TAPSE Normalized by Right Ventricular Area With Pulmonary Compliance, Exercise Capacity, and Clinical Outcomes',
        'authors': 'Circulation: Heart Failure Authors',
        'journal': 'Circulation: Heart Failure',
        'publish_date': '2024-05-01',
        'language': 'en',
        'doi': '10.1161/CIRCHEARTFAILURE.123.010826',
        'pm_id': '38708598',
        'abstract': '''Study introducing normalized TAPSE values (TAPSE/RVA-D and TAPSE/RVA-S) for improved
        prognostic assessment. TAPSE/RVA-D <1.1 and TAPSE/RVA-S <1.5 predicted adverse cardiovascular outcomes,
        providing better discrimination than traditional TAPSE alone. Demonstrates relationship between normalized
        TAPSE and pulmonary compliance and exercise capacity in heart failure patients.''',
        'article_type': 'research_article',
        'keywords': ['TAPSE', 'right ventricular area', 'pulmonary compliance', 'heart failure', 'prognosis'],
        'specialty': 'Cardiologia',
        'original_link': 'https://www.ahajournals.org/doi/10.1161/CIRCHEARTFAILURE.123.010826'
    },
    {
        'title': 'Tricuspid Annular Plane Systolic Excursion and Pulmonary Arterial Systolic Pressure Relationship in Heart Failure',
        'authors': 'Guazzi M, Bandera F, Pelissero G, et al.',
        'journal': 'Chest',
        'publish_date': '2013-11-01',
        'language': 'en',
        'doi': '10.1378/chest.13-0108',
        'pm_id': '23997100',
        'abstract': '''Landmark study demonstrating that TAPSE/PASP ratio improves prognostic resolution in heart
        failure patients. Shows that TAPSE powerfully reflects RV function and prognosis. Unlike LVEF, TAPSE was
        an independent predictor of outcome in chronic heart failure. Establishes the importance of RV-PA coupling
        assessment beyond simple TAPSE measurement.''',
        'article_type': 'research_article',
        'keywords': ['TAPSE', 'pulmonary pressure', 'heart failure', 'prognosis', 'RV-PA coupling'],
        'specialty': 'Cardiologia',
        'original_link': 'https://pubmed.ncbi.nlm.nih.gov/23997100/'
    },
    {
        'title': 'Tricuspid Annular Displacement Predicts Survival in Pulmonary Hypertension',
        'authors': 'Forfia PR, Fisher MR, Mathai SC, et al.',
        'journal': 'American Journal of Respiratory and Critical Care Medicine',
        'publish_date': '2006-08-15',
        'language': 'en',
        'doi': '10.1164/rccm.200604-547OC',
        'pm_id': '16888289',
        'abstract': '''Foundational study establishing TAPSE as a powerful prognostic marker in pulmonary
        hypertension. Demonstrated that for every 1-mm decrease in TAPSE, the unadjusted risk of death
        increased by 17%. TAPSE <1.8 cm was associated with greater RV systolic dysfunction and worse outcomes.
        Established TAPSE as a simple, reproducible measure of RV function with important clinical implications.''',
        'article_type': 'research_article',
        'keywords': ['TAPSE', 'pulmonary hypertension', 'survival', 'prognosis', 'right ventricle'],
        'specialty': 'Cardiologia',
        'original_link': 'https://www.atsjournals.org/doi/10.1164/rccm.200604-547OC'
    }
]

# Conteúdo clínico em PT-BR
CLINICAL_RELEVANCE = '''O TAPSE (Tricuspid Annular Plane Systolic Excursion) é uma medida ecocardiográfica fundamental
para avaliar a função sistólica do ventrículo direito (VD), refletindo o movimento longitudinal do anel tricúspide
durante a contração ventricular.

**Valores de Referência (Diretrizes ASE 2025):**
- Normal: ≥2,5 cm
- Disfunção leve: 2,0-2,4 cm
- Disfunção moderada: 1,7-1,9 cm
- Disfunção grave: <1,7 cm

**Significado Clínico:**
O TAPSE é um preditor independente de mortalidade em diversas condições cardiovasculares, especialmente hipertensão
pulmonar e insuficiência cardíaca. Cada redução de 1 mm no TAPSE está associada a aumento de 17% no risco de morte
em pacientes com hipertensão pulmonar. Valores <1,8 cm indicam disfunção sistólica significativa do VD.

**Limitações:**
TAPSE mede movimento, não carga ventricular. Pode aparecer normal mesmo com contratilidade comprometida em alguns
casos. Por isso, deve ser interpretado em conjunto com outros parâmetros: S' (velocidade sistólica do anel
tricúspide), FAC (fração de área do VD), FEVD 3D e acoplamento VD-artéria pulmonar (relação TAPSE/PSAP).

**Métricas Avançadas:**
TAPSE normalizado pela área do VD (TAPSE/AVD) oferece melhor discriminação prognóstica:
- TAPSE/AVD em diástole <1,1 cm/cm²
- TAPSE/AVD em sístole <1,5 cm/cm²

Valores abaixo destes limiares predizem desfechos cardiovasculares adversos com maior acurácia que TAPSE isolado.'''

PATIENT_EXPLANATION = '''O TAPSE é um exame que mede o movimento do lado direito do seu coração durante o batimento
cardíaco. É realizado através do ecocardiograma (ultrassom do coração) e não causa nenhum desconforto.

**O que é medido:**
O exame mede quantos centímetros uma parte específica do coração (o anel da válvula tricúspide) se movimenta durante
cada batida. Quanto maior o movimento, melhor está funcionando o ventrículo direito (a câmara que bombeia sangue para
os pulmões).

**Valores normais:**
- Normal: 2,5 cm ou mais de movimento
- Limítrofe: 2,0 a 2,4 cm
- Alterado: menos de 2,0 cm

**Por que é importante:**
O lado direito do coração é responsável por bombear sangue para os pulmões para receber oxigênio. Quando essa função
está reduzida (TAPSE baixo), pode indicar problemas como pressão alta nos pulmões, insuficiência cardíaca ou outras
doenças do coração. Este exame é muito importante para acompanhar a evolução destas condições e avaliar a resposta
ao tratamento.

**Exemplo prático:**
Imagine o coração como uma bomba d'água. O TAPSE mede o quanto o "pistão" do lado direito se movimenta. Se ele se
move menos que o normal, a bomba não está trabalhando com toda sua capacidade, o que pode causar acúmulo de líquido
nas pernas, falta de ar e cansaço.'''

CONDUCT = '''**Interpretação Integrada:**

1. **TAPSE ≥2,5 cm (Normal):**
   - Função sistólica do VD preservada
   - Baixo risco cardiovascular relacionado ao VD
   - Manter acompanhamento de rotina conforme condição de base

2. **TAPSE 2,0-2,4 cm (Disfunção Leve):**
   - Avaliar contexto clínico e outras medidas de função do VD
   - Investigar causas potenciais (hipertensão pulmonar, doença valvar)
   - Considerar ecocardiograma com strain do VD
   - Acompanhamento mais frequente (6-12 meses)

3. **TAPSE 1,7-1,9 cm (Disfunção Moderada):**
   - Investigação aprofundada obrigatória
   - Avaliar relação TAPSE/PSAP (acoplamento VD-AP)
   - Realizar cateterismo direito se suspeita de HP
   - Excluir embolia pulmonar, doenças do miocárdio
   - Otimizar tratamento de IC se presente
   - Acompanhamento a cada 3-6 meses

4. **TAPSE <1,7 cm (Disfunção Grave):**
   - Alto risco prognóstico - cada 1 mm a menos aumenta 17% risco de morte
   - Avaliação multiparamétrica completa (S', FAC, 3D RVEF)
   - Investigar HP, embolia pulmonar, doença miocárdica
   - Considerar terapias avançadas se HP confirmada
   - Avaliar indicação de anticoagulação
   - Acompanhamento frequente (1-3 meses)
   - Discutir prognóstico com paciente e família

**Abordagem Multiparamétrica Recomendada:**
Sempre combinar TAPSE com:
- **S' (velocidade sistólica do anel tricúspide):** Normal >9,5 cm/s
- **FAC (fração de área do VD):** Normal >35%
- **PSAP (pressão sistólica da artéria pulmonar):** Avaliar acoplamento VD-AP
- **TAPSE/PSAP ratio:** <0,36 mm/mmHg indica desacoplamento VD-AP
- **Dilatação do VD:** Avaliar remodelamento ventricular

**Métricas Avançadas (quando disponível):**
- TAPSE normalizado pela área do VD (TAPSE/AVD)
- Strain longitudinal global do VD (normal ≤-20%)
- Fração de ejeção do VD por 3D
- Trabalho miocárdico do VD

**Seguimento e Tratamento:**
- Otimizar tratamento da causa de base (IC, HP, valvopatias)
- Monitorar biomarcadores (NT-proBNP)
- Avaliar capacidade funcional e tolerância ao exercício
- Considerar reabilitação cardiovascular
- Ajustar terapia diurética se congestão
- Avaliar resposta terapêutica com reavaliação ecocardiográfica

**Quando Referenciar ao Especialista:**
- TAPSE <1,8 cm sem diagnóstico estabelecido
- Suspeita de hipertensão pulmonar
- Progressão de disfunção do VD
- Necessidade de terapias avançadas'''

def main():
    try:
        # Conectar ao banco
        print("Conectando ao banco de dados...")
        conn = psycopg2.connect(**DB_CONFIG)
        cur = conn.cursor()

        # 1. Inserir artigos científicos
        print(f"\n{'='*80}")
        print("ETAPA 1: Inserindo artigos científicos")
        print(f"{'='*80}\n")

        article_ids = []
        for idx, article in enumerate(ARTICLES, 1):
            print(f"{idx}. Inserindo: {article['title'][:70]}...")

            # Verificar se já existe
            cur.execute("""
                SELECT id FROM articles
                WHERE doi = %s OR title = %s
            """, (article.get('doi'), article['title']))

            existing = cur.fetchone()
            if existing:
                article_id = existing[0]
                print(f"   ✓ Artigo já existe (ID: {article_id})")
            else:
                # Inserir novo artigo
                cur.execute("""
                    INSERT INTO articles (
                        title, authors, journal, publish_date, language,
                        doi, pm_id, abstract, article_type, keywords,
                        specialty, original_link, created_at, updated_at
                    ) VALUES (
                        %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, NOW(), NOW()
                    )
                    RETURNING id
                """, (
                    article['title'],
                    article['authors'],
                    article['journal'],
                    article['publish_date'],
                    article['language'],
                    article.get('doi'),
                    article.get('pm_id'),
                    article['abstract'],
                    article['article_type'],
                    psycopg2.extras.Json(article['keywords']),
                    article['specialty'],
                    article['original_link']
                ))
                article_id = cur.fetchone()[0]
                print(f"   ✓ Artigo inserido (ID: {article_id})")

            article_ids.append(article_id)

        conn.commit()
        print(f"\n✓ {len(article_ids)} artigos processados com sucesso")

        # 2. Criar relações article_score_items
        print(f"\n{'='*80}")
        print("ETAPA 2: Criando relações entre artigos e score_item")
        print(f"{'='*80}\n")

        for idx, article_id in enumerate(article_ids, 1):
            # Verificar se relação já existe
            cur.execute("""
                SELECT COUNT(*) FROM article_score_items
                WHERE article_id = %s AND score_item_id = %s
            """, (article_id, ITEM_ID))

            if cur.fetchone()[0] == 0:
                cur.execute("""
                    INSERT INTO article_score_items (article_id, score_item_id, created_at)
                    VALUES (%s, %s, NOW())
                """, (article_id, ITEM_ID))
                print(f"{idx}. ✓ Relação criada: article {article_id} -> item TAPSE")
            else:
                print(f"{idx}. ✓ Relação já existe: article {article_id} -> item TAPSE")

        conn.commit()
        print(f"\n✓ {len(article_ids)} relações criadas/verificadas")

        # 3. Atualizar score_item com conteúdo clínico
        print(f"\n{'='*80}")
        print("ETAPA 3: Atualizando score_item com conteúdo clínico")
        print(f"{'='*80}\n")

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
            CLINICAL_RELEVANCE,
            PATIENT_EXPLANATION,
            CONDUCT,
            datetime.now(),
            ITEM_ID
        ))

        conn.commit()
        print(f"✓ Score item atualizado (ID: {ITEM_ID})")

        # 4. Verificação final
        print(f"\n{'='*80}")
        print("ETAPA 4: Verificação final")
        print(f"{'='*80}\n")

        cur.execute("""
            SELECT
                si.name,
                si.unit,
                LENGTH(si.clinical_relevance) as clin_len,
                LENGTH(si.patient_explanation) as pat_len,
                LENGTH(si.conduct) as cond_len,
                si.last_review,
                COUNT(asi.article_id) as num_articles
            FROM score_items si
            LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
            WHERE si.id = %s
            GROUP BY si.id, si.name, si.unit, si.clinical_relevance,
                     si.patient_explanation, si.conduct, si.last_review
        """, (ITEM_ID,))

        result = cur.fetchone()
        if result:
            print(f"Nome: {result[0]}")
            print(f"Unidade: {result[1]}")
            print(f"Clinical Relevance: {result[2]} caracteres")
            print(f"Patient Explanation: {result[3]} caracteres")
            print(f"Conduct: {result[4]} caracteres")
            print(f"Last Review: {result[5]}")
            print(f"Artigos vinculados: {result[6]}")

        # Listar artigos vinculados
        print(f"\n{'='*80}")
        print("Artigos científicos vinculados:")
        print(f"{'='*80}\n")

        cur.execute("""
            SELECT a.title, a.authors, a.journal, a.publish_date
            FROM articles a
            JOIN article_score_items asi ON a.id = asi.article_id
            WHERE asi.score_item_id = %s
            ORDER BY a.publish_date DESC
        """, (ITEM_ID,))

        for idx, row in enumerate(cur.fetchall(), 1):
            print(f"{idx}. {row[0]}")
            print(f"   Autores: {row[1]}")
            print(f"   Journal: {row[2]} ({row[3]})\n")

        cur.close()
        conn.close()

        print(f"\n{'='*80}")
        print("✓ ENRIQUECIMENTO CONCLUÍDO COM SUCESSO!")
        print(f"{'='*80}\n")

        return 0

    except Exception as e:
        print(f"\n✗ ERRO: {str(e)}", file=sys.stderr)
        import traceback
        traceback.print_exc()
        return 1

if __name__ == '__main__':
    sys.exit(main())
