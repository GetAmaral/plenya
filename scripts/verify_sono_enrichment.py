#!/usr/bin/env python3
"""
Script de verificação do enriquecimento dos 20 items de SONO
"""

import requests
import json

API_URL = "http://localhost:3001/api/v1"
EMAIL = "import@plenya.com"
PASSWORD = "Import@123456"

def login():
    response = requests.post(
        f"{API_URL}/auth/login",
        json={"email": EMAIL, "password": PASSWORD}
    )
    return response.json()["accessToken"]

def get_item(token, item_id):
    headers = {"Authorization": f"Bearer {token}"}
    response = requests.get(
        f"{API_URL}/score-items/{item_id}",
        headers=headers
    )
    return response.json()

# Alguns items para verificar
SAMPLE_ITEMS = [
    ("339da4f9-89ef-4990-bbc5-e15de8eb96be", "Campos eletromagnéticos"),
    ("b4d6985c-3c5d-4709-ad60-43b2a44cd079", "Bruxismo"),
    ("29f615c2-9608-409f-aa26-2c81aec67dda", "Apneias"),
]

print("=== Verificação de Enriquecimento ===\n")
token = login()

for item_id, name in SAMPLE_ITEMS:
    print(f"\n{'='*60}")
    print(f"Item: {name}")
    print(f"ID: {item_id}")
    print('='*60)

    item = get_item(token, item_id)

    # Verificar campos
    cr = item.get("clinicalRelevance", "")
    pe = item.get("patientExplanation", "")
    co = item.get("conduct", "")

    print(f"\nClinical Relevance: {len(cr)} caracteres")
    print(f"Primeiras 300 caracteres:")
    print(cr[:300] + "...")

    print(f"\nPatient Explanation: {len(pe)} caracteres")
    print(f"Primeiras 200 caracteres:")
    print(pe[:200] + "...")

    print(f"\nConduct: {len(co)} caracteres")
    print(f"Primeiras 300 caracteres:")
    print(co[:300] + "...")

    # Verificar artigos vinculados
    articles = item.get("articles", [])
    print(f"\nArtigos vinculados: {len(articles)}")
    for article in articles:
        print(f"  - {article.get('title', 'N/A')}")

print("\n" + "="*60)
print("Verificação concluída!")
