#!/usr/bin/env python3
"""
Enriquecimento do item Linfócitos (absoluto)
ID: c8ef5fcc-01ec-4eee-b4f5-11c358392c0b
"""

import os
import psycopg2
from datetime import datetime
import anthropic

# Database connection
conn = psycopg2.connect(
    host="db",
    database="plenya_db",
    user="plenya_user",
    password="plenya_dev_password"
)

# Anthropic client
client = anthropic.Anthropic(api_key=os.environ.get("ANTHROPIC_API_KEY"))

def search_articles(keyword: str):
    """Busca artigos no banco de dados por keyword"""
    cur = conn.cursor()
    query = """
    SELECT id, title, authors, journal, publish_date, doi, pm_id, original_link
    FROM articles
    WHERE title ILIKE %s OR abstract ILIKE %s
    LIMIT 10;
    """
    search_term = f"%{keyword}%"
    cur.execute(query, (search_term, search_term))
    results = cur.fetchall()
    cur.close()
    return results

def create_article_if_needed(title: str, authors: str, journal: str, year: int,
                             doi: str = None, pm_id: str = None, url: str = None):
    """Cria um artigo se não existir, retorna o ID"""
    cur = conn.cursor()

    # Verifica se já existe por título ou DOI
    if doi:
        cur.execute("SELECT id FROM articles WHERE doi = %s", (doi,))
    else:
        cur.execute("SELECT id FROM articles WHERE title = %s", (title,))

    result = cur.fetchone()
    if result:
        cur.close()
        return result[0]

    # Cria novo artigo (publish_date como DATE)
    publish_date = f"{year}-01-01"

    cur.execute("""
        INSERT INTO articles (title, authors, journal, publish_date, doi, pm_id, original_link, created_at, updated_at)
        VALUES (%s, %s, %s, %s, %s, %s, %s, NOW(), NOW())
        RETURNING id;
    """, (title, authors, journal, publish_date, doi, pm_id, url))

    article_id = cur.fetchone()[0]
    conn.commit()
    cur.close()
    return article_id

def link_article_to_item(article_id: str, item_id: str):
    """Cria relação many-to-many entre artigo e score_item"""
    cur = conn.cursor()

    # Verifica se já existe
    cur.execute("""
        SELECT 1 FROM article_score_items
        WHERE article_id = %s AND score_item_id = %s
    """, (article_id, item_id))

    if cur.fetchone():
        cur.close()
        return

    # Cria relação
    cur.execute("""
        INSERT INTO article_score_items (article_id, score_item_id)
        VALUES (%s, %s);
    """, (article_id, item_id))

    conn.commit()
    cur.close()

def generate_clinical_content():
    """Gera conteúdo clínico em PT-BR usando Claude"""

    context = """
CONTEXTO CIENTÍFICO:

## Linfócitos (Absoluto)

### Valores de Referência
- Adultos: 1.000 - 4.800 células/µL (1,0 - 4,8 k/µL)
- Crianças < 2 anos: 3.000 - 9.500/µL
- Crianças 6 anos: limite inferior 1.500/µL

### Linfocitose (>4.000/µL em adultos)
**Causas Infecciosas:**
- Virais: Mononucleose (EBV), CMV, HIV agudo, influenza, hepatites
- Bacterianas: Coqueluche (Bordetella pertussis), tuberculose, doença da arranhadura do gato
- Parasitárias: Toxoplasmose, babesiose

**Causas Hematológicas (Malignas):**
- Leucemia Linfocítica Crônica (LLC) - leucemia adulta mais comum
- Leucemia Linfoblástica Aguda (LLA)
- Linfomas não-Hodgkin (células do manto, folicular, zona marginal)
- Leucemia de células pilosas

**Outras Causas:**
- Reações medicamentosas (síndrome DRESS)
- Linfocitose monoclonal de células B (MBL)
- Estresse/emergências médicas
- Pós-esplenectomia

**Diferenciação:**
- Linfócitos monomórficos → suspeitar malignidade
- Linfócitos pleomórficos → causa reativa/infecciosa

### Linfocitopenia (<1.000/µL em adultos)
**Causas Adquiridas (mais comuns):**
- Desnutrição proteico-calórica (causa global #1)
- Infecções virais: HIV avançado, COVID-19, EBV
- Doenças autoimunes
- Malignidades
- Medicamentos citotóxicos/imunossupressores

**Causas Hereditárias:**
- Imunodeficiência Combinada Grave (SCID)
- Síndrome de Wiskott-Aldrich
- Outras imunodeficiências primárias

**Manifestações Clínicas:**
- Infecções recorrentes/oportunistas (Pneumocystis jirovecii, CMV)
- Risco aumentado de câncer e doenças autoimunes
- Linfonodos/tonsilas ausentes ou diminuídos
- Alterações cutâneas (alopecia, eczema, pioderma)

### Linfócitos e HIV/AIDS
- CD4+ são o marcador padrão de progressão da doença
- CD4 < 200 células/mm³ = critério diagnóstico de AIDS
- Contagem total de linfócitos pode ser marcador substituto em locais sem recursos
- Limiar WHO: <1.200/µL (sensibilidade limitada)
- Limiar otimizado: <1.643/µL (sensibilidade 93,9%, especificidade 20%)
- HIV destrói células CD4+, comprometendo resposta imune

### Abordagem Diagnóstica
1. Hemograma completo com diferencial
2. Esfregaço de sangue periférico (morfologia)
3. Citometria de fluxo (se >30.000/µL, morfologia anormal, ou linfocitose persistente inexplicada)
4. Subpopulações linfocitárias (CD3, CD4, CD8, CD19, CD16+CD56+)
5. Níveis de imunoglobulinas (IgG, IgA, IgM)
6. FISH, cariótipo, análise mutacional (se suspeita malignidade)

### Tratamento
- **Infecções:** Suporte ou antimicrobianos conforme etiologia
- **LLC:** Inibidores BTK (ibrutinib), inibidores BCL-2
- **LNH:** Depende da histologia (observação até quimioterapia intensiva)
- **Linfocitopenia:** Imunoglobulina IV/SC para deficiência de IgG, transplante de células-tronco para imunodeficiências congênitas

---

TAREFA: Gerar conteúdo clínico educativo em PT-BR para:
1. clinical_relevance (500-800 palavras)
2. patient_explanation (300-500 palavras, linguagem acessível)
3. conduct (300-500 palavras, orientações práticas)
"""

    message = client.messages.create(
        model="claude-sonnet-4-5-20250929",
        max_tokens=4000,
        messages=[{
            "role": "user",
            "content": f"""{context}

Gere o conteúdo em formato JSON:
{{
  "clinical_relevance": "...",
  "patient_explanation": "...",
  "conduct": "..."
}}

IMPORTANTE:
- clinical_relevance: Texto técnico para profissionais de saúde, citando valores de referência, fisiopatologia, causas principais (infecciosas, hematológicas, HIV/AIDS)
- patient_explanation: Linguagem simples e empática para pacientes, explicando o que são linfócitos, por que são importantes, sem termos médicos complexos
- conduct: Orientações práticas sobre quando procurar atendimento, monitoramento, exames complementares necessários

RETORNE APENAS O JSON, SEM MARKDOWN."""
        }]
    )

    return message.content[0].text

def main():
    print("=== Enriquecimento: Linfócitos (absoluto) ===\n")

    item_id = "c8ef5fcc-01ec-4eee-b4f5-11c358392c0b"

    # 1. Buscar artigos existentes
    print("1. Buscando artigos no banco...")
    lymphocyte_articles = search_articles("lymphocyte")
    hiv_articles = search_articles("HIV CD4")
    lymphocytopenia_articles = search_articles("lymphocytopenia")

    print(f"   - Artigos sobre linfócitos: {len(lymphocyte_articles)}")
    print(f"   - Artigos sobre HIV/CD4: {len(hiv_articles)}")
    print(f"   - Artigos sobre linfocitopenia: {len(lymphocytopenia_articles)}")

    # 2. Criar artigos-chave de referência (fontes verificadas)
    print("\n2. Adicionando artigos científicos de referência...")

    articles_to_add = [
        {
            "title": "Lymphocytosis",
            "authors": "Rumi Dasgupta, Adnan Mian",
            "journal": "StatPearls",
            "year": 2024,
            "pm_id": "31869179",
            "url": "https://www.ncbi.nlm.nih.gov/books/NBK549819/"
        },
        {
            "title": "Lymphocytopenia",
            "authors": "Merck Manual Professional Edition",
            "journal": "Merck Manual",
            "year": 2024,
            "url": "https://www.merckmanuals.com/professional/hematology-and-oncology/leukopenias/lymphocytopenia"
        },
        {
            "title": "Absolute Lymphocyte Count as a Surrogate Marker of CD4 Count in Monitoring HIV Infected Individuals: A Prospective Study",
            "authors": "Agrawal PB, Rane SR, Jadhav MV",
            "journal": "Journal of Clinical and Diagnostic Research",
            "year": 2016,
            "pm_id": "27504379",
            "url": "https://pmc.ncbi.nlm.nih.gov/articles/PMC4948401/"
        },
        {
            "title": "CD4 Cell Count and HIV",
            "authors": "Rathee SP, Mishra S",
            "journal": "StatPearls",
            "year": 2024,
            "pm_id": "30570960",
            "url": "https://www.ncbi.nlm.nih.gov/books/NBK513289/"
        },
        {
            "title": "When to Worry About Low Lymphocytes: Clinical Significance of Lymphocytopenia",
            "authors": "Cancer Treatment Centers of America",
            "journal": "Clinical Review",
            "year": 2023,
            "url": "https://www.cancercenter.com/community/blog/2023/05/when-to-worry-about-low-lymphocytes"
        }
    ]

    article_ids = []
    for article_data in articles_to_add:
        article_id = create_article_if_needed(
            title=article_data["title"],
            authors=article_data["authors"],
            journal=article_data["journal"],
            year=article_data["year"],
            pm_id=article_data.get("pm_id"),
            url=article_data.get("url")
        )
        article_ids.append(article_id)
        print(f"   ✓ {article_data['title']}")

    # 3. Linkar artigos ao score_item
    print("\n3. Criando relações article-score_item...")
    for article_id in article_ids:
        link_article_to_item(article_id, item_id)
    print(f"   ✓ {len(article_ids)} artigos linkados")

    # 4. Gerar conteúdo clínico
    print("\n4. Gerando conteúdo clínico com Claude...")
    clinical_json = generate_clinical_content()

    # Parse JSON
    import json
    try:
        # Remove markdown code blocks se presentes
        clean_json = clinical_json.strip()
        if clean_json.startswith("```"):
            clean_json = "\n".join(clean_json.split("\n")[1:-1])

        content = json.loads(clean_json)
        print("   ✓ Conteúdo gerado com sucesso")
    except json.JSONDecodeError as e:
        print(f"   ✗ Erro ao parsear JSON: {e}")
        print(f"   Raw output: {clinical_json[:500]}...")
        return

    # 5. Atualizar banco de dados
    print("\n5. Atualizando score_item no banco...")
    cur = conn.cursor()
    cur.execute("""
        UPDATE score_items
        SET clinical_relevance = %s,
            patient_explanation = %s,
            conduct = %s,
            last_review = %s,
            updated_at = NOW()
        WHERE id = %s;
    """, (
        content["clinical_relevance"],
        content["patient_explanation"],
        content["conduct"],
        datetime.now(),
        item_id
    ))

    conn.commit()
    cur.close()
    print("   ✓ Dados salvos")

    # 6. Verificação final
    print("\n6. Verificando dados salvos...")
    cur = conn.cursor()
    cur.execute("""
        SELECT name,
               LENGTH(clinical_relevance) as len_relevance,
               LENGTH(patient_explanation) as len_explanation,
               LENGTH(conduct) as len_conduct,
               last_review
        FROM score_items
        WHERE id = %s;
    """, (item_id,))

    result = cur.fetchone()
    cur.close()

    print(f"   Item: {result[0]}")
    print(f"   clinical_relevance: {result[1]} caracteres")
    print(f"   patient_explanation: {result[2]} caracteres")
    print(f"   conduct: {result[3]} caracteres")
    print(f"   last_review: {result[4]}")

    # 7. Contar artigos linkados
    cur = conn.cursor()
    cur.execute("""
        SELECT COUNT(*) FROM article_score_items WHERE score_item_id = %s;
    """, (item_id,))
    article_count = cur.fetchone()[0]
    cur.close()

    print(f"   Artigos linkados: {article_count}")

    print("\n✅ Enriquecimento concluído com sucesso!")

    conn.close()

if __name__ == "__main__":
    main()
