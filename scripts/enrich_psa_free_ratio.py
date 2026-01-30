#!/usr/bin/env python3
"""
Script para enriquecer PSA Livre/Total com artigos científicos e conteúdo clínico PT-BR
Item ID: 1ce8bba1-6377-46de-bba3-93a62b6950e5
Grupo: Exames > Laboratoriais
"""

import psycopg2
import os
import sys
from datetime import datetime

# Configuração do banco de dados
DB_CONFIG = {
    'host': os.getenv('DB_HOST', 'localhost'),
    'port': os.getenv('DB_PORT', '5432'),
    'database': os.getenv('DB_NAME', 'plenya_db'),
    'user': os.getenv('DB_USER', 'plenya_user'),
    'password': os.getenv('DB_PASSWORD', 'plenya_password')
}

ITEM_ID = '1ce8bba1-6377-46de-bba3-93a62b6950e5'

def get_db_connection():
    """Estabelece conexão com o banco de dados"""
    return psycopg2.connect(**DB_CONFIG)

def create_article(conn, article_data):
    """Cria um artigo científico e retorna seu ID"""
    cursor = conn.cursor()

    # Converter keywords de lista para JSONB
    import json
    keywords_json = json.dumps(article_data['keywords'])

    # publish_date no formato 'YYYY-01-01' (início do ano)
    publish_date = f"{article_data['publication_year']}-01-01"

    query = """
    INSERT INTO articles
    (title, authors, journal, publish_date, doi, original_link, abstract, keywords, article_type, created_at, updated_at)
    VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, NOW(), NOW())
    ON CONFLICT (doi) DO UPDATE
    SET title = EXCLUDED.title,
        authors = EXCLUDED.authors,
        abstract = EXCLUDED.abstract,
        keywords = EXCLUDED.keywords,
        updated_at = NOW()
    RETURNING id
    """

    cursor.execute(query, (
        article_data['title'],
        article_data['authors'],
        article_data['journal'],
        publish_date,
        article_data['doi'],
        article_data['url'],
        article_data['abstract'],
        keywords_json,
        'research_article'
    ))

    article_id = cursor.fetchone()[0]
    cursor.close()
    return article_id

def link_article_to_item(conn, article_id, score_item_id):
    """Cria relação entre artigo e score_item"""
    cursor = conn.cursor()

    query = """
    INSERT INTO article_score_items
    (score_item_id, article_id)
    VALUES (%s, %s)
    ON CONFLICT (score_item_id, article_id) DO NOTHING
    """

    cursor.execute(query, (score_item_id, article_id))
    cursor.close()

def update_score_item(conn, item_id, clinical_data):
    """Atualiza score_item com conteúdo clínico"""
    cursor = conn.cursor()

    query = """
    UPDATE score_items
    SET
        clinical_relevance = %s,
        patient_explanation = %s,
        conduct = %s,
        last_review = %s,
        updated_at = NOW()
    WHERE id = %s
    """

    cursor.execute(query, (
        clinical_data['clinical_relevance'],
        clinical_data['patient_explanation'],
        clinical_data['conduct'],
        clinical_data['last_review'],
        item_id
    ))
    cursor.close()

def main():
    """Função principal"""
    print(f"[{datetime.now().isoformat()}] Iniciando enriquecimento PSA Livre/Total")

    # Artigos científicos
    articles = [
        {
            'title': 'Usefulness of free PSA ratio to enhance detection of clinically significant prostate cancer in patients with PI-RADS<3 and PSA≤10',
            'authors': 'Heo JE, Han HH, Jang WS, Ham WS, Han WK, Choi YD, Lee J',
            'journal': 'Prostate International',
            'publication_year': 2024,
            'doi': '10.1016/j.prnil.2024.12.001',
            'url': 'https://pmc.ncbi.nlm.nih.gov/articles/PMC12223521/',
            'abstract': 'Study evaluating free PSA ratio for detecting clinically significant prostate cancer. At cutoff 17.6%, achieved sensitivity 86.5% and specificity 63.7% (AUC 0.757). Clinically significant cancer found in 34% with %fPSA <17.6% vs only 4% with %fPSA ≥17.6%.',
            'keywords': ['free PSA ratio', 'prostate cancer', 'sensitivity', 'specificity', 'PI-RADS', 'biomarker']
        },
        {
            'title': 'Actual Contribution of Free to Total PSA Ratio in Prostate Diseases Differentiation',
            'authors': 'Prcic A, Begic E, Hiros M',
            'journal': 'Medical Archives',
            'publication_year': 2016,
            'doi': '10.5455/medarh.2016.70.288-292',
            'url': 'https://pmc.ncbi.nlm.nih.gov/articles/PMC5034994/',
            'abstract': 'Study of 220 patients examining %fPSA cutoffs. At ≤16%: 72.3% sensitivity, 50.4% specificity. At <7%: 8.4% sensitivity, 97.8% specificity. Cancer patients had significantly lower %fPSA than benign prostatic hyperplasia patients.',
            'keywords': ['free PSA', 'prostate cancer', 'benign prostatic hyperplasia', 'cutoff', 'diagnostic accuracy']
        },
        {
            'title': 'Using the Free-to-total Prostate-specific Antigen Ratio to Detect Prostate Cancer in Men with Nonspecific Elevations of Prostate-specific Antigen Levels',
            'authors': 'Hoffman RM, Clanon DL, Littenberg B, Frank JJ, Peirce JC',
            'journal': 'Journal of General Internal Medicine',
            'publication_year': 2000,
            'doi': '10.1046/j.1525-1497.2000.90907.x',
            'url': 'https://pmc.ncbi.nlm.nih.gov/articles/PMC1495603/',
            'abstract': 'Meta-analysis of 21 studies on free/total PSA ratio in PSA 4.0-10.0 ng/ml gray zone. Median likelihood ratio positive 1.76, negative 0.27. At 25% pretest probability, negative test reduced posttest probability to 8%. Modest discriminating power (AUC 0.68).',
            'keywords': ['free total PSA ratio', 'meta-analysis', 'gray zone', 'likelihood ratio', 'prostate cancer screening']
        }
    ]

    # Conteúdo clínico em PT-BR
    clinical_data = {
        'clinical_relevance': '''A relação PSA livre/total (% PSA livre) é um marcador complementar fundamental para refinar a especificidade do rastreamento de câncer de próstata, especialmente na "zona cinzenta" (PSA 4-10 ng/mL).

**Princípio biológico:**
- Câncer de próstata produz predominantemente PSA ligado a proteínas (menor % PSA livre)
- Hiperplasia prostática benigna (HPB) libera proporcionalmente mais PSA livre (maior % PSA livre)

**Performance diagnóstica (estudos recentes 2024):**
- Cutoff 17.6%: Sensibilidade 86.5%, Especificidade 63.7% (AUC 0.757)
- Cutoff ≤16%: Sensibilidade 72.3%, Especificidade 50.4%
- Cutoff <7%: Especificidade 97.8% (alta certeza de malignidade)

**Interpretação por faixas:**
- **>25%**: Fortemente sugestivo de condição benigna (95% dos cânceres excluídos)
- **10-25%**: Zona cinzenta - correlacionar com toque retal, idade, história familiar
- **<10%**: Alto risco de malignidade - biópsia fortemente recomendada

**Impacto clínico:**
- Reduz biópsias desnecessárias em ~20% dos casos
- Em pacientes com PI-RADS <3 e % PSA livre ≥17.6%, apenas 4% têm câncer clinicamente significativo
- Valor preditivo negativo permite observação ao invés de biópsia imediata em casos selecionados

**Limitações:**
- Menos útil em próstatas muito aumentadas (>40-50g)
- Variabilidade inter-laboratorial (padronização de ensaio importante)
- Não substitui toque retal e avaliação clínica completa''',

        'patient_explanation': '''O **PSA Livre/Total** é um exame de sangue que ajuda o médico a entender melhor o que pode estar causando a elevação do seu PSA.

**Como funciona:**
O PSA no sangue existe em duas formas:
- **PSA livre**: circula solto no sangue
- **PSA ligado**: está preso a outras proteínas

Quando há câncer de próstata, o PSA tende a estar mais "ligado" (percentual de PSA livre baixo). Já em problemas benignos como aumento da próstata, o PSA livre é maior.

**Interpretando seu resultado:**
- **Acima de 25%**: Provavelmente benigno - indica que o aumento do PSA é mais provável ser por crescimento natural da próstata
- **Entre 10-25%**: Zona intermediária - seu médico pode pedir outros exames como ressonância magnética
- **Abaixo de 10%**: Maior preocupação - geralmente indica necessidade de biópsia para investigar possível câncer

**Por que é importante:**
Este exame ajuda a **evitar biópsias desnecessárias** em muitos homens. Se seu PSA total está elevado, mas o PSA livre também está alto (acima de 25%), há grandes chances de ser apenas crescimento benigno da próstata.

**Exemplo prático:**
Imagine dois homens com PSA total de 6 ng/mL:
- Homem A: PSA livre/total = 30% → Provavelmente benigno, observação
- Homem B: PSA livre/total = 8% → Risco aumentado, biópsia recomendada

Sempre discuta seus resultados com seu médico, que considerará também seu exame de toque, idade e histórico familiar.''',

        'conduct': '''**Protocolo de interpretação e conduta:**

**1. Indicação do exame:**
- PSA total elevado (>4.0 ng/mL) com toque retal normal ou duvidoso
- PSA na zona cinzenta (4-10 ng/mL) - maior utilidade clínica
- Paciente relutante/ansioso quanto à biópsia
- Acompanhamento de vigilância ativa

**2. Interpretação por cutoff:**

**% PSA livre ≥25%:**
- ✓ Risco baixo de malignidade (95% dos cânceres excluídos)
- ✓ Conduta: Considerar observação clínica
- ✓ Reavaliação: PSA total a cada 6-12 meses
- ✓ Toque retal anual
- ✓ Considerar RM próstata se PSA persistentemente elevado

**% PSA livre 10-25% (zona cinzenta):**
- ⚠️ Risco intermediário
- ⚠️ Conduta depende de fatores adicionais:
  - Idade <65 anos → considerar biópsia ou RM próstata
  - História familiar positiva → biópsia indicada
  - Toque retal suspeito → biópsia obrigatória
  - Paciente >70 anos sem comorbidades → discutir observação vs investigação
- ⚠️ RM multiparamétrica (mpMRI) recomendada antes da biópsia

**% PSA livre <10%:**
- ⚡ Alto risco de malignidade
- ⚡ Conduta: **Biópsia prostática indicada**
- ⚡ RM próstata pré-biópsia para guiar áreas suspeitas
- ⚡ Encaminhamento urológico prioritário

**3. Correlação com PI-RADS (quando disponível RM):**
- PI-RADS ≥3 + %fPSA <17.6% → Biópsia obrigatória (risco 34% câncer significativo)
- PI-RADS <3 + %fPSA ≥17.6% → Observação segura (risco 4%)
- PI-RADS ≥4 → Biópsia independente de %fPSA

**4. Fatores que limitam interpretação:**
- Próstata >50g (volume reduz acurácia)
- Uso de finasterida/dutasterida (ajustar PSA x2)
- Manipulação prostática recente (<48h)
- Prostatite aguda ou ITU ativa
- Ejaculação nas últimas 48h

**5. Seguimento:**
- **Observação clínica:** PSA total + % livre a cada 6 meses no 1º ano, depois anual
- **Pós-biópsia negativa:** Repetir % PSA livre + RM em 12-24 meses
- **Vigilância ativa:** % PSA livre trimestral + RM anual

**6. Documentação obrigatória:**
- Valor absoluto de PSA livre e total (não apenas %)
- Volume prostático (TRUS ou RM)
- Status de uso de 5-alfa redutase inibidores
- Data do último toque retal e achados

**Atenção:** % PSA livre **não substitui** o toque retal. Sempre correlacionar clinicamente.''',

        'last_review': datetime.now().date()
    }

    try:
        conn = get_db_connection()
        conn.autocommit = False

        # 1. Criar artigos
        article_ids = []
        for i, article in enumerate(articles, 1):
            print(f"[{datetime.now().isoformat()}] Criando artigo {i}/3: {article['title'][:60]}...")
            article_id = create_article(conn, article)
            article_ids.append(article_id)
            print(f"   ✓ Artigo criado com ID: {article_id}")

        # 2. Vincular artigos ao score_item
        print(f"\n[{datetime.now().isoformat()}] Vinculando artigos ao score_item {ITEM_ID}...")
        for article_id in article_ids:
            link_article_to_item(conn, article_id, ITEM_ID)
        print(f"   ✓ {len(article_ids)} artigos vinculados")

        # 3. Atualizar conteúdo clínico do score_item
        print(f"\n[{datetime.now().isoformat()}] Atualizando conteúdo clínico do score_item...")
        update_score_item(conn, ITEM_ID, clinical_data)
        print(f"   ✓ Conteúdo clínico atualizado")

        # Commit
        conn.commit()
        print(f"\n[{datetime.now().isoformat()}] ✅ SUCESSO! Enriquecimento concluído.")
        print(f"\n📊 Resumo:")
        print(f"   - Item atualizado: {ITEM_ID}")
        print(f"   - Artigos adicionados: {len(article_ids)}")
        print(f"   - Clinical relevance: {len(clinical_data['clinical_relevance'])} caracteres")
        print(f"   - Patient explanation: {len(clinical_data['patient_explanation'])} caracteres")
        print(f"   - Conduct: {len(clinical_data['conduct'])} caracteres")

        conn.close()
        return 0

    except Exception as e:
        if conn:
            conn.rollback()
            conn.close()
        print(f"\n❌ ERRO: {str(e)}", file=sys.stderr)
        import traceback
        traceback.print_exc()
        return 1

if __name__ == '__main__':
    sys.exit(main())
