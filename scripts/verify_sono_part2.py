#!/usr/bin/env python3
"""
Verificação dos items SONO - Parte 2
Confirma que o enriquecimento foi aplicado corretamente
"""

import requests
import json

API_BASE = "http://localhost:3001/api/v1"
LOGIN_EMAIL = "import@plenya.com"
LOGIN_PASSWORD = "Import@123456"

# Amostra de items para verificar
SAMPLE_IDS = [
    ("073e9b8b-f3b8-411d-996c-e4b6ae6b71e1", "Campos eletromagnéticos"),
    ("339da4f9-89ef-4990-bbc5-e15de8eb96be", "Colchão"),
    ("428052bb-0cbd-4ab2-977c-a6a39cedc3fa", "Estimulantes noturnos (cafeína)"),
    ("88f10871-8821-42db-bee4-01c31f2187f8", "Higiene do sono"),
    ("b256e241-ac54-4043-82a3-0ec5087702bc", "Motivos"),
]

def login():
    response = requests.post(
        f"{API_BASE}/auth/login",
        json={"email": LOGIN_EMAIL, "password": LOGIN_PASSWORD}
    )
    return response.json()["accessToken"]

def verify_item(token, item_id, expected_name):
    headers = {"Authorization": f"Bearer {token}"}
    response = requests.get(f"{API_BASE}/score-items/{item_id}", headers=headers)
    data = response.json()
    if "data" in data:
        data = data["data"]

    print(f"\n{'='*80}")
    print(f"ITEM: {data['name']}")
    print(f"{'='*80}")

    # Verificar campos enriquecidos
    fields = [
        ("clinical_relevance", "Relevância Clínica"),
        ("interpretation_guide", "Guia de Interpretação"),
        ("scientific_references", "Referências Científicas"),
        ("low_description", "Descrição Low"),
        ("medium_description", "Descrição Medium"),
        ("high_description", "Descrição High"),
    ]

    all_present = True
    for field, label in fields:
        value = data.get(field)
        if value and len(str(value)) > 50:
            status = "✓"
            preview = str(value)[:80] + "..."
        else:
            status = "✗"
            preview = "AUSENTE ou VAZIO"
            all_present = False

        print(f"{status} {label:25s}: {preview}")

    return all_present

def main():
    print("VERIFICAÇÃO DE ENRIQUECIMENTO - SONO PARTE 2")
    print("="*80)

    token = login()
    print("✓ Login realizado")

    total = len(SAMPLE_IDS)
    success = 0

    for item_id, expected_name in SAMPLE_IDS:
        if verify_item(token, item_id, expected_name):
            success += 1

    print(f"\n{'='*80}")
    print(f"RESULTADO: {success}/{total} items verificados com sucesso")
    print(f"{'='*80}")

if __name__ == "__main__":
    main()
