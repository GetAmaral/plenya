#!/usr/bin/env python3
"""
Enriquecimento: Hemácias - Homens
ID: c04e1997-0846-4557-8990-baffb1f7542a
Grupo: Exames > Laboratoriais
"""

import psycopg2
from datetime import datetime

# Dados do item
ITEM_ID = "c04e1997-0846-4557-8990-baffb1f7542a"

# Conteúdo clínico em PT-BR
clinical_relevance = """A contagem de hemácias (eritrócitos) é um dos parâmetros fundamentais do hemograma completo, refletindo a capacidade de transporte de oxigênio do sangue. Em homens adultos saudáveis, valores normais situam-se entre 4,5-5,9 milhões/µL, com variações individuais dependentes de fatores como genética, altitude de residência, nível de condicionamento físico e tabagismo.

Valores elevados (policitemia/eritrocitose) podem indicar condições primárias como policitemia vera - caracterizada pela mutação JAK2 V617F presente em >95% dos casos - ou secundárias como hipóxia crônica, tumores produtores de eritropoietina, ou desidratação relativa. Os critérios diagnósticos da OMS 2022 estabelecem hemoglobina >165 g/L (hematócrito >49%) como threshold para investigação de policitemia em homens ao nível do mar.

Valores reduzidos sugerem anemia, classificada pela análise integrada com VCM (volume corpuscular médio) em microcítica (deficiência de ferro), normocítica (doença crônica, hemólise) ou macrocítica (deficiência de B12/folato). A interpretação requer correlação com hemoglobina, hematócrito e índices eritrocitários, considerando que alterações no volume plasmático (gravidez, desidratação) afetam valores relativos sem refletir mudanças absolutas na massa eritrocitária.

Evidências recentes (2024-2025) destacam a importância das alterações na deformabilidade e tempo de vida das hemácias em condições como COVID-19, doença falciforme e lesões de estocagem, cada vez mais reconhecidas em decisões diagnósticas e terapêuticas."""

patient_explanation = """As hemácias (glóbulos vermelhos) são as células do sangue responsáveis por levar oxigênio para todos os órgãos e tecidos do corpo. Este exame conta quantas hemácias existem em uma pequena amostra do seu sangue.

Em homens adultos, valores normais ficam geralmente entre 4,5 a 5,9 milhões de hemácias por microlitro de sangue. Esse número pode variar um pouco dependendo de onde você mora (pessoas em lugares altos têm mais hemácias), seu condicionamento físico, genética e se você fuma.

Hemácias AUMENTADAS podem indicar:
- Policitemia vera: condição onde a medula óssea produz hemácias demais
- Problemas respiratórios crônicos ou viver em grandes altitudes
- Desidratação (que concentra o sangue)
- Alguns tumores raros

Hemácias DIMINUÍDAS indicam anemia, que pode ser causada por:
- Falta de ferro (anemia ferropriva)
- Deficiência de vitaminas B12 ou ácido fólico
- Doenças crônicas (rins, inflamações)
- Perda de sangue ou destruição aumentada de hemácias

Importante: este valor sempre deve ser analisado junto com outros dados do hemograma (hemoglobina, hematócrito, tamanho das hemácias) para um diagnóstico correto. Seu médico irá avaliar o conjunto de resultados considerando seus sintomas e histórico clínico."""

conduct = """VALORES NORMAIS (4,5-5,9 milhões/µL em homens):
- Manter acompanhamento de rotina conforme indicação clínica
- Reforçar hábitos saudáveis: hidratação adequada, dieta balanceada, atividade física regular
- Avaliar contexto individual (altitude, tabagismo, condicionamento físico)

ERITROCITOSE/POLICITEMIA (>5,9 milhões/µL):
Investigação inicial:
- Repetir hemograma completo com reticulócitos
- Gasometria arterial (avaliar saturação de O2)
- Dosagem de eritropoietina sérica
- Pesquisa de mutação JAK2 V617F (se Hb >165 g/L ou Hct >49%)
- Ultrassom abdominal (avaliar rins, baço, fígado)

Critérios OMS 2022 para Policitemia Vera:
Maiores: (1) Hb >165 g/L ou Hct >49%, (2) Biópsia de medula com hipercelularidade, (3) JAK2 V617F positivo
Menor: Eritropoietina sérica baixa/normal-baixa
Diagnóstico: 3 critérios maiores OU 2 maiores + 1 menor

Conduta conforme etiologia:
- Policitemia vera: encaminhar para hematologia (flebotomias, antiagregação, eventual hidroxiureia)
- Secundária: tratar causa base (DPOC, apneia do sono, tumor)
- Relativa: reidratação, investigar causas de hemoconcentração

ANEMIA (<4,5 milhões/µL):
Classificar por VCM:
- Microcítica (VCM <80 fL): investigar ferro sérico, ferritina, capacidade de ligação do ferro, talassemia
- Normocítica (VCM 80-100 fL): avaliar função renal, doenças crônicas, reticulócitos (hemólise?)
- Macrocítica (VCM >100 fL): dosar vitamina B12, ácido fólico, TSH, considerar mielograma se persistente

Conduta adicional:
- Sintomas (fadiga, dispneia, palidez): urgência na investigação
- Hemoglobina <70 g/L: considerar transfusão e internação
- Anemia crônica leve-moderada: investigação ambulatorial, reposição conforme deficiência identificada

MONITORAMENTO:
- Eritrocitose: hemograma a cada 2-3 meses até estabilização
- Anemia: repetir após 4-6 semanas de tratamento específico
- Avaliar sempre em conjunto: hemoglobina, hematócrito, VCM, HCM, RDW
- Atenção a sintomas novos: cefaleias, prurido aquagênico (policitemia), fadiga extrema, taquicardia (anemia)"""

# Artigos científicos (URLs das buscas realizadas)
articles_data = [
    {
        "title": "Normal and Abnormal Complete Blood Count With Differential",
        "url": "https://www.ncbi.nlm.nih.gov/books/NBK604207/",
        "source": "StatPearls - NCBI Bookshelf",
        "context": "Comprehensive review of CBC parameters including RBC count normal ranges and clinical interpretation"
    },
    {
        "title": "Investigation and management of erythrocytosis",
        "url": "https://pmc.ncbi.nlm.nih.gov/articles/PMC7829024/",
        "source": "CMAJ - Canadian Medical Association Journal",
        "context": "Clinical guidelines for diagnostic approach to elevated RBC counts and polycythemia management"
    },
    {
        "title": "Polycythemia Vera - Clinical Features and Diagnosis",
        "url": "https://www.ncbi.nlm.nih.gov/books/NBK557660/",
        "source": "StatPearls - NCBI Bookshelf",
        "context": "Updated WHO 2022 diagnostic criteria for polycythemia vera including JAK2 mutation testing"
    },
    {
        "title": "Red Cell Indices - Clinical Significance",
        "url": "https://www.ncbi.nlm.nih.gov/books/NBK260/",
        "source": "Clinical Methods - NCBI Bookshelf",
        "context": "Interpretation of RBC indices in anemia classification and diagnostic workup"
    },
    {
        "title": "Erythrocytosis: Diagnosis and investigation",
        "url": "https://onlinelibrary.wiley.com/doi/10.1111/ijlh.14298",
        "source": "International Journal of Laboratory Hematology",
        "context": "2024 review on diagnostic algorithms for primary and secondary erythrocytosis"
    }
]

def insert_articles_and_relationships(cursor):
    """Insere artigos e cria relacionamentos com o score_item"""
    article_ids = []

    for article in articles_data:
        # Inserir artigo
        cursor.execute("""
            INSERT INTO articles (title, url, source, created_at, updated_at)
            VALUES (%s, %s, %s, NOW(), NOW())
            ON CONFLICT (url) DO UPDATE
            SET title = EXCLUDED.title,
                source = EXCLUDED.source,
                updated_at = NOW()
            RETURNING id
        """, (article['title'], article['url'], article['source']))

        article_id = cursor.fetchone()[0]
        article_ids.append(article_id)

        # Criar relacionamento com score_item
        cursor.execute("""
            INSERT INTO article_score_items (score_item_id, article_id)
            VALUES (%s, %s)
            ON CONFLICT (score_item_id, article_id) DO NOTHING
        """, (ITEM_ID, article_id))

    return article_ids

def update_score_item(cursor):
    """Atualiza o score_item com conteúdo clínico"""
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
        clinical_relevance,
        patient_explanation,
        conduct,
        datetime.now().date(),
        ITEM_ID
    ))

    return cursor.rowcount

def main():
    try:
        # Conectar ao banco
        conn = psycopg2.connect(
            host="localhost",
            port=5432,
            database="plenya_db",
            user="plenya_user",
            password="plenya_password"
        )
        cursor = conn.cursor()

        print(f"🔬 Enriquecendo: Hemácias - Homens")
        print(f"📋 ID: {ITEM_ID}\n")

        # Verificar se item existe
        cursor.execute("SELECT name FROM score_items WHERE id = %s", (ITEM_ID,))
        result = cursor.fetchone()

        if not result:
            print(f"❌ ERRO: Item {ITEM_ID} não encontrado no banco!")
            return

        name = result[0]
        print(f"✅ Item encontrado: {name}\n")

        # Inserir artigos e relacionamentos
        print("📚 Inserindo artigos científicos...")
        article_ids = insert_articles_and_relationships(cursor)
        print(f"   ✅ {len(article_ids)} artigos processados\n")

        # Atualizar score_item
        print("💾 Atualizando conteúdo clínico...")
        rows_updated = update_score_item(cursor)

        if rows_updated > 0:
            print(f"   ✅ Item atualizado com sucesso!\n")
        else:
            print(f"   ⚠️  Nenhuma linha atualizada\n")

        # Commit
        conn.commit()

        # Verificar resultado
        cursor.execute("""
            SELECT
                name,
                LENGTH(clinical_relevance) as clin_len,
                LENGTH(patient_explanation) as patient_len,
                LENGTH(conduct) as conduct_len,
                last_review,
                (SELECT COUNT(*) FROM article_score_items WHERE score_item_id = %s) as articles_count
            FROM score_items
            WHERE id = %s
        """, (ITEM_ID, ITEM_ID))

        result = cursor.fetchone()
        print("📊 RESULTADO FINAL:")
        print(f"   Nome: {result[0]}")
        print(f"   Clinical Relevance: {result[1]} caracteres")
        print(f"   Patient Explanation: {result[2]} caracteres")
        print(f"   Conduct: {result[3]} caracteres")
        print(f"   Last Review: {result[4]}")
        print(f"   Artigos vinculados: {result[5]}")
        print("\n✅ ENRIQUECIMENTO COMPLETO!")

        cursor.close()
        conn.close()

    except Exception as e:
        print(f"\n❌ ERRO: {str(e)}")
        if 'conn' in locals():
            conn.rollback()
        raise

if __name__ == "__main__":
    main()
