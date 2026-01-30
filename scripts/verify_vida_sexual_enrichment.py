#!/usr/bin/env python3
"""
Verificar qualidade do enriquecimento de VIDA SEXUAL
Busca items aleatórios e exibe conteúdo completo
"""

import requests
import json

API_BASE = "http://localhost:3001/api/v1"
LOGIN_EMAIL = "import@plenya.com"
LOGIN_PASSWORD = "Import@123456"

# Selecionar alguns items para verificação detalhada
SAMPLE_IDS = [
    "70babe19-54b9-4d12-9f9e-5dd3fc1e6df0",  # Primeiro
    "86f1dbce-c9c5-4c96-9b0b-4c6452c5dd1e",  # Segundo
    "3a1759ed-cc26-4fb6-99ad-dd12f4e9e919",  # Meio
    "7cc28981-bd20-4e5e-b7c0-bd312f6a7ed4",  # Penúltimo
    "6d18ecfb-040a-4467-9c04-2dc347979727"   # Último
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
    print("="*80)
    print("VERIFICAÇÃO DE QUALIDADE - ENRIQUECIMENTO VIDA SEXUAL")
    print("="*80 + "\n")

    token = login()
    print("Autenticado com sucesso\n")

    for idx, item_id in enumerate(SAMPLE_IDS, 1):
        print(f"\n{'='*80}")
        print(f"ITEM {idx}/{len(SAMPLE_IDS)}: {item_id}")
        print(f"{'='*80}\n")

        item = get_item(token, item_id)

        print(f"PERGUNTA:")
        print(f"  {item.get('question', 'N/A')}\n")

        print(f"GRUPO:")
        print(f"  {item.get('groupName', 'N/A')}\n")

        print(f"RELEVÂNCIA CLÍNICA:")
        rel = item.get('clinicalRelevance', 'N/A')
        print(f"  Caracteres: {len(rel)}")
        print(f"  {rel[:300]}...\n")

        print(f"AÇÃO CLÍNICA:")
        action = item.get('clinicalAction', 'N/A')
        print(f"  Caracteres: {len(action)}")
        print(f"  {action[:300]}...\n")

        print(f"FONTES CIENTÍFICAS:")
        sources = item.get('scientificSources', [])
        print(f"  Total: {len(sources)}")
        for i, source in enumerate(sources[:3], 1):
            print(f"  {i}. {source[:150]}...")

        print("\n")

if __name__ == "__main__":
    main()
