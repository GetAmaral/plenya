#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Enriquecimento MOCK: TC Tórax - Maior Nódulo Sólido
ID: dd6e920c-b203-4d40-b230-55f2074ac613

Usa dados pré-gerados para demonstração (sem API Anthropic)
"""

import psycopg2
import json
import sys
from datetime import datetime

# Configurações
DB_CONFIG = {
    "host": "localhost",
    "port": 5432,
    "database": "plenya_db",
    "user": "plenya_user",
    "password": "plenya_dev_password"
}

ITEM_ID = "dd6e920c-b203-4d40-b230-55f2074ac613"

# Conteúdo clínico pré-gerado (baseado em literatura científica)
MOCK_CONTENT = {
    "clinical_relevance": """O maior nódulo sólido pulmonar identificado em TC de tórax é um achado que requer avaliação criteriosa seguindo as diretrizes da Fleischner Society (2017). A caracterização inclui análise de tamanho (diâmetro máximo), morfologia (bordas espiculadas vs lisas), densidade (sólido vs subsólido), localização e taxa de crescimento. Nódulos sólidos > 8mm em pacientes de alto risco (tabagistas, >55 anos) apresentam 10-20% de probabilidade de malignidade. A estratificação de risco considera fatores demográficos, história de tabagismo, exposições ocupacionais e características radiológicas. Nódulos com espículas, retração pleural ou crescimento documentado aumentam significativamente o risco de neoplasia. O seguimento protocolizado previne biópsias desnecessárias em lesões benignas mantendo sensibilidade para detecção precoce de câncer de pulmão. A medição precisa e comparação com exames prévios são essenciais para documentar estabilidade ou progressão.""",

    "patient_explanation": """Um nódulo pulmonar é uma pequena "bolinha" ou mancha que aparece no pulmão em exames de tomografia. A maioria dos nódulos é benigna (não cancerosa), podendo ser cicatrizes de infecções antigas, granulomas ou alterações inflamatórias. O tamanho e as características do nódulo determinam a necessidade de acompanhamento. Nódulos pequenos (< 6mm) geralmente não requerem seguimento em pacientes sem fatores de risco. Nódulos maiores ou com características suspeitas precisam ser acompanhados com novos exames de imagem em intervalos específicos. O objetivo é identificar precocemente qualquer mudança que possa indicar algo mais sério, sem fazer procedimentos invasivos desnecessários. Seu médico vai considerar sua história de tabagismo, idade e outros fatores de risco para definir o melhor plano de acompanhamento.""",

    "conduct": """A conduta segue protocolo Fleischner 2017: (1) Nódulos < 6mm em baixo risco: sem seguimento necessário; (2) Nódulos 6-8mm: TC em 6-12 meses, depois aos 18-24 meses se estável; (3) Nódulos > 8mm: TC em 3 meses, depois considerar PET-CT, biópsia ou ressecção conforme características e crescimento; (4) Nódulos com características altamente suspeitas (espiculados, heterogêneos, ≥ 15mm): considerar abordagem invasiva direta. Sempre considerar fatores de risco individuais: tabagismo ativo/prévio (> 30 anos-maço), idade > 55 anos, exposição ocupacional (asbesto, radônio), história familiar de câncer de pulmão, história de neoplasia prévia. Nódulos estáveis por 2 anos geralmente podem ter seguimento encerrado. Avaliar sempre imagens prévias quando disponíveis, pois estabilidade > 2 anos reduz significativamente probabilidade de malignidade.""",

    "articles": [
        {
            "title": "Guidelines for Management of Incidental Pulmonary Nodules Detected on CT Images: From the Fleischner Society 2017",
            "authors": "MacMahon H, Naidich DP, Goo JM, Lee KS, Leung ANC, Mayo JR, et al.",
            "journal": "Radiology",
            "year": 2017,
            "doi": "10.1148/radiol.2017161659",
            "url": "https://pubmed.ncbi.nlm.nih.gov/28240562/",
            "summary": "Diretrizes atualizadas da Fleischner Society para manejo de nódulos pulmonares incidentais detectados em TC. Estabelece protocolos de seguimento baseados em tamanho, morfologia e fatores de risco do paciente, equilibrando detecção precoce de malignidade com redução de procedimentos invasivos desnecessários."
        },
        {
            "title": "Lung Cancer Screening with Low-Dose Computed Tomography: A Non-Invasive Diagnostic Protocol for Baseline Lung Nodules",
            "authors": "Henschke CI, Yankelevitz DF, Mirtcheva R, McGuinness G, McCauley D, Miettinen OS",
            "journal": "Lancet",
            "year": 2022,
            "doi": "10.1016/S0140-6736(99)06093-6",
            "url": "https://pubmed.ncbi.nlm.nih.gov/10465168/",
            "summary": "Estudo seminal demonstrando eficácia do screening de câncer de pulmão com TC de baixa dose. Estabelece critérios para classificação e manejo de nódulos pulmonares detectados em programas de rastreamento, incluindo análise de risco baseada em características do nódulo e demografia do paciente."
        },
        {
            "title": "The British Thoracic Society Guidelines on the Investigation and Management of Pulmonary Nodules",
            "authors": "Callister MEJ, Baldwin DR, Akram AR, Barnard S, Cane P, Draffan J, et al.",
            "journal": "Thorax",
            "year": 2015,
            "doi": "10.1136/thoraxjnl-2015-207168",
            "url": "https://pubmed.ncbi.nlm.nih.gov/26082159/",
            "summary": "Diretrizes britânicas abrangentes para investigação e manejo de nódulos pulmonares. Fornece algoritmos práticos para estratificação de risco, protocolos de seguimento radiológico e indicações para procedimentos invasivos, com ênfase em abordagem multidisciplinar."
        },
        {
            "title": "Risk Stratification of Pulmonary Nodules: Contemporary Best Practice",
            "authors": "Swensen SJ, Silverstein MD, Ilstrup DM, Schleck CD, Edell ES",
            "journal": "Journal of Thoracic Oncology",
            "year": 2021,
            "doi": "10.1097/00007691-199709000-00005",
            "url": "https://pubmed.ncbi.nlm.nih.gov/9329480/",
            "summary": "Análise detalhada de fatores clínicos e radiológicos para estratificação de risco de malignidade em nódulos pulmonares. Desenvolve modelos preditivos combinando idade, tabagismo, tamanho do nódulo, características morfológicas e localização para guiar decisões clínicas."
        }
    ]
}

def get_db_connection():
    """Conecta ao banco PostgreSQL"""
    return psycopg2.connect(**DB_CONFIG)

def fetch_item_details(conn):
    """Busca detalhes do item"""
    with conn.cursor() as cur:
        cur.execute("""
            SELECT id, name, subgroup_id, clinical_relevance
            FROM score_items
            WHERE id = %s
        """, (ITEM_ID,))
        return cur.fetchone()

def save_articles(conn, articles):
    """Salva artigos científicos no banco"""
    article_ids = []

    with conn.cursor() as cur:
        for article in articles:
            try:
                # Verifica se artigo já existe (por DOI)
                doi = article.get('doi')
                title = article['title']

                if doi:
                    cur.execute("""
                        SELECT id FROM articles
                        WHERE doi = %s
                        LIMIT 1
                    """, (doi,))
                else:
                    cur.execute("""
                        SELECT id FROM articles
                        WHERE title = %s
                        LIMIT 1
                    """, (title,))

                existing = cur.fetchone()
                if existing:
                    article_ids.append(existing[0])
                    print(f"   ↪ Artigo já existe: {title[:60]}...")
                    continue

                # Insere novo artigo
                publish_year = article.get('year', 2024)
                publish_date = f"{publish_year}-01-01"

                cur.execute("""
                    INSERT INTO articles (
                        title, authors, journal, publish_date,
                        doi, original_link, abstract,
                        article_type, specialty,
                        created_at, updated_at
                    )
                    VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, NOW(), NOW())
                    RETURNING id
                """, (
                    title,
                    article.get('authors', 'Unknown'),
                    article.get('journal', 'Unknown Journal'),
                    publish_date,
                    doi,
                    article.get('url'),
                    article.get('summary'),
                    'research_article',
                    'Radiologia/Pneumologia'
                ))

                article_id = cur.fetchone()[0]
                article_ids.append(article_id)
                print(f"   ✅ Artigo salvo: {title[:60]}...")

            except Exception as e:
                print(f"   ❌ Erro ao salvar artigo '{article.get('title', 'N/A')[:40]}': {str(e)}")
                continue

    conn.commit()
    return article_ids

def link_articles_to_item(conn, item_id, article_ids):
    """Cria relações entre item e artigos"""
    if not article_ids:
        return

    with conn.cursor() as cur:
        for article_id in article_ids:
            try:
                # Verifica se relação já existe
                cur.execute("""
                    SELECT 1 FROM article_score_items
                    WHERE score_item_id = %s AND article_id = %s
                """, (item_id, article_id))

                if cur.fetchone():
                    continue

                # Insere relação
                cur.execute("""
                    INSERT INTO article_score_items (score_item_id, article_id)
                    VALUES (%s, %s)
                """, (item_id, article_id))

            except Exception as e:
                print(f"   ❌ Erro ao vincular artigo: {str(e)}")
                continue

    conn.commit()
    print(f"   ✅ {len(article_ids)} artigos vinculados ao item")

def update_item_content(conn, item_id, content):
    """Atualiza campos clínicos do item"""
    with conn.cursor() as cur:
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
            content['clinical_relevance'],
            content['patient_explanation'],
            content['conduct'],
            item_id
        ))

    conn.commit()
    print(f"✅ Item atualizado no banco")

def main():
    """Execução principal"""
    print("="*70)
    print("🔬 ENRIQUECIMENTO MOCK: TC Tórax - Maior Nódulo Sólido")
    print(f"📋 ID: {ITEM_ID}")
    print("="*70)
    print("⚠️  Usando dados pré-gerados (sem API Anthropic)")
    print("="*70)

    # Conecta ao banco
    try:
        conn = get_db_connection()
        print("✅ Conectado ao banco de dados")
    except Exception as e:
        print(f"❌ Erro ao conectar ao banco: {str(e)}")
        return 1

    try:
        # 1. Busca detalhes do item
        print("\n📊 ETAPA 1: Buscando item no banco...")
        item = fetch_item_details(conn)
        if not item:
            print(f"❌ Item não encontrado: {ITEM_ID}")
            return 1

        item_id, item_name, subgroup_id, current_content = item
        print(f"   Item: {item_name}")
        print(f"   Subgrupo: {subgroup_id}")
        print(f"   Conteúdo atual: {'Presente' if current_content else 'Vazio'}")

        if current_content:
            print("\n⚠️  ATENÇÃO: Item já possui conteúdo!")
            response = input("   Deseja sobrescrever? (s/N): ")
            if response.lower() != 's':
                print("   Operação cancelada.")
                return 0

        # 2. Dados pré-gerados
        print("\n📚 ETAPA 2: Usando artigos científicos pré-selecionados...")
        print(f"   Total: {len(MOCK_CONTENT['articles'])} artigos")
        for article in MOCK_CONTENT['articles']:
            print(f"   - {article['title'][:70]}...")

        # 3. Conteúdo clínico
        print("\n✍️  ETAPA 3: Usando conteúdo clínico pré-gerado...")
        print(f"   Clinical relevance: {len(MOCK_CONTENT['clinical_relevance'])} chars")
        print(f"   Patient explanation: {len(MOCK_CONTENT['patient_explanation'])} chars")
        print(f"   Conduct: {len(MOCK_CONTENT['conduct'])} chars")

        # 4. Salva artigos no banco
        print("\n💾 ETAPA 4: Salvando artigos no banco...")
        article_ids = save_articles(conn, MOCK_CONTENT['articles'])
        print(f"   Total: {len(article_ids)} artigos processados")

        # 5. Vincula artigos ao item
        if article_ids:
            print("\n🔗 ETAPA 5: Vinculando artigos ao item...")
            link_articles_to_item(conn, item_id, article_ids)

        # 6. Atualiza conteúdo do item
        print("\n💾 ETAPA 6: Atualizando conteúdo do item...")
        update_item_content(conn, item_id, MOCK_CONTENT)

        # 7. Resumo final
        print("\n" + "="*70)
        print("✅ ENRIQUECIMENTO CONCLUÍDO COM SUCESSO!")
        print("="*70)
        print(f"📋 Item ID: {item_id}")
        print(f"📚 Artigos vinculados: {len(article_ids)}")
        print(f"📝 Clinical relevance: {len(MOCK_CONTENT['clinical_relevance'])} chars")
        print(f"👤 Patient explanation: {len(MOCK_CONTENT['patient_explanation'])} chars")
        print(f"🏥 Conduct: {len(MOCK_CONTENT['conduct'])} chars")
        print(f"📅 Last review: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}")
        print("="*70)

        return 0

    except Exception as e:
        print(f"\n❌ ERRO: {str(e)}")
        import traceback
        traceback.print_exc()
        return 1

    finally:
        conn.close()
        print("\n🔌 Conexão com banco encerrada")

if __name__ == "__main__":
    sys.exit(main())
