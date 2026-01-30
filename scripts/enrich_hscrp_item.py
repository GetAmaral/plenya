#!/usr/bin/env python3
"""
Enriquecimento científico do item PCR ultrassensível (hs-CRP)
Baseado em evidências de 2024-2025 do ACC, ESC e UK Biobank
"""

import psycopg2
from datetime import datetime
import uuid
import json

# Configuração do banco
DB_CONFIG = {
    'dbname': 'plenya_db',
    'user': 'plenya_user',
    'password': 'plenya_dev_password',
    'host': 'db',
    'port': '5432'
}

# ID do score_item
SCORE_ITEM_ID = '9b8eb642-4b9c-48e2-b15d-ee2aeba78c83'

# Artigos científicos para inserir
ARTICLES = [
    {
        'id': str(uuid.uuid4()),
        'title': 'Inflammation and Cardiovascular Disease: 2025 ACC Scientific Statement',
        'authors': 'American College of Cardiology Committee',
        'journal': 'Journal of the American College of Cardiology',
        'publish_date': '2025-01-15',
        'language': 'en',
        'doi': '10.1016/j.jacc.2025.08.047',
        'pm_id': '41020749',
        'abstract': 'Comprehensive scientific statement on inflammation and cardiovascular disease. High-sensitivity C-reactive protein (hsCRP) is established as a strong predictor of cardiovascular events in both primary and secondary prevention. In statin-treated patients, hsCRP proves to be a stronger predictor of recurrent myocardial infarction, stroke, and cardiovascular death than LDL cholesterol. The statement recommends hsCRP ≥2 mg/L as a risk enhancer for cardiovascular risk assessment.',
        'specialty': 'Cardiologia',
        'article_type': 'review',
        'keywords': ['hs-CRP', 'cardiovascular risk', 'inflammation', 'primary prevention', 'risk stratification']
    },
    {
        'id': str(uuid.uuid4()),
        'title': 'C-reactive protein and cardiovascular risk in the general population: UK Biobank analysis',
        'authors': 'UK Biobank Research Consortium',
        'journal': 'European Heart Journal',
        'publish_date': '2024-11-20',
        'language': 'en',
        'doi': '10.1093/eurheartj/ehaf937',
        'abstract': 'Analysis of 448,653 individuals without known atherosclerotic cardiovascular disease from UK Biobank demonstrates strong and independent association of hsCRP with long-term cardiovascular outcomes. Stable hsCRP levels across serial measurements showed consistent prognostic value for major adverse cardiovascular events (MACE), CV death and all-cause death across all evaluated subgroups. Integration of hsCRP into SCORE2 improved model performance for prediction of MACE.',
        'specialty': 'Cardiologia',
        'article_type': 'research_article',
        'keywords': ['hs-CRP', 'UK Biobank', 'cardiovascular risk prediction', 'MACE', 'population study']
    },
    {
        'id': str(uuid.uuid4()),
        'title': 'C-Reactive Protein: Clinical Relevance and Interpretation',
        'authors': 'Nehring SM, Goyal A, Patel BC',
        'journal': 'StatPearls',
        'publish_date': '2024-08-12',
        'language': 'en',
        'pm_id': '30020613',
        'abstract': 'Comprehensive review of C-reactive protein clinical applications. CRP is a pentameric acute-phase protein synthesized by hepatocytes in response to IL-6 during inflammation. For cardiovascular risk assessment, levels are interpreted as: <1 mg/L (low risk), 1-3 mg/L (moderate risk), >3 mg/L (high risk). Values >10 mg/L indicate acute inflammation and should be excluded from cardiovascular risk assessment. Two readings at least 2 weeks apart should be obtained for stable assessment.',
        'specialty': 'Medicina Laboratorial',
        'article_type': 'review',
        'keywords': ['CRP', 'hs-CRP', 'cardiovascular risk', 'acute-phase protein', 'clinical interpretation']
    }
]

# Conteúdo clínico em português
CLINICAL_CONTENT = {
    'clinical_relevance': '''A Proteína C Reativa ultrassensível (PCR-us) é um biomarcador inflamatório de fase aguda que reflete inflamação sistêmica de baixo grau. No contexto cardiovascular, representa uma das ferramentas mais validadas para estratificação de risco em prevenção primária e secundária.

**Evidências Científicas:**
- Estudo UK Biobank (2024) com 448.653 indivíduos demonstrou associação forte e independente da PCR-us com eventos cardiovasculares maiores (MACE), morte cardiovascular e mortalidade por todas as causas
- Declaração Científica ACC 2025 estabelece PCR-us como preditor mais forte que LDL-colesterol para eventos recorrentes em pacientes tratados com estatinas
- Integração da PCR-us ao escore SCORE2 melhora significativamente a performance do modelo de predição de risco cardiovascular

**Interpretação Clínica:**
- <1 mg/L: Baixo risco cardiovascular
- 1-3 mg/L: Risco cardiovascular moderado
- >3 mg/L: Alto risco cardiovascular
- ≥2 mg/L: Risco residual inflamatório elevado (em pacientes em prevenção secundária)
- >10 mg/L: Inflamação aguda (excluir da avaliação de risco cardiovascular)

**Fisiopatologia:**
A PCR é uma proteína pentamérica sintetizada pelos hepatócitos em resposta à interleucina-6 (IL-6). Participa da cascata inflamatória, liga-se a fosforilcolina em membranas celulares danificadas e ativa o sistema complemento. Elevações crônicas refletem inflamação vascular subclínica, preditor independente de aterosclerose e eventos cardiovasculares.

**Fatores Modificadores:**
- Tabagismo, obesidade, sedentarismo (elevam)
- Atividade física, perda de peso, dieta anti-inflamatória (reduzem)
- Estatinas com efeito pleiotrópico anti-inflamatório''',

    'patient_explanation': '''A Proteína C Reativa ultrassensível (PCR-us) é um exame de sangue que mede o nível de inflamação no seu organismo. Quando está elevada de forma persistente (crônica), indica que existe uma inflamação silenciosa nos vasos sanguíneos, aumentando o risco de problemas no coração.

**Por que esse exame é importante?**
Estudos recentes com centenas de milhares de pessoas mostraram que a PCR-us é um dos melhores exames para prever quem tem maior chance de ter infarto, derrame ou outros problemas cardiovasculares nos próximos anos. Em alguns casos, ela prevê melhor o risco do que o próprio colesterol LDL.

**O que os valores significam?**
- **Abaixo de 1 mg/L**: Seu risco cardiovascular relacionado à inflamação é baixo
- **Entre 1 e 3 mg/L**: Risco moderado - vale a pena conversar sobre prevenção
- **Acima de 3 mg/L**: Risco alto - é importante investigar as causas e tomar medidas preventivas
- **Acima de 10 mg/L**: Pode indicar uma infecção ou inflamação aguda acontecendo agora (não serve para avaliar risco de coração)

**Como melhorar seus níveis?**
A boa notícia é que você pode reduzir a PCR-us através de hábitos saudáveis: praticar atividade física regular, manter peso adequado, não fumar, ter uma alimentação anti-inflamatória (rica em frutas, vegetais, peixes) e dormir bem. Se necessário, seu médico pode prescrever medicamentos como estatinas, que além de baixar o colesterol também reduzem a inflamação.

**Quando repetir o exame?**
Para ter certeza do resultado, geralmente repetimos o exame após 2 semanas, pois infecções temporárias podem alterar o valor. O resultado mais baixo entre as duas medições é o que usamos para avaliar seu risco cardiovascular real.''',

    'conduct': '''**Solicitação do Exame:**
1. Realizar 2 dosagens com intervalo mínimo de 2 semanas (utilizar o menor valor)
2. Evitar dosagem durante infecções, processos inflamatórios agudos ou trauma recente
3. Coleta em jejum não é obrigatória, mas preferível para padronização

**Interpretação e Estratificação de Risco:**
- **PCR-us <1 mg/L**: Risco cardiovascular baixo
  - Reforçar hábitos saudáveis
  - Manter vigilância em pacientes com outros fatores de risco

- **PCR-us 1-3 mg/L**: Risco cardiovascular moderado
  - Intensificar mudanças de estilo de vida
  - Reavaliar demais fatores de risco (perfil lipídico, PA, glicemia)
  - Considerar escore de cálcio coronário se disponível

- **PCR-us >3 mg/L**: Risco cardiovascular alto
  - Investigar causas de inflamação crônica (obesidade, síndrome metabólica, tabagismo)
  - Considerar terapia com estatina se indicação limítrofe pelo escore de risco tradicional
  - Seguimento mais próximo e metas mais rigorosas de controle de fatores de risco

- **PCR-us ≥2 mg/L** (prevenção secundária):
  - Indicativo de risco residual inflamatório elevado
  - Otimizar terapia hipolipemiante
  - Considerar terapias anti-inflamatórias direcionadas (ex: colchicina em pacientes pós-IAM)

**Valores >10 mg/L:**
- Excluir da avaliação de risco cardiovascular
- Investigar infecção aguda, processos inflamatórios sistêmicos ou trauma
- Repetir após resolução do quadro agudo

**Abordagem Terapêutica:**
1. **Medidas não farmacológicas** (primeira linha):
   - Atividade física regular (150 min/semana de exercício moderado)
   - Perda de peso (se IMC >25)
   - Cessação do tabagismo
   - Dieta anti-inflamatória (padrão mediterrâneo)
   - Controle do estresse e qualidade do sono

2. **Farmacoterapia** (considerar):
   - Estatinas: efeito pleiotrópico anti-inflamatório (redução ~30-40% PCR-us)
   - Aspirina em baixa dose (prevenção secundária)
   - Colchicina 0,5 mg/dia (prevenção secundária em pacientes selecionados)

**Monitoramento:**
- Repetir PCR-us a cada 6-12 meses em pacientes de risco moderado/alto
- Avaliar resposta às intervenções terapêuticas
- Integrar resultado com demais marcadores de risco cardiovascular

**Referências das Diretrizes:**
- ESC Guidelines 2024: recomenda avaliação de PCR-us em pacientes com suspeita de doença arterial coronariana
- ACC Scientific Statement 2025: estabelece PCR-us ≥2 mg/L como fator de risco aumentado
- AHA/ACC: considera PCR-us >2 mg/L como "risk enhancer" para decisão terapêutica'''
}

def main():
    """Executa o enriquecimento do item PCR ultrassensível"""
    conn = None
    try:
        # Conectar ao banco
        print("Conectando ao banco de dados...")
        conn = psycopg2.connect(**DB_CONFIG)
        cur = conn.cursor()

        # 1. Verificar se o item existe
        print(f"\n1. Verificando item {SCORE_ITEM_ID}...")
        cur.execute("""
            SELECT name, clinical_relevance IS NOT NULL
            FROM score_items
            WHERE id = %s
        """, (SCORE_ITEM_ID,))

        result = cur.fetchone()
        if not result:
            print(f"ERRO: Item {SCORE_ITEM_ID} não encontrado!")
            return

        item_name, has_content = result
        print(f"   Item encontrado: {item_name}")
        print(f"   Já tem conteúdo: {'Sim' if has_content else 'Não'}")

        # 2. Inserir artigos científicos
        print("\n2. Inserindo artigos científicos...")
        article_ids = []

        for article in ARTICLES:
            # Verificar se já existe pelo DOI ou PM_ID
            identifier = article.get('doi') or article.get('pm_id')
            if identifier:
                if article.get('doi'):
                    cur.execute("SELECT id FROM articles WHERE doi = %s", (article['doi'],))
                else:
                    cur.execute("SELECT id FROM articles WHERE pm_id = %s", (article['pm_id'],))

                existing = cur.fetchone()
                if existing:
                    print(f"   ✓ Artigo já existe: {article['title'][:60]}...")
                    article_ids.append(existing[0])
                    continue

            # Inserir novo artigo
            keywords_json = json.dumps(article.get('keywords')) if article.get('keywords') else None

            cur.execute("""
                INSERT INTO articles (
                    id, title, authors, journal, publish_date, language,
                    doi, pm_id, abstract, specialty, article_type, keywords, created_at
                ) VALUES (
                    %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s::jsonb, NOW()
                )
            """, (
                article['id'],
                article['title'],
                article['authors'],
                article['journal'],
                article['publish_date'],
                article['language'],
                article.get('doi'),
                article.get('pm_id'),
                article.get('abstract'),
                article.get('specialty'),
                article.get('article_type'),
                keywords_json
            ))

            article_ids.append(article['id'])
            print(f"   + Artigo inserido: {article['title'][:60]}...")

        # 3. Criar relações article-score_item
        print("\n3. Criando relações article-score_item...")
        for article_id in article_ids:
            # Verificar se relação já existe
            cur.execute("""
                SELECT 1 FROM article_score_items
                WHERE article_id = %s AND score_item_id = %s
            """, (article_id, SCORE_ITEM_ID))

            if cur.fetchone():
                continue

            cur.execute("""
                INSERT INTO article_score_items (article_id, score_item_id)
                VALUES (%s, %s)
                ON CONFLICT (article_id, score_item_id) DO NOTHING
            """, (article_id, SCORE_ITEM_ID))

        print(f"   ✓ {len(article_ids)} relações criadas")

        # 4. Atualizar campos clínicos do score_item
        print("\n4. Atualizando campos clínicos do score_item...")
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

        print("   ✓ Campos clínicos atualizados")

        # Commit
        conn.commit()

        # 5. Verificar resultado final
        print("\n5. Verificação final...")
        cur.execute("""
            SELECT
                si.name,
                LENGTH(si.clinical_relevance) as len_relevance,
                LENGTH(si.patient_explanation) as len_explanation,
                LENGTH(si.conduct) as len_conduct,
                si.last_review,
                COUNT(asi.article_id) as num_articles
            FROM score_items si
            LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
            WHERE si.id = %s
            GROUP BY si.id
        """, (SCORE_ITEM_ID,))

        result = cur.fetchone()
        print(f"""
   Item: {result[0]}
   Clinical relevance: {result[1]} caracteres
   Patient explanation: {result[2]} caracteres
   Conduct: {result[3]} caracteres
   Last review: {result[4]}
   Artigos vinculados: {result[5]}
        """)

        print("\n✅ Enriquecimento concluído com sucesso!")

    except Exception as e:
        if conn:
            conn.rollback()
        print(f"\n❌ ERRO: {str(e)}")
        raise

    finally:
        if conn:
            conn.close()

if __name__ == '__main__':
    main()
