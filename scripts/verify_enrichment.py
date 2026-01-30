#!/usr/bin/env python3
"""Verifica se o enriquecimento foi salvo corretamente na API"""

import requests
import json

API_BASE = "http://localhost:3001/api/v1"

# Login
response = requests.post(
    f"{API_BASE}/auth/login",
    json={"email": "import@plenya.com", "password": "Import@123456"}
)

token = response.json()["accessToken"]

# Testar 3 items diferentes (diabetes, cardiovascular, obesidade)
test_items = [
    ("4e960d9f-a04a-4524-9049-f4559170db14", "Diabetes"),
    ("8fb2e7de-c341-4ff5-814f-25d7695fd1cf", "Cardiovascular"),
    ("77f78798-6a3a-4eef-b092-7ae1c277df4e", "Obesidade")
]

print("=" * 70)
print("VERIFICAÇÃO DE ENRIQUECIMENTO")
print("=" * 70)

for item_id, disease_name in test_items:
    print(f"\n{disease_name} ({item_id}):")

    response = requests.get(
        f"{API_BASE}/score-items/{item_id}",
        headers={"Authorization": f"Bearer {token}"}
    )

    if response.status_code == 200:
        item = response.json()

        has_relevance = bool(item.get("clinicalRelevance"))
        has_conduct = bool(item.get("conduct"))

        if has_relevance:
            relevance = item["clinicalRelevance"]
            print(f"  ✓ Clinical Relevance: {len(relevance)} caracteres")
            print(f"    Preview: {relevance[:150]}...")
        else:
            print(f"  ✗ Clinical Relevance: VAZIO")

        if has_conduct:
            conduct = item["conduct"]
            print(f"  ✓ Conduct: {len(conduct)} caracteres")
            print(f"    Preview: {conduct[:150]}...")
        else:
            print(f"  ✗ Conduct: VAZIO")
    else:
        print(f"  ✗ Erro ao buscar item: {response.status_code}")

print("\n" + "=" * 70)
