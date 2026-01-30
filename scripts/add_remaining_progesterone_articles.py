#!/usr/bin/env python3
"""
Add the 2 remaining progesterone articles
"""
import requests
import time

API_BASE_URL = "http://localhost:3001/api/v1"
EMAIL, PASSWORD = "import@plenya.com", "Import@123456"
SCORE_ITEM_ID = "325b61c6-11f6-49af-b3c9-b211d8bd6c28"

ARTICLES = [
    {
        "title": "The impact of micronized progesterone on breast cancer risk: a systematic review",
        "authors": "Stute P, Wildt L, Neulen J",
        "journal": "Climacteric",
        "publicationYear": 2018,
        "volume": "21",
        "issue": "2",
        "pages": "111-122",
        "doi": "10.1080/13697137.2017.1421925",
        "pmid": "29384406",
        "url": "https://www.tandfonline.com/doi/full/10.1080/13697137.2017.1421925",
        "keyFindings": "Revisão sistemática demonstrou que progesterona micronizada associada a estrogênio tem risco de câncer de mama significativamente menor comparado a progestágenos sintéticos. Recomenda-se como primeira escolha na terapia hormonal da menopausa.",
        "studyType": "Revisão sistemática"
    },
    {
        "title": "Progesterone vs. synthetic progestins and the risk of breast cancer: a systematic review and meta-analysis",
        "authors": "Asi N, Mohammed K, Haydour Q, et al",
        "journal": "Systematic Reviews",
        "publicationYear": 2016,
        "volume": "5",
        "issue": "1",
        "pages": "121",
        "doi": "10.1186/s13643-016-0294-5",
        "pmid": "27439390",
        "url": "https://pmc.ncbi.nlm.nih.gov/articles/PMC4960754/",
        "keyFindings": "Meta-análise demonstrou que progesterona foi associada a menor risco de câncer de mama comparado a progestinas sintéticas quando combinados com estrogênio, com risco relativo de 0,67 (IC 95% 0,55-0,81). Evidência de qualidade moderada.",
        "studyType": "Meta-análise"
    }
]

def main():
    # Login
    session = requests.Session()
    print("Fazendo login...")
    r = session.post(f"{API_BASE_URL}/auth/login",
                    json={"email": EMAIL, "password": PASSWORD})
    token = r.json()["accessToken"]
    session.headers.update({"Authorization": f"Bearer {token}"})
    print("✓ Login OK\n")

    print("="*80)
    print("CRIANDO ARTIGOS RESTANTES")
    print("="*80 + "\n")

    article_ids = []
    for idx, article_data in enumerate(ARTICLES, 1):
        try:
            print(f"[{idx}/2] {article_data['title'][:60]}...")

            # Check if already exists
            r = session.get(f"{API_BASE_URL}/articles/search",
                          params={"q": article_data['title'][:30], "pageSize": 5})

            existing = None
            if r.status_code == 200:
                results = r.json()
                for art in results.get("articles", []):
                    if art.get("doi") == article_data["doi"]:
                        existing = art
                        break

            if existing:
                print(f"    ⚠ Artigo já existe: {existing['id']}")
                article_ids.append(existing['id'])
            else:
                # Create new article
                payload = {
                    "title": article_data["title"],
                    "authors": article_data["authors"],
                    "journal": article_data["journal"],
                    "publicationYear": article_data["publicationYear"],
                    "volume": article_data["volume"],
                    "issue": article_data["issue"],
                    "pages": article_data["pages"],
                    "doi": article_data["doi"],
                    "pmid": article_data["pmid"],
                    "url": article_data["url"],
                    "keyFindings": article_data["keyFindings"],
                    "studyType": article_data["studyType"]
                }

                r = session.post(f"{API_BASE_URL}/articles", json=payload)
                r.raise_for_status()
                article_id = r.json()["id"]
                article_ids.append(article_id)
                print(f"    ✓ Criado: {article_id}")

            time.sleep(0.3)
            print()

        except Exception as e:
            print(f"    ✗ Erro: {e}\n")

    # Link to score item
    print("="*80)
    print("VINCULANDO AO SCORE ITEM")
    print("="*80 + "\n")

    for idx, article_id in enumerate(article_ids, 1):
        try:
            print(f"[{idx}/{len(article_ids)}] Vinculando {article_id}...")
            r = session.post(f"{API_BASE_URL}/articles/{article_id}/score-items",
                           json={"scoreItemIds": [SCORE_ITEM_ID]})

            if r.status_code in [200, 201]:
                print(f"    ✓ Vinculado\n")
            elif r.status_code == 409:
                print(f"    ⚠ Já vinculado\n")
            else:
                r.raise_for_status()

            time.sleep(0.2)
        except Exception as e:
            print(f"    ✗ Erro: {e}\n")

    print("="*80)
    print(f"CONCLUÍDO: {len(article_ids)} artigos processados")
    print("="*80)

if __name__ == "__main__":
    main()
