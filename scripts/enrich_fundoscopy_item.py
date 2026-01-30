#!/usr/bin/env python3
"""
Script para enriquecer o item Fundoscopia - Retinopatia Diabética
ID: dbbb0258-ebcb-4bb2-ace3-73d0d7b93a52
Grupo: Exames > Imagem
"""

import psycopg2
from psycopg2.extras import Json
import os
from datetime import datetime

# Dados do item
ITEM_ID = "dbbb0258-ebcb-4bb2-ace3-73d0d7b93a52"

# Artigos científicos para vincular
ARTICLES = [
    {
        "title": "Retinopathy, Neuropathy, and Foot Care: Standards of Care in Diabetes—2026",
        "authors": "American Diabetes Association",
        "journal": "Diabetes Care",
        "publish_date": "2026-01-01",
        "doi": "10.2337/dc26-S012",
        "url": "https://pmc.ncbi.nlm.nih.gov/articles/PMC12690177/",
        "pm_id": "39651974"
    },
    {
        "title": "Classification of diabetic retinopathy: Past, present and future",
        "authors": "Wilkinson CP, Ferris FL 3rd, Klein RE, et al.",
        "journal": "Eye",
        "publish_date": "2023-01-01",
        "doi": "10.1038/s41433-022-02387-7",
        "url": "https://pmc.ncbi.nlm.nih.gov/articles/PMC9800497/",
        "pm_id": "36609536"
    },
    {
        "title": "Update on the Management of Diabetic Retinopathy: Anti-VEGF Agents for the Prevention of Complications and Progression",
        "authors": "Cheung N, Mitchell P, Wong TY",
        "journal": "Journal of Clinical Medicine",
        "publish_date": "2023-05-01",
        "doi": "10.3390/jcm12103370",
        "url": "https://pmc.ncbi.nlm.nih.gov/articles/PMC10222751/",
        "pm_id": "37240555"
    }
]

# Conteúdo clínico enriquecido em PT-BR
CLINICAL_CONTENT = {
    "clinical_relevance": """A fundoscopia para avaliação de retinopatia diabética é essencial no acompanhamento de pacientes com diabetes mellitus tipo 1 e tipo 2. A retinopatia diabética é uma das principais causas de cegueira em adultos em idade produtiva.

**Classificação ETDRS (Early Treatment Diabetic Retinopathy Study):**

**Retinopatia Diabética Não Proliferativa (RDNP):**
- RDNP Leve: Presença de microaneurismas isolados
- RDNP Moderada: Hemorragias ou microaneurismas em 1-3 quadrantes retinianos, exsudatos duros, exsudatos algodonosos ou tortuosidade venosa
- RDNP Grave: Regra 4:2:1 - Hemorragias em 4 quadrantes ou tortuosidade venosa em 2+ quadrantes ou AMIR (Anomalias Microvasculares Intrarretinianas) em 1+ quadrante

**Retinopatia Diabética Proliferativa (RDP):**
- Neovascularização do disco óptico (NVD) ou retina (NVE)
- Hemorragia vítrea ou pré-retiniana
- Descolamento de retina tracional

**Edema Macular Diabético (EMD):**
- Pode ocorrer em qualquer estágio
- Causa principal de perda visual em diabéticos

As diretrizes ADA 2026 recomendam rastreamento anual com fundoscopia dilatada ou fotografia de retina validada. Exames a cada 1-2 anos podem ser custo-efetivos após um ou mais exames normais.""",

    "patient_explanation": """A fundoscopia para retinopatia diabética é um exame oftalmológico que avalia as alterações nos vasos sanguíneos da retina causadas pelo diabetes.

**Por que é importante:**
O diabetes pode danificar os pequenos vasos da retina ao longo do tempo, levando a problemas de visão e, se não tratado, à cegueira. O exame detecta essas alterações precocemente, quando o tratamento é mais efetivo.

**Tipos de alterações:**
- **Fase inicial (não proliferativa)**: Pequenos vasos se dilatam formando microaneurismas, podem ocorrer pequenas hemorragias e acúmulo de gordura (exsudatos)
- **Fase avançada (proliferativa)**: Formação de vasos sanguíneos anormais que podem sangrar dentro do olho ou causar descolamento de retina

**Quando fazer:**
- Diabetes tipo 1: Começar 5 anos após o diagnóstico
- Diabetes tipo 2: Iniciar no momento do diagnóstico
- Periodicidade: Anualmente ou a cada 1-2 anos se exames normais
- Mais frequente se houver alterações detectadas

**Preparo:**
O médico irá dilatar suas pupilas com colírios para examinar toda a retina. Sua visão ficará embaçada por 4-6 horas após o exame.

**Tratamento disponível:**
Caso sejam encontradas alterações, existem tratamentos eficazes como laser (fotocoagulação) e injeções intraoculares (anti-VEGF) que podem prevenir a progressão e perda visual.""",

    "conduct": """**Indicações de Rastreamento (ADA 2026):**

**Diabetes Tipo 1:**
- Início: 5 anos após o diagnóstico
- Periodicidade: Anualmente

**Diabetes Tipo 2:**
- Início: No momento do diagnóstico
- Periodicidade: Anualmente
- Intervalos de 1-2 anos podem ser considerados após 1+ exames normais

**Gestantes com Diabetes:**
- Avaliar no primeiro trimestre
- Monitoramento trimestral e até 1 ano pós-parto

**Método de Rastreamento:**
1. Fundoscopia dilatada por oftalmologista
2. Fotografia de retina com 7 campos estereoscópicos (padrão-ouro)
3. Fotografia digital de retina não-midriática validada
4. Sistemas de IA aprovados pelo FDA (EyeArt, AEYE-DS, LumineticsCore)

**Classificação dos Achados:**

**Sem Retinopatia:** Repetir anualmente

**RDNP Leve a Moderada:**
- Controle glicêmico rigoroso (HbA1c <7%)
- Controle pressórico (<130/80 mmHg)
- Acompanhamento oftalmológico semestral a anual

**RDNP Grave:**
- Referência urgente ao retinólogo
- Considerar fotocoagulação panretiniana (PRP) profilática
- Considerar anti-VEGF intravítreo
- Seguimento a cada 3-4 meses

**RDP:**
- Encaminhamento imediato ao retinólogo
- Tratamento: PRP ou anti-VEGF intravítreo
- Anti-VEGF (ranibizumabe, aflibercepte, bevacizumabe) demonstram melhores resultados visuais em curto prazo
- PRP é tratamento estabelecido, durável e custo-efetivo
- Terapia combinada pode ser indicada

**Edema Macular Diabético:**
- Primeira linha: Anti-VEGF intravítreo
- Alternativa: Laser focal/grid ou implante de corticoide
- OCT para acompanhamento

**Urgências:**
- Hemorragia vítrea: Encaminhamento imediato
- Descolamento de retina: Emergência cirúrgica
- Perda visual súbita: Avaliar em <24h

**Prevenção:**
- Controle glicêmico intensivo reduz risco em 76% (DCCT/UKPDS)
- Controle pressórico adequado
- Controle lipídico (estatinas)
- Evitar tabagismo
- Atividade física regular"""
}

def connect_db():
    """Conecta ao banco de dados PostgreSQL"""
    return psycopg2.connect(
        host=os.getenv("DB_HOST", "localhost"),
        port=os.getenv("DB_PORT", "5432"),
        database=os.getenv("DB_NAME", "plenya_db"),
        user=os.getenv("DB_USER", "plenya_user"),
        password=os.getenv("DB_PASSWORD", "senha_segura")
    )

def upsert_article(cursor, article):
    """Insere ou atualiza um artigo científico"""
    cursor.execute("""
        INSERT INTO articles (
            title, authors, journal, publish_date, doi, original_link, pm_id,
            created_at, updated_at
        ) VALUES (
            %s, %s, %s, %s, %s, %s, %s, NOW(), NOW()
        )
        ON CONFLICT (doi)
        DO UPDATE SET
            title = EXCLUDED.title,
            authors = EXCLUDED.authors,
            journal = EXCLUDED.journal,
            publish_date = EXCLUDED.publish_date,
            original_link = EXCLUDED.original_link,
            pm_id = EXCLUDED.pm_id,
            updated_at = NOW()
        RETURNING id
    """, (
        article["title"],
        article["authors"],
        article["journal"],
        article["publish_date"],
        article["doi"],
        article["url"],
        article["pm_id"]
    ))

    result = cursor.fetchone()
    return result[0]

def link_article_to_score_item(cursor, score_item_id, article_id):
    """Cria a relação entre score_item e article"""
    cursor.execute("""
        INSERT INTO article_score_items (article_id, score_item_id)
        VALUES (%s, %s)
        ON CONFLICT (score_item_id, article_id) DO NOTHING
    """, (article_id, score_item_id))

def update_score_item_content(cursor, item_id, content):
    """Atualiza o conteúdo clínico do score_item"""
    cursor.execute("""
        UPDATE score_items
        SET
            clinical_relevance = %s,
            patient_explanation = %s,
            conduct = %s,
            updated_at = NOW()
        WHERE id = %s
    """, (
        content["clinical_relevance"],
        content["patient_explanation"],
        content["conduct"],
        item_id
    ))

def main():
    print(f"🚀 Iniciando enriquecimento do item: Fundoscopia - Retinopatia Diabética")
    print(f"ID: {ITEM_ID}\n")

    conn = connect_db()
    cursor = conn.cursor()

    try:
        # 1. Verificar se o item existe
        cursor.execute("SELECT id, name FROM score_items WHERE id = %s", (ITEM_ID,))
        item = cursor.fetchone()

        if not item:
            print(f"❌ Item não encontrado: {ITEM_ID}")
            return

        print(f"✅ Item encontrado: {item[1]}\n")

        # 2. Inserir/atualizar artigos científicos
        print("📚 Processando artigos científicos...")
        article_ids = []

        for i, article in enumerate(ARTICLES, 1):
            article_id = upsert_article(cursor, article)
            article_ids.append(article_id)
            print(f"  {i}. ✅ {article['title'][:60]}... (ID: {article_id})")

        print()

        # 3. Vincular artigos ao score_item
        print("🔗 Vinculando artigos ao item...")
        for article_id in article_ids:
            link_article_to_score_item(cursor, ITEM_ID, article_id)
        print(f"  ✅ {len(article_ids)} artigos vinculados\n")

        # 4. Atualizar conteúdo clínico
        print("📝 Atualizando conteúdo clínico em PT-BR...")
        update_score_item_content(cursor, ITEM_ID, CLINICAL_CONTENT)
        print("  ✅ clinical_relevance")
        print("  ✅ patient_explanation")
        print("  ✅ conduct\n")

        # 5. Verificar resultado
        cursor.execute("""
            SELECT
                si.id,
                si.name,
                LENGTH(si.clinical_relevance) as clinical_len,
                LENGTH(si.patient_explanation) as patient_len,
                LENGTH(si.conduct) as conduct_len,
                COUNT(asi.article_id) as article_count
            FROM score_items si
            LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
            WHERE si.id = %s
            GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct
        """, (ITEM_ID,))

        result = cursor.fetchone()

        print("=" * 70)
        print("📊 RESULTADO FINAL")
        print("=" * 70)
        print(f"Item: {result[1]}")
        print(f"ID: {result[0]}")
        print(f"Clinical Relevance: {result[2]} caracteres")
        print(f"Patient Explanation: {result[3]} caracteres")
        print(f"Conduct: {result[4]} caracteres")
        print(f"Artigos vinculados: {result[5]}")
        print("=" * 70)

        # Commit das alterações
        conn.commit()
        print("\n✅ Enriquecimento concluído com sucesso!")

    except Exception as e:
        conn.rollback()
        print(f"\n❌ Erro durante o enriquecimento: {e}")
        raise

    finally:
        cursor.close()
        conn.close()

if __name__ == "__main__":
    main()
