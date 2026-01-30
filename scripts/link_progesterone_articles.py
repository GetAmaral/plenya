#!/usr/bin/env python3
"""
Link existing progesterone articles to score item
Since articles already exist (duplicate DOI error), this script finds them and links them.
"""
import requests
import time

API_BASE_URL = "http://localhost:3001/api/v1"
EMAIL, PASSWORD = "import@plenya.com", "Import@123456"
SCORE_ITEM_ID = "325b61c6-11f6-49af-b3c9-b211d8bd6c28"

# DOIs dos artigos já criados
DOIs = [
    "10.1371/journal.pmed.1004033",
    "10.1097/GME.0000000000002333",
    "10.1080/13697137.2017.1421925",
    "10.1080/13697137.2022.2085226"
]

def main():
    # Login
    session = requests.Session()
    print("Fazendo login...")
    r = session.post(f"{API_BASE_URL}/auth/login",
                    json={"email": EMAIL, "password": PASSWORD})
    r.raise_for_status()
    token = r.json()["accessToken"]
    session.headers.update({"Authorization": f"Bearer {token}"})
    print("✓ Login OK\n")

    # Find articles by searching
    print("="*80)
    print("BUSCANDO ARTIGOS EXISTENTES")
    print("="*80 + "\n")

    article_ids = []
    search_terms = [
        "Menopausal Hormone Therapy Formulation",
        "menopausal hormone therapy beyond age 65",
        "micronized progesterone breast cancer",
        "micronized progesterone estradiol lipids"
    ]

    for idx, search_term in enumerate(search_terms, 1):
        print(f"[{idx}/4] Buscando: {search_term}")
        try:
            r = session.get(f"{API_BASE_URL}/articles/search",
                          params={"q": search_term, "page": 1, "pageSize": 10})
            r.raise_for_status()
            results = r.json()

            if results.get("articles") and len(results["articles"]) > 0:
                article = results["articles"][0]  # Take first result
                article_ids.append(article["id"])
                print(f"    ✓ Encontrado: {article['title'][:60]}...")
                print(f"      ID: {article['id']}\n")
            else:
                print(f"    ✗ Não encontrado\n")
        except Exception as e:
            print(f"    ✗ Erro ao buscar: {e}\n")

    print(f"Total encontrado: {len(article_ids)}/4\n")

    # Link to score item
    print("="*80)
    print("VINCULANDO ARTIGOS AO SCORE ITEM")
    print("="*80 + "\n")

    success = 0
    for idx, article_id in enumerate(article_ids, 1):
        try:
            print(f"[{idx}/{len(article_ids)}] Vinculando {article_id}...")
            r = session.post(f"{API_BASE_URL}/articles/{article_id}/score-items",
                           json={"scoreItemIds": [SCORE_ITEM_ID]})
            r.raise_for_status()
            print(f"    ✓ Vinculado\n")
            success += 1
            time.sleep(0.2)
        except requests.exceptions.HTTPError as e:
            if e.response.status_code == 409:
                print(f"    ⚠ Já vinculado\n")
                success += 1
            else:
                print(f"    ✗ Erro: {e}\n")
        except Exception as e:
            print(f"    ✗ Erro: {e}\n")

    # Verify
    print("="*80)
    print("VERIFICAÇÃO")
    print("="*80 + "\n")

    for idx, article_id in enumerate(article_ids, 1):
        try:
            r = session.get(f"{API_BASE_URL}/articles/{article_id}/score-items")
            r.raise_for_status()
            score_items = r.json()
            linked = any(item.get('id') == SCORE_ITEM_ID for item in score_items)

            r2 = session.get(f"{API_BASE_URL}/articles/{article_id}")
            r2.raise_for_status()
            article = r2.json()

            print(f"[{idx}] {article['title'][:60]}...")
            print(f"    Vinculado: {'✓ SIM' if linked else '✗ NÃO'}\n")
        except Exception as e:
            print(f"[{idx}] Erro ao verificar: {e}\n")

    print("="*80)
    print(f"CONCLUÍDO: {success}/{len(article_ids)} artigos vinculados")
    print("="*80)

if __name__ == "__main__":
    main()
