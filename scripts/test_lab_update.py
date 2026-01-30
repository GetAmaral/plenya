#!/usr/bin/env python3
"""
Script de teste para verificar atualização de lab items
"""

import requests
import json

API_BASE = "http://localhost:3001/api/v1"
LOGIN_EMAIL = "import@plenya.com"
LOGIN_PASSWORD = "Import@123456"
TEST_ITEM_ID = "62d59b40-6dab-46c8-a6f5-53bafe3b7cae"  # CA-125

# Login
response = requests.post(
    f"{API_BASE}/auth/login",
    json={"email": LOGIN_EMAIL, "password": LOGIN_PASSWORD}
)
token = response.json()["accessToken"]
headers = {
    "Authorization": f"Bearer {token}",
    "Content-Type": "application/json"
}

# Get current item
print("Buscando item atual...")
response = requests.get(
    f"{API_BASE}/lab-tests/definitions/{TEST_ITEM_ID}",
    headers=headers
)
current_item = response.json()
print(f"Status: {response.status_code}")
print(f"Item atual (primeiros campos):")
print(json.dumps({
    "id": current_item.get("id"),
    "name": current_item.get("name"),
    "clinicalSignificance": current_item.get("clinicalSignificance", "")[:50] if current_item.get("clinicalSignificance") else None,
    "longevityContext": current_item.get("longevityContext", "")[:50] if current_item.get("longevityContext") else None,
    "clinicalRecommendations": current_item.get("clinicalRecommendations", "")[:50] if current_item.get("clinicalRecommendations") else None,
}, indent=2))

# Update with test data
test_enrichment = {
    "clinicalSignificance": "TESTE: Este é um teste de significância clínica.",
    "longevityContext": "TESTE: Este é um teste de contexto de longevidade.",
    "clinicalRecommendations": "TESTE: Estas são recomendações clínicas de teste."
}

updated_item = {**current_item, **test_enrichment}

print("\n\nEnviando atualização...")
print(f"Dados sendo enviados (clinicalSignificance): {updated_item.get('clinicalSignificance')}")

response = requests.put(
    f"{API_BASE}/lab-tests/definitions/{TEST_ITEM_ID}",
    headers=headers,
    json=updated_item
)

print(f"Status da atualização: {response.status_code}")
if response.status_code != 200:
    print(f"Erro: {response.text}")
else:
    result = response.json()
    print(f"Item atualizado:")
    print(json.dumps({
        "id": result.get("id"),
        "name": result.get("name"),
        "clinicalSignificance": result.get("clinicalSignificance", "")[:100] if result.get("clinicalSignificance") else None,
        "longevityContext": result.get("longevityContext", "")[:100] if result.get("longevityContext") else None,
        "clinicalRecommendations": result.get("clinicalRecommendations", "")[:100] if result.get("clinicalRecommendations") else None,
    }, indent=2))

# Verificar no banco
print("\n\nVerificando no banco...")
import subprocess
result = subprocess.run([
    "docker", "compose", "exec", "-T", "db",
    "psql", "-U", "plenya_user", "-d", "plenya_db",
    "-c", f"SELECT name, clinical_significance, longevity_context, clinical_recommendations FROM lab_test_definitions WHERE id = '{TEST_ITEM_ID}'"
], capture_output=True, text=True)
print(result.stdout)
