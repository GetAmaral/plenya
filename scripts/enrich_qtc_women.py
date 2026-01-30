#!/usr/bin/env python3
"""
Enriquecimento do item ECG - QTc (Bazett) - Mulheres
ID: 2e3c06ce-bcb6-4649-984e-8c30a92e26f4

Baseado em pesquisa científica:
- Salama G, Bett GCL. Sex differences in the mechanisms underlying long QT syndrome. Am J Physiol Heart Circ Physiol. 2014.
- Rabkin SW. Impact of Age and Sex on QT Prolongation in Patients Receiving Psychotropics. Can J Psychiatry. 2015.
"""

import psycopg2
from datetime import datetime
import os

# Configuração do banco de dados
DB_CONFIG = {
    'host': os.getenv('DB_HOST', 'localhost'),
    'port': os.getenv('DB_PORT', '5432'),
    'database': os.getenv('DB_NAME', 'plenya_db'),
    'user': os.getenv('DB_USER', 'plenya_user'),
    'password': os.getenv('DB_PASSWORD', 'plenya_password')
}

ITEM_ID = '2e3c06ce-bcb6-4649-984e-8c30a92e26f4'

# Artigos científicos relacionados
ARTICLES = [
    {
        'title': 'Sex differences in the mechanisms underlying long QT syndrome',
        'authors': 'Guy Salama, Glenna C L Bett',
        'journal': 'American Journal of Physiology - Heart and Circulatory Physiology',
        'year': 2014,
        'doi': '10.1152/ajpheart.00864.2013',
        'url': 'https://pmc.ncbi.nlm.nih.gov/articles/PMC4187395/'
    },
    {
        'title': 'Impact of Age and Sex on QT Prolongation in Patients Receiving Psychotropics',
        'authors': 'Simon W Rabkin',
        'journal': 'Canadian Journal of Psychiatry',
        'year': 2015,
        'doi': '10.1177/070674371506000502',
        'url': 'https://pmc.ncbi.nlm.nih.gov/articles/PMC4484689/'
    }
]

# Conteúdo clínico em português
CLINICAL_CONTENT = {
    'clinical_relevance': """O intervalo QTc (QT corrigido pela frequência cardíaca) representa o tempo de repolarização ventricular no eletrocardiograma. Mulheres apresentam QTc naturalmente mais prolongado que homens após a puberdade, devido à influência hormonal do estrogênio sobre os canais de cálcio cardíacos.

**Valores de referência:**
- Normal em mulheres: <460 ms
- Limítrofe: 460-470 ms
- Prolongado: >470 ms

**Importância clínica:**
O QTc prolongado em mulheres está associado a maior risco de:
- Torsades de Pointes (arritmia ventricular polimórfica)
- Morte súbita cardíaca
- Eventos arrítmicos induzidos por medicamentos

Mulheres representam 70% dos casos de Torsades de Pointes associados ao uso de medicamentos cardiovasculares que prolongam o QT. Para cada aumento de 10 ms no QTc, há um incremento de ~5% no risco de eventos arrítmicos.

**Fatores de risco adicionais:**
- Idade >65 anos
- Medicamentos que prolongam QT (antiarrítmicos, psicotrópicos, antibióticos)
- Hipocalemia, hipomagnesemia
- Bradicardia
- História familiar de síndrome do QT longo

**Mecanismos fisiopatológicos:**
O estrogênio aumenta a expressão de canais de cálcio tipo-L no miocárdio ventricular feminino, criando heterogeneidades regionais que facilitam arritmias. A testosterona, por outro lado, encurta o intervalo QT.""",

    'patient_explanation': """O QTc é uma medida do tempo que o coração leva para se recarregar eletricamente entre um batimento e outro. É avaliado através do eletrocardiograma (ECG).

**Por que mulheres têm valores diferentes?**
Após a puberdade, as mulheres naturalmente apresentam QTc um pouco mais longo que os homens, devido à influência dos hormônios femininos (especialmente o estrogênio) no coração.

**O que significam os valores:**
- **Normal:** Abaixo de 460 milissegundos
- **Limítrofe:** Entre 460-470 ms (requer atenção)
- **Prolongado:** Acima de 470 ms (maior risco)

**Por que é importante monitorar:**
Um QTc prolongado aumenta o risco de uma arritmia cardíaca perigosa chamada "Torsades de Pointes", que pode causar desmaios ou, em casos graves, morte súbita.

**Fatores que podem piorar:**
- Alguns medicamentos (antiarrítmicos, antibióticos, antidepressivos)
- Falta de potássio ou magnésio no sangue
- Idade acima de 65 anos
- Frequência cardíaca muito baixa

**Cuidados especiais:**
Informe sempre seu médico sobre todos os medicamentos que usa, pois alguns podem aumentar ainda mais o QTc. Evite combinar múltiplos remédios que prolongam o intervalo QT sem orientação médica.""",

    'conduct': """**Avaliação inicial:**
1. Calcular QTc pela fórmula de Bazett: QTc = QT / √RR
2. Confirmar medida em múltiplas derivações (especialmente DII e V5)
3. Repetir ECG se ritmo irregular ou dúvida na medição
4. Revisar histórico medicamentoso completo
5. Solicitar eletrólitos (K+, Mg2+, Ca2+)

**Condutas por faixa de valor:**

**QTc <460 ms (Normal):**
- Manter monitoramento de rotina
- Revisar periodicamente se em uso de medicamentos que prolongam QT
- ECG anual ou conforme indicação clínica

**QTc 460-470 ms (Limítrofe):**
- Investigar causas reversíveis (medicamentos, distúrbios eletrolíticos)
- Corrigir hipocalemia (<4.0 mEq/L) e hipomagnesemia (<2.0 mg/dL)
- Revisar medicações: considerar alternativas se possível
- Repetir ECG em 1-3 meses
- Considerar Holter 24h se sintomas (palpitações, síncopes)
- Avaliar história familiar de morte súbita ou QT longo congênito

**QTc >470 ms (Prolongado):**
- **Suspender imediatamente** medicamentos que prolongam QT (se possível)
- Corrigir urgentemente distúrbios eletrolíticos
- Solicitar Holter 24h e ecocardiograma
- Avaliar necessidade de teste ergométrico (pode haver QT longo congênito)
- Considerar teste genético se:
  - História familiar positiva
  - QTc >480 ms sem causa aparente
  - Episódios de síncope inexplicada
- **Referência ao cardiologista/arritmologista** obrigatória

**QTc >500 ms (Alto risco):**
- **Emergência cardiológica**
- Internação hospitalar para monitorização contínua
- Suspensão de todos medicamentos que prolongam QT
- Reposição agressiva de K+ (alvo 4.5-5.0 mEq/L) e Mg2+ (>2.0 mg/dL)
- Considerar betabloqueador (propranolol, nadolol)
- Avaliar necessidade de cardioversor-desfibrilador implantável (CDI)

**Monitoramento de medicamentos de risco:**
- Antiarrítmicos: amiodarona, sotalol, quinidina
- Psicotrópicos: haloperidol, citalopram, escitalopram
- Antibióticos: azitromicina, fluoroquinolonas, macrolídeos
- Antifúngicos: fluconazol, ketoconazol
- Antieméticos: ondansetrona, metoclopramida

**Educação da paciente:**
- Explicar importância do QTc e riscos de prolongamento
- Orientar sobre sintomas de alerta (tonturas, desmaios, palpitações)
- Fornecer lista de medicamentos a evitar
- Recomendar sempre informar QTc prolongado a novos médicos

**Seguimento:**
- QTc limítrofe: ECG a cada 3-6 meses
- QTc prolongado: ECG mensal até estabilização, depois trimestral
- Repetir eletrólitos conforme necessidade clínica"""
}

def get_or_create_article(cursor, article_data):
    """Busca ou cria um artigo científico no banco."""

    # Verificar se artigo já existe (por DOI ou título)
    cursor.execute("""
        SELECT id FROM articles
        WHERE doi = %s OR title = %s
        LIMIT 1
    """, (article_data['doi'], article_data['title']))

    result = cursor.fetchone()
    if result:
        print(f"✓ Artigo já existe: {article_data['title'][:60]}...")
        return result[0]

    # Criar novo artigo
    # Converter ano para data (usando 1 de janeiro do ano)
    publish_date = f"{article_data['year']}-01-01"

    cursor.execute("""
        INSERT INTO articles (
            title, authors, journal, publish_date, doi, original_link, article_type,
            created_at, updated_at
        )
        VALUES (%s, %s, %s, %s, %s, %s, 'research_article', NOW(), NOW())
        RETURNING id
    """, (
        article_data['title'],
        article_data['authors'],
        article_data['journal'],
        publish_date,
        article_data['doi'],
        article_data['url']
    ))

    article_id = cursor.fetchone()[0]
    print(f"✓ Artigo criado: {article_data['title'][:60]}...")
    return article_id

def create_article_relation(cursor, article_id, item_id):
    """Cria relação many-to-many entre artigo e score_item."""

    # Verificar se relação já existe
    cursor.execute("""
        SELECT 1 FROM article_score_items
        WHERE article_id = %s AND score_item_id = %s
    """, (article_id, item_id))

    if cursor.fetchone():
        print(f"  → Relação já existe")
        return

    # Criar relação
    cursor.execute("""
        INSERT INTO article_score_items (article_id, score_item_id)
        VALUES (%s, %s)
    """, (article_id, item_id))

    print(f"  → Relação criada com score_item")

def update_score_item(cursor, item_id, content):
    """Atualiza os campos clínicos do score_item."""

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
        content['clinical_relevance'],
        content['patient_explanation'],
        content['conduct'],
        datetime.now(),
        item_id
    ))

    print(f"✓ Score item atualizado com conteúdo clínico")

def main():
    """Executa o enriquecimento completo do item."""

    print("=" * 80)
    print("ENRIQUECIMENTO: ECG - QTc (Bazett) - Mulheres")
    print("=" * 80)
    print()

    try:
        # Conectar ao banco
        conn = psycopg2.connect(**DB_CONFIG)
        cursor = conn.cursor()

        print("✓ Conectado ao banco de dados")
        print()

        # Verificar se item existe
        cursor.execute("SELECT name FROM score_items WHERE id = %s", (ITEM_ID,))
        result = cursor.fetchone()
        if not result:
            raise Exception(f"Item {ITEM_ID} não encontrado!")

        print(f"✓ Item encontrado: {result[0]}")
        print()

        # Processar artigos
        print("ETAPA 1: Processando artigos científicos")
        print("-" * 80)
        article_ids = []
        for article in ARTICLES:
            article_id = get_or_create_article(cursor, article)
            create_article_relation(cursor, article_id, ITEM_ID)
            article_ids.append(article_id)

        print()
        print(f"✓ Total de artigos vinculados: {len(article_ids)}")
        print()

        # Atualizar conteúdo clínico
        print("ETAPA 2: Atualizando conteúdo clínico")
        print("-" * 80)
        update_score_item(cursor, ITEM_ID, CLINICAL_CONTENT)
        print()

        # Commit das alterações
        conn.commit()

        print("=" * 80)
        print("✓ ENRIQUECIMENTO CONCLUÍDO COM SUCESSO!")
        print("=" * 80)
        print()
        print("Resumo:")
        print(f"  - Item ID: {ITEM_ID}")
        print(f"  - Artigos vinculados: {len(article_ids)}")
        print(f"  - Campos atualizados: clinical_relevance, patient_explanation, conduct")
        print(f"  - Data de revisão: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}")
        print()

    except Exception as e:
        if 'conn' in locals():
            conn.rollback()
        print(f"\n❌ ERRO: {str(e)}")
        raise

    finally:
        if 'cursor' in locals():
            cursor.close()
        if 'conn' in locals():
            conn.close()

if __name__ == '__main__':
    main()
