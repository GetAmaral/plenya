#!/usr/bin/env python3
"""
Verificar vinculação de artigos aos score items
"""

import requests

API_URL = "http://localhost:3001/api/v1"
EMAIL = "import@plenya.com"
PASSWORD = "Import@123456"

def login():
    response = requests.post(
        f"{API_URL}/auth/login",
        json={"email": EMAIL, "password": PASSWORD}
    )
    return response.json()["accessToken"]

def link_article(token, article_id, item_id):
    """Vincula um artigo a um score item."""
    headers = {"Authorization": f"Bearer {token}"}
    response = requests.post(
        f"{API_URL}/articles/{article_id}/score-items",
        json={"score_item_id": item_id},
        headers=headers
    )

    print(f"Status: {response.status_code}")
    if response.status_code not in [200, 201, 409]:
        print(f"Erro: {response.text}")
    return response.status_code in [200, 201, 409]

# Testar com um item
ARTICLE_ID = "6e63eb6e-9895-4f13-9c98-fa675bd1020e"  # Ritmo Circadiano I
ITEM_ID = "339da4f9-89ef-4990-bbc5-e15de8eb96be"  # Campos eletromagnéticos

print("Testando vinculação de artigo...")
print(f"Article ID: {ARTICLE_ID}")
print(f"Item ID: {ITEM_ID}")

token = login()
result = link_article(token, ARTICLE_ID, ITEM_ID)

print(f"\nResultado: {'Sucesso' if result else 'Falha'}")

# Verificar se o endpoint retorna os artigos vinculados
print("\nVerificando item após vinculação...")
headers = {"Authorization": f"Bearer {token}"}
response = requests.get(
    f"{API_URL}/score-items/{ITEM_ID}",
    headers=headers
)

item = response.json()
articles = item.get("articles", [])
print(f"Artigos vinculados: {len(articles)}")

if articles:
    for article in articles:
        print(f"  - {article.get('title', 'N/A')}")
else:
    print("Nenhum artigo vinculado encontrado no response")

# Testar endpoint alternativo - buscar artigos do item
print("\nTestando endpoint alternativo...")
response2 = requests.get(
    f"{API_URL}/score-items/{ITEM_ID}/articles",
    headers=headers
)

if response2.status_code == 200:
    articles2 = response2.json()
    print(f"Artigos (endpoint alternativo): {len(articles2)}")
else:
    print(f"Endpoint alternativo não existe (status {response2.status_code})")
