#!/usr/bin/env python3
"""
Verify progesterone score item enrichment
"""
import requests

API_BASE_URL = "http://localhost:3001/api/v1"
EMAIL, PASSWORD = "import@plenya.com", "Import@123456"
SCORE_ITEM_ID = "325b61c6-11f6-49af-b3c9-b211d8bd6c28"

def main():
    # Login
    session = requests.Session()
    r = session.post(f"{API_BASE_URL}/auth/login",
                    json={"email": EMAIL, "password": PASSWORD})
    token = r.json()["accessToken"]
    session.headers.update({"Authorization": f"Bearer {token}"})

    print("="*80)
    print("VERIFICAÇÃO DE ENRIQUECIMENTO: PROGESTERONA - MULHERES PÓS-MENOPAUSA")
    print("="*80 + "\n")

    # Get score item
    r = session.get(f"{API_BASE_URL}/score-items/{SCORE_ITEM_ID}")
    r.raise_for_status()
    item = r.json()

    print("SCORE ITEM:")
    print(f"  ID: {item.get('id', 'N/A')}")
    print(f"  Nome: {item.get('itemName', 'N/A')}")
    print(f"  Grupo: {item.get('groupName', 'N/A')}")
    print()

    print("CONTEÚDO CLÍNICO:")
    cr = item.get('clinicalRelevance', '')
    pe = item.get('patientExplanation', '')
    cd = item.get('conduct', '')

    print(f"  Clinical Relevance: {len(cr)} caracteres")
    if cr:
        print(f"    Início: {cr[:100]}...")
    else:
        print("    ⚠ VAZIO")

    print(f"\n  Patient Explanation: {len(pe)} caracteres")
    if pe:
        print(f"    Início: {pe[:100]}...")
    else:
        print("    ⚠ VAZIO")

    print(f"\n  Conduct: {len(cd)} caracteres")
    if cd:
        print(f"    Início: {cd[:100]}...")
    else:
        print("    ⚠ VAZIO")

    # Search for linked articles
    print("\n" + "="*80)
    print("ARTIGOS VINCULADOS")
    print("="*80 + "\n")

    search_terms = [
        "progesterone",
        "menopausal hormone therapy",
        "micronized progesterone"
    ]

    all_article_ids = set()
    for term in search_terms:
        try:
            r = session.get(f"{API_BASE_URL}/articles/search",
                          params={"q": term, "page": 1, "pageSize": 20})
            r.raise_for_status()
            results = r.json()

            for article in results.get("articles", []):
                # Check if linked to our score item
                try:
                    r2 = session.get(f"{API_BASE_URL}/articles/{article['id']}/score-items")
                    r2.raise_for_status()
                    score_items = r2.json()

                    if any(si.get('id') == SCORE_ITEM_ID for si in score_items):
                        all_article_ids.add(article['id'])
                except:
                    pass
        except:
            pass

    linked_articles = list(all_article_ids)

    if linked_articles:
        print(f"Total de artigos vinculados: {len(linked_articles)}\n")
        for idx, article_id in enumerate(linked_articles, 1):
            try:
                r = session.get(f"{API_BASE_URL}/articles/{article_id}")
                r.raise_for_status()
                article = r.json()

                print(f"[{idx}] {article.get('title', 'N/A')}")
                print(f"    Autores: {article.get('authors', 'N/A')}")
                print(f"    Revista: {article.get('journal', 'N/A')} ({article.get('publicationYear', 'N/A')})")
                print(f"    DOI: {article.get('doi', 'N/A')}")
                print(f"    PMID: {article.get('pmId', 'N/A')}")
                print()
            except:
                print(f"[{idx}] Erro ao buscar artigo {article_id}\n")
    else:
        print("⚠ Nenhum artigo vinculado encontrado\n")

    # Summary
    print("="*80)
    print("RESUMO")
    print("="*80)
    print(f"✓ Conteúdo clínico: {'OK' if cr and pe and cd else '⚠ INCOMPLETO'}")
    print(f"✓ Clinical Relevance: {len(cr)}/1500-2000 caracteres (alvo)")
    print(f"✓ Patient Explanation: {len(pe)}/1000-1500 caracteres (alvo)")
    print(f"✓ Conduct: {len(cd)}/1500-2500 caracteres (alvo)")
    print(f"✓ Artigos vinculados: {len(linked_articles)}/2-4 (alvo)")
    print("="*80)

if __name__ == "__main__":
    main()
