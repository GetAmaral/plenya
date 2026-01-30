#!/usr/bin/env python3
"""Amostrar items de VIDA SEXUAL para verificar qualidade"""

import requests
import json

API_BASE = "http://localhost:3001/api/v1"
LOGIN_EMAIL = "import@plenya.com"
LOGIN_PASSWORD = "Import@123456"

# Amostras representativas
SAMPLES = [
    "70babe19-54b9-4d12-9f9e-5dd3fc1e6df0",  # ASEX
    "bd7913cc-c28c-4c61-852d-4e4251cf3e83",  # FSFI
    "7cc28981-bd20-4e5e-b7c0-bd312f6a7ed4",  # IIEF-5
    "b884c9c9-8279-4bc8-a56b-41bec7497c68",  # Traumas
    "7b587635-f23b-4b9a-9076-5b7f5e61920c"   # Menopausa
]

def login():
    response = requests.post(
        f"{API_BASE}/auth/login",
        json={"email": LOGIN_EMAIL, "password": LOGIN_PASSWORD}
    )
    response.raise_for_status()
    return response.json()['accessToken']

def get_item(token, item_id):
    response = requests.get(
        f"{API_BASE}/score-items/{item_id}",
        headers={'Authorization': f'Bearer {token}'}
    )
    response.raise_for_status()
    return response.json()

def main():
    print("="*100)
    print("AMOSTRAS DE ITEMS ENRIQUECIDOS - VIDA SEXUAL")
    print("="*100 + "\n")

    token = login()

    for idx, item_id in enumerate(SAMPLES, 1):
        item = get_item(token, item_id)

        print(f"\n{'='*100}")
        print(f"AMOSTRA {idx}/{len(SAMPLES)}")
        print(f"{'='*100}\n")

        print(f"ID: {item['id']}")
        print(f"Nome: {item['name']}")
        print(f"Grupo: {item.get('subgroup', {}).get('group', {}).get('name', 'N/A')}")
        print(f"Subgrupo: {item.get('subgroup', {}).get('name', 'N/A')}\n")

        print(f"RELEVÂNCIA CLÍNICA ({len(item.get('clinicalRelevance', ''))} chars):")
        print(f"{item.get('clinicalRelevance', 'N/A')}\n")

        print(f"CONDUTA ({len(item.get('conduct', ''))} chars):")
        print(f"{item.get('conduct', 'N/A')}\n")

        print(f"EXPLICAÇÃO PACIENTE ({len(item.get('patientExplanation', ''))} chars):")
        print(f"{item.get('patientExplanation', 'N/A')}\n")

        if item.get('lastReview'):
            print(f"Última Revisão: {item['lastReview']}\n")

if __name__ == "__main__":
    main()
